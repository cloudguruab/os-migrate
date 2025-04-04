package openstack_test

import (
	"testing"
	"fmt"
	"os-migrate/os_migrate/plugins/module_utils/test_helpers"
)

func TestOpenStackAuth(t *testing.T) {
	module, providerClient, ctx, cancel := test_helpers.GetTestProviderClient(t)
	defer cancel()

	t.Run("Test OpenStackAuth", func(t *testing.T) {
		fmt.Printf("ProviderClient: %+v\n", providerClient.Token())
		fmt.Printf("Module: %+v\n", module)
		fmt.Printf("ProviderClient: %+v\n", ctx)
		fmt.Printf("Cancel: %+v\n", cancel)
	})
}

