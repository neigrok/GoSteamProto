// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: service_steamvrvoicechat.proto

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

type CSteamVR_Vector3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X *float32 `protobuf:"fixed32,1,opt,name=x" json:"x,omitempty"`
	Y *float32 `protobuf:"fixed32,2,opt,name=y" json:"y,omitempty"`
	Z *float32 `protobuf:"fixed32,3,opt,name=z" json:"z,omitempty"`
}

func (x *CSteamVR_Vector3) Reset() {
	*x = CSteamVR_Vector3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_Vector3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_Vector3) ProtoMessage() {}

func (x *CSteamVR_Vector3) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_Vector3.ProtoReflect.Descriptor instead.
func (*CSteamVR_Vector3) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{0}
}

func (x *CSteamVR_Vector3) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *CSteamVR_Vector3) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *CSteamVR_Vector3) GetZ() float32 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

type CSteamVR_VoiceChat_Active_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CSteamVR_VoiceChat_Active_Notification) Reset() {
	*x = CSteamVR_VoiceChat_Active_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_Active_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_Active_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_Active_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_Active_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_Active_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{1}
}

type CSteamVR_VoiceChat_GroupName_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (x *CSteamVR_VoiceChat_GroupName_Notification) Reset() {
	*x = CSteamVR_VoiceChat_GroupName_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_GroupName_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_GroupName_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_GroupName_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_GroupName_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_GroupName_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{2}
}

func (x *CSteamVR_VoiceChat_GroupName_Notification) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type CSteamVR_VoiceChat_Inactive_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CSteamVR_VoiceChat_Inactive_Notification) Reset() {
	*x = CSteamVR_VoiceChat_Inactive_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_Inactive_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_Inactive_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_Inactive_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_Inactive_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_Inactive_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{3}
}

type CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatGroupId     *uint64 `protobuf:"varint,1,opt,name=chat_group_id,json=chatGroupId" json:"chat_group_id,omitempty"`
	ChatRoomId      *uint64 `protobuf:"varint,2,opt,name=chat_room_id,json=chatRoomId" json:"chat_room_id,omitempty"`
	SenderAccountid *uint32 `protobuf:"varint,3,opt,name=sender_accountid,json=senderAccountid" json:"sender_accountid,omitempty"`
	Timestamp       *uint32 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Ordinal         *uint32 `protobuf:"varint,5,opt,name=ordinal" json:"ordinal,omitempty"`
	Message         *string `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) Reset() {
	*x = CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{4}
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) GetChatGroupId() uint64 {
	if x != nil && x.ChatGroupId != nil {
		return *x.ChatGroupId
	}
	return 0
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) GetChatRoomId() uint64 {
	if x != nil && x.ChatRoomId != nil {
		return *x.ChatRoomId
	}
	return 0
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) GetSenderAccountid() uint32 {
	if x != nil && x.SenderAccountid != nil {
		return *x.SenderAccountid
	}
	return 0
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) GetOrdinal() uint32 {
	if x != nil && x.Ordinal != nil {
		return *x.Ordinal
	}
	return 0
}

func (x *CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type CSteamVR_VoiceChat_PerUserGainValue_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accountid *uint32  `protobuf:"varint,1,opt,name=accountid" json:"accountid,omitempty"`
	Muted     *bool    `protobuf:"varint,2,opt,name=muted" json:"muted,omitempty"`
	Gain      *float32 `protobuf:"fixed32,3,opt,name=gain" json:"gain,omitempty"`
}

func (x *CSteamVR_VoiceChat_PerUserGainValue_Notification) Reset() {
	*x = CSteamVR_VoiceChat_PerUserGainValue_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_PerUserGainValue_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_PerUserGainValue_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_PerUserGainValue_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_PerUserGainValue_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_PerUserGainValue_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{5}
}

func (x *CSteamVR_VoiceChat_PerUserGainValue_Notification) GetAccountid() uint32 {
	if x != nil && x.Accountid != nil {
		return *x.Accountid
	}
	return 0
}

func (x *CSteamVR_VoiceChat_PerUserGainValue_Notification) GetMuted() bool {
	if x != nil && x.Muted != nil {
		return *x.Muted
	}
	return false
}

func (x *CSteamVR_VoiceChat_PerUserGainValue_Notification) GetGain() float32 {
	if x != nil && x.Gain != nil {
		return *x.Gain
	}
	return 0
}

type CSteamVR_VoiceChat_PerUserVoiceStatus_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accountid          *uint32 `protobuf:"varint,1,opt,name=accountid" json:"accountid,omitempty"`
	MicMutedLocally    *bool   `protobuf:"varint,2,opt,name=mic_muted_locally,json=micMutedLocally" json:"mic_muted_locally,omitempty"`
	OutputMutedLocally *bool   `protobuf:"varint,3,opt,name=output_muted_locally,json=outputMutedLocally" json:"output_muted_locally,omitempty"`
}

func (x *CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) Reset() {
	*x = CSteamVR_VoiceChat_PerUserVoiceStatus_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_PerUserVoiceStatus_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{6}
}

func (x *CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) GetAccountid() uint32 {
	if x != nil && x.Accountid != nil {
		return *x.Accountid
	}
	return 0
}

func (x *CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) GetMicMutedLocally() bool {
	if x != nil && x.MicMutedLocally != nil {
		return *x.MicMutedLocally
	}
	return false
}

func (x *CSteamVR_VoiceChat_PerUserVoiceStatus_Notification) GetOutputMutedLocally() bool {
	if x != nil && x.OutputMutedLocally != nil {
		return *x.OutputMutedLocally
	}
	return false
}

type CSteamVR_VoiceChat_SetDefaultSession_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatGroupId *uint64 `protobuf:"varint,1,opt,name=chat_group_id,json=chatGroupId" json:"chat_group_id,omitempty"`
	ChatRoomId  *uint64 `protobuf:"varint,2,opt,name=chat_room_id,json=chatRoomId" json:"chat_room_id,omitempty"`
}

func (x *CSteamVR_VoiceChat_SetDefaultSession_Notification) Reset() {
	*x = CSteamVR_VoiceChat_SetDefaultSession_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_SetDefaultSession_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_SetDefaultSession_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_SetDefaultSession_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_SetDefaultSession_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_SetDefaultSession_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{7}
}

func (x *CSteamVR_VoiceChat_SetDefaultSession_Notification) GetChatGroupId() uint64 {
	if x != nil && x.ChatGroupId != nil {
		return *x.ChatGroupId
	}
	return 0
}

func (x *CSteamVR_VoiceChat_SetDefaultSession_Notification) GetChatRoomId() uint64 {
	if x != nil && x.ChatRoomId != nil {
		return *x.ChatRoomId
	}
	return 0
}

type CSteamVR_VoiceChat_SetSpatialAudioListener_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position *CSteamVR_Vector3 `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	Forward  *CSteamVR_Vector3 `protobuf:"bytes,2,opt,name=forward" json:"forward,omitempty"`
	Up       *CSteamVR_Vector3 `protobuf:"bytes,3,opt,name=up" json:"up,omitempty"`
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) Reset() {
	*x = CSteamVR_VoiceChat_SetSpatialAudioListener_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_SetSpatialAudioListener_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{8}
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) GetPosition() *CSteamVR_Vector3 {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) GetForward() *CSteamVR_Vector3 {
	if x != nil {
		return x.Forward
	}
	return nil
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioListener_Notification) GetUp() *CSteamVR_Vector3 {
	if x != nil {
		return x.Up
	}
	return nil
}

type CSteamVR_VoiceChat_SetSpatialAudioSource_Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steamid  *uint64           `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	Position *CSteamVR_Vector3 `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioSource_Notification) Reset() {
	*x = CSteamVR_VoiceChat_SetSpatialAudioSource_Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_steamvrvoicechat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioSource_Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CSteamVR_VoiceChat_SetSpatialAudioSource_Notification) ProtoMessage() {}

