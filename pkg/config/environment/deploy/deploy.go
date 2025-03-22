package deploy

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"instascale/pkg/converters"
)

// Deploy nests target-related configuration.
type Deploy struct {
	Target  Target   `yaml:"target" json:"target"`
	Regions []Region `yaml:"regions,omitempty" json:"regions,omitempty"`
}

func (d *Deploy) JSONSchema() *jsonschema.Schema {
	var targetRegionRules []*jsonschema.Schema

	for _, tgt := range Targets {
		targetRegionRules = append(targetRegionRules, &jsonschema.Schema{
			If: &jsonschema.Schema{
				Type: "object",
				Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
					"target": {Const: tgt},
				}),
			},
			Then: &jsonschema.Schema{
				Type: "object",
				Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
					"regions": func() *jsonschema.Schema {
						if tgt == TargetLocal {
							// Disallow any regions for local
							return &jsonschema.Schema{
								MaxItems: converters.ToUint64Ptr(0),
							}
						}
						// Otherwise, restrict to allowed regions for this target
						return &jsonschema.Schema{
							Enum: converters.ToAnySlice(allowedRegions(tgt)),
						}
					}(),
				}),
				// Only require "regions" if not local
				Required: func() []string {
					if tgt == TargetLocal {
						return nil
					}
					return []string{"regions"}
				}(),
			},
		})
	}

	return &jsonschema.Schema{
		Type:        "object",
		Description: "Deploy configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"target": {
				Type:        "string",
				Description: "The deployment target.",
				Enum:        converters.ToAnySlice(Targets),
			},
			"regions": {
				Type:        "array",
				Description: "List of regions to deploy to.",
				Items:       &jsonschema.Schema{Type: "string"},
			},
		}),
		Required: []string{"target"},
		AllOf:    targetRegionRules,
	}
}

// SetDefaults initializes any missing Deploy fields with default values.
func (d *Deploy) SetDefaults() {
	if d.Target == "" {
		d.Target = TargetLocal
	}
}

// Validate checks that the Deploy configuration is valid.
func (d *Deploy) Validate() error {
	if err := validateTarget(d.Target); err != nil {
		return err
	}
	if err := validateRegions(d.Target, d.Regions); err != nil {
		return err
	}
	return nil
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
	allowedRegions := allowedRegions(t)

	switch t {
	case TargetLocal:
		if len(regions) > 0 {
			return fmt.Errorf("regions are not allowedRegions for target %q", t)
		}
	case TargetAWS:
		if len(regions) == 0 {
			return fmt.Errorf("at least one region must be provided for target %q", t)
		}

		for _, r := range regions {
			valid := false
			for _, ar := range allowedRegions {
				if r == ar {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("region %q is not allowedRegions for target %q", r, t)
			}
		}
	}
	return nil
}
