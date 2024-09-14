// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: dota_fighting_game_p2p_messages.proto

package dota2

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

type CP2P_FightingGame_GameData_EState int32

const (
	CP2P_FightingGame_GameData_k_ChoosingCharacter CP2P_FightingGame_GameData_EState = 1
	CP2P_FightingGame_GameData_k_Loaded            CP2P_FightingGame_GameData_EState = 2
	CP2P_FightingGame_GameData_k_Fighting          CP2P_FightingGame_GameData_EState = 3
)

// Enum value maps for CP2P_FightingGame_GameData_EState.
var (
	CP2P_FightingGame_GameData_EState_name = map[int32]string{
		1: "k_ChoosingCharacter",
		2: "k_Loaded",
		3: "k_Fighting",
	}
	CP2P_FightingGame_GameData_EState_value = map[string]int32{
		"k_ChoosingCharacter": 1,
		"k_Loaded":            2,
		"k_Fighting":          3,
	}
)

func (x CP2P_FightingGame_GameData_EState) Enum() *CP2P_FightingGame_GameData_EState {
	p := new(CP2P_FightingGame_GameData_EState)
	*p = x
	return p
}

func (x CP2P_FightingGame_GameData_EState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CP2P_FightingGame_GameData_EState) Descriptor() protoreflect.EnumDescriptor {
	return file_dota_fighting_game_p2p_messages_proto_enumTypes[0].Descriptor()
}

func (CP2P_FightingGame_GameData_EState) Type() protoreflect.EnumType {
	return &file_dota_fighting_game_p2p_messages_proto_enumTypes[0]
}

func (x CP2P_FightingGame_GameData_EState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CP2P_FightingGame_GameData_EState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CP2P_FightingGame_GameData_EState(num)
	return nil
}

// Deprecated: Use CP2P_FightingGame_GameData_EState.Descriptor instead.
func (CP2P_FightingGame_GameData_EState) EnumDescriptor() ([]byte, []int) {
	return file_dota_fighting_game_p2p_messages_proto_rawDescGZIP(), []int{3, 0}
}

type CMsgFightingGame_GameData_Fighting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastAckedFrame  *int32                                            `protobuf:"varint,1,opt,name=last_acked_frame,json=lastAckedFrame" json:"last_acked_frame,omitempty"`
	PlayerId        *uint32                                           `protobuf:"varint,2,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
	LastCrcFrame    *int32                                            `protobuf:"varint,3,opt,name=last_crc_frame,json=lastCrcFrame" json:"last_crc_frame,omitempty"`
	LastCrcValue    *uint32                                           `protobuf:"varint,4,opt,name=last_crc_value,json=lastCrcValue" json:"last_crc_value,omitempty"`
	Now             *float32                                          `protobuf:"fixed32,5,opt,name=now" json:"now,omitempty"`
	PeerAckTime     *float32                                          `protobuf:"fixed32,6,opt,name=peer_ack_time,json=peerAckTime" json:"peer_ack_time,omitempty"`
	InputStartFrame *int32                                            `protobuf:"varint,7,opt,name=input_start_frame,json=inputStartFrame" json:"input_start_frame,omitempty"`
	InputSample     []*CMsgFightingGame_GameData_Fighting_InputSample `protobuf:"bytes,8,rep,name=input_sample,json=inputSample" json:"input_sample,omitempty"`
}

func (x *CMsgFightingGame_GameData_Fighting) Reset() {
	*x = CMsgFightingGame_GameData_Fighting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgFightingGame_GameData_Fighting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgFightingGame_GameData_Fighting) ProtoMessage() {}

func (x *CMsgFightingGame_GameData_Fighting) ProtoReflect() protoreflect.Message {
	mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgFightingGame_GameData_Fighting.ProtoReflect.Descriptor instead.
func (*CMsgFightingGame_GameData_Fighting) Descriptor() ([]byte, []int) {
	return file_dota_fighting_game_p2p_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgFightingGame_GameData_Fighting) GetLastAckedFrame() int32 {
	if x != nil && x.LastAckedFrame != nil {
		return *x.LastAckedFrame
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Fighting) GetPlayerId() uint32 {
	if x != nil && x.PlayerId != nil {
		return *x.PlayerId
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Fighting) GetLastCrcFrame() int32 {
	if x != nil && x.LastCrcFrame != nil {
		return *x.LastCrcFrame
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Fighting) GetLastCrcValue() uint32 {
	if x != nil && x.LastCrcValue != nil {
		return *x.LastCrcValue
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Fighting) GetNow() float32 {
	if x != nil && x.Now != nil {
		return *x.Now
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Fighting) GetPeerAckTime() float32 {
	if x != nil && x.PeerAckTime != nil {
		return *x.PeerAckTime
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Fighting) GetInputStartFrame() int32 {
	if x != nil && x.InputStartFrame != nil {
		return *x.InputStartFrame
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Fighting) GetInputSample() []*CMsgFightingGame_GameData_Fighting_InputSample {
	if x != nil {
		return x.InputSample
	}
	return nil
}

type CMsgFightingGame_GameData_CharacterSelect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CursorIndex    *uint32                                           `protobuf:"varint,1,opt,name=cursor_index,json=cursorIndex" json:"cursor_index,omitempty"`
	SelectedHeroId *int32                                            `protobuf:"varint,2,opt,name=selected_hero_id,json=selectedHeroId" json:"selected_hero_id,omitempty"`
	SelectedStyle  *uint32                                           `protobuf:"varint,3,opt,name=selected_style,json=selectedStyle" json:"selected_style,omitempty"`
	EconItemRefs   []*CMsgFightingGame_GameData_CharacterSelect_Item `protobuf:"bytes,4,rep,name=econ_item_refs,json=econItemRefs" json:"econ_item_refs,omitempty"`
	MessageAck     *int64                                            `protobuf:"varint,5,opt,name=message_ack,json=messageAck" json:"message_ack,omitempty"`
	ConfirmedStyle *bool                                             `protobuf:"varint,6,opt,name=confirmed_style,json=confirmedStyle" json:"confirmed_style,omitempty"`
}

func (x *CMsgFightingGame_GameData_CharacterSelect) Reset() {
	*x = CMsgFightingGame_GameData_CharacterSelect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgFightingGame_GameData_CharacterSelect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgFightingGame_GameData_CharacterSelect) ProtoMessage() {}

func (x *CMsgFightingGame_GameData_CharacterSelect) ProtoReflect() protoreflect.Message {
	mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgFightingGame_GameData_CharacterSelect.ProtoReflect.Descriptor instead.
func (*CMsgFightingGame_GameData_CharacterSelect) Descriptor() ([]byte, []int) {
	return file_dota_fighting_game_p2p_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgFightingGame_GameData_CharacterSelect) GetCursorIndex() uint32 {
	if x != nil && x.CursorIndex != nil {
		return *x.CursorIndex
	}
	return 0
}

func (x *CMsgFightingGame_GameData_CharacterSelect) GetSelectedHeroId() int32 {
	if x != nil && x.SelectedHeroId != nil {
		return *x.SelectedHeroId
	}
	return 0
}

func (x *CMsgFightingGame_GameData_CharacterSelect) GetSelectedStyle() uint32 {
	if x != nil && x.SelectedStyle != nil {
		return *x.SelectedStyle
	}
	return 0
}

func (x *CMsgFightingGame_GameData_CharacterSelect) GetEconItemRefs() []*CMsgFightingGame_GameData_CharacterSelect_Item {
	if x != nil {
		return x.EconItemRefs
	}
	return nil
}

func (x *CMsgFightingGame_GameData_CharacterSelect) GetMessageAck() int64 {
	if x != nil && x.MessageAck != nil {
		return *x.MessageAck
	}
	return 0
}

func (x *CMsgFightingGame_GameData_CharacterSelect) GetConfirmedStyle() bool {
	if x != nil && x.ConfirmedStyle != nil {
		return *x.ConfirmedStyle
	}
	return false
}

type CMsgFightingGame_GameData_Loaded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Now               *float32 `protobuf:"fixed32,1,opt,name=now" json:"now,omitempty"`
	PeerAckTime       *float32 `protobuf:"fixed32,2,opt,name=peer_ack_time,json=peerAckTime" json:"peer_ack_time,omitempty"`
	ProposedStartTime *float32 `protobuf:"fixed32,3,opt,name=proposed_start_time,json=proposedStartTime" json:"proposed_start_time,omitempty"`
	AcceptedStartTime *float32 `protobuf:"fixed32,4,opt,name=accepted_start_time,json=acceptedStartTime" json:"accepted_start_time,omitempty"`
}

func (x *CMsgFightingGame_GameData_Loaded) Reset() {
	*x = CMsgFightingGame_GameData_Loaded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgFightingGame_GameData_Loaded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgFightingGame_GameData_Loaded) ProtoMessage() {}

func (x *CMsgFightingGame_GameData_Loaded) ProtoReflect() protoreflect.Message {
	mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgFightingGame_GameData_Loaded.ProtoReflect.Descriptor instead.
func (*CMsgFightingGame_GameData_Loaded) Descriptor() ([]byte, []int) {
	return file_dota_fighting_game_p2p_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgFightingGame_GameData_Loaded) GetNow() float32 {
	if x != nil && x.Now != nil {
		return *x.Now
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Loaded) GetPeerAckTime() float32 {
	if x != nil && x.PeerAckTime != nil {
		return *x.PeerAckTime
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Loaded) GetProposedStartTime() float32 {
	if x != nil && x.ProposedStartTime != nil {
		return *x.ProposedStartTime
	}
	return 0
}

func (x *CMsgFightingGame_GameData_Loaded) GetAcceptedStartTime() float32 {
	if x != nil && x.AcceptedStartTime != nil {
		return *x.AcceptedStartTime
	}
	return 0
}

type CP2P_FightingGame_GameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State *CP2P_FightingGame_GameData_EState `protobuf:"varint,1,opt,name=state,enum=CP2P_FightingGame_GameData_EState,def=1" json:"state,omitempty"`
	// Types that are assignable to StateData:
	//
	//	*CP2P_FightingGame_GameData_Fight
	//	*CP2P_FightingGame_GameData_CharacterSelect
	//	*CP2P_FightingGame_GameData_Loaded
	StateData isCP2P_FightingGame_GameData_StateData `protobuf_oneof:"state_data"`
}

// Default values for CP2P_FightingGame_GameData fields.
const (
	Default_CP2P_FightingGame_GameData_State = CP2P_FightingGame_GameData_k_ChoosingCharacter
)

func (x *CP2P_FightingGame_GameData) Reset() {
	*x = CP2P_FightingGame_GameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CP2P_FightingGame_GameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CP2P_FightingGame_GameData) ProtoMessage() {}

func (x *CP2P_FightingGame_GameData) ProtoReflect() protoreflect.Message {
	mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CP2P_FightingGame_GameData.ProtoReflect.Descriptor instead.
func (*CP2P_FightingGame_GameData) Descriptor() ([]byte, []int) {
	return file_dota_fighting_game_p2p_messages_proto_rawDescGZIP(), []int{3}
}

func (x *CP2P_FightingGame_GameData) GetState() CP2P_FightingGame_GameData_EState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return Default_CP2P_FightingGame_GameData_State
}

func (m *CP2P_FightingGame_GameData) GetStateData() isCP2P_FightingGame_GameData_StateData {
	if m != nil {
		return m.StateData
	}
	return nil
}

func (x *CP2P_FightingGame_GameData) GetFight() *CMsgFightingGame_GameData_Fighting {
	if x, ok := x.GetStateData().(*CP2P_FightingGame_GameData_Fight); ok {
		return x.Fight
	}
	return nil
}

func (x *CP2P_FightingGame_GameData) GetCharacterSelect() *CMsgFightingGame_GameData_CharacterSelect {
	if x, ok := x.GetStateData().(*CP2P_FightingGame_GameData_CharacterSelect); ok {
		return x.CharacterSelect
	}
	return nil
}

func (x *CP2P_FightingGame_GameData) GetLoaded() *CMsgFightingGame_GameData_Loaded {
	if x, ok := x.GetStateData().(*CP2P_FightingGame_GameData_Loaded); ok {
		return x.Loaded
	}
	return nil
}

type isCP2P_FightingGame_GameData_StateData interface {
	isCP2P_FightingGame_GameData_StateData()
}

type CP2P_FightingGame_GameData_Fight struct {
	Fight *CMsgFightingGame_GameData_Fighting `protobuf:"bytes,2,opt,name=fight,oneof"`
}

type CP2P_FightingGame_GameData_CharacterSelect struct {
	CharacterSelect *CMsgFightingGame_GameData_CharacterSelect `protobuf:"bytes,3,opt,name=character_select,json=characterSelect,oneof"`
}

type CP2P_FightingGame_GameData_Loaded struct {
	Loaded *CMsgFightingGame_GameData_Loaded `protobuf:"bytes,4,opt,name=loaded,oneof"`
}

func (*CP2P_FightingGame_GameData_Fight) isCP2P_FightingGame_GameData_StateData() {}

func (*CP2P_FightingGame_GameData_CharacterSelect) isCP2P_FightingGame_GameData_StateData() {}

func (*CP2P_FightingGame_GameData_Loaded) isCP2P_FightingGame_GameData_StateData() {}

type CMsgFightingGame_GameData_Fighting_InputSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ButtonMask *uint32 `protobuf:"varint,1,opt,name=button_mask,json=buttonMask" json:"button_mask,omitempty"`
}

func (x *CMsgFightingGame_GameData_Fighting_InputSample) Reset() {
	*x = CMsgFightingGame_GameData_Fighting_InputSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgFightingGame_GameData_Fighting_InputSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgFightingGame_GameData_Fighting_InputSample) ProtoMessage() {}

func (x *CMsgFightingGame_GameData_Fighting_InputSample) ProtoReflect() protoreflect.Message {
	mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgFightingGame_GameData_Fighting_InputSample.ProtoReflect.Descriptor instead.
func (*CMsgFightingGame_GameData_Fighting_InputSample) Descriptor() ([]byte, []int) {
	return file_dota_fighting_game_p2p_messages_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CMsgFightingGame_GameData_Fighting_InputSample) GetButtonMask() uint32 {
	if x != nil && x.ButtonMask != nil {
		return *x.ButtonMask
	}
	return 0
}

type CMsgFightingGame_GameData_CharacterSelect_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemDef    *uint32 `protobuf:"varint,1,opt,name=item_def,json=itemDef" json:"item_def,omitempty"`
	StyleIndex *uint32 `protobuf:"varint,2,opt,name=style_index,json=styleIndex,def=255" json:"style_index,omitempty"`
}

// Default values for CMsgFightingGame_GameData_CharacterSelect_Item fields.
const (
	Default_CMsgFightingGame_GameData_CharacterSelect_Item_StyleIndex = uint32(255)
)

func (x *CMsgFightingGame_GameData_CharacterSelect_Item) Reset() {
	*x = CMsgFightingGame_GameData_CharacterSelect_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgFightingGame_GameData_CharacterSelect_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgFightingGame_GameData_CharacterSelect_Item) ProtoMessage() {}

func (x *CMsgFightingGame_GameData_CharacterSelect_Item) ProtoReflect() protoreflect.Message {
	mi := &file_dota_fighting_game_p2p_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgFightingGame_GameData_CharacterSelect_Item.ProtoReflect.Descriptor instead.
func (*CMsgFightingGame_GameData_CharacterSelect_Item) Descriptor() ([]byte, []int) {
	return file_dota_fighting_game_p2p_messages_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CMsgFightingGame_GameData_CharacterSelect_Item) GetItemDef() uint32 {
	if x != nil && x.ItemDef != nil {
		return *x.ItemDef
	}
	return 0
}

func (x *CMsgFightingGame_GameData_CharacterSelect_Item) GetStyleIndex() uint32 {
	if x != nil && x.StyleIndex != nil {
		return *x.StyleIndex
	}
	return Default_CMsgFightingGame_GameData_CharacterSelect_Item_StyleIndex
}

var File_dota_fighting_game_p2p_messages_proto protoreflect.FileDescriptor

var file_dota_fighting_game_p2p_messages_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x66, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x32, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x65, 0x74, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9d, 0x03, 0x0a, 0x22, 0x43, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x67, 0x68, 0x74,
	0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x5f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x41, 0x63, 0x6b, 0x65, 0x64, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x72, 0x63, 0x5f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x72,
	0x63, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63,
	0x72, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x43, 0x72, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6e, 0x6f, 0x77, 0x12, 0x22,
	0x0a, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x41, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x52,
	0x0a, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x67, 0x68, 0x74,
	0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x5f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x1a, 0x2e, 0x0a, 0x0b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x4d, 0x61,
	0x73, 0x6b, 0x22, 0x89, 0x03, 0x0a, 0x29, 0x43, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x67, 0x68, 0x74,
	0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x5f, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f,
	0x68, 0x65, 0x72, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x48, 0x65, 0x72, 0x6f, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x79, 0x6c, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x65, 0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x43,
	0x4d, 0x73, 0x67, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x5f,
	0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x65,
	0x63, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x66, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64,
	0x53, 0x74, 0x79, 0x6c, 0x65, 0x1a, 0x47, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x66, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x74, 0x79, 0x6c,
	0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x03, 0x32,
	0x35, 0x35, 0x52, 0x0a, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0xb8,
	0x01, 0x0a, 0x20, 0x43, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x47,
	0x61, 0x6d, 0x65, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x4c, 0x6f, 0x61,
	0x64, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6e, 0x6f, 0x77, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x65,
	0x65, 0x72, 0x41, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8d, 0x03, 0x0a, 0x1a, 0x43, 0x50,
	0x32, 0x50, 0x5f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x5f,
	0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x4d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x43, 0x50, 0x32, 0x50, 0x5f, 0x46,
	0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x47, 0x61, 0x6d, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3a, 0x13, 0x6b, 0x5f, 0x43,
	0x68, 0x6f, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x66, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x67,
	0x68, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x5f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x05, 0x66,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x57, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x43, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d,
	0x65, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a,
	0x06, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x43, 0x4d, 0x73, 0x67, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65,
	0x5f, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64,
	0x48, 0x00, 0x52, 0x06, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x06, 0x45, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x6b, 0x5f, 0x43, 0x68, 0x6f, 0x6f, 0x73, 0x69,
	0x6e, 0x67, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x6b, 0x5f, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x6b,
	0x5f, 0x46, 0x69, 0x67, 0x68, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x42, 0x0c, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61,
}

var (
	file_dota_fighting_game_p2p_messages_proto_rawDescOnce sync.Once
	file_dota_fighting_game_p2p_messages_proto_rawDescData = file_dota_fighting_game_p2p_messages_proto_rawDesc
)

func file_dota_fighting_game_p2p_messages_proto_rawDescGZIP() []byte {
	file_dota_fighting_game_p2p_messages_proto_rawDescOnce.Do(func() {
		file_dota_fighting_game_p2p_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_dota_fighting_game_p2p_messages_proto_rawDescData)
	})
	return file_dota_fighting_game_p2p_messages_proto_rawDescData
}

var file_dota_fighting_game_p2p_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dota_fighting_game_p2p_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dota_fighting_game_p2p_messages_proto_goTypes = []any{
	(CP2P_FightingGame_GameData_EState)(0),                 // 0: CP2P_FightingGame_GameData.EState
	(*CMsgFightingGame_GameData_Fighting)(nil),             // 1: CMsgFightingGame_GameData_Fighting
	(*CMsgFightingGame_GameData_CharacterSelect)(nil),      // 2: CMsgFightingGame_GameData_CharacterSelect
	(*CMsgFightingGame_GameData_Loaded)(nil),               // 3: CMsgFightingGame_GameData_Loaded
	(*CP2P_FightingGame_GameData)(nil),                     // 4: CP2P_FightingGame_GameData
	(*CMsgFightingGame_GameData_Fighting_InputSample)(nil), // 5: CMsgFightingGame_GameData_Fighting.InputSample
	(*CMsgFightingGame_GameData_CharacterSelect_Item)(nil), // 6: CMsgFightingGame_GameData_CharacterSelect.Item
}
var file_dota_fighting_game_p2p_messages_proto_depIdxs = []int32{
	5, // 0: CMsgFightingGame_GameData_Fighting.input_sample:type_name -> CMsgFightingGame_GameData_Fighting.InputSample
	6, // 1: CMsgFightingGame_GameData_CharacterSelect.econ_item_refs:type_name -> CMsgFightingGame_GameData_CharacterSelect.Item
	0, // 2: CP2P_FightingGame_GameData.state:type_name -> CP2P_FightingGame_GameData.EState
	1, // 3: CP2P_FightingGame_GameData.fight:type_name -> CMsgFightingGame_GameData_Fighting
	2, // 4: CP2P_FightingGame_GameData.character_select:type_name -> CMsgFightingGame_GameData_CharacterSelect
	3, // 5: CP2P_FightingGame_GameData.loaded:type_name -> CMsgFightingGame_GameData_Loaded
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_dota_fighting_game_p2p_messages_proto_init() }
func file_dota_fighting_game_p2p_messages_proto_init() {
	if File_dota_fighting_game_p2p_messages_proto != nil {
		return
	}
	file_netmessages_proto_init()
	file_networkbasetypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dota_fighting_game_p2p_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgFightingGame_GameData_Fighting); i {
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
		file_dota_fighting_game_p2p_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgFightingGame_GameData_CharacterSelect); i {
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
		file_dota_fighting_game_p2p_messages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgFightingGame_GameData_Loaded); i {
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
		file_dota_fighting_game_p2p_messages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CP2P_FightingGame_GameData); i {
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
		file_dota_fighting_game_p2p_messages_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgFightingGame_GameData_Fighting_InputSample); i {
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
		file_dota_fighting_game_p2p_messages_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgFightingGame_GameData_CharacterSelect_Item); i {
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
	file_dota_fighting_game_p2p_messages_proto_msgTypes[3].OneofWrappers = []any{
		(*CP2P_FightingGame_GameData_Fight)(nil),
		(*CP2P_FightingGame_GameData_CharacterSelect)(nil),
		(*CP2P_FightingGame_GameData_Loaded)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dota_fighting_game_p2p_messages_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota_fighting_game_p2p_messages_proto_goTypes,
		DependencyIndexes: file_dota_fighting_game_p2p_messages_proto_depIdxs,
		EnumInfos:         file_dota_fighting_game_p2p_messages_proto_enumTypes,
		MessageInfos:      file_dota_fighting_game_p2p_messages_proto_msgTypes,
	}.Build()
	File_dota_fighting_game_p2p_messages_proto = out.File
	file_dota_fighting_game_p2p_messages_proto_rawDesc = nil
	file_dota_fighting_game_p2p_messages_proto_goTypes = nil
	file_dota_fighting_game_p2p_messages_proto_depIdxs = nil
}
