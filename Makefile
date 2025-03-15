.PHONY: build run test clean package

VERSION := $(shell cat ./Version)
LDFLAGS := -ldflags "-X main.Version=$(VERSION)"

build:
	@mkdir -p ./dist
	GOPROXY=https://goproxy.cn go install github.com/goreleaser/goreleaser@latest
	go build -o ./dist/ding ./cmd/main.go

build-all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64

linux-amd64:
	@mkdir -p ./dist/linux-amd64
	GOOS=linux GOARCH=amd64 go build -o ./dist/linux-amd64/ding ./cmd/main.go

linux-arm64:
	@mkdir -p ./dist/linux-arm64
	GOOS=linux GOARCH=arm64 go build -o ./dist/linux-arm64/ding ./cmd/main.go

darwin-amd64:
	@mkdir -p ./dist/darwin-amd64
	GOOS=darwin GOARCH=amd64 go build -o ./dist/darwin-amd64/ding ./cmd/main.go

darwin-arm64:
	@mkdir -p ./dist/darwin-arm64
	GOOS=darwin GOARCH=arm64 go build -o ./dist/darwin-arm64/ding ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v ./...

package: build
	GOPROXY=https://goproxy.cn go run github.com/goreleaser/goreleaser@latest release --clean --snapshot

clean:
	rm -rf ./dist