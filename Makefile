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
	# Update REST doc's version
	sed -i "s/\(version:\) .*/\1 $(VERSION)/" ./modules/doc/data/public/restapi/openapi/openapi.yml

	# Update MQTT doc's version
	sed -i "s/\(version:\) .*/\1 $(VERSION)/" ./modules/doc/data/public/mqttapi/asyncapi/asyncapi.yml

data:
	$(info ************ BUILDING DATA FILES ************)
	# Config example
	cp ./.e-backend.example ./build/.e-backend.example

	# Module CV
	mkdir -p ./build/modules/cv/data
	cp -r ./modules/cv/data/* ./build/modules/cv/data
	mkdir -p ./build/modules/cv/templates
	cp -r ./modules/cv/templates/* ./build/modules/cv/templates

	# Module Doc
	mkdir -p ./build/modules/doc/data
	cp -r ./modules/doc/data/* ./build/modules/doc/data

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
