// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: service_mobiledevice.proto

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

type CMobileDevice_DeregisterMobileDevice_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deviceid *string `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
}

func (x *CMobileDevice_DeregisterMobileDevice_Notification) Reset() {
	*x = CMobileDevice_DeregisterMobileDevice_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mobiledevice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMobileDevice_DeregisterMobileDevice_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMobileDevice_DeregisterMobileDevice_Notification) ProtoMessage() {}

func (x *CMobileDevice_DeregisterMobileDevice_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_mobiledevice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMobileDevice_DeregisterMobileDevice_Notification.ProtoReflect.Descriptor instead.
func (*CMobileDevice_DeregisterMobileDevice_Notification) Descriptor() ([]byte, []int) {
	return file_service_mobiledevice_proto_rawDescGZIP(), []int{0}
}

func (x *CMobileDevice_DeregisterMobileDevice_Notification) GetDeviceid() string {
	if x != nil && x.Deviceid != nil {
		return *x.Deviceid
	}
	return ""
}

type CMobileDevice_HasMobileDevice_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppType         *int32  `protobuf:"varint,1,opt,name=app_type,json=appType" json:"app_type,omitempty"`
	PushEnabledOnly *bool   `protobuf:"varint,2,opt,name=push_enabled_only,json=pushEnabledOnly" json:"push_enabled_only,omitempty"`
	MinimumVersion  *string `protobuf:"bytes,3,opt,name=minimum_version,json=minimumVersion" json:"minimum_version,omitempty"`
}

func (x *CMobileDevice_HasMobileDevice_Request) Reset() {
	*x = CMobileDevice_HasMobileDevice_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mobiledevice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMobileDevice_HasMobileDevice_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMobileDevice_HasMobileDevice_Request) ProtoMessage() {}

func (x *CMobileDevice_HasMobileDevice_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_mobiledevice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMobileDevice_HasMobileDevice_Request.ProtoReflect.Descriptor instead.
func (*CMobileDevice_HasMobileDevice_Request) Descriptor() ([]byte, []int) {
	return file_service_mobiledevice_proto_rawDescGZIP(), []int{1}
}

func (x *CMobileDevice_HasMobileDevice_Request) GetAppType() int32 {
	if x != nil && x.AppType != nil {
		return *x.AppType
	}
	return 0
}

func (x *CMobileDevice_HasMobileDevice_Request) GetPushEnabledOnly() bool {
	if x != nil && x.PushEnabledOnly != nil {
		return *x.PushEnabledOnly
	}
	return false
}

func (x *CMobileDevice_HasMobileDevice_Request) GetMinimumVersion() string {
	if x != nil && x.MinimumVersion != nil {
		return *x.MinimumVersion
	}
	return ""
}

type CMobileDevice_HasMobileDevice_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoundDevice *bool `protobuf:"varint,1,opt,name=found_device,json=foundDevice" json:"found_device,omitempty"`
	UpToDate    *bool `protobuf:"varint,2,opt,name=up_to_date,json=upToDate" json:"up_to_date,omitempty"`
}

func (x *CMobileDevice_HasMobileDevice_Response) Reset() {
	*x = CMobileDevice_HasMobileDevice_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mobiledevice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMobileDevice_HasMobileDevice_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMobileDevice_HasMobileDevice_Response) ProtoMessage() {}

func (x *CMobileDevice_HasMobileDevice_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_mobiledevice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMobileDevice_HasMobileDevice_Response.ProtoReflect.Descriptor instead.
func (*CMobileDevice_HasMobileDevice_Response) Descriptor() ([]byte, []int) {
	return file_service_mobiledevice_proto_rawDescGZIP(), []int{2}
}

func (x *CMobileDevice_HasMobileDevice_Response) GetFoundDevice() bool {
	if x != nil && x.FoundDevice != nil {
		return *x.FoundDevice
	}
	return false
}

func (x *CMobileDevice_HasMobileDevice_Response) GetUpToDate() bool {
	if x != nil && x.UpToDate != nil {
		return *x.UpToDate
	}
	return false
}

type CMobileDevice_RegisterMobileDevice_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deviceid                  *string `protobuf:"bytes,1,opt,name=deviceid" json:"deviceid,omitempty"`
	Language                  *string `protobuf:"bytes,2,opt,name=language" json:"language,omitempty"`
	PushEnabled               *bool   `protobuf:"varint,3,opt,name=push_enabled,json=pushEnabled" json:"push_enabled,omitempty"`
	AppVersion                *string `protobuf:"bytes,4,opt,name=app_version,json=appVersion" json:"app_version,omitempty"`
	OsVersion                 *string `protobuf:"bytes,5,opt,name=os_version,json=osVersion" json:"os_version,omitempty"`
	DeviceModel               *string `protobuf:"bytes,6,opt,name=device_model,json=deviceModel" json:"device_model,omitempty"`
	TwofactorDeviceIdentifier *string `protobuf:"bytes,7,opt,name=twofactor_device_identifier,json=twofactorDeviceIdentifier" json:"twofactor_device_identifier,omitempty"`
	MobileApp                 *int32  `protobuf:"varint,8,opt,name=mobile_app,json=mobileApp" json:"mobile_app,omitempty"`
}

func (x *CMobileDevice_RegisterMobileDevice_Request) Reset() {
	*x = CMobileDevice_RegisterMobileDevice_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mobiledevice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMobileDevice_RegisterMobileDevice_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMobileDevice_RegisterMobileDevice_Request) ProtoMessage() {}

func (x *CMobileDevice_RegisterMobileDevice_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_mobiledevice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMobileDevice_RegisterMobileDevice_Request.ProtoReflect.Descriptor instead.
func (*CMobileDevice_RegisterMobileDevice_Request) Descriptor() ([]byte, []int) {
	return file_service_mobiledevice_proto_rawDescGZIP(), []int{3}
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetDeviceid() string {
	if x != nil && x.Deviceid != nil {
		return *x.Deviceid
	}
	return ""
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetPushEnabled() bool {
	if x != nil && x.PushEnabled != nil {
		return *x.PushEnabled
	}
	return false
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetAppVersion() string {
	if x != nil && x.AppVersion != nil {
		return *x.AppVersion
	}
	return ""
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetOsVersion() string {
	if x != nil && x.OsVersion != nil {
		return *x.OsVersion
	}
	return ""
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetDeviceModel() string {
	if x != nil && x.DeviceModel != nil {
		return *x.DeviceModel
	}
	return ""
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetTwofactorDeviceIdentifier() string {
	if x != nil && x.TwofactorDeviceIdentifier != nil {
		return *x.TwofactorDeviceIdentifier
	}
	return ""
}

func (x *CMobileDevice_RegisterMobileDevice_Request) GetMobileApp() int32 {
	if x != nil && x.MobileApp != nil {
		return *x.MobileApp
	}
	return 0
}

type CMobileDevice_RegisterMobileDevice_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UniqueDeviceid *uint32 `protobuf:"varint,2,opt,name=unique_deviceid,json=uniqueDeviceid" json:"unique_deviceid,omitempty"`
}

func (x *CMobileDevice_RegisterMobileDevice_Response) Reset() {
	*x = CMobileDevice_RegisterMobileDevice_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_mobiledevice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMobileDevice_RegisterMobileDevice_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMobileDevice_RegisterMobileDevice_Response) ProtoMessage() {}

func (x *CMobileDevice_RegisterMobileDevice_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_mobiledevice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMobileDevice_RegisterMobileDevice_Response.ProtoReflect.Descriptor instead.
func (*CMobileDevice_RegisterMobileDevice_Response) Descriptor() ([]byte, []int) {
	return file_service_mobiledevice_proto_rawDescGZIP(), []int{4}
}

func (x *CMobileDevice_RegisterMobileDevice_Response) GetUniqueDeviceid() uint32 {
	if x != nil && x.UniqueDeviceid != nil {
		return *x.UniqueDeviceid
	}
	return 0
}

var File_service_mobiledevice_proto protoreflect.FileDescriptor

var file_service_mobiledevice_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4f, 0x0a, 0x31, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64,
	0x22, 0xa1, 0x01, 0x0a, 0x25, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x61, 0x70,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82, 0xb5,
	0x18, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x52, 0x07, 0x61, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f,
	0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x75, 0x73, 0x68,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x6d,
	0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x26, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x75, 0x70, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x70, 0x54, 0x6f, 0x44, 0x61, 0x74, 0x65, 0x22,
	0xd3, 0x02, 0x0a, 0x2a, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x75,
	0x73, 0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x73,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3e, 0x0a, 0x1b,
	0x74, 0x77, 0x6f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x19, 0x74, 0x77, 0x6f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0a,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x08, 0x82, 0xb5, 0x18, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x41, 0x70, 0x70, 0x22, 0x56, 0x0a, 0x2b, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x64, 0x32, 0xc0, 0x02,
	0x0a, 0x0c, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59,
	0x0a, 0x16, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x2e, 0x43, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x44, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x4e,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0f, 0x48, 0x61, 0x73,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x43,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x48, 0x61, 0x73,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a,
	0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x43, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
}

var (
	file_service_mobiledevice_proto_rawDescOnce sync.Once
	file_service_mobiledevice_proto_rawDescData = file_service_mobiledevice_proto_rawDesc
)

func file_service_mobiledevice_proto_rawDescGZIP() []byte {
	file_service_mobiledevice_proto_rawDescOnce.Do(func() {
		file_service_mobiledevice_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_mobiledevice_proto_rawDescData)
	})
	return file_service_mobiledevice_proto_rawDescData
}

var file_service_mobiledevice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_mobiledevice_proto_goTypes = []any{
	(*CMobileDevice_DeregisterMobileDevice_Notification)(nil), // 0: CMobileDevice_DeregisterMobileDevice_Notification
	(*CMobileDevice_HasMobileDevice_Request)(nil),             // 1: CMobileDevice_HasMobileDevice_Request
	(*CMobileDevice_HasMobileDevice_Response)(nil),            // 2: CMobileDevice_HasMobileDevice_Response
	(*CMobileDevice_RegisterMobileDevice_Request)(nil),        // 3: CMobileDevice_RegisterMobileDevice_Request
	(*CMobileDevice_RegisterMobileDevice_Response)(nil),       // 4: CMobileDevice_RegisterMobileDevice_Response
	(*NoResponse)(nil), // 5: NoResponse
}
var file_service_mobiledevice_proto_depIdxs = []int32{
	0, // 0: MobileDevice.DeregisterMobileDevice:input_type -> CMobileDevice_DeregisterMobileDevice_Notification
	1, // 1: MobileDevice.HasMobileDevice:input_type -> CMobileDevice_HasMobileDevice_Request
	3, // 2: MobileDevice.RegisterMobileDevice:input_type -> CMobileDevice_RegisterMobileDevice_Request
	5, // 3: MobileDevice.DeregisterMobileDevice:output_type -> NoResponse
	2, // 4: MobileDevice.HasMobileDevice:output_type -> CMobileDevice_HasMobileDevice_Response
	4, // 5: MobileDevice.RegisterMobileDevice:output_type -> CMobileDevice_RegisterMobileDevice_Response
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_mobiledevice_proto_init() }
func file_service_mobiledevice_proto_init() {
	if File_service_mobiledevice_proto != nil {
		return
	}
	file_common_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_mobiledevice_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMobileDevice_DeregisterMobileDevice_Notification); i {
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
		file_service_mobiledevice_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMobileDevice_HasMobileDevice_Request); i {
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
		file_service_mobiledevice_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CMobileDevice_HasMobileDevice_Response); i {
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
		file_service_mobiledevice_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CMobileDevice_RegisterMobileDevice_Request); i {
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
		file_service_mobiledevice_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CMobileDevice_RegisterMobileDevice_Response); i {
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
			RawDescriptor: file_service_mobiledevice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_mobiledevice_proto_goTypes,
		DependencyIndexes: file_service_mobiledevice_proto_depIdxs,
		MessageInfos:      file_service_mobiledevice_proto_msgTypes,
	}.Build()
	File_service_mobiledevice_proto = out.File
	file_service_mobiledevice_proto_rawDesc = nil
	file_service_mobiledevice_proto_goTypes = nil
	file_service_mobiledevice_proto_depIdxs = nil
}