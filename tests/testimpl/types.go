package testimpl

import "github.com/launchbynttdata/lcaf-component-terratest/types"

// ThisTFModuleConfig represents the configuration for the AWS Service Discovery Service module
type ThisTFModuleConfig struct {
	types.GenericTFModuleConfig
}

// ServiceDiscoveryServiceConfig represents the expected configuration for a Service Discovery Service
type ServiceDiscoveryServiceConfig struct {
	Name                        string
	NamespaceID                 string
	DNSRecordType               string
	TTL                         int64
	RoutingPolicy               string
	HealthCheckFailureThreshold int32
	Tags                        map[string]string
}
