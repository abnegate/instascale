package config

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"instascale/pkg/config/application"
	"instascale/pkg/config/environment"
	"instascale/pkg/config/pipeline"
	"instascale/pkg/config/vcs"
	"instascale/pkg/converters"
)

// Section defines the minimal behavior every configuration section must implement.
type Section interface {
	SetDefaults()
	Validate() error
	JSONSchema() *jsonschema.Schema
}

// Config defines the project configuration.
type Config struct {
	Name         string                             `yaml:"name" json:"name"`
	Application  application.Application            `yaml:"application" json:"application"`
	Environments map[string]environment.Environment `yaml:"environments" json:"environments"`
	VCS          *vcs.VCS                           `yaml:"vcs,omitempty" json:"vcs,omitempty"`
	Pipeline     *pipeline.Pipeline                 `yaml:"pipeline,omitempty" json:"pipeline,omitempty"`
}

// JSONSchema returns a JSON schema for the Config configuration.
func (c *Config) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Project configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"name": {
				Type:        "string",
				Description: "The project name.",
			},
			"vcs":         (&vcs.VCS{}).JSONSchema(),
			"pipeline":    (&pipeline.Pipeline{}).JSONSchema(),
			"application": (&application.Application{}).JSONSchema(),
			"environments": {
				Type:                 "object",
				Description:          "The project environments.",
				AdditionalProperties: (&environment.Environment{}).JSONSchema(),
			},
		}),
		Required: []string{
			"name",
			"application",
			"environments",
		},
	}
}

// SetDefaults sets the default values for the configuration.
func (c *Config) SetDefaults() {
	c.Application.SetDefaults()

	for k, e := range c.Environments {
		e.SetDefaults()
		c.Environments[k] = e
	}

	if c.VCS != nil {
		c.VCS.SetDefaults()
	}
	if c.Pipeline != nil {
		c.Pipeline.SetDefaults()
	}
}

// Validate calls the Validate methods on nested sections.
func (c *Config) Validate() error {
	if err := validateName(c.Name); err != nil {
		return err
	}
	if err := validateApplication(c.Application); err != nil {
		return err
	}
	if err := validateEnvironments(c.Environments); err != nil {
		return err
	}
	if err := validateVCS(c.VCS); err != nil {
		return err
	}
	if err := validatePipeline(c.Pipeline); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("name is required")
	}
	return nil
}

func validateApplication(a application.Application) error {
	if err := a.Validate(); err != nil {
		return err
	}
	return nil
}

func validateEnvironments(envs map[string]environment.Environment) error {
	for key, env := range envs {
		if err := env.Validate(); err != nil {
			return fmt.Errorf("invalid environment %q: %s", key, err)
		}
	}
	return nil
}

func validateVCS(v *vcs.VCS) error {
	if v == nil {
		return nil
	}
	if err := (*v).Validate(); err != nil {
		return err
	}
	return nil
}

func validatePipeline(p *pipeline.Pipeline) error {
	if p == nil {
		return nil
	}
	if err := (*p).Validate(); err != nil {
		return err
	}
	return nil
}
