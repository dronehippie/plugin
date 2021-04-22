include .bingo/Variables.mk

NAME := plugin
IMPORT := github.com/dronehippie/$(NAME)
SHELL := bash

ifeq ($(UNAME), Darwin)
	GOBUILD ?= CGO_ENABLED=0 go build -i
else
	GOBUILD ?= CGO_ENABLED=0 go build
endif

PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f)

TAGS ?= netgo

LDFLAGS += -s -w -extldflags "-static"

.PHONY: all
all: build

.PHONY: clean
clean:
	go clean -i ./...

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: staticcheck
staticcheck: $(STATICCHECK)
	$(STATICCHECK) -tags '$(TAGS)' $(PACKAGES)

.PHONY: lint
lint: $(GOLINT)
	for PKG in $(PACKAGES); do $(GOLINT) -set_exit_status $$PKG || exit 1; done;

.PHONY: test
test:
	go test -coverprofile coverage.out $(PACKAGES)

.PHONY: build
build:
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o /dev/null ./...
