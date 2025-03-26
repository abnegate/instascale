package deploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploy_JSONSchema(t *testing.T) {
	t.Run("Schema generation", func(t *testing.T) {
		d := &Deploy{}
		schema := d.JSONSchema()
		assert.NotNil(t, schema)
		assert.Equal(t, "Deploy configuration.", schema.Description)

		_, hasTarget := schema.Properties.Get("target")
		assert.True(t, hasTarget, "Expected 'target' property in schema")

		_, hasRegions := schema.Properties.Get("regions")
		assert.True(t, hasRegions, "Expected 'regions' property in schema")
	})
}

func TestDeploy_SetDefaults(t *testing.T) {
	t.Run("Sets default target to local if empty", func(t *testing.T) {
		d := Deploy{}
		d.SetDefaults()
		assert.Equal(t, TargetLocal, d.Target)
	})
}

func TestDeploy_Validate(t *testing.T) {
	t.Run("Valid local deploy with no regions", func(t *testing.T) {
		d := Deploy{Target: TargetLocal}
		err := d.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid local deploy with regions", func(t *testing.T) {
		d := Deploy{
			Target:  TargetLocal,
			Regions: []Region{RegionAWSUSEast1},
		}
		err := d.Validate()
		assert.Error(t, err)
	})

	t.Run("Valid AWS deploy with region", func(t *testing.T) {
		d := Deploy{
			Target:  TargetAWS,
			Regions: []Region{RegionAWSUSEast1},
		}
		err := d.Validate()
		assert.NoError(t, err)
	})

	t.Run("Invalid AWS deploy with no regions", func(t *testing.T) {
		d := Deploy{Target: TargetAWS}
		err := d.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid region for AWS", func(t *testing.T) {
		d := Deploy{
			Target:  TargetAWS,
			Regions: []Region{"unknown-region"},
		}
		err := d.Validate()
		assert.Error(t, err)
	})

	t.Run("Invalid target", func(t *testing.T) {
		d := Deploy{Target: "invalid-target"}
		err := d.Validate()
		assert.Error(t, err)
	})
}
