#!/bin/bash

# Diretórios
EPUB_DIR="data/documents/epub"
TXT_DIR="data/documents/txt"
PDF_DIR="data/documents/pdf"

# Cria os diretórios se não existirem
mkdir -p "$TXT_DIR" "$PDF_DIR"

# Percorre todos os arquivos .epub
for epub_file in "$EPUB_DIR"/*.epub; do
  # Remove o caminho e extensão
  base_name=$(basename "$epub_file" .epub)

  # Caminhos esperados
  txt_file="$TXT_DIR/$base_name.txt"
  pdf_file="$PDF_DIR/$base_name.pdf"

  # Verifica e converte para .txt
  if [ ! -f "$txt_file" ]; then
    echo "Convertendo $base_name.epub para TXT..."
    ebook-convert "$epub_file" "$txt_file"
  else
    echo "[✓] TXT já existe: $txt_file"
  fi

  # Verifica e converte para .pdf
  if [ ! -f "$pdf_file" ]; then
    echo "Convertendo $base_name.epub para PDF..."
    ebook-convert "$epub_file" "$pdf_file"
  else
    echo "[✓] PDF já existe: $pdf_file"
  fi

  echo "----------------------------------"
done
