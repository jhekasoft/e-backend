#!/usr/bin/env make
# e-backend

BUILD_TIME=$(shell date -Iseconds)
VERSION=$(shell cat ./VERSION)
LDFLAGS=-s -w -X 'e-backend/internal.BuildTime=$(BUILD_TIME)' -X 'e-backend/internal.Version=$(VERSION)'
TAGS=all

all: clean build doc data

build:
	$(info ************ BUILDING EXECUTABLE FILE ************)
	go build -ldflags "$(LDFLAGS)" -tags="$(TAGS)" -o ./build/e-backend

doc:
	$(info ************ BUILDING DOC ************)
	# Update doc's version
	sed -i "s/\(version:\) .*/\1 $(VERSION)/" ./internal/modules/doc/data/public/restapi/openapi/openapi.yml

data:
	$(info ************ BUILDING DATA FILES ************)
	# Module CV
	mkdir -p ./build/internal/modules/cv/data
	cp -r ./internal/modules/cv/data/* ./build/internal/modules/cv/data

	# Module Doc
	mkdir -p ./build/internal/modules/doc/data
	cp -r ./internal/modules/doc/data/* ./build/internal/modules/doc/data

clean:
	$(info ************ CLEANING ************)
	rm -rf ./build

run:
	$(info ************ RUNNING ************)
	go run -ldflags "$(LDFLAGS)" -tags="$(TAGS)" main.go serve

test:
	$(info ************ RUNNING TESTS ************)
	go test ./...

.PHONY: all doc run
