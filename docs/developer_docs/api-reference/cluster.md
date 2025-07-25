---
title: Cluster API Reference
description: Cluster API Reference
keywords: [cluster, api]
sidebar_position: 1
sidebar_label: Cluster
---
<br />
<p>Packages:</p>
<ul>
<li>
<a href="#apps.kubeblocks.io%2fv1">apps.kubeblocks.io/v1</a>
</li>
<li>
<a href="#apps.kubeblocks.io%2fv1alpha1">apps.kubeblocks.io/v1alpha1</a>
</li>
<li>
<a href="#apps.kubeblocks.io%2fv1beta1">apps.kubeblocks.io/v1beta1</a>
</li>
<li>
<a href="#workloads.kubeblocks.io%2fv1">workloads.kubeblocks.io/v1</a>
</li>
<li>
<a href="#workloads.kubeblocks.io%2fv1alpha1">workloads.kubeblocks.io/v1alpha1</a>
</li>
</ul>
<h2 id="apps.kubeblocks.io/v1">apps.kubeblocks.io/v1</h2>
<div>
</div>
Resource Types:
<ul><li>
<a href="#apps.kubeblocks.io/v1.Cluster">Cluster</a>
</li><li>
<a href="#apps.kubeblocks.io/v1.ClusterDefinition">ClusterDefinition</a>
</li><li>
<a href="#apps.kubeblocks.io/v1.Component">Component</a>
</li><li>
<a href="#apps.kubeblocks.io/v1.ComponentDefinition">ComponentDefinition</a>
</li><li>
<a href="#apps.kubeblocks.io/v1.ComponentVersion">ComponentVersion</a>
</li><li>
<a href="#apps.kubeblocks.io/v1.ServiceDescriptor">ServiceDescriptor</a>
</li><li>
<a href="#apps.kubeblocks.io/v1.ShardingDefinition">ShardingDefinition</a>
</li><li>
<a href="#apps.kubeblocks.io/v1.SidecarDefinition">SidecarDefinition</a>
</li></ul>
<h3 id="apps.kubeblocks.io/v1.Cluster">Cluster
</h3>
<div>
<p>Cluster offers a unified management interface for a wide variety of database and storage systems:</p>
<ul>
<li>Relational databases: MySQL, PostgreSQL, MariaDB</li>
<li>NoSQL databases: Redis, MongoDB</li>
<li>KV stores: ZooKeeper, etcd</li>
<li>Analytics systems: ElasticSearch, OpenSearch, ClickHouse, Doris, StarRocks, Solr</li>
<li>Message queues: Kafka, Pulsar</li>
<li>Distributed SQL: TiDB, OceanBase</li>
<li>Vector databases: Qdrant, Milvus, Weaviate</li>
<li>Object storage: Minio</li>
</ul>
<p>KubeBlocks utilizes an abstraction layer to encapsulate the characteristics of these diverse systems.
A Cluster is composed of multiple Components, each defined by vendors or KubeBlocks Addon developers via ComponentDefinition,
arranged in Directed Acyclic Graph (DAG) topologies.
The topologies, defined in a ClusterDefinition, coordinate reconciliation across Cluster&rsquo;s lifecycle phases:
Creating, Running, Updating, Stopping, Stopped, Deleting.
Lifecycle management ensures that each Component operates in harmony, executing appropriate actions at each lifecycle stage.</p>
<p>For sharded-nothing architecture, the Cluster supports managing multiple shards,
each shard managed by a separate Component, supporting dynamic resharding.</p>
<p>The Cluster object is aimed to maintain the overall integrity and availability of a database cluster,
serves as the central control point, abstracting the complexity of multiple-component management,
and providing a unified interface for cluster-wide operations.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Cluster</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterSpec">
ClusterSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>clusterDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterDefinition to use when creating a Cluster.</p>
<p>This field enables users to create a Cluster based on a specific ClusterDefinition.
Which, in conjunction with the <code>topology</code> field, determine:</p>
<ul>
<li>The Components to be included in the Cluster.</li>
<li>The sequences in which the Components are created, updated, and terminate.</li>
</ul>
<p>This facilitates multiple-components management with predefined ClusterDefinition.</p>
<p>Users with advanced requirements can bypass this general setting and specify more precise control over
the composition of the Cluster by directly referencing specific ComponentDefinitions for each component
within <code>componentSpecs[*].componentDef</code>.</p>
<p>If this field is not provided, each component must be explicitly defined in <code>componentSpecs[*].componentDef</code>.</p>
<p>Note: Once set, this field cannot be modified; it is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>topology</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterTopology to be used when creating the Cluster.</p>
<p>This field defines which set of Components, as outlined in the ClusterDefinition, will be used to
construct the Cluster based on the named topology.
The ClusterDefinition may list multiple topologies under <code>clusterdefinition.spec.topologies[*]</code>,
each tailored to different use cases or environments.</p>
<p>If <code>topology</code> is not specified, the Cluster will use the default topology defined in the ClusterDefinition.</p>
<p>Note: Once set during the Cluster creation, the <code>topology</code> field cannot be modified.
It establishes the initial composition and structure of the Cluster and is intended for one-time configuration.</p>
</td>
</tr>
<tr>
<td>
<code>terminationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TerminationPolicyType">
TerminationPolicyType
</a>
</em>
</td>
<td>
<p>Specifies the behavior when a Cluster is deleted.
It defines how resources, data, and backups associated with a Cluster are managed during termination.
Choose a policy based on the desired level of resource cleanup and data preservation:</p>
<ul>
<li><code>DoNotTerminate</code>: Prevents deletion of the Cluster. This policy ensures that all resources remain intact.</li>
<li><code>Delete</code>: Deletes all runtime resources belong to the Cluster.</li>
<li><code>WipeOut</code>: An aggressive policy that deletes all Cluster resources, including volume snapshots and
backups in external storage.
This results in complete data removal and should be used cautiously, primarily in non-production environments
to avoid irreversible data loss.</li>
</ul>
<p>Warning: Choosing an inappropriate termination policy can result in data loss.
The <code>WipeOut</code> policy is particularly risky in production environments due to its irreversible nature.</p>
</td>
</tr>
<tr>
<td>
<code>componentSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">
[]ClusterComponentSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ClusterComponentSpec objects used to define the individual Components that make up a Cluster.
This field allows for detailed configuration of each Component within the Cluster.</p>
<p>Note: <code>shardings</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>shardings</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterSharding">
[]ClusterSharding
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ClusterSharding objects that manage the sharding topology for Cluster Components.
Each ClusterSharding organizes components into shards, with each shard corresponding to a Component.
Components within a shard are all based on a common ClusterComponentSpec template, ensuring uniform configurations.</p>
<p>This field supports dynamic resharding by facilitating the addition or removal of shards
through the <code>shards</code> field in ClusterSharding.</p>
<p>Note: <code>shardings</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies runtimeClassName for all Pods managed by this Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterService">
[]ClusterService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of additional Services that are exposed by a Cluster.
This field allows Services of selected Components, either from <code>componentSpecs</code> or <code>shardings</code> to be exposed,
alongside Services defined with ComponentService.</p>
<p>Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterBackup">
ClusterBackup
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the backup configuration of the Cluster.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterStatus">
ClusterStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterDefinition">ClusterDefinition
</h3>
<div>
<p>ClusterDefinition defines the topology for databases or storage systems,
offering a variety of topological configurations to meet diverse deployment needs and scenarios.</p>
<p>It includes a list of Components and/or Shardings, each linked to a ComponentDefinition or a ShardingDefinition,
which enhances reusability and reduce redundancy.
For example, widely used components such as etcd and Zookeeper can be defined once and reused across multiple ClusterDefinitions,
simplifying the setup of new systems.</p>
<p>Additionally, ClusterDefinition also specifies the sequence of startup, upgrade, and shutdown between Components and/or Shardings,
ensuring a controlled and predictable management of cluster lifecycles.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ClusterDefinition</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterDefinitionSpec">
ClusterDefinitionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>topologies</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterTopology">
[]ClusterTopology
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Topologies defines all possible topologies within the cluster.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterDefinitionStatus">
ClusterDefinitionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Component">Component
</h3>
<div>
<p>Component is a fundamental building block of a Cluster object.
For example, a Redis Cluster can include Components like &lsquo;redis&rsquo;, &lsquo;sentinel&rsquo;, and potentially a proxy like &lsquo;twemproxy&rsquo;.</p>
<p>The Component object is responsible for managing the lifecycle of all replicas within a Cluster component,
It supports a wide range of operations including provisioning, stopping, restarting, termination, upgrading,
configuration changes, vertical and horizontal scaling, failover, switchover, cross-node migration,
scheduling configuration, exposing Services, managing system accounts, enabling/disabling exporter,
and configuring log collection.</p>
<p>Component is an internal sub-object derived from the user-submitted Cluster object.
It is designed primarily to be used by the KubeBlocks controllers,
users are discouraged from modifying Component objects directly and should use them only for monitoring Component statuses.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Component</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentSpec">
ComponentSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>terminationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TerminationPolicyType">
TerminationPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the behavior when a Component is deleted.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this Component.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRef">
[]ServiceRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of ServiceRef for a Component, enabling access to both external services and
Services provided by other Clusters.</p>
<p>Types of services:</p>
<ul>
<li>External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;
Require a ServiceDescriptor for connection details.</li>
<li>Services provided by a Cluster: Managed by the same KubeBlocks operator;
identified using Cluster, Component and Service names.</li>
</ul>
<p>ServiceRefs with identical <code>serviceRef.name</code> in the same Cluster are considered the same.</p>
<p>Example:</p>
<pre><code class="language-yaml">serviceRefs:
  - name: &quot;redis-sentinel&quot;
    serviceDescriptor:
      name: &quot;external-redis-sentinel&quot;
  - name: &quot;postgres-cluster&quot;
    clusterServiceSelector:
      cluster: &quot;my-postgres-cluster&quot;
      service:
        component: &quot;postgresql&quot;
</code></pre>
<p>The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Labels to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Annotations to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of environment variables to add.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources required by the Component.
It allows defining the CPU, memory requirements and limits for the Component&rsquo;s containers.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimTemplate">
[]PersistentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for the Component.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>persistentVolumeClaimRetentionPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicy">
PersistentVolumeClaimRetentionPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>persistentVolumeClaimRetentionPolicy describes the lifecycle of persistent
volume claims created from volumeClaimTemplates. By default, all persistent
volume claims are created as needed and retained until manually deleted. This
policy allows the lifecycle to be altered, for example by deleting persistent
volume claims when their workload is deleted, or when their pod is scaled
down.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes to override.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentNetwork">
ComponentNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the network configuration for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides Services defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentSystemAccount">
[]ComponentSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides system accounts defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentConfig">
[]ClusterComponentConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration content of a config template.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceAccount required by the running Component.
This ServiceAccount is used to grant necessary permissions for the Component&rsquo;s Pods to interact
with other Kubernetes resources, such as modifying Pod labels or sending events.</p>
<p>If not specified, KubeBlocks automatically creates a default ServiceAccount named
&ldquo;kb-&#123;componentdefinition.name&#125;&rdquo;, bound to a role with rules defined in ComponentDefinition&rsquo;s
<code>policyRules</code> field. If needed (currently this means if any lifecycleAction is enabled),
it will also be bound to a default role named
&ldquo;kubeblocks-cluster-pod-role&rdquo;, which is installed together with KubeBlocks.
If multiple components use the same ComponentDefinition, they will share one ServiceAccount.</p>
<p>If the field is not empty, the specified ServiceAccount will be used, and KubeBlocks will not
create a ServiceAccount. But KubeBlocks does create RoleBindings for the specified ServiceAccount.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides fine-grained control over the spec update process of all instances.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>tlsConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TLSConfig">
TLSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the TLS configuration for the Component, including:</p>
<ul>
<li>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication.</li>
<li>An optional field that specifies the configuration for the TLS certificates issuer when TLS is enabled.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows for the customization of configuration values for each instance within a Component.
An Instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).
While instances typically share a common configuration as defined in the ClusterComponentSpec,
they can require unique settings in various scenarios:</p>
<p>For example:
- A database Component might require different resource allocations for primary and secondary instances,
  with primaries needing more resources.
- During a rolling upgrade, a Component may first update the image for one or a few instances,
and then update the remaining instances after verifying that the updated instances are functioning correctly.</p>
<p>InstanceTemplate allows for specifying these unique configurations per instance.
Each instance&rsquo;s name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),
starting with an ordinal of 0.
It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the Component.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>flatInstanceOrdinal</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>flatInstanceOrdinal controls whether the naming of instances(pods) under this component uses a flattened,
globally uniquely ordinal scheme, regardless of the instance template.</p>
<p>Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the Cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines runtimeClassName for all Pods managed by this Component.</p>
</td>
</tr>
<tr>
<td>
<code>disableExporter</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines whether metrics exporter information is annotated on the Component&rsquo;s headless Service.</p>
<p>If set to true, the following annotations will not be patched into the Service:</p>
<ul>
<li>&ldquo;monitor.kubeblocks.io/path&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/port&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/scheme&rdquo;</li>
</ul>
<p>These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.</p>
</td>
</tr>
<tr>
<td>
<code>stop</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Stop the Component.
If set, all the computing resources will be released.</p>
</td>
</tr>
<tr>
<td>
<code>sidecars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Sidecar">
[]Sidecar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the sidecars to be injected into the Component.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentStatus">
ComponentStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentDefinition">ComponentDefinition
</h3>
<div>
<p>ComponentDefinition serves as a reusable blueprint for creating Components,
encapsulating essential static settings such as Component description,
Pod templates, configuration file templates, scripts, parameter lists,
injected environment variables and their sources, and event handlers.
ComponentDefinition works in conjunction with dynamic settings from the ClusterComponentSpec,
to instantiate Components during Cluster creation.</p>
<p>Key aspects that can be defined in a ComponentDefinition include:</p>
<ul>
<li>PodSpec template: Specifies the PodSpec template used by the Component.</li>
<li>Configuration templates: Specify the configuration file templates required by the Component.</li>
<li>Scripts: Provide the necessary scripts for Component management and operations.</li>
<li>Storage volumes: Specify the storage volumes and their configurations for the Component.</li>
<li>Pod roles: Outlines various roles of Pods within the Component along with their capabilities.</li>
<li>Exposed Kubernetes Services: Specify the Services that need to be exposed by the Component.</li>
<li>System accounts: Define the system accounts required for the Component.</li>
<li>Monitoring and logging: Configure the exporter and logging settings for the Component.</li>
</ul>
<p>ComponentDefinitions also enable defining reactive behaviors of the Component in response to events,
such as member join/leave, Component addition/deletion, role changes, switch over, and more.
This allows for automatic event handling, thus encapsulating complex behaviors within the Component.</p>
<p>Referencing a ComponentDefinition when creating individual Components ensures inheritance of predefined configurations,
promoting reusability and consistency across different deployments and cluster topologies.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ComponentDefinition</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">
ComponentDefinitionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>provider</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the Component provider, typically the vendor or developer name.
It identifies the entity responsible for creating and maintaining the Component.</p>
<p>When specifying the provider name, consider the following guidelines:</p>
<ul>
<li>Keep the name concise and relevant to the Component.</li>
<li>Use a consistent naming convention across Components from the same provider.</li>
<li>Avoid using trademarked or copyrighted names without proper permission.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a brief and concise explanation of the Component&rsquo;s purpose, functionality, and any relevant details.
It serves as a quick reference for users to understand the Component&rsquo;s role and characteristics.</p>
</td>
</tr>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the type of well-known service protocol that the Component provides.
It specifies the standard or widely recognized protocol used by the Component to offer its Services.</p>
<p>The <code>serviceKind</code> field allows users to quickly identify the type of Service provided by the Component
based on common protocols or service types. This information helps in understanding the compatibility,
interoperability, and usage of the Component within a system.</p>
<p>Some examples of well-known service protocols include:</p>
<ul>
<li>&ldquo;MySQL&rdquo;: Indicates that the Component provides a MySQL database service.</li>
<li>&ldquo;PostgreSQL&rdquo;: Indicates that the Component offers a PostgreSQL database service.</li>
<li>&ldquo;Redis&rdquo;: Signifies that the Component functions as a Redis key-value store.</li>
<li>&ldquo;ETCD&rdquo;: Denotes that the Component serves as an ETCD distributed key-value store.</li>
</ul>
<p>The <code>serviceKind</code> value is case-insensitive, allowing for flexibility in specifying the protocol name.</p>
<p>When specifying the <code>serviceKind</code>, consider the following guidelines:</p>
<ul>
<li>Use well-established and widely recognized protocol names or service types.</li>
<li>Ensure that the <code>serviceKind</code> accurately represents the primary service type offered by the Component.</li>
<li>If the Component provides multiple services, choose the most prominent or commonly used protocol.</li>
<li>Limit the <code>serviceKind</code> to a maximum of 32 characters for conciseness and readability.</li>
</ul>
<p>Note: The <code>serviceKind</code> field is optional and can be left empty if the Component does not fit into a well-known
service category or if the protocol is not widely recognized. It is primarily used to convey information about
the Component&rsquo;s service type to users and facilitate discovery and integration.</p>
<p>The <code>serviceKind</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the version of the Service provided by the Component.
It follows the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
<p>The Semantic Versioning specification defines a version number format of X.Y.Z (MAJOR.MINOR.PATCH), where:</p>
<ul>
<li>X represents the major version and indicates incompatible API changes.</li>
<li>Y represents the minor version and indicates added functionality in a backward-compatible manner.</li>
<li>Z represents the patch version and indicates backward-compatible bug fixes.</li>
</ul>
<p>Additional labels for pre-release and build metadata are available as extensions to the X.Y.Z format:</p>
<ul>
<li>Use pre-release labels (e.g., -alpha, -beta) for versions that are not yet stable or ready for production use.</li>
<li>Use build metadata (e.g., +build.1) for additional version information if needed.</li>
</ul>
<p>Examples of valid ServiceVersion values:</p>
<ul>
<li>&ldquo;1.0.0&rdquo;</li>
<li>&ldquo;2.3.1&rdquo;</li>
<li>&ldquo;3.0.0-alpha.1&rdquo;</li>
<li>&ldquo;4.5.2+build.1&rdquo;</li>
</ul>
<p>The <code>serviceVersion</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static labels that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If a label key in the <code>labels</code> field conflicts with any system labels or user-specified labels,
it will be silently ignored to avoid overriding higher-priority labels.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static annotations that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If an annotation key in the <code>annotations</code> field conflicts with any system annotations
or user-specified annotations, it will be silently ignored to avoid overriding higher-priority annotations.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>runtime</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<p>Specifies the PodSpec template used in the Component.
It includes the following elements:</p>
<ul>
<li>Init containers</li>
<li>Containers
<ul>
<li>Image</li>
<li>Commands</li>
<li>Args</li>
<li>Envs</li>
<li>Mounts</li>
<li>Ports</li>
<li>Security context</li>
<li>Probes</li>
<li>Lifecycle</li>
</ul></li>
<li>Volumes</li>
</ul>
<p>This field is intended to define static settings that remain consistent across all instantiated Components.
Dynamic settings such as CPU and memory resource limits, as well as scheduling settings (affinity,
toleration, priority), may vary among different instantiated Components.
They should be specified in the <code>cluster.spec.componentSpecs</code> (ClusterComponentSpec).</p>
<p>Specific instances of a Component may override settings defined here, such as using a different container image
or modifying environment variable values.
These instance-specific overrides can be specified in <code>cluster.spec.componentSpecs[*].instances</code>.</p>
<p>This field is immutable and cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.EnvVar">
[]EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines variables which are determined after Cluster instantiation and reflect
dynamic or runtime attributes of instantiated Clusters.
These variables serve as placeholders for setting environment variables in Pods and Actions,
or for rendering configuration and script templates before actual values are finalized.</p>
<p>These variables are placed in front of the environment variables declared in the Pod if used as
environment variables.</p>
<p>Variable values can be sourced from:</p>
<ul>
<li>ConfigMap: Select and extract a value from a specific key within a ConfigMap.</li>
<li>Secret: Select and extract a value from a specific key within a Secret.</li>
<li>HostNetwork: Retrieves values (including ports) from host-network resources.</li>
<li>Service: Retrieves values (including address, port, NodePort) from a selected Service.
Intended to obtain the address of a ComponentService within the same Cluster.</li>
<li>Credential: Retrieves account name and password from a SystemAccount variable.</li>
<li>ServiceRef: Retrieves address, port, account name and password from a selected ServiceRefDeclaration.
Designed to obtain the address bound to a ServiceRef, such as a ClusterService or
ComponentService of another cluster or an external service.</li>
<li>Component: Retrieves values from a selected Component, including replicas and instance name list.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVolume">
[]ComponentVolume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the volumes used by the Component and some static attributes of the volumes.
After defining the volumes here, user can reference them in the
<code>cluster.spec.componentSpecs[*].volumeClaimTemplates</code> field to configure dynamic properties such as
volume capacity and storage class.</p>
<p>This field allows you to specify the following:</p>
<ul>
<li>Snapshot behavior: Determines whether a snapshot of the volume should be taken when performing
a snapshot backup of the Component.</li>
<li>Disk high watermark: Sets the high watermark for the volume&rsquo;s disk usage.
When the disk usage reaches the specified threshold, it triggers an alert or action.</li>
</ul>
<p>By configuring these volume behaviors, you can control how the volumes are managed and monitored within the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.HostNetwork">
HostNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the host network configuration for the Component.</p>
<p>When <code>hostNetwork</code> option is enabled, the Pods share the host&rsquo;s network namespace and can directly access
the host&rsquo;s network interfaces.
This means that if multiple Pods need to use the same port, they cannot run on the same host simultaneously
due to port conflicts.</p>
<p>The DNSPolicy field in the Pod spec determines how containers within the Pod perform DNS resolution.
When using hostNetwork, the operator will set the DNSPolicy to &lsquo;ClusterFirstWithHostNet&rsquo;.
With this policy, DNS queries will first go through the K8s cluster&rsquo;s DNS service.
If the query fails, it will fall back to the host&rsquo;s DNS settings.</p>
<p>If set, the DNS policy will be automatically set to &ldquo;ClusterFirstWithHostNet&rdquo;.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines additional Services to expose the Component&rsquo;s endpoints.</p>
<p>A default headless Service, named <code>&#123;cluster.name&#125;-&#123;component.name&#125;-headless</code>, is automatically created
for internal Cluster communication.</p>
<p>This field enables customization of additional Services to expose the Component&rsquo;s endpoints to
other Components within the same or different Clusters, and to external applications.
Each Service entry in this list can include properties such as ports, type, and selectors.</p>
<ul>
<li>For intra-Cluster access, Components can reference Services using variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceVarRef</code>.</li>
<li>For inter-Cluster access, reference Services use variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceRefVarRef</code>,
and bind Services at Cluster creation time with <code>clusterComponentSpec.ServiceRef[*].clusterServiceSelector</code>.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>disableDefaultHeadlessService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to create the default headless service.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the config file templates and volume mount parameters used by the Component.</p>
<p>This field specifies a list of templates that will be rendered into Component containers&rsquo; config files.
Each template is represented as a ConfigMap and may contain multiple config files, with each file being a key in the ConfigMap.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>scripts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies groups of scripts, each provided via a ConfigMap, to be mounted as volumes in the container.
These scripts can be executed during container startup or via specific actions.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>logConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.LogConfig">
[]LogConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the types of logs generated by instances of the Component and their corresponding file paths.
These logs can be collected for further analysis and monitoring.</p>
<p>The <code>logConfigs</code> field is an optional list of LogConfig objects, where each object represents
a specific log type and its configuration.
It allows you to specify multiple log types and their respective file paths for the Component.</p>
<p>Examples:</p>
<pre><code class="language-yaml"> logConfigs:
 - filePathPattern: /data/mysql/log/mysqld-error.log
   name: error
 - filePathPattern: /data/mysql/log/mysqld.log
   name: general
 - filePathPattern: /data/mysql/log/mysqld-slowquery.log
   name: slow
</code></pre>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SystemAccount">
[]SystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An array of <code>SystemAccount</code> objects that define the system accounts needed
for the management operations of the Component.</p>
<p>Each <code>SystemAccount</code> includes:</p>
<ul>
<li>Account name.</li>
<li>The SQL statement template: Used to create the system account.</li>
<li>Password Source: Either generated based on certain rules or retrieved from a Secret.</li>
</ul>
<p>Use cases for system accounts typically involve tasks like system initialization, backups, monitoring,
health checks, replication, and other system-level operations.</p>
<p>System accounts are distinct from user accounts, although both are database accounts.</p>
<ul>
<li><strong>System Accounts</strong>: Created during Cluster setup by the KubeBlocks operator,
these accounts have higher privileges for system management and are fully managed
through a declarative API by the operator.</li>
<li><strong>User Accounts</strong>: Managed by users or administrator.
User account permissions should follow the principle of least privilege,
granting only the necessary access rights to complete their required tasks.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TLS">
TLS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the TLS configuration for the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>replicasLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ReplicasLimit">
ReplicasLimit
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the upper limit of the number of replicas supported by the Component.</p>
<p>It defines the maximum number of replicas that can be created for the Component.
This field allows you to set a limit on the scalability of the Component, preventing it from exceeding a certain number of replicas.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>available</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailable">
ComponentAvailable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategies for determining the available status of the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enumerate all possible roles assigned to each replica of the Component, influencing its behavior.</p>
<p>A replica can have zero or one role.
KubeBlocks operator determines the role of each replica by invoking the <code>lifecycleActions.roleProbe</code> method.
This action returns the role for each replica, and the returned role must be predefined here.</p>
<p>The roles assigned to a replica can influence various aspects of the Component&rsquo;s behavior, such as:</p>
<ul>
<li>Service selection: The Component&rsquo;s exposed Services may target replicas based on their roles using <code>roleSelector</code>.</li>
<li>Update order: The roles can determine the order in which replicas are updated during a Component update.
For instance, replicas with a &ldquo;follower&rdquo; role can be updated first, while the replica with the &ldquo;leader&rdquo;
role is updated last. This helps minimize the number of leader changes during the update process.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p><code>minReadySeconds</code> is the minimum duration in seconds that a new Pod should remain in the ready
state without any of its containers crashing to be considered available.
This ensures the Pod&rsquo;s stability and readiness to serve requests.</p>
<p>A default value of 0 seconds means the Pod is considered available as soon as it enters the ready state.</p>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the concurrency strategy for updating multiple instances of the Component.
Available strategies:</p>
<ul>
<li><code>Serial</code>: Updates replicas one at a time, ensuring minimal downtime by waiting for each replica to become ready
before updating the next.</li>
<li><code>Parallel</code>: Updates all replicas simultaneously, optimizing for speed but potentially reducing availability
during the update.</li>
<li><code>BestEffortParallel</code>: Updates replicas concurrently with a limit on simultaneous updates to ensure a minimum
number of operational replicas for maintaining quorum.
 For example, in a 5-replica component, updating a maximum of 2 replicas simultaneously keeps
at least 3 operational for quorum.</li>
</ul>
<p>This field is immutable and defaults to &lsquo;Serial&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InstanceSet controls the creation of pods during initial scale up, replacement of pods on nodes, and scaling down.</p>
<ul>
<li><code>OrderedReady</code>: Creates pods in increasing order (pod-0, then pod-1, etc). The controller waits until each pod
is ready before continuing. Pods are removed in reverse order when scaling down.</li>
<li><code>Parallel</code>: Creates pods in parallel to match the desired scale without waiting. All pods are deleted at once
when scaling down.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>policyRules</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#policyrule-v1-rbac">
[]Kubernetes rbac/v1.PolicyRule
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the namespaced policy rules required by the Component.</p>
<p>The <code>policyRules</code> field is an array of <code>rbacv1.PolicyRule</code> objects that define the policy rules
needed by the Component to operate within a namespace.
These policy rules determine the permissions and verbs the Component is allowed to perform on
Kubernetes resources within the namespace.</p>
<p>The purpose of this field is to automatically generate the necessary RBAC roles
for the Component based on the specified policy rules.
This ensures that the Pods in the Component has appropriate permissions to function.</p>
<p>To prevent privilege escalation, only permissions already owned by KubeBlocks can be added here.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>lifecycleActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentLifecycleActions">
ComponentLifecycleActions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of hooks and procedures that customize the behavior of a Component throughout its lifecycle.
Actions are triggered at specific lifecycle stages:</p>
<ul>
<li><code>postProvision</code>: Defines the hook to be executed after the creation of a Component,
with <code>preCondition</code> specifying when the action should be fired relative to the Component&rsquo;s lifecycle stages:
<code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.</li>
<li><code>preTerminate</code>: Defines the hook to be executed before terminating a Component.</li>
<li><code>roleProbe</code>: Defines the procedure which is invoked regularly to assess the role of replicas.</li>
<li><code>switchover</code>: Defines the procedure for a controlled transition of a role to a new replica.
This approach aims to minimize downtime and maintain availability in systems with a leader-follower topology,
such as before planned maintenance or upgrades on the current leader node.</li>
<li><code>memberJoin</code>: Defines the procedure to add a new replica to the replication group.</li>
<li><code>memberLeave</code>: Defines the method to remove a replica from the replication group.</li>
<li><code>readOnly</code>: Defines the procedure to switch a replica into the read-only state.</li>
<li><code>readWrite</code>: transition a replica from the read-only state back to the read-write state.</li>
<li><code>dataDump</code>: Defines the procedure to export the data from a replica.</li>
<li><code>dataLoad</code>: Defines the procedure to import data into a replica.</li>
<li><code>reconfigure</code>: Defines the procedure that update a replica with new configuration file.</li>
<li><code>accountProvision</code>: Defines the procedure to generate a new database account.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefDeclarations</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefDeclaration">
[]ServiceRefDeclaration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists external service dependencies of the Component, including services from other Clusters or outside the K8s environment.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>exporter</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Exporter">
Exporter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the built-in metrics exporter container.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentDefinitionStatus">
ComponentDefinitionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVersion">ComponentVersion
</h3>
<div>
<p>ComponentVersion is the Schema for the componentversions API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ComponentVersion</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVersionSpec">
ComponentVersionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>compatibilityRules</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVersionCompatibilityRule">
[]ComponentVersionCompatibilityRule
</a>
</em>
</td>
<td>
<p>CompatibilityRules defines compatibility rules between sets of component definitions and releases.</p>
</td>
</tr>
<tr>
<td>
<code>releases</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVersionRelease">
[]ComponentVersionRelease
</a>
</em>
</td>
<td>
<p>Releases represents different releases of component instances within this ComponentVersion.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVersionStatus">
ComponentVersionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceDescriptor">ServiceDescriptor
</h3>
<div>
<p>ServiceDescriptor describes a service provided by external sources.
It contains the necessary details such as the service&rsquo;s address and connection credentials.
To enable a Cluster to access this service, the ServiceDescriptor&rsquo;s name should be specified
in the Cluster configuration under <code>clusterComponent.serviceRefs[*].serviceDescriptor</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ServiceDescriptor</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceDescriptorSpec">
ServiceDescriptorSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the type of database service provided by the external service.
For example, &ldquo;mysql&rdquo;, &ldquo;redis&rdquo;, &ldquo;mongodb&rdquo;.
This field categorizes databases by their functionality, protocol and compatibility, facilitating appropriate
service integration based on their unique capabilities.</p>
<p>This field is case-insensitive.</p>
<p>It also supports abbreviations for some well-known databases:
- &ldquo;pg&rdquo;, &ldquo;pgsql&rdquo;, &ldquo;postgres&rdquo;, &ldquo;postgresql&rdquo;: PostgreSQL service
- &ldquo;zk&rdquo;, &ldquo;zookeeper&rdquo;: ZooKeeper service
- &ldquo;es&rdquo;, &ldquo;elasticsearch&rdquo;: Elasticsearch service
- &ldquo;mongo&rdquo;, &ldquo;mongodb&rdquo;: MongoDB service
- &ldquo;ch&rdquo;, &ldquo;clickhouse&rdquo;: ClickHouse service</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the version of the service provided by the external service.
This is crucial for ensuring compatibility between different components of the system,
as different versions of a service may have varying features.</p>
</td>
</tr>
<tr>
<td>
<code>endpoint</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the endpoint of the external service.</p>
<p>If the service is exposed via a cluster, the endpoint will be provided in the format of <code>host:port</code>.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the service or IP address of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the port of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>podFQDNs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the pod FQDNs of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ConnectionCredentialAuth">
ConnectionCredentialAuth
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the authentication credentials required for accessing an external service.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceDescriptorStatus">
ServiceDescriptorStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardingDefinition">ShardingDefinition
</h3>
<div>
<p>ShardingDefinition is the Schema for the shardingdefinitions API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ShardingDefinition</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingDefinitionSpec">
ShardingDefinitionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingTemplate">
ShardingTemplate
</a>
</em>
</td>
<td>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>shardsLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardsLimit">
ShardsLimit
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the upper limit of the number of shards supported by the sharding.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>provisionStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategy for provisioning shards of the sharding. Only <code>Serial</code> and <code>Parallel</code> are supported.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategy for updating shards of the sharding. Only <code>Serial</code> and <code>Parallel</code> are supported.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>lifecycleActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingLifecycleActions">
ShardingLifecycleActions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of hooks and procedures that customize the behavior of a sharding throughout its lifecycle.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingSystemAccount">
[]ShardingSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the system accounts for the sharding.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingTLS">
ShardingTLS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the TLS for the sharding.</p>
<p>This field is immutable.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingDefinitionStatus">
ShardingDefinitionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.SidecarDefinition">SidecarDefinition
</h3>
<div>
<p>SidecarDefinition is the Schema for the sidecardefinitions API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>SidecarDefinition</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SidecarDefinitionSpec">
SidecarDefinitionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the sidecar.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the component definition that the sidecar belongs to.</p>
<p>For a specific cluster object, if there is any components provided by the component definition of @owner,
the sidecar will be created and injected into the components which are provided by
the component definition of @selectors automatically.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>selectors</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Specifies the component definition of components that the sidecar along with.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>containers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<p>List of containers for the sidecar.</p>
<p>Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.EnvVar">
[]EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines variables which are needed by the sidecar.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration file templates used by the Sidecar.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>scripts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scripts used by the Sidecar.</p>
<p>This field is immutable.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SidecarDefinitionStatus">
SidecarDefinitionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Action">Action
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentConfig">ClusterComponentConfig</a>, <a href="#apps.kubeblocks.io/v1.ComponentLifecycleActions">ComponentLifecycleActions</a>, <a href="#apps.kubeblocks.io/v1.Probe">Probe</a>, <a href="#apps.kubeblocks.io/v1.ShardingLifecycleActions">ShardingLifecycleActions</a>, <a href="#apps.kubeblocks.io/v1alpha1.RolloutPromoteCondition">RolloutPromoteCondition</a>, <a href="#workloads.kubeblocks.io/v1.ConfigTemplate">ConfigTemplate</a>, <a href="#workloads.kubeblocks.io/v1.MembershipReconfiguration">MembershipReconfiguration</a>)
</p>
<div>
<p>Action defines a customizable hook or procedure tailored for different database engines,
designed to be invoked at predetermined points within the lifecycle of a Component instance.
It provides a modular and extensible way to customize a Component&rsquo;s behavior through the execution of defined actions.</p>
<p>Available Action triggers include:</p>
<ul>
<li><code>postProvision</code>: Defines the hook to be executed after the creation of a Component,
with <code>preCondition</code> specifying when the action should be fired relative to the Component&rsquo;s lifecycle stages:
<code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.</li>
<li><code>preTerminate</code>: Defines the hook to be executed before terminating a Component.</li>
<li><code>roleProbe</code>: Defines the procedure which is invoked regularly to assess the role of replicas.</li>
<li><code>switchover</code>: Defines the procedure for a controlled transition of a role to a new replica.</li>
<li><code>memberJoin</code>: Defines the procedure to add a new replica to the replication group.</li>
<li><code>memberLeave</code>: Defines the method to remove a replica from the replication group.</li>
<li><code>readOnly</code>: Defines the procedure to switch a replica into the read-only state.</li>
<li><code>readWrite</code>: Defines the procedure to transition a replica from the read-only state back to the read-write state.</li>
<li><code>dataDump</code>: Defines the procedure to export the data from a replica.</li>
<li><code>dataLoad</code>: Defines the procedure to import data into a replica.</li>
<li><code>reconfigure</code>: Defines the procedure that update a replica with new configuration.</li>
<li><code>accountProvision</code>: Defines the procedure to generate a new database account.</li>
</ul>
<p>Actions can be executed in different ways:</p>
<ul>
<li>ExecAction: Executes a command inside a container.
A set of predefined environment variables are available and can be leveraged within the <code>exec.command</code>
to access context information such as details about pods, components, the overall cluster state,
or database connection credentials.
These variables provide a dynamic and context-aware mechanism for script execution.</li>
<li>HTTPAction: Performs an HTTP request.
HTTPAction is to be implemented in future version.</li>
<li>GRPCAction: In future version, Actions will support initiating gRPC calls.
This allows developers to implement Actions using plugins written in programming language like Go,
providing greater flexibility and extensibility.</li>
</ul>
<p>An action is considered successful on returning 0, or HTTP 200 for status HTTP(s) Actions.
Any other return value or HTTP status codes indicate failure,
and the action may be retried based on the configured retry policy.</p>
<ul>
<li>If an action exceeds the specified timeout duration, it will be terminated, and the action is considered failed.</li>
<li>If an action produces any data as output, it should be written to stdout,
or included in the HTTP response payload for HTTP(s) actions.</li>
<li>If an action encounters any errors, error messages should be written to stderr,
or detailed in the HTTP response with the appropriate non-200 status code.</li>
</ul>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>exec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ExecAction">
ExecAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the command to run.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the maximum duration in seconds that the Action is allowed to run.</p>
<p>If the Action does not complete within this time frame, it will be terminated.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>retryPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.RetryPolicy">
RetryPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the strategy to be taken when retrying the Action after a failure.</p>
<p>It specifies the conditions under which the Action should be retried and the limits to apply,
such as the maximum number of retries and backoff strategy.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>preCondition</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PreConditionType">
PreConditionType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the state that the cluster must reach before the Action is executed.
Currently, this is only applicable to the <code>postProvision</code> action.</p>
<p>The conditions are as follows:</p>
<ul>
<li><code>Immediately</code>: Executed right after the Component object is created.
The readiness of the Component and its resources is not guaranteed at this stage.</li>
<li><code>RuntimeReady</code>: The Action is triggered after the Component object has been created and all associated
runtime resources (e.g. Pods) are in a ready state.</li>
<li><code>ComponentReady</code>: The Action is triggered after the Component itself is in a ready state.
This process does not affect the readiness state of the Component or the Cluster.</li>
<li><code>ClusterReady</code>: The Action is executed after the Cluster is in a ready state.
This execution does not alter the Component or the Cluster&rsquo;s state of readiness.</li>
</ul>
<p>This field cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ActionAssertion">ActionAssertion
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentAvailableProbeAssertion">ComponentAvailableProbeAssertion</a>)
</p>
<div>
<p>ActionAssertion defines the custom assertions for evaluating the success or failure of an action.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>succeed</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Whether the action should succeed or fail.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>stdout</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ActionOutputMatcher">
ActionOutputMatcher
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the stdout matcher for the action.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>stderr</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ActionOutputMatcher">
ActionOutputMatcher
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the stderr matcher for the action.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ActionOutputMatcher">ActionOutputMatcher
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ActionAssertion">ActionAssertion</a>)
</p>
<div>
<p>ActionOutputMatcher defines the matcher for the output of an action.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>equalTo</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The output of the action should be equal to the specified value.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>contains</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The output of the action should contain the specified value.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterBackup">ClusterBackup
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether automated backup is enabled for the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPeriod</code><br/>
<em>
github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1.RetentionPeriod
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines the duration to retain backups. Backups older than this period are automatically removed.</p>
<p>For example, RetentionPeriod of <code>30d</code> will keep only the backups of last 30 days.
Sample duration format:</p>
<ul>
<li>years: 	2y</li>
<li>months: 	6mo</li>
<li>days: 		30d</li>
<li>hours: 	12h</li>
<li>minutes: 	30m</li>
</ul>
<p>You can also combine the above durations. For example: 30d12h30m.
Default value is 7d.</p>
</td>
</tr>
<tr>
<td>
<code>method</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the backup method to use, as defined in backupPolicy.</p>
</td>
</tr>
<tr>
<td>
<code>cronExpression</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The cron expression for the schedule. The timezone is in UTC. See <a href="https://en.wikipedia.org/wiki/Cron">https://en.wikipedia.org/wiki/Cron</a>.</p>
</td>
</tr>
<tr>
<td>
<code>startingDeadlineMinutes</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the maximum time in minutes that the system will wait to start a missed backup job.
If the scheduled backup time is missed for any reason, the backup job must start within this deadline.
Values must be between 0 (immediate execution) and 1440 (one day).</p>
</td>
</tr>
<tr>
<td>
<code>repoName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the backupRepo. If not set, the default backupRepo will be used.</p>
</td>
</tr>
<tr>
<td>
<code>pitrEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to enable point-in-time recovery.</p>
</td>
</tr>
<tr>
<td>
<code>continuousMethod</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the backup method to use, if not set, use the first continuous method.</p>
</td>
</tr>
<tr>
<td>
<code>incrementalBackupEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to enable incremental backup.</p>
</td>
</tr>
<tr>
<td>
<code>incrementalCronExpression</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The cron expression for the incremental backup schedule. The timezone is in UTC. See <a href="https://en.wikipedia.org/wiki/Cron">https://en.wikipedia.org/wiki/Cron</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterComponentConfig">ClusterComponentConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
<p>ClusterComponentConfig represents a configuration for a component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the config.</p>
</td>
</tr>
<tr>
<td>
<code>variables</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Variables are key-value pairs for dynamic configuration values that can be provided by the user.</p>
</td>
</tr>
<tr>
<td>
<code>ClusterComponentConfigSource</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentConfigSource">
ClusterComponentConfigSource
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterComponentConfigSource</code> are embedded into this type.)
</p>
<p>The external source for the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>reconfigure</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The custom reconfigure action to reload the service configuration whenever changes to this config are detected.</p>
<p>The container executing this action has access to following variables:</p>
<ul>
<li>KB_CONFIG_FILES_CREATED: file1,file2&hellip;</li>
<li>KB_CONFIG_FILES_REMOVED: file1,file2&hellip;</li>
<li>KB_CONFIG_FILES_UPDATED: file1:checksum1,file2:checksum2&hellip;</li>
</ul>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>externalManaged</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>ExternalManaged indicates whether the configuration is managed by an external system.
When set to true, the controller will use the user-provided template and reconfigure action,
ignoring the default template and update behavior.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterComponentConfigSource">ClusterComponentConfigSource
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentConfig">ClusterComponentConfig</a>)
</p>
<div>
<p>ClusterComponentConfigSource represents the source of a configuration for a component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>configMap</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#configmapvolumesource-v1-core">
Kubernetes core/v1.ConfigMapVolumeSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigMap source for the config.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterComponentService">ClusterComponentService
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>References the ComponentService name defined in the <code>componentDefinition.spec.services[*].name</code>.</p>
</td>
</tr>
<tr>
<td>
<code>serviceType</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#servicetype-v1-core">
Kubernetes core/v1.ServiceType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines how the Service is exposed. Valid options are <code>ClusterIP</code>, <code>NodePort</code>, and <code>LoadBalancer</code>.</p>
<ul>
<li><code>ClusterIP</code> allocates a Cluster-internal IP address for load-balancing to endpoints.
Endpoints are determined by the selector or if that is not specified,
they are determined by manual construction of an Endpoints object or EndpointSlice objects.</li>
<li><code>NodePort</code> builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the ClusterIP.</li>
<li><code>LoadBalancer</code> builds on NodePort and creates an external load-balancer (if supported in the current cloud)
which routes to the same endpoints as the ClusterIP.</li>
</ul>
<p>Note: although K8s Service type allows the &lsquo;ExternalName&rsquo; type, it is not a valid option for ClusterComponentService.</p>
<p>For more info, see:
<a href="https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types">https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types</a>.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If ServiceType is LoadBalancer, cloud provider related parameters can be put here.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer">https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer</a>.</p>
</td>
</tr>
<tr>
<td>
<code>podService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to generate individual Services for each Pod.
If set to true, a separate Service will be created for each Pod in the Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterSharding">ClusterSharding</a>, <a href="#apps.kubeblocks.io/v1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ClusterComponentSpec defines the specification of a Component within a Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the Component&rsquo;s name.
It&rsquo;s part of the Service DNS name and must comply with the IANA service naming rule.
The name is optional when ClusterComponentSpec is used as a template (e.g., in <code>clusterSharding</code>),
but required otherwise.</p>
</td>
</tr>
<tr>
<td>
<code>componentDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the ComponentDefinition custom resource (CR) that defines the Component&rsquo;s characteristics and behavior.</p>
<p>Supports three different ways to specify the ComponentDefinition:</p>
<ul>
<li>the regular expression - recommended</li>
<li>the full name - recommended</li>
<li>the name prefix</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this Component.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).
If no version is specified, the latest available version will be used.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRef">
[]ServiceRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of ServiceRef for a Component, enabling access to both external services and
Services provided by other Clusters.</p>
<p>Types of services:</p>
<ul>
<li>External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;
Require a ServiceDescriptor for connection details.</li>
<li>Services provided by a Cluster: Managed by the same KubeBlocks operator;
identified using Cluster, Component and Service names.</li>
</ul>
<p>ServiceRefs with identical <code>serviceRef.name</code> in the same Cluster are considered the same.</p>
<p>Example:</p>
<pre><code class="language-yaml">serviceRefs:
  - name: &quot;redis-sentinel&quot;
    serviceDescriptor:
      name: &quot;external-redis-sentinel&quot;
  - name: &quot;postgres-cluster&quot;
    clusterServiceSelector:
      cluster: &quot;my-postgres-cluster&quot;
      service:
        component: &quot;postgresql&quot;
</code></pre>
<p>The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Labels to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Annotations to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of environment variables to add.
These environment variables will be placed after the environment variables declared in the Pod.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.
If defined, it will overwrite the scheduling policy defined in ClusterSpec.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources required by the Component.
It allows defining the CPU, memory requirements and limits for the Component&rsquo;s containers.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimTemplate">
[]PersistentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>persistentVolumeClaimRetentionPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicy">
PersistentVolumeClaimRetentionPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>persistentVolumeClaimRetentionPolicy describes the lifecycle of persistent
volume claims created from volumeClaimTemplates. By default, all persistent
volume claims are created as needed and retained until manually deleted. This
policy allows the lifecycle to be altered, for example by deleting persistent
volume claims when their workload is deleted, or when their pod is scaled
down.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes to override.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentNetwork">
ComponentNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the network configuration for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentService">
[]ClusterComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides services defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentSystemAccount">
[]ComponentSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides system accounts defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentConfig">
[]ClusterComponentConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration content of a config template.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)
for secure communication.
When set to true, the Component will be configured to use TLS encryption for its network connections.
This ensures that the data transmitted between the Component and its clients or other Components is encrypted
and protected from unauthorized access.
If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys,
to properly set up the secure communication channel.</p>
</td>
</tr>
<tr>
<td>
<code>issuer</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Issuer">
Issuer
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration for the TLS certificates issuer.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.
Required when TLS is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceAccount required by the running Component.
This ServiceAccount is used to grant necessary permissions for the Component&rsquo;s Pods to interact
with other Kubernetes resources, such as modifying Pod labels or sending events.</p>
<p>If not specified, KubeBlocks automatically creates a default ServiceAccount named
&ldquo;kb-&#123;componentdefinition.name&#125;&rdquo;, bound to a role with rules defined in ComponentDefinition&rsquo;s
<code>policyRules</code> field. If needed (currently this means if any lifecycleAction is enabled),
it will also be bound to a default role named
&ldquo;kubeblocks-cluster-pod-role&rdquo;, which is installed together with KubeBlocks.
If multiple components use the same ComponentDefinition, they will share one ServiceAccount.</p>
<p>If the field is not empty, the specified ServiceAccount will be used, and KubeBlocks will not
create a ServiceAccount. But KubeBlocks does create RoleBindings for the specified ServiceAccount.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides fine-grained control over the spec update process of all instances.</p>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows for the customization of configuration values for each instance within a Component.
An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).
While instances typically share a common configuration as defined in the ClusterComponentSpec,
they can require unique settings in various scenarios:</p>
<p>For example:
- A database Component might require different resource allocations for primary and secondary instances,
  with primaries needing more resources.
- During a rolling upgrade, a Component may first update the image for one or a few instances,
and then update the remaining instances after verifying that the updated instances are functioning correctly.</p>
<p>InstanceTemplate allows for specifying these unique configurations per instance.
Each instance&rsquo;s name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),
starting with an ordinal of 0.
It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>flatInstanceOrdinal</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>flatInstanceOrdinal controls whether the naming of instances(pods) under this component uses a flattened,
globally uniquely ordinal scheme, regardless of the instance template.</p>
<p>Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>disableExporter</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines whether metrics exporter information is annotated on the Component&rsquo;s headless Service.</p>
<p>If set to true, the following annotations will not be patched into the Service:</p>
<ul>
<li>&ldquo;monitor.kubeblocks.io/path&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/port&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/scheme&rdquo;</li>
</ul>
<p>These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.</p>
</td>
</tr>
<tr>
<td>
<code>stop</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Stop the Component.
If set, all the computing resources will be released.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterComponentStatus">ClusterComponentStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterStatus">ClusterStatus</a>)
</p>
<div>
<p>ClusterComponentStatus records Component status.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentPhase">
ComponentPhase
</a>
</em>
</td>
<td>
<p>Specifies the current state of the Component.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Records detailed information about the Component in its current phase.
The keys are either podName, deployName, or statefulSetName, formatted as &lsquo;ObjectKind/Name&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterDefinitionSpec">ClusterDefinitionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterDefinition">ClusterDefinition</a>)
</p>
<div>
<p>ClusterDefinitionSpec defines the desired state of ClusterDefinition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>topologies</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterTopology">
[]ClusterTopology
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Topologies defines all possible topologies within the cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterDefinitionStatus">ClusterDefinitionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterDefinition">ClusterDefinition</a>)
</p>
<div>
<p>ClusterDefinitionStatus defines the observed state of ClusterDefinition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the most recent generation observed for this ClusterDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Phase">
Phase
</a>
</em>
</td>
<td>
<p>Specifies the current phase of the ClusterDefinition. Valid values are <code>empty</code>, <code>Available</code>, <code>Unavailable</code>.
When <code>Available</code>, the ClusterDefinition is ready and can be referenced by related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
<tr>
<td>
<code>topologies</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Topologies this ClusterDefinition supported.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterObjectReference">ClusterObjectReference
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentVarSelector">ComponentVarSelector</a>, <a href="#apps.kubeblocks.io/v1.CredentialVarSelector">CredentialVarSelector</a>, <a href="#apps.kubeblocks.io/v1.HostNetworkVarSelector">HostNetworkVarSelector</a>, <a href="#apps.kubeblocks.io/v1.ResourceVarSelector">ResourceVarSelector</a>, <a href="#apps.kubeblocks.io/v1.ServiceRefVarSelector">ServiceRefVarSelector</a>, <a href="#apps.kubeblocks.io/v1.ServiceVarSelector">ServiceVarSelector</a>, <a href="#apps.kubeblocks.io/v1.TLSVarSelector">TLSVarSelector</a>)
</p>
<div>
<p>ClusterObjectReference defines information to let you locate the referenced object inside the same Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the exact name, name prefix, or regular expression pattern for matching the name of the ComponentDefinition
custom resource (CR) used by the component that the referent object resident in.</p>
<p>If not specified, the component itself will be used.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Name of the referent object.</p>
</td>
</tr>
<tr>
<td>
<code>optional</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specify whether the object must be defined.</p>
</td>
</tr>
<tr>
<td>
<code>multipleClusterObjectOption</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.MultipleClusterObjectOption">
MultipleClusterObjectOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>This option defines the behavior when multiple component objects match the specified @CompDef.
If not provided, an error will be raised when handling multiple matches.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterPhase">ClusterPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterStatus">ClusterStatus</a>)
</p>
<div>
<p>ClusterPhase defines the phase of the Cluster within the .status.phase field.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Abnormal&#34;</p></td>
<td><p>AbnormalClusterPhase represents some components are in <code>Failed</code> phase, indicates that the cluster is in
a fragile state and troubleshooting is required.</p>
</td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td><p>CreatingClusterPhase represents all components are in <code>Creating</code> phase.</p>
</td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td><p>DeletingClusterPhase indicates the cluster is being deleted.</p>
</td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td><p>FailedClusterPhase represents all components are in <code>Failed</code> phase, indicates that the cluster is unavailable.</p>
</td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td><p>RunningClusterPhase represents all components are in <code>Running</code> phase, indicates that the cluster is functioning properly.</p>
</td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td><p>StoppedClusterPhase represents all components are in <code>Stopped</code> phase, indicates that the cluster has stopped and
is not providing any functionality.</p>
</td>
</tr><tr><td><p>&#34;Stopping&#34;</p></td>
<td><p>StoppingClusterPhase represents at least one component is in <code>Stopping</code> phase, indicates that the cluster is in
the process of stopping.</p>
</td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td><p>UpdatingClusterPhase represents all components are in <code>Creating</code>, <code>Running</code> or <code>Updating</code> phase, and at least one
component is in <code>Creating</code> or <code>Updating</code> phase, indicates that the cluster is undergoing an update.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterService">ClusterService
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ClusterService defines a service that is exposed externally, allowing entities outside the cluster to access it.
For example, external applications, or other Clusters.
And another Cluster managed by the same KubeBlocks operator can resolve the address exposed by a ClusterService
using the <code>serviceRef</code> field.</p>
<p>When a Component needs to access another Cluster&rsquo;s ClusterService using the <code>serviceRef</code> field,
it must also define the service type and version information in the <code>componentDefinition.spec.serviceRefDeclarations</code>
section.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Service</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Service">
Service
</a>
</em>
</td>
<td>
<p>
(Members of <code>Service</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>componentSelector</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extends the ServiceSpec.Selector by allowing the specification of components, to be used as a selector for the service.</p>
<p>If the <code>componentSelector</code> is set as the name of a sharding, the service will be exposed to all components in the sharding.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterSharding">ClusterSharding
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ClusterSharding defines how KubeBlocks manage dynamic provisioned shards.
A typical design pattern for distributed databases is to distribute data across multiple shards,
with each shard consisting of multiple replicas.
Therefore, KubeBlocks supports representing a shard with a Component and dynamically instantiating Components
using a template when shards are added.
When shards are removed, the corresponding Components are also deleted.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the common parent part of all shard names.</p>
<p>This identifier is included as part of the Service DNS name and must comply with IANA service naming rules.
It is used to generate the names of underlying Components following the pattern <code>$(clusterSharding.name)-$(ShardID)</code>.
ShardID is a random string that is appended to the Name to generate unique identifiers for each shard.
For example, if the sharding specification name is &ldquo;my-shard&rdquo; and the ShardID is &ldquo;abc&rdquo;, the resulting Component name
would be &ldquo;my-shard-abc&rdquo;.</p>
<p>Note that the name defined in Component template(<code>clusterSharding.template.name</code>) will be disregarded
when generating the Component names of the shards. The <code>clusterSharding.name</code> field takes precedence.</p>
</td>
</tr>
<tr>
<td>
<code>shardingDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the ShardingDefinition custom resource (CR) that defines the sharding&rsquo;s characteristics and behavior.</p>
<p>The full name or regular expression is supported to match the ShardingDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>shards</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of shards.</p>
<p>Users can declare the desired number of shards through this field.
KubeBlocks dynamically creates and deletes Components based on the difference
between the desired and actual number of shards.
KubeBlocks provides lifecycle management for sharding, including:</p>
<ul>
<li>Executing the shardProvision Action defined in the ShardingDefinition when the number of shards increases.
This allows for custom actions to be performed after a new shard is provisioned.</li>
<li>Executing the shardTerminate Action defined in the ShardingDefinition when the number of shards decreases.
This enables custom cleanup or data migration tasks to be executed before a shard is terminated.
Resources and data associated with the corresponding Component will also be deleted.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">
ClusterComponentSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The default template for generating Components for shards, where each shard consists of one Component.</p>
<p>This field is of type ClusterComponentSpec, which encapsulates all the required details and
definitions for creating and managing the Components.
KubeBlocks uses this template to generate a set of identical Components of shards.
All the generated Components will have the same specifications and definitions as specified in the <code>template</code> field.</p>
<p>This allows for the creation of multiple Components with consistent configurations,
enabling sharding and distribution of workloads across Components.</p>
</td>
</tr>
<tr>
<td>
<code>shardTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardTemplate">
[]ShardTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of heterogeneous shard templates, allowing different groups of shards
to be created with distinct configurations.</p>
</td>
</tr>
<tr>
<td>
<code>offline</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of shards (components) to be transitioned to offline status.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterSpec">ClusterSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Cluster">Cluster</a>)
</p>
<div>
<p>ClusterSpec defines the desired state of Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterDefinition to use when creating a Cluster.</p>
<p>This field enables users to create a Cluster based on a specific ClusterDefinition.
Which, in conjunction with the <code>topology</code> field, determine:</p>
<ul>
<li>The Components to be included in the Cluster.</li>
<li>The sequences in which the Components are created, updated, and terminate.</li>
</ul>
<p>This facilitates multiple-components management with predefined ClusterDefinition.</p>
<p>Users with advanced requirements can bypass this general setting and specify more precise control over
the composition of the Cluster by directly referencing specific ComponentDefinitions for each component
within <code>componentSpecs[*].componentDef</code>.</p>
<p>If this field is not provided, each component must be explicitly defined in <code>componentSpecs[*].componentDef</code>.</p>
<p>Note: Once set, this field cannot be modified; it is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>topology</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterTopology to be used when creating the Cluster.</p>
<p>This field defines which set of Components, as outlined in the ClusterDefinition, will be used to
construct the Cluster based on the named topology.
The ClusterDefinition may list multiple topologies under <code>clusterdefinition.spec.topologies[*]</code>,
each tailored to different use cases or environments.</p>
<p>If <code>topology</code> is not specified, the Cluster will use the default topology defined in the ClusterDefinition.</p>
<p>Note: Once set during the Cluster creation, the <code>topology</code> field cannot be modified.
It establishes the initial composition and structure of the Cluster and is intended for one-time configuration.</p>
</td>
</tr>
<tr>
<td>
<code>terminationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TerminationPolicyType">
TerminationPolicyType
</a>
</em>
</td>
<td>
<p>Specifies the behavior when a Cluster is deleted.
It defines how resources, data, and backups associated with a Cluster are managed during termination.
Choose a policy based on the desired level of resource cleanup and data preservation:</p>
<ul>
<li><code>DoNotTerminate</code>: Prevents deletion of the Cluster. This policy ensures that all resources remain intact.</li>
<li><code>Delete</code>: Deletes all runtime resources belong to the Cluster.</li>
<li><code>WipeOut</code>: An aggressive policy that deletes all Cluster resources, including volume snapshots and
backups in external storage.
This results in complete data removal and should be used cautiously, primarily in non-production environments
to avoid irreversible data loss.</li>
</ul>
<p>Warning: Choosing an inappropriate termination policy can result in data loss.
The <code>WipeOut</code> policy is particularly risky in production environments due to its irreversible nature.</p>
</td>
</tr>
<tr>
<td>
<code>componentSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">
[]ClusterComponentSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ClusterComponentSpec objects used to define the individual Components that make up a Cluster.
This field allows for detailed configuration of each Component within the Cluster.</p>
<p>Note: <code>shardings</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>shardings</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterSharding">
[]ClusterSharding
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ClusterSharding objects that manage the sharding topology for Cluster Components.
Each ClusterSharding organizes components into shards, with each shard corresponding to a Component.
Components within a shard are all based on a common ClusterComponentSpec template, ensuring uniform configurations.</p>
<p>This field supports dynamic resharding by facilitating the addition or removal of shards
through the <code>shards</code> field in ClusterSharding.</p>
<p>Note: <code>shardings</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies runtimeClassName for all Pods managed by this Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterService">
[]ClusterService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of additional Services that are exposed by a Cluster.
This field allows Services of selected Components, either from <code>componentSpecs</code> or <code>shardings</code> to be exposed,
alongside Services defined with ComponentService.</p>
<p>Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterBackup">
ClusterBackup
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the backup configuration of the Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterStatus">ClusterStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Cluster">Cluster</a>)
</p>
<div>
<p>ClusterStatus defines the observed state of the Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>The most recent generation number of the Cluster object that has been observed by the controller.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterPhase">
ClusterPhase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The current phase of the Cluster includes:
<code>Creating</code>, <code>Running</code>, <code>Updating</code>, <code>Stopping</code>, <code>Stopped</code>, <code>Deleting</code>, <code>Failed</code>, <code>Abnormal</code>.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentStatus">
map[string]github.com/apecloud/kubeblocks/apis/apps/v1.ClusterComponentStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Records the current status information of all Components within the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>shardings</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentStatus">
map[string]github.com/apecloud/kubeblocks/apis/apps/v1.ClusterComponentStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Records the current status information of all shardings within the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a list of detailed status of the Cluster object.
Each condition in the list provides real-time information about certain aspect of the Cluster object.</p>
<p>This field is crucial for administrators and developers to monitor and respond to changes within the Cluster.
It provides a history of state transitions and a snapshot of the current state that can be used for
automated logic or direct inspection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterTopology">ClusterTopology
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterDefinitionSpec">ClusterDefinitionSpec</a>)
</p>
<div>
<p>ClusterTopology represents the definition for a specific cluster topology.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the unique identifier for the cluster topology.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterTopologyComponent">
[]ClusterTopologyComponent
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Components specifies the components in the topology.</p>
</td>
</tr>
<tr>
<td>
<code>shardings</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterTopologySharding">
[]ClusterTopologySharding
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Shardings specifies the shardings in the topology.</p>
</td>
</tr>
<tr>
<td>
<code>orders</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterTopologyOrders">
ClusterTopologyOrders
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the sequence in which components within a cluster topology are
started, stopped, and upgraded.
This ordering is crucial for maintaining the correct dependencies and operational flow across components.</p>
</td>
</tr>
<tr>
<td>
<code>default</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Default indicates whether this topology serves as the default configuration.
When set to true, this topology is automatically used unless another is explicitly specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterTopologyComponent">ClusterTopologyComponent
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterTopology">ClusterTopology</a>)
</p>
<div>
<p>ClusterTopologyComponent defines a Component within a ClusterTopology.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the unique identifier of the component within the cluster topology.</p>
<p>It follows IANA Service naming rules and is used as part of the Service&rsquo;s DNS name.
The name must start with a lowercase letter, can contain lowercase letters, numbers,
and hyphens, and must end with a lowercase letter or number.</p>
<p>If the @template field is set to true, the name will be used as a prefix to match the specific components dynamically created.</p>
<p>Cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the exact name, name prefix, or regular expression pattern for matching the name of the ComponentDefinition
custom resource (CR) that defines the Component&rsquo;s characteristics and behavior.</p>
<p>The system selects the ComponentDefinition CR with the latest version that matches the pattern.
This approach allows:</p>
<ol>
<li>Precise selection by providing the exact name of a ComponentDefinition CR.</li>
<li>Flexible and automatic selection of the most up-to-date ComponentDefinition CR
by specifying a name prefix or regular expression pattern.</li>
</ol>
<p>Cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the topology component will be considered as a template for instantiating components upon user requests dynamically.</p>
<p>Cannot be updated once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterTopologyOrders">ClusterTopologyOrders
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterTopology">ClusterTopology</a>)
</p>
<div>
<p>ClusterTopologyOrders manages the lifecycle of components within a cluster by defining their provisioning,
terminating, and updating sequences.
It organizes components into stages or groups, where each group indicates a set of components
that can be managed concurrently.
These groups are processed sequentially, allowing precise control based on component dependencies and requirements.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>provision</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the order for creating and initializing entities.
This is designed for entities that depend on one another. Entities without dependencies can be grouped together.</p>
<p>Entities that can be provisioned independently or have no dependencies can be listed together in the same stage,
separated by commas.</p>
</td>
</tr>
<tr>
<td>
<code>terminate</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Outlines the order for stopping and deleting entities.
This sequence is designed for entities that require a graceful shutdown or have interdependencies.</p>
<p>Entities that can be terminated independently or have no dependencies can be listed together in the same stage,
separated by commas.</p>
</td>
</tr>
<tr>
<td>
<code>update</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Update determines the order for updating entities&rsquo; specifications, such as image upgrades or resource scaling.
This sequence is designed for entities that have dependencies or require specific update procedures.</p>
<p>Entities that can be updated independently or have no dependencies can be listed together in the same stage,
separated by commas.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterTopologySharding">ClusterTopologySharding
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterTopology">ClusterTopology</a>)
</p>
<div>
<p>ClusterTopologySharding defines a sharding within a ClusterTopology.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the unique identifier of the sharding within the cluster topology.
It follows IANA Service naming rules and is used as part of the Service&rsquo;s DNS name.
The name must start with a lowercase letter, can contain lowercase letters, numbers,
and hyphens, and must end with a lowercase letter or number.</p>
<p>Cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>shardingDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the sharding definition that defines the characteristics and behavior of the sharding.</p>
<p>The system selects the ShardingDefinition CR with the latest version that matches the pattern.
This approach allows:</p>
<ol>
<li>Precise selection by providing the exact name of a ShardingDefinition CR.</li>
<li>Flexible and automatic selection of the most up-to-date ShardingDefinition CR
by specifying a regular expression pattern.</li>
</ol>
<p>Once set, this field cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterVarSelector">ClusterVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>ClusterVarSelector selects a var from a Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterVars">
ClusterVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ClusterVars">ClusterVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterVarSelector">ClusterVarSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>namespace</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the namespace of the Cluster object.</p>
</td>
</tr>
<tr>
<td>
<code>clusterName</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the name of the Cluster object.</p>
</td>
</tr>
<tr>
<td>
<code>clusterUID</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the UID of the Cluster object.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentAvailable">ComponentAvailable
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ComponentAvailable defines the strategies for determining whether the component is available.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>withPhases</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the phases that the component will go through to be considered available.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>withRole</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the role that the component will go through to be considered available.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>withProbe</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableWithProbe">
ComponentAvailableWithProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategies for determining whether the component is available based on the available probe.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentAvailableCondition">ComponentAvailableCondition
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentAvailableWithProbe">ComponentAvailableWithProbe</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ComponentAvailableExpression</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableExpression">
ComponentAvailableExpression
</a>
</em>
</td>
<td>
<p>
(Members of <code>ComponentAvailableExpression</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>and</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableExpression">
[]ComponentAvailableExpression
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logical And to combine multiple expressions.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>or</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableExpression">
[]ComponentAvailableExpression
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logical Or to combine multiple expressions.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>not</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableExpression">
ComponentAvailableExpression
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logical Not to negate the expression.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentAvailableExpression">ComponentAvailableExpression
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentAvailableCondition">ComponentAvailableCondition</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>all</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableProbeAssertion">
ComponentAvailableProbeAssertion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>All replicas must satisfy the assertion.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>any</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableProbeAssertion">
ComponentAvailableProbeAssertion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>At least one replica must satisfy the assertion.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>none</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableProbeAssertion">
ComponentAvailableProbeAssertion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>None of the replicas must satisfy the assertion.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>majority</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableProbeAssertion">
ComponentAvailableProbeAssertion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Majority replicas must satisfy the assertion.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentAvailableProbeAssertion">ComponentAvailableProbeAssertion
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentAvailableExpression">ComponentAvailableExpression</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ActionAssertion</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ActionAssertion">
ActionAssertion
</a>
</em>
</td>
<td>
<p>
(Members of <code>ActionAssertion</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>and</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ActionAssertion">
[]ActionAssertion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logical And to combine multiple assertions.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>or</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ActionAssertion">
[]ActionAssertion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logical Or to combine multiple assertions.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>not</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ActionAssertion">
ActionAssertion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Logical Not to negate the assertions.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>strict</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether apply the assertions strictly to all replicas.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentAvailableWithProbe">ComponentAvailableWithProbe
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentAvailable">ComponentAvailable</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>timeWindowSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>condition</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailableCondition">
ComponentAvailableCondition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the conditions that the component will go through to be considered available.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A brief description for the condition when the component is available.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinition">ComponentDefinition</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>provider</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the Component provider, typically the vendor or developer name.
It identifies the entity responsible for creating and maintaining the Component.</p>
<p>When specifying the provider name, consider the following guidelines:</p>
<ul>
<li>Keep the name concise and relevant to the Component.</li>
<li>Use a consistent naming convention across Components from the same provider.</li>
<li>Avoid using trademarked or copyrighted names without proper permission.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a brief and concise explanation of the Component&rsquo;s purpose, functionality, and any relevant details.
It serves as a quick reference for users to understand the Component&rsquo;s role and characteristics.</p>
</td>
</tr>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the type of well-known service protocol that the Component provides.
It specifies the standard or widely recognized protocol used by the Component to offer its Services.</p>
<p>The <code>serviceKind</code> field allows users to quickly identify the type of Service provided by the Component
based on common protocols or service types. This information helps in understanding the compatibility,
interoperability, and usage of the Component within a system.</p>
<p>Some examples of well-known service protocols include:</p>
<ul>
<li>&ldquo;MySQL&rdquo;: Indicates that the Component provides a MySQL database service.</li>
<li>&ldquo;PostgreSQL&rdquo;: Indicates that the Component offers a PostgreSQL database service.</li>
<li>&ldquo;Redis&rdquo;: Signifies that the Component functions as a Redis key-value store.</li>
<li>&ldquo;ETCD&rdquo;: Denotes that the Component serves as an ETCD distributed key-value store.</li>
</ul>
<p>The <code>serviceKind</code> value is case-insensitive, allowing for flexibility in specifying the protocol name.</p>
<p>When specifying the <code>serviceKind</code>, consider the following guidelines:</p>
<ul>
<li>Use well-established and widely recognized protocol names or service types.</li>
<li>Ensure that the <code>serviceKind</code> accurately represents the primary service type offered by the Component.</li>
<li>If the Component provides multiple services, choose the most prominent or commonly used protocol.</li>
<li>Limit the <code>serviceKind</code> to a maximum of 32 characters for conciseness and readability.</li>
</ul>
<p>Note: The <code>serviceKind</code> field is optional and can be left empty if the Component does not fit into a well-known
service category or if the protocol is not widely recognized. It is primarily used to convey information about
the Component&rsquo;s service type to users and facilitate discovery and integration.</p>
<p>The <code>serviceKind</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the version of the Service provided by the Component.
It follows the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
<p>The Semantic Versioning specification defines a version number format of X.Y.Z (MAJOR.MINOR.PATCH), where:</p>
<ul>
<li>X represents the major version and indicates incompatible API changes.</li>
<li>Y represents the minor version and indicates added functionality in a backward-compatible manner.</li>
<li>Z represents the patch version and indicates backward-compatible bug fixes.</li>
</ul>
<p>Additional labels for pre-release and build metadata are available as extensions to the X.Y.Z format:</p>
<ul>
<li>Use pre-release labels (e.g., -alpha, -beta) for versions that are not yet stable or ready for production use.</li>
<li>Use build metadata (e.g., +build.1) for additional version information if needed.</li>
</ul>
<p>Examples of valid ServiceVersion values:</p>
<ul>
<li>&ldquo;1.0.0&rdquo;</li>
<li>&ldquo;2.3.1&rdquo;</li>
<li>&ldquo;3.0.0-alpha.1&rdquo;</li>
<li>&ldquo;4.5.2+build.1&rdquo;</li>
</ul>
<p>The <code>serviceVersion</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static labels that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If a label key in the <code>labels</code> field conflicts with any system labels or user-specified labels,
it will be silently ignored to avoid overriding higher-priority labels.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static annotations that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If an annotation key in the <code>annotations</code> field conflicts with any system annotations
or user-specified annotations, it will be silently ignored to avoid overriding higher-priority annotations.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>runtime</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<p>Specifies the PodSpec template used in the Component.
It includes the following elements:</p>
<ul>
<li>Init containers</li>
<li>Containers
<ul>
<li>Image</li>
<li>Commands</li>
<li>Args</li>
<li>Envs</li>
<li>Mounts</li>
<li>Ports</li>
<li>Security context</li>
<li>Probes</li>
<li>Lifecycle</li>
</ul></li>
<li>Volumes</li>
</ul>
<p>This field is intended to define static settings that remain consistent across all instantiated Components.
Dynamic settings such as CPU and memory resource limits, as well as scheduling settings (affinity,
toleration, priority), may vary among different instantiated Components.
They should be specified in the <code>cluster.spec.componentSpecs</code> (ClusterComponentSpec).</p>
<p>Specific instances of a Component may override settings defined here, such as using a different container image
or modifying environment variable values.
These instance-specific overrides can be specified in <code>cluster.spec.componentSpecs[*].instances</code>.</p>
<p>This field is immutable and cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.EnvVar">
[]EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines variables which are determined after Cluster instantiation and reflect
dynamic or runtime attributes of instantiated Clusters.
These variables serve as placeholders for setting environment variables in Pods and Actions,
or for rendering configuration and script templates before actual values are finalized.</p>
<p>These variables are placed in front of the environment variables declared in the Pod if used as
environment variables.</p>
<p>Variable values can be sourced from:</p>
<ul>
<li>ConfigMap: Select and extract a value from a specific key within a ConfigMap.</li>
<li>Secret: Select and extract a value from a specific key within a Secret.</li>
<li>HostNetwork: Retrieves values (including ports) from host-network resources.</li>
<li>Service: Retrieves values (including address, port, NodePort) from a selected Service.
Intended to obtain the address of a ComponentService within the same Cluster.</li>
<li>Credential: Retrieves account name and password from a SystemAccount variable.</li>
<li>ServiceRef: Retrieves address, port, account name and password from a selected ServiceRefDeclaration.
Designed to obtain the address bound to a ServiceRef, such as a ClusterService or
ComponentService of another cluster or an external service.</li>
<li>Component: Retrieves values from a selected Component, including replicas and instance name list.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVolume">
[]ComponentVolume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the volumes used by the Component and some static attributes of the volumes.
After defining the volumes here, user can reference them in the
<code>cluster.spec.componentSpecs[*].volumeClaimTemplates</code> field to configure dynamic properties such as
volume capacity and storage class.</p>
<p>This field allows you to specify the following:</p>
<ul>
<li>Snapshot behavior: Determines whether a snapshot of the volume should be taken when performing
a snapshot backup of the Component.</li>
<li>Disk high watermark: Sets the high watermark for the volume&rsquo;s disk usage.
When the disk usage reaches the specified threshold, it triggers an alert or action.</li>
</ul>
<p>By configuring these volume behaviors, you can control how the volumes are managed and monitored within the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.HostNetwork">
HostNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the host network configuration for the Component.</p>
<p>When <code>hostNetwork</code> option is enabled, the Pods share the host&rsquo;s network namespace and can directly access
the host&rsquo;s network interfaces.
This means that if multiple Pods need to use the same port, they cannot run on the same host simultaneously
due to port conflicts.</p>
<p>The DNSPolicy field in the Pod spec determines how containers within the Pod perform DNS resolution.
When using hostNetwork, the operator will set the DNSPolicy to &lsquo;ClusterFirstWithHostNet&rsquo;.
With this policy, DNS queries will first go through the K8s cluster&rsquo;s DNS service.
If the query fails, it will fall back to the host&rsquo;s DNS settings.</p>
<p>If set, the DNS policy will be automatically set to &ldquo;ClusterFirstWithHostNet&rdquo;.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines additional Services to expose the Component&rsquo;s endpoints.</p>
<p>A default headless Service, named <code>&#123;cluster.name&#125;-&#123;component.name&#125;-headless</code>, is automatically created
for internal Cluster communication.</p>
<p>This field enables customization of additional Services to expose the Component&rsquo;s endpoints to
other Components within the same or different Clusters, and to external applications.
Each Service entry in this list can include properties such as ports, type, and selectors.</p>
<ul>
<li>For intra-Cluster access, Components can reference Services using variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceVarRef</code>.</li>
<li>For inter-Cluster access, reference Services use variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceRefVarRef</code>,
and bind Services at Cluster creation time with <code>clusterComponentSpec.ServiceRef[*].clusterServiceSelector</code>.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>disableDefaultHeadlessService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to create the default headless service.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the config file templates and volume mount parameters used by the Component.</p>
<p>This field specifies a list of templates that will be rendered into Component containers&rsquo; config files.
Each template is represented as a ConfigMap and may contain multiple config files, with each file being a key in the ConfigMap.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>scripts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies groups of scripts, each provided via a ConfigMap, to be mounted as volumes in the container.
These scripts can be executed during container startup or via specific actions.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>logConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.LogConfig">
[]LogConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the types of logs generated by instances of the Component and their corresponding file paths.
These logs can be collected for further analysis and monitoring.</p>
<p>The <code>logConfigs</code> field is an optional list of LogConfig objects, where each object represents
a specific log type and its configuration.
It allows you to specify multiple log types and their respective file paths for the Component.</p>
<p>Examples:</p>
<pre><code class="language-yaml"> logConfigs:
 - filePathPattern: /data/mysql/log/mysqld-error.log
   name: error
 - filePathPattern: /data/mysql/log/mysqld.log
   name: general
 - filePathPattern: /data/mysql/log/mysqld-slowquery.log
   name: slow
</code></pre>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SystemAccount">
[]SystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An array of <code>SystemAccount</code> objects that define the system accounts needed
for the management operations of the Component.</p>
<p>Each <code>SystemAccount</code> includes:</p>
<ul>
<li>Account name.</li>
<li>The SQL statement template: Used to create the system account.</li>
<li>Password Source: Either generated based on certain rules or retrieved from a Secret.</li>
</ul>
<p>Use cases for system accounts typically involve tasks like system initialization, backups, monitoring,
health checks, replication, and other system-level operations.</p>
<p>System accounts are distinct from user accounts, although both are database accounts.</p>
<ul>
<li><strong>System Accounts</strong>: Created during Cluster setup by the KubeBlocks operator,
these accounts have higher privileges for system management and are fully managed
through a declarative API by the operator.</li>
<li><strong>User Accounts</strong>: Managed by users or administrator.
User account permissions should follow the principle of least privilege,
granting only the necessary access rights to complete their required tasks.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TLS">
TLS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the TLS configuration for the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>replicasLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ReplicasLimit">
ReplicasLimit
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the upper limit of the number of replicas supported by the Component.</p>
<p>It defines the maximum number of replicas that can be created for the Component.
This field allows you to set a limit on the scalability of the Component, preventing it from exceeding a certain number of replicas.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>available</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentAvailable">
ComponentAvailable
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategies for determining the available status of the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enumerate all possible roles assigned to each replica of the Component, influencing its behavior.</p>
<p>A replica can have zero or one role.
KubeBlocks operator determines the role of each replica by invoking the <code>lifecycleActions.roleProbe</code> method.
This action returns the role for each replica, and the returned role must be predefined here.</p>
<p>The roles assigned to a replica can influence various aspects of the Component&rsquo;s behavior, such as:</p>
<ul>
<li>Service selection: The Component&rsquo;s exposed Services may target replicas based on their roles using <code>roleSelector</code>.</li>
<li>Update order: The roles can determine the order in which replicas are updated during a Component update.
For instance, replicas with a &ldquo;follower&rdquo; role can be updated first, while the replica with the &ldquo;leader&rdquo;
role is updated last. This helps minimize the number of leader changes during the update process.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p><code>minReadySeconds</code> is the minimum duration in seconds that a new Pod should remain in the ready
state without any of its containers crashing to be considered available.
This ensures the Pod&rsquo;s stability and readiness to serve requests.</p>
<p>A default value of 0 seconds means the Pod is considered available as soon as it enters the ready state.</p>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the concurrency strategy for updating multiple instances of the Component.
Available strategies:</p>
<ul>
<li><code>Serial</code>: Updates replicas one at a time, ensuring minimal downtime by waiting for each replica to become ready
before updating the next.</li>
<li><code>Parallel</code>: Updates all replicas simultaneously, optimizing for speed but potentially reducing availability
during the update.</li>
<li><code>BestEffortParallel</code>: Updates replicas concurrently with a limit on simultaneous updates to ensure a minimum
number of operational replicas for maintaining quorum.
 For example, in a 5-replica component, updating a maximum of 2 replicas simultaneously keeps
at least 3 operational for quorum.</li>
</ul>
<p>This field is immutable and defaults to &lsquo;Serial&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InstanceSet controls the creation of pods during initial scale up, replacement of pods on nodes, and scaling down.</p>
<ul>
<li><code>OrderedReady</code>: Creates pods in increasing order (pod-0, then pod-1, etc). The controller waits until each pod
is ready before continuing. Pods are removed in reverse order when scaling down.</li>
<li><code>Parallel</code>: Creates pods in parallel to match the desired scale without waiting. All pods are deleted at once
when scaling down.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>policyRules</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#policyrule-v1-rbac">
[]Kubernetes rbac/v1.PolicyRule
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the namespaced policy rules required by the Component.</p>
<p>The <code>policyRules</code> field is an array of <code>rbacv1.PolicyRule</code> objects that define the policy rules
needed by the Component to operate within a namespace.
These policy rules determine the permissions and verbs the Component is allowed to perform on
Kubernetes resources within the namespace.</p>
<p>The purpose of this field is to automatically generate the necessary RBAC roles
for the Component based on the specified policy rules.
This ensures that the Pods in the Component has appropriate permissions to function.</p>
<p>To prevent privilege escalation, only permissions already owned by KubeBlocks can be added here.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>lifecycleActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentLifecycleActions">
ComponentLifecycleActions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of hooks and procedures that customize the behavior of a Component throughout its lifecycle.
Actions are triggered at specific lifecycle stages:</p>
<ul>
<li><code>postProvision</code>: Defines the hook to be executed after the creation of a Component,
with <code>preCondition</code> specifying when the action should be fired relative to the Component&rsquo;s lifecycle stages:
<code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.</li>
<li><code>preTerminate</code>: Defines the hook to be executed before terminating a Component.</li>
<li><code>roleProbe</code>: Defines the procedure which is invoked regularly to assess the role of replicas.</li>
<li><code>switchover</code>: Defines the procedure for a controlled transition of a role to a new replica.
This approach aims to minimize downtime and maintain availability in systems with a leader-follower topology,
such as before planned maintenance or upgrades on the current leader node.</li>
<li><code>memberJoin</code>: Defines the procedure to add a new replica to the replication group.</li>
<li><code>memberLeave</code>: Defines the method to remove a replica from the replication group.</li>
<li><code>readOnly</code>: Defines the procedure to switch a replica into the read-only state.</li>
<li><code>readWrite</code>: transition a replica from the read-only state back to the read-write state.</li>
<li><code>dataDump</code>: Defines the procedure to export the data from a replica.</li>
<li><code>dataLoad</code>: Defines the procedure to import data into a replica.</li>
<li><code>reconfigure</code>: Defines the procedure that update a replica with new configuration file.</li>
<li><code>accountProvision</code>: Defines the procedure to generate a new database account.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefDeclarations</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefDeclaration">
[]ServiceRefDeclaration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists external service dependencies of the Component, including services from other Clusters or outside the K8s environment.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>exporter</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Exporter">
Exporter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the built-in metrics exporter container.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentDefinitionStatus">ComponentDefinitionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinition">ComponentDefinition</a>)
</p>
<div>
<p>ComponentDefinitionStatus defines the observed state of ComponentDefinition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the most recent generation that has been observed for the ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the current status of the ComponentDefinition. Valid values include `<code>,</code>Available<code>, and</code>Unavailable<code>.
When the status is</code>Available`, the ComponentDefinition is ready and can be utilized by related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentFileTemplate">ComponentFileTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#apps.kubeblocks.io/v1.SidecarDefinitionSpec">SidecarDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the template.</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the referenced template ConfigMap object.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the namespace of the referenced template ConfigMap object.</p>
</td>
</tr>
<tr>
<td>
<code>volumeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the volume name of PodTemplate. The file produced through the template will be mounted to
the corresponding volume. Must be a DNS_LABEL name.
The volume name must be defined in podSpec.containers[*].volumeMounts.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMode</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The operator attempts to set default file permissions (0444).</p>
<p>Must be specified as an octal value between 0000 and 0777 (inclusive),
or as a decimal value between 0 and 511 (inclusive).
YAML supports both octal and decimal values for file permissions.</p>
<p>Please note that this setting only affects the permissions of the files themselves.
Directories within the specified path are not impacted by this setting.
It&rsquo;s important to be aware that this setting might conflict with other options
that influence the file mode, such as fsGroup.
In such cases, the resulting file mode may have additional bits set.
Refers to documents of k8s.ConfigMapVolumeSource.defaultMode for more information.</p>
</td>
</tr>
<tr>
<td>
<code>externalManaged</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>ExternalManaged indicates whether the configuration is managed by an external system.
When set to true, the controller will ignore the management of this configuration.</p>
</td>
</tr>
<tr>
<td>
<code>restartOnFileChange</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to restart the pod when the file changes.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentLifecycleActions">ComponentLifecycleActions
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ComponentLifecycleActions defines a collection of Actions for customizing the behavior of a Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>postProvision</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed after a component&rsquo;s creation.</p>
<p>By setting <code>postProvision.customHandler.preCondition</code>, you can determine the specific lifecycle stage
at which the action should trigger: <code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.
with <code>ComponentReady</code> being the default.</p>
<p>The PostProvision Action is intended to run only once.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>preTerminate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed prior to terminating a component.</p>
<p>The PreTerminate Action is intended to run only once.</p>
<p>This action is executed immediately when a scale-down operation for the Component is initiated.
The actual termination and cleanup of the Component and its associated resources will not proceed
until the PreTerminate action has completed successfully.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>roleProbe</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Probe">
Probe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure which is invoked regularly to assess the role of replicas.</p>
<p>This action is periodically triggered at the specified interval to determine the role of each replica.
Upon successful execution, the action&rsquo;s output designates the role of the replica,
which should match one of the predefined role names within <code>componentDefinition.spec.roles</code>.
The output is then compared with the previous successful execution result.
If a role change is detected, an event is generated to inform the controller,
which initiates an update of the replica&rsquo;s role.</p>
<p>Defining a RoleProbe Action for a Component is required if roles are defined for the Component.
It ensures replicas are correctly labeled with their respective roles.
Without this, services that rely on roleSelectors might improperly direct traffic to wrong replicas.</p>
<p>The container executing this action has access to following variables:</p>
<ul>
<li>KB_POD_FQDN: The FQDN of the Pod whose role is being assessed.</li>
</ul>
<p>Expected output of this action:
- On Success: The determined role of the replica, which must align with one of the roles specified
  in the component definition.
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>availableProbe</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Probe">
Probe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure which is invoked regularly to assess the availability of the component.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>switchover</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure for a controlled transition of a role to a new replica.
This approach aims to minimize downtime and maintain availability
during events such as planned maintenance or when performing stop, shutdown, restart, or upgrade operations.
In a typical consensus system, this action is used to transfer leader role to another replica.</p>
<p>The container executing this action has access to following variables:</p>
<ul>
<li>KB_SWITCHOVER_CANDIDATE_NAME: The name of the pod of the new role&rsquo;s candidate, which may not be specified (empty).</li>
<li>KB_SWITCHOVER_CANDIDATE_FQDN: The FQDN of the pod of the new role&rsquo;s candidate, which may not be specified (empty).</li>
<li>KB_SWITCHOVER_CURRENT_NAME: The name of the pod of the current role.</li>
<li>KB_SWITCHOVER_CURRENT_FQDN: The FQDN of the pod of the current role.</li>
<li>KB_SWITCHOVER_ROLE: The role that will be transferred to another replica.
This variable can be empty if, for example, role probe does not succeed.
It depends on the addon implementation what to do under such cases.</li>
</ul>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>memberJoin</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to add a new replica to the replication group.</p>
<p>This action is initiated after a replica pod becomes ready.</p>
<p>The role of the replica (e.g., primary, secondary) will be determined and assigned as part of the action command
implementation, or automatically by the database kernel or a sidecar utility like Patroni that implements
a consensus algorithm.</p>
<p>The container executing this action has access to following variables:</p>
<ul>
<li>KB_JOIN_MEMBER_POD_FQDN: The pod FQDN of the replica being added to the group.</li>
<li>KB_JOIN_MEMBER_POD_NAME: The pod name of the replica being added to the group.</li>
</ul>
<p>Expected action output:
- On Failure: An error message detailing the reason for any failure encountered
during the addition of the new member.</p>
<p>For example, to add a new OBServer to an OceanBase Cluster in &lsquo;zone1&rsquo;, the following command may be used:</p>
<pre><code class="language-yaml">command:
- bash
- -c
- |
   CLIENT=&quot;mysql -u $SERVICE_USER -p$SERVICE_PASSWORD -P $SERVICE_PORT -h $SERVICE_HOST -e&quot;
	  $CLIENT &quot;ALTER SYSTEM ADD SERVER '$KB_POD_FQDN:$SERVICE_PORT' ZONE 'zone1'&quot;
</code></pre>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>memberLeave</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to remove a replica from the replication group.</p>
<p>This action is initiated before remove a replica from the group.
The operator will wait for MemberLeave to complete successfully before releasing the replica and cleaning up
related Kubernetes resources.</p>
<p>The process typically includes updating configurations and informing other group members about the removal.
Data migration is generally not part of this action and should be handled separately if needed.</p>
<p>The container executing this action has access to following variables:</p>
<ul>
<li>KB_LEAVE_MEMBER_POD_FQDN: The pod name of the replica being removed from the group.</li>
<li>KB_LEAVE_MEMBER_POD_NAME: The pod name of the replica being removed from the group.</li>
</ul>
<p>Expected action output:
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>For example, to remove an OBServer from an OceanBase Cluster in &lsquo;zone1&rsquo;, the following command can be executed:</p>
<pre><code class="language-yaml">command:
- bash
- -c
- |
   CLIENT=&quot;mysql -u $SERVICE_USER -p$SERVICE_PASSWORD -P $SERVICE_PORT -h $SERVICE_HOST -e&quot;
	  $CLIENT &quot;ALTER SYSTEM DELETE SERVER '$KB_POD_FQDN:$SERVICE_PORT' ZONE 'zone1'&quot;
</code></pre>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>readonly</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to switch a replica into the read-only state.</p>
<p>Use Case:
This action is invoked when the database&rsquo;s volume capacity nears its upper limit and space is about to be exhausted.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_POD_FQDN: The FQDN of the replica pod whose role is being checked.</li>
</ul>
<p>Expected action output:
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>readwrite</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to transition a replica from the read-only state back to the read-write state.</p>
<p>Use Case:
This action is used to bring back a replica that was previously in a read-only state,
which restricted write operations, to its normal operational state where it can handle
both read and write operations.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_POD_FQDN: The FQDN of the replica pod whose role is being checked.</li>
</ul>
<p>Expected action output:
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>dataDump</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure for exporting the data from a replica.</p>
<p>Use Case:
This action is intended for initializing a newly created replica with data. It involves exporting data
from an existing replica and importing it into the new, empty replica. This is essential for synchronizing
the state of replicas across the system.</p>
<p>Applicability:
Some database engines or associated sidecar applications (e.g., Patroni) may already provide this functionality.
In such cases, this action may not be required.</p>
<p>The output should be a valid data dump streamed to stdout. It must exclude any irrelevant information to ensure
that only the necessary data is exported for import into the new replica.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_TARGET_POD_NAME: The name of the replica pod into which the data will be loaded.</li>
</ul>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>dataLoad</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure for importing data into a replica.</p>
<p>Use Case:
This action is intended for initializing a newly created replica with data. It involves exporting data
from an existing replica and importing it into the new, empty replica. This is essential for synchronizing
the state of replicas across the system.</p>
<p>Some database engines or associated sidecar applications (e.g., Patroni) may already provide this functionality.
In such cases, this action may not be required.</p>
<p>Data should be received through stdin. If any error occurs during the process,
the action must be able to guarantee idempotence to allow for retries from the beginning.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>reconfigure</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure that update a replica with new configuration.</p>
<p>Note: This field is immutable once it has been set.</p>
<p>This Action is reserved for future versions.</p>
</td>
</tr>
<tr>
<td>
<code>accountProvision</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to generate a new database account.</p>
<p>Use Case:
This action is designed to create system accounts that are utilized for replication, monitoring, backup,
and other administrative tasks.</p>
<p>The container executing this action has access to following variables:</p>
<ul>
<li>KB_ACCOUNT_NAME: The name of the system account to be manipulated.</li>
<li>KB_ACCOUNT_PASSWORD: The password for the system account.</li>
<li>KB_ACCOUNT_STATEMENT: The statement used to manipulate the system account.</li>
</ul>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentNetwork">ComponentNetwork
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>hostNetwork</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Host networking requested for this pod. Use the host&rsquo;s network namespace.</p>
</td>
</tr>
<tr>
<td>
<code>hostAliases</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#hostalias-v1-core">
[]Kubernetes core/v1.HostAlias
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>HostAliases is an optional list of hosts and IPs that will be injected into the pod&rsquo;s hosts file if specified.
This is only valid for non-hostNetwork pods.</p>
</td>
</tr>
<tr>
<td>
<code>dnsPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#dnspolicy-v1-core">
Kubernetes core/v1.DNSPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Set DNS policy for the pod.
Defaults to &ldquo;ClusterFirst&rdquo;. If the hostNetwork is enabled, the default policy will be set to &ldquo;ClusterFirstWithHostNet&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>dnsConfig</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#poddnsconfig-v1-core">
Kubernetes core/v1.PodDNSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the DNS parameters of a pod.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentPhase">ComponentPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentStatus">ClusterComponentStatus</a>, <a href="#apps.kubeblocks.io/v1.ComponentStatus">ComponentStatus</a>)
</p>
<div>
<p>ComponentPhase defines the phase of the Component within the .status.phase field.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Creating&#34;</p></td>
<td><p>CreatingComponentPhase indicates the component is currently being created.</p>
</td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td><p>DeletingComponentPhase indicates the component is currently being deleted.</p>
</td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td><p>FailedComponentPhase indicates that there are some pods of the component not in a &lsquo;Running&rsquo; state.</p>
</td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td><p>RunningComponentPhase indicates that all pods of the component are up-to-date and in a &lsquo;Running&rsquo; state.</p>
</td>
</tr><tr><td><p>&#34;Starting&#34;</p></td>
<td><p>StartingComponentPhase indicates the component is currently being started.</p>
</td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td><p>StoppedComponentPhase indicates the component is stopped.</p>
</td>
</tr><tr><td><p>&#34;Stopping&#34;</p></td>
<td><p>StoppingComponentPhase indicates the component is currently being stopped.</p>
</td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td><p>UpdatingComponentPhase indicates the component is currently being updated.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentService">ComponentService
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
<p>ComponentService defines a service that would be exposed as an inter-component service within a Cluster.
A Service defined in the ComponentService is expected to be accessed by other Components within the same Cluster.</p>
<p>When a Component needs to use a ComponentService provided by another Component within the same Cluster,
it can declare a variable in the <code>componentDefinition.spec.vars</code> section and bind it to the specific exposed address
of the ComponentService using the <code>serviceVarRef</code> field.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Service</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Service">
Service
</a>
</em>
</td>
<td>
<p>
(Members of <code>Service</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>podService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to create a corresponding Service for each Pod of the selected Component.
When set to true, a set of Services will be automatically generated for each Pod,
and the <code>roleSelector</code> field will be ignored.</p>
<p>The names of the generated Services will follow the same suffix naming pattern: <code>$(serviceName)-$(podOrdinal)</code>.
The total number of generated Services will be equal to the number of replicas specified for the Component.</p>
<p>Example usage:</p>
<pre><code class="language-yaml">name: my-service
serviceName: my-service
podService: true
disableAutoProvision: true
spec:
  type: NodePort
  ports:
  - name: http
    port: 80
    targetPort: 8080
</code></pre>
<p>In this example, if the Component has 3 replicas, three Services will be generated:
- my-service-0: Points to the first Pod (podOrdinal: 0)
- my-service-1: Points to the second Pod (podOrdinal: 1)
- my-service-2: Points to the third Pod (podOrdinal: 2)</p>
<p>Each generated Service will have the specified spec configuration and will target its respective Pod.</p>
<p>This feature is useful when you need to expose each Pod of a Component individually, allowing external access
to specific instances of the Component.</p>
</td>
</tr>
<tr>
<td>
<code>disableAutoProvision</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether the automatic provisioning of the service should be disabled.</p>
<p>If set to true, the service will not be automatically created at the component provisioning.
Instead, you can enable the creation of this service by specifying it explicitly in the cluster API.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Component">Component</a>)
</p>
<div>
<p>ComponentSpec defines the desired state of Component</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>terminationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TerminationPolicyType">
TerminationPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the behavior when a Component is deleted.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this Component.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRef">
[]ServiceRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of ServiceRef for a Component, enabling access to both external services and
Services provided by other Clusters.</p>
<p>Types of services:</p>
<ul>
<li>External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;
Require a ServiceDescriptor for connection details.</li>
<li>Services provided by a Cluster: Managed by the same KubeBlocks operator;
identified using Cluster, Component and Service names.</li>
</ul>
<p>ServiceRefs with identical <code>serviceRef.name</code> in the same Cluster are considered the same.</p>
<p>Example:</p>
<pre><code class="language-yaml">serviceRefs:
  - name: &quot;redis-sentinel&quot;
    serviceDescriptor:
      name: &quot;external-redis-sentinel&quot;
  - name: &quot;postgres-cluster&quot;
    clusterServiceSelector:
      cluster: &quot;my-postgres-cluster&quot;
      service:
        component: &quot;postgresql&quot;
</code></pre>
<p>The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Labels to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Annotations to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of environment variables to add.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources required by the Component.
It allows defining the CPU, memory requirements and limits for the Component&rsquo;s containers.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimTemplate">
[]PersistentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for the Component.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>persistentVolumeClaimRetentionPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicy">
PersistentVolumeClaimRetentionPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>persistentVolumeClaimRetentionPolicy describes the lifecycle of persistent
volume claims created from volumeClaimTemplates. By default, all persistent
volume claims are created as needed and retained until manually deleted. This
policy allows the lifecycle to be altered, for example by deleting persistent
volume claims when their workload is deleted, or when their pod is scaled
down.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes to override.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentNetwork">
ComponentNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the network configuration for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides Services defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentSystemAccount">
[]ComponentSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides system accounts defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterComponentConfig">
[]ClusterComponentConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration content of a config template.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceAccount required by the running Component.
This ServiceAccount is used to grant necessary permissions for the Component&rsquo;s Pods to interact
with other Kubernetes resources, such as modifying Pod labels or sending events.</p>
<p>If not specified, KubeBlocks automatically creates a default ServiceAccount named
&ldquo;kb-&#123;componentdefinition.name&#125;&rdquo;, bound to a role with rules defined in ComponentDefinition&rsquo;s
<code>policyRules</code> field. If needed (currently this means if any lifecycleAction is enabled),
it will also be bound to a default role named
&ldquo;kubeblocks-cluster-pod-role&rdquo;, which is installed together with KubeBlocks.
If multiple components use the same ComponentDefinition, they will share one ServiceAccount.</p>
<p>If the field is not empty, the specified ServiceAccount will be used, and KubeBlocks will not
create a ServiceAccount. But KubeBlocks does create RoleBindings for the specified ServiceAccount.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides fine-grained control over the spec update process of all instances.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>tlsConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TLSConfig">
TLSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the TLS configuration for the Component, including:</p>
<ul>
<li>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication.</li>
<li>An optional field that specifies the configuration for the TLS certificates issuer when TLS is enabled.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows for the customization of configuration values for each instance within a Component.
An Instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).
While instances typically share a common configuration as defined in the ClusterComponentSpec,
they can require unique settings in various scenarios:</p>
<p>For example:
- A database Component might require different resource allocations for primary and secondary instances,
  with primaries needing more resources.
- During a rolling upgrade, a Component may first update the image for one or a few instances,
and then update the remaining instances after verifying that the updated instances are functioning correctly.</p>
<p>InstanceTemplate allows for specifying these unique configurations per instance.
Each instance&rsquo;s name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),
starting with an ordinal of 0.
It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the Component.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>flatInstanceOrdinal</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>flatInstanceOrdinal controls whether the naming of instances(pods) under this component uses a flattened,
globally uniquely ordinal scheme, regardless of the instance template.</p>
<p>Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the Cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines runtimeClassName for all Pods managed by this Component.</p>
</td>
</tr>
<tr>
<td>
<code>disableExporter</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines whether metrics exporter information is annotated on the Component&rsquo;s headless Service.</p>
<p>If set to true, the following annotations will not be patched into the Service:</p>
<ul>
<li>&ldquo;monitor.kubeblocks.io/path&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/port&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/scheme&rdquo;</li>
</ul>
<p>These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.</p>
</td>
</tr>
<tr>
<td>
<code>stop</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Stop the Component.
If set, all the computing resources will be released.</p>
</td>
</tr>
<tr>
<td>
<code>sidecars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Sidecar">
[]Sidecar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the sidecars to be injected into the Component.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentStatus">ComponentStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Component">Component</a>)
</p>
<div>
<p>ComponentStatus represents the observed state of a Component within the Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the most recent generation observed for this Component object.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a list of detailed status of the Component object.
Each condition in the list provides real-time information about certain aspect of the Component object.</p>
<p>This field is crucial for administrators and developers to monitor and respond to changes within the Component.
It provides a history of state transitions and a snapshot of the current state that can be used for
automated logic or direct inspection.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentPhase">
ComponentPhase
</a>
</em>
</td>
<td>
<p>Indicates the current phase of the Component, with each phase indicating specific conditions:</p>
<ul>
<li>Creating: The initial phase for new Components, transitioning from &lsquo;empty&rsquo;(&ldquo;&rdquo;).</li>
<li>Running: All Pods are up-to-date and in a Running state.</li>
<li>Updating: The Component is currently being updated, with no failed Pods present.</li>
<li>Failed: A significant number of Pods have failed.</li>
<li>Stopping: All Pods are being terminated, with current replica count at zero.</li>
<li>Stopped: All associated Pods have been successfully deleted.</li>
<li>Starting: Pods are being started.</li>
<li>Deleting: The Component is being deleted.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A map that stores detailed message about the Component.
Each entry in the map provides insights into specific elements of the Component, such as Pods or workloads.</p>
<p>Keys in this map are formatted as <code>ObjectKind/Name</code>, where <code>ObjectKind</code> could be a type like Pod,
and <code>Name</code> is the specific name of the object.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentSystemAccount">ComponentSystemAccount
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the system account.</p>
</td>
</tr>
<tr>
<td>
<code>disabled</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the system account is disabled.</p>
</td>
</tr>
<tr>
<td>
<code>passwordConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PasswordConfig">
PasswordConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the policy for generating the account&rsquo;s password.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>secretRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ProvisionSecretRef">
ProvisionSecretRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the secret from which data will be copied to create the new account.</p>
<p>For user-specified passwords, the maximum length is limited to 64 bytes.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVarSelector">ComponentVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>ComponentVarSelector selects a var from a Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Component to select from.</p>
</td>
</tr>
<tr>
<td>
<code>ComponentVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVars">
ComponentVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ComponentVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVars">ComponentVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentVarSelector">ComponentVarSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>componentName</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the name of the Component object.</p>
</td>
</tr>
<tr>
<td>
<code>shortName</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the short name of the Component object.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the replicas of the component.</p>
</td>
</tr>
<tr>
<td>
<code>podNames</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the pod name list of the component.
and the value will be presented in the following format: name1,name2,&hellip;</p>
</td>
</tr>
<tr>
<td>
<code>podFQDNs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the pod FQDN list of the component.
The value will be presented in the following format: FQDN1,FQDN2,&hellip;</p>
</td>
</tr>
<tr>
<td>
<code>podNamesForRole</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.RoledVar">
RoledVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the pod name list of the component that have a specific role.
The value will be presented in the following format: name1,name2,&hellip;</p>
</td>
</tr>
<tr>
<td>
<code>podFQDNsForRole</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.RoledVar">
RoledVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the pod FQDN list of the component that have a specific role.
The value will be presented in the following format: FQDN1,FQDN2,&hellip;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVersionCompatibilityRule">ComponentVersionCompatibilityRule
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentVersionSpec">ComponentVersionSpec</a>)
</p>
<div>
<p>ComponentVersionCompatibilityRule defines the compatibility between a set of component definitions and a set of releases.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compDefs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CompDefs specifies names for the component definitions associated with this ComponentVersion.
Each name in the list can represent an exact name, a name prefix, or a regular expression pattern.</p>
<p>For example:</p>
<ul>
<li>&ldquo;mysql-8.0.30-v1alpha1&rdquo;: Matches the exact name &ldquo;mysql-8.0.30-v1alpha1&rdquo;</li>
<li>&ldquo;mysql-8.0.30&rdquo;: Matches all names starting with &ldquo;mysql-8.0.30&rdquo;</li>
<li>&rdquo;^mysql-8.0.\d&#123;1,2&#125;$&ldquo;: Matches all names starting with &ldquo;mysql-8.0.&rdquo; followed by one or two digits.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>releases</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Releases is a list of identifiers for the releases.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVersionRelease">ComponentVersionRelease
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentVersionSpec">ComponentVersionSpec</a>)
</p>
<div>
<p>ComponentVersionRelease represents a release of component instances within a ComponentVersion.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is a unique identifier for this release.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>changes</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Changes provides information about the changes made in this release.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceVersion defines the version of the well-known service that the component provides.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).
If the release is used, it will serve as the service version for component instances, overriding the one defined in the component definition.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>images</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Images define the new images for containers, actions or external applications within the release.</p>
<p>If an image is specified for a lifecycle action, the key should be the field name (case-insensitive) of
the action in the LifecycleActions struct.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVersionSpec">ComponentVersionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentVersion">ComponentVersion</a>)
</p>
<div>
<p>ComponentVersionSpec defines the desired state of ComponentVersion</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityRules</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVersionCompatibilityRule">
[]ComponentVersionCompatibilityRule
</a>
</em>
</td>
<td>
<p>CompatibilityRules defines compatibility rules between sets of component definitions and releases.</p>
</td>
</tr>
<tr>
<td>
<code>releases</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVersionRelease">
[]ComponentVersionRelease
</a>
</em>
</td>
<td>
<p>Releases represents different releases of component instances within this ComponentVersion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVersionStatus">ComponentVersionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentVersion">ComponentVersion</a>)
</p>
<div>
<p>ComponentVersionStatus defines the observed state of ComponentVersion</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ObservedGeneration is the most recent generation observed for this ComponentVersion.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Phase valid values are `<code>,</code>Available<code>, 'Unavailable</code>.
Available is ComponentVersion become available, and can be used for co-related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extra message for current phase.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersions</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersions represent the supported service versions of this ComponentVersion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ComponentVolume">ComponentVolume
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the volume.
It must be a DNS_LABEL and unique within the pod.
More info can be found at: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names">https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</a>
Note: This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>needSnapshot</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the creation of a snapshot of this volume is necessary when performing a backup of the Component.</p>
<p>Note: This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>highWatermark</code><br/>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sets the critical threshold for volume space utilization as a percentage (0-100).</p>
<p>Exceeding this percentage triggers the system to switch the volume to read-only mode as specified in
<code>componentDefinition.spec.lifecycleActions.readOnly</code>.
This precaution helps prevent space depletion while maintaining read-only access.
If the space utilization later falls below this threshold, the system reverts the volume to read-write mode
as defined in <code>componentDefinition.spec.lifecycleActions.readWrite</code>, restoring full functionality.</p>
<p>Note: This field cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ConnectionCredentialAuth">ConnectionCredentialAuth
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceDescriptorSpec">ServiceDescriptorSpec</a>)
</p>
<div>
<p>ConnectionCredentialAuth specifies the authentication credentials required for accessing an external service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>username</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the username for the external service.</p>
</td>
</tr>
<tr>
<td>
<code>password</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the password for the external service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ContainerVars">ContainerVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.HostNetworkVars">HostNetworkVars</a>)
</p>
<div>
<p>ContainerVars defines the vars that can be referenced from a Container.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the container.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.NamedVar">
NamedVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Container port to reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.CredentialVar">CredentialVar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ConnectionCredentialAuth">ConnectionCredentialAuth</a>, <a href="#apps.kubeblocks.io/v1.ServiceDescriptorSpec">ServiceDescriptorSpec</a>)
</p>
<div>
<p>CredentialVar represents a variable that retrieves its value either directly from a specified expression
or from a source defined in <code>valueFrom</code>.
Only one of these options may be used at a time.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Holds a direct string or an expression that can be evaluated to a string.</p>
<p>It can include variables denoted by $(VAR_NAME).
These variables are expanded to the value of the environment variables defined in the container.
If a variable cannot be resolved, it remains unchanged in the output.</p>
<p>To escape variable expansion and retain the literal value, use double $ characters.</p>
<p>For example:</p>
<ul>
<li>&rdquo;$(VAR_NAME)&rdquo; will be expanded to the value of the environment variable VAR_NAME.</li>
<li>&rdquo;$$(VAR_NAME)&rdquo; will result in &ldquo;$(VAR_NAME)&rdquo; in the output, without any variable expansion.</li>
</ul>
<p>Default value is an empty string.</p>
</td>
</tr>
<tr>
<td>
<code>valueFrom</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvarsource-v1-core">
Kubernetes core/v1.EnvVarSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the source for the variable&rsquo;s value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.CredentialVarSelector">CredentialVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>CredentialVarSelector selects a var from a Credential (SystemAccount).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Credential (SystemAccount) to select from.</p>
</td>
</tr>
<tr>
<td>
<code>CredentialVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVars">
CredentialVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>CredentialVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.CredentialVars">CredentialVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.CredentialVarSelector">CredentialVarSelector</a>, <a href="#apps.kubeblocks.io/v1.ServiceRefVars">ServiceRefVars</a>)
</p>
<div>
<p>CredentialVars defines the vars that can be referenced from a Credential (SystemAccount).
!!!!! CredentialVars will only be used as environment variables for Pods &amp; Actions, and will not be used to render the templates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>username</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>password</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.EnvVar">EnvVar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#apps.kubeblocks.io/v1.SidecarDefinitionSpec">SidecarDefinitionSpec</a>)
</p>
<div>
<p>EnvVar represents a variable present in the env of Pod/Action or the template of config/script.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the variable. Must be a C_IDENTIFIER.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Variable references <code>$(VAR_NAME)</code> are expanded using the previously defined variables in the current context.</p>
<p>If a variable cannot be resolved, the reference in the input string will be unchanged.
Double <code>$$</code> are reduced to a single <code>$</code>, which allows for escaping the <code>$(VAR_NAME)</code> syntax: i.e.</p>
<ul>
<li><code>$$(VAR_NAME)</code> will produce the string literal <code>$(VAR_NAME)</code>.</li>
</ul>
<p>Escaped references will never be expanded, regardless of whether the variable exists or not.
Defaults to &ldquo;&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>valueFrom</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarSource">
VarSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Source for the variable&rsquo;s value. Cannot be used if value is not empty.</p>
</td>
</tr>
<tr>
<td>
<code>expression</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A Go template expression that will be applied to the resolved value of the var.</p>
<p>The expression will only be evaluated if the var is successfully resolved to a non-credential value.</p>
<p>The resolved value can be accessed by its name within the expression, system vars and other user-defined
non-credential vars can be used within the expression in the same way.
Notice that, when accessing vars by its name, you should replace all the &ldquo;-&rdquo; in the name with &ldquo;_&rdquo;, because of
that &ldquo;-&rdquo; is not a valid identifier in Go.</p>
<p>All expressions are evaluated in the order the vars are defined. If a var depends on any vars that also
have expressions defined, be careful about the evaluation order as it may use intermediate values.</p>
<p>The result of evaluation will be used as the final value of the var. If the expression fails to evaluate,
the resolving of var will also be considered failed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ExecAction">ExecAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Action">Action</a>)
</p>
<div>
<p>ExecAction describes an Action that executes a command inside a container.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the container image to be used for running the Action.</p>
<p>When specified, a dedicated container will be created using this image to execute the Action.
All actions with same image will share the same container.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a list of environment variables that will be injected into the container.
These variables enable the container to adapt its behavior based on the environment it&rsquo;s running in.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>command</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the command to be executed inside the container.
The working directory for this command is the container&rsquo;s root directory(&lsquo;/&rsquo;).
Commands are executed directly without a shell environment, meaning shell-specific syntax (&lsquo;|&rsquo;, etc.) is not supported.
If the shell is required, it must be explicitly invoked in the command.</p>
<p>A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.</p>
</td>
</tr>
<tr>
<td>
<code>args</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Args represents the arguments that are passed to the <code>command</code> for execution.</p>
</td>
</tr>
<tr>
<td>
<code>targetPodSelector</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TargetPodSelector">
TargetPodSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the criteria used to select the target Pod(s) for executing the Action.
This is useful when there is no default target replica identified.
It allows for precise control over which Pod(s) the Action should run in.</p>
<p>If not specified, the Action will be executed in the pod where the Action is triggered, such as the pod
to be removed or added; or a random pod if the Action is triggered at the component level, such as
post-provision or pre-terminate of the component.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>matchingKey</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used in conjunction with the <code>targetPodSelector</code> field to refine the selection of target pod(s) for Action execution.
The impact of this field depends on the <code>targetPodSelector</code> value:</p>
<ul>
<li>When <code>targetPodSelector</code> is set to <code>Any</code> or <code>All</code>, this field will be ignored.</li>
<li>When <code>targetPodSelector</code> is set to <code>Role</code>, only those replicas whose role matches the <code>matchingKey</code>
will be selected for the Action.</li>
</ul>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>container</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the container within the same pod whose resources will be shared with the action.
This allows the action to utilize the specified container&rsquo;s resources without executing within it.</p>
<p>The name must match one of the containers defined in <code>componentDefinition.spec.runtime</code>.</p>
<p>The resources that can be shared are included:</p>
<ul>
<li>volume mounts</li>
</ul>
<p>This field cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Exporter">Exporter
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>containerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the built-in metrics exporter container.</p>
</td>
</tr>
<tr>
<td>
<code>scrapePath</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the http/https url path to scrape for metrics.
If empty, Prometheus uses the default value (e.g. <code>/metrics</code>).</p>
</td>
</tr>
<tr>
<td>
<code>scrapePort</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the port name to scrape for metrics.</p>
</td>
</tr>
<tr>
<td>
<code>scrapeScheme</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PrometheusScheme">
PrometheusScheme
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the schema to use for scraping.
<code>http</code> and <code>https</code> are the expected values unless you rewrite the <code>__scheme__</code> label via relabeling.
If empty, Prometheus uses the default value <code>http</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.HostNetwork">HostNetwork
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>containerPorts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.HostNetworkContainerPort">
[]HostNetworkContainerPort
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The list of container ports that are required by the component.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.HostNetworkContainerPort">HostNetworkContainerPort
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.HostNetwork">HostNetwork</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>container</code><br/>
<em>
string
</em>
</td>
<td>
<p>Container specifies the target container within the Pod.</p>
</td>
</tr>
<tr>
<td>
<code>ports</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Ports are named container ports within the specified container.
These container ports must be defined in the container for proper port allocation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.HostNetworkVarSelector">HostNetworkVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>HostNetworkVarSelector selects a var from host-network resources.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The component to select from.</p>
</td>
</tr>
<tr>
<td>
<code>HostNetworkVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.HostNetworkVars">
HostNetworkVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>HostNetworkVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.HostNetworkVars">HostNetworkVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.HostNetworkVarSelector">HostNetworkVarSelector</a>)
</p>
<div>
<p>HostNetworkVars defines the vars that can be referenced from host-network resources.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>container</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ContainerVars">
ContainerVars
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.InstanceTemplate">InstanceTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ShardTemplate">ShardTemplate</a>)
</p>
<div>
<p>InstanceTemplate allows customization of individual replica configurations in a Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name specifies the unique name of the instance Pod created using this InstanceTemplate.
This name is constructed by concatenating the Component&rsquo;s name, the template&rsquo;s name, and the instance&rsquo;s ordinal
using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0.
The name can&rsquo;t be empty.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this InstanceTemplate.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>canary</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicate whether the instances belonging to this template are canary instances.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of instances (Pods) to create from this InstanceTemplate.
This field allows setting how many replicated instances of the Component,
with the specific overrides in the InstanceTemplate, are created.
The default value is 1. A value of 0 disables instance creation.</p>
</td>
</tr>
<tr>
<td>
<code>ordinals</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Ordinals">
Ordinals
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the desired Ordinals of this InstanceTemplate.
The Ordinals used to specify the ordinal of the instance (pod) names to be generated under this InstanceTemplate.
If Ordinals are defined, their number must be equal to or more than the corresponding replicas.</p>
<p>For example, if Ordinals is &#123;ranges: [&#123;start: 0, end: 1&#125;], discrete: [7]&#125;,
then the instance names generated under this InstanceTemplate would be
$(cluster.name)-$(component.name)-$(template.name)-0、$(cluster.name)-$(component.name)-$(template.name)-1 and
$(cluster.name)-$(component.name)-$(template.name)-7</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs to be merged into the Pod&rsquo;s existing annotations.
Existing keys will have their values overwritten, while new keys will be added to the annotations.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs that will be merged into the Pod&rsquo;s existing labels.
Values for existing keys will be overwritten, and new keys will be added.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the instance.
If defined, it will overwrite the scheduling policy defined in ClusterSpec and/or ClusterComponentSpec.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the resource requirements of the first container in the Pod.
This field allows for customizing resource allocation (CPU, memory, etc.) for the container.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines Env to override.
Add new or override existing envs.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimTemplate">
[]PersistentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the storage requirements of the instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.InstanceUpdateStrategy">InstanceUpdateStrategy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>, <a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>InstanceUpdateStrategy defines fine-grained control over the spec update process of all instances.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategyType">
InstanceUpdateStrategyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the type of the update strategy.
Default is RollingUpdate.</p>
</td>
</tr>
<tr>
<td>
<code>rollingUpdate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.RollingUpdate">
RollingUpdate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies how the rolling update should be applied.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.InstanceUpdateStrategyType">InstanceUpdateStrategyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategy">InstanceUpdateStrategy</a>)
</p>
<div>
<p>InstanceUpdateStrategyType is a string enumeration type that enumerates
all possible update strategies for the KubeBlocks controllers.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;OnDelete&#34;</p></td>
<td><p>OnDeleteStrategyType indicates that ordered rolling restarts are disabled. Instances are recreated
when they are manually deleted.</p>
</td>
</tr><tr><td><p>&#34;RollingUpdate&#34;</p></td>
<td><p>RollingUpdateStrategyType indicates that update will be
applied to all Instances with respect to the workload
ordering constraints.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Issuer">Issuer
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.TLSConfig">TLSConfig</a>)
</p>
<div>
<p>Issuer defines the TLS certificates issuer for the Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.IssuerName">
IssuerName
</a>
</em>
</td>
<td>
<p>The issuer for TLS certificates.
It only allows two enum values: <code>KubeBlocks</code> and <code>UserProvided</code>.</p>
<ul>
<li><code>KubeBlocks</code> indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used.</li>
<li><code>UserProvided</code> means that the user is responsible for providing their own CA, Cert, and Key.
In this case, the user-provided CA certificate, server certificate, and private key will be used
for TLS communication.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>secretRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TLSSecretRef">
TLSSecretRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SecretRef is the reference to the secret that contains user-provided certificates.
It is required when the issuer is set to <code>UserProvided</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.IssuerName">IssuerName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Issuer">Issuer</a>)
</p>
<div>
<p>IssuerName defines the name of the TLS certificates issuer.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;KubeBlocks&#34;</p></td>
<td><p>IssuerKubeBlocks represents certificates that are signed by the KubeBlocks Operator.</p>
</td>
</tr><tr><td><p>&#34;UserProvided&#34;</p></td>
<td><p>IssuerUserProvided indicates that the user has provided their own CA-signed certificates.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.LetterCase">LetterCase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.PasswordConfig">PasswordConfig</a>)
</p>
<div>
<p>LetterCase defines the available cases to be used in password generation.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;LowerCases&#34;</p></td>
<td><p>LowerCases represents the use of lower case letters only.</p>
</td>
</tr><tr><td><p>&#34;MixedCases&#34;</p></td>
<td><p>MixedCases represents the use of a mix of both lower and upper case letters.</p>
</td>
</tr><tr><td><p>&#34;UpperCases&#34;</p></td>
<td><p>UpperCases represents the use of upper case letters only.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.LogConfig">LogConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies a descriptive label for the log type, such as &lsquo;slow&rsquo; for a MySQL slow log file.
It provides a clear identification of the log&rsquo;s purpose and content.</p>
</td>
</tr>
<tr>
<td>
<code>filePathPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the paths or patterns identifying where the log files are stored.
This field allows the system to locate and manage log files effectively.</p>
<p>Examples:</p>
<ul>
<li>/home/postgres/pgdata/pgroot/data/log/postgresql-*</li>
<li>/data/mysql/log/mysqld-error.log</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.MultipleClusterObjectCombinedOption">MultipleClusterObjectCombinedOption
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.MultipleClusterObjectOption">MultipleClusterObjectOption</a>)
</p>
<div>
<p>MultipleClusterObjectCombinedOption defines options for handling combined variables.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>newVarSuffix</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If set, the existing variable will be kept, and a new variable will be defined with the specified suffix
in pattern: $(var.name)_$(suffix).
The new variable will be auto-created and placed behind the existing one.
If not set, the existing variable will be reused with the value format defined below.</p>
</td>
</tr>
<tr>
<td>
<code>valueFormat</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.MultipleClusterObjectValueFormat">
MultipleClusterObjectValueFormat
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The format of the value that the operator will use to compose values from multiple components.</p>
</td>
</tr>
<tr>
<td>
<code>flattenFormat</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.MultipleClusterObjectValueFormatFlatten">
MultipleClusterObjectValueFormatFlatten
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The flatten format, default is: $(comp-name-1):value,$(comp-name-2):value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.MultipleClusterObjectOption">MultipleClusterObjectOption
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterObjectReference">ClusterObjectReference</a>)
</p>
<div>
<p>MultipleClusterObjectOption defines the options for handling multiple cluster objects matched.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>requireAllComponentObjects</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>RequireAllComponentObjects controls whether all component objects must exist before resolving.
If set to true, resolving will only proceed if all component objects are present.</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.MultipleClusterObjectStrategy">
MultipleClusterObjectStrategy
</a>
</em>
</td>
<td>
<p>Define the strategy for handling multiple cluster objects.</p>
</td>
</tr>
<tr>
<td>
<code>combinedOption</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.MultipleClusterObjectCombinedOption">
MultipleClusterObjectCombinedOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Define the options for handling combined variables.
Valid only when the strategy is set to &ldquo;combined&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.MultipleClusterObjectStrategy">MultipleClusterObjectStrategy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.MultipleClusterObjectOption">MultipleClusterObjectOption</a>)
</p>
<div>
<p>MultipleClusterObjectStrategy defines the strategy for handling multiple cluster objects.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;combined&#34;</p></td>
<td><p>MultipleClusterObjectStrategyCombined - the values from all matched components will be combined into a single
variable using the specified option.</p>
</td>
</tr><tr><td><p>&#34;individual&#34;</p></td>
<td><p>MultipleClusterObjectStrategyIndividual - each matched component will have its individual variable with its name
as the suffix.
This is required when referencing credential variables that cannot be passed by values.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.MultipleClusterObjectValueFormat">MultipleClusterObjectValueFormat
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.MultipleClusterObjectCombinedOption">MultipleClusterObjectCombinedOption</a>)
</p>
<div>
<p>MultipleClusterObjectValueFormat defines the format details for the value.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Flatten&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.MultipleClusterObjectValueFormatFlatten">MultipleClusterObjectValueFormatFlatten
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.MultipleClusterObjectCombinedOption">MultipleClusterObjectCombinedOption</a>)
</p>
<div>
<p>MultipleClusterObjectValueFormatFlatten defines the flatten format for the value.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>delimiter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Pair delimiter.</p>
</td>
</tr>
<tr>
<td>
<code>keyValueDelimiter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key-value delimiter.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.NamedVar">NamedVar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ContainerVars">ContainerVars</a>, <a href="#apps.kubeblocks.io/v1.ResourceVars">ResourceVars</a>, <a href="#apps.kubeblocks.io/v1.ServiceVars">ServiceVars</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>option</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Ordinals">Ordinals
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.InstanceTemplate">InstanceTemplate</a>, <a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>, <a href="#workloads.kubeblocks.io/v1.InstanceTemplate">InstanceTemplate</a>)
</p>
<div>
<p>Ordinals represents a combination of continuous segments and individual values.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ranges</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Range">
[]Range
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>discrete</code><br/>
<em>
[]int32
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.PasswordConfig">PasswordConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentSystemAccount">ComponentSystemAccount</a>, <a href="#apps.kubeblocks.io/v1.SystemAccount">SystemAccount</a>)
</p>
<div>
<p>PasswordConfig helps provide to customize complexity of password generation pattern.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>length</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The length of the password.</p>
</td>
</tr>
<tr>
<td>
<code>numDigits</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of digits in the password.</p>
</td>
</tr>
<tr>
<td>
<code>numSymbols</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of symbols in the password.</p>
</td>
</tr>
<tr>
<td>
<code>symbolCharacters</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The set of symbols allowed when generating password. If empty, kubeblocks will
use a default symbol set, which is &ldquo;!@#&amp;*&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>letterCase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.LetterCase">
LetterCase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The case of the letters in the password.</p>
</td>
</tr>
<tr>
<td>
<code>seed</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Seed to generate the account&rsquo;s password.
Cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicy">PersistentVolumeClaimRetentionPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>, <a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>PersistentVolumeClaimRetentionPolicy describes the policy used for PVCs created from the VolumeClaimTemplates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>whenDeleted</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicyType">
PersistentVolumeClaimRetentionPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>WhenDeleted specifies what happens to PVCs created from VolumeClaimTemplates when the workload is deleted.
The <code>Retain</code> policy causes PVCs to not be affected by workload deletion.
The default policy of <code>Delete</code> causes those PVCs to be deleted.</p>
</td>
</tr>
<tr>
<td>
<code>whenScaled</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicyType">
PersistentVolumeClaimRetentionPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>WhenScaled specifies what happens to PVCs created from VolumeClaimTemplates when the workload is scaled down.
The <code>Retain</code> policy causes PVCs to not be affected by a scale down.
The default policy of <code>Delete</code> causes the associated PVCs for pods scaled down to be deleted.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicyType">PersistentVolumeClaimRetentionPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicy">PersistentVolumeClaimRetentionPolicy</a>)
</p>
<div>
<p>PersistentVolumeClaimRetentionPolicyType is a string enumeration of the policies that will determine
when volumes from the VolumeClaimTemplates will be deleted when the controlling StatefulSet is
deleted or scaled down.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Delete&#34;</p></td>
<td><p>DeletePersistentVolumeClaimRetentionPolicyType specifies that PersistentVolumeClaims associated with
VolumeClaimTemplates will be deleted in the scenario specified in PersistentVolumeClaimRetentionPolicy.</p>
</td>
</tr><tr><td><p>&#34;Retain&#34;</p></td>
<td><p>RetainPersistentVolumeClaimRetentionPolicyType is the default PersistentVolumeClaimRetentionPolicy
and specifies that PersistentVolumeClaims associated with VolumeClaimTemplates will not be deleted.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.PersistentVolumeClaimTemplate">PersistentVolumeClaimTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.InstanceTemplate">InstanceTemplate</a>, <a href="#apps.kubeblocks.io/v1.ShardTemplate">ShardTemplate</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Refers to the name of a volumeMount defined in either:</p>
<ul>
<li><code>componentDefinition.spec.runtime.containers[*].volumeMounts</code></li>
</ul>
<p>The value of <code>name</code> must match the <code>name</code> field of a volumeMount specified in the corresponding <code>volumeMounts</code> array.</p>
</td>
</tr>
<tr>
<td>
<code>persistentVolumeClaimName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the prefix of the PVC name for the volume.</p>
<p>For each replica, the final name of the PVC will be in format: <persistentVolumeClaimName>-<ordinal></p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the labels for the PVC of the volume.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the annotations for the PVC of the volume.</p>
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeclaimspec-v1-core">
Kubernetes core/v1.PersistentVolumeClaimSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume
with the mount name specified in the <code>name</code> field.</p>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>accessModes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeaccessmode-v1-core">
[]Kubernetes core/v1.PersistentVolumeAccessMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>accessModes contains the desired access modes the volume should have.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1</a></p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>selector is a label query over volumes to consider for binding.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volumeresourcerequirements-v1-core">
Kubernetes core/v1.VolumeResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>resources represents the minimum resources the volume should have.
If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements
that are lower than previous value but must still be higher than capacity recorded in the
status field of the claim.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources">https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources</a></p>
</td>
</tr>
<tr>
<td>
<code>volumeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>volumeName is the binding reference to the PersistentVolume backing this claim.</p>
</td>
</tr>
<tr>
<td>
<code>storageClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>storageClassName is the name of the StorageClass required by the claim.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1</a></p>
</td>
</tr>
<tr>
<td>
<code>volumeMode</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumemode-v1-core">
Kubernetes core/v1.PersistentVolumeMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>volumeMode defines what type of volume is required by the claim.
Value of Filesystem is implied when not included in claim spec.</p>
</td>
</tr>
<tr>
<td>
<code>dataSource</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#typedlocalobjectreference-v1-core">
Kubernetes core/v1.TypedLocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>dataSource field can be used to specify either:
* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
* An existing PVC (PersistentVolumeClaim)
If the provisioner or an external controller can support the specified data source,
it will create a new volume based on the contents of the specified data source.
When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,
and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.
If the namespace is specified, then dataSourceRef will not be copied to dataSource.</p>
</td>
</tr>
<tr>
<td>
<code>dataSourceRef</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#typedobjectreference-v1-core">
Kubernetes core/v1.TypedObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>dataSourceRef specifies the object from which to populate the volume with data, if a non-empty
volume is desired. This may be any object from a non-empty API group (non
core object) or a PersistentVolumeClaim object.
When this field is specified, volume binding will only succeed if the type of
the specified object matches some installed volume populator or dynamic
provisioner.
This field will replace the functionality of the dataSource field and as such
if both fields are non-empty, they must have the same value. For backwards
compatibility, when namespace isn&rsquo;t specified in dataSourceRef,
both fields (dataSource and dataSourceRef) will be set to the same
value automatically if one of them is empty and the other is non-empty.
When namespace is specified in dataSourceRef,
dataSource isn&rsquo;t set to the same value and must be empty.
There are three important differences between dataSource and dataSourceRef:
* While dataSource only allows two specific types of objects, dataSourceRef
  allows any non-core object, as well as PersistentVolumeClaim objects.
* While dataSource ignores disallowed values (dropping them), dataSourceRef
  preserves all values, and generates an error if a disallowed value is
  specified.
* While dataSource only allows local objects, dataSourceRef allows objects
  in any namespaces.
(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.
(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.</p>
</td>
</tr>
<tr>
<td>
<code>volumeAttributesClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.
If specified, the CSI driver will create or update the volume with the attributes defined
in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,
it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass
will be applied to the claim but it&rsquo;s not allowed to reset this field to empty string once it is set.
If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass
will be set by the persistentvolume controller if it exists.
If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be
set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource
exists.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass">https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass</a>
(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Phase">Phase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterDefinitionStatus">ClusterDefinitionStatus</a>, <a href="#apps.kubeblocks.io/v1.ComponentDefinitionStatus">ComponentDefinitionStatus</a>, <a href="#apps.kubeblocks.io/v1.ComponentVersionStatus">ComponentVersionStatus</a>, <a href="#apps.kubeblocks.io/v1.ServiceDescriptorStatus">ServiceDescriptorStatus</a>, <a href="#apps.kubeblocks.io/v1.ShardingDefinitionStatus">ShardingDefinitionStatus</a>, <a href="#apps.kubeblocks.io/v1.SidecarDefinitionStatus">SidecarDefinitionStatus</a>)
</p>
<div>
<p>Phase represents the status of a CR.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Available&#34;</p></td>
<td><p>AvailablePhase indicates that a CR is in an available state.</p>
</td>
</tr><tr><td><p>&#34;Unavailable&#34;</p></td>
<td><p>UnavailablePhase indicates that a CR is in an unavailable state.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.PodUpdatePolicyType">PodUpdatePolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>, <a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>PodUpdatePolicyType indicates how pods should be updated</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;PreferInPlace&#34;</p></td>
<td><p>PreferInPlacePodUpdatePolicyType indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.</p>
</td>
</tr><tr><td><p>&#34;StrictInPlace&#34;</p></td>
<td><p>StrictInPlacePodUpdatePolicyType indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.PreConditionType">PreConditionType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Action">Action</a>)
</p>
<div>
<p>PreConditionType defines the preCondition type of the action execution.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ClusterReady&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ComponentReady&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Immediately&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RuntimeReady&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Probe">Probe
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentLifecycleActions">ComponentLifecycleActions</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Action</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<p>
(Members of <code>Action</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>initialDelaySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of seconds to wait after the container has started before the RoleProbe
begins to detect the container&rsquo;s role.</p>
</td>
</tr>
<tr>
<td>
<code>periodSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the frequency at which the probe is conducted. This value is expressed in seconds.
Default to 60 seconds. Minimum value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>successThreshold</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Minimum value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>failureThreshold</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.PrometheusScheme">PrometheusScheme
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Exporter">Exporter</a>)
</p>
<div>
<p>PrometheusScheme defines the protocol of prometheus scrape metrics.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;https&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ProvisionSecretRef">ProvisionSecretRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentSystemAccount">ComponentSystemAccount</a>)
</p>
<div>
<p>ProvisionSecretRef represents the reference to a secret.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The unique identifier of the secret.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>The namespace where the secret is located.</p>
</td>
</tr>
<tr>
<td>
<code>password</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The key in the secret data that contains the password.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Range">Range
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Ordinals">Ordinals</a>)
</p>
<div>
<p>Range represents a range with a start and an end value. Both start and end are included.
It is used to define a continuous segment.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>start</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>end</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ReplicaRole">ReplicaRole
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>, <a href="#workloads.kubeblocks.io/v1.MemberStatus">MemberStatus</a>)
</p>
<div>
<p>ReplicaRole represents a role that can be assigned to a component instance, defining its behavior and responsibilities.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name defines the role&rsquo;s unique identifier. This value is used to set the &ldquo;apps.kubeblocks.io/role&rdquo; label
on the corresponding object to identify its role.</p>
<p>For example, common role names include:
- &ldquo;leader&rdquo;: The primary/master instance that handles write operations
- &ldquo;follower&rdquo;: Secondary/replica instances that replicate data from the leader
- &ldquo;learner&rdquo;: Read-only instances that don&rsquo;t participate in elections</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>updatePriority</code><br/>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>UpdatePriority determines the order in which pods with different roles are updated.
Pods are sorted by this priority (higher numbers = higher priority) and updated accordingly.
Roles with the highest priority will be updated last.
The default priority is 0.</p>
<p>For example:
- Leader role may have priority 2 (updated last)
- Follower role may have priority 1 (updated before leader)
- Learner role may have priority 0 (updated first)</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>participatesInQuorum</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>ParticipatesInQuorum indicates if pods with this role are counted when determining quorum.
This affects update strategies that need to maintain quorum for availability. Roles participate
in quorum should have higher update priority than roles do not participate in quorum.
The default value is false.</p>
<p>For example, in a 5-pod component where:
- 2 learner pods (participatesInQuorum=false)
- 2 follower pods (participatesInQuorum=true)
- 1 leader pod (participatesInQuorum=true)
The quorum size would be 3 (based on the 3 participating pods), allowing parallel updates
of 2 learners and 1 follower while maintaining quorum.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ReplicasLimit">ReplicasLimit
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ReplicasLimit defines the valid range of number of replicas supported.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The minimum limit of replicas.</p>
</td>
</tr>
<tr>
<td>
<code>maxReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The maximum limit of replicas.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ResourceVarSelector">ResourceVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>ResourceVarSelector selects a var from a kind of resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Component to select from.</p>
</td>
</tr>
<tr>
<td>
<code>ResourceVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ResourceVars">
ResourceVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ResourceVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ResourceVars">ResourceVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ResourceVarSelector">ResourceVarSelector</a>)
</p>
<div>
<p>ResourceVars defines the vars that can be referenced from resources.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cpu</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>cpuLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>memory</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>memoryLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.NamedVar">
NamedVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.RetryPolicy">RetryPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Action">Action</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxRetries</code><br/>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the maximum number of retry attempts that should be made for a given Action.
This value is set to 0 by default, indicating that no retries will be made.</p>
</td>
</tr>
<tr>
<td>
<code>retryInterval</code><br/>
<em>
time.Duration
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the duration of time to wait between each retry attempt.
This value is set to 0 by default, indicating that there will be no delay between retry attempts.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.RoledVar">RoledVar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentVars">ComponentVars</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>role</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>option</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.RollingUpdate">RollingUpdate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategy">InstanceUpdateStrategy</a>)
</p>
<div>
<p>RollingUpdate specifies how the rolling update should be applied.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>replicas</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the number of instances that should be updated during a rolling update.
The remaining instances will remain untouched. This is helpful in defining how many instances
should participate in the update process.
Value can be an absolute number (ex: 5) or a percentage of desired instances (ex: 10%).
Absolute number is calculated from percentage by rounding up.
The default value is ComponentSpec.Replicas (i.e., update all instances).</p>
</td>
</tr>
<tr>
<td>
<code>maxUnavailable</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The maximum number of instances that can be unavailable during the update.
Value can be an absolute number (ex: 5) or a percentage of desired instances (ex: 10%).
Absolute number is calculated from percentage by rounding up. This can not be 0.
Defaults to 1. The field applies to all instances. That means if there is any unavailable pod,
it will be counted towards MaxUnavailable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.SchedulingPolicy">SchedulingPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ClusterSpec">ClusterSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.InstanceTemplate">InstanceTemplate</a>, <a href="#apps.kubeblocks.io/v1.ShardTemplate">ShardTemplate</a>, <a href="#workloads.kubeblocks.io/v1.InstanceTemplate">InstanceTemplate</a>)
</p>
<div>
<p>SchedulingPolicy defines the scheduling policy for instances.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>schedulerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the Pod will be dispatched by specified scheduler.
If not specified, the Pod will be dispatched by default scheduler.</p>
</td>
</tr>
<tr>
<td>
<code>nodeSelector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeSelector is a selector which must be true for the Pod to fit on a node.
Selector which must match a node&rsquo;s labels for the Pod to be scheduled on that node.
More info: <a href="https://kubernetes.io/docs/concepts/configuration/assign-pod-node/">https://kubernetes.io/docs/concepts/configuration/assign-pod-node/</a></p>
</td>
</tr>
<tr>
<td>
<code>nodeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,
the scheduler simply schedules this Pod onto that node, assuming that it fits resource
requirements.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#affinity-v1-core">
Kubernetes core/v1.Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows Pods to be scheduled onto nodes with matching taints.
Each toleration in the array allows the Pod to tolerate node taints based on
specified <code>key</code>, <code>value</code>, <code>effect</code>, and <code>operator</code>.</p>
<ul>
<li>The <code>key</code>, <code>value</code>, and <code>effect</code> identify the taint that the toleration matches.</li>
<li>The <code>operator</code> determines how the toleration matches the taint.</li>
</ul>
<p>Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.</p>
</td>
</tr>
<tr>
<td>
<code>topologySpreadConstraints</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#topologyspreadconstraint-v1-core">
[]Kubernetes core/v1.TopologySpreadConstraint
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TopologySpreadConstraints describes how a group of Pods ought to spread across topology
domains. Scheduler will schedule Pods in a way which abides by the constraints.
All topologySpreadConstraints are ANDed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Service">Service
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterService">ClusterService</a>, <a href="#apps.kubeblocks.io/v1.ComponentService">ComponentService</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name defines the name of the service.
otherwise, it indicates the name of the service.
Others can refer to this service by its name. (e.g., connection credential)
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>serviceName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceName defines the name of the underlying service object.
If not specified, the default service name with different patterns will be used:</p>
<ul>
<li>CLUSTER_NAME: for cluster-level services</li>
<li>CLUSTER_NAME-COMPONENT_NAME: for component-level services</li>
</ul>
<p>Only one default service name is allowed.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If ServiceType is LoadBalancer, cloud provider related parameters can be put here
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer">https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer</a>.</p>
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#servicespec-v1-core">
Kubernetes core/v1.ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Spec defines the behavior of a service.
<a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</a></p>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>ports</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceport-v1-core">
[]Kubernetes core/v1.ServicePort
</a>
</em>
</td>
<td>
<p>The list of ports that are exposed by this service.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Route service traffic to pods with label keys and values matching this
selector. If empty or not present, the service is assumed to have an
external process managing its endpoints, which Kubernetes will not
modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.
Ignored if type is ExternalName.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/">https://kubernetes.io/docs/concepts/services-networking/service/</a></p>
</td>
</tr>
<tr>
<td>
<code>clusterIP</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>clusterIP is the IP address of the service and is usually assigned
randomly. If an address is specified manually, is in-range (as per
system configuration), and is not in use, it will be allocated to the
service; otherwise creation of the service will fail. This field may not
be changed through updates unless the type field is also being changed
to ExternalName (which requires this field to be blank) or the type
field is being changed from ExternalName (in which case this field may
optionally be specified, as describe above).  Valid values are &ldquo;None&rdquo;,
empty string (&ldquo;&rdquo;), or a valid IP address. Setting this to &ldquo;None&rdquo; makes a
&ldquo;headless service&rdquo; (no virtual IP), which is useful when direct endpoint
connections are preferred and proxying is not required.  Only applies to
types ClusterIP, NodePort, and LoadBalancer. If this field is specified
when creating a Service of type ExternalName, creation will fail. This
field will be wiped when updating a Service to type ExternalName.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>clusterIPs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ClusterIPs is a list of IP addresses assigned to this service, and are
usually assigned randomly.  If an address is specified manually, is
in-range (as per system configuration), and is not in use, it will be
allocated to the service; otherwise creation of the service will fail.
This field may not be changed through updates unless the type field is
also being changed to ExternalName (which requires this field to be
empty) or the type field is being changed from ExternalName (in which
case this field may optionally be specified, as describe above).  Valid
values are &ldquo;None&rdquo;, empty string (&ldquo;&rdquo;), or a valid IP address.  Setting
this to &ldquo;None&rdquo; makes a &ldquo;headless service&rdquo; (no virtual IP), which is
useful when direct endpoint connections are preferred and proxying is
not required.  Only applies to types ClusterIP, NodePort, and
LoadBalancer. If this field is specified when creating a Service of type
ExternalName, creation will fail. This field will be wiped when updating
a Service to type ExternalName.  If this field is not specified, it will
be initialized from the clusterIP field.  If this field is specified,
clients must ensure that clusterIPs[0] and clusterIP have the same
value.</p>
<p>This field may hold a maximum of two entries (dual-stack IPs, in either order).
These IPs must correspond to the values of the ipFamilies field. Both
clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#servicetype-v1-core">
Kubernetes core/v1.ServiceType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>type determines how the Service is exposed. Defaults to ClusterIP. Valid
options are ExternalName, ClusterIP, NodePort, and LoadBalancer.
&ldquo;ClusterIP&rdquo; allocates a cluster-internal IP address for load-balancing
to endpoints. Endpoints are determined by the selector or if that is not
specified, by manual construction of an Endpoints object or
EndpointSlice objects. If clusterIP is &ldquo;None&rdquo;, no virtual IP is
allocated and the endpoints are published as a set of endpoints rather
than a virtual IP.
&ldquo;NodePort&rdquo; builds on ClusterIP and allocates a port on every node which
routes to the same endpoints as the clusterIP.
&ldquo;LoadBalancer&rdquo; builds on NodePort and creates an external load-balancer
(if supported in the current cloud) which routes to the same endpoints
as the clusterIP.
&ldquo;ExternalName&rdquo; aliases this service to the specified externalName.
Several other fields do not apply to ExternalName services.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types">https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types</a></p>
</td>
</tr>
<tr>
<td>
<code>externalIPs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>externalIPs is a list of IP addresses for which nodes in the cluster
will also accept traffic for this service.  These IPs are not managed by
Kubernetes.  The user is responsible for ensuring that traffic arrives
at a node with this IP.  A common example is external load-balancers
that are not part of the Kubernetes system.</p>
</td>
</tr>
<tr>
<td>
<code>sessionAffinity</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceaffinity-v1-core">
Kubernetes core/v1.ServiceAffinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Supports &ldquo;ClientIP&rdquo; and &ldquo;None&rdquo;. Used to maintain session affinity.
Enable client IP based session affinity.
Must be ClientIP or None.
Defaults to None.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerIP</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Only applies to Service Type: LoadBalancer.
This feature depends on whether the underlying cloud-provider supports specifying
the loadBalancerIP when a load balancer is created.
This field will be ignored if the cloud-provider does not support the feature.
Deprecated: This field was under-specified and its meaning varies across implementations.
Using it is non-portable and it may not support dual-stack.
Users are encouraged to use implementation-specific annotations when available.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSourceRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified and supported by the platform, this will restrict traffic through the cloud-provider
load-balancer will be restricted to the specified client IPs. This field will be ignored if the
cloud-provider does not support the feature.&rdquo;
More info: <a href="https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/">https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/</a></p>
</td>
</tr>
<tr>
<td>
<code>externalName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>externalName is the external reference that discovery mechanisms will
return as an alias for this service (e.g. a DNS CNAME record). No
proxying will be involved.  Must be a lowercase RFC-1123 hostname
(<a href="https://tools.ietf.org/html/rfc1123">https://tools.ietf.org/html/rfc1123</a>) and requires <code>type</code> to be &ldquo;ExternalName&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>externalTrafficPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceexternaltrafficpolicy-v1-core">
Kubernetes core/v1.ServiceExternalTrafficPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>externalTrafficPolicy describes how nodes distribute service traffic they
receive on one of the Service&rsquo;s &ldquo;externally-facing&rdquo; addresses (NodePorts,
ExternalIPs, and LoadBalancer IPs). If set to &ldquo;Local&rdquo;, the proxy will configure
the service in a way that assumes that external load balancers will take care
of balancing the service traffic between nodes, and so each node will deliver
traffic only to the node-local endpoints of the service, without masquerading
the client source IP. (Traffic mistakenly sent to a node with no endpoints will
be dropped.) The default value, &ldquo;Cluster&rdquo;, uses the standard behavior of
routing to all endpoints evenly (possibly modified by topology and other
features). Note that traffic sent to an External IP or LoadBalancer IP from
within the cluster will always get &ldquo;Cluster&rdquo; semantics, but clients sending to
a NodePort from within the cluster may need to take traffic policy into account
when picking a node.</p>
</td>
</tr>
<tr>
<td>
<code>healthCheckNodePort</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>healthCheckNodePort specifies the healthcheck nodePort for the service.
This only applies when type is set to LoadBalancer and
externalTrafficPolicy is set to Local. If a value is specified, is
in-range, and is not in use, it will be used.  If not specified, a value
will be automatically allocated.  External systems (e.g. load-balancers)
can use this port to determine if a given node holds endpoints for this
service or not.  If this field is specified when creating a Service
which does not need it, creation will fail. This field will be wiped
when updating a Service to no longer need it (e.g. changing type).
This field cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>publishNotReadyAddresses</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>publishNotReadyAddresses indicates that any agent which deals with endpoints for this
Service should disregard any indications of ready/not-ready.
The primary use case for setting this field is for a StatefulSet&rsquo;s Headless Service to
propagate SRV DNS records for its Pods for the purpose of peer discovery.
The Kubernetes controllers that generate Endpoints and EndpointSlice resources for
Services interpret this to mean that all endpoints are considered &ldquo;ready&rdquo; even if the
Pods themselves are not. Agents which consume only Kubernetes generated endpoints
through the Endpoints or EndpointSlice resources can safely assume this behavior.</p>
</td>
</tr>
<tr>
<td>
<code>sessionAffinityConfig</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#sessionaffinityconfig-v1-core">
Kubernetes core/v1.SessionAffinityConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>sessionAffinityConfig contains the configurations of session affinity.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilies</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#ipfamily-v1-core">
[]Kubernetes core/v1.IPFamily
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this
service. This field is usually assigned automatically based on cluster
configuration and the ipFamilyPolicy field. If this field is specified
manually, the requested family is available in the cluster,
and ipFamilyPolicy allows it, it will be used; otherwise creation of
the service will fail. This field is conditionally mutable: it allows
for adding or removing a secondary IP family, but it does not allow
changing the primary IP family of the Service. Valid values are &ldquo;IPv4&rdquo;
and &ldquo;IPv6&rdquo;.  This field only applies to Services of types ClusterIP,
NodePort, and LoadBalancer, and does apply to &ldquo;headless&rdquo; services.
This field will be wiped when updating a Service to type ExternalName.</p>
<p>This field may hold a maximum of two entries (dual-stack families, in
either order).  These families must correspond to the values of the
clusterIPs field, if specified. Both clusterIPs and ipFamilies are
governed by the ipFamilyPolicy field.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilyPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#ipfamilypolicy-v1-core">
Kubernetes core/v1.IPFamilyPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPFamilyPolicy represents the dual-stack-ness requested or required by
this Service. If there is no value provided, then this field will be set
to SingleStack. Services can be &ldquo;SingleStack&rdquo; (a single IP family),
&ldquo;PreferDualStack&rdquo; (two IP families on dual-stack configured clusters or
a single IP family on single-stack clusters), or &ldquo;RequireDualStack&rdquo;
(two IP families on dual-stack configured clusters, otherwise fail). The
ipFamilies and clusterIPs fields depend on the value of this field. This
field will be wiped when updating a service to type ExternalName.</p>
</td>
</tr>
<tr>
<td>
<code>allocateLoadBalancerNodePorts</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>allocateLoadBalancerNodePorts defines if NodePorts will be automatically
allocated for services with type LoadBalancer.  Default is &ldquo;true&rdquo;. It
may be set to &ldquo;false&rdquo; if the cluster load-balancer does not rely on
NodePorts.  If the caller requests specific NodePorts (by specifying a
value), those requests will be respected, regardless of this field.
This field may only be set for services with type LoadBalancer and will
be cleared if the type is changed to any other type.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerClass</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>loadBalancerClass is the class of the load balancer implementation this Service belongs to.
If specified, the value of this field must be a label-style identifier, with an optional prefix,
e.g. &ldquo;internal-vip&rdquo; or &ldquo;example.com/internal-vip&rdquo;. Unprefixed names are reserved for end-users.
This field can only be set when the Service type is &lsquo;LoadBalancer&rsquo;. If not set, the default load
balancer implementation is used, today this is typically done through the cloud provider integration,
but should apply for any default implementation. If set, it is assumed that a load balancer
implementation is watching for Services with a matching class. Any default load balancer
implementation (e.g. cloud providers) should ignore Services that set this field.
This field can only be set when creating or updating a Service to type &lsquo;LoadBalancer&rsquo;.
Once set, it can not be changed. This field will be wiped when a service is updated to a non &lsquo;LoadBalancer&rsquo; type.</p>
</td>
</tr>
<tr>
<td>
<code>internalTrafficPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceinternaltrafficpolicy-v1-core">
Kubernetes core/v1.ServiceInternalTrafficPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InternalTrafficPolicy describes how nodes distribute service traffic they
receive on the ClusterIP. If set to &ldquo;Local&rdquo;, the proxy will assume that pods
only want to talk to endpoints of the service on the same node as the pod,
dropping the traffic if there are no local endpoints. The default value,
&ldquo;Cluster&rdquo;, uses the standard behavior of routing to all endpoints evenly
(possibly modified by topology and other features).</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>roleSelector</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extends the above <code>serviceSpec.selector</code> by allowing you to specify defined role as selector for the service.
When <code>roleSelector</code> is set, it adds a label selector &ldquo;kubeblocks.io/role: &#123;roleSelector&#125;&rdquo;
to the <code>serviceSpec.selector</code>.
Example usage:</p>
<pre><code>  roleSelector: &quot;leader&quot;
</code></pre>
<p>In this example, setting <code>roleSelector</code> to &ldquo;leader&rdquo; will add a label selector
&ldquo;kubeblocks.io/role: leader&rdquo; to the <code>serviceSpec.selector</code>.
This means that the service will select and route traffic to Pods with the label
&ldquo;kubeblocks.io/role&rdquo; set to &ldquo;leader&rdquo;.</p>
<p>Note that if <code>podService</code> sets to true, RoleSelector will be ignored.
The <code>podService</code> flag takes precedence over <code>roleSelector</code> and generates a service for each Pod.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceDescriptorSpec">ServiceDescriptorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceDescriptor">ServiceDescriptor</a>)
</p>
<div>
<p>ServiceDescriptorSpec defines the desired state of ServiceDescriptor</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the type of database service provided by the external service.
For example, &ldquo;mysql&rdquo;, &ldquo;redis&rdquo;, &ldquo;mongodb&rdquo;.
This field categorizes databases by their functionality, protocol and compatibility, facilitating appropriate
service integration based on their unique capabilities.</p>
<p>This field is case-insensitive.</p>
<p>It also supports abbreviations for some well-known databases:
- &ldquo;pg&rdquo;, &ldquo;pgsql&rdquo;, &ldquo;postgres&rdquo;, &ldquo;postgresql&rdquo;: PostgreSQL service
- &ldquo;zk&rdquo;, &ldquo;zookeeper&rdquo;: ZooKeeper service
- &ldquo;es&rdquo;, &ldquo;elasticsearch&rdquo;: Elasticsearch service
- &ldquo;mongo&rdquo;, &ldquo;mongodb&rdquo;: MongoDB service
- &ldquo;ch&rdquo;, &ldquo;clickhouse&rdquo;: ClickHouse service</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the version of the service provided by the external service.
This is crucial for ensuring compatibility between different components of the system,
as different versions of a service may have varying features.</p>
</td>
</tr>
<tr>
<td>
<code>endpoint</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the endpoint of the external service.</p>
<p>If the service is exposed via a cluster, the endpoint will be provided in the format of <code>host:port</code>.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the service or IP address of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the port of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>podFQDNs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the pod FQDNs of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ConnectionCredentialAuth">
ConnectionCredentialAuth
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the authentication credentials required for accessing an external service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceDescriptorStatus">ServiceDescriptorStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceDescriptor">ServiceDescriptor</a>)
</p>
<div>
<p>ServiceDescriptorStatus defines the observed state of ServiceDescriptor</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the generation number that has been processed by the controller.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the current lifecycle phase of the ServiceDescriptor. This can be either &lsquo;Available&rsquo; or &lsquo;Unavailable&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a human-readable explanation detailing the reason for the current phase of the ServiceConnectionCredential.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRef">ServiceRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the identifier of the service reference declaration.
It corresponds to the serviceRefDeclaration name defined in either:</p>
<ul>
<li><code>componentDefinition.spec.serviceRefDeclarations[*].name</code></li>
<li><code>clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name</code> (deprecated)</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object.
If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current
Cluster by default.</p>
</td>
</tr>
<tr>
<td>
<code>clusterServiceSelector</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefClusterSelector">
ServiceRefClusterSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>References a service provided by another KubeBlocks Cluster.
It specifies the ClusterService and the account credentials needed for access.
The <code>ServiceKind</code> and <code>ServiceVersion</code> specified in the service reference within the
ClusterDefinition are not validated when using this approach.</p>
<p>If both <code>clusterServiceSelector</code> and <code>serviceDescriptor</code> are specified, the <code>clusterServiceSelector</code> takes precedence.</p>
</td>
</tr>
<tr>
<td>
<code>serviceDescriptor</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.</p>
<p>When referencing a service provided by external sources, a ServiceDescriptor object is required to establish
the service binding.
The <code>serviceDescriptor.spec.serviceKind</code> and <code>serviceDescriptor.spec.serviceVersion</code> should match the serviceKind
and serviceVersion declared in the definition.</p>
<p>If both <code>clusterServiceSelector</code> and <code>serviceDescriptor</code> are specified, the <code>clusterServiceSelector</code> takes precedence.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefClusterSelector">ServiceRefClusterSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceRef">ServiceRef</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cluster</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the Cluster being referenced.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefServiceSelector">
ServiceRefServiceSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Identifies a ClusterService from the list of Services defined in <code>cluster.spec.services</code> of the referenced Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>podFQDNs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefPodFQDNsSelector">
ServiceRefPodFQDNsSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>credential</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefCredentialSelector">
ServiceRefCredentialSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster.
The SystemAccount should be defined in <code>componentDefinition.spec.systemAccounts</code>
of the Component providing the service in the referenced Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefCredentialSelector">ServiceRefCredentialSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceRefClusterSelector">ServiceRefClusterSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>component</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the Component where the credential resides in.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the credential (SystemAccount) to reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefDeclaration">ServiceRefDeclaration
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ServiceRefDeclaration represents a reference to a service that can be either provided by a KubeBlocks Cluster
or an external service.
It acts as a placeholder for the actual service reference, which is determined later when a Cluster is created.</p>
<p>The purpose of ServiceRefDeclaration is to declare a service dependency without specifying the concrete details
of the service.
It allows for flexibility and abstraction in defining service references within a Component.
By using ServiceRefDeclaration, you can define service dependencies in a declarative manner, enabling loose coupling
and easier management of service references across different components and clusters.</p>
<p>Upon Cluster creation, the ServiceRefDeclaration is bound to an actual service through the ServiceRef field,
effectively resolving and connecting to the specified service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the ServiceRefDeclaration.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefDeclarationSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefDeclarationSpec">
[]ServiceRefDeclarationSpec
</a>
</em>
</td>
<td>
<p>Defines a list of constraints and requirements for services that can be bound to this ServiceRefDeclaration
upon Cluster creation.
Each ServiceRefDeclarationSpec defines a ServiceKind and ServiceVersion,
outlining the acceptable service types and versions that are compatible.</p>
<p>This flexibility allows a ServiceRefDeclaration to be fulfilled by any one of the provided specs.
For example, if it requires an OLTP database, specs for both MySQL and PostgreSQL are listed,
either MySQL or PostgreSQL services can be used when binding.</p>
</td>
</tr>
<tr>
<td>
<code>optional</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the service reference can be optional.</p>
<p>For an optional service-ref, the component can still be created even if the service-ref is not provided.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefDeclarationSpec">ServiceRefDeclarationSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceRefDeclaration">ServiceRefDeclaration</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the type or nature of the service. This should be a well-known application cluster type, such as
&#123;mysql, redis, mongodb&#125;.
The field is case-insensitive and supports abbreviations for some well-known databases.
For instance, both <code>zk</code> and <code>zookeeper</code> are considered as a ZooKeeper cluster, while <code>pg</code>, <code>postgres</code>, <code>postgresql</code>
are all recognized as a PostgreSQL cluster.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the service version of the service reference. This is a regular expression that matches a version number pattern.
For instance, <code>^8.0.8$</code>, <code>8.0.\d&#123;1,2&#125;$</code>, <code>^[v\-]*?(\d&#123;1,2&#125;\.)&#123;0,3&#125;\d&#123;1,2&#125;$</code> are all valid patterns.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefPodFQDNsSelector">ServiceRefPodFQDNsSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceRefClusterSelector">ServiceRefClusterSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>component</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the Component where the pods reside in.</p>
</td>
</tr>
<tr>
<td>
<code>role</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The role of the pods to reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefServiceSelector">ServiceRefServiceSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceRefClusterSelector">ServiceRefClusterSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>component</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the Component where the Service resides in.</p>
<p>It is required when referencing a Component&rsquo;s Service.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the Service to be referenced.</p>
<p>Leave it empty to reference the default Service. Set it to &ldquo;headless&rdquo; to reference the default headless Service.</p>
<p>If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,
and the resolved value will be presented in the following format: service1.name,service2.name&hellip;</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The port name of the Service to be referenced.</p>
<p>If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.</p>
<p>If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,
and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2&hellip;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefVarSelector">ServiceRefVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>ServiceRefVarSelector selects a var from a ServiceRefDeclaration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The ServiceRefDeclaration to select from.</p>
</td>
</tr>
<tr>
<td>
<code>ServiceRefVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefVars">
ServiceRefVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ServiceRefVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceRefVars">ServiceRefVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceRefVarSelector">ServiceRefVarSelector</a>)
</p>
<div>
<p>ServiceRefVars defines the vars that can be referenced from a ServiceRef.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>endpoint</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>podFQDNs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>CredentialVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVars">
CredentialVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>CredentialVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceVarSelector">ServiceVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>ServiceVarSelector selects a var from a Service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Service to select from.
It can be referenced from the default headless service by setting the name to &ldquo;headless&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>ServiceVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceVars">
ServiceVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ServiceVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ServiceVars">ServiceVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ServiceVarSelector">ServiceVarSelector</a>)
</p>
<div>
<p>ServiceVars defines the vars that can be referenced from a Service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceType</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceType references the type of the service.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>loadBalancer</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LoadBalancer represents the LoadBalancer ingress point of the service.</p>
<p>If multiple ingress points are available, the first one will be used automatically, choosing between IP and Hostname.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.NamedVar">
NamedVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Port references a port or node-port defined in the service.</p>
<p>If the referenced service is a pod-service, there will be multiple service objects matched,
and the value will be presented in the following format: service1.name:port1,service2.name:port2&hellip;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardTemplate">ShardTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterSharding">ClusterSharding</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The unique name of this ShardTemplate.</p>
<p>The name can&rsquo;t be empty.</p>
</td>
</tr>
<tr>
<td>
<code>shardingDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the ShardingDefinition custom resource (CR) that defines the sharding&rsquo;s characteristics and behavior.</p>
<p>The full name or regular expression is supported to match the ShardingDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>shards</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of shards to create from this ShardTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>shardIDs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the shard IDs to take over from the existing shards.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this template.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Labels to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Annotations to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines Env to override.
Add new or override existing envs.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the desired number of replicas for the shard which are created from this template.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the shard.
If defined, it will overwrite the scheduling policy defined in ClusterSpec and/or default template.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the resource requirements of the shard.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimTemplate">
[]PersistentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the storage requirements of the shard.</p>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the custom instances of the shard.</p>
</td>
</tr>
<tr>
<td>
<code>flatInstanceOrdinal</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the instance naming of the shard.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardingDefinitionSpec">ShardingDefinitionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ShardingDefinition">ShardingDefinition</a>)
</p>
<div>
<p>ShardingDefinitionSpec defines the desired state of ShardingDefinition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingTemplate">
ShardingTemplate
</a>
</em>
</td>
<td>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>shardsLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardsLimit">
ShardsLimit
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the upper limit of the number of shards supported by the sharding.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>provisionStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategy for provisioning shards of the sharding. Only <code>Serial</code> and <code>Parallel</code> are supported.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategy for updating shards of the sharding. Only <code>Serial</code> and <code>Parallel</code> are supported.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>lifecycleActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingLifecycleActions">
ShardingLifecycleActions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of hooks and procedures that customize the behavior of a sharding throughout its lifecycle.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingSystemAccount">
[]ShardingSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the system accounts for the sharding.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ShardingTLS">
ShardingTLS
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the TLS for the sharding.</p>
<p>This field is immutable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardingDefinitionStatus">ShardingDefinitionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ShardingDefinition">ShardingDefinition</a>)
</p>
<div>
<p>ShardingDefinitionStatus defines the observed state of ShardingDefinition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the most recent generation that has been observed for the ShardingDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the current status of the ShardingDefinition. Valid values include `<code>,</code>Available<code>, and</code>Unavailable<code>.
When the status is</code>Available`, the ShardingDefinition is ready and can be utilized by related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardingLifecycleActions">ShardingLifecycleActions
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ShardingDefinitionSpec">ShardingDefinitionSpec</a>)
</p>
<div>
<p>ShardingLifecycleActions defines a collection of Actions for customizing the behavior of a sharding.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>postProvision</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed after a sharding&rsquo;s creation.</p>
<p>By setting <code>postProvision.preCondition</code>, you can determine the specific lifecycle stage at which
the action should trigger, available conditions for sharding include: <code>Immediately</code>, <code>ComponentReady</code>,
and <code>ClusterReady</code>. For sharding, the <code>ComponentReady</code> condition means all components of the sharding are ready.</p>
<p>With <code>ComponentReady</code> being the default.</p>
<p>The PostProvision Action is intended to run only once.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>preTerminate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed prior to terminating a sharding.</p>
<p>The PreTerminate Action is intended to run only once.</p>
<p>This action is executed immediately when a terminate operation for the sharding is initiated.
The actual termination and cleanup of the sharding and its associated resources will not proceed
until the PreTerminate action has completed successfully.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>shardAdd</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed after a shard added.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>shardRemove</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed prior to remove a shard.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardingSystemAccount">ShardingSystemAccount
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ShardingDefinitionSpec">ShardingDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the system account defined in the sharding template.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>shared</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the account is shared across all shards in the sharding.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardingTLS">ShardingTLS
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ShardingDefinitionSpec">ShardingDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>shared</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the TLS configuration is shared across all shards in the sharding.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardingTemplate">ShardingTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ShardingDefinitionSpec">ShardingDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>The component definition(s) that the sharding is based on.</p>
<p>The component definition can be specified using one of the following:</p>
<ul>
<li>the full name</li>
<li>the regular expression pattern (&lsquo;^&rsquo; will be added to the beginning of the pattern automatically)</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.ShardsLimit">ShardsLimit
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ShardingDefinitionSpec">ShardingDefinitionSpec</a>)
</p>
<div>
<p>ShardsLimit defines the valid range of number of shards supported.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minShards</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The minimum limit of shards.</p>
</td>
</tr>
<tr>
<td>
<code>maxShards</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The maximum limit of shards.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.Sidecar">Sidecar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name specifies the unique name of the sidecar.</p>
<p>The name will be used as the name of the sidecar container in the Pod.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the exact component definition that the sidecar belongs to.</p>
<p>A sidecar will be updated when the owner component definition is updated only.</p>
</td>
</tr>
<tr>
<td>
<code>sidecarDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the sidecar definition CR to be used to create the sidecar.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.SidecarDefinitionSpec">SidecarDefinitionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.SidecarDefinition">SidecarDefinition</a>)
</p>
<div>
<p>SidecarDefinitionSpec defines the desired state of SidecarDefinition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the sidecar.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the component definition that the sidecar belongs to.</p>
<p>For a specific cluster object, if there is any components provided by the component definition of @owner,
the sidecar will be created and injected into the components which are provided by
the component definition of @selectors automatically.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>selectors</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Specifies the component definition of components that the sidecar along with.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>containers</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<p>List of containers for the sidecar.</p>
<p>Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.EnvVar">
[]EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines variables which are needed by the sidecar.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration file templates used by the Sidecar.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>scripts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentFileTemplate">
[]ComponentFileTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scripts used by the Sidecar.</p>
<p>This field is immutable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.SidecarDefinitionStatus">SidecarDefinitionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.SidecarDefinition">SidecarDefinition</a>)
</p>
<div>
<p>SidecarDefinitionStatus defines the observed state of SidecarDefinition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the most recent generation that has been observed for the SidecarDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the current status of the SidecarDefinition. Valid values include `<code>,</code>Available<code>, and</code>Unavailable<code>.
When the status is</code>Available`, the SidecarDefinition is ready and can be utilized by related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
<tr>
<td>
<code>owners</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Resolved owners of the SidecarDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>selectors</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Resolved selectors of the SidecarDefinition.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.SystemAccount">SystemAccount
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the unique identifier for the account. This name is used by other entities to reference the account.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>initAccount</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates if this account is a system initialization account (e.g., MySQL root).</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>statement</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SystemAccountStatement">
SystemAccountStatement
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the statements used to create, delete, and update the account.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>passwordGenerationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PasswordConfig">
PasswordConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the policy for generating the account&rsquo;s password.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.SystemAccountStatement">SystemAccountStatement
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.SystemAccount">SystemAccount</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>create</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The statement to create a new account with the necessary privileges.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>delete</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The statement to delete a account.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>update</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The statement to update an existing account.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.TLS">TLS
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>volumeName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the volume name for the TLS secret.
The controller will create a volume object with the specified name and add it to the pod when the TLS is enabled.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>mountPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the mount path for the TLS secret to be mounted.
Similar to the volume, the controller will mount the created volume to the specified path within containers when the TLS is enabled.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMode</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The permissions for the mounted path. Defaults to 0600.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>caFile</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The CA file of the TLS.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>certFile</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The certificate file of the TLS.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>keyFile</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The key file of the TLS.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.TLSConfig">TLSConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enable</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)
for secure communication.
When set to true, the Component will be configured to use TLS encryption for its network connections.
This ensures that the data transmitted between the Component and its clients or other Components is encrypted
and protected from unauthorized access.
If TLS is enabled, the Component may require additional configuration,
such as specifying TLS certificates and keys, to properly set up the secure communication channel.</p>
</td>
</tr>
<tr>
<td>
<code>issuer</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Issuer">
Issuer
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration for the TLS certificates issuer.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.
Required when TLS is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.TLSSecretRef">TLSSecretRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.Issuer">Issuer</a>)
</p>
<div>
<p>TLSSecretRef defines the Secret that contains TLS certs.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The namespace where the secret is located.
If not provided, the secret is assumed to be in the same namespace as the Cluster object.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the Secret that contains user-provided certificates.</p>
</td>
</tr>
<tr>
<td>
<code>ca</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key of CA cert in Secret</p>
</td>
</tr>
<tr>
<td>
<code>cert</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key of Cert in Secret</p>
</td>
</tr>
<tr>
<td>
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key of TLS private key in Secret</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.TLSVarSelector">TLSVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.VarSource">VarSource</a>)
</p>
<div>
<p>TLSVarSelector selects a var from the TLS.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Component to select from.</p>
</td>
</tr>
<tr>
<td>
<code>TLSVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TLSVars">
TLSVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>TLSVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.TLSVars">TLSVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.TLSVarSelector">TLSVarSelector</a>)
</p>
<div>
<p>TLSVars defines the vars that can be referenced from the TLS.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.TargetPodSelector">TargetPodSelector
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ExecAction">ExecAction</a>)
</p>
<div>
<p>TargetPodSelector defines how to select pod(s) to execute an Action.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Any&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ordinal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Role&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.TerminationPolicyType">TerminationPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterSpec">ClusterSpec</a>, <a href="#apps.kubeblocks.io/v1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
<p>TerminationPolicyType defines termination policy types.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Delete&#34;</p></td>
<td><p>Delete will delete all runtime resources belong to the cluster.</p>
</td>
</tr><tr><td><p>&#34;DoNotTerminate&#34;</p></td>
<td><p>DoNotTerminate will block delete operation.</p>
</td>
</tr><tr><td><p>&#34;WipeOut&#34;</p></td>
<td><p>WipeOut is based on Delete and wipe out all volume snapshots and snapshot data from backup storage location.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.UpdateStrategy">UpdateStrategy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#apps.kubeblocks.io/v1.ShardingDefinitionSpec">ShardingDefinitionSpec</a>)
</p>
<div>
<p>UpdateStrategy defines the update strategy for cluster components. This strategy determines how updates are applied
across the cluster.
The available strategies are <code>Serial</code>, <code>BestEffortParallel</code>, and <code>Parallel</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;BestEffortParallel&#34;</p></td>
<td><p>BestEffortParallelStrategy indicates that the replicas are updated in parallel, with the operator making
a best-effort attempt to update as many replicas as possible concurrently
while maintaining the component&rsquo;s availability.
Unlike the <code>Parallel</code> strategy, the <code>BestEffortParallel</code> strategy aims to ensure that a minimum number
of replicas remain available during the update process to maintain the component&rsquo;s quorum and functionality.</p>
<p>For example, consider a component with 5 replicas. To maintain the component&rsquo;s availability and quorum,
the operator may allow a maximum of 2 replicas to be simultaneously updated. This ensures that at least
3 replicas (a quorum) remain available and functional during the update process.</p>
<p>The <code>BestEffortParallel</code> strategy strikes a balance between update speed and component availability.</p>
</td>
</tr><tr><td><p>&#34;Parallel&#34;</p></td>
<td><p>ParallelStrategy indicates that updates are applied simultaneously to all Pods of a Component.
The replicas are updated in parallel, with the operator updating all replicas concurrently.
This strategy provides the fastest update time but may lead to a period of reduced availability or
capacity during the update process.</p>
</td>
</tr><tr><td><p>&#34;Serial&#34;</p></td>
<td><p>SerialStrategy indicates that updates are applied one at a time in a sequential manner.
The operator waits for each replica to be updated and ready before proceeding to the next one.
This ensures that only one replica is unavailable at a time during the update process.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1.VarOption">VarOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.ClusterVars">ClusterVars</a>, <a href="#apps.kubeblocks.io/v1.ComponentVars">ComponentVars</a>, <a href="#apps.kubeblocks.io/v1.CredentialVars">CredentialVars</a>, <a href="#apps.kubeblocks.io/v1.NamedVar">NamedVar</a>, <a href="#apps.kubeblocks.io/v1.ResourceVars">ResourceVars</a>, <a href="#apps.kubeblocks.io/v1.RoledVar">RoledVar</a>, <a href="#apps.kubeblocks.io/v1.ServiceRefVars">ServiceRefVars</a>, <a href="#apps.kubeblocks.io/v1.ServiceVars">ServiceVars</a>, <a href="#apps.kubeblocks.io/v1.TLSVars">TLSVars</a>)
</p>
<div>
<p>VarOption defines whether a variable is required or optional.</p>
</div>
<h3 id="apps.kubeblocks.io/v1.VarSource">VarSource
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1.EnvVar">EnvVar</a>)
</p>
<div>
<p>VarSource represents a source for the value of an EnvVar.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>configMapKeyRef</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#configmapkeyselector-v1-core">
Kubernetes core/v1.ConfigMapKeySelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a key of a ConfigMap.</p>
</td>
</tr>
<tr>
<td>
<code>secretKeyRef</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#secretkeyselector-v1-core">
Kubernetes core/v1.SecretKeySelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a key of a Secret.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetworkVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.HostNetworkVarSelector">
HostNetworkVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of host-network resources.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceVarSelector">
ServiceVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a Service.</p>
</td>
</tr>
<tr>
<td>
<code>credentialVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.CredentialVarSelector">
CredentialVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a Credential (SystemAccount).</p>
</td>
</tr>
<tr>
<td>
<code>tlsVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.TLSVarSelector">
TLSVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of the TLS.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ServiceRefVarSelector">
ServiceRefVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a ServiceRef.</p>
</td>
</tr>
<tr>
<td>
<code>resourceVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ResourceVarSelector">
ResourceVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a kind of resource.</p>
</td>
</tr>
<tr>
<td>
<code>componentVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ComponentVarSelector">
ComponentVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a Component.</p>
</td>
</tr>
<tr>
<td>
<code>clusterVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ClusterVarSelector">
ClusterVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a Cluster.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="apps.kubeblocks.io/v1alpha1">apps.kubeblocks.io/v1alpha1</h2>
<div>
</div>
Resource Types:
<ul><li>
<a href="#apps.kubeblocks.io/v1alpha1.Cluster">Cluster</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinition">ClusterDefinition</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.Component">Component</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinition">ComponentDefinition</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVersion">ComponentVersion</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraint">ConfigConstraint</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.Configuration">Configuration</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.Rollout">Rollout</a>
</li><li>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptor">ServiceDescriptor</a>
</li></ul>
<h3 id="apps.kubeblocks.io/v1alpha1.Cluster">Cluster
</h3>
<div>
<p>Cluster offers a unified management interface for a wide variety of database and storage systems:</p>
<ul>
<li>Relational databases: MySQL, PostgreSQL, MariaDB</li>
<li>NoSQL databases: Redis, MongoDB</li>
<li>KV stores: ZooKeeper, etcd</li>
<li>Analytics systems: ElasticSearch, OpenSearch, ClickHouse, Doris, StarRocks, Solr</li>
<li>Message queues: Kafka, Pulsar</li>
<li>Distributed SQL: TiDB, OceanBase</li>
<li>Vector databases: Qdrant, Milvus, Weaviate</li>
<li>Object storage: Minio</li>
</ul>
<p>KubeBlocks utilizes an abstraction layer to encapsulate the characteristics of these diverse systems.
A Cluster is composed of multiple Components, each defined by vendors or KubeBlocks Addon developers via ComponentDefinition,
arranged in Directed Acyclic Graph (DAG) topologies.
The topologies, defined in a ClusterDefinition, coordinate reconciliation across Cluster&rsquo;s lifecycle phases:
Creating, Running, Updating, Stopping, Stopped, Deleting.
Lifecycle management ensures that each Component operates in harmony, executing appropriate actions at each lifecycle stage.</p>
<p>For sharded-nothing architecture, the Cluster supports managing multiple shards,
each shard managed by a separate Component, supporting dynamic resharding.</p>
<p>The Cluster object is aimed to maintain the overall integrity and availability of a database cluster,
serves as the central control point, abstracting the complexity of multiple-component management,
and providing a unified interface for cluster-wide operations.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Cluster</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">
ClusterSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>clusterDefinitionRef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterDefinition to use when creating a Cluster.</p>
<p>This field enables users to create a Cluster based on a specific ClusterDefinition.
Which, in conjunction with the <code>topology</code> field, determine:</p>
<ul>
<li>The Components to be included in the Cluster.</li>
<li>The sequences in which the Components are created, updated, and terminate.</li>
</ul>
<p>This facilitates multiple-components management with predefined ClusterDefinition.</p>
<p>Users with advanced requirements can bypass this general setting and specify more precise control over
the composition of the Cluster by directly referencing specific ComponentDefinitions for each component
within <code>componentSpecs[*].componentDef</code>.</p>
<p>If this field is not provided, each component must be explicitly defined in <code>componentSpecs[*].componentDef</code>.</p>
<p>Note: Once set, this field cannot be modified; it is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>clusterVersionRef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the ClusterVersion name.</p>
<p>Deprecated since v0.9, use ComponentVersion instead.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>topology</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterTopology to be used when creating the Cluster.</p>
<p>This field defines which set of Components, as outlined in the ClusterDefinition, will be used to
construct the Cluster based on the named topology.
The ClusterDefinition may list multiple topologies under <code>clusterdefinition.spec.topologies[*]</code>,
each tailored to different use cases or environments.</p>
<p>If <code>topology</code> is not specified, the Cluster will use the default topology defined in the ClusterDefinition.</p>
<p>Note: Once set during the Cluster creation, the <code>topology</code> field cannot be modified.
It establishes the initial composition and structure of the Cluster and is intended for one-time configuration.</p>
</td>
</tr>
<tr>
<td>
<code>terminationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TerminationPolicyType">
TerminationPolicyType
</a>
</em>
</td>
<td>
<p>Specifies the behavior when a Cluster is deleted.
It defines how resources, data, and backups associated with a Cluster are managed during termination.
Choose a policy based on the desired level of resource cleanup and data preservation:</p>
<ul>
<li><code>DoNotTerminate</code>: Prevents deletion of the Cluster. This policy ensures that all resources remain intact.</li>
<li><code>Halt</code>: Deletes Cluster resources like Pods and Services but retains Persistent Volume Claims (PVCs),
allowing for data preservation while stopping other operations.
Warning: Halt policy is deprecated in 0.9.1 and will have same meaning as DoNotTerminate.</li>
<li><code>Delete</code>: Extends the <code>Halt</code> policy by also removing PVCs, leading to a thorough cleanup while
removing all persistent data.</li>
<li><code>WipeOut</code>: An aggressive policy that deletes all Cluster resources, including volume snapshots and
backups in external storage.
This results in complete data removal and should be used cautiously, primarily in non-production environments
to avoid irreversible data loss.</li>
</ul>
<p>Warning: Choosing an inappropriate termination policy can result in data loss.
The <code>WipeOut</code> policy is particularly risky in production environments due to its irreversible nature.</p>
</td>
</tr>
<tr>
<td>
<code>shardingSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ShardingSpec">
[]ShardingSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ShardingSpec objects that manage the sharding topology for Cluster Components.
Each ShardingSpec organizes components into shards, with each shard corresponding to a Component.
Components within a shard are all based on a common ClusterComponentSpec template, ensuring uniform configurations.</p>
<p>This field supports dynamic resharding by facilitating the addition or removal of shards
through the <code>shards</code> field in ShardingSpec.</p>
<p>Note: <code>shardingSpecs</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>componentSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">
[]ClusterComponentSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ClusterComponentSpec objects used to define the individual Components that make up a Cluster.
This field allows for detailed configuration of each Component within the Cluster.</p>
<p>Note: <code>shardingSpecs</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterService">
[]ClusterService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of additional Services that are exposed by a Cluster.
This field allows Services of selected Components, either from <code>componentSpecs</code> or <code>shardingSpecs</code> to be exposed,
alongside Services defined with ComponentService.</p>
<p>Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Affinity">
Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of node affinity scheduling rules for the Cluster&rsquo;s Pods.
This field helps control the placement of Pods on nodes within the Cluster.</p>
<p>Deprecated since v0.10. Use the <code>schedulingPolicy</code> field instead.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An array that specifies tolerations attached to the Cluster&rsquo;s Pods,
allowing them to be scheduled onto nodes with matching taints.</p>
<p>Deprecated since v0.10. Use the <code>schedulingPolicy</code> field instead.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies runtimeClassName for all Pods managed by this Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterBackup">
ClusterBackup
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the backup configuration of the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>tenancy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TenancyType">
TenancyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Describes how Pods are distributed across node.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>availabilityPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.AvailabilityPolicyType">
AvailabilityPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Describes the availability policy, including zone, node, and none.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the replicas of the first componentSpec, if the replicas of the first componentSpec is specified,
this value will be ignored.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterResources">
ClusterResources
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources of the first componentSpec, if the resources of the first componentSpec is specified,
this value will be ignored.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterStorage">
ClusterStorage
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the storage of the first componentSpec, if the storage of the first componentSpec is specified,
this value will be ignored.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterNetwork">
ClusterNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The configuration of network.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterStatus">
ClusterStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterDefinition">ClusterDefinition
</h3>
<div>
<p>ClusterDefinition defines the topology for databases or storage systems,
offering a variety of topological configurations to meet diverse deployment needs and scenarios.</p>
<p>It includes a list of Components, each linked to a ComponentDefinition, which enhances reusability and reduce redundancy.
For example, widely used components such as etcd and Zookeeper can be defined once and reused across multiple ClusterDefinitions,
simplifying the setup of new systems.</p>
<p>Additionally, ClusterDefinition also specifies the sequence of startup, upgrade, and shutdown for Components,
ensuring a controlled and predictable management of component lifecycles.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ClusterDefinition</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionSpec">
ClusterDefinitionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the well-known database type, such as mysql, redis, or mongodb.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>componentDefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">
[]ClusterComponentDefinition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides the definitions for the cluster components.</p>
<p>Deprecated since v0.9.
Components should now be individually defined using ComponentDefinition and
collectively referenced via <code>topology.components</code>.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>connectionCredential</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Connection credential template used for creating a connection credential secret for cluster objects.</p>
<p>Built-in objects are:</p>
<ul>
<li><code>$(RANDOM_PASSWD)</code> random 8 characters.</li>
<li><code>$(STRONG_RANDOM_PASSWD)</code> random 16 characters, with mixed cases, digits and symbols.</li>
<li><code>$(UUID)</code> generate a random UUID v4 string.</li>
<li><code>$(UUID_B64)</code> generate a random UUID v4 BASE64 encoded string.</li>
<li><code>$(UUID_STR_B64)</code> generate a random UUID v4 string then BASE64 encoded.</li>
<li><code>$(UUID_HEX)</code> generate a random UUID v4 HEX representation.</li>
<li><code>$(HEADLESS_SVC_FQDN)</code> headless service FQDN placeholder, value pattern is <code>$(CLUSTER_NAME)-$(1ST_COMP_NAME)-headless.$(NAMESPACE).svc</code>,
where 1ST_COMP_NAME is the 1st component that provide <code>ClusterDefinition.spec.componentDefs[].service</code> attribute;</li>
<li><code>$(SVC_FQDN)</code> service FQDN placeholder, value pattern is <code>$(CLUSTER_NAME)-$(1ST_COMP_NAME).$(NAMESPACE).svc</code>,
where 1ST_COMP_NAME is the 1st component that provide <code>ClusterDefinition.spec.componentDefs[].service</code> attribute;</li>
<li><code>$(SVC_PORT_&#123;PORT-NAME&#125;)</code> is ServicePort&rsquo;s port value with specified port name, i.e, a servicePort JSON struct:
<code>&#123;&quot;name&quot;: &quot;mysql&quot;, &quot;targetPort&quot;: &quot;mysqlContainerPort&quot;, &quot;port&quot;: 3306&#125;</code>, and <code>$(SVC_PORT_mysql)</code> in the
connection credential value is 3306.</li>
</ul>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>topologies</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterTopology">
[]ClusterTopology
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Topologies defines all possible topologies within the cluster.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionStatus">
ClusterDefinitionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Component">Component
</h3>
<div>
<p>Component is a fundamental building block of a Cluster object.
For example, a Redis Cluster can include Components like &lsquo;redis&rsquo;, &lsquo;sentinel&rsquo;, and potentially a proxy like &lsquo;twemproxy&rsquo;.</p>
<p>The Component object is responsible for managing the lifecycle of all replicas within a Cluster component,
It supports a wide range of operations including provisioning, stopping, restarting, termination, upgrading,
configuration changes, vertical and horizontal scaling, failover, switchover, cross-node migration,
scheduling configuration, exposing Services, managing system accounts, enabling/disabling exporter,
and configuring log collection.</p>
<p>Component is an internal sub-object derived from the user-submitted Cluster object.
It is designed primarily to be used by the KubeBlocks controllers,
users are discouraged from modifying Component objects directly and should use them only for monitoring Component statuses.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Component</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">
ComponentSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this Component.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRef">
[]ServiceRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of ServiceRef for a Component, enabling access to both external services and
Services provided by other Clusters.</p>
<p>Types of services:</p>
<ul>
<li>External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;
Require a ServiceDescriptor for connection details.</li>
<li>Services provided by a Cluster: Managed by the same KubeBlocks operator;
identified using Cluster, Component and Service names.</li>
</ul>
<p>ServiceRefs with identical <code>serviceRef.name</code> in the same Cluster are considered the same.</p>
<p>Example:</p>
<pre><code class="language-yaml">serviceRefs:
  - name: &quot;redis-sentinel&quot;
    serviceDescriptor:
      name: &quot;external-redis-sentinel&quot;
  - name: &quot;postgres-cluster&quot;
    clusterServiceSelector:
      cluster: &quot;my-postgres-cluster&quot;
      service:
        component: &quot;postgresql&quot;
</code></pre>
<p>The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Labels to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Annotations to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of environment variables to add.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources required by the Component.
It allows defining the CPU, memory requirements and limits for the Component&rsquo;s containers.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentVolumeClaimTemplate">
[]ClusterComponentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for the Component.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes to override.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides Services defined in referenced ComponentDefinition and exposes endpoints that can be accessed
by clients.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentSystemAccount">
[]ComponentSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides system accounts defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentConfig">
[]ClusterComponentConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration content of a config template.</p>
</td>
</tr>
<tr>
<td>
<code>enabledLogs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies which types of logs should be collected for the Cluster.
The log types are defined in the <code>componentDefinition.spec.logConfigs</code> field with the LogConfig entries.</p>
<p>The elements in the <code>enabledLogs</code> array correspond to the names of the LogConfig entries.
For example, if the <code>componentDefinition.spec.logConfigs</code> defines LogConfig entries with
names &ldquo;slow_query_log&rdquo; and &ldquo;error_log&rdquo;,
you can enable the collection of these logs by including their names in the <code>enabledLogs</code> array:</p>
<pre><code class="language-yaml">enabledLogs:
- slow_query_log
- error_log
</code></pre>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceAccount required by the running Component.
This ServiceAccount is used to grant necessary permissions for the Component&rsquo;s Pods to interact
with other Kubernetes resources, such as modifying Pod labels or sending events.</p>
<p>Defaults:
If not specified, KubeBlocks automatically assigns a default ServiceAccount named &ldquo;kb-&#123;cluster.name&#125;&rdquo;,
bound to a default role defined during KubeBlocks installation.</p>
<p>Future Changes:
Future versions might change the default ServiceAccount creation strategy to one per Component,
potentially revising the naming to &ldquo;kb-&#123;cluster.name&#125;-&#123;component.name&#125;&rdquo;.</p>
<p>Users can override the automatic ServiceAccount assignment by explicitly setting the name of
an existed ServiceAccount in this field.</p>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the InstanceUpdateStrategy that will be
employed to update Pods in the InstanceSet when a revision is made to
Template.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Affinity">
Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a group of affinity scheduling rules for the Component.
It allows users to control how the Component&rsquo;s Pods are scheduled onto nodes in the Cluster.</p>
<p>Deprecated since v0.10, replaced by the <code>schedulingPolicy</code> field.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows Pods to be scheduled onto nodes with matching taints.
Each toleration in the array allows the Pod to tolerate node taints based on
specified <code>key</code>, <code>value</code>, <code>effect</code>, and <code>operator</code>.</p>
<ul>
<li>The <code>key</code>, <code>value</code>, and <code>effect</code> identify the taint that the toleration matches.</li>
<li>The <code>operator</code> determines how the toleration matches the taint.</li>
</ul>
<p>Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.</p>
<p>Deprecated since v0.10, replaced by the <code>schedulingPolicy</code> field.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>tlsConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TLSConfig">
TLSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the TLS configuration for the Component, including:</p>
<ul>
<li>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication.</li>
<li>An optional field that specifies the configuration for the TLS certificates issuer when TLS is enabled.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows for the customization of configuration values for each instance within a Component.
An Instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).
While instances typically share a common configuration as defined in the ClusterComponentSpec,
they can require unique settings in various scenarios:</p>
<p>For example:
- A database Component might require different resource allocations for primary and secondary instances,
  with primaries needing more resources.
- During a rolling upgrade, a Component may first update the image for one or a few instances,
and then update the remaining instances after verifying that the updated instances are functioning correctly.</p>
<p>InstanceTemplate allows for specifying these unique configurations per instance.
Each instance&rsquo;s name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),
starting with an ordinal of 0.
It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the Component.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the Cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines runtimeClassName for all Pods managed by this Component.</p>
</td>
</tr>
<tr>
<td>
<code>disableExporter</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines whether metrics exporter information is annotated on the Component&rsquo;s headless Service.</p>
<p>If set to true, the following annotations will not be patched into the Service:</p>
<ul>
<li>&ldquo;monitor.kubeblocks.io/path&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/port&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/scheme&rdquo;</li>
</ul>
<p>These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.</p>
</td>
</tr>
<tr>
<td>
<code>stop</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Stop the Component.
If set, all the computing resources will be released.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentStatus">
ComponentStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentDefinition">ComponentDefinition
</h3>
<div>
<p>ComponentDefinition serves as a reusable blueprint for creating Components,
encapsulating essential static settings such as Component description,
Pod templates, configuration file templates, scripts, parameter lists,
injected environment variables and their sources, and event handlers.
ComponentDefinition works in conjunction with dynamic settings from the ClusterComponentSpec,
to instantiate Components during Cluster creation.</p>
<p>Key aspects that can be defined in a ComponentDefinition include:</p>
<ul>
<li>PodSpec template: Specifies the PodSpec template used by the Component.</li>
<li>Configuration templates: Specify the configuration file templates required by the Component.</li>
<li>Scripts: Provide the necessary scripts for Component management and operations.</li>
<li>Storage volumes: Specify the storage volumes and their configurations for the Component.</li>
<li>Pod roles: Outlines various roles of Pods within the Component along with their capabilities.</li>
<li>Exposed Kubernetes Services: Specify the Services that need to be exposed by the Component.</li>
<li>System accounts: Define the system accounts required for the Component.</li>
<li>Monitoring and logging: Configure the exporter and logging settings for the Component.</li>
</ul>
<p>ComponentDefinitions also enable defining reactive behaviors of the Component in response to events,
such as member join/leave, Component addition/deletion, role changes, switch over, and more.
This allows for automatic event handling, thus encapsulating complex behaviors within the Component.</p>
<p>Referencing a ComponentDefinition when creating individual Components ensures inheritance of predefined configurations,
promoting reusability and consistency across different deployments and cluster topologies.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ComponentDefinition</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">
ComponentDefinitionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>provider</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the Component provider, typically the vendor or developer name.
It identifies the entity responsible for creating and maintaining the Component.</p>
<p>When specifying the provider name, consider the following guidelines:</p>
<ul>
<li>Keep the name concise and relevant to the Component.</li>
<li>Use a consistent naming convention across Components from the same provider.</li>
<li>Avoid using trademarked or copyrighted names without proper permission.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a brief and concise explanation of the Component&rsquo;s purpose, functionality, and any relevant details.
It serves as a quick reference for users to understand the Component&rsquo;s role and characteristics.</p>
</td>
</tr>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the type of well-known service protocol that the Component provides.
It specifies the standard or widely recognized protocol used by the Component to offer its Services.</p>
<p>The <code>serviceKind</code> field allows users to quickly identify the type of Service provided by the Component
based on common protocols or service types. This information helps in understanding the compatibility,
interoperability, and usage of the Component within a system.</p>
<p>Some examples of well-known service protocols include:</p>
<ul>
<li>&ldquo;MySQL&rdquo;: Indicates that the Component provides a MySQL database service.</li>
<li>&ldquo;PostgreSQL&rdquo;: Indicates that the Component offers a PostgreSQL database service.</li>
<li>&ldquo;Redis&rdquo;: Signifies that the Component functions as a Redis key-value store.</li>
<li>&ldquo;ETCD&rdquo;: Denotes that the Component serves as an ETCD distributed key-value store.</li>
</ul>
<p>The <code>serviceKind</code> value is case-insensitive, allowing for flexibility in specifying the protocol name.</p>
<p>When specifying the <code>serviceKind</code>, consider the following guidelines:</p>
<ul>
<li>Use well-established and widely recognized protocol names or service types.</li>
<li>Ensure that the <code>serviceKind</code> accurately represents the primary service type offered by the Component.</li>
<li>If the Component provides multiple services, choose the most prominent or commonly used protocol.</li>
<li>Limit the <code>serviceKind</code> to a maximum of 32 characters for conciseness and readability.</li>
</ul>
<p>Note: The <code>serviceKind</code> field is optional and can be left empty if the Component does not fit into a well-known
service category or if the protocol is not widely recognized. It is primarily used to convey information about
the Component&rsquo;s service type to users and facilitate discovery and integration.</p>
<p>The <code>serviceKind</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the version of the Service provided by the Component.
It follows the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
<p>The Semantic Versioning specification defines a version number format of X.Y.Z (MAJOR.MINOR.PATCH), where:</p>
<ul>
<li>X represents the major version and indicates incompatible API changes.</li>
<li>Y represents the minor version and indicates added functionality in a backward-compatible manner.</li>
<li>Z represents the patch version and indicates backward-compatible bug fixes.</li>
</ul>
<p>Additional labels for pre-release and build metadata are available as extensions to the X.Y.Z format:</p>
<ul>
<li>Use pre-release labels (e.g., -alpha, -beta) for versions that are not yet stable or ready for production use.</li>
<li>Use build metadata (e.g., +build.1) for additional version information if needed.</li>
</ul>
<p>Examples of valid ServiceVersion values:</p>
<ul>
<li>&ldquo;1.0.0&rdquo;</li>
<li>&ldquo;2.3.1&rdquo;</li>
<li>&ldquo;3.0.0-alpha.1&rdquo;</li>
<li>&ldquo;4.5.2+build.1&rdquo;</li>
</ul>
<p>The <code>serviceVersion</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>runtime</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<p>Specifies the PodSpec template used in the Component.
It includes the following elements:</p>
<ul>
<li>Init containers</li>
<li>Containers
<ul>
<li>Image</li>
<li>Commands</li>
<li>Args</li>
<li>Envs</li>
<li>Mounts</li>
<li>Ports</li>
<li>Security context</li>
<li>Probes</li>
<li>Lifecycle</li>
</ul></li>
<li>Volumes</li>
</ul>
<p>This field is intended to define static settings that remain consistent across all instantiated Components.
Dynamic settings such as CPU and memory resource limits, as well as scheduling settings (affinity,
toleration, priority), may vary among different instantiated Components.
They should be specified in the <code>cluster.spec.componentSpecs</code> (ClusterComponentSpec).</p>
<p>Specific instances of a Component may override settings defined here, such as using a different container image
or modifying environment variable values.
These instance-specific overrides can be specified in <code>cluster.spec.componentSpecs[*].instances</code>.</p>
<p>This field is immutable and cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>monitor</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MonitorConfig">
MonitorConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Deprecated since v0.9
monitor is monitoring config which provided by provider.</p>
</td>
</tr>
<tr>
<td>
<code>exporter</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Exporter">
Exporter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the built-in metrics exporter container.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.EnvVar">
[]EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines variables which are determined after Cluster instantiation and reflect
dynamic or runtime attributes of instantiated Clusters.
These variables serve as placeholders for setting environment variables in Pods and Actions,
or for rendering configuration and script templates before actual values are finalized.</p>
<p>These variables are placed in front of the environment variables declared in the Pod if used as
environment variables.</p>
<p>Variable values can be sourced from:</p>
<ul>
<li>ConfigMap: Select and extract a value from a specific key within a ConfigMap.</li>
<li>Secret: Select and extract a value from a specific key within a Secret.</li>
<li>HostNetwork: Retrieves values (including ports) from host-network resources.</li>
<li>Service: Retrieves values (including address, port, NodePort) from a selected Service.
Intended to obtain the address of a ComponentService within the same Cluster.</li>
<li>Credential: Retrieves account name and password from a SystemAccount variable.</li>
<li>ServiceRef: Retrieves address, port, account name and password from a selected ServiceRefDeclaration.
Designed to obtain the address bound to a ServiceRef, such as a ClusterService or
ComponentService of another cluster or an external service.</li>
<li>Component: Retrieves values from a selected Component, including replicas and instance name list.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVolume">
[]ComponentVolume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the volumes used by the Component and some static attributes of the volumes.
After defining the volumes here, user can reference them in the
<code>cluster.spec.componentSpecs[*].volumeClaimTemplates</code> field to configure dynamic properties such as
volume capacity and storage class.</p>
<p>This field allows you to specify the following:</p>
<ul>
<li>Snapshot behavior: Determines whether a snapshot of the volume should be taken when performing
a snapshot backup of the Component.</li>
<li>Disk high watermark: Sets the high watermark for the volume&rsquo;s disk usage.
When the disk usage reaches the specified threshold, it triggers an alert or action.</li>
</ul>
<p>By configuring these volume behaviors, you can control how the volumes are managed and monitored within the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HostNetwork">
HostNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the host network configuration for the Component.</p>
<p>When <code>hostNetwork</code> option is enabled, the Pods share the host&rsquo;s network namespace and can directly access
the host&rsquo;s network interfaces.
This means that if multiple Pods need to use the same port, they cannot run on the same host simultaneously
due to port conflicts.</p>
<p>The DNSPolicy field in the Pod spec determines how containers within the Pod perform DNS resolution.
When using hostNetwork, the operator will set the DNSPolicy to &lsquo;ClusterFirstWithHostNet&rsquo;.
With this policy, DNS queries will first go through the K8s cluster&rsquo;s DNS service.
If the query fails, it will fall back to the host&rsquo;s DNS settings.</p>
<p>If set, the DNS policy will be automatically set to &ldquo;ClusterFirstWithHostNet&rdquo;.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines additional Services to expose the Component&rsquo;s endpoints.</p>
<p>A default headless Service, named <code>&#123;cluster.name&#125;-&#123;component.name&#125;-headless</code>, is automatically created
for internal Cluster communication.</p>
<p>This field enables customization of additional Services to expose the Component&rsquo;s endpoints to
other Components within the same or different Clusters, and to external applications.
Each Service entry in this list can include properties such as ports, type, and selectors.</p>
<ul>
<li>For intra-Cluster access, Components can reference Services using variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceVarRef</code>.</li>
<li>For inter-Cluster access, reference Services use variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceRefVarRef</code>,
and bind Services at Cluster creation time with <code>clusterComponentSpec.ServiceRef[*].clusterServiceSelector</code>.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">
[]ComponentConfigSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration file templates and volume mount parameters used by the Component.
It also includes descriptions of the parameters in the ConfigMaps, such as value range limitations.</p>
<p>This field specifies a list of templates that will be rendered into Component containers&rsquo; configuration files.
Each template is represented as a ConfigMap and may contain multiple configuration files,
with each file being a key in the ConfigMap.</p>
<p>The rendered configuration files will be mounted into the Component&rsquo;s containers
according to the specified volume mount parameters.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>logConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LogConfig">
[]LogConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the types of logs generated by instances of the Component and their corresponding file paths.
These logs can be collected for further analysis and monitoring.</p>
<p>The <code>logConfigs</code> field is an optional list of LogConfig objects, where each object represents
a specific log type and its configuration.
It allows you to specify multiple log types and their respective file paths for the Component.</p>
<p>Examples:</p>
<pre><code class="language-yaml"> logConfigs:
 - filePathPattern: /data/mysql/log/mysqld-error.log
   name: error
 - filePathPattern: /data/mysql/log/mysqld.log
   name: general
 - filePathPattern: /data/mysql/log/mysqld-slowquery.log
   name: slow
</code></pre>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>scripts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentTemplateSpec">
[]ComponentTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies groups of scripts, each provided via a ConfigMap, to be mounted as volumes in the container.
These scripts can be executed during container startup or via specific actions.</p>
<p>Each script group is encapsulated in a ComponentTemplateSpec that includes:</p>
<ul>
<li>The ConfigMap containing the scripts.</li>
<li>The mount point where the scripts will be mounted inside the container.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>policyRules</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#policyrule-v1-rbac">
[]Kubernetes rbac/v1.PolicyRule
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the namespaced policy rules required by the Component.</p>
<p>The <code>policyRules</code> field is an array of <code>rbacv1.PolicyRule</code> objects that define the policy rules
needed by the Component to operate within a namespace.
These policy rules determine the permissions and verbs the Component is allowed to perform on
Kubernetes resources within the namespace.</p>
<p>The purpose of this field is to automatically generate the necessary RBAC roles
for the Component based on the specified policy rules.
This ensures that the Pods in the Component has appropriate permissions to function.</p>
<p>Note: This field is currently non-functional and is reserved for future implementation.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static labels that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If a label key in the <code>labels</code> field conflicts with any system labels or user-specified labels,
it will be silently ignored to avoid overriding higher-priority labels.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static annotations that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If an annotation key in the <code>annotations</code> field conflicts with any system annotations
or user-specified annotations, it will be silently ignored to avoid overriding higher-priority annotations.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>replicasLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReplicasLimit">
ReplicasLimit
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the upper limit of the number of replicas supported by the Component.</p>
<p>It defines the maximum number of replicas that can be created for the Component.
This field allows you to set a limit on the scalability of the Component, preventing it from exceeding a certain number of replicas.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SystemAccount">
[]SystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An array of <code>SystemAccount</code> objects that define the system accounts needed
for the management operations of the Component.</p>
<p>Each <code>SystemAccount</code> includes:</p>
<ul>
<li>Account name.</li>
<li>The SQL statement template: Used to create the system account.</li>
<li>Password Source: Either generated based on certain rules or retrieved from a Secret.</li>
</ul>
<p>Use cases for system accounts typically involve tasks like system initialization, backups, monitoring,
health checks, replication, and other system-level operations.</p>
<p>System accounts are distinct from user accounts, although both are database accounts.</p>
<ul>
<li><strong>System Accounts</strong>: Created during Cluster setup by the KubeBlocks operator,
these accounts have higher privileges for system management and are fully managed
through a declarative API by the operator.</li>
<li><strong>User Accounts</strong>: Managed by users or administrator.
User account permissions should follow the principle of least privilege,
granting only the necessary access rights to complete their required tasks.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the concurrency strategy for updating multiple instances of the Component.
Available strategies:</p>
<ul>
<li><code>Serial</code>: Updates replicas one at a time, ensuring minimal downtime by waiting for each replica to become ready
before updating the next.</li>
<li><code>Parallel</code>: Updates all replicas simultaneously, optimizing for speed but potentially reducing availability
during the update.</li>
<li><code>BestEffortParallel</code>: Updates replicas concurrently with a limit on simultaneous updates to ensure a minimum
number of operational replicas for maintaining quorum.
 For example, in a 5-replica component, updating a maximum of 2 replicas simultaneously keeps
at least 3 operational for quorum.</li>
</ul>
<p>This field is immutable and defaults to &lsquo;Serial&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InstanceSet controls the creation of pods during initial scale up, replacement of pods on nodes, and scaling down.</p>
<ul>
<li><code>OrderedReady</code>: Creates pods in increasing order (pod-0, then pod-1, etc). The controller waits until each pod
is ready before continuing. Pods are removed in reverse order when scaling down.</li>
<li><code>Parallel</code>: Creates pods in parallel to match the desired scale without waiting. All pods are deleted at once
when scaling down.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enumerate all possible roles assigned to each replica of the Component, influencing its behavior.</p>
<p>A replica can have zero to multiple roles.
KubeBlocks operator determines the roles of each replica by invoking the <code>lifecycleActions.roleProbe</code> method.
This action returns a list of roles for each replica, and the returned roles must be predefined in the <code>roles</code> field.</p>
<p>The roles assigned to a replica can influence various aspects of the Component&rsquo;s behavior, such as:</p>
<ul>
<li>Service selection: The Component&rsquo;s exposed Services may target replicas based on their roles using <code>roleSelector</code>.</li>
<li>Update order: The roles can determine the order in which replicas are updated during a Component update.
For instance, replicas with a &ldquo;follower&rdquo; role can be updated first, while the replica with the &ldquo;leader&rdquo;
role is updated last. This helps minimize the number of leader changes during the update process.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>roleArbitrator</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RoleArbitrator">
RoleArbitrator
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>This field has been deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>lifecycleActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentLifecycleActions">
ComponentLifecycleActions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of hooks and procedures that customize the behavior of a Component throughout its lifecycle.
Actions are triggered at specific lifecycle stages:</p>
<ul>
<li><code>postProvision</code>: Defines the hook to be executed after the creation of a Component,
with <code>preCondition</code> specifying when the action should be fired relative to the Component&rsquo;s lifecycle stages:
<code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.</li>
<li><code>preTerminate</code>: Defines the hook to be executed before terminating a Component.</li>
<li><code>roleProbe</code>: Defines the procedure which is invoked regularly to assess the role of replicas.</li>
<li><code>switchover</code>: Defines the procedure for a controlled transition of leadership from the current leader to a new replica.
This approach aims to minimize downtime and maintain availability in systems with a leader-follower topology,
such as before planned maintenance or upgrades on the current leader node.</li>
<li><code>memberJoin</code>: Defines the procedure to add a new replica to the replication group.</li>
<li><code>memberLeave</code>: Defines the method to remove a replica from the replication group.</li>
<li><code>readOnly</code>: Defines the procedure to switch a replica into the read-only state.</li>
<li><code>readWrite</code>: transition a replica from the read-only state back to the read-write state.</li>
<li><code>dataDump</code>: Defines the procedure to export the data from a replica.</li>
<li><code>dataLoad</code>: Defines the procedure to import data into a replica.</li>
<li><code>reconfigure</code>: Defines the procedure that update a replica with new configuration file.</li>
<li><code>accountProvision</code>: Defines the procedure to generate a new database account.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefDeclarations</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefDeclaration">
[]ServiceRefDeclaration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists external service dependencies of the Component, including services from other Clusters or outside the K8s environment.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p><code>minReadySeconds</code> is the minimum duration in seconds that a new Pod should remain in the ready
state without any of its containers crashing to be considered available.
This ensures the Pod&rsquo;s stability and readiness to serve requests.</p>
<p>A default value of 0 seconds means the Pod is considered available as soon as it enters the ready state.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionStatus">
ComponentDefinitionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVersion">ComponentVersion
</h3>
<div>
<p>ComponentVersion is the Schema for the componentversions API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ComponentVersion</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionSpec">
ComponentVersionSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>compatibilityRules</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionCompatibilityRule">
[]ComponentVersionCompatibilityRule
</a>
</em>
</td>
<td>
<p>CompatibilityRules defines compatibility rules between sets of component definitions and releases.</p>
</td>
</tr>
<tr>
<td>
<code>releases</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionRelease">
[]ComponentVersionRelease
</a>
</em>
</td>
<td>
<p>Releases represents different releases of component instances within this ComponentVersion.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionStatus">
ComponentVersionStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigConstraint">ConfigConstraint
</h3>
<div>
<p>ConfigConstraint manages the parameters across multiple configuration files contained in a single configure template.
These configuration files should have the same format (e.g. ini, xml, properties, json).</p>
<p>It provides the following functionalities:</p>
<ol>
<li><strong>Parameter Value Validation</strong>: Validates and ensures compliance of parameter values with defined constraints.</li>
<li><strong>Dynamic Reload on Modification</strong>: Monitors parameter changes and triggers dynamic reloads to apply updates.</li>
<li><strong>Parameter Rendering in Templates</strong>: Injects parameters into templates to generate up-to-date configuration files.</li>
</ol>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ConfigConstraint</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">
ConfigConstraintSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>reloadOptions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReloadOptions">
ReloadOptions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the dynamic reload action supported by the engine.
When set, the controller executes the method defined here to execute hot parameter updates.</p>
<p>Dynamic reloading is triggered only if both of the following conditions are met:</p>
<ol>
<li>The modified parameters are listed in the <code>dynamicParameters</code> field.
If <code>reloadStaticParamsBeforeRestart</code> is set to true, modifications to <code>staticParameters</code>
can also trigger a reload.</li>
<li><code>reloadOptions</code> is set.</li>
</ol>
<p>If <code>reloadOptions</code> is not set or the modified parameters are not listed in <code>dynamicParameters</code>,
dynamic reloading will not be triggered.</p>
<p>Example:</p>
<pre><code class="language-yaml">reloadOptions:
 tplScriptTrigger:
   namespace: kb-system
   scriptConfigMapRef: mysql-reload-script
   sync: true
</code></pre>
</td>
</tr>
<tr>
<td>
<code>dynamicActionCanBeMerged</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to consolidate dynamic reload and restart actions into a single restart.</p>
<ul>
<li>If true, updates requiring both actions will result in only a restart, merging the actions.</li>
<li>If false, updates will trigger both actions executed sequentially: first dynamic reload, then restart.</li>
</ul>
<p>This flag allows for more efficient handling of configuration changes by potentially eliminating
an unnecessary reload step.</p>
</td>
</tr>
<tr>
<td>
<code>reloadStaticParamsBeforeRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Configures whether the dynamic reload specified in <code>reloadOptions</code> applies only to dynamic parameters or
to all parameters (including static parameters).</p>
<ul>
<li>false (default): Only modifications to the dynamic parameters listed in <code>dynamicParameters</code>
will trigger a dynamic reload.</li>
<li>true: Modifications to both dynamic parameters listed in <code>dynamicParameters</code> and static parameters
listed in <code>staticParameters</code> will trigger a dynamic reload.
The &ldquo;true&rdquo; option is for certain engines that require static parameters to be set
via SQL statements before they can take effect on restart.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>toolsImageSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ToolsSetup">
ToolsSetup
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the tools container image used by ShellTrigger for dynamic reload.
If the dynamic reload action is triggered by a ShellTrigger, this field is required.
This image must contain all necessary tools for executing the ShellTrigger scripts.</p>
<p>Usually the specified image is referenced by the init container,
which is then responsible for copy the tools from the image to a bin volume.
This ensures that the tools are available to the &lsquo;config-manager&rsquo; sidecar.</p>
</td>
</tr>
<tr>
<td>
<code>downwardAPIOptions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.DownwardAPIChangeTriggeredAction">
[]DownwardAPIChangeTriggeredAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of actions to execute specified commands based on Pod labels.</p>
<p>It utilizes the K8s Downward API to mount label information as a volume into the pod.
The &lsquo;config-manager&rsquo; sidecar container watches for changes in the role label and dynamically invoke
registered commands (usually execute some SQL statements) when a change is detected.</p>
<p>It is designed for scenarios where:</p>
<ul>
<li>Replicas with different roles have different configurations, such as Redis primary &amp; secondary replicas.</li>
<li>After a role switch (e.g., from secondary to primary), some changes in configuration are needed
to reflect the new role.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>scriptConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ScriptConfig">
[]ScriptConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A list of ScriptConfig Object.</p>
<p>Each ScriptConfig object specifies a ConfigMap that contains script files that should be mounted inside the pod.
The scripts are mounted as volumes and can be referenced and executed by the dynamic reload
and DownwardAction to perform specific tasks or configurations.</p>
</td>
</tr>
<tr>
<td>
<code>cfgSchemaTopLevelName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the top-level key in the &lsquo;configurationSchema.cue&rsquo; that organizes the validation rules for parameters.
This key must exist within the CUE script defined in &lsquo;configurationSchema.cue&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>configurationSchema</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CustomParametersValidation">
CustomParametersValidation
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of parameters including their names, default values, descriptions,
types, and constraints (permissible values or the range of valid values).</p>
</td>
</tr>
<tr>
<td>
<code>staticParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List static parameters.
Modifications to any of these parameters require a restart of the process to take effect.</p>
</td>
</tr>
<tr>
<td>
<code>dynamicParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List dynamic parameters.
Modifications to these parameters trigger a configuration reload without requiring a process restart.</p>
</td>
</tr>
<tr>
<td>
<code>immutableParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists the parameters that cannot be modified once set.
Attempting to change any of these parameters will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to match labels on the pod to determine whether a dynamic reload should be performed.</p>
<p>In some scenarios, only specific pods (e.g., primary replicas) need to undergo a dynamic reload.
The <code>selector</code> allows you to specify label selectors to target the desired pods for the reload process.</p>
<p>If the <code>selector</code> is not specified or is nil, all pods managed by the workload will be considered for the dynamic
reload.</p>
</td>
</tr>
<tr>
<td>
<code>formatterConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.FileFormatConfig">
FileFormatConfig
</a>
</em>
</td>
<td>
<p>Specifies the format of the configuration file and any associated parameters that are specific to the chosen format.
Supported formats include <code>ini</code>, <code>xml</code>, <code>yaml</code>, <code>json</code>, <code>hcl</code>, <code>dotenv</code>, <code>properties</code>, and <code>toml</code>.</p>
<p>Each format may have its own set of parameters that can be configured.
For instance, when using the <code>ini</code> format, you can specify the section name.</p>
<p>Example:</p>
<pre><code>formatterConfig:
 format: ini
 iniConfig:
   sectionName: mysqld
</code></pre>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintStatus">
ConfigConstraintStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Configuration">Configuration
</h3>
<div>
<p>Configuration represents the complete set of configurations for a specific Component of a Cluster.
This includes templates for each configuration file, their corresponding ConfigConstraints, volume mounts,
and other relevant details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Configuration</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigurationSpec">
ConfigurationSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>clusterRef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the Cluster that this configuration is associated with.</p>
</td>
</tr>
<tr>
<td>
<code>componentName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the name of the Component that this configuration pertains to.</p>
</td>
</tr>
<tr>
<td>
<code>configItemDetails</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetail">
[]ConfigurationItemDetail
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigItemDetails is an array of ConfigurationItemDetail objects.</p>
<p>Each ConfigurationItemDetail corresponds to a configuration template,
which is a ConfigMap that contains multiple configuration files.
Each configuration file is stored as a key-value pair within the ConfigMap.</p>
<p>The ConfigurationItemDetail includes information such as:</p>
<ul>
<li>The configuration template (a ConfigMap)</li>
<li>The corresponding ConfigConstraint (constraints and validation rules for the configuration)</li>
<li>Volume mounts (for mounting the configuration files)</li>
</ul>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigurationStatus">
ConfigurationStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Rollout">Rollout
</h3>
<div>
<p>Rollout is the Schema for the rollouts API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Rollout</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutSpec">
RolloutSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>clusterName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the target cluster of the Rollout.</p>
</td>
</tr>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutComponent">
[]RolloutComponent
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target components to be rolled out.</p>
</td>
</tr>
<tr>
<td>
<code>shardings</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutSharding">
[]RolloutSharding
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target shardings to be rolled out.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutStatus">
RolloutStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceDescriptor">ServiceDescriptor
</h3>
<div>
<p>ServiceDescriptor describes a service provided by external sources.
It contains the necessary details such as the service&rsquo;s address and connection credentials.
To enable a Cluster to access this service, the ServiceDescriptor&rsquo;s name should be specified
in the Cluster configuration under <code>clusterComponent.serviceRefs[*].serviceDescriptor</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ServiceDescriptor</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptorSpec">
ServiceDescriptorSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the type of database service provided by the external service.
For example, &ldquo;mysql&rdquo;, &ldquo;redis&rdquo;, &ldquo;mongodb&rdquo;.
This field categorizes databases by their functionality, protocol and compatibility, facilitating appropriate
service integration based on their unique capabilities.</p>
<p>This field is case-insensitive.</p>
<p>It also supports abbreviations for some well-known databases:
- &ldquo;pg&rdquo;, &ldquo;pgsql&rdquo;, &ldquo;postgres&rdquo;, &ldquo;postgresql&rdquo;: PostgreSQL service
- &ldquo;zk&rdquo;, &ldquo;zookeeper&rdquo;: ZooKeeper service
- &ldquo;es&rdquo;, &ldquo;elasticsearch&rdquo;: Elasticsearch service
- &ldquo;mongo&rdquo;, &ldquo;mongodb&rdquo;: MongoDB service
- &ldquo;ch&rdquo;, &ldquo;clickhouse&rdquo;: ClickHouse service</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the version of the service provided by the external service.
This is crucial for ensuring compatibility between different components of the system,
as different versions of a service may have varying features.</p>
</td>
</tr>
<tr>
<td>
<code>endpoint</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the endpoint of the external service.</p>
<p>If the service is exposed via a cluster, the endpoint will be provided in the format of <code>host:port</code>.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the service or IP address of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the port of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConnectionCredentialAuth">
ConnectionCredentialAuth
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the authentication credentials required for accessing an external service.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptorStatus">
ServiceDescriptorStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.AccessMode">AccessMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConsensusMember">ConsensusMember</a>)
</p>
<div>
<p>AccessMode defines the modes of access granted to the SVC.
The modes can be <code>None</code>, <code>Readonly</code>, or <code>ReadWrite</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td><p>None implies no access.</p>
</td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td><p>ReadWrite permits both read and write operations.</p>
</td>
</tr><tr><td><p>&#34;Readonly&#34;</p></td>
<td><p>Readonly allows only read operations.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.AccountName">AccountName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.SystemAccountConfig">SystemAccountConfig</a>)
</p>
<div>
<p>AccountName defines system account names.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;kbadmin&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kbdataprotection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kbmonitoring&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kbprobe&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kbreplicator&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Action">Action
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentSwitchover">ComponentSwitchover</a>, <a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">LifecycleActionHandler</a>, <a href="#apps.kubeblocks.io/v1alpha1.Probe">Probe</a>)
</p>
<div>
<p>Action defines a customizable hook or procedure tailored for different database engines,
designed to be invoked at predetermined points within the lifecycle of a Component instance.
It provides a modular and extensible way to customize a Component&rsquo;s behavior through the execution of defined actions.</p>
<p>Available Action triggers include:</p>
<ul>
<li><code>postProvision</code>: Defines the hook to be executed after the creation of a Component,
with <code>preCondition</code> specifying when the action should be fired relative to the Component&rsquo;s lifecycle stages:
<code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.</li>
<li><code>preTerminate</code>: Defines the hook to be executed before terminating a Component.</li>
<li><code>roleProbe</code>: Defines the procedure which is invoked regularly to assess the role of replicas.</li>
<li><code>switchover</code>: Defines the procedure for a controlled transition of leadership from the current leader to a new replica.
This approach aims to minimize downtime and maintain availability in systems with a leader-follower topology,
such as during planned maintenance or upgrades on the current leader node.</li>
<li><code>memberJoin</code>: Defines the procedure to add a new replica to the replication group.</li>
<li><code>memberLeave</code>: Defines the method to remove a replica from the replication group.</li>
<li><code>readOnly</code>: Defines the procedure to switch a replica into the read-only state.</li>
<li><code>readWrite</code>: Defines the procedure to transition a replica from the read-only state back to the read-write state.</li>
<li><code>dataDump</code>: Defines the procedure to export the data from a replica.</li>
<li><code>dataLoad</code>: Defines the procedure to import data into a replica.</li>
<li><code>reconfigure</code>: Defines the procedure that update a replica with new configuration.</li>
<li><code>accountProvision</code>: Defines the procedure to generate a new database account.</li>
</ul>
<p>Actions can be executed in different ways:</p>
<ul>
<li>ExecAction: Executes a command inside a container.
which may run as a K8s job or be executed inside the Lorry sidecar container, depending on the implementation.
Future implementations will standardize execution within Lorry.
A set of predefined environment variables are available and can be leveraged within the <code>exec.command</code>
to access context information such as details about pods, components, the overall cluster state,
or database connection credentials.
These variables provide a dynamic and context-aware mechanism for script execution.</li>
<li>HTTPAction: Performs an HTTP request.
HTTPAction is to be implemented in future version.</li>
<li>GRPCAction: In future version, Actions will support initiating gRPC calls.
This allows developers to implement Actions using plugins written in programming language like Go,
providing greater flexibility and extensibility.</li>
</ul>
<p>An action is considered successful on returning 0, or HTTP 200 for status HTTP(s) Actions.
Any other return value or HTTP status codes indicate failure,
and the action may be retried based on the configured retry policy.</p>
<ul>
<li>If an action exceeds the specified timeout duration, it will be terminated, and the action is considered failed.</li>
<li>If an action produces any data as output, it should be written to stdout,
or included in the HTTP response payload for HTTP(s) actions.</li>
<li>If an action encounters any errors, error messages should be written to stderr,
or detailed in the HTTP response with the appropriate non-200 status code.</li>
</ul>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the container image to be used for running the Action.</p>
<p>When specified, a dedicated container will be created using this image to execute the Action.
This field is mutually exclusive with the <code>container</code> field; only one of them should be provided.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>exec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ExecAction">
ExecAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the command to run.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>http</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HTTPAction">
HTTPAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the HTTP request to perform.</p>
<p>This field cannot be updated.</p>
<p>Note: HTTPAction is to be implemented in future version.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a list of environment variables that will be injected into the container.
These variables enable the container to adapt its behavior based on the environment it&rsquo;s running in.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>targetPodSelector</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TargetPodSelector">
TargetPodSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the criteria used to select the target Pod(s) for executing the Action.
This is useful when there is no default target replica identified.
It allows for precise control over which Pod(s) the Action should run in.</p>
<p>This field cannot be updated.</p>
<p>Note: This field is reserved for future use and is not currently active.</p>
</td>
</tr>
<tr>
<td>
<code>matchingKey</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used in conjunction with the <code>targetPodSelector</code> field to refine the selection of target pod(s) for Action execution.
The impact of this field depends on the <code>targetPodSelector</code> value:</p>
<ul>
<li>When <code>targetPodSelector</code> is set to <code>Any</code> or <code>All</code>, this field will be ignored.</li>
<li>When <code>targetPodSelector</code> is set to <code>Role</code>, only those replicas whose role matches the <code>matchingKey</code>
will be selected for the Action.</li>
</ul>
<p>This field cannot be updated.</p>
<p>Note: This field is reserved for future use and is not currently active.</p>
</td>
</tr>
<tr>
<td>
<code>container</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the name of the container within the target Pod where the action will be executed.</p>
<p>This name must correspond to one of the containers defined in <code>componentDefinition.spec.runtime</code>.
If this field is not specified, the default behavior is to use the first container listed in
<code>componentDefinition.spec.runtime</code>.</p>
<p>This field cannot be updated.</p>
<p>Note: This field is reserved for future use and is not currently active.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the maximum duration in seconds that the Action is allowed to run.</p>
<p>If the Action does not complete within this time frame, it will be terminated.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>retryPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RetryPolicy">
RetryPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the strategy to be taken when retrying the Action after a failure.</p>
<p>It specifies the conditions under which the Action should be retried and the limits to apply,
such as the maximum number of retries and backoff strategy.</p>
<p>This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>preCondition</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PreConditionType">
PreConditionType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the state that the cluster must reach before the Action is executed.
Currently, this is only applicable to the <code>postProvision</code> action.</p>
<p>The conditions are as follows:</p>
<ul>
<li><code>Immediately</code>: Executed right after the Component object is created.
The readiness of the Component and its resources is not guaranteed at this stage.</li>
<li><code>RuntimeReady</code>: The Action is triggered after the Component object has been created and all associated
runtime resources (e.g. Pods) are in a ready state.</li>
<li><code>ComponentReady</code>: The Action is triggered after the Component itself is in a ready state.
This process does not affect the readiness state of the Component or the Cluster.</li>
<li><code>ClusterReady</code>: The Action is executed after the Cluster is in a ready state.
This execution does not alter the Component or the Cluster&rsquo;s state of readiness.</li>
</ul>
<p>This field cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Affinity">Affinity
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>podAntiAffinity</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PodAntiAffinity">
PodAntiAffinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the anti-affinity level of Pods within a Component.
It determines how pods should be spread across nodes to improve availability and performance.
It can have the following values: <code>Preferred</code> and <code>Required</code>.
The default value is <code>Preferred</code>.</p>
</td>
</tr>
<tr>
<td>
<code>topologyKeys</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the key of node labels used to define the topology domain for Pod anti-affinity
and Pod spread constraints.</p>
<p>In K8s, a topology domain is a set of nodes that have the same value for a specific label key.
Nodes with labels containing any of the specified TopologyKeys and identical values are considered
to be in the same topology domain.</p>
<p>Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of
topology in the ClusterDefinition.</p>
<p>When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the
Pod on nodes with different values for the specified TopologyKeys.
This ensures that Pods are spread across different topology domains, promoting high availability and
reducing the impact of node failures.</p>
<p>Some well-known label keys, such as <code>kubernetes.io/hostname</code> and <code>topology.kubernetes.io/zone</code>,
are often used as TopologyKey.
These keys represent the hostname and zone of a node, respectively.
By including these keys in the TopologyKeys list, Pods will be spread across nodes with
different hostnames or zones.</p>
<p>In addition to the well-known keys, users can also specify custom label keys as TopologyKeys.
This allows for more flexible and custom topology definitions based on the specific needs
of the application or environment.</p>
<p>The TopologyKeys field is a slice of strings, where each string represents a label key.
The order of the keys in the slice does not matter.</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the node labels that must be present on nodes for pods to be scheduled on them.
It is a map where the keys are the label keys and the values are the corresponding label values.
Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.</p>
<p>For example, if NodeLabels is set to &#123;&ldquo;nodeType&rdquo;: &ldquo;ssd&rdquo;, &ldquo;environment&rdquo;: &ldquo;production&rdquo;&#125;,
pods will only be scheduled on nodes that have both the &ldquo;nodeType&rdquo; label with value &ldquo;ssd&rdquo;
and the &ldquo;environment&rdquo; label with value &ldquo;production&rdquo;.</p>
<p>This field allows users to control Pod placement based on specific node labels.
It can be used to ensure that Pods are scheduled on nodes with certain characteristics,
such as specific hardware (e.g., SSD), environment (e.g., production, staging),
or any other custom labels assigned to nodes.</p>
</td>
</tr>
<tr>
<td>
<code>tenancy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TenancyType">
TenancyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines the level of resource isolation between Pods.
It can have the following values: <code>SharedNode</code> and <code>DedicatedNode</code>.</p>
<ul>
<li>SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s.</li>
<li>DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node.
In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node.
Which provides a higher level of isolation and resource guarantee for Pods.</li>
</ul>
<p>The default value is <code>SharedNode</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.AvailabilityPolicyType">AvailabilityPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>AvailabilityPolicyType defines the type of availability policy to be applied for cluster affinity, influencing how
resources are distributed across zones or nodes for high availability and resilience.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;node&#34;</p></td>
<td><p>AvailabilityPolicyNode specifies that resources should be distributed across different nodes within the same zone.
This policy aims to provide resilience against node failures, ensuring that the failure of a single node does not
impact the overall service availability.</p>
</td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td><p>AvailabilityPolicyNone specifies that no specific availability policy is applied.
Resources may not be explicitly distributed for high availability, potentially concentrating them in a single
zone or node based on other scheduling decisions.</p>
</td>
</tr><tr><td><p>&#34;zone&#34;</p></td>
<td><p>AvailabilityPolicyZone specifies that resources should be distributed across different availability zones.
This policy aims to ensure high availability and protect against zone failures, spreading the resources to reduce
the risk of simultaneous downtime.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.BackupStatusUpdateStage">BackupStatusUpdateStage
(<code>string</code> alias)</h3>
<div>
<p>BackupStatusUpdateStage defines the stage of backup status update.</p>
</div>
<h3 id="apps.kubeblocks.io/v1alpha1.BaseBackupType">BaseBackupType
(<code>string</code> alias)</h3>
<div>
<p>BaseBackupType the base backup type, keep synchronized with the BaseBackupType of the data protection API.</p>
</div>
<h3 id="apps.kubeblocks.io/v1alpha1.BuiltinActionHandlerType">BuiltinActionHandlerType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">LifecycleActionHandler</a>)
</p>
<div>
<p>BuiltinActionHandlerType defines build-in action handlers provided by Lorry, including:</p>
<ul>
<li><code>mysql</code></li>
<li><code>wesql</code></li>
<li><code>oceanbase</code></li>
<li><code>redis</code></li>
<li><code>mongodb</code></li>
<li><code>etcd</code></li>
<li><code>postgresql</code></li>
<li><code>vanilla-postgresql</code></li>
<li><code>apecloud-postgresql</code></li>
<li><code>polardbx</code></li>
<li><code>custom</code></li>
<li><code>unknown</code></li>
</ul>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;apecloud-postgresql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;custom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;etcd&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;mongodb&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;mysql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;oceanbase&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;polardbx&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;postgresql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;redis&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;unknown&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;vanilla-postgresql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wesql&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClassDefRef">ClassDefRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>)
</p>
<div>
<p>ClassDefRef is deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ComponentClassDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>class</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the name of the class that is defined in the ComponentClassDefinition.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterBackup">ClusterBackup
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether automated backup is enabled for the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPeriod</code><br/>
<em>
github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1.RetentionPeriod
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines the duration to retain backups. Backups older than this period are automatically removed.</p>
<p>For example, RetentionPeriod of <code>30d</code> will keep only the backups of last 30 days.
Sample duration format:</p>
<ul>
<li>years: 	2y</li>
<li>months: 	6mo</li>
<li>days: 		30d</li>
<li>hours: 	12h</li>
<li>minutes: 	30m</li>
</ul>
<p>You can also combine the above durations. For example: 30d12h30m.
Default value is 7d.</p>
</td>
</tr>
<tr>
<td>
<code>method</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the backup method to use, as defined in backupPolicy.</p>
</td>
</tr>
<tr>
<td>
<code>cronExpression</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The cron expression for the schedule. The timezone is in UTC. See <a href="https://en.wikipedia.org/wiki/Cron">https://en.wikipedia.org/wiki/Cron</a>.</p>
</td>
</tr>
<tr>
<td>
<code>startingDeadlineMinutes</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the maximum time in minutes that the system will wait to start a missed backup job.
If the scheduled backup time is missed for any reason, the backup job must start within this deadline.
Values must be between 0 (immediate execution) and 1440 (one day).</p>
</td>
</tr>
<tr>
<td>
<code>repoName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the backupRepo. If not set, the default backupRepo will be used.</p>
</td>
</tr>
<tr>
<td>
<code>pitrEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to enable point-in-time recovery.</p>
</td>
</tr>
<tr>
<td>
<code>continuousMethod</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the backup method to use, if not set, use the first continuous method.</p>
</td>
</tr>
<tr>
<td>
<code>incrementalBackupEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to enable incremental backup.</p>
</td>
</tr>
<tr>
<td>
<code>incrementalCronExpression</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The cron expression for the incremental backup schedule. The timezone is in UTC. See <a href="https://en.wikipedia.org/wiki/Cron">https://en.wikipedia.org/wiki/Cron</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentConfig">ClusterComponentConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
<p>ClusterComponentConfig represents a config with its source bound.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the config.</p>
</td>
</tr>
<tr>
<td>
<code>ClusterComponentConfigSource</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentConfigSource">
ClusterComponentConfigSource
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterComponentConfigSource</code> are embedded into this type.)
</p>
<p>The source of the config.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentConfigSource">ClusterComponentConfigSource
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentConfig">ClusterComponentConfig</a>)
</p>
<div>
<p>ClusterComponentConfigSource represents the source of a config.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>configMap</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#configmapvolumesource-v1-core">
Kubernetes core/v1.ConfigMapVolumeSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigMap source for the config.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionSpec">ClusterDefinitionSpec</a>)
</p>
<div>
<p>ClusterComponentDefinition defines a Component within a ClusterDefinition but is deprecated and
has been replaced by ComponentDefinition.</p>
<p>Deprecated: Use ComponentDefinition instead. This type is deprecated as of version 0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>This name could be used as default name of <code>cluster.spec.componentSpecs.name</code>, and needs to conform with same
validation rules as <code>cluster.spec.componentSpecs.name</code>, currently complying with IANA Service Naming rule.
This name will apply to cluster objects as the value of label &ldquo;apps.kubeblocks.io/component-name&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Description of the component definition.</p>
</td>
</tr>
<tr>
<td>
<code>workloadType</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.WorkloadType">
WorkloadType
</a>
</em>
</td>
<td>
<p>Defines the type of the workload.</p>
<ul>
<li><code>Stateless</code> describes stateless applications.</li>
<li><code>Stateful</code> describes common stateful applications.</li>
<li><code>Consensus</code> describes applications based on consensus protocols, such as raft and paxos.</li>
<li><code>Replication</code> describes applications based on the primary-secondary data replication protocol.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>characterType</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines well-known database component name, such as mongos(mongodb), proxy(redis), mariadb(mysql).</p>
</td>
</tr>
<tr>
<td>
<code>configSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">
[]ComponentConfigSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the template of configurations.</p>
</td>
</tr>
<tr>
<td>
<code>scriptSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentTemplateSpec">
[]ComponentTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the template of scripts.</p>
</td>
</tr>
<tr>
<td>
<code>probes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbes">
ClusterDefinitionProbes
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Settings for health checks.</p>
</td>
</tr>
<tr>
<td>
<code>logConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LogConfig">
[]LogConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specify the logging files which can be observed and configured by cluster users.</p>
</td>
</tr>
<tr>
<td>
<code>podSpec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the pod spec template of component.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceSpec">
ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the service spec.</p>
</td>
</tr>
<tr>
<td>
<code>statelessSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.StatelessSetSpec">
StatelessSetSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines spec for <code>Stateless</code> workloads.</p>
</td>
</tr>
<tr>
<td>
<code>statefulSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.StatefulSetSpec">
StatefulSetSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines spec for <code>Stateful</code> workloads.</p>
</td>
</tr>
<tr>
<td>
<code>consensusSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConsensusSetSpec">
ConsensusSetSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines spec for <code>Consensus</code> workloads. It&rsquo;s required if the workload type is <code>Consensus</code>.</p>
</td>
</tr>
<tr>
<td>
<code>replicationSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReplicationSetSpec">
ReplicationSetSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines spec for <code>Replication</code> workloads.</p>
</td>
</tr>
<tr>
<td>
<code>rsmSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RSMSpec">
RSMSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines workload spec of this component.
From KB 0.7.0, RSM(InstanceSetSpec) will be the underlying CR which powers all kinds of workload in KB.
RSM is an enhanced stateful workload extension dedicated for heavy-state workloads like databases.</p>
</td>
</tr>
<tr>
<td>
<code>horizontalScalePolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HorizontalScalePolicy">
HorizontalScalePolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the behavior of horizontal scale.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SystemAccountSpec">
SystemAccountSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines system accounts needed to manage the component, and the statement to create them.</p>
</td>
</tr>
<tr>
<td>
<code>volumeTypes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VolumeTypeSpec">
[]VolumeTypeSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to describe the purpose of the volumes mapping the name of the VolumeMounts in the PodSpec.Container field,
such as data volume, log volume, etc. When backing up the volume, the volume can be correctly backed up according
to the volumeType.</p>
<p>For example:</p>
<ul>
<li><code>name: data, type: data</code> means that the volume named <code>data</code> is used to store <code>data</code>.</li>
<li><code>name: binlog, type: log</code> means that the volume named <code>binlog</code> is used to store <code>log</code>.</li>
</ul>
<p>NOTE: When volumeTypes is not defined, the backup function will not be supported, even if a persistent volume has
been specified.</p>
</td>
</tr>
<tr>
<td>
<code>customLabelSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CustomLabelSpec">
[]CustomLabelSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used for custom label tags which you want to add to the component resources.</p>
</td>
</tr>
<tr>
<td>
<code>switchoverSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SwitchoverSpec">
SwitchoverSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines command to do switchover.
In particular, when workloadType=Replication, the command defined in switchoverSpec will only be executed under
the condition of cluster.componentSpecs[x].SwitchPolicy.type=Noop.</p>
</td>
</tr>
<tr>
<td>
<code>postStartSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PostStartAction">
PostStartAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the command to be executed when the component is ready, and the command will only be executed once after
the component becomes ready.</p>
</td>
</tr>
<tr>
<td>
<code>volumeProtectionSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VolumeProtectionSpec">
VolumeProtectionSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines settings to do volume protect.</p>
</td>
</tr>
<tr>
<td>
<code>componentDefRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentDefRef">
[]ComponentDefRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to inject values from other components into the current component. Values will be saved and updated in a
configmap and mounted to the current component.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefDeclarations</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefDeclaration">
[]ServiceRefDeclaration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to declare the service reference of the current component.</p>
</td>
</tr>
<tr>
<td>
<code>exporter</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Exporter">
Exporter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the metrics exporter.</p>
</td>
</tr>
<tr>
<td>
<code>monitor</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MonitorConfig">
MonitorConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Deprecated since v0.9
monitor is monitoring config which provided by provider.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentPhase">ClusterComponentPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentStatus">ClusterComponentStatus</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentStatus">ComponentStatus</a>)
</p>
<div>
<p>ClusterComponentPhase defines the phase of a cluster component as represented in cluster.status.components.phase field.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Abnormal&#34;</p></td>
<td><p>AbnormalClusterCompPhase indicates the component has more than zero replicas, but there are some failed pods.
The component is functioning, but it is in a fragile state.</p>
</td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td><p>CreatingClusterCompPhase indicates the component is being created.</p>
</td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td><p>DeletingClusterCompPhase indicates the component is currently being deleted.</p>
</td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td><p>FailedClusterCompPhase indicates the component has more than zero replicas, but there are some failed pods.
The component is not functioning.</p>
</td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td><p>RunningClusterCompPhase indicates the component has more than zero replicas, and all pods are up-to-date and
in a &lsquo;Running&rsquo; state.</p>
</td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td><p>StoppedClusterCompPhase indicates the component has zero replicas, and all pods have been deleted.</p>
</td>
</tr><tr><td><p>&#34;Stopping&#34;</p></td>
<td><p>StoppingClusterCompPhase indicates the component has zero replicas, and there are pods that are terminating.</p>
</td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td><p>UpdatingClusterCompPhase indicates the component has more than zero replicas, and there are no failed pods,
it is currently being updated.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentService">ClusterComponentService
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>References the ComponentService name defined in the <code>componentDefinition.spec.services[*].name</code>.</p>
</td>
</tr>
<tr>
<td>
<code>serviceType</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#servicetype-v1-core">
Kubernetes core/v1.ServiceType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines how the Service is exposed. Valid options are <code>ClusterIP</code>, <code>NodePort</code>, and <code>LoadBalancer</code>.</p>
<ul>
<li><code>ClusterIP</code> allocates a Cluster-internal IP address for load-balancing to endpoints.
Endpoints are determined by the selector or if that is not specified,
they are determined by manual construction of an Endpoints object or EndpointSlice objects.</li>
<li><code>NodePort</code> builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the ClusterIP.</li>
<li><code>LoadBalancer</code> builds on NodePort and creates an external load-balancer (if supported in the current cloud)
which routes to the same endpoints as the ClusterIP.</li>
</ul>
<p>Note: although K8s Service type allows the &lsquo;ExternalName&rsquo; type, it is not a valid option for ClusterComponentService.</p>
<p>For more info, see:
<a href="https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types">https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types</a>.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If ServiceType is LoadBalancer, cloud provider related parameters can be put here.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer">https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer</a>.</p>
</td>
</tr>
<tr>
<td>
<code>podService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to generate individual Services for each Pod.
If set to true, a separate Service will be created for each Pod in the Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ShardingSpec">ShardingSpec</a>)
</p>
<div>
<p>ClusterComponentSpec defines the specification of a Component within a Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the Component&rsquo;s name.
It&rsquo;s part of the Service DNS name and must comply with the IANA service naming rule.
The name is optional when ClusterComponentSpec is used as a template (e.g., in <code>shardingSpec</code>),
but required otherwise.</p>
</td>
</tr>
<tr>
<td>
<code>componentDefRef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>References a ClusterComponentDefinition defined in the <code>clusterDefinition.spec.componentDef</code> field.
Must comply with the IANA service naming rule.</p>
<p>Deprecated since v0.9,
because defining Components in <code>clusterDefinition.spec.componentDef</code> field has been deprecated.
This field is replaced by the <code>componentDef</code> field, use <code>componentDef</code> instead.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>componentDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the exact name, name prefix, or regular expression pattern for matching the name of the ComponentDefinition
custom resource (CR) that defines the Component&rsquo;s characteristics and behavior.</p>
<p>If both <code>componentDefRef</code> and <code>componentDef</code> are provided,
the <code>componentDef</code> will take precedence over <code>componentDefRef</code>.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this Component.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).
If no version is specified, the latest available version will be used.</p>
</td>
</tr>
<tr>
<td>
<code>classDefRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClassDefRef">
ClassDefRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>References the class defined in ComponentClassDefinition.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRef">
[]ServiceRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of ServiceRef for a Component, enabling access to both external services and
Services provided by other Clusters.</p>
<p>Types of services:</p>
<ul>
<li>External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;
Require a ServiceDescriptor for connection details.</li>
<li>Services provided by a Cluster: Managed by the same KubeBlocks operator;
identified using Cluster, Component and Service names.</li>
</ul>
<p>ServiceRefs with identical <code>serviceRef.name</code> in the same Cluster are considered the same.</p>
<p>Example:</p>
<pre><code class="language-yaml">serviceRefs:
  - name: &quot;redis-sentinel&quot;
    serviceDescriptor:
      name: &quot;external-redis-sentinel&quot;
  - name: &quot;postgres-cluster&quot;
    clusterServiceSelector:
      cluster: &quot;my-postgres-cluster&quot;
      service:
        component: &quot;postgresql&quot;
</code></pre>
<p>The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>enabledLogs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies which types of logs should be collected for the Component.
The log types are defined in the <code>componentDefinition.spec.logConfigs</code> field with the LogConfig entries.</p>
<p>The elements in the <code>enabledLogs</code> array correspond to the names of the LogConfig entries.
For example, if the <code>componentDefinition.spec.logConfigs</code> defines LogConfig entries with
names &ldquo;slow_query_log&rdquo; and &ldquo;error_log&rdquo;,
you can enable the collection of these logs by including their names in the <code>enabledLogs</code> array:</p>
<pre><code class="language-yaml">enabledLogs:
- slow_query_log
- error_log
</code></pre>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Labels to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Annotations to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of environment variables to add.
These environment variables will be placed after the environment variables declared in the Pod.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Affinity">
Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a group of affinity scheduling rules for the Component.
It allows users to control how the Component&rsquo;s Pods are scheduled onto nodes in the K8s cluster.</p>
<p>Deprecated since v0.10, replaced by the <code>schedulingPolicy</code> field.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows Pods to be scheduled onto nodes with matching taints.
Each toleration in the array allows the Pod to tolerate node taints based on
specified <code>key</code>, <code>value</code>, <code>effect</code>, and <code>operator</code>.</p>
<ul>
<li>The <code>key</code>, <code>value</code>, and <code>effect</code> identify the taint that the toleration matches.</li>
<li>The <code>operator</code> determines how the toleration matches the taint.</li>
</ul>
<p>Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.</p>
<p>Deprecated since v0.10, replaced by the <code>schedulingPolicy</code> field.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources required by the Component.
It allows defining the CPU, memory requirements and limits for the Component&rsquo;s containers.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentVolumeClaimTemplate">
[]ClusterComponentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes to override.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentService">
[]ClusterComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentSystemAccount">
[]ComponentSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides system accounts defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentConfig">
[]ClusterComponentConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration content of a config template.</p>
</td>
</tr>
<tr>
<td>
<code>switchPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterSwitchPolicy">
ClusterSwitchPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the strategy for switchover and failover.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)
for secure communication.
When set to true, the Component will be configured to use TLS encryption for its network connections.
This ensures that the data transmitted between the Component and its clients or other Components is encrypted
and protected from unauthorized access.
If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys,
to properly set up the secure communication channel.</p>
</td>
</tr>
<tr>
<td>
<code>issuer</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Issuer">
Issuer
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration for the TLS certificates issuer.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.
Required when TLS is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceAccount required by the running Component.
This ServiceAccount is used to grant necessary permissions for the Component&rsquo;s Pods to interact
with other Kubernetes resources, such as modifying Pod labels or sending events.</p>
<p>Defaults:
To perform certain operational tasks, agent sidecars running in Pods require specific RBAC permissions.
The service account will be bound to a default role named &ldquo;kubeblocks-cluster-pod-role&rdquo; which is installed together with KubeBlocks.
If not specified, KubeBlocks automatically assigns a default ServiceAccount named &ldquo;kb-&#123;cluster.name&#125;&rdquo;</p>
<p>Future Changes:
Future versions might change the default ServiceAccount creation strategy to one per Component,
potentially revising the naming to &ldquo;kb-&#123;cluster.name&#125;-&#123;component.name&#125;&rdquo;.</p>
<p>Users can override the automatic ServiceAccount assignment by explicitly setting the name of
an existed ServiceAccount in this field.</p>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the update strategy for the Component.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the InstanceUpdateStrategy that will be
employed to update Pods in the InstanceSet when a revision is made to
Template.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>userResourceRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.UserResourceRefs">
UserResourceRefs
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows users to specify custom ConfigMaps and Secrets to be mounted as volumes
in the Cluster&rsquo;s Pods.
This is useful in scenarios where users need to provide additional resources to the Cluster, such as:</p>
<ul>
<li>Mounting custom scripts or configuration files during Cluster startup.</li>
<li>Mounting Secrets as volumes to provide sensitive information, like S3 AK/SK, to the Cluster.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows for the customization of configuration values for each instance within a Component.
An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).
While instances typically share a common configuration as defined in the ClusterComponentSpec,
they can require unique settings in various scenarios:</p>
<p>For example:
- A database Component might require different resource allocations for primary and secondary instances,
  with primaries needing more resources.
- During a rolling upgrade, a Component may first update the image for one or a few instances,
and then update the remaining instances after verifying that the updated instances are functioning correctly.</p>
<p>InstanceTemplate allows for specifying these unique configurations per instance.
Each instance&rsquo;s name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),
starting with an ordinal of 0.
It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the Cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>disableExporter</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines whether metrics exporter information is annotated on the Component&rsquo;s headless Service.</p>
<p>If set to true, the following annotations will not be patched into the Service:</p>
<ul>
<li>&ldquo;monitor.kubeblocks.io/path&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/port&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/scheme&rdquo;</li>
</ul>
<p>These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.</p>
</td>
</tr>
<tr>
<td>
<code>monitor</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Deprecated since v0.9
Determines whether metrics exporter information is annotated on the Component&rsquo;s headless Service.</p>
<p>If set to true, the following annotations will be patched into the Service:</p>
<ul>
<li>&ldquo;monitor.kubeblocks.io/path&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/port&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/scheme&rdquo;</li>
</ul>
<p>These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.</p>
</td>
</tr>
<tr>
<td>
<code>stop</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Stop the Component.
If set, all the computing resources will be released.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentStatus">ClusterComponentStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterStatus">ClusterStatus</a>)
</p>
<div>
<p>ClusterComponentStatus records Component status.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentPhase">
ClusterComponentPhase
</a>
</em>
</td>
<td>
<p>Specifies the current state of the Component.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentMessageMap">
ComponentMessageMap
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Records detailed information about the Component in its current phase.
The keys are either podName, deployName, or statefulSetName, formatted as &lsquo;ObjectKind/Name&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>podsReady</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Checks if all Pods of the Component are ready.</p>
</td>
</tr>
<tr>
<td>
<code>podsReadyTime</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the time when all Component Pods became ready.
This is the readiness time of the last Component Pod.</p>
</td>
</tr>
<tr>
<td>
<code>membersStatus</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MemberStatus">
[]MemberStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the status of the members.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterComponentVolumeClaimTemplate">ClusterComponentVolumeClaimTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.InstanceTemplate">InstanceTemplate</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Refers to the name of a volumeMount defined in either:</p>
<ul>
<li><code>componentDefinition.spec.runtime.containers[*].volumeMounts</code></li>
<li><code>clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts</code> (deprecated)</li>
</ul>
<p>The value of <code>name</code> must match the <code>name</code> field of a volumeMount specified in the corresponding <code>volumeMounts</code> array.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the labels for the PVC of the volume.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the annotations for the PVC of the volume.</p>
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PersistentVolumeClaimSpec">
PersistentVolumeClaimSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume
with the mount name specified in the <code>name</code> field.</p>
<p>When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification
defined in the <code>spec</code> field. The PVC will be associated with the volume mount specified by the <code>name</code> field.</p>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>accessModes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeaccessmode-v1-core">
[]Kubernetes core/v1.PersistentVolumeAccessMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Contains the desired access modes the volume should have.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1</a>.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volumeresourcerequirements-v1-core">
Kubernetes core/v1.VolumeResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the minimum resources the volume should have.
If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that
are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources">https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources</a>.</p>
</td>
</tr>
<tr>
<td>
<code>storageClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the StorageClass required by the claim.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1</a>.</p>
</td>
</tr>
<tr>
<td>
<code>volumeMode</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumemode-v1-core">
Kubernetes core/v1.PersistentVolumeMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines what type of volume is required by the claim, either Block or Filesystem.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbe">ClusterDefinitionProbe
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbes">ClusterDefinitionProbes</a>)
</p>
<div>
<p>ClusterDefinitionProbe is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>periodSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<p>How often (in seconds) to perform the probe.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Number of seconds after which the probe times out. Defaults to 1 second.</p>
</td>
</tr>
<tr>
<td>
<code>failureThreshold</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Minimum consecutive failures for the probe to be considered failed after having succeeded.</p>
</td>
</tr>
<tr>
<td>
<code>commands</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbeCMDs">
ClusterDefinitionProbeCMDs
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Commands used to execute for probe.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbeCMDs">ClusterDefinitionProbeCMDs
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbe">ClusterDefinitionProbe</a>)
</p>
<div>
<p>ClusterDefinitionProbeCMDs is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>writes</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines write checks that are executed on the probe sidecar.</p>
</td>
</tr>
<tr>
<td>
<code>queries</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines read checks that are executed on the probe sidecar.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbes">ClusterDefinitionProbes
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>ClusterDefinitionProbes is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>runningProbe</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbe">
ClusterDefinitionProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the probe used for checking the running status of the component.</p>
</td>
</tr>
<tr>
<td>
<code>statusProbe</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbe">
ClusterDefinitionProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the probe used for checking the status of the component.</p>
</td>
</tr>
<tr>
<td>
<code>roleProbe</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionProbe">
ClusterDefinitionProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the probe used for checking the role of the component.</p>
</td>
</tr>
<tr>
<td>
<code>roleProbeTimeoutAfterPodsReady</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the timeout (in seconds) for the role probe after all pods of the component are ready.
The system will check if the application is available in the pod.
If pods exceed the InitializationTimeoutSeconds time without a role label, this component will enter the
Failed/Abnormal phase.</p>
<p>Note that this configuration will only take effect if the component supports RoleProbe
and will not affect the life cycle of the pod. default values are 60 seconds.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterDefinitionSpec">ClusterDefinitionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinition">ClusterDefinition</a>)
</p>
<div>
<p>ClusterDefinitionSpec defines the desired state of ClusterDefinition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the well-known database type, such as mysql, redis, or mongodb.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>componentDefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">
[]ClusterComponentDefinition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides the definitions for the cluster components.</p>
<p>Deprecated since v0.9.
Components should now be individually defined using ComponentDefinition and
collectively referenced via <code>topology.components</code>.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>connectionCredential</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Connection credential template used for creating a connection credential secret for cluster objects.</p>
<p>Built-in objects are:</p>
<ul>
<li><code>$(RANDOM_PASSWD)</code> random 8 characters.</li>
<li><code>$(STRONG_RANDOM_PASSWD)</code> random 16 characters, with mixed cases, digits and symbols.</li>
<li><code>$(UUID)</code> generate a random UUID v4 string.</li>
<li><code>$(UUID_B64)</code> generate a random UUID v4 BASE64 encoded string.</li>
<li><code>$(UUID_STR_B64)</code> generate a random UUID v4 string then BASE64 encoded.</li>
<li><code>$(UUID_HEX)</code> generate a random UUID v4 HEX representation.</li>
<li><code>$(HEADLESS_SVC_FQDN)</code> headless service FQDN placeholder, value pattern is <code>$(CLUSTER_NAME)-$(1ST_COMP_NAME)-headless.$(NAMESPACE).svc</code>,
where 1ST_COMP_NAME is the 1st component that provide <code>ClusterDefinition.spec.componentDefs[].service</code> attribute;</li>
<li><code>$(SVC_FQDN)</code> service FQDN placeholder, value pattern is <code>$(CLUSTER_NAME)-$(1ST_COMP_NAME).$(NAMESPACE).svc</code>,
where 1ST_COMP_NAME is the 1st component that provide <code>ClusterDefinition.spec.componentDefs[].service</code> attribute;</li>
<li><code>$(SVC_PORT_&#123;PORT-NAME&#125;)</code> is ServicePort&rsquo;s port value with specified port name, i.e, a servicePort JSON struct:
<code>&#123;&quot;name&quot;: &quot;mysql&quot;, &quot;targetPort&quot;: &quot;mysqlContainerPort&quot;, &quot;port&quot;: 3306&#125;</code>, and <code>$(SVC_PORT_mysql)</code> in the
connection credential value is 3306.</li>
</ul>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>topologies</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterTopology">
[]ClusterTopology
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Topologies defines all possible topologies within the cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterDefinitionStatus">ClusterDefinitionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinition">ClusterDefinition</a>)
</p>
<div>
<p>ClusterDefinitionStatus defines the observed state of ClusterDefinition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the most recent generation observed for this ClusterDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Phase">
Phase
</a>
</em>
</td>
<td>
<p>Specifies the current phase of the ClusterDefinition. Valid values are <code>empty</code>, <code>Available</code>, <code>Unavailable</code>.
When <code>Available</code>, the ClusterDefinition is ready and can be referenced by related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
<tr>
<td>
<code>topologies</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Topologies this ClusterDefinition supported.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefs</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The service references declared by this ClusterDefinition.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterNetwork">ClusterNetwork
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ClusterNetwork is deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>hostNetworkAccessible</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether the host network can be accessed. By default, this is set to false.</p>
</td>
</tr>
<tr>
<td>
<code>publiclyAccessible</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether the network is accessible to the public. By default, this is set to false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterObjectReference">ClusterObjectReference
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentVarSelector">ComponentVarSelector</a>, <a href="#apps.kubeblocks.io/v1alpha1.CredentialVarSelector">CredentialVarSelector</a>, <a href="#apps.kubeblocks.io/v1alpha1.HostNetworkVarSelector">HostNetworkVarSelector</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceRefVarSelector">ServiceRefVarSelector</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceVarSelector">ServiceVarSelector</a>)
</p>
<div>
<p>ClusterObjectReference defines information to let you locate the referenced object inside the same Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the exact name, name prefix, or regular expression pattern for matching the name of the ComponentDefinition
custom resource (CR) used by the component that the referent object resident in.</p>
<p>If not specified, the component itself will be used.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Name of the referent object.</p>
</td>
</tr>
<tr>
<td>
<code>optional</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specify whether the object must be defined.</p>
</td>
</tr>
<tr>
<td>
<code>multipleClusterObjectOption</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectOption">
MultipleClusterObjectOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>This option defines the behavior when multiple component objects match the specified @CompDef.
If not provided, an error will be raised when handling multiple matches.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterPhase">ClusterPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterStatus">ClusterStatus</a>)
</p>
<div>
<p>ClusterPhase defines the phase of the Cluster within the .status.phase field.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Abnormal&#34;</p></td>
<td><p>AbnormalClusterPhase represents some components are in <code>Failed</code> or <code>Abnormal</code> phase, indicates that the cluster
is in a fragile state and troubleshooting is required.</p>
</td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td><p>CreatingClusterPhase represents all components are in <code>Creating</code> phase.</p>
</td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td><p>DeletingClusterPhase indicates the cluster is being deleted.</p>
</td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td><p>FailedClusterPhase represents all components are in <code>Failed</code> phase, indicates that the cluster is unavailable.</p>
</td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td><p>RunningClusterPhase represents all components are in <code>Running</code> phase, indicates that the cluster is functioning properly.</p>
</td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td><p>StoppedClusterPhase represents all components are in <code>Stopped</code> phase, indicates that the cluster has stopped and
is not providing any functionality.</p>
</td>
</tr><tr><td><p>&#34;Stopping&#34;</p></td>
<td><p>StoppingClusterPhase represents at least one component is in <code>Stopping</code> phase, indicates that the cluster is in
the process of stopping.</p>
</td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td><p>UpdatingClusterPhase represents all components are in <code>Creating</code>, <code>Running</code> or <code>Updating</code> phase, and at least one
component is in <code>Creating</code> or <code>Updating</code> phase, indicates that the cluster is undergoing an update.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterResources">ClusterResources
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ClusterResources is deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cpu</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#quantity-resource-core">
Kubernetes resource.Quantity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the amount of CPU resource the Cluster needs.
For more information, refer to: <a href="https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/">https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/</a></p>
</td>
</tr>
<tr>
<td>
<code>memory</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#quantity-resource-core">
Kubernetes resource.Quantity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the amount of memory resource the Cluster needs.
For more information, refer to: <a href="https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/">https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterService">ClusterService
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ClusterService defines a service that is exposed externally, allowing entities outside the cluster to access it.
For example, external applications, or other Clusters.
And another Cluster managed by the same KubeBlocks operator can resolve the address exposed by a ClusterService
using the <code>serviceRef</code> field.</p>
<p>When a Component needs to access another Cluster&rsquo;s ClusterService using the <code>serviceRef</code> field,
it must also define the service type and version information in the <code>componentDefinition.spec.serviceRefDeclarations</code>
section.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Service</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Service">
Service
</a>
</em>
</td>
<td>
<p>
(Members of <code>Service</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>shardingSelector</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extends the ServiceSpec.Selector by allowing the specification of a sharding name, which is defined in
<code>cluster.spec.shardingSpecs[*].name</code>, to be used as a selector for the service.
Note that this and the <code>componentSelector</code> are mutually exclusive and cannot be set simultaneously.</p>
</td>
</tr>
<tr>
<td>
<code>componentSelector</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extends the ServiceSpec.Selector by allowing the specification of a component, to be used as a selector for the service.
Note that this and the <code>shardingSelector</code> are mutually exclusive and cannot be set simultaneously.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Cluster">Cluster</a>)
</p>
<div>
<p>ClusterSpec defines the desired state of Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterDefinitionRef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterDefinition to use when creating a Cluster.</p>
<p>This field enables users to create a Cluster based on a specific ClusterDefinition.
Which, in conjunction with the <code>topology</code> field, determine:</p>
<ul>
<li>The Components to be included in the Cluster.</li>
<li>The sequences in which the Components are created, updated, and terminate.</li>
</ul>
<p>This facilitates multiple-components management with predefined ClusterDefinition.</p>
<p>Users with advanced requirements can bypass this general setting and specify more precise control over
the composition of the Cluster by directly referencing specific ComponentDefinitions for each component
within <code>componentSpecs[*].componentDef</code>.</p>
<p>If this field is not provided, each component must be explicitly defined in <code>componentSpecs[*].componentDef</code>.</p>
<p>Note: Once set, this field cannot be modified; it is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>clusterVersionRef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the ClusterVersion name.</p>
<p>Deprecated since v0.9, use ComponentVersion instead.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>topology</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ClusterTopology to be used when creating the Cluster.</p>
<p>This field defines which set of Components, as outlined in the ClusterDefinition, will be used to
construct the Cluster based on the named topology.
The ClusterDefinition may list multiple topologies under <code>clusterdefinition.spec.topologies[*]</code>,
each tailored to different use cases or environments.</p>
<p>If <code>topology</code> is not specified, the Cluster will use the default topology defined in the ClusterDefinition.</p>
<p>Note: Once set during the Cluster creation, the <code>topology</code> field cannot be modified.
It establishes the initial composition and structure of the Cluster and is intended for one-time configuration.</p>
</td>
</tr>
<tr>
<td>
<code>terminationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TerminationPolicyType">
TerminationPolicyType
</a>
</em>
</td>
<td>
<p>Specifies the behavior when a Cluster is deleted.
It defines how resources, data, and backups associated with a Cluster are managed during termination.
Choose a policy based on the desired level of resource cleanup and data preservation:</p>
<ul>
<li><code>DoNotTerminate</code>: Prevents deletion of the Cluster. This policy ensures that all resources remain intact.</li>
<li><code>Halt</code>: Deletes Cluster resources like Pods and Services but retains Persistent Volume Claims (PVCs),
allowing for data preservation while stopping other operations.
Warning: Halt policy is deprecated in 0.9.1 and will have same meaning as DoNotTerminate.</li>
<li><code>Delete</code>: Extends the <code>Halt</code> policy by also removing PVCs, leading to a thorough cleanup while
removing all persistent data.</li>
<li><code>WipeOut</code>: An aggressive policy that deletes all Cluster resources, including volume snapshots and
backups in external storage.
This results in complete data removal and should be used cautiously, primarily in non-production environments
to avoid irreversible data loss.</li>
</ul>
<p>Warning: Choosing an inappropriate termination policy can result in data loss.
The <code>WipeOut</code> policy is particularly risky in production environments due to its irreversible nature.</p>
</td>
</tr>
<tr>
<td>
<code>shardingSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ShardingSpec">
[]ShardingSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ShardingSpec objects that manage the sharding topology for Cluster Components.
Each ShardingSpec organizes components into shards, with each shard corresponding to a Component.
Components within a shard are all based on a common ClusterComponentSpec template, ensuring uniform configurations.</p>
<p>This field supports dynamic resharding by facilitating the addition or removal of shards
through the <code>shards</code> field in ShardingSpec.</p>
<p>Note: <code>shardingSpecs</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>componentSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">
[]ClusterComponentSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of ClusterComponentSpec objects used to define the individual Components that make up a Cluster.
This field allows for detailed configuration of each Component within the Cluster.</p>
<p>Note: <code>shardingSpecs</code> and <code>componentSpecs</code> cannot both be empty; at least one must be defined to configure a Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterService">
[]ClusterService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of additional Services that are exposed by a Cluster.
This field allows Services of selected Components, either from <code>componentSpecs</code> or <code>shardingSpecs</code> to be exposed,
alongside Services defined with ComponentService.</p>
<p>Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Affinity">
Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of node affinity scheduling rules for the Cluster&rsquo;s Pods.
This field helps control the placement of Pods on nodes within the Cluster.</p>
<p>Deprecated since v0.10. Use the <code>schedulingPolicy</code> field instead.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An array that specifies tolerations attached to the Cluster&rsquo;s Pods,
allowing them to be scheduled onto nodes with matching taints.</p>
<p>Deprecated since v0.10. Use the <code>schedulingPolicy</code> field instead.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies runtimeClassName for all Pods managed by this Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterBackup">
ClusterBackup
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the backup configuration of the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>tenancy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TenancyType">
TenancyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Describes how Pods are distributed across node.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>availabilityPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.AvailabilityPolicyType">
AvailabilityPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Describes the availability policy, including zone, node, and none.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the replicas of the first componentSpec, if the replicas of the first componentSpec is specified,
this value will be ignored.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterResources">
ClusterResources
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources of the first componentSpec, if the resources of the first componentSpec is specified,
this value will be ignored.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterStorage">
ClusterStorage
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the storage of the first componentSpec, if the storage of the first componentSpec is specified,
this value will be ignored.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterNetwork">
ClusterNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The configuration of network.</p>
<p>Deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterStatus">ClusterStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Cluster">Cluster</a>)
</p>
<div>
<p>ClusterStatus defines the observed state of the Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>The most recent generation number of the Cluster object that has been observed by the controller.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterPhase">
ClusterPhase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The current phase of the Cluster includes:
<code>Creating</code>, <code>Running</code>, <code>Updating</code>, <code>Stopping</code>, <code>Stopped</code>, <code>Deleting</code>, <code>Failed</code>, <code>Abnormal</code>.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentStatus">
map[string]github.com/apecloud/kubeblocks/apis/apps/v1alpha1.ClusterComponentStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Records the current status information of all Components within the Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>clusterDefGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the generation number of the referenced ClusterDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a list of detailed status of the Cluster object.
Each condition in the list provides real-time information about certain aspect of the Cluster object.</p>
<p>This field is crucial for administrators and developers to monitor and respond to changes within the Cluster.
It provides a history of state transitions and a snapshot of the current state that can be used for
automated logic or direct inspection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterStorage">ClusterStorage
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ClusterStorage is deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>size</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#quantity-resource-core">
Kubernetes resource.Quantity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the amount of storage the Cluster needs.
For more information, refer to: <a href="https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/">https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterSwitchPolicy">ClusterSwitchPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>)
</p>
<div>
<p>ClusterSwitchPolicy defines the switch policy for a Cluster.</p>
<p>Deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SwitchPolicyType">
SwitchPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Type specifies the type of switch policy to be applied.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterTopology">ClusterTopology
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionSpec">ClusterDefinitionSpec</a>)
</p>
<div>
<p>ClusterTopology represents the definition for a specific cluster topology.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the unique identifier for the cluster topology.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterTopologyComponent">
[]ClusterTopologyComponent
</a>
</em>
</td>
<td>
<p>Components specifies the components in the topology.</p>
</td>
</tr>
<tr>
<td>
<code>orders</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterTopologyOrders">
ClusterTopologyOrders
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the sequence in which components within a cluster topology are
started, stopped, and upgraded.
This ordering is crucial for maintaining the correct dependencies and operational flow across components.</p>
</td>
</tr>
<tr>
<td>
<code>default</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Default indicates whether this topology serves as the default configuration.
When set to true, this topology is automatically used unless another is explicitly specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterTopologyComponent">ClusterTopologyComponent
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterTopology">ClusterTopology</a>)
</p>
<div>
<p>ClusterTopologyComponent defines a Component within a ClusterTopology.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the unique identifier of the component within the cluster topology.
It follows IANA Service naming rules and is used as part of the Service&rsquo;s DNS name.
The name must start with a lowercase letter, can contain lowercase letters, numbers,
and hyphens, and must end with a lowercase letter or number.</p>
<p>Cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the exact name, name prefix, or regular expression pattern for matching the name of the ComponentDefinition
custom resource (CR) that defines the Component&rsquo;s characteristics and behavior.</p>
<p>The system selects the ComponentDefinition CR with the latest version that matches the pattern.
This approach allows:</p>
<ol>
<li>Precise selection by providing the exact name of a ComponentDefinition CR.</li>
<li>Flexible and automatic selection of the most up-to-date ComponentDefinition CR
by specifying a name prefix or regular expression pattern.</li>
</ol>
<p>Once set, this field cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ClusterTopologyOrders">ClusterTopologyOrders
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterTopology">ClusterTopology</a>)
</p>
<div>
<p>ClusterTopologyOrders manages the lifecycle of components within a cluster by defining their provisioning,
terminating, and updating sequences.
It organizes components into stages or groups, where each group indicates a set of components
that can be managed concurrently.
These groups are processed sequentially, allowing precise control based on component dependencies and requirements.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>provision</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the order for creating and initializing components.
This is designed for components that depend on one another. Components without dependencies can be grouped together.</p>
<p>Components that can be provisioned independently or have no dependencies can be listed together in the same stage,
separated by commas.</p>
</td>
</tr>
<tr>
<td>
<code>terminate</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Outlines the order for stopping and deleting components.
This sequence is designed for components that require a graceful shutdown or have interdependencies.</p>
<p>Components that can be terminated independently or have no dependencies can be listed together in the same stage,
separated by commas.</p>
</td>
</tr>
<tr>
<td>
<code>update</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Update determines the order for updating components&rsquo; specifications, such as image upgrades or resource scaling.
This sequence is designed for components that have dependencies or require specific update procedures.</p>
<p>Components that can be updated independently or have no dependencies can be listed together in the same stage,
separated by commas.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CmdExecutorConfig">CmdExecutorConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.PostStartAction">PostStartAction</a>, <a href="#apps.kubeblocks.io/v1alpha1.SwitchoverAction">SwitchoverAction</a>, <a href="#apps.kubeblocks.io/v1alpha1.SystemAccountSpec">SystemAccountSpec</a>)
</p>
<div>
<p>CmdExecutorConfig specifies how to perform creation and deletion statements.</p>
<p>Deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>CommandExecutorEnvItem</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CommandExecutorEnvItem">
CommandExecutorEnvItem
</a>
</em>
</td>
<td>
<p>
(Members of <code>CommandExecutorEnvItem</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>CommandExecutorItem</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CommandExecutorItem">
CommandExecutorItem
</a>
</em>
</td>
<td>
<p>
(Members of <code>CommandExecutorItem</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CommandExecutorEnvItem">CommandExecutorEnvItem
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.CmdExecutorConfig">CmdExecutorConfig</a>)
</p>
<div>
<p>CommandExecutorEnvItem is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the image used to execute the command.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A list of environment variables that will be injected into the command execution context.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CommandExecutorItem">CommandExecutorItem
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.CmdExecutorConfig">CmdExecutorConfig</a>)
</p>
<div>
<p>CommandExecutorItem is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>command</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>The command to be executed.</p>
</td>
</tr>
<tr>
<td>
<code>args</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional parameters used in the execution of the command.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">ComponentConfigSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetail">ConfigurationItemDetail</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ComponentTemplateSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentTemplateSpec">
ComponentTemplateSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ComponentTemplateSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>keys</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration files within the ConfigMap that support dynamic updates.</p>
<p>A configuration template (provided in the form of a ConfigMap) may contain templates for multiple
configuration files.
Each configuration file corresponds to a key in the ConfigMap.
Some of these configuration files may support dynamic modification and reloading without requiring
a pod restart.</p>
<p>If empty or omitted, all configuration files in the ConfigMap are assumed to support dynamic updates,
and ConfigConstraint applies to all keys.</p>
</td>
</tr>
<tr>
<td>
<code>legacyRenderedConfigSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LegacyRenderedTemplateSpec">
LegacyRenderedTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the secondary rendered config spec for pod-specific customization.</p>
<p>The template is rendered inside the pod (by the &ldquo;config-manager&rdquo; sidecar container) and merged with the main
template&rsquo;s render result to generate the final configuration file.</p>
<p>This field is intended to handle scenarios where different pods within the same Component have
varying configurations. It allows for pod-specific customization of the configuration.</p>
<p>Note: This field will be deprecated in future versions, and the functionality will be moved to
<code>cluster.spec.componentSpecs[*].instances[*]</code>.</p>
</td>
</tr>
<tr>
<td>
<code>constraintRef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the referenced configuration constraints object.</p>
</td>
</tr>
<tr>
<td>
<code>asEnvFrom</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the containers to inject the ConfigMap parameters as environment variables.</p>
<p>This is useful when application images accept parameters through environment variables and
generate the final configuration file in the startup script based on these variables.</p>
<p>This field allows users to specify a list of container names, and KubeBlocks will inject the environment
variables converted from the ConfigMap into these designated containers. This provides a flexible way to
pass the configuration items from the ConfigMap to the container without modifying the image.</p>
<p>Deprecated: <code>asEnvFrom</code> has been deprecated since 0.9.0 and will be removed in 0.10.0.
Use <code>injectEnvTo</code> instead.</p>
</td>
</tr>
<tr>
<td>
<code>injectEnvTo</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the containers to inject the ConfigMap parameters as environment variables.</p>
<p>This is useful when application images accept parameters through environment variables and
generate the final configuration file in the startup script based on these variables.</p>
<p>This field allows users to specify a list of container names, and KubeBlocks will inject the environment
variables converted from the ConfigMap into these designated containers. This provides a flexible way to
pass the configuration items from the ConfigMap to the container without modifying the image.</p>
</td>
</tr>
<tr>
<td>
<code>reRenderResourceTypes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RerenderResourceType">
[]RerenderResourceType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the configuration needs to be re-rendered after v-scale or h-scale operations to reflect changes.</p>
<p>In some scenarios, the configuration may need to be updated to reflect the changes in resource allocation
or cluster topology. Examples:</p>
<ul>
<li>Redis: adjust maxmemory after v-scale operation.</li>
<li>MySQL: increase max connections after v-scale operation.</li>
<li>Zookeeper: update zoo.cfg with new node addresses after h-scale operation.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>asSecret</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Whether to store the final rendered parameters as a secret.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentDefRef">ComponentDefRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>ComponentDefRef is used to select the component and its fields to be referenced.</p>
<p>Deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>componentDefName</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the componentDef to be selected.</p>
</td>
</tr>
<tr>
<td>
<code>failurePolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.FailurePolicyType">
FailurePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the policy to be followed in case of a failure in finding the component.</p>
</td>
</tr>
<tr>
<td>
<code>componentRefEnv</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentRefEnv">
[]ComponentRefEnv
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The values that are to be injected as environment variables into each component.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinition">ComponentDefinition</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>provider</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the Component provider, typically the vendor or developer name.
It identifies the entity responsible for creating and maintaining the Component.</p>
<p>When specifying the provider name, consider the following guidelines:</p>
<ul>
<li>Keep the name concise and relevant to the Component.</li>
<li>Use a consistent naming convention across Components from the same provider.</li>
<li>Avoid using trademarked or copyrighted names without proper permission.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a brief and concise explanation of the Component&rsquo;s purpose, functionality, and any relevant details.
It serves as a quick reference for users to understand the Component&rsquo;s role and characteristics.</p>
</td>
</tr>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the type of well-known service protocol that the Component provides.
It specifies the standard or widely recognized protocol used by the Component to offer its Services.</p>
<p>The <code>serviceKind</code> field allows users to quickly identify the type of Service provided by the Component
based on common protocols or service types. This information helps in understanding the compatibility,
interoperability, and usage of the Component within a system.</p>
<p>Some examples of well-known service protocols include:</p>
<ul>
<li>&ldquo;MySQL&rdquo;: Indicates that the Component provides a MySQL database service.</li>
<li>&ldquo;PostgreSQL&rdquo;: Indicates that the Component offers a PostgreSQL database service.</li>
<li>&ldquo;Redis&rdquo;: Signifies that the Component functions as a Redis key-value store.</li>
<li>&ldquo;ETCD&rdquo;: Denotes that the Component serves as an ETCD distributed key-value store.</li>
</ul>
<p>The <code>serviceKind</code> value is case-insensitive, allowing for flexibility in specifying the protocol name.</p>
<p>When specifying the <code>serviceKind</code>, consider the following guidelines:</p>
<ul>
<li>Use well-established and widely recognized protocol names or service types.</li>
<li>Ensure that the <code>serviceKind</code> accurately represents the primary service type offered by the Component.</li>
<li>If the Component provides multiple services, choose the most prominent or commonly used protocol.</li>
<li>Limit the <code>serviceKind</code> to a maximum of 32 characters for conciseness and readability.</li>
</ul>
<p>Note: The <code>serviceKind</code> field is optional and can be left empty if the Component does not fit into a well-known
service category or if the protocol is not widely recognized. It is primarily used to convey information about
the Component&rsquo;s service type to users and facilitate discovery and integration.</p>
<p>The <code>serviceKind</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the version of the Service provided by the Component.
It follows the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
<p>The Semantic Versioning specification defines a version number format of X.Y.Z (MAJOR.MINOR.PATCH), where:</p>
<ul>
<li>X represents the major version and indicates incompatible API changes.</li>
<li>Y represents the minor version and indicates added functionality in a backward-compatible manner.</li>
<li>Z represents the patch version and indicates backward-compatible bug fixes.</li>
</ul>
<p>Additional labels for pre-release and build metadata are available as extensions to the X.Y.Z format:</p>
<ul>
<li>Use pre-release labels (e.g., -alpha, -beta) for versions that are not yet stable or ready for production use.</li>
<li>Use build metadata (e.g., +build.1) for additional version information if needed.</li>
</ul>
<p>Examples of valid ServiceVersion values:</p>
<ul>
<li>&ldquo;1.0.0&rdquo;</li>
<li>&ldquo;2.3.1&rdquo;</li>
<li>&ldquo;3.0.0-alpha.1&rdquo;</li>
<li>&ldquo;4.5.2+build.1&rdquo;</li>
</ul>
<p>The <code>serviceVersion</code> field is immutable and cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>runtime</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podspec-v1-core">
Kubernetes core/v1.PodSpec
</a>
</em>
</td>
<td>
<p>Specifies the PodSpec template used in the Component.
It includes the following elements:</p>
<ul>
<li>Init containers</li>
<li>Containers
<ul>
<li>Image</li>
<li>Commands</li>
<li>Args</li>
<li>Envs</li>
<li>Mounts</li>
<li>Ports</li>
<li>Security context</li>
<li>Probes</li>
<li>Lifecycle</li>
</ul></li>
<li>Volumes</li>
</ul>
<p>This field is intended to define static settings that remain consistent across all instantiated Components.
Dynamic settings such as CPU and memory resource limits, as well as scheduling settings (affinity,
toleration, priority), may vary among different instantiated Components.
They should be specified in the <code>cluster.spec.componentSpecs</code> (ClusterComponentSpec).</p>
<p>Specific instances of a Component may override settings defined here, such as using a different container image
or modifying environment variable values.
These instance-specific overrides can be specified in <code>cluster.spec.componentSpecs[*].instances</code>.</p>
<p>This field is immutable and cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>monitor</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MonitorConfig">
MonitorConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Deprecated since v0.9
monitor is monitoring config which provided by provider.</p>
</td>
</tr>
<tr>
<td>
<code>exporter</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Exporter">
Exporter
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the built-in metrics exporter container.</p>
</td>
</tr>
<tr>
<td>
<code>vars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.EnvVar">
[]EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines variables which are determined after Cluster instantiation and reflect
dynamic or runtime attributes of instantiated Clusters.
These variables serve as placeholders for setting environment variables in Pods and Actions,
or for rendering configuration and script templates before actual values are finalized.</p>
<p>These variables are placed in front of the environment variables declared in the Pod if used as
environment variables.</p>
<p>Variable values can be sourced from:</p>
<ul>
<li>ConfigMap: Select and extract a value from a specific key within a ConfigMap.</li>
<li>Secret: Select and extract a value from a specific key within a Secret.</li>
<li>HostNetwork: Retrieves values (including ports) from host-network resources.</li>
<li>Service: Retrieves values (including address, port, NodePort) from a selected Service.
Intended to obtain the address of a ComponentService within the same Cluster.</li>
<li>Credential: Retrieves account name and password from a SystemAccount variable.</li>
<li>ServiceRef: Retrieves address, port, account name and password from a selected ServiceRefDeclaration.
Designed to obtain the address bound to a ServiceRef, such as a ClusterService or
ComponentService of another cluster or an external service.</li>
<li>Component: Retrieves values from a selected Component, including replicas and instance name list.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVolume">
[]ComponentVolume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the volumes used by the Component and some static attributes of the volumes.
After defining the volumes here, user can reference them in the
<code>cluster.spec.componentSpecs[*].volumeClaimTemplates</code> field to configure dynamic properties such as
volume capacity and storage class.</p>
<p>This field allows you to specify the following:</p>
<ul>
<li>Snapshot behavior: Determines whether a snapshot of the volume should be taken when performing
a snapshot backup of the Component.</li>
<li>Disk high watermark: Sets the high watermark for the volume&rsquo;s disk usage.
When the disk usage reaches the specified threshold, it triggers an alert or action.</li>
</ul>
<p>By configuring these volume behaviors, you can control how the volumes are managed and monitored within the Component.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HostNetwork">
HostNetwork
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the host network configuration for the Component.</p>
<p>When <code>hostNetwork</code> option is enabled, the Pods share the host&rsquo;s network namespace and can directly access
the host&rsquo;s network interfaces.
This means that if multiple Pods need to use the same port, they cannot run on the same host simultaneously
due to port conflicts.</p>
<p>The DNSPolicy field in the Pod spec determines how containers within the Pod perform DNS resolution.
When using hostNetwork, the operator will set the DNSPolicy to &lsquo;ClusterFirstWithHostNet&rsquo;.
With this policy, DNS queries will first go through the K8s cluster&rsquo;s DNS service.
If the query fails, it will fall back to the host&rsquo;s DNS settings.</p>
<p>If set, the DNS policy will be automatically set to &ldquo;ClusterFirstWithHostNet&rdquo;.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines additional Services to expose the Component&rsquo;s endpoints.</p>
<p>A default headless Service, named <code>&#123;cluster.name&#125;-&#123;component.name&#125;-headless</code>, is automatically created
for internal Cluster communication.</p>
<p>This field enables customization of additional Services to expose the Component&rsquo;s endpoints to
other Components within the same or different Clusters, and to external applications.
Each Service entry in this list can include properties such as ports, type, and selectors.</p>
<ul>
<li>For intra-Cluster access, Components can reference Services using variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceVarRef</code>.</li>
<li>For inter-Cluster access, reference Services use variables declared in
<code>componentDefinition.spec.vars[*].valueFrom.serviceRefVarRef</code>,
and bind Services at Cluster creation time with <code>clusterComponentSpec.ServiceRef[*].clusterServiceSelector</code>.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">
[]ComponentConfigSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration file templates and volume mount parameters used by the Component.
It also includes descriptions of the parameters in the ConfigMaps, such as value range limitations.</p>
<p>This field specifies a list of templates that will be rendered into Component containers&rsquo; configuration files.
Each template is represented as a ConfigMap and may contain multiple configuration files,
with each file being a key in the ConfigMap.</p>
<p>The rendered configuration files will be mounted into the Component&rsquo;s containers
according to the specified volume mount parameters.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>logConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LogConfig">
[]LogConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the types of logs generated by instances of the Component and their corresponding file paths.
These logs can be collected for further analysis and monitoring.</p>
<p>The <code>logConfigs</code> field is an optional list of LogConfig objects, where each object represents
a specific log type and its configuration.
It allows you to specify multiple log types and their respective file paths for the Component.</p>
<p>Examples:</p>
<pre><code class="language-yaml"> logConfigs:
 - filePathPattern: /data/mysql/log/mysqld-error.log
   name: error
 - filePathPattern: /data/mysql/log/mysqld.log
   name: general
 - filePathPattern: /data/mysql/log/mysqld-slowquery.log
   name: slow
</code></pre>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>scripts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentTemplateSpec">
[]ComponentTemplateSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies groups of scripts, each provided via a ConfigMap, to be mounted as volumes in the container.
These scripts can be executed during container startup or via specific actions.</p>
<p>Each script group is encapsulated in a ComponentTemplateSpec that includes:</p>
<ul>
<li>The ConfigMap containing the scripts.</li>
<li>The mount point where the scripts will be mounted inside the container.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>policyRules</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#policyrule-v1-rbac">
[]Kubernetes rbac/v1.PolicyRule
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the namespaced policy rules required by the Component.</p>
<p>The <code>policyRules</code> field is an array of <code>rbacv1.PolicyRule</code> objects that define the policy rules
needed by the Component to operate within a namespace.
These policy rules determine the permissions and verbs the Component is allowed to perform on
Kubernetes resources within the namespace.</p>
<p>The purpose of this field is to automatically generate the necessary RBAC roles
for the Component based on the specified policy rules.
This ensures that the Pods in the Component has appropriate permissions to function.</p>
<p>Note: This field is currently non-functional and is reserved for future implementation.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static labels that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If a label key in the <code>labels</code> field conflicts with any system labels or user-specified labels,
it will be silently ignored to avoid overriding higher-priority labels.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies static annotations that will be patched to all Kubernetes resources created for the Component.</p>
<p>Note: If an annotation key in the <code>annotations</code> field conflicts with any system annotations
or user-specified annotations, it will be silently ignored to avoid overriding higher-priority annotations.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>replicasLimit</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReplicasLimit">
ReplicasLimit
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the upper limit of the number of replicas supported by the Component.</p>
<p>It defines the maximum number of replicas that can be created for the Component.
This field allows you to set a limit on the scalability of the Component, preventing it from exceeding a certain number of replicas.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SystemAccount">
[]SystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>An array of <code>SystemAccount</code> objects that define the system accounts needed
for the management operations of the Component.</p>
<p>Each <code>SystemAccount</code> includes:</p>
<ul>
<li>Account name.</li>
<li>The SQL statement template: Used to create the system account.</li>
<li>Password Source: Either generated based on certain rules or retrieved from a Secret.</li>
</ul>
<p>Use cases for system accounts typically involve tasks like system initialization, backups, monitoring,
health checks, replication, and other system-level operations.</p>
<p>System accounts are distinct from user accounts, although both are database accounts.</p>
<ul>
<li><strong>System Accounts</strong>: Created during Cluster setup by the KubeBlocks operator,
these accounts have higher privileges for system management and are fully managed
through a declarative API by the operator.</li>
<li><strong>User Accounts</strong>: Managed by users or administrator.
User account permissions should follow the principle of least privilege,
granting only the necessary access rights to complete their required tasks.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the concurrency strategy for updating multiple instances of the Component.
Available strategies:</p>
<ul>
<li><code>Serial</code>: Updates replicas one at a time, ensuring minimal downtime by waiting for each replica to become ready
before updating the next.</li>
<li><code>Parallel</code>: Updates all replicas simultaneously, optimizing for speed but potentially reducing availability
during the update.</li>
<li><code>BestEffortParallel</code>: Updates replicas concurrently with a limit on simultaneous updates to ensure a minimum
number of operational replicas for maintaining quorum.
 For example, in a 5-replica component, updating a maximum of 2 replicas simultaneously keeps
at least 3 operational for quorum.</li>
</ul>
<p>This field is immutable and defaults to &lsquo;Serial&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InstanceSet controls the creation of pods during initial scale up, replacement of pods on nodes, and scaling down.</p>
<ul>
<li><code>OrderedReady</code>: Creates pods in increasing order (pod-0, then pod-1, etc). The controller waits until each pod
is ready before continuing. Pods are removed in reverse order when scaling down.</li>
<li><code>Parallel</code>: Creates pods in parallel to match the desired scale without waiting. All pods are deleted at once
when scaling down.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enumerate all possible roles assigned to each replica of the Component, influencing its behavior.</p>
<p>A replica can have zero to multiple roles.
KubeBlocks operator determines the roles of each replica by invoking the <code>lifecycleActions.roleProbe</code> method.
This action returns a list of roles for each replica, and the returned roles must be predefined in the <code>roles</code> field.</p>
<p>The roles assigned to a replica can influence various aspects of the Component&rsquo;s behavior, such as:</p>
<ul>
<li>Service selection: The Component&rsquo;s exposed Services may target replicas based on their roles using <code>roleSelector</code>.</li>
<li>Update order: The roles can determine the order in which replicas are updated during a Component update.
For instance, replicas with a &ldquo;follower&rdquo; role can be updated first, while the replica with the &ldquo;leader&rdquo;
role is updated last. This helps minimize the number of leader changes during the update process.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>roleArbitrator</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RoleArbitrator">
RoleArbitrator
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>This field has been deprecated since v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>lifecycleActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentLifecycleActions">
ComponentLifecycleActions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a set of hooks and procedures that customize the behavior of a Component throughout its lifecycle.
Actions are triggered at specific lifecycle stages:</p>
<ul>
<li><code>postProvision</code>: Defines the hook to be executed after the creation of a Component,
with <code>preCondition</code> specifying when the action should be fired relative to the Component&rsquo;s lifecycle stages:
<code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.</li>
<li><code>preTerminate</code>: Defines the hook to be executed before terminating a Component.</li>
<li><code>roleProbe</code>: Defines the procedure which is invoked regularly to assess the role of replicas.</li>
<li><code>switchover</code>: Defines the procedure for a controlled transition of leadership from the current leader to a new replica.
This approach aims to minimize downtime and maintain availability in systems with a leader-follower topology,
such as before planned maintenance or upgrades on the current leader node.</li>
<li><code>memberJoin</code>: Defines the procedure to add a new replica to the replication group.</li>
<li><code>memberLeave</code>: Defines the method to remove a replica from the replication group.</li>
<li><code>readOnly</code>: Defines the procedure to switch a replica into the read-only state.</li>
<li><code>readWrite</code>: transition a replica from the read-only state back to the read-write state.</li>
<li><code>dataDump</code>: Defines the procedure to export the data from a replica.</li>
<li><code>dataLoad</code>: Defines the procedure to import data into a replica.</li>
<li><code>reconfigure</code>: Defines the procedure that update a replica with new configuration file.</li>
<li><code>accountProvision</code>: Defines the procedure to generate a new database account.</li>
</ul>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefDeclarations</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefDeclaration">
[]ServiceRefDeclaration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists external service dependencies of the Component, including services from other Clusters or outside the K8s environment.</p>
<p>This field is immutable.</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p><code>minReadySeconds</code> is the minimum duration in seconds that a new Pod should remain in the ready
state without any of its containers crashing to be considered available.
This ensures the Pod&rsquo;s stability and readiness to serve requests.</p>
<p>A default value of 0 seconds means the Pod is considered available as soon as it enters the ready state.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentDefinitionStatus">ComponentDefinitionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinition">ComponentDefinition</a>)
</p>
<div>
<p>ComponentDefinitionStatus defines the observed state of ComponentDefinition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the most recent generation that has been observed for the ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the current status of the ComponentDefinition. Valid values include `<code>,</code>Available<code>, and</code>Unavailable<code>.
When the status is</code>Available`, the ComponentDefinition is ready and can be utilized by related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the current phase.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentLifecycleActions">ComponentLifecycleActions
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ComponentLifecycleActions defines a collection of Actions for customizing the behavior of a Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>postProvision</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed after a component&rsquo;s creation.</p>
<p>By setting <code>postProvision.customHandler.preCondition</code>, you can determine the specific lifecycle stage
at which the action should trigger: <code>Immediately</code>, <code>RuntimeReady</code>, <code>ComponentReady</code>, and <code>ClusterReady</code>.
with <code>ComponentReady</code> being the default.</p>
<p>The PostProvision Action is intended to run only once.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li><p>KB_CLUSTER_POD_IP_LIST: Comma-separated list of the cluster&rsquo;s pod IP addresses (e.g., &ldquo;podIp1,podIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_POD_NAME_LIST: Comma-separated list of the cluster&rsquo;s pod names (e.g., &ldquo;pod1,pod2&rdquo;).</p></li>
<li><p>KB_CLUSTER_POD_HOST_NAME_LIST: Comma-separated list of host names, each corresponding to a pod in
KB_CLUSTER_POD_NAME_LIST (e.g., &ldquo;hostName1,hostName2&rdquo;).</p></li>
<li><p>KB_CLUSTER_POD_HOST_IP_LIST: Comma-separated list of host IP addresses, each corresponding to a pod in
KB_CLUSTER_POD_NAME_LIST (e.g., &ldquo;hostIp1,hostIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_NAME_LIST: Comma-separated list of all pod names within the component
(e.g., &ldquo;pod1,pod2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_IP_LIST: Comma-separated list of pod IP addresses,
matching the order of pods in KB_CLUSTER_COMPONENT_POD_NAME_LIST (e.g., &ldquo;podIp1,podIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_HOST_NAME_LIST: Comma-separated list of host names for each pod,
matching the order of pods in KB_CLUSTER_COMPONENT_POD_NAME_LIST (e.g., &ldquo;hostName1,hostName2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_HOST_IP_LIST: Comma-separated list of host IP addresses for each pod,
matching the order of pods in KB_CLUSTER_COMPONENT_POD_NAME_LIST (e.g., &ldquo;hostIp1,hostIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_LIST: Comma-separated list of all cluster components (e.g., &ldquo;comp1,comp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_DELETING_LIST: Comma-separated list of components that are currently being deleted
(e.g., &ldquo;comp1,comp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_UNDELETED_LIST: Comma-separated list of components that are not being deleted
(e.g., &ldquo;comp1,comp2&rdquo;).</p></li>
</ul>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>preTerminate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the hook to be executed prior to terminating a component.</p>
<p>The PreTerminate Action is intended to run only once.</p>
<p>This action is executed immediately when a scale-down operation for the Component is initiated.
The actual termination and cleanup of the Component and its associated resources will not proceed
until the PreTerminate action has completed successfully.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li><p>KB_CLUSTER_POD_IP_LIST: Comma-separated list of the cluster&rsquo;s pod IP addresses (e.g., &ldquo;podIp1,podIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_POD_NAME_LIST: Comma-separated list of the cluster&rsquo;s pod names (e.g., &ldquo;pod1,pod2&rdquo;).</p></li>
<li><p>KB_CLUSTER_POD_HOST_NAME_LIST: Comma-separated list of host names, each corresponding to a pod in
KB_CLUSTER_POD_NAME_LIST (e.g., &ldquo;hostName1,hostName2&rdquo;).</p></li>
<li><p>KB_CLUSTER_POD_HOST_IP_LIST: Comma-separated list of host IP addresses, each corresponding to a pod in
KB_CLUSTER_POD_NAME_LIST (e.g., &ldquo;hostIp1,hostIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_NAME_LIST: Comma-separated list of all pod names within the component
(e.g., &ldquo;pod1,pod2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_IP_LIST: Comma-separated list of pod IP addresses,
matching the order of pods in KB_CLUSTER_COMPONENT_POD_NAME_LIST (e.g., &ldquo;podIp1,podIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_HOST_NAME_LIST: Comma-separated list of host names for each pod,
matching the order of pods in KB_CLUSTER_COMPONENT_POD_NAME_LIST (e.g., &ldquo;hostName1,hostName2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_POD_HOST_IP_LIST: Comma-separated list of host IP addresses for each pod,
matching the order of pods in KB_CLUSTER_COMPONENT_POD_NAME_LIST (e.g., &ldquo;hostIp1,hostIp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_LIST: Comma-separated list of all cluster components (e.g., &ldquo;comp1,comp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_DELETING_LIST: Comma-separated list of components that are currently being deleted
(e.g., &ldquo;comp1,comp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_UNDELETED_LIST: Comma-separated list of components that are not being deleted
(e.g., &ldquo;comp1,comp2&rdquo;).</p></li>
<li><p>KB_CLUSTER_COMPONENT_IS_SCALING_IN: Indicates whether the component is currently scaling in.
If this variable is present and set to &ldquo;true&rdquo;, it denotes that the component is undergoing a scale-in operation.
During scale-in, data rebalancing is necessary to maintain cluster integrity.
Contrast this with a cluster deletion scenario where data rebalancing is not required as the entire cluster
is being cleaned up.</p></li>
</ul>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>roleProbe</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RoleProbe">
RoleProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure which is invoked regularly to assess the role of replicas.</p>
<p>This action is periodically triggered by Lorry at the specified interval to determine the role of each replica.
Upon successful execution, the action&rsquo;s output designates the role of the replica,
which should match one of the predefined role names within <code>componentDefinition.spec.roles</code>.
The output is then compared with the previous successful execution result.
If a role change is detected, an event is generated to inform the controller,
which initiates an update of the replica&rsquo;s role.</p>
<p>Defining a RoleProbe Action for a Component is required if roles are defined for the Component.
It ensures replicas are correctly labeled with their respective roles.
Without this, services that rely on roleSelectors might improperly direct traffic to wrong replicas.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_POD_FQDN: The FQDN of the Pod whose role is being assessed.</li>
<li>KB_SERVICE_PORT: The port used by the database service.</li>
<li>KB_SERVICE_USER: The username with the necessary permissions to interact with the database service.</li>
<li>KB_SERVICE_PASSWORD: The corresponding password for KB_SERVICE_USER to authenticate with the database service.</li>
</ul>
<p>Expected output of this action:
- On Success: The determined role of the replica, which must align with one of the roles specified
  in the component definition.
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>switchover</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentSwitchover">
ComponentSwitchover
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure for a controlled transition of leadership from the current leader to a new replica.
This approach aims to minimize downtime and maintain availability in systems with a leader-follower topology,
during events such as planned maintenance or when performing stop, shutdown, restart, or upgrade operations
involving the current leader node.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_SWITCHOVER_CANDIDATE_NAME: The name of the pod for the new leader candidate, which may not be specified (empty).</li>
<li>KB_SWITCHOVER_CANDIDATE_FQDN: The FQDN of the new leader candidate&rsquo;s pod, which may not be specified (empty).</li>
<li>KB_LEADER_POD_IP: The IP address of the current leader&rsquo;s pod prior to the switchover.</li>
<li>KB_LEADER_POD_NAME: The name of the current leader&rsquo;s pod prior to the switchover.</li>
<li>KB_LEADER_POD_FQDN: The FQDN of the current leader&rsquo;s pod prior to the switchover.</li>
</ul>
<p>The environment variables with the following prefixes are deprecated and will be removed in future releases:</p>
<ul>
<li>KB_REPLICATION_PRIMARY<em>POD</em></li>
<li>KB_CONSENSUS_LEADER<em>POD</em></li>
</ul>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>memberJoin</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to add a new replica to the replication group.</p>
<p>This action is initiated after a replica pod becomes ready.</p>
<p>The role of the replica (e.g., primary, secondary) will be determined and assigned as part of the action command
implementation, or automatically by the database kernel or a sidecar utility like Patroni that implements
a consensus algorithm.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_SERVICE_PORT: The port used by the database service.</li>
<li>KB_SERVICE_USER: The username with the necessary permissions to interact with the database service.</li>
<li>KB_SERVICE_PASSWORD: The corresponding password for KB_SERVICE_USER to authenticate with the database service.</li>
<li>KB_PRIMARY_POD_FQDN: The FQDN of the primary Pod within the replication group.</li>
<li>KB_MEMBER_ADDRESSES: A comma-separated list of Pod addresses for all replicas in the group.</li>
<li>KB_NEW_MEMBER_POD_NAME: The pod name of the replica being added to the group.</li>
<li>KB_NEW_MEMBER_POD_IP: The IP address of the replica being added to the group.</li>
</ul>
<p>Expected action output:
- On Failure: An error message detailing the reason for any failure encountered
during the addition of the new member.</p>
<p>For example, to add a new OBServer to an OceanBase Cluster in &lsquo;zone1&rsquo;, the following command may be used:</p>
<pre><code class="language-yaml">command:
- bash
- -c
- |
   ADDRESS=$(KB_MEMBER_ADDRESSES%%,*)
   HOST=$(echo $ADDRESS | cut -d ':' -f 1)
   PORT=$(echo $ADDRESS | cut -d ':' -f 2)
   CLIENT=&quot;mysql -u $KB_SERVICE_USER -p$KB_SERVICE_PASSWORD -P $PORT -h $HOST -e&quot;
       $CLIENT &quot;ALTER SYSTEM ADD SERVER '$KB_NEW_MEMBER_POD_IP:$KB_SERVICE_PORT' ZONE 'zone1'&quot;
</code></pre>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>memberLeave</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to remove a replica from the replication group.</p>
<p>This action is initiated before remove a replica from the group.
The operator will wait for MemberLeave to complete successfully before releasing the replica and cleaning up
related Kubernetes resources.</p>
<p>The process typically includes updating configurations and informing other group members about the removal.
Data migration is generally not part of this action and should be handled separately if needed.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_SERVICE_PORT: The port used by the database service.</li>
<li>KB_SERVICE_USER: The username with the necessary permissions to interact with the database service.</li>
<li>KB_SERVICE_PASSWORD: The corresponding password for KB_SERVICE_USER to authenticate with the database service.</li>
<li>KB_PRIMARY_POD_FQDN: The FQDN of the primary Pod within the replication group.</li>
<li>KB_MEMBER_ADDRESSES: A comma-separated list of Pod addresses for all replicas in the group.</li>
<li>KB_LEAVE_MEMBER_POD_NAME: The pod name of the replica being removed from the group.</li>
<li>KB_LEAVE_MEMBER_POD_IP: The IP address of the replica being removed from the group.</li>
</ul>
<p>Expected action output:
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>For example, to remove an OBServer from an OceanBase Cluster in &lsquo;zone1&rsquo;, the following command can be executed:</p>
<pre><code class="language-yaml">command:
- bash
- -c
- |
   ADDRESS=$(KB_MEMBER_ADDRESSES%%,*)
   HOST=$(echo $ADDRESS | cut -d ':' -f 1)
   PORT=$(echo $ADDRESS | cut -d ':' -f 2)
   CLIENT=&quot;mysql -u $KB_SERVICE_USER  -p$KB_SERVICE_PASSWORD -P $PORT -h $HOST -e&quot;
       $CLIENT &quot;ALTER SYSTEM DELETE SERVER '$KB_LEAVE_MEMBER_POD_IP:$KB_SERVICE_PORT' ZONE 'zone1'&quot;
</code></pre>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>readonly</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to switch a replica into the read-only state.</p>
<p>Use Case:
This action is invoked when the database&rsquo;s volume capacity nears its upper limit and space is about to be exhausted.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_POD_FQDN: The FQDN of the replica pod whose role is being checked.</li>
<li>KB_SERVICE_PORT: The port used by the database service.</li>
<li>KB_SERVICE_USER: The username with the necessary permissions to interact with the database service.</li>
<li>KB_SERVICE_PASSWORD: The corresponding password for KB_SERVICE_USER to authenticate with the database service.</li>
</ul>
<p>Expected action output:
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>readwrite</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to transition a replica from the read-only state back to the read-write state.</p>
<p>Use Case:
This action is used to bring back a replica that was previously in a read-only state,
which restricted write operations, to its normal operational state where it can handle
both read and write operations.</p>
<p>The container executing this action has access to following environment variables:</p>
<ul>
<li>KB_POD_FQDN: The FQDN of the replica pod whose role is being checked.</li>
<li>KB_SERVICE_PORT: The port used by the database service.</li>
<li>KB_SERVICE_USER: The username with the necessary permissions to interact with the database service.</li>
<li>KB_SERVICE_PASSWORD: The corresponding password for KB_SERVICE_USER to authenticate with the database service.</li>
</ul>
<p>Expected action output:
- On Failure: An error message, if applicable, indicating why the action failed.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>dataDump</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure for exporting the data from a replica.</p>
<p>Use Case:
This action is intended for initializing a newly created replica with data. It involves exporting data
from an existing replica and importing it into the new, empty replica. This is essential for synchronizing
the state of replicas across the system.</p>
<p>Applicability:
Some database engines or associated sidecar applications (e.g., Patroni) may already provide this functionality.
In such cases, this action may not be required.</p>
<p>The output should be a valid data dump streamed to stdout. It must exclude any irrelevant information to ensure
that only the necessary data is exported for import into the new replica.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>dataLoad</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure for importing data into a replica.</p>
<p>Use Case:
This action is intended for initializing a newly created replica with data. It involves exporting data
from an existing replica and importing it into the new, empty replica. This is essential for synchronizing
the state of replicas across the system.</p>
<p>Some database engines or associated sidecar applications (e.g., Patroni) may already provide this functionality.
In such cases, this action may not be required.</p>
<p>Data should be received through stdin. If any error occurs during the process,
the action must be able to guarantee idempotence to allow for retries from the beginning.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
<tr>
<td>
<code>reconfigure</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure that update a replica with new configuration.</p>
<p>Note: This field is immutable once it has been set.</p>
<p>This Action is reserved for future versions.</p>
</td>
</tr>
<tr>
<td>
<code>accountProvision</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure to generate a new database account.</p>
<p>Use Case:
This action is designed to create system accounts that are utilized for replication, monitoring, backup,
and other administrative tasks.</p>
<p>Note: This field is immutable once it has been set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentMessageMap">ComponentMessageMap
(<code>map[string]string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentStatus">ClusterComponentStatus</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentStatus">ComponentStatus</a>)
</p>
<div>
</div>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentRefEnv">ComponentRefEnv
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefRef">ComponentDefRef</a>)
</p>
<div>
<p>ComponentRefEnv specifies name and value of an env.</p>
<p>Deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the env, it must be a C identifier.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The value of the env.</p>
</td>
</tr>
<tr>
<td>
<code>valueFrom</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentValueFrom">
ComponentValueFrom
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The source from which the value of the env.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentService">ComponentService
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
<p>ComponentService defines a service that would be exposed as an inter-component service within a Cluster.
A Service defined in the ComponentService is expected to be accessed by other Components within the same Cluster.</p>
<p>When a Component needs to use a ComponentService provided by another Component within the same Cluster,
it can declare a variable in the <code>componentDefinition.spec.vars</code> section and bind it to the specific exposed address
of the ComponentService using the <code>serviceVarRef</code> field.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Service</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Service">
Service
</a>
</em>
</td>
<td>
<p>
(Members of <code>Service</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>podService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to create a corresponding Service for each Pod of the selected Component.
When set to true, a set of Services will be automatically generated for each Pod,
and the <code>roleSelector</code> field will be ignored.</p>
<p>The names of the generated Services will follow the same suffix naming pattern: <code>$(serviceName)-$(podOrdinal)</code>.
The total number of generated Services will be equal to the number of replicas specified for the Component.</p>
<p>Example usage:</p>
<pre><code class="language-yaml">name: my-service
serviceName: my-service
podService: true
disableAutoProvision: true
spec:
  type: NodePort
  ports:
  - name: http
    port: 80
    targetPort: 8080
</code></pre>
<p>In this example, if the Component has 3 replicas, three Services will be generated:
- my-service-0: Points to the first Pod (podOrdinal: 0)
- my-service-1: Points to the second Pod (podOrdinal: 1)
- my-service-2: Points to the third Pod (podOrdinal: 2)</p>
<p>Each generated Service will have the specified spec configuration and will target its respective Pod.</p>
<p>This feature is useful when you need to expose each Pod of a Component individually, allowing external access
to specific instances of the Component.</p>
</td>
</tr>
<tr>
<td>
<code>disableAutoProvision</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether the automatic provisioning of the service should be disabled.</p>
<p>If set to true, the service will not be automatically created at the component provisioning.
Instead, you can enable the creation of this service by specifying it explicitly in the cluster API.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Component">Component</a>)
</p>
<div>
<p>ComponentSpec defines the desired state of Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersion specifies the version of the Service expected to be provisioned by this Component.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRef">
[]ServiceRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of ServiceRef for a Component, enabling access to both external services and
Services provided by other Clusters.</p>
<p>Types of services:</p>
<ul>
<li>External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator;
Require a ServiceDescriptor for connection details.</li>
<li>Services provided by a Cluster: Managed by the same KubeBlocks operator;
identified using Cluster, Component and Service names.</li>
</ul>
<p>ServiceRefs with identical <code>serviceRef.name</code> in the same Cluster are considered the same.</p>
<p>Example:</p>
<pre><code class="language-yaml">serviceRefs:
  - name: &quot;redis-sentinel&quot;
    serviceDescriptor:
      name: &quot;external-redis-sentinel&quot;
  - name: &quot;postgres-cluster&quot;
    clusterServiceSelector:
      cluster: &quot;my-postgres-cluster&quot;
      service:
        component: &quot;postgresql&quot;
</code></pre>
<p>The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Labels to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies Annotations to override or add for underlying Pods, PVCs, Account &amp; TLS Secrets, Services Owned by Component.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of environment variables to add.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the resources required by the Component.
It allows defining the CPU, memory requirements and limits for the Component&rsquo;s containers.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentVolumeClaimTemplate">
[]ClusterComponentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for the Component.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of volumes to override.</p>
</td>
</tr>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentService">
[]ComponentService
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides Services defined in referenced ComponentDefinition and exposes endpoints that can be accessed
by clients.</p>
</td>
</tr>
<tr>
<td>
<code>systemAccounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentSystemAccount">
[]ComponentSystemAccount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides system accounts defined in referenced ComponentDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentConfig">
[]ClusterComponentConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration content of a config template.</p>
</td>
</tr>
<tr>
<td>
<code>enabledLogs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies which types of logs should be collected for the Cluster.
The log types are defined in the <code>componentDefinition.spec.logConfigs</code> field with the LogConfig entries.</p>
<p>The elements in the <code>enabledLogs</code> array correspond to the names of the LogConfig entries.
For example, if the <code>componentDefinition.spec.logConfigs</code> defines LogConfig entries with
names &ldquo;slow_query_log&rdquo; and &ldquo;error_log&rdquo;,
you can enable the collection of these logs by including their names in the <code>enabledLogs</code> array:</p>
<pre><code class="language-yaml">enabledLogs:
- slow_query_log
- error_log
</code></pre>
</td>
</tr>
<tr>
<td>
<code>serviceAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceAccount required by the running Component.
This ServiceAccount is used to grant necessary permissions for the Component&rsquo;s Pods to interact
with other Kubernetes resources, such as modifying Pod labels or sending events.</p>
<p>Defaults:
If not specified, KubeBlocks automatically assigns a default ServiceAccount named &ldquo;kb-&#123;cluster.name&#125;&rdquo;,
bound to a default role defined during KubeBlocks installation.</p>
<p>Future Changes:
Future versions might change the default ServiceAccount creation strategy to one per Component,
potentially revising the naming to &ldquo;kb-&#123;cluster.name&#125;-&#123;component.name&#125;&rdquo;.</p>
<p>Users can override the automatic ServiceAccount assignment by explicitly setting the name of
an existed ServiceAccount in this field.</p>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the InstanceUpdateStrategy that will be
employed to update Pods in the InstanceSet when a revision is made to
Template.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Affinity">
Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a group of affinity scheduling rules for the Component.
It allows users to control how the Component&rsquo;s Pods are scheduled onto nodes in the Cluster.</p>
<p>Deprecated since v0.10, replaced by the <code>schedulingPolicy</code> field.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows Pods to be scheduled onto nodes with matching taints.
Each toleration in the array allows the Pod to tolerate node taints based on
specified <code>key</code>, <code>value</code>, <code>effect</code>, and <code>operator</code>.</p>
<ul>
<li>The <code>key</code>, <code>value</code>, and <code>effect</code> identify the taint that the toleration matches.</li>
<li>The <code>operator</code> determines how the toleration matches the taint.</li>
</ul>
<p>Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.</p>
<p>Deprecated since v0.10, replaced by the <code>schedulingPolicy</code> field.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>tlsConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TLSConfig">
TLSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the TLS configuration for the Component, including:</p>
<ul>
<li>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication.</li>
<li>An optional field that specifies the configuration for the TLS certificates issuer when TLS is enabled.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows for the customization of configuration values for each instance within a Component.
An Instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps).
While instances typically share a common configuration as defined in the ClusterComponentSpec,
they can require unique settings in various scenarios:</p>
<p>For example:
- A database Component might require different resource allocations for primary and secondary instances,
  with primaries needing more resources.
- During a rolling upgrade, a Component may first update the image for one or a few instances,
and then update the remaining instances after verifying that the updated instances are functioning correctly.</p>
<p>InstanceTemplate allows for specifying these unique configurations per instance.
Each instance&rsquo;s name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal),
starting with an ordinal of 0.
It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the Component.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the Cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>runtimeClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines runtimeClassName for all Pods managed by this Component.</p>
</td>
</tr>
<tr>
<td>
<code>disableExporter</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines whether metrics exporter information is annotated on the Component&rsquo;s headless Service.</p>
<p>If set to true, the following annotations will not be patched into the Service:</p>
<ul>
<li>&ldquo;monitor.kubeblocks.io/path&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/port&rdquo;</li>
<li>&ldquo;monitor.kubeblocks.io/scheme&rdquo;</li>
</ul>
<p>These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.</p>
</td>
</tr>
<tr>
<td>
<code>stop</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Stop the Component.
If set, all the computing resources will be released.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentStatus">ComponentStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Component">Component</a>)
</p>
<div>
<p>ComponentStatus represents the observed state of a Component within the Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the most recent generation observed for this Component object.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a list of detailed status of the Component object.
Each condition in the list provides real-time information about certain aspect of the Component object.</p>
<p>This field is crucial for administrators and developers to monitor and respond to changes within the Component.
It provides a history of state transitions and a snapshot of the current state that can be used for
automated logic or direct inspection.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentPhase">
ClusterComponentPhase
</a>
</em>
</td>
<td>
<p>Indicates the current phase of the Component, with each phase indicating specific conditions:</p>
<ul>
<li>Creating: The initial phase for new Components, transitioning from &lsquo;empty&rsquo;(&ldquo;&rdquo;).</li>
<li>Running: All Pods in a Running state.</li>
<li>Updating: The Component is currently being updated, with no failed Pods present.</li>
<li>Abnormal: Some Pods have failed, indicating a potentially unstable state.
However, the cluster remains available as long as a quorum of members is functioning.</li>
<li>Failed: A significant number of Pods or critical Pods have failed
The cluster may be non-functional or may offer only limited services (e.g, read-only).</li>
<li>Stopping: All Pods are being terminated, with current replica count at zero.</li>
<li>Stopped: All associated Pods have been successfully deleted.</li>
<li>Deleting: The Component is being deleted.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentMessageMap">
ComponentMessageMap
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A map that stores detailed message about the Component.
Each entry in the map provides insights into specific elements of the Component, such as Pods or workloads.</p>
<p>Keys in this map are formatted as <code>ObjectKind/Name</code>, where <code>ObjectKind</code> could be a type like Pod,
and <code>Name</code> is the specific name of the object.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentSwitchover">ComponentSwitchover
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentLifecycleActions">ComponentLifecycleActions</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>withCandidate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the switchover process for a specified candidate primary or leader instance.
Note that only Action.Exec is currently supported, while Action.HTTP is not.</p>
</td>
</tr>
<tr>
<td>
<code>withoutCandidate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a switchover process that does not involve a specific candidate primary or leader instance.
As with the previous field, only Action.Exec is currently supported, not Action.HTTP.</p>
</td>
</tr>
<tr>
<td>
<code>scriptSpecSelectors</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ScriptSpecSelector">
[]ScriptSpecSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to define the selectors for the scriptSpecs that need to be referenced.
If this field is set, the scripts defined under the &lsquo;scripts&rsquo; field can be invoked or referenced within an Action.</p>
<p>This field is deprecated from v0.9.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentSystemAccount">ComponentSystemAccount
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the system account.</p>
</td>
</tr>
<tr>
<td>
<code>passwordConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PasswordConfig">
PasswordConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the policy for generating the account&rsquo;s password.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>secretRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProvisionSecretRef">
ProvisionSecretRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the secret from which data will be copied to create the new account.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentTemplateSpec">ComponentTemplateSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">ComponentConfigSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the configuration template.</p>
</td>
</tr>
<tr>
<td>
<code>templateRef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the referenced configuration template ConfigMap object.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the namespace of the referenced configuration template ConfigMap object.
An empty namespace is equivalent to the &ldquo;default&rdquo; namespace.</p>
</td>
</tr>
<tr>
<td>
<code>volumeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the volume name of PodTemplate. The configuration file produced through the configuration
template will be mounted to the corresponding volume. Must be a DNS_LABEL name.
The volume name must be defined in podSpec.containers[*].volumeMounts.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMode</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The operator attempts to set default file permissions for scripts (0555) and configurations (0444).
However, certain database engines may require different file permissions.
You can specify the desired file permissions here.</p>
<p>Must be specified as an octal value between 0000 and 0777 (inclusive),
or as a decimal value between 0 and 511 (inclusive).
YAML supports both octal and decimal values for file permissions.</p>
<p>Please note that this setting only affects the permissions of the files themselves.
Directories within the specified path are not impacted by this setting.
It&rsquo;s important to be aware that this setting might conflict with other options
that influence the file mode, such as fsGroup.
In such cases, the resulting file mode may have additional bits set.
Refers to documents of k8s.ConfigMapVolumeSource.defaultMode for more information.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentValueFrom">ComponentValueFrom
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentRefEnv">ComponentRefEnv</a>)
</p>
<div>
<p>ComponentValueFrom is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentValueFromType">
ComponentValueFromType
</a>
</em>
</td>
<td>
<p>Specifies the source to select. It can be one of three types: <code>FieldRef</code>, <code>ServiceRef</code>, <code>HeadlessServiceRef</code>.</p>
</td>
</tr>
<tr>
<td>
<code>fieldPath</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The jsonpath of the source to select when the Type is <code>FieldRef</code>.
Two objects are registered in the jsonpath: <code>componentDef</code> and <code>components</code>:</p>
<ul>
<li><code>componentDef</code> is the component definition object specified in <code>componentRef.componentDefName</code>.</li>
<li><code>components</code> are the component list objects referring to the component definition object.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the format of each headless service address.
Three builtin variables can be used as placeholders: <code>$POD_ORDINAL</code>, <code>$POD_FQDN</code>, <code>$POD_NAME</code></p>
<ul>
<li><code>$POD_ORDINAL</code> represents the ordinal of the pod.</li>
<li><code>$POD_FQDN</code> represents the fully qualified domain name of the pod.</li>
<li><code>$POD_NAME</code> represents the name of the pod.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>joinWith</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The string used to join the values of headless service addresses.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentValueFromType">ComponentValueFromType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentValueFrom">ComponentValueFrom</a>)
</p>
<div>
<p>ComponentValueFromType specifies the type of component value from which the data is derived.</p>
<p>Deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;FieldRef&#34;</p></td>
<td><p>FromFieldRef refers to the value of a specific field in the object.</p>
</td>
</tr><tr><td><p>&#34;HeadlessServiceRef&#34;</p></td>
<td><p>FromHeadlessServiceRef refers to a headless service within the same namespace as the object.</p>
</td>
</tr><tr><td><p>&#34;ServiceRef&#34;</p></td>
<td><p>FromServiceRef refers to a service within the same namespace as the object.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVarSelector">ComponentVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.VarSource">VarSource</a>)
</p>
<div>
<p>ComponentVarSelector selects a var from a Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Component to select from.</p>
</td>
</tr>
<tr>
<td>
<code>ComponentVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVars">
ComponentVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ComponentVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVars">ComponentVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentVarSelector">ComponentVarSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>componentName</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the name of the Component object.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the replicas of the component.</p>
</td>
</tr>
<tr>
<td>
<code>instanceNames</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the pod name list of the component.
and the value will be presented in the following format: name1,name2,&hellip;</p>
</td>
</tr>
<tr>
<td>
<code>podFQDNs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Reference to the pod FQDN list of the component.
The value will be presented in the following format: FQDN1,FQDN2,&hellip;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVersionCompatibilityRule">ComponentVersionCompatibilityRule
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionSpec">ComponentVersionSpec</a>)
</p>
<div>
<p>ComponentVersionCompatibilityRule defines the compatibility between a set of component definitions and a set of releases.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compDefs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CompDefs specifies names for the component definitions associated with this ComponentVersion.
Each name in the list can represent an exact name, a name prefix, or a regular expression pattern.</p>
<p>For example:</p>
<ul>
<li>&ldquo;mysql-8.0.30-v1alpha1&rdquo;: Matches the exact name &ldquo;mysql-8.0.30-v1alpha1&rdquo;</li>
<li>&ldquo;mysql-8.0.30&rdquo;: Matches all names starting with &ldquo;mysql-8.0.30&rdquo;</li>
<li>&rdquo;^mysql-8.0.\d&#123;1,2&#125;$&ldquo;: Matches all names starting with &ldquo;mysql-8.0.&rdquo; followed by one or two digits.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>releases</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Releases is a list of identifiers for the releases.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVersionRelease">ComponentVersionRelease
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionSpec">ComponentVersionSpec</a>)
</p>
<div>
<p>ComponentVersionRelease represents a release of component instances within a ComponentVersion.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is a unique identifier for this release.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>changes</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Changes provides information about the changes made in this release.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceVersion defines the version of the well-known service that the component provides.
The version should follow the syntax and semantics of the &ldquo;Semantic Versioning&rdquo; specification (<a href="http://semver.org/">http://semver.org/</a>).
If the release is used, it will serve as the service version for component instances, overriding the one defined in the component definition.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>images</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Images define the new images for different containers within the release.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVersionSpec">ComponentVersionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentVersion">ComponentVersion</a>)
</p>
<div>
<p>ComponentVersionSpec defines the desired state of ComponentVersion</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityRules</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionCompatibilityRule">
[]ComponentVersionCompatibilityRule
</a>
</em>
</td>
<td>
<p>CompatibilityRules defines compatibility rules between sets of component definitions and releases.</p>
</td>
</tr>
<tr>
<td>
<code>releases</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionRelease">
[]ComponentVersionRelease
</a>
</em>
</td>
<td>
<p>Releases represents different releases of component instances within this ComponentVersion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVersionStatus">ComponentVersionStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentVersion">ComponentVersion</a>)
</p>
<div>
<p>ComponentVersionStatus defines the observed state of ComponentVersion</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>ObservedGeneration is the most recent generation observed for this ComponentVersion.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Phase valid values are `<code>,</code>Available<code>, 'Unavailable</code>.
Available is ComponentVersion become available, and can be used for co-related objects.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extra message for current phase.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersions</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceVersions represent the supported service versions of this ComponentVersion.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ComponentVolume">ComponentVolume
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the volume.
It must be a DNS_LABEL and unique within the pod.
More info can be found at: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names">https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names</a>
Note: This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>needSnapshot</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the creation of a snapshot of this volume is necessary when performing a backup of the Component.</p>
<p>Note: This field cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>highWatermark</code><br/>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>Sets the critical threshold for volume space utilization as a percentage (0-100).</p>
<p>Exceeding this percentage triggers the system to switch the volume to read-only mode as specified in
<code>componentDefinition.spec.lifecycleActions.readOnly</code>.
This precaution helps prevent space depletion while maintaining read-only access.
If the space utilization later falls below this threshold, the system reverts the volume to read-write mode
as defined in <code>componentDefinition.spec.lifecycleActions.readWrite</code>, restoring full functionality.</p>
<p>Note: This field cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">ConfigConstraintSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraint">ConfigConstraint</a>)
</p>
<div>
<p>ConfigConstraintSpec defines the desired state of ConfigConstraint</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reloadOptions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReloadOptions">
ReloadOptions
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the dynamic reload action supported by the engine.
When set, the controller executes the method defined here to execute hot parameter updates.</p>
<p>Dynamic reloading is triggered only if both of the following conditions are met:</p>
<ol>
<li>The modified parameters are listed in the <code>dynamicParameters</code> field.
If <code>reloadStaticParamsBeforeRestart</code> is set to true, modifications to <code>staticParameters</code>
can also trigger a reload.</li>
<li><code>reloadOptions</code> is set.</li>
</ol>
<p>If <code>reloadOptions</code> is not set or the modified parameters are not listed in <code>dynamicParameters</code>,
dynamic reloading will not be triggered.</p>
<p>Example:</p>
<pre><code class="language-yaml">reloadOptions:
 tplScriptTrigger:
   namespace: kb-system
   scriptConfigMapRef: mysql-reload-script
   sync: true
</code></pre>
</td>
</tr>
<tr>
<td>
<code>dynamicActionCanBeMerged</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to consolidate dynamic reload and restart actions into a single restart.</p>
<ul>
<li>If true, updates requiring both actions will result in only a restart, merging the actions.</li>
<li>If false, updates will trigger both actions executed sequentially: first dynamic reload, then restart.</li>
</ul>
<p>This flag allows for more efficient handling of configuration changes by potentially eliminating
an unnecessary reload step.</p>
</td>
</tr>
<tr>
<td>
<code>reloadStaticParamsBeforeRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Configures whether the dynamic reload specified in <code>reloadOptions</code> applies only to dynamic parameters or
to all parameters (including static parameters).</p>
<ul>
<li>false (default): Only modifications to the dynamic parameters listed in <code>dynamicParameters</code>
will trigger a dynamic reload.</li>
<li>true: Modifications to both dynamic parameters listed in <code>dynamicParameters</code> and static parameters
listed in <code>staticParameters</code> will trigger a dynamic reload.
The &ldquo;true&rdquo; option is for certain engines that require static parameters to be set
via SQL statements before they can take effect on restart.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>toolsImageSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ToolsSetup">
ToolsSetup
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the tools container image used by ShellTrigger for dynamic reload.
If the dynamic reload action is triggered by a ShellTrigger, this field is required.
This image must contain all necessary tools for executing the ShellTrigger scripts.</p>
<p>Usually the specified image is referenced by the init container,
which is then responsible for copy the tools from the image to a bin volume.
This ensures that the tools are available to the &lsquo;config-manager&rsquo; sidecar.</p>
</td>
</tr>
<tr>
<td>
<code>downwardAPIOptions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.DownwardAPIChangeTriggeredAction">
[]DownwardAPIChangeTriggeredAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of actions to execute specified commands based on Pod labels.</p>
<p>It utilizes the K8s Downward API to mount label information as a volume into the pod.
The &lsquo;config-manager&rsquo; sidecar container watches for changes in the role label and dynamically invoke
registered commands (usually execute some SQL statements) when a change is detected.</p>
<p>It is designed for scenarios where:</p>
<ul>
<li>Replicas with different roles have different configurations, such as Redis primary &amp; secondary replicas.</li>
<li>After a role switch (e.g., from secondary to primary), some changes in configuration are needed
to reflect the new role.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>scriptConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ScriptConfig">
[]ScriptConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A list of ScriptConfig Object.</p>
<p>Each ScriptConfig object specifies a ConfigMap that contains script files that should be mounted inside the pod.
The scripts are mounted as volumes and can be referenced and executed by the dynamic reload
and DownwardAction to perform specific tasks or configurations.</p>
</td>
</tr>
<tr>
<td>
<code>cfgSchemaTopLevelName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the top-level key in the &lsquo;configurationSchema.cue&rsquo; that organizes the validation rules for parameters.
This key must exist within the CUE script defined in &lsquo;configurationSchema.cue&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>configurationSchema</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CustomParametersValidation">
CustomParametersValidation
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of parameters including their names, default values, descriptions,
types, and constraints (permissible values or the range of valid values).</p>
</td>
</tr>
<tr>
<td>
<code>staticParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List static parameters.
Modifications to any of these parameters require a restart of the process to take effect.</p>
</td>
</tr>
<tr>
<td>
<code>dynamicParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List dynamic parameters.
Modifications to these parameters trigger a configuration reload without requiring a process restart.</p>
</td>
</tr>
<tr>
<td>
<code>immutableParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists the parameters that cannot be modified once set.
Attempting to change any of these parameters will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to match labels on the pod to determine whether a dynamic reload should be performed.</p>
<p>In some scenarios, only specific pods (e.g., primary replicas) need to undergo a dynamic reload.
The <code>selector</code> allows you to specify label selectors to target the desired pods for the reload process.</p>
<p>If the <code>selector</code> is not specified or is nil, all pods managed by the workload will be considered for the dynamic
reload.</p>
</td>
</tr>
<tr>
<td>
<code>formatterConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.FileFormatConfig">
FileFormatConfig
</a>
</em>
</td>
<td>
<p>Specifies the format of the configuration file and any associated parameters that are specific to the chosen format.
Supported formats include <code>ini</code>, <code>xml</code>, <code>yaml</code>, <code>json</code>, <code>hcl</code>, <code>dotenv</code>, <code>properties</code>, and <code>toml</code>.</p>
<p>Each format may have its own set of parameters that can be configured.
For instance, when using the <code>ini</code> format, you can specify the section name.</p>
<p>Example:</p>
<pre><code>formatterConfig:
 format: ini
 iniConfig:
   sectionName: mysqld
</code></pre>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigConstraintStatus">ConfigConstraintStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraint">ConfigConstraint</a>)
</p>
<div>
<p>ConfigConstraintStatus represents the observed state of a ConfigConstraint.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintPhase">
ConfigConstraintPhase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the status of the configuration template.
When set to CCAvailablePhase, the ConfigConstraint can be referenced by ClusterDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides descriptions for abnormal states.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the most recent generation observed for this ConfigConstraint. This value is updated by the API Server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigMapRef">ConfigMapRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.UserResourceRefs">UserResourceRefs</a>)
</p>
<div>
<p>ConfigMapRef defines a reference to a ConfigMap.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ResourceMeta</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ResourceMeta">
ResourceMeta
</a>
</em>
</td>
<td>
<p>
(Members of <code>ResourceMeta</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>configMap</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#configmapvolumesource-v1-core">
Kubernetes core/v1.ConfigMapVolumeSource
</a>
</em>
</td>
<td>
<p>ConfigMap specifies the ConfigMap to be mounted as a volume.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigParams">ConfigParams
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetail">ConfigurationItemDetail</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>content</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Holds the configuration keys and values. This field is a workaround for issues found in kubebuilder and code-generator.
Refer to <a href="https://github.com/kubernetes-sigs/kubebuilder/issues/528">https://github.com/kubernetes-sigs/kubebuilder/issues/528</a> and <a href="https://github.com/kubernetes/code-generator/issues/50">https://github.com/kubernetes/code-generator/issues/50</a> for more details.</p>
<p>Represents the content of the configuration file.</p>
</td>
</tr>
<tr>
<td>
<code>parameters</code><br/>
<em>
map[string]*string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the updated parameters for a single configuration file.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigTemplateExtension">ConfigTemplateExtension
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetail">ConfigurationItemDetail</a>, <a href="#apps.kubeblocks.io/v1alpha1.LegacyRenderedTemplateSpec">LegacyRenderedTemplateSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>templateRef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the referenced configuration template ConfigMap object.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the namespace of the referenced configuration template ConfigMap object.
An empty namespace is equivalent to the &ldquo;default&rdquo; namespace.</p>
</td>
</tr>
<tr>
<td>
<code>policy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MergedPolicy">
MergedPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the strategy for merging externally imported templates into component templates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigurationItemDetail">ConfigurationItemDetail
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigurationSpec">ConfigurationSpec</a>)
</p>
<div>
<p>ConfigurationItemDetail corresponds to settings of a configuration template (a ConfigMap).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the unique identifier of the configuration template.</p>
<p>It must be a string of maximum 63 characters, and can only include lowercase alphanumeric characters,
hyphens, and periods.
The name must start and end with an alphanumeric character.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Deprecated: No longer used. Please use &lsquo;Payload&rsquo; instead. Previously represented the version of the configuration template.</p>
</td>
</tr>
<tr>
<td>
<code>payload</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Payload">
Payload
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>External controllers can trigger a configuration rerender by modifying this field.</p>
<p>Note: Currently, the <code>payload</code> field is opaque and its content is not interpreted by the system.
Modifying this field will cause a rerender, regardless of the specific content of this field.</p>
</td>
</tr>
<tr>
<td>
<code>configSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">
ComponentConfigSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the configuration template (a ConfigMap), ConfigConstraint, and other miscellaneous options.</p>
<p>The configuration template is a ConfigMap that contains multiple configuration files.
Each configuration file is stored as a key-value pair within the ConfigMap.</p>
<p>ConfigConstraint allows defining constraints and validation rules for configuration parameters.
It ensures that the configuration adheres to certain requirements and limitations.</p>
</td>
</tr>
<tr>
<td>
<code>importTemplateRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigTemplateExtension">
ConfigTemplateExtension
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the user-defined configuration template.</p>
<p>When provided, the <code>importTemplateRef</code> overrides the default configuration template
specified in <code>configSpec.templateRef</code>.
This allows users to customize the configuration template according to their specific requirements.</p>
</td>
</tr>
<tr>
<td>
<code>configFileParams</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigParams">
map[string]github.com/apecloud/kubeblocks/apis/apps/v1alpha1.ConfigParams
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the user-defined configuration parameters.</p>
<p>When provided, the parameter values in <code>configFileParams</code> override the default configuration parameters.
This allows users to override the default configuration according to their specific needs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigurationItemDetailStatus">ConfigurationItemDetailStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigurationStatus">ConfigurationStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the configuration template. It is a required field and must be a string of maximum 63 characters.
The name should only contain lowercase alphanumeric characters, hyphens, or periods. It should start and end with an alphanumeric character.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigurationPhase">
ConfigurationPhase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the current status of the configuration item.</p>
<p>Possible values include &ldquo;Creating&rdquo;, &ldquo;Init&rdquo;, &ldquo;Running&rdquo;, &ldquo;Pending&rdquo;, &ldquo;Merged&rdquo;, &ldquo;MergeFailed&rdquo;, &ldquo;FailedAndPause&rdquo;,
&ldquo;Upgrading&rdquo;, &ldquo;Deleting&rdquo;, &ldquo;FailedAndRetry&rdquo;, &ldquo;Finished&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>lastDoneRevision</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the last completed revision of the configuration item. This field is optional.</p>
</td>
</tr>
<tr>
<td>
<code>updateRevision</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the updated revision of the configuration item. This field is optional.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a description of any abnormal status. This field is optional.</p>
</td>
</tr>
<tr>
<td>
<code>reconcileDetail</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ReconcileDetail">
ReconcileDetail
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides detailed information about the execution of the configuration change. This field is optional.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigurationPhase">ConfigurationPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetailStatus">ConfigurationItemDetailStatus</a>)
</p>
<div>
<p>ConfigurationPhase defines the Configuration FSM phase</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FailedAndPause&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FailedAndRetry&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Finished&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Init&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MergeFailed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Merged&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Upgrading&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigurationSpec">ConfigurationSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Configuration">Configuration</a>)
</p>
<div>
<p>ConfigurationSpec defines the desired state of a Configuration resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterRef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the Cluster that this configuration is associated with.</p>
</td>
</tr>
<tr>
<td>
<code>componentName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the name of the Component that this configuration pertains to.</p>
</td>
</tr>
<tr>
<td>
<code>configItemDetails</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetail">
[]ConfigurationItemDetail
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigItemDetails is an array of ConfigurationItemDetail objects.</p>
<p>Each ConfigurationItemDetail corresponds to a configuration template,
which is a ConfigMap that contains multiple configuration files.
Each configuration file is stored as a key-value pair within the ConfigMap.</p>
<p>The ConfigurationItemDetail includes information such as:</p>
<ul>
<li>The configuration template (a ConfigMap)</li>
<li>The corresponding ConfigConstraint (constraints and validation rules for the configuration)</li>
<li>Volume mounts (for mounting the configuration files)</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConfigurationStatus">ConfigurationStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Configuration">Configuration</a>)
</p>
<div>
<p>ConfigurationStatus represents the observed state of a Configuration resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a description of any abnormal status.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the latest generation observed for this
ClusterDefinition. It corresponds to the ConfigConstraint&rsquo;s generation, which is
updated by the API Server.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides detailed status information for opsRequest.</p>
</td>
</tr>
<tr>
<td>
<code>configurationStatus</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetailStatus">
[]ConfigurationItemDetailStatus
</a>
</em>
</td>
<td>
<p>Provides the status of each component undergoing reconfiguration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConnectionCredentialAuth">ConnectionCredentialAuth
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptorSpec">ServiceDescriptorSpec</a>)
</p>
<div>
<p>ConnectionCredentialAuth specifies the authentication credentials required for accessing an external service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>username</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the username for the external service.</p>
</td>
</tr>
<tr>
<td>
<code>password</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the password for the external service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConsensusMember">ConsensusMember
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConsensusSetSpec">ConsensusSetSpec</a>)
</p>
<div>
<p>ConsensusMember is deprecated since v0.7.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the consensus member.</p>
</td>
</tr>
<tr>
<td>
<code>accessMode</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.AccessMode">
AccessMode
</a>
</em>
</td>
<td>
<p>Specifies the services that this member is capable of providing.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the number of Pods that perform this role.
The default is 1 for <code>Leader</code>, 0 for <code>Learner</code>, others for <code>Followers</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ConsensusSetSpec">ConsensusSetSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>ConsensusSetSpec is deprecated since v0.7.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>StatefulSetSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.StatefulSetSpec">
StatefulSetSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>StatefulSetSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>leader</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConsensusMember">
ConsensusMember
</a>
</em>
</td>
<td>
<p>Represents a single leader in the consensus set.</p>
</td>
</tr>
<tr>
<td>
<code>followers</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConsensusMember">
[]ConsensusMember
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Members of the consensus set that have voting rights but are not the leader.</p>
</td>
</tr>
<tr>
<td>
<code>learner</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConsensusMember">
ConsensusMember
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a member of the consensus set that does not have voting rights.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ContainerVars">ContainerVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.HostNetworkVars">HostNetworkVars</a>)
</p>
<div>
<p>ContainerVars defines the vars that can be referenced from a Container.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the container.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.NamedVar">
NamedVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Container port to reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CredentialVar">CredentialVar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConnectionCredentialAuth">ConnectionCredentialAuth</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptorSpec">ServiceDescriptorSpec</a>)
</p>
<div>
<p>CredentialVar represents a variable that retrieves its value either directly from a specified expression
or from a source defined in <code>valueFrom</code>.
Only one of these options may be used at a time.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Holds a direct string or an expression that can be evaluated to a string.</p>
<p>It can include variables denoted by $(VAR_NAME).
These variables are expanded to the value of the environment variables defined in the container.
If a variable cannot be resolved, it remains unchanged in the output.</p>
<p>To escape variable expansion and retain the literal value, use double $ characters.</p>
<p>For example:</p>
<ul>
<li>&rdquo;$(VAR_NAME)&rdquo; will be expanded to the value of the environment variable VAR_NAME.</li>
<li>&rdquo;$$(VAR_NAME)&rdquo; will result in &ldquo;$(VAR_NAME)&rdquo; in the output, without any variable expansion.</li>
</ul>
<p>Default value is an empty string.</p>
</td>
</tr>
<tr>
<td>
<code>valueFrom</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvarsource-v1-core">
Kubernetes core/v1.EnvVarSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the source for the variable&rsquo;s value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CredentialVarSelector">CredentialVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.VarSource">VarSource</a>)
</p>
<div>
<p>CredentialVarSelector selects a var from a Credential (SystemAccount).</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Credential (SystemAccount) to select from.</p>
</td>
</tr>
<tr>
<td>
<code>CredentialVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVars">
CredentialVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>CredentialVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CredentialVars">CredentialVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.CredentialVarSelector">CredentialVarSelector</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceRefVars">ServiceRefVars</a>)
</p>
<div>
<p>CredentialVars defines the vars that can be referenced from a Credential (SystemAccount).
!!!!! CredentialVars will only be used as environment variables for Pods &amp; Actions, and will not be used to render the templates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>username</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>password</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CustomLabelSpec">CustomLabelSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>CustomLabelSpec is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<p>The key of the label.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>The value of the label.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.GVKResource">
[]GVKResource
</a>
</em>
</td>
<td>
<p>The resources that will be patched with the label.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.CustomParametersValidation">CustomParametersValidation
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">ConfigConstraintSpec</a>)
</p>
<div>
<p>CustomParametersValidation Defines a list of configuration items with their names, default values, descriptions,
types, and constraints.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cue</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Hold a string that contains a script written in CUE language that defines a list of configuration items.
Each item is detailed with its name, default value, description, type (e.g. string, integer, float),
and constraints (permissible values or the valid range of values).</p>
<p>CUE (Configure, Unify, Execute) is a declarative language designed for defining and validating
complex data configurations.
It is particularly useful in environments like K8s where complex configurations and validation rules are common.</p>
<p>This script functions as a validator for user-provided configurations, ensuring compliance with
the established specifications and constraints.</p>
</td>
</tr>
<tr>
<td>
<code>schema</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#jsonschemaprops-v1-apiextensions-k8s-io">
Kubernetes api extensions v1.JSONSchemaProps
</a>
</em>
</td>
<td>
<p>Generated from the &lsquo;cue&rsquo; field and transformed into a JSON format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.EnvVar">EnvVar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>EnvVar represents a variable present in the env of Pod/Action or the template of config/script.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the variable. Must be a C_IDENTIFIER.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Variable references <code>$(VAR_NAME)</code> are expanded using the previously defined variables in the current context.</p>
<p>If a variable cannot be resolved, the reference in the input string will be unchanged.
Double <code>$$</code> are reduced to a single <code>$</code>, which allows for escaping the <code>$(VAR_NAME)</code> syntax: i.e.</p>
<ul>
<li><code>$$(VAR_NAME)</code> will produce the string literal <code>$(VAR_NAME)</code>.</li>
</ul>
<p>Escaped references will never be expanded, regardless of whether the variable exists or not.
Defaults to &ldquo;&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>valueFrom</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarSource">
VarSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Source for the variable&rsquo;s value. Cannot be used if value is not empty.</p>
</td>
</tr>
<tr>
<td>
<code>expression</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A Go template expression that will be applied to the resolved value of the var.</p>
<p>The expression will only be evaluated if the var is successfully resolved to a non-credential value.</p>
<p>The resolved value can be accessed by its name within the expression, system vars and other user-defined
non-credential vars can be used within the expression in the same way.
Notice that, when accessing vars by its name, you should replace all the &ldquo;-&rdquo; in the name with &ldquo;_&rdquo;, because of
that &ldquo;-&rdquo; is not a valid identifier in Go.</p>
<p>All expressions are evaluated in the order the vars are defined. If a var depends on any vars that also
have expressions defined, be careful about the evaluation order as it may use intermediate values.</p>
<p>The result of evaluation will be used as the final value of the var. If the expression fails to evaluate,
the resolving of var will also be considered failed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ExecAction">ExecAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Action">Action</a>)
</p>
<div>
<p>ExecAction describes an Action that executes a command inside a container.
Which may run as a K8s job or be executed inside the Lorry sidecar container, depending on the implementation.
Future implementations will standardize execution within Lorry.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>command</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the command to be executed inside the container.
The working directory for this command is the container&rsquo;s root directory(&lsquo;/&rsquo;).
Commands are executed directly without a shell environment, meaning shell-specific syntax (&lsquo;|&rsquo;, etc.) is not supported.
If the shell is required, it must be explicitly invoked in the command.</p>
<p>A successful execution is indicated by an exit status of 0; any non-zero status signifies a failure.</p>
</td>
</tr>
<tr>
<td>
<code>args</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Args represents the arguments that are passed to the <code>command</code> for execution.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Exporter">Exporter
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>containerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the built-in metrics exporter container.</p>
</td>
</tr>
<tr>
<td>
<code>scrapePath</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the http/https url path to scrape for metrics.
If empty, Prometheus uses the default value (e.g. <code>/metrics</code>).</p>
</td>
</tr>
<tr>
<td>
<code>scrapePort</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the port name to scrape for metrics.</p>
</td>
</tr>
<tr>
<td>
<code>scrapeScheme</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PrometheusScheme">
PrometheusScheme
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the schema to use for scraping.
<code>http</code> and <code>https</code> are the expected values unless you rewrite the <code>__scheme__</code> label via relabeling.
If empty, Prometheus uses the default value <code>http</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ExporterConfig">ExporterConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.MonitorConfig">MonitorConfig</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>scrapePort</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<p>scrapePort is exporter port for Time Series Database to scrape metrics.</p>
</td>
</tr>
<tr>
<td>
<code>scrapePath</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>scrapePath is exporter url path for Time Series Database to scrape metrics.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.FailurePolicyType">FailurePolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefRef">ComponentDefRef</a>)
</p>
<div>
<p>FailurePolicyType specifies the type of failure policy.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Fail&#34;</p></td>
<td><p>FailurePolicyFail means that an error will be reported.</p>
</td>
</tr><tr><td><p>&#34;Ignore&#34;</p></td>
<td><p>FailurePolicyIgnore means that an error will be ignored but logged.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.GVKResource">GVKResource
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.CustomLabelSpec">CustomLabelSpec</a>)
</p>
<div>
<p>GVKResource is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>gvk</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the GVK of a resource, such as &ldquo;v1/Pod&rdquo;, &ldquo;apps/v1/StatefulSet&rdquo;, etc.
When a resource matching this is found by the selector, a custom label will be added if it doesn&rsquo;t already exist,
or updated if it does.</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A label query used to filter a set of resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.HScaleDataClonePolicyType">HScaleDataClonePolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.HorizontalScalePolicy">HorizontalScalePolicy</a>)
</p>
<div>
<p>HScaleDataClonePolicyType defines the data clone policy to be used during horizontal scaling.
This policy determines how data is handled when new nodes are added to the cluster.
The policy can be set to <code>None</code>, <code>CloneVolume</code>, or <code>Snapshot</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CloneVolume&#34;</p></td>
<td><p>HScaleDataClonePolicyCloneVolume indicates that data will be cloned from existing volumes during horizontal scaling.</p>
</td>
</tr><tr><td><p>&#34;Snapshot&#34;</p></td>
<td><p>HScaleDataClonePolicyFromSnapshot indicates that data will be cloned from a snapshot during horizontal scaling.</p>
</td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td><p>HScaleDataClonePolicyNone indicates that no data cloning will occur during horizontal scaling.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.HTTPAction">HTTPAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Action">Action</a>)
</p>
<div>
<p>HTTPAction describes an Action that triggers HTTP requests.
HTTPAction is to be implemented in future version.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the endpoint to be requested on the HTTP server.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<p>Specifies the target port for the HTTP request.
It can be specified either as a numeric value in the range of 1 to 65535,
or as a named port that meets the IANA_SVC_NAME specification.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the server&rsquo;s domain name or IP address. Defaults to the Pod&rsquo;s IP.
Prefer setting the &ldquo;Host&rdquo; header in httpHeaders when needed.</p>
</td>
</tr>
<tr>
<td>
<code>scheme</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#urischeme-v1-core">
Kubernetes core/v1.URIScheme
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Designates the protocol used to make the request, such as HTTP or HTTPS.
If not specified, HTTP is used by default.</p>
</td>
</tr>
<tr>
<td>
<code>method</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the type of HTTP request to be made, such as &ldquo;GET,&rdquo; &ldquo;POST,&rdquo; &ldquo;PUT,&rdquo; etc.
If not specified, &ldquo;GET&rdquo; is the default method.</p>
</td>
</tr>
<tr>
<td>
<code>httpHeaders</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#httpheader-v1-core">
[]Kubernetes core/v1.HTTPHeader
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows for the inclusion of custom headers in the request.
HTTP permits the use of repeated headers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.HorizontalScalePolicy">HorizontalScalePolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>HorizontalScalePolicy is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HScaleDataClonePolicyType">
HScaleDataClonePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines the data synchronization method when a component scales out.
The policy can be one of the following: &#123;None, CloneVolume&#125;. The default policy is <code>None</code>.</p>
<ul>
<li><code>None</code>: This is the default policy. It creates an empty volume without data cloning.</li>
<li><code>CloneVolume</code>: This policy clones data to newly scaled pods. It first tries to use a volume snapshot.
If volume snapshot is not enabled, it will attempt to use a backup tool. If neither method works, it will report an error.</li>
<li><code>Snapshot</code>: This policy is deprecated and is an alias for CloneVolume.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>backupPolicyTemplateName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the backup policy template.</p>
</td>
</tr>
<tr>
<td>
<code>volumeMountsName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the volumeMount of the container to backup.
This only works if Type is not None. If not specified, the first volumeMount will be selected.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.HostNetwork">HostNetwork
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>containerPorts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HostNetworkContainerPort">
[]HostNetworkContainerPort
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The list of container ports that are required by the component.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.HostNetworkContainerPort">HostNetworkContainerPort
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.HostNetwork">HostNetwork</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>container</code><br/>
<em>
string
</em>
</td>
<td>
<p>Container specifies the target container within the Pod.</p>
</td>
</tr>
<tr>
<td>
<code>ports</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Ports are named container ports within the specified container.
These container ports must be defined in the container for proper port allocation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.HostNetworkVarSelector">HostNetworkVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.VarSource">VarSource</a>)
</p>
<div>
<p>HostNetworkVarSelector selects a var from host-network resources.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The component to select from.</p>
</td>
</tr>
<tr>
<td>
<code>HostNetworkVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HostNetworkVars">
HostNetworkVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>HostNetworkVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.HostNetworkVars">HostNetworkVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.HostNetworkVarSelector">HostNetworkVarSelector</a>)
</p>
<div>
<p>HostNetworkVars defines the vars that can be referenced from host-network resources.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>container</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ContainerVars">
ContainerVars
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.InstanceMeta">InstanceMeta
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutInstanceMeta">RolloutInstanceMeta</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.InstanceTemplate">InstanceTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
<p>InstanceTemplate allows customization of individual replica configurations in a Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name specifies the unique name of the instance Pod created using this InstanceTemplate.
This name is constructed by concatenating the Component&rsquo;s name, the template&rsquo;s name, and the instance&rsquo;s ordinal
using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0.
The specified name overrides any default naming conventions or patterns.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of instances (Pods) to create from this InstanceTemplate.
This field allows setting how many replicated instances of the Component,
with the specific overrides in the InstanceTemplate, are created.
The default value is 1. A value of 0 disables instance creation.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs to be merged into the Pod&rsquo;s existing annotations.
Existing keys will have their values overwritten, while new keys will be added to the annotations.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs that will be merged into the Pod&rsquo;s existing labels.
Values for existing keys will be overwritten, and new keys will be added.</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the first container&rsquo;s image in the Pod.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the resource requirements of the first container in the Pod.
This field allows for customizing resource allocation (CPU, memory, etc.) for the container.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines Env to override.
Add new or override existing envs.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines Volumes to override.
Add new or override existing volumes.</p>
</td>
</tr>
<tr>
<td>
<code>volumeMounts</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volumemount-v1-core">
[]Kubernetes core/v1.VolumeMount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines VolumeMounts to override.
Add new or override existing volume mounts of the first container in the Pod.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentVolumeClaimTemplate">
[]ClusterComponentVolumeClaimTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines VolumeClaimTemplates to override.
Add new or override existing volume claim templates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">InstanceUpdateStrategy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
<p>InstanceUpdateStrategy indicates the strategy that the InstanceSet
controller will use to perform updates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>partition</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Partition indicates the number of pods that should be updated during a rolling update.
The remaining pods will remain untouched. This is helpful in defining how many pods
should participate in the update process. The update process will follow the order
of pod names in descending lexicographical (dictionary) order. The default value is
ComponentSpec.Replicas (i.e., update all pods).</p>
</td>
</tr>
<tr>
<td>
<code>maxUnavailable</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The maximum number of pods that can be unavailable during the update.
Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
Absolute number is calculated from percentage by rounding up. This can not be 0.
Defaults to 1. The field applies to all pods. That means if there is any unavailable pod,
it will be counted towards MaxUnavailable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Issuer">Issuer
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.TLSConfig">TLSConfig</a>)
</p>
<div>
<p>Issuer defines the TLS certificates issuer for the Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.IssuerName">
IssuerName
</a>
</em>
</td>
<td>
<p>The issuer for TLS certificates.
It only allows two enum values: <code>KubeBlocks</code> and <code>UserProvided</code>.</p>
<ul>
<li><code>KubeBlocks</code> indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used.</li>
<li><code>UserProvided</code> means that the user is responsible for providing their own CA, Cert, and Key.
In this case, the user-provided CA certificate, server certificate, and private key will be used
for TLS communication.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>secretRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.TLSSecretRef">
TLSSecretRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SecretRef is the reference to the secret that contains user-provided certificates.
It is required when the issuer is set to <code>UserProvided</code>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.IssuerName">IssuerName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Issuer">Issuer</a>)
</p>
<div>
<p>IssuerName defines the name of the TLS certificates issuer.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;KubeBlocks&#34;</p></td>
<td><p>IssuerKubeBlocks represents certificates that are signed by the KubeBlocks Operator.</p>
</td>
</tr><tr><td><p>&#34;UserProvided&#34;</p></td>
<td><p>IssuerUserProvided indicates that the user has provided their own CA-signed certificates.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.LegacyRenderedTemplateSpec">LegacyRenderedTemplateSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">ComponentConfigSpec</a>)
</p>
<div>
<p>LegacyRenderedTemplateSpec describes the configuration extension for the lazy rendered template.
Deprecated: LegacyRenderedTemplateSpec has been deprecated since 0.9.0 and will be removed in 0.10.0</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ConfigTemplateExtension</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigTemplateExtension">
ConfigTemplateExtension
</a>
</em>
</td>
<td>
<p>
(Members of <code>ConfigTemplateExtension</code> are embedded into this type.)
</p>
<p>Extends the configuration template.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.LetterCase">LetterCase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.PasswordConfig">PasswordConfig</a>)
</p>
<div>
<p>LetterCase defines the available cases to be used in password generation.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;LowerCases&#34;</p></td>
<td><p>LowerCases represents the use of lower case letters only.</p>
</td>
</tr><tr><td><p>&#34;MixedCases&#34;</p></td>
<td><p>MixedCases represents the use of a mix of both lower and upper case letters.</p>
</td>
</tr><tr><td><p>&#34;UpperCases&#34;</p></td>
<td><p>UpperCases represents the use of upper case letters only.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">LifecycleActionHandler
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentLifecycleActions">ComponentLifecycleActions</a>, <a href="#apps.kubeblocks.io/v1alpha1.RoleProbe">RoleProbe</a>)
</p>
<div>
<p>LifecycleActionHandler describes the implementation of a specific lifecycle action.</p>
<p>Each action is deemed successful if it returns an exit code of 0 for command executions,
or an HTTP 200 status for HTTP(s) actions.
Any other exit code or HTTP status is considered an indication of failure.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>builtinHandler</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.BuiltinActionHandlerType">
BuiltinActionHandlerType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the predefined action handler to be invoked for lifecycle actions.</p>
<p>Lorry, as a sidecar agent co-located with the database container in the same Pod,
includes a suite of built-in action implementations that are tailored to different database engines.
These are known as &ldquo;builtin&rdquo; handlers, includes: <code>mysql</code>, <code>redis</code>, <code>mongodb</code>, <code>etcd</code>,
<code>postgresql</code>, <code>vanilla-postgresql</code>, <code>apecloud-postgresql</code>, <code>wesql</code>, <code>oceanbase</code>, <code>polardbx</code>.</p>
<p>If the <code>builtinHandler</code> field is specified, it instructs Lorry to utilize its internal built-in action handler
to execute the specified lifecycle actions.</p>
<p>The <code>builtinHandler</code> field is of type <code>BuiltinActionHandlerType</code>,
which represents the name of the built-in handler.
The <code>builtinHandler</code> specified within the same <code>ComponentLifecycleActions</code> should be consistent across all
actions.
This means that if you specify a built-in handler for one action, you should use the same handler
for all other actions throughout the entire <code>ComponentLifecycleActions</code> collection.</p>
<p>If you need to define lifecycle actions for database engines not covered by the existing built-in support,
or when the pre-existing built-in handlers do not meet your specific needs,
you can use the <code>customHandler</code> field to define your own action implementation.</p>
<p>Deprecation Notice:</p>
<ul>
<li>In the future, the <code>builtinHandler</code> field will be deprecated in favor of using the <code>customHandler</code> field
for configuring all lifecycle actions.</li>
<li>Instead of using a name to indicate the built-in action implementations in Lorry,
the recommended approach will be to explicitly invoke the desired action implementation through
a gRPC interface exposed by the sidecar agent.</li>
<li>Developers will have the flexibility to either use the built-in action implementations provided by Lorry
or develop their own sidecar agent to implement custom actions and expose them via gRPC interfaces.</li>
<li>This change will allow for greater customization and extensibility of lifecycle actions,
as developers can create their own &ldquo;builtin&rdquo; implementations tailored to their specific requirements.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>customHandler</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a user-defined hook or procedure that is called to perform the specific lifecycle action.
It offers a flexible and expandable approach for customizing the behavior of a Component by leveraging
tailored actions.</p>
<p>An Action can be implemented as either an ExecAction or an HTTPAction, with future versions planning
to support GRPCAction,
thereby accommodating unique logic for different database systems within the Action&rsquo;s framework.</p>
<p>In future iterations, all built-in handlers are expected to transition to GRPCAction.
This change means that Lorry or other sidecar agents will expose the implementation of actions
through a GRPC interface for external invocation.
Then the controller will interact with these actions via GRPCAction calls.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.LogConfig">LogConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies a descriptive label for the log type, such as &lsquo;slow&rsquo; for a MySQL slow log file.
It provides a clear identification of the log&rsquo;s purpose and content.</p>
</td>
</tr>
<tr>
<td>
<code>filePathPattern</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the paths or patterns identifying where the log files are stored.
This field allows the system to locate and manage log files effectively.</p>
<p>Examples:</p>
<ul>
<li>/home/postgres/pgdata/pgroot/data/log/postgresql-*</li>
<li>/data/mysql/log/mysqld-error.log</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.MergedPolicy">MergedPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigTemplateExtension">ConfigTemplateExtension</a>)
</p>
<div>
<p>MergedPolicy defines how to merge external imported templates into component templates.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;add&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;patch&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;replace&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.MonitorConfig">MonitorConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>builtIn</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>builtIn is a switch to enable KubeBlocks builtIn monitoring.
If BuiltIn is set to true, monitor metrics will be scraped automatically.
If BuiltIn is set to false, the provider should set ExporterConfig and Sidecar container own.</p>
</td>
</tr>
<tr>
<td>
<code>exporterConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ExporterConfig">
ExporterConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>exporterConfig provided by provider, which specify necessary information to Time Series Database.
exporterConfig is valid when builtIn is false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.MultipleClusterObjectCombinedOption">MultipleClusterObjectCombinedOption
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectOption">MultipleClusterObjectOption</a>)
</p>
<div>
<p>MultipleClusterObjectCombinedOption defines options for handling combined variables.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>newVarSuffix</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If set, the existing variable will be kept, and a new variable will be defined with the specified suffix
in pattern: $(var.name)_$(suffix).
The new variable will be auto-created and placed behind the existing one.
If not set, the existing variable will be reused with the value format defined below.</p>
</td>
</tr>
<tr>
<td>
<code>valueFormat</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectValueFormat">
MultipleClusterObjectValueFormat
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The format of the value that the operator will use to compose values from multiple components.</p>
</td>
</tr>
<tr>
<td>
<code>flattenFormat</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectValueFormatFlatten">
MultipleClusterObjectValueFormatFlatten
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The flatten format, default is: $(comp-name-1):value,$(comp-name-2):value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.MultipleClusterObjectOption">MultipleClusterObjectOption
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterObjectReference">ClusterObjectReference</a>)
</p>
<div>
<p>MultipleClusterObjectOption defines the options for handling multiple cluster objects matched.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>requireAllComponentObjects</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>RequireAllComponentObjects controls whether all component objects must exist before resolving.
If set to true, resolving will only proceed if all component objects are present.</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectStrategy">
MultipleClusterObjectStrategy
</a>
</em>
</td>
<td>
<p>Define the strategy for handling multiple cluster objects.</p>
</td>
</tr>
<tr>
<td>
<code>combinedOption</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectCombinedOption">
MultipleClusterObjectCombinedOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Define the options for handling combined variables.
Valid only when the strategy is set to &ldquo;combined&rdquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.MultipleClusterObjectStrategy">MultipleClusterObjectStrategy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectOption">MultipleClusterObjectOption</a>)
</p>
<div>
<p>MultipleClusterObjectStrategy defines the strategy for handling multiple cluster objects.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;combined&#34;</p></td>
<td><p>MultipleClusterObjectStrategyCombined - the values from all matched components will be combined into a single
variable using the specified option.</p>
</td>
</tr><tr><td><p>&#34;individual&#34;</p></td>
<td><p>MultipleClusterObjectStrategyIndividual - each matched component will have its individual variable with its name
as the suffix.
This is required when referencing credential variables that cannot be passed by values.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.MultipleClusterObjectValueFormat">MultipleClusterObjectValueFormat
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectCombinedOption">MultipleClusterObjectCombinedOption</a>)
</p>
<div>
<p>MultipleClusterObjectValueFormat defines the format details for the value.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Flatten&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.MultipleClusterObjectValueFormatFlatten">MultipleClusterObjectValueFormatFlatten
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.MultipleClusterObjectCombinedOption">MultipleClusterObjectCombinedOption</a>)
</p>
<div>
<p>MultipleClusterObjectValueFormatFlatten defines the flatten format for the value.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>delimiter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Pair delimiter.</p>
</td>
</tr>
<tr>
<td>
<code>keyValueDelimiter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key-value delimiter.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.NamedVar">NamedVar
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ContainerVars">ContainerVars</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceVars">ServiceVars</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>option</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.PasswordConfig">PasswordConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentSystemAccount">ComponentSystemAccount</a>, <a href="#apps.kubeblocks.io/v1alpha1.SystemAccount">SystemAccount</a>, <a href="#apps.kubeblocks.io/v1alpha1.SystemAccountSpec">SystemAccountSpec</a>)
</p>
<div>
<p>PasswordConfig helps provide to customize complexity of password generation pattern.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>length</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The length of the password.</p>
</td>
</tr>
<tr>
<td>
<code>numDigits</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of digits in the password.</p>
</td>
</tr>
<tr>
<td>
<code>numSymbols</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of symbols in the password.</p>
</td>
</tr>
<tr>
<td>
<code>letterCase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LetterCase">
LetterCase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The case of the letters in the password.</p>
</td>
</tr>
<tr>
<td>
<code>seed</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Seed to generate the account&rsquo;s password.
Cannot be updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Payload">Payload
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetail">ConfigurationItemDetail</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>-</code><br/>
<em>
map[string]any
</em>
</td>
<td>
<em>(Optional)</em>
<p>Holds the payload data. This field is optional and can contain any type of data.
Not included in the JSON representation of the object.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.PersistentVolumeClaimSpec">PersistentVolumeClaimSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentVolumeClaimTemplate">ClusterComponentVolumeClaimTemplate</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessModes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeaccessmode-v1-core">
[]Kubernetes core/v1.PersistentVolumeAccessMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Contains the desired access modes the volume should have.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1</a>.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volumeresourcerequirements-v1-core">
Kubernetes core/v1.VolumeResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the minimum resources the volume should have.
If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that
are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources">https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources</a>.</p>
</td>
</tr>
<tr>
<td>
<code>storageClassName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the StorageClass required by the claim.
More info: <a href="https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1">https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1</a>.</p>
</td>
</tr>
<tr>
<td>
<code>volumeMode</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumemode-v1-core">
Kubernetes core/v1.PersistentVolumeMode
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines what type of volume is required by the claim, either Block or Filesystem.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Phase">Phase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterDefinitionStatus">ClusterDefinitionStatus</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionStatus">ComponentDefinitionStatus</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentVersionStatus">ComponentVersionStatus</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptorStatus">ServiceDescriptorStatus</a>)
</p>
<div>
<p>Phase represents the current status of the ClusterDefinition CR.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Available&#34;</p></td>
<td><p>AvailablePhase indicates that the object is in an available state.</p>
</td>
</tr><tr><td><p>&#34;Unavailable&#34;</p></td>
<td><p>UnavailablePhase indicates that the object is in an unavailable state.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.PodAntiAffinity">PodAntiAffinity
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Affinity">Affinity</a>)
</p>
<div>
<p>PodAntiAffinity defines the pod anti-affinity strategy.</p>
<p>This strategy determines how pods are scheduled in relation to other pods, with the aim of either spreading pods
across nodes (Preferred) or ensuring that certain pods do not share a node (Required).</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Preferred&#34;</p></td>
<td><p>Preferred indicates that the scheduler will try to enforce the anti-affinity rules, but it will not guarantee it.</p>
</td>
</tr><tr><td><p>&#34;Required&#34;</p></td>
<td><p>Required indicates that the scheduler must enforce the anti-affinity rules and will not schedule the pods unless
the rules are met.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.PodAvailabilityPolicy">PodAvailabilityPolicy
(<code>string</code> alias)</h3>
<div>
<p>PodAvailabilityPolicy pod availability strategy.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Available&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UnAvailable&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.PostStartAction">PostStartAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>PostStartAction is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cmdExecutorConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CmdExecutorConfig">
CmdExecutorConfig
</a>
</em>
</td>
<td>
<p>Specifies the  post-start command to be executed.</p>
</td>
</tr>
<tr>
<td>
<code>scriptSpecSelectors</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ScriptSpecSelector">
[]ScriptSpecSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to select the script that need to be referenced.
When defined, the scripts defined in scriptSpecs can be referenced within the CmdExecutorConfig.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.PreConditionType">PreConditionType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Action">Action</a>)
</p>
<div>
<p>PreConditionType defines the preCondition type of the action execution.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ClusterReady&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ComponentReady&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Immediately&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RuntimeReady&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Probe">Probe
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Action</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<p>
(Members of <code>Action</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>initialDelaySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of seconds to wait after the container has started before the RoleProbe
begins to detect the container&rsquo;s role.</p>
</td>
</tr>
<tr>
<td>
<code>periodSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the frequency at which the probe is conducted. This value is expressed in seconds.
Default to 10 seconds. Minimum value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>successThreshold</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Minimum consecutive successes for the probe to be considered successful after having failed.
Defaults to 1. Minimum value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>failureThreshold</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Minimum consecutive failures for the probe to be considered failed after having succeeded.
Defaults to 3. Minimum value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.PrometheusScheme">PrometheusScheme
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Exporter">Exporter</a>)
</p>
<div>
<p>PrometheusScheme defines the protocol of prometheus scrape metrics.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;https&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ProtectedVolume">ProtectedVolume
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.VolumeProtectionSpec">VolumeProtectionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The Name of the volume to protect.</p>
</td>
</tr>
<tr>
<td>
<code>highWatermark</code><br/>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the high watermark threshold for the volume, it will override the component level threshold.
If the value is invalid, it will be ignored and the component level threshold will be used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ProvisionPolicy">ProvisionPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.SystemAccountConfig">SystemAccountConfig</a>)
</p>
<div>
<p>ProvisionPolicy defines the policy details for creating accounts.</p>
<p>Deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProvisionPolicyType">
ProvisionPolicyType
</a>
</em>
</td>
<td>
<p>Specifies the method to provision an account.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProvisionScope">
ProvisionScope
</a>
</em>
</td>
<td>
<p>Defines the scope within which the account is provisioned.</p>
</td>
</tr>
<tr>
<td>
<code>statements</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProvisionStatements">
ProvisionStatements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The statement to provision an account.</p>
</td>
</tr>
<tr>
<td>
<code>secretRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProvisionSecretRef">
ProvisionSecretRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The external secret to refer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ProvisionPolicyType">ProvisionPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ProvisionPolicy">ProvisionPolicy</a>)
</p>
<div>
<p>ProvisionPolicyType defines the policy for creating accounts.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CreateByStmt&#34;</p></td>
<td><p>CreateByStmt will create account w.r.t. deletion and creation statement given by provider.</p>
</td>
</tr><tr><td><p>&#34;ReferToExisting&#34;</p></td>
<td><p>ReferToExisting will not create account, but create a secret by copying data from referred secret file.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ProvisionScope">ProvisionScope
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ProvisionPolicy">ProvisionPolicy</a>)
</p>
<div>
<p>ProvisionScope defines the scope of provision within a component.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AllPods&#34;</p></td>
<td><p>AllPods indicates that accounts will be created for all pods within the component.</p>
</td>
</tr><tr><td><p>&#34;AnyPods&#34;</p></td>
<td><p>AnyPods indicates that accounts will be created only on a single pod within the component.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ProvisionSecretRef">ProvisionSecretRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentSystemAccount">ComponentSystemAccount</a>, <a href="#apps.kubeblocks.io/v1alpha1.ProvisionPolicy">ProvisionPolicy</a>, <a href="#apps.kubeblocks.io/v1alpha1.SystemAccount">SystemAccount</a>)
</p>
<div>
<p>ProvisionSecretRef represents the reference to a secret.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The unique identifier of the secret.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>The namespace where the secret is located.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ProvisionStatements">ProvisionStatements
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ProvisionPolicy">ProvisionPolicy</a>)
</p>
<div>
<p>ProvisionStatements defines the statements used to create accounts.</p>
<p>Deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>creation</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the statement required to create a new account with the necessary privileges.</p>
</td>
</tr>
<tr>
<td>
<code>update</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the statement required to update the password of an existing account.</p>
</td>
</tr>
<tr>
<td>
<code>deletion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the statement required to delete an existing account.
Typically used in conjunction with the creation statement to delete an account before recreating it.
For example, one might use a <code>drop user if exists</code> statement followed by a <code>create user</code> statement to ensure a fresh account.</p>
<p>Deprecated: This field is deprecated and the update statement should be used instead.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RSMSpec">RSMSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>RSMSpec is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of roles defined within the system.</p>
</td>
</tr>
<tr>
<td>
<code>roleProbe</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.RoleProbe">
RoleProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the method used to probe a role.</p>
</td>
</tr>
<tr>
<td>
<code>membershipReconfiguration</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MembershipReconfiguration">
MembershipReconfiguration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the actions required for dynamic membership reconfiguration.</p>
</td>
</tr>
<tr>
<td>
<code>memberUpdateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MemberUpdateStrategy">
MemberUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Describes the strategy for updating Members (Pods).</p>
<ul>
<li><code>Serial</code>: Updates Members sequentially to ensure minimum component downtime.</li>
<li><code>BestEffortParallel</code>: Updates Members in parallel to ensure minimum component write downtime.</li>
<li><code>Parallel</code>: Forces parallel updates.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ReconcileDetail">ReconcileDetail
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigurationItemDetailStatus">ConfigurationItemDetailStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>policy</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the policy applied during the most recent execution.</p>
</td>
</tr>
<tr>
<td>
<code>execResult</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the outcome of the most recent execution.</p>
</td>
</tr>
<tr>
<td>
<code>currentRevision</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the current revision of the configuration item.</p>
</td>
</tr>
<tr>
<td>
<code>succeedCount</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the number of pods where configuration changes were successfully applied.</p>
</td>
</tr>
<tr>
<td>
<code>expectedCount</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the total number of pods that require execution of configuration changes.</p>
</td>
</tr>
<tr>
<td>
<code>errMessage</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the error message generated when the execution of configuration changes fails.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ReloadOptions">ReloadOptions
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">ConfigConstraintSpec</a>)
</p>
<div>
<p>ReloadOptions defines the mechanisms available for dynamically reloading a process within K8s without requiring a restart.</p>
<p>Only one of the mechanisms can be specified at a time.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>unixSignalTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.UnixSignalTrigger">
UnixSignalTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to trigger a reload by sending a specific Unix signal to the process.</p>
</td>
</tr>
<tr>
<td>
<code>shellTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ShellTrigger">
ShellTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows to execute a custom shell script to reload the process.</p>
</td>
</tr>
<tr>
<td>
<code>tplScriptTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.TPLScriptTrigger">
TPLScriptTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enables reloading process using a Go template script.</p>
</td>
</tr>
<tr>
<td>
<code>autoTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.AutoTrigger">
AutoTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Automatically perform the reload when specified conditions are met.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ReplicaRole">ReplicaRole
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ReplicaRole represents a role that can be assumed by a component instance.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the role&rsquo;s identifier. It is used to set the &ldquo;apps.kubeblocks.io/role&rdquo; label value
on the corresponding object.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>serviceable</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether a replica assigned this role is capable of providing services.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>writable</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines if a replica in this role has the authority to perform write operations.
A writable replica can modify data, handle update operations.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>votable</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether a replica with this role has voting rights.
In distributed systems, this typically means the replica can participate in consensus decisions,
configuration changes, or other processes that require a quorum.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ReplicasLimit">ReplicasLimit
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ReplicasLimit defines the valid range of number of replicas supported.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The minimum limit of replicas.</p>
</td>
</tr>
<tr>
<td>
<code>maxReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The maximum limit of replicas.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ReplicationSetSpec">ReplicationSetSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>ReplicationSetSpec is deprecated since v0.7.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>StatefulSetSpec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.StatefulSetSpec">
StatefulSetSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>StatefulSetSpec</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RerenderResourceType">RerenderResourceType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentConfigSpec">ComponentConfigSpec</a>)
</p>
<div>
<p>RerenderResourceType defines the resource requirements for a component.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;hscale&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;vscale&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;shardingHScale&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ResourceMeta">ResourceMeta
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigMapRef">ConfigMapRef</a>, <a href="#apps.kubeblocks.io/v1alpha1.SecretRef">SecretRef</a>)
</p>
<div>
<p>ResourceMeta encapsulates metadata and configuration for referencing ConfigMaps and Secrets as volumes.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.</p>
</td>
</tr>
<tr>
<td>
<code>mountPoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>MountPoint is the filesystem path where the volume will be mounted.</p>
</td>
</tr>
<tr>
<td>
<code>subPath</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SubPath specifies a path within the volume from which to mount.</p>
</td>
</tr>
<tr>
<td>
<code>asVolumeFrom</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>AsVolumeFrom lists the names of containers in which the volume should be mounted.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RetryPolicy">RetryPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Action">Action</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxRetries</code><br/>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the maximum number of retry attempts that should be made for a given Action.
This value is set to 0 by default, indicating that no retries will be made.</p>
</td>
</tr>
<tr>
<td>
<code>retryInterval</code><br/>
<em>
time.Duration
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the duration of time to wait between each retry attempt.
This value is set to 0 by default, indicating that there will be no delay between retry attempts.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RoleArbitrator">RoleArbitrator
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>RoleArbitrator defines how to arbitrate the role of replicas.</p>
<p>Deprecated since v0.9</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;External&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Lorry&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RoleProbe">RoleProbe
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentLifecycleActions">ComponentLifecycleActions</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>LifecycleActionHandler</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.LifecycleActionHandler">
LifecycleActionHandler
</a>
</em>
</td>
<td>
<p>
(Members of <code>LifecycleActionHandler</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>initialDelaySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of seconds to wait after the container has started before the RoleProbe
begins to detect the container&rsquo;s role.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of seconds after which the probe times out.
Defaults to 1 second. Minimum value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>periodSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the frequency at which the probe is conducted. This value is expressed in seconds.
Default to 10 seconds. Minimum value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutComponent">RolloutComponent
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutSpec">RolloutSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the component.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target ServiceVersion of the component.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target ComponentDefinition of the component.</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategy">
RolloutStrategy
</a>
</em>
</td>
<td>
<p>Specifies the rollout strategy for the component.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of instances to be rolled out.</p>
</td>
</tr>
<tr>
<td>
<code>instanceMeta</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutInstanceMeta">
RolloutInstanceMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional meta for the instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutComponentStatus">RolloutComponentStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutStatus">RolloutStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the component.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>The ServiceVersion of the component before the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>The ComponentDefinition of the component before the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The replicas the component has before the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>newReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The new replicas the component has been created successfully.</p>
</td>
</tr>
<tr>
<td>
<code>rolledOutReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The replicas the component has been rolled out successfully.</p>
</td>
</tr>
<tr>
<td>
<code>canaryReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of canary replicas the component has.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The instances that are scaled down.</p>
</td>
</tr>
<tr>
<td>
<code>lastScaleUpTimestamp</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The last time a component replica was scaled up successfully.</p>
</td>
</tr>
<tr>
<td>
<code>lastScaleDownTimestamp</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The last time a component replica was scaled down successfully.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutInstanceMeta">RolloutInstanceMeta
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutComponent">RolloutComponent</a>, <a href="#apps.kubeblocks.io/v1alpha1.RolloutSharding">RolloutSharding</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>canary</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.InstanceMeta">
InstanceMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Meta added to the new instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutPromoteCondition">RolloutPromoteCondition
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutPromotion">RolloutPromotion</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>prev</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The condition before promoting the new instances.</p>
<p>If specified, the new instances will be promoted only when the condition is met.</p>
</td>
</tr>
<tr>
<td>
<code>post</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The condition after promoting the new instances successfully.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutPromotion">RolloutPromotion
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategyCreate">RolloutStrategyCreate</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>auto</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to automatically promote the new instances.</p>
</td>
</tr>
<tr>
<td>
<code>delaySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The delay seconds before promoting the new instances.</p>
</td>
</tr>
<tr>
<td>
<code>condition</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutPromoteCondition">
RolloutPromoteCondition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The condition for promoting the new instances.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownDelaySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The delay seconds before scaling down the old instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutSharding">RolloutSharding
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutSpec">RolloutSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the sharding.</p>
</td>
</tr>
<tr>
<td>
<code>shardingDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target ShardingDefinition of the sharding.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target ServiceVersion of the sharding.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target ComponentDefinition of the sharding.</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategy">
RolloutStrategy
</a>
</em>
</td>
<td>
<p>Specifies the rollout strategy for the sharding.</p>
</td>
</tr>
<tr>
<td>
<code>instanceMeta</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutInstanceMeta">
RolloutInstanceMeta
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional meta for the instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutShardingStatus">RolloutShardingStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutStatus">RolloutStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the sharding.</p>
</td>
</tr>
<tr>
<td>
<code>shardingDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>The ShardingDefinition of the sharding before the rollout.</p>
<p>optional</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>The ServiceVersion of the sharding before the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>compDef</code><br/>
<em>
string
</em>
</td>
<td>
<p>The ComponentDefinition of the sharding before the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The replicas the sharding has before the rollout.</p>
</td>
</tr>
<tr>
<td>
<code>newReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The new replicas the sharding has been created successfully.</p>
</td>
</tr>
<tr>
<td>
<code>rolledOutReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The replicas the sharding has been rolled out successfully.</p>
</td>
</tr>
<tr>
<td>
<code>canaryReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of canary replicas the sharding has.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The instances that are scaled down.</p>
</td>
</tr>
<tr>
<td>
<code>lastScaleUpTimestamp</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The last time a sharding replica was scaled up successfully.</p>
</td>
</tr>
<tr>
<td>
<code>lastScaleDownTimestamp</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The last time a sharding replica was scaled down successfully.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutSpec">RolloutSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Rollout">Rollout</a>)
</p>
<div>
<p>RolloutSpec defines the desired state of Rollout</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clusterName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the target cluster of the Rollout.</p>
</td>
</tr>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutComponent">
[]RolloutComponent
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target components to be rolled out.</p>
</td>
</tr>
<tr>
<td>
<code>shardings</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutSharding">
[]RolloutSharding
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the target shardings to be rolled out.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutState">RolloutState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutStatus">RolloutStatus</a>)
</p>
<div>
<p>RolloutState defines the state of the Rollout within the .status.state field.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Error&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Rolling&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutStatus">RolloutStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Rollout">Rollout</a>)
</p>
<div>
<p>RolloutStatus defines the observed state of Rollout</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>The most recent generation number of the Rollout object that has been observed by the controller.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutState">
RolloutState
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The current state of the Rollout.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides additional information about the phase.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents a list of detailed status of the Rollout object.</p>
</td>
</tr>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutComponentStatus">
[]RolloutComponentStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Records the status information of all components within the Rollout.</p>
</td>
</tr>
<tr>
<td>
<code>shardings</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutShardingStatus">
[]RolloutShardingStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Records the status information of all shardings within the Rollout.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutStrategy">RolloutStrategy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutComponent">RolloutComponent</a>, <a href="#apps.kubeblocks.io/v1alpha1.RolloutSharding">RolloutSharding</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>inplace</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategyInplace">
RolloutStrategyInplace
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>In-place rollout strategy.</p>
<p>If specified, the rollout will be performed in-place (delete and then create).</p>
</td>
</tr>
<tr>
<td>
<code>replace</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategyReplace">
RolloutStrategyReplace
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Replace rollout strategy.</p>
<p>If specified, the rollout will be performed by replacing the old instances with new instances one by one (create and then delete).</p>
</td>
</tr>
<tr>
<td>
<code>create</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategyCreate">
RolloutStrategyCreate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Create rollout strategy.</p>
<p>If specified, the rollout will be performed by creating new instances.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutStrategyCreate">RolloutStrategyCreate
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategy">RolloutStrategy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>canary</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Whether to decorate the new instances as canary instances.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the new instance.</p>
</td>
</tr>
<tr>
<td>
<code>promotion</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.RolloutPromotion">
RolloutPromotion
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the promotion strategy for the component.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutStrategyInplace">RolloutStrategyInplace
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategy">RolloutStrategy</a>)
</p>
<div>
</div>
<h3 id="apps.kubeblocks.io/v1alpha1.RolloutStrategyReplace">RolloutStrategyReplace
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategy">RolloutStrategy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the new instance.</p>
</td>
</tr>
<tr>
<td>
<code>perInstanceIntervalSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of seconds to wait between rolling out two instances.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownDelaySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The number of seconds to wait before scaling down an old instance, after the new instance becomes ready.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SchedulingPolicy">SchedulingPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.InstanceTemplate">InstanceTemplate</a>, <a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategyCreate">RolloutStrategyCreate</a>, <a href="#apps.kubeblocks.io/v1alpha1.RolloutStrategyReplace">RolloutStrategyReplace</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>schedulerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the Pod will be dispatched by specified scheduler.
If not specified, the Pod will be dispatched by default scheduler.</p>
</td>
</tr>
<tr>
<td>
<code>nodeSelector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeSelector is a selector which must be true for the Pod to fit on a node.
Selector which must match a node&rsquo;s labels for the Pod to be scheduled on that node.
More info: <a href="https://kubernetes.io/docs/concepts/configuration/assign-pod-node/">https://kubernetes.io/docs/concepts/configuration/assign-pod-node/</a></p>
</td>
</tr>
<tr>
<td>
<code>nodeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,
the scheduler simply schedules this Pod onto that node, assuming that it fits resource
requirements.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#affinity-v1-core">
Kubernetes core/v1.Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows Pods to be scheduled onto nodes with matching taints.
Each toleration in the array allows the Pod to tolerate node taints based on
specified <code>key</code>, <code>value</code>, <code>effect</code>, and <code>operator</code>.</p>
<ul>
<li>The <code>key</code>, <code>value</code>, and <code>effect</code> identify the taint that the toleration matches.</li>
<li>The <code>operator</code> determines how the toleration matches the taint.</li>
</ul>
<p>Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.</p>
</td>
</tr>
<tr>
<td>
<code>topologySpreadConstraints</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#topologyspreadconstraint-v1-core">
[]Kubernetes core/v1.TopologySpreadConstraint
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TopologySpreadConstraints describes how a group of Pods ought to spread across topology
domains. Scheduler will schedule Pods in a way which abides by the constraints.
All topologySpreadConstraints are ANDed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ScriptSpecSelector">ScriptSpecSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentSwitchover">ComponentSwitchover</a>, <a href="#apps.kubeblocks.io/v1alpha1.PostStartAction">PostStartAction</a>, <a href="#apps.kubeblocks.io/v1alpha1.SwitchoverAction">SwitchoverAction</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the name of the ScriptSpec referent.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SecretRef">SecretRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.UserResourceRefs">UserResourceRefs</a>)
</p>
<div>
<p>SecretRef defines a reference to a Secret.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ResourceMeta</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ResourceMeta">
ResourceMeta
</a>
</em>
</td>
<td>
<p>
(Members of <code>ResourceMeta</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#secretvolumesource-v1-core">
Kubernetes core/v1.SecretVolumeSource
</a>
</em>
</td>
<td>
<p>Secret specifies the Secret to be mounted as a volume.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.Service">Service
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterService">ClusterService</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentService">ComponentService</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name defines the name of the service.
otherwise, it indicates the name of the service.
Others can refer to this service by its name. (e.g., connection credential)
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>serviceName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ServiceName defines the name of the underlying service object.
If not specified, the default service name with different patterns will be used:</p>
<ul>
<li>CLUSTER_NAME: for cluster-level services</li>
<li>CLUSTER_NAME-COMPONENT_NAME: for component-level services</li>
</ul>
<p>Only one default service name is allowed.
Cannot be updated.</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If ServiceType is LoadBalancer, cloud provider related parameters can be put here
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer">https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer</a>.</p>
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#servicespec-v1-core">
Kubernetes core/v1.ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Spec defines the behavior of a service.
<a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</a></p>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>ports</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceport-v1-core">
[]Kubernetes core/v1.ServicePort
</a>
</em>
</td>
<td>
<p>The list of ports that are exposed by this service.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Route service traffic to pods with label keys and values matching this
selector. If empty or not present, the service is assumed to have an
external process managing its endpoints, which Kubernetes will not
modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.
Ignored if type is ExternalName.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/">https://kubernetes.io/docs/concepts/services-networking/service/</a></p>
</td>
</tr>
<tr>
<td>
<code>clusterIP</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>clusterIP is the IP address of the service and is usually assigned
randomly. If an address is specified manually, is in-range (as per
system configuration), and is not in use, it will be allocated to the
service; otherwise creation of the service will fail. This field may not
be changed through updates unless the type field is also being changed
to ExternalName (which requires this field to be blank) or the type
field is being changed from ExternalName (in which case this field may
optionally be specified, as describe above).  Valid values are &ldquo;None&rdquo;,
empty string (&ldquo;&rdquo;), or a valid IP address. Setting this to &ldquo;None&rdquo; makes a
&ldquo;headless service&rdquo; (no virtual IP), which is useful when direct endpoint
connections are preferred and proxying is not required.  Only applies to
types ClusterIP, NodePort, and LoadBalancer. If this field is specified
when creating a Service of type ExternalName, creation will fail. This
field will be wiped when updating a Service to type ExternalName.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>clusterIPs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ClusterIPs is a list of IP addresses assigned to this service, and are
usually assigned randomly.  If an address is specified manually, is
in-range (as per system configuration), and is not in use, it will be
allocated to the service; otherwise creation of the service will fail.
This field may not be changed through updates unless the type field is
also being changed to ExternalName (which requires this field to be
empty) or the type field is being changed from ExternalName (in which
case this field may optionally be specified, as describe above).  Valid
values are &ldquo;None&rdquo;, empty string (&ldquo;&rdquo;), or a valid IP address.  Setting
this to &ldquo;None&rdquo; makes a &ldquo;headless service&rdquo; (no virtual IP), which is
useful when direct endpoint connections are preferred and proxying is
not required.  Only applies to types ClusterIP, NodePort, and
LoadBalancer. If this field is specified when creating a Service of type
ExternalName, creation will fail. This field will be wiped when updating
a Service to type ExternalName.  If this field is not specified, it will
be initialized from the clusterIP field.  If this field is specified,
clients must ensure that clusterIPs[0] and clusterIP have the same
value.</p>
<p>This field may hold a maximum of two entries (dual-stack IPs, in either order).
These IPs must correspond to the values of the ipFamilies field. Both
clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#servicetype-v1-core">
Kubernetes core/v1.ServiceType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>type determines how the Service is exposed. Defaults to ClusterIP. Valid
options are ExternalName, ClusterIP, NodePort, and LoadBalancer.
&ldquo;ClusterIP&rdquo; allocates a cluster-internal IP address for load-balancing
to endpoints. Endpoints are determined by the selector or if that is not
specified, by manual construction of an Endpoints object or
EndpointSlice objects. If clusterIP is &ldquo;None&rdquo;, no virtual IP is
allocated and the endpoints are published as a set of endpoints rather
than a virtual IP.
&ldquo;NodePort&rdquo; builds on ClusterIP and allocates a port on every node which
routes to the same endpoints as the clusterIP.
&ldquo;LoadBalancer&rdquo; builds on NodePort and creates an external load-balancer
(if supported in the current cloud) which routes to the same endpoints
as the clusterIP.
&ldquo;ExternalName&rdquo; aliases this service to the specified externalName.
Several other fields do not apply to ExternalName services.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types">https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types</a></p>
</td>
</tr>
<tr>
<td>
<code>externalIPs</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>externalIPs is a list of IP addresses for which nodes in the cluster
will also accept traffic for this service.  These IPs are not managed by
Kubernetes.  The user is responsible for ensuring that traffic arrives
at a node with this IP.  A common example is external load-balancers
that are not part of the Kubernetes system.</p>
</td>
</tr>
<tr>
<td>
<code>sessionAffinity</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceaffinity-v1-core">
Kubernetes core/v1.ServiceAffinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Supports &ldquo;ClientIP&rdquo; and &ldquo;None&rdquo;. Used to maintain session affinity.
Enable client IP based session affinity.
Must be ClientIP or None.
Defaults to None.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerIP</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Only applies to Service Type: LoadBalancer.
This feature depends on whether the underlying cloud-provider supports specifying
the loadBalancerIP when a load balancer is created.
This field will be ignored if the cloud-provider does not support the feature.
Deprecated: This field was under-specified and its meaning varies across implementations.
Using it is non-portable and it may not support dual-stack.
Users are encouraged to use implementation-specific annotations when available.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSourceRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified and supported by the platform, this will restrict traffic through the cloud-provider
load-balancer will be restricted to the specified client IPs. This field will be ignored if the
cloud-provider does not support the feature.&rdquo;
More info: <a href="https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/">https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/</a></p>
</td>
</tr>
<tr>
<td>
<code>externalName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>externalName is the external reference that discovery mechanisms will
return as an alias for this service (e.g. a DNS CNAME record). No
proxying will be involved.  Must be a lowercase RFC-1123 hostname
(<a href="https://tools.ietf.org/html/rfc1123">https://tools.ietf.org/html/rfc1123</a>) and requires <code>type</code> to be &ldquo;ExternalName&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>externalTrafficPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceexternaltrafficpolicy-v1-core">
Kubernetes core/v1.ServiceExternalTrafficPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>externalTrafficPolicy describes how nodes distribute service traffic they
receive on one of the Service&rsquo;s &ldquo;externally-facing&rdquo; addresses (NodePorts,
ExternalIPs, and LoadBalancer IPs). If set to &ldquo;Local&rdquo;, the proxy will configure
the service in a way that assumes that external load balancers will take care
of balancing the service traffic between nodes, and so each node will deliver
traffic only to the node-local endpoints of the service, without masquerading
the client source IP. (Traffic mistakenly sent to a node with no endpoints will
be dropped.) The default value, &ldquo;Cluster&rdquo;, uses the standard behavior of
routing to all endpoints evenly (possibly modified by topology and other
features). Note that traffic sent to an External IP or LoadBalancer IP from
within the cluster will always get &ldquo;Cluster&rdquo; semantics, but clients sending to
a NodePort from within the cluster may need to take traffic policy into account
when picking a node.</p>
</td>
</tr>
<tr>
<td>
<code>healthCheckNodePort</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>healthCheckNodePort specifies the healthcheck nodePort for the service.
This only applies when type is set to LoadBalancer and
externalTrafficPolicy is set to Local. If a value is specified, is
in-range, and is not in use, it will be used.  If not specified, a value
will be automatically allocated.  External systems (e.g. load-balancers)
can use this port to determine if a given node holds endpoints for this
service or not.  If this field is specified when creating a Service
which does not need it, creation will fail. This field will be wiped
when updating a Service to no longer need it (e.g. changing type).
This field cannot be updated once set.</p>
</td>
</tr>
<tr>
<td>
<code>publishNotReadyAddresses</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>publishNotReadyAddresses indicates that any agent which deals with endpoints for this
Service should disregard any indications of ready/not-ready.
The primary use case for setting this field is for a StatefulSet&rsquo;s Headless Service to
propagate SRV DNS records for its Pods for the purpose of peer discovery.
The Kubernetes controllers that generate Endpoints and EndpointSlice resources for
Services interpret this to mean that all endpoints are considered &ldquo;ready&rdquo; even if the
Pods themselves are not. Agents which consume only Kubernetes generated endpoints
through the Endpoints or EndpointSlice resources can safely assume this behavior.</p>
</td>
</tr>
<tr>
<td>
<code>sessionAffinityConfig</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#sessionaffinityconfig-v1-core">
Kubernetes core/v1.SessionAffinityConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>sessionAffinityConfig contains the configurations of session affinity.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilies</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#ipfamily-v1-core">
[]Kubernetes core/v1.IPFamily
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this
service. This field is usually assigned automatically based on cluster
configuration and the ipFamilyPolicy field. If this field is specified
manually, the requested family is available in the cluster,
and ipFamilyPolicy allows it, it will be used; otherwise creation of
the service will fail. This field is conditionally mutable: it allows
for adding or removing a secondary IP family, but it does not allow
changing the primary IP family of the Service. Valid values are &ldquo;IPv4&rdquo;
and &ldquo;IPv6&rdquo;.  This field only applies to Services of types ClusterIP,
NodePort, and LoadBalancer, and does apply to &ldquo;headless&rdquo; services.
This field will be wiped when updating a Service to type ExternalName.</p>
<p>This field may hold a maximum of two entries (dual-stack families, in
either order).  These families must correspond to the values of the
clusterIPs field, if specified. Both clusterIPs and ipFamilies are
governed by the ipFamilyPolicy field.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilyPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#ipfamilypolicy-v1-core">
Kubernetes core/v1.IPFamilyPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>IPFamilyPolicy represents the dual-stack-ness requested or required by
this Service. If there is no value provided, then this field will be set
to SingleStack. Services can be &ldquo;SingleStack&rdquo; (a single IP family),
&ldquo;PreferDualStack&rdquo; (two IP families on dual-stack configured clusters or
a single IP family on single-stack clusters), or &ldquo;RequireDualStack&rdquo;
(two IP families on dual-stack configured clusters, otherwise fail). The
ipFamilies and clusterIPs fields depend on the value of this field. This
field will be wiped when updating a service to type ExternalName.</p>
</td>
</tr>
<tr>
<td>
<code>allocateLoadBalancerNodePorts</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>allocateLoadBalancerNodePorts defines if NodePorts will be automatically
allocated for services with type LoadBalancer.  Default is &ldquo;true&rdquo;. It
may be set to &ldquo;false&rdquo; if the cluster load-balancer does not rely on
NodePorts.  If the caller requests specific NodePorts (by specifying a
value), those requests will be respected, regardless of this field.
This field may only be set for services with type LoadBalancer and will
be cleared if the type is changed to any other type.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerClass</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>loadBalancerClass is the class of the load balancer implementation this Service belongs to.
If specified, the value of this field must be a label-style identifier, with an optional prefix,
e.g. &ldquo;internal-vip&rdquo; or &ldquo;example.com/internal-vip&rdquo;. Unprefixed names are reserved for end-users.
This field can only be set when the Service type is &lsquo;LoadBalancer&rsquo;. If not set, the default load
balancer implementation is used, today this is typically done through the cloud provider integration,
but should apply for any default implementation. If set, it is assumed that a load balancer
implementation is watching for Services with a matching class. Any default load balancer
implementation (e.g. cloud providers) should ignore Services that set this field.
This field can only be set when creating or updating a Service to type &lsquo;LoadBalancer&rsquo;.
Once set, it can not be changed. This field will be wiped when a service is updated to a non &lsquo;LoadBalancer&rsquo; type.</p>
</td>
</tr>
<tr>
<td>
<code>internalTrafficPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#serviceinternaltrafficpolicy-v1-core">
Kubernetes core/v1.ServiceInternalTrafficPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>InternalTrafficPolicy describes how nodes distribute service traffic they
receive on the ClusterIP. If set to &ldquo;Local&rdquo;, the proxy will assume that pods
only want to talk to endpoints of the service on the same node as the pod,
dropping the traffic if there are no local endpoints. The default value,
&ldquo;Cluster&rdquo;, uses the standard behavior of routing to all endpoints evenly
(possibly modified by topology and other features).</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>roleSelector</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extends the above <code>serviceSpec.selector</code> by allowing you to specify defined role as selector for the service.
When <code>roleSelector</code> is set, it adds a label selector &ldquo;kubeblocks.io/role: &#123;roleSelector&#125;&rdquo;
to the <code>serviceSpec.selector</code>.
Example usage:</p>
<pre><code>  roleSelector: &quot;leader&quot;
</code></pre>
<p>In this example, setting <code>roleSelector</code> to &ldquo;leader&rdquo; will add a label selector
&ldquo;kubeblocks.io/role: leader&rdquo; to the <code>serviceSpec.selector</code>.
This means that the service will select and route traffic to Pods with the label
&ldquo;kubeblocks.io/role&rdquo; set to &ldquo;leader&rdquo;.</p>
<p>Note that if <code>podService</code> sets to true, RoleSelector will be ignored.
The <code>podService</code> flag takes precedence over <code>roleSelector</code> and generates a service for each Pod.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceDescriptorSpec">ServiceDescriptorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptor">ServiceDescriptor</a>)
</p>
<div>
<p>ServiceDescriptorSpec defines the desired state of ServiceDescriptor.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the type of database service provided by the external service.
For example, &ldquo;mysql&rdquo;, &ldquo;redis&rdquo;, &ldquo;mongodb&rdquo;.
This field categorizes databases by their functionality, protocol and compatibility, facilitating appropriate
service integration based on their unique capabilities.</p>
<p>This field is case-insensitive.</p>
<p>It also supports abbreviations for some well-known databases:
- &ldquo;pg&rdquo;, &ldquo;pgsql&rdquo;, &ldquo;postgres&rdquo;, &ldquo;postgresql&rdquo;: PostgreSQL service
- &ldquo;zk&rdquo;, &ldquo;zookeeper&rdquo;: ZooKeeper service
- &ldquo;es&rdquo;, &ldquo;elasticsearch&rdquo;: Elasticsearch service
- &ldquo;mongo&rdquo;, &ldquo;mongodb&rdquo;: MongoDB service
- &ldquo;ch&rdquo;, &ldquo;clickhouse&rdquo;: ClickHouse service</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>Describes the version of the service provided by the external service.
This is crucial for ensuring compatibility between different components of the system,
as different versions of a service may have varying features.</p>
</td>
</tr>
<tr>
<td>
<code>endpoint</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the endpoint of the external service.</p>
<p>If the service is exposed via a cluster, the endpoint will be provided in the format of <code>host:port</code>.</p>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the service or IP address of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the port of the external service.</p>
</td>
</tr>
<tr>
<td>
<code>auth</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConnectionCredentialAuth">
ConnectionCredentialAuth
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the authentication credentials required for accessing an external service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceDescriptorStatus">ServiceDescriptorStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceDescriptor">ServiceDescriptor</a>)
</p>
<div>
<p>ServiceDescriptorStatus defines the observed state of ServiceDescriptor</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Phase">
Phase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates the current lifecycle phase of the ServiceDescriptor. This can be either &lsquo;Available&rsquo; or &lsquo;Unavailable&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides a human-readable explanation detailing the reason for the current phase of the ServiceConnectionCredential.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the generation number that has been processed by the controller.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServicePort">ServicePort
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceSpec">ServiceSpec</a>)
</p>
<div>
<p>ServicePort is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of this port within the service. This must be a DNS_LABEL.
All ports within a ServiceSpec must have unique names. When considering
the endpoints for a Service, this must match the &lsquo;name&rsquo; field in the
EndpointPort.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#protocol-v1-core">
Kubernetes core/v1.Protocol
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The IP protocol for this port. Supports &ldquo;TCP&rdquo;, &ldquo;UDP&rdquo;, and &ldquo;SCTP&rdquo;.
Default is TCP.</p>
</td>
</tr>
<tr>
<td>
<code>appProtocol</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The application protocol for this port.
This field follows standard Kubernetes label syntax.
Un-prefixed names are reserved for IANA standard service names (as per
RFC-6335 and <a href="https://www.iana.org/assignments/service-names)">https://www.iana.org/assignments/service-names)</a>.
Non-standard protocols should use prefixed names such as
mycompany.com/my-custom-protocol.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int32
</em>
</td>
<td>
<p>The port that will be exposed by this service.</p>
</td>
</tr>
<tr>
<td>
<code>targetPort</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Number or name of the port to access on the pods targeted by the service.</p>
<p>Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.</p>
<ul>
<li>If this is a string, it will be looked up as a named port in the target Pod&rsquo;s container ports.</li>
<li>If this is not specified, the value of the <code>port</code> field is used (an identity map).</li>
</ul>
<p>This field is ignored for services with clusterIP=None, and should be
omitted or set equal to the <code>port</code> field.</p>
<p>More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service">https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRef">ServiceRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the identifier of the service reference declaration.
It corresponds to the serviceRefDeclaration name defined in either:</p>
<ul>
<li><code>componentDefinition.spec.serviceRefDeclarations[*].name</code></li>
<li><code>clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name</code> (deprecated)</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object.
If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current
Cluster by default.</p>
</td>
</tr>
<tr>
<td>
<code>cluster</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the KubeBlocks Cluster being referenced.
This is used when services from another KubeBlocks Cluster are consumed.</p>
<p>By default, the referenced KubeBlocks Cluster&rsquo;s <code>clusterDefinition.spec.connectionCredential</code>
will be utilized to bind to the current Component. This credential should include:
<code>endpoint</code>, <code>port</code>, <code>username</code>, and <code>password</code>.</p>
<p>Note:</p>
<ul>
<li>The <code>ServiceKind</code> and <code>ServiceVersion</code> specified in the service reference within the
ClusterDefinition are not validated when using this approach.</li>
<li>If both <code>cluster</code> and <code>serviceDescriptor</code> are present, <code>cluster</code> will take precedence.</li>
</ul>
<p>Deprecated since v0.9 since <code>clusterDefinition.spec.connectionCredential</code> is deprecated,
use <code>clusterServiceSelector</code> instead.
This field is maintained for backward compatibility and its use is discouraged.
Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.</p>
</td>
</tr>
<tr>
<td>
<code>clusterServiceSelector</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefClusterSelector">
ServiceRefClusterSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>References a service provided by another KubeBlocks Cluster.
It specifies the ClusterService and the account credentials needed for access.</p>
</td>
</tr>
<tr>
<td>
<code>serviceDescriptor</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.</p>
<p>When referencing a service provided by external sources, a ServiceDescriptor object is required to establish
the service binding.
The <code>serviceDescriptor.spec.serviceKind</code> and <code>serviceDescriptor.spec.serviceVersion</code> should match the serviceKind
and serviceVersion declared in the definition.</p>
<p>If both <code>cluster</code> and <code>serviceDescriptor</code> are specified, the <code>cluster</code> takes precedence.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRefClusterSelector">ServiceRefClusterSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceRef">ServiceRef</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cluster</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the Cluster being referenced.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefServiceSelector">
ServiceRefServiceSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Identifies a ClusterService from the list of Services defined in <code>cluster.spec.services</code> of the referenced Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>credential</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefCredentialSelector">
ServiceRefCredentialSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster.
The SystemAccount should be defined in <code>componentDefinition.spec.systemAccounts</code>
of the Component providing the service in the referenced Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRefCredentialSelector">ServiceRefCredentialSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceRefClusterSelector">ServiceRefClusterSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>component</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the Component where the credential resides in.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the credential (SystemAccount) to reference.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRefDeclaration">ServiceRefDeclaration
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
<p>ServiceRefDeclaration represents a reference to a service that can be either provided by a KubeBlocks Cluster
or an external service.
It acts as a placeholder for the actual service reference, which is determined later when a Cluster is created.</p>
<p>The purpose of ServiceRefDeclaration is to declare a service dependency without specifying the concrete details
of the service.
It allows for flexibility and abstraction in defining service references within a Component.
By using ServiceRefDeclaration, you can define service dependencies in a declarative manner, enabling loose coupling
and easier management of service references across different components and clusters.</p>
<p>Upon Cluster creation, the ServiceRefDeclaration is bound to an actual service through the ServiceRef field,
effectively resolving and connecting to the specified service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the ServiceRefDeclaration.</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefDeclarationSpecs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefDeclarationSpec">
[]ServiceRefDeclarationSpec
</a>
</em>
</td>
<td>
<p>Defines a list of constraints and requirements for services that can be bound to this ServiceRefDeclaration
upon Cluster creation.
Each ServiceRefDeclarationSpec defines a ServiceKind and ServiceVersion,
outlining the acceptable service types and versions that are compatible.</p>
<p>This flexibility allows a ServiceRefDeclaration to be fulfilled by any one of the provided specs.
For example, if it requires an OLTP database, specs for both MySQL and PostgreSQL are listed,
either MySQL or PostgreSQL services can be used when binding.</p>
</td>
</tr>
<tr>
<td>
<code>optional</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether the service reference can be optional.</p>
<p>For an optional service-ref, the component can still be created even if the service-ref is not provided.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRefDeclarationSpec">ServiceRefDeclarationSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceRefDeclaration">ServiceRefDeclaration</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceKind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the type or nature of the service. This should be a well-known application cluster type, such as
&#123;mysql, redis, mongodb&#125;.
The field is case-insensitive and supports abbreviations for some well-known databases.
For instance, both <code>zk</code> and <code>zookeeper</code> are considered as a ZooKeeper cluster, while <code>pg</code>, <code>postgres</code>, <code>postgresql</code>
are all recognized as a PostgreSQL cluster.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the service version of the service reference. This is a regular expression that matches a version number pattern.
For instance, <code>^8.0.8$</code>, <code>8.0.\d&#123;1,2&#125;$</code>, <code>^[v\-]*?(\d&#123;1,2&#125;\.)&#123;0,3&#125;\d&#123;1,2&#125;$</code> are all valid patterns.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRefServiceSelector">ServiceRefServiceSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceRefClusterSelector">ServiceRefClusterSelector</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>component</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the Component where the Service resides in.</p>
<p>It is required when referencing a Component&rsquo;s Service.</p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the Service to be referenced.</p>
<p>Leave it empty to reference the default Service. Set it to &ldquo;headless&rdquo; to reference the default headless Service.</p>
<p>If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,
and the resolved value will be presented in the following format: service1.name,service2.name&hellip;</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The port name of the Service to be referenced.</p>
<p>If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.</p>
<p>If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched,
and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2&hellip;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRefVarSelector">ServiceRefVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.VarSource">VarSource</a>)
</p>
<div>
<p>ServiceRefVarSelector selects a var from a ServiceRefDeclaration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The ServiceRefDeclaration to select from.</p>
</td>
</tr>
<tr>
<td>
<code>ServiceRefVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefVars">
ServiceRefVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ServiceRefVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceRefVars">ServiceRefVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceRefVarSelector">ServiceRefVarSelector</a>)
</p>
<div>
<p>ServiceRefVars defines the vars that can be referenced from a ServiceRef.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>endpoint</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>CredentialVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVars">
CredentialVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>CredentialVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceSpec">ServiceSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>ServiceSpec is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ports</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServicePort">
[]ServicePort
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The list of ports that are exposed by this service.
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies">https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceVarSelector">ServiceVarSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.VarSource">VarSource</a>)
</p>
<div>
<p>ServiceVarSelector selects a var from a Service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ClusterObjectReference</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterObjectReference">
ClusterObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>ClusterObjectReference</code> are embedded into this type.)
</p>
<p>The Service to select from.
It can be referenced from the default headless service by setting the name to &ldquo;headless&rdquo;.</p>
</td>
</tr>
<tr>
<td>
<code>ServiceVars</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceVars">
ServiceVars
</a>
</em>
</td>
<td>
<p>
(Members of <code>ServiceVars</code> are embedded into this type.)
</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ServiceVars">ServiceVars
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ServiceVarSelector">ServiceVarSelector</a>)
</p>
<div>
<p>ServiceVars defines the vars that can be referenced from a Service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>host</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
</td>
</tr>
<tr>
<td>
<code>loadBalancer</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VarOption">
VarOption
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>LoadBalancer represents the LoadBalancer ingress point of the service.</p>
<p>If multiple ingress points are available, the first one will be used automatically, choosing between IP and Hostname.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.NamedVar">
NamedVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Port references a port or node-port defined in the service.</p>
<p>If the referenced service is a pod-service, there will be multiple service objects matched,
and the value will be presented in the following format: service1.name:port1,service2.name:port2&hellip;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.ShardingSpec">ShardingSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>ShardingSpec defines how KubeBlocks manage dynamic provisioned shards.
A typical design pattern for distributed databases is to distribute data across multiple shards,
with each shard consisting of multiple replicas.
Therefore, KubeBlocks supports representing a shard with a Component and dynamically instantiating Components
using a template when shards are added.
When shards are removed, the corresponding Components are also deleted.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the common parent part of all shard names.
This identifier is included as part of the Service DNS name and must comply with IANA service naming rules.
It is used to generate the names of underlying Components following the pattern <code>$(shardingSpec.name)-$(ShardID)</code>.
ShardID is a random string that is appended to the Name to generate unique identifiers for each shard.
For example, if the sharding specification name is &ldquo;my-shard&rdquo; and the ShardID is &ldquo;abc&rdquo;, the resulting Component name
would be &ldquo;my-shard-abc&rdquo;.</p>
<p>Note that the name defined in Component template(<code>shardingSpec.template.name</code>) will be disregarded
when generating the Component names of the shards. The <code>shardingSpec.name</code> field takes precedence.</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">
ClusterComponentSpec
</a>
</em>
</td>
<td>
<p>The template for generating Components for shards, where each shard consists of one Component.
This field is of type ClusterComponentSpec, which encapsulates all the required details and
definitions for creating and managing the Components.
KubeBlocks uses this template to generate a set of identical Components or shards.
All the generated Components will have the same specifications and definitions as specified in the <code>template</code> field.</p>
<p>This allows for the creation of multiple Components with consistent configurations,
enabling sharding and distribution of workloads across Components.</p>
</td>
</tr>
<tr>
<td>
<code>shards</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Specifies the desired number of shards.
Users can declare the desired number of shards through this field.
KubeBlocks dynamically creates and deletes Components based on the difference
between the desired and actual number of shards.
KubeBlocks provides lifecycle management for sharding, including:</p>
<ul>
<li>Executing the postProvision Action defined in the ComponentDefinition when the number of shards increases.
This allows for custom actions to be performed after a new shard is provisioned.</li>
<li>Executing the preTerminate Action defined in the ComponentDefinition when the number of shards decreases.
This enables custom cleanup or data migration tasks to be executed before a shard is terminated.
Resources and data associated with the corresponding Component will also be deleted.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.StatefulSetSpec">StatefulSetSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>, <a href="#apps.kubeblocks.io/v1alpha1.ConsensusSetSpec">ConsensusSetSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ReplicationSetSpec">ReplicationSetSpec</a>)
</p>
<div>
<p>StatefulSetSpec is deprecated since v0.7.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.UpdateStrategy">
UpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the strategy for updating Pods.
For workloadType=<code>Consensus</code>, the update strategy can be one of the following:</p>
<ul>
<li><code>Serial</code>: Updates Members sequentially to minimize component downtime.</li>
<li><code>BestEffortParallel</code>: Updates Members in parallel to minimize component write downtime. Majority remains online
at all times.</li>
<li><code>Parallel</code>: Forces parallel updates.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>llPodManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the creation of pods during initial scale up, replacement of pods on nodes, and scaling down.</p>
<ul>
<li><code>OrderedReady</code>: Creates pods in increasing order (pod-0, then pod-1, etc). The controller waits until each pod
is ready before continuing. Pods are removed in reverse order when scaling down.</li>
<li><code>Parallel</code>: Creates pods in parallel to match the desired scale without waiting. All pods are deleted at once
when scaling down.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>llUpdateStrategy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#statefulsetupdatestrategy-v1-apps">
Kubernetes apps/v1.StatefulSetUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the low-level StatefulSetUpdateStrategy to be used when updating Pods in the StatefulSet upon a
revision to the Template.
<code>UpdateStrategy</code> will be ignored if this is provided.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.StatefulSetWorkload">StatefulSetWorkload
</h3>
<div>
<p>StatefulSetWorkload interface</p>
</div>
<h3 id="apps.kubeblocks.io/v1alpha1.StatelessSetSpec">StatelessSetSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>StatelessSetSpec is deprecated since v0.7.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#deploymentstrategy-v1-apps">
Kubernetes apps/v1.DeploymentStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the deployment strategy that will be used to replace existing pods with new ones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SwitchPolicyType">SwitchPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSwitchPolicy">ClusterSwitchPolicy</a>)
</p>
<div>
<p>SwitchPolicyType defines the types of switch policies that can be applied to a cluster.</p>
<p>Currently, only the Noop policy is supported. Support for MaximumAvailability and MaximumDataProtection policies is
planned for future releases.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;MaximumAvailability&#34;</p></td>
<td><p>MaximumAvailability represents a switch policy that aims for maximum availability. This policy will switch if the
primary is active and the synchronization delay is 0 according to the user-defined lagProbe data delay detection
logic. If the primary is down, it will switch immediately.
This policy is intended for future support.</p>
</td>
</tr><tr><td><p>&#34;MaximumDataProtection&#34;</p></td>
<td><p>MaximumDataProtection represents a switch policy focused on maximum data protection. This policy will only switch
if the primary is active and the synchronization delay is 0, based on the user-defined lagProbe data lag detection
logic. If the primary is down, it will switch only if it can be confirmed that the primary and secondary data are
consistent. Otherwise, it will not switch.
This policy is planned for future implementation.</p>
</td>
</tr><tr><td><p>&#34;Noop&#34;</p></td>
<td><p>Noop indicates that KubeBlocks will not perform any high-availability switching for the components. Users are
required to implement their own HA solution or integrate an existing open-source HA solution.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SwitchoverAction">SwitchoverAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.SwitchoverSpec">SwitchoverSpec</a>)
</p>
<div>
<p>SwitchoverAction is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cmdExecutorConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CmdExecutorConfig">
CmdExecutorConfig
</a>
</em>
</td>
<td>
<p>Specifies the switchover command.</p>
</td>
</tr>
<tr>
<td>
<code>scriptSpecSelectors</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ScriptSpecSelector">
[]ScriptSpecSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to select the script that need to be referenced.
When defined, the scripts defined in scriptSpecs can be referenced within the SwitchoverAction.CmdExecutorConfig.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SwitchoverSpec">SwitchoverSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>SwitchoverSpec is deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>withCandidate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SwitchoverAction">
SwitchoverAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the action of switching over to a specified candidate primary or leader instance.</p>
</td>
</tr>
<tr>
<td>
<code>withoutCandidate</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SwitchoverAction">
SwitchoverAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the action of switching over without specifying a candidate primary or leader instance.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SystemAccount">SystemAccount
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the unique identifier for the account. This name is used by other entities to reference the account.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>initAccount</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates if this account is a system initialization account (e.g., MySQL root).</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>statement</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the statement used to create the account with the necessary privileges.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>passwordGenerationPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PasswordConfig">
PasswordConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the policy for generating the account&rsquo;s password.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
<tr>
<td>
<code>secretRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProvisionSecretRef">
ProvisionSecretRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the secret from which data will be copied to create the new account.</p>
<p>This field is immutable once set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SystemAccountConfig">SystemAccountConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.SystemAccountSpec">SystemAccountSpec</a>)
</p>
<div>
<p>SystemAccountConfig specifies how to create and delete system accounts.</p>
<p>Deprecated since v0.9.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.AccountName">
AccountName
</a>
</em>
</td>
<td>
<p>The unique identifier of a system account.</p>
</td>
</tr>
<tr>
<td>
<code>provisionPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProvisionPolicy">
ProvisionPolicy
</a>
</em>
</td>
<td>
<p>Outlines the strategy for creating the account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.SystemAccountSpec">SystemAccountSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>SystemAccountSpec specifies information to create system accounts.</p>
<p>Deprecated since v0.8, be replaced by <code>componentDefinition.spec.systemAccounts</code> and
<code>componentDefinition.spec.lifecycleActions.accountProvision</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>cmdExecutorConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CmdExecutorConfig">
CmdExecutorConfig
</a>
</em>
</td>
<td>
<p>Configures how to obtain the client SDK and execute statements.</p>
</td>
</tr>
<tr>
<td>
<code>passwordConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.PasswordConfig">
PasswordConfig
</a>
</em>
</td>
<td>
<p>Defines the pattern used to generate passwords for system accounts.</p>
</td>
</tr>
<tr>
<td>
<code>accounts</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SystemAccountConfig">
[]SystemAccountConfig
</a>
</em>
</td>
<td>
<p>Defines the configuration settings for system accounts.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.TLSConfig">TLSConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enable</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>A boolean flag that indicates whether the Component should use Transport Layer Security (TLS)
for secure communication.
When set to true, the Component will be configured to use TLS encryption for its network connections.
This ensures that the data transmitted between the Component and its clients or other Components is encrypted
and protected from unauthorized access.
If TLS is enabled, the Component may require additional configuration,
such as specifying TLS certificates and keys, to properly set up the secure communication channel.</p>
</td>
</tr>
<tr>
<td>
<code>issuer</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.Issuer">
Issuer
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the configuration for the TLS certificates issuer.
It allows defining the issuer name and the reference to the secret containing the TLS certificates and key.
The secret should contain the CA certificate, TLS certificate, and private key in the specified keys.
Required when TLS is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.TLSSecretRef">TLSSecretRef
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Issuer">Issuer</a>)
</p>
<div>
<p>TLSSecretRef defines Secret contains Tls certs</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name of the Secret that contains user-provided certificates.</p>
</td>
</tr>
<tr>
<td>
<code>ca</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key of CA cert in Secret</p>
</td>
</tr>
<tr>
<td>
<code>cert</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key of Cert in Secret</p>
</td>
</tr>
<tr>
<td>
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key of TLS private key in Secret</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.TargetPodSelector">TargetPodSelector
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Action">Action</a>)
</p>
<div>
<p>TargetPodSelector defines how to select pod(s) to execute an Action.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;All&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Any&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ordinal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Role&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.TenancyType">TenancyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.Affinity">Affinity</a>, <a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>TenancyType defines the type of tenancy for cluster tenant resources.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;DedicatedNode&#34;</p></td>
<td><p>DedicatedNode means each pod runs on their own dedicated node.</p>
</td>
</tr><tr><td><p>&#34;SharedNode&#34;</p></td>
<td><p>SharedNode means multiple pods may share the same node.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.TerminationPolicyType">TerminationPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterSpec">ClusterSpec</a>)
</p>
<div>
<p>TerminationPolicyType defines termination policy types.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Delete&#34;</p></td>
<td><p>Delete is based on Halt and deletes PVCs.</p>
</td>
</tr><tr><td><p>&#34;DoNotTerminate&#34;</p></td>
<td><p>DoNotTerminate will block delete operation.</p>
</td>
</tr><tr><td><p>&#34;Halt&#34;</p></td>
<td><p>Halt will delete workload resources such as statefulset, deployment workloads but keep PVCs.</p>
</td>
</tr><tr><td><p>&#34;WipeOut&#34;</p></td>
<td><p>WipeOut is based on Delete and wipe out all volume snapshots and snapshot data from backup storage location.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.UpdateStrategy">UpdateStrategy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentDefinitionSpec">ComponentDefinitionSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.StatefulSetSpec">StatefulSetSpec</a>)
</p>
<div>
<p>UpdateStrategy defines the update strategy for cluster components. This strategy determines how updates are applied
across the cluster.
The available strategies are <code>Serial</code>, <code>BestEffortParallel</code>, and <code>Parallel</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;BestEffortParallel&#34;</p></td>
<td><p>BestEffortParallelStrategy indicates that the replicas are updated in parallel, with the operator making
a best-effort attempt to update as many replicas as possible concurrently
while maintaining the component&rsquo;s availability.
Unlike the <code>Parallel</code> strategy, the <code>BestEffortParallel</code> strategy aims to ensure that a minimum number
of replicas remain available during the update process to maintain the component&rsquo;s quorum and functionality.</p>
<p>For example, consider a component with 5 replicas. To maintain the component&rsquo;s availability and quorum,
the operator may allow a maximum of 2 replicas to be simultaneously updated. This ensures that at least
3 replicas (a quorum) remain available and functional during the update process.</p>
<p>The <code>BestEffortParallel</code> strategy strikes a balance between update speed and component availability.</p>
</td>
</tr><tr><td><p>&#34;Parallel&#34;</p></td>
<td><p>ParallelStrategy indicates that updates are applied simultaneously to all Pods of a Component.
The replicas are updated in parallel, with the operator updating all replicas concurrently.
This strategy provides the fastest update time but may lead to a period of reduced availability or
capacity during the update process.</p>
</td>
</tr><tr><td><p>&#34;Serial&#34;</p></td>
<td><p>SerialStrategy indicates that updates are applied one at a time in a sequential manner.
The operator waits for each replica to be updated and ready before proceeding to the next one.
This ensures that only one replica is unavailable at a time during the update process.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.UpgradePolicy">UpgradePolicy
(<code>string</code> alias)</h3>
<div>
<p>UpgradePolicy defines the policy of reconfiguring.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;autoReload&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;dynamicReloadBeginRestart&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;simple&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;parallel&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;rolling&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;operatorSyncUpdate&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.UserResourceRefs">UserResourceRefs
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>)
</p>
<div>
<p>UserResourceRefs defines references to user-defined Secrets and ConfigMaps.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secretRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.SecretRef">
[]SecretRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>SecretRefs defines the user-defined Secrets.</p>
</td>
</tr>
<tr>
<td>
<code>configMapRefs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ConfigMapRef">
[]ConfigMapRef
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigMapRefs defines the user-defined ConfigMaps.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.VarOption">VarOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ComponentVars">ComponentVars</a>, <a href="#apps.kubeblocks.io/v1alpha1.CredentialVars">CredentialVars</a>, <a href="#apps.kubeblocks.io/v1alpha1.NamedVar">NamedVar</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceRefVars">ServiceRefVars</a>, <a href="#apps.kubeblocks.io/v1alpha1.ServiceVars">ServiceVars</a>)
</p>
<div>
<p>VarOption defines whether a variable is required or optional.</p>
</div>
<h3 id="apps.kubeblocks.io/v1alpha1.VarSource">VarSource
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.EnvVar">EnvVar</a>)
</p>
<div>
<p>VarSource represents a source for the value of an EnvVar.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>configMapKeyRef</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#configmapkeyselector-v1-core">
Kubernetes core/v1.ConfigMapKeySelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a key of a ConfigMap.</p>
</td>
</tr>
<tr>
<td>
<code>secretKeyRef</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#secretkeyselector-v1-core">
Kubernetes core/v1.SecretKeySelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a key of a Secret.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetworkVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.HostNetworkVarSelector">
HostNetworkVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of host-network resources.</p>
</td>
</tr>
<tr>
<td>
<code>serviceVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceVarSelector">
ServiceVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a Service.</p>
</td>
</tr>
<tr>
<td>
<code>credentialVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.CredentialVarSelector">
CredentialVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a Credential (SystemAccount).</p>
</td>
</tr>
<tr>
<td>
<code>serviceRefVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ServiceRefVarSelector">
ServiceRefVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a ServiceRef.</p>
</td>
</tr>
<tr>
<td>
<code>componentVarRef</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ComponentVarSelector">
ComponentVarSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Selects a defined var of a Component.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.VolumeProtectionSpec">VolumeProtectionSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>VolumeProtectionSpec is deprecated since v0.9, replaced with ComponentVolume.HighWatermark.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>highWatermark</code><br/>
<em>
int
</em>
</td>
<td>
<em>(Optional)</em>
<p>The high watermark threshold for volume space usage.
If there is any specified volumes who&rsquo;s space usage is over the threshold, the pre-defined &ldquo;LOCK&rdquo; action
will be triggered to degrade the service to protect volume from space exhaustion, such as to set the instance
as read-only. And after that, if all volumes&rsquo; space usage drops under the threshold later, the pre-defined
&ldquo;UNLOCK&rdquo; action will be performed to recover the service normally.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.ProtectedVolume">
[]ProtectedVolume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The Volumes to be protected.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.VolumeType">VolumeType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.VolumeTypeSpec">VolumeTypeSpec</a>)
</p>
<div>
<p>VolumeType defines the type of volume, specifically distinguishing between volumes used for backup data and those used for logs.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;data&#34;</p></td>
<td><p>VolumeTypeData indicates a volume designated for storing backup data. This type of volume is optimized for the
storage and retrieval of data backups, ensuring data persistence and reliability.</p>
</td>
</tr><tr><td><p>&#34;log&#34;</p></td>
<td><p>VolumeTypeLog indicates a volume designated for storing logs. This type of volume is optimized for log data,
facilitating efficient log storage, retrieval, and management.</p>
</td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.VolumeTypeSpec">VolumeTypeSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>VolumeTypeSpec is deprecated since v0.9, replaced with ComponentVolume.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Corresponds to the name of the VolumeMounts field in PodSpec.Container.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1alpha1.VolumeType">
VolumeType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Type of data the volume will persistent.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1alpha1.WorkloadType">WorkloadType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentDefinition">ClusterComponentDefinition</a>)
</p>
<div>
<p>WorkloadType defines the type of workload for the components of the ClusterDefinition.
It can be one of the following: <code>Stateless</code>, <code>Stateful</code>, <code>Consensus</code>, or <code>Replication</code>.</p>
<p>Deprecated since v0.8.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Consensus&#34;</p></td>
<td><p>Consensus represents a workload type involving distributed consensus algorithms for coordinated decision-making.</p>
</td>
</tr><tr><td><p>&#34;Replication&#34;</p></td>
<td><p>Replication represents a workload type that involves replication, typically used for achieving high availability
and fault tolerance.</p>
</td>
</tr><tr><td><p>&#34;Stateful&#34;</p></td>
<td><p>Stateful represents a workload type where components maintain state, and each instance has a unique identity.</p>
</td>
</tr><tr><td><p>&#34;Stateless&#34;</p></td>
<td><p>Stateless represents a workload type where components do not maintain state, and instances are interchangeable.</p>
</td>
</tr></tbody>
</table>
<hr/>
<h2 id="apps.kubeblocks.io/v1beta1">apps.kubeblocks.io/v1beta1</h2>
<div>
</div>
Resource Types:
<ul><li>
<a href="#apps.kubeblocks.io/v1beta1.ConfigConstraint">ConfigConstraint</a>
</li></ul>
<h3 id="apps.kubeblocks.io/v1beta1.ConfigConstraint">ConfigConstraint
</h3>
<div>
<p>ConfigConstraint manages the parameters across multiple configuration files contained in a single configure template.
These configuration files should have the same format (e.g. ini, xml, properties, json).</p>
<p>It provides the following functionalities:</p>
<ol>
<li><strong>Parameter Value Validation</strong>: Validates and ensures compliance of parameter values with defined constraints.</li>
<li><strong>Dynamic Reload on Modification</strong>: Monitors parameter changes and triggers dynamic reloads to apply updates.</li>
<li><strong>Parameter Rendering in Templates</strong>: Injects parameters into templates to generate up-to-date configuration files.</li>
</ol>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>apps.kubeblocks.io/v1beta1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>ConfigConstraint</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintSpec">
ConfigConstraintSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>reloadAction</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ReloadAction">
ReloadAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the dynamic reload (dynamic reconfiguration) actions supported by the engine.
When set, the controller executes the scripts defined in these actions to handle dynamic parameter updates.</p>
<p>Dynamic reloading is triggered only if both of the following conditions are met:</p>
<ol>
<li>The modified parameters are listed in the <code>dynamicParameters</code> field.
If <code>dynamicParameterSelectedPolicy</code> is set to &ldquo;all&rdquo;, modifications to <code>staticParameters</code>
can also trigger a reload.</li>
<li><code>reloadAction</code> is set.</li>
</ol>
<p>If <code>reloadAction</code> is not set or the modified parameters are not listed in <code>dynamicParameters</code>,
dynamic reloading will not be triggered.</p>
<p>Example:</p>
<pre><code class="language-yaml">dynamicReloadAction:
 tplScriptTrigger:
   namespace: kb-system
   scriptConfigMapRef: mysql-reload-script
   sync: true
</code></pre>
</td>
</tr>
<tr>
<td>
<code>mergeReloadAndRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to consolidate dynamic reload and restart actions into a single restart.</p>
<ul>
<li>If true, updates requiring both actions will result in only a restart, merging the actions.</li>
<li>If false, updates will trigger both actions executed sequentially: first dynamic reload, then restart.</li>
</ul>
<p>This flag allows for more efficient handling of configuration changes by potentially eliminating
an unnecessary reload step.</p>
</td>
</tr>
<tr>
<td>
<code>reloadStaticParamsBeforeRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Configures whether the dynamic reload specified in <code>reloadAction</code> applies only to dynamic parameters or
to all parameters (including static parameters).</p>
<ul>
<li>false (default): Only modifications to the dynamic parameters listed in <code>dynamicParameters</code>
will trigger a dynamic reload.</li>
<li>true: Modifications to both dynamic parameters listed in <code>dynamicParameters</code> and static parameters
listed in <code>staticParameters</code> will trigger a dynamic reload.
The &ldquo;all&rdquo; option is for certain engines that require static parameters to be set
via SQL statements before they can take effect on restart.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>downwardAPIChangeTriggeredActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.DownwardAPIChangeTriggeredAction">
[]DownwardAPIChangeTriggeredAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of actions to execute specified commands based on Pod labels.</p>
<p>It utilizes the K8s Downward API to mount label information as a volume into the pod.
The &lsquo;config-manager&rsquo; sidecar container watches for changes in the role label and dynamically invoke
registered commands (usually execute some SQL statements) when a change is detected.</p>
<p>It is designed for scenarios where:</p>
<ul>
<li>Replicas with different roles have different configurations, such as Redis primary &amp; secondary replicas.</li>
<li>After a role switch (e.g., from secondary to primary), some changes in configuration are needed
to reflect the new role.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>parametersSchema</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ParametersSchema">
ParametersSchema
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of parameters including their names, default values, descriptions,
types, and constraints (permissible values or the range of valid values).</p>
</td>
</tr>
<tr>
<td>
<code>staticParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List static parameters.
Modifications to any of these parameters require a restart of the process to take effect.</p>
</td>
</tr>
<tr>
<td>
<code>dynamicParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List dynamic parameters.
Modifications to these parameters trigger a configuration reload without requiring a process restart.</p>
</td>
</tr>
<tr>
<td>
<code>immutableParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists the parameters that cannot be modified once set.
Attempting to change any of these parameters will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>fileFormatConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.FileFormatConfig">
FileFormatConfig
</a>
</em>
</td>
<td>
<p>Specifies the format of the configuration file and any associated parameters that are specific to the chosen format.
Supported formats include <code>ini</code>, <code>xml</code>, <code>yaml</code>, <code>json</code>, <code>hcl</code>, <code>dotenv</code>, <code>properties</code>, and <code>toml</code>.</p>
<p>Each format may have its own set of parameters that can be configured.
For instance, when using the <code>ini</code> format, you can specify the section name.</p>
<p>Example:</p>
<pre><code>fileFormatConfig:
 format: ini
 iniConfig:
   sectionName: mysqld
</code></pre>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintStatus">
ConfigConstraintStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.AutoTrigger">AutoTrigger
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ReloadOptions">ReloadOptions</a>, <a href="#apps.kubeblocks.io/v1beta1.ReloadAction">ReloadAction</a>)
</p>
<div>
<p>AutoTrigger automatically perform the reload when specified conditions are met.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>processName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the process.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.CfgFileFormat">CfgFileFormat
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.FileFormatConfig">FileFormatConfig</a>)
</p>
<div>
<p>CfgFileFormat defines formatter of configuration files.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;dotenv&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;hcl&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ini&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;json&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;properties&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;props-plus&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;redis&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;toml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;xml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;yaml&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ConfigConstraintPhase">ConfigConstraintPhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintStatus">ConfigConstraintStatus</a>, <a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintStatus">ConfigConstraintStatus</a>)
</p>
<div>
<p>ConfigConstraintPhase defines the ConfigConstraint  CR .status.phase</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Available&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unavailable&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ConfigConstraintSpec">ConfigConstraintSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.ConfigConstraint">ConfigConstraint</a>)
</p>
<div>
<p>ConfigConstraintSpec defines the desired state of ConfigConstraint</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reloadAction</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ReloadAction">
ReloadAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the dynamic reload (dynamic reconfiguration) actions supported by the engine.
When set, the controller executes the scripts defined in these actions to handle dynamic parameter updates.</p>
<p>Dynamic reloading is triggered only if both of the following conditions are met:</p>
<ol>
<li>The modified parameters are listed in the <code>dynamicParameters</code> field.
If <code>dynamicParameterSelectedPolicy</code> is set to &ldquo;all&rdquo;, modifications to <code>staticParameters</code>
can also trigger a reload.</li>
<li><code>reloadAction</code> is set.</li>
</ol>
<p>If <code>reloadAction</code> is not set or the modified parameters are not listed in <code>dynamicParameters</code>,
dynamic reloading will not be triggered.</p>
<p>Example:</p>
<pre><code class="language-yaml">dynamicReloadAction:
 tplScriptTrigger:
   namespace: kb-system
   scriptConfigMapRef: mysql-reload-script
   sync: true
</code></pre>
</td>
</tr>
<tr>
<td>
<code>mergeReloadAndRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether to consolidate dynamic reload and restart actions into a single restart.</p>
<ul>
<li>If true, updates requiring both actions will result in only a restart, merging the actions.</li>
<li>If false, updates will trigger both actions executed sequentially: first dynamic reload, then restart.</li>
</ul>
<p>This flag allows for more efficient handling of configuration changes by potentially eliminating
an unnecessary reload step.</p>
</td>
</tr>
<tr>
<td>
<code>reloadStaticParamsBeforeRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Configures whether the dynamic reload specified in <code>reloadAction</code> applies only to dynamic parameters or
to all parameters (including static parameters).</p>
<ul>
<li>false (default): Only modifications to the dynamic parameters listed in <code>dynamicParameters</code>
will trigger a dynamic reload.</li>
<li>true: Modifications to both dynamic parameters listed in <code>dynamicParameters</code> and static parameters
listed in <code>staticParameters</code> will trigger a dynamic reload.
The &ldquo;all&rdquo; option is for certain engines that require static parameters to be set
via SQL statements before they can take effect on restart.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>downwardAPIChangeTriggeredActions</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.DownwardAPIChangeTriggeredAction">
[]DownwardAPIChangeTriggeredAction
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of actions to execute specified commands based on Pod labels.</p>
<p>It utilizes the K8s Downward API to mount label information as a volume into the pod.
The &lsquo;config-manager&rsquo; sidecar container watches for changes in the role label and dynamically invoke
registered commands (usually execute some SQL statements) when a change is detected.</p>
<p>It is designed for scenarios where:</p>
<ul>
<li>Replicas with different roles have different configurations, such as Redis primary &amp; secondary replicas.</li>
<li>After a role switch (e.g., from secondary to primary), some changes in configuration are needed
to reflect the new role.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>parametersSchema</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ParametersSchema">
ParametersSchema
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a list of parameters including their names, default values, descriptions,
types, and constraints (permissible values or the range of valid values).</p>
</td>
</tr>
<tr>
<td>
<code>staticParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List static parameters.
Modifications to any of these parameters require a restart of the process to take effect.</p>
</td>
</tr>
<tr>
<td>
<code>dynamicParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>List dynamic parameters.
Modifications to these parameters trigger a configuration reload without requiring a process restart.</p>
</td>
</tr>
<tr>
<td>
<code>immutableParameters</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Lists the parameters that cannot be modified once set.
Attempting to change any of these parameters will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>fileFormatConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.FileFormatConfig">
FileFormatConfig
</a>
</em>
</td>
<td>
<p>Specifies the format of the configuration file and any associated parameters that are specific to the chosen format.
Supported formats include <code>ini</code>, <code>xml</code>, <code>yaml</code>, <code>json</code>, <code>hcl</code>, <code>dotenv</code>, <code>properties</code>, and <code>toml</code>.</p>
<p>Each format may have its own set of parameters that can be configured.
For instance, when using the <code>ini</code> format, you can specify the section name.</p>
<p>Example:</p>
<pre><code>fileFormatConfig:
 format: ini
 iniConfig:
   sectionName: mysqld
</code></pre>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ConfigConstraintStatus">ConfigConstraintStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.ConfigConstraint">ConfigConstraint</a>)
</p>
<div>
<p>ConfigConstraintStatus represents the observed state of a ConfigConstraint.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintPhase">
ConfigConstraintPhase
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the status of the configuration template.
When set to CCAvailablePhase, the ConfigConstraint can be referenced by ClusterDefinition.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides descriptions for abnormal states.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the most recent generation observed for this ConfigConstraint. This value is updated by the API Server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.DownwardAPIChangeTriggeredAction">DownwardAPIChangeTriggeredAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">ConfigConstraintSpec</a>, <a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintSpec">ConfigConstraintSpec</a>)
</p>
<div>
<p>DownwardAPIChangeTriggeredAction defines an action that triggers specific commands in response to changes in Pod labels.
For example, a command might be executed when the &lsquo;role&rsquo; label of the Pod is updated.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the field. It must be a string of maximum length 63.
The name should match the regex pattern <code>^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$</code>.</p>
</td>
</tr>
<tr>
<td>
<code>mountPoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the mount point of the Downward API volume.</p>
</td>
</tr>
<tr>
<td>
<code>items</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#downwardapivolumefile-v1-core">
[]Kubernetes core/v1.DownwardAPIVolumeFile
</a>
</em>
</td>
<td>
<p>Represents a list of files under the Downward API volume.</p>
</td>
</tr>
<tr>
<td>
<code>command</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the command to be triggered when changes are detected in Downward API volume files.
It relies on the inotify mechanism in the config-manager sidecar to monitor file changes.</p>
</td>
</tr>
<tr>
<td>
<code>scriptConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ScriptConfig">
ScriptConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ScriptConfig object specifies a ConfigMap that contains script files that should be mounted inside the pod.
The scripts are mounted as volumes and can be referenced and executed by the DownwardAction to perform specific tasks or configurations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.DynamicParameterSelectedPolicy">DynamicParameterSelectedPolicy
(<code>string</code> alias)</h3>
<div>
<p>DynamicParameterSelectedPolicy determines how to select the parameters of dynamic reload actions</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;all&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;dynamic&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.DynamicReloadType">DynamicReloadType
(<code>string</code> alias)</h3>
<div>
<p>DynamicReloadType defines reload method.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;auto&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;sql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;exec&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;tpl&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;signal&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.FileFormatConfig">FileFormatConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">ConfigConstraintSpec</a>, <a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintSpec">ConfigConstraintSpec</a>)
</p>
<div>
<p>FileFormatConfig specifies the format of the configuration file and any associated parameters
that are specific to the chosen format.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>FormatterAction</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.FormatterAction">
FormatterAction
</a>
</em>
</td>
<td>
<p>
(Members of <code>FormatterAction</code> are embedded into this type.)
</p>
<em>(Optional)</em>
<p>Each format may have its own set of parameters that can be configured.
For instance, when using the <code>ini</code> format, you can specify the section name.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.CfgFileFormat">
CfgFileFormat
</a>
</em>
</td>
<td>
<p>The config file format. Valid values are <code>ini</code>, <code>xml</code>, <code>yaml</code>, <code>json</code>,
<code>hcl</code>, <code>dotenv</code>, <code>properties</code> and <code>toml</code>. Each format has its own characteristics and use cases.</p>
<ul>
<li>ini: is a text-based content with a structure and syntax comprising key–value pairs for properties, reference wiki: <a href="https://en.wikipedia.org/wiki/INI_file">https://en.wikipedia.org/wiki/INI_file</a></li>
<li>xml: refers to wiki: <a href="https://en.wikipedia.org/wiki/XML">https://en.wikipedia.org/wiki/XML</a></li>
<li>yaml: supports for complex data types and structures.</li>
<li>json: refers to wiki: <a href="https://en.wikipedia.org/wiki/JSON">https://en.wikipedia.org/wiki/JSON</a></li>
<li>hcl: The HashiCorp Configuration Language (HCL) is a configuration language authored by HashiCorp, reference url: <a href="https://www.linode.com/docs/guides/introduction-to-hcl/">https://www.linode.com/docs/guides/introduction-to-hcl/</a></li>
<li>dotenv: is a plain text file with simple key–value pairs, reference wiki: <a href="https://en.wikipedia.org/wiki/Configuration_file#MS-DOS">https://en.wikipedia.org/wiki/Configuration_file#MS-DOS</a></li>
<li>properties: a file extension mainly used in Java, reference wiki: <a href="https://en.wikipedia.org/wiki/.properties">https://en.wikipedia.org/wiki/.properties</a></li>
<li>toml: refers to wiki: <a href="https://en.wikipedia.org/wiki/TOML">https://en.wikipedia.org/wiki/TOML</a></li>
<li>props-plus: a file extension mainly used in Java, supports CamelCase(e.g: brokerMaxConnectionsPerIp)</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.FormatterAction">FormatterAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.FileFormatConfig">FileFormatConfig</a>)
</p>
<div>
<p>FormatterAction configures format-specific options for different configuration file format.
Note: Only one of its members should be specified at any given time.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>iniConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.IniConfig">
IniConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Holds options specific to the &lsquo;ini&rsquo; file format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.IniConfig">IniConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.FormatterAction">FormatterAction</a>)
</p>
<div>
<p>IniConfig holds options specific to the &lsquo;ini&rsquo; file format.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sectionName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A string that describes the name of the ini section.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ParametersSchema">ParametersSchema
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintSpec">ConfigConstraintSpec</a>)
</p>
<div>
<p>ParametersSchema Defines a list of configuration items with their names, default values, descriptions,
types, and constraints.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>topLevelKey</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the top-level key in the &lsquo;configSchema.cue&rsquo; that organizes the validation rules for parameters.
This key must exist within the CUE script defined in &lsquo;configSchema.cue&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>cue</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Hold a string that contains a script written in CUE language that defines a list of configuration items.
Each item is detailed with its name, default value, description, type (e.g. string, integer, float),
and constraints (permissible values or the valid range of values).</p>
<p>CUE (Configure, Unify, Execute) is a declarative language designed for defining and validating
complex data configurations.
It is particularly useful in environments like K8s where complex configurations and validation rules are common.</p>
<p>This script functions as a validator for user-provided configurations, ensuring compliance with
the established specifications and constraints.</p>
</td>
</tr>
<tr>
<td>
<code>schemaInJSON</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#jsonschemaprops-v1-apiextensions-k8s-io">
Kubernetes api extensions v1.JSONSchemaProps
</a>
</em>
</td>
<td>
<p>Generated from the &lsquo;cue&rsquo; field and transformed into a JSON format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ReloadAction">ReloadAction
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.ConfigConstraintSpec">ConfigConstraintSpec</a>)
</p>
<div>
<p>ReloadAction defines the mechanisms available for dynamically reloading a process within K8s without requiring a restart.</p>
<p>Only one of the mechanisms can be specified at a time.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>unixSignalTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.UnixSignalTrigger">
UnixSignalTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to trigger a reload by sending a specific Unix signal to the process.</p>
</td>
</tr>
<tr>
<td>
<code>shellTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ShellTrigger">
ShellTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows to execute a custom shell script to reload the process.</p>
</td>
</tr>
<tr>
<td>
<code>tplScriptTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.TPLScriptTrigger">
TPLScriptTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Enables reloading process using a Go template script.</p>
</td>
</tr>
<tr>
<td>
<code>autoTrigger</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.AutoTrigger">
AutoTrigger
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Automatically perform the reload when specified conditions are met.</p>
</td>
</tr>
<tr>
<td>
<code>targetPodSelector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Used to match labels on the pod to determine whether a dynamic reload should be performed.</p>
<p>In some scenarios, only specific pods (e.g., primary replicas) need to undergo a dynamic reload.
The <code>reloadedPodSelector</code> allows you to specify label selectors to target the desired pods for the reload process.</p>
<p>If the <code>reloadedPodSelector</code> is not specified or is nil, all pods managed by the workload will be considered for the dynamic
reload.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ScriptConfig">ScriptConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">ConfigConstraintSpec</a>, <a href="#apps.kubeblocks.io/v1beta1.DownwardAPIChangeTriggeredAction">DownwardAPIChangeTriggeredAction</a>, <a href="#apps.kubeblocks.io/v1beta1.ShellTrigger">ShellTrigger</a>, <a href="#apps.kubeblocks.io/v1beta1.TPLScriptTrigger">TPLScriptTrigger</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>scriptConfigMapRef</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the reference to the ConfigMap containing the scripts.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the namespace for the ConfigMap.
If not specified, it defaults to the &ldquo;default&rdquo; namespace.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ShellTrigger">ShellTrigger
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ReloadOptions">ReloadOptions</a>, <a href="#apps.kubeblocks.io/v1beta1.ReloadAction">ReloadAction</a>)
</p>
<div>
<p>ShellTrigger allows to execute a custom shell script to reload the process.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>command</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Specifies the command to execute in order to reload the process. It should be a valid shell command.</p>
</td>
</tr>
<tr>
<td>
<code>sync</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines the synchronization mode of parameter updates with &ldquo;config-manager&rdquo;.</p>
<ul>
<li>&lsquo;True&rsquo;: Executes reload actions synchronously, pausing until completion.</li>
<li>&lsquo;False&rsquo;: Executes reload actions asynchronously, without waiting for completion.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>batchReload</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls whether parameter updates are processed individually or collectively in a batch:</p>
<ul>
<li>&lsquo;True&rsquo;: Processes all changes in one batch reload.</li>
<li>&lsquo;False&rsquo;: Processes each change individually.</li>
</ul>
<p>Defaults to &lsquo;False&rsquo; if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>batchParamsFormatterTemplate</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a Go template string for formatting batch input data.
It&rsquo;s used when <code>batchReload</code> is &lsquo;True&rsquo; to format data passed into STDIN of the script.
The template accesses key-value pairs of updated parameters via the &lsquo;$&rsquo; variable.
This allows for custom formatting of the input data.</p>
<p>Example template:</p>
<pre><code class="language-yaml">batchParamsFormatterTemplate: |-
&#123;&#123;- range $pKey, $pValue := $ &#125;&#125;
&#123;&#123; printf &quot;%s:%s&quot; $pKey $pValue &#125;&#125;
&#123;&#123;- end &#125;&#125;
</code></pre>
<p>This example generates batch input data in a key:value format, sorted by keys.</p>
<pre><code>key1:value1
key2:value2
key3:value3
</code></pre>
<p>If not specified, the default format is key=value, sorted by keys, for each updated parameter.</p>
<pre><code>key1=value1
key2=value2
key3=value3
</code></pre>
</td>
</tr>
<tr>
<td>
<code>toolsSetup</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ToolsSetup">
ToolsSetup
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the tools container image used by ShellTrigger for dynamic reload.
If the dynamic reload action is triggered by a ShellTrigger, this field is required.
This image must contain all necessary tools for executing the ShellTrigger scripts.</p>
<p>Usually the specified image is referenced by the init container,
which is then responsible for copy the tools from the image to a bin volume.
This ensures that the tools are available to the &lsquo;config-manager&rsquo; sidecar.</p>
</td>
</tr>
<tr>
<td>
<code>scriptConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ScriptConfig">
ScriptConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ScriptConfig object specifies a ConfigMap that contains script files that should be mounted inside the pod.
The scripts are mounted as volumes and can be referenced and executed by the dynamic reload.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.SignalType">SignalType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.UnixSignalTrigger">UnixSignalTrigger</a>)
</p>
<div>
<p>SignalType defines which signals are valid.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;SIGABRT&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGALRM&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGBUS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGCHLD&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGCONT&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGFPE&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGHUP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGILL&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGINT&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGIO&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGKILL&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGPIPE&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGPROF&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGPWR&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGQUIT&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGSEGV&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGSTKFLT&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGSTOP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGSYS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGTERM&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGTRAP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGTSTP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGTTIN&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGTTOU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGURG&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGUSR1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGUSR2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGVTALRM&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGWINCH&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGXCPU&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SIGXFSZ&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.TPLScriptTrigger">TPLScriptTrigger
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ReloadOptions">ReloadOptions</a>, <a href="#apps.kubeblocks.io/v1beta1.ReloadAction">ReloadAction</a>)
</p>
<div>
<p>TPLScriptTrigger Enables reloading process using a Go template script.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ScriptConfig</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ScriptConfig">
ScriptConfig
</a>
</em>
</td>
<td>
<p>
(Members of <code>ScriptConfig</code> are embedded into this type.)
</p>
<p>Specifies the ConfigMap that contains the script to be executed for reload.</p>
</td>
</tr>
<tr>
<td>
<code>sync</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines whether parameter updates should be synchronized with the &ldquo;config-manager&rdquo;.
Specifies the controller&rsquo;s reload strategy:</p>
<ul>
<li>If set to &lsquo;True&rsquo;, the controller executes the reload action in synchronous mode,
pausing execution until the reload completes.</li>
<li>If set to &lsquo;False&rsquo;, the controller executes the reload action in asynchronous mode,
updating the ConfigMap without waiting for the reload process to finish.</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ToolConfig">ToolConfig
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1beta1.ToolsSetup">ToolsSetup</a>)
</p>
<div>
<p>ToolConfig specifies the settings of an init container that prepare tools for dynamic reload.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the name of the init container.</p>
</td>
</tr>
<tr>
<td>
<code>asContainerImage</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether the tool image should be used as the container image for a sidecar.
This is useful for large tool images, such as those for C++ tools, which may depend on
numerous libraries (e.g., *.so files).</p>
<p>If enabled, the tool image is deployed as a sidecar container image.</p>
<p>Examples:</p>
<pre><code class="language-yaml"> toolsSetup::
   mountPoint: /kb_tools
   toolConfigs:
     - name: kb-tools
       asContainerImage: true
       image:  apecloud/oceanbase:4.2.0.0-100010032023083021
</code></pre>
<p>generated containers:</p>
<pre><code class="language-yaml">initContainers:
 - name: install-config-manager-tool
   image: apecloud/kubeblocks-tools:$&#123;version&#125;
   command:
   - cp
   - /bin/config_render
   - /opt/tools
   volumemounts:
   - name: kb-tools
     mountpath: /opt/tools
containers:
 - name: config-manager
   image: apecloud/oceanbase:4.2.0.0-100010032023083021
   imagePullPolicy: IfNotPresent
	  command:
   - /opt/tools/reloader
   - --log-level
   - info
   - --operator-update-enable
   - --tcp
   - &quot;9901&quot;
   - --config
   - /opt/config-manager/config-manager.yaml
   volumemounts:
   - name: kb-tools
     mountpath: /opt/tools
</code></pre>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the tool container image.</p>
</td>
</tr>
<tr>
<td>
<code>command</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the command to be executed by the init container.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.ToolsSetup">ToolsSetup
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ConfigConstraintSpec">ConfigConstraintSpec</a>, <a href="#apps.kubeblocks.io/v1beta1.ShellTrigger">ShellTrigger</a>)
</p>
<div>
<p>ToolsSetup prepares the tools for dynamic reloads used in ShellTrigger from a specified container image.</p>
<p>Example:</p>
<pre><code class="language-yaml">
toolsSetup:
	 mountPoint: /kb_tools
	 toolConfigs:
	   - name: kb-tools
	     command:
	       - cp
	       - /bin/ob-tools
	       - /kb_tools/obtools
	     image: docker.io/apecloud/obtools
</code></pre>
<p>This example copies the &ldquo;/bin/ob-tools&rdquo; binary from the image to &ldquo;/kb_tools/obtools&rdquo;.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>mountPoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Specifies the directory path in the container where the tools-related files are to be copied.
This field is typically used with an emptyDir volume to ensure a temporary, empty directory is provided at pod creation.</p>
</td>
</tr>
<tr>
<td>
<code>toolConfigs</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.ToolConfig">
[]ToolConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of settings of init containers that prepare tools for dynamic reload.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apps.kubeblocks.io/v1beta1.UnixSignalTrigger">UnixSignalTrigger
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ReloadOptions">ReloadOptions</a>, <a href="#apps.kubeblocks.io/v1beta1.ReloadAction">ReloadAction</a>)
</p>
<div>
<p>UnixSignalTrigger is used to trigger a reload by sending a specific Unix signal to the process.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>signal</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1beta1.SignalType">
SignalType
</a>
</em>
</td>
<td>
<p>Specifies a valid Unix signal to be sent.
For a comprehensive list of all Unix signals, see: ../../pkg/configuration/configmap/handler.go:allUnixSignals</p>
</td>
</tr>
<tr>
<td>
<code>processName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Identifies the name of the process to which the Unix signal will be sent.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="workloads.kubeblocks.io/v1">workloads.kubeblocks.io/v1</h2>
<div>
</div>
Resource Types:
<ul><li>
<a href="#workloads.kubeblocks.io/v1.InstanceSet">InstanceSet</a>
</li></ul>
<h3 id="workloads.kubeblocks.io/v1.InstanceSet">InstanceSet
</h3>
<div>
<p>InstanceSet is the Schema for the instancesets API.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>workloads.kubeblocks.io/v1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>InstanceSet</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<p>Contains the metadata for the particular object, such as name, namespace, labels, and annotations.</p>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">
InstanceSetSpec
</a>
</em>
</td>
<td>
<p>Defines the desired state of the state machine. It includes the configuration details for the state machine.</p>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the desired number of replicas of the given Template.
These replicas are instantiations of the same Template, with each having a consistent identity.
Defaults to 1 if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTemplateOrdinals</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Ordinals">
Ordinals
</a>
</em>
</td>
<td>
<p>Specifies the desired Ordinals of the default template.
The Ordinals used to specify the ordinal of the instance (pod) names to be generated under the default template.
If Ordinals are defined, their number must be equal to or more than the corresponding replicas.</p>
<p>For example, if Ordinals is &#123;ranges: [&#123;start: 0, end: 1&#125;], discrete: [7]&#125;,
then the instance names generated under the default template would be
$(cluster.name)-$(component.name)-0、$(cluster.name)-$(component.name)-1 and $(cluster.name)-$(component.name)-7</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the minimum number of seconds a newly created pod should be ready
without any of its container crashing to be considered available.
Defaults to 0, meaning the pod will be considered available as soon as it is ready.</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<p>Represents a label query over pods that should match the desired replica count indicated by the <code>replica</code> field.
It must match the labels defined in the pod template.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors">https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors</a></p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides values in default Template.</p>
<p>Instance is the fundamental unit managed by KubeBlocks.
It represents a Pod with additional objects such as PVCs, Services, ConfigMaps, etc.
An InstanceSet manages instances with a total count of Replicas,
and by default, all these instances are generated from the same template.
The InstanceTemplate provides a way to override values in the default template,
allowing the InstanceSet to manage instances from different templates.</p>
<p>By default, the ordinal starts from 0 for each InstanceTemplate.
It is important to ensure that the Name of each InstanceTemplate is unique.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the InstanceSet.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>flatInstanceOrdinal</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>flatInstanceOrdinal controls whether the naming of instances(pods) under this component uses a flattened,
globally uniquely ordinal scheme, regardless of the instance template.</p>
<p>Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The cluster administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeclaim-v1-core">
[]Kubernetes core/v1.PersistentVolumeClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for each replica.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for replicas upon their creation.
The final name of each PVC is generated by appending the pod&rsquo;s identifier to the name specified in volumeClaimTemplates[*].name.</p>
</td>
</tr>
<tr>
<td>
<code>persistentVolumeClaimRetentionPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicy">
PersistentVolumeClaimRetentionPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>persistentVolumeClaimRetentionPolicy describes the lifecycle of persistent
volume claims created from volumeClaimTemplates. By default, all persistent
volume claims are created as needed and retained until manually deleted. This
policy allows the lifecycle to be altered, for example by deleting persistent
volume claims when their workload is deleted, or when their pod is scaled
down.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls how pods are created during initial scale up,
when replacing pods on nodes, or when scaling down.</p>
<p>The default policy is <code>OrderedReady</code>, where pods are created in increasing order and the controller waits until each pod is ready before
continuing. When scaling down, the pods are removed in the opposite order.
The alternative policy is <code>Parallel</code> which will create pods in parallel
to match the desired scale without waiting, and on scale down will delete
all pods at once.</p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides fine-grained control over the spec update process of all instances.</p>
</td>
</tr>
<tr>
<td>
<code>memberUpdateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.MemberUpdateStrategy">
MemberUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Members(Pods) update strategy.</p>
<ul>
<li>serial: update Members one by one that guarantee minimum component unavailable time.</li>
<li>parallel: force parallel</li>
<li>bestEffortParallel: update Members in parallel that guarantee minimum component un-writable time.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A list of roles defined in the system. Instanceset obtains role through pods&rsquo; role label <code>kubeblocks.io/role</code>.</p>
</td>
</tr>
<tr>
<td>
<code>membershipReconfiguration</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.MembershipReconfiguration">
MembershipReconfiguration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides actions to do membership dynamic reconfiguration.</p>
</td>
</tr>
<tr>
<td>
<code>templateVars</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides variables which are used to call Actions.</p>
</td>
</tr>
<tr>
<td>
<code>paused</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates that the InstanceSet is paused, meaning the reconciliation of this InstanceSet object will be paused.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.ConfigTemplate">
[]ConfigTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Describe the configs to be reconfigured.</p>
</td>
</tr>
<tr>
<td>
<code>disableDefaultHeadlessService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to create the default headless service.</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.InstanceSetStatus">
InstanceSetStatus
</a>
</em>
</td>
<td>
<p>Represents the current information about the state machine. This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.ConditionType">ConditionType
(<code>string</code> alias)</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;InstanceAvailable&#34;</p></td>
<td><p>InstanceAvailable ConditionStatus will be True if all instances(pods) are in the ready condition
and continue for &ldquo;MinReadySeconds&rdquo; seconds. Otherwise, it will be set to False.</p>
</td>
</tr><tr><td><p>&#34;InstanceFailure&#34;</p></td>
<td><p>InstanceFailure is added in an instance set when at least one of its instances(pods) is in a <code>Failed</code> phase.</p>
</td>
</tr><tr><td><p>&#34;InstanceReady&#34;</p></td>
<td><p>InstanceReady is added in an instance set when at least one of its instances(pods) is in a Ready condition.
ConditionStatus will be True if all its instances(pods) are in a Ready condition.
Or, a NotReady reason with not ready instances encoded in the Message filed will be set.</p>
</td>
</tr><tr><td><p>&#34;InstanceUpdateRestricted&#34;</p></td>
<td><p>InstanceUpdateRestricted represents a ConditionType that indicates updates to an InstanceSet are blocked(when the
PodUpdatePolicy is set to StrictInPlace but the pods cannot be updated in-place).</p>
</td>
</tr></tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.ConfigTemplate">ConfigTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the config.</p>
</td>
</tr>
<tr>
<td>
<code>generation</code><br/>
<em>
int64
</em>
</td>
<td>
<p>The generation of the config.</p>
</td>
</tr>
<tr>
<td>
<code>reconfigure</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The custom reconfigure action.</p>
</td>
</tr>
<tr>
<td>
<code>reconfigureActionName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The name of the custom reconfigure action.</p>
<p>An empty name indicates that the reconfigure action is the default one defined by lifecycle actions.</p>
</td>
</tr>
<tr>
<td>
<code>parameters</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The parameters to call the reconfigure action.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.InstanceConfigStatus">InstanceConfigStatus
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceStatus">InstanceStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>The name of the config.</p>
</td>
</tr>
<tr>
<td>
<code>generation</code><br/>
<em>
int64
</em>
</td>
<td>
<p>The generation of the config.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSet">InstanceSet</a>)
</p>
<div>
<p>InstanceSetSpec defines the desired state of InstanceSet</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the desired number of replicas of the given Template.
These replicas are instantiations of the same Template, with each having a consistent identity.
Defaults to 1 if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTemplateOrdinals</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Ordinals">
Ordinals
</a>
</em>
</td>
<td>
<p>Specifies the desired Ordinals of the default template.
The Ordinals used to specify the ordinal of the instance (pod) names to be generated under the default template.
If Ordinals are defined, their number must be equal to or more than the corresponding replicas.</p>
<p>For example, if Ordinals is &#123;ranges: [&#123;start: 0, end: 1&#125;], discrete: [7]&#125;,
then the instance names generated under the default template would be
$(cluster.name)-$(component.name)-0、$(cluster.name)-$(component.name)-1 and $(cluster.name)-$(component.name)-7</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the minimum number of seconds a newly created pod should be ready
without any of its container crashing to be considered available.
Defaults to 0, meaning the pod will be considered available as soon as it is ready.</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<p>Represents a label query over pods that should match the desired replica count indicated by the <code>replica</code> field.
It must match the labels defined in the pod template.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors">https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors</a></p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides values in default Template.</p>
<p>Instance is the fundamental unit managed by KubeBlocks.
It represents a Pod with additional objects such as PVCs, Services, ConfigMaps, etc.
An InstanceSet manages instances with a total count of Replicas,
and by default, all these instances are generated from the same template.
The InstanceTemplate provides a way to override values in the default template,
allowing the InstanceSet to manage instances from different templates.</p>
<p>By default, the ordinal starts from 0 for each InstanceTemplate.
It is important to ensure that the Name of each InstanceTemplate is unique.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the InstanceSet.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>flatInstanceOrdinal</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>flatInstanceOrdinal controls whether the naming of instances(pods) under this component uses a flattened,
globally uniquely ordinal scheme, regardless of the instance template.</p>
<p>Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The cluster administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeclaim-v1-core">
[]Kubernetes core/v1.PersistentVolumeClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for each replica.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for replicas upon their creation.
The final name of each PVC is generated by appending the pod&rsquo;s identifier to the name specified in volumeClaimTemplates[*].name.</p>
</td>
</tr>
<tr>
<td>
<code>persistentVolumeClaimRetentionPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PersistentVolumeClaimRetentionPolicy">
PersistentVolumeClaimRetentionPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>persistentVolumeClaimRetentionPolicy describes the lifecycle of persistent
volume claims created from volumeClaimTemplates. By default, all persistent
volume claims are created as needed and retained until manually deleted. This
policy allows the lifecycle to be altered, for example by deleting persistent
volume claims when their workload is deleted, or when their pod is scaled
down.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls how pods are created during initial scale up,
when replacing pods on nodes, or when scaling down.</p>
<p>The default policy is <code>OrderedReady</code>, where pods are created in increasing order and the controller waits until each pod is ready before
continuing. When scaling down, the pods are removed in the opposite order.
The alternative policy is <code>Parallel</code> which will create pods in parallel
to match the desired scale without waiting, and on scale down will delete
all pods at once.</p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>instanceUpdateStrategy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides fine-grained control over the spec update process of all instances.</p>
</td>
</tr>
<tr>
<td>
<code>memberUpdateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.MemberUpdateStrategy">
MemberUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Members(Pods) update strategy.</p>
<ul>
<li>serial: update Members one by one that guarantee minimum component unavailable time.</li>
<li>parallel: force parallel</li>
<li>bestEffortParallel: update Members in parallel that guarantee minimum component un-writable time.</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A list of roles defined in the system. Instanceset obtains role through pods&rsquo; role label <code>kubeblocks.io/role</code>.</p>
</td>
</tr>
<tr>
<td>
<code>membershipReconfiguration</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.MembershipReconfiguration">
MembershipReconfiguration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides actions to do membership dynamic reconfiguration.</p>
</td>
</tr>
<tr>
<td>
<code>templateVars</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides variables which are used to call Actions.</p>
</td>
</tr>
<tr>
<td>
<code>paused</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates that the InstanceSet is paused, meaning the reconciliation of this InstanceSet object will be paused.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.ConfigTemplate">
[]ConfigTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Describe the configs to be reconfigured.</p>
</td>
</tr>
<tr>
<td>
<code>disableDefaultHeadlessService</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies whether to create the default headless service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.InstanceSetStatus">InstanceSetStatus
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSet">InstanceSet</a>)
</p>
<div>
<p>InstanceSetStatus defines the observed state of InstanceSet</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>observedGeneration is the most recent generation observed for this InstanceSet. It corresponds to the
InstanceSet&rsquo;s generation, which is updated on mutation by the API Server.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>replicas is the number of instances created by the InstanceSet controller.</p>
</td>
</tr>
<tr>
<td>
<code>ordinals</code><br/>
<em>
[]int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Ordinals is the ordinals used by the instances of the InstanceSet except the template instances.</p>
</td>
</tr>
<tr>
<td>
<code>readyReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>readyReplicas is the number of instances created for this InstanceSet with a Ready Condition.</p>
</td>
</tr>
<tr>
<td>
<code>currentReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>currentReplicas is the number of instances created by the InstanceSet controller from the InstanceSet version
indicated by CurrentRevisions.</p>
</td>
</tr>
<tr>
<td>
<code>updatedReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>updatedReplicas is the number of instances created by the InstanceSet controller from the InstanceSet version
indicated by UpdateRevisions.</p>
</td>
</tr>
<tr>
<td>
<code>currentRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>currentRevision, if not empty, indicates the version of the InstanceSet used to generate instances in the
sequence [0,currentReplicas).</p>
</td>
</tr>
<tr>
<td>
<code>updateRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>updateRevision, if not empty, indicates the version of the InstanceSet used to generate instances in the sequence
[replicas-updatedReplicas,replicas)</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the latest available observations of an instanceset&rsquo;s current state.
Known .status.conditions.type are: &ldquo;InstanceFailure&rdquo;, &ldquo;InstanceReady&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>availableReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Total number of available instances (ready for at least minReadySeconds) targeted by this InstanceSet.</p>
</td>
</tr>
<tr>
<td>
<code>initReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the initial number of instances when the cluster is first initialized.
This value is set to spec.Replicas at the time of object creation and remains constant thereafter.
Used only when spec.roles set.</p>
</td>
</tr>
<tr>
<td>
<code>readyInitReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the number of instances that have already reached the MembersStatus during the cluster initialization stage.
This value remains constant once it equals InitReplicas.
Used only when spec.roles set.</p>
</td>
</tr>
<tr>
<td>
<code>membersStatus</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.MemberStatus">
[]MemberStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides the status of each member in the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>instanceStatus</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.InstanceStatus">
[]InstanceStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides the status of each instance in the ITS.</p>
</td>
</tr>
<tr>
<td>
<code>currentRevisions</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>currentRevisions, if not empty, indicates the old version of the InstanceSet used to generate the underlying workload.
key is the pod name, value is the revision.</p>
</td>
</tr>
<tr>
<td>
<code>updateRevisions</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>updateRevisions, if not empty, indicates the new version of the InstanceSet used to generate the underlying workload.
key is the pod name, value is the revision.</p>
</td>
</tr>
<tr>
<td>
<code>templatesStatus</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.InstanceTemplateStatus">
[]InstanceTemplateStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TemplatesStatus represents status of each instance generated by InstanceTemplates</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.InstanceStatus">InstanceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSetStatus">InstanceSetStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>podName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the name of the pod.</p>
</td>
</tr>
<tr>
<td>
<code>configs</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1.InstanceConfigStatus">
[]InstanceConfigStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The status of configs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.InstanceTemplate">InstanceTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>InstanceTemplate allows customization of individual replica configurations in a Component.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name specifies the unique name of the instance Pod created using this InstanceTemplate.
This name is constructed by concatenating the Component&rsquo;s name, the template&rsquo;s name, and the instance&rsquo;s ordinal
using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0.
The specified name overrides any default naming conventions or patterns.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of instances (Pods) to create from this InstanceTemplate.
This field allows setting how many replicated instances of the Component,
with the specific overrides in the InstanceTemplate, are created.
The default value is 1. A value of 0 disables instance creation.</p>
</td>
</tr>
<tr>
<td>
<code>ordinals</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Ordinals">
Ordinals
</a>
</em>
</td>
<td>
<p>Specifies the desired Ordinals of this InstanceTemplate.
The Ordinals used to specify the ordinal of the instance (pod) names to be generated under this InstanceTemplate.</p>
<p>For example, if Ordinals is &#123;ranges: [&#123;start: 0, end: 1&#125;], discrete: [7]&#125;,
then the instance names generated under this InstanceTemplate would be
$(cluster.name)-$(component.name)-$(template.name)-0、$(cluster.name)-$(component.name)-$(template.name)-1 and
$(cluster.name)-$(component.name)-$(template.name)-7</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs to be merged into the Pod&rsquo;s existing annotations.
Existing keys will have their values overwritten, while new keys will be added to the annotations.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs that will be merged into the Pod&rsquo;s existing labels.
Values for existing keys will be overwritten, and new keys will be added.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the resource requirements of the first container in the Pod.
This field allows for customizing resource allocation (CPU, memory, etc.) for the container.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines Env to override.
Add new or override existing envs.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeclaim-v1-core">
[]Kubernetes core/v1.PersistentVolumeClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the storage requirements of the instances.</p>
</td>
</tr>
<tr>
<td>
<code>images</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Images for the containers of the instance template.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.InstanceTemplateStatus">InstanceTemplateStatus
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSetStatus">InstanceSetStatus</a>)
</p>
<div>
<p>InstanceTemplateStatus aggregates the status of replicas for each InstanceTemplate</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name, the name of the InstanceTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Replicas is the number of replicas of the InstanceTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>ordinals</code><br/>
<em>
[]int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Ordinals is the ordinals used by the instances of the InstanceTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>readyReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReadyReplicas is the number of Pods that have a Ready Condition.</p>
</td>
</tr>
<tr>
<td>
<code>availableReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>AvailableReplicas is the number of Pods that ready for at least minReadySeconds.</p>
</td>
</tr>
<tr>
<td>
<code>currentReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>currentReplicas is the number of instances created by the InstanceSet controller from the InstanceSet version
indicated by CurrentRevisions.</p>
</td>
</tr>
<tr>
<td>
<code>updatedReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>UpdatedReplicas is the number of Pods created by the InstanceSet controller from the InstanceSet version
indicated by UpdateRevisions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.MemberStatus">MemberStatus
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSetStatus">InstanceSetStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>podName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the name of the pod.</p>
</td>
</tr>
<tr>
<td>
<code>role</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.ReplicaRole">
ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the role of the replica in the cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.MemberUpdateStrategy">MemberUpdateStrategy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>MemberUpdateStrategy defines Cluster Component update strategy.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;BestEffortParallel&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Parallel&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Serial&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1.MembershipReconfiguration">MembershipReconfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>switchover</code><br/>
<em>
<a href="#apps.kubeblocks.io/v1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the procedure for a controlled transition of a role to a new replica.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<h2 id="workloads.kubeblocks.io/v1alpha1">workloads.kubeblocks.io/v1alpha1</h2>
<div>
</div>
Resource Types:
<ul><li>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceSet">InstanceSet</a>
</li></ul>
<h3 id="workloads.kubeblocks.io/v1alpha1.InstanceSet">InstanceSet
</h3>
<div>
<p>InstanceSet is the Schema for the instancesets API.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>workloads.kubeblocks.io/v1alpha1</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>InstanceSet</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
<p>Contains the metadata for the particular object, such as name, namespace, labels, and annotations.</p>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">
InstanceSetSpec
</a>
</em>
</td>
<td>
<p>Defines the desired state of the state machine. It includes the configuration details for the state machine.</p>
<br/>
<br/>
<table>
<tbody>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the desired number of replicas of the given Template.
These replicas are instantiations of the same Template, with each having a consistent identity.
Defaults to 1 if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTemplateOrdinals</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Ordinals">
Ordinals
</a>
</em>
</td>
<td>
<p>Specifies the desired Ordinals of the default template.
The Ordinals used to specify the ordinal of the instance (pod) names to be generated under the default template.</p>
<p>For example, if Ordinals is &#123;ranges: [&#123;start: 0, end: 1&#125;], discrete: [7]&#125;,
then the instance names generated under the default template would be
$(cluster.name)-$(component.name)-0、$(cluster.name)-$(component.name)-1 and $(cluster.name)-$(component.name)-7</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the minimum number of seconds a newly created pod should be ready
without any of its container crashing to be considered available.
Defaults to 0, meaning the pod will be considered available as soon as it is ready.</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<p>Represents a label query over pods that should match the desired replica count indicated by the <code>replica</code> field.
It must match the labels defined in the pod template.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors">https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors</a></p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#service-v1-core">
Kubernetes core/v1.Service
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the behavior of a service spec.
Provides read-write service.
<a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</a></p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides values in default Template.</p>
<p>Instance is the fundamental unit managed by KubeBlocks.
It represents a Pod with additional objects such as PVCs, Services, ConfigMaps, etc.
An InstanceSet manages instances with a total count of Replicas,
and by default, all these instances are generated from the same template.
The InstanceTemplate provides a way to override values in the default template,
allowing the InstanceSet to manage instances from different templates.</p>
<p>The naming convention for instances (pods) based on the InstanceSet Name, InstanceTemplate Name, and ordinal.
The constructed instance name follows the pattern: $(instance_set.name)-$(template.name)-$(ordinal).
By default, the ordinal starts from 0 for each InstanceTemplate.
It is important to ensure that the Name of each InstanceTemplate is unique.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the InstanceSet.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The cluster administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeclaim-v1-core">
[]Kubernetes core/v1.PersistentVolumeClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for each replica.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for replicas upon their creation.
The final name of each PVC is generated by appending the pod&rsquo;s identifier to the name specified in volumeClaimTemplates[*].name.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls how pods are created during initial scale up,
when replacing pods on nodes, or when scaling down.</p>
<p>The default policy is <code>OrderedReady</code>, where pods are created in increasing order and the controller waits until each pod is ready before
continuing. When scaling down, the pods are removed in the opposite order.
The alternative policy is <code>Parallel</code> which will create pods in parallel
to match the desired scale without waiting, and on scale down will delete
all pods at once.</p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<p>Indicates the StatefulSetUpdateStrategy that will be
employed to update Pods in the InstanceSet when a revision is made to
Template.</p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A list of roles defined in the system.</p>
</td>
</tr>
<tr>
<td>
<code>roleProbe</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.RoleProbe">
RoleProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides method to probe role.</p>
</td>
</tr>
<tr>
<td>
<code>membershipReconfiguration</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MembershipReconfiguration">
MembershipReconfiguration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides actions to do membership dynamic reconfiguration.</p>
</td>
</tr>
<tr>
<td>
<code>memberUpdateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MemberUpdateStrategy">
MemberUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Members(Pods) update strategy.</p>
<ul>
<li>serial: update Members one by one that guarantee minimum component unavailable time.</li>
<li>bestEffortParallel: update Members in parallel that guarantee minimum component un-writable time.</li>
<li>parallel: force parallel</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>paused</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates that the InstanceSet is paused, meaning the reconciliation of this InstanceSet object will be paused.</p>
</td>
</tr>
<tr>
<td>
<code>credential</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Credential">
Credential
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Credential used to connect to DB engine</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetStatus">
InstanceSetStatus
</a>
</em>
</td>
<td>
<p>Represents the current information about the state machine. This data may be out of date.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.AccessMode">AccessMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.ReplicaRole">ReplicaRole</a>)
</p>
<div>
<p>AccessMode defines SVC access mode enums.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadWrite&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Readonly&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.Action">Action
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.MembershipReconfiguration">MembershipReconfiguration</a>, <a href="#workloads.kubeblocks.io/v1alpha1.RoleProbe">RoleProbe</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Refers to the utility image that contains the command which can be utilized to retrieve or process role information.</p>
</td>
</tr>
<tr>
<td>
<code>command</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>A set of instructions that will be executed within the Container to retrieve or process role information. This field is required.</p>
</td>
</tr>
<tr>
<td>
<code>args</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional parameters used to perform specific statements. This field is optional.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.ConditionType">ConditionType
(<code>string</code> alias)</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;InstanceAvailable&#34;</p></td>
<td><p>InstanceAvailable ConditionStatus will be True if all instances(pods) are in the ready condition
and continue for &ldquo;MinReadySeconds&rdquo; seconds. Otherwise, it will be set to False.</p>
</td>
</tr><tr><td><p>&#34;InstanceFailure&#34;</p></td>
<td><p>InstanceFailure is added in an instance set when at least one of its instances(pods) is in a <code>Failed</code> phase.</p>
</td>
</tr><tr><td><p>&#34;InstanceReady&#34;</p></td>
<td><p>InstanceReady is added in an instance set when at least one of its instances(pods) is in a Ready condition.
ConditionStatus will be True if all its instances(pods) are in a Ready condition.
Or, a NotReady reason with not ready instances encoded in the Message filed will be set.</p>
</td>
</tr><tr><td><p>&#34;InstanceUpdateRestricted&#34;</p></td>
<td><p>InstanceUpdateRestricted represents a ConditionType that indicates updates to an InstanceSet are blocked(when the
PodUpdatePolicy is set to StrictInPlace but the pods cannot be updated in-place).</p>
</td>
</tr></tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.Credential">Credential
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>username</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<p>Defines the user&rsquo;s name for the credential.
The corresponding environment variable will be KB_ITS_USERNAME.</p>
</td>
</tr>
<tr>
<td>
<code>password</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.CredentialVar">
CredentialVar
</a>
</em>
</td>
<td>
<p>Represents the user&rsquo;s password for the credential.
The corresponding environment variable will be KB_ITS_PASSWORD.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.CredentialVar">CredentialVar
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.Credential">Credential</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the value of the environment variable. This field is optional and defaults to an empty string.
The value can include variable references in the format $(VAR_NAME) which will be expanded using previously defined environment variables in the container and any service environment variables.</p>
<p>If a variable cannot be resolved, the reference in the input string will remain unchanged.
Double $$ can be used to escape the $(VAR_NAME) syntax, resulting in a single $ and producing the string literal &ldquo;$(VAR_NAME)&rdquo;.
Escaped references will not be expanded, regardless of whether the variable exists or not.</p>
</td>
</tr>
<tr>
<td>
<code>valueFrom</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvarsource-v1-core">
Kubernetes core/v1.EnvVarSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the source for the environment variable&rsquo;s value. This field is optional and cannot be used if the &lsquo;Value&rsquo; field is not empty.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceSet">InstanceSet</a>)
</p>
<div>
<p>InstanceSetSpec defines the desired state of InstanceSet</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the desired number of replicas of the given Template.
These replicas are instantiations of the same Template, with each having a consistent identity.
Defaults to 1 if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>defaultTemplateOrdinals</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Ordinals">
Ordinals
</a>
</em>
</td>
<td>
<p>Specifies the desired Ordinals of the default template.
The Ordinals used to specify the ordinal of the instance (pod) names to be generated under the default template.</p>
<p>For example, if Ordinals is &#123;ranges: [&#123;start: 0, end: 1&#125;], discrete: [7]&#125;,
then the instance names generated under the default template would be
$(cluster.name)-$(component.name)-0、$(cluster.name)-$(component.name)-1 and $(cluster.name)-$(component.name)-7</p>
</td>
</tr>
<tr>
<td>
<code>minReadySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the minimum number of seconds a newly created pod should be ready
without any of its container crashing to be considered available.
Defaults to 0, meaning the pod will be considered available as soon as it is ready.</p>
</td>
</tr>
<tr>
<td>
<code>selector</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#labelselector-v1-meta">
Kubernetes meta/v1.LabelSelector
</a>
</em>
</td>
<td>
<p>Represents a label query over pods that should match the desired replica count indicated by the <code>replica</code> field.
It must match the labels defined in the pod template.
More info: <a href="https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors">https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors</a></p>
</td>
</tr>
<tr>
<td>
<code>service</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#service-v1-core">
Kubernetes core/v1.Service
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the behavior of a service spec.
Provides read-write service.
<a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status</a></p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>template</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podtemplatespec-v1-core">
Kubernetes core/v1.PodTemplateSpec
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceTemplate">
[]InstanceTemplate
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Overrides values in default Template.</p>
<p>Instance is the fundamental unit managed by KubeBlocks.
It represents a Pod with additional objects such as PVCs, Services, ConfigMaps, etc.
An InstanceSet manages instances with a total count of Replicas,
and by default, all these instances are generated from the same template.
The InstanceTemplate provides a way to override values in the default template,
allowing the InstanceSet to manage instances from different templates.</p>
<p>The naming convention for instances (pods) based on the InstanceSet Name, InstanceTemplate Name, and ordinal.
The constructed instance name follows the pattern: $(instance_set.name)-$(template.name)-$(ordinal).
By default, the ordinal starts from 0 for each InstanceTemplate.
It is important to ensure that the Name of each InstanceTemplate is unique.</p>
<p>The sum of replicas across all InstanceTemplates should not exceed the total number of Replicas specified for the InstanceSet.
Any remaining replicas will be generated using the default template and will follow the default naming rules.</p>
</td>
</tr>
<tr>
<td>
<code>offlineInstances</code><br/>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the names of instances to be transitioned to offline status.</p>
<p>Marking an instance as offline results in the following:</p>
<ol>
<li>The associated pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential
future reuse or data recovery, but it is no longer actively used.</li>
<li>The ordinal number assigned to this instance is preserved, ensuring it remains unique
and avoiding conflicts with new instances.</li>
</ol>
<p>Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining
ordinal consistency within the cluster.
Note that offline instances and their associated resources, such as PVCs, are not automatically deleted.
The cluster administrator must manually manage the cleanup and removal of these resources when they are no longer needed.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeclaim-v1-core">
[]Kubernetes core/v1.PersistentVolumeClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a list of PersistentVolumeClaim templates that define the storage requirements for each replica.
Each template specifies the desired characteristics of a persistent volume, such as storage class,
size, and access modes.
These templates are used to dynamically provision persistent volumes for replicas upon their creation.
The final name of each PVC is generated by appending the pod&rsquo;s identifier to the name specified in volumeClaimTemplates[*].name.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls how pods are created during initial scale up,
when replacing pods on nodes, or when scaling down.</p>
<p>The default policy is <code>OrderedReady</code>, where pods are created in increasing order and the controller waits until each pod is ready before
continuing. When scaling down, the pods are removed in the opposite order.
The alternative policy is <code>Parallel</code> which will create pods in parallel
to match the desired scale without waiting, and on scale down will delete
all pods at once.</p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>parallelPodManagementConcurrency</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Controls the concurrency of pods during initial scale up, when replacing pods on nodes,
or when scaling down. It only used when <code>PodManagementPolicy</code> is set to <code>Parallel</code>.
The default Concurrency is 100%.</p>
</td>
</tr>
<tr>
<td>
<code>podUpdatePolicy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.PodUpdatePolicyType">
PodUpdatePolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodUpdatePolicy indicates how pods should be updated</p>
<ul>
<li><code>StrictInPlace</code> indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</li>
<li><code>PreferInPlace</code> indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.
Default value is &ldquo;PreferInPlace&rdquo;</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>updateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">
InstanceUpdateStrategy
</a>
</em>
</td>
<td>
<p>Indicates the StatefulSetUpdateStrategy that will be
employed to update Pods in the InstanceSet when a revision is made to
Template.</p>
<p>Note: This field will be removed in future version.</p>
</td>
</tr>
<tr>
<td>
<code>roles</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.ReplicaRole">
[]ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>A list of roles defined in the system.</p>
</td>
</tr>
<tr>
<td>
<code>roleProbe</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.RoleProbe">
RoleProbe
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides method to probe role.</p>
</td>
</tr>
<tr>
<td>
<code>membershipReconfiguration</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MembershipReconfiguration">
MembershipReconfiguration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides actions to do membership dynamic reconfiguration.</p>
</td>
</tr>
<tr>
<td>
<code>memberUpdateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MemberUpdateStrategy">
MemberUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Members(Pods) update strategy.</p>
<ul>
<li>serial: update Members one by one that guarantee minimum component unavailable time.</li>
<li>bestEffortParallel: update Members in parallel that guarantee minimum component un-writable time.</li>
<li>parallel: force parallel</li>
</ul>
</td>
</tr>
<tr>
<td>
<code>paused</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates that the InstanceSet is paused, meaning the reconciliation of this InstanceSet object will be paused.</p>
</td>
</tr>
<tr>
<td>
<code>credential</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Credential">
Credential
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Credential used to connect to DB engine</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.InstanceSetStatus">InstanceSetStatus
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceSet">InstanceSet</a>)
</p>
<div>
<p>InstanceSetStatus defines the observed state of InstanceSet</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>observedGeneration is the most recent generation observed for this InstanceSet. It corresponds to the
InstanceSet&rsquo;s generation, which is updated on mutation by the API Server.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>replicas is the number of instances created by the InstanceSet controller.</p>
</td>
</tr>
<tr>
<td>
<code>readyReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>readyReplicas is the number of instances created for this InstanceSet with a Ready Condition.</p>
</td>
</tr>
<tr>
<td>
<code>currentReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>currentReplicas is the number of instances created by the InstanceSet controller from the InstanceSet version
indicated by CurrentRevisions.</p>
</td>
</tr>
<tr>
<td>
<code>updatedReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>updatedReplicas is the number of instances created by the InstanceSet controller from the InstanceSet version
indicated by UpdateRevisions.</p>
</td>
</tr>
<tr>
<td>
<code>currentRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>currentRevision, if not empty, indicates the version of the InstanceSet used to generate instances in the
sequence [0,currentReplicas).</p>
</td>
</tr>
<tr>
<td>
<code>updateRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>updateRevision, if not empty, indicates the version of the InstanceSet used to generate instances in the sequence
[replicas-updatedReplicas,replicas)</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#condition-v1-meta">
[]Kubernetes meta/v1.Condition
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the latest available observations of an instanceset&rsquo;s current state.
Known .status.conditions.type are: &ldquo;InstanceFailure&rdquo;, &ldquo;InstanceReady&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>availableReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Total number of available instances (ready for at least minReadySeconds) targeted by this InstanceSet.</p>
</td>
</tr>
<tr>
<td>
<code>initReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the initial number of instances when the cluster is first initialized.
This value is set to spec.Replicas at the time of object creation and remains constant thereafter.
Used only when spec.roles set.</p>
</td>
</tr>
<tr>
<td>
<code>readyInitReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Represents the number of instances that have already reached the MembersStatus during the cluster initialization stage.
This value remains constant once it equals InitReplicas.
Used only when spec.roles set.</p>
</td>
</tr>
<tr>
<td>
<code>membersStatus</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MemberStatus">
[]MemberStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Provides the status of each member in the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>readyWithoutPrimary</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates whether it is required for the InstanceSet to have at least one primary instance ready.</p>
</td>
</tr>
<tr>
<td>
<code>currentRevisions</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>currentRevisions, if not empty, indicates the old version of the InstanceSet used to generate the underlying workload.
key is the pod name, value is the revision.</p>
</td>
</tr>
<tr>
<td>
<code>updateRevisions</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>updateRevisions, if not empty, indicates the new version of the InstanceSet used to generate the underlying workload.
key is the pod name, value is the revision.</p>
</td>
</tr>
<tr>
<td>
<code>templatesStatus</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.InstanceTemplateStatus">
[]InstanceTemplateStatus
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TemplatesStatus represents status of each instance generated by InstanceTemplates</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.InstanceTemplate">InstanceTemplate
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>InstanceTemplate allows customization of individual replica configurations within a Component,
without altering the base component template defined in ClusterComponentSpec.
It enables the application of distinct settings to specific instances (replicas),
providing flexibility while maintaining a common configuration baseline.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name specifies the unique name of the instance Pod created using this InstanceTemplate.
This name is constructed by concatenating the component&rsquo;s name, the template&rsquo;s name, and the instance&rsquo;s ordinal
using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0.
The specified name overrides any default naming conventions or patterns.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of instances (Pods) to create from this InstanceTemplate.
This field allows setting how many replicated instances of the component,
with the specific overrides in the InstanceTemplate, are created.
The default value is 1. A value of 0 disables instance creation.</p>
</td>
</tr>
<tr>
<td>
<code>ordinals</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Ordinals">
Ordinals
</a>
</em>
</td>
<td>
<p>Specifies the desired Ordinals of this InstanceTemplate.
The Ordinals used to specify the ordinal of the instance (pod) names to be generated under this InstanceTemplate.</p>
<p>For example, if Ordinals is &#123;ranges: [&#123;start: 0, end: 1&#125;], discrete: [7]&#125;,
then the instance names generated under this InstanceTemplate would be
$(cluster.name)-$(component.name)-$(template.name)-0、$(cluster.name)-$(component.name)-$(template.name)-1 and
$(cluster.name)-$(component.name)-$(template.name)-7</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs to be merged into the Pod&rsquo;s existing annotations.
Existing keys will have their values overwritten, while new keys will be added to the annotations.</p>
</td>
</tr>
<tr>
<td>
<code>labels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a map of key-value pairs that will be merged into the Pod&rsquo;s existing labels.
Values for existing keys will be overwritten, and new keys will be added.</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the first container&rsquo;s image in the pod.</p>
</td>
</tr>
<tr>
<td>
<code>schedulingPolicy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.SchedulingPolicy">
SchedulingPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the scheduling policy for the Component.</p>
</td>
</tr>
<tr>
<td>
<code>resources</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies an override for the resource requirements of the first container in the Pod.
This field allows for customizing resource allocation (CPU, memory, etc.) for the container.</p>
</td>
</tr>
<tr>
<td>
<code>env</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines Env to override.
Add new or override existing envs.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines Volumes to override.
Add new or override existing volumes.</p>
</td>
</tr>
<tr>
<td>
<code>volumeMounts</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#volumemount-v1-core">
[]Kubernetes core/v1.VolumeMount
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines VolumeMounts to override.
Add new or override existing volume mounts of the first container in the pod.</p>
</td>
</tr>
<tr>
<td>
<code>volumeClaimTemplates</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#persistentvolumeclaim-v1-core">
[]Kubernetes core/v1.PersistentVolumeClaim
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines VolumeClaimTemplates to override.
Add new or override existing volume claim templates.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.InstanceTemplateStatus">InstanceTemplateStatus
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetStatus">InstanceSetStatus</a>)
</p>
<div>
<p>InstanceTemplateStatus aggregates the status of replicas for each InstanceTemplate</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name, the name of the InstanceTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Replicas is the number of replicas of the InstanceTemplate.</p>
</td>
</tr>
<tr>
<td>
<code>readyReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>ReadyReplicas is the number of Pods that have a Ready Condition.</p>
</td>
</tr>
<tr>
<td>
<code>availableReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>AvailableReplicas is the number of Pods that ready for at least minReadySeconds.</p>
</td>
</tr>
<tr>
<td>
<code>currentReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<p>currentReplicas is the number of instances created by the InstanceSet controller from the InstanceSet version
indicated by CurrentRevisions.</p>
</td>
</tr>
<tr>
<td>
<code>updatedReplicas</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>UpdatedReplicas is the number of Pods created by the InstanceSet controller from the InstanceSet version
indicated by UpdateRevisions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">InstanceUpdateStrategy
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>InstanceUpdateStrategy indicates the strategy that the InstanceSet
controller will use to perform updates. It includes any additional parameters
necessary to perform the update for the indicated strategy.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>partition</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Partition indicates the number of pods that should be updated during a rolling update.
The remaining pods will remain untouched. This is helpful in defining how many pods
should participate in the update process. The update process will follow the order
of pod names in descending lexicographical (dictionary) order. The default value is
Replicas (i.e., update all pods).</p>
</td>
</tr>
<tr>
<td>
<code>maxUnavailable</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/util/intstr#IntOrString">
Kubernetes api utils intstr.IntOrString
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>The maximum number of pods that can be unavailable during the update.
Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
Absolute number is calculated from percentage by rounding up. This can not be 0.
Defaults to 1. The field applies to all pods. That means if there is any unavailable pod,
it will be counted towards MaxUnavailable.</p>
</td>
</tr>
<tr>
<td>
<code>memberUpdateStrategy</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.MemberUpdateStrategy">
MemberUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Members(Pods) update strategy.</p>
<ul>
<li>serial: update Members one by one that guarantee minimum component unavailable time.</li>
<li>bestEffortParallel: update Members in parallel that guarantee minimum component un-writable time.</li>
<li>parallel: force parallel</li>
</ul>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.MemberStatus">MemberStatus
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentStatus">ClusterComponentStatus</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetStatus">InstanceSetStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>podName</code><br/>
<em>
string
</em>
</td>
<td>
<p>Represents the name of the pod.</p>
</td>
</tr>
<tr>
<td>
<code>role</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.ReplicaRole">
ReplicaRole
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the role of the replica in the cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.MemberUpdateStrategy">MemberUpdateStrategy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RSMSpec">RSMSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceUpdateStrategy">InstanceUpdateStrategy</a>)
</p>
<div>
<p>MemberUpdateStrategy defines Cluster Component update strategy.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;BestEffortParallel&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Parallel&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Serial&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.MembershipReconfiguration">MembershipReconfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RSMSpec">RSMSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>switchoverAction</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the environment variables that can be used in all following Actions:
- KB_ITS_USERNAME: Represents the username part of the credential
- KB_ITS_PASSWORD: Represents the password part of the credential
- KB_ITS_LEADER_HOST: Represents the leader host
- KB_ITS_TARGET_HOST: Represents the target host
- KB_ITS_SERVICE_PORT: Represents the service port</p>
<p>Defines the action to perform a switchover.
If the Image is not configured, the latest <a href="https://busybox.net/">BusyBox</a> image will be used.</p>
</td>
</tr>
<tr>
<td>
<code>memberJoinAction</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the action to add a member.
If the Image is not configured, the Image from the previous non-nil action will be used.</p>
</td>
</tr>
<tr>
<td>
<code>memberLeaveAction</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the action to remove a member.
If the Image is not configured, the Image from the previous non-nil action will be used.</p>
</td>
</tr>
<tr>
<td>
<code>logSyncAction</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the action to trigger the new member to start log syncing.
If the Image is not configured, the Image from the previous non-nil action will be used.</p>
</td>
</tr>
<tr>
<td>
<code>promoteAction</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Action">
Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines the action to inform the cluster that the new member can join voting now.
If the Image is not configured, the Image from the previous non-nil action will be used.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.Ordinals">Ordinals
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceTemplate">InstanceTemplate</a>)
</p>
<div>
<p>Ordinals represents a combination of continuous segments and individual values.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ranges</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Range">
[]Range
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>discrete</code><br/>
<em>
[]int32
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.PodUpdatePolicyType">PodUpdatePolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.ClusterComponentSpec">ClusterComponentSpec</a>, <a href="#apps.kubeblocks.io/v1alpha1.ComponentSpec">ComponentSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;PreferInPlace&#34;</p></td>
<td><p>PreferInPlacePodUpdatePolicyType indicates that we will first attempt an in-place upgrade of the Pod.
If that fails, it will fall back to the ReCreate, where pod will be recreated.</p>
</td>
</tr><tr><td><p>&#34;StrictInPlace&#34;</p></td>
<td><p>StrictInPlacePodUpdatePolicyType indicates that only allows in-place upgrades.
Any attempt to modify other fields will be rejected.</p>
</td>
</tr></tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.Range">Range
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.Ordinals">Ordinals</a>)
</p>
<div>
<p>Range represents a range with a start and an end value.
It is used to define a continuous segment.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>start</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>end</code><br/>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.ReplicaRole">ReplicaRole
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RSMSpec">RSMSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.MemberStatus">MemberStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Defines the role name of the replica.</p>
</td>
</tr>
<tr>
<td>
<code>accessMode</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.AccessMode">
AccessMode
</a>
</em>
</td>
<td>
<p>Specifies the service capabilities of this member.</p>
</td>
</tr>
<tr>
<td>
<code>canVote</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Indicates if this member has voting rights.</p>
</td>
</tr>
<tr>
<td>
<code>isLeader</code><br/>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Determines if this member is the leader.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.RoleProbe">RoleProbe
</h3>
<p>
(<em>Appears on:</em><a href="#apps.kubeblocks.io/v1alpha1.RSMSpec">RSMSpec</a>, <a href="#workloads.kubeblocks.io/v1alpha1.InstanceSetSpec">InstanceSetSpec</a>)
</p>
<div>
<p>RoleProbe defines how to observe role</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>builtinHandlerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the builtin handler name to use to probe the role of the main container.
Available handlers include: mysql, postgres, mongodb, redis, etcd, kafka.
Use CustomHandler to define a custom role probe function if none of the built-in handlers meet the requirement.</p>
</td>
</tr>
<tr>
<td>
<code>customHandler</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.Action">
[]Action
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Defines a custom method for role probing.
Actions defined here are executed in series.
Upon completion of all actions, the final output should be a single string representing the role name defined in spec.Roles.
The latest <a href="https://busybox.net/">BusyBox</a> image will be used if Image is not configured.
Environment variables can be used in Command:
- v_KB_ITS_LAST<em>STDOUT: stdout from the last action, watch for &lsquo;v</em>&rsquo; prefix
- KB_ITS_USERNAME: username part of the credential
- KB_ITS_PASSWORD: password part of the credential</p>
</td>
</tr>
<tr>
<td>
<code>initialDelaySeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of seconds to wait after the container has started before initiating role probing.</p>
</td>
</tr>
<tr>
<td>
<code>timeoutSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the number of seconds after which the probe times out.</p>
</td>
</tr>
<tr>
<td>
<code>periodSeconds</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the frequency (in seconds) of probe execution.</p>
</td>
</tr>
<tr>
<td>
<code>successThreshold</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the minimum number of consecutive successes for the probe to be considered successful after having failed.</p>
</td>
</tr>
<tr>
<td>
<code>failureThreshold</code><br/>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the minimum number of consecutive failures for the probe to be considered failed after having succeeded.</p>
</td>
</tr>
<tr>
<td>
<code>roleUpdateMechanism</code><br/>
<em>
<a href="#workloads.kubeblocks.io/v1alpha1.RoleUpdateMechanism">
RoleUpdateMechanism
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies the method for updating the pod role label.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.RoleUpdateMechanism">RoleUpdateMechanism
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.RoleProbe">RoleProbe</a>)
</p>
<div>
<p>RoleUpdateMechanism defines the way how pod role label being updated.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;DirectAPIServerEventUpdate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReadinessProbeEventUpdate&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="workloads.kubeblocks.io/v1alpha1.SchedulingPolicy">SchedulingPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#workloads.kubeblocks.io/v1alpha1.InstanceTemplate">InstanceTemplate</a>)
</p>
<div>
<p>SchedulingPolicy the scheduling policy.
Deprecated: Unify with apps/v1alpha1.SchedulingPolicy</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>schedulerName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>If specified, the Pod will be dispatched by specified scheduler.
If not specified, the Pod will be dispatched by default scheduler.</p>
</td>
</tr>
<tr>
<td>
<code>nodeSelector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeSelector is a selector which must be true for the Pod to fit on a node.
Selector which must match a node&rsquo;s labels for the Pod to be scheduled on that node.
More info: <a href="https://kubernetes.io/docs/concepts/configuration/assign-pod-node/">https://kubernetes.io/docs/concepts/configuration/assign-pod-node/</a></p>
</td>
</tr>
<tr>
<td>
<code>nodeName</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeName is a request to schedule this Pod onto a specific node. If it is non-empty,
the scheduler simply schedules this Pod onto that node, assuming that it fits resource
requirements.</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#affinity-v1-core">
Kubernetes core/v1.Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Allows Pods to be scheduled onto nodes with matching taints.
Each toleration in the array allows the Pod to tolerate node taints based on
specified <code>key</code>, <code>value</code>, <code>effect</code>, and <code>operator</code>.</p>
<ul>
<li>The <code>key</code>, <code>value</code>, and <code>effect</code> identify the taint that the toleration matches.</li>
<li>The <code>operator</code> determines how the toleration matches the taint.</li>
</ul>
<p>Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.</p>
</td>
</tr>
<tr>
<td>
<code>topologySpreadConstraints</code><br/>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.25/#topologyspreadconstraint-v1-core">
[]Kubernetes core/v1.TopologySpreadConstraint
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>TopologySpreadConstraints describes how a group of Pods ought to spread across topology
domains. Scheduler will schedule Pods in a way which abides by the constraints.
All topologySpreadConstraints are ANDed.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
</em></p>