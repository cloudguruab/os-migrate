package main

import (
	"fmt"
	"os"
	"clients/pkg"
)

// TestModule is a concrete implementation of OpenStackModule for testing
type TestModule struct {
	*pkg.OpenStackModule
}

// Run implements the ModuleRunner interface
func (m *TestModule) Run() error {
	// Test different log levels
	m.Log("This is an INFO message", "INFO")
	m.Log("This is a DEBUG message", "DEBUG")
	m.Log("This is a WARNING message", "WARNING")
	m.Log("This is an ERROR message", "ERROR")
	
	// Test default level (should be INFO)
	m.Log("This is a message with no level specified", "")

	return nil
}

// Execute overrides the base module's Execute method
func (m *TestModule) Execute() error {
	if err := m.OpenStackCloudFromModule(); err != nil {
		return fmt.Errorf("failed to setup cloud connection: %v", err)
	}

	if err := m.Run(); err != nil {
		return fmt.Errorf("module execution failed: %v", err)
	}

	return nil
}

func main() {

	// Create a new test module with initialized base module
	module := &TestModule{
		OpenStackModule: pkg.NewModule(),
	}
	module.ModuleName = "test_module"
		
	// Set some test parameters using the CloudConfig type
	module.Params = pkg.CloudConfig{
		AuthType: "password",
		Auth: map[string]interface{}{
			"auth_url":         "http://192.168.122.45/identity",
			"username":        "demo",
			"password":        "password",
			"project_name":    "demo",
			"project_domain_id": "default",
			"user_domain_id":   "default",
		},
		RegionName: "RegionOne",
		Interface:  "public",
		Verify:     true,
	}

	// Execute the module
	if err := module.Execute(); err != nil {
		fmt.Printf("Error executing module: %v\n", err)
		os.Exit(1)
	}
} 