// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: steammessages_oauth.steamworkssdk.proto

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

type COAuthToken_ImplicitGrantNoPrompt_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clientid *string `protobuf:"bytes,1,opt,name=clientid" json:"clientid,omitempty"`
}

func (x *COAuthToken_ImplicitGrantNoPrompt_Request) Reset() {
	*x = COAuthToken_ImplicitGrantNoPrompt_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_oauth_steamworkssdk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COAuthToken_ImplicitGrantNoPrompt_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COAuthToken_ImplicitGrantNoPrompt_Request) ProtoMessage() {}

func (x *COAuthToken_ImplicitGrantNoPrompt_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_oauth_steamworkssdk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COAuthToken_ImplicitGrantNoPrompt_Request.ProtoReflect.Descriptor instead.
func (*COAuthToken_ImplicitGrantNoPrompt_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_oauth_steamworkssdk_proto_rawDescGZIP(), []int{0}
}

func (x *COAuthToken_ImplicitGrantNoPrompt_Request) GetClientid() string {
	if x != nil && x.Clientid != nil {
		return *x.Clientid
	}
	return ""
}

type COAuthToken_ImplicitGrantNoPrompt_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken *string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RedirectUri *string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri" json:"redirect_uri,omitempty"`
}

func (x *COAuthToken_ImplicitGrantNoPrompt_Response) Reset() {
	*x = COAuthToken_ImplicitGrantNoPrompt_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_oauth_steamworkssdk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *COAuthToken_ImplicitGrantNoPrompt_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*COAuthToken_ImplicitGrantNoPrompt_Response) ProtoMessage() {}

func (x *COAuthToken_ImplicitGrantNoPrompt_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_oauth_steamworkssdk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use COAuthToken_ImplicitGrantNoPrompt_Response.ProtoReflect.Descriptor instead.
func (*COAuthToken_ImplicitGrantNoPrompt_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_oauth_steamworkssdk_proto_rawDescGZIP(), []int{1}
}

func (x *COAuthToken_ImplicitGrantNoPrompt_Response) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

func (x *COAuthToken_ImplicitGrantNoPrompt_Response) GetRedirectUri() string {
	if x != nil && x.RedirectUri != nil {
		return *x.RedirectUri
	}
	return ""
}

var File_steammessages_oauth_steamworkssdk_proto protoreflect.FileDescriptor

var file_steammessages_oauth_steamworkssdk_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x29, 0x43, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3c, 0x82, 0xb5, 0x18, 0x38, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x49, 0x44, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x77, 0x68, 0x69,
	0x63, 0x68, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64,
	0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69,
	0x64, 0x22, 0xd1, 0x01, 0x0a, 0x2a, 0x43, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4e,
	0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x23, 0x82, 0xb5, 0x18, 0x1f, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x20, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x20, 0x6f, 0x6e, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x5b, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38,
	0x82, 0xb5, 0x18, 0x34, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x55, 0x52, 0x49, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x64, 0x20, 0x64, 0x75, 0x72,
	0x69, 0x6e, 0x67, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x55, 0x72, 0x69, 0x32, 0xb1, 0x02, 0x0a, 0x0a, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0xeb, 0x01, 0x0a, 0x15, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x2a,
	0x2e, 0x43, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x49, 0x6d, 0x70,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x6f, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x43, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69,
	0x74, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x4e, 0x6f, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x82, 0xb5, 0x18, 0x75, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x20,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x28, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x27, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x27, 0x29,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x20, 0x49, 0x44, 0x20, 0x6f, 0x6e, 0x20,
	0x62, 0x65, 0x68, 0x61, 0x6c, 0x66, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72,
	0x20, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x69,
	0x6e, 0x67, 0x1a, 0x35, 0x82, 0xb5, 0x18, 0x31, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x20, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
}

var (
	file_steammessages_oauth_steamworkssdk_proto_rawDescOnce sync.Once
	file_steammessages_oauth_steamworkssdk_proto_rawDescData = file_steammessages_oauth_steamworkssdk_proto_rawDesc
)

func file_steammessages_oauth_steamworkssdk_proto_rawDescGZIP() []byte {
	file_steammessages_oauth_steamworkssdk_proto_rawDescOnce.Do(func() {
		file_steammessages_oauth_steamworkssdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_oauth_steamworkssdk_proto_rawDescData)
	})
	return file_steammessages_oauth_steamworkssdk_proto_rawDescData
}

var file_steammessages_oauth_steamworkssdk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_steammessages_oauth_steamworkssdk_proto_goTypes = []any{
	(*COAuthToken_ImplicitGrantNoPrompt_Request)(nil),  // 0: COAuthToken_ImplicitGrantNoPrompt_Request
	(*COAuthToken_ImplicitGrantNoPrompt_Response)(nil), // 1: COAuthToken_ImplicitGrantNoPrompt_Response
}
var file_steammessages_oauth_steamworkssdk_proto_depIdxs = []int32{
	0, // 0: OAuthToken.ImplicitGrantNoPrompt:input_type -> COAuthToken_ImplicitGrantNoPrompt_Request
	1, // 1: OAuthToken.ImplicitGrantNoPrompt:output_type -> COAuthToken_ImplicitGrantNoPrompt_Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steammessages_oauth_steamworkssdk_proto_init() }
func file_steammessages_oauth_steamworkssdk_proto_init() {
	if File_steammessages_oauth_steamworkssdk_proto != nil {
		return
	}
	file_steammessages_unified_base_steamworkssdk_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_oauth_steamworkssdk_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*COAuthToken_ImplicitGrantNoPrompt_Request); i {
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
		file_steammessages_oauth_steamworkssdk_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*COAuthToken_ImplicitGrantNoPrompt_Response); i {
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
			RawDescriptor: file_steammessages_oauth_steamworkssdk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steammessages_oauth_steamworkssdk_proto_goTypes,
		DependencyIndexes: file_steammessages_oauth_steamworkssdk_proto_depIdxs,
		MessageInfos:      file_steammessages_oauth_steamworkssdk_proto_msgTypes,
	}.Build()
	File_steammessages_oauth_steamworkssdk_proto = out.File
	file_steammessages_oauth_steamworkssdk_proto_rawDesc = nil
	file_steammessages_oauth_steamworkssdk_proto_goTypes = nil
	file_steammessages_oauth_steamworkssdk_proto_depIdxs = nil
}
