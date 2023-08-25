NAME := walter
DOCKER_PREFIX = benmatselby
DOCKER_RELEASE ?= latest
DOCKER_PLATFORM ?= --platform=amd64
BUILD_DIR ?= build
GOOS ?=
ARCH ?=
OUT_PATH=$(BUILD_DIR)/$(NAME)-$(GOOS)-$(GOARCH)

.PHONY: explain
explain:
	### Welcome
	#
	#    _   _   _   _   _   _
	#   / \ / \ / \ / \ / \ /
	#  ( W | a | l | t | e | r )
	#   \_/ \_/ \_/ \_/ \_/ \_/
	#
	#
	### Installation
	#
	# $$ make all
	#
	### Targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

GITCOMMIT := $(shell git rev-parse --short HEAD)

.PHONY: clean
clean: ## Clean the local dependencies
	rm -fr vendor
	rm -fr $(BUILD_DIR) && mkdir -p $(BUILD_DIR)

.PHONY: install
install: ## Install the local dependencies
	go install github.com/golang/mock/mockgen@master
	go install github.com/securego/gosec/cmd/gosec@master
	go install golang.org/x/lint/golint@master
	go get ./...

.PHONY: lint
lint: ## Vet the code
	golangci-lint run

.PHONY: security
security: ## Inspect the code
	gosec ./...

.PHONY: build
build: ## Build the application
	go build .

.PHONY: static
static: ## Build the application
	CGO_ENABLED=0 go build -ldflags "-extldflags -static -X github.com/benmatselby/$(NAME)/version.GITCOMMIT=$(GITCOMMIT)" -o $(NAME) .

.PHONY: static-named
static-named: ## Build the application with named outputs
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 \
		go build \
		-ldflags "-extldflags -static -X github.com/benmatselby/$(NAME)/version.GITCOMMIT=$(GITCOMMIT)" \
		-o $(OUT_PATH) .

	md5sum $(OUT_PATH) > $(OUT_PATH).md5 || md5 $(OUT_PATH) > $(OUT_PATH).md5
	sha256sum $(OUT_PATH) > $(OUT_PATH).sha256 || shasum $(OUT_PATH) > $(OUT_PATH).sha256

.PHONY: test
test: ## Run the unit tests
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: test-cov
test-cov: test ## Run the unit tests with coverage
	go tool cover -html=coverage.out

.PHONY: mocks
mocks: ## Generate the mocks for the tests
	mockgen -source jira/jira.go

.PHONY: all ## Run everything
all: clean install build test

.PHONY: static-all ## Run everything
static-all: clean install static test


###
# Docker
###

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_PREFIX)/$(NAME):$(DOCKER_RELEASE) $(DOCKER_PLATFORM) .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_PREFIX)/$(NAME):$(DOCKER_RELEASE)

.PHONY: docker-scan
docker-scan: ## Scan the docker image
	docker scout recommendations $(DOCKER_PREFIX)/$(NAME)
	docker scout quickview $(DOCKER_PREFIX)/$(NAME)
