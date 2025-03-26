package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		p := &Pipeline{}
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

func TestPipeline_SetDefaults(t *testing.T) {
	t.Run("Sets defaults for lint, test, coverage to true", func(t *testing.T) {
		p := Pipeline{}
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
		p := Pipeline{Lint: &disable}
		p.SetDefaults()
		assert.False(t, *p.Lint)
	})
}

func TestPipeline_Validate(t *testing.T) {
	t.Run("Valid provider", func(t *testing.T) {
		p := Pipeline{
			Provider: ProviderGitHub,
			Owner:    "some-owner",
			Repo:     "some-repo",
		}
		err := p.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid provider", func(t *testing.T) {
		p := Pipeline{
			Provider: "unknown",
			Owner:    "some-owner",
			Repo:     "some-repo",
		}
		err := p.Validate()
		assert.Error(t, err)
	})
}
