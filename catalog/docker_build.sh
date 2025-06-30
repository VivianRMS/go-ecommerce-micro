#!/bin/zsh

set -e

echo "Building Go binary using Bazel..."
bazel build --platforms=//:linux_amd64 //catalog/cmd/catalog:catalog

# Get bazel-bin path (absolute)
BIN_ROOT=$(bazel info bazel-bin)

# Find the binary under bazel-bin
FOUND=$(find "$BIN_ROOT/catalog/cmd/catalog" -type f -name 'catalog' | head -n 1)

if [ ! -f "$FOUND" ]; then
  echo "Error: Bazel build succeeded but binary not found."
  exit 1
fi

echo "Found binary at: $FOUND"

# Copy to location Dockerfile expects
cp "$FOUND" app-bazel

echo "Building Docker image..."
docker build -f Dockerfile -t catalog:latest .

echo "Done. Docker image 'catalog:latest' is ready."
