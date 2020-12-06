cmake_minimum_required(VERSION 2.8)
project(func)
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -DEXPORT")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -O3")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -g -Wl,-rpath=./lib:../lib")
include_directories(
        ${PROJECT_SOURCE_DIR}/include
)

SET(LIBRARY_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/lib) 
add_library(faissengine SHARED ${PROJECT_SOURCE_DIR}/src/faissengine.cpp)