func (x *CSteamVR_VoiceChat_SetSpatialAudioSource_Notification) ProtoReflect() protoreflect.Message {
	mi := &file_service_steamvrvoicechat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CSteamVR_VoiceChat_SetSpatialAudioSource_Notification.ProtoReflect.Descriptor instead.
func (*CSteamVR_VoiceChat_SetSpatialAudioSource_Notification) Descriptor() ([]byte, []int) {
	return file_service_steamvrvoicechat_proto_rawDescGZIP(), []int{9}
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioSource_Notification) GetSteamid() uint64 {
	if x != nil && x.Steamid != nil {
		return *x.Steamid
	}
	return 0
}

func (x *CSteamVR_VoiceChat_SetSpatialAudioSource_Notification) GetPosition() *CSteamVR_Vector3 {
	if x != nil {
		return x.Position
	}
	return nil
}

var File_service_steamvrvoicechat_proto protoreflect.FileDescriptor

var file_service_steamvrvoicechat_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x76,
	0x72, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x10, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x7a, 0x22, 0x28, 0x0a, 0x26, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x29, 0x43,
	0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x5f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x28,
	0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x5f, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf9, 0x01, 0x0a, 0x34, 0x43, 0x53, 0x74,
	0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f,
	0x4e, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x7a, 0x0a, 0x30, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52,
	0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x50, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x75, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x67, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x67, 0x61, 0x69, 0x6e,
	0x22, 0xb0, 0x01, 0x0a, 0x32, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x50, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x75, 0x74,
	0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x6d, 0x69, 0x63, 0x4d, 0x75, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x6c,
	0x79, 0x12, 0x30, 0x0a, 0x14, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x6d, 0x75, 0x74, 0x65,
	0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x4d, 0x75, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x6c, 0x79, 0x22, 0x79, 0x0a, 0x31, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x53, 0x65, 0x74, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x63, 0x68, 0x61, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x63, 0x68, 0x61, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0xb8,
	0x01, 0x0a, 0x37, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x53, 0x65, 0x74, 0x53, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c,
	0x41, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43,
	0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x07, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x53, 0x74,
	0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x07, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x02, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x33, 0x52, 0x02, 0x75, 0x70, 0x22, 0x80, 0x01, 0x0a, 0x35, 0x43, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x5f, 0x53, 0x65, 0x74, 0x53, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x12, 0x2d, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x33, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xfa, 0x05, 0x0a,
	0x10, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x12, 0x3e, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x27, 0x2e, 0x43, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x5f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x5f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x4e, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x49, 0x6e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x29, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b,
	0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x14, 0x4e,
	0x65, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x41, 0x64,
	0x64, 0x65, 0x64, 0x12, 0x35, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56,
	0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x4e, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x41, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x4e, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x31, 0x2e, 0x43, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x5f, 0x50, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x69, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b,
	0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x12, 0x50,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x33, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x50, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61,
	0x6d, 0x56, 0x52, 0x5f, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x53, 0x65,
	0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x4e,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x17, 0x53, 0x65, 0x74,
	0x53, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x53, 0x65, 0x74, 0x53, 0x70, 0x61,
	0x74, 0x69, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b,
	0x2e, 0x4e, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x53,
	0x65, 0x74, 0x53, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x36, 0x2e, 0x43, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x56, 0x52, 0x5f,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x74, 0x5f, 0x53, 0x65, 0x74, 0x53, 0x70, 0x61,
	0x74, 0x69, 0x61, 0x6c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0b, 0x2e, 0x4e,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
}

var (
	file_service_steamvrvoicechat_proto_rawDescOnce sync.Once
	file_service_steamvrvoicechat_proto_rawDescData = file_service_steamvrvoicechat_proto_rawDesc
)

func file_service_steamvrvoicechat_proto_rawDescGZIP() []byte {
	file_service_steamvrvoicechat_proto_rawDescOnce.Do(func() {
		file_service_steamvrvoicechat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_steamvrvoicechat_proto_rawDescData)
	})
	return file_service_steamvrvoicechat_proto_rawDescData
}

