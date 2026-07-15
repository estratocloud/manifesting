package internal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// NewWorkingDirectory Ensure an absolute path is used without modification
func Test_NewWorkingDirectory1(t *testing.T) {
	wd, err := NewWorkingDirectory("/app")
	require.NoError(t, err)
	assert.Equal(t, "/app", wd.GetPath())
}

// NewWorkingDirectory Ensure the path is prefixed with the current working directory
func Test_NewWorkingDirectory2(t *testing.T) {
	overrideFS(t)
	fs.Getwd = func() (string, error) {
		return "/app", nil
	}

	wd, err := NewWorkingDirectory("internal")
	require.NoError(t, err)
	assert.Equal(t, "/app/internal", wd.GetPath())
}

// NewWorkingDirectory Ensure an error is returned if the path doesn't exist
func Test_NewWorkingDirectory3(t *testing.T) {
	got, err := NewWorkingDirectory("/does-not-exist")
	assert.Empty(t, got)
	assert.EqualError(t, err, "stat /does-not-exist: no such file or directory")
}

// NewWorkingDirectory Ensure an error is returned if the path doesn't exist (after prefixing)
func Test_NewWorkingDirectory4(t *testing.T) {
	overrideFS(t)
	fs.Getwd = func() (string, error) {
		return "/app", nil
	}

	got, err := NewWorkingDirectory("does-not-exist")
	assert.Empty(t, got)
	assert.EqualError(t, err, "stat /app/does-not-exist: no such file or directory")
}

// NewWorkingDirectory Ensure errors from Getwd() are passed along
func Test_NewWorkingDirectory5(t *testing.T) {
	overrideFS(t)
	fs.Getwd = func() (string, error) {
		return "", errors.New("where are we?")
	}

	got, err := NewWorkingDirectory("internal")
	assert.Empty(t, got)
	assert.EqualError(t, err, "where are we?")
}

// NewPath Ensure we can create a new path struct
func Test_NewPath1(t *testing.T) {
	wd, err := NewWorkingDirectory("/app")
	require.NoError(t, err)
	got := wd.NewPath("internal")
	assert.Equal(t, "/app/internal", got.GetFullyQualifiedPath())
}
