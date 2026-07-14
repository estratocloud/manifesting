package config

import (
	"testing"

	"errors"

	"github.com/estratocloud/manifesting/internal"
	"github.com/estratocloud/manifesting/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	corev1 "k8s.io/api/core/v1"
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

// GetEnvVars Ensure we get an empty map if there is no envFrom defined
func Test_GetEnvVars1(t *testing.T) {

	wd, err := internal.NewWorkingDirectory("/app/tests/samples")
	require.NoError(t, err)

	var conf Config
	data := []byte(`environments:
  - name: "nonprod"
    envFrom: "envvars.yaml"`)
	err = yaml.Unmarshal(data, &conf)
	require.NoError(t, err)

	got, err := conf.Environments[0].GetEnvVars(wd)
	require.NoError(t, err)
	assert.Equal(t, map[string]corev1.EnvVar{
		"CONFIG_KEY_1": {Name: "CONFIG_KEY_1", Value: "value1"},
		"CONFIG_KEY_2": {Name: "CONFIG_KEY_2", Value: "value2"},
	}, got)
}

// GetEnvVars Ensure we get an empty map if there is no envFrom defined
func Test_GetEnvVars2(t *testing.T) {

	wd, err := internal.NewWorkingDirectory("/tmp")
	require.NoError(t, err)

	var conf Config
	data := []byte(`environments:
  - name: "nonprod"`)
	err = yaml.Unmarshal(data, &conf)
	require.NoError(t, err)

	got, err := conf.Environments[0].GetEnvVars(wd)
	require.NoError(t, err)
	assert.Equal(t, map[string]corev1.EnvVar{}, got)
}

// GetEnvVars Ensure we get an error if the declared file doesn't exist
func Test_GetEnvVars3(t *testing.T) {

	wd, err := internal.NewWorkingDirectory("/tmp")
	require.NoError(t, err)

	var conf Config
	data := []byte(`environments:
  - name: "nonprod"
    envFrom: "does-not-exist.yaml"`)
	err = yaml.Unmarshal(data, &conf)
	require.NoError(t, err)

	got, err := conf.Environments[0].GetEnvVars(wd)
	assert.Nil(t, got)
	assert.EqualError(t, err, "unable to find the envFrom file for nonprod at '/tmp/does-not-exist.yaml'")
}

// GetEnvVars Ensure we pass along errors from ReadFile
func Test_GetEnvVars4(t *testing.T) {

	wd := &mocks.WorkingDirectory{
		NewPathFunc: func(path string) internal.PathInterface {
			return &mocks.Path{
				Path: path,
				ReadFileFunc: func() ([]byte, error) {
					return nil, errors.New("cannot read")
				},
			}
		},
	}

	var conf Config
	data := []byte(`environments:
  - name: "nonprod"
    envFrom: "file.yaml"`)
	err := yaml.Unmarshal(data, &conf)
	require.NoError(t, err)

	got, err := conf.Environments[0].GetEnvVars(wd)
	assert.Nil(t, got)
	assert.EqualError(t, err, "unable to read the envFrom file for nonprod at 'file.yaml': cannot read")
}

// GetEnvVars Ensure we return an error if we can't parse the envvars file
func Test_GetEnvVars5(t *testing.T) {

	wd, err := internal.NewWorkingDirectory("/app/tests/samples")
	require.NoError(t, err)

	var conf Config
	data := []byte(`environments:
  - name: "nonprod"
    envFrom: "badsyntax.yaml"`)
	err = yaml.Unmarshal(data, &conf)
	require.NoError(t, err)

	got, err := conf.Environments[0].GetEnvVars(wd)
	assert.Nil(t, got)
	assert.EqualError(t, err, "unable to parse the envFrom file for nonprod at '/app/tests/samples/badsyntax.yaml': error converting YAML to JSON: yaml: line 3: found unexpected end of stream")
}
