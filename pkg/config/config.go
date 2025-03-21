package config

// Config defines the project configuration.
type Config struct {
	Version      string       `yaml:"version" json:"version"`
	Name         string       `yaml:"name" json:"name" json:"name"`
	Deploy       Deploy       `yaml:"deploy" json:"deploy"`
	Application  Application  `yaml:"application" json:"application"`
	CI           CI           `yaml:"ci" json:"ci" jsonschema:"enum=github"`
	Orchestrator Orchestrator `yaml:"orchestrator" json:"orchestrator" jsonschema:"enum=kubernetes"`
}

// Validate calls the Validate methods on nested sections.
func (cfg *Config) Validate() error {
	if err := cfg.Deploy.Validate(); err != nil {
		return err
	}
	if err := cfg.Application.Validate(); err != nil {
		return err
	}
	if err := cfg.CI.Validate(); err != nil {
		return err
	}
	if err := cfg.Orchestrator.Validate(); err != nil {
		return err
	}
	return nil
}
