#!/usr/bin/env bash

env GOOS=linux
env GOARCH=amd64

go build -o bin/linux/amd64/fiverm