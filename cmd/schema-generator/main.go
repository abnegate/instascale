package main

import (
	"encoding/json"
	"fmt"
	"instascale/pkg/config"
	"os"
)

// GenerateSchemaFile writes the Config JSON schema to the specified file.
func GenerateSchemaFile(filename string) error {
	schema := (&config.Config{}).JSONSchema()

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling schema: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("error writing schema file: %w", err)
	}

	return nil
}

func main() {
	const outFile = "schemas/instascale.schema.json"

	if err := GenerateSchemaFile(outFile); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Println("Schema generated successfully in instascale.schema.json")
}
