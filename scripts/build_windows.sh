#!/usr/bin/env bash

env GOOS=windows
env GOARCH=amd64

go build -o bin/windows/amd64/fiverm.exe