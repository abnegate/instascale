package deploy

// Region defines a deployment region.
type Region string

const (
	RegionAWSUSEast1 Region = "us-east-1"
	RegionAWSUSWest1 Region = "us-west-1"
)

var Regions = []Region{
	RegionAWSUSEast1,
	RegionAWSUSWest1,
}

var AWSREgions = []Region{
	RegionAWSUSEast1,
	RegionAWSUSWest1,
}
