/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package parameters

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	appsv1beta1 "github.com/apecloud/kubeblocks/apis/apps/v1beta1"
	cfgcore "github.com/apecloud/kubeblocks/pkg/configuration/core"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	"github.com/apecloud/kubeblocks/pkg/generics"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
	testutil "github.com/apecloud/kubeblocks/pkg/testutil/k8s"
)

var _ = Describe("ConfigWrapper util test", func() {
	var (
		// ctrl       *gomock.Controller
		// mockClient *mock_client.MockClient
		k8sMockClient *testutil.K8sClientMockHelper

		reqCtx = intctrlutil.RequestCtx{
			Ctx: ctx,
			Log: log.FromContext(ctx).WithValues("reconfigure_for_test", testCtx.DefaultNamespace),
		}
	)

	var (
		configMapObj        *corev1.ConfigMap
		configConstraintObj *appsv1beta1.ConfigConstraint
		compDefObj          *appsv1.ComponentDefinition
	)

	cleanEnv := func() {
		// must wait till resources deleted and no longer existed before the testcases start,
		// otherwise if later it needs to create some new resource objects with the same name,
		// in race conditions, it will find the existence of old objects, resulting failure to
		// create the new objects.
		By("clean resources")

		// delete rest mocked objects
		inNS := client.InNamespace(testCtx.DefaultNamespace)
		ml := client.HasLabels{testCtx.TestObjLabelKey}
		// namespaced
		testapps.ClearResources(&testCtx, generics.ConfigMapSignature, inNS, ml)
		// non-namespaced
		testapps.ClearResources(&testCtx, generics.ComponentDefinitionSignature, ml)
		testapps.ClearResources(&testCtx, generics.ConfigConstraintSignature, ml)
	}

	BeforeEach(func() {
		cleanEnv()

		// Add any setup steps that needs to be executed before each test
		k8sMockClient = testutil.NewK8sMockClient()

		By("creating a cluster")
		configMapObj = testapps.CreateCustomizedObj(&testCtx,
			"resources/mysql-config-template.yaml", &corev1.ConfigMap{},
			testCtx.UseDefaultNamespace())

		configConstraintObj = testapps.CreateCustomizedObj(&testCtx,
			"resources/mysql-config-constraint.yaml",
			&appsv1beta1.ConfigConstraint{})

		By("Create a componentDefinition obj")
		compDefObj = testapps.NewComponentDefinitionFactory(compDefName).
			WithRandomName().
			SetDefaultSpec().
			AddConfigTemplate(configSpecName, configMapObj.Name, configConstraintObj.Name, testCtx.DefaultNamespace, configVolumeName).
			Create(&testCtx).
			GetObject()
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		cleanEnv()

		k8sMockClient.Finish()
	})

	Context("ComponentDefinition CR test", func() {
		It("Should success without error", func() {
			availableTPL := configConstraintObj.DeepCopy()
			availableTPL.Status.Phase = appsv1beta1.CCAvailablePhase

			k8sMockClient.MockPatchMethod(testutil.WithSucceed())
			k8sMockClient.MockListMethod(testutil.WithSucceed())
			k8sMockClient.MockGetMethod(testutil.WithGetReturned(testutil.WithConstructSequenceResult(
				map[client.ObjectKey][]testutil.MockGetReturned{
					client.ObjectKeyFromObject(configMapObj): {{
						Object: nil,
						Err:    cfgcore.MakeError("failed to get cc object"),
					}, {
						Object: configMapObj,
						Err:    nil,
					}},
					client.ObjectKeyFromObject(configConstraintObj): {{
						Object: nil,
						Err:    cfgcore.MakeError("failed to get cc object"),
					}, {
						Object: configConstraintObj,
						Err:    nil,
					}, {
						Object: availableTPL,
						Err:    nil,
					}},
				},
			), testutil.WithAnyTimes()))

			_, err := checkConfigTemplate(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).ShouldNot(Succeed())
			Expect(err.Error()).Should(ContainSubstring("failed to get cc object"))

			_, err = checkConfigTemplate(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).ShouldNot(Succeed())
			Expect(err.Error()).Should(ContainSubstring("failed to get cc object"))

			_, err = checkConfigTemplate(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).ShouldNot(Succeed())
			Expect(err.Error()).Should(ContainSubstring("status not ready"))

			ok, err := checkConfigTemplate(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).Should(Succeed())
			Expect(ok).Should(BeTrue())

			ok, err = updateLabelsByConfigSpec(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).Should(Succeed())
			Expect(ok).Should(BeTrue())

			_, err = updateLabelsByConfigSpec(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).Should(Succeed())

			err = DeleteConfigMapFinalizer(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).Should(Succeed())
		})
	})

	Context("ComponentDefinition CR test without config Constraints", func() {
		It("Should success without error", func() {
			// remove ConfigConstraintRef
			_, err := handleConfigTemplate(compDefObj, func(templates []appsv1.ComponentConfigSpec) (bool, error) {
				return true, nil
			}, func(compDef *appsv1.ComponentDefinition) error {
				if len(compDef.Spec.Configs) == 0 {
					return nil
				}
				for i := range compDef.Spec.Configs {
					tpl := &compDef.Spec.Configs[i]
					tpl.ConfigConstraintRef = ""
				}
				return nil
			})
			Expect(err).Should(Succeed())

			availableTPL := configConstraintObj.DeepCopy()
			availableTPL.Status.Phase = appsv1beta1.CCAvailablePhase

			k8sMockClient.MockGetMethod(testutil.WithGetReturned(testutil.WithConstructSequenceResult(
				map[client.ObjectKey][]testutil.MockGetReturned{
					client.ObjectKeyFromObject(configMapObj): {{
						Object: nil,
						Err:    cfgcore.MakeError("failed to get cc object"),
					}, {
						Object: configMapObj,
						Err:    nil,
					}}},
			), testutil.WithAnyTimes()))

			_, err = checkConfigTemplate(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).ShouldNot(Succeed())
			Expect(err.Error()).Should(ContainSubstring("failed to get cc object"))

			ok, err := checkConfigTemplate(k8sMockClient.Client(), reqCtx, compDefObj)
			Expect(err).Should(Succeed())
			Expect(ok).Should(BeTrue())
		})
	})
})
