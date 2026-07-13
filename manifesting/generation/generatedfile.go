package generation

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/estratocloud/manifesting/internal"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

type GeneratedFile struct {
	path    *internal.Path
	content bytes.Buffer
}

func NewGeneratedFile(path *internal.Path) *GeneratedFile {
	return &GeneratedFile{
		path: path,
	}
}

func (f *GeneratedFile) AppendObject(resource runtime.Object) error {
	data, err := yaml.Marshal(resource)
	if err != nil {
		return err
	}

	if f.content.Len() > 0 {
		f.content.WriteString("---\n")
	}

	f.content.Write(data)

	return nil
}

func (f *GeneratedFile) Write() error {
	err := os.MkdirAll(filepath.Dir(f.path.GetFullyQualifiedPath()), 0755)
	if err != nil {
		return err
	}

	return f.path.WriteFile(f.content.Bytes())
}
