#!/bin/bash

# Script to run all Go learning examples

echo "================================"
echo "Go Learning Examples Runner"
echo "================================"
echo ""

for dir in 01_hello_world 02_variables 03_functions 04_structs 05_interfaces 06_concurrency 07_error_handling; do
    echo "=========================================="
    echo "Running: $dir"
    echo "=========================================="
    (cd "$dir" && go run main.go)
    echo ""
done

echo "================================"
echo "All examples completed!"
echo "================================"
