#!/usr/bin/env bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then

  echo "Building for Windows"
  env GOOS=windows GOARCH=amd64 go build -o build/pong.exe .

  echo "Building for Linux"
  env GOOS=linux GOARCH=amd64 go build -o build/pong .

  echo "Building for Web"
  cp $(go env $GOROOT)/misc/wasm/wasm_exec.js public/
  env GOOS=js GOARCH=wasm go build -o public/pong.wasm
  zip build/pong-web-go.zip public/*

elif [[ "$OSTYPE" == "darwin"* ]]; then

  echo "Building for Mac Intel"
  env GOOS=darwin GOARCH=amd64 go build -o build/pong-intel .

  echo "Building for Mac Apple Silicon"
  env GOOS=darwin GOARCH=arm64 go build -o build/pong-apple-silicon .

else
  echo "Unsupported OS"
fi
