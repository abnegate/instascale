package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestFramework_UnmarshalYAML(t *testing.T) {
	t.Run("Unmarshal as string", func(t *testing.T) {
		var f Framework
		yml := "express"
		err := yaml.Unmarshal([]byte(yml), &f)
		assert.NoError(t, err)
		assert.Equal(t, FrameworkExpress, f.Name)
		assert.Equal(t, "", f.Version)
	})

	t.Run("Unmarshal as map", func(t *testing.T) {
		var f Framework
		yml := "name: express\nversion: 4.17.1"
		err := yaml.Unmarshal([]byte(yml), &f)
		assert.NoError(t, err)
		assert.Equal(t, FrameworkExpress, f.Name)
		assert.Equal(t, FrameworkExpressLatest, f.Version) // Because 4.17.1 == FrameworkExpressLatest
	})
}

func TestFramework_SetDefaults(t *testing.T) {
	t.Run("Sets default for express if version is empty", func(t *testing.T) {
		f := Framework{Name: FrameworkExpress}
		f.SetDefaults()
		assert.Equal(t, FrameworkExpressLatest, f.Version)
	})

	t.Run("Does not overwrite existing version", func(t *testing.T) {
		f := Framework{Name: FrameworkExpress, Version: "4.16.0"}
		f.SetDefaults()
		assert.Equal(t, "4.16.0", f.Version)
	})
}

func TestFramework_Validate(t *testing.T) {
	t.Run("Valid express framework", func(t *testing.T) {
		f := Framework{Name: FrameworkExpress, Version: "4.17.1"}
		err := f.Validate()
		assert.NoError(t, err)
	})

	t.Run("Empty version for express is invalid", func(t *testing.T) {
		f := Framework{Name: FrameworkExpress, Version: ""}
		err := f.Validate()
		assert.Error(t, err)
	})

	t.Run("Unsupported framework", func(t *testing.T) {
		f := Framework{Name: "some-other-framework", Version: "1.0.0"}
		err := f.Validate()
		assert.Error(t, err)
	})
}

func TestFramework_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		f := &Framework{}
		schema := f.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Framework to use for the application.", schema.Description)
	})
}
