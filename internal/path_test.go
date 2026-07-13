package internal

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func wd(t *testing.T) *WorkingDirectory {
	wd, err := NewWorkingDirectory("/app")
	require.NoError(t, err)
	return wd
}

func overrideFS(t *testing.T) {
	old := fs
	t.Cleanup(func() {
		fs = old
	})
}

// NewWorkingDirectory Ensure an absolute path is used without modification
func Test_NewWorkingDirectory1(t *testing.T) {
	wd, err := NewWorkingDirectory("/app")
	require.NoError(t, err)
	assert.Equal(t, "/app", wd.path)
}

// NewWorkingDirectory Ensure the path is prefixed with the current working directory
func Test_NewWorkingDirectory2(t *testing.T) {
	overrideFS(t)
	fs.Getwd = func() (string, error) {
		return "/app", nil
	}

	wd, err := NewWorkingDirectory("internal")
	require.NoError(t, err)
	assert.Equal(t, "/app/internal", wd.path)
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

// GetFullyQualifiedPath Ensure a relative path is prefixed with the working directory
func Test_GetFullyQualifiedPath1(t *testing.T) {
	path := Path{Path: "manifesting.yaml", WorkingDirectory: wd(t)}
	got := path.GetFullyQualifiedPath()
	assert.Equal(t, "/app/manifesting.yaml", got)
}

// GetFullyQualifiedPath Ensure a fully qualified path is returned without modification
func Test_GetFullyQualifiedPath2(t *testing.T) {
	path := Path{Path: "/manifesting.yaml", WorkingDirectory: wd(t)}
	got := path.GetFullyQualifiedPath()
	assert.Equal(t, "/manifesting.yaml", got)
}

// Exists Ensure true is returned for fully qualified files
func Test_Exists1(t *testing.T) {
	path := Path{Path: "/app/internal/path_test.go", WorkingDirectory: wd(t)}
	got, err := path.Exists()
	require.NoError(t, err)
	assert.Equal(t, true, got)
}

// Exists Ensure true is returned for files in the working directory
func Test_Exists2(t *testing.T) {
	path := Path{Path: "internal/path_test.go", WorkingDirectory: wd(t)}
	got, err := path.Exists()
	require.NoError(t, err)
	assert.Equal(t, true, got)
}

// Exists Ensure false is returned for fully qualified files
func Test_Exists3(t *testing.T) {
	path := Path{Path: "/app/internal/does-not-exist.go", WorkingDirectory: wd(t)}
	got, err := path.Exists()
	require.NoError(t, err)
	assert.Equal(t, false, got)
}

// Exists Ensure true is returned for files in the working directory
func Test_Exists4(t *testing.T) {
	path := Path{Path: "internal/does-not-exist.go", WorkingDirectory: wd(t)}
	got, err := path.Exists()
	require.NoError(t, err)
	assert.Equal(t, false, got)
}

// Exists Ensure an error is returned for fully qualified files
func Test_Exists5(t *testing.T) {
	path := Path{Path: "/bad-perms", WorkingDirectory: wd(t)}

	overrideFS(t)
	fs.Stat = func(name string) (os.FileInfo, error) {
		return nil, fmt.Errorf("can't read file '%s'", name)
	}

	got, err := path.Exists()
	assert.Empty(t, got)
	assert.EqualError(t, err, "can't read file '/bad-perms'")
}

// Exists Ensure true is returned for files in the working directory
func Test_Exists6(t *testing.T) {
	path := Path{Path: "internal/bad-perms", WorkingDirectory: wd(t)}

	overrideFS(t)
	fs.Stat = func(name string) (os.FileInfo, error) {
		return nil, fmt.Errorf("can't read file '%s'", name)
	}

	got, err := path.Exists()
	assert.Empty(t, got)
	assert.EqualError(t, err, "can't read file '/app/internal/bad-perms'")
}

// ExistsOrError Ensure no error is returned for fully qualified files
func Test_ExistsOrError1(t *testing.T) {
	path := Path{Path: "/app/internal/path_test.go", WorkingDirectory: wd(t)}
	err := path.ExistsOrError("")
	require.NoError(t, err)
}

// ExistsOrError Ensure no error is returned for files in the working directory
func Test_ExistsOrError2(t *testing.T) {
	path := Path{Path: "internal/path_test.go", WorkingDirectory: wd(t)}
	err := path.ExistsOrError("")
	require.NoError(t, err)
}

// ExistsOrError Ensure the default error is returned for fully qualified files
func Test_ExistsOrError3(t *testing.T) {
	path := Path{Path: "/app/internal/does-not-exist.go", WorkingDirectory: wd(t)}
	err := path.ExistsOrError("")
	assert.EqualError(t, err, "file '/app/internal/does-not-exist.go' does not exist")
}

// ExistsOrError Ensure the default error is returned for files in the working directory
func Test_ExistsOrError4(t *testing.T) {
	path := Path{Path: "internal/does-not-exist.go", WorkingDirectory: wd(t)}
	err := path.ExistsOrError("")
	assert.EqualError(t, err, "file '/app/internal/does-not-exist.go' does not exist")
}

// ExistsOrError Ensure a custom error is returned for fully qualified files
func Test_ExistsOrError5(t *testing.T) {
	path := Path{Path: "/app/internal/does-not-exist.go", WorkingDirectory: wd(t)}
	err := path.ExistsOrError("bad file %s")
	assert.EqualError(t, err, "bad file /app/internal/does-not-exist.go")
}

// ExistsOrError Ensure a custom error is returned for files in the working directory
func Test_ExistsOrError6(t *testing.T) {
	path := Path{Path: "internal/does-not-exist.go", WorkingDirectory: wd(t)}
	err := path.ExistsOrError(fmt.Sprintf("where is %s file %%s", "my"))
	assert.EqualError(t, err, "where is my file /app/internal/does-not-exist.go")
}

// ExistsOrError Ensure internal errors are passed along fully qualified files
func Test_ExistsOrError7(t *testing.T) {
	path := Path{Path: "/bad-perms", WorkingDirectory: wd(t)}

	overrideFS(t)
	fs.Stat = func(name string) (os.FileInfo, error) {
		return nil, fmt.Errorf("can't read file '%s'", name)
	}

	err := path.ExistsOrError("")
	assert.EqualError(t, err, "can't read file '/bad-perms'")
}

// ExistsOrError Ensure internal errors are passed along for files in the working directory
func Test_ExistsOrError8(t *testing.T) {
	path := Path{Path: "internal/bad-perms", WorkingDirectory: wd(t)}

	overrideFS(t)
	fs.Stat = func(name string) (os.FileInfo, error) {
		return nil, fmt.Errorf("can't read file '%s'", name)
	}

	err := path.ExistsOrError("")
	assert.EqualError(t, err, "can't read file '/app/internal/bad-perms'")
}

// Open Ensure we pass the fully qualified path to the underlying function
func Test_Open1(t *testing.T) {
	path := Path{Path: "file.txt", WorkingDirectory: wd(t)}

	var gotName string
	wantFile := &os.File{}
	wantErr := errors.New("cannot open")
	overrideFS(t)
	fs.Open = func(name string) (*os.File, error) {
		gotName = name
		return wantFile, wantErr
	}

	gotFile, gotErr := path.Open()
	assert.ErrorIs(t, gotErr, wantErr)
	assert.Equal(t, gotFile, wantFile)
	assert.Equal(t, "/app/file.txt", gotName)
}

// ReadFile Ensure we pass the fully qualified path to the underlying function
func Test_ReadFile1(t *testing.T) {
	path := Path{Path: "file.txt", WorkingDirectory: wd(t)}

	var gotName string
	wantData := []byte("ok")
	wantErr := errors.New("cannot read")
	overrideFS(t)
	fs.ReadFile = func(name string) ([]byte, error) {
		gotName = name
		return wantData, wantErr
	}

	gotData, gotErr := path.ReadFile()
	assert.ErrorIs(t, gotErr, wantErr)
	assert.Equal(t, gotData, wantData)
	assert.Equal(t, "/app/file.txt", gotName)
}

// WriteFile Ensure we pass the fully qualified path to the underlying function
func Test_WriteFile1(t *testing.T) {
	path := Path{Path: "file.txt", WorkingDirectory: wd(t)}

	var gotName string
	var gotData []byte
	var gotMode os.FileMode
	wantErr := errors.New("cannot write")
	overrideFS(t)
	fs.WriteFile = func(name string, data []byte, mode os.FileMode) error {
		gotName = name
		gotData = data
		gotMode = mode
		return wantErr
	}

	gotErr := path.WriteFile([]byte("ok"))
	assert.ErrorIs(t, gotErr, wantErr)
	assert.Equal(t, "/app/file.txt", gotName)
	assert.Equal(t, []byte("ok"), gotData)
	assert.Equal(t, os.FileMode(0644), gotMode)
}
