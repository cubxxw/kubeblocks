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

package component

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	"github.com/apecloud/kubeblocks/pkg/constant"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	"github.com/apecloud/kubeblocks/pkg/generics"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
)

var _ = Describe("service references", func() {
	cleanEnv := func() {
		// must wait till resources deleted and no longer existed before the testcases start,
		// otherwise if later it needs to create some new resource objects with the same name,
		// in race conditions, it will find the existence of old objects, resulting failure to
		// create the new objects.
		By("clean resources")

		inNS := client.InNamespace(testCtx.DefaultNamespace)
		ml := client.HasLabels{testCtx.TestObjLabelKey}

		// resources should be released in following order
		// non-namespaced
		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.ClusterDefinitionSignature, true, ml)

		// namespaced
		testapps.ClearResourcesWithRemoveFinalizerOption(&testCtx, generics.ConfigMapSignature, true, inNS, ml)
	}

	var (
		namespace   = "default"
		clusterName = "cluster"
	)

	BeforeEach(func() {
		cleanEnv()
	})

	AfterEach(func() {
		cleanEnv()
	})

	Context("service descriptor", func() {
		It("service version regex validation test", func() {
			type versionCmp struct {
				serviceRefDeclRegex      string
				serviceDescriptorVersion string
			}
			tests := []struct {
				name   string
				fields versionCmp
				want   bool
			}{{
				name: "version string test true",
				fields: versionCmp{
					serviceRefDeclRegex:      "8.0.8",
					serviceDescriptorVersion: "8.0.8",
				},
				want: true,
			}, {
				name: "version string test false",
				fields: versionCmp{
					serviceRefDeclRegex:      "8.0.8",
					serviceDescriptorVersion: "8.0.7",
				},
				want: false,
			}, {
				name: "version string test false",
				fields: versionCmp{
					serviceRefDeclRegex:      "^8.0.8$",
					serviceDescriptorVersion: "v8.0.8",
				},
				want: false,
			}, {
				name: "version string test true",
				fields: versionCmp{
					serviceRefDeclRegex:      "8.0.\\d{1,2}$",
					serviceDescriptorVersion: "8.0.6",
				},
				want: true,
			}, {
				name: "version string test false",
				fields: versionCmp{
					serviceRefDeclRegex:      "8.0.\\d{1,2}$",
					serviceDescriptorVersion: "8.0.8.8.8",
				},
				want: false,
			}, {
				name: "version string test true",
				fields: versionCmp{
					serviceRefDeclRegex:      "^[v\\-]*?(\\d{1,2}\\.){0,3}\\d{1,2}$",
					serviceDescriptorVersion: "v-8.0.8.0",
				},
				want: true,
			}, {
				name: "version string test false",
				fields: versionCmp{
					serviceRefDeclRegex:      "^[v\\-]*?(\\d{1,2}\\.){0,3}\\d{1,2}$",
					serviceDescriptorVersion: "mysql-8.0.8",
				},
				want: false,
			}}
			for _, tt := range tests {
				match := verifyServiceVersion(tt.fields.serviceDescriptorVersion, tt.fields.serviceRefDeclRegex)
				Expect(match).Should(Equal(tt.want))
			}
		})
	})

	Context("service reference from new cluster objects", func() {
		const (
			etcd          = "etcd"
			etcdVersion   = "v3.5.6"
			etcdCluster   = "etcd"
			etcdComponent = "etcd"
		)

		var (
			compDef               *appsv1.ComponentDefinition
			comp                  *appsv1.Component
			synthesizedComp       *SynthesizedComponent
			serviceRefDeclaration = appsv1.ServiceRefDeclaration{
				Name: etcd,
				ServiceRefDeclarationSpecs: []appsv1.ServiceRefDeclarationSpec{
					{
						ServiceKind:    etcd,
						ServiceVersion: etcdVersion,
					},
				},
			}
		)

		BeforeEach(func() {
			compDef = &appsv1.ComponentDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "compdef",
				},
				Spec: appsv1.ComponentDefinitionSpec{
					ServiceRefDeclarations: []appsv1.ServiceRefDeclaration{serviceRefDeclaration},
				},
			}
			comp = &appsv1.Component{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: namespace,
					Name:      "comp",
				},
				Spec: appsv1.ComponentSpec{
					ServiceRefs: []appsv1.ServiceRef{},
				},
			}
			synthesizedComp = &SynthesizedComponent{
				Namespace:   namespace,
				ClusterName: clusterName,
			}
		})

		It("has service-ref not defined", func() {
			comp.Spec.CompDef = compDef.GetName()
			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, testCtx.Cli, synthesizedComp, compDef, comp)
			Expect(err).ShouldNot(Succeed())
			Expect(err.Error()).Should(ContainSubstring("service-ref for %s is not defined", serviceRefDeclaration.Name))

			// set the service-ref as optional
			compDef.Spec.ServiceRefDeclarations[0].Optional = func() *bool { optional := true; return &optional }()
			err = buildServiceReferencesWithoutResolve(testCtx.Ctx, testCtx.Cli, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())
			Expect(synthesizedComp.ServiceReferences).Should(HaveLen(0))
		})

		It("service vars - cluster service", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name: serviceRefDeclaration.Name,
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						Service: &appsv1.ServiceRefServiceSelector{
							Service: "client",
							Port:    "client",
						},
					},
				},
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					&corev1.Service{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      constant.GenerateClusterServiceName(etcdCluster, "client"),
						},
						Spec: corev1.ServiceSpec{
							Ports: []corev1.ServicePort{
								{
									Name: "peer",
									Port: 2380,
								},
								{
									Name: "client",
									Port: 2379,
								},
							},
						},
					},
				},
			}

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())

			Expect(synthesizedComp.ServiceReferences).Should(HaveKey(serviceRefDeclaration.Name))
			serviceDescriptor := synthesizedComp.ServiceReferences[serviceRefDeclaration.Name]
			Expect(serviceDescriptor).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint.Value).Should(Equal(fmt.Sprintf("%s:%s", reader.objs[0].GetName(), "2379")))
			Expect(serviceDescriptor.Spec.Host).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Host.Value).Should(Equal(reader.objs[0].GetName()))
			Expect(serviceDescriptor.Spec.Port).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Port.Value).Should(Equal("2379"))
			Expect(serviceDescriptor.Spec.Auth).Should(BeNil())
		})

		It("service vars - component service", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name: serviceRefDeclaration.Name,
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						Service: &appsv1.ServiceRefServiceSelector{
							Component: etcdComponent,
							Service:   "", // default service
							Port:      "peer",
						},
					},
				},
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					&corev1.Service{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      constant.GenerateComponentServiceName(etcdCluster, etcdComponent, ""),
						},
						Spec: corev1.ServiceSpec{
							Ports: []corev1.ServicePort{
								{
									Name: "peer",
									Port: 2380,
								},
								{
									Name: "client",
									Port: 2379,
								},
							},
						},
					},
				},
			}

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())

			Expect(synthesizedComp.ServiceReferences).Should(HaveKey(serviceRefDeclaration.Name))
			serviceDescriptor := synthesizedComp.ServiceReferences[serviceRefDeclaration.Name]
			Expect(serviceDescriptor).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint.Value).Should(Equal(fmt.Sprintf("%s:%s", reader.objs[0].GetName(), "2380")))
			Expect(serviceDescriptor.Spec.Host).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Host.Value).Should(Equal(reader.objs[0].GetName()))
			Expect(serviceDescriptor.Spec.Port).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Port.Value).Should(Equal("2380"))
			Expect(serviceDescriptor.Spec.Auth).Should(BeNil())
		})

		It("service vars - pod service", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name: serviceRefDeclaration.Name,
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						Service: &appsv1.ServiceRefServiceSelector{
							Component: etcdComponent,
							Service:   "peer",
							Port:      "peer",
						},
					},
				},
			}
			newPodService := func(ordinal int) *corev1.Service {
				return &corev1.Service{
					ObjectMeta: metav1.ObjectMeta{
						Namespace: namespace,
						Name:      fmt.Sprintf("%s-%d", constant.GenerateComponentServiceName(etcdCluster, etcdComponent, "peer"), ordinal),
						Labels:    constant.GetCompLabels(etcdCluster, etcdComponent),
					},
					Spec: corev1.ServiceSpec{
						Ports: []corev1.ServicePort{
							{
								Name: "peer",
								Port: 2380,
							},
							{
								Name: "client",
								Port: 2379,
							},
						},
					},
				}
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					newPodService(0),
					newPodService(1),
					newPodService(2),
					&appsv1.Component{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      FullName(etcdCluster, etcdComponent),
						},
						Spec: appsv1.ComponentSpec{
							CompDef: "test-compdef",
						},
					},
					&appsv1.ComponentDefinition{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test-compdef",
						},
						Spec: appsv1.ComponentDefinitionSpec{
							Services: []appsv1.ComponentService{
								{
									Service: appsv1.Service{
										Name:        "peer",
										ServiceName: "peer",
									},
								},
							},
						},
					},
				},
			}

			hosts, ports := make([]string, 0), make([]string, 0)
			for i := 0; i < 3; i++ {
				hosts = append(hosts, reader.objs[i].GetName())
				ports = append(ports, fmt.Sprintf("%s:%s", reader.objs[i].GetName(), "2380"))
			}
			expectedHostValue, expectedPortValue := strings.Join(hosts, ","), strings.Join(ports, ",")

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())

			Expect(synthesizedComp.ServiceReferences).Should(HaveKey(serviceRefDeclaration.Name))
			serviceDescriptor := synthesizedComp.ServiceReferences[serviceRefDeclaration.Name]
			Expect(serviceDescriptor).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint.Value).Should(Equal(expectedPortValue))
			Expect(serviceDescriptor.Spec.Host).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Host.Value).Should(Equal(expectedHostValue))
			Expect(serviceDescriptor.Spec.Port).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Port.Value).Should(Equal(expectedPortValue))
			Expect(serviceDescriptor.Spec.Auth).Should(BeNil())
		})

		It("service vars - different namespace", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name:      serviceRefDeclaration.Name,
					Namespace: "external",
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						Service: &appsv1.ServiceRefServiceSelector{
							Service: "client",
							Port:    "client",
						},
					},
				},
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					&corev1.Service{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "external",
							Name:      constant.GenerateClusterServiceName(etcdCluster, "client"),
						},
						Spec: corev1.ServiceSpec{
							Ports: []corev1.ServicePort{
								{
									Name: "peer",
									Port: 2380,
								},
								{
									Name: "client",
									Port: 2379,
								},
							},
						},
					},
				},
			}

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())

			svcFQDN := intctrlutil.ServiceFQDN("external", reader.objs[0].GetName())

			Expect(synthesizedComp.ServiceReferences).Should(HaveKey(serviceRefDeclaration.Name))
			serviceDescriptor := synthesizedComp.ServiceReferences[serviceRefDeclaration.Name]
			Expect(serviceDescriptor).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Endpoint.Value).Should(Equal(fmt.Sprintf("%s:%s", svcFQDN, "2379")))
			Expect(serviceDescriptor.Spec.Host).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Host.Value).Should(Equal(svcFQDN))
			Expect(serviceDescriptor.Spec.Port).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Port.Value).Should(Equal("2379"))
			Expect(serviceDescriptor.Spec.Auth).Should(BeNil())
		})

		It("credential vars - same namespace", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name:      serviceRefDeclaration.Name,
					Namespace: namespace,
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						Credential: &appsv1.ServiceRefCredentialSelector{
							Component: etcdComponent,
							Name:      "default",
						},
					},
				},
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					&corev1.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      constant.GenerateAccountSecretName(etcdCluster, etcdComponent, "default"),
						},
						Data: map[string][]byte{
							constant.AccountNameForSecret:   []byte("username"),
							constant.AccountPasswdForSecret: []byte("password"),
						},
					},
				},
			}

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())

			Expect(synthesizedComp.ServiceReferences).Should(HaveKey(serviceRefDeclaration.Name))
			serviceDescriptor := synthesizedComp.ServiceReferences[serviceRefDeclaration.Name]
			Expect(serviceDescriptor).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Username).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Username.ValueFrom).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Username.ValueFrom).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Username.ValueFrom.SecretKeyRef).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Username.ValueFrom.SecretKeyRef.Name).Should(Equal(reader.objs[0].GetName()))
			Expect(serviceDescriptor.Spec.Auth.Username.ValueFrom.SecretKeyRef.Key).Should(Equal(constant.AccountNameForSecret))
			Expect(serviceDescriptor.Spec.Auth.Password).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Password.ValueFrom).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Password.ValueFrom.SecretKeyRef).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.Auth.Password.ValueFrom.SecretKeyRef.Name).Should(Equal(reader.objs[0].GetName()))
			Expect(serviceDescriptor.Spec.Auth.Password.ValueFrom.SecretKeyRef.Key).Should(Equal(constant.AccountPasswdForSecret))
			Expect(serviceDescriptor.Spec.Endpoint).Should(BeNil())
			Expect(serviceDescriptor.Spec.Host).Should(BeNil())
			Expect(serviceDescriptor.Spec.Port).Should(BeNil())
		})

		It("credential vars - different namespace", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name:      serviceRefDeclaration.Name,
					Namespace: "external",
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						Credential: &appsv1.ServiceRefCredentialSelector{
							Component: etcdComponent,
							Name:      "default",
						},
					},
				},
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					&corev1.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: "external",
							Name:      constant.GenerateAccountSecretName(etcdCluster, etcdComponent, "default"),
						},
						Data: map[string][]byte{
							constant.AccountNameForSecret:   []byte("username"),
							constant.AccountPasswdForSecret: []byte("password"),
						},
					},
				},
			}

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).ShouldNot(Succeed())
			Expect(err.Error()).Should(ContainSubstring("prohibits referencing credential variables from different namespaces"))
		})

		It("component vars - pod FQDNs", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name: serviceRefDeclaration.Name,
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						PodFQDNs: &appsv1.ServiceRefPodFQDNsSelector{
							Component: etcdComponent,
						},
					},
				},
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					&appsv1.Component{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      constant.GenerateClusterComponentName(etcdCluster, etcdComponent),
						},
						Spec: appsv1.ComponentSpec{
							Replicas: 2,
						},
					},
				},
			}

			etcdComp := reader.objs[0].(*appsv1.Component)
			podNamePrefix := constant.GenerateWorkloadNamePattern(etcdCluster, etcdComponent) + "-"
			expectedPodFQDNs := strings.Join([]string{
				intctrlutil.PodFQDN(namespace, etcdComp.Name, podNamePrefix+"0"),
				intctrlutil.PodFQDN(namespace, etcdComp.Name, podNamePrefix+"1"),
			}, ",")

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())

			Expect(synthesizedComp.ServiceReferences).Should(HaveKey(serviceRefDeclaration.Name))
			serviceDescriptor := synthesizedComp.ServiceReferences[serviceRefDeclaration.Name]
			Expect(serviceDescriptor).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.PodFQDNs).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.PodFQDNs.Value).Should(Equal(expectedPodFQDNs))
			Expect(serviceDescriptor.Spec.Endpoint).Should(BeNil())
			Expect(serviceDescriptor.Spec.Host).Should(BeNil())
			Expect(serviceDescriptor.Spec.Port).Should(BeNil())
			Expect(serviceDescriptor.Spec.Auth).Should(BeNil())
		})

		It("component vars - pod FQDNs with role", func() {
			comp.Spec.ServiceRefs = []appsv1.ServiceRef{
				{
					Name: serviceRefDeclaration.Name,
					ClusterServiceSelector: &appsv1.ServiceRefClusterSelector{
						Cluster: etcdCluster,
						PodFQDNs: &appsv1.ServiceRefPodFQDNsSelector{
							Component: etcdComponent,
							Role:      &[]string{"leader"}[0],
						},
					},
				},
			}
			reader := &mockReader{
				cli: testCtx.Cli,
				objs: []client.Object{
					&corev1.Pod{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      fmt.Sprintf("%s-%s-%d", etcdCluster, etcdComponent, 0),
							Labels: map[string]string{
								constant.AppManagedByLabelKey:   constant.AppName,
								constant.AppInstanceLabelKey:    etcdCluster,
								constant.KBAppComponentLabelKey: etcdComponent,
								constant.RoleLabelKey:           "follower",
							},
						},
						Spec: corev1.PodSpec{},
					},
					&corev1.Pod{
						ObjectMeta: metav1.ObjectMeta{
							Namespace: namespace,
							Name:      fmt.Sprintf("%s-%s-%d", etcdCluster, etcdComponent, 1),
							Labels: map[string]string{
								constant.AppManagedByLabelKey:   constant.AppName,
								constant.AppInstanceLabelKey:    etcdCluster,
								constant.KBAppComponentLabelKey: etcdComponent,
								constant.RoleLabelKey:           "leader",
							},
						},
						Spec: corev1.PodSpec{},
					},
				},
			}

			compName := constant.GenerateClusterComponentName(etcdCluster, etcdComponent)
			expectedPodFQDNs := intctrlutil.PodFQDN(namespace, compName, reader.objs[1].GetName())

			err := buildServiceReferencesWithoutResolve(testCtx.Ctx, reader, synthesizedComp, compDef, comp)
			Expect(err).Should(Succeed())

			Expect(synthesizedComp.ServiceReferences).Should(HaveKey(serviceRefDeclaration.Name))
			serviceDescriptor := synthesizedComp.ServiceReferences[serviceRefDeclaration.Name]
			Expect(serviceDescriptor).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.PodFQDNs).Should(Not(BeNil()))
			Expect(serviceDescriptor.Spec.PodFQDNs.Value).Should(Equal(expectedPodFQDNs))
			Expect(serviceDescriptor.Spec.Endpoint).Should(BeNil())
			Expect(serviceDescriptor.Spec.Host).Should(BeNil())
			Expect(serviceDescriptor.Spec.Port).Should(BeNil())
			Expect(serviceDescriptor.Spec.Auth).Should(BeNil())
		})
	})
})
