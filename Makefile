#!/usr/bin/env make
# e-backend

BUILD_TIME=$(shell date -Iseconds)
VERSION=$(shell cat ./VERSION)
LDFLAGS=-s -w -X 'e-backend/internal.BuildTime=$(BUILD_TIME)' -X 'e-backend/internal.Version=$(VERSION)'
TAGS=all

build:
	$(info ************ BUILDING EXECUTABLE FILE ************)
	go build -ldflags "$(LDFLAGS)" -tags="$(TAGS)" -o ./build/e-backend

clean:
	$(info ************ CLEANING ************)
	rm -rf build/*
	rmdir build

test:
	$(info ************ RUNNING TESTS ************)
	go test ./...
