/*
Copyright 2015 The Kubernetes Authors.

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

package proxy

import (
	"fmt"
	"net"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/proxy/config"
)

// Provider is the interface provided by proxier implementations.
type Provider interface {
	config.EndpointsHandler
	config.EndpointSliceHandler
	config.ServiceHandler
	config.NodeHandler

	// Sync immediately synchronizes the Provider's current state to proxy rules.
	Sync()
	// SyncLoop runs periodic work.
	// This is expected to run as a goroutine or as the main loop of the app.
	// It does not return.
	SyncLoop()
}

// ServicePortName carries a namespace + name + portname.  This is the unique
// identifier for a load-balanced service.
type ServicePortName struct {
	types.NamespacedName
	Port     string
	Protocol v1.Protocol
}

func (spn ServicePortName) String() string {
	return fmt.Sprintf("%s%s", spn.NamespacedName.String(), fmtPortName(spn.Port))
}

func fmtPortName(in string) string {
	if in == "" {
		return ""
	}
	return fmt.Sprintf(":%s", in)
}

// ServicePort is an interface which abstracts information about a service.
type ServicePort interface {
	// String returns service string.  An example format can be: `IP:Port/Protocol`.
	String() string
	// GetClusterIP returns service cluster IP in net.IP format.
	ClusterIP() net.IP
	// GetPort returns service port if present. If return 0 means not present.
	Port() int
	// GetSessionAffinityType returns service session affinity type
	SessionAffinityType() v1.ServiceAffinity
	// GetStickyMaxAgeSeconds returns service max connection age
	StickyMaxAgeSeconds() int
	// ExternalIPStrings returns service ExternalIPs as a string array.
	ExternalIPStrings() []string
	// LoadBalancerIPStrings returns service LoadBalancerIPs as a string array.
	LoadBalancerIPStrings() []string
	// GetProtocol returns service protocol.
	Protocol() v1.Protocol
	// LoadBalancerSourceRanges returns service LoadBalancerSourceRanges if present empty array if not
	LoadBalancerSourceRanges() []string
	// GetHealthCheckNodePort returns service health check node port if present.  If return 0, it means not present.
	HealthCheckNodePort() int
	// GetNodePort returns a service Node port if present. If return 0, it means not present.
	NodePort() int
	// GetOnlyNodeLocalEndpoints returns if a service has only node local endpoints
	OnlyNodeLocalEndpoints() bool
	// TopologyKeys returns service TopologyKeys as a string array.
	TopologyKeys() []string
}

// Endpoint in an interface which abstracts information about an endpoint.
// TODO: Rename functions to be consistent with ServicePort.
type Endpoint interface {
	// String returns endpoint string.  An example format can be: `IP:Port`.
	// We take the returned value as ServiceEndpoint.Endpoint.
	String() string
	// IsReady returns true if the endpoint came from the NotReadyAddresses list in Endpoints
	// or any endpoints in EndpointSlice with Ready=false
	IsReady() bool
	// IsTerminating returns true if an endpoint is terminating
	IsTerminating() bool
	// GetIsLocal returns true if the endpoint is running in same host as kube-proxy, otherwise returns false.
	GetIsLocal() bool
	// GetTopology returns the topology information of the endpoint.
	GetTopology() map[string]string
	// IP returns IP part of the endpoint.
	IP() string
	// Port returns the Port part of the endpoint.
	Port() (int, error)
	// Equal checks if two endpoints are equal.
	Equal(Endpoint) bool
}

// ServiceEndpoint is used to identify a service and one of its endpoint pair.
type ServiceEndpoint struct {
	Endpoint        string
	ServicePortName ServicePortName
}
