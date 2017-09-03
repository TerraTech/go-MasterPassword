PROJ=go-MasterPassword
export CLI=gompw
ORG_PATH=github.com/TerraTech
REPO_PATH=$(ORG_PATH)/$(PROJ)
CMD_PATH=$(REPO_PATH)/cmd
export PATH := $(PWD)/bin:$(PATH)

GOFILES := $(filter-out ./vendor/% ./@_VERSION_@.go,$(shell find ./ -type f -name '*.go' -print))

BUILDHOST ?= $(shell hostname -s)

$( shell mkdir -p bin )

user=$(shell id -u -n)
group=$(shell id -g -n)

LD_FLAGS="-w -X main.BUILDHOST=$(BUILDHOST)"

build: bin/$(CLI)

bin/$(CLI): $(GOFILES)
	@GOBIN=$(PWD)/bin go install -v -ldflags $(LD_FLAGS) $(CMD_PATH)/$(CLI)

.PHONY: install
install: build
	@scripts/install.sh

.PHONY: test
test:
	@go test -v -i $(shell go list ./... | grep -v '/vendor/')
	@go test -v $(shell go list ./... | grep -v '/vendor/')

.PHONY: vet
vet:
	@go vet $(shell go list ./... | grep -v '/vendor/')

.PHONY: fmt
fmt:
	@go fmt $(shell go list ./... | grep -v '/vendor/')

.PHONY: clean
clean:
	@scripts/clean.sh

.PHONY: testall
testall: test vet fmt

FORCE:
