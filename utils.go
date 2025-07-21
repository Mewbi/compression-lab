package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"compression-lab/compressor"
)

type Result struct {
	File      string
	Type      compressor.CompressionType
	OrigSize  int
	CompSize  int
	CompRatio float64
	EncTime   time.Duration
	DecTime   time.Duration
	Correct   bool
}

type Results struct {
	results []Result
}

func (r *Results) SaveResultsCSV(name string) {
	fileName := fmt.Sprintf("results_%s_.csv", name)
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"File", "Type", "OrigSize", "CompSize", "Ratio%", "EncTime(ms)", "DecTime(ms)", "Correct"})

	for _, result := range r.results {
		w.Write([]string{
			result.File,
			result.Type.String(),
			strconv.Itoa(result.OrigSize),
			strconv.Itoa(result.CompSize),
			fmt.Sprintf("%.2f", result.CompRatio),
			fmt.Sprintf("%.3f", float64(result.EncTime.Microseconds())/1000),
			fmt.Sprintf("%.3f", float64(result.DecTime.Microseconds())/1000),
			strconv.FormatBool(result.Correct),
		})
	}

	fmt.Println("✅ Results saved to", fileName)
}

func (r *Result) Print() {
	fmt.Println("📁 File:", r.File)
	fmt.Println("📦 Type:", r.Type)
	fmt.Println("📏 Original Size:", r.OrigSize)
	fmt.Println("📉 Compressed Size:", r.CompSize)
	fmt.Printf("📊 Compression Ratio: %.2f%%\n", r.CompRatio)
	fmt.Println("⏱️ Encode Time:", r.EncTime)
	fmt.Println("⏱️ Decode Time:", r.DecTime)
	fmt.Println("✅ Correct:", r.Correct)
}
