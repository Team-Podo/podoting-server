#!/usr/bin/env bash
set -xe

go get

GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"