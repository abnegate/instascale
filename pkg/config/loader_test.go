package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoader(t *testing.T) {
	t.Run("Loads config", func(t *testing.T) {
		os.Setenv("POSTGRES_USER", "test")
		os.Setenv("POSTGRES_PASSWORD", "test")

		cfg, err := LoadConfig("../../example.full.yaml")

		assert.NoError(t, err, "Expected no error when loading valid config")
		assert.NotNil(t, cfg, "Expected a non-nil config")
		assert.Equal(t, "test-project", cfg.Name)
		assert.NotEmpty(t, cfg.Environments, "Expected environments to be loaded")
		assert.NotEmpty(t, cfg.Application.Version, "Expected version to be loaded")
		assert.NotEmpty(t, cfg.Application.Services, "Expected services to be loaded")
		assert.NotEmpty(t, cfg.Application.Language, "Expected language to be loaded")
		assert.NotEmpty(t, cfg.Application.Language.Framework, "Expected framework to be loaded")
	})

	t.Run("Fails to load config", func(t *testing.T) {
		_, err := LoadConfig("non-existent.yml")
		assert.Error(t, err, "Expected an error for missing file")
	})

	t.Run("Invalid YAML", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "invalid-config-*.yaml")
		assert.NoError(t, err)
		defer os.Remove(tmpFile.Name())

		// Write some garbage YAML
		_, err = tmpFile.Write([]byte("not: : valid: : yaml"))
		assert.NoError(t, err)
		tmpFile.Close()

		_, err = LoadConfig(tmpFile.Name())
		assert.Error(t, err, "Expected an error for invalid YAML")
	})

}
