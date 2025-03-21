package config

import "fmt"

// Orchestrator defines the orchestration schema-generator.
type Orchestrator string

const (
	OrchestratorKubernetes Orchestrator = "kubernetes"
)

func (o *Orchestrator) Validate() error {
	switch *o {
	case OrchestratorKubernetes:
		return nil
	default:
		return fmt.Errorf("unsupported orchestrator: %q", o)
	}
}
