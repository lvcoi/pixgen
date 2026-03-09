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

	for _, s := range doc.Sprites {
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
