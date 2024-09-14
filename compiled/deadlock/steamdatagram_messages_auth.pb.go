// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: steamdatagram_messages_auth.proto

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

type CMsgSteamDatagramRelayAuthTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeExpiry                           *uint32                                        `protobuf:"fixed32,1,opt,name=time_expiry,json=timeExpiry" json:"time_expiry,omitempty"`
	AuthorizedClientIdentityString       *string                                        `protobuf:"bytes,14,opt,name=authorized_client_identity_string,json=authorizedClientIdentityString" json:"authorized_client_identity_string,omitempty"`
	GameserverIdentityString             *string                                        `protobuf:"bytes,15,opt,name=gameserver_identity_string,json=gameserverIdentityString" json:"gameserver_identity_string,omitempty"`
	AuthorizedPublicIp                   *uint32                                        `protobuf:"fixed32,3,opt,name=authorized_public_ip,json=authorizedPublicIp" json:"authorized_public_ip,omitempty"`
	GameserverAddress                    []byte                                         `protobuf:"bytes,11,opt,name=gameserver_address,json=gameserverAddress" json:"gameserver_address,omitempty"`
	AppId                                *uint32                                        `protobuf:"varint,7,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	VirtualPort                          *uint32                                        `protobuf:"varint,10,opt,name=virtual_port,json=virtualPort" json:"virtual_port,omitempty"`
	ExtraFields                          []*CMsgSteamDatagramRelayAuthTicket_ExtraField `protobuf:"bytes,8,rep,name=extra_fields,json=extraFields" json:"extra_fields,omitempty"`
	LegacyAuthorizedSteamId              *uint64                                        `protobuf:"fixed64,2,opt,name=legacy_authorized_steam_id,json=legacyAuthorizedSteamId" json:"legacy_authorized_steam_id,omitempty"`
	LegacyGameserverSteamId              *uint64                                        `protobuf:"fixed64,4,opt,name=legacy_gameserver_steam_id,json=legacyGameserverSteamId" json:"legacy_gameserver_steam_id,omitempty"`
	LegacyGameserverPopId                *uint32                                        `protobuf:"fixed32,9,opt,name=legacy_gameserver_pop_id,json=legacyGameserverPopId" json:"legacy_gameserver_pop_id,omitempty"`
	LegacyAuthorizedClientIdentityBinary []byte                                         `protobuf:"bytes,12,opt,name=legacy_authorized_client_identity_binary,json=legacyAuthorizedClientIdentityBinary" json:"legacy_authorized_client_identity_binary,omitempty"`
	LegacyGameserverIdentityBinary       []byte                                         `protobuf:"bytes,13,opt,name=legacy_gameserver_identity_binary,json=legacyGameserverIdentityBinary" json:"legacy_gameserver_identity_binary,omitempty"`
}

func (x *CMsgSteamDatagramRelayAuthTicket) Reset() {
	*x = CMsgSteamDatagramRelayAuthTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steamdatagram_messages_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramRelayAuthTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramRelayAuthTicket) ProtoMessage() {}

func (x *CMsgSteamDatagramRelayAuthTicket) ProtoReflect() protoreflect.Message {
	mi := &file_steamdatagram_messages_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramRelayAuthTicket.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramRelayAuthTicket) Descriptor() ([]byte, []int) {
	return file_steamdatagram_messages_auth_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetTimeExpiry() uint32 {
	if x != nil && x.TimeExpiry != nil {
		return *x.TimeExpiry
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetAuthorizedClientIdentityString() string {
	if x != nil && x.AuthorizedClientIdentityString != nil {
		return *x.AuthorizedClientIdentityString
	}
	return ""
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetGameserverIdentityString() string {
	if x != nil && x.GameserverIdentityString != nil {
		return *x.GameserverIdentityString
	}
	return ""
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetAuthorizedPublicIp() uint32 {
	if x != nil && x.AuthorizedPublicIp != nil {
		return *x.AuthorizedPublicIp
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetGameserverAddress() []byte {
	if x != nil {
		return x.GameserverAddress
	}
	return nil
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetVirtualPort() uint32 {
	if x != nil && x.VirtualPort != nil {
		return *x.VirtualPort
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetExtraFields() []*CMsgSteamDatagramRelayAuthTicket_ExtraField {
	if x != nil {
		return x.ExtraFields
	}
	return nil
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetLegacyAuthorizedSteamId() uint64 {
	if x != nil && x.LegacyAuthorizedSteamId != nil {
		return *x.LegacyAuthorizedSteamId
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetLegacyGameserverSteamId() uint64 {
	if x != nil && x.LegacyGameserverSteamId != nil {
		return *x.LegacyGameserverSteamId
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetLegacyGameserverPopId() uint32 {
	if x != nil && x.LegacyGameserverPopId != nil {
		return *x.LegacyGameserverPopId
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetLegacyAuthorizedClientIdentityBinary() []byte {
	if x != nil {
		return x.LegacyAuthorizedClientIdentityBinary
	}
	return nil
}

func (x *CMsgSteamDatagramRelayAuthTicket) GetLegacyGameserverIdentityBinary() []byte {
	if x != nil {
		return x.LegacyGameserverIdentityBinary
	}
	return nil
}

type CMsgSteamDatagramSignedRelayAuthTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReservedDoNotUse *uint64                               `protobuf:"fixed64,1,opt,name=reserved_do_not_use,json=reservedDoNotUse" json:"reserved_do_not_use,omitempty"`
	Ticket           []byte                                `protobuf:"bytes,3,opt,name=ticket" json:"ticket,omitempty"`
	Signature        []byte                                `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	KeyId            *uint64                               `protobuf:"fixed64,2,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	Certs            []*CMsgSteamDatagramCertificateSigned `protobuf:"bytes,5,rep,name=certs" json:"certs,omitempty"`
}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) Reset() {
	*x = CMsgSteamDatagramSignedRelayAuthTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steamdatagram_messages_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramSignedRelayAuthTicket) ProtoMessage() {}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) ProtoReflect() protoreflect.Message {
	mi := &file_steamdatagram_messages_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramSignedRelayAuthTicket.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramSignedRelayAuthTicket) Descriptor() ([]byte, []int) {
	return file_steamdatagram_messages_auth_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) GetReservedDoNotUse() uint64 {
	if x != nil && x.ReservedDoNotUse != nil {
		return *x.ReservedDoNotUse
	}
	return 0
}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) GetTicket() []byte {
	if x != nil {
		return x.Ticket
	}
	return nil
}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) GetKeyId() uint64 {
	if x != nil && x.KeyId != nil {
		return *x.KeyId
	}
	return 0
}

func (x *CMsgSteamDatagramSignedRelayAuthTicket) GetCerts() []*CMsgSteamDatagramCertificateSigned {
	if x != nil {
		return x.Certs
	}
	return nil
}

type CMsgSteamDatagramCachedCredentialsForApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKey   []byte   `protobuf:"bytes,1,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	Cert         []byte   `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
	RelayTickets [][]byte `protobuf:"bytes,3,rep,name=relay_tickets,json=relayTickets" json:"relay_tickets,omitempty"`
}

