package config

import (
	"fmt"
	"maps"
)

type Resource struct {
	Name         string         `yaml:"name"`
	Template     string         `yaml:"template"`
	Environments []string       `yaml:"environments"`
	Vars         map[string]any `yaml:"vars"`
}

func (r *Resource) GetVars(environment *Environment, config *Config) (map[string]any, error) {

	reserved := map[string]any{
		"RESOURCE_NAME": r.Name,
		"ENVIRONMENT":   environment.Name,
	}

	// Create a new map, allocating enough memory for all the expected keys
	vars := make(map[string]any, len(config.Vars)+len(r.Vars)+len(reserved))

	// Populate any default vars from the main config file
	maps.Copy(vars, config.Vars)

	// Override the defaults with any specific to this resource
	maps.Copy(vars, r.Vars)

	for key, value := range reserved {
		_, ok := vars[key]
		if ok {
			return nil, fmt.Errorf("invalid variable (%s) defined, uppercase variable names are reserved for Manifesting variables", key)
		}
		vars[key] = value
	}

	return vars, nil
}
