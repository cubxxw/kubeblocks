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

package dataprotection

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	testclocks "k8s.io/utils/clock/testing"
	"sigs.k8s.io/controller-runtime/pkg/client"

	dpv1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	"github.com/apecloud/kubeblocks/pkg/constant"
	dprestore "github.com/apecloud/kubeblocks/pkg/dataprotection/restore"
	dptypes "github.com/apecloud/kubeblocks/pkg/dataprotection/types"
	dputils "github.com/apecloud/kubeblocks/pkg/dataprotection/utils"
	"github.com/apecloud/kubeblocks/pkg/generics"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
	testdp "github.com/apecloud/kubeblocks/pkg/testutil/dataprotection"
	viper "github.com/apecloud/kubeblocks/pkg/viperx"
)

var _ = Describe("Restore Controller test", func() {
	const namespace2 = "test2"

	cleanEnv := func() {
		// must wait till resources deleted and no longer existed before the testcases start,
		// otherwise if later it needs to create some new resource objects with the same name,
		// in race conditions, it will find the existence of old objects, resulting failure to
		// create the new objects.
		By("clean resources")

		ml := client.HasLabels{testCtx.TestObjLabelKey}

		cleanNamespaced := func(namespace string) {
			// delete rest mocked objects
			inNS := client.InNamespace(namespace)

			// namespaced
			testapps.ClearResources(&testCtx, generics.ClusterSignature, inNS, ml)
			testapps.ClearResources(&testCtx, generics.PodSignature, inNS, ml)
			testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.BackupSignature, true, inNS)

			// wait all backup to be deleted, otherwise the controller maybe create
			// job to delete the backup between the ClearResources function delete
			// the job and get the job list, resulting the ClearResources panic.
			Eventually(testapps.List(&testCtx, generics.BackupSignature, inNS)).Should(HaveLen(0))

			testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.JobSignature, true, inNS)
			testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.RestoreSignature, true, inNS)
			testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.PersistentVolumeClaimSignature, true, inNS)
			testapps.ClearResources(&testCtx, generics.SecretSignature, inNS, ml)
		}

		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.BackupRepoSignature, true, ml)

		cleanNamespaced(testCtx.DefaultNamespace)
		cleanNamespaced(namespace2)

		// non-namespaced
		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.ActionSetSignature, true, ml)
		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.StorageClassSignature, true, ml)
		testapps.ClearResources(&testCtx, generics.StorageProviderSignature, ml)
		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.PersistentVolumeSignature, true, ml)
	}

	ensureNamespace := func(name string) {
		Eventually(func(g Gomega) {
			obj := &corev1.Namespace{}
			obj.Name = name
			err := testCtx.Cli.Get(testCtx.Ctx, client.ObjectKeyFromObject(obj), &corev1.Namespace{})
			if err == nil {
				return
			}
			g.Expect(client.IgnoreNotFound(err)).Should(Succeed())
			err = testCtx.Cli.Create(testCtx.Ctx, obj)
			g.Expect(err).Should(Succeed())
		}).Should(Succeed())
	}

	BeforeEach(func() {
		cleanEnv()
		ensureNamespace(namespace2)
	})

	AfterEach(func() {
		cleanEnv()
	})

	When("restore controller test", func() {
		var (
			repo        *dpv1alpha1.BackupRepo
			repoPVCName string
			actionSet   *dpv1alpha1.ActionSet
			nodeName    = "minikube"
		)

		BeforeEach(func() {
			By("creating an actionSet")
			actionSet = testdp.NewFakeActionSet(&testCtx, nil)

			By("creating storage provider")
			_ = testdp.NewFakeStorageProvider(&testCtx, nil)

			By("creating a backupRepo")
			repo, repoPVCName = testdp.NewFakeBackupRepo(&testCtx, nil)
		})

		initResourcesAndWaitRestore := func(
			mockBackupCompleted,
			useVolumeSnapshot,
			isSerialPolicy bool,
			backupType dpv1alpha1.BackupType,
			expectRestorePhase dpv1alpha1.RestorePhase,
			change func(f *testdp.MockRestoreFactory),
			changeBackupStatus func(b *dpv1alpha1.Backup),
			backupNames ...string,
		) *dpv1alpha1.Restore {
			By("create a completed backup")
			backup := mockBackupForRestore(actionSet.Name, repo.Name, repoPVCName, mockBackupCompleted,
				useVolumeSnapshot, backupType, backupNames...)
			if changeBackupStatus != nil {
				Expect(testapps.ChangeObjStatus(&testCtx, backup, func() {
					changeBackupStatus(backup)
				})).Should(Succeed())
			}
			By("create restore ")
			schedulingSpec := dpv1alpha1.SchedulingSpec{
				NodeName: nodeName,
			}
			restoreFactory := testdp.NewRestoreFactory(testCtx.DefaultNamespace, testdp.RestoreName).
				SetBackup(backup.Name, testCtx.DefaultNamespace).
				SetSchedulingSpec(schedulingSpec)

			change(restoreFactory)

			if isSerialPolicy {
				restoreFactory.SetVolumeClaimRestorePolicy(dpv1alpha1.VolumeClaimRestorePolicySerial)
			}
			restore := restoreFactory.Create(&testCtx).GetObject()

			By(fmt.Sprintf("wait for restore is %s", expectRestorePhase))
			restoreKey := client.ObjectKeyFromObject(restore)
			Eventually(testapps.CheckObj(&testCtx, restoreKey, func(g Gomega, r *dpv1alpha1.Restore) {
				g.Expect(r.Status.Phase).Should(Equal(expectRestorePhase))
			})).Should(Succeed())
			return restore
		}

		checkJobAndPVCSCount := func(restore *dpv1alpha1.Restore, jobReplicas, pvcReplicas, startingIndex int) {
			Eventually(testapps.List(&testCtx, generics.JobSignature,
				client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
				client.InNamespace(testCtx.DefaultNamespace))).Should(HaveLen(jobReplicas))

			pvcMatchingLabels := client.MatchingLabels{constant.AppManagedByLabelKey: "restore"}
			Eventually(testapps.List(&testCtx, generics.PersistentVolumeClaimSignature, pvcMatchingLabels,
				client.InNamespace(testCtx.DefaultNamespace))).Should(HaveLen(pvcReplicas))

			By(fmt.Sprintf("pvc index should greater than or equal to %d", startingIndex))
			pvcList := &corev1.PersistentVolumeClaimList{}
			Expect(k8sClient.List(ctx, pvcList, pvcMatchingLabels,
				client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())
			for _, v := range pvcList.Items {
				indexStr := string(v.Name[len(v.Name)-1])
				index, _ := strconv.Atoi(indexStr)
				Expect(index >= startingIndex).Should(BeTrue())
			}
		}

		checkJobSA := func(restore *dpv1alpha1.Restore, saName string) {
			jobList := &batchv1.JobList{}
			Expect(k8sClient.List(ctx, jobList,
				client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
				client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())
			for _, v := range jobList.Items {
				Expect(v.Spec.Template.Spec.ServiceAccountName).WithOffset(1).
					Should(Equal(saName))
			}
		}

		mockRestoreJobsCompleted := func(restore *dpv1alpha1.Restore) {
			jobList := &batchv1.JobList{}
			Expect(k8sClient.List(ctx, jobList,
				client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
				client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())
			for _, v := range jobList.Items {
				testdp.PatchK8sJobStatus(&testCtx, client.ObjectKeyFromObject(&v), batchv1.JobComplete)
			}
		}

		mockAndCheckRestoreCompleted := func(restore *dpv1alpha1.Restore) {
			By("mock jobs are completed")
			mockRestoreJobsCompleted(restore)

			By("wait for restore is completed")
			Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(restore), func(g Gomega, r *dpv1alpha1.Restore) {
				g.Expect(r.Status.Phase).Should(Equal(dpv1alpha1.RestorePhaseCompleted))
			})).Should(Succeed())
		}

		testRestoreWithVolumeClaimsTemplate := func(replicas, startingIndex int) {
			restore := initResourcesAndWaitRestore(true, false, false, "", dpv1alpha1.RestorePhaseRunning,
				func(f *testdp.MockRestoreFactory) {
					f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
						testdp.DataVolumeMountPath, "", int32(replicas), int32(startingIndex), nil)
					// Note: should ignore this policy when podSelectionStrategy is Any of the source target.
					f.SetPrepareDataRequiredPolicy(dpv1alpha1.OneToOneRestorePolicy, "")
				}, nil)

			By("expect restore jobs and pvcs are created")
			checkJobAndPVCSCount(restore, replicas, replicas, startingIndex)

			checkJobSA(restore, viper.GetString(dptypes.CfgKeyWorkerServiceAccountName))

			By("mock jobs are completed and wait for restore is completed")
			mockAndCheckRestoreCompleted(restore)
		}
		checkJobParametersEnv := func(restore *dpv1alpha1.Restore) {

			By("check parameters env in restore jobs")
			jobList := &batchv1.JobList{}
			Expect(k8sClient.List(ctx, jobList,
				client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
				client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())
			for _, job := range jobList.Items {
				Expect(len(job.Spec.Template.Spec.Containers)).ShouldNot(BeZero())
				for _, c := range job.Spec.Template.Spec.Containers {
					if c.Name != dprestore.Restore {
						continue
					}
					count := 0
					for _, env := range c.Env {
						for _, param := range testdp.TestParameters {
							if param.Name == env.Name && param.Value == env.Value {
								count++
							}
						}
					}
					Expect(count).To(Equal(len(testdp.TestParameters)))
				}
			}
		}
		Context("with restore fails", func() {
			It("test restore is Failed when backup is not completed", func() {
				By("expect for restore is Failed ")
				initResourcesAndWaitRestore(false, false, true, "", dpv1alpha1.RestorePhaseFailed,
					func(f *testdp.MockRestoreFactory) {
						f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
							testdp.DataVolumeMountPath, "", int32(3), int32(0), nil)
					}, nil)
			})

			It("test restore is failed when check failed in new action", func() {
				By("expect for restore is Failed")
				restore := initResourcesAndWaitRestore(true, false, true, "", dpv1alpha1.RestorePhaseFailed,
					func(f *testdp.MockRestoreFactory) {
						f.Get().Spec.Backup.Name = "wrongBackup"
					}, nil)
				By("check status.conditions")
				Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(restore), func(g Gomega, r *dpv1alpha1.Restore) {
					val := meta.IsStatusConditionFalse(r.Status.Conditions, dprestore.ConditionTypeRestoreCheckBackupRepo)
					g.Expect(val).Should(BeTrue())
				})).Should(Succeed())
			})

			It("test restore is failed when validate failed in new action", func() {
				By("expect for restore is Failed")
				restore := initResourcesAndWaitRestore(false, false, true, "", dpv1alpha1.RestorePhaseFailed, func(f *testdp.MockRestoreFactory) {
					f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
						testdp.DataVolumeMountPath, "", int32(3), int32(0), nil)
				}, nil)
				By("check status.conditions")
				Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(restore), func(g Gomega, r *dpv1alpha1.Restore) {
					val := meta.IsStatusConditionFalse(r.Status.Conditions, dprestore.ConditionTypeRestoreValidationPassed)
					g.Expect(val).Should(BeTrue())
				})).Should(Succeed())
			})

			It("test restore is Failed when restore job is not Failed", func() {
				By("expect for restore is Failed ")
				restore := initResourcesAndWaitRestore(true, false, true, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
							testdp.DataVolumeMountPath, "", int32(3), int32(0), nil)
					}, nil)

				By("wait for creating first job and pvc")
				checkJobAndPVCSCount(restore, 1, 1, 0)

				By("mock restore job is Failed")
				jobList := &batchv1.JobList{}
				Expect(k8sClient.List(ctx, jobList,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())

				for _, v := range jobList.Items {
					testdp.PatchK8sJobStatus(&testCtx, client.ObjectKeyFromObject(&v), batchv1.JobFailed)
				}

				Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(restore), func(g Gomega, r *dpv1alpha1.Restore) {
					g.Expect(r.Status.Phase).Should(Equal(dpv1alpha1.RestorePhaseFailed))
				})).Should(Succeed())
			})
		})

		Context("test prepareData stage", func() {
			It("test volumeClaimsTemplate when startingIndex is 0", func() {
				testRestoreWithVolumeClaimsTemplate(3, 0)
			})

			It("test volumeClaimsTemplate when startingIndex is 1", func() {
				testRestoreWithVolumeClaimsTemplate(2, 1)
			})
			It("test restore parameters", func() {
				By("set schema and parameters in actionSet")
				testdp.MockActionSetWithSchema(&testCtx, actionSet)
				replicas := 3
				startingIndex := 0
				restore := initResourcesAndWaitRestore(true, false, false, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
							testdp.DataVolumeMountPath, "", int32(replicas), int32(startingIndex), nil)
						// Note: should ignore this policy when podSelectionStrategy is Any of the source target.
						f.SetPrepareDataRequiredPolicy(dpv1alpha1.OneToOneRestorePolicy, "")
						f.SetParameters(testdp.TestParameters)
					}, nil)

				By("expect restore jobs and pvcs are created")
				checkJobAndPVCSCount(restore, replicas, replicas, startingIndex)
				By("expect parameters env in restore jobs")
				checkJobParametersEnv(restore)
			})
			It("test volumeClaimsTemplate when volumeClaimRestorePolicy is Serial", func() {
				replicas := 2
				startingIndex := 1
				restore := initResourcesAndWaitRestore(true, false, true, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
							testdp.DataVolumeMountPath, "", int32(replicas), int32(startingIndex), nil)
					}, nil)

				By("wait for creating first job and pvc")
				checkJobAndPVCSCount(restore, 1, 1, startingIndex)

				By("mock jobs are completed")
				mockRestoreJobsCompleted(restore)

				var firstJobName string
				Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(restore), func(g Gomega, r *dpv1alpha1.Restore) {
					g.Expect(r.Status.Actions.PrepareData).ShouldNot(BeEmpty())
					g.Expect(r.Status.Actions.PrepareData[0].Status).Should(Equal(dpv1alpha1.RestoreActionCompleted))
					firstJobName = strings.ReplaceAll(r.Status.Actions.PrepareData[0].ObjectKey, "Job/", "")
				})).Should(Succeed())

				By("wait for deleted first job")
				Eventually(testapps.CheckObjExists(&testCtx,
					types.NamespacedName{Name: firstJobName, Namespace: testCtx.DefaultNamespace}, &batchv1.Job{}, false)).Should(Succeed())

				By("after the first job is completed, next job will be created")
				checkJobAndPVCSCount(restore, 1, replicas, startingIndex)

				jobList := &batchv1.JobList{}
				Expect(k8sClient.List(ctx, jobList,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())

				for _, v := range jobList.Items {
					Expect(v.Labels[constant.AppManagedByLabelKey]).Should(Equal(dptypes.AppName))
					finished, _, _ := dputils.IsJobFinished(&v)
					Expect(finished).Should(BeFalse())
				}

				By("mock jobs are completed and wait for restore is completed")
				mockAndCheckRestoreCompleted(restore)
			})

			It("test dataSourceRef", func() {
				initResourcesAndWaitRestore(true, true, false, "", dpv1alpha1.RestorePhaseAsDataSource,
					func(f *testdp.MockRestoreFactory) {
						f.SetDataSourceRef(testdp.DataVolumeName, testdp.DataVolumeMountPath)
					}, nil)
			})

			It("test when dataRestorePolicy is OneToOne", func() {
				startingIndex := 0
				restoredReplicas := 2
				restore := initResourcesAndWaitRestore(true, false, false, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
							testdp.DataVolumeMountPath, "", int32(restoredReplicas), int32(startingIndex), nil)
						f.SetPrepareDataRequiredPolicy(dpv1alpha1.OneToOneRestorePolicy, "")
					}, func(b *dpv1alpha1.Backup) {
						b.Status.Target.PodSelector.Strategy = dpv1alpha1.PodSelectionStrategyAll
						b.Status.Target.SelectedTargetPods = []string{"pod-0", "pod-1"}
					})

				By("wait to create two jobs and pvcs")
				checkJobAndPVCSCount(restore, restoredReplicas, restoredReplicas, 0)

				jobList := &batchv1.JobList{}
				Expect(k8sClient.List(ctx, jobList,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())
				for _, v := range jobList.Items {
					var checkBackupBasePathPass bool
					index := v.Name[strings.LastIndex(v.Name, "-")+1:]
					// checks if the backupBasePath exits
					for _, env := range v.Spec.Template.Spec.Containers[0].Env {
						if env.Name == dptypes.DPBackupBasePath && strings.Contains(env.Value, fmt.Sprintf("pod-%s", index)) {
							checkBackupBasePathPass = true
							break
						}
					}
					Expect(checkBackupBasePathPass).Should(BeTrue())
				}

				By("mock jobs are completed and wait for restore is completed")
				mockAndCheckRestoreCompleted(restore)
			})

			It("test when dataRestorePolicy is OneToMany and sourceTargetPod is pod-0", func() {
				startingIndex := 0
				restoredReplicas := 2
				sourcePodName := "pod-0"
				restore := initResourcesAndWaitRestore(true, false, false, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
							testdp.DataVolumeMountPath, "", int32(restoredReplicas), int32(startingIndex), nil)
						f.SetPrepareDataRequiredPolicy(dpv1alpha1.OneToManyRestorePolicy, sourcePodName)
					}, func(b *dpv1alpha1.Backup) {
						b.Status.Target.PodSelector.Strategy = dpv1alpha1.PodSelectionStrategyAll
						b.Status.Target.SelectedTargetPods = []string{sourcePodName, "pod-1"}
					})

				By("wait to create two jobs and pvcs")
				checkJobAndPVCSCount(restore, restoredReplicas, restoredReplicas, 0)
				jobList := &batchv1.JobList{}
				Expect(k8sClient.List(ctx, jobList,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())
				for _, v := range jobList.Items {
					var checkBackupBasePathPass bool
					// checks if the backupBasePath exits
					for _, env := range v.Spec.Template.Spec.Containers[0].Env {
						if env.Name == dptypes.DPBackupBasePath && strings.Contains(env.Value, sourcePodName) {
							checkBackupBasePathPass = true
							break
						}
					}
					Expect(checkBackupBasePathPass).Should(BeTrue())
				}

				By("mock jobs are completed and wait for restore is completed")
				mockAndCheckRestoreCompleted(restore)
			})

		})

		Context("test postReady stage", func() {
			var _ *testdp.BackupClusterInfo
			BeforeEach(func() {
				By("fake a new cluster")
				_ = testdp.NewFakeCluster(&testCtx)
			})

			It("test post ready actions", func() {
				By("remove the prepareData stage for testing post ready actions")
				Expect(testapps.ChangeObj(&testCtx, actionSet, func(set *dpv1alpha1.ActionSet) {
					set.Spec.Restore.PrepareData = nil
				})).Should(Succeed())

				matchLabels := map[string]string{
					constant.AppInstanceLabelKey: testdp.ClusterName,
				}
				restore := initResourcesAndWaitRestore(true, false, false, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetConnectCredential(testdp.ClusterName).SetJobActionConfig(matchLabels).SetExecActionConfig(matchLabels)
					}, nil)

				By("wait for creating two exec jobs with the matchLabels")
				Eventually(testapps.List(&testCtx, generics.JobSignature,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(HaveLen(2))

				checkJobSA(restore, viper.GetString(dptypes.CfgKeyExecWorkerServiceAccountName))

				By("mock exec jobs are completed")
				mockRestoreJobsCompleted(restore)

				By("wait for creating a job of jobAction with the matchLabels, expect jobs count is 3(2+1)")
				Eventually(testapps.List(&testCtx, generics.JobSignature,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(HaveLen(3))

				By("mock jobs are completed and wait for restore is completed")
				mockAndCheckRestoreCompleted(restore)

				By("test deleting restore")
				Expect(k8sClient.Delete(ctx, restore)).Should(Succeed())
				Eventually(testapps.CheckObjExists(&testCtx, client.ObjectKeyFromObject(restore), restore, false)).Should(Succeed())
			})

			It("test jobAction env", func() {
				By("remove the prepareData stage for testing post ready actions")
				Expect(testapps.ChangeObj(&testCtx, actionSet, func(set *dpv1alpha1.ActionSet) {
					set.Spec.Restore.PrepareData = nil
				})).Should(Succeed())

				matchLabels := map[string]string{
					constant.AppInstanceLabelKey: testdp.ClusterName,
				}

				restore := initResourcesAndWaitRestore(true, false, false, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetJobActionConfig(matchLabels).SetExecActionConfig(matchLabels)
					}, func(b *dpv1alpha1.Backup) {
						b.Status.Target.ConnectionCredential = nil
					})

				getJobKey := func(jobIndex int) client.ObjectKey {
					return client.ObjectKey{
						Name:      fmt.Sprintf("restore-post-ready-%s-%s-%d-%d", restore.UID[:8], restore.Spec.Backup.Name, 1, jobIndex),
						Namespace: restore.Namespace,
					}
				}

				getDPDBPortEnv := func(container *corev1.Container) corev1.EnvVar {
					for _, env := range container.Env {
						if env.Name == dptypes.DPDBPort {
							return env
						}
					}
					return corev1.EnvVar{}
				}
				By("wait for creating two exec jobs with the matchLabels")
				a := testapps.List(&testCtx, generics.JobSignature,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))
				Eventually(a).Should(HaveLen(2))

				checkJobSA(restore, viper.GetString(dptypes.CfgKeyExecWorkerServiceAccountName))

				By("mock exec jobs are completed")
				mockRestoreJobsCompleted(restore)

				By("wait for creating a job of jobAction with the matchLabels, expect jobs count is 3(2+1)")
				Eventually(testapps.List(&testCtx, generics.JobSignature,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(HaveLen(3))

				By("check backup job's port env")
				Eventually(testapps.CheckObj(&testCtx, getJobKey(0), func(g Gomega, fetched *batchv1.Job) {
					g.Expect(getDPDBPortEnv(&fetched.Spec.Template.Spec.Containers[0]).Value).Should(Equal(strconv.Itoa(testdp.PortNum)))
				})).Should(Succeed())

			})
			It("test parameters env", func() {
				By("set schema and parameters in actionSet")
				testdp.MockActionSetWithSchema(&testCtx, actionSet)
				By("remove the prepareData stage for testing post ready actions")
				Expect(testapps.ChangeObj(&testCtx, actionSet, func(set *dpv1alpha1.ActionSet) {
					set.Spec.Restore.PrepareData = nil
				})).Should(Succeed())

				matchLabels := map[string]string{
					constant.AppInstanceLabelKey: testdp.ClusterName,
				}

				restore := initResourcesAndWaitRestore(true, false, false, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetJobActionConfig(matchLabels).SetExecActionConfig(matchLabels)
						f.SetParameters(testdp.TestParameters)
					}, func(b *dpv1alpha1.Backup) {
						b.Status.Target.ConnectionCredential = nil
					})
				By("expect parameters env in restore jobs")
				checkJobParametersEnv(restore)
			})
		})

		Context("test cross namespace", func() {
			It("should wait for preparation of backup repo", func() {
				By("creating a restore in a different namespace from backup")
				initResourcesAndWaitRestore(true, false, true, "", dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetNamespace(namespace2)
					}, nil)
			})
		})

		Context("test restore from incremental backup", func() {
			var (
				baseBackup       *dpv1alpha1.Backup
				parentBackupName string
				ancestorBackups  = []*dpv1alpha1.Backup{}
				cnt              = 0
				testClock        = testclocks.NewFakeClock(time.Now().Add(time.Hour))
			)

			genIncBackupName := func() string {
				cnt++
				return fmt.Sprintf("inc-backup-%d", cnt)
			}

			changeTimeRange := func(backup *dpv1alpha1.Backup) {
				backup.Status.TimeRange = &dpv1alpha1.BackupTimeRange{
					Start: &metav1.Time{},
					End:   &metav1.Time{},
				}
				// testClock is only used to mock timestamps in ascending order, Step() don't affect the real or logic time of test
				testClock.Step(time.Minute)
				backup.Status.TimeRange.Start.Time = testClock.Now()
				testClock.Step(time.Minute)
				backup.Status.TimeRange.End.Time = testClock.Now()
			}

			BeforeEach(func() {
				By("mock completed full backup and parent incremental backup")
				baseBackup = mockBackupForRestore(actionSet.Name, repo.Name, repoPVCName, true, false, dpv1alpha1.BackupTypeFull)
				Expect(testapps.ChangeObjStatus(&testCtx, baseBackup, func() { changeTimeRange(baseBackup) })).Should(Succeed())
				actionSet = testdp.NewFakeIncActionSet(&testCtx)
				parentBackupName = baseBackup.Name
				for i := 0; i < 3; i++ {
					backup := mockBackupForRestore(actionSet.Name, repo.Name, repoPVCName, true, false, dpv1alpha1.BackupTypeIncremental,
						genIncBackupName(), parentBackupName, baseBackup.Name)
					Expect(testapps.ChangeObjStatus(&testCtx, backup, func() { changeTimeRange(backup) })).Should(Succeed())
					ancestorBackups = append(ancestorBackups, backup)
					parentBackupName = backup.Name
				}
			})

			AfterEach(func() {
				ancestorBackups = []*dpv1alpha1.Backup{}
				cnt = 0
			})

			It("test restore from incremental backup", func() {
				replicas, startingIndex := 3, 0
				restore := initResourcesAndWaitRestore(true, false, false, dpv1alpha1.BackupTypeIncremental, dpv1alpha1.RestorePhaseRunning,
					func(f *testdp.MockRestoreFactory) {
						f.SetVolumeClaimsTemplate(testdp.MysqlTemplateName, testdp.DataVolumeName,
							testdp.DataVolumeMountPath, "", int32(replicas), int32(startingIndex), nil)
						f.SetPrepareDataRequiredPolicy(dpv1alpha1.OneToOneRestorePolicy, "")
					}, changeTimeRange, genIncBackupName(), parentBackupName, baseBackup.Name)

				By("wait for creating jobs and pvcs")
				checkJobAndPVCSCount(restore, replicas, replicas, 0)
				By("check job env")
				ancestorIncrementalBackupNames := []string{}
				for _, backup := range ancestorBackups {
					ancestorIncrementalBackupNames = append(ancestorIncrementalBackupNames, backup.Name)
				}
				expectedEnv := map[string]string{
					dptypes.DPAncestorIncrementalBackupNames: strings.Join(ancestorIncrementalBackupNames, ","),
					dptypes.DPBaseBackupName:                 baseBackup.Name,
				}
				jobList := &batchv1.JobList{}
				Expect(k8sClient.List(ctx, jobList,
					client.MatchingLabels{dprestore.DataProtectionRestoreLabelKey: restore.Name},
					client.InNamespace(testCtx.DefaultNamespace))).Should(Succeed())
				for _, job := range jobList.Items {
					cnt := 0
					for _, env := range job.Spec.Template.Spec.Containers[0].Env {
						if value, ok := expectedEnv[env.Name]; ok {
							Expect(env.Value).Should(Equal(value))
							cnt++
						}
					}
					Expect(cnt).To(Equal(len(expectedEnv)))
				}
				By("mock jobs are completed and wait for restore is completed")
				mockAndCheckRestoreCompleted(restore)
			})
		})
	})
})

