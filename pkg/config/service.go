package config

import "fmt"

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
	Name    string                 `yaml:"name" json:"name"`
	Type    string                 `yaml:"type" json:"type" jsonschema:"enum=http,enum=postgres,enum=mysql,enum=mongo,enum=redis"`
	Version string                 `yaml:"version,omitempty" json:"version,omitempty"`
	Port    uint16                 `yaml:"port,omitempty" json:"port,omitempty"`
	Env     map[string]interface{} `yaml:"env,omitempty" json:"env,omitempty"`
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
