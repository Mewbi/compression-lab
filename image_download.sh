#!/bin/bash

DOWNLOAD_DIR="data/images/jpg"
mkdir -p "$DOWNLOAD_DIR"

echo "Baixando 100 imagens aleatórias do Lorem Picsum..."

for i in $(seq 1 100); do
  filename=$(printf "%03d.jpg" "$i")
  filepath="$DOWNLOAD_DIR/$filename"

  if [ ! -f "$filepath" ]; then
    curl -Ls "https://picsum.photos/seed/$i/800/600" -o "$filepath"
    echo "Baixada: $filepath"
  fi
done

echo "Download concluído! As imagens estão na pasta '$DOWNLOAD_DIR'."
