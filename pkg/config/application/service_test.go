package application

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestService_UnmarshalYAML(t *testing.T) {
	t.Run("Unmarshal with direct env values", func(t *testing.T) {
		yml := `
name: my-service
type: http
env:
  KEY: $TEST_ENV
`
		os.Setenv("TEST_ENV", "test-value")
		defer os.Unsetenv("TEST_ENV")

		var s Service
		err := yaml.Unmarshal([]byte(yml), &s)
		assert.NoError(t, err)
		assert.Equal(t, "my-service", s.Name)
		assert.Equal(t, "test-value", s.Env["KEY"])
	})

	t.Run("Unmarshal with missing env variable", func(t *testing.T) {
		yml := `
name: my-service
type: http
env:
  KEY: $MISSING_ENV
`
		var s Service
		err := yaml.Unmarshal([]byte(yml), &s)
		assert.Error(t, err)
	})
}

func TestService_SetDefaults(t *testing.T) {
	t.Run("Sets default port for Postgres", func(t *testing.T) {
		s := Service{Type: ServiceTypePostgres}
		s.SetDefaults()
		assert.Equal(t, uint16(5432), s.Port)
	})

	t.Run("Sets default version to latest if empty", func(t *testing.T) {
		s := Service{Name: "my-service"}
		s.SetDefaults()
		assert.Equal(t, "latest", s.Version)
	})

	t.Run("Initializes env map if nil", func(t *testing.T) {
		s := Service{}
		s.SetDefaults()
		assert.NotNil(t, s.Env)
	})
}

func TestService_Validate(t *testing.T) {
	t.Run("Valid service with http type and port", func(t *testing.T) {
		s := Service{
			Name:    "my-http-service",
			Type:    ServiceTypeHttp,
			Version: "1.0.0",
			Port:    8080,
		}
		err := s.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid empty name", func(t *testing.T) {
		s := Service{Type: ServiceTypeHttp, Port: 8080, Version: "1.0.0"}
		err := s.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid port range", func(t *testing.T) {
		s := Service{
			Name:    "invalid-port-service",
			Type:    ServiceTypeHttp,
			Port:    80,
			Version: "1.0.0",
		}
		err := s.Validate()
		assert.Error(t, err)
	})

	t.Run("Empty version not allowed", func(t *testing.T) {
		s := Service{
			Name: "no-version",
			Type: ServiceTypeHttp,
			Port: 8080,
		}
		err := s.Validate()
		assert.Error(t, err)
	})
}

func TestService_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		s := &Service{}
		schema := s.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Service configuration.", schema.Description)
		
		_, hasName := schema.Properties.Get("name")
		assert.True(t, hasName, "Expected 'name' property in schema")

		_, hasType := schema.Properties.Get("type")
		assert.True(t, hasType, "Expected 'type' property in schema")

		_, hasPort := schema.Properties.Get("port")
		assert.True(t, hasPort, "Expected 'port' property in schema")
	})
}
