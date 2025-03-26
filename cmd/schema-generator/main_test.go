package main

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSchemaFile(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tmpDir := t.TempDir()
		outFile := filepath.Join(tmpDir, "test-schema.json")

		err := GenerateSchemaFile(outFile)
		assert.NoError(t, err, "Expected no error generating schema")

		info, statErr := os.Stat(outFile)
		assert.NoError(t, statErr, "Expected the schema file to be created")
		assert.False(t, info.IsDir(), "Expected a file, not a directory")

		data, readErr := os.ReadFile(outFile)
		assert.NoError(t, readErr)

		var v any
		jsonErr := json.Unmarshal(data, &v)
		assert.NoError(t, jsonErr, "Expected valid JSON content")
	})

	t.Run("Failed write", func(t *testing.T) {
		tmpDir := t.TempDir()
		dirPath := filepath.Join(tmpDir, "subdir")

		require.NoError(t, os.MkdirAll(dirPath, 0755))

		err := GenerateSchemaFile(dirPath)
		require.Error(t, err, "Expected an error writing to a directory")
		assert.Contains(t, err.Error(), "error writing schema file", "Should wrap the write error")
	})
}
