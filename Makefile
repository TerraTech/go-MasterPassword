PROJ := go-MasterPassword
VERSION := $(shell git describe --tags --dirty)
export CLI := gompw
ORG_PATH := github.com/TerraTech
REPO_PATH := $(ORG_PATH)/$(PROJ)
CMD_PATH := $(REPO_PATH)/cmd
export PATH := $(PWD)/bin:$(PATH)

FQGOLIBS_PATH := $(GOPATH)/src/futurequest.net/FQgolibs-Public/
VENDOR_DST := futurequest.net/FQgolibs
VENDOR_SUBPKGS_FQ := FQdebug FQfile FQtesting FQversion

GOFILES := $(filter-out ./vendor/% ./@_VERSION_@.go,$(shell find ./ -type f -name '*.go' -print))

BUILDHOST ?= $(shell hostname -s)

$( shell mkdir -p bin )

user=$(shell id -u -n)
group=$(shell id -g -n)

LD_FLAGS="-w -X main.VERSION=$(VERSION) -X main.BUILDHOST=$(BUILDHOST)"

build: bin/$(CLI)

bin/$(CLI): $(GOFILES)
	@GOBIN=$(PWD)/bin go install -v -ldflags $(LD_FLAGS) $(CMD_PATH)/$(CLI)

.PHONY: help
help:
	@cat files/make_help.txt

.PHONY: install
install: build
	@scripts/install.sh

.PHONY: test
test:
	@go test -v $(shell go list ./... | grep -v '/vendor/')

.PHONY: vendor
vendor:
	@scripts/update_vendor.sh $(FQGOLIBS_PATH) $(VENDOR_DST) $(VENDOR_SUBPKGS_FQ)

.PHONY: vendorDry
vendorDry:
	@scripts/update_vendor-dryrun.sh $(FQGOLIBS_PATH) $(VENDOR_DST) $(VENDOR_SUBPKGS_FQ)

.PHONY: glide
glide:
	@scripts/update_glide.sh
	@make vendor

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
