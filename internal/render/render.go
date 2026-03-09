package render

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"pixgen/internal/schema"
)

func RenderSprite(sprite schema.Sprite, palette map[string]string, cellW, cellH int) (*image.RGBA, error) {
	img := image.NewRGBA(image.Rect(0, 0, cellW, cellH))
	for y, row := range sprite.Pixels {
		for x, c := range []rune(row) {
			rgba, err := parseRGBA(palette[string(c)])
			if err != nil {
				return nil, fmt.Errorf("sprite %s parse color %q: %w", sprite.ID, string(c), err)
			}
			img.Set(x, y, rgba)
		}
	}
	return img, nil
}

func SavePNG(path string, img image.Image) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create png: %w", err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return fmt.Errorf("encode png: %w", err)
	}
	return nil
}

func parseRGBA(hex string) (color.RGBA, error) {
	h := strings.TrimPrefix(hex, "#")
	if len(h) == 6 {
		h += "ff"
	}
	if len(h) != 8 {
		return color.RGBA{}, fmt.Errorf("expected #RRGGBBAA or #RRGGBB")
	}
	n, err := strconv.ParseUint(h, 16, 32)
	if err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{R: uint8(n >> 24), G: uint8(n >> 16), B: uint8(n >> 8), A: uint8(n)}, nil
}
