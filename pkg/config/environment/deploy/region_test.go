package deploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegionsConstants(t *testing.T) {
	assert.Contains(t, Regions, RegionAWSUSEast1, "Expected RegionAWSUSEast1 to be in Regions slice")
}
