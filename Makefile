# Code originates from github.com/woodpecker-ci/woodpecker

GO_PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

TARGETOS ?= $(shell go env GOOS)
TARGETARCH ?= $(shell go env GOARCH)

BIN_SUFFIX :=
ifeq ($(TARGETOS),windows)
	BIN_SUFFIX := .exe
endif

VERSION ?= next
VERSION_NUMBER ?= 0.0.0
CI_COMMIT_SHA ?= $(shell git rev-parse HEAD)

# it's a tagged release
ifneq ($(CI_COMMIT_TAG),)
	VERSION := $(CI_COMMIT_TAG:v%=%)
	VERSION_NUMBER := ${CI_COMMIT_TAG:v%=%}
else
	# append commit-sha to next version
	ifeq ($(VERSION),next)
		VERSION := $(shell echo "next-$(shell echo ${CI_COMMIT_SHA} | cut -c -10)")
	endif
	# append commit-sha to release branch version
	ifeq ($(shell echo ${CI_COMMIT_BRANCH} | cut -c -9),release/v)
		VERSION := $(shell echo "$(shell echo ${CI_COMMIT_BRANCH} | cut -c 10-)-$(shell echo ${CI_COMMIT_SHA} | cut -c -10)")
	endif
endif

LDFLAGS := -X go.codycody31.dev/vanity/version.Version=${VERSION}
STATIC_BUILD ?= true
ifeq ($(STATIC_BUILD),true)
	LDFLAGS := -s -w -extldflags "-static" $(LDFLAGS)
endif
CGO_ENABLED ?= 1 # only used to compile server

HAS_GO = $(shell hash go > /dev/null 2>&1 && echo "GO" || echo "NOGO" )
ifeq ($(HAS_GO),GO)
	XGO_VERSION ?= go-1.20.x
	CGO_CFLAGS ?= $(shell go env CGO_CFLAGS)
endif
CGO_CFLAGS ?=

# If the first argument is "in_docker"...
ifeq (in_docker,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "in_docker"
  MAKE_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # Ignore the next args
  $(eval $(MAKE_ARGS):;@:)

  in_docker:
	@[ "1" -eq "$(shell docker image ls woodpecker/make:local -a | wc -l)" ] && docker buildx build -f ./docker/Dockerfile.make -t woodpecker/make:local --load . || echo reuse existing docker image
	@echo run in docker:
	@docker run -it \
		--user $(shell id -u):$(shell id -g) \
		-e VERSION="$(VERSION)" \
		-e CI_COMMIT_SHA="$(CI_COMMIT_SHA)" \
		-e TARGETOS="$(TARGETOS)" \
		-e TARGETARCH="$(TARGETARCH)" \
		-e CGO_ENABLED="$(CGO_ENABLED)" \
		-e GOPATH=/tmp/go \
		-e HOME=/tmp/home \
		-v $(PWD):/build --rm woodpecker/make:local make $(MAKE_ARGS)
else

# Proceed with normal make

##@ General

.PHONY: all
all: help

.PHONY: version
version: ## Print the current version
	@echo ${VERSION}

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: vendor
vendor: ## Update the vendor directory
	go mod tidy
	go mod vendor

format: install-tools ## Format source code
	@gofumpt -extra -w .

.PHONY: clean
clean: ## Clean build artifacts
	go clean -i ./...
	rm -rf build
	@[ "1" != "$(shell docker image ls woodpecker/make:local -a | wc -l)" ] && docker image rm woodpecker/make:local || echo no docker image to clean

.PHONY: clean-all
clean-all: clean ## Clean all artifacts
	rm -rf dist

check-xgo: ## Check if xgo is installed
	@hash xgo > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) install src.techknowlogick.com/xgo@latest; \
	fi

install-tools: ## Install development tools
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest ; \
	fi ; \
	hash gofumpt > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go install mvdan.cc/gofumpt@latest; \
	fi

##@ Test

.PHONY: lint
lint: install-tools ## Lint code
	@echo "Running golangci-lint"
	golangci-lint run

test: ## Test code
	go test -race -cover -coverprofile agent-coverage.out -timeout 30s go.codycody31.dev/vanity/...

##@ Build

build: ## Build
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags '${LDFLAGS}' -o dist/vanity${BIN_SUFFIX} go.codycody31.dev/vanity

build-tarball: ## Build tar archive
	mkdir -p dist && tar chzvf dist/vanity-src.tar.gz \
	  --exclude="*.exe" \
	  --exclude="./.pnpm-store" \
	  --exclude="node_modules" \
	  --exclude="./dist" \
	  --exclude="./data" \
	  --exclude="./build" \
	  --exclude="./.git" \
	  .

release: ## Create binaries for release
	# compile
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '${LDFLAGS}' -tags 'grpcnotrace $(TAGS)' -o dist/linux_amd64       go.codycody31.dev/vanity
	GOOS=linux   GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '${LDFLAGS}' -tags 'grpcnotrace $(TAGS)' -o dist/linux_arm64       go.codycody31.dev/vanity
	GOOS=linux   GOARCH=arm   CGO_ENABLED=0 go build -ldflags '${LDFLAGS}' -tags 'grpcnotrace $(TAGS)' -o dist/linux_arm         go.codycody31.dev/vanity
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '${LDFLAGS}' -tags 'grpcnotrace $(TAGS)' -o dist/windows_amd64.exe go.codycody31.dev/vanity
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '${LDFLAGS}' -tags 'grpcnotrace $(TAGS)' -o dist/darwin_amd64      go.codycody31.dev/vanity
	GOOS=darwin  GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '${LDFLAGS}' -tags 'grpcnotrace $(TAGS)' -o dist/darwin_arm64      go.codycody31.dev/vanity
	# tar binary files
	tar -cvzf dist/vanity_linux_amd64.tar.gz   -C dist/linux_amd64   vanity
	tar -cvzf dist/vanity_linux_arm64.tar.gz   -C dist/linux_arm64   vanity
	tar -cvzf dist/vanity_linux_arm.tar.gz     -C dist/linux_arm     vanity
	tar -cvzf dist/vanity_windows_amd64.tar.gz -C dist/windows_amd64 vanity.exe
	tar -cvzf dist/vanity_darwin_amd64.tar.gz  -C dist/darwin_amd64  vanity
	tar -cvzf dist/vanity_darwin_arm64.tar.gz  -C dist/darwin_arm64  vanity

release-checksums: ## Create checksums for all release files
	# generate shas for tar files
	(cd dist/; sha256sum *.* > checksums.txt)
endif