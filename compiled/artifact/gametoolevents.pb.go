// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: gametoolevents.proto

package artifact

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

type ChangeMapToolEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mapname *string `protobuf:"bytes,1,opt,name=mapname" json:"mapname,omitempty"`
}

func (x *ChangeMapToolEvent) Reset() {
	*x = ChangeMapToolEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeMapToolEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeMapToolEvent) ProtoMessage() {}

func (x *ChangeMapToolEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeMapToolEvent.ProtoReflect.Descriptor instead.
func (*ChangeMapToolEvent) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeMapToolEvent) GetMapname() string {
	if x != nil && x.Mapname != nil {
		return *x.Mapname
	}
	return ""
}

type TraceRayServerToolEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *CMsgVector `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   *CMsgVector `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (x *TraceRayServerToolEvent) Reset() {
	*x = TraceRayServerToolEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceRayServerToolEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceRayServerToolEvent) ProtoMessage() {}

func (x *TraceRayServerToolEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceRayServerToolEvent.ProtoReflect.Descriptor instead.
func (*TraceRayServerToolEvent) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{1}
}

func (x *TraceRayServerToolEvent) GetStart() *CMsgVector {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *TraceRayServerToolEvent) GetEnd() *CMsgVector {
	if x != nil {
		return x.End
	}
	return nil
}

type ToolTraceRayResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hit      *bool       `protobuf:"varint,1,opt,name=hit" json:"hit,omitempty"`
	Impact   *CMsgVector `protobuf:"bytes,2,opt,name=impact" json:"impact,omitempty"`
	Normal   *CMsgVector `protobuf:"bytes,3,opt,name=normal" json:"normal,omitempty"`
	Distance *float32    `protobuf:"fixed32,4,opt,name=distance" json:"distance,omitempty"`
	Fraction *float32    `protobuf:"fixed32,5,opt,name=fraction" json:"fraction,omitempty"`
	Ehandle  *int32      `protobuf:"varint,6,opt,name=ehandle" json:"ehandle,omitempty"`
}

func (x *ToolTraceRayResult) Reset() {
	*x = ToolTraceRayResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolTraceRayResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolTraceRayResult) ProtoMessage() {}

func (x *ToolTraceRayResult) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolTraceRayResult.ProtoReflect.Descriptor instead.
func (*ToolTraceRayResult) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{2}
}

func (x *ToolTraceRayResult) GetHit() bool {
	if x != nil && x.Hit != nil {
		return *x.Hit
	}
	return false
}

func (x *ToolTraceRayResult) GetImpact() *CMsgVector {
	if x != nil {
		return x.Impact
	}
	return nil
}

func (x *ToolTraceRayResult) GetNormal() *CMsgVector {
	if x != nil {
		return x.Normal
	}
	return nil
}

func (x *ToolTraceRayResult) GetDistance() float32 {
	if x != nil && x.Distance != nil {
		return *x.Distance
	}
	return 0
}

func (x *ToolTraceRayResult) GetFraction() float32 {
	if x != nil && x.Fraction != nil {
		return *x.Fraction
	}
	return 0
}

func (x *ToolTraceRayResult) GetEhandle() int32 {
	if x != nil && x.Ehandle != nil {
		return *x.Ehandle
	}
	return 0
}

type SpawnEntityToolEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityKeyvalues  []byte `protobuf:"bytes,1,opt,name=entity_keyvalues,json=entityKeyvalues" json:"entity_keyvalues,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
}

func (x *SpawnEntityToolEvent) Reset() {
	*x = SpawnEntityToolEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpawnEntityToolEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpawnEntityToolEvent) ProtoMessage() {}

func (x *SpawnEntityToolEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpawnEntityToolEvent.ProtoReflect.Descriptor instead.
func (*SpawnEntityToolEvent) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{3}
}

func (x *SpawnEntityToolEvent) GetEntityKeyvalues() []byte {
	if x != nil {
		return x.EntityKeyvalues
	}
	return nil
}

func (x *SpawnEntityToolEvent) GetClientsideentity() bool {
	if x != nil && x.Clientsideentity != nil {
		return *x.Clientsideentity
	}
	return false
}

type SpawnEntityToolEventResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ehandle *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
}

func (x *SpawnEntityToolEventResult) Reset() {
	*x = SpawnEntityToolEventResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpawnEntityToolEventResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpawnEntityToolEventResult) ProtoMessage() {}

func (x *SpawnEntityToolEventResult) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpawnEntityToolEventResult.ProtoReflect.Descriptor instead.
func (*SpawnEntityToolEventResult) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{4}
}

func (x *SpawnEntityToolEventResult) GetEhandle() int32 {
	if x != nil && x.Ehandle != nil {
		return *x.Ehandle
	}
	return 0
}

type DestroyEntityToolEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ehandle *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
}

func (x *DestroyEntityToolEvent) Reset() {
	*x = DestroyEntityToolEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyEntityToolEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyEntityToolEvent) ProtoMessage() {}

func (x *DestroyEntityToolEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyEntityToolEvent.ProtoReflect.Descriptor instead.
func (*DestroyEntityToolEvent) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{5}
}

func (x *DestroyEntityToolEvent) GetEhandle() int32 {
	if x != nil && x.Ehandle != nil {
		return *x.Ehandle
	}
	return 0
}

type DestroyAllEntitiesToolEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DestroyAllEntitiesToolEvent) Reset() {
	*x = DestroyAllEntitiesToolEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestroyAllEntitiesToolEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestroyAllEntitiesToolEvent) ProtoMessage() {}

func (x *DestroyAllEntitiesToolEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestroyAllEntitiesToolEvent.ProtoReflect.Descriptor instead.
func (*DestroyAllEntitiesToolEvent) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{6}
}

type RestartMapToolEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RestartMapToolEvent) Reset() {
	*x = RestartMapToolEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartMapToolEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartMapToolEvent) ProtoMessage() {}

func (x *RestartMapToolEvent) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartMapToolEvent.ProtoReflect.Descriptor instead.
func (*RestartMapToolEvent) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{7}
}

type ToolEvent_GetEntityInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
}

func (x *ToolEvent_GetEntityInfo) Reset() {
	*x = ToolEvent_GetEntityInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolEvent_GetEntityInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEvent_GetEntityInfo) ProtoMessage() {}

func (x *ToolEvent_GetEntityInfo) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEvent_GetEntityInfo.ProtoReflect.Descriptor instead.
func (*ToolEvent_GetEntityInfo) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{8}
}

func (x *ToolEvent_GetEntityInfo) GetEhandle() int32 {
	if x != nil && x.Ehandle != nil {
		return *x.Ehandle
	}
	return 0
}

func (x *ToolEvent_GetEntityInfo) GetClientsideentity() bool {
	if x != nil && x.Clientsideentity != nil {
		return *x.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cppclass  *string     `protobuf:"bytes,1,opt,name=cppclass,def=shithead" json:"cppclass,omitempty"`
	Classname *string     `protobuf:"bytes,2,opt,name=classname" json:"classname,omitempty"`
	Name      *string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Origin    *CMsgVector `protobuf:"bytes,4,opt,name=origin" json:"origin,omitempty"`
	Mins      *CMsgVector `protobuf:"bytes,5,opt,name=mins" json:"mins,omitempty"`
	Maxs      *CMsgVector `protobuf:"bytes,6,opt,name=maxs" json:"maxs,omitempty"`
}

// Default values for ToolEvent_GetEntityInfoResult fields.
const (
	Default_ToolEvent_GetEntityInfoResult_Cppclass = string("shithead")
)

func (x *ToolEvent_GetEntityInfoResult) Reset() {
	*x = ToolEvent_GetEntityInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolEvent_GetEntityInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEvent_GetEntityInfoResult) ProtoMessage() {}

func (x *ToolEvent_GetEntityInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEvent_GetEntityInfoResult.ProtoReflect.Descriptor instead.
func (*ToolEvent_GetEntityInfoResult) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{9}
}

func (x *ToolEvent_GetEntityInfoResult) GetCppclass() string {
	if x != nil && x.Cppclass != nil {
		return *x.Cppclass
	}
	return Default_ToolEvent_GetEntityInfoResult_Cppclass
}

func (x *ToolEvent_GetEntityInfoResult) GetClassname() string {
	if x != nil && x.Classname != nil {
		return *x.Classname
	}
	return ""
}

func (x *ToolEvent_GetEntityInfoResult) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ToolEvent_GetEntityInfoResult) GetOrigin() *CMsgVector {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *ToolEvent_GetEntityInfoResult) GetMins() *CMsgVector {
	if x != nil {
		return x.Mins
	}
	return nil
}

func (x *ToolEvent_GetEntityInfoResult) GetMaxs() *CMsgVector {
	if x != nil {
		return x.Maxs
	}
	return nil
}

type ToolEvent_GetEntityInputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ehandle          *int32 `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool  `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
}

func (x *ToolEvent_GetEntityInputs) Reset() {
	*x = ToolEvent_GetEntityInputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolEvent_GetEntityInputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEvent_GetEntityInputs) ProtoMessage() {}

func (x *ToolEvent_GetEntityInputs) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEvent_GetEntityInputs.ProtoReflect.Descriptor instead.
func (*ToolEvent_GetEntityInputs) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{10}
}

func (x *ToolEvent_GetEntityInputs) GetEhandle() int32 {
	if x != nil && x.Ehandle != nil {
		return *x.Ehandle
	}
	return 0
}

func (x *ToolEvent_GetEntityInputs) GetClientsideentity() bool {
	if x != nil && x.Clientsideentity != nil {
		return *x.Clientsideentity
	}
	return false
}

type ToolEvent_GetEntityInputsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputList []string `protobuf:"bytes,1,rep,name=input_list,json=inputList" json:"input_list,omitempty"`
}

func (x *ToolEvent_GetEntityInputsResult) Reset() {
	*x = ToolEvent_GetEntityInputsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolEvent_GetEntityInputsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEvent_GetEntityInputsResult) ProtoMessage() {}

func (x *ToolEvent_GetEntityInputsResult) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEvent_GetEntityInputsResult.ProtoReflect.Descriptor instead.
func (*ToolEvent_GetEntityInputsResult) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{11}
}

func (x *ToolEvent_GetEntityInputsResult) GetInputList() []string {
	if x != nil {
		return x.InputList
	}
	return nil
}

type ToolEvent_FireEntityInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ehandle          *int32  `protobuf:"varint,1,opt,name=ehandle" json:"ehandle,omitempty"`
	Clientsideentity *bool   `protobuf:"varint,2,opt,name=clientsideentity" json:"clientsideentity,omitempty"`
	InputName        *string `protobuf:"bytes,3,opt,name=input_name,json=inputName" json:"input_name,omitempty"`
	InputParam       *string `protobuf:"bytes,4,opt,name=input_param,json=inputParam" json:"input_param,omitempty"`
}

func (x *ToolEvent_FireEntityInput) Reset() {
	*x = ToolEvent_FireEntityInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolEvent_FireEntityInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEvent_FireEntityInput) ProtoMessage() {}

func (x *ToolEvent_FireEntityInput) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEvent_FireEntityInput.ProtoReflect.Descriptor instead.
func (*ToolEvent_FireEntityInput) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{12}
}

func (x *ToolEvent_FireEntityInput) GetEhandle() int32 {
	if x != nil && x.Ehandle != nil {
		return *x.Ehandle
	}
	return 0
}

func (x *ToolEvent_FireEntityInput) GetClientsideentity() bool {
	if x != nil && x.Clientsideentity != nil {
		return *x.Clientsideentity
	}
	return false
}

func (x *ToolEvent_FireEntityInput) GetInputName() string {
	if x != nil && x.InputName != nil {
		return *x.InputName
	}
	return ""
}

func (x *ToolEvent_FireEntityInput) GetInputParam() string {
	if x != nil && x.InputParam != nil {
		return *x.InputParam
	}
	return ""
}

type ToolEvent_SFMRecordingStateChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isrecording *bool `protobuf:"varint,1,opt,name=isrecording" json:"isrecording,omitempty"`
}

func (x *ToolEvent_SFMRecordingStateChanged) Reset() {
	*x = ToolEvent_SFMRecordingStateChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolEvent_SFMRecordingStateChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEvent_SFMRecordingStateChanged) ProtoMessage() {}

func (x *ToolEvent_SFMRecordingStateChanged) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEvent_SFMRecordingStateChanged.ProtoReflect.Descriptor instead.
func (*ToolEvent_SFMRecordingStateChanged) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{13}
}

func (x *ToolEvent_SFMRecordingStateChanged) GetIsrecording() bool {
	if x != nil && x.Isrecording != nil {
		return *x.Isrecording
	}
	return false
}

type ToolEvent_SFMToolActiveStateChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isactive *bool `protobuf:"varint,1,opt,name=isactive" json:"isactive,omitempty"`
}

func (x *ToolEvent_SFMToolActiveStateChanged) Reset() {
	*x = ToolEvent_SFMToolActiveStateChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gametoolevents_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ToolEvent_SFMToolActiveStateChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEvent_SFMToolActiveStateChanged) ProtoMessage() {}

func (x *ToolEvent_SFMToolActiveStateChanged) ProtoReflect() protoreflect.Message {
	mi := &file_gametoolevents_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEvent_SFMToolActiveStateChanged.ProtoReflect.Descriptor instead.
func (*ToolEvent_SFMToolActiveStateChanged) Descriptor() ([]byte, []int) {
	return file_gametoolevents_proto_rawDescGZIP(), []int{14}
}

func (x *ToolEvent_SFMToolActiveStateChanged) GetIsactive() bool {
	if x != nil && x.Isactive != nil {
		return *x.Isactive
	}
	return false
}

var File_gametoolevents_proto protoreflect.FileDescriptor

var file_gametoolevents_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x6f, 0x6f, 0x6c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e,
	0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5b,
	0x0a, 0x17, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x03,
	0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x12,
	0x54, 0x6f, 0x6f, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x61, 0x79, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x68, 0x69, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x06, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x6e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67,
	0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x66, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x22, 0x6d, 0x0a, 0x14, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64,
	0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x36, 0x0a, 0x1a, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f,
	0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x44,
	0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x61, 0x70, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x5f, 0x0a, 0x17, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x69, 0x64, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0xde, 0x01, 0x0a, 0x1d, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x24, 0x0a, 0x08, 0x63, 0x70, 0x70, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x08, 0x73, 0x68, 0x69, 0x74, 0x68, 0x65, 0x61, 0x64,
	0x52, 0x08, 0x63, 0x70, 0x70, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43,
	0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6d, 0x69,
	0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x61, 0x78, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x04, 0x6d,
	0x61, 0x78, 0x73, 0x22, 0x61, 0x0a, 0x19, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x40, 0x0a, 0x1f, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa1, 0x01, 0x0a, 0x19, 0x54, 0x6f, 0x6f,
	0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x46, 0x69, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x69, 0x64, 0x65, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x46, 0x0a, 0x22,
	0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x53, 0x46, 0x4d, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x22, 0x41, 0x0a, 0x23, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x53, 0x46, 0x4d, 0x54, 0x6f, 0x6f, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x03, 0x80, 0x01, 0x00,
}

var (
	file_gametoolevents_proto_rawDescOnce sync.Once
	file_gametoolevents_proto_rawDescData = file_gametoolevents_proto_rawDesc
)

func file_gametoolevents_proto_rawDescGZIP() []byte {
	file_gametoolevents_proto_rawDescOnce.Do(func() {
		file_gametoolevents_proto_rawDescData = protoimpl.X.CompressGZIP(file_gametoolevents_proto_rawDescData)
	})
	return file_gametoolevents_proto_rawDescData
}

var file_gametoolevents_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_gametoolevents_proto_goTypes = []any{
	(*ChangeMapToolEvent)(nil),                  // 0: ChangeMapToolEvent
	(*TraceRayServerToolEvent)(nil),             // 1: TraceRayServerToolEvent
	(*ToolTraceRayResult)(nil),                  // 2: ToolTraceRayResult
	(*SpawnEntityToolEvent)(nil),                // 3: SpawnEntityToolEvent
	(*SpawnEntityToolEventResult)(nil),          // 4: SpawnEntityToolEventResult
	(*DestroyEntityToolEvent)(nil),              // 5: DestroyEntityToolEvent
	(*DestroyAllEntitiesToolEvent)(nil),         // 6: DestroyAllEntitiesToolEvent
	(*RestartMapToolEvent)(nil),                 // 7: RestartMapToolEvent
	(*ToolEvent_GetEntityInfo)(nil),             // 8: ToolEvent_GetEntityInfo
	(*ToolEvent_GetEntityInfoResult)(nil),       // 9: ToolEvent_GetEntityInfoResult
	(*ToolEvent_GetEntityInputs)(nil),           // 10: ToolEvent_GetEntityInputs
	(*ToolEvent_GetEntityInputsResult)(nil),     // 11: ToolEvent_GetEntityInputsResult
	(*ToolEvent_FireEntityInput)(nil),           // 12: ToolEvent_FireEntityInput
	(*ToolEvent_SFMRecordingStateChanged)(nil),  // 13: ToolEvent_SFMRecordingStateChanged
	(*ToolEvent_SFMToolActiveStateChanged)(nil), // 14: ToolEvent_SFMToolActiveStateChanged
	(*CMsgVector)(nil),                          // 15: CMsgVector
}
var file_gametoolevents_proto_depIdxs = []int32{
	15, // 0: TraceRayServerToolEvent.start:type_name -> CMsgVector
	15, // 1: TraceRayServerToolEvent.end:type_name -> CMsgVector
	15, // 2: ToolTraceRayResult.impact:type_name -> CMsgVector
	15, // 3: ToolTraceRayResult.normal:type_name -> CMsgVector
	15, // 4: ToolEvent_GetEntityInfoResult.origin:type_name -> CMsgVector
	15, // 5: ToolEvent_GetEntityInfoResult.mins:type_name -> CMsgVector
	15, // 6: ToolEvent_GetEntityInfoResult.maxs:type_name -> CMsgVector
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_gametoolevents_proto_init() }
func file_gametoolevents_proto_init() {
	if File_gametoolevents_proto != nil {
		return
	}
	file_networkbasetypes_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gametoolevents_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChangeMapToolEvent); i {
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
		file_gametoolevents_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TraceRayServerToolEvent); i {
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
		file_gametoolevents_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ToolTraceRayResult); i {
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
		file_gametoolevents_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SpawnEntityToolEvent); i {
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
		file_gametoolevents_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SpawnEntityToolEventResult); i {
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
		file_gametoolevents_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DestroyEntityToolEvent); i {
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
		file_gametoolevents_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DestroyAllEntitiesToolEvent); i {
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
		file_gametoolevents_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RestartMapToolEvent); i {
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
		file_gametoolevents_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ToolEvent_GetEntityInfo); i {
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
		file_gametoolevents_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ToolEvent_GetEntityInfoResult); i {
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
		file_gametoolevents_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ToolEvent_GetEntityInputs); i {
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
		file_gametoolevents_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ToolEvent_GetEntityInputsResult); i {
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
		file_gametoolevents_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*ToolEvent_FireEntityInput); i {
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
		file_gametoolevents_proto_msgTypes[13].Exporter = func(v any, i int) any {
			switch v := v.(*ToolEvent_SFMRecordingStateChanged); i {
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
		file_gametoolevents_proto_msgTypes[14].Exporter = func(v any, i int) any {
			switch v := v.(*ToolEvent_SFMToolActiveStateChanged); i {
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
			RawDescriptor: file_gametoolevents_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gametoolevents_proto_goTypes,
		DependencyIndexes: file_gametoolevents_proto_depIdxs,
		MessageInfos:      file_gametoolevents_proto_msgTypes,
	}.Build()
	File_gametoolevents_proto = out.File
	file_gametoolevents_proto_rawDesc = nil
	file_gametoolevents_proto_goTypes = nil
	file_gametoolevents_proto_depIdxs = nil
}