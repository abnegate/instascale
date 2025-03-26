package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication_SetDefaults(t *testing.T) {
	t.Run("Sets default version when empty", func(t *testing.T) {
		app := Application{}
		app.SetDefaults()
		assert.Equal(t, "1.0.0", app.Version)
	})

	t.Run("Does not overwrite existing version", func(t *testing.T) {
		app := Application{Version: "2.0.0"}
		app.SetDefaults()
		assert.Equal(t, "2.0.0", app.Version)
	})

	t.Run("Calls SetDefaults on language and services", func(t *testing.T) {
		app := Application{
			Services: []Service{{Name: "service1"}},
		}
		app.SetDefaults()
		assert.Equal(t, "latest", app.Services[0].Version)
	})
}

func TestApplication_Validate(t *testing.T) {
	t.Run("Valid Application", func(t *testing.T) {
		app := Application{
			Language: Language{
				Name:    LanguageJavaScript,
				Version: LanguageVersionES6,
				Framework: Framework{
					Name:    FrameworkExpress,
					Version: FrameworkExpressLatest,
				},
			},
			Services: []Service{
				{
					Name:    "my-service",
					Type:    ServiceTypeHttp,
					Version: "v1",
					Port:    8080,
				},
			},
		}
		err := app.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid language", func(t *testing.T) {
		app := Application{
			Language: Language{
				Name:    "some-other-lang",
				Version: "v1",
				Framework: Framework{
					Name:    FrameworkExpress,
					Version: FrameworkExpressLatest,
				},
			},
			Services: []Service{
				{
					Name:    "my-service",
					Type:    ServiceTypeHttp,
					Port:    8080,
					Version: "1.0.0",
				},
			},
		}
		err := app.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid service", func(t *testing.T) {
		app := Application{
			Language: Language{
				Name:    LanguageJavaScript,
				Version: LanguageVersionES6,
				Framework: Framework{
					Name:    FrameworkExpress,
					Version: FrameworkExpressLatest,
				},
			},
			Services: []Service{
				{
					Name: "",
					Type: ServiceTypeHttp,
				},
			},
		}
		err := app.Validate()
		assert.Error(t, err)
	})
}

func TestApplication_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		app := &Application{}
		schema := app.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Application configuration.", schema.Description)

		_, hasVersion := schema.Properties.Get("version")
		assert.True(t, hasVersion, "Expected 'version' property in schema")

		_, hasServices := schema.Properties.Get("services")
		assert.True(t, hasServices, "Expected 'services' property in schema")
	})
}
