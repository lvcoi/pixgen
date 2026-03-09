package validator

import (
	"fmt"
	"sort"
	"strings"

	"pixgen/internal/schema"
)

type Result struct {
	Valid    bool     `json:"valid"`
	Errors   []string `json:"errors"`
	Warnings []string `json:"warnings"`
}

func Validate(doc schema.Document) Result {
	r := Result{Valid: true}

	if doc.Sheet.Name == "" {
		r.Errors = append(r.Errors, "sheet.name is required")
	}
	if doc.Sheet.SpriteWidth <= 0 || doc.Sheet.SpriteHeight <= 0 {
		r.Errors = append(r.Errors, "sheet sprite dimensions must be > 0")
	}
	if doc.Sheet.Columns <= 0 || doc.Sheet.Rows <= 0 {
		r.Errors = append(r.Errors, "sheet grid dimensions must be > 0")
	}
	if len(doc.Palette) == 0 {
		r.Errors = append(r.Errors, "palette must define at least one color")
	}

	ids := map[string]struct{}{}
	for i, s := range doc.Sprites {
		if s.ID == "" {
			r.Errors = append(r.Errors, fmt.Sprintf("sprites[%d].id is required", i))
		}
		if _, ok := ids[s.ID]; ok {
			r.Errors = append(r.Errors, fmt.Sprintf("duplicate sprite id %q", s.ID))
		}
		ids[s.ID] = struct{}{}

		if s.X < 0 || s.X >= doc.Sheet.Columns || s.Y < 0 || s.Y >= doc.Sheet.Rows {
			r.Errors = append(r.Errors, fmt.Sprintf("sprite %q coordinates (%d,%d) out of sheet bounds", s.ID, s.X, s.Y))
		}
		if len(s.Pixels) != doc.Sheet.SpriteHeight {
			r.Errors = append(r.Errors, fmt.Sprintf("sprite %q has %d pixel rows, expected %d", s.ID, len(s.Pixels), doc.Sheet.SpriteHeight))
			continue
		}
		for row, line := range s.Pixels {
			if len(line) != doc.Sheet.SpriteWidth {
				r.Errors = append(r.Errors, fmt.Sprintf("sprite %q row %d has width %d, expected %d", s.ID, row, len(line), doc.Sheet.SpriteWidth))
				continue
			}
			for _, c := range line {
				if _, ok := doc.Palette[string(c)]; !ok {
					r.Errors = append(r.Errors, fmt.Sprintf("sprite %q row %d uses undefined palette key %q", s.ID, row, string(c)))
				}
			}
		}

		if allTransparent(s, ".") {
			r.Warnings = append(r.Warnings, fmt.Sprintf("sprite %q is fully transparent", s.ID))
		}
	}

	if len(doc.Sprites) == 0 {
		r.Warnings = append(r.Warnings, "document has no sprites")
	}

	r.Valid = len(r.Errors) == 0
	sort.Strings(r.Errors)
	sort.Strings(r.Warnings)
	return r
}

func allTransparent(s schema.Sprite, key string) bool {
	for _, row := range s.Pixels {
		if strings.Trim(row, key) != "" {
			return false
		}
	}
	return true
}
