#!/bin/bash

echo "Compiling 01knapsack go lang solution"
go test -v
go clean

echo
echo
echo "Compiling 01knapsack C++ solution"
g++ -Wall -g -std=c++11 01knapsack.c++ 01knapsack_test.c++  -lgtest -lgtest_main -lpthread -o 01knapsack_test.out
ret=$?
[[ $ret -eq 0 ]] && ./01knapsack_test.out && rm -rf 01knapsack_test.out 01knapsack_test.out.dSYM 
