// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: service_phone.proto

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

type CPhone_AddPhoneToAccount_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success         *bool  `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	PhoneNumberType *int32 `protobuf:"varint,2,opt,name=phone_number_type,json=phoneNumberType" json:"phone_number_type,omitempty"`
}

func (x *CPhone_AddPhoneToAccount_Response) Reset() {
	*x = CPhone_AddPhoneToAccount_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_AddPhoneToAccount_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_AddPhoneToAccount_Response) ProtoMessage() {}

func (x *CPhone_AddPhoneToAccount_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_AddPhoneToAccount_Response.ProtoReflect.Descriptor instead.
func (*CPhone_AddPhoneToAccount_Response) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{0}
}

func (x *CPhone_AddPhoneToAccount_Response) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *CPhone_AddPhoneToAccount_Response) GetPhoneNumberType() int32 {
	if x != nil && x.PhoneNumberType != nil {
		return *x.PhoneNumberType
	}
	return 0
}

type CPhone_ConfirmAddPhoneToAccount_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steamid *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	Stoken  *string `protobuf:"bytes,2,opt,name=stoken" json:"stoken,omitempty"`
}

func (x *CPhone_ConfirmAddPhoneToAccount_Request) Reset() {
	*x = CPhone_ConfirmAddPhoneToAccount_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_ConfirmAddPhoneToAccount_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_ConfirmAddPhoneToAccount_Request) ProtoMessage() {}

func (x *CPhone_ConfirmAddPhoneToAccount_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_ConfirmAddPhoneToAccount_Request.ProtoReflect.Descriptor instead.
func (*CPhone_ConfirmAddPhoneToAccount_Request) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{1}
}

func (x *CPhone_ConfirmAddPhoneToAccount_Request) GetSteamid() uint64 {
	if x != nil && x.Steamid != nil {
		return *x.Steamid
	}
	return 0
}

func (x *CPhone_ConfirmAddPhoneToAccount_Request) GetStoken() string {
	if x != nil && x.Stoken != nil {
		return *x.Stoken
	}
	return ""
}

type CPhone_IsAccountWaitingForEmailConfirmation_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Request) Reset() {
	*x = CPhone_IsAccountWaitingForEmailConfirmation_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_IsAccountWaitingForEmailConfirmation_Request) ProtoMessage() {}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_IsAccountWaitingForEmailConfirmation_Request.ProtoReflect.Descriptor instead.
func (*CPhone_IsAccountWaitingForEmailConfirmation_Request) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{2}
}

type CPhone_IsAccountWaitingForEmailConfirmation_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AwaitingEmailConfirmation *bool   `protobuf:"varint,1,opt,name=awaiting_email_confirmation,json=awaitingEmailConfirmation" json:"awaiting_email_confirmation,omitempty"`
	SecondsToWait             *uint32 `protobuf:"varint,2,opt,name=seconds_to_wait,json=secondsToWait" json:"seconds_to_wait,omitempty"`
}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Response) Reset() {
	*x = CPhone_IsAccountWaitingForEmailConfirmation_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_IsAccountWaitingForEmailConfirmation_Response) ProtoMessage() {}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_IsAccountWaitingForEmailConfirmation_Response.ProtoReflect.Descriptor instead.
func (*CPhone_IsAccountWaitingForEmailConfirmation_Response) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{3}
}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Response) GetAwaitingEmailConfirmation() bool {
	if x != nil && x.AwaitingEmailConfirmation != nil {
		return *x.AwaitingEmailConfirmation
	}
	return false
}

func (x *CPhone_IsAccountWaitingForEmailConfirmation_Response) GetSecondsToWait() uint32 {
	if x != nil && x.SecondsToWait != nil {
		return *x.SecondsToWait
	}
	return 0
}

type CPhone_SendPhoneVerificationCode_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *uint32 `protobuf:"varint,1,opt,name=language" json:"language,omitempty"`
}

func (x *CPhone_SendPhoneVerificationCode_Request) Reset() {
	*x = CPhone_SendPhoneVerificationCode_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_SendPhoneVerificationCode_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_SendPhoneVerificationCode_Request) ProtoMessage() {}

func (x *CPhone_SendPhoneVerificationCode_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_SendPhoneVerificationCode_Request.ProtoReflect.Descriptor instead.
func (*CPhone_SendPhoneVerificationCode_Request) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{4}
}

func (x *CPhone_SendPhoneVerificationCode_Request) GetLanguage() uint32 {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return 0
}

type CPhone_SendPhoneVerificationCode_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CPhone_SendPhoneVerificationCode_Response) Reset() {
	*x = CPhone_SendPhoneVerificationCode_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_SendPhoneVerificationCode_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_SendPhoneVerificationCode_Response) ProtoMessage() {}

func (x *CPhone_SendPhoneVerificationCode_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_SendPhoneVerificationCode_Response.ProtoReflect.Descriptor instead.
func (*CPhone_SendPhoneVerificationCode_Response) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{5}
}

type CPhone_SetAccountPhoneNumber_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNumber      *string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
	PhoneCountryCode *string `protobuf:"bytes,2,opt,name=phone_country_code,json=phoneCountryCode" json:"phone_country_code,omitempty"`
}

func (x *CPhone_SetAccountPhoneNumber_Request) Reset() {
	*x = CPhone_SetAccountPhoneNumber_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_SetAccountPhoneNumber_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_SetAccountPhoneNumber_Request) ProtoMessage() {}

func (x *CPhone_SetAccountPhoneNumber_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_SetAccountPhoneNumber_Request.ProtoReflect.Descriptor instead.
func (*CPhone_SetAccountPhoneNumber_Request) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{6}
}

func (x *CPhone_SetAccountPhoneNumber_Request) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *CPhone_SetAccountPhoneNumber_Request) GetPhoneCountryCode() string {
	if x != nil && x.PhoneCountryCode != nil {
		return *x.PhoneCountryCode
	}
	return ""
}

type CPhone_SetAccountPhoneNumber_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfirmationEmailAddress *string `protobuf:"bytes,1,opt,name=confirmation_email_address,json=confirmationEmailAddress" json:"confirmation_email_address,omitempty"`
	PhoneNumberFormatted     *string `protobuf:"bytes,2,opt,name=phone_number_formatted,json=phoneNumberFormatted" json:"phone_number_formatted,omitempty"`
}

func (x *CPhone_SetAccountPhoneNumber_Response) Reset() {
	*x = CPhone_SetAccountPhoneNumber_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_SetAccountPhoneNumber_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_SetAccountPhoneNumber_Response) ProtoMessage() {}

func (x *CPhone_SetAccountPhoneNumber_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_SetAccountPhoneNumber_Response.ProtoReflect.Descriptor instead.
func (*CPhone_SetAccountPhoneNumber_Response) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{7}
}

func (x *CPhone_SetAccountPhoneNumber_Response) GetConfirmationEmailAddress() string {
	if x != nil && x.ConfirmationEmailAddress != nil {
		return *x.ConfirmationEmailAddress
	}
	return ""
}

func (x *CPhone_SetAccountPhoneNumber_Response) GetPhoneNumberFormatted() string {
	if x != nil && x.PhoneNumberFormatted != nil {
		return *x.PhoneNumberFormatted
	}
	return ""
}

type CPhone_VerifyAccountPhoneWithCode_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code *string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
}

func (x *CPhone_VerifyAccountPhoneWithCode_Request) Reset() {
	*x = CPhone_VerifyAccountPhoneWithCode_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_VerifyAccountPhoneWithCode_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_VerifyAccountPhoneWithCode_Request) ProtoMessage() {}

func (x *CPhone_VerifyAccountPhoneWithCode_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_VerifyAccountPhoneWithCode_Request.ProtoReflect.Descriptor instead.
func (*CPhone_VerifyAccountPhoneWithCode_Request) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{8}
}

func (x *CPhone_VerifyAccountPhoneWithCode_Request) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

type CPhone_VerifyAccountPhoneWithCode_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CPhone_VerifyAccountPhoneWithCode_Response) Reset() {
	*x = CPhone_VerifyAccountPhoneWithCode_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_phone_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPhone_VerifyAccountPhoneWithCode_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPhone_VerifyAccountPhoneWithCode_Response) ProtoMessage() {}

func (x *CPhone_VerifyAccountPhoneWithCode_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_phone_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPhone_VerifyAccountPhoneWithCode_Response.ProtoReflect.Descriptor instead.
func (*CPhone_VerifyAccountPhoneWithCode_Response) Descriptor() ([]byte, []int) {
	return file_service_phone_proto_rawDescGZIP(), []int{9}
}

var File_service_phone_proto protoreflect.FileDescriptor

var file_service_phone_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x21, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x22, 0x5b, 0x0a, 0x27, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x35, 0x0a,
	0x33, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x49, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x34, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x49, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67,
	0x46, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x1b, 0x61, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x19, 0x61, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x77, 0x61, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x54,
	0x6f, 0x57, 0x61, 0x69, 0x74, 0x22, 0x46, 0x0a, 0x28, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a,
	0x29, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x77, 0x0a, 0x24, 0x43, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x25, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x53,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x64, 0x22, 0x3f, 0x0a, 0x29, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x2c, 0x0a, 0x2a, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xda, 0x04, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x68, 0x0a, 0x18, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54,
	0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x41, 0x64, 0x64, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x24, 0x49, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x2e,
	0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x49, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x49, 0x73, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x57, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x19, 0x53, 0x65,
	0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x53, 0x65, 0x6e, 0x64,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66,
	0x0a, 0x15, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x53, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x1a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x43, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x43, 0x6f, 0x64, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
}

var (
	file_service_phone_proto_rawDescOnce sync.Once
	file_service_phone_proto_rawDescData = file_service_phone_proto_rawDesc
)

func file_service_phone_proto_rawDescGZIP() []byte {
	file_service_phone_proto_rawDescOnce.Do(func() {
		file_service_phone_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_phone_proto_rawDescData)
	})
	return file_service_phone_proto_rawDescData
}

var file_service_phone_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_phone_proto_goTypes = []any{
	(*CPhone_AddPhoneToAccount_Response)(nil),                    // 0: CPhone_AddPhoneToAccount_Response
	(*CPhone_ConfirmAddPhoneToAccount_Request)(nil),              // 1: CPhone_ConfirmAddPhoneToAccount_Request
	(*CPhone_IsAccountWaitingForEmailConfirmation_Request)(nil),  // 2: CPhone_IsAccountWaitingForEmailConfirmation_Request
	(*CPhone_IsAccountWaitingForEmailConfirmation_Response)(nil), // 3: CPhone_IsAccountWaitingForEmailConfirmation_Response
	(*CPhone_SendPhoneVerificationCode_Request)(nil),             // 4: CPhone_SendPhoneVerificationCode_Request
	(*CPhone_SendPhoneVerificationCode_Response)(nil),            // 5: CPhone_SendPhoneVerificationCode_Response
	(*CPhone_SetAccountPhoneNumber_Request)(nil),                 // 6: CPhone_SetAccountPhoneNumber_Request
	(*CPhone_SetAccountPhoneNumber_Response)(nil),                // 7: CPhone_SetAccountPhoneNumber_Response
	(*CPhone_VerifyAccountPhoneWithCode_Request)(nil),            // 8: CPhone_VerifyAccountPhoneWithCode_Request
	(*CPhone_VerifyAccountPhoneWithCode_Response)(nil),           // 9: CPhone_VerifyAccountPhoneWithCode_Response
}
var file_service_phone_proto_depIdxs = []int32{
	1, // 0: Phone.ConfirmAddPhoneToAccount:input_type -> CPhone_ConfirmAddPhoneToAccount_Request
	2, // 1: Phone.IsAccountWaitingForEmailConfirmation:input_type -> CPhone_IsAccountWaitingForEmailConfirmation_Request
	4, // 2: Phone.SendPhoneVerificationCode:input_type -> CPhone_SendPhoneVerificationCode_Request
	6, // 3: Phone.SetAccountPhoneNumber:input_type -> CPhone_SetAccountPhoneNumber_Request
	8, // 4: Phone.VerifyAccountPhoneWithCode:input_type -> CPhone_VerifyAccountPhoneWithCode_Request
	0, // 5: Phone.ConfirmAddPhoneToAccount:output_type -> CPhone_AddPhoneToAccount_Response
	3, // 6: Phone.IsAccountWaitingForEmailConfirmation:output_type -> CPhone_IsAccountWaitingForEmailConfirmation_Response
	5, // 7: Phone.SendPhoneVerificationCode:output_type -> CPhone_SendPhoneVerificationCode_Response
	7, // 8: Phone.SetAccountPhoneNumber:output_type -> CPhone_SetAccountPhoneNumber_Response
	9, // 9: Phone.VerifyAccountPhoneWithCode:output_type -> CPhone_VerifyAccountPhoneWithCode_Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_phone_proto_init() }
func file_service_phone_proto_init() {
	if File_service_phone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_phone_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_AddPhoneToAccount_Response); i {
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
		file_service_phone_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_ConfirmAddPhoneToAccount_Request); i {
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
		file_service_phone_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_IsAccountWaitingForEmailConfirmation_Request); i {
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
		file_service_phone_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_IsAccountWaitingForEmailConfirmation_Response); i {
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
		file_service_phone_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_SendPhoneVerificationCode_Request); i {
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
		file_service_phone_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_SendPhoneVerificationCode_Response); i {
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
		file_service_phone_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_SetAccountPhoneNumber_Request); i {
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
		file_service_phone_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_SetAccountPhoneNumber_Response); i {
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
		file_service_phone_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_VerifyAccountPhoneWithCode_Request); i {
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
		file_service_phone_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CPhone_VerifyAccountPhoneWithCode_Response); i {
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
			RawDescriptor: file_service_phone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_phone_proto_goTypes,
		DependencyIndexes: file_service_phone_proto_depIdxs,
		MessageInfos:      file_service_phone_proto_msgTypes,
	}.Build()
	File_service_phone_proto = out.File
	file_service_phone_proto_rawDesc = nil
	file_service_phone_proto_goTypes = nil
	file_service_phone_proto_depIdxs = nil
}
