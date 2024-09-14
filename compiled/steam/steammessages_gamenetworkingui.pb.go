// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: steammessages_gamenetworkingui.proto

package steam

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

type CGameNetworkingUI_GlobalState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CGameNetworkingUI_GlobalState) Reset() {
	*x = CGameNetworkingUI_GlobalState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_gamenetworkingui_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGameNetworkingUI_GlobalState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGameNetworkingUI_GlobalState) ProtoMessage() {}

func (x *CGameNetworkingUI_GlobalState) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_gamenetworkingui_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGameNetworkingUI_GlobalState.ProtoReflect.Descriptor instead.
func (*CGameNetworkingUI_GlobalState) Descriptor() ([]byte, []int) {
	return file_steammessages_gamenetworkingui_proto_rawDescGZIP(), []int{0}
}

type CGameNetworkingUI_ConnectionState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionKey                       *string                             `protobuf:"bytes,1,opt,name=connection_key,json=connectionKey" json:"connection_key,omitempty"`
	Appid                               *uint32                             `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	ConnectionIdLocal                   *uint32                             `protobuf:"fixed32,3,opt,name=connection_id_local,json=connectionIdLocal" json:"connection_id_local,omitempty"`
	IdentityLocal                       *string                             `protobuf:"bytes,4,opt,name=identity_local,json=identityLocal" json:"identity_local,omitempty"`
	IdentityRemote                      *string                             `protobuf:"bytes,5,opt,name=identity_remote,json=identityRemote" json:"identity_remote,omitempty"`
	ConnectionState                     *uint32                             `protobuf:"varint,10,opt,name=connection_state,json=connectionState" json:"connection_state,omitempty"`
	StartTime                           *uint32                             `protobuf:"varint,12,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	CloseTime                           *uint32                             `protobuf:"varint,13,opt,name=close_time,json=closeTime" json:"close_time,omitempty"`
	CloseReason                         *uint32                             `protobuf:"varint,14,opt,name=close_reason,json=closeReason" json:"close_reason,omitempty"`
	CloseMessage                        *string                             `protobuf:"bytes,15,opt,name=close_message,json=closeMessage" json:"close_message,omitempty"`
	StatusLocToken                      *string                             `protobuf:"bytes,16,opt,name=status_loc_token,json=statusLocToken" json:"status_loc_token,omitempty"`
	TransportKind                       *uint32                             `protobuf:"varint,20,opt,name=transport_kind,json=transportKind" json:"transport_kind,omitempty"`
	SdrpopidLocal                       *string                             `protobuf:"bytes,21,opt,name=sdrpopid_local,json=sdrpopidLocal" json:"sdrpopid_local,omitempty"`
	SdrpopidRemote                      *string                             `protobuf:"bytes,22,opt,name=sdrpopid_remote,json=sdrpopidRemote" json:"sdrpopid_remote,omitempty"`
	AddressRemote                       *string                             `protobuf:"bytes,23,opt,name=address_remote,json=addressRemote" json:"address_remote,omitempty"`
	P2PRouting                          *CMsgSteamDatagramP2PRoutingSummary `protobuf:"bytes,24,opt,name=p2p_routing,json=p2pRouting" json:"p2p_routing,omitempty"`
	PingInterior                        *uint32                             `protobuf:"varint,25,opt,name=ping_interior,json=pingInterior" json:"ping_interior,omitempty"`
	PingRemoteFront                     *uint32                             `protobuf:"varint,26,opt,name=ping_remote_front,json=pingRemoteFront" json:"ping_remote_front,omitempty"`
	PingDefaultInternetRoute            *uint32                             `protobuf:"varint,27,opt,name=ping_default_internet_route,json=pingDefaultInternetRoute" json:"ping_default_internet_route,omitempty"`
	E2EQualityLocal                     *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,30,opt,name=e2e_quality_local,json=e2eQualityLocal" json:"e2e_quality_local,omitempty"`
	E2EQualityRemote                    *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,31,opt,name=e2e_quality_remote,json=e2eQualityRemote" json:"e2e_quality_remote,omitempty"`
	E2EQualityRemoteInstantaneousTime   *uint64                             `protobuf:"varint,32,opt,name=e2e_quality_remote_instantaneous_time,json=e2eQualityRemoteInstantaneousTime" json:"e2e_quality_remote_instantaneous_time,omitempty"`
	E2EQualityRemoteLifetimeTime        *uint64                             `protobuf:"varint,33,opt,name=e2e_quality_remote_lifetime_time,json=e2eQualityRemoteLifetimeTime" json:"e2e_quality_remote_lifetime_time,omitempty"`
	FrontQualityLocal                   *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,40,opt,name=front_quality_local,json=frontQualityLocal" json:"front_quality_local,omitempty"`
	FrontQualityRemote                  *CMsgSteamDatagramConnectionQuality `protobuf:"bytes,41,opt,name=front_quality_remote,json=frontQualityRemote" json:"front_quality_remote,omitempty"`
	FrontQualityRemoteInstantaneousTime *uint64                             `protobuf:"varint,42,opt,name=front_quality_remote_instantaneous_time,json=frontQualityRemoteInstantaneousTime" json:"front_quality_remote_instantaneous_time,omitempty"`
	FrontQualityRemoteLifetimeTime      *uint64                             `protobuf:"varint,43,opt,name=front_quality_remote_lifetime_time,json=frontQualityRemoteLifetimeTime" json:"front_quality_remote_lifetime_time,omitempty"`
}

