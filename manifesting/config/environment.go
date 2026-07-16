package config

import (
	"fmt"

	"github.com/estratocloud/manifesting/internal"
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
