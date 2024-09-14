// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: valveextensions.proto

package deadlock

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EProtoDebugVisiblity int32

const (
	EProtoDebugVisiblity_k_EProtoDebugVisibility_Always      EProtoDebugVisiblity = 0
	EProtoDebugVisiblity_k_EProtoDebugVisibility_Server      EProtoDebugVisiblity = 70
	EProtoDebugVisiblity_k_EProtoDebugVisibility_ValveServer EProtoDebugVisiblity = 80
	EProtoDebugVisiblity_k_EProtoDebugVisibility_GC          EProtoDebugVisiblity = 90
	EProtoDebugVisiblity_k_EProtoDebugVisibility_Never       EProtoDebugVisiblity = 100
)

// Enum value maps for EProtoDebugVisiblity.
var (
	EProtoDebugVisiblity_name = map[int32]string{
		0:   "k_EProtoDebugVisibility_Always",
		70:  "k_EProtoDebugVisibility_Server",
		80:  "k_EProtoDebugVisibility_ValveServer",
		90:  "k_EProtoDebugVisibility_GC",
		100: "k_EProtoDebugVisibility_Never",
	}
	EProtoDebugVisiblity_value = map[string]int32{
		"k_EProtoDebugVisibility_Always":      0,
		"k_EProtoDebugVisibility_Server":      70,
		"k_EProtoDebugVisibility_ValveServer": 80,
		"k_EProtoDebugVisibility_GC":          90,
		"k_EProtoDebugVisibility_Never":       100,
	}
)

func (x EProtoDebugVisiblity) Enum() *EProtoDebugVisiblity {
	p := new(EProtoDebugVisiblity)
	*p = x
	return p
}

func (x EProtoDebugVisiblity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EProtoDebugVisiblity) Descriptor() protoreflect.EnumDescriptor {
	return file_valveextensions_proto_enumTypes[0].Descriptor()
}

func (EProtoDebugVisiblity) Type() protoreflect.EnumType {
	return &file_valveextensions_proto_enumTypes[0]
}

func (x EProtoDebugVisiblity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EProtoDebugVisiblity) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EProtoDebugVisiblity(num)
	return nil
}

// Deprecated: Use EProtoDebugVisiblity.Descriptor instead.
func (EProtoDebugVisiblity) EnumDescriptor() ([]byte, []int) {
	return file_valveextensions_proto_rawDescGZIP(), []int{0}
}

var file_valveextensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61000,
		Name:          "valve_map_field",
		Tag:           "varint,61000,opt,name=valve_map_field,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61001,
		Name:          "valve_map_key",
		Tag:           "varint,61001,opt,name=valve_map_key,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         61002,
		Name:          "diff_encode_field",
		Tag:           "varint,61002,opt,name=diff_encode_field,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61003,
		Name:          "delta_ignore",
		Tag:           "varint,61003,opt,name=delta_ignore,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         61004,
		Name:          "steamml_max_entries",
		Tag:           "varint,61004,opt,name=steamml_max_entries,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         61005,
		Name:          "steamml_is_timestamp",
		Tag:           "varint,61005,opt,name=steamml_is_timestamp,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*uint32)(nil),
		Field:         61006,
		Name:          "steamlearn_count",
		Tag:           "varint,61006,opt,name=steamlearn_count,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*EProtoDebugVisiblity)(nil),
		Field:         61007,
		Name:          "debugprint_visibility",
		Tag:           "varint,61007,opt,name=debugprint_visibility,enum=EProtoDebugVisiblity,def=0",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1000,
		Name:          "schema_friendly_name",
		Tag:           "bytes,1000,opt,name=schema_friendly_name",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1001,
		Name:          "schema_description",
		Tag:           "bytes,1001,opt,name=schema_description",
		Filename:      "valveextensions.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1002,
		Name:          "schema_suppress_enumerator",
		Tag:           "varint,1002,opt,name=schema_suppress_enumerator",
		Filename:      "valveextensions.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool valve_map_field = 61000;
	E_ValveMapField = &file_valveextensions_proto_extTypes[0]
	// optional bool valve_map_key = 61001;
	E_ValveMapKey = &file_valveextensions_proto_extTypes[1]
	// optional int32 diff_encode_field = 61002;
	E_DiffEncodeField = &file_valveextensions_proto_extTypes[2]
	// optional bool delta_ignore = 61003;
	E_DeltaIgnore = &file_valveextensions_proto_extTypes[3]
	// optional uint32 steamml_max_entries = 61004;
	E_SteammlMaxEntries = &file_valveextensions_proto_extTypes[4]
	// optional bool steamml_is_timestamp = 61005;
	E_SteammlIsTimestamp = &file_valveextensions_proto_extTypes[5]
	// optional uint32 steamlearn_count = 61006;
	E_SteamlearnCount = &file_valveextensions_proto_extTypes[6]
	// optional EProtoDebugVisiblity debugprint_visibility = 61007;
	E_DebugprintVisibility = &file_valveextensions_proto_extTypes[7]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string schema_friendly_name = 1000;
	E_SchemaFriendlyName = &file_valveextensions_proto_extTypes[8]
	// optional string schema_description = 1001;
	E_SchemaDescription = &file_valveextensions_proto_extTypes[9]
	// optional bool schema_suppress_enumerator = 1002;
	E_SchemaSuppressEnumerator = &file_valveextensions_proto_extTypes[10]
)

