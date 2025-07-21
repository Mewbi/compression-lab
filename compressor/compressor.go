package compressor

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io"

	"github.com/dgryski/go-bitstream"
	huffman "github.com/dgryski/go-huff"
)

// CompressionType defines supported compression algorithms
type CompressionType int

const (
	Unknown CompressionType = iota
	LZW
	Huffman
)

func (c CompressionType) String() string {
	switch c {
	case LZW:
		return "LZW"
	case Huffman:
		return "Huffman"
	default:
		return "Unknown"
	}
}

type Compressor struct {
	Type           CompressionType
	CompressedData []byte
	tree           []byte
}

func NewCompressor(t CompressionType) *Compressor {
	return &Compressor{
		Type: t,
	}
}

func (c *Compressor) Compress(data []byte) error {
	defer func() error {
		if r := recover(); r != nil {
			msg := "❌ Recovered from panic during compression"
			fmt.Println(msg)
			return fmt.Errorf(msg, r)
		}
		return nil
	}()

	switch c.Type {
	case LZW:
		c.compressLZW(data)
	case Huffman:
		c.compressHuffman(data)
	default:
		fmt.Println("Invalid compression type")
	}
	return nil
}

func (c *Compressor) Decompress() ([]byte, error) {
	switch c.Type {
	case LZW:
		return c.decompressLZW(), nil
	case Huffman:
		return c.decompressHuffman()
	default:
		return nil, fmt.Errorf("invalid compression type")
	}
}

func (c *Compressor) compressHuffman(data []byte) {
	counts := make([]int, 256)

	for _, v := range data {
		counts[v]++
	}

	e := huffman.NewEncoder(counts)

	var b bytes.Buffer

	w := e.Writer(&b)

	for _, v := range data {
		w.WriteSymbol(uint32(v))
	}
	w.WriteSymbol(huffman.EOF)
	w.Close()

	c.CompressedData = b.Bytes()
	c.tree = e.CodebookBytes()
}

func (c *Compressor) decompressHuffman() ([]byte, error) {
	d, err := huffman.NewDecoder(c.tree)
	if err != nil {
		return nil, err
	}

	br := bitstream.NewReader(bytes.NewReader(c.CompressedData))

	var uncompressed []byte

	for {
		b, err := d.ReadSymbol(br)
		if err == io.EOF {
			return nil, err
		}
		if b == huffman.EOF {
			break
		}

		uncompressed = append(uncompressed, byte(b))
		if err != nil {
			return nil, err
		}
	}

	return uncompressed, nil
}

func (c *Compressor) compressLZW(data []byte) {
	var buf bytes.Buffer

	writer := lzw.NewWriter(&buf, lzw.LSB, 8)
	writer.Write(data)
	writer.Close()

	c.CompressedData = buf.Bytes()
}

func (c *Compressor) decompressLZW() []byte {
	buf := bytes.NewReader(c.CompressedData)

	reader := lzw.NewReader(buf, lzw.LSB, 8)
	defer reader.Close()
	result, _ := io.ReadAll(reader)

	return result
}
