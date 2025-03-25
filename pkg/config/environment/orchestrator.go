package environment

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"gopkg.in/yaml.v3"
	"instascale/pkg/converters"
)

// OrchestratorType defines the supported orchestrator types.
type OrchestratorType string

const (
	OrchestratorTypeKubernetes OrchestratorType = "kubernetes"
	OrchestratorTypeCompose    OrchestratorType = "compose"
	OrchestratorTypeSwarm      OrchestratorType = "swarm"
)

var OrchestratorTypes = []OrchestratorType{
	OrchestratorTypeKubernetes,
	OrchestratorTypeCompose,
	OrchestratorTypeSwarm,
}

// Orchestrator defines the orchestration configuration.
type Orchestrator struct {
	Type OrchestratorType `yaml:"type" json:"type"`
}

// JSONSchema returns a JSON schema for the Orchestrator configuration.
func (o *Orchestrator) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Orchestrator configuration.",
		OneOf: []*jsonschema.Schema{
			{
				Type:        "string",
				Description: "The orchestrator.",
				Enum:        converters.ToAnySlice(OrchestratorTypes),
			},
			{
				Type:        "object",
				Description: "The orchestrator.",
				Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
					"type": {
						Type:        "string",
						Description: "The orchestrator type.",
						Enum:        converters.ToAnySlice(OrchestratorTypes),
					},
					"version": {
						Type:        "string",
						Description: "The language-specific framework version.",
					},
				}),
				Required: []string{"type"},
			},
		},
	}
}

func (o *Orchestrator) UnmarshalYAML(value *yaml.Node) error {
	switch value.Tag {
	case "!!str":
		if err := value.Decode(&o.Type); err != nil {
			return err
		}
	case "!!map":
		type rawOrchestrator Orchestrator
		var raw rawOrchestrator
		if err := value.Decode(&raw); err != nil {
			return err
		}
		*o = Orchestrator(raw)
	default:
		return fmt.Errorf("unsupported YAML type for orchestrator: %v", value.Tag)
	}
	return nil
}

// SetDefaults sets the default values for the Orchestrator configuration.
func (o *Orchestrator) SetDefaults() {
	if o.Type == "" {
		o.Type = OrchestratorTypeCompose
	}
}

// Validate checks that the Orchestrator configuration is valid.
func (o *Orchestrator) Validate() error {
	switch o.Type {
	case OrchestratorTypeKubernetes, OrchestratorTypeCompose, OrchestratorTypeSwarm:
		return nil
	default:
		return fmt.Errorf("unsupported orchestrator: %q", o.Type)
	}
}
