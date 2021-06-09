.DEFAULT_GOAL := build

BIN_FILE=bin/http-advices

build:
	@go build -o "${BIN_FILE}"

run:
	./"${BIN_FILE}"