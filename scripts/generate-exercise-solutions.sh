#!/usr/bin/env bash
# generate-exercise-solutions.sh — Regenerate topic-specific exercise solutions.
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"
go run ./scripts/cmd/generate-exercise-solutions "$ROOT"
