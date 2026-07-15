package templates

import (
	text "text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/group/all"
)

func getTextTemplate(content []byte) (*text.Template, error) {

	tmpl := text.New("template")

	tmpl.Funcs(sprout.New(sprout.WithGroups(all.RegistryGroup())).Build())

	_, err := tmpl.Parse(string(content))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
