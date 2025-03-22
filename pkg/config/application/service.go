package application

import (
	"fmt"
	"github.com/invopop/jsonschema"
	"gopkg.in/yaml.v3"
	"instascale/pkg/converters"
	"os"
	"strings"
)

// ServiceType defines the available service type.
type ServiceType string

const (
	ServiceTypeHttp     ServiceType = "http"
	ServiceTypePostgres ServiceType = "postgres"
	ServiceTypeMysql    ServiceType = "mysql"
	ServiceTypeMongo    ServiceType = "mongo"
	ServiceTypeRedis    ServiceType = "redis"
)

var ServiceTypes = []ServiceType{
	ServiceTypeHttp,
	ServiceTypePostgres,
	ServiceTypeMysql,
	ServiceTypeMongo,
	ServiceTypeRedis,
}

// Service represents a single service in your application.
type Service struct {
	Name        string            `yaml:"name" json:"name"`
	Type        ServiceType       `yaml:"type" json:"type"`
	Description string            `yaml:"description,omitempty" json:"description,omitempty"`
	Version     string            `yaml:"version,omitempty" json:"version,omitempty"`
	Port        uint16            `yaml:"port,omitempty" json:"port,omitempty"`
	Env         map[string]string `yaml:"env,omitempty" json:"env,omitempty"`
}

// JSONSchema returns a JSON schema for the Service configuration.
func (s *Service) JSONSchema() *jsonschema.Schema {
	return &jsonschema.Schema{
		Type:        "object",
		Description: "Service configuration.",
		Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
			"name": {
				Type:        "string",
				Description: "The service name.",
			},
			"type": {
				Type:        "string",
				Description: "The service type.",
				Enum:        converters.ToAnySlice(ServiceTypes),
			},
			"description": {
				Type:        "string",
				Description: "The service description.",
			},
			"version": {
				Type:        "string",
				Description: "The service version.",
			},
			"port": {
				Type:        "integer",
				Description: "The service port.",
				Minimum:     "1024",
				Maximum:     "65535",
			},
			"env": {
				Type:        "object",
				Description: "The service environment variables.",
				AdditionalProperties: &jsonschema.Schema{
					Type: "string",
				},
			},
		}),
		Required: []string{"name"},
		AllOf: []*jsonschema.Schema{
			{
				If: &jsonschema.Schema{
					Properties: converters.ToOrderedProps(map[string]*jsonschema.Schema{
						"type": {
							Const: ServiceTypeHttp,
						},
					}),
				},
				Then: &jsonschema.Schema{
					Required: []string{"port"},
				},
			},
		},
	}
}

// UnmarshalYAML is a custom unmarshaler that expands env variables
func (s *Service) UnmarshalYAML(node *yaml.Node) error {
	type rawService Service
	var raw rawService

	if err := node.Decode(&raw); err != nil {
		return err
	}

	*s = Service(raw)

	for key, value := range s.Env {
		if !strings.HasPrefix(value, "$") {
			continue
		}
		name := value[1:] // "$POSTGRES_USER" -> "POSTGRES_USER"
		expanded := os.Getenv(name)
		if expanded == "" {
			return fmt.Errorf("environment variable %q is not set", name)
		}
		s.Env[key] = expanded
	}

	return nil
}

// SetDefaults initializes any missing Service fields with default values.
func (s *Service) SetDefaults() {
	switch s.Type {
	case ServiceTypePostgres:
		if s.Port == 0 {
			s.Port = 5432
		}
	case ServiceTypeMysql:
		if s.Port == 0 {
			s.Port = 3306
		}
	case ServiceTypeMongo:
		if s.Port == 0 {
			s.Port = 27017
		}
	case ServiceTypeRedis:
		if s.Port == 0 {
			s.Port = 6379
		}
	}

	if s.Version == "" {
		s.Version = "latest"
	}

	if s.Env == nil {
		s.Env = make(map[string]string)
	}
}

// Validate checks the Service configuration for basic correctness.
func (s *Service) Validate() error {
	if err := validateName(s.Name); err != nil {
		return err
	}
	if err := validateType(*s, s.Type); err != nil {
		return err
	}
	if err := validateVersion(s.Version); err != nil {
		return err
	}
	if err := validatePort(s.Port); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}

func validateType(s Service, t ServiceType) error {
	switch t {
	case ServiceTypeHttp:
		if s.Port == 0 {
			return fmt.Errorf("service port cannot be empty for type %q", t)
		}
		return nil
	case ServiceTypePostgres, ServiceTypeMysql, ServiceTypeMongo, ServiceTypeRedis:
		return nil
	default:
		return fmt.Errorf("unsupported service type: %q", t)
	}
}

func validateVersion(v string) error {
	if v == "" {
		return fmt.Errorf("version cannot be empty")
	}
	return nil
}

func validatePort(p uint16) error {
	if p < 1024 || p > 65535 {
		return fmt.Errorf("port %d is out of range, must be between 1024 and 65535", p)
	}
	return nil
}