var File_valveextensions_proto protoreflect.FileDescriptor

var file_valveextensions_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x61, 0x6c, 0x76, 0x65, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xca, 0x01, 0x0a, 0x14, 0x45, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x22, 0x0a, 0x1e, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x41, 0x6c,
	0x77, 0x61, 0x79, 0x73, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x10, 0x46, 0x12, 0x27, 0x0a, 0x23, 0x6b, 0x5f,
	0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x56, 0x61, 0x6c, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x10, 0x50, 0x12, 0x1e, 0x0a, 0x1a, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x47,
	0x43, 0x10, 0x5a, 0x12, 0x21, 0x0a, 0x1d, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x4e,
	0x65, 0x76, 0x65, 0x72, 0x10, 0x64, 0x3a, 0x4e, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x76, 0x65, 0x5f,
	0x6d, 0x61, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc8, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0d, 0x76, 0x61, 0x6c, 0x76, 0x65, 0x4d, 0x61,
	0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x4a, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x76, 0x65, 0x5f,
	0x6d, 0x61, 0x70, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc9, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0b, 0x76, 0x61, 0x6c, 0x76, 0x65, 0x4d, 0x61, 0x70, 0x4b,
	0x65, 0x79, 0x3a, 0x4e, 0x0a, 0x11, 0x64, 0x69, 0x66, 0x66, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xca, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x01,
	0x30, 0x52, 0x0f, 0x64, 0x69, 0x66, 0x66, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x3a, 0x49, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5f, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xcb, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65,
	0x52, 0x0b, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x3a, 0x52, 0x0a,
	0x13, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x6c, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xcc, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x30, 0x52, 0x11,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x6c, 0x4d, 0x61, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x3a, 0x58, 0x0a, 0x14, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x6c, 0x5f, 0x69, 0x73, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcd, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x12, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x6c,
	0x49, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x3a, 0x4d, 0x0a, 0x10, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xce,
	0xdc, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x30, 0x52, 0x0f, 0x73, 0x74, 0x65, 0x61, 0x6d,
	0x6c, 0x65, 0x61, 0x72, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x8b, 0x01, 0x0a, 0x15, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xcf, 0xdc, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x45, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x62, 0x75, 0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x69,
	0x74, 0x79, 0x3a, 0x1e, 0x6b, 0x5f, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x44, 0x65, 0x62, 0x75,
	0x67, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x41, 0x6c, 0x77, 0x61,
	0x79, 0x73, 0x52, 0x14, 0x64, 0x65, 0x62, 0x75, 0x67, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x56, 0x69,
	0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x3a, 0x54, 0x0a, 0x14, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x51,
	0x0a, 0x12, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x60, 0x0a, 0x1a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x73, 0x75, 0x70, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72,
}

var (
	file_valveextensions_proto_rawDescOnce sync.Once
	file_valveextensions_proto_rawDescData = file_valveextensions_proto_rawDesc
)

func file_valveextensions_proto_rawDescGZIP() []byte {
	file_valveextensions_proto_rawDescOnce.Do(func() {
		file_valveextensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_valveextensions_proto_rawDescData)
	})
	return file_valveextensions_proto_rawDescData
}

var file_valveextensions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_valveextensions_proto_goTypes = []any{
	(EProtoDebugVisiblity)(0),             // 0: EProtoDebugVisiblity
	(*descriptorpb.FieldOptions)(nil),     // 1: google.protobuf.FieldOptions
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_valveextensions_proto_depIdxs = []int32{
	1,  // 0: valve_map_field:extendee -> google.protobuf.FieldOptions
	1,  // 1: valve_map_key:extendee -> google.protobuf.FieldOptions
	1,  // 2: diff_encode_field:extendee -> google.protobuf.FieldOptions
	1,  // 3: delta_ignore:extendee -> google.protobuf.FieldOptions
	1,  // 4: steamml_max_entries:extendee -> google.protobuf.FieldOptions
	1,  // 5: steamml_is_timestamp:extendee -> google.protobuf.FieldOptions
	1,  // 6: steamlearn_count:extendee -> google.protobuf.FieldOptions
	1,  // 7: debugprint_visibility:extendee -> google.protobuf.FieldOptions
	2,  // 8: schema_friendly_name:extendee -> google.protobuf.EnumValueOptions
	2,  // 9: schema_description:extendee -> google.protobuf.EnumValueOptions
	2,  // 10: schema_suppress_enumerator:extendee -> google.protobuf.EnumValueOptions
	0,  // 11: debugprint_visibility:type_name -> EProtoDebugVisiblity
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	11, // [11:12] is the sub-list for extension type_name
	0,  // [0:11] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_valveextensions_proto_init() }
func file_valveextensions_proto_init() {
	if File_valveextensions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_valveextensions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 11,
			NumServices:   0,
		},
		GoTypes:           file_valveextensions_proto_goTypes,
		DependencyIndexes: file_valveextensions_proto_depIdxs,
		EnumInfos:         file_valveextensions_proto_enumTypes,
		ExtensionInfos:    file_valveextensions_proto_extTypes,
	}.Build()
	File_valveextensions_proto = out.File
	file_valveextensions_proto_rawDesc = nil
	file_valveextensions_proto_goTypes = nil
	file_valveextensions_proto_depIdxs = nil
}
