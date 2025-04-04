package keystone

import (
	"testing"
	"os-migrate/os_migrate/plugins/module_utils/test_helpers"
)

func TestListProjects(t *testing.T) {
	_, providerClient, ctx, cancel := test_helpers.GetTestProviderClient(t)
	defer cancel()

	project := ProjectOpts{
		ID: "123",
		Name: "test",
		DomainID: "123",
		IsDomain: false,
		Enabled: true,
	}

	changed, err := ListProjects(providerClient, project)
	if err != nil {
		t.Errorf("ListProjects failed: %v", err)
	}
	if !changed {
		t.Error("Expected changed to be true")
	}
} 