#!/bin/bash

# --- Configuration ---
MAIN_PACKAGE="github.com/debarbarinantoine/go-enum-generate"
VERSION_FILE="./version"
BINARY_NAME="go-enum-generate"

# Define target platforms
PLATFORMS=(
  "linux/amd64"
  "linux/arm64"
  "darwin/amd64"
  "darwin/arm64"
  "windows/amd64"
)

# Output directory
BUILD_DIR="build"
mkdir -p "$BUILD_DIR"

# --- Read version from file ---
if [ ! -f "$VERSION_FILE" ]; then
    echo "Error: Version file '$VERSION_FILE' not found!"
    exit 1
fi
read -r APP_VERSION < "$VERSION_FILE"

# --- Build flags for release ---
# Define LDFLAGS without extra internal quotes, it will be quoted when passed to go build
LDFLAGS="-s -w"

# --- Build ---
echo "Building ${BINARY_NAME} v${APP_VERSION}..."

for PLATFORM in "${PLATFORMS[@]}"; do
  GOOS=$(echo "$PLATFORM" | cut -d'/' -f1)
  GOARCH=$(echo "$PLATFORM" | cut -d'/' -f2)

  OUTPUT_NAME="go-enum-generate-${GOOS}-${GOARCH}"
  if [ "$GOOS" = "windows" ]; then
    OUTPUT_NAME+=".exe"
  fi

  echo "Building for ${GOOS}/${GOARCH}..."
  env GOOS="$GOOS" GOARCH="$GOARCH" go build \
    -o "${BUILD_DIR}/${OUTPUT_NAME}" \
    -ldflags "${LDFLAGS}" \
    -trimpath \
    ${MAIN_PACKAGE}

  if [ $? -ne 0 ]; then
    echo "Error building for ${PLATFORM}. Aborting."
    exit 1
  fi
done

echo "Build successful: ./${BINARY_NAME} v${APP_VERSION}"