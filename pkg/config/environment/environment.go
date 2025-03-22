package environment

import (
	"github.com/invopop/jsonschema"
	"instascale/pkg/config/environment/deploy"
	"instascale/pkg/converters"
)

// Environment defines overrides for a given environment (e.g., dev, staging, prod).
type Environment struct {
	Deploy       deploy.Deploy `yaml:"deploy" json:"deploy"`
	Orchestrator Orchestrator  `yaml:"orchestrator" json:"orchestrator"`
}

func (e *Environment) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Environment configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"deploy":       (&deploy.Deploy{}).JSONSchema(),
			"orchestrator": (&Orchestrator{}).JSONSchema(),
		}),
		Required: []string{"deploy", "orchestrator"},
	}
}

// SetDefaults initializes any missing Environment fields with default values.
func (e *Environment) SetDefaults() {
	e.Deploy.SetDefaults()
	e.Orchestrator.SetDefaults()
}

// Validate calls the Validate methods on nested sections.
func (e *Environment) Validate() error {
	if err := e.Deploy.Validate(); err != nil {
		return err
	}
	if err := e.Orchestrator.Validate(); err != nil {
		return err
	}
	return nil
}
