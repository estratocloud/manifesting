package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// GetArgs Ensure the working directory passed is used (and the default config file name)
func Test_GetArgs1(t *testing.T) {
	args, err := GetArgs([]string{
		"--working-dir=/app/tests/samples/config",
	})
	require.NoError(t, err)
	assert.Equal(t, "/app/tests/samples/config/manifesting.yaml", args.configPath.GetFullyQualifiedPath())
}

// GetArgs Ensure it uses the provided config file (and the current directory by default)
func Test_GetArgs2(t *testing.T) {
	args, err := GetArgs([]string{
		"--config=../../tests/samples/config/manifesting.yaml",
	})
	require.NoError(t, err)
	assert.Equal(t, "/app/manifesting/cli", args.workingDirectory.GetPath())
}

// GetArgs Ensure it returns an error if there's no config file
func Test_GetArgs3(t *testing.T) {
	got, err := GetArgs([]string{})
	assert.Nil(t, got)
	assert.EqualError(t, err, "unable to find the manifesting --config file at '/app/manifesting/cli/manifesting.yaml'")
}

// GetArgs Ensure it returns an error if the working directory doesn't exist
func Test_GetArgs4(t *testing.T) {
	got, err := GetArgs([]string{
		"--working-dir=/does-not-exist",
	})
	assert.Nil(t, got)
	assert.EqualError(t, err, "unable to use the --working-dir '/does-not-exist': stat /does-not-exist: no such file or directory")
}

// GetArgs Ensure the generated directory passed is used
func Test_GetArgs5(t *testing.T) {
	args, err := GetArgs([]string{
		"--generated-dir=.my-generated-dir",
		"--working-dir=/app/tests/samples/config",
	})
	require.NoError(t, err)
	assert.Equal(t, ".my-generated-dir", args.generatedDirectory.GetPath())
}

// GetArgs Ensure the generated directory can be passed as empty
func Test_GetArgs6(t *testing.T) {
	args, err := GetArgs([]string{
		"--generated-dir=",
		"--working-dir=/app/tests/samples/config",
	})
	require.NoError(t, err)
	assert.Equal(t, "", args.generatedDirectory.GetPath())
}

// GetArgs Ensure the generated directory defaults to .generated if the flag isn't supplied
func Test_GetArgs7(t *testing.T) {
	args, err := GetArgs([]string{
		"--working-dir=/app/tests/samples/config",
	})
	require.NoError(t, err)
	assert.Equal(t, ".generated", args.generatedDirectory.GetPath())
}
