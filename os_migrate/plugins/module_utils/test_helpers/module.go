package test_helpers

import (
	"context"
	"os-migrate/os_migrate/plugins/module_utils/openstack"
	"github.com/gophercloud/gophercloud/v2"
)

// TestModule is a concrete implementation of OpenStackModule for testing
type TestModule struct {
	*openstack.OpenStackModule
}

// NewTestModule creates a new test module instance with default auth configuration
func NewTestModule(cloud string) *TestModule {
	baseModule := openstack.NewModule(cloud)
	
	module := &TestModule{
		OpenStackModule: baseModule,
	}
	
	// Set up default test values in the argument spec
	module.ArgumentSpec.Auth.Default = map[string]interface{}{
		"auth_url": "http://192.168.122.45/identity",
		"username": "demo",
		"password": "password",
		"project_name": "demo",
		"project_domain_id": "default",
		"user_domain_id": "default",
	}
	
	return module
}

// GetTestProviderClient creates a new context and provider client for testing
func GetTestProviderClient(t interface{}) (*TestModule, *gophercloud.ProviderClient, context.Context, context.CancelFunc) {
	module := NewTestModule("src")
	ctx, cancel := context.WithCancel(context.Background())
	
	providerClient, err := module.OpenStackAuth(ctx)
	if err != nil {
		panic(err) // In tests, we want to fail fast
	}
	
	return module, providerClient, ctx, cancel
} 