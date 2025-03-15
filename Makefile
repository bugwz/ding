.PHONY: build run test clean package

VERSION := $(shell cat ./Version)
COMMIT := $(shell git rev-parse --short HEAD)
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.Commit=$(COMMIT)"
GOPROXY := GOPROXY=https://goproxy.cn

DIST_DIR := ./dist

build:
	@$(GOPROXY) go install github.com/goreleaser/goreleaser@latest
	@mkdir -p $(DIST_DIR)
	go build $(LDFLAGS) -o $(DIST_DIR)/ding ./cmd/main.go

define build_cross
	@mkdir -p $(DIST_DIR)/$(1)-$(2)
	GOOS=$(1) GOARCH=$(2) $(GOPROXY) go build $(LDFLAGS) -o $(DIST_DIR)/$(1)-$(2)/ding ./cmd/main.go
endef

build-all: linux-amd64 linux-arm64 darwin-amd64 darwin-arm64

linux-amd64:
	$(call build_cross,linux,amd64)

linux-arm64:
	$(call build_cross,linux,arm64)

darwin-amd64:
	$(call build_cross,darwin,amd64)

darwin-arm64:
	$(call build_cross,darwin,arm64)

run:
	go run ./cmd/main.go

test:
	go test -v -race ./...

package: build
	 VERSION=$(VERSION) COMMIT=$(COMMIT) $(GOPROXY) go run github.com/goreleaser/goreleaser@latest release --clean --snapshot

clean:
	rm -rf $(DIST_DIR)