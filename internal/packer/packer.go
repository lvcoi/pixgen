package packer

import (
	"fmt"
	"image"

	"pixgen/internal/render"
	"pixgen/internal/schema"
)

func BuildSheet(doc schema.Document) (*image.RGBA, error) {
	sw, sh := doc.Sheet.SpriteWidth, doc.Sheet.SpriteHeight
	out := image.NewRGBA(image.Rect(0, 0, sw*doc.Sheet.Columns, sh*doc.Sheet.Rows))

	// Track occupied grid cells to detect overlapping sprites.
	occupied := make(map[[2]int]string, len(doc.Sprites))

	for _, s := range doc.Sprites {
		// Bounds check: ensure sprite grid coordinates are within the sheet.
		if s.X < 0 || s.X >= doc.Sheet.Columns || s.Y < 0 || s.Y >= doc.Sheet.Rows {
			return nil, fmt.Errorf("sprite %q coordinates (%d,%d) are out of sheet bounds (%dx%d)",
				s.ID, s.X, s.Y, doc.Sheet.Columns, doc.Sheet.Rows)
		}
		// Overlap check: ensure no two sprites occupy the same grid cell.
		cell := [2]int{s.X, s.Y}
		if prev, ok := occupied[cell]; ok {
			return nil, fmt.Errorf("sprite %q overlaps with sprite %q at grid cell (%d,%d)", s.ID, prev, s.X, s.Y)
		}
		occupied[cell] = s.ID

		spriteImg, err := render.RenderSprite(s, doc.Palette, sw, sh)
		if err != nil {
			return nil, fmt.Errorf("render sprite %q: %w", s.ID, err)
		}
		offX, offY := s.X*sw, s.Y*sh
		for y := 0; y < sh; y++ {
			for x := 0; x < sw; x++ {
				out.Set(offX+x, offY+y, spriteImg.At(x, y))
			}
		}
	}
	return out, nil
}
