package config

type Config struct {
	Environments []*Environment
	Vars         map[string]any `yaml:"vars"`
	Resources    []*Resource    `yaml:"resources"`
}
