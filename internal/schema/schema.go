package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Document is the canonical authoring model for sprite sheets.
type Document struct {
	Sheet   Sheet             `json:"sheet"`
	Palette map[string]string `json:"palette"`
	Sprites []Sprite          `json:"sprites"`
}

type Sheet struct {
	Name         string `json:"name"`
	SpriteWidth  int    `json:"sprite_width"`
	SpriteHeight int    `json:"sprite_height"`
	Columns      int    `json:"columns"`
	Rows         int    `json:"rows"`
}

type Sprite struct {
	ID     string   `json:"id"`
	X      int      `json:"x"`
	Y      int      `json:"y"`
	Tags   []string `json:"tags,omitempty"`
	Pixels []string `json:"pixels"`
}

func Load(path string) (Document, error) {
	var doc Document
	b, err := os.ReadFile(path)
	if err != nil {
		return doc, fmt.Errorf("read input: %w", err)
	}
	if err := json.Unmarshal(b, &doc); err != nil {
		return doc, fmt.Errorf("decode json: %w", err)
	}
	return doc, nil
}

func Save(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}
	if err := os.WriteFile(path, b, 0o644); err != nil {
		return fmt.Errorf("write json: %w", err)
	}
	return nil
}
