// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: service_clan.proto

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

type CClan_GetDraftAndRecentPartnerEventSnippet_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steamid         *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	RtimeOldestDate *uint32 `protobuf:"varint,2,opt,name=rtime_oldest_date,json=rtimeOldestDate" json:"rtime_oldest_date,omitempty"`
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Request) Reset() {
	*x = CClan_GetDraftAndRecentPartnerEventSnippet_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_GetDraftAndRecentPartnerEventSnippet_Request) ProtoMessage() {}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_GetDraftAndRecentPartnerEventSnippet_Request.ProtoReflect.Descriptor instead.
func (*CClan_GetDraftAndRecentPartnerEventSnippet_Request) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{0}
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Request) GetSteamid() uint64 {
	if x != nil && x.Steamid != nil {
		return *x.Steamid
	}
	return 0
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Request) GetRtimeOldestDate() uint32 {
	if x != nil && x.RtimeOldestDate != nil {
		return *x.RtimeOldestDate
	}
	return 0
}

type CClan_GetDraftAndRecentPartnerEventSnippet_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snippets []*CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData `protobuf:"bytes,1,rep,name=snippets" json:"snippets,omitempty"`
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response) Reset() {
	*x = CClan_GetDraftAndRecentPartnerEventSnippet_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_GetDraftAndRecentPartnerEventSnippet_Response) ProtoMessage() {}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_GetDraftAndRecentPartnerEventSnippet_Response.ProtoReflect.Descriptor instead.
func (*CClan_GetDraftAndRecentPartnerEventSnippet_Response) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{1}
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response) GetSnippets() []*CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData {
	if x != nil {
		return x.Snippets
	}
	return nil
}

type CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gid              *uint64 `protobuf:"fixed64,1,opt,name=gid" json:"gid,omitempty"`
	AnnouncementGid  *uint64 `protobuf:"fixed64,2,opt,name=announcement_gid,json=announcementGid" json:"announcement_gid,omitempty"`
	Hidden           *bool   `protobuf:"varint,3,opt,name=hidden" json:"hidden,omitempty"`
	Published        *bool   `protobuf:"varint,4,opt,name=published" json:"published,omitempty"`
	Rtime32StartTime *uint32 `protobuf:"varint,5,opt,name=rtime32_start_time,json=rtime32StartTime" json:"rtime32_start_time,omitempty"`
	EventName        *string `protobuf:"bytes,6,opt,name=event_name,json=eventName" json:"event_name,omitempty"`
	EventType        *int32  `protobuf:"varint,7,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) Reset() {
	*x = CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) ProtoMessage() {}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData.ProtoReflect.Descriptor instead.
func (*CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{2}
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) GetGid() uint64 {
	if x != nil && x.Gid != nil {
		return *x.Gid
	}
	return 0
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) GetAnnouncementGid() uint64 {
	if x != nil && x.AnnouncementGid != nil {
		return *x.AnnouncementGid
	}
	return 0
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) GetHidden() bool {
	if x != nil && x.Hidden != nil {
		return *x.Hidden
	}
	return false
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) GetPublished() bool {
	if x != nil && x.Published != nil {
		return *x.Published
	}
	return false
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) GetRtime32StartTime() uint32 {
	if x != nil && x.Rtime32StartTime != nil {
		return *x.Rtime32StartTime
	}
	return 0
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) GetEventName() string {
	if x != nil && x.EventName != nil {
		return *x.EventName
	}
	return ""
}

func (x *CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

type CClan_GetPartnerEventsByBuildIDRange_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests []*CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange `protobuf:"bytes,1,rep,name=requests" json:"requests,omitempty"`
	Cursor   *string                                                        `protobuf:"bytes,2,opt,name=cursor" json:"cursor,omitempty"`
	Count    *uint32                                                        `protobuf:"varint,3,opt,name=count,def=100" json:"count,omitempty"`
}

// Default values for CClan_GetPartnerEventsByBuildIDRange_Request fields.
const (
	Default_CClan_GetPartnerEventsByBuildIDRange_Request_Count = uint32(100)
)

func (x *CClan_GetPartnerEventsByBuildIDRange_Request) Reset() {
	*x = CClan_GetPartnerEventsByBuildIDRange_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_GetPartnerEventsByBuildIDRange_Request) ProtoMessage() {}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_GetPartnerEventsByBuildIDRange_Request.ProtoReflect.Descriptor instead.
func (*CClan_GetPartnerEventsByBuildIDRange_Request) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{3}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request) GetRequests() []*CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request) GetCursor() string {
	if x != nil && x.Cursor != nil {
		return *x.Cursor
	}
	return ""
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return Default_CClan_GetPartnerEventsByBuildIDRange_Request_Count
}

type CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid        *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	StartBuildId *uint32 `protobuf:"varint,2,opt,name=start_build_id,json=startBuildId" json:"start_build_id,omitempty"`
	EndBuildId   *uint32 `protobuf:"varint,3,opt,name=end_build_id,json=endBuildId" json:"end_build_id,omitempty"`
	Branch       *string `protobuf:"bytes,4,opt,name=branch" json:"branch,omitempty"`
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) Reset() {
	*x = CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) ProtoMessage() {}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange.ProtoReflect.Descriptor instead.
func (*CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{4}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) GetStartBuildId() uint32 {
	if x != nil && x.StartBuildId != nil {
		return *x.StartBuildId
	}
	return 0
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) GetEndBuildId() uint32 {
	if x != nil && x.EndBuildId != nil {
		return *x.EndBuildId
	}
	return 0
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange) GetBranch() string {
	if x != nil && x.Branch != nil {
		return *x.Branch
	}
	return ""
}

type CClan_GetPartnerEventsByBuildIDRange_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches         []*CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc `protobuf:"bytes,1,rep,name=matches" json:"matches,omitempty"`
	NumTotalResults *uint32                                                         `protobuf:"varint,2,opt,name=num_total_results,json=numTotalResults" json:"num_total_results,omitempty"`
	NextCursor      *string                                                         `protobuf:"bytes,3,opt,name=next_cursor,json=nextCursor" json:"next_cursor,omitempty"`
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response) Reset() {
	*x = CClan_GetPartnerEventsByBuildIDRange_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_GetPartnerEventsByBuildIDRange_Response) ProtoMessage() {}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_GetPartnerEventsByBuildIDRange_Response.ProtoReflect.Descriptor instead.
func (*CClan_GetPartnerEventsByBuildIDRange_Response) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{5}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response) GetMatches() []*CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc {
	if x != nil {
		return x.Matches
	}
	return nil
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response) GetNumTotalResults() uint32 {
	if x != nil && x.NumTotalResults != nil {
		return *x.NumTotalResults
	}
	return 0
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response) GetNextCursor() string {
	if x != nil && x.NextCursor != nil {
		return *x.NextCursor
	}
	return ""
}

type CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid         *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	BuildId       *uint32 `protobuf:"varint,2,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	Branch        *string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	ClanEventGid  *uint64 `protobuf:"fixed64,4,opt,name=clan_event_gid,json=clanEventGid" json:"clan_event_gid,omitempty"`
	ClanAccountId *uint32 `protobuf:"varint,5,opt,name=clan_account_id,json=clanAccountId" json:"clan_account_id,omitempty"`
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) Reset() {
	*x = CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) ProtoMessage() {}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc.ProtoReflect.Descriptor instead.
func (*CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{6}
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) GetBuildId() uint32 {
	if x != nil && x.BuildId != nil {
		return *x.BuildId
	}
	return 0
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) GetBranch() string {
	if x != nil && x.Branch != nil {
		return *x.Branch
	}
	return ""
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) GetClanEventGid() uint64 {
	if x != nil && x.ClanEventGid != nil {
		return *x.ClanEventGid
	}
	return 0
}

func (x *CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc) GetClanAccountId() uint32 {
	if x != nil && x.ClanAccountId != nil {
		return *x.ClanAccountId
	}
	return 0
}

type CClan_RespondToClanInvite_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steamid *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	Accept  *bool   `protobuf:"varint,2,opt,name=accept" json:"accept,omitempty"`
}

