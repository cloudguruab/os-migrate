package pkg

// CustomVarParams defines parameters that are custom variables
var CustomVarParams = []string{"min_ver", "max_ver"}

// Overrides maps deprecated module names to their new names
var Overrides = make(map[string]string)

// GetOpenStackFullArgumentSpec returns the full OpenStack argument specification
func GetOpenStackFullArgumentSpec(kwargs map[string]interface{}) map[string]interface{} {
	spec := map[string]interface{}{
		"cloud": map[string]interface{}{
			"type": "raw",
		},
		"auth_type": map[string]interface{}{},
		"auth": map[string]interface{}{
			"type":    "dict",
			"no_log": true,
		},
		"region_name": map[string]interface{}{},
		"validate_certs": map[string]interface{}{
			"type":    "bool",
			"aliases": []string{"verify"},
		},
		"ca_cert": map[string]interface{}{
			"aliases": []string{"cacert"},
		},
		"client_cert": map[string]interface{}{
			"aliases": []string{"cert"},
		},
		"client_key": map[string]interface{}{
			"no_log": true,
			"aliases": []string{"key"},
		},
		"wait": map[string]interface{}{
			"default": true,
			"type":    "bool",
		},
		"timeout": map[string]interface{}{
			"default": 180,
			"type":    "int",
		},
		"api_timeout": map[string]interface{}{
			"type": "int",
		},
		"interface": map[string]interface{}{
			"default": "public",
			"choices": []string{"public", "internal", "admin"},
			"aliases": []string{"endpoint_type"},
		},
		"sdk_log_path": map[string]interface{}{},
		"sdk_log_level": map[string]interface{}{
			"default": "INFO",
			"choices": []string{"INFO", "DEBUG"},
		},
	}

	// Filter out custom parameters
	kwargsCopy := make(map[string]interface{})
	for k, v := range kwargs {
		kwargsCopy[k] = v
	}

	for _, param := range CustomVarParams {
		for _, v := range kwargsCopy {
			if m, ok := v.(map[string]interface{}); ok {
				delete(m, param)
			}
		}
	}

	// Update spec with kwargs
	for k, v := range kwargsCopy {
		spec[k] = v
	}

	return spec
}

// GetOpenStackModuleKwargs returns module kwargs
func GetOpenStackModuleKwargs(kwargs map[string]interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	for _, key := range []string{"mutually_exclusive", "required_together", "required_one_of"} {
		if v, ok := kwargs[key]; ok {
			if existing, exists := ret[key]; exists {
				if existingSlice, ok := existing.([]interface{}); ok {
					if newSlice, ok := v.([]interface{}); ok {
						ret[key] = append(existingSlice, newSlice...)
					}
				}
			} else {
				ret[key] = v
			}
		}
	}
	return ret
}
