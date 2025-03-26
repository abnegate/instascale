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

	c.VCS.SetDefaults()
	c.Pipeline.SetDefaults()
}

// Validate calls the Validate methods on nested sections.
func (c *Config) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	if err := c.Application.Validate(); err != nil {
		return err
	}
	for _, env := range c.Environments {
		if err := env.Validate(); err != nil {
			return fmt.Errorf("invalid environment %s: %s", c.Name, err)
		}
	}

	if c.VCS != nil {
		if err := c.VCS.Validate(); err != nil {
			return err
		}
	}
	if c.Pipeline != nil {
		if err := c.Pipeline.Validate(); err != nil {
			return err
		}
	}
	return nil
}