func (x *CMsgSteamDatagramCachedCredentialsForApp) Reset() {
	*x = CMsgSteamDatagramCachedCredentialsForApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steamdatagram_messages_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramCachedCredentialsForApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramCachedCredentialsForApp) ProtoMessage() {}

func (x *CMsgSteamDatagramCachedCredentialsForApp) ProtoReflect() protoreflect.Message {
	mi := &file_steamdatagram_messages_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramCachedCredentialsForApp.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramCachedCredentialsForApp) Descriptor() ([]byte, []int) {
	return file_steamdatagram_messages_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgSteamDatagramCachedCredentialsForApp) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *CMsgSteamDatagramCachedCredentialsForApp) GetCert() []byte {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *CMsgSteamDatagramCachedCredentialsForApp) GetRelayTickets() [][]byte {
	if x != nil {
		return x.RelayTickets
	}
	return nil
}

type CMsgSteamDatagramGameCoordinatorServerLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeGenerated        *uint32 `protobuf:"varint,1,opt,name=time_generated,json=timeGenerated" json:"time_generated,omitempty"`
	Appid                *uint32 `protobuf:"varint,2,opt,name=appid" json:"appid,omitempty"`
	Routing              []byte  `protobuf:"bytes,3,opt,name=routing" json:"routing,omitempty"`
	Appdata              []byte  `protobuf:"bytes,4,opt,name=appdata" json:"appdata,omitempty"`
	LegacyIdentityBinary []byte  `protobuf:"bytes,5,opt,name=legacy_identity_binary,json=legacyIdentityBinary" json:"legacy_identity_binary,omitempty"`
	IdentityString       *string `protobuf:"bytes,6,opt,name=identity_string,json=identityString" json:"identity_string,omitempty"`
	DummySteamId         *uint64 `protobuf:"fixed64,99,opt,name=dummy_steam_id,json=dummySteamId" json:"dummy_steam_id,omitempty"`
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) Reset() {
	*x = CMsgSteamDatagramGameCoordinatorServerLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steamdatagram_messages_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramGameCoordinatorServerLogin) ProtoMessage() {}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) ProtoReflect() protoreflect.Message {
	mi := &file_steamdatagram_messages_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramGameCoordinatorServerLogin.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramGameCoordinatorServerLogin) Descriptor() ([]byte, []int) {
	return file_steamdatagram_messages_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) GetTimeGenerated() uint32 {
	if x != nil && x.TimeGenerated != nil {
		return *x.TimeGenerated
	}
	return 0
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) GetRouting() []byte {
	if x != nil {
		return x.Routing
	}
	return nil
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) GetAppdata() []byte {
	if x != nil {
		return x.Appdata
	}
	return nil
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) GetLegacyIdentityBinary() []byte {
	if x != nil {
		return x.LegacyIdentityBinary
	}
	return nil
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) GetIdentityString() string {
	if x != nil && x.IdentityString != nil {
		return *x.IdentityString
	}
	return ""
}

