#!/usr/bin/env bash

# Build for Windows
env GOOS=windows GOARCH=amd64 go build -o build/pong.exe .

# Build for Linux
env GOOS=linux GOARCH=amd64 go build -o build/pong .

# Build for Web
cp $(go env GOROOT)/misc/wasm/wasm_exec.js public/
env GOOS=js GOARCH=wasm go build -o public/pong.wasm
zip build/pong-web-go.zip public/*
