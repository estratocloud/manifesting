package config

type Resource struct {
	Name         string         `yaml:"name"`
	Template     string         `yaml:"template"`
	Environments []string       `yaml:"environments"`
	Vars         map[string]any `yaml:"vars"`
}
