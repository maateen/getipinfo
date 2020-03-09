# Go parameters
GOCMD=$(shell which go)
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=getipinfo
GOBIN=bin
GOOS=darwin linux
GOARCH=amd64 386
VERSION=v0.1

all: prepare build
prepare:
		$(GOCMD) mod tidy
		$(GOCMD) mod verify
build:
		mkdir -p $(GOBIN); \
		for goos in $(GOOS); do \
			for goarch in $(GOARCH); do \
				GOOS=$$goos GOARCH=$$goarch $(GOBUILD) cmd/getipinfo/getipinfo.go; \
				upx --brute getipinfo; \
				mv getipinfo $(GOBIN)/$(BINARY_NAME)-$$goos-$$goarch-$(VERSION); \
			done \
		done
clean: 
		$(GOCLEAN) cmd/getipinfo/main.go
		rm -rf $(GOBIN);