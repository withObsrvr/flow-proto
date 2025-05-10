package main

import (
	"fmt"

	"github.com/withObsrvr/flow-proto/go/gen/consumer"
	"github.com/withObsrvr/flow-proto/go/gen/processor"
	"github.com/withObsrvr/flow-proto/go/gen/source"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
	// Create metadata struct for processor
	procMetadata, _ := structpb.NewStruct(map[string]interface{}{
		"test": "value",
	})

	// Create some example proto messages
	procMsg := &processor.DataMessage{
		Payload:  []byte("hello"),
		Metadata: procMetadata,
	}

	sourceMsg := &source.CapabilitiesRequest{}

	// Create config struct for consumer
	consumerConfig, _ := structpb.NewStruct(map[string]interface{}{
		"key": "value",
	})

	consumerMsg := &consumer.ConfigureRequest{
		Config: consumerConfig,
	}

	fmt.Printf("Created processor message: %v\n", procMsg)
	fmt.Printf("Created source message: %v\n", sourceMsg)
	fmt.Printf("Created consumer message: %v\n", consumerMsg)
}
