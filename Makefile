BUILD_VERSION ?= $(shell git describe --always --tags)

help:
	@echo "Targets are:\n"
	@echo "run"
	@echo " run the fidipides locally"
	@echo
	@echo "build"
	@echo " build the fidipides client"

all: help

run:
	@go run ./src/main.go run

build:
	@go build -ldflags \
	    "-X ./src/cmd/version.Version=$(BUILD_VERSION)" \
	    -o ./pkg/fidipides \
	    ./src
