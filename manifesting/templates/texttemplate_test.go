package templates

import (
	"bytes"
	"testing"

	"github.com/estratocloud/manifesting/manifesting/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// getTextTemplate Ensure we can render a template with variables
func Test_getTextTemplate1(t *testing.T) {
	tmpl, err := getTextTemplate([]byte("Hello {{ .Name }}!"), nil)
	require.NoError(t, err)

	var got bytes.Buffer
	err = tmpl.Execute(&got, map[string]string{
		"Name": "Craig",
	})
	require.NoError(t, err)

	assert.Equal(t, "Hello Craig!", got.String())
}

// getTextTemplate Ensure we can use the sprout functions
func Test_getTextTemplate2(t *testing.T) {
	tmpl, err := getTextTemplate([]byte("Lowest = {{ min .One .Two }}"), nil)
	require.NoError(t, err)

	var got bytes.Buffer
	err = tmpl.Execute(&got, map[string]int{
		"One": 9,
		"Two": 4,
	})
	require.NoError(t, err)

	assert.Equal(t, "Lowest = 4", got.String())
}

// getTextTemplate Ensure errors are returned
func Test_getTextTemplate3(t *testing.T) {
	tmpl, err := getTextTemplate([]byte("Invalid {{ unknown .Bad }}"), nil)
	assert.Nil(t, tmpl)
	assert.EqualError(t, err, `template: template:1: function "unknown" not defined`)
}

// getTextTemplate Ensure we can use the perEnvironment function
func Test_getTextTemplate4(t *testing.T) {
	environemnt := &config.Environment{Name: "nonprod"}
	tmpl, err := getTextTemplate([]byte("debug = {{ perEnvironment .Debug }}"), environemnt)
	require.NoError(t, err)

	var got bytes.Buffer
	err = tmpl.Execute(&got, map[string]any{
		"Debug": map[string]int{
			"production": 0,
			"nonprod":    1,
		},
	})
	require.NoError(t, err)

	assert.Equal(t, "debug = 1", got.String())
}
