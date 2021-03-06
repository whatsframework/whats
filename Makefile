DIST := dist
EXECUTABLE := whats
GOFMT ?= gofmt "-s"
GO ?= go

TARGETS ?= linux darwin windows
ARCHS ?= amd64 386
PACKAGES ?= $(shell $(GO) list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f)
TAGS ?=

VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)

LDFLAGS ?= -X 'main.Version=$(VERSION)'

EXTLDFLAGS = -extldflags "-static" $(null)

all: build

#c: vet misspell-check sec
c: vet lint misspell-check sec
 
fmt:
	$(GOFMT) -w $(SOURCES)

vet:
	$(GO) vet $(PACKAGES)

lint:
	@hash revive > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mgechev/revive; \
	fi
	revive -config .revive.toml ./... || exit 1

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -error $(SOURCES)

.PHONY: misspell
misspell:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -w $(SOURCES)

sec:
	@hash gosec > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/securego/gosec/v2/cmd/gosec; \
	fi
	gosec -exclude=G401,G404,G501 ./...

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(SOURCES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

verify: vet misspell-check fmt-check

test: fmt-check
	@$(GO) test -v -cover -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

build:
	$(GO) build -v -tags "$(TAGS)" -ldflags "$(EXTLDFLAGS) -s -w $(LDFLAGS)"  -o $(EXECUTABLE) .
