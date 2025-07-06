#include "huffman.h"
#include <queue>
#include <array>
#include <cstdint>
#include <stdexcept>

struct Node {
    int freq;
    unsigned char ch;
    Node* left;
    Node* right;
    Node(int f, unsigned char c) : freq(f), ch(c), left(nullptr), right(nullptr) {}
    Node(Node* l, Node* r) : freq(l->freq + r->freq), ch(0), left(l), right(r) {}
};

struct NodeCmp {
    bool operator()(Node* a, Node* b) const { return a->freq > b->freq; }
};

static void build_codes(Node* node, const std::string& prefix, std::array<std::string,256>& codes) {
    if (!node) return;
    if (!node->left && !node->right) {
        codes[node->ch] = prefix.empty() ? "0" : prefix;
    } else {
        build_codes(node->left, prefix + "0", codes);
        build_codes(node->right, prefix + "1", codes);
    }
}

static void free_tree(Node* node) {
    if (!node) return;
    free_tree(node->left);
    free_tree(node->right);
    delete node;
}

std::vector<uint8_t> huffman_compress(const std::vector<uint8_t>& data) {
    if (data.empty()) return {};
    std::array<int,256> freq{};
    for (uint8_t b : data) freq[b]++;
    std::priority_queue<Node*, std::vector<Node*>, NodeCmp> pq;
    for (int i=0;i<256;i++) if (freq[i]) pq.push(new Node(freq[i], static_cast<unsigned char>(i)));
    if (pq.empty()) pq.push(new Node(1,'\0'));
    while (pq.size()>1) {
        Node* a=pq.top(); pq.pop();
        Node* b=pq.top(); pq.pop();
        pq.push(new Node(a,b));
    }
    Node* root=pq.top();
    std::array<std::string,256> codes;
    build_codes(root, "", codes);
    std::vector<uint8_t> out;
    // header: frequency table
    out.resize(256*4);
    for (int i=0;i<256;i++) {
        uint32_t f=freq[i];
        for(int j=0;j<4;j++) out[i*4+3-j]=(f>>(j*8))&0xFF;
    }
    // generate bitstring
    std::string bits;
    bits.reserve(data.size()*2);
    for (uint8_t b: data) bits+=codes[b];
    uint64_t bitlen=bits.size();
    for(int i=7;i>=0;i--) out.push_back((bitlen>>(i*8))&0xFF);
    uint8_t byte=0; int count=0;
    for(char bit: bits){
        byte=(byte<<1)|(bit=='1');
        count++;
        if(count==8){ out.push_back(byte); byte=0; count=0; }
    }
    if(count){ byte <<= (8-count); out.push_back(byte); }
    free_tree(root);
    return out;
}

static Node* build_tree_from_freq(const std::array<int,256>& freq){
    std::priority_queue<Node*, std::vector<Node*>, NodeCmp> pq;
    for(int i=0;i<256;i++) if(freq[i]) pq.push(new Node(freq[i], static_cast<unsigned char>(i)));
    if(pq.empty()) pq.push(new Node(1,'\0'));
    while(pq.size()>1){ Node* a=pq.top(); pq.pop(); Node* b=pq.top(); pq.pop(); pq.push(new Node(a,b)); }
    return pq.top();
}

std::vector<uint8_t> huffman_decompress(const std::vector<uint8_t>& data){
    if(data.size()<256*4+8) throw std::runtime_error("Invalid data");
    std::array<int,256> freq{};
    for(int i=0;i<256;i++){
        uint32_t f=0;
        for(int j=0;j<4;j++) f=(f<<8)|data[i*4+j];
        freq[i]=f;
    }
    size_t idx=256*4;
    uint64_t bitlen=0;
    for(int i=0;i<8;i++) bitlen=(bitlen<<8)|data[idx++];
    Node* root=build_tree_from_freq(freq);
    std::vector<uint8_t> out;
    Node* node=root;
    uint64_t bits_read=0;
    while(bits_read<bitlen){
        uint8_t byte=data[idx + bits_read/8];
        int bit= (byte >> (7-(bits_read%8))) & 1;
        if(bit==0) node=node->left; else node=node->right;
        bits_read++;
        if(!node->left && !node->right){ out.push_back(node->ch); node=root; }
    }
    free_tree(root);
    return out;
}
