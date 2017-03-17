/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_APIVersion = map[string]string{
	"":     "An APIVersion represents a single concrete version of an object model.",
	"name": "Name of this version (e.g. 'v1').",
}

func (APIVersion) SwaggerDoc() map[string]string {
	return map_APIVersion
}

var map_CustomMetricCurrentStatus = map[string]string{
	"name":  "Custom Metric name.",
	"value": "Custom Metric value (average).",
}

func (CustomMetricCurrentStatus) SwaggerDoc() map[string]string {
	return map_CustomMetricCurrentStatus
}

var map_CustomMetricTarget = map[string]string{
	"":      "Alpha-level support for Custom Metrics in HPA (as annotations).",
	"name":  "Custom Metric name.",
	"value": "Custom Metric value (average).",
}

func (CustomMetricTarget) SwaggerDoc() map[string]string {
	return map_CustomMetricTarget
}

var map_DaemonSet = map[string]string{
	"":         "DaemonSet represents the configuration of a daemon set.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec defines the desired behavior of this daemon set. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status":   "Status is the current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (DaemonSet) SwaggerDoc() map[string]string {
	return map_DaemonSet
}

var map_DaemonSetList = map[string]string{
	"":         "DaemonSetList is a collection of daemon sets.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is a list of daemon sets.",
}

func (DaemonSetList) SwaggerDoc() map[string]string {
	return map_DaemonSetList
}

var map_DaemonSetSpec = map[string]string{
	"":                   "DaemonSetSpec is the specification of a daemon set.",
	"selector":           "Selector is a label query over pods that are managed by the daemon set. Must match in order to be controlled. If empty, defaulted to labels on Pod template. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors",
	"template":           "Template is the object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: http://kubernetes.io/docs/user-guide/replication-controller#pod-template",
	"updateStrategy":     "UpdateStrategy to replace existing DaemonSet pods with new pods.",
	"minReadySeconds":    "The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready).",
	"templateGeneration": "A sequence number representing a specific generation of the template. Populated by the system. It can be set only during the creation.",
}

func (DaemonSetSpec) SwaggerDoc() map[string]string {
	return map_DaemonSetSpec
}

var map_DaemonSetStatus = map[string]string{
	"": "DaemonSetStatus represents the current status of a daemon set.",
	"currentNumberScheduled": "CurrentNumberScheduled is the number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md",
	"numberMisscheduled":     "NumberMisscheduled is the number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md",
	"desiredNumberScheduled": "DesiredNumberScheduled is the total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md",
	"numberReady":            "NumberReady is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.",
	"observedGeneration":     "ObservedGeneration is the most recent generation observed by the daemon set controller.",
	"updatedNumberScheduled": "UpdatedNumberScheduled is the total number of nodes that are running updated daemon pod",
	"numberAvailable":        "NumberAvailable is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least minReadySeconds)",
	"numberUnavailable":      "NumberUnavailable is the number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least minReadySeconds)",
}

func (DaemonSetStatus) SwaggerDoc() map[string]string {
	return map_DaemonSetStatus
}

var map_DaemonSetUpdateStrategy = map[string]string{
	"type":          "Type of daemon set update. Can be \"RollingUpdate\" or \"OnDelete\". Default is OnDelete.",
	"rollingUpdate": "Rolling update config params. Present only if DaemonSetUpdateStrategy = RollingUpdate.",
}

func (DaemonSetUpdateStrategy) SwaggerDoc() map[string]string {
	return map_DaemonSetUpdateStrategy
}

var map_Deployment = map[string]string{
	"":         "Deployment enables declarative updates for Pods and ReplicaSets.",
	"metadata": "Standard object metadata.",
	"spec":     "Specification of the desired behavior of the Deployment.",
	"status":   "Most recently observed status of the Deployment.",
}

func (Deployment) SwaggerDoc() map[string]string {
	return map_Deployment
}

var map_DeploymentCondition = map[string]string{
	"":                   "DeploymentCondition describes the state of a deployment at a certain point.",
	"type":               "Type of deployment condition.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastUpdateTime":     "The last time this condition was updated.",
	"lastTransitionTime": "Last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (DeploymentCondition) SwaggerDoc() map[string]string {
	return map_DeploymentCondition
}

var map_DeploymentList = map[string]string{
	"":         "DeploymentList is a list of Deployments.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of Deployments.",
}

func (DeploymentList) SwaggerDoc() map[string]string {
	return map_DeploymentList
}

var map_DeploymentRollback = map[string]string{
	"":                   "DeploymentRollback stores the information required to rollback a deployment.",
	"name":               "Required: This must match the Name of a deployment.",
	"updatedAnnotations": "The annotations to be updated to a deployment",
	"rollbackTo":         "The config of this deployment rollback.",
}

func (DeploymentRollback) SwaggerDoc() map[string]string {
	return map_DeploymentRollback
}

var map_DeploymentSpec = map[string]string{
	"":                        "DeploymentSpec is the specification of the desired behavior of the Deployment.",
	"replicas":                "Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",
	"selector":                "Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment.",
	"template":                "Template describes the pods that will be created.",
	"strategy":                "The deployment strategy to use to replace existing pods with new ones.",
	"minReadySeconds":         "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
	"revisionHistoryLimit":    "The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified.",
	"paused":                  "Indicates that the deployment is paused and will not be processed by the deployment controller.",
	"rollbackTo":              "The config this deployment is rolling back to. Will be cleared after rollback is done.",
	"progressDeadlineSeconds": "The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Once autoRollback is implemented, the deployment controller will automatically rollback failed deployments. Note that progress will not be estimated during the time a deployment is paused. This is not set by default.",
}

func (DeploymentSpec) SwaggerDoc() map[string]string {
	return map_DeploymentSpec
}

var map_DeploymentStatus = map[string]string{
	"":                    "DeploymentStatus is the most recently observed status of the Deployment.",
	"observedGeneration":  "The generation observed by the deployment controller.",
	"replicas":            "Total number of non-terminated pods targeted by this deployment (their labels match the selector).",
	"updatedReplicas":     "Total number of non-terminated pods targeted by this deployment that have the desired template spec.",
	"readyReplicas":       "Total number of ready pods targeted by this deployment.",
	"availableReplicas":   "Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.",
	"unavailableReplicas": "Total number of unavailable pods targeted by this deployment.",
	"conditions":          "Represents the latest available observations of a deployment's current state.",
}

func (DeploymentStatus) SwaggerDoc() map[string]string {
	return map_DeploymentStatus
}

var map_DeploymentStrategy = map[string]string{
	"":              "DeploymentStrategy describes how to replace existing pods with new ones.",
	"type":          "Type of deployment. Can be \"Recreate\" or \"RollingUpdate\". Default is RollingUpdate.",
	"rollingUpdate": "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.",
}

func (DeploymentStrategy) SwaggerDoc() map[string]string {
	return map_DeploymentStrategy
}

var map_FSGroupStrategyOptions = map[string]string{
	"":       "FSGroupStrategyOptions defines the strategy type and options used to create the strategy.",
	"rule":   "Rule is the strategy that will dictate what FSGroup is used in the SecurityContext.",
	"ranges": "Ranges are the allowed ranges of fs groups.  If you would like to force a single fs group then supply a single range with the same start and end.",
}

func (FSGroupStrategyOptions) SwaggerDoc() map[string]string {
	return map_FSGroupStrategyOptions
}

var map_HTTPIngressPath = map[string]string{
	"":        "HTTPIngressPath associates a path regex with a backend. Incoming urls matching the path are forwarded to the backend.",
	"path":    "Path is an extended POSIX regex as defined by IEEE Std 1003.1, (i.e this follows the egrep/unix syntax, not the perl syntax) matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \"path\" part of a URL as defined by RFC 3986. Paths must begin with a '/'. If unspecified, the path defaults to a catch all sending traffic to the backend.",
	"backend": "Backend defines the referenced service endpoint to which the traffic will be forwarded to.",
}

func (HTTPIngressPath) SwaggerDoc() map[string]string {
	return map_HTTPIngressPath
}

var map_HTTPIngressRuleValue = map[string]string{
	"":      "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
	"paths": "A collection of paths that map requests to backends.",
}

func (HTTPIngressRuleValue) SwaggerDoc() map[string]string {
	return map_HTTPIngressRuleValue
}

var map_HostPortRange = map[string]string{
	"":    "Host Port Range defines a range of host ports that will be enabled by a policy for pods to use.  It requires both the start and end to be defined.",
	"min": "min is the start of the range, inclusive.",
	"max": "max is the end of the range, inclusive.",
}

func (HostPortRange) SwaggerDoc() map[string]string {
	return map_HostPortRange
}

var map_IDRange = map[string]string{
	"":    "ID Range provides a min/max of an allowed range of IDs.",
	"min": "Min is the start of the range, inclusive.",
	"max": "Max is the end of the range, inclusive.",
}

func (IDRange) SwaggerDoc() map[string]string {
	return map_IDRange
}

var map_Ingress = map[string]string{
	"":         "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec is the desired state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status":   "Status is the current state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (Ingress) SwaggerDoc() map[string]string {
	return map_Ingress
}

var map_IngressBackend = map[string]string{
	"":            "IngressBackend describes all endpoints for a given service and port.",
	"serviceName": "Specifies the name of the referenced service.",
	"servicePort": "Specifies the port of the referenced service.",
}

func (IngressBackend) SwaggerDoc() map[string]string {
	return map_IngressBackend
}

var map_IngressList = map[string]string{
	"":         "IngressList is a collection of Ingress.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of Ingress.",
}

func (IngressList) SwaggerDoc() map[string]string {
	return map_IngressList
}

var map_IngressRule = map[string]string{
	"":     "IngressRule represents the rules mapping the paths under a specified host to the related backend services. Incoming requests are first evaluated for a host match, then routed to the backend associated with the matching IngressRuleValue.",
	"host": "Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the \"host\" part of the URI as defined in the RFC: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the\n\t  IP in the Spec of the parent Ingress.\n2. The `:` delimiter is not respected because ports are not allowed.\n\t  Currently the port of an Ingress is implicitly :80 for http and\n\t  :443 for https.\nBoth these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.",
}

func (IngressRule) SwaggerDoc() map[string]string {
	return map_IngressRule
}

var map_IngressRuleValue = map[string]string{
	"": "IngressRuleValue represents a rule to apply against incoming requests. If the rule is satisfied, the request is routed to the specified backend. Currently mixing different types of rules in a single Ingress is disallowed, so exactly one of the following must be set.",
}

func (IngressRuleValue) SwaggerDoc() map[string]string {
	return map_IngressRuleValue
}

var map_IngressSpec = map[string]string{
	"":        "IngressSpec describes the Ingress the user wishes to exist.",
	"backend": "A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.",
	"tls":     "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
	"rules":   "A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
}

func (IngressSpec) SwaggerDoc() map[string]string {
	return map_IngressSpec
}

var map_IngressStatus = map[string]string{
	"":             "IngressStatus describe the current state of the Ingress.",
	"loadBalancer": "LoadBalancer contains the current status of the load-balancer.",
}

func (IngressStatus) SwaggerDoc() map[string]string {
	return map_IngressStatus
}

var map_IngressTLS = map[string]string{
	"":           "IngressTLS describes the transport layer security associated with an Ingress.",
	"hosts":      "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
	"secretName": "SecretName is the name of the secret used to terminate SSL traffic on 443. Field is left optional to allow SSL routing based on SNI hostname alone. If the SNI host in a listener conflicts with the \"Host\" header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
}

func (IngressTLS) SwaggerDoc() map[string]string {
	return map_IngressTLS
}

var map_NetworkPolicy = map[string]string{
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Specification of the desired behavior for this NetworkPolicy.",
}

func (NetworkPolicy) SwaggerDoc() map[string]string {
	return map_NetworkPolicy
}

var map_NetworkPolicyIngressRule = map[string]string{
	"":      "This NetworkPolicyIngressRule matches traffic if and only if the traffic matches both ports AND from.",
	"ports": "List of ports which should be made accessible on the pods selected for this rule. Each item in this list is combined using a logical OR. If this field is not provided, this rule matches all ports (traffic not restricted by port). If this field is empty, this rule matches no ports (no traffic matches). If this field is present and contains at least one item, then this rule allows traffic only if the traffic matches at least one port in the list.",
	"from":  "List of sources which should be able to access the pods selected for this rule. Items in this list are combined using a logical OR operation. If this field is not provided, this rule matches all sources (traffic not restricted by source). If this field is empty, this rule matches no sources (no traffic matches). If this field is present and contains at least on item, this rule allows traffic only if the traffic matches at least one item in the from list.",
}

func (NetworkPolicyIngressRule) SwaggerDoc() map[string]string {
	return map_NetworkPolicyIngressRule
}

var map_NetworkPolicyList = map[string]string{
	"":         "Network Policy List is a list of NetworkPolicy objects.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is a list of schema objects.",
}

func (NetworkPolicyList) SwaggerDoc() map[string]string {
	return map_NetworkPolicyList
}

var map_NetworkPolicyPeer = map[string]string{
	"podSelector":       "This is a label selector which selects Pods in this namespace. This field follows standard label selector semantics. If not provided, this selector selects no pods. If present but empty, this selector selects all pods in this namespace.",
	"namespaceSelector": "Selects Namespaces using cluster scoped-labels.  This matches all pods in all namespaces selected by this label selector. This field follows standard label selector semantics. If omitted, this selector selects no namespaces. If present but empty, this selector selects all namespaces.",
}

func (NetworkPolicyPeer) SwaggerDoc() map[string]string {
	return map_NetworkPolicyPeer
}

var map_NetworkPolicyPort = map[string]string{
	"protocol": "Optional.  The protocol (TCP or UDP) which traffic must match. If not specified, this field defaults to TCP.",
	"port":     "If specified, the port on the given protocol.  This can either be a numerical or named port on a pod.  If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.",
}

func (NetworkPolicyPort) SwaggerDoc() map[string]string {
	return map_NetworkPolicyPort
}

var map_NetworkPolicySpec = map[string]string{
	"podSelector": "Selects the pods to which this NetworkPolicy object applies.  The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods.  In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.",
	"ingress":     "List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if namespace.networkPolicy.ingress.isolation is undefined and cluster policy allows it, OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not affect ingress isolation. If this field is present and contains at least one rule, this policy allows any traffic which matches at least one of the ingress rules in this list.",
}

func (NetworkPolicySpec) SwaggerDoc() map[string]string {
	return map_NetworkPolicySpec
}

var map_PodSecurityPolicy = map[string]string{
	"":         "Pod Security Policy governs the ability to make requests that affect the Security Context that will be applied to a pod and container.",
	"metadata": "Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "spec defines the policy enforced.",
}

func (PodSecurityPolicy) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicy
}

var map_PodSecurityPolicyList = map[string]string{
	"":         "Pod Security Policy List is a list of PodSecurityPolicy objects.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is a list of schema objects.",
}

func (PodSecurityPolicyList) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicyList
}

var map_PodSecurityPolicySpec = map[string]string{
	"":                         "Pod Security Policy Spec defines the policy enforced.",
	"privileged":               "privileged determines if a pod can request to be run as privileged.",
	"defaultAddCapabilities":   "DefaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.  You may not list a capabiility in both DefaultAddCapabilities and RequiredDropCapabilities.",
	"requiredDropCapabilities": "RequiredDropCapabilities are the capabilities that will be dropped from the container.  These are required to be dropped and cannot be added.",
	"allowedCapabilities":      "AllowedCapabilities is a list of capabilities that can be requested to add to the container. Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both AllowedCapabilities and RequiredDropCapabilities.",
	"volumes":                  "volumes is a white list of allowed volume plugins.  Empty indicates that all plugins may be used.",
	"hostNetwork":              "hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.",
	"hostPorts":                "hostPorts determines which host port ranges are allowed to be exposed.",
	"hostPID":                  "hostPID determines if the policy allows the use of HostPID in the pod spec.",
	"hostIPC":                  "hostIPC determines if the policy allows the use of HostIPC in the pod spec.",
	"seLinux":                  "seLinux is the strategy that will dictate the allowable labels that may be set.",
	"runAsUser":                "runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.",
	"supplementalGroups":       "SupplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.",
	"fsGroup":                  "FSGroup is the strategy that will dictate what fs group is used by the SecurityContext.",
	"readOnlyRootFilesystem":   "ReadOnlyRootFilesystem when set to true will force containers to run with a read only root file system.  If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.",
}

func (PodSecurityPolicySpec) SwaggerDoc() map[string]string {
	return map_PodSecurityPolicySpec
}

