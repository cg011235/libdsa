#!/bin/bash

# Change directory to the script's location
cd "$(dirname "$0")"

# Variables for file names and compiler flags
GO_TEST_FILES="trie_test.go trie.go"
CPP_SOURCE_FILES="trie_test.cc trie.cc"
CPP_OUTPUT_FILE="trie_test_cc"
CPP_FLAGS="-pthread -lgtest -std=c++11"

# Function to run Go tests
run_go_tests() {
    echo "Running Go tests..."
    go test $GO_TEST_FILES -v
    if [ $? -ne 0 ]; then
        echo "Go tests failed, aborting."
        exit 1
    fi
    echo "Go tests passed."
}

# Function to compile and run C++ tests
compile_and_run_cpp_tests() {
    echo "Compiling C++ tests..."
    g++ $CPP_SOURCE_FILES $CPP_FLAGS -o $CPP_OUTPUT_FILE
    if [ $? -ne 0 ]; then
        echo "C++ compilation failed, aborting."
        exit 1
    fi

    echo "Running C++ tests..."
    ./$CPP_OUTPUT_FILE
    if [ $? -ne 0 ]; then
        echo "C++ tests failed."
        exit 1
    else
        echo "C++ tests passed."
    fi
}

# Main script execution
run_go_tests
compile_and_run_cpp_tests

# Clean up
rm $CPP_OUTPUT_FILE
