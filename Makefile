TARGET=apiserver  
.PHONY: build
build:
	go build -v -o ${TARGET} ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
