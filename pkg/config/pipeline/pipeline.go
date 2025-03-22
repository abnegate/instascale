package pipeline

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"instascale/pkg/converters"
)

type Provider string

const (
	ProviderGitHub Provider = "github"
	ProviderGitLab Provider = "gitlab"
)

var Providers = []Provider{ProviderGitHub, ProviderGitLab}

// Pipelines represents CI/CD pipeline configuration.
type Pipelines struct {
	Provider Provider `yaml:"provider" json:"provider"`
	Owner    string   `yaml:"owner" json:"owner"`
	Repo     string   `yaml:"repo" json:"repo"`
	Lint     *bool    `yaml:"lint" json:"lint"`
	Test     *bool    `yaml:"test" json:"test"`
	Coverage *bool    `yaml:"coverage" json:"coverage"`
}

// JSONSchema returns a JSON schema for the Pipelines configuration.
func (p *Pipelines) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "CI/CD pipeline configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"provider": {
				Type:        "string",
				Description: "The CI/CD provider.",
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
			"lint": {
				Type:        "boolean",
				Description: "Enable linting.",
			},
			"test": {
				Type:        "boolean",
				Description: "Enable testing.",
			},
			"coverage": {
				Type:        "boolean",
				Description: "Enable coverage.",
			},
		}),
		Required: []string{"provider", "owner", "repo"},
	}
}

// SetDefaults sets the default values for the Pipelines configuration.
func (p *Pipelines) SetDefaults() {
	if p.Lint == nil {
		p.Lint = new(bool)
		*p.Lint = true
	}
	if p.Test == nil {
		p.Test = new(bool)
		*p.Test = true
	}
	if p.Coverage == nil {
		p.Coverage = new(bool)
		*p.Coverage = true
	}
}

// Validate returns an error if the Pipelines configuration is invalid.
func (p *Pipelines) Validate() error {
	if err := validateProvider(p.Provider); err != nil {
		return err
	}
	return nil
}

func validateProvider(p Provider) error {
	switch p {
	case ProviderGitHub, ProviderGitLab:
		return nil
	default:
		return fmt.Errorf("invalid value for provider: %q", p)
	}
}
