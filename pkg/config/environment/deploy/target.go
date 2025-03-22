package deploy

// Target defines the deployment target.
type Target string

const (
	TargetLocal Target = "local"
	TargetAWS   Target = "aws"
)

var Targets = []Target{TargetLocal, TargetAWS}

func allowedRegions(target Target) []Region {
	switch target {
	case TargetAWS:
		return []Region{RegionAWSUSEast1}
	default:
		return nil
	}
}
