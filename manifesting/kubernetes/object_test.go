package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
)

// GetObject Ensure we can get a deployment struct
func Test_GetObject1(t *testing.T) {
	input := []byte("apiVersion: apps/v1\nkind: Deployment")
	got, err := GetObject(input)
	require.NoError(t, err)
	assert.IsType(t, &appsv1.Deployment{}, got)
}

// GetObject Ensure we get an error for a bad manifest
func Test_GetObject2(t *testing.T) {
	input := []byte("apiVersion: apps/v1")
	got, err := GetObject(input)
	assert.Nil(t, got)
	assert.EqualError(t, err, "Object 'Kind' is missing in 'apiVersion: apps/v1'")
}
