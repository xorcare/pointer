# Based on https://git.io/fjkGc

# The full path to the main package is use in the
# imports tool to format imports correctly.
NAMESPACE = github.com/xorcare/pointer

# The name of the file recommended in the standard
# documentation go test -cover and used codecov.io
# to check code coverage.
COVER_FILE ?= coverage.out

# Main targets.
.DEFAULT_GOAL := help

.PHONY: build
build: ## Build the project binary
	@go build ./...

.PHONY: ci
ci: check ## Target for integration with ci pipeline

.PHONY: check
check: static test build ## Check project with static checks and unit tests

.PHONY: help
help: ## Print this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: imports
imports: tools ## Check and fix import section by import rules
	@test -z $$(for d in $$(go list -f {{.Dir}} ./...); do goimports -e -l -local $(NAMESPACE) -w $$d/*.go; done)

.PHONY: lint
lint: tools ## Check the project with lint
	@golint -set_exit_status ./...

.PHONY: static
static: imports vet lint ## Run static checks (lint, imports, vet, ...) all over the project

.PHONY: test
test: ## Run unit tests
	@go test ./... -count=1 -race
	@go test ./... -count=1 -coverprofile=$(COVER_FILE) -covermode=atomic $d
	@go tool cover -func=$(COVER_FILE) | grep ^total

.PHONY: tools
tools: ## Install all needed tools, e.g. for static checks
	@go install golang.org/x/lint/golint@v0.0.0-20210508222113-6edffad5e616
	@go install golang.org/x/tools/cmd/goimports@v0.0.0-20200130002326-2f3ba24bd6e7

.PHONY: vet
vet: ## Check the project with vet
	@go vet ./...
