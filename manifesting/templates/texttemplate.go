package templates

import (
	text "text/template"

	"github.com/estratocloud/manifesting/manifesting/config"
	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/group/all"
)

func getTextTemplate(content []byte, environment *config.Environment) (*text.Template, error) {

	tmpl := text.New("template")

	tmpl.Funcs(sprout.New(sprout.WithGroups(all.RegistryGroup())).Build())

	tmpl.Funcs(text.FuncMap{
		"perEnvironment": environment.PerEnvironment,
	})

	_, err := tmpl.Parse(string(content))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
