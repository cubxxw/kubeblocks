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
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/configuration/core"
	cfgutil "github.com/apecloud/kubeblocks/pkg/configuration/util"
	"github.com/apecloud/kubeblocks/pkg/controller/render"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
)

var _ = Describe("Configuration Controller", func() {

	BeforeEach(cleanEnv)

	AfterEach(cleanEnv)

	Context("When updating configuration", func() {
		It("Should reconcile success", func() {
			_, _, clusterObj, componentObj, synthesizedComp := mockReconcileResource()

			cfgKey := client.ObjectKey{
				Name:      core.GenerateComponentConfigurationName(clusterName, defaultCompName),
				Namespace: testCtx.DefaultNamespace,
			}
			checkCfgStatus := func(phase appsv1alpha1.ConfigurationPhase) func() bool {
				return func() bool {
					cfg := &appsv1alpha1.Configuration{}
					Expect(k8sClient.Get(ctx, cfgKey, cfg)).Should(Succeed())
					itemStatus := cfg.Status.GetItemStatus(configSpecName)
					return itemStatus != nil && itemStatus.Phase == phase
				}
			}

			By("wait for configuration status to be init phase.")
			Eventually(checkCfgStatus(appsv1alpha1.CInitPhase)).Should(BeFalse())
			Expect(initConfiguration(&render.ResourceCtx{
				Client:        k8sClient,
				Context:       ctx,
				Namespace:     testCtx.DefaultNamespace,
				ClusterName:   clusterName,
				ComponentName: defaultCompName,
			}, synthesizedComp, clusterObj, componentObj)).Should(Succeed())

			Eventually(checkCfgStatus(appsv1alpha1.CFinishedPhase)).Should(BeTrue())

			By("reconfiguring parameters.")
			Eventually(testapps.GetAndChangeObj(&testCtx, cfgKey, func(cfg *appsv1alpha1.Configuration) {
				cfg.Spec.GetConfigurationItem(configSpecName).ConfigFileParams = map[string]appsv1alpha1.ConfigParams{
					"my.cnf": {
						Parameters: map[string]*string{
							"max_connections": cfgutil.ToPointer("1000"),
							"gtid_mode":       cfgutil.ToPointer("ON"),
						},
					},
				}
			})).Should(Succeed())

			Eventually(func(g Gomega) {
				cfg := &appsv1alpha1.Configuration{}
				g.Expect(k8sClient.Get(ctx, cfgKey, cfg)).Should(Succeed())
				itemStatus := cfg.Status.GetItemStatus(configSpecName)
				g.Expect(itemStatus).ShouldNot(BeNil())
				g.Expect(itemStatus.UpdateRevision).Should(BeEquivalentTo("2"))
				g.Expect(itemStatus.Phase).Should(BeEquivalentTo(appsv1alpha1.CFinishedPhase))
			}, time.Second*60, time.Second*1).Should(Succeed())
		})

		It("Invalid component test", func() {
			_, _, clusterObj, componentObj, synthesizedComp := mockReconcileResource()

			cfgKey := client.ObjectKey{
				Name:      core.GenerateComponentConfigurationName(clusterName, "invalid-component"),
				Namespace: testCtx.DefaultNamespace,
			}

			Expect(initConfiguration(&render.ResourceCtx{
				Client:        k8sClient,
				Context:       ctx,
				Namespace:     testCtx.DefaultNamespace,
				ClusterName:   clusterName,
				ComponentName: "invalid-component",
			}, synthesizedComp, clusterObj, componentObj)).Should(Succeed())

			Eventually(func(g Gomega) {
				cfg := &appsv1alpha1.Configuration{}
				g.Expect(k8sClient.Get(ctx, cfgKey, cfg)).Should(Succeed())
				g.Expect(cfg.Status.Message).Should(ContainSubstring("not found cluster component"))
			}, time.Second*60, time.Second*1).Should(Succeed())
		})
	})

	Context("When updating configuration with injectEnvTo", func() {
		It("Should reconcile success", func() {
			_, _, clusterObj, componentObj, synthesizedComp := mockReconcileResource()
			synthesizedComp.ConfigTemplates[0].AsSecret = cfgutil.ToPointer(true)
			synthesizedComp.ConfigTemplates[0].InjectEnvTo = []string{"mock-container"}

			cfgKey := client.ObjectKey{
				Name:      core.GenerateComponentConfigurationName(clusterName, defaultCompName),
				Namespace: testCtx.DefaultNamespace,
			}
			renderedKey := client.ObjectKey{
				Name:      core.GetComponentCfgName(synthesizedComp.ClusterName, synthesizedComp.Name, synthesizedComp.ConfigTemplates[0].Name),
				Namespace: testCtx.DefaultNamespace,
			}
			checkCfgStatus := func(phase appsv1alpha1.ConfigurationPhase) func() bool {
				return func() bool {
					cfg := &appsv1alpha1.Configuration{}
					Expect(k8sClient.Get(ctx, cfgKey, cfg)).Should(Succeed())
					itemStatus := cfg.Status.GetItemStatus(configSpecName)
					return itemStatus != nil && itemStatus.Phase == phase
				}
			}

			By("wait for configuration status to be init phase.")
			Eventually(checkCfgStatus(appsv1alpha1.CInitPhase)).Should(BeFalse())
			Expect(initConfiguration(&render.ResourceCtx{
				Client:        k8sClient,
				Context:       ctx,
				Namespace:     testCtx.DefaultNamespace,
				ClusterName:   clusterName,
				ComponentName: defaultCompName,
			}, synthesizedComp, clusterObj, componentObj)).Should(Succeed())

			Eventually(checkCfgStatus(appsv1alpha1.CFinishedPhase)).Should(BeTrue())

			Eventually(testapps.CheckObjExists(&testCtx, renderedKey, &corev1.ConfigMap{}, false)).Should(Succeed())
			Eventually(testapps.CheckObjExists(&testCtx, client.ObjectKey{
				Name:      core.GenerateEnvFromName(renderedKey.Name),
				Namespace: renderedKey.Namespace,
			}, &corev1.Secret{}, true)).Should(Succeed())

			By("reconfiguring parameters.")
			Eventually(testapps.GetAndChangeObj(&testCtx, cfgKey, func(cfg *appsv1alpha1.Configuration) {
				cfg.Spec.GetConfigurationItem(configSpecName).ConfigFileParams = map[string]appsv1alpha1.ConfigParams{
					"my.cnf": {
						Parameters: map[string]*string{
							"max_connections": cfgutil.ToPointer("1000"),
							"gtid_mode":       cfgutil.ToPointer("ON"),
						},
					},
				}
			})).Should(Succeed())

			Eventually(func(g Gomega) {
				cfg := &appsv1alpha1.Configuration{}
				g.Expect(k8sClient.Get(ctx, cfgKey, cfg)).Should(Succeed())
				itemStatus := cfg.Status.GetItemStatus(configSpecName)
				g.Expect(itemStatus).ShouldNot(BeNil())
				g.Expect(itemStatus.UpdateRevision).Should(BeEquivalentTo("2"))
				g.Expect(itemStatus.Phase).Should(BeEquivalentTo(appsv1alpha1.CFinishedPhase))
			}, time.Second*60, time.Second*1).Should(Succeed())

			Eventually(testapps.CheckObjExists(&testCtx, renderedKey, &corev1.ConfigMap{}, false)).Should(Succeed())
			Eventually(testapps.CheckObjExists(&testCtx, client.ObjectKey{
				Name:      core.GenerateEnvFromName(renderedKey.Name),
				Namespace: renderedKey.Namespace,
			}, &corev1.Secret{}, true)).Should(Succeed())
		})

	})
})
