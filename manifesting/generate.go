package manifesting

import (
	"github.com/estratocloud/manifesting/internal"
	"github.com/estratocloud/manifesting/manifesting/config"
	"github.com/estratocloud/manifesting/manifesting/generation"
	"github.com/estratocloud/manifesting/manifesting/kubernetes"
	"github.com/estratocloud/manifesting/manifesting/templates"
)

func GenerateManifest(environment *config.Environment, conf *config.Config, wd internal.WorkingDirectoryInterface) error {

	envvars, err := environment.GetEnvVars(wd)
	if err != nil {
		return err
	}

	output := generation.NewGeneratedFile(environment.GetOutputPath(wd))
	for _, resource := range conf.GetResources(environment) {
		template, err := templates.NewTemplate(resource, environment, wd)
		if err != nil {
			return err
		}

		vars, err := resource.GetVars(environment, conf)
		if err != nil {
			return err
		}

		object, err := template.Render(vars, environment, wd)
		if err != nil {
			return err
		}
		kubernetes.SetEnvironmentVariableDefaults(object, envvars)

		err = output.AppendObject(object)
		if err != nil {
			return err
		}
	}

	return output.Write()
}
