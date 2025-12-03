# Flow Proto - Protocol Definitions for Flowctl

This repository contains the Protocol Buffer definitions for the flowctl data pipeline framework. These definitions serve as the contract between flowctl components (sources, processors, and consumers).

## Structure

```
proto/
├── flowctl/v1/                # Core flow protocol (general-purpose)
│   ├── common.proto               # Event, ComponentInfo, Health
│   ├── source.proto               # SourceService interface
│   ├── processor.proto            # ProcessorService interface
│   ├── consumer.proto             # ConsumerService interface
│   └── control_plane.proto        # ControlPlaneService (registration & discovery)
│
└── stellar/v1/                # Stellar blockchain extensions
    ├── raw_ledger.proto           # RawLedgerService (ledger streaming)
    ├── contract_events.proto      # ContractEventService (Soroban events)
    ├── contract_invocation.proto  # ContractInvocationEvent (Soroban invocations with execution context)
    ├── token_transfers.proto      # TokenTransferEvent (SEP-41 token transfers)
    ├── account_balances.proto     # AccountBalanceService
    └── contract_data.proto        # ContractDataService
```

## Overview

Flow Proto defines the protocol for the flowctl data pipeline framework, providing:

**Core Framework (`flowctl/v1`)**:
- Standard `Event` envelope for all data flowing through pipelines
- Component service interfaces (Source, Processor, Consumer)
- Control plane for component registration and discovery
- Health monitoring and metrics

**Stellar Integration (`stellar/v1`)**:
- Raw ledger streaming (XDR format)
- Contract events and invocations with rich execution context
- SEP-41 compliant token transfer events
- Account balance tracking
- Contract data access

**Recent Additions**:
- Contract invocation events with diagnostic events, cross-contract calls, and state changes
- Token transfer events supporting native XLM, issued assets, and Soroban tokens

## Quick Start

### Using in a Go Project

```bash
# Add the dependency
go get github.com/withObsrvr/flow-proto@latest

# Update go.mod
go mod tidy
```

### Import in Go Code

```go
import (
    // Core flow types (always import)
    flowctlv1 "github.com/withObsrvr/flow-proto/go/gen/flowctl/v1"

    // Stellar extensions (if building Stellar components)
    stellarv1 "github.com/withObsrvr/flow-proto/go/gen/stellar/v1"
)
```

## Usage Examples

### Building a Processor

```go
package main

import (
    "context"
    flowctlv1 "github.com/withObsrvr/flow-proto/go/gen/flowctl/v1"
    "google.golang.org/grpc"
)

type myProcessor struct {
    flowctlv1.UnimplementedProcessorServiceServer
}

func (p *myProcessor) Process(stream flowctlv1.ProcessorService_ProcessServer) error {
    for {
        event, err := stream.Recv()
        if err != nil {
            return err
        }

        // Process event
        processed := processEvent(event)

        // Send result
        if err := stream.Send(processed); err != nil {
            return err
        }
    }
}

func (p *myProcessor) GetInfo(ctx context.Context, _ *emptypb.Empty) (*flowctlv1.ComponentInfo, error) {
    return &flowctlv1.ComponentInfo{
        Id:                "my-processor",
        Name:              "My Processor",
        Type:              flowctlv1.ComponentType_COMPONENT_TYPE_PROCESSOR,
        InputEventTypes:   []string{"stellar.ledger.v1"},
        OutputEventTypes:  []string{"contract.event.v1"},
    }, nil
}
```

### Building a Stellar Source

```go
package main

import (
    flowctlv1 "github.com/withObsrvr/flow-proto/go/gen/flowctl/v1"
    stellarv1 "github.com/withObsrvr/flow-proto/go/gen/stellar/v1"
)

type stellarSource struct {
    stellarv1.UnimplementedRawLedgerServiceServer
}

func (s *stellarSource) StreamRawLedgers(req *stellarv1.StreamLedgersRequest, stream stellarv1.RawLedgerService_StreamRawLedgersServer) error {
    for ledger := req.StartLedger; ; ledger++ {
        rawLedger := fetchLedger(ledger)

        if err := stream.Send(&stellarv1.RawLedger{
            Sequence:             ledger,
            LedgerCloseMetaXdr:   rawLedger,
        }); err != nil {
            return err
        }
    }
}
```

### Using the Standard Event Envelope

All components use the same `Event` message:

```go
event := &flowctlv1.Event{
    Id:                "evt-12345",
    Type:              "stellar.ledger.v1",
    Payload:           ledgerData,
    Metadata:          map[string]string{
        "ledger_sequence": "1000",
        "network": "testnet",
    },
    SourceComponentId: "stellar-source",
    ContentType:       "stellar/xdr",
    StellarCursor:     &flowctlv1.StellarCursor{
        LedgerSequence: 1000,
    },
}
```

### Working with Contract Invocation Events

```go
import stellarv1 "github.com/withObsrvr/flow-proto/go/gen/stellar/v1"

// Contract invocation with full execution context
invocation := &stellarv1.ContractInvocationEvent{
    LedgerSequence:    1000,
    TransactionHash:   "abc123...",
    ContractId:        "CDLZFC3SYJYDZT7K67VZ75HPJVIEUVNIXF47ZG2FB2RMQQVU2HHGCYSC",
    InvokingAccount:   "GABC...",
    FunctionName:      "transfer",
    Arguments:         []string{`{"from":"GABC..."}`, `{"to":"GXYZ..."}`, `{"amount":"1000000"}`},
    Successful:        true,
    DiagnosticEvents:  []*stellarv1.DiagnosticEvent{...},  // Debug/monitoring events
    ContractCalls:     []*stellarv1.ContractCall{...},     // Cross-contract call tree
    StateChanges:      []*stellarv1.StateChange{...},      // Storage modifications
}
```

