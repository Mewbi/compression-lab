#include "lzw.h"
#include "huffman.h"
#include <fstream>
#include <iostream>
#include <chrono>
#include <sys/resource.h>
#include <unistd.h>
#include <cstring>

static std::vector<uint8_t> read_file(const std::string& path){
    std::ifstream f(path, std::ios::binary);
    return std::vector<uint8_t>((std::istreambuf_iterator<char>(f)), std::istreambuf_iterator<char>());
}

static size_t get_peak_rss(){
    struct rusage usage; getrusage(RUSAGE_SELF, &usage); return (size_t)usage.ru_maxrss * 1024L; }

static void test_lzw(const std::vector<uint8_t>& data){
    auto start_peak=get_peak_rss();
    auto t1=std::chrono::steady_clock::now();
    auto compressed=lzw_compress(data);
    auto t2=std::chrono::steady_clock::now();
    auto mid_peak=get_peak_rss();
    auto decompressed=lzw_decompress(compressed);
    auto t3=std::chrono::steady_clock::now();
    auto end_peak=get_peak_rss();
    double comp_time=std::chrono::duration<double>(t2-t1).count();
    double decomp_time=std::chrono::duration<double>(t3-t2).count();
    size_t comp_mem=mid_peak-start_peak;
    size_t decomp_mem=end_peak-mid_peak;
    double ratio=(double)(compressed.size()*sizeof(int))/data.size();
    std::cout<<"LZW\n";
    std::cout<<" Compression time: "<<comp_time<<"s\n";
    std::cout<<" Decompression time: "<<decomp_time<<"s\n";
    std::cout<<" Compression memory: "<<comp_mem/1024.0<<" KB\n";
    std::cout<<" Decompression memory: "<<decomp_mem/1024.0<<" KB\n";
    std::cout<<" Compression ratio: "<<ratio<<"\n";
    std::cout<<" Correct: "<<(data==decompressed?"yes":"no")<<"\n\n";
}

static void test_huffman(const std::vector<uint8_t>& data){
    auto start_peak=get_peak_rss();
    auto t1=std::chrono::steady_clock::now();
    auto compressed=huffman_compress(data);
    auto t2=std::chrono::steady_clock::now();
    auto mid_peak=get_peak_rss();
    auto decompressed=huffman_decompress(compressed);
    auto t3=std::chrono::steady_clock::now();
    auto end_peak=get_peak_rss();
    double comp_time=std::chrono::duration<double>(t2-t1).count();
    double decomp_time=std::chrono::duration<double>(t3-t2).count();
    size_t comp_mem=mid_peak-start_peak;
    size_t decomp_mem=end_peak-mid_peak;
    double ratio=(double)compressed.size()/data.size();
    std::cout<<"Huffman\n";
    std::cout<<" Compression time: "<<comp_time<<"s\n";
    std::cout<<" Decompression time: "<<decomp_time<<"s\n";
    std::cout<<" Compression memory: "<<comp_mem/1024.0<<" KB\n";
    std::cout<<" Decompression memory: "<<decomp_mem/1024.0<<" KB\n";
    std::cout<<" Compression ratio: "<<ratio<<"\n";
    std::cout<<" Correct: "<<(data==decompressed?"yes":"no")<<"\n\n";
}

int main(int argc, char** argv){
    if(argc<2){
        std::cerr<<"Usage: "<<argv[0]<<" <file1> [file2 ...]"<<std::endl;
        return 1;
    }
    for(int i=1;i<argc;i++){
        std::string path=argv[i];
        auto data=read_file(path);
        std::cout<<"Testing file: "<<path<<" ("<<data.size()<<" bytes)"<<std::endl;
        test_lzw(data);
        test_huffman(data);
    }
    return 0;
}
