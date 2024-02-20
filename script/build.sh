#!/bin/bash

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

build_dir="release"
build_name=$(basename "$(pwd)")

build() {
  local OS=$1
  local ARCH=$2
  local DIR=$build_dir/$build_name-$OS-$ARCH
  local OUT=$DIR/$build_name
  if [ $OS = "windows" ]; then OUT+=".exe"; fi

  echo "Building $OS $ARCH..."
  mkdir -p "$DIR"
  GOOS=$OS GOARCH=$ARCH go build -o "$OUT" ./cmd/main.go 2>/dev/null
  if [ $? -ne 0 ]; then
    echo -e "${red}Error: ${plain}Build $OS $ARCH failed."
    rm -rf "$DIR"
  else
    echo -e "${green}Success: ${plain}Build $OS $ARCH success."
  fi
  echo ""
}

compress() {
  local FILE=$1
  echo "Compressing $FILE..."
  zip -r "$FILE.zip" "$FILE" >/dev/null
  echo -e "${green}Success: ${plain}Compress $FILE success."
  rm -rf "$FILE"
  echo ""
}

rm -rf "$build_dir"
mkdir -p "$build_dir"

build linux arm
build linux 386
build linux amd64
build windows 386
build windows amd64
build darwin amd64
build darwin arm64
build freebsd 386
build freebsd amd64
build freebsd arm
build openbsd 386
build openbsd amd64
build openbsd arm
build netbsd 386
build netbsd amd64
build netbsd arm
build dragonfly amd64

ls -R "$build_dir"

cd "$build_dir"
echo ""
for file in *; do
  compress "$file"
done
echo -e "${green}Success: ${plain}Compress all success. Has been removed all uncompressed files. Total $(ls | wc -l) files."

echo -e "${green}Success: ${plain}All build success. See $build_dir."