func mockBackupForRestore(
	actionSetName, repoName, backupPVCName string,
	mockBackupCompleted, useVolumeSnapshotBackup bool,
	backupType dpv1alpha1.BackupType,
	backupNames ...string,
) *dpv1alpha1.Backup {
	backup := testdp.NewFakeBackup(&testCtx, func(backup *dpv1alpha1.Backup) {
		if len(backupNames) > 0 {
			backup.Name = backupNames[0]
		}
		if backupType == dpv1alpha1.BackupTypeIncremental {
			if len(backupNames) > 1 {
				backup.Spec.ParentBackupName = backupNames[1]
			}
			backup.Spec.BackupMethod = testdp.IncBackupMethodName
		}
	})
	// wait for backup is failed by backup controller.
	// it will be failed if the backupPolicy is not created.
	Eventually(testapps.CheckObj(&testCtx, client.ObjectKeyFromObject(backup), func(g Gomega, tmpBackup *dpv1alpha1.Backup) {
		g.Expect(tmpBackup.Status.Phase).Should(Equal(dpv1alpha1.BackupPhaseFailed))
	})).Should(Succeed())

	if mockBackupCompleted {
		// then mock backup to completed
		Expect(testapps.ChangeObjStatus(&testCtx, backup, func() {
			backupMethodName := testdp.BackupMethodName
			if useVolumeSnapshotBackup {
				backupMethodName = testdp.VSBackupMethodName
				testdp.MockBackupVSStatusActions(backup)
			}
			if backupType == dpv1alpha1.BackupTypeIncremental {
				backupMethodName = testdp.IncBackupMethodName
				if len(backupNames) > 2 {
					backup.Status.ParentBackupName = backupNames[1]
					backup.Status.BaseBackupName = backupNames[2]
				}
			}
			backup.Status.Path = "/backup-data" + "/" + backup.Name
			backup.Status.Phase = dpv1alpha1.BackupPhaseCompleted
			backup.Status.BackupRepoName = repoName
			backup.Status.PersistentVolumeClaimName = backupPVCName
			testdp.MockBackupStatusTarget(backup, dpv1alpha1.PodSelectionStrategyAny)
			backup.Status.Target.ContainerPort = &dpv1alpha1.ContainerPort{
				ContainerName: testdp.ContainerName + "-1",
				PortName:      testdp.PortName,
			}
			testdp.MockBackupStatusMethod(backup, backupMethodName, testdp.DataVolumeName, actionSetName)
		})).Should(Succeed())
	}
	return backup
}