func (x *CGameNetworkingUI_ConnectionState) Reset() {
	*x = CGameNetworkingUI_ConnectionState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_gamenetworkingui_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGameNetworkingUI_ConnectionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGameNetworkingUI_ConnectionState) ProtoMessage() {}

func (x *CGameNetworkingUI_ConnectionState) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_gamenetworkingui_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGameNetworkingUI_ConnectionState.ProtoReflect.Descriptor instead.
func (*CGameNetworkingUI_ConnectionState) Descriptor() ([]byte, []int) {
	return file_steammessages_gamenetworkingui_proto_rawDescGZIP(), []int{1}
}

func (x *CGameNetworkingUI_ConnectionState) GetConnectionKey() string {
	if x != nil && x.ConnectionKey != nil {
		return *x.ConnectionKey
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetConnectionIdLocal() uint32 {
	if x != nil && x.ConnectionIdLocal != nil {
		return *x.ConnectionIdLocal
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetIdentityLocal() string {
	if x != nil && x.IdentityLocal != nil {
		return *x.IdentityLocal
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetIdentityRemote() string {
	if x != nil && x.IdentityRemote != nil {
		return *x.IdentityRemote
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetConnectionState() uint32 {
	if x != nil && x.ConnectionState != nil {
		return *x.ConnectionState
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetStartTime() uint32 {
	if x != nil && x.StartTime != nil {
		return *x.StartTime
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetCloseTime() uint32 {
	if x != nil && x.CloseTime != nil {
		return *x.CloseTime
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetCloseReason() uint32 {
	if x != nil && x.CloseReason != nil {
		return *x.CloseReason
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetCloseMessage() string {
	if x != nil && x.CloseMessage != nil {
		return *x.CloseMessage
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetStatusLocToken() string {
	if x != nil && x.StatusLocToken != nil {
		return *x.StatusLocToken
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetTransportKind() uint32 {
	if x != nil && x.TransportKind != nil {
		return *x.TransportKind
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetSdrpopidLocal() string {
	if x != nil && x.SdrpopidLocal != nil {
		return *x.SdrpopidLocal
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetSdrpopidRemote() string {
	if x != nil && x.SdrpopidRemote != nil {
		return *x.SdrpopidRemote
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetAddressRemote() string {
	if x != nil && x.AddressRemote != nil {
		return *x.AddressRemote
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionState) GetP2PRouting() *CMsgSteamDatagramP2PRoutingSummary {
	if x != nil {
		return x.P2PRouting
	}
	return nil
}

func (x *CGameNetworkingUI_ConnectionState) GetPingInterior() uint32 {
	if x != nil && x.PingInterior != nil {
		return *x.PingInterior
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetPingRemoteFront() uint32 {
	if x != nil && x.PingRemoteFront != nil {
		return *x.PingRemoteFront
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetPingDefaultInternetRoute() uint32 {
	if x != nil && x.PingDefaultInternetRoute != nil {
		return *x.PingDefaultInternetRoute
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetE2EQualityLocal() *CMsgSteamDatagramConnectionQuality {
	if x != nil {
		return x.E2EQualityLocal
	}
	return nil
}

func (x *CGameNetworkingUI_ConnectionState) GetE2EQualityRemote() *CMsgSteamDatagramConnectionQuality {
	if x != nil {
		return x.E2EQualityRemote
	}
	return nil
}

func (x *CGameNetworkingUI_ConnectionState) GetE2EQualityRemoteInstantaneousTime() uint64 {
	if x != nil && x.E2EQualityRemoteInstantaneousTime != nil {
		return *x.E2EQualityRemoteInstantaneousTime
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetE2EQualityRemoteLifetimeTime() uint64 {
	if x != nil && x.E2EQualityRemoteLifetimeTime != nil {
		return *x.E2EQualityRemoteLifetimeTime
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetFrontQualityLocal() *CMsgSteamDatagramConnectionQuality {
	if x != nil {
		return x.FrontQualityLocal
	}
	return nil
}

func (x *CGameNetworkingUI_ConnectionState) GetFrontQualityRemote() *CMsgSteamDatagramConnectionQuality {
	if x != nil {
		return x.FrontQualityRemote
	}
	return nil
}

func (x *CGameNetworkingUI_ConnectionState) GetFrontQualityRemoteInstantaneousTime() uint64 {
	if x != nil && x.FrontQualityRemoteInstantaneousTime != nil {
		return *x.FrontQualityRemoteInstantaneousTime
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionState) GetFrontQualityRemoteLifetimeTime() uint64 {
	if x != nil && x.FrontQualityRemoteLifetimeTime != nil {
		return *x.FrontQualityRemoteLifetimeTime
	}
	return 0
}

type CGameNetworkingUI_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionState []*CGameNetworkingUI_ConnectionState `protobuf:"bytes,1,rep,name=connection_state,json=connectionState" json:"connection_state,omitempty"`
}

func (x *CGameNetworkingUI_Message) Reset() {
	*x = CGameNetworkingUI_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_gamenetworkingui_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGameNetworkingUI_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGameNetworkingUI_Message) ProtoMessage() {}

func (x *CGameNetworkingUI_Message) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_gamenetworkingui_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGameNetworkingUI_Message.ProtoReflect.Descriptor instead.
func (*CGameNetworkingUI_Message) Descriptor() ([]byte, []int) {
	return file_steammessages_gamenetworkingui_proto_rawDescGZIP(), []int{2}
}

func (x *CGameNetworkingUI_Message) GetConnectionState() []*CGameNetworkingUI_ConnectionState {
	if x != nil {
		return x.ConnectionState
	}
	return nil
}

type CGameNetworkingUI_ConnectionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransportKind            *uint32  `protobuf:"varint,1,opt,name=transport_kind,json=transportKind" json:"transport_kind,omitempty"`
	ConnectionState          *uint32  `protobuf:"varint,8,opt,name=connection_state,json=connectionState" json:"connection_state,omitempty"`
	SdrpopLocal              *string  `protobuf:"bytes,2,opt,name=sdrpop_local,json=sdrpopLocal" json:"sdrpop_local,omitempty"`
	SdrpopRemote             *string  `protobuf:"bytes,3,opt,name=sdrpop_remote,json=sdrpopRemote" json:"sdrpop_remote,omitempty"`
	PingMs                   *uint32  `protobuf:"varint,4,opt,name=ping_ms,json=pingMs" json:"ping_ms,omitempty"`
	PacketLoss               *float32 `protobuf:"fixed32,5,opt,name=packet_loss,json=packetLoss" json:"packet_loss,omitempty"`
	PingDefaultInternetRoute *uint32  `protobuf:"varint,6,opt,name=ping_default_internet_route,json=pingDefaultInternetRoute" json:"ping_default_internet_route,omitempty"`
	IpWasShared              *bool    `protobuf:"varint,7,opt,name=ip_was_shared,json=ipWasShared" json:"ip_was_shared,omitempty"`
}

func (x *CGameNetworkingUI_ConnectionSummary) Reset() {
	*x = CGameNetworkingUI_ConnectionSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_gamenetworkingui_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGameNetworkingUI_ConnectionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGameNetworkingUI_ConnectionSummary) ProtoMessage() {}

func (x *CGameNetworkingUI_ConnectionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_gamenetworkingui_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGameNetworkingUI_ConnectionSummary.ProtoReflect.Descriptor instead.
func (*CGameNetworkingUI_ConnectionSummary) Descriptor() ([]byte, []int) {
	return file_steammessages_gamenetworkingui_proto_rawDescGZIP(), []int{3}
}

func (x *CGameNetworkingUI_ConnectionSummary) GetTransportKind() uint32 {
	if x != nil && x.TransportKind != nil {
		return *x.TransportKind
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionSummary) GetConnectionState() uint32 {
	if x != nil && x.ConnectionState != nil {
		return *x.ConnectionState
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionSummary) GetSdrpopLocal() string {
	if x != nil && x.SdrpopLocal != nil {
		return *x.SdrpopLocal
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionSummary) GetSdrpopRemote() string {
	if x != nil && x.SdrpopRemote != nil {
		return *x.SdrpopRemote
	}
	return ""
}

func (x *CGameNetworkingUI_ConnectionSummary) GetPingMs() uint32 {
	if x != nil && x.PingMs != nil {
		return *x.PingMs
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionSummary) GetPacketLoss() float32 {
	if x != nil && x.PacketLoss != nil {
		return *x.PacketLoss
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionSummary) GetPingDefaultInternetRoute() uint32 {
	if x != nil && x.PingDefaultInternetRoute != nil {
		return *x.PingDefaultInternetRoute
	}
	return 0
}

func (x *CGameNetworkingUI_ConnectionSummary) GetIpWasShared() bool {
	if x != nil && x.IpWasShared != nil {
		return *x.IpWasShared
	}
	return false
}

type CGameNetworkingUI_AppSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid                    *uint32                              `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	IpWasSharedWithFriend    *bool                                `protobuf:"varint,10,opt,name=ip_was_shared_with_friend,json=ipWasSharedWithFriend" json:"ip_was_shared_with_friend,omitempty"`
	IpWasSharedWithNonfriend *bool                                `protobuf:"varint,11,opt,name=ip_was_shared_with_nonfriend,json=ipWasSharedWithNonfriend" json:"ip_was_shared_with_nonfriend,omitempty"`
	ActiveConnections        *uint32                              `protobuf:"varint,20,opt,name=active_connections,json=activeConnections" json:"active_connections,omitempty"`
	MainCxn                  *CGameNetworkingUI_ConnectionSummary `protobuf:"bytes,30,opt,name=main_cxn,json=mainCxn" json:"main_cxn,omitempty"`
}

func (x *CGameNetworkingUI_AppSummary) Reset() {
	*x = CGameNetworkingUI_AppSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_gamenetworkingui_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CGameNetworkingUI_AppSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CGameNetworkingUI_AppSummary) ProtoMessage() {}

func (x *CGameNetworkingUI_AppSummary) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_gamenetworkingui_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CGameNetworkingUI_AppSummary.ProtoReflect.Descriptor instead.
func (*CGameNetworkingUI_AppSummary) Descriptor() ([]byte, []int) {
	return file_steammessages_gamenetworkingui_proto_rawDescGZIP(), []int{4}
}

func (x *CGameNetworkingUI_AppSummary) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CGameNetworkingUI_AppSummary) GetIpWasSharedWithFriend() bool {
	if x != nil && x.IpWasSharedWithFriend != nil {
		return *x.IpWasSharedWithFriend
	}
	return false
}

func (x *CGameNetworkingUI_AppSummary) GetIpWasSharedWithNonfriend() bool {
	if x != nil && x.IpWasSharedWithNonfriend != nil {
		return *x.IpWasSharedWithNonfriend
	}
	return false
}

func (x *CGameNetworkingUI_AppSummary) GetActiveConnections() uint32 {
	if x != nil && x.ActiveConnections != nil {
		return *x.ActiveConnections
	}
	return 0
}

func (x *CGameNetworkingUI_AppSummary) GetMainCxn() *CGameNetworkingUI_ConnectionSummary {
	if x != nil {
		return x.MainCxn
	}
	return nil
}

var File_steammessages_gamenetworkingui_proto protoreflect.FileDescriptor

var file_steammessages_gamenetworkingui_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x75, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x73, 0x64, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1f, 0x0a, 0x1d, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x55, 0x49, 0x5f, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0xbb, 0x0b, 0x0a, 0x21, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x55, 0x49, 0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6c,
	0x6f, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x64, 0x72, 0x70, 0x6f, 0x70, 0x69,
	0x64, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x64, 0x72, 0x70, 0x6f, 0x70, 0x69, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x64, 0x72, 0x70, 0x6f, 0x70, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x64, 0x72, 0x70, 0x6f, 0x70, 0x69, 0x64, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x0b,
	0x70, 0x32, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x18, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x67, 0x72, 0x61, 0x6d, 0x50, 0x32, 0x50, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0a, 0x70, 0x32, 0x70, 0x52, 0x6f, 0x75, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x69, 0x6f, 0x72, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x69, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x69, 0x6e, 0x67, 0x5f,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x70, 0x69, 0x6e, 0x67, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x4f, 0x0a, 0x11, 0x65, 0x32, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0f, 0x65, 0x32, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x12, 0x51, 0x0a, 0x12, 0x65, 0x32, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67,
	0x72, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x10, 0x65, 0x32, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x25, 0x65, 0x32, 0x65, 0x5f, 0x71, 0x75,
	0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x04, 0x52, 0x21, 0x65, 0x32, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x61, 0x6e,
	0x65, 0x6f, 0x75, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x20, 0x65, 0x32, 0x65, 0x5f,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c,
	0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x1c, 0x65, 0x32, 0x65, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x53, 0x0a, 0x13, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x11, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x55, 0x0a, 0x14, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x29, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x12, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x51,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x54, 0x0a, 0x27,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x61, 0x6e, 0x65, 0x6f,
	0x75, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x23, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x22, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1e,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6a,
	0x0a, 0x19, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x55, 0x49, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x43, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x49, 0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xdc, 0x02, 0x0a, 0x23, 0x43,
	0x47, 0x61, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x49,
	0x5f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61,
	0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x64, 0x72, 0x70, 0x6f, 0x70, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x64, 0x72, 0x70,
	0x6f, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x64, 0x72, 0x70, 0x6f,
	0x70, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x64, 0x72, 0x70, 0x6f, 0x70, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70,
	0x69, 0x6e, 0x67, 0x4d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x6c, 0x6f, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x4c, 0x6f, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x70, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x70, 0x5f, 0x77, 0x61, 0x73, 0x5f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x70,
	0x57, 0x61, 0x73, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x22, 0x9e, 0x02, 0x0a, 0x1c, 0x43, 0x47,
	0x61, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x49, 0x5f,
	0x41, 0x70, 0x70, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64,
	0x12, 0x38, 0x0a, 0x19, 0x69, 0x70, 0x5f, 0x77, 0x61, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x70, 0x57, 0x61, 0x73, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x1c, 0x69, 0x70,
	0x5f, 0x77, 0x61, 0x73, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68,
	0x5f, 0x6e, 0x6f, 0x6e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x18, 0x69, 0x70, 0x57, 0x61, 0x73, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x57, 0x69, 0x74,
	0x68, 0x4e, 0x6f, 0x6e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x61, 0x69,
	0x6e, 0x5f, 0x63, 0x78, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x43, 0x47,
	0x61, 0x6d, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x55, 0x49, 0x5f,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x07, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x78, 0x6e, 0x42, 0x05, 0x48, 0x01, 0x80, 0x01,
	0x01,
}

var (
	file_steammessages_gamenetworkingui_proto_rawDescOnce sync.Once
	file_steammessages_gamenetworkingui_proto_rawDescData = file_steammessages_gamenetworkingui_proto_rawDesc
)

func file_steammessages_gamenetworkingui_proto_rawDescGZIP() []byte {
	file_steammessages_gamenetworkingui_proto_rawDescOnce.Do(func() {
		file_steammessages_gamenetworkingui_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_gamenetworkingui_proto_rawDescData)
	})
	return file_steammessages_gamenetworkingui_proto_rawDescData
}

var file_steammessages_gamenetworkingui_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_steammessages_gamenetworkingui_proto_goTypes = []any{
	(*CGameNetworkingUI_GlobalState)(nil),       // 0: CGameNetworkingUI_GlobalState
	(*CGameNetworkingUI_ConnectionState)(nil),   // 1: CGameNetworkingUI_ConnectionState
	(*CGameNetworkingUI_Message)(nil),           // 2: CGameNetworkingUI_Message
	(*CGameNetworkingUI_ConnectionSummary)(nil), // 3: CGameNetworkingUI_ConnectionSummary
	(*CGameNetworkingUI_AppSummary)(nil),        // 4: CGameNetworkingUI_AppSummary
	(*CMsgSteamDatagramP2PRoutingSummary)(nil),  // 5: CMsgSteamDatagramP2PRoutingSummary
	(*CMsgSteamDatagramConnectionQuality)(nil),  // 6: CMsgSteamDatagramConnectionQuality
}
var file_steammessages_gamenetworkingui_proto_depIdxs = []int32{
	5, // 0: CGameNetworkingUI_ConnectionState.p2p_routing:type_name -> CMsgSteamDatagramP2PRoutingSummary
	6, // 1: CGameNetworkingUI_ConnectionState.e2e_quality_local:type_name -> CMsgSteamDatagramConnectionQuality
	6, // 2: CGameNetworkingUI_ConnectionState.e2e_quality_remote:type_name -> CMsgSteamDatagramConnectionQuality
	6, // 3: CGameNetworkingUI_ConnectionState.front_quality_local:type_name -> CMsgSteamDatagramConnectionQuality
	6, // 4: CGameNetworkingUI_ConnectionState.front_quality_remote:type_name -> CMsgSteamDatagramConnectionQuality
	1, // 5: CGameNetworkingUI_Message.connection_state:type_name -> CGameNetworkingUI_ConnectionState
	3, // 6: CGameNetworkingUI_AppSummary.main_cxn:type_name -> CGameNetworkingUI_ConnectionSummary
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_steammessages_gamenetworkingui_proto_init() }
func file_steammessages_gamenetworkingui_proto_init() {
	if File_steammessages_gamenetworkingui_proto != nil {
		return
	}
	file_steamnetworkingsockets_messages_proto_init()
	file_steamdatagram_messages_sdr_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_gamenetworkingui_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CGameNetworkingUI_GlobalState); i {
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
		file_steammessages_gamenetworkingui_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CGameNetworkingUI_ConnectionState); i {
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
		file_steammessages_gamenetworkingui_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CGameNetworkingUI_Message); i {
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
		file_steammessages_gamenetworkingui_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CGameNetworkingUI_ConnectionSummary); i {
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
		file_steammessages_gamenetworkingui_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CGameNetworkingUI_AppSummary); i {
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
			RawDescriptor: file_steammessages_gamenetworkingui_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steammessages_gamenetworkingui_proto_goTypes,
		DependencyIndexes: file_steammessages_gamenetworkingui_proto_depIdxs,
		MessageInfos:      file_steammessages_gamenetworkingui_proto_msgTypes,
	}.Build()
	File_steammessages_gamenetworkingui_proto = out.File
	file_steammessages_gamenetworkingui_proto_rawDesc = nil
	file_steammessages_gamenetworkingui_proto_goTypes = nil
	file_steammessages_gamenetworkingui_proto_depIdxs = nil
}