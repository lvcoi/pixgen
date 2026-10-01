package packer

import (
	"testing"

	"pixgen/internal/schema"
)

func TestBuildSheetSize(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 2, SpriteHeight: 2, Columns: 2, Rows: 1},
		Palette: map[string]string{".": "#00000000", "A": "#ff0000"},
		Sprites: []schema.Sprite{{ID: "a", X: 0, Y: 0, Pixels: []string{"AA", "AA"}}},
	}
	img, err := BuildSheet(doc)
	if err != nil {
		t.Fatal(err)
	}
	if gotW, gotH := img.Bounds().Dx(), img.Bounds().Dy(); gotW != 4 || gotH != 2 {
		t.Fatalf("got size %dx%d", gotW, gotH)
	}
}

func TestBuildSheetOutOfBoundsError(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 2, SpriteHeight: 2, Columns: 1, Rows: 1},
		Palette: map[string]string{".": "#00000000", "A": "#ff0000"},
		Sprites: []schema.Sprite{{ID: "a", X: 2, Y: 0, Pixels: []string{"AA", "AA"}}},
	}
	_, err := BuildSheet(doc)
	if err == nil {
		t.Fatal("expected error for out-of-bounds sprite")
	}
}

func TestBuildSheetOverlapError(t *testing.T) {
	doc := schema.Document{
		Sheet:   schema.Sheet{Name: "x", SpriteWidth: 2, SpriteHeight: 2, Columns: 2, Rows: 1},
		Palette: map[string]string{".": "#00000000", "A": "#ff0000"},
		Sprites: []schema.Sprite{
			{ID: "a", X: 0, Y: 0, Pixels: []string{"AA", "AA"}},
			{ID: "b", X: 0, Y: 0, Pixels: []string{"..", ".."}},
		},
	}
	_, err := BuildSheet(doc)
	if err == nil {
		t.Fatal("expected error for overlapping sprites")
	}
}
