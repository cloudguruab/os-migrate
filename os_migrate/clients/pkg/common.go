package pkg

import "encoding/json"
import "fmt"
import "os"
import "runtime/debug"

func returnResponse(responseBody ModuleResult) {
	var response []byte
	var err error
	response, err = json.Marshal(responseBody)
	if err != nil {
		response, _ = json.Marshal(ModuleResult{Data: map[string]interface{}{"response": "Invalid json object"}})
		responseBody.Error = err
	}
	fmt.Println(string(response))
	if responseBody.Error != nil {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

func NewModule() *OpenStackModule {
	// Set the module 
	module := &OpenStackModule{
		Results: ModuleResult{Changed: false},
		ArgumentSpec: make(map[string]ArgumentSpec),
		ModuleName: "Ansible OpenStack Module (Gophercloud)",
	}

	// Get the SDK version
	info, _ := debug.ReadBuildInfo()
	for _, dep := range info.Deps {
		if dep.Path == "github.com/gophercloud/gophercloud" {
			module.SDKVersion = dep.Version
		}
	}
	return module
}

func ExitJson(responseBody ModuleResult) {
	returnResponse(responseBody)
}

func FailJson(responseBody ModuleResult) {
	responseBody.Changed = true
	returnResponse(responseBody)
}

func (m *OpenStackModule) Log(msg string, level string) {
	// Default to INFO if level not specified
	if level == "" {
		level = m.ArgumentSpec["sdk_log_level"].Default.(string)
	}

	switch level {
	case "DEBUG":
		FailJson(ModuleResult{Data: map[string]interface{}{fmt.Sprintf("[%s]", level): msg}})
	case "INFO":
		ExitJson(ModuleResult{Data: map[string]interface{}{fmt.Sprintf("[%s]", level): msg}})
	case "WARNING":
		ExitJson(ModuleResult{Data: map[string]interface{}{fmt.Sprintf("[%s]", level): msg}})
	case "ERROR":
		FailJson(ModuleResult{Data: map[string]interface{}{fmt.Sprintf("[%s]", level): msg}})
	default:
		ExitJson(ModuleResult{Data: map[string]interface{}{fmt.Sprintf("[%s]", level): msg}})
	}
}

// Sets up a gophercloud ProviderClient from module arguments
func (m *OpenStackModule) OpenStackCloudFromModule() error {
	return nil
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
			Aliases: []string{"verify"},
		},
		CACert: ArgumentSpec{
			Aliases: []string{"cacert"},
		},
		ClientCert: ArgumentSpec{
			Aliases: []string{"cert"},
		},
		ClientKey: ArgumentSpec{
			NoLog:   true,
			Aliases: []string{"key"},
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
