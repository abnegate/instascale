package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestLanguage_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		l := &Language{}
		schema := l.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Language configuration.", schema.Description)

		_, hasName := schema.Properties.Get("name")
		assert.True(t, hasName, "Expected 'name' property in schema")

		_, hasFramework := schema.Properties.Get("framework")
		assert.True(t, hasFramework, "Expected 'framework' property in schema")
	})
}

func TestLanguage_UnmarshalYAML(t *testing.T) {
	t.Run("Unmarshal as string", func(t *testing.T) {
		var l Language
		yml := "javascript"
		err := yaml.Unmarshal([]byte(yml), &l)
		assert.NoError(t, err)
		assert.Equal(t, LanguageJavaScript, l.Name)
		assert.Equal(t, LanguageVersion(""), l.Version)
	})

	t.Run("Unmarshal as map", func(t *testing.T) {
		var l Language
		yml := "name: javascript\nversion: es6\nframework:\n  name: express"
		err := yaml.Unmarshal([]byte(yml), &l)
		assert.NoError(t, err)
		assert.Equal(t, LanguageJavaScript, l.Name)
		assert.Equal(t, LanguageVersionES6, l.Version)
		assert.Equal(t, FrameworkExpress, l.Framework.Name)
	})

	t.Run("Unmarshal as invalid type", func(t *testing.T) {
		var l Language
		yml := `100`
		err := yaml.Unmarshal([]byte(yml), &l)
		assert.Error(t, err)
	})
}

func TestLanguage_SetDefaults(t *testing.T) {
	t.Run("Sets default version for JavaScript if empty", func(t *testing.T) {
		l := Language{Name: LanguageJavaScript}
		l.SetDefaults()
		assert.Equal(t, LanguageVersionES6, l.Version)
	})

	t.Run("Calls SetDefaults on framework", func(t *testing.T) {
		l := Language{
			Name: LanguageJavaScript,
			Framework: Framework{
				Name: FrameworkExpress,
			},
		}
		l.SetDefaults()
		assert.Equal(t, FrameworkExpressLatest, l.Framework.Version)
	})
}

func TestLanguage_Validate(t *testing.T) {
	t.Run("Valid language and framework", func(t *testing.T) {
		l := Language{
			Name:      LanguageJavaScript,
			Version:   LanguageVersionES6,
			Framework: Framework{Name: FrameworkExpress, Version: "4.17.1"},
		}
		err := l.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid framework for language", func(t *testing.T) {
		l := Language{
			Name:      LanguageJavaScript,
			Version:   LanguageVersionES6,
			Framework: Framework{Name: "other-framework", Version: "1.0.0"},
		}
		err := l.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid version for language", func(t *testing.T) {
		l := Language{
			Name:      LanguageJavaScript,
			Version:   "es7",
			Framework: Framework{Name: FrameworkExpress, Version: "4.17.1"},
		}
		err := l.Validate()
		assert.Error(t, err)
	})
}
