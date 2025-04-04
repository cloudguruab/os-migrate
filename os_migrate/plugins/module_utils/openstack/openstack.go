package openstack


import "context"
import "github.com/gophercloud/gophercloud/v2/openstack"
import "github.com/gophercloud/gophercloud/v2"


type ArgumentSpec struct {
	// ArgumentSpec represents the specification for an argument
	Type        string        `json:"type,omitempty"`
	Default     interface{}   `json:"default,omitempty"`
	Required    bool          `json:"required,omitempty"`
	Choices     []string      `json:"choices,omitempty"`
	Aliases     []string      `json:"aliases,omitempty"`
	NoLog       bool          `json:"no_log,omitempty"`
}

type ModuleResult struct {
	// ModuleResult represents the result of a module execution
	Changed bool                   `json:"changed"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   error                  `json:"error,omitempty"`
}

type OpenStackArgumentSpec struct {
	Cloud           ArgumentSpec `json:"cloud"`
	AuthType        ArgumentSpec `json:"auth_type"`
	Auth            ArgumentSpec `json:"auth"`
	RegionName      ArgumentSpec `json:"region_name"`
	ValidateCerts   ArgumentSpec `json:"validate_certs"`
	CACert          ArgumentSpec `json:"ca_cert"`
	ClientCert      ArgumentSpec `json:"client_cert"`
	ClientKey       ArgumentSpec `json:"client_key"`
	Wait            ArgumentSpec `json:"wait"`
	Timeout         ArgumentSpec `json:"timeout"`
	APITimeout      ArgumentSpec `json:"api_timeout"`
	Interface       ArgumentSpec `json:"interface"`
	SDKLogPath      ArgumentSpec `json:"sdk_log_path"`
	SDKLogLevel     ArgumentSpec `json:"sdk_log_level"`
	Path            ArgumentSpec `json:"path"`
	Name            ArgumentSpec `json:"name"`
}

type OpenStackModule struct {
	ModuleName       string 
	Results          ModuleResult
	Conn             *gophercloud.ProviderClient
	ArgumentSpec     OpenStackArgumentSpec
}

func NewModule(cloud string) *OpenStackModule {
	// Set the module 
	module := &OpenStackModule{
		Results: ModuleResult{Changed: false},
		ArgumentSpec: GetOpenStackArgumentSpec(),
		ModuleName: cloud,
	}

	return module
}

// Sets up a gophercloud ProviderClient from module arguments
func (m *OpenStackModule) OpenStackAuth(ctx context.Context) (*gophercloud.ProviderClient, error) {

	// Get auth options from the module's argument spec
	authMap := m.ArgumentSpec.Auth.Default.(map[string]interface{})
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: authMap["auth_url"].(string),
		Username:         authMap["username"].(string),
		Password:         authMap["password"].(string),
		TenantName:       authMap["project_name"].(string),
		DomainID:         authMap["user_domain_id"].(string),
	}

    // Create the provider client
	providerClient, err := openstack.AuthenticatedClient(context.TODO(), authOpts)
	if err != nil {
		return nil, err
	}

	return providerClient, nil
}

// NOTE: not needed, helpful for tracking module arguments
func GetOpenStackArgumentSpec() OpenStackArgumentSpec {
	return OpenStackArgumentSpec{
		Cloud: ArgumentSpec{
			Type: "raw",
		},
		AuthType: ArgumentSpec{},
		Auth: ArgumentSpec{
			Type:  "dict",
			NoLog: true,
		},
		RegionName: ArgumentSpec{},
		ValidateCerts: ArgumentSpec{
			Type:    "bool",
			Aliases: []string{"validate_certs"},
		},
		CACert: ArgumentSpec{
			Aliases: []string{"ca_cert"},
		},
		ClientCert: ArgumentSpec{
			Aliases: []string{"cert"},
		},
		ClientKey: ArgumentSpec{
			NoLog:   true,
			Aliases: []string{"client_key"},
		},
		Wait: ArgumentSpec{
			Default: true,
			Type:    "bool",
		},
		Timeout: ArgumentSpec{
			Default: 180,
			Type:    "int",
		},
		APITimeout: ArgumentSpec{
			Type: "int",
		},
		Interface: ArgumentSpec{
			Default: "public",
			Choices: []string{"public", "internal", "admin"},
			Aliases: []string{"endpoint_type"},
		},
		SDKLogPath: ArgumentSpec{},
		SDKLogLevel: ArgumentSpec{
			Default: "INFO",
			Choices: []string{"INFO", "DEBUG"},
		},
		Path: ArgumentSpec{
			Type:     "str",
			Required: true,
		},
		Name: ArgumentSpec{
			Type:     "str",
			Required: true,
		},
	}
}
