package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"pixgen/internal/exporter"
	"pixgen/internal/pipeline"
	"pixgen/internal/schema"
	"pixgen/internal/validator"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "validate":
		runValidate(os.Args[2:])
	case "export":
		runExport(os.Args[2:])
	case "run":
		runPipeline(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}
}

func runValidate(args []string) {
	fs := flag.NewFlagSet("validate", flag.ExitOnError)
	in := fs.String("in", "", "path to sheet json")
	_ = fs.Parse(args)
	doc := mustLoad(*in)
	out := validator.Validate(doc)
	printJSON(out)
	if !out.Valid {
		os.Exit(2)
	}
}

func runExport(args []string) {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	in := fs.String("in", "", "path to sheet json")
	outDir := fs.String("out", "out", "output directory")
	_ = fs.Parse(args)
	doc := mustLoad(*in)
	if err := exporter.Export(doc, *outDir); err != nil {
		fmt.Fprintln(os.Stderr, "export failed:", err)
		os.Exit(1)
	}
	fmt.Println("exported assets to", filepath.Clean(*outDir))
}

func runPipeline(args []string) {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	in := fs.String("in", "", "path to sheet json")
	outDir := fs.String("out", "out", "output directory")
	_ = fs.Parse(args)
	doc := mustLoad(*in)
	res, err := pipeline.Run(doc, *outDir)
	printJSON(res)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !res.Validation.Valid {
		os.Exit(2)
	}
}

func mustLoad(path string) schema.Document {
	if path == "" {
		fmt.Fprintln(os.Stderr, "-in is required")
		os.Exit(1)
	}
	doc, err := schema.Load(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return doc
}

func printJSON(v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}

func usage() {
	fmt.Println("pixgen <command>\n\nCommands:\n  validate -in <sheet.json>\n  export -in <sheet.json> -out <dir>\n  run -in <sheet.json> -out <dir>")
}
