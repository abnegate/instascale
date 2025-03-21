package config

import "fmt"

// CI defines the continuous integration schema-generator.
type CI string

const (
	CIGitHub CI = "github"
)

func (ci *CI) Validate() error {
	switch *ci {
	case CIGitHub:
		return nil
	default:
		return fmt.Errorf("unsupported CI: %q", ci)
	}
}