func (x *CMsgSteamDatagramGameCoordinatorServerLogin) GetDummySteamId() uint64 {
	if x != nil && x.DummySteamId != nil {
		return *x.DummySteamId
	}
	return 0
}

type CMsgSteamDatagramSignedGameCoordinatorServerLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert      *CMsgSteamDatagramCertificateSigned `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
	Login     []byte                              `protobuf:"bytes,2,opt,name=login" json:"login,omitempty"`
	Signature []byte                              `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (x *CMsgSteamDatagramSignedGameCoordinatorServerLogin) Reset() {
	*x = CMsgSteamDatagramSignedGameCoordinatorServerLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steamdatagram_messages_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramSignedGameCoordinatorServerLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramSignedGameCoordinatorServerLogin) ProtoMessage() {}

func (x *CMsgSteamDatagramSignedGameCoordinatorServerLogin) ProtoReflect() protoreflect.Message {
	mi := &file_steamdatagram_messages_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramSignedGameCoordinatorServerLogin.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramSignedGameCoordinatorServerLogin) Descriptor() ([]byte, []int) {
	return file_steamdatagram_messages_auth_proto_rawDescGZIP(), []int{4}
}

func (x *CMsgSteamDatagramSignedGameCoordinatorServerLogin) GetCert() *CMsgSteamDatagramCertificateSigned {
	if x != nil {
		return x.Cert
	}
	return nil
}

func (x *CMsgSteamDatagramSignedGameCoordinatorServerLogin) GetLogin() []byte {
	if x != nil {
		return x.Login
	}
	return nil
}

