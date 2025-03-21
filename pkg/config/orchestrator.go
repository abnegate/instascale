package config

import "fmt"

// Orchestrator defines the orchestration schema-generator.
type Orchestrator string

const (
	OrchestratorKubernetes Orchestrator = "kubernetes"
	OrchestratorCompose    Orchestrator = "compose"
	OrchestratorSwarm      Orchestrator = "swarm"
)

func (o *Orchestrator) Validate() error {
	switch *o {
	case OrchestratorKubernetes, OrchestratorCompose, OrchestratorSwarm:
		return nil
	default:
		return fmt.Errorf("unsupported orchestrator: %q", o)
	}
}