func (x *CClan_RespondToClanInvite_Request) Reset() {
	*x = CClan_RespondToClanInvite_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_RespondToClanInvite_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_RespondToClanInvite_Request) ProtoMessage() {}

func (x *CClan_RespondToClanInvite_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_RespondToClanInvite_Request.ProtoReflect.Descriptor instead.
func (*CClan_RespondToClanInvite_Request) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{7}
}

func (x *CClan_RespondToClanInvite_Request) GetSteamid() uint64 {
	if x != nil && x.Steamid != nil {
		return *x.Steamid
	}
	return 0
}

func (x *CClan_RespondToClanInvite_Request) GetAccept() bool {
	if x != nil && x.Accept != nil {
		return *x.Accept
	}
	return false
}

type CClan_RespondToClanInvite_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CClan_RespondToClanInvite_Response) Reset() {
	*x = CClan_RespondToClanInvite_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_clan_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CClan_RespondToClanInvite_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CClan_RespondToClanInvite_Response) ProtoMessage() {}

func (x *CClan_RespondToClanInvite_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_clan_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CClan_RespondToClanInvite_Response.ProtoReflect.Descriptor instead.
func (*CClan_RespondToClanInvite_Response) Descriptor() ([]byte, []int) {
	return file_service_clan_proto_rawDescGZIP(), []int{8}
}

var File_service_clan_proto protoreflect.FileDescriptor

var file_service_clan_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x32, 0x43, 0x43, 0x6c, 0x61, 0x6e,
	0x5f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61, 0x66, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07,
	0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x72, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x4f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x33, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65,
	0x74, 0x44, 0x72, 0x61, 0x66, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x08, 0x73,
	0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x46, 0x2e,
	0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61, 0x66, 0x74, 0x41, 0x6e,
	0x64, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x43, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x73, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x73, 0x22,
	0xb0, 0x02, 0x0a, 0x45, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61,
	0x66, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x43, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x61,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x47, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x72, 0x74, 0x69, 0x6d, 0x65, 0x33, 0x32, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x33,
	0x32, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0x82,
	0xb5, 0x18, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xbb, 0x01, 0x0a, 0x2c, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65,
	0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79,
	0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x3a, 0x03, 0x31, 0x30, 0x30, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xb3, 0x01, 0x0a, 0x3b, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c,
	0x65, 0x6e, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x22, 0xd5, 0x01, 0x0a, 0x2d, 0x43, 0x43, 0x6c, 0x61, 0x6e,
	0x5f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x42, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x43, 0x43, 0x6c, 0x61,
	0x6e, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4e,
	0x6f, 0x74, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6e, 0x75,
	0x6d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0xd5,
	0x01, 0x0a, 0x3c, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x44, 0x65, 0x73, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x6e,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x0c, 0x63, 0x6c, 0x61, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0f, 0x63, 0x6c, 0x61, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6c, 0x61, 0x6e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x21, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x43, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x22, 0x24, 0x0a,
	0x22, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x54, 0x6f,
	0x43, 0x6c, 0x61, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xfb, 0x02, 0x0a, 0x04, 0x43, 0x6c, 0x61, 0x6e, 0x12, 0x91, 0x01, 0x0a,
	0x24, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61, 0x66, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6e,
	0x69, 0x70, 0x70, 0x65, 0x74, 0x12, 0x33, 0x2e, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65,
	0x74, 0x44, 0x72, 0x61, 0x66, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x6e, 0x69, 0x70, 0x70,
	0x65, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x43, 0x43, 0x6c,
	0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x44, 0x72, 0x61, 0x66, 0x74, 0x41, 0x6e, 0x64, 0x52, 0x65,
	0x63, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x6e, 0x69, 0x70, 0x70, 0x65, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x7f, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x2d, 0x2e, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x43, 0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x49, 0x44, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5e, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x43, 0x6c,
	0x61, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x43, 0x43, 0x6c, 0x61, 0x6e,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x43, 0x6c, 0x61, 0x6e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x43,
	0x43, 0x6c, 0x61, 0x6e, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x43, 0x6c,
	0x61, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65,
}

