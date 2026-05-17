package file

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
)

type Read struct{}

func NewRead() *Read { return &Read{} }

func (r *Read) Name() string { return "file_read" }

func (r *Read) Description() string {
	return "Read the contents of a file at the given path. Returns the file contents as text."
}

func (r *Read) Schema() json.RawMessage {
	return json.RawMessage(`{
  "type": "object",
  "properties": {
    "path": {
      "type": "string",
      "description": "Absolute or relative path to the file."
    }
  },
  "required": ["path"]
}`)
}

type readInput struct {
	Path string `json:"path"`
}

func (r *Read) Run(ctx context.Context, input json.RawMessage) (string, error) {
	var in readInput
	if err := json.Unmarshal(input, &in); err != nil {
		return "", fmt.Errorf("invalid input: %w", err)
	}
	if in.Path == "" {
		return "", fmt.Errorf("invalid input: path is required")
	}
	data, err := os.ReadFile(in.Path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
