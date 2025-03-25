package config

import (
	"instascale/pkg/config/environment/deploy"
	"testing"

	"github.com/stretchr/testify/assert"
	"instascale/pkg/config/application"
	"instascale/pkg/config/environment"
)

func TestConfig_Validate(t *testing.T) {
	t.Run("Valid config", func(t *testing.T) {
		cfg := Config{
			Name: "my-project",
			Application: application.Application{
				Language: application.Language{
					Name:      application.LanguageJavaScript,
					Version:   application.LanguageVersionES6,
					Framework: application.Framework{Name: application.FrameworkExpress, Version: application.FrameworkExpressLatest},
				},
				Services: []application.Service{
					{
						Name:    "svc",
						Type:    application.ServiceTypeHttp,
						Version: "1.0.0",
						Port:    8080,
					},
				},
			},
			Environments: map[string]environment.Environment{
				"dev": {
					Deploy: deploy.Deploy{
						Target: "local",
					},
					Orchestrator: environment.Orchestrator{Type: "compose"},
				},
			},
		}
		err := cfg.Validate()
		assert.NoError(t, err)
	})

	t.Run("Missing project name", func(t *testing.T) {
		cfg := Config{}
		err := cfg.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid environment", func(t *testing.T) {
		cfg := Config{
			Name: "my-project",
			Environments: map[string]environment.Environment{
				"dev": {
					Deploy: deploy.Deploy{
						Target:  "aws",
						Regions: []deploy.Region{},
					},
					Orchestrator: environment.Orchestrator{Type: "compose"},
				},
			},
		}
		err := cfg.Validate()
		assert.Error(t, err)
	})
}

func TestConfig_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		cfg := &Config{}
		schema := cfg.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Project configuration.", schema.Description)

		_, hasName := schema.Properties.Get("name")
		assert.True(t, hasName, "Expected 'name' property in schema")

		_, hasApplication := schema.Properties.Get("application")
		assert.True(t, hasApplication, "Expected 'application' property in schema")

		_, hasEnvironments := schema.Properties.Get("environments")
		assert.True(t, hasEnvironments, "Expected 'environments' property in schema")
	})
}
