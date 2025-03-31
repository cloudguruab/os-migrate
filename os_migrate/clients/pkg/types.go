package pkg

import "github.com/gophercloud/gophercloud"


// ============== GENERICS ==============
type CloudConfig struct {
	// CloudConfig represents the configuration for connecting to an OpenStack cloud
	Cloud       string                 `json:"cloud,omitempty"`
	AuthType    string                 `json:"auth_type,omitempty"`
	Auth        map[string]interface{} `json:"auth,omitempty"`
	RegionName  string                 `json:"region_name,omitempty"`
	Verify      bool                   `json:"verify,omitempty"`
	CACert      string                 `json:"ca_cert,omitempty"`
	Key         string                 `json:"key,omitempty"`
	Cert        string                 `json:"cert,omitempty"`
	APITimeout  int                    `json:"api_timeout,omitempty"`
	Interface   string                 `json:"interface,omitempty"`
}

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


// ============== OPENSTACK =============
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
	Params           CloudConfig
	ModuleName       string 
	SDKVersion       string
	Results          ModuleResult
	Conn             *gophercloud.ProviderClient
	ArgumentSpec     map[string]ArgumentSpec
}


// ============== KEYSTONE ==============
type IdentityMiddleware interface {
	Authenticate() error
	GetToken() string
	IsTokenValid() bool
	RefreshToken() error
	OpenStackCloudFromModule() error
}

type KeystoneStruct struct {
	AuthURL      string `json:"auth_url"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	TenantName   string `json:"tenant_name"`
	RegionName   string `json:"region_name"`
	Token        string `json:"token"`
	ExpiresAt    string `json:"expires_at"`
}