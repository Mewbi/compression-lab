package main

import (
	"flag"
	"fmt"

	"compression-lab/compressor"
)

func main() {
	file := flag.String("file", "", "Path to file to compress")
	comp := flag.String("type", "", "Compression type: lzw or huffman")
	bench := flag.String("benchmark", "", "Benchmark directory: books, sensors-data, images")
	flag.Parse()

	if *file != "" && *comp != "" {
		switch *comp {

		case "lzw":
			RunSingleTest(*file, compressor.LZW)
		case "huffman":
			RunSingleTest(*file, compressor.Huffman)
		default:
			panic("❌ Invalid compression type. Use 'lzw' or 'huffman'.")
		}

		return
	}

	if *bench != "" {
		RunBenchmark(*bench)
		return
	}

	fmt.Println("Usage:")
	fmt.Println("  ./main -file <path> -type <lzw|huffman>")
	fmt.Println("  ./main -benchmark <books|sensors-data|images>")
}
