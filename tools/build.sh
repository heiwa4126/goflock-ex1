#!/bin/sh -ex
cd $(dirname $0)/..
BIN=$(basename $PWD)
MOD=$(grep module go.mod | cut -d' ' -f2)

go mod tidy
go fmt ./...

# test
go clean -testcache
go test ./...

# lint
go vet ./...
shadow ./...
golangci-lint run

# build
VER=$(git tag --sort=-v:refname | grep '^v' | head -1 | sed 's/^v//')
REV=$(git rev-parse --short HEAD)
go build -ldflags "-s -w -X $MOD/cli.Version=$VER -X $MOD/cli.Revision=$REV" -trimpath

# omake
go version -m "$BIN"
upx "$BIN"
