#ifndef HUFFMAN_H
#define HUFFMAN_H
#include <vector>
#include <cstdint>

std::vector<uint8_t> huffman_compress(const std::vector<uint8_t>& data);
std::vector<uint8_t> huffman_decompress(const std::vector<uint8_t>& data);

#endif
