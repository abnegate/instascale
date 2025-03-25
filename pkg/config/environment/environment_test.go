package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"instascale/pkg/config/environment/deploy"
)

func TestEnvironment_SetDefaults(t *testing.T) {
	t.Run("Sets defaults for deploy and orchestrator", func(t *testing.T) {
		env := Environment{}
		env.SetDefaults()
		assert.Equal(t, deploy.TargetLocal, env.Deploy.Target)
		assert.Equal(t, OrchestratorTypeCompose, env.Orchestrator.Type)
	})
}

func TestEnvironment_Validate(t *testing.T) {
	t.Run("Valid environment", func(t *testing.T) {
		env := Environment{
			Deploy: deploy.Deploy{
				Target:  deploy.TargetAWS,
				Regions: []deploy.Region{deploy.RegionAWSUSEast1},
			},
			Orchestrator: Orchestrator{Type: OrchestratorTypeCompose},
		}
		err := env.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid deploy", func(t *testing.T) {
		env := Environment{
			Deploy: deploy.Deploy{
				Target:  deploy.TargetAWS,
				Regions: []deploy.Region{},
			},
			Orchestrator: Orchestrator{Type: OrchestratorTypeCompose},
		}
		err := env.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid orchestrator", func(t *testing.T) {
		env := Environment{
			Deploy: deploy.Deploy{
				Target:  deploy.TargetLocal,
				Regions: []deploy.Region{},
			},
			Orchestrator: Orchestrator{Type: "unknown"},
		}
		err := env.Validate()
		assert.Error(t, err)
	})
}

func TestEnvironment_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		env := &Environment{}
		schema := env.JSONSchema()
		assert.NotNil(t, schema)

		_, hasDeploy := schema.Properties.Get("deploy")
		assert.True(t, hasDeploy, "Expected 'deploy' property in schema")

		_, hasOrchestrator := schema.Properties.Get("orchestrator")
		assert.True(t, hasOrchestrator, "Expected 'orchestrator' property in schema")
	})
}
