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

package operations

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/api/meta"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	opsv1alpha1 "github.com/apecloud/kubeblocks/apis/operations/v1alpha1"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	"github.com/apecloud/kubeblocks/pkg/generics"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
	testops "github.com/apecloud/kubeblocks/pkg/testutil/operations"
)

var _ = Describe("Restart OpsRequest", func() {

	var (
		randomStr   = testCtx.GetRandomStr()
		compDefName = "test-compdef-" + randomStr
		clusterName = "test-cluster-" + randomStr
	)

	cleanEnv := func() {
		// must wait till resources deleted and no longer existed before the testcases start,
		// otherwise if later it needs to create some new resource objects with the same name,
		// in race conditions, it will find the existence of old objects, resulting failure to
		// create the new objects.
		By("clean resources")

		// delete cluster(and all dependent sub-resources), cluster definition
		testapps.ClearClusterResourcesWithRemoveFinalizerOption(&testCtx)

		// delete rest resources
		inNS := client.InNamespace(testCtx.DefaultNamespace)
		ml := client.HasLabels{testCtx.TestObjLabelKey}
		// namespaced
		testapps.ClearResources(&testCtx, generics.OpsRequestSignature, inNS, ml)
		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.InstanceSetSignature, true, inNS, ml)
	}

	BeforeEach(cleanEnv)

	AfterEach(cleanEnv)

	Context("Test OpsRequest", func() {
		var (
			opsRes  *OpsResource
			cluster *appsv1.Cluster
			reqCtx  intctrlutil.RequestCtx
		)

		BeforeEach(func() {
			reqCtx = intctrlutil.RequestCtx{Ctx: testCtx.Ctx}
		})

		It("Test restart OpsRequest", func() {
			By("init operations resources ")
			opsRes, _, cluster = initOperationsResources(compDefName, clusterName)
			By("create Restart opsRequest")
			opsRes.OpsRequest = createRestartOpsObj(clusterName, "restart-ops-"+randomStr)
			mockComponentIsOperating(opsRes.Cluster, appsv1.UpdatingComponentPhase, defaultCompName)

			By("mock restart OpsRequest to Creating")
			_, err := GetOpsManager().Do(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(testops.GetOpsRequestPhase(&testCtx, client.ObjectKeyFromObject(opsRes.OpsRequest))).Should(Equal(opsv1alpha1.OpsCreatingPhase))

			By("test restart action and reconcile function")
			rHandler := restartOpsHandler{}
			_ = rHandler.Action(reqCtx, k8sClient, opsRes)

			_, err = GetOpsManager().Reconcile(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("expect failed when cluster is stopped", func() {
			By("init operations resources ")
			opsRes, _, cluster = initOperationsResources(compDefName, clusterName)
			By("mock cluster is stopped")
			Expect(testapps.ChangeObjStatus(&testCtx, cluster, func() {
				cluster.Status.Phase = appsv1.StoppedClusterPhase
			})).Should(Succeed())
			By("create Restart opsRequest")
			opsRes.OpsRequest = createRestartOpsObj(clusterName, "restart-ops-"+randomStr)
			_, err := GetOpsManager().Do(reqCtx, k8sClient, opsRes)
			Expect(err).ShouldNot(HaveOccurred())
			Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(opsRes.OpsRequest),
				func(g Gomega, fetched *opsv1alpha1.OpsRequest) {
					g.Expect(fetched.Status.Phase).To(Equal(opsv1alpha1.OpsFailedPhase))
					condition := meta.FindStatusCondition(fetched.Status.Conditions, opsv1alpha1.ConditionTypeValidated)
					g.Expect(condition.Message).Should(Equal("OpsRequest.spec.type=Restart is forbidden when Cluster.status.phase=Stopped"))
				})).Should(Succeed())
		})
	})
})

func createRestartOpsObj(clusterName, restartOpsName string, compNames ...string) *opsv1alpha1.OpsRequest {
	ops := testops.NewOpsRequestObj(restartOpsName, testCtx.DefaultNamespace,
		clusterName, opsv1alpha1.RestartType)
	if len(compNames) == 0 {
		ops.Spec.RestartList = []opsv1alpha1.ComponentOps{
			{ComponentName: defaultCompName},
		}
	} else {
		for _, compName := range compNames {
			ops.Spec.RestartList = append(ops.Spec.RestartList, opsv1alpha1.ComponentOps{
				ComponentName: compName,
			})
		}
	}
	opsRequest := testops.CreateOpsRequest(ctx, testCtx, ops)
	opsRequest.Status.Phase = opsv1alpha1.OpsPendingPhase
	return opsRequest
}
