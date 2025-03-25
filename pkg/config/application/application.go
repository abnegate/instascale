package application

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"instascale/pkg/converters"
)

// Application nests application-specific configuration.
type Application struct {
	Language Language  `yaml:"language" json:"language"`
	Services []Service `yaml:"services,omitempty" json:"services,omitempty"`
	Version  string    `yaml:"version,omitempty" json:"version,omitempty"`
}

func (a *Application) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Application configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"version": {
				Type:        "string",
				Description: "The application version.",
			},
			"services": {
				Type:        "array",
				Description: "List of services in the application.",
				Items:       (&Service{}).JSONSchema(),
			},
		}),
		Required: []string{"language", "services"},
	}
}

// SetDefaults initializes any missing Application fields with default values.
func (a *Application) SetDefaults() {
	a.Language.SetDefaults()

	for i := range a.Services {
		a.Services[i].SetDefaults()
	}

	if a.Version == "" {
		a.Version = "1.0.0"
	}
}

// Validate checks that the Application configuration is valid.
func (a *Application) Validate() error {
	if err := validateLanguage(a.Language); err != nil {
		return err
	}
	if err := validateServices(a.Services); err != nil {
		return err
	}
	return nil
}

func validateLanguage(l Language) error {
	if err := l.Validate(); err != nil {
		return err
	}
	return nil

}

func validateServices(s []Service) error {
	for i := range s {
		if err := s[i].Validate(); err != nil {
			return fmt.Errorf("invalid service %q: %w", s[i].Name, err)
		}
	}
	return nil
}
