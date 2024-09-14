// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: service_mobileapp.proto

package webui

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

type CMobileApp_GetMobileSummary_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthenticatorGid *uint64 `protobuf:"fixed64,1,opt,name=authenticator_gid,json=authenticatorGid" json:"authenticator_gid,omitempty"`
}

func (x *CMobileApp_GetMobileSummary_Request) Reset() {
	*x = CMobileApp_GetMobileSummary_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mobileapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMobileApp_GetMobileSummary_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMobileApp_GetMobileSummary_Request) ProtoMessage() {}

func (x *CMobileApp_GetMobileSummary_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_mobileapp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMobileApp_GetMobileSummary_Request.ProtoReflect.Descriptor instead.
func (*CMobileApp_GetMobileSummary_Request) Descriptor() ([]byte, []int) {
	return file_service_mobileapp_proto_rawDescGZIP(), []int{0}
}

func (x *CMobileApp_GetMobileSummary_Request) GetAuthenticatorGid() uint64 {
	if x != nil && x.AuthenticatorGid != nil {
		return *x.AuthenticatorGid
	}
	return 0
}

type CMobileApp_GetMobileSummary_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StaleTimeSeconds     *uint32 `protobuf:"varint,1,opt,name=stale_time_seconds,json=staleTimeSeconds" json:"stale_time_seconds,omitempty"`
	IsAuthenticatorValid *bool   `protobuf:"varint,2,opt,name=is_authenticator_valid,json=isAuthenticatorValid" json:"is_authenticator_valid,omitempty"`
	OwnedGames           *uint32 `protobuf:"varint,3,opt,name=owned_games,json=ownedGames" json:"owned_games,omitempty"`
	FriendCount          *uint32 `protobuf:"varint,4,opt,name=friend_count,json=friendCount" json:"friend_count,omitempty"`
	WalletBalance        *string `protobuf:"bytes,5,opt,name=wallet_balance,json=walletBalance" json:"wallet_balance,omitempty"`
	Language             *string `protobuf:"bytes,6,opt,name=language" json:"language,omitempty"`
}

func (x *CMobileApp_GetMobileSummary_Response) Reset() {
	*x = CMobileApp_GetMobileSummary_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mobileapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMobileApp_GetMobileSummary_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMobileApp_GetMobileSummary_Response) ProtoMessage() {}

func (x *CMobileApp_GetMobileSummary_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_mobileapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMobileApp_GetMobileSummary_Response.ProtoReflect.Descriptor instead.
func (*CMobileApp_GetMobileSummary_Response) Descriptor() ([]byte, []int) {
	return file_service_mobileapp_proto_rawDescGZIP(), []int{1}
}

func (x *CMobileApp_GetMobileSummary_Response) GetStaleTimeSeconds() uint32 {
	if x != nil && x.StaleTimeSeconds != nil {
		return *x.StaleTimeSeconds
	}
	return 0
}

func (x *CMobileApp_GetMobileSummary_Response) GetIsAuthenticatorValid() bool {
	if x != nil && x.IsAuthenticatorValid != nil {
		return *x.IsAuthenticatorValid
	}
	return false
}

func (x *CMobileApp_GetMobileSummary_Response) GetOwnedGames() uint32 {
	if x != nil && x.OwnedGames != nil {
		return *x.OwnedGames
	}
	return 0
}

func (x *CMobileApp_GetMobileSummary_Response) GetFriendCount() uint32 {
	if x != nil && x.FriendCount != nil {
		return *x.FriendCount
	}
	return 0
}

func (x *CMobileApp_GetMobileSummary_Response) GetWalletBalance() string {
	if x != nil && x.WalletBalance != nil {
		return *x.WalletBalance
	}
	return ""
}

func (x *CMobileApp_GetMobileSummary_Response) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

var File_service_mobileapp_proto protoreflect.FileDescriptor

var file_service_mobileapp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x23, 0x43, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x5f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x10, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x47, 0x69, 0x64, 0x22, 0x91, 0x02,
	0x0a, 0x24, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x5f, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x77,
	0x6e, 0x65, 0x64, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x32, 0x6c, 0x0a, 0x09, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x12, 0x5f,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x24, 0x2e, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x5f,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x43, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x41, 0x70, 0x70, 0x5f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
}

var (
	file_service_mobileapp_proto_rawDescOnce sync.Once
	file_service_mobileapp_proto_rawDescData = file_service_mobileapp_proto_rawDesc
)

func file_service_mobileapp_proto_rawDescGZIP() []byte {
	file_service_mobileapp_proto_rawDescOnce.Do(func() {
		file_service_mobileapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_mobileapp_proto_rawDescData)
	})
	return file_service_mobileapp_proto_rawDescData
}

var file_service_mobileapp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_mobileapp_proto_goTypes = []any{
	(*CMobileApp_GetMobileSummary_Request)(nil),  // 0: CMobileApp_GetMobileSummary_Request
	(*CMobileApp_GetMobileSummary_Response)(nil), // 1: CMobileApp_GetMobileSummary_Response
}
var file_service_mobileapp_proto_depIdxs = []int32{
	0, // 0: MobileApp.GetMobileSummary:input_type -> CMobileApp_GetMobileSummary_Request
	1, // 1: MobileApp.GetMobileSummary:output_type -> CMobileApp_GetMobileSummary_Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_mobileapp_proto_init() }
func file_service_mobileapp_proto_init() {
	if File_service_mobileapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_mobileapp_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMobileApp_GetMobileSummary_Request); i {
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
		file_service_mobileapp_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMobileApp_GetMobileSummary_Response); i {
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
			RawDescriptor: file_service_mobileapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_mobileapp_proto_goTypes,
		DependencyIndexes: file_service_mobileapp_proto_depIdxs,
		MessageInfos:      file_service_mobileapp_proto_msgTypes,
	}.Build()
	File_service_mobileapp_proto = out.File
	file_service_mobileapp_proto_rawDesc = nil
	file_service_mobileapp_proto_goTypes = nil
	file_service_mobileapp_proto_depIdxs = nil
}