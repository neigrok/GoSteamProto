// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: dota_modifiers.proto

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

type DOTA_MODIFIER_ENTRY_TYPE int32

const (
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE  DOTA_MODIFIER_ENTRY_TYPE = 1
	DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_REMOVED DOTA_MODIFIER_ENTRY_TYPE = 2
)

// Enum value maps for DOTA_MODIFIER_ENTRY_TYPE.
var (
	DOTA_MODIFIER_ENTRY_TYPE_name = map[int32]string{
		1: "DOTA_MODIFIER_ENTRY_TYPE_ACTIVE",
		2: "DOTA_MODIFIER_ENTRY_TYPE_REMOVED",
	}
	DOTA_MODIFIER_ENTRY_TYPE_value = map[string]int32{
		"DOTA_MODIFIER_ENTRY_TYPE_ACTIVE":  1,
		"DOTA_MODIFIER_ENTRY_TYPE_REMOVED": 2,
	}
)

func (x DOTA_MODIFIER_ENTRY_TYPE) Enum() *DOTA_MODIFIER_ENTRY_TYPE {
	p := new(DOTA_MODIFIER_ENTRY_TYPE)
	*p = x
	return p
}

func (x DOTA_MODIFIER_ENTRY_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DOTA_MODIFIER_ENTRY_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_dota_modifiers_proto_enumTypes[0].Descriptor()
}

func (DOTA_MODIFIER_ENTRY_TYPE) Type() protoreflect.EnumType {
	return &file_dota_modifiers_proto_enumTypes[0]
}

func (x DOTA_MODIFIER_ENTRY_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DOTA_MODIFIER_ENTRY_TYPE) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DOTA_MODIFIER_ENTRY_TYPE(num)
	return nil
}

// Deprecated: Use DOTA_MODIFIER_ENTRY_TYPE.Descriptor instead.
func (DOTA_MODIFIER_ENTRY_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_dota_modifiers_proto_rawDescGZIP(), []int{0}
}

type CDOTAModifierBuffTableEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryType           *DOTA_MODIFIER_ENTRY_TYPE `protobuf:"varint,1,req,name=entry_type,json=entryType,enum=DOTA_MODIFIER_ENTRY_TYPE,def=1" json:"entry_type,omitempty"`
	Parent              *uint32                   `protobuf:"varint,2,req,name=parent,def=16777215" json:"parent,omitempty"`
	Index               *int32                    `protobuf:"varint,3,req,name=index" json:"index,omitempty"`
	SerialNum           *int32                    `protobuf:"varint,4,req,name=serial_num,json=serialNum" json:"serial_num,omitempty"`
	ModifierClass       *int32                    `protobuf:"varint,5,opt,name=modifier_class,json=modifierClass" json:"modifier_class,omitempty"`
	AbilityLevel        *int32                    `protobuf:"varint,6,opt,name=ability_level,json=abilityLevel" json:"ability_level,omitempty"`
	StackCount          *int32                    `protobuf:"varint,7,opt,name=stack_count,json=stackCount" json:"stack_count,omitempty"`
	CreationTime        *float32                  `protobuf:"fixed32,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	Duration            *float32                  `protobuf:"fixed32,9,opt,name=duration,def=-1" json:"duration,omitempty"`
	Caster              *uint32                   `protobuf:"varint,10,opt,name=caster,def=16777215" json:"caster,omitempty"`
	Ability             *uint32                   `protobuf:"varint,11,opt,name=ability,def=16777215" json:"ability,omitempty"`
	Armor               *int32                    `protobuf:"varint,12,opt,name=armor" json:"armor,omitempty"`
	FadeTime            *float32                  `protobuf:"fixed32,13,opt,name=fade_time,json=fadeTime" json:"fade_time,omitempty"`
	Subtle              *bool                     `protobuf:"varint,14,opt,name=subtle" json:"subtle,omitempty"`
	ChannelTime         *float32                  `protobuf:"fixed32,15,opt,name=channel_time,json=channelTime" json:"channel_time,omitempty"`
	VStart              *CMsgVector               `protobuf:"bytes,16,opt,name=v_start,json=vStart" json:"v_start,omitempty"`
	VEnd                *CMsgVector               `protobuf:"bytes,17,opt,name=v_end,json=vEnd" json:"v_end,omitempty"`
	PortalLoopAppear    *string                   `protobuf:"bytes,18,opt,name=portal_loop_appear,json=portalLoopAppear" json:"portal_loop_appear,omitempty"`
	PortalLoopDisappear *string                   `protobuf:"bytes,19,opt,name=portal_loop_disappear,json=portalLoopDisappear" json:"portal_loop_disappear,omitempty"`
	HeroLoopAppear      *string                   `protobuf:"bytes,20,opt,name=hero_loop_appear,json=heroLoopAppear" json:"hero_loop_appear,omitempty"`
	HeroLoopDisappear   *string                   `protobuf:"bytes,21,opt,name=hero_loop_disappear,json=heroLoopDisappear" json:"hero_loop_disappear,omitempty"`
	MovementSpeed       *int32                    `protobuf:"varint,22,opt,name=movement_speed,json=movementSpeed" json:"movement_speed,omitempty"`
	Aura                *bool                     `protobuf:"varint,23,opt,name=aura" json:"aura,omitempty"`
	Activity            *int32                    `protobuf:"varint,24,opt,name=activity" json:"activity,omitempty"`
	Damage              *int32                    `protobuf:"varint,25,opt,name=damage" json:"damage,omitempty"`
	Range               *int32                    `protobuf:"varint,26,opt,name=range" json:"range,omitempty"`
	DdModifierIndex     *int32                    `protobuf:"varint,27,opt,name=dd_modifier_index,json=ddModifierIndex" json:"dd_modifier_index,omitempty"`
	DdAbilityId         *int32                    `protobuf:"varint,28,opt,name=dd_ability_id,json=ddAbilityId,def=-1" json:"dd_ability_id,omitempty"`
	IllusionLabel       *string                   `protobuf:"bytes,29,opt,name=illusion_label,json=illusionLabel" json:"illusion_label,omitempty"`
	Active              *bool                     `protobuf:"varint,30,opt,name=active" json:"active,omitempty"`
	PlayerIds           *string                   `protobuf:"bytes,31,opt,name=player_ids,json=playerIds" json:"player_ids,omitempty"`
	LuaName             *string                   `protobuf:"bytes,32,opt,name=lua_name,json=luaName" json:"lua_name,omitempty"`
	AttackSpeed         *int32                    `protobuf:"varint,33,opt,name=attack_speed,json=attackSpeed" json:"attack_speed,omitempty"`
	AuraOwner           *uint32                   `protobuf:"varint,34,opt,name=aura_owner,json=auraOwner,def=16777215" json:"aura_owner,omitempty"`
	BonusAllStats       *int32                    `protobuf:"varint,35,opt,name=bonus_all_stats,json=bonusAllStats" json:"bonus_all_stats,omitempty"`
	BonusHealth         *int32                    `protobuf:"varint,36,opt,name=bonus_health,json=bonusHealth" json:"bonus_health,omitempty"`
	BonusMana           *int32                    `protobuf:"varint,37,opt,name=bonus_mana,json=bonusMana" json:"bonus_mana,omitempty"`
	CustomEntity        *uint32                   `protobuf:"varint,38,opt,name=custom_entity,json=customEntity,def=16777215" json:"custom_entity,omitempty"`
	AuraWithinRange     *bool                     `protobuf:"varint,39,opt,name=aura_within_range,json=auraWithinRange" json:"aura_within_range,omitempty"`
}

// Default values for CDOTAModifierBuffTableEntry fields.
const (
	Default_CDOTAModifierBuffTableEntry_EntryType    = DOTA_MODIFIER_ENTRY_TYPE_DOTA_MODIFIER_ENTRY_TYPE_ACTIVE
	Default_CDOTAModifierBuffTableEntry_Parent       = uint32(16777215)
	Default_CDOTAModifierBuffTableEntry_Duration     = float32(-1)
	Default_CDOTAModifierBuffTableEntry_Caster       = uint32(16777215)
	Default_CDOTAModifierBuffTableEntry_Ability      = uint32(16777215)
	Default_CDOTAModifierBuffTableEntry_DdAbilityId  = int32(-1)
	Default_CDOTAModifierBuffTableEntry_AuraOwner    = uint32(16777215)
	Default_CDOTAModifierBuffTableEntry_CustomEntity = uint32(16777215)
)

func (x *CDOTAModifierBuffTableEntry) Reset() {
	*x = CDOTAModifierBuffTableEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_modifiers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDOTAModifierBuffTableEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDOTAModifierBuffTableEntry) ProtoMessage() {}

func (x *CDOTAModifierBuffTableEntry) ProtoReflect() protoreflect.Message {
	mi := &file_dota_modifiers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDOTAModifierBuffTableEntry.ProtoReflect.Descriptor instead.
func (*CDOTAModifierBuffTableEntry) Descriptor() ([]byte, []int) {
	return file_dota_modifiers_proto_rawDescGZIP(), []int{0}
}

func (x *CDOTAModifierBuffTableEntry) GetEntryType() DOTA_MODIFIER_ENTRY_TYPE {
	if x != nil && x.EntryType != nil {
		return *x.EntryType
	}
	return Default_CDOTAModifierBuffTableEntry_EntryType
}

func (x *CDOTAModifierBuffTableEntry) GetParent() uint32 {
	if x != nil && x.Parent != nil {
		return *x.Parent
	}
	return Default_CDOTAModifierBuffTableEntry_Parent
}

func (x *CDOTAModifierBuffTableEntry) GetIndex() int32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetSerialNum() int32 {
	if x != nil && x.SerialNum != nil {
		return *x.SerialNum
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetModifierClass() int32 {
	if x != nil && x.ModifierClass != nil {
		return *x.ModifierClass
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetAbilityLevel() int32 {
	if x != nil && x.AbilityLevel != nil {
		return *x.AbilityLevel
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetStackCount() int32 {
	if x != nil && x.StackCount != nil {
		return *x.StackCount
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetCreationTime() float32 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetDuration() float32 {
	if x != nil && x.Duration != nil {
		return *x.Duration
	}
	return Default_CDOTAModifierBuffTableEntry_Duration
}

func (x *CDOTAModifierBuffTableEntry) GetCaster() uint32 {
	if x != nil && x.Caster != nil {
		return *x.Caster
	}
	return Default_CDOTAModifierBuffTableEntry_Caster
}

func (x *CDOTAModifierBuffTableEntry) GetAbility() uint32 {
	if x != nil && x.Ability != nil {
		return *x.Ability
	}
	return Default_CDOTAModifierBuffTableEntry_Ability
}

func (x *CDOTAModifierBuffTableEntry) GetArmor() int32 {
	if x != nil && x.Armor != nil {
		return *x.Armor
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetFadeTime() float32 {
	if x != nil && x.FadeTime != nil {
		return *x.FadeTime
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetSubtle() bool {
	if x != nil && x.Subtle != nil {
		return *x.Subtle
	}
	return false
}

func (x *CDOTAModifierBuffTableEntry) GetChannelTime() float32 {
	if x != nil && x.ChannelTime != nil {
		return *x.ChannelTime
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetVStart() *CMsgVector {
	if x != nil {
		return x.VStart
	}
	return nil
}

func (x *CDOTAModifierBuffTableEntry) GetVEnd() *CMsgVector {
	if x != nil {
		return x.VEnd
	}
	return nil
}

func (x *CDOTAModifierBuffTableEntry) GetPortalLoopAppear() string {
	if x != nil && x.PortalLoopAppear != nil {
		return *x.PortalLoopAppear
	}
	return ""
}

func (x *CDOTAModifierBuffTableEntry) GetPortalLoopDisappear() string {
	if x != nil && x.PortalLoopDisappear != nil {
		return *x.PortalLoopDisappear
	}
	return ""
}

func (x *CDOTAModifierBuffTableEntry) GetHeroLoopAppear() string {
	if x != nil && x.HeroLoopAppear != nil {
		return *x.HeroLoopAppear
	}
	return ""
}

func (x *CDOTAModifierBuffTableEntry) GetHeroLoopDisappear() string {
	if x != nil && x.HeroLoopDisappear != nil {
		return *x.HeroLoopDisappear
	}
	return ""
}

func (x *CDOTAModifierBuffTableEntry) GetMovementSpeed() int32 {
	if x != nil && x.MovementSpeed != nil {
		return *x.MovementSpeed
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetAura() bool {
	if x != nil && x.Aura != nil {
		return *x.Aura
	}
	return false
}

func (x *CDOTAModifierBuffTableEntry) GetActivity() int32 {
	if x != nil && x.Activity != nil {
		return *x.Activity
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetDamage() int32 {
	if x != nil && x.Damage != nil {
		return *x.Damage
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetRange() int32 {
	if x != nil && x.Range != nil {
		return *x.Range
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetDdModifierIndex() int32 {
	if x != nil && x.DdModifierIndex != nil {
		return *x.DdModifierIndex
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetDdAbilityId() int32 {
	if x != nil && x.DdAbilityId != nil {
		return *x.DdAbilityId
	}
	return Default_CDOTAModifierBuffTableEntry_DdAbilityId
}

func (x *CDOTAModifierBuffTableEntry) GetIllusionLabel() string {
	if x != nil && x.IllusionLabel != nil {
		return *x.IllusionLabel
	}
	return ""
}

func (x *CDOTAModifierBuffTableEntry) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *CDOTAModifierBuffTableEntry) GetPlayerIds() string {
	if x != nil && x.PlayerIds != nil {
		return *x.PlayerIds
	}
	return ""
}

func (x *CDOTAModifierBuffTableEntry) GetLuaName() string {
	if x != nil && x.LuaName != nil {
		return *x.LuaName
	}
	return ""
}

func (x *CDOTAModifierBuffTableEntry) GetAttackSpeed() int32 {
	if x != nil && x.AttackSpeed != nil {
		return *x.AttackSpeed
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetAuraOwner() uint32 {
	if x != nil && x.AuraOwner != nil {
		return *x.AuraOwner
	}
	return Default_CDOTAModifierBuffTableEntry_AuraOwner
}

func (x *CDOTAModifierBuffTableEntry) GetBonusAllStats() int32 {
	if x != nil && x.BonusAllStats != nil {
		return *x.BonusAllStats
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetBonusHealth() int32 {
	if x != nil && x.BonusHealth != nil {
		return *x.BonusHealth
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetBonusMana() int32 {
	if x != nil && x.BonusMana != nil {
		return *x.BonusMana
	}
	return 0
}

func (x *CDOTAModifierBuffTableEntry) GetCustomEntity() uint32 {
	if x != nil && x.CustomEntity != nil {
		return *x.CustomEntity
	}
	return Default_CDOTAModifierBuffTableEntry_CustomEntity
}

func (x *CDOTAModifierBuffTableEntry) GetAuraWithinRange() bool {
	if x != nil && x.AuraWithinRange != nil {
		return *x.AuraWithinRange
	}
	return false
}

type CDOTALuaModifierEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifierType     *int32  `protobuf:"varint,1,req,name=modifier_type,json=modifierType" json:"modifier_type,omitempty"`
	ModifierFilename *string `protobuf:"bytes,2,req,name=modifier_filename,json=modifierFilename" json:"modifier_filename,omitempty"`
}

func (x *CDOTALuaModifierEntry) Reset() {
	*x = CDOTALuaModifierEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dota_modifiers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDOTALuaModifierEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDOTALuaModifierEntry) ProtoMessage() {}

func (x *CDOTALuaModifierEntry) ProtoReflect() protoreflect.Message {
	mi := &file_dota_modifiers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDOTALuaModifierEntry.ProtoReflect.Descriptor instead.
func (*CDOTALuaModifierEntry) Descriptor() ([]byte, []int) {
	return file_dota_modifiers_proto_rawDescGZIP(), []int{1}
}

func (x *CDOTALuaModifierEntry) GetModifierType() int32 {
	if x != nil && x.ModifierType != nil {
		return *x.ModifierType
	}
	return 0
}

func (x *CDOTALuaModifierEntry) GetModifierFilename() string {
	if x != nil && x.ModifierFilename != nil {
		return *x.ModifierFilename
	}
	return ""
}

var File_dota_modifiers_proto protoreflect.FileDescriptor

var file_dota_modifiers_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x0b, 0x0a, 0x1b, 0x43, 0x44, 0x4f, 0x54, 0x41, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x42, 0x75, 0x66, 0x66, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x59,
	0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49,
	0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x3a, 0x1f, 0x44,
	0x4f, 0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54,
	0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x52, 0x09,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x3a, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37,
	0x32, 0x31, 0x35, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x04, 0x20, 0x02, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d,
	0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x02, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x06, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x3a, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37, 0x32, 0x31, 0x35, 0x52, 0x06, 0x63, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37, 0x32, 0x31, 0x35, 0x52,
	0x07, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x6d, 0x6f,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x72, 0x6d, 0x6f, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x61, 0x64, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x66, 0x61, 0x64, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x75, 0x62, 0x74, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x75, 0x62,
	0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x76, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x05,
	0x76, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d,
	0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x76, 0x45, 0x6e, 0x64, 0x12, 0x2c,
	0x0a, 0x12, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x61, 0x70,
	0x70, 0x65, 0x61, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x4c, 0x6f, 0x6f, 0x70, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x12, 0x32, 0x0a, 0x15,
	0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x64, 0x69, 0x73, 0x61,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x4c, 0x6f, 0x6f, 0x70, 0x44, 0x69, 0x73, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72,
	0x12, 0x28, 0x0a, 0x10, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x61, 0x70,
	0x70, 0x65, 0x61, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x65, 0x72, 0x6f,
	0x4c, 0x6f, 0x6f, 0x70, 0x41, 0x70, 0x70, 0x65, 0x61, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x65,
	0x72, 0x6f, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x5f, 0x64, 0x69, 0x73, 0x61, 0x70, 0x70, 0x65, 0x61,
	0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68, 0x65, 0x72, 0x6f, 0x4c, 0x6f, 0x6f,
	0x70, 0x44, 0x69, 0x73, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x6f,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x6d, 0x6f, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x65,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x72, 0x61, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x61, 0x75, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x18, 0x18, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e,
	0x67, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x64, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x64, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0d, 0x64,
	0x64, 0x5f, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x05, 0x3a, 0x02, 0x2d, 0x31, 0x52, 0x0b, 0x64, 0x64, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6c, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6c, 0x6c,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x75, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x20, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x75, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x21, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12,
	0x27, 0x0a, 0x0a, 0x61, 0x75, 0x72, 0x61, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x22, 0x20,
	0x01, 0x28, 0x0d, 0x3a, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37, 0x32, 0x31, 0x35, 0x52, 0x09, 0x61,
	0x75, 0x72, 0x61, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x6f, 0x6e, 0x75,
	0x73, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x23, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x18, 0x24, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x18, 0x25, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6f, 0x6e, 0x75, 0x73, 0x4d, 0x61,
	0x6e, 0x61, 0x12, 0x2d, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x26, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x08, 0x31, 0x36, 0x37, 0x37, 0x37,
	0x32, 0x31, 0x35, 0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x72, 0x61, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e,
	0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x75,
	0x72, 0x61, 0x57, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x69, 0x0a,
	0x15, 0x43, 0x44, 0x4f, 0x54, 0x41, 0x4c, 0x75, 0x61, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0c, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x2a, 0x65, 0x0a, 0x18, 0x44, 0x4f, 0x54, 0x41,
	0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x4f, 0x54, 0x41, 0x5f, 0x4d, 0x4f, 0x44,
	0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x4f, 0x54,
	0x41, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x45, 0x4e, 0x54, 0x52, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x02,
}

var (
	file_dota_modifiers_proto_rawDescOnce sync.Once
	file_dota_modifiers_proto_rawDescData = file_dota_modifiers_proto_rawDesc
)

func file_dota_modifiers_proto_rawDescGZIP() []byte {
	file_dota_modifiers_proto_rawDescOnce.Do(func() {
		file_dota_modifiers_proto_rawDescData = protoimpl.X.CompressGZIP(file_dota_modifiers_proto_rawDescData)
	})
	return file_dota_modifiers_proto_rawDescData
}

var file_dota_modifiers_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dota_modifiers_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dota_modifiers_proto_goTypes = []any{
	(DOTA_MODIFIER_ENTRY_TYPE)(0),       // 0: DOTA_MODIFIER_ENTRY_TYPE
	(*CDOTAModifierBuffTableEntry)(nil), // 1: CDOTAModifierBuffTableEntry
	(*CDOTALuaModifierEntry)(nil),       // 2: CDOTALuaModifierEntry
	(*CMsgVector)(nil),                  // 3: CMsgVector
}
var file_dota_modifiers_proto_depIdxs = []int32{
	0, // 0: CDOTAModifierBuffTableEntry.entry_type:type_name -> DOTA_MODIFIER_ENTRY_TYPE
	3, // 1: CDOTAModifierBuffTableEntry.v_start:type_name -> CMsgVector
	3, // 2: CDOTAModifierBuffTableEntry.v_end:type_name -> CMsgVector
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_dota_modifiers_proto_init() }
func file_dota_modifiers_proto_init() {
	if File_dota_modifiers_proto != nil {
		return
	}
	file_networkbasetypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dota_modifiers_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CDOTAModifierBuffTableEntry); i {
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
		file_dota_modifiers_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CDOTALuaModifierEntry); i {
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
			RawDescriptor: file_dota_modifiers_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota_modifiers_proto_goTypes,
		DependencyIndexes: file_dota_modifiers_proto_depIdxs,
		EnumInfos:         file_dota_modifiers_proto_enumTypes,
		MessageInfos:      file_dota_modifiers_proto_msgTypes,
	}.Build()
	File_dota_modifiers_proto = out.File
	file_dota_modifiers_proto_rawDesc = nil
	file_dota_modifiers_proto_goTypes = nil
	file_dota_modifiers_proto_depIdxs = nil
}
