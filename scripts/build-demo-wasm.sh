#!/usr/bin/env bash
# Compile the landing page's runnable Emo samples to WebAssembly for the
# in-browser runner. Run this after editing anything under
# app/views/home/samples/; the produced .wasm files are committed so a plain
# `go build` never needs the Emo toolchain.
#
# Set EMO_BIN to point at a specific emo binary (defaults to `emo` on PATH).
set -euo pipefail

EMO_BIN="${EMO_BIN:-emo}"
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SAMPLES="$ROOT/app/views/home/samples"
OUT="$ROOT/app/assets/public/demo"
STAGE="$ROOT/tmp/demo-wasm"

if ! command -v "$EMO_BIN" >/dev/null 2>&1; then
  echo "build-demo-wasm: emo not found; set EMO_BIN to the emo binary" >&2
  exit 1
fi

mkdir -p "$OUT" "$STAGE"

for dir in "$SAMPLES"/*/; do
  name="$(basename "$dir")"
  [ -f "$dir/main.emo" ] || continue
  (cd "$dir" && "$EMO_BIN" build main.emo --target wasm -o "$STAGE/$name.wasm")
  cp "$STAGE/$name.wasm" "$OUT/$name.wasm"
  echo "wrote $OUT/$name.wasm"
done
