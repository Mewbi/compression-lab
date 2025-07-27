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
	splitted := flag.Bool("splitted", false, "Process file line-by-line instead of full content")
	flag.Parse()

	if *file != "" && *comp != "" {

		var cType compressor.CompressionType

		switch *comp {
		case "lzw":
			cType = compressor.LZW
		case "huffman":
			cType = compressor.Huffman
		default:
			panic("❌ Invalid compression type. Use 'lzw' or 'huffman'.")
		}

		if *splitted {
			RunSingleTestSplitted(*file, cType)
			return
		}

		RunSingleTest(*file, cType)

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
