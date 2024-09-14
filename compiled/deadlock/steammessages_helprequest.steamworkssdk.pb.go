// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: steammessages_helprequest.steamworkssdk.proto

package deadlock

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

type CHelpRequestLogs_UploadUserApplicationLog_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid         *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	LogType       *string `protobuf:"bytes,2,opt,name=log_type,json=logType" json:"log_type,omitempty"`
	VersionString *string `protobuf:"bytes,3,opt,name=version_string,json=versionString" json:"version_string,omitempty"`
	LogContents   *string `protobuf:"bytes,4,opt,name=log_contents,json=logContents" json:"log_contents,omitempty"`
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Request) Reset() {
	*x = CHelpRequestLogs_UploadUserApplicationLog_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_helprequest_steamworkssdk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CHelpRequestLogs_UploadUserApplicationLog_Request) ProtoMessage() {}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Request) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_helprequest_steamworkssdk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CHelpRequestLogs_UploadUserApplicationLog_Request.ProtoReflect.Descriptor instead.
func (*CHelpRequestLogs_UploadUserApplicationLog_Request) Descriptor() ([]byte, []int) {
	return file_steammessages_helprequest_steamworkssdk_proto_rawDescGZIP(), []int{0}
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Request) GetLogType() string {
	if x != nil && x.LogType != nil {
		return *x.LogType
	}
	return ""
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Request) GetVersionString() string {
	if x != nil && x.VersionString != nil {
		return *x.VersionString
	}
	return ""
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Request) GetLogContents() string {
	if x != nil && x.LogContents != nil {
		return *x.LogContents
	}
	return ""
}

type CHelpRequestLogs_UploadUserApplicationLog_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Response) Reset() {
	*x = CHelpRequestLogs_UploadUserApplicationLog_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_helprequest_steamworkssdk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CHelpRequestLogs_UploadUserApplicationLog_Response) ProtoMessage() {}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Response) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_helprequest_steamworkssdk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CHelpRequestLogs_UploadUserApplicationLog_Response.ProtoReflect.Descriptor instead.
func (*CHelpRequestLogs_UploadUserApplicationLog_Response) Descriptor() ([]byte, []int) {
	return file_steammessages_helprequest_steamworkssdk_proto_rawDescGZIP(), []int{1}
}

func (x *CHelpRequestLogs_UploadUserApplicationLog_Response) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

var File_steammessages_helprequest_steamworkssdk_proto protoreflect.FileDescriptor

var file_steammessages_helprequest_steamworkssdk_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x68, 0x65, 0x6c, 0x70, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x75,
	0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xae, 0x01, 0x0a, 0x31, 0x43, 0x48, 0x65, 0x6c, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x5f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c,
	0x6f, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x44, 0x0a, 0x32, 0x43, 0x48, 0x65, 0x6c, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x73, 0x5f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xee, 0x01, 0x0a, 0x0f, 0x48, 0x65, 0x6c, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0xa8, 0x01, 0x0a, 0x18, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x32, 0x2e, 0x43, 0x48, 0x65, 0x6c, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x5f, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x43, 0x48,
	0x65, 0x6c, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x5f, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x23, 0x82, 0xb5, 0x18, 0x1f, 0x55, 0x73, 0x65, 0x72, 0x20, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x6c, 0x6f, 0x67, 0x73, 0x1a, 0x30, 0x82, 0xb5, 0x18, 0x2c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x64, 0x65, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x20, 0x77,
	0x69, 0x74, 0x68, 0x20, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x64, 0x20, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x03, 0x80, 0x01, 0x01,
}

var (
	file_steammessages_helprequest_steamworkssdk_proto_rawDescOnce sync.Once
	file_steammessages_helprequest_steamworkssdk_proto_rawDescData = file_steammessages_helprequest_steamworkssdk_proto_rawDesc
)

func file_steammessages_helprequest_steamworkssdk_proto_rawDescGZIP() []byte {
	file_steammessages_helprequest_steamworkssdk_proto_rawDescOnce.Do(func() {
		file_steammessages_helprequest_steamworkssdk_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_helprequest_steamworkssdk_proto_rawDescData)
	})
	return file_steammessages_helprequest_steamworkssdk_proto_rawDescData
}

var file_steammessages_helprequest_steamworkssdk_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_steammessages_helprequest_steamworkssdk_proto_goTypes = []any{
	(*CHelpRequestLogs_UploadUserApplicationLog_Request)(nil),  // 0: CHelpRequestLogs_UploadUserApplicationLog_Request
	(*CHelpRequestLogs_UploadUserApplicationLog_Response)(nil), // 1: CHelpRequestLogs_UploadUserApplicationLog_Response
}
var file_steammessages_helprequest_steamworkssdk_proto_depIdxs = []int32{
	0, // 0: HelpRequestLogs.UploadUserApplicationLog:input_type -> CHelpRequestLogs_UploadUserApplicationLog_Request
	1, // 1: HelpRequestLogs.UploadUserApplicationLog:output_type -> CHelpRequestLogs_UploadUserApplicationLog_Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_steammessages_helprequest_steamworkssdk_proto_init() }
func file_steammessages_helprequest_steamworkssdk_proto_init() {
	if File_steammessages_helprequest_steamworkssdk_proto != nil {
		return
	}
	file_steammessages_unified_base_steamworkssdk_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_helprequest_steamworkssdk_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CHelpRequestLogs_UploadUserApplicationLog_Request); i {
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
		file_steammessages_helprequest_steamworkssdk_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CHelpRequestLogs_UploadUserApplicationLog_Response); i {
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
			RawDescriptor: file_steammessages_helprequest_steamworkssdk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_steammessages_helprequest_steamworkssdk_proto_goTypes,
		DependencyIndexes: file_steammessages_helprequest_steamworkssdk_proto_depIdxs,
		MessageInfos:      file_steammessages_helprequest_steamworkssdk_proto_msgTypes,
	}.Build()
	File_steammessages_helprequest_steamworkssdk_proto = out.File
	file_steammessages_helprequest_steamworkssdk_proto_rawDesc = nil
	file_steammessages_helprequest_steamworkssdk_proto_goTypes = nil
	file_steammessages_helprequest_steamworkssdk_proto_depIdxs = nil
}
