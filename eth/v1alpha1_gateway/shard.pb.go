// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: eth/v1alpha1/shard.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DataCommitment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point  []byte `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
	Length uint64 `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *DataCommitment) Reset() {
	*x = DataCommitment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataCommitment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataCommitment) ProtoMessage() {}

func (x *DataCommitment) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataCommitment.ProtoReflect.Descriptor instead.
func (*DataCommitment) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{0}
}

func (x *DataCommitment) GetPoint() []byte {
	if x != nil {
		return x.Point
	}
	return nil
}

func (x *DataCommitment) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ShardHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot        uint64          `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Shard       uint64          `protobuf:"varint,2,opt,name=shard,proto3" json:"shard,omitempty"`
	Commitment  *DataCommitment `protobuf:"bytes,3,opt,name=commitment,proto3" json:"commitment,omitempty"`
	LengthProof []byte          `protobuf:"bytes,4,opt,name=length_proof,json=lengthProof,proto3" json:"length_proof,omitempty"`
}

func (x *ShardHeader) Reset() {
	*x = ShardHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardHeader) ProtoMessage() {}

func (x *ShardHeader) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardHeader.ProtoReflect.Descriptor instead.
func (*ShardHeader) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{1}
}

func (x *ShardHeader) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *ShardHeader) GetShard() uint64 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *ShardHeader) GetCommitment() *DataCommitment {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *ShardHeader) GetLengthProof() []byte {
	if x != nil {
		return x.LengthProof
	}
	return nil
}

var File_eth_v1alpha1_shard_proto protoreflect.FileDescriptor

var file_eth_v1alpha1_shard_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x51, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a,
	0x22, 0x34, 0x38, 0x22, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x22, 0xb4, 0x01, 0x0a, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x45, 0x0a,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0c, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d,
	0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x34, 0x38, 0x22, 0x52, 0x0b, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x42, 0x50, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e,
	0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v1alpha1_shard_proto_rawDescOnce sync.Once
	file_eth_v1alpha1_shard_proto_rawDescData = file_eth_v1alpha1_shard_proto_rawDesc
)

func file_eth_v1alpha1_shard_proto_rawDescGZIP() []byte {
	file_eth_v1alpha1_shard_proto_rawDescOnce.Do(func() {
		file_eth_v1alpha1_shard_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1alpha1_shard_proto_rawDescData)
	})
	return file_eth_v1alpha1_shard_proto_rawDescData
}

var file_eth_v1alpha1_shard_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eth_v1alpha1_shard_proto_goTypes = []interface{}{
	(*DataCommitment)(nil), // 0: ethereum.eth.v1alpha1.DataCommitment
	(*ShardHeader)(nil),    // 1: ethereum.eth.v1alpha1.ShardHeader
}
var file_eth_v1alpha1_shard_proto_depIdxs = []int32{
	0, // 0: ethereum.eth.v1alpha1.ShardHeader.commitment:type_name -> ethereum.eth.v1alpha1.DataCommitment
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_eth_v1alpha1_shard_proto_init() }
func file_eth_v1alpha1_shard_proto_init() {
	if File_eth_v1alpha1_shard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v1alpha1_shard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataCommitment); i {
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
		file_eth_v1alpha1_shard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardHeader); i {
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
			RawDescriptor: file_eth_v1alpha1_shard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1alpha1_shard_proto_goTypes,
		DependencyIndexes: file_eth_v1alpha1_shard_proto_depIdxs,
		MessageInfos:      file_eth_v1alpha1_shard_proto_msgTypes,
	}.Build()
	File_eth_v1alpha1_shard_proto = out.File
	file_eth_v1alpha1_shard_proto_rawDesc = nil
	file_eth_v1alpha1_shard_proto_goTypes = nil
	file_eth_v1alpha1_shard_proto_depIdxs = nil
}
