#include "lzw.h"
#include <stdexcept>
#include <unordered_map>

std::vector<int> lzw_compress(const std::vector<uint8_t>& input) {
    std::unordered_map<std::string, int> dictionary;
    for (int i = 0; i < 256; ++i) {
        dictionary[std::string(1, static_cast<char>(i))] = i;
    }
    std::string w;
    std::vector<int> result;
    int dictSize = 256;
    for (uint8_t c : input) {
        std::string wc = w + static_cast<char>(c);
        if (dictionary.count(wc)) {
            w = wc;
        } else {
            if (!w.empty())
                result.push_back(dictionary[w]);
            dictionary[wc] = dictSize++;
            w = std::string(1, static_cast<char>(c));
        }
    }
    if (!w.empty())
        result.push_back(dictionary[w]);
    return result;
}

std::vector<uint8_t> lzw_decompress(const std::vector<int>& compressed) {
    std::unordered_map<int, std::string> dictionary;
    for (int i = 0; i < 256; ++i) {
        dictionary[i] = std::string(1, static_cast<char>(i));
    }
    std::string w(1, static_cast<char>(compressed[0]));
    std::vector<uint8_t> result(w.begin(), w.end());
    int dictSize = 256;
    std::string entry;
    for (size_t i = 1; i < compressed.size(); ++i) {
        int k = compressed[i];
        if (dictionary.count(k)) {
            entry = dictionary[k];
        } else if (k == dictSize) {
            entry = w + w[0];
        } else {
            throw std::runtime_error("Bad compressed k");
        }
        result.insert(result.end(), entry.begin(), entry.end());
        dictionary[dictSize++] = w + entry[0];
        w = entry;
    }
    return result;
}
