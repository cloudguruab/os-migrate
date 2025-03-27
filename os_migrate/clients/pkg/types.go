package pkg

// ModuleRunner defines the interface that all OpenStack modules must implement
type ModuleRunner interface {
	Run() error
}

// CloudConfig represents the configuration for connecting to an OpenStack cloud
type CloudConfig struct {
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

// ModuleResult represents the result of a module execution
type ModuleResult struct {
	Changed bool                   `json:"changed"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   error                  `json:"error,omitempty"`
}

// ArgumentSpec represents the specification for a module argument
type ArgumentSpec struct {
	Type     string        `json:"type,omitempty"`
	Default  interface{}   `json:"default,omitempty"`
	Required bool          `json:"required,omitempty"`
	Choices  []string      `json:"choices,omitempty"`
	Aliases  []string      `json:"aliases,omitempty"`
	NoLog    bool          `json:"no_log,omitempty"`
	MinVer   string        `json:"min_ver,omitempty"`
	MaxVer   string        `json:"max_ver,omitempty"`
} 
