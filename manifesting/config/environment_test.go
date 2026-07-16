package config

import (
	"testing"

	"github.com/estratocloud/manifesting/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// GetOutputPath Ensure we can read an output file path
func Test_GetOutputPath1(t *testing.T) {

	wd, err := internal.NewWorkingDirectory("/tmp")
	require.NoError(t, err)

	var conf Config
	data := []byte(`environments:
  - name: "production"
    output: ".generated/production.yaml"`)
	err = yaml.Unmarshal(data, &conf)
	require.NoError(t, err)

	got := conf.Environments[0].GetOutputPath(wd).GetFullyQualifiedPath()
	assert.Equal(t, "/tmp/.generated/production.yaml", got)
}

// GetOutputPath Ensure we can use a default output file path
func Test_GetOutputPath2(t *testing.T) {

	wd, err := internal.NewWorkingDirectory("/tmp")
	require.NoError(t, err)

	var conf Config
	data := []byte(`environments:
  - name: "nonprod"`)
	err = yaml.Unmarshal(data, &conf)
	require.NoError(t, err)

	got := conf.Environments[0].GetOutputPath(wd).GetFullyQualifiedPath()
	assert.Equal(t, "/tmp/.generated/nonprod.yaml", got)
}
