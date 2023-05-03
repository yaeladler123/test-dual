BUILD_DIR := $(CURDIR)/build
BIN_DIR := $(BUILD_DIR)/bin
LOGS_DIR := $(BUILD_DIR)/logs

IMAGE_NAME := "potato-service"
BINARY_NAME := "potato-service"

PROTOC = $(BIN_DIR)/protoc
PROTOC_VERSION = 21.12

GOLANGCI_LINT := $(BIN_DIR)/golangci-lint
GOLANGCI_LINT_VERSION := v1.51.1

OS := $(shell uname -s)

.DEFAULT_GOAL := help

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

init:
	@mkdir -p "$(BUILD_DIR)" "$(BIN_DIR)"

##@ Build

.PHONY: build
build: init ## Build and install the binary.
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./main.go

.PHONY: test
test: init build ## Execute unit tests.
	@echo "➡️Running tests..."
	go test ./...

.PHONY: lint
lint: init $(GOLANGCI_LINT) ## Lint the go code.
	@$(GOLANGCI_LINT) run .

.PHONY: run
run: init build ## Locally run the server.
	@echo "➡️Launching server..."
	go run ./main.go -configFile=./local_config.json

.PHONY: container
container: init build lint test ## Build the container image
	@echo "➡️Launching server..."
	docker build -t $(IMAGE_NAME):latest .

.PHONY: all
all: clean lint test build ## Run all the checks and tests.
	@echo "✅ Done!"

.PHONY: clean
clean: ## Delete the build directory.
	@rm -rf $(BUILD_DIR) $(LOGS_DIR)

$(GOLANGCI_LINT):
	@$(call install-golangci-lint-version, $(GOLANGCI_LINT), $(GOLANGCI_LINT_VERSION), $(BIN_DIR))

# go-install-tool will 'go install' any package $2 and install it to $1.
define go-install-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp 2>/dev/null ;\
echo "Downloading $(2)" ;\
GOBIN=$(BIN_DIR) go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

# install-golangci-lint-version will check the installed golangci-lint (located in $1) version. If version does not equal $2, then install in $3
define install-golangci-lint-version
@LINT_RESULT=$$($(1) --version 2>/dev/null);\
LINT_VERSION=$$(echo "$$LINT_RESULT" | cut -d ' ' -f 4); \
if [ "$$LINT_RESULT" != "0" ] || [ "$$LINT_VERSION" != "v$(2)" ]; then \
	echo "➡️Installing locally golangci-lint $(2)"; \
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(3) $(2); \
fi
endef