// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: completeTx.proto

package util

import (
	api "github.com/ranjbar-dev/tron-wallet/grpcClient/proto/api"
	core "github.com/ranjbar-dev/tron-wallet/grpcClient/proto/core"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type ProtoCompleteTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber  uint64                    `protobuf:"varint,1,opt,name=BlockNumber,proto3" json:"BlockNumber,omitempty"`
	BlockTime    uint64                    `protobuf:"varint,2,opt,name=BlockTime,proto3" json:"BlockTime,omitempty"`
	Tx           *api.TransactionExtention `protobuf:"bytes,3,opt,name=Tx,proto3" json:"Tx,omitempty"`
	Info         *core.TransactionInfo     `protobuf:"bytes,4,opt,name=Info,proto3" json:"Info,omitempty"`
	ExchangeInfo *core.Exchange            `protobuf:"bytes,5,opt,name=ExchangeInfo,proto3" json:"ExchangeInfo,omitempty"`
	AssetInfo    *core.AssetIssueContract  `protobuf:"bytes,6,opt,name=AssetInfo,proto3" json:"AssetInfo,omitempty"`
}

func (x *ProtoCompleteTransaction) Reset() {
	*x = ProtoCompleteTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_completeTx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoCompleteTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoCompleteTransaction) ProtoMessage() {}

func (x *ProtoCompleteTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_completeTx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoCompleteTransaction.ProtoReflect.Descriptor instead.
func (*ProtoCompleteTransaction) Descriptor() ([]byte, []int) {
	return file_completeTx_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoCompleteTransaction) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *ProtoCompleteTransaction) GetBlockTime() uint64 {
	if x != nil {
		return x.BlockTime
	}
	return 0
}

func (x *ProtoCompleteTransaction) GetTx() *api.TransactionExtention {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *ProtoCompleteTransaction) GetInfo() *core.TransactionInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ProtoCompleteTransaction) GetExchangeInfo() *core.Exchange {
	if x != nil {
		return x.ExchangeInfo
	}
	return nil
}

func (x *ProtoCompleteTransaction) GetAssetInfo() *core.AssetIssueContract {
	if x != nil {
		return x.AssetInfo
	}
	return nil
}

var File_completeTx_proto protoreflect.FileDescriptor

var file_completeTx_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x0f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x54, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x02, 0x54, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x02, 0x54, 0x78, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0c, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x09, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x09, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x62, 0x73, 0x6f, 0x62, 0x72, 0x65, 0x69, 0x72, 0x61, 0x2f,
	0x67, 0x6f, 0x74, 0x72, 0x6f, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_completeTx_proto_rawDescOnce sync.Once
	file_completeTx_proto_rawDescData = file_completeTx_proto_rawDesc
)

func file_completeTx_proto_rawDescGZIP() []byte {
	file_completeTx_proto_rawDescOnce.Do(func() {
		file_completeTx_proto_rawDescData = protoimpl.X.CompressGZIP(file_completeTx_proto_rawDescData)
	})
	return file_completeTx_proto_rawDescData
}

var file_completeTx_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_completeTx_proto_goTypes = []interface{}{
	(*ProtoCompleteTransaction)(nil), // 0: protocol.ProtoCompleteTransaction
	(*api.TransactionExtention)(nil), // 1: protocol.TransactionExtention
	(*core.TransactionInfo)(nil),     // 2: protocol.TransactionInfo
	(*core.Exchange)(nil),            // 3: protocol.Exchange
	(*core.AssetIssueContract)(nil),  // 4: protocol.AssetIssueContract
}
var file_completeTx_proto_depIdxs = []int32{
	1, // 0: protocol.ProtoCompleteTransaction.Tx:type_name -> protocol.TransactionExtention
	2, // 1: protocol.ProtoCompleteTransaction.Info:type_name -> protocol.TransactionInfo
	3, // 2: protocol.ProtoCompleteTransaction.ExchangeInfo:type_name -> protocol.Exchange
	4, // 3: protocol.ProtoCompleteTransaction.AssetInfo:type_name -> protocol.AssetIssueContract
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_completeTx_proto_init() }
func file_completeTx_proto_init() {
	if File_completeTx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_completeTx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoCompleteTransaction); i {
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
			RawDescriptor: file_completeTx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_completeTx_proto_goTypes,
		DependencyIndexes: file_completeTx_proto_depIdxs,
		MessageInfos:      file_completeTx_proto_msgTypes,
	}.Build()
	File_completeTx_proto = out.File
	file_completeTx_proto_rawDesc = nil
	file_completeTx_proto_goTypes = nil
	file_completeTx_proto_depIdxs = nil
}