func (x *CMsgSteamDatagramSignedGameCoordinatorServerLogin) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type CMsgSteamDatagramHostedServerAddressPlaintext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ipv4            *uint32 `protobuf:"fixed32,1,opt,name=ipv4" json:"ipv4,omitempty"`
	Ipv6            []byte  `protobuf:"bytes,2,opt,name=ipv6" json:"ipv6,omitempty"`
	Port            *uint32 `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	RoutingSecret   *uint64 `protobuf:"fixed64,4,opt,name=routing_secret,json=routingSecret" json:"routing_secret,omitempty"`
	ProtocolVersion *uint32 `protobuf:"varint,5,opt,name=protocol_version,json=protocolVersion" json:"protocol_version,omitempty"`
}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) Reset() {
	*x = CMsgSteamDatagramHostedServerAddressPlaintext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steamdatagram_messages_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramHostedServerAddressPlaintext) ProtoMessage() {}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) ProtoReflect() protoreflect.Message {
	mi := &file_steamdatagram_messages_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramHostedServerAddressPlaintext.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramHostedServerAddressPlaintext) Descriptor() ([]byte, []int) {
	return file_steamdatagram_messages_auth_proto_rawDescGZIP(), []int{5}
}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) GetIpv4() uint32 {
	if x != nil && x.Ipv4 != nil {
		return *x.Ipv4
	}
	return 0
}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) GetIpv6() []byte {
	if x != nil {
		return x.Ipv6
	}
	return nil
}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) GetRoutingSecret() uint64 {
	if x != nil && x.RoutingSecret != nil {
		return *x.RoutingSecret
	}
	return 0
}

func (x *CMsgSteamDatagramHostedServerAddressPlaintext) GetProtocolVersion() uint32 {
	if x != nil && x.ProtocolVersion != nil {
		return *x.ProtocolVersion
	}
	return 0
}

type CMsgSteamDatagramRelayAuthTicket_ExtraField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StringValue  *string `protobuf:"bytes,2,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	Int64Value   *int64  `protobuf:"zigzag64,3,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Fixed64Value *uint64 `protobuf:"fixed64,5,opt,name=fixed64_value,json=fixed64Value" json:"fixed64_value,omitempty"`
}

func (x *CMsgSteamDatagramRelayAuthTicket_ExtraField) Reset() {
	*x = CMsgSteamDatagramRelayAuthTicket_ExtraField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steamdatagram_messages_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgSteamDatagramRelayAuthTicket_ExtraField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgSteamDatagramRelayAuthTicket_ExtraField) ProtoMessage() {}

func (x *CMsgSteamDatagramRelayAuthTicket_ExtraField) ProtoReflect() protoreflect.Message {
	mi := &file_steamdatagram_messages_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgSteamDatagramRelayAuthTicket_ExtraField.ProtoReflect.Descriptor instead.
func (*CMsgSteamDatagramRelayAuthTicket_ExtraField) Descriptor() ([]byte, []int) {
	return file_steamdatagram_messages_auth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetInt64Value() int64 {
	if x != nil && x.Int64Value != nil {
		return *x.Int64Value
	}
	return 0
}

func (x *CMsgSteamDatagramRelayAuthTicket_ExtraField) GetFixed64Value() uint64 {
	if x != nil && x.Fixed64Value != nil {
		return *x.Fixed64Value
	}
	return 0
}

var File_steamdatagram_messages_auth_proto protoreflect.FileDescriptor

var file_steamdatagram_messages_auth_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x64, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9a, 0x07, 0x0a, 0x20, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61,
	0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x21, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x3c, 0x0a, 0x1a, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x30, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x07, 0x52, 0x12, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x70, 0x12, 0x2d, 0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x4f, 0x0a, 0x0c, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x67, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x0b,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x3b, 0x0a, 0x1a, 0x6c,
	0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x5f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52,
	0x17, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x64, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x6c, 0x65, 0x67, 0x61,
	0x63, 0x79, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x17, 0x6c, 0x65,
	0x67, 0x61, 0x63, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x07, 0x52, 0x15, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x47,
	0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x56,
	0x0a, 0x28, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x24, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x49, 0x0a, 0x21, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79,
	0x5f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x1e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x1a, 0x89, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x72, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0a, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x06, 0x52,
	0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xdf, 0x01,
	0x0a, 0x26, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67,
	0x72, 0x61, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x10, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x44,
	0x6f, 0x4e, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x05, 0x6b,
	0x65, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x05, 0x63, 0x65, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44,
	0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x52, 0x05, 0x63, 0x65, 0x72, 0x74, 0x73, 0x22,
	0x84, 0x01, 0x0a, 0x28, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x67, 0x72, 0x61, 0x6d, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x46, 0x6f, 0x72, 0x41, 0x70, 0x70, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x63, 0x65, 0x72,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x22, 0xa3, 0x02, 0x0a, 0x2b, 0x43, 0x4d, 0x73, 0x67, 0x53,
	0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x47, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x74, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x70, 0x70, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x27, 0x0a,
	0x0f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x5f,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x63, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x22, 0xa0, 0x01, 0x0a,
	0x31, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x37, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61,
	0x67, 0x72, 0x61, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0xbd, 0x01, 0x0a, 0x2d, 0x43, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74,
	0x61, 0x67, 0x72, 0x61, 0x6d, 0x48, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52,
	0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0d, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x05, 0x48, 0x01, 0x80, 0x01, 0x00,
}

var (
	file_steamdatagram_messages_auth_proto_rawDescOnce sync.Once
	file_steamdatagram_messages_auth_proto_rawDescData = file_steamdatagram_messages_auth_proto_rawDesc
)

func file_steamdatagram_messages_auth_proto_rawDescGZIP() []byte {
	file_steamdatagram_messages_auth_proto_rawDescOnce.Do(func() {
		file_steamdatagram_messages_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_steamdatagram_messages_auth_proto_rawDescData)
	})
	return file_steamdatagram_messages_auth_proto_rawDescData
}

