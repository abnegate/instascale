package vcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVCS_SetDefaults(t *testing.T) {
	t.Run("Sets default visibility to public", func(t *testing.T) {
		v := VCS{}
		v.SetDefaults()
		assert.Equal(t, "public", v.Visibility)
	})

	t.Run("Does not overwrite existing visibility", func(t *testing.T) {
		v := VCS{Visibility: "private"}
		v.SetDefaults()
		assert.Equal(t, "private", v.Visibility)
	})
}

func TestVCS_Validate(t *testing.T) {
	t.Run("Valid provider", func(t *testing.T) {
		v := VCS{
			Provider: ProviderGitHub,
			Owner:    "some-owner",
			Repo:     "some-repo",
		}
		err := v.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid provider", func(t *testing.T) {
		v := VCS{
			Provider: "unknown",
			Owner:    "some-owner",
			Repo:     "some-repo",
		}
		err := v.Validate()
		assert.Error(t, err)
	})
}

func TestVCS_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		v := &VCS{}
		schema := v.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Version Control System configuration.", schema.Description)

		_, hasProvider := schema.Properties.Get("provider")
		assert.True(t, hasProvider, "Expected 'provider' property in schema")

		_, hasOwner := schema.Properties.Get("owner")
		assert.True(t, hasOwner, "Expected 'owner' property in schema")

		_, hasRepo := schema.Properties.Get("repo")
		assert.True(t, hasRepo, "Expected 'repo' property in schema")
	})
}
