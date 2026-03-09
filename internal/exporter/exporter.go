package exporter

import (
	"fmt"
	"path/filepath"

	"pixgen/internal/packer"
	"pixgen/internal/render"
	"pixgen/internal/revision"
	"pixgen/internal/schema"
	"pixgen/internal/validator"
)

type Metadata struct {
	SheetName   string              `json:"sheet_name"`
	SpriteCount int                 `json:"sprite_count"`
	TagsByID    map[string][]string `json:"tags_by_id"`
}

func Export(doc schema.Document, outDir string) error {
	v := validator.Validate(doc)
	if !v.Valid {
		return fmt.Errorf("cannot export invalid document: %v", v.Errors)
	}
	if err := schema.Save(filepath.Join(outDir, "sheet.json"), doc); err != nil {
		return err
	}

	sheet, err := packer.BuildSheet(doc)
	if err != nil {
		return err
	}
	if err := render.SavePNG(filepath.Join(outDir, "spritesheet.png"), sheet); err != nil {
		return err
	}

	for _, s := range doc.Sprites {
		img, err := render.RenderSprite(s, doc.Palette, doc.Sheet.SpriteWidth, doc.Sheet.SpriteHeight)
		if err != nil {
			return err
		}
		if err := render.SavePNG(filepath.Join(outDir, "sprites", s.ID+".png"), img); err != nil {
			return err
		}
	}

	md := Metadata{SheetName: doc.Sheet.Name, SpriteCount: len(doc.Sprites), TagsByID: map[string][]string{}}
	for _, s := range doc.Sprites {
		md.TagsByID[s.ID] = s.Tags
	}
	if err := schema.Save(filepath.Join(outDir, "metadata.json"), md); err != nil {
		return err
	}

	review := revision.Review(doc)
	if err := schema.Save(filepath.Join(outDir, "review.json"), review); err != nil {
		return err
	}
	return nil
}
