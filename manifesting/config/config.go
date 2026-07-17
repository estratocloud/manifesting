package config

import (
	"slices"
)

type Config struct {
	Environments []*Environment
	Vars         map[string]any `yaml:"vars"`
	Resources    []*Resource    `yaml:"resources"`
}

func (c *Config) GetResources(environment *Environment) []*Resource {

	resources := make([]*Resource, 0, len(c.Resources))

	for _, resource := range c.Resources {
		if len(resource.Environments) > 0 {
			if !slices.Contains(resource.Environments, environment.Name) {
				continue
			}
		}

		resources = append(resources, resource)
	}

	return resources
}
