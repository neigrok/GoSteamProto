// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: citadel_gamemessages.proto

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

type ECitadelGameMessages int32

const (
	ECitadelGameMessages_k_EMsgGameServerToClientConnectionStatus ECitadelGameMessages = 10
	ECitadelGameMessages_k_EMsgGameServerToClientInitialGameState ECitadelGameMessages = 12
	ECitadelGameMessages_k_EMsgGameServerToClientGameCompleted    ECitadelGameMessages = 13
	ECitadelGameMessages_k_EMsgGameServerToClientGoodbye          ECitadelGameMessages = 15
)

// Enum value maps for ECitadelGameMessages.
var (
	ECitadelGameMessages_name = map[int32]string{
		10: "k_EMsgGameServerToClientConnectionStatus",
		12: "k_EMsgGameServerToClientInitialGameState",
		13: "k_EMsgGameServerToClientGameCompleted",
		15: "k_EMsgGameServerToClientGoodbye",
	}
	ECitadelGameMessages_value = map[string]int32{
		"k_EMsgGameServerToClientConnectionStatus": 10,
		"k_EMsgGameServerToClientInitialGameState": 12,
		"k_EMsgGameServerToClientGameCompleted":    13,
		"k_EMsgGameServerToClientGoodbye":          15,
	}
)

func (x ECitadelGameMessages) Enum() *ECitadelGameMessages {
	p := new(ECitadelGameMessages)
	*p = x
	return p
}

func (x ECitadelGameMessages) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECitadelGameMessages) Descriptor() protoreflect.EnumDescriptor {
	return file_citadel_gamemessages_proto_enumTypes[0].Descriptor()
}

func (ECitadelGameMessages) Type() protoreflect.EnumType {
	return &file_citadel_gamemessages_proto_enumTypes[0]
}

func (x ECitadelGameMessages) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ECitadelGameMessages) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ECitadelGameMessages(num)
	return nil
}

// Deprecated: Use ECitadelGameMessages.Descriptor instead.
func (ECitadelGameMessages) EnumDescriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{0}
}

type ECitadelDisconnectReason int32

const (
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_UserLeaveMatch            ECitadelDisconnectReason = 1001
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_UserQuitApp               ECitadelDisconnectReason = 1002
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_UserCancel                ECitadelDisconnectReason = 1003
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_Goodbye                   ECitadelDisconnectReason = 1004
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_BadMessage                ECitadelDisconnectReason = 2001
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_GameDestroyedUnexpectedly ECitadelDisconnectReason = 2002
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_ChangingServer            ECitadelDisconnectReason = 2003
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_OldConnection             ECitadelDisconnectReason = 2004
	ECitadelDisconnectReason_k_ECitadelDisconnectReason_GoodbyeUnrecognizedGame   ECitadelDisconnectReason = 2005
)

// Enum value maps for ECitadelDisconnectReason.
var (
	ECitadelDisconnectReason_name = map[int32]string{
		1001: "k_ECitadelDisconnectReason_UserLeaveMatch",
		1002: "k_ECitadelDisconnectReason_UserQuitApp",
		1003: "k_ECitadelDisconnectReason_UserCancel",
		1004: "k_ECitadelDisconnectReason_Goodbye",
		2001: "k_ECitadelDisconnectReason_BadMessage",
		2002: "k_ECitadelDisconnectReason_GameDestroyedUnexpectedly",
		2003: "k_ECitadelDisconnectReason_ChangingServer",
		2004: "k_ECitadelDisconnectReason_OldConnection",
		2005: "k_ECitadelDisconnectReason_GoodbyeUnrecognizedGame",
	}
	ECitadelDisconnectReason_value = map[string]int32{
		"k_ECitadelDisconnectReason_UserLeaveMatch":            1001,
		"k_ECitadelDisconnectReason_UserQuitApp":               1002,
		"k_ECitadelDisconnectReason_UserCancel":                1003,
		"k_ECitadelDisconnectReason_Goodbye":                   1004,
		"k_ECitadelDisconnectReason_BadMessage":                2001,
		"k_ECitadelDisconnectReason_GameDestroyedUnexpectedly": 2002,
		"k_ECitadelDisconnectReason_ChangingServer":            2003,
		"k_ECitadelDisconnectReason_OldConnection":             2004,
		"k_ECitadelDisconnectReason_GoodbyeUnrecognizedGame":   2005,
	}
)

func (x ECitadelDisconnectReason) Enum() *ECitadelDisconnectReason {
	p := new(ECitadelDisconnectReason)
	*p = x
	return p
}

func (x ECitadelDisconnectReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECitadelDisconnectReason) Descriptor() protoreflect.EnumDescriptor {
	return file_citadel_gamemessages_proto_enumTypes[1].Descriptor()
}

func (ECitadelDisconnectReason) Type() protoreflect.EnumType {
	return &file_citadel_gamemessages_proto_enumTypes[1]
}

func (x ECitadelDisconnectReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ECitadelDisconnectReason) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ECitadelDisconnectReason(num)
	return nil
}

// Deprecated: Use ECitadelDisconnectReason.Descriptor instead.
func (ECitadelDisconnectReason) EnumDescriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{1}
}

type CMsgGameServerToClientConnectionStatus_EStatus int32

const (
	CMsgGameServerToClientConnectionStatus_k_EConnected    CMsgGameServerToClientConnectionStatus_EStatus = 1
	CMsgGameServerToClientConnectionStatus_k_EDisconnected CMsgGameServerToClientConnectionStatus_EStatus = 2
)

// Enum value maps for CMsgGameServerToClientConnectionStatus_EStatus.
var (
	CMsgGameServerToClientConnectionStatus_EStatus_name = map[int32]string{
		1: "k_EConnected",
		2: "k_EDisconnected",
	}
	CMsgGameServerToClientConnectionStatus_EStatus_value = map[string]int32{
		"k_EConnected":    1,
		"k_EDisconnected": 2,
	}
)

func (x CMsgGameServerToClientConnectionStatus_EStatus) Enum() *CMsgGameServerToClientConnectionStatus_EStatus {
	p := new(CMsgGameServerToClientConnectionStatus_EStatus)
	*p = x
	return p
}

func (x CMsgGameServerToClientConnectionStatus_EStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMsgGameServerToClientConnectionStatus_EStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_citadel_gamemessages_proto_enumTypes[2].Descriptor()
}

func (CMsgGameServerToClientConnectionStatus_EStatus) Type() protoreflect.EnumType {
	return &file_citadel_gamemessages_proto_enumTypes[2]
}

func (x CMsgGameServerToClientConnectionStatus_EStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CMsgGameServerToClientConnectionStatus_EStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CMsgGameServerToClientConnectionStatus_EStatus(num)
	return nil
}

// Deprecated: Use CMsgGameServerToClientConnectionStatus_EStatus.Descriptor instead.
func (CMsgGameServerToClientConnectionStatus_EStatus) EnumDescriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{3, 0}
}

type CMsgClientServerHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameInstanceId   *uint64               `protobuf:"varint,1,opt,name=game_instance_id,json=gameInstanceId" json:"game_instance_id,omitempty"`
	LocalPlayerIndex *uint32               `protobuf:"varint,2,opt,name=local_player_index,json=localPlayerIndex" json:"local_player_index,omitempty"`
	Payload          []byte                `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	MsgId            *ECitadelGameMessages `protobuf:"varint,4,opt,name=msg_id,json=msgId,enum=ECitadelGameMessages,def=10" json:"msg_id,omitempty"`
}

// Default values for CMsgClientServerHeader fields.
const (
	Default_CMsgClientServerHeader_MsgId = ECitadelGameMessages_k_EMsgGameServerToClientConnectionStatus
)

func (x *CMsgClientServerHeader) Reset() {
	*x = CMsgClientServerHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citadel_gamemessages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientServerHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientServerHeader) ProtoMessage() {}

func (x *CMsgClientServerHeader) ProtoReflect() protoreflect.Message {
	mi := &file_citadel_gamemessages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientServerHeader.ProtoReflect.Descriptor instead.
func (*CMsgClientServerHeader) Descriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgClientServerHeader) GetGameInstanceId() uint64 {
	if x != nil && x.GameInstanceId != nil {
		return *x.GameInstanceId
	}
	return 0
}

func (x *CMsgClientServerHeader) GetLocalPlayerIndex() uint32 {
	if x != nil && x.LocalPlayerIndex != nil {
		return *x.LocalPlayerIndex
	}
	return 0
}

func (x *CMsgClientServerHeader) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *CMsgClientServerHeader) GetMsgId() ECitadelGameMessages {
	if x != nil && x.MsgId != nil {
		return *x.MsgId
	}
	return Default_CMsgClientServerHeader_MsgId
}

type CMsgGameServerToClientGameCompleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CMsgGameServerToClientGameCompleted) Reset() {
	*x = CMsgGameServerToClientGameCompleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citadel_gamemessages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgGameServerToClientGameCompleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgGameServerToClientGameCompleted) ProtoMessage() {}

func (x *CMsgGameServerToClientGameCompleted) ProtoReflect() protoreflect.Message {
	mi := &file_citadel_gamemessages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgGameServerToClientGameCompleted.ProtoReflect.Descriptor instead.
func (*CMsgGameServerToClientGameCompleted) Descriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{1}
}

type CMsgGameServerToClientGoodbye struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CMsgGameServerToClientGoodbye) Reset() {
	*x = CMsgGameServerToClientGoodbye{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citadel_gamemessages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgGameServerToClientGoodbye) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgGameServerToClientGoodbye) ProtoMessage() {}

func (x *CMsgGameServerToClientGoodbye) ProtoReflect() protoreflect.Message {
	mi := &file_citadel_gamemessages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgGameServerToClientGoodbye.ProtoReflect.Descriptor instead.
func (*CMsgGameServerToClientGoodbye) Descriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{2}
}

type CMsgGameServerToClientConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*CMsgGameServerToClientConnectionStatus_Player `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
}

func (x *CMsgGameServerToClientConnectionStatus) Reset() {
	*x = CMsgGameServerToClientConnectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citadel_gamemessages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgGameServerToClientConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgGameServerToClientConnectionStatus) ProtoMessage() {}

func (x *CMsgGameServerToClientConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_citadel_gamemessages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgGameServerToClientConnectionStatus.ProtoReflect.Descriptor instead.
func (*CMsgGameServerToClientConnectionStatus) Descriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{3}
}

func (x *CMsgGameServerToClientConnectionStatus) GetPlayers() []*CMsgGameServerToClientConnectionStatus_Player {
	if x != nil {
		return x.Players
	}
	return nil
}

type CClientReconnectInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerSteamId        *uint64 `protobuf:"fixed64,1,opt,name=server_steam_id,json=serverSteamId" json:"server_steam_id,omitempty"`
	LobbyId              *uint64 `protobuf:"varint,2,opt,name=lobby_id,json=lobbyId" json:"lobby_id,omitempty"`
	TimeUpdated          *uint32 `protobuf:"varint,3,opt,name=time_updated,json=timeUpdated" json:"time_updated,omitempty"`
	UdpConnectIp         *uint32 `protobuf:"varint,4,opt,name=udp_connect_ip,json=udpConnectIp" json:"udp_connect_ip,omitempty"`
	UdpConnectPort       *uint32 `protobuf:"varint,5,opt,name=udp_connect_port,json=udpConnectPort" json:"udp_connect_port,omitempty"`
	CompatibilityVersion *uint32 `protobuf:"varint,6,opt,name=compatibility_version,json=compatibilityVersion" json:"compatibility_version,omitempty"`
}

func (x *CClientReconnectInfo) Reset() {
	*x = CClientReconnectInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citadel_gamemessages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClientReconnectInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClientReconnectInfo) ProtoMessage() {}

func (x *CClientReconnectInfo) ProtoReflect() protoreflect.Message {
	mi := &file_citadel_gamemessages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClientReconnectInfo.ProtoReflect.Descriptor instead.
func (*CClientReconnectInfo) Descriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{4}
}

func (x *CClientReconnectInfo) GetServerSteamId() uint64 {
	if x != nil && x.ServerSteamId != nil {
		return *x.ServerSteamId
	}
	return 0
}

func (x *CClientReconnectInfo) GetLobbyId() uint64 {
	if x != nil && x.LobbyId != nil {
		return *x.LobbyId
	}
	return 0
}

func (x *CClientReconnectInfo) GetTimeUpdated() uint32 {
	if x != nil && x.TimeUpdated != nil {
		return *x.TimeUpdated
	}
	return 0
}

func (x *CClientReconnectInfo) GetUdpConnectIp() uint32 {
	if x != nil && x.UdpConnectIp != nil {
		return *x.UdpConnectIp
	}
	return 0
}

func (x *CClientReconnectInfo) GetUdpConnectPort() uint32 {
	if x != nil && x.UdpConnectPort != nil {
		return *x.UdpConnectPort
	}
	return 0
}

func (x *CClientReconnectInfo) GetCompatibilityVersion() uint32 {
	if x != nil && x.CompatibilityVersion != nil {
		return *x.CompatibilityVersion
	}
	return 0
}

type CMsgClientAccountSyncStorageFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *uint32  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	Ids     []uint32 `protobuf:"varint,2,rep,name=ids" json:"ids,omitempty"`
	Values  []uint32 `protobuf:"varint,3,rep,name=values" json:"values,omitempty"`
}

func (x *CMsgClientAccountSyncStorageFile) Reset() {
	*x = CMsgClientAccountSyncStorageFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citadel_gamemessages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgClientAccountSyncStorageFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgClientAccountSyncStorageFile) ProtoMessage() {}

func (x *CMsgClientAccountSyncStorageFile) ProtoReflect() protoreflect.Message {
	mi := &file_citadel_gamemessages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgClientAccountSyncStorageFile.ProtoReflect.Descriptor instead.
func (*CMsgClientAccountSyncStorageFile) Descriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{5}
}

func (x *CMsgClientAccountSyncStorageFile) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *CMsgClientAccountSyncStorageFile) GetIds() []uint32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CMsgClientAccountSyncStorageFile) GetValues() []uint32 {
	if x != nil {
		return x.Values
	}
	return nil
}

type CMsgGameServerToClientConnectionStatus_Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerSlot                *int32                                          `protobuf:"varint,1,opt,name=player_slot,json=playerSlot,def=-1" json:"player_slot,omitempty"`
	Status                    *CMsgGameServerToClientConnectionStatus_EStatus `protobuf:"varint,2,opt,name=status,enum=CMsgGameServerToClientConnectionStatus_EStatus,def=1" json:"status,omitempty"`
	InactivityTicking         *bool                                           `protobuf:"varint,3,opt,name=inactivity_ticking,json=inactivityTicking" json:"inactivity_ticking,omitempty"`
	InactivityMsRemaining     *uint32                                         `protobuf:"varint,4,opt,name=inactivity_ms_remaining,json=inactivityMsRemaining" json:"inactivity_ms_remaining,omitempty"`
	InactivityAnimMsRemaining *uint32                                         `protobuf:"varint,5,opt,name=inactivity_anim_ms_remaining,json=inactivityAnimMsRemaining" json:"inactivity_anim_ms_remaining,omitempty"`
}

// Default values for CMsgGameServerToClientConnectionStatus_Player fields.
const (
	Default_CMsgGameServerToClientConnectionStatus_Player_PlayerSlot = int32(-1)
	Default_CMsgGameServerToClientConnectionStatus_Player_Status     = CMsgGameServerToClientConnectionStatus_k_EConnected
)

func (x *CMsgGameServerToClientConnectionStatus_Player) Reset() {
	*x = CMsgGameServerToClientConnectionStatus_Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_citadel_gamemessages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgGameServerToClientConnectionStatus_Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgGameServerToClientConnectionStatus_Player) ProtoMessage() {}

func (x *CMsgGameServerToClientConnectionStatus_Player) ProtoReflect() protoreflect.Message {
	mi := &file_citadel_gamemessages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgGameServerToClientConnectionStatus_Player.ProtoReflect.Descriptor instead.
func (*CMsgGameServerToClientConnectionStatus_Player) Descriptor() ([]byte, []int) {
	return file_citadel_gamemessages_proto_rawDescGZIP(), []int{3, 0}
}

func (x *CMsgGameServerToClientConnectionStatus_Player) GetPlayerSlot() int32 {
	if x != nil && x.PlayerSlot != nil {
		return *x.PlayerSlot
	}
	return Default_CMsgGameServerToClientConnectionStatus_Player_PlayerSlot
}

func (x *CMsgGameServerToClientConnectionStatus_Player) GetStatus() CMsgGameServerToClientConnectionStatus_EStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Default_CMsgGameServerToClientConnectionStatus_Player_Status
}

func (x *CMsgGameServerToClientConnectionStatus_Player) GetInactivityTicking() bool {
	if x != nil && x.InactivityTicking != nil {
		return *x.InactivityTicking
	}
	return false
}

func (x *CMsgGameServerToClientConnectionStatus_Player) GetInactivityMsRemaining() uint32 {
	if x != nil && x.InactivityMsRemaining != nil {
		return *x.InactivityMsRemaining
	}
	return 0
}

func (x *CMsgGameServerToClientConnectionStatus_Player) GetInactivityAnimMsRemaining() uint32 {
	if x != nil && x.InactivityAnimMsRemaining != nil {
		return *x.InactivityAnimMsRemaining
	}
	return 0
}

var File_citadel_gamemessages_proto protoreflect.FileDescriptor

var file_citadel_gamemessages_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x5f, 0x67, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01,
	0x0a, 0x16, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x56, 0x0a, 0x06, 0x6d, 0x73,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x45, 0x43, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x3a, 0x28, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x64, 0x22, 0x25, 0x0a, 0x23, 0x43, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x6d, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x43, 0x4d, 0x73,
	0x67, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x22, 0xd3, 0x03, 0x0a, 0x26, 0x43,
	0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x1a,
	0xac, 0x02, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x02, 0x2d, 0x31, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12,
	0x55, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x45, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x3a, 0x0c, 0x6b, 0x5f, 0x45, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x69,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x6d, 0x73, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x4d, 0x73, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x3f, 0x0a,
	0x1c, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x6e, 0x69, 0x6d,
	0x5f, 0x6d, 0x73, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x19, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x41,
	0x6e, 0x69, 0x6d, 0x4d, 0x73, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0x30,
	0x0a, 0x07, 0x45, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x6b, 0x5f, 0x45,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x6b,
	0x5f, 0x45, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02,
	0x22, 0x81, 0x02, 0x0a, 0x14, 0x43, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x75, 0x64, 0x70, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x75, 0x64, 0x70, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x49, 0x70, 0x12, 0x28, 0x0a, 0x10, 0x75, 0x64, 0x70, 0x5f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0e, 0x75, 0x64, 0x70, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x33, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x20, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2a, 0xc2, 0x01, 0x0a,
	0x14, 0x45, 0x43, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x28, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x10, 0x0a, 0x12, 0x2c, 0x0a, 0x28, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10,
	0x0c, 0x12, 0x29, 0x0a, 0x25, 0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x6d,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x0d, 0x12, 0x23, 0x0a, 0x1f,
	0x6b, 0x5f, 0x45, 0x4d, 0x73, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x6f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x10,
	0x0f, 0x2a, 0xcb, 0x03, 0x0a, 0x18, 0x45, 0x43, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x29, 0x6b, 0x5f, 0x45, 0x43, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x10, 0xe9, 0x07, 0x12, 0x2b,
	0x0a, 0x26, 0x6b, 0x5f, 0x45, 0x43, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x55, 0x73, 0x65,
	0x72, 0x51, 0x75, 0x69, 0x74, 0x41, 0x70, 0x70, 0x10, 0xea, 0x07, 0x12, 0x2a, 0x0a, 0x25, 0x6b,
	0x5f, 0x45, 0x43, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x10, 0xeb, 0x07, 0x12, 0x27, 0x0a, 0x22, 0x6b, 0x5f, 0x45, 0x43, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x10, 0xec, 0x07,
	0x12, 0x2a, 0x0a, 0x25, 0x6b, 0x5f, 0x45, 0x43, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x42,
	0x61, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0xd1, 0x0f, 0x12, 0x39, 0x0a, 0x34,
	0x6b, 0x5f, 0x45, 0x43, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x65, 0x64, 0x55, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x6c, 0x79, 0x10, 0xd2, 0x0f, 0x12, 0x2e, 0x0a, 0x29, 0x6b, 0x5f, 0x45, 0x43, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x10, 0xd3, 0x0f, 0x12, 0x2d, 0x0a, 0x28, 0x6b, 0x5f, 0x45, 0x43, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x4f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x10, 0xd4, 0x0f, 0x12, 0x37, 0x0a, 0x32, 0x6b, 0x5f, 0x45, 0x43, 0x69, 0x74,
	0x61, 0x64, 0x65, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x55, 0x6e, 0x72, 0x65,
	0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x10, 0xd5, 0x0f,
}

var (
	file_citadel_gamemessages_proto_rawDescOnce sync.Once
	file_citadel_gamemessages_proto_rawDescData = file_citadel_gamemessages_proto_rawDesc
)

func file_citadel_gamemessages_proto_rawDescGZIP() []byte {
	file_citadel_gamemessages_proto_rawDescOnce.Do(func() {
		file_citadel_gamemessages_proto_rawDescData = protoimpl.X.CompressGZIP(file_citadel_gamemessages_proto_rawDescData)
	})
	return file_citadel_gamemessages_proto_rawDescData
}

var file_citadel_gamemessages_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_citadel_gamemessages_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_citadel_gamemessages_proto_goTypes = []any{
	(ECitadelGameMessages)(0),                             // 0: ECitadelGameMessages
	(ECitadelDisconnectReason)(0),                         // 1: ECitadelDisconnectReason
	(CMsgGameServerToClientConnectionStatus_EStatus)(0),   // 2: CMsgGameServerToClientConnectionStatus.EStatus
	(*CMsgClientServerHeader)(nil),                        // 3: CMsgClientServerHeader
	(*CMsgGameServerToClientGameCompleted)(nil),           // 4: CMsgGameServerToClientGameCompleted
	(*CMsgGameServerToClientGoodbye)(nil),                 // 5: CMsgGameServerToClientGoodbye
	(*CMsgGameServerToClientConnectionStatus)(nil),        // 6: CMsgGameServerToClientConnectionStatus
	(*CClientReconnectInfo)(nil),                          // 7: CClientReconnectInfo
	(*CMsgClientAccountSyncStorageFile)(nil),              // 8: CMsgClientAccountSyncStorageFile
	(*CMsgGameServerToClientConnectionStatus_Player)(nil), // 9: CMsgGameServerToClientConnectionStatus.Player
}
var file_citadel_gamemessages_proto_depIdxs = []int32{
	0, // 0: CMsgClientServerHeader.msg_id:type_name -> ECitadelGameMessages
	9, // 1: CMsgGameServerToClientConnectionStatus.players:type_name -> CMsgGameServerToClientConnectionStatus.Player
	2, // 2: CMsgGameServerToClientConnectionStatus.Player.status:type_name -> CMsgGameServerToClientConnectionStatus.EStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_citadel_gamemessages_proto_init() }
func file_citadel_gamemessages_proto_init() {
	if File_citadel_gamemessages_proto != nil {
		return
	}
	file_citadel_gcmessages_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_citadel_gamemessages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientServerHeader); i {
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
		file_citadel_gamemessages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgGameServerToClientGameCompleted); i {
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
		file_citadel_gamemessages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgGameServerToClientGoodbye); i {
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
		file_citadel_gamemessages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgGameServerToClientConnectionStatus); i {
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
		file_citadel_gamemessages_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CClientReconnectInfo); i {
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
		file_citadel_gamemessages_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgClientAccountSyncStorageFile); i {
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
		file_citadel_gamemessages_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgGameServerToClientConnectionStatus_Player); i {
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
			RawDescriptor: file_citadel_gamemessages_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_citadel_gamemessages_proto_goTypes,
		DependencyIndexes: file_citadel_gamemessages_proto_depIdxs,
		EnumInfos:         file_citadel_gamemessages_proto_enumTypes,
		MessageInfos:      file_citadel_gamemessages_proto_msgTypes,
	}.Build()
	File_citadel_gamemessages_proto = out.File
	file_citadel_gamemessages_proto_rawDesc = nil
	file_citadel_gamemessages_proto_goTypes = nil
	file_citadel_gamemessages_proto_depIdxs = nil
}
