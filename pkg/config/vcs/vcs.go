package vcs

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"instascale/pkg/converters"
)

type Provider string

const (
	ProviderGitHub Provider = "github"
)

var Providers = []Provider{ProviderGitHub}

// VCS represents Version Control System configuration.
type VCS struct {
	Provider   Provider `yaml:"provider" json:"provider"`
	Owner      string   `yaml:"owner" json:"owner"`
	Repo       string   `yaml:"repo" json:"repo"`
	Visibility string   `yaml:"visibility" json:"visibility"`
}

// JSONSchema returns a JSON schema for the VCS configuration.
func (v *VCS) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Version Control System configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"provider": {
				Type:        "string",
				Description: "The VCS provider.",
				Enum:        converters.ToAnySlice(Providers),
			},
			"owner": {
				Type:        "string",
				Description: "The repository owner.",
			},
			"repo": {
				Type:        "string",
				Description: "The repository name.",
			},
			"visibility": {
				Type:        "string",
				Description: "The repository visibility.",
			},
		}),
		Required: []string{"provider", "owner", "repo"},
	}
}

// SetDefaults sets the default values for the VCS configuration.
func (v *VCS) SetDefaults() {
	if v.Visibility == "" {
		v.Visibility = "public"
	}
}

// Validate the VCS configuration.
func (v *VCS) Validate() error {
	switch v.Provider {
	case ProviderGitHub:
		return nil
	default:
		return fmt.Errorf("invalid provider: %s", v.Provider)
	}
}
