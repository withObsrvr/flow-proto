# Flow Proto Repository

This repository contains the protobuf definitions for the Flow microservices architecture. These definitions serve as the contract between Flow Core and its distributed components (Sources, Processors, and Consumers).

## Contents

- `proto/processor`: Protocol definitions for processor services
- `proto/source`: Protocol definitions for source services
- `proto/consumer`: Protocol definitions for consumer services

## Usage

### Adding to a Go Project

```bash
# Add the dependency to your project
go get github.com/withObsrvr/flow-proto@latest

# Update your go.mod file
go mod tidy
```

### Importing in Go Code

Import the specific proto package you need:

```go
import (
    // Import processor definitions
    processorpb "github.com/withObsrvr/flow-proto/proto/processor"
    
    // Import source definitions
    sourcepb "github.com/withObsrvr/flow-proto/proto/source"
    
    // Import consumer definitions
    consumerpb "github.com/withObsrvr/flow-proto/proto/consumer"
)
```

### Example Usage

```go
// Create a processor client
client := processorpb.NewProcessorServiceClient(conn)

// Call GetCapabilities
capabilities, err := client.GetCapabilities(ctx, &processorpb.CapabilitiesRequest{})
if err != nil {
    log.Fatalf("Failed to get capabilities: %v", err)
}

// Process data with bidirectional streaming
stream, err := client.Process(ctx)
if err != nil {
    log.Fatalf("Error opening stream: %v", err)
}

// Send data messages
msg := &processorpb.DataMessage{
    Payload:  []byte("hello"),
    Metadata: map[string]string{"key": "value"},
}
if err := stream.Send(msg); err != nil {
    log.Fatalf("Failed to send message: %v", err)
}

// Receive processed messages
resp, err := stream.Recv()
if err != nil {
    log.Fatalf("Error receiving message: %v", err)
}
```

## Regenerating Proto Code

```bash
# Install protoc compiler and Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code from proto definitions
protoc --go_out=. --go-grpc_out=. proto/processor/processor.proto proto/source/source.proto proto/consumer/consumer.proto
```

## Development

When making changes to proto definitions:

1. Update the proto files
2. Regenerate code
3. Increment version following semver
4. Push changes to the repository

## Integration with Flow Core

For local development of Flow Core, you can use a replace directive in go.mod:

```
replace github.com/withObsrvr/flow-proto => ../flow-proto
```