var file_service_steamvrvoicechat_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_steamvrvoicechat_proto_goTypes = []any{
	(*CSteamVR_Vector3)(nil),                                        // 0: CSteamVR_Vector3
	(*CSteamVR_VoiceChat_Active_Notification)(nil),                  // 1: CSteamVR_VoiceChat_Active_Notification
	(*CSteamVR_VoiceChat_GroupName_Notification)(nil),               // 2: CSteamVR_VoiceChat_GroupName_Notification
	(*CSteamVR_VoiceChat_Inactive_Notification)(nil),                // 3: CSteamVR_VoiceChat_Inactive_Notification
	(*CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification)(nil),    // 4: CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification
	(*CSteamVR_VoiceChat_PerUserGainValue_Notification)(nil),        // 5: CSteamVR_VoiceChat_PerUserGainValue_Notification
	(*CSteamVR_VoiceChat_PerUserVoiceStatus_Notification)(nil),      // 6: CSteamVR_VoiceChat_PerUserVoiceStatus_Notification
	(*CSteamVR_VoiceChat_SetDefaultSession_Notification)(nil),       // 7: CSteamVR_VoiceChat_SetDefaultSession_Notification
	(*CSteamVR_VoiceChat_SetSpatialAudioListener_Notification)(nil), // 8: CSteamVR_VoiceChat_SetSpatialAudioListener_Notification
	(*CSteamVR_VoiceChat_SetSpatialAudioSource_Notification)(nil),   // 9: CSteamVR_VoiceChat_SetSpatialAudioSource_Notification
	(*NoResponse)(nil), // 10: NoResponse
}
var file_service_steamvrvoicechat_proto_depIdxs = []int32{
	0,  // 0: CSteamVR_VoiceChat_SetSpatialAudioListener_Notification.position:type_name -> CSteamVR_Vector3
	0,  // 1: CSteamVR_VoiceChat_SetSpatialAudioListener_Notification.forward:type_name -> CSteamVR_Vector3
	0,  // 2: CSteamVR_VoiceChat_SetSpatialAudioListener_Notification.up:type_name -> CSteamVR_Vector3
	0,  // 3: CSteamVR_VoiceChat_SetSpatialAudioSource_Notification.position:type_name -> CSteamVR_Vector3
	1,  // 4: SteamVRVoiceChat.Active:input_type -> CSteamVR_VoiceChat_Active_Notification
	2,  // 5: SteamVRVoiceChat.GroupName:input_type -> CSteamVR_VoiceChat_GroupName_Notification
	3,  // 6: SteamVRVoiceChat.Inactive:input_type -> CSteamVR_VoiceChat_Inactive_Notification
	4,  // 7: SteamVRVoiceChat.NewGroupChatMsgAdded:input_type -> CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification
	5,  // 8: SteamVRVoiceChat.PerUserGainValue:input_type -> CSteamVR_VoiceChat_PerUserGainValue_Notification
	6,  // 9: SteamVRVoiceChat.PerUserVoiceStatus:input_type -> CSteamVR_VoiceChat_PerUserVoiceStatus_Notification
	7,  // 10: SteamVRVoiceChat.SetDefaultSession:input_type -> CSteamVR_VoiceChat_SetDefaultSession_Notification
	8,  // 11: SteamVRVoiceChat.SetSpatialAudioListener:input_type -> CSteamVR_VoiceChat_SetSpatialAudioListener_Notification
	9,  // 12: SteamVRVoiceChat.SetSpatialAudioSource:input_type -> CSteamVR_VoiceChat_SetSpatialAudioSource_Notification
	10, // 13: SteamVRVoiceChat.Active:output_type -> NoResponse
	10, // 14: SteamVRVoiceChat.GroupName:output_type -> NoResponse
	10, // 15: SteamVRVoiceChat.Inactive:output_type -> NoResponse
	10, // 16: SteamVRVoiceChat.NewGroupChatMsgAdded:output_type -> NoResponse
	10, // 17: SteamVRVoiceChat.PerUserGainValue:output_type -> NoResponse
	10, // 18: SteamVRVoiceChat.PerUserVoiceStatus:output_type -> NoResponse
	10, // 19: SteamVRVoiceChat.SetDefaultSession:output_type -> NoResponse
	10, // 20: SteamVRVoiceChat.SetSpatialAudioListener:output_type -> NoResponse
	10, // 21: SteamVRVoiceChat.SetSpatialAudioSource:output_type -> NoResponse
	13, // [13:22] is the sub-list for method output_type
	4,  // [4:13] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_service_steamvrvoicechat_proto_init() }
func file_service_steamvrvoicechat_proto_init() {
	if File_service_steamvrvoicechat_proto != nil {
		return
	}
	file_common_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_steamvrvoicechat_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_Vector3); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_Active_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_GroupName_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_Inactive_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_NewGroupChatMsgAdded_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_PerUserGainValue_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_PerUserVoiceStatus_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_SetDefaultSession_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_SetSpatialAudioListener_Notification); i {
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
		file_service_steamvrvoicechat_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*CSteamVR_VoiceChat_SetSpatialAudioSource_Notification); i {
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
			RawDescriptor: file_service_steamvrvoicechat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_steamvrvoicechat_proto_goTypes,
		DependencyIndexes: file_service_steamvrvoicechat_proto_depIdxs,
		MessageInfos:      file_service_steamvrvoicechat_proto_msgTypes,
	}.Build()
	File_service_steamvrvoicechat_proto = out.File
	file_service_steamvrvoicechat_proto_rawDesc = nil
	file_service_steamvrvoicechat_proto_goTypes = nil
	file_service_steamvrvoicechat_proto_depIdxs = nil
}