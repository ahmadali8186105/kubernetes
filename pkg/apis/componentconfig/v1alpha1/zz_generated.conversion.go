// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentconfig "k8s.io/kubernetes/pkg/apis/componentconfig"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration,
		Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration,
		Convert_v1alpha1_GroupResource_To_componentconfig_GroupResource,
		Convert_componentconfig_GroupResource_To_v1alpha1_GroupResource,
		Convert_v1alpha1_KubeControllerManagerConfiguration_To_componentconfig_KubeControllerManagerConfiguration,
		Convert_componentconfig_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration,
		Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration,
		Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration,
		Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration,
		Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration,
		Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
		Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
		Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_componentconfig_PersistentVolumeRecyclerConfiguration,
		Convert_componentconfig_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration,
		Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource,
		Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource,
		Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource,
		Convert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource,
		Convert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource,
		Convert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource,
		Convert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource,
		Convert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource,
		Convert_v1alpha1_VolumeConfiguration_To_componentconfig_VolumeConfiguration,
		Convert_componentconfig_VolumeConfiguration_To_v1alpha1_VolumeConfiguration,
	)
}

func autoConvert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *componentconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}

// Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *componentconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *componentconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}

// Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *componentconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_GroupResource_To_componentconfig_GroupResource(in *GroupResource, out *componentconfig.GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_v1alpha1_GroupResource_To_componentconfig_GroupResource is an autogenerated conversion function.
func Convert_v1alpha1_GroupResource_To_componentconfig_GroupResource(in *GroupResource, out *componentconfig.GroupResource, s conversion.Scope) error {
	return autoConvert_v1alpha1_GroupResource_To_componentconfig_GroupResource(in, out, s)
}

func autoConvert_componentconfig_GroupResource_To_v1alpha1_GroupResource(in *componentconfig.GroupResource, out *GroupResource, s conversion.Scope) error {
	out.Group = in.Group
	out.Resource = in.Resource
	return nil
}

// Convert_componentconfig_GroupResource_To_v1alpha1_GroupResource is an autogenerated conversion function.
func Convert_componentconfig_GroupResource_To_v1alpha1_GroupResource(in *componentconfig.GroupResource, out *GroupResource, s conversion.Scope) error {
	return autoConvert_componentconfig_GroupResource_To_v1alpha1_GroupResource(in, out, s)
}

func autoConvert_v1alpha1_KubeControllerManagerConfiguration_To_componentconfig_KubeControllerManagerConfiguration(in *KubeControllerManagerConfiguration, out *componentconfig.KubeControllerManagerConfiguration, s conversion.Scope) error {
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	out.Port = in.Port
	out.Address = in.Address
	out.UseServiceAccountCredentials = in.UseServiceAccountCredentials
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.ExternalCloudVolumePlugin = in.ExternalCloudVolumePlugin
	out.AllowUntaggedCloud = in.AllowUntaggedCloud
	out.ConcurrentEndpointSyncs = in.ConcurrentEndpointSyncs
	out.ConcurrentRSSyncs = in.ConcurrentRSSyncs
	out.ConcurrentRCSyncs = in.ConcurrentRCSyncs
	out.ConcurrentServiceSyncs = in.ConcurrentServiceSyncs
	out.ConcurrentResourceQuotaSyncs = in.ConcurrentResourceQuotaSyncs
	out.ConcurrentDeploymentSyncs = in.ConcurrentDeploymentSyncs
	out.ConcurrentDaemonSetSyncs = in.ConcurrentDaemonSetSyncs
	out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
	out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
	out.ConcurrentSATokenSyncs = in.ConcurrentSATokenSyncs
	out.NodeSyncPeriod = in.NodeSyncPeriod
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
	out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
	out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
	out.MinResyncPeriod = in.MinResyncPeriod
	out.TerminatedPodGCThreshold = in.TerminatedPodGCThreshold
	out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
	out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
	out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
	out.HorizontalPodAutoscalerTolerance = in.HorizontalPodAutoscalerTolerance
	out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
	out.PodEvictionTimeout = in.PodEvictionTimeout
	out.DeletingPodsQps = in.DeletingPodsQps
	out.DeletingPodsBurst = in.DeletingPodsBurst
	out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
	out.RegisterRetryCount = in.RegisterRetryCount
	out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	out.ServiceAccountKeyFile = in.ServiceAccountKeyFile
	out.ClusterSigningCertFile = in.ClusterSigningCertFile
	out.ClusterSigningKeyFile = in.ClusterSigningKeyFile
	out.ClusterSigningDuration = in.ClusterSigningDuration
	out.EnableProfiling = in.EnableProfiling
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.ClusterName = in.ClusterName
	out.ClusterCIDR = in.ClusterCIDR
	out.ServiceCIDR = in.ServiceCIDR
	out.NodeCIDRMaskSize = in.NodeCIDRMaskSize
	out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
	out.CIDRAllocatorType = in.CIDRAllocatorType
	if err := v1.Convert_Pointer_bool_To_bool(&in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes, s); err != nil {
		return err
	}
	out.RootCAFile = in.RootCAFile
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VolumeConfiguration_To_componentconfig_VolumeConfiguration(&in.VolumeConfiguration, &out.VolumeConfiguration, s); err != nil {
		return err
	}
	out.ControllerStartInterval = in.ControllerStartInterval
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
		return err
	}
	out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
	out.GCIgnoredResources = *(*[]componentconfig.GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
	out.NodeEvictionRate = in.NodeEvictionRate
	out.SecondaryNodeEvictionRate = in.SecondaryNodeEvictionRate
	out.LargeClusterSizeThreshold = in.LargeClusterSizeThreshold
	out.UnhealthyZoneThreshold = in.UnhealthyZoneThreshold
	out.DisableAttachDetachReconcilerSync = in.DisableAttachDetachReconcilerSync
	out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableTaintManager, &out.EnableTaintManager, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.HorizontalPodAutoscalerUseRESTClients, &out.HorizontalPodAutoscalerUseRESTClients, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_KubeControllerManagerConfiguration_To_componentconfig_KubeControllerManagerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeControllerManagerConfiguration_To_componentconfig_KubeControllerManagerConfiguration(in *KubeControllerManagerConfiguration, out *componentconfig.KubeControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeControllerManagerConfiguration_To_componentconfig_KubeControllerManagerConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in *componentconfig.KubeControllerManagerConfiguration, out *KubeControllerManagerConfiguration, s conversion.Scope) error {
	out.Controllers = *(*[]string)(unsafe.Pointer(&in.Controllers))
	out.Port = in.Port
	out.Address = in.Address
	out.UseServiceAccountCredentials = in.UseServiceAccountCredentials
	out.CloudProvider = in.CloudProvider
	out.CloudConfigFile = in.CloudConfigFile
	out.ExternalCloudVolumePlugin = in.ExternalCloudVolumePlugin
	out.AllowUntaggedCloud = in.AllowUntaggedCloud
	out.ConcurrentEndpointSyncs = in.ConcurrentEndpointSyncs
	out.ConcurrentRSSyncs = in.ConcurrentRSSyncs
	out.ConcurrentRCSyncs = in.ConcurrentRCSyncs
	out.ConcurrentServiceSyncs = in.ConcurrentServiceSyncs
	out.ConcurrentResourceQuotaSyncs = in.ConcurrentResourceQuotaSyncs
	out.ConcurrentDeploymentSyncs = in.ConcurrentDeploymentSyncs
	out.ConcurrentDaemonSetSyncs = in.ConcurrentDaemonSetSyncs
	out.ConcurrentJobSyncs = in.ConcurrentJobSyncs
	out.ConcurrentNamespaceSyncs = in.ConcurrentNamespaceSyncs
	out.ConcurrentSATokenSyncs = in.ConcurrentSATokenSyncs
	out.NodeSyncPeriod = in.NodeSyncPeriod
	out.RouteReconciliationPeriod = in.RouteReconciliationPeriod
	out.ResourceQuotaSyncPeriod = in.ResourceQuotaSyncPeriod
	out.NamespaceSyncPeriod = in.NamespaceSyncPeriod
	out.PVClaimBinderSyncPeriod = in.PVClaimBinderSyncPeriod
	out.MinResyncPeriod = in.MinResyncPeriod
	out.TerminatedPodGCThreshold = in.TerminatedPodGCThreshold
	out.HorizontalPodAutoscalerSyncPeriod = in.HorizontalPodAutoscalerSyncPeriod
	out.HorizontalPodAutoscalerUpscaleForbiddenWindow = in.HorizontalPodAutoscalerUpscaleForbiddenWindow
	out.HorizontalPodAutoscalerDownscaleForbiddenWindow = in.HorizontalPodAutoscalerDownscaleForbiddenWindow
	out.HorizontalPodAutoscalerTolerance = in.HorizontalPodAutoscalerTolerance
	out.DeploymentControllerSyncPeriod = in.DeploymentControllerSyncPeriod
	out.PodEvictionTimeout = in.PodEvictionTimeout
	out.DeletingPodsQps = in.DeletingPodsQps
	out.DeletingPodsBurst = in.DeletingPodsBurst
	out.NodeMonitorGracePeriod = in.NodeMonitorGracePeriod
	out.RegisterRetryCount = in.RegisterRetryCount
	out.NodeStartupGracePeriod = in.NodeStartupGracePeriod
	out.NodeMonitorPeriod = in.NodeMonitorPeriod
	out.ServiceAccountKeyFile = in.ServiceAccountKeyFile
	out.ClusterSigningCertFile = in.ClusterSigningCertFile
	out.ClusterSigningKeyFile = in.ClusterSigningKeyFile
	out.ClusterSigningDuration = in.ClusterSigningDuration
	out.EnableProfiling = in.EnableProfiling
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.ClusterName = in.ClusterName
	out.ClusterCIDR = in.ClusterCIDR
	out.ServiceCIDR = in.ServiceCIDR
	out.NodeCIDRMaskSize = in.NodeCIDRMaskSize
	out.AllocateNodeCIDRs = in.AllocateNodeCIDRs
	out.CIDRAllocatorType = in.CIDRAllocatorType
	if err := v1.Convert_bool_To_Pointer_bool(&in.ConfigureCloudRoutes, &out.ConfigureCloudRoutes, s); err != nil {
		return err
	}
	out.RootCAFile = in.RootCAFile
	out.ContentType = in.ContentType
	out.KubeAPIQPS = in.KubeAPIQPS
	out.KubeAPIBurst = in.KubeAPIBurst
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(&in.VolumeConfiguration, &out.VolumeConfiguration, s); err != nil {
		return err
	}
	out.ControllerStartInterval = in.ControllerStartInterval
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableGarbageCollector, &out.EnableGarbageCollector, s); err != nil {
		return err
	}
	out.ConcurrentGCSyncs = in.ConcurrentGCSyncs
	out.GCIgnoredResources = *(*[]GroupResource)(unsafe.Pointer(&in.GCIgnoredResources))
	out.NodeEvictionRate = in.NodeEvictionRate
	out.SecondaryNodeEvictionRate = in.SecondaryNodeEvictionRate
	out.LargeClusterSizeThreshold = in.LargeClusterSizeThreshold
	out.UnhealthyZoneThreshold = in.UnhealthyZoneThreshold
	out.DisableAttachDetachReconcilerSync = in.DisableAttachDetachReconcilerSync
	out.ReconcilerSyncLoopPeriod = in.ReconcilerSyncLoopPeriod
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableTaintManager, &out.EnableTaintManager, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.HorizontalPodAutoscalerUseRESTClients, &out.HorizontalPodAutoscalerUseRESTClients, s); err != nil {
		return err
	}
	return nil
}

// Convert_componentconfig_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in *componentconfig.KubeControllerManagerConfiguration, out *KubeControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeControllerManagerConfiguration_To_v1alpha1_KubeControllerManagerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	out.SchedulerName = in.SchedulerName
	if err := Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	if err := Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.FailureDomains = in.FailureDomains
	out.CacheDebugResyncInterval = in.CacheDebugResyncInterval
	return nil
}

// Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	out.SchedulerName = in.SchedulerName
	if err := Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	if err := Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.FailureDomains = in.FailureDomains
	out.CacheDebugResyncInterval = in.CacheDebugResyncInterval
	return nil
}

// Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(in *KubeSchedulerLeaderElectionConfiguration, out *componentconfig.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(in *KubeSchedulerLeaderElectionConfiguration, out *componentconfig.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *componentconfig.KubeSchedulerLeaderElectionConfiguration, out *KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *componentconfig.KubeSchedulerLeaderElectionConfiguration, out *KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	return nil
}

// Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	return nil
}

// Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_componentconfig_PersistentVolumeRecyclerConfiguration(in *PersistentVolumeRecyclerConfiguration, out *componentconfig.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	out.MaximumRetry = in.MaximumRetry
	out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
	out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
	out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
	out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
	out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
	out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
	return nil
}

// Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_componentconfig_PersistentVolumeRecyclerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_componentconfig_PersistentVolumeRecyclerConfiguration(in *PersistentVolumeRecyclerConfiguration, out *componentconfig.PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_componentconfig_PersistentVolumeRecyclerConfiguration(in, out, s)
}

func autoConvert_componentconfig_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in *componentconfig.PersistentVolumeRecyclerConfiguration, out *PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	out.MaximumRetry = in.MaximumRetry
	out.MinimumTimeoutNFS = in.MinimumTimeoutNFS
	out.PodTemplateFilePathNFS = in.PodTemplateFilePathNFS
	out.IncrementTimeoutNFS = in.IncrementTimeoutNFS
	out.PodTemplateFilePathHostPath = in.PodTemplateFilePathHostPath
	out.MinimumTimeoutHostPath = in.MinimumTimeoutHostPath
	out.IncrementTimeoutHostPath = in.IncrementTimeoutHostPath
	return nil
}

// Convert_componentconfig_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in *componentconfig.PersistentVolumeRecyclerConfiguration, out *PersistentVolumeRecyclerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(in *SchedulerAlgorithmSource, out *componentconfig.SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*componentconfig.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(in *SchedulerAlgorithmSource, out *componentconfig.SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *componentconfig.SchedulerAlgorithmSource, out *SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *componentconfig.SchedulerAlgorithmSource, out *SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource(in *SchedulerPolicyConfigMapSource, out *componentconfig.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource(in *SchedulerPolicyConfigMapSource, out *componentconfig.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *componentconfig.SchedulerPolicyConfigMapSource, out *SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *componentconfig.SchedulerPolicyConfigMapSource, out *SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource(in *SchedulerPolicyFileSource, out *componentconfig.SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource(in *SchedulerPolicyFileSource, out *componentconfig.SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *componentconfig.SchedulerPolicyFileSource, out *SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *componentconfig.SchedulerPolicyFileSource, out *SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource(in *SchedulerPolicySource, out *componentconfig.SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*componentconfig.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*componentconfig.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource(in *SchedulerPolicySource, out *componentconfig.SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource(in, out, s)
}

func autoConvert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *componentconfig.SchedulerPolicySource, out *SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *componentconfig.SchedulerPolicySource, out *SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in, out, s)
}

func autoConvert_v1alpha1_VolumeConfiguration_To_componentconfig_VolumeConfiguration(in *VolumeConfiguration, out *componentconfig.VolumeConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableHostPathProvisioning, &out.EnableHostPathProvisioning, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_bool_To_bool(&in.EnableDynamicProvisioning, &out.EnableDynamicProvisioning, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PersistentVolumeRecyclerConfiguration_To_componentconfig_PersistentVolumeRecyclerConfiguration(&in.PersistentVolumeRecyclerConfiguration, &out.PersistentVolumeRecyclerConfiguration, s); err != nil {
		return err
	}
	out.FlexVolumePluginDir = in.FlexVolumePluginDir
	return nil
}

// Convert_v1alpha1_VolumeConfiguration_To_componentconfig_VolumeConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_VolumeConfiguration_To_componentconfig_VolumeConfiguration(in *VolumeConfiguration, out *componentconfig.VolumeConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_VolumeConfiguration_To_componentconfig_VolumeConfiguration(in, out, s)
}

func autoConvert_componentconfig_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in *componentconfig.VolumeConfiguration, out *VolumeConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableHostPathProvisioning, &out.EnableHostPathProvisioning, s); err != nil {
		return err
	}
	if err := v1.Convert_bool_To_Pointer_bool(&in.EnableDynamicProvisioning, &out.EnableDynamicProvisioning, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_PersistentVolumeRecyclerConfiguration_To_v1alpha1_PersistentVolumeRecyclerConfiguration(&in.PersistentVolumeRecyclerConfiguration, &out.PersistentVolumeRecyclerConfiguration, s); err != nil {
		return err
	}
	out.FlexVolumePluginDir = in.FlexVolumePluginDir
	return nil
}

// Convert_componentconfig_VolumeConfiguration_To_v1alpha1_VolumeConfiguration is an autogenerated conversion function.
func Convert_componentconfig_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in *componentconfig.VolumeConfiguration, out *VolumeConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_VolumeConfiguration_To_v1alpha1_VolumeConfiguration(in, out, s)
}
