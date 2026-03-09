# AGENT.md

## Project mission

Build a local-first sprite-sheet generation pipeline that lets a user describe pixel art in natural language, generate structured sprite data, validate it, render previews, revise weak sprites, and export a final sprite sheet plus metadata.

The system must be designed for iterative AI-assisted creation, not one-shot image generation.

---

## Primary objective

Scaffold the repository into a usable development baseline with:

1. a strict sprite authoring schema
2. a validator
3. a renderer
4. a sheet packer
5. an export pipeline
6. a repair/revision loop contract
7. a minimal CLI
8. a minimal preview app or static preview output
9. tests, fixtures, and sample assets
10. documentation that makes the workflow easy to continue

---

## Default technical direction

Unless the repo already dictates otherwise, use this default stack:

- **Go** for the core engine, CLI, validation, packing, export, and asset processing
- **TypeScript** only for a small browser preview tool if needed
- **JSON** as the canonical authored data format
- **PNG** as the main preview/export format
- **optional indexed palette format** for compact authoring
- keep the system local-first and file-based

Do not overbuild the first scaffold. The first pass is a strong foundation, not a full studio.

---

## Non-goals for the first scaffold

Do not start with:

- a full web platform
- user accounts
- cloud storage
- database-backed persistence
- distributed workers
- complex plugin systems
- model-specific API wiring unless already requested
- MCP server implementation unless the repo is explicitly for that

The first scaffold should make local generation, validation, preview, and export possible.

---

## Core product shape

The scaffold should support this workflow:

1. user describes a sprite sheet or character set
2. AI produces structured sprite data in JSON
3. validator checks dimensions, palette, layout, and schema correctness
4. renderer produces preview images
5. packer assembles sprites into a sheet
6. review step identifies missing, weak, or duplicate sprites
7. revision step updates only failed sprites
8. exporter emits final sheet and metadata

---

## Canonical file formats

### 1. Authoring format

Define a canonical JSON schema for sprite sheets and sprite assets.

The scaffold must support:

- sheet-level metadata
- sprite dimensions
- palette
- transparent color representation
- per-sprite identity and placement
- per-sprite pixel content
- optional tags, state, animation frame, direction

### 2. Export formats

The scaffold must be able to emit:

- `spritesheet.png`
- `sheet.json`
- `sprites/` individual PNGs
- `metadata.json`

### 3. Internal format rules

Prefer an indexed palette representation over raw full hex per pixel when authoring. The schema should still allow lossless export.

Example shape:

```json
{
  "sheet": {
    "name": "wasteland_npcs",
    "sprite_width": 16,
    "sprite_height": 16,
    "columns": 8,
    "rows": 4
  },
  "palette": {
    ".": "#00000000",
    "A": "#1b1b1b",
    "B": "#4d3b2f",
    "C": "#d7b98e"
  },
  "sprites": [
    {
      "id": "npc_scavenger_idle_south_0",
      "x": 0,
      "y": 0,
      "tags": ["npc", "idle", "south"],
      "pixels": [
        "................",
        "......AA........",
        ".....ABBA......."
      ]
    }
  ]
}
