package config

import "fmt"

// Target defines the deployment target.
type Target string

const (
	TargetLocal Target = "local"
	TargetAWS   Target = "aws"
)

// Region defines a deployment region.
type Region string

const (
	RegionAWSUSEast1 Region = "us-east-1"
)

// Deploy nests target-related configuration.
type Deploy struct {
	Target  Target   `yaml:"target" json:"target" jsonschema:"enum=local,enum=aws"`
	Regions []Region `yaml:"regions,omitempty" json:"regions,omitempty" jsonschema:"omitempty,uniqueItems=true,enum=us-east-1"`
}

// Validate checks that the Deploy configuration is valid.
func (d *Deploy) Validate() error {
	if err := validateTarget(d.Target); err != nil {
		return err
	}

	return validateRegions(d.Target, d.Regions)
}

func validateTarget(t Target) error {
	switch t {
	case TargetLocal, TargetAWS:
		return nil
	default:
		return fmt.Errorf("invalid value for target: %q", t)
	}
}

func validateRegions(t Target, regions []Region) error {
	switch t {
	case TargetLocal:
		if len(regions) > 0 {
			return fmt.Errorf("regions are not allowed for target %q", t)
		}
	case TargetAWS:
		if len(regions) == 0 {
			return fmt.Errorf("at least one region must be provided for target %q", t)
		}

		allowed := allowedRegionsForTarget(t)
		for _, r := range regions {
			valid := false
			for _, ar := range allowed {
				if r == ar {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("region %q is not allowed for target %q", r, t)
			}
		}
	}
	return nil
}

// allowedRegionsForTarget returns the list of valid regions for a given target.
func allowedRegionsForTarget(t Target) []Region {
	switch t {
	case TargetAWS:
		return []Region{RegionAWSUSEast1}
	default:
		return []Region{}
	}
}
