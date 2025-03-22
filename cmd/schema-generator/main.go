package main

import (
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
	"instascale/pkg/config"
	"os"
)

func main() {
	schema := jsonschema.Reflect(&config.Config{})

	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error marshaling schema: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("schemas/instascale.schema.json", data, 0644); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error writing schema file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Schema generated successfully in instascale.schema.json")
}
