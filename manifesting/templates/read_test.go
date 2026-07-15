package templates

import (
	"errors"
	"os"
	"testing"

	"github.com/estratocloud/manifesting/internal"
	"github.com/estratocloud/manifesting/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newWorkingDirectory(t *testing.T) internal.WorkingDirectoryInterface {
	wd, err := internal.NewWorkingDirectory("/app/tests/samples")
	require.NoError(t, err)
	return wd
}

// readTemplate Ensure we can read a template with a header
func Test_readTemplate1(t *testing.T) {
	doc1, doc2, err := readTemplate("templates/with-header-doc.yaml.gotmpl", newWorkingDirectory(t))
	require.NoError(t, err)
	assert.Equal(t, "extends: base\n", string(doc1))
	assert.Equal(t, "apiVersion: batch/v1\nkind: CronJob\n", string(doc2))
}

// readTemplate Ensure we can default to the templates/ path and the file extension
func Test_readTemplate2(t *testing.T) {
	doc1, doc2, err := readTemplate("with-header-doc", newWorkingDirectory(t))
	require.NoError(t, err)
	assert.Equal(t, "extends: base\n", string(doc1))
	assert.Equal(t, "apiVersion: batch/v1\nkind: CronJob\n", string(doc2))
}

func Test_readTemplate3(t *testing.T) {
	doc1, doc2, err := readTemplate("templates/no-header-doc.yaml.gotmpl", newWorkingDirectory(t))
	require.NoError(t, err)
	assert.Equal(t, "", string(doc1))
	assert.Equal(t, "apiVersion: batch/v1\nkind: CronJob\n", string(doc2))
}

// readTemplate Ensure we get an error if the file doesn't exist
func Test_readTemplate4(t *testing.T) {
	doc1, doc2, err := readTemplate("does-not-exist", newWorkingDirectory(t))
	assert.Nil(t, doc1)
	assert.Nil(t, doc2)
	assert.EqualError(t, err, "file '/app/tests/samples/templates/does-not-exist.yaml.gotmpl' does not exist")
}

// readTemplate Ensure we get an error if the file doesn't exist
func Test_readTemplate5(t *testing.T) {
	doc1, doc2, err := readTemplate("does-not-exist", newWorkingDirectory(t))
	assert.Nil(t, doc1)
	assert.Nil(t, doc2)
	assert.EqualError(t, err, "file '/app/tests/samples/templates/does-not-exist.yaml.gotmpl' does not exist")
}

// readTemplate Ensure we get an error if the file is empty
func Test_readTemplate6(t *testing.T) {
	doc1, doc2, err := readTemplate("empty.yaml", newWorkingDirectory(t))
	assert.Nil(t, doc1)
	assert.Nil(t, doc2)
	assert.EqualError(t, err, "unable to read first yaml doc from template file (/app/tests/samples/empty.yaml): EOF")
}

// readTemplate Ensure errors from path.Exists() are passed along
func Test_readTemplate7(t *testing.T) {
	wd := &mocks.WorkingDirectory{
		NewPathFunc: func(path string) internal.PathInterface {
			return &mocks.Path{
				Path: path,
				ExistsFunc: func() (bool, error) {
					return false, errors.New("internal problem")
				},
			}
		},
	}
	doc1, doc2, err := readTemplate("file.yaml", wd)
	assert.Nil(t, doc1)
	assert.Nil(t, doc2)
	assert.EqualError(t, err, "unable to check if template file (file.yaml) exists: internal problem")
}

// readTemplate Ensure errors from path.Open() are passed along
func Test_readTemplate8(t *testing.T) {
	wd := &mocks.WorkingDirectory{
		NewPathFunc: func(path string) internal.PathInterface {
			return &mocks.Path{
				Path: path,
				ExistsFunc: func() (bool, error) {
					return true, nil
				},
				OpenFunc: func() (*os.File, error) {
					return nil, errors.New("cannot open")
				},
			}
		},
	}
	doc1, doc2, err := readTemplate("file.yaml", wd)
	assert.Nil(t, doc1)
	assert.Nil(t, doc2)
	assert.EqualError(t, err, "unable to open template file (file.yaml): cannot open")
}
