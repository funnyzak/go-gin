#!/bin/bash

build_dir="release"
build_name=$(basename "$(pwd)")

build() {
  local OS=$1
  local ARCH=$2
  local DIR=$build_dir/$build_name-$OS-$ARCH
  local OUT=$DIR/$build_name
  if [ $OS = "windows" ]; then OUT+=".exe"; fi

  echo "Building for $OS $ARCH"
  mkdir -p "$DIR"
  GOOS=$OS GOARCH=$ARCH go build -o "$OUT" ./cmd/main.go
}

compress() {
  local FILE=$1
  zip -r "$FILE.zip" "$FILE"
  rm -rf "$FILE"
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

ls -R "$build_dir"

cd "$build_dir"
for file in *; do
  compress "$file"
done
