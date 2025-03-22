package deploy

// Region defines a deployment region.
type Region string

const (
	RegionAWSUSEast1 Region = "us-east-1"
)

var Regions = []Region{RegionAWSUSEast1}
