PROTO_DIR := ./proto
GO_SRC_DICR := ./go
OUT_DIR   := .$(GO_SRC_DIR)/gen   # you’re generating next to the .proto source
GO_PKGS   := $(shell go list ./... | grep -v /vendor/)

# List every .proto file so ‘make’ knows what to re‑generate
PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto')

.PHONY: proto tidy test clean

## Regenerate *.pb.go and *_grpc.pb.go
proto: $(PROTO_FILES)
	@echo "→ Generating Go stubs…"
	@protoc \
	  -I=$(PROTO_DIR) \
	  -I=. \
	  --go_out=$(OUT_DIR) \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=$(OUT_DIR) \
	  --go-grpc_opt=paths=source_relative \
	  $(PROTO_FILES)
	@echo "✓ Done"

## Go module tidy check
tidy:
	go mod tidy && go vet ./...

## Run the tiny demo program (cmd/test)
test: proto tidy
	go run ./cmd/test

## Delete generated code
clean:
	rm -f $(shell find $(PROTO_DIR) -name '*.pb.go')

