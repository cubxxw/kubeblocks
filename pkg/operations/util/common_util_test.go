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

package util

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	dpv1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	opsv1alpha1 "github.com/apecloud/kubeblocks/apis/operations/v1alpha1"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/constant"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
	testops "github.com/apecloud/kubeblocks/pkg/testutil/operations"
)

var _ = Describe("OpsRequest Controller", func() {

	var (
		randomStr       = testCtx.GetRandomStr()
		compDefName     = "test-compdef-" + randomStr
		clusterName     = "test-cluster-" + randomStr
		defaultCompName = "mysql"
	)

	cleanupObjects := func() {
		err := k8sClient.DeleteAllOf(ctx, &appsv1.Cluster{}, client.InNamespace(testCtx.DefaultNamespace), client.HasLabels{testCtx.TestObjLabelKey})
		Expect(err).NotTo(HaveOccurred())
		err = k8sClient.DeleteAllOf(ctx, &opsv1alpha1.OpsRequest{}, client.InNamespace(testCtx.DefaultNamespace), client.HasLabels{testCtx.TestObjLabelKey})
		Expect(err).NotTo(HaveOccurred())
	}

	BeforeEach(func() {
		// Add any steup steps that needs to be executed before each test
		cleanupObjects()
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		cleanupObjects()
	})

	Context("Test OpsRequest", func() {
		It("Should Test all OpsRequest", func() {
			cluster := testapps.CreateDefaultMysqlCluster(&testCtx, clusterName, compDefName, defaultCompName)
			By("init restart OpsRequest")
			testOpsName := "restart-" + randomStr
			ops := testops.NewOpsRequestObj(testOpsName, testCtx.DefaultNamespace,
				clusterName, opsv1alpha1.RestartType)
			ops.Spec.RestartList = []opsv1alpha1.ComponentOps{
				{ComponentName: defaultCompName},
			}
			testops.CreateOpsRequest(ctx, testCtx, ops)

			By("test PatchOpsRequestReconcileAnnotation function")
			Expect(PatchOpsRequestReconcileAnnotation(ctx, k8sClient, cluster.Namespace, testOpsName)).Should(Succeed())
			opsRecordSlice := []opsv1alpha1.OpsRecorder{
				{
					Name: testOpsName,
					Type: opsv1alpha1.RestartType,
				},
				{
					Name: "not-exists-ops",
					Type: opsv1alpha1.RestartType,
				},
			}
			Expect(UpdateClusterOpsAnnotations(ctx, k8sClient, cluster, opsRecordSlice)).Should(Succeed())

			By("test GetOpsRequestSliceFromCluster function")
			opsRecordSlice, _ = GetOpsRequestSliceFromCluster(cluster)
			Expect(len(opsRecordSlice) == 2 && opsRecordSlice[0].Name == testOpsName).Should(BeTrue())

			By("test no OpsRequest annotation in cluster")
			Expect(UpdateClusterOpsAnnotations(ctx, k8sClient, cluster, nil)).Should(Succeed())
			opsRecordSlice, _ = GetOpsRequestSliceFromCluster(cluster)
			Expect(len(opsRecordSlice) == 0).Should(BeTrue())
		})

		It("Should Test Backup OpsRequest", func() {
			By("test GetOpsRequestFromBackup function")
			backup := &dpv1alpha1.Backup{}
			backup.Labels = map[string]string{
				intctrlutil.OpsRequestNameLabelKey: "backup-ops",
				intctrlutil.OpsRequestTypeLabelKey: string(opsv1alpha1.BackupType),
			}
			Expect(GetOpsRequestFromBackup(backup)).ShouldNot(BeNil())

			By("test GetOpsRequestFromBackup function without ops type")
			backup.Labels = map[string]string{
				intctrlutil.OpsRequestNameLabelKey: "backup-ops",
			}
			Expect(GetOpsRequestFromBackup(backup)).Should(BeNil())

			By("test GetOpsRequestFromBackup function without ops name")
			backup.Labels = map[string]string{
				intctrlutil.OpsRequestTypeLabelKey: string(opsv1alpha1.BackupType),
			}
			Expect(GetOpsRequestFromBackup(backup)).Should(BeNil())
		})
	})
})
