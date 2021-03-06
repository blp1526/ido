export GO111MODULE=on
export GOBIN=${PWD}/bin

REVISION = $(shell git rev-parse --short HEAD)
LDFLAGS = -ldflags="-s -w -X 'github.com/blp1526/ido.revision=$(REVISION)'"

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -rf bin/
	@echo

.PHONY: mod
mod:
	go mod tidy
	go get github.com/editorconfig-checker/editorconfig-checker/cmd/editorconfig-checker
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	@echo

.PHONY: editorconfig
editorconfig: mod
	./bin/editorconfig-checker
	@echo

.PHONY: lint
lint: editorconfig
	./bin/golangci-lint run ./...
	@echo

.PHONY: test
test: lint
	go test ./... -v --cover -race -covermode=atomic -coverprofile=coverage.txt
	@echo

.PHONY: build
build: test
	go build $(LDFLAGS) -o bin/ido ./cmd/ido
	@echo
