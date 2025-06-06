// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v4.24.4
// source: flow/v1/common.proto

package flowv1

import (
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

// * Stable position in the ledger stream.
type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LedgerSeq uint64 `protobuf:"varint,1,opt,name=ledger_seq,json=ledgerSeq,proto3" json:"ledger_seq,omitempty"`
	Index     uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"` // index inside ledger (0 for meta)
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_flow_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *Cursor) GetLedgerSeq() uint64 {
	if x != nil {
		return x.LedgerSeq
	}
	return 0
}

func (x *Cursor) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

// * Used by sources & processors to resume after crash / reconnect.
type ResumeToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor      *Cursor                `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Opaque      []byte                 `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"` // vendor‑specific checkpoint
	GeneratedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=generated_at,json=generatedAt,proto3" json:"generated_at,omitempty"`
}

func (x *ResumeToken) Reset() {
	*x = ResumeToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flow_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResumeToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResumeToken) ProtoMessage() {}

func (x *ResumeToken) ProtoReflect() protoreflect.Message {
	mi := &file_flow_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResumeToken.ProtoReflect.Descriptor instead.
func (*ResumeToken) Descriptor() ([]byte, []int) {
	return file_flow_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *ResumeToken) GetCursor() *Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ResumeToken) GetOpaque() []byte {
	if x != nil {
		return x.Opaque
	}
	return nil
}

func (x *ResumeToken) GetGeneratedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.GeneratedAt
	}
	return nil
}

var File_flow_v1_common_proto protoreflect.FileDescriptor

var file_flow_v1_common_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x53, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x8d, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x27, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65,
	0x12, 0x3d, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69,
	0x74, 0x68, 0x4f, 0x62, 0x73, 0x72, 0x76, 0x72, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76,
	0x31, 0x3b, 0x66, 0x6c, 0x6f, 0x77, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flow_v1_common_proto_rawDescOnce sync.Once
	file_flow_v1_common_proto_rawDescData = file_flow_v1_common_proto_rawDesc
)

func file_flow_v1_common_proto_rawDescGZIP() []byte {
	file_flow_v1_common_proto_rawDescOnce.Do(func() {
		file_flow_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_flow_v1_common_proto_rawDescData)
	})
	return file_flow_v1_common_proto_rawDescData
}

var file_flow_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flow_v1_common_proto_goTypes = []interface{}{
	(*Cursor)(nil),                // 0: flow.v1.Cursor
	(*ResumeToken)(nil),           // 1: flow.v1.ResumeToken
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_flow_v1_common_proto_depIdxs = []int32{
	0, // 0: flow.v1.ResumeToken.cursor:type_name -> flow.v1.Cursor
	2, // 1: flow.v1.ResumeToken.generated_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_flow_v1_common_proto_init() }
func file_flow_v1_common_proto_init() {
	if File_flow_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flow_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
		file_flow_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResumeToken); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flow_v1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flow_v1_common_proto_goTypes,
		DependencyIndexes: file_flow_v1_common_proto_depIdxs,
		MessageInfos:      file_flow_v1_common_proto_msgTypes,
	}.Build()
	File_flow_v1_common_proto = out.File
	file_flow_v1_common_proto_rawDesc = nil
	file_flow_v1_common_proto_goTypes = nil
	file_flow_v1_common_proto_depIdxs = nil
}
