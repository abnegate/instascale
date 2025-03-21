package config

import "fmt"

// Language defines the programming language.
type Language string

const (
	LanguageJavaScript Language = "javascript"
)

// Framework defines the application framework.
type Framework string

const (
	FrameworkExpress Framework = "express"
)

// Application nests application-specific configuration.
type Application struct {
	Language  Language  `yaml:"language" json:"language" jsonschema:"enum=javascript"`
	Framework Framework `yaml:"framework" json:"framework" jsonschema:"enum=express"`
	Services  []string  `yaml:"services,omitempty" json:"services,omitempty"`
}

// Validate checks that the Application configuration is valid.
func (a *Application) Validate() error {
	if err := validateLanguage(a.Language); err != nil {
		return err
	}
	return validateFramework(a.Language, a.Framework)
}

func validateLanguage(lang Language) error {
	switch lang {
	case LanguageJavaScript:
		return nil
	default:
		return fmt.Errorf("unsupported language: %q", lang)
	}
}

func validateFramework(lang Language, fw Framework) error {
	switch lang {
	case LanguageJavaScript:
		switch fw {
		case FrameworkExpress:
			return nil
		default:
			return fmt.Errorf("invalid framework %q for language %q", fw, lang)
		}
	default:
		return fmt.Errorf("unsupported language for framework validation: %q", lang)
	}
}
