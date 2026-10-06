package config

import (
	"fmt"
	"path/filepath"
	"reflect"

	"github.com/estratocloud/manifesting/internal"
	"github.com/estratocloud/manifesting/internal/deprecations"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

type Environment struct {
	Name               string `yaml:"name"`
	Output             string `yaml:"output"`
	DefaultEnvVarsFile string `yaml:"defaultEnvVarsFile"`
}

func (e *Environment) GetOutputPath(gd internal.PathInterface, wd internal.WorkingDirectoryInterface) internal.PathInterface {
	if e.Output == "" {
		return wd.NewPath(filepath.Join(gd.GetPath(), e.Name+".yaml"))
	}
	return wd.NewPath(e.Output)
}

func (e *Environment) GetEnvVars(wd internal.WorkingDirectoryInterface, d *deprecations.Checker) (map[string]corev1.EnvVar, error) {

	envvars := map[string]corev1.EnvVar{}

	if e.DefaultEnvVarsFile == "" {
		return envvars, nil
	}

	path := wd.NewPath(e.DefaultEnvVarsFile)
	err := path.ExistsOrError(fmt.Sprintf("unable to find the defaultEnvVarsFile file for %s at '%%s'", e.Name))
	if err != nil {
		return nil, err
	}

	data, err := path.ReadFile()
	if err != nil {
		return nil, fmt.Errorf("unable to read the defaultEnvVarsFile file for %s at '%s': %w", e.Name, path.GetFullyQualifiedPath(), err)
	}

	var objects []corev1.EnvVar
	if err := yaml.Unmarshal(data, &objects); err != nil {
		return nil, fmt.Errorf("unable to parse the defaultEnvVarsFile file for %s at '%s': %w", e.Name, path.GetFullyQualifiedPath(), err)
	}

	envvars = make(map[string]corev1.EnvVar, len(objects))
	for _, env := range objects {
		envvars[env.Name] = env
	}

	return envvars, nil
}

func (e *Environment) PerEnvironment(values any, defaultValue any) any {
	if values == nil {
		return defaultValue
	}

	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Map && v.Type().Key().Kind() == reflect.String {
		value := v.MapIndex(reflect.ValueOf(e.Name))
		if value.IsValid() {
			return value.Interface()
		}
		return defaultValue
	}

	return values
}
