#!/bin/bash 
curPath=$(readlink -f "$(dirname "$0")")
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:${curPath}/searchengine/lib

mkdir searchengine/build
cd searchengine/build
cmake ..
make -j20
cd ../../
go build