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
	for _, s := range doc.Sprites {
		if weakSprite(s) {
			r.NeedsRevision = append(r.NeedsRevision, s.ID)
			r.Reasons[s.ID] = "sprite has <=1 opaque pixel row; likely incomplete"
		}
	}
	return r
}

func weakSprite(s schema.Sprite) bool {
	opaqueRows := 0
	for _, row := range s.Pixels {
		if strings.Trim(row, ".") != "" {
			opaqueRows++
		}
	}
	return opaqueRows <= 1
}
