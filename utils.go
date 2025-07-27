package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"compression-lab/compressor"
)

const _REPETITION_FACTOR_K = 3

type Result struct {
	File             string
	Type             compressor.CompressionType
	OrigSize         int
	CompSize         int
	CompRatio        float64
	EncTime          time.Duration
	DecTime          time.Duration
	Correct          bool
	ShannonEntropy   float64
	MaxEntropy       float64
	RepetitionFactor float64
}

type Results struct {
	results []Result
}

func (r *Results) SaveResultsCSV(name string) {
	fileName := filepath.Join("results", fmt.Sprintf("results_%s.csv", name))

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	w.Write([]string{"File", "Type", "OrigSize", "CompSize", "Ratio%", "EncTime(ms)", "DecTime(ms)", "Correct", "ShannonEntropy", "MaxEntropy", "RepetitionFactor"})

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
			fmt.Sprintf("%.4f", result.ShannonEntropy),
			fmt.Sprintf("%.4f", result.MaxEntropy),
			fmt.Sprintf("%.4f", result.RepetitionFactor),
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
	fmt.Printf("🎲 Shannon Entropy: %.4f\n", r.ShannonEntropy)
	fmt.Printf("🎲 Max Entropy: %.4f\n", r.MaxEntropy)
	fmt.Printf("🔁 Repetition Factor: %.4f\n", r.RepetitionFactor)
}

func shannonEntropy(data []byte) (float64, float64) {
	freq := make(map[byte]int)
	for _, b := range data {
		freq[b]++
	}

	total := float64(len(data))
	var entropy float64
	for _, count := range freq {
		p := float64(count) / total
		entropy -= p * math.Log2(p)
	}

	numSymbols := float64(len(freq))
	maxEntropy := math.Log2(numSymbols)

	return entropy, maxEntropy
}

func repetitionFactor(data []byte, k int) float64 {
	if len(data) < k {
		return 0
	}
	substrCount := make(map[string]int)
	totalSubstrings := 0

	for i := 0; i <= len(data)-k; i++ {
		sub := string(data[i : i+k])
		substrCount[sub]++
		totalSubstrings++
	}

	distinctSubstrings := len(substrCount)
	redundancy := 1 - float64(distinctSubstrings)/float64(totalSubstrings)
	return redundancy
}
