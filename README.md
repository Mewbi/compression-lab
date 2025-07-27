# 🗜️ Compression Lab

This project serves as a small laboratory for experimenting with lossless compression algorithms in the **Data Structures & Algorithms** class at UFABC.  It provides Go implementations of **LZW** and **Huffman** encoders/decoders along with a set of scripts and sample data so their performance can be measured.

## 📂 Data

Benchmark files live under the `data` directory and come from several open sources:

- **Books** – downloaded from the Project Gutenberg catalog using the helper scripts in `scripts/`.
- **Sensor data** – compressed archives retrieved from the following public datasets:
  - [MIT Lab Data](https://db.csail.mit.edu/labdata/labdata.html)
  - [Gas Sensor Array under Dynamic Gas Mixtures](https://archive.ics.uci.edu/dataset/322/gas+sensor+array+under+dynamic+gas+mixtures)
  - [Historical Hourly Weather Data](https://www.kaggle.com/datasets/selfishgene/historical-hourly-weather-data)
- **Images** – a collection of photos from [Lorem Picsum](https://picsum.photos) converted to `jpg`, `png` and `webp` formats.

### 🔍 Observation

The data from `Gas Sensor Array under Dynamic Gas Mixtures` are two big files that can't be directly upload to git (even compressed).

In that way, the two files from this source (`ethylene_CO.txt` and `ethylene_methane.txt`) are splitted in multiple files that can be
recovered using the following commands:

```sh
cat data/sensor/ethylene_CO_part_* > data/sensor/ethylene_CO.txt
cat data/sensor/ethylene_methane_part_* > data/sensor/ethylene_methane.txt
```

## 🛠️ Scripts

The `scripts` folder contains utilities for gathering and preparing the datasets:

- `book_download.py` – downloads the top 100 Project Gutenberg books (EPUB and TXT).
- `book_converter.sh` – converts the downloaded EPUBs to TXT and PDF using `ebook-convert`.
- `image_download.sh` – fetches 100 random images as JPEG files.
- `image_converter.sh` – converts those JPEGs to PNG and WEBP via `convert`.

Running these scripts will populate the subdirectories of `data/` accordingly.

## 🚀 Using the Go program

The entry point for the compression experiments is `main.go`.  It can compress a single file or benchmark all files of a given category.  Build it with:

```bash
go build -o compression-tester
```

### 🔎 Single file test

```bash
./compression-lab -file path/to/file -type lzw      # or huffman
```
For large text files you can enable **splitted** mode to process the input line by line and save results to `results/<file>_splitted_<type>.csv`.

```bash
./compression-lab -file path/to/file -type lzw -splitted
```

Example output:

```text
📁 File: data/sensor/humidity.csv
📦 Type: LZW
📏 Original Size: 9075077
📉 Compressed Size: 2318213
📊 Compression Ratio: 25.54%
⏱️ Encode Time: 104.371078ms
⏱️ Decode Time: 74.844891ms
✅ Correct: true
🎲 Shannon Entropy: 3.3087
🎲 Max Entropy: 5.8826
🔁 Repetition Factor: 0.9999
```


### 🏁 Benchmark mode

```bash
./compression-lab -benchmark books
./compression-lab -benchmark sensors-data
./compression-lab -benchmark images
```

Benchmark runs output per-file statistics and also generate a CSV report summarizing the results.
