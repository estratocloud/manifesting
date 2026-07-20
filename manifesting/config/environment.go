package config

import (
	"fmt"
	"reflect"

	"github.com/estratocloud/manifesting/internal"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

type Environment struct {
	Name    string `yaml:"name"`
	Output  string `yaml:"output"`
	EnvFrom string `yaml:"envFrom"`
}

func (e *Environment) GetOutputPath(wd internal.WorkingDirectoryInterface) internal.PathInterface {
	if e.Output == "" {
		return wd.NewPath(fmt.Sprintf(".generated/%s.yaml", e.Name))
	}
	return wd.NewPath(e.Output)
}

func (e *Environment) GetEnvVars(wd internal.WorkingDirectoryInterface) (map[string]corev1.EnvVar, error) {

	envvars := map[string]corev1.EnvVar{}

	if e.EnvFrom == "" {
		return envvars, nil
	}

	path := wd.NewPath(e.EnvFrom)
	err := path.ExistsOrError(fmt.Sprintf("unable to find the envFrom file for %s at '%%s'", e.Name))
	if err != nil {
		return nil, err
	}

	data, err := path.ReadFile()
	if err != nil {
		return nil, fmt.Errorf("unable to read the envFrom file for %s at '%s': %w", e.Name, path.GetFullyQualifiedPath(), err)
	}

	var objects []corev1.EnvVar
	if err := yaml.Unmarshal(data, &objects); err != nil {
		return nil, fmt.Errorf("unable to parse the envFrom file for %s at '%s': %w", e.Name, path.GetFullyQualifiedPath(), err)
	}

	envvars = make(map[string]corev1.EnvVar, len(objects))
	for _, env := range objects {
		envvars[env.Name] = env
	}

	return envvars, nil
}

func (e *Environment) PerEnvironment(values any) any {
	v := reflect.ValueOf(values)

	if v.Kind() == reflect.Map && v.Type().Key().Kind() == reflect.String {
		value := v.MapIndex(reflect.ValueOf(e.Name))
		if value.IsValid() {
			return value.Interface()
		}
		return nil
	}

	return values
}
