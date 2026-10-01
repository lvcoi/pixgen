package validator

import (
	"testing"

	"pixgen/internal/schema"
)

func TestValidatePassesValidDocument(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 2, SpriteHeight: 2, Columns: 1, Rows: 1},
		Palette: map[string]string{".": "#00000000", "A": "#ffffff"},
		Sprites: []schema.Sprite{{ID: "a", X: 0, Y: 0, Pixels: []string{"..", "AA"}}},
	}
	res := Validate(doc)
	if !res.Valid {
		t.Fatalf("expected valid, got errors: %v", res.Errors)
	}
}

func TestValidateCatchesBadPaletteUse(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 1, SpriteHeight: 1, Columns: 1, Rows: 1},
		Palette: map[string]string{".": "#00000000"},
		Sprites: []schema.Sprite{{ID: "a", X: 0, Y: 0, Pixels: []string{"A"}}},
	}
	res := Validate(doc)
	if res.Valid {
		t.Fatal("expected invalid doc")
	}
}

func TestValidateRejectsColorWithoutHash(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 1, SpriteHeight: 1, Columns: 1, Rows: 1},
		Palette: map[string]string{".": "00000000"},
		Sprites: []schema.Sprite{{ID: "a", X: 0, Y: 0, Pixels: []string{"."}}},
	}
	res := Validate(doc)
	if res.Valid {
		t.Fatal("expected invalid: color without leading '#' should be rejected")
	}
}

func TestValidateAcceptsColorWithHash(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 2, SpriteHeight: 1, Columns: 1, Rows: 1},
		Palette: map[string]string{".": "#00000000", "A": "#ff0000"},
		Sprites: []schema.Sprite{{ID: "a", X: 0, Y: 0, Pixels: []string{".A"}}},
	}
	res := Validate(doc)
	if !res.Valid {
		t.Fatalf("expected valid, got errors: %v", res.Errors)
	}
}
