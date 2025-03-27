# flow-proto

A shared protocol buffer (protobuf) definitions repository for Flow's microservices architecture.

## Overview

The `flow-proto` repository contains the protocol buffer definitions that form the contract between Flow core and its distributed components (Sources, Processors, and Consumers). This repository serves as the central source of truth for the RPC interfaces, enabling language-agnostic service development.

## Purpose

This repository:
- Defines the data models and service methods for the Flow ecosystem
- Enables version control of the RPC interface independent of implementation
- Provides the foundation for generating client/server code in multiple programming languages
- Ensures consistency and compatibility across the Flow ecosystem

## Architecture

Flow has been redesigned with a decoupled architecture similar to some IaC tools:

- **Flow Core**: Focuses on configuration parsing, state management, and orchestration
- **Services**: Independent microservices (Sources, Processors, Consumers) that implement the RPC interfaces defined in this repository
- **Service Discovery**: Flexible discovery mechanisms for finding and connecting to services

The proto definitions in this repository define the following interfaces:

- **Source Services**: For data ingestion components
- **Processor Services**: For data transformation and business logic
- **Consumer Services**: For outputting data to external systems

## Usage

### Generating Code

To generate code from these proto definitions:

```bash
# Install protoc compiler
# For Go:
protoc --go_out=. --go-grpc_out=. ./path/to/proto/files
# For other languages, use the appropriate plugin
```

### Implementing a Service

1. Generate code using the above instructions
2. Implement the required service interface
3. Register your service with the appropriate discovery mechanism

## Development

### Prerequisites

- Protocol Buffers compiler (protoc)
- Language-specific protoc plugins

### Adding or Modifying Definitions

1. Update the .proto files
2. Run tests to ensure compatibility
3. Update version numbers according to semantic versioning

## Integration

Flow core and its microservices import this shared API module to ensure consistency across the ecosystem. The RPC interface enables:

- **Capability Negotiation**: Services can describe their capabilities
- **Configuration**: Services can be configured at runtime
- **Data Processing**: Bidirectional streaming for efficient data handling
- **Health Checks**: Monitoring service status

## Contributing

Contributions are welcome! Please ensure that any changes to the proto definitions are backward compatible unless a major version change is planned.

## License

[License information here]
