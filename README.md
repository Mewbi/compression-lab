# compression-lab
Compression lab for DSA class at UFABC

## Compression Test Program

This repository now includes a small C++ program that tests LZW and Huffman
compression on sample PDF and PNG files. To compile and run the program:

```bash
# build
g++ -std=c++17 src/*.cpp -o compression_test

# execute on provided samples
./compression_test data/sample.pdf data/sample.png
```

The program reports compression time, decompression time, memory usage and
compression ratio for each algorithm.
