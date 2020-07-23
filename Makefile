BINARY = dunder-dbcall

GO_BIN = $(GOPATH)/bin
GO_PKG = $(GOPATH)/pkg
GO_SRC = $(GOPATH)/src

GO = go
GOLINT = $(GO_BIN)/revive

PROJECT_BASE = $(GO_SRC)/$(BINARY)

VENDOR_FOLDER = $(PROJECT_BASE)/vendor

.PHONY: lint
lint: check-gosetup ## Run linter on a project
ifeq (, $(shell which ${GOLINT}))
	go get -u github.com/mgechev/revive
endif
	@for source in $(shell find ${PROJECT_BASE} -type f -name '*.go' -not -path '*/vendor/*'); do \
		${GOLINT} -config config.toml -formatter stylish $$source; \
	done

.PHONY: fmt
fmt: ## Run go fmt on a project
	go fmt ./...

.PHONY: unit-test
unit-test: ## Run unit  tests
	@go test -covermode=count -coverprofile=coverage.out ./... -v -count=1

.PHONY: check-gosetup
check-gosetup:
ifndef GOPATH
	$(error GOPATH is undefined)
endif

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

