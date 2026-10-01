package revision

import (
	"strings"

	"pixgen/internal/schema"
)

type Report struct {
	NeedsRevision []string          `json:"needs_revision"`
	Reasons       map[string]string `json:"reasons"`
}

func Review(doc schema.Document) Report {
	r := Report{Reasons: map[string]string{}}
	// Derive the transparent key from the palette; if "." is absent, any
	// non-empty row is treated as opaque (no symbol is assumed transparent).
	transparentKey := ""
	if _, ok := doc.Palette["."]; ok {
		transparentKey = "."
	}
	for _, s := range doc.Sprites {
		if weakSprite(s, transparentKey) {
			r.NeedsRevision = append(r.NeedsRevision, s.ID)
			r.Reasons[s.ID] = "sprite has <=1 opaque pixel row; likely incomplete"
		}
	}
	return r
}

func weakSprite(s schema.Sprite, transparentKey string) bool {
	opaqueRows := 0
	for _, row := range s.Pixels {
		trimmed := row
		if transparentKey != "" {
			trimmed = strings.Trim(row, transparentKey)
		}
		if trimmed != "" {
			opaqueRows++
		}
	}
	return opaqueRows <= 1
}
