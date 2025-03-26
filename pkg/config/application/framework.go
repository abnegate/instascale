package application

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"gopkg.in/yaml.v3"
	"instascale/pkg/converters"
)

// FrameworkName defines the application framework name.
type FrameworkName string

const (
	FrameworkExpress       FrameworkName = "express"
	FrameworkExpressLatest string        = "4.17.1"
)

var Frameworks = []FrameworkName{
	FrameworkExpress,
}

var JavaScriptFrameworks = []FrameworkName{
	FrameworkExpress,
}

// Framework defines the application framework.
type Framework struct {
	Name    FrameworkName `yaml:"name" json:"name"`
	Version string        `yaml:"version" json:"version"`
}

// JSONSchema returns a JSON schema for the Framework configuration.
func (f *Framework) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Framework to use for the application.",
		OneOf: []*jsonschema.Schema{
			{
				Type:        "string",
				Description: "The language-specific framework.",
				Enum:        converters.ToAnySlice(Frameworks),
			},
			{
				Type:        "object",
				Description: "The language-specific framework.",
				Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
					"name": {
						Type:        "string",
						Description: "The language-specific framework name.",
						Enum:        converters.ToAnySlice(Frameworks),
					},
					"version": {
						Type:        "string",
						Description: "The language-specific framework version.",
					},
				}),
				Required: []string{"name"},
			},
		},
	}
}

// UnmarshalYAML a YAML node into a Framework.
func (f *Framework) UnmarshalYAML(value *yaml.Node) error {
	switch value.Tag {
	case "!!str":
		if err := value.Decode(&f.Name); err != nil {
			return err
		}
		f.Version = ""
	case "!!map":
		type rawFramework Framework
		var raw rawFramework
		if err := value.Decode(&raw); err != nil {
			return err
		}
		*f = Framework(raw)
	default:
		return fmt.Errorf("unsupported YAML type for framework: %v", value.Tag)
	}
	return nil
}

// SetDefaults initializes any missing Framework fields with default values.
func (f *Framework) SetDefaults() {
	switch f.Name {
	case FrameworkExpress:
		if f.Version == "" {
			f.Version = FrameworkExpressLatest
		}
	}
}

// Validate checks that the Framework configuration is valid.
func (f *Framework) Validate() error {
	switch f.Name {
	case FrameworkExpress:
		if f.Version == "" {
			return fmt.Errorf("framework version cannot be empty")
		}
		return nil
	default:
		return fmt.Errorf("unsupported framework: %q", f.Name)
	}
}
