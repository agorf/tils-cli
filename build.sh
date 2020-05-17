#!/usr/bin/env bash

version=$(grep 'const Version =' version/version.go | tr -d ' ' | cut -d = -f 2 | tr -d '"')

rm -fv tils-$version-*

for goos in linux darwin windows; do
  for goarch in amd64 386; do
    filename="tils-$version-$goos-$goarch"
    echo "Building $filename ..."
    env GOOS=$goos GOARCH=$goarch go build -o "$filename"
    ls -lh "$filename"
  done
done
