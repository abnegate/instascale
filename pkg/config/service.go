package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type Type string

const (
	TypeHttp     Type = "http"
	TypePostgres Type = "postgres"
	TypeMysql    Type = "mysql"
	TypeMongo    Type = "mongo"
	TypeRedis    Type = "redis"
)

// Service represents a single microservice in your application.
type Service struct {
	Name        string            `yaml:"name" json:"name"`
	Description string            `yaml:"description" json:"description"`
	Type        string            `yaml:"type" json:"type" jsonschema:"enum=http,enum=postgres,enum=mysql,enum=mongo,enum=redis"`
	Version     string            `yaml:"version,omitempty" json:"version,omitempty"`
	Port        uint16            `yaml:"port,omitempty" json:"port,omitempty"`
	Env         map[string]string `yaml:"env,omitempty" json:"env,omitempty"`
}

// Validate checks the Service configuration for basic correctness.
func (s *Service) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("service name cannot be empty")
	}
	if s.Type == "" {
		return fmt.Errorf("service type cannot be empty")
	}
	if s.Version == "" {
		return fmt.Errorf("service version cannot be empty")
	}
	return nil
}

// UnmarshalYAML is a custom unmarshaler that expands env variables
func (s *Service) UnmarshalYAML(node *yaml.Node) error {
	// Create a temporary type alias to avoid recursion in UnmarshalYAML.
	type rawService Service
	var raw rawService

	if err := node.Decode(&raw); err != nil {
		return err
	}

	*s = Service(raw)

	// Now do environment variable expansion on s.Env.
	for key, value := range s.Env {
		if !strings.HasPrefix(value, "$") {
			continue
		}
		name := value[1:] // "$POSTGRES_USER" -> "POSTGRES_USER"
		expanded := os.Getenv(name)
		if expanded == "" {
			// If the env var is not set, you might want to throw an error or leave it blank.
			// Here we just warn, but you could decide differently:
			return fmt.Errorf("environment variable %q is not set", name)
		}
		s.Env[key] = expanded
	}
	return nil
}
