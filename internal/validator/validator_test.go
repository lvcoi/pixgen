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

func TestValidateRejectsMultiCharPaletteKey(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 1, SpriteHeight: 1, Columns: 1, Rows: 1},
		Palette: map[string]string{".": "#00000000", "AB": "#ff0000"},
		Sprites: []schema.Sprite{{ID: "a", X: 0, Y: 0, Pixels: []string{"."}}},
	}
	res := Validate(doc)
	if res.Valid {
		t.Fatal("expected invalid: multi-character palette key should be rejected")
	}
}

func TestValidateRejectsEmptyPaletteKey(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 1, SpriteHeight: 1, Columns: 1, Rows: 1},
		Palette: map[string]string{".": "#00000000", "": "#ff0000"},
		Sprites: []schema.Sprite{{ID: "a", X: 0, Y: 0, Pixels: []string{"."}}},
	}
	res := Validate(doc)
	if res.Valid {
		t.Fatal("expected invalid: empty palette key should be rejected")
	}
}

func TestValidateEmptyIDNoDuplicateConfusion(t *testing.T) {
	// Two sprites with empty IDs should each produce a "required" error,
	// but NOT a "duplicate sprite id" error for the empty string.
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 1, SpriteHeight: 1, Columns: 2, Rows: 1},
		Palette: map[string]string{".": "#00000000"},
		Sprites: []schema.Sprite{
			{ID: "", X: 0, Y: 0, Pixels: []string{"."}},
			{ID: "", X: 1, Y: 0, Pixels: []string{"."}},
		},
	}
	res := Validate(doc)
	if res.Valid {
		t.Fatal("expected invalid due to missing IDs")
	}
	for _, e := range res.Errors {
		if e == `duplicate sprite id ""` {
			t.Errorf("unexpected duplicate-id error for empty string: %q", e)
		}
	}
}
