#!/bin/bash

INPUT_DIR="data/images/jpg"
OUTPUT_PNG="data/images/png"
OUTPUT_WEBP="data/images/webp"

mkdir -p "$OUTPUT_PNG" "$OUTPUT_WEBP"

for img in "$INPUT_DIR"/*; do
  name=$(basename "$img" | cut -d. -f1)

  png_path="$OUTPUT_PNG/$name.png"
  webp_path="$OUTPUT_WEBP/$name.webp"

  if [ ! -f "$png_path" ]; then
    convert "$img" "$png_path" && echo "Convertido para PNG: $png_path"
  fi

  if [ ! -f "$webp_path" ]; then
    convert "$img" "$webp_path" && echo "Convertido para WEBP: $webp_path"
  fi
done
