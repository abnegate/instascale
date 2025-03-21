package config

import "fmt"

// Service represents a single microservice in your application.
type Service struct {
	Name  string `yaml:"name" json:"name"`
	Image string `yaml:"image,omitempty" json:"image,omitempty"`
}

// Validate checks the Service configuration for basic correctness.
func (s *Service) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("service name cannot be empty")
	}
	return nil
}
