package generation

import (
	"errors"
	"os"
	"testing"

	"github.com/estratocloud/manifesting/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newPath(t *testing.T, path string) *internal.Path {
	wd, err := internal.NewWorkingDirectory("/app")
	require.NoError(t, err)
	return &internal.Path{
		Path:             path,
		WorkingDirectory: wd,
	}
}

type BadObject struct {
	metav1.Status
}

func (b *BadObject) MarshalJSON() ([]byte, error) {
	return nil, errors.New("bad object")
}

// NewGeneratedFile Ensure we can create a new instance
func Test_NewGeneratedFile1(t *testing.T) {
	path := newPath(t, "output.txt")
	f := NewGeneratedFile(path)
	assert.IsType(t, GeneratedFile{}, *f)
}

// AppendObject Ensure failures are returned
func Test_AppendObject1(t *testing.T) {
	f := NewGeneratedFile(newPath(t, "output.txt"))

	err := f.AppendObject(&BadObject{})
	assert.EqualError(t, err, "error marshaling into JSON: json: error calling MarshalJSON for type *generation.BadObject: bad object")
}

// AppendObject Ensure yaml separators are added
func Test_AppendObject2(t *testing.T) {
	f := NewGeneratedFile(newPath(t, "output.txt"))

	err := f.AppendObject(&metav1.Status{
		TypeMeta: metav1.TypeMeta{
			Kind: "Object1",
		},
	})
	require.NoError(t, err)

	err = f.AppendObject(&metav1.Status{
		TypeMeta: metav1.TypeMeta{
			Kind: "Object2",
		},
	})
	require.NoError(t, err)

	want := "kind: Object1\nmetadata: {}\n"
	want += "---\n"
	want += "kind: Object2\nmetadata: {}\n"
	assert.Equal(t, want, f.content.String())
}

// Write Ensure directories are created on the file
func Test_Write1(t *testing.T) {
	f := NewGeneratedFile(newPath(t, "/tmp/one/two/three.txt"))

	err := f.AppendObject(&metav1.Status{
		TypeMeta: metav1.TypeMeta{
			Kind: "Test1",
		},
	})
	require.NoError(t, err)

	err = f.Write()
	require.NoError(t, err)

	data, err := os.ReadFile("/tmp/one/two/three.txt")
	require.NoError(t, err)
	assert.Equal(t, "kind: Test1\nmetadata: {}\n", string(data))
}

// Write Ensure errors are returned on directory creation failure
func Test_Write2(t *testing.T) {
	f := NewGeneratedFile(newPath(t, "/tmp/\x00/not-valid"))
	err := f.Write()
	assert.EqualError(t, err, "mkdir /tmp/\x00: invalid argument")
}
