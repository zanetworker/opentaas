BIN_DIR := $(GOPATH)/bin
VERSION ?= latest
OS ?= darwin

# support multi-package testing which go does not support natively. 
GOVERAGE := $(BIN_DIR)/goverage 
GOMETALINTER := $(BIN_DIR)/gometalinter
TAAS_BUILD_DIR := $(GOPATH)/github.com/zanetworker/taas

PKGS := $(shell go list ./... | grep -v /vendor)

BINARY := taas
PLATFORMS := darwin windows linux 
os = $(word 1, $@)

.PHONY: test 
test: $(GOVERAGE) lint 
	go test $(PKGS)
	goverage -race -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

.PHONY: testnolint
test: $(GOVERAGE)
	go test $(PKGS)
	goverage -race -coverprofile=coverage.out ./...

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null

$(GOVERAGE):
	go get -u github.com/haya14busa/goverage


.PHONY: lint 
lint: $(GOMETALINTER)
	gometalinter ./... --vendor --errors


release: 
	mkdir -p release 

.PHONY: $(PLATFORMS) 
$(PLATFORMS): 
	@- cd cmd && GOOS=$(os) GOARCH=amd64 go build -o ../release/$(BINARY)-$(VERSION)-$(os)-amd64


.PHONY: releases
releases: release darwin windows linux 


.PHONY: install 
install: releases
ifeq ($(OS),darwin)
	cp release/$(BINARY)-$(VERSION)-darwin-amd64 $(GOPATH)/bin/taas
else ifeq ($(OS),linux)
	cp release/$(BINARY)-$(VERSION)-linux-amd64 $(GOPATH)/bin/taas
else 
	cp release/$(BINARY)-$(VERSION)-windows-amd64 $(GOPATH)/bin/taas
endif

# For developers to do dry-run builds
.PHONY: dry
dry: 
	@- cd cmd && GOOS=$(OS) GOARCH=amd64 go build -o ../$(BINARY)

.PHONY: doc 
doc: dry 
	@-./$(BINARY)  > /dev/null  2>&1 || true  