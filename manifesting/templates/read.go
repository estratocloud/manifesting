package templates

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/estratocloud/manifesting/internal"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func readTemplate(filename string, wd internal.WorkingDirectoryInterface) ([]byte, []byte, error) {

	path := wd.NewPath(filename)
	exists, err := path.Exists()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to check if template file (%s) exists: %w", path.GetFullyQualifiedPath(), err)
	}
	if !exists {
		path = wd.NewPath(fmt.Sprintf("templates/%s.yaml.gotmpl", filename))
		err = path.ExistsOrError("")
		if err != nil {
			return nil, nil, err
		}
	}

	f, err := path.Open()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to open template file (%s): %w", path.GetFullyQualifiedPath(), err)
	}
	defer f.Close()

	yamlDecoder := yaml.NewYAMLReader(bufio.NewReader(f))

	doc1, err := yamlDecoder.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read first yaml doc from template file (%s): %w", path.GetFullyQualifiedPath(), err)
	}

	doc2, err := yamlDecoder.Read()
	if err != nil {
		// If there's only one doc then assume it's the k8s template and there was no header provided
		if errors.Is(err, io.EOF) {
			return []byte{}, doc1, nil
		}
		return nil, nil, fmt.Errorf("unable to read second yaml doc from template file (%s): %w", path.GetFullyQualifiedPath(), err)
	}

	return doc1, doc2, nil
}
