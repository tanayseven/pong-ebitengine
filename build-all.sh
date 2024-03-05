#!/usr/bin/env bash
env GOOS=windows GOARCH=amd64 go build -o build/pong.exe .
env GOOS=linux GOARCH=amd64 go build -o build/pong .
