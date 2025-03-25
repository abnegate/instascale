package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestOrchestrator_UnmarshalYAML(t *testing.T) {
	t.Run("Unmarshal as string", func(t *testing.T) {
		var o Orchestrator
		yml := "compose"
		err := yaml.Unmarshal([]byte(yml), &o)
		assert.NoError(t, err)
		assert.Equal(t, OrchestratorTypeCompose, o.Type)
	})

	t.Run("Unmarshal as map", func(t *testing.T) {
		var o Orchestrator
		yml := "type: swarm"
		err := yaml.Unmarshal([]byte(yml), &o)
		assert.NoError(t, err)
		assert.Equal(t, OrchestratorTypeSwarm, o.Type)
	})
}

func TestOrchestrator_SetDefaults(t *testing.T) {
	t.Run("Sets default orchestrator type to compose", func(t *testing.T) {
		o := Orchestrator{}
		o.SetDefaults()
		assert.Equal(t, OrchestratorTypeCompose, o.Type)
	})

	t.Run("Does not overwrite existing type", func(t *testing.T) {
		o := Orchestrator{Type: OrchestratorTypeKubernetes}
		o.SetDefaults()
		assert.Equal(t, OrchestratorTypeKubernetes, o.Type)
	})
}

func TestOrchestrator_Validate(t *testing.T) {
	t.Run("Valid orchestrator", func(t *testing.T) {
		o := Orchestrator{Type: OrchestratorTypeKubernetes}
		err := o.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid orchestrator", func(t *testing.T) {
		o := Orchestrator{Type: "invalid-type"}
		err := o.Validate()
		assert.Error(t, err)
	})
}

func TestOrchestrator_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		o := &Orchestrator{}
		schema := o.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Orchestrator configuration.", schema.Description)
	})
}
