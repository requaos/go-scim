.PHONY: all build generate assets binary test

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

NOW = $(shell date -u '+%Y%m%d%I%M%S')

DOCKER := docker
PKGER := pkger
GO := go
GO_ENV := $(shell $(GO) env GOOS GOARCH)
GOOS ?= $(word 1,$(GO_ENV))
GOARCH ?= $(word 2,$(GO_ENV))
GOFLAGS ?= $(GOFLAGS:)
ROOT_DIR := $(realpath .)

# GOOS/GOARCH of the build host, used to determine whether we're cross-compiling or not
BUILDER_GOOS_GOARCH="$(GOOS)_$(GOARCH)"

ifneq ($(GOOS), darwin)
	# EXTLDFLAGS = -extldflags "-lm -lstdc++ -static"
	EXTLDFLAGS =
else
	EXTLDFLAGS =
endif

GO_LINKER_FLAGS ?= --ldflags \
	'$(EXTLDFLAGS) -s -w '

all: build

build: generate binary

generate:
	@echo "$(OK_COLOR)==> Fetching dependencies...$(NO_COLOR)"
	$(GO) install github.com/markbates/pkger/cmd/pkger
	$(GO) mod download
	@echo "$(OK_COLOR)==> Generating files via go generate...$(NO_COLOR)"
	$(GO) generate $(GOFLAGS) ./...

assets:
	@echo "$(OK_COLOR)==> Generating Compile-Time Assets...$(NO_COLOR)"
	@echo "Assets to be compiled in:"
	@$(PKGER) list -include /schemas
	@$(PKGER) -include /schemas

binary:
	@echo "$(OK_COLOR)==> Building binary ($(GOOS)/$(GOARCH))...$(NO_COLOR)"
	$(GO) build -a $(GOFLAGS) $(GO_LINKER_FLAGS) -o bin/$(GOOS)_$(GOARCH)/scim

test: generate
	@echo "$(OK_COLOR)==> Running tests...$(NO_COLOR)"
	$(GO) test $(GOFLAGS) -race ./...
