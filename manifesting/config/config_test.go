package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// Unmarshal Ensure we can parse a full config file
func Test_Unmarshal1(t *testing.T) {

	var got Config
	data, err := os.ReadFile("/app/tests/samples/config/full.yaml")
	require.NoError(t, err)
	err = yaml.Unmarshal(data, &got)
	require.NoError(t, err)

	assert.Equal(t, Config{
		Environments: []*Environment{
			{
				Name:    "production",
				Output:  ".generated/production.yaml",
				EnvFrom: "envvars/production.yaml",
			},
			{
				Name:    "nonprod",
				Output:  ".generated/nonprod.yaml",
				EnvFrom: "envvars/nonprod.yaml",
			},
		},
		Vars: map[string]any{
			"Priority": "first",
			"CPU":      50,
		},
		Resources: []*Resource{
			{
				Name:         "awesome-daemon",
				Template:     "daemon",
				Environments: []string{"production"},
				Vars: map[string]any{
					"Program": "awesome.php",
				},
			},
			{
				Name:         "boring-daemon",
				Template:     "daemon",
				Environments: []string{"nonprod"},
				Vars: map[string]any{
					"Program": "boring.py",
				},
			},
		},
	}, got)
}

// GetResources Ensure we can get all the resources
func Test_GetResources1(t *testing.T) {

	conf := &Config{
		Resources: []*Resource{
			{Name: "myapp"},
			{Name: "yourapp"},
			{Name: "theirapp"},
		},
	}

	got := conf.GetResources(&Environment{})

	assert.Equal(t, []*Resource{
		{Name: "myapp"},
		{Name: "yourapp"},
		{Name: "theirapp"},
	}, got)
}

// GetResources Ensure we can get filter the resources by environment
func Test_GetResources2(t *testing.T) {

	conf := &Config{
		Resources: []*Resource{
			{Name: "myapp"},
			{Name: "yourapp", Environments: []string{"uk"}},
			{Name: "theirapp", Environments: []string{"canada"}},
		},
	}

	resources := conf.GetResources(&Environment{Name: "canada"})
	var got []string
	for _, resource := range resources {
		got = append(got, resource.Name)
	}

	assert.Equal(t, []string{
		"myapp",
		"theirapp",
	}, got)
}