### Working with Token Transfer Events

```go
// SEP-41 compliant token transfer
transfer := &stellarv1.TokenTransferEvent{
    Meta: &stellarv1.TokenTransferEventMeta{
        LedgerSequence:   1000,
        TxHash:          "abc123...",
        TransactionIndex: 1,
        ContractAddress:  "CDLZFC3SYJYDZT7K67VZ75HPJVIEUVNIXF47ZG2FB2RMQQVU2HHGCYSC",
    },
    Event: &stellarv1.TokenTransferEvent_Transfer{
        Transfer: &stellarv1.Transfer{
            From:   "GABC...",
            To:     "GXYZ...",
            Asset:  &stellarv1.Asset{Asset: &stellarv1.Asset_Native{Native: true}},
            Amount: "1000000",
        },
    },
}
```

## Development

### Prerequisites

- Go 1.21+
- protoc compiler
- protoc-gen-go and protoc-gen-go-grpc plugins

Or use Nix (recommended):

```bash
nix develop  # Enters development environment with all tools
```

The Nix flake provides:
- Go 1.25.4 toolchain
- protoc 32.1 with protoc-gen-go and protoc-gen-go-grpc plugins
- Git and GNU Make

### Generating Proto Code

```bash
# Generate flowctl/v1 and stellar/v1 protos
make proto-new

# Generate everything including legacy protos
make proto

# Verify proto files compile
make verify

# Clean generated code
make clean
```

### Available Make Targets

```bash
make proto-new    # Generate flowctl/v1 and stellar/v1 protos (default)
make proto-legacy # Generate legacy protos for backward compatibility
make proto        # Generate both new and legacy protos
make verify       # Verify proto files compile
make clean        # Remove generated code
make tidy         # Generate protos and tidy go modules
make list         # List all proto files
make help         # Show help message
```

## Core Types

### Event

The standard envelope for all data flowing through flowctl:

```protobuf
message Event {
  string id = 1;                      // Unique identifier
  string type = 2;                    // Event type for routing
  bytes payload = 3;                  // Raw event data
  map<string, string> metadata = 4;   // Key-value metadata
  google.protobuf.Timestamp timestamp = 5;
  string source_component_id = 6;
  string content_type = 7;
  optional StellarCursor stellar_cursor = 8;  // For Stellar events
}
```

### ComponentInfo

Describes a component's capabilities:

```protobuf
message ComponentInfo {
  string id = 1;
  string name = 2;
  string description = 3;
  string version = 4;
  ComponentType type = 5;
  repeated string input_event_types = 6;
  repeated string output_event_types = 7;
  string endpoint = 8;
  map<string, string> metadata = 9;
}
```

## Services

### Core Services (flowctl/v1)

- **SourceService**: Streams events to downstream components
- **ProcessorService**: Transforms events via bidirectional streaming
- **ConsumerService**: Consumes and persists events
- **ControlPlaneService**: Component registration and discovery

### Stellar Services (stellar/v1)

- **RawLedgerService**: Streams Stellar ledger data (XDR format)
- **ContractEventService**: Streams Soroban contract events with filtering
- **ContractInvocationEvent**: Rich contract invocation data with execution context (diagnostic events, cross-contract calls, state changes)
- **TokenTransferEvent**: SEP-41 compliant token transfer events (transfer, mint, burn, clawback, fee)
- **AccountBalanceService**: Streams account balance changes
- **ContractDataService**: Provides contract data via Arrow Flight

## Versioning

All packages use `/v1/` suffix for future evolution:
- `flowctl.v1` - Core flow protocol
- `stellar.v1` - Stellar blockchain extensions

This allows backward-compatible evolution and parallel version support.

## Migration from Legacy Protos

If you're using old proto imports:

**Before:**
```go
import processorpb "github.com/withObsrvr/flow-proto/gen/processor"
import sourcepb "github.com/withObsrvr/flow-proto/gen/source"
```

**After:**
```go
import flowctlv1 "github.com/withObsrvr/flow-proto/go/gen/flowctl/v1"
```

See [MIGRATION.md](./MIGRATION.md) for detailed migration guide.

## Integration with Flowctl SDK

Use these protos with [flowctl-sdk](https://github.com/withObsrvr/flowctl-sdk) for easy component development:

```go
import (
    "github.com/withObsrvr/flowctl-sdk/pkg/processor"
    flowctlv1 "github.com/withObsrvr/flow-proto/go/gen/flowctl/v1"
)

func main() {
    proc := processor.New(processor.DefaultConfig())

    proc.OnProcess(func(ctx context.Context, event *flowctlv1.Event) (*flowctlv1.Event, error) {
        // Just write business logic - SDK handles gRPC
        return processEvent(event), nil
    })

    proc.Start(context.Background())
}
```

## Contributing

1. Update proto files in `proto/flowctl/v1/` or `proto/stellar/v1/`
2. Run `make verify` to check syntax
3. Run `make proto-new` to generate code
4. Update version in `go.mod` following semver
5. Push changes and create a release

## License

Apache License 2.0 - see LICENSE file for details.
