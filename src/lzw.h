#ifndef LZW_H
#define LZW_H
#include <vector>
#include <cstdint>
#include <string>

std::vector<int> lzw_compress(const std::vector<uint8_t>& input);
std::vector<uint8_t> lzw_decompress(const std::vector<int>& compressed);

#endif
