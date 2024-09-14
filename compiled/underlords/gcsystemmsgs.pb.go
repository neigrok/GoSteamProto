// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: gcsystemmsgs.proto

package underlords

import (
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

type ESOMsg int32

const (
	ESOMsg_k_ESOMsg_Create                   ESOMsg = 21
	ESOMsg_k_ESOMsg_Update                   ESOMsg = 22
	ESOMsg_k_ESOMsg_Destroy                  ESOMsg = 23
	ESOMsg_k_ESOMsg_CacheSubscribed          ESOMsg = 24
	ESOMsg_k_ESOMsg_CacheUnsubscribed        ESOMsg = 25
	ESOMsg_k_ESOMsg_UpdateMultiple           ESOMsg = 26
	ESOMsg_k_ESOMsg_CacheSubscriptionRefresh ESOMsg = 28
	ESOMsg_k_ESOMsg_CacheSubscribedUpToDate  ESOMsg = 29
)

// Enum value maps for ESOMsg.
var (
	ESOMsg_name = map[int32]string{
		21: "k_ESOMsg_Create",
		22: "k_ESOMsg_Update",
		23: "k_ESOMsg_Destroy",
		24: "k_ESOMsg_CacheSubscribed",
		25: "k_ESOMsg_CacheUnsubscribed",
		26: "k_ESOMsg_UpdateMultiple",
		28: "k_ESOMsg_CacheSubscriptionRefresh",
		29: "k_ESOMsg_CacheSubscribedUpToDate",
	}
	ESOMsg_value = map[string]int32{
		"k_ESOMsg_Create":                   21,
		"k_ESOMsg_Update":                   22,
		"k_ESOMsg_Destroy":                  23,
		"k_ESOMsg_CacheSubscribed":          24,
		"k_ESOMsg_CacheUnsubscribed":        25,
		"k_ESOMsg_UpdateMultiple":           26,
		"k_ESOMsg_CacheSubscriptionRefresh": 28,
		"k_ESOMsg_CacheSubscribedUpToDate":  29,
	}
)

func (x ESOMsg) Enum() *ESOMsg {
	p := new(ESOMsg)
	*p = x
	return p
}

func (x ESOMsg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ESOMsg) Descriptor() protoreflect.EnumDescriptor {
	return file_gcsystemmsgs_proto_enumTypes[0].Descriptor()
}

func (ESOMsg) Type() protoreflect.EnumType {
	return &file_gcsystemmsgs_proto_enumTypes[0]
}

func (x ESOMsg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ESOMsg) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ESOMsg(num)
	return nil
}

// Deprecated: Use ESOMsg.Descriptor instead.
func (ESOMsg) EnumDescriptor() ([]byte, []int) {
	return file_gcsystemmsgs_proto_rawDescGZIP(), []int{0}
}

type EGCBaseClientMsg int32

const (
	EGCBaseClientMsg_k_EMsgGCPingRequest                  EGCBaseClientMsg = 3001
	EGCBaseClientMsg_k_EMsgGCPingResponse                 EGCBaseClientMsg = 3002
	EGCBaseClientMsg_k_EMsgGCToClientPollConvarRequest    EGCBaseClientMsg = 3003
	EGCBaseClientMsg_k_EMsgGCToClientPollConvarResponse   EGCBaseClientMsg = 3004
	EGCBaseClientMsg_k_EMsgGCCompressedMsgToClient        EGCBaseClientMsg = 3005
	EGCBaseClientMsg_k_EMsgGCCompressedMsgToClient_Legacy EGCBaseClientMsg = 523
	EGCBaseClientMsg_k_EMsgGCToClientRequestDropped       EGCBaseClientMsg = 3006
	EGCBaseClientMsg_k_EMsgGCClientWelcome                EGCBaseClientMsg = 4004
	EGCBaseClientMsg_k_EMsgGCServerWelcome                EGCBaseClientMsg = 4005
	EGCBaseClientMsg_k_EMsgGCClientHello                  EGCBaseClientMsg = 4006
	EGCBaseClientMsg_k_EMsgGCServerHello                  EGCBaseClientMsg = 4007
	EGCBaseClientMsg_k_EMsgGCClientConnectionStatus       EGCBaseClientMsg = 4009
	EGCBaseClientMsg_k_EMsgGCServerConnectionStatus       EGCBaseClientMsg = 4010
)

// Enum value maps for EGCBaseClientMsg.
var (
	EGCBaseClientMsg_name = map[int32]string{
		3001: "k_EMsgGCPingRequest",
		3002: "k_EMsgGCPingResponse",
		3003: "k_EMsgGCToClientPollConvarRequest",
		3004: "k_EMsgGCToClientPollConvarResponse",
		3005: "k_EMsgGCCompressedMsgToClient",
		523:  "k_EMsgGCCompressedMsgToClient_Legacy",
		3006: "k_EMsgGCToClientRequestDropped",
		4004: "k_EMsgGCClientWelcome",
		4005: "k_EMsgGCServerWelcome",
		4006: "k_EMsgGCClientHello",
		4007: "k_EMsgGCServerHello",
		4009: "k_EMsgGCClientConnectionStatus",
		4010: "k_EMsgGCServerConnectionStatus",
	}
	EGCBaseClientMsg_value = map[string]int32{
		"k_EMsgGCPingRequest":                  3001,
		"k_EMsgGCPingResponse":                 3002,
		"k_EMsgGCToClientPollConvarRequest":    3003,
		"k_EMsgGCToClientPollConvarResponse":   3004,
		"k_EMsgGCCompressedMsgToClient":        3005,
		"k_EMsgGCCompressedMsgToClient_Legacy": 523,
		"k_EMsgGCToClientRequestDropped":       3006,
		"k_EMsgGCClientWelcome":                4004,
		"k_EMsgGCServerWelcome":                4005,
		"k_EMsgGCClientHello":                  4006,
		"k_EMsgGCServerHello":                  4007,
		"k_EMsgGCClientConnectionStatus":       4009,
		"k_EMsgGCServerConnectionStatus":       4010,
	}
)

func (x EGCBaseClientMsg) Enum() *EGCBaseClientMsg {
	p := new(EGCBaseClientMsg)
	*p = x
	return p
}

func (x EGCBaseClientMsg) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EGCBaseClientMsg) Descriptor() protoreflect.EnumDescriptor {
	return file_gcsystemmsgs_proto_enumTypes[1].Descriptor()
}

func (EGCBaseClientMsg) Type() protoreflect.EnumType {
	return &file_gcsystemmsgs_proto_enumTypes[1]
}

func (x EGCBaseClientMsg) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EGCBaseClientMsg) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EGCBaseClientMsg(num)
	return nil
}

// Deprecated: Use EGCBaseClientMsg.Descriptor instead.
func (EGCBaseClientMsg) EnumDescriptor() ([]byte, []int) {
	return file_gcsystemmsgs_proto_rawDescGZIP(), []int{1}
}

var File_gcsystemmsgs_proto protoreflect.FileDescriptor

var file_gcsystemmsgs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x63, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x6d, 0x73, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf0, 0x01, 0x0a, 0x06, 0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67, 0x12,
	0x13, 0x0a, 0x0f, 0x6b, 0x5f, 0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67, 0x5f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x10, 0x15, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x5f, 0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67,
	0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x16, 0x12, 0x14, 0x0a, 0x10, 0x6b, 0x5f, 0x45,
	0x53, 0x4f, 0x4d, 0x73, 0x67, 0x5f, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x10, 0x17, 0x12,
	0x1c, 0x0a, 0x18, 0x6b, 0x5f, 0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67, 0x5f, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x10, 0x18, 0x12, 0x1e, 0x0a,
	0x1a, 0x6b, 0x5f, 0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67, 0x5f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x55,
	0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x10, 0x19, 0x12, 0x1b, 0x0a,
	0x17, 0x6b, 0x5f, 0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67, 0x5f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x10, 0x1a, 0x12, 0x25, 0x0a, 0x21, 0x6b, 0x5f,
	0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67, 0x5f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x10,
	0x1c, 0x12, 0x24, 0x0a, 0x20, 0x6b, 0x5f, 0x45, 0x53, 0x4f, 0x4d, 0x73, 0x67, 0x5f, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x55, 0x70, 0x54,
	0x6f, 0x44, 0x61, 0x74, 0x65, 0x10, 0x1d, 0x2a, 0xc2, 0x03, 0x0a, 0x10, 0x45, 0x47, 0x43, 0x42,
	0x61, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x13,
	0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x43, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x10, 0xb9, 0x17, 0x12, 0x19, 0x0a, 0x14, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67,
	0x47, 0x43, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0xba,
	0x17, 0x12, 0x26, 0x0a, 0x21, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x43, 0x54, 0x6f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x76, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0xbb, 0x17, 0x12, 0x27, 0x0a, 0x22, 0x6b, 0x5f, 0x45,
	0x4d, 0x73, 0x67, 0x47, 0x43, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c,
	0x6c, 0x43, 0x6f, 0x6e, 0x76, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10,
	0xbc, 0x17, 0x12, 0x22, 0x0a, 0x1d, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x43, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x10, 0xbd, 0x17, 0x12, 0x29, 0x0a, 0x24, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67,
	0x47, 0x43, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x4d, 0x73, 0x67, 0x54,
	0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x4c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x10, 0x8b,
	0x04, 0x12, 0x23, 0x0a, 0x1e, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x43, 0x54, 0x6f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x72, 0x6f, 0x70,
	0x70, 0x65, 0x64, 0x10, 0xbe, 0x17, 0x12, 0x1a, 0x0a, 0x15, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67,
	0x47, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x10,
	0xa4, 0x1f, 0x12, 0x1a, 0x0a, 0x15, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x43, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x57, 0x65, 0x6c, 0x63, 0x6f, 0x6d, 0x65, 0x10, 0xa5, 0x1f, 0x12, 0x18,
	0x0a, 0x13, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x10, 0xa6, 0x1f, 0x12, 0x18, 0x0a, 0x13, 0x6b, 0x5f, 0x45, 0x4d,
	0x73, 0x67, 0x47, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x10,
	0xa7, 0x1f, 0x12, 0x23, 0x0a, 0x1e, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x43, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x10, 0xa9, 0x1f, 0x12, 0x23, 0x0a, 0x1e, 0x6b, 0x5f, 0x45, 0x4d, 0x73,
	0x67, 0x47, 0x43, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0xaa, 0x1f, 0x42, 0x05, 0x48, 0x01,
	0x80, 0x01, 0x00,
}

var (
	file_gcsystemmsgs_proto_rawDescOnce sync.Once
	file_gcsystemmsgs_proto_rawDescData = file_gcsystemmsgs_proto_rawDesc
)

func file_gcsystemmsgs_proto_rawDescGZIP() []byte {
	file_gcsystemmsgs_proto_rawDescOnce.Do(func() {
		file_gcsystemmsgs_proto_rawDescData = protoimpl.X.CompressGZIP(file_gcsystemmsgs_proto_rawDescData)
	})
	return file_gcsystemmsgs_proto_rawDescData
}

var file_gcsystemmsgs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gcsystemmsgs_proto_goTypes = []any{
	(ESOMsg)(0),           // 0: ESOMsg
	(EGCBaseClientMsg)(0), // 1: EGCBaseClientMsg
}
var file_gcsystemmsgs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gcsystemmsgs_proto_init() }
func file_gcsystemmsgs_proto_init() {
	if File_gcsystemmsgs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gcsystemmsgs_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gcsystemmsgs_proto_goTypes,
		DependencyIndexes: file_gcsystemmsgs_proto_depIdxs,
		EnumInfos:         file_gcsystemmsgs_proto_enumTypes,
	}.Build()
	File_gcsystemmsgs_proto = out.File
	file_gcsystemmsgs_proto_rawDesc = nil
	file_gcsystemmsgs_proto_goTypes = nil
	file_gcsystemmsgs_proto_depIdxs = nil
}
