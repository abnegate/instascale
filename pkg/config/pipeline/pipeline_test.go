package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipelines_SetDefaults(t *testing.T) {
	t.Run("Sets defaults for lint, test, coverage to true", func(t *testing.T) {
		p := Pipelines{}
		p.SetDefaults()
		assert.NotNil(t, p.Lint)
		assert.True(t, *p.Lint)
		assert.NotNil(t, p.Test)
		assert.True(t, *p.Test)
		assert.NotNil(t, p.Coverage)
		assert.True(t, *p.Coverage)
	})

	t.Run("Does not overwrite existing values", func(t *testing.T) {
		disable := false
		p := Pipelines{Lint: &disable}
		p.SetDefaults()
		assert.False(t, *p.Lint)
	})
}

func TestPipelines_Validate(t *testing.T) {
	t.Run("Valid provider", func(t *testing.T) {
		p := Pipelines{
			Provider: ProviderGitHub,
			Owner:    "some-owner",
			Repo:     "some-repo",
		}
		err := p.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid provider", func(t *testing.T) {
		p := Pipelines{
			Provider: "unknown",
			Owner:    "some-owner",
			Repo:     "some-repo",
		}
		err := p.Validate()
		assert.Error(t, err)
	})
}

func TestPipelines_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		p := &Pipelines{}
		schema := p.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "CI/CD pipeline configuration.", schema.Description)

		_, hasProvider := schema.Properties.Get("provider")
		assert.True(t, hasProvider, "Expected 'provider' property in schema")

		_, hasOwner := schema.Properties.Get("owner")
		assert.True(t, hasOwner, "Expected 'owner' property in schema")

		_, hasRepo := schema.Properties.Get("repo")
		assert.True(t, hasRepo, "Expected 'repo' property in schema")
	})
}
