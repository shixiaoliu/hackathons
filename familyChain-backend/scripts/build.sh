#!/bin/bash

# Exit on error
set -e

# Print commands
set -x

# Set variables
APP_NAME="eth-for-babies-backend"
BUILD_DIR="./build"
MAIN_PATH="./cmd/server/main.go"

# Create build directory if it doesn't exist
mkdir -p $BUILD_DIR

# Clean previous builds
echo "Cleaning previous builds..."
rm -f $BUILD_DIR/$APP_NAME

# Get dependencies
echo "Getting dependencies..."
go mod tidy

# Run tests
echo "Running tests..."
go test ./...

# Build the application
echo "Building application..."
go build -o $BUILD_DIR/$APP_NAME $MAIN_PATH

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "Build successful! Binary is at $BUILD_DIR/$APP_NAME"
else
    echo "Build failed!"
    exit 1
fi

# Make binary executable
chmod +x $BUILD_DIR/$APP_NAME

echo "Build process completed successfully!" 