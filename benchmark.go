package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"compression-lab/compressor"
)

func RunSingleTest(file string, cType compressor.CompressionType) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	res := run(file, data, cType)

	res.Print()
}

func RunBenchmark(bType string) {
	switch bType {
	case "books":
		runBooksBenchmark()
	case "sensors-data":
		runSensorsBenchmark()
	case "images":
		runImagesBenchmark()
	default:
		log.Fatal("❌ Invalid benchmark type.")
	}
}

func runBooksBenchmark() {
	booksTypes := []string{
		"epub",
		"pdf",
		"txt",
	}

	for _, bType := range booksTypes {
		path := filepath.Join("data", "books", bType)

		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}

		runFiles(path, files, fmt.Sprintf("books_%s", bType))
	}
}

func runSensorsBenchmark() {
	path := filepath.Join("data", "sensor")

	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	runFiles(path, files, "sensors_data")
}

func runImagesBenchmark() {
	imagesTypes := []string{
		"jpg",
		"png",
		"webp",
	}

	for _, iType := range imagesTypes {
		path := filepath.Join("data", "images", iType)

		files, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}

		runFiles(path, files, fmt.Sprintf("images_%s", iType))
	}
}

func runFiles(path string, files []fs.DirEntry, resultFile string) {
	var results Results

	for _, typ := range []compressor.CompressionType{compressor.LZW, compressor.Huffman} {
		for _, f := range files {
			if f.IsDir() {
				continue
			}

			ext := strings.ToLower(filepath.Ext(f.Name()))

			// Skip compressed files
			if ext == ".zip" || ext == ".gz" {
				continue
			}

			fp := filepath.Join(path, f.Name())

			fmt.Println("🔍 Benchmarking", f.Name(), typ)
			data, err := os.ReadFile(fp)
			if err != nil {
				panic(err)
			}

			res := run(f.Name(), data, typ)
			results.results = append(results.results, res)
		}
	}

	results.SaveResultsCSV(resultFile)
}

func run(file string, data []byte, cType compressor.CompressionType) Result {
	var res Result
	res.File = file
	res.Type = cType
	res.OrigSize = len(data)

	var decomp []byte
	encStart := time.Now()

	c := compressor.NewCompressor(cType)

	err := c.Compress(data)
	if err != nil {
		fmt.Println("Error: ", err)
		res.Correct = false
		return res
	}

	res.EncTime = time.Since(encStart)

	decStart := time.Now()
	decomp, err = c.Decompress()
	if err != nil {
		fmt.Println("Error: ", err)
		res.Correct = false
		return res
	}

	res.DecTime = time.Since(decStart)
	res.CompSize = len(c.CompressedData)
	res.CompRatio = float64(res.CompSize) / float64(res.OrigSize) * 100
	res.Correct = string(decomp) == string(data)

	return res
}
