package application

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"gopkg.in/yaml.v3"
	"instascale/pkg/converters"
)

// LanguageName defines the application programming language name.
type LanguageName string

const (
	LanguageJavaScript LanguageName = "javascript"
)

var LanguageNames = []LanguageName{LanguageJavaScript}

// LanguageVersion defines the application programming language version.
type LanguageVersion string

const (
	LanguageVersionES6 LanguageVersion = "es6"
)

var LanguageVersions = []LanguageVersion{LanguageVersionES6}

// Language defines the application programming language.
type Language struct {
	Name      LanguageName    `yaml:"name" json:"name"`
	Version   LanguageVersion `yaml:"version,omitempty" json:"version,omitempty"`
	Framework Framework       `yaml:"framework" json:"framework"`
}

func (l *Language) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Language configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"name": {
				Type:        "string",
				Description: "The programming language name.",
				Enum:        converters.ToAnySlice(LanguageNames),
			},
			"version": {
				Type:        "string",
				Description: "The programming language version.",
				Enum:        converters.ToAnySlice(LanguageVersions),
			},
			"framework": (&Framework{}).JSONSchema(),
		}),
		Required: []string{"name", "framework"},
		AllOf: []*jsonschema.Schema{
			{
				If: &jsonschema.Schema{
					Type: "object",
					Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
						"name": {
							Const: LanguageJavaScript,
						},
					}),
				},
				Then: &jsonschema.Schema{
					Type: "object",
					Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
						"version": {
							Type: "string",
							Enum: []any{LanguageVersionES6},
						},
					}),
				},
			},
		},
	}
}

// UnmarshalYAML a YAML node into a Framework.
func (l *Language) UnmarshalYAML(value *yaml.Node) error {
	switch value.Tag {
	case "!!str":
		if err := value.Decode(&l.Name); err != nil {
			return err
		}
		l.Version = ""
		l.Framework = Framework{}
	case "!!map":
		type rawLanguage Language
		var raw rawLanguage
		if err := value.Decode(&raw); err != nil {
			return err
		}
		*l = Language(raw)
	default:
		return fmt.Errorf("unsupported YAML type for framework: %v", value.Tag)
	}
	return nil
}

// SetDefaults initializes any missing Language fields with default values.
func (l *Language) SetDefaults() {
	switch l.Name {
	case LanguageJavaScript:
		if l.Version == "" {
			l.Version = LanguageVersionES6
		}
	}

	l.Framework.SetDefaults()
}

// Validate checks that the Language configuration is valid.
func (l *Language) Validate() error {
	allowedFrameworks := allowedFrameworksForLanguage(l)
	for _, allowedFramework := range allowedFrameworks {
		if l.Framework.Name == allowedFramework {
			break
		}
		return fmt.Errorf("framework %q is not allowed for language %q", l.Framework.Name, l.Name)
	}

	allowedVersions := allowedVersionsForLanguage(l)
	for _, allowedVersion := range allowedVersions {
		if l.Version == allowedVersion {
			break
		}
		return fmt.Errorf("version %q is not allowed for language %q", l.Version, l.Name)
	}

	if err := l.Framework.Validate(); err != nil {
		return err
	}

	return nil
}

func allowedVersionsForLanguage(l *Language) []LanguageVersion {
	switch l.Name {
	case LanguageJavaScript:
		return []LanguageVersion{LanguageVersionES6}
	default:
		return []LanguageVersion{}
	}
}

func allowedFrameworksForLanguage(l *Language) []FrameworkName {
	switch l.Name {
	case LanguageJavaScript:
		return []FrameworkName{FrameworkExpress}
	default:
		return []FrameworkName{}
	}
}