var file_steamdatagram_messages_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_steamdatagram_messages_auth_proto_goTypes = []any{
	(*CMsgSteamDatagramRelayAuthTicket)(nil),                  // 0: CMsgSteamDatagramRelayAuthTicket
	(*CMsgSteamDatagramSignedRelayAuthTicket)(nil),            // 1: CMsgSteamDatagramSignedRelayAuthTicket
	(*CMsgSteamDatagramCachedCredentialsForApp)(nil),          // 2: CMsgSteamDatagramCachedCredentialsForApp
	(*CMsgSteamDatagramGameCoordinatorServerLogin)(nil),       // 3: CMsgSteamDatagramGameCoordinatorServerLogin
	(*CMsgSteamDatagramSignedGameCoordinatorServerLogin)(nil), // 4: CMsgSteamDatagramSignedGameCoordinatorServerLogin
	(*CMsgSteamDatagramHostedServerAddressPlaintext)(nil),     // 5: CMsgSteamDatagramHostedServerAddressPlaintext
	(*CMsgSteamDatagramRelayAuthTicket_ExtraField)(nil),       // 6: CMsgSteamDatagramRelayAuthTicket.ExtraField
	(*CMsgSteamDatagramCertificateSigned)(nil),                // 7: CMsgSteamDatagramCertificateSigned
}
var file_steamdatagram_messages_auth_proto_depIdxs = []int32{
	6, // 0: CMsgSteamDatagramRelayAuthTicket.extra_fields:type_name -> CMsgSteamDatagramRelayAuthTicket.ExtraField
	7, // 1: CMsgSteamDatagramSignedRelayAuthTicket.certs:type_name -> CMsgSteamDatagramCertificateSigned
	7, // 2: CMsgSteamDatagramSignedGameCoordinatorServerLogin.cert:type_name -> CMsgSteamDatagramCertificateSigned
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_steamdatagram_messages_auth_proto_init() }
func file_steamdatagram_messages_auth_proto_init() {
	if File_steamdatagram_messages_auth_proto != nil {
		return
	}
	file_steamnetworkingsockets_messages_certs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steamdatagram_messages_auth_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgSteamDatagramRelayAuthTicket); i {
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
		file_steamdatagram_messages_auth_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgSteamDatagramSignedRelayAuthTicket); i {
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
		file_steamdatagram_messages_auth_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgSteamDatagramCachedCredentialsForApp); i {
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
		file_steamdatagram_messages_auth_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgSteamDatagramGameCoordinatorServerLogin); i {
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
		file_steamdatagram_messages_auth_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgSteamDatagramSignedGameCoordinatorServerLogin); i {
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
		file_steamdatagram_messages_auth_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgSteamDatagramHostedServerAddressPlaintext); i {
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
		file_steamdatagram_messages_auth_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgSteamDatagramRelayAuthTicket_ExtraField); i {
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
			RawDescriptor: file_steamdatagram_messages_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steamdatagram_messages_auth_proto_goTypes,
		DependencyIndexes: file_steamdatagram_messages_auth_proto_depIdxs,
		MessageInfos:      file_steamdatagram_messages_auth_proto_msgTypes,
	}.Build()
	File_steamdatagram_messages_auth_proto = out.File
	file_steamdatagram_messages_auth_proto_rawDesc = nil
	file_steamdatagram_messages_auth_proto_goTypes = nil
	file_steamdatagram_messages_auth_proto_depIdxs = nil
}
