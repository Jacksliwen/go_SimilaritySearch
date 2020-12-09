cmake_minimum_required(VERSION 2.8)
project(faissengine)
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -DEXPORT")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O3")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -g -Wl,-rpath=./lib:../lib")
include_directories(
        ${PROJECT_SOURCE_DIR}/include
        /usr/local/cuda/include
)


link_directories(
        ${PROJECT_SOURCE_DIR}/lib/faiss
        /usr/local/cuda/lib64
        /usr/lib64
        /usr/local/lib

)
set (CMAKE_VERBOSE_MAKEFILE on)
SET(LIBRARY_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/lib) 
add_library(faissengine SHARED ${PROJECT_SOURCE_DIR}/src/faissengine.cpp)

target_link_libraries(
        faissengine
        faiss
        pthread
        cuda
        #boost_regex 
        #boost_system
)
