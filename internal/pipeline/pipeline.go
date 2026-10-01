package pipeline

import (
	"pixgen/internal/exporter"
	"pixgen/internal/revision"
	"pixgen/internal/schema"
	"pixgen/internal/validator"
)

type RunResult struct {
	Validation validator.Result `json:"validation"`
	Review     revision.Report  `json:"review"`
}

func Run(doc schema.Document, outDir string) (RunResult, error) {
	v := validator.Validate(doc)
	r := revision.Review(doc)
	if v.Valid {
		if err := exporter.Export(doc, outDir); err != nil {
			return RunResult{Validation: v, Review: r}, err
		}
	}
	return RunResult{Validation: v, Review: r}, nil
}