var map_ReplicaSet = map[string]string{
	"":         "ReplicaSet represents the configuration of a ReplicaSet.",
	"metadata": "If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"spec":     "Spec defines the specification of the desired behavior of the ReplicaSet. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
	"status":   "Status is the most recently observed status of the ReplicaSet. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status",
}

func (ReplicaSet) SwaggerDoc() map[string]string {
	return map_ReplicaSet
}

var map_ReplicaSetCondition = map[string]string{
	"":                   "ReplicaSetCondition describes the state of a replica set at a certain point.",
	"type":               "Type of replica set condition.",
	"status":             "Status of the condition, one of True, False, Unknown.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (ReplicaSetCondition) SwaggerDoc() map[string]string {
	return map_ReplicaSetCondition
}

var map_ReplicaSetList = map[string]string{
	"":         "ReplicaSetList is a collection of ReplicaSets.",
	"metadata": "Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds",
	"items":    "List of ReplicaSets. More info: http://kubernetes.io/docs/user-guide/replication-controller",
}

func (ReplicaSetList) SwaggerDoc() map[string]string {
	return map_ReplicaSetList
}

var map_ReplicaSetSpec = map[string]string{
	"":                "ReplicaSetSpec is the specification of a ReplicaSet.",
	"replicas":        "Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller",
	"minReadySeconds": "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
	"selector":        "Selector is a label query over pods that should match the replica count. If the selector is empty, it is defaulted to the labels present on the pod template. Label keys and values that must match in order to be controlled by this replica set. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors",
	"template":        "Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: http://kubernetes.io/docs/user-guide/replication-controller#pod-template",
}

func (ReplicaSetSpec) SwaggerDoc() map[string]string {
	return map_ReplicaSetSpec
}

var map_ReplicaSetStatus = map[string]string{
	"":                     "ReplicaSetStatus represents the current status of a ReplicaSet.",
	"replicas":             "Replicas is the most recently oberved number of replicas. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller",
	"fullyLabeledReplicas": "The number of pods that have labels matching the labels of the pod template of the replicaset.",
	"readyReplicas":        "The number of ready replicas for this replica set.",
	"availableReplicas":    "The number of available replicas (ready for at least minReadySeconds) for this replica set.",
	"observedGeneration":   "ObservedGeneration reflects the generation of the most recently observed ReplicaSet.",
	"conditions":           "Represents the latest available observations of a replica set's current state.",
}

func (ReplicaSetStatus) SwaggerDoc() map[string]string {
	return map_ReplicaSetStatus
}

var map_ReplicationControllerDummy = map[string]string{
	"": "Dummy definition",
}

func (ReplicationControllerDummy) SwaggerDoc() map[string]string {
	return map_ReplicationControllerDummy
}

var map_RollbackConfig = map[string]string{
	"revision": "The revision to rollback to. If set to 0, rollbck to the last revision.",
}

func (RollbackConfig) SwaggerDoc() map[string]string {
	return map_RollbackConfig
}

var map_RollingUpdateDaemonSet = map[string]string{
	"":               "Spec to control the desired behavior of daemon set rolling update.",
	"maxUnavailable": "The maximum number of DaemonSet pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total number of DaemonSet pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up. This cannot be 0. Default value is 1. Example: when this is set to 30%, at most 30% of the total number of nodes that should be running the daemon pod (i.e. DesiredNumberScheduled in DaemonSetStatus) can have their pods stopped for an update at any given time. The update starts by stopping at most 30% of those DaemonSet pods and then brings up new DaemonSet pods in their place. Once the new pods are available, it then proceeds onto other DaemonSet pods, thus ensuring that at least 70% of original number of DaemonSet pods are available at all times during the update.",
}

func (RollingUpdateDaemonSet) SwaggerDoc() map[string]string {
	return map_RollingUpdateDaemonSet
}

var map_RollingUpdateDeployment = map[string]string{
	"":               "Spec to control the desired behavior of rolling update.",
	"maxUnavailable": "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. By default, a fixed value of 1 is used. Example: when this is set to 30%, the old RC can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
	"maxSurge":       "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. By default, a value of 1 is used. Example: when this is set to 30%, the new RC can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of desired pods.",
}

func (RollingUpdateDeployment) SwaggerDoc() map[string]string {
	return map_RollingUpdateDeployment
}

var map_RunAsUserStrategyOptions = map[string]string{
	"":       "Run A sUser Strategy Options defines the strategy type and any options used to create the strategy.",
	"rule":   "Rule is the strategy that will dictate the allowable RunAsUser values that may be set.",
	"ranges": "Ranges are the allowed ranges of uids that may be used.",
}

func (RunAsUserStrategyOptions) SwaggerDoc() map[string]string {
	return map_RunAsUserStrategyOptions
}

var map_SELinuxStrategyOptions = map[string]string{
	"":               "SELinux  Strategy Options defines the strategy type and any options used to create the strategy.",
	"rule":           "type is the strategy that will dictate the allowable labels that may be set.",
	"seLinuxOptions": "seLinuxOptions required to run as; required for MustRunAs More info: http://releases.k8s.io/HEAD/docs/design/security_context.md#security-context",
}

func (SELinuxStrategyOptions) SwaggerDoc() map[string]string {
	return map_SELinuxStrategyOptions
}

var map_Scale = map[string]string{
	"":         "represents a scaling request for a resource.",
	"metadata": "Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.",
	"spec":     "defines the behavior of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.",
	"status":   "current status of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status. Read-only.",
}

func (Scale) SwaggerDoc() map[string]string {
	return map_Scale
}

var map_ScaleSpec = map[string]string{
	"":         "describes the attributes of a scale subresource",
	"replicas": "desired number of instances for the scaled object.",
}

func (ScaleSpec) SwaggerDoc() map[string]string {
	return map_ScaleSpec
}

var map_ScaleStatus = map[string]string{
	"":               "represents the current status of a scale subresource.",
	"replicas":       "actual number of observed instances of the scaled object.",
	"selector":       "label query over pods that should match the replicas count. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors",
	"targetSelector": "label selector for pods that should match the replicas count. This is a serializated version of both map-based and more expressive set-based selectors. This is done to avoid introspection in the clients. The string will be in the same format as the query-param syntax. If the target type only supports map-based selectors, both this field and map-based selector field are populated. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors",
}

func (ScaleStatus) SwaggerDoc() map[string]string {
	return map_ScaleStatus
}

var map_SupplementalGroupsStrategyOptions = map[string]string{
	"":       "SupplementalGroupsStrategyOptions defines the strategy type and options used to create the strategy.",
	"rule":   "Rule is the strategy that will dictate what supplemental groups is used in the SecurityContext.",
	"ranges": "Ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end.",
}

func (SupplementalGroupsStrategyOptions) SwaggerDoc() map[string]string {
	return map_SupplementalGroupsStrategyOptions
}

var map_ThirdPartyResource = map[string]string{
	"":            "A ThirdPartyResource is a generic representation of a resource, it is used by add-ons and plugins to add new resource types to the API.  It consists of one or more Versions of the api.",
	"metadata":    "Standard object metadata",
	"description": "Description is the description of this object.",
	"versions":    "Versions are versions for this third party object",
}

func (ThirdPartyResource) SwaggerDoc() map[string]string {
	return map_ThirdPartyResource
}

var map_ThirdPartyResourceData = map[string]string{
	"":         "An internal object, used for versioned storage in etcd.  Not exposed to the end user.",
	"metadata": "Standard object metadata.",
	"data":     "Data is the raw JSON data for this data.",
}

func (ThirdPartyResourceData) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceData
}

var map_ThirdPartyResourceDataList = map[string]string{
	"":         "ThirdPartyResrouceDataList is a list of ThirdPartyResourceData.",
	"metadata": "Standard list metadata More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata",
	"items":    "Items is the list of ThirdpartyResourceData.",
}

func (ThirdPartyResourceDataList) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceDataList
}

var map_ThirdPartyResourceList = map[string]string{
	"":         "ThirdPartyResourceList is a list of ThirdPartyResources.",
	"metadata": "Standard list metadata.",
	"items":    "Items is the list of ThirdPartyResources.",
}

func (ThirdPartyResourceList) SwaggerDoc() map[string]string {
	return map_ThirdPartyResourceList
}

// AUTO-GENERATED FUNCTIONS END HERE
