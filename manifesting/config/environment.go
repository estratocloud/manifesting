package config

type Environment struct {
	Name    string `yaml:"name"`
	Output  string `yaml:"output"`
	EnvFrom string `yaml:"envFrom"`
}
