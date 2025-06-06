// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: flow/v1/envelope.proto

package flowv1

import (
	token_transfer "github.com/withObsrvr/flow-proto/proto/ingest/processors/token_transfer"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Kind int32

const (
	Kind_KIND_UNSPECIFIED Kind = 0
	Kind_LEDGER_META      Kind = 1 // raw LedgerCloseMeta XDR
	Kind_NETWORK_EVENT    Kind = 2 // Soroban Event, payment op, etc. – raw XDR
	Kind_TOKEN_TRANSFER   Kind = 3 // decoded, schema‑safe event
	Kind_RESERVED4        Kind = 4 // leave holes for future native kinds
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "KIND_UNSPECIFIED",
		1: "LEDGER_META",
		2: "NETWORK_EVENT",
		3: "TOKEN_TRANSFER",
		4: "RESERVED4",
	}
	Kind_value = map[string]int32{
		"KIND_UNSPECIFIED": 0,
		"LEDGER_META":      1,
		"NETWORK_EVENT":    2,
		"TOKEN_TRANSFER":   3,
		"RESERVED4":        4,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_flow_v1_envelope_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_flow_v1_envelope_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_flow_v1_envelope_proto_rawDescGZIP(), []int{0}
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * What is inside. Helps routers decide fast.
	Kind Kind `protobuf:"varint,1,opt,name=kind,proto3,enum=flow.v1.Kind" json:"kind,omitempty"`
	// * Where in the chain this item belongs.
	Cursor *Cursor `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// * Nanosecond wall‑clock when the source observed it.
	ObservedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=observed_at,json=observedAt,proto3" json:"observed_at,omitempty"`
	// * One‑of payload.  TokenTransferEvent is typed; others stay raw.
	// Types that are assignable to Payload:
	//
	//	*Envelope_LedgerMetaXdr
	//	*Envelope_EventXdr
	//	*Envelope_TokenTransfer
	Payload isEnvelope_Payload `protobuf_oneof:"payload"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_v1_envelope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_envelope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_flow_v1_envelope_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_KIND_UNSPECIFIED
}

func (x *Envelope) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *Envelope) GetObservedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ObservedAt
	}
	return nil
}

func (m *Envelope) GetPayload() isEnvelope_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Envelope) GetLedgerMetaXdr() []byte {
	if x, ok := x.GetPayload().(*Envelope_LedgerMetaXdr); ok {
		return x.LedgerMetaXdr
	}
	return nil
}

func (x *Envelope) GetEventXdr() []byte {
	if x, ok := x.GetPayload().(*Envelope_EventXdr); ok {
		return x.EventXdr
	}
	return nil
}

func (x *Envelope) GetTokenTransfer() *token_transfer.TokenTransferEvent {
	if x, ok := x.GetPayload().(*Envelope_TokenTransfer); ok {
		return x.TokenTransfer
	}
	return nil
}

type isEnvelope_Payload interface {
	isEnvelope_Payload()
}

type Envelope_LedgerMetaXdr struct {
	LedgerMetaXdr []byte `protobuf:"bytes,10,opt,name=ledger_meta_xdr,json=ledgerMetaXdr,proto3,oneof"` // Kind = LEDGER_META
}

type Envelope_EventXdr struct {
	EventXdr []byte `protobuf:"bytes,11,opt,name=event_xdr,json=eventXdr,proto3,oneof"` // Kind = NETWORK_EVENT
}

type Envelope_TokenTransfer struct {
	TokenTransfer *token_transfer.TokenTransferEvent `protobuf:"bytes,12,opt,name=token_transfer,json=tokenTransfer,proto3,oneof"` // Kind = TOKEN_TRANSFER
}

func (*Envelope_LedgerMetaXdr) isEnvelope_Payload() {}

func (*Envelope_EventXdr) isEnvelope_Payload() {}

func (*Envelope_TokenTransfer) isEnvelope_Payload() {}

var File_flow_v1_envelope_proto protoreflect.FileDescriptor

var file_flow_v1_envelope_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76,
	0x31, 0x1a, 0x14, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x3b,
	0x0a, 0x0b, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x78, 0x64, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x58, 0x64, 0x72, 0x12, 0x1d, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x78,
	0x64, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x58, 0x64, 0x72, 0x12, 0x4b, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a, 0x63, 0x0a, 0x04,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45,
	0x44, 0x47, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x34, 0x10,
	0x04, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x77, 0x69, 0x74, 0x68, 0x4f, 0x62, 0x73, 0x72, 0x76, 0x72, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x76, 0x31, 0x3b, 0x66, 0x6c, 0x6f, 0x77, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_flow_v1_envelope_proto_rawDescOnce sync.Once
	file_flow_v1_envelope_proto_rawDescData = file_flow_v1_envelope_proto_rawDesc
)

func file_flow_v1_envelope_proto_rawDescGZIP() []byte {
	file_flow_v1_envelope_proto_rawDescOnce.Do(func() {
		file_flow_v1_envelope_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_v1_envelope_proto_rawDescData)
	})
	return file_flow_v1_envelope_proto_rawDescData
}

var file_flow_v1_envelope_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_flow_v1_envelope_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flow_v1_envelope_proto_goTypes = []interface{}{
	(Kind)(0),                     // 0: flow.v1.Kind
	(*Envelope)(nil),              // 1: flow.v1.Envelope
	(*Cursor)(nil),                // 2: flow.v1.Cursor
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*token_transfer.TokenTransferEvent)(nil), // 4: token_transfer.TokenTransferEvent
}
var file_flow_v1_envelope_proto_depIdxs = []int32{
	0, // 0: flow.v1.Envelope.kind:type_name -> flow.v1.Kind
	2, // 1: flow.v1.Envelope.cursor:type_name -> flow.v1.Cursor
	3, // 2: flow.v1.Envelope.observed_at:type_name -> google.protobuf.Timestamp
	4, // 3: flow.v1.Envelope.token_transfer:type_name -> token_transfer.TokenTransferEvent
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_flow_v1_envelope_proto_init() }
func file_flow_v1_envelope_proto_init() {
	if File_flow_v1_envelope_proto != nil {
		return
	}
	file_flow_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flow_v1_envelope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_flow_v1_envelope_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Envelope_LedgerMetaXdr)(nil),
		(*Envelope_EventXdr)(nil),
		(*Envelope_TokenTransfer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flow_v1_envelope_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_v1_envelope_proto_goTypes,
		DependencyIndexes: file_flow_v1_envelope_proto_depIdxs,
		EnumInfos:         file_flow_v1_envelope_proto_enumTypes,
		MessageInfos:      file_flow_v1_envelope_proto_msgTypes,
	}.Build()
	File_flow_v1_envelope_proto = out.File
	file_flow_v1_envelope_proto_rawDesc = nil
	file_flow_v1_envelope_proto_goTypes = nil
	file_flow_v1_envelope_proto_depIdxs = nil
}
