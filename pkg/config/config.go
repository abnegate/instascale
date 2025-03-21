package config

// Config defines the project configuration.
type Config struct {
	Name         string                 `yaml:"name,omitempty" json:"name,omitempty"`
	Application  Application            `yaml:"application" json:"application"`
	Environments map[string]Environment `yaml:"environments" json:"environments"`
	CI           CI                     `yaml:"ci" json:"ci" jsonschema:"enum=github"`
}

// Validate calls the Validate methods on nested sections.
func (cfg *Config) Validate() error {
	if err := cfg.Application.Validate(); err != nil {
		return err
	}
	for _, env := range cfg.Environments {
		if err := env.Validate(); err != nil {
			return err
		}
	}
	if err := cfg.CI.Validate(); err != nil {
		return err
	}
	return nil
}