var (
	file_service_clan_proto_rawDescOnce sync.Once
	file_service_clan_proto_rawDescData = file_service_clan_proto_rawDesc
)

func file_service_clan_proto_rawDescGZIP() []byte {
	file_service_clan_proto_rawDescOnce.Do(func() {
		file_service_clan_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_clan_proto_rawDescData)
	})
	return file_service_clan_proto_rawDescData
}

var file_service_clan_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_clan_proto_goTypes = []any{
	(*CClan_GetDraftAndRecentPartnerEventSnippet_Request)(nil),                    // 0: CClan_GetDraftAndRecentPartnerEventSnippet_Request
	(*CClan_GetDraftAndRecentPartnerEventSnippet_Response)(nil),                   // 1: CClan_GetDraftAndRecentPartnerEventSnippet_Response
	(*CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData)(nil), // 2: CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData
	(*CClan_GetPartnerEventsByBuildIDRange_Request)(nil),                          // 3: CClan_GetPartnerEventsByBuildIDRange_Request
	(*CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange)(nil),           // 4: CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange
	(*CClan_GetPartnerEventsByBuildIDRange_Response)(nil),                         // 5: CClan_GetPartnerEventsByBuildIDRange_Response
	(*CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc)(nil),          // 6: CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc
	(*CClan_RespondToClanInvite_Request)(nil),                                     // 7: CClan_RespondToClanInvite_Request
	(*CClan_RespondToClanInvite_Response)(nil),                                    // 8: CClan_RespondToClanInvite_Response
}
var file_service_clan_proto_depIdxs = []int32{
	2, // 0: CClan_GetDraftAndRecentPartnerEventSnippet_Response.snippets:type_name -> CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData
	4, // 1: CClan_GetPartnerEventsByBuildIDRange_Request.requests:type_name -> CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange
	6, // 2: CClan_GetPartnerEventsByBuildIDRange_Response.matches:type_name -> CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc
	0, // 3: Clan.GetDraftAndRecentPartnerEventSnippet:input_type -> CClan_GetDraftAndRecentPartnerEventSnippet_Request
	3, // 4: Clan.GetPartnerEventsByBuildIDRange:input_type -> CClan_GetPartnerEventsByBuildIDRange_Request
	7, // 5: Clan.RespondToClanInvite:input_type -> CClan_RespondToClanInvite_Request
	1, // 6: Clan.GetDraftAndRecentPartnerEventSnippet:output_type -> CClan_GetDraftAndRecentPartnerEventSnippet_Response
	5, // 7: Clan.GetPartnerEventsByBuildIDRange:output_type -> CClan_GetPartnerEventsByBuildIDRange_Response
	8, // 8: Clan.RespondToClanInvite:output_type -> CClan_RespondToClanInvite_Response
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_clan_proto_init() }
func file_service_clan_proto_init() {
	if File_service_clan_proto != nil {
		return
	}
	file_common_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_clan_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_GetDraftAndRecentPartnerEventSnippet_Request); i {
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
		file_service_clan_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_GetDraftAndRecentPartnerEventSnippet_Response); i {
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
		file_service_clan_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_GetDraftAndRecentPartnerEventSnippet_Response_CEventSnippetData); i {
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
		file_service_clan_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_GetPartnerEventsByBuildIDRange_Request); i {
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
		file_service_clan_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_GetPartnerEventsByBuildIDRange_Request_PatchNoteRange); i {
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
		file_service_clan_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_GetPartnerEventsByBuildIDRange_Response); i {
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
		file_service_clan_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_GetPartnerEventsByBuildIDRange_Response_PatchNotesDesc); i {
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
		file_service_clan_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_RespondToClanInvite_Request); i {
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
		file_service_clan_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*CClan_RespondToClanInvite_Response); i {
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
			RawDescriptor: file_service_clan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_clan_proto_goTypes,
		DependencyIndexes: file_service_clan_proto_depIdxs,
		MessageInfos:      file_service_clan_proto_msgTypes,
	}.Build()
	File_service_clan_proto = out.File
	file_service_clan_proto_rawDesc = nil
	file_service_clan_proto_goTypes = nil
	file_service_clan_proto_depIdxs = nil
}
