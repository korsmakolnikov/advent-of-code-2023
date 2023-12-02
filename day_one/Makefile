BINARY=out
INPUT=input
.DEFAULT_GOAL := build
.PHONY: test watch

build:
	go build -o ${BINARY} cmd/main.go

run: build
	./${BINARY} < ${INPUT}

watch:
	gotestsum --watch

test:
	gotestsum
