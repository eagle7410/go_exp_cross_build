#!/usr/bin/env bash
# Run as sudo
# build variants serch in https://dave.cheney.net/2012/09/08/an-introduction-to-cross-compilation-with-go
apt-get install gcc-multilib
apt-get install gcc-mingw-w64
# example cmd GOOS=windows GOARCH=amd64 go build main.go
