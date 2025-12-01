PROTO_DIR := proto
GO_OUT := go/gen

# Find all flowctl and stellar proto files
FLOWCTL_PROTOS := $(shell find $(PROTO_DIR)/flowctl -name '*.proto' 2>/dev/null)
STELLAR_PROTOS := $(shell find $(PROTO_DIR)/stellar -name '*.proto' 2>/dev/null)
ALL_NEW_PROTOS := $(FLOWCTL_PROTOS) $(STELLAR_PROTOS)

# Legacy protos (for backward compatibility during transition)
LEGACY_PROTOS := $(shell find $(PROTO_DIR)/processor $(PROTO_DIR)/source $(PROTO_DIR)/consumer $(PROTO_DIR)/flow -name '*.proto' 2>/dev/null)

.PHONY: all proto proto-new proto-legacy clean verify help

# Default: generate new proto structure
all: proto-new

## Generate all new proto definitions (flowctl/v1 + stellar/v1)
proto-new:
	@echo "→ Generating flowctl/v1 and stellar/v1 Go stubs..."
	@mkdir -p $(GO_OUT)
	@protoc \
	  -I=$(PROTO_DIR) \
	  --go_out=$(GO_OUT) \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=$(GO_OUT) \
	  --go-grpc_opt=paths=source_relative \
	  $(ALL_NEW_PROTOS)
	@echo "✓ Generated $(words $(ALL_NEW_PROTOS)) proto files"
	@echo "✓ Output: $(GO_OUT)/flowctl/v1/ and $(GO_OUT)/stellar/v1/"

## Generate legacy protos (for backward compatibility)
proto-legacy:
	@echo "→ Generating legacy proto stubs..."
	@mkdir -p $(GO_OUT)
	@if [ -n "$(LEGACY_PROTOS)" ]; then \
		protoc \
		  -I=$(PROTO_DIR) \
		  --go_out=$(GO_OUT) \
		  --go_opt=paths=source_relative \
		  --go-grpc_out=$(GO_OUT) \
		  --go-grpc_opt=paths=source_relative \
		  $(LEGACY_PROTOS); \
		echo "✓ Generated $(words $(LEGACY_PROTOS)) legacy proto files"; \
	else \
		echo "No legacy protos found"; \
	fi

## Generate both new and legacy protos
proto: proto-new proto-legacy

## Verify proto files compile without errors
verify:
	@echo "→ Verifying proto files..."
	@protoc \
	  -I=$(PROTO_DIR) \
	  --descriptor_set_out=/dev/null \
	  $(ALL_NEW_PROTOS)
	@echo "✓ All proto files are valid"

## Clean generated code
clean:
	@echo "→ Cleaning generated code..."
	@rm -rf $(GO_OUT)
	@echo "✓ Cleaned $(GO_OUT)"

## Run go mod tidy after generation
tidy: proto-new
	@echo "→ Running go mod tidy..."
	@cd go && go mod tidy
	@echo "✓ Go modules tidied"

## List all proto files
list:
	@echo "Flowctl protos:"
	@echo "$(FLOWCTL_PROTOS)" | tr ' ' '\n'
	@echo "\nStellar protos:"
	@echo "$(STELLAR_PROTOS)" | tr ' ' '\n'

## Show help
help:
	@echo "Available targets:"
	@echo "  make proto-new    - Generate flowctl/v1 and stellar/v1 protos (default)"
	@echo "  make proto-legacy - Generate legacy protos for backward compatibility"
	@echo "  make proto        - Generate both new and legacy protos"
	@echo "  make verify       - Verify proto files compile"
	@echo "  make clean        - Remove generated code"
	@echo "  make tidy         - Generate protos and tidy go modules"
	@echo "  make list         - List all proto files"
	@echo "  make help         - Show this help message"
