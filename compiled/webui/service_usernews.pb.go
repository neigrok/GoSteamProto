// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: service_usernews.proto

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

type CUserNews_Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eventtype          *uint32  `protobuf:"varint,1,opt,name=eventtype" json:"eventtype,omitempty"`
	Eventtime          *uint32  `protobuf:"varint,2,opt,name=eventtime" json:"eventtime,omitempty"`
	SteamidActor       *uint64  `protobuf:"fixed64,3,opt,name=steamid_actor,json=steamidActor" json:"steamid_actor,omitempty"`
	SteamidTarget      *uint64  `protobuf:"fixed64,4,opt,name=steamid_target,json=steamidTarget" json:"steamid_target,omitempty"`
	Gameid             *uint64  `protobuf:"fixed64,5,opt,name=gameid" json:"gameid,omitempty"`
	Packageid          *uint32  `protobuf:"varint,6,opt,name=packageid" json:"packageid,omitempty"`
	Shortcutid         *uint32  `protobuf:"varint,7,opt,name=shortcutid" json:"shortcutid,omitempty"`
	AchievementNames   []string `protobuf:"bytes,8,rep,name=achievement_names,json=achievementNames" json:"achievement_names,omitempty"`
	ClanEventid        *uint64  `protobuf:"fixed64,9,opt,name=clan_eventid,json=clanEventid" json:"clan_eventid,omitempty"`
	ClanAnnouncementid *uint64  `protobuf:"fixed64,10,opt,name=clan_announcementid,json=clanAnnouncementid" json:"clan_announcementid,omitempty"`
	Publishedfileid    *uint64  `protobuf:"fixed64,11,opt,name=publishedfileid" json:"publishedfileid,omitempty"`
	EventLastModTime   *uint32  `protobuf:"varint,12,opt,name=event_last_mod_time,json=eventLastModTime" json:"event_last_mod_time,omitempty"`
	Appids             []uint32 `protobuf:"varint,13,rep,name=appids" json:"appids,omitempty"`
	EventPostTime      *uint32  `protobuf:"varint,14,opt,name=event_post_time,json=eventPostTime" json:"event_post_time,omitempty"`
}

func (x *CUserNews_Event) Reset() {
	*x = CUserNews_Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNews_Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNews_Event) ProtoMessage() {}

func (x *CUserNews_Event) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNews_Event.ProtoReflect.Descriptor instead.
func (*CUserNews_Event) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{0}
}

func (x *CUserNews_Event) GetEventtype() uint32 {
	if x != nil && x.Eventtype != nil {
		return *x.Eventtype
	}
	return 0
}

func (x *CUserNews_Event) GetEventtime() uint32 {
	if x != nil && x.Eventtime != nil {
		return *x.Eventtime
	}
	return 0
}

func (x *CUserNews_Event) GetSteamidActor() uint64 {
	if x != nil && x.SteamidActor != nil {
		return *x.SteamidActor
	}
	return 0
}

func (x *CUserNews_Event) GetSteamidTarget() uint64 {
	if x != nil && x.SteamidTarget != nil {
		return *x.SteamidTarget
	}
	return 0
}

func (x *CUserNews_Event) GetGameid() uint64 {
	if x != nil && x.Gameid != nil {
		return *x.Gameid
	}
	return 0
}

func (x *CUserNews_Event) GetPackageid() uint32 {
	if x != nil && x.Packageid != nil {
		return *x.Packageid
	}
	return 0
}

func (x *CUserNews_Event) GetShortcutid() uint32 {
	if x != nil && x.Shortcutid != nil {
		return *x.Shortcutid
	}
	return 0
}

func (x *CUserNews_Event) GetAchievementNames() []string {
	if x != nil {
		return x.AchievementNames
	}
	return nil
}

func (x *CUserNews_Event) GetClanEventid() uint64 {
	if x != nil && x.ClanEventid != nil {
		return *x.ClanEventid
	}
	return 0
}

func (x *CUserNews_Event) GetClanAnnouncementid() uint64 {
	if x != nil && x.ClanAnnouncementid != nil {
		return *x.ClanAnnouncementid
	}
	return 0
}

func (x *CUserNews_Event) GetPublishedfileid() uint64 {
	if x != nil && x.Publishedfileid != nil {
		return *x.Publishedfileid
	}
	return 0
}

func (x *CUserNews_Event) GetEventLastModTime() uint32 {
	if x != nil && x.EventLastModTime != nil {
		return *x.EventLastModTime
	}
	return 0
}

func (x *CUserNews_Event) GetAppids() []uint32 {
	if x != nil {
		return x.Appids
	}
	return nil
}

func (x *CUserNews_Event) GetEventPostTime() uint32 {
	if x != nil && x.EventPostTime != nil {
		return *x.EventPostTime
	}
	return 0
}

type CUserNews_GetAppDetailsSpotlight_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid              *uint32 `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	IncludeAlreadySeen *bool   `protobuf:"varint,2,opt,name=include_already_seen,json=includeAlreadySeen" json:"include_already_seen,omitempty"`
}

func (x *CUserNews_GetAppDetailsSpotlight_Request) Reset() {
	*x = CUserNews_GetAppDetailsSpotlight_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNews_GetAppDetailsSpotlight_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNews_GetAppDetailsSpotlight_Request) ProtoMessage() {}

func (x *CUserNews_GetAppDetailsSpotlight_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNews_GetAppDetailsSpotlight_Request.ProtoReflect.Descriptor instead.
func (*CUserNews_GetAppDetailsSpotlight_Request) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{1}
}

func (x *CUserNews_GetAppDetailsSpotlight_Request) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CUserNews_GetAppDetailsSpotlight_Request) GetIncludeAlreadySeen() bool {
	if x != nil && x.IncludeAlreadySeen != nil {
		return *x.IncludeAlreadySeen
	}
	return false
}

type CUserNews_GetAppDetailsSpotlight_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
}

func (x *CUserNews_GetAppDetailsSpotlight_Response) Reset() {
	*x = CUserNews_GetAppDetailsSpotlight_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNews_GetAppDetailsSpotlight_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNews_GetAppDetailsSpotlight_Response) ProtoMessage() {}

func (x *CUserNews_GetAppDetailsSpotlight_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNews_GetAppDetailsSpotlight_Response.ProtoReflect.Descriptor instead.
func (*CUserNews_GetAppDetailsSpotlight_Response) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{2}
}

func (x *CUserNews_GetAppDetailsSpotlight_Response) GetEvents() []*CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

type CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType           *uint32 `protobuf:"varint,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	EventTime           *uint32 `protobuf:"varint,2,opt,name=event_time,json=eventTime" json:"event_time,omitempty"`
	ClanId              *uint64 `protobuf:"fixed64,3,opt,name=clan_id,json=clanId" json:"clan_id,omitempty"`
	ClanAnnouncementid  *uint64 `protobuf:"fixed64,4,opt,name=clan_announcementid,json=clanAnnouncementid" json:"clan_announcementid,omitempty"`
	Appid               *uint32 `protobuf:"varint,5,opt,name=appid" json:"appid,omitempty"`
	Rtime32LastModified *uint32 `protobuf:"varint,6,opt,name=rtime32_last_modified,json=rtime32LastModified" json:"rtime32_last_modified,omitempty"`
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) Reset() {
	*x = CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) ProtoMessage() {}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent.ProtoReflect.Descriptor instead.
func (*CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{3}
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) GetEventType() uint32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) GetEventTime() uint32 {
	if x != nil && x.EventTime != nil {
		return *x.EventTime
	}
	return 0
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) GetClanId() uint64 {
	if x != nil && x.ClanId != nil {
		return *x.ClanId
	}
	return 0
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) GetClanAnnouncementid() uint64 {
	if x != nil && x.ClanAnnouncementid != nil {
		return *x.ClanAnnouncementid
	}
	return 0
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent) GetRtime32LastModified() uint32 {
	if x != nil && x.Rtime32LastModified != nil {
		return *x.Rtime32LastModified
	}
	return 0
}

type CUserNews_GetUserNews_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       *uint32 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Starttime   *uint32 `protobuf:"varint,2,opt,name=starttime" json:"starttime,omitempty"`
	Endtime     *uint32 `protobuf:"varint,3,opt,name=endtime" json:"endtime,omitempty"`
	Language    *string `protobuf:"bytes,4,opt,name=language" json:"language,omitempty"`
	Filterflags *uint32 `protobuf:"varint,5,opt,name=filterflags" json:"filterflags,omitempty"`
	Filterappid *uint32 `protobuf:"varint,6,opt,name=filterappid" json:"filterappid,omitempty"`
}

func (x *CUserNews_GetUserNews_Request) Reset() {
	*x = CUserNews_GetUserNews_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNews_GetUserNews_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNews_GetUserNews_Request) ProtoMessage() {}

func (x *CUserNews_GetUserNews_Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNews_GetUserNews_Request.ProtoReflect.Descriptor instead.
func (*CUserNews_GetUserNews_Request) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{4}
}

func (x *CUserNews_GetUserNews_Request) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *CUserNews_GetUserNews_Request) GetStarttime() uint32 {
	if x != nil && x.Starttime != nil {
		return *x.Starttime
	}
	return 0
}

func (x *CUserNews_GetUserNews_Request) GetEndtime() uint32 {
	if x != nil && x.Endtime != nil {
		return *x.Endtime
	}
	return 0
}

func (x *CUserNews_GetUserNews_Request) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *CUserNews_GetUserNews_Request) GetFilterflags() uint32 {
	if x != nil && x.Filterflags != nil {
		return *x.Filterflags
	}
	return 0
}

func (x *CUserNews_GetUserNews_Request) GetFilterappid() uint32 {
	if x != nil && x.Filterappid != nil {
		return *x.Filterappid
	}
	return 0
}

type CUserNews_GetUserNews_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	News                   []*CUserNews_Event                 `protobuf:"bytes,1,rep,name=news" json:"news,omitempty"`
	AchievementDisplayData []*CUserNewsAchievementDisplayData `protobuf:"bytes,2,rep,name=achievement_display_data,json=achievementDisplayData" json:"achievement_display_data,omitempty"`
}

func (x *CUserNews_GetUserNews_Response) Reset() {
	*x = CUserNews_GetUserNews_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNews_GetUserNews_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNews_GetUserNews_Response) ProtoMessage() {}

func (x *CUserNews_GetUserNews_Response) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNews_GetUserNews_Response.ProtoReflect.Descriptor instead.
func (*CUserNews_GetUserNews_Response) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{5}
}

func (x *CUserNews_GetUserNews_Response) GetNews() []*CUserNews_Event {
	if x != nil {
		return x.News
	}
	return nil
}

func (x *CUserNews_GetUserNews_Response) GetAchievementDisplayData() []*CUserNewsAchievementDisplayData {
	if x != nil {
		return x.AchievementDisplayData
	}
	return nil
}

type CUserNewsAchievementDisplayData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid        *uint32                                         `protobuf:"varint,1,opt,name=appid" json:"appid,omitempty"`
	Achievements []*CUserNewsAchievementDisplayData_CAchievement `protobuf:"bytes,2,rep,name=achievements" json:"achievements,omitempty"`
}

func (x *CUserNewsAchievementDisplayData) Reset() {
	*x = CUserNewsAchievementDisplayData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNewsAchievementDisplayData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNewsAchievementDisplayData) ProtoMessage() {}

func (x *CUserNewsAchievementDisplayData) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNewsAchievementDisplayData.ProtoReflect.Descriptor instead.
func (*CUserNewsAchievementDisplayData) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{6}
}

func (x *CUserNewsAchievementDisplayData) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *CUserNewsAchievementDisplayData) GetAchievements() []*CUserNewsAchievementDisplayData_CAchievement {
	if x != nil {
		return x.Achievements
	}
	return nil
}

type CUserNewsAchievementDisplayData_CAchievement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DisplayName        *string  `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	DisplayDescription *string  `protobuf:"bytes,3,opt,name=display_description,json=displayDescription" json:"display_description,omitempty"`
	Icon               *string  `protobuf:"bytes,4,opt,name=icon" json:"icon,omitempty"`
	UnlockedPct        *float32 `protobuf:"fixed32,5,opt,name=unlocked_pct,json=unlockedPct" json:"unlocked_pct,omitempty"`
	Hidden             *bool    `protobuf:"varint,6,opt,name=hidden" json:"hidden,omitempty"`
}

func (x *CUserNewsAchievementDisplayData_CAchievement) Reset() {
	*x = CUserNewsAchievementDisplayData_CAchievement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_usernews_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CUserNewsAchievementDisplayData_CAchievement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CUserNewsAchievementDisplayData_CAchievement) ProtoMessage() {}

func (x *CUserNewsAchievementDisplayData_CAchievement) ProtoReflect() protoreflect.Message {
	mi := &file_service_usernews_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CUserNewsAchievementDisplayData_CAchievement.ProtoReflect.Descriptor instead.
func (*CUserNewsAchievementDisplayData_CAchievement) Descriptor() ([]byte, []int) {
	return file_service_usernews_proto_rawDescGZIP(), []int{7}
}

func (x *CUserNewsAchievementDisplayData_CAchievement) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CUserNewsAchievementDisplayData_CAchievement) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *CUserNewsAchievementDisplayData_CAchievement) GetDisplayDescription() string {
	if x != nil && x.DisplayDescription != nil {
		return *x.DisplayDescription
	}
	return ""
}

func (x *CUserNewsAchievementDisplayData_CAchievement) GetIcon() string {
	if x != nil && x.Icon != nil {
		return *x.Icon
	}
	return ""
}

func (x *CUserNewsAchievementDisplayData_CAchievement) GetUnlockedPct() float32 {
	if x != nil && x.UnlockedPct != nil {
		return *x.UnlockedPct
	}
	return 0
}

func (x *CUserNewsAchievementDisplayData_CAchievement) GetHidden() bool {
	if x != nil && x.Hidden != nil {
		return *x.Hidden
	}
	return false
}

var File_service_usernews_proto protoreflect.FileDescriptor

var file_service_usernews_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x65,
	0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x04, 0x0a, 0x0f, 0x43, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x65, 0x61,
	0x6d, 0x69, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x06, 0x52,
	0x0c, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0d, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x69, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x63, 0x75, 0x74, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x6e, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0b, 0x63,
	0x6c, 0x61, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6c,
	0x61, 0x6e, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x06, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x6e, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x69, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x06, 0x52, 0x0f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x66,
	0x69, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x69, 0x64, 0x73, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x70, 0x70, 0x69, 0x64, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x72, 0x0a, 0x28, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77,
	0x73, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x53,
	0x70, 0x6f, 0x74, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x41, 0x6c, 0x72,
	0x65, 0x61, 0x64, 0x79, 0x53, 0x65, 0x65, 0x6e, 0x22, 0x7d, 0x0a, 0x29, 0x43, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x53, 0x70, 0x6f, 0x74, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77,
	0x73, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x53,
	0x70, 0x6f, 0x74, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x8b, 0x02, 0x0a, 0x37, 0x43, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x53, 0x70, 0x6f, 0x74, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x06, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6c,
	0x61, 0x6e, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x06, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x6e, 0x41, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x69,
	0x64, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x33, 0x32, 0x5f, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x13, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x33, 0x32, 0x4c, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0xcd, 0x01, 0x0a, 0x1d, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x65, 0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x70, 0x70,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x22, 0xa2, 0x01, 0x0a, 0x1e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e,
	0x65, 0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65,
	0x77, 0x73, 0x5f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x12, 0x5a,
	0x0a, 0x18, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x16, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x22, 0x8a, 0x01, 0x0a, 0x1f, 0x43,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x69, 0x64, 0x12, 0x51, 0x0a, 0x0c, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x43, 0x55, 0x73,
	0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x43, 0x41, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x61, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xe5, 0x01, 0x0a, 0x2c, 0x43, 0x55, 0x73, 0x65,
	0x72, 0x4e, 0x65, 0x77, 0x73, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x5f, 0x43, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x5f, 0x70, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x50, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x32,
	0xcb, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x6f, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x53, 0x70, 0x6f,
	0x74, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x29, 0x2e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65,
	0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x53, 0x70, 0x6f, 0x74, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x43, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x53, 0x70, 0x6f, 0x74, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x1e, 0x2e, 0x43,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x65, 0x77, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x43,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x65, 0x77, 0x73, 0x5f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x65, 0x77, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
}

var (
	file_service_usernews_proto_rawDescOnce sync.Once
	file_service_usernews_proto_rawDescData = file_service_usernews_proto_rawDesc
)

func file_service_usernews_proto_rawDescGZIP() []byte {
	file_service_usernews_proto_rawDescOnce.Do(func() {
		file_service_usernews_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_usernews_proto_rawDescData)
	})
	return file_service_usernews_proto_rawDescData
}

var file_service_usernews_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_usernews_proto_goTypes = []any{
	(*CUserNews_Event)(nil),                                         // 0: CUserNews_Event
	(*CUserNews_GetAppDetailsSpotlight_Request)(nil),                // 1: CUserNews_GetAppDetailsSpotlight_Request
	(*CUserNews_GetAppDetailsSpotlight_Response)(nil),               // 2: CUserNews_GetAppDetailsSpotlight_Response
	(*CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent)(nil), // 3: CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent
	(*CUserNews_GetUserNews_Request)(nil),                           // 4: CUserNews_GetUserNews_Request
	(*CUserNews_GetUserNews_Response)(nil),                          // 5: CUserNews_GetUserNews_Response
	(*CUserNewsAchievementDisplayData)(nil),                         // 6: CUserNewsAchievementDisplayData
	(*CUserNewsAchievementDisplayData_CAchievement)(nil),            // 7: CUserNewsAchievementDisplayData_CAchievement
}
var file_service_usernews_proto_depIdxs = []int32{
	3, // 0: CUserNews_GetAppDetailsSpotlight_Response.events:type_name -> CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent
	0, // 1: CUserNews_GetUserNews_Response.news:type_name -> CUserNews_Event
	6, // 2: CUserNews_GetUserNews_Response.achievement_display_data:type_name -> CUserNewsAchievementDisplayData
	7, // 3: CUserNewsAchievementDisplayData.achievements:type_name -> CUserNewsAchievementDisplayData_CAchievement
	1, // 4: UserNews.GetAppDetailsSpotlight:input_type -> CUserNews_GetAppDetailsSpotlight_Request
	4, // 5: UserNews.GetUserNews:input_type -> CUserNews_GetUserNews_Request
	2, // 6: UserNews.GetAppDetailsSpotlight:output_type -> CUserNews_GetAppDetailsSpotlight_Response
	5, // 7: UserNews.GetUserNews:output_type -> CUserNews_GetUserNews_Response
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_usernews_proto_init() }
func file_service_usernews_proto_init() {
	if File_service_usernews_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_usernews_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNews_Event); i {
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
		file_service_usernews_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNews_GetAppDetailsSpotlight_Request); i {
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
		file_service_usernews_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNews_GetAppDetailsSpotlight_Response); i {
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
		file_service_usernews_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNews_GetAppDetailsSpotlight_Response_FeaturedEvent); i {
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
		file_service_usernews_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNews_GetUserNews_Request); i {
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
		file_service_usernews_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNews_GetUserNews_Response); i {
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
		file_service_usernews_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNewsAchievementDisplayData); i {
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
		file_service_usernews_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CUserNewsAchievementDisplayData_CAchievement); i {
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
			RawDescriptor: file_service_usernews_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_usernews_proto_goTypes,
		DependencyIndexes: file_service_usernews_proto_depIdxs,
		MessageInfos:      file_service_usernews_proto_msgTypes,
	}.Build()
	File_service_usernews_proto = out.File
	file_service_usernews_proto_rawDesc = nil
	file_service_usernews_proto_goTypes = nil
	file_service_usernews_proto_depIdxs = nil
}
