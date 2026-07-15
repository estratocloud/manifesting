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
