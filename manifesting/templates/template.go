package templates

import (
	"bytes"
	"fmt"
	text "text/template"

	"github.com/estratocloud/manifesting/internal"
	"github.com/estratocloud/manifesting/manifesting/config"
	"github.com/estratocloud/manifesting/manifesting/kubernetes"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

type Template struct {
	Extends  string
	Template *text.Template
}

func NewTemplate(resource *config.Resource, environment *config.Environment, wd internal.WorkingDirectoryInterface) (*Template, error) {

	doc1, doc2, err := readTemplate(resource.Template, wd)
	if err != nil {
		return nil, fmt.Errorf("unable to read template for resource '%s': %w", resource.Name, err)
	}

	var template Template
	err = yaml.Unmarshal(doc1, &template)
	if err != nil {
		return nil, fmt.Errorf("unable to parse header for template '%s': %w", resource.Template, err)
	}

	template.Template, err = getTextTemplate(doc2, environment)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the template '%s': %w", resource.Template, err)
	}

	return &template, nil
}

func (t *Template) Render(vars map[string]any, environment *config.Environment, wd internal.WorkingDirectoryInterface) (runtime.Object, error) {

	var buffer bytes.Buffer

	err := t.Template.Execute(&buffer, vars)
	if err != nil {
		return nil, err
	}

	content := buffer.Bytes()

	if t.Extends != "" {
		_, base, err := readTemplate(t.Extends, wd)
		if err != nil {
			return nil, err
		}
		tmpl, err := getTextTemplate(base, environment)
		if err != nil {
			return nil, err
		}

		var baseBuffer bytes.Buffer
		err = tmpl.Execute(&baseBuffer, vars)
		if err != nil {
			return nil, err
		}

		content, err = internal.MergeYAML(baseBuffer.Bytes(), content)
		if err != nil {
			return nil, err
		}
	}

	return kubernetes.GetObject(content)
}
