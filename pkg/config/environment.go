package config

// Environment defines overrides for a given environment (e.g., dev, staging, prod).
type Environment struct {
	Deploy       Deploy       `yaml:"deploy" json:"deploy"`
	Orchestrator Orchestrator `yaml:"orchestrator" json:"orchestrator" jsonschema:"enum=kubernetes,enum=compose,enum=swarm"`
}

// Validate calls the Validate methods on nested sections.
func (cfg *Environment) Validate() error {
	if err := cfg.Deploy.Validate(); err != nil {
		return err
	}
	if err := cfg.Orchestrator.Validate(); err != nil {
		return err
	}
	return nil
}
