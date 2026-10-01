# pixgen

`pixgen` is a local-first sprite sheet generation scaffold for iterative AI-assisted workflows.

## What is included

- Canonical JSON schema for sheet metadata, indexed palette, and sprite pixel grids.
- Validator that checks dimensions, palette usage, IDs, and bounds.
- Renderer + sheet packer using Go's image packages.
- Export pipeline that emits:
  - `spritesheet.png`
  - `sheet.json`
  - `sprites/*.png`
  - `metadata.json`
  - `review.json` (revision loop contract)
- Minimal CLI (`validate`, `export`, `run`).
- Sample fixture and tests.

## Authoring format

See `fixtures/sample_sheet.json` for a complete example.

## Quickstart

```bash
go test ./...
go run ./cmd/pixgen validate -in fixtures/sample_sheet.json
go run ./cmd/pixgen export -in fixtures/sample_sheet.json -out out
```

## CLI

### Validate only

```bash
go run ./cmd/pixgen validate -in fixtures/sample_sheet.json
```

### Export artifacts

```bash
go run ./cmd/pixgen export -in fixtures/sample_sheet.json -out out
```

### Run full pipeline

```bash
go run ./cmd/pixgen run -in fixtures/sample_sheet.json -out out
```

## Preview

Open `out/spritesheet.png` in any image viewer. A lightweight static preview shell is included at `preview/index.html`.
