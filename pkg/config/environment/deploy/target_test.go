package deploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllowedRegions(t *testing.T) {
	t.Run("AWS returns known regions", func(t *testing.T) {
		res := allowedRegions(TargetAWS)
		assert.Contains(t, res, RegionAWSUSEast1)
	})

	t.Run("Local returns nil or empty", func(t *testing.T) {
		res := allowedRegions(TargetLocal)
		assert.Empty(t, res)
	})
}
