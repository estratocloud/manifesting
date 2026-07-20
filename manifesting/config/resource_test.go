package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// GetVars Ensure we can get a simple resource vars map
func Test_GetVars1(t *testing.T) {

	conf := &Config{}
	environment := &Environment{Name: "jupiter"}
	resource := &Resource{
		Name: "webapp",
		Vars: map[string]any{
			"One": 1,
			"Two": "2",
		},
	}

	got, err := resource.GetVars(environment, conf)
	require.NoError(t, err)
	assert.Equal(t, map[string]any{
		"One":           1,
		"Two":           "2",
		"ENVIRONMENT":   "jupiter",
		"RESOURCE_NAME": "webapp",
	}, got)
}

// GetVars Ensure we can use default vars from the config file
func Test_GetVars2(t *testing.T) {

	conf := &Config{
		Vars: map[string]any{
			"One": 1,
			"Two": "2",
		},
	}
	environment := &Environment{Name: "saturn"}
	resource := &Resource{
		Name: "webapp",
		Vars: map[string]any{
			"Two":   "222",
			"Three": "3",
		},
	}

	got, err := resource.GetVars(environment, conf)
	require.NoError(t, err)
	assert.Equal(t, map[string]any{
		"One":           1,
		"Two":           "222",
		"Three":         "3",
		"ENVIRONMENT":   "saturn",
		"RESOURCE_NAME": "webapp",
	}, got)
}

// GetVars Ensure empty structs are handled
func Test_GetVars3(t *testing.T) {

	conf := &Config{}
	environment := &Environment{}
	resource := &Resource{}

	got, err := resource.GetVars(environment, conf)
	require.NoError(t, err)
	assert.Equal(t, map[string]any{
		"ENVIRONMENT":   "",
		"RESOURCE_NAME": "",
	}, got)
}

// GetVars Ensure we get an error if we try to use a reserved name (in resource vars)
func Test_GetVars4(t *testing.T) {

	conf := &Config{}
	environment := &Environment{Name: "mars"}
	resource := &Resource{
		Name: "webapp",
		Vars: map[string]any{
			"ENVIRONMENT": "override",
		},
	}

	got, err := resource.GetVars(environment, conf)
	assert.Nil(t, got)
	assert.EqualError(t, err, "invalid variable (ENVIRONMENT) defined, uppercase variable names are reserved for Manifesting variables")
}

// GetVars Ensure we get an error if we try to use a reserved name (in config vars)
func Test_GetVars5(t *testing.T) {

	conf := &Config{
		Vars: map[string]any{
			"RESOURCE_NAME": "override",
		},
	}
	environment := &Environment{Name: "pluto"}
	resource := &Resource{
		Name: "webapp",
	}

	got, err := resource.GetVars(environment, conf)
	assert.Nil(t, got)
	assert.EqualError(t, err, "invalid variable (RESOURCE_NAME) defined, uppercase variable names are reserved for Manifesting variables")
}
