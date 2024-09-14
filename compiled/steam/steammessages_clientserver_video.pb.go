// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: steammessages_clientserver_video.proto

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

type CMsgVideoGameRecordingRepresentation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepresentationName   *string                            `protobuf:"bytes,2,opt,name=representation_name,json=representationName" json:"representation_name,omitempty"`
	HorizontalResolution *uint32                            `protobuf:"varint,3,opt,name=horizontal_resolution,json=horizontalResolution" json:"horizontal_resolution,omitempty"`
	VerticalResolution   *uint32                            `protobuf:"varint,4,opt,name=vertical_resolution,json=verticalResolution" json:"vertical_resolution,omitempty"`
	FrameRate            *float64                           `protobuf:"fixed64,5,opt,name=frame_rate,json=frameRate" json:"frame_rate,omitempty"`
	Bandwidth            *uint32                            `protobuf:"varint,6,opt,name=bandwidth" json:"bandwidth,omitempty"`
	AudioSampleRate      *uint32                            `protobuf:"varint,7,opt,name=audio_sample_rate,json=audioSampleRate" json:"audio_sample_rate,omitempty"`
	FrameRateString      *string                            `protobuf:"bytes,8,opt,name=frame_rate_string,json=frameRateString" json:"frame_rate_string,omitempty"`
	Codec                *string                            `protobuf:"bytes,9,opt,name=codec" json:"codec,omitempty"`
	AudioChannelConfig   *uint32                            `protobuf:"varint,10,opt,name=audio_channel_config,json=audioChannelConfig" json:"audio_channel_config,omitempty"`
	SegmentInfo          []*CVideo_GameRecordingSegmentInfo `protobuf:"bytes,11,rep,name=segment_info,json=segmentInfo" json:"segment_info,omitempty"`
}

func (x *CMsgVideoGameRecordingRepresentation) Reset() {
	*x = CMsgVideoGameRecordingRepresentation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgVideoGameRecordingRepresentation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgVideoGameRecordingRepresentation) ProtoMessage() {}

func (x *CMsgVideoGameRecordingRepresentation) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgVideoGameRecordingRepresentation.ProtoReflect.Descriptor instead.
func (*CMsgVideoGameRecordingRepresentation) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_video_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgVideoGameRecordingRepresentation) GetRepresentationName() string {
	if x != nil && x.RepresentationName != nil {
		return *x.RepresentationName
	}
	return ""
}

func (x *CMsgVideoGameRecordingRepresentation) GetHorizontalResolution() uint32 {
	if x != nil && x.HorizontalResolution != nil {
		return *x.HorizontalResolution
	}
	return 0
}

func (x *CMsgVideoGameRecordingRepresentation) GetVerticalResolution() uint32 {
	if x != nil && x.VerticalResolution != nil {
		return *x.VerticalResolution
	}
	return 0
}

func (x *CMsgVideoGameRecordingRepresentation) GetFrameRate() float64 {
	if x != nil && x.FrameRate != nil {
		return *x.FrameRate
	}
	return 0
}

func (x *CMsgVideoGameRecordingRepresentation) GetBandwidth() uint32 {
	if x != nil && x.Bandwidth != nil {
		return *x.Bandwidth
	}
	return 0
}

func (x *CMsgVideoGameRecordingRepresentation) GetAudioSampleRate() uint32 {
	if x != nil && x.AudioSampleRate != nil {
		return *x.AudioSampleRate
	}
	return 0
}

func (x *CMsgVideoGameRecordingRepresentation) GetFrameRateString() string {
	if x != nil && x.FrameRateString != nil {
		return *x.FrameRateString
	}
	return ""
}

func (x *CMsgVideoGameRecordingRepresentation) GetCodec() string {
	if x != nil && x.Codec != nil {
		return *x.Codec
	}
	return ""
}

func (x *CMsgVideoGameRecordingRepresentation) GetAudioChannelConfig() uint32 {
	if x != nil && x.AudioChannelConfig != nil {
		return *x.AudioChannelConfig
	}
	return 0
}

func (x *CMsgVideoGameRecordingRepresentation) GetSegmentInfo() []*CVideo_GameRecordingSegmentInfo {
	if x != nil {
		return x.SegmentInfo
	}
	return nil
}

type CMsgVideoGameRecordingComponent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComponentName   *string                                 `protobuf:"bytes,1,opt,name=component_name,json=componentName" json:"component_name,omitempty"`
	Contents        *uint32                                 `protobuf:"varint,2,opt,name=contents" json:"contents,omitempty"`
	SegmentSize     *uint32                                 `protobuf:"varint,3,opt,name=segment_size,json=segmentSize" json:"segment_size,omitempty"`
	FileType        *string                                 `protobuf:"bytes,4,opt,name=file_type,json=fileType" json:"file_type,omitempty"`
	Representations []*CMsgVideoGameRecordingRepresentation `protobuf:"bytes,5,rep,name=representations" json:"representations,omitempty"`
}

func (x *CMsgVideoGameRecordingComponent) Reset() {
	*x = CMsgVideoGameRecordingComponent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgVideoGameRecordingComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgVideoGameRecordingComponent) ProtoMessage() {}

func (x *CMsgVideoGameRecordingComponent) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgVideoGameRecordingComponent.ProtoReflect.Descriptor instead.
func (*CMsgVideoGameRecordingComponent) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_video_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgVideoGameRecordingComponent) GetComponentName() string {
	if x != nil && x.ComponentName != nil {
		return *x.ComponentName
	}
	return ""
}

func (x *CMsgVideoGameRecordingComponent) GetContents() uint32 {
	if x != nil && x.Contents != nil {
		return *x.Contents
	}
	return 0
}

func (x *CMsgVideoGameRecordingComponent) GetSegmentSize() uint32 {
	if x != nil && x.SegmentSize != nil {
		return *x.SegmentSize
	}
	return 0
}

func (x *CMsgVideoGameRecordingComponent) GetFileType() string {
	if x != nil && x.FileType != nil {
		return *x.FileType
	}
	return ""
}

func (x *CMsgVideoGameRecordingComponent) GetRepresentations() []*CMsgVideoGameRecordingRepresentation {
	if x != nil {
		return x.Representations
	}
	return nil
}

type CMsgVideoGameRecordingDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Steamid                  *uint64                            `protobuf:"varint,1,opt,name=steamid" json:"steamid,omitempty"`
	AppId                    *uint32                            `protobuf:"varint,2,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	NumSegments              *uint32                            `protobuf:"varint,3,opt,name=num_segments,json=numSegments" json:"num_segments,omitempty"`
	LengthMilliseconds       *uint32                            `protobuf:"varint,4,opt,name=length_milliseconds,json=lengthMilliseconds" json:"length_milliseconds,omitempty"`
	SegmentDurationTimescale *uint32                            `protobuf:"varint,5,opt,name=segment_duration_timescale,json=segmentDurationTimescale" json:"segment_duration_timescale,omitempty"`
	SegmentDuration          *uint32                            `protobuf:"varint,6,opt,name=segment_duration,json=segmentDuration" json:"segment_duration,omitempty"`
	Components               []*CMsgVideoGameRecordingComponent `protobuf:"bytes,7,rep,name=components" json:"components,omitempty"`
	StartTimeMs              *uint32                            `protobuf:"varint,8,opt,name=start_time_ms,json=startTimeMs" json:"start_time_ms,omitempty"`
	StartOffsetInTimelineMs  *uint32                            `protobuf:"varint,9,opt,name=start_offset_in_timeline_ms,json=startOffsetInTimelineMs" json:"start_offset_in_timeline_ms,omitempty"`
}

func (x *CMsgVideoGameRecordingDef) Reset() {
	*x = CMsgVideoGameRecordingDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CMsgVideoGameRecordingDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgVideoGameRecordingDef) ProtoMessage() {}

func (x *CMsgVideoGameRecordingDef) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgVideoGameRecordingDef.ProtoReflect.Descriptor instead.
func (*CMsgVideoGameRecordingDef) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_video_proto_rawDescGZIP(), []int{2}
}

func (x *CMsgVideoGameRecordingDef) GetSteamid() uint64 {
	if x != nil && x.Steamid != nil {
		return *x.Steamid
	}
	return 0
}

func (x *CMsgVideoGameRecordingDef) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *CMsgVideoGameRecordingDef) GetNumSegments() uint32 {
	if x != nil && x.NumSegments != nil {
		return *x.NumSegments
	}
	return 0
}

func (x *CMsgVideoGameRecordingDef) GetLengthMilliseconds() uint32 {
	if x != nil && x.LengthMilliseconds != nil {
		return *x.LengthMilliseconds
	}
	return 0
}

func (x *CMsgVideoGameRecordingDef) GetSegmentDurationTimescale() uint32 {
	if x != nil && x.SegmentDurationTimescale != nil {
		return *x.SegmentDurationTimescale
	}
	return 0
}

func (x *CMsgVideoGameRecordingDef) GetSegmentDuration() uint32 {
	if x != nil && x.SegmentDuration != nil {
		return *x.SegmentDuration
	}
	return 0
}

func (x *CMsgVideoGameRecordingDef) GetComponents() []*CMsgVideoGameRecordingComponent {
	if x != nil {
		return x.Components
	}
	return nil
}

func (x *CMsgVideoGameRecordingDef) GetStartTimeMs() uint32 {
	if x != nil && x.StartTimeMs != nil {
		return *x.StartTimeMs
	}
	return 0
}

func (x *CMsgVideoGameRecordingDef) GetStartOffsetInTimelineMs() uint32 {
	if x != nil && x.StartOffsetInTimelineMs != nil {
		return *x.StartOffsetInTimelineMs
	}
	return 0
}

type CVideo_GameRecordingSegmentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SegmentNumber      *uint32 `protobuf:"varint,1,opt,name=segment_number,json=segmentNumber" json:"segment_number,omitempty"`
	SegmentSizeBytes   *uint64 `protobuf:"varint,2,opt,name=segment_size_bytes,json=segmentSizeBytes" json:"segment_size_bytes,omitempty"`
	ComponentName      *string `protobuf:"bytes,3,opt,name=component_name,json=componentName" json:"component_name,omitempty"`
	RepresentationName *string `protobuf:"bytes,4,opt,name=representation_name,json=representationName" json:"representation_name,omitempty"`
}

func (x *CVideo_GameRecordingSegmentInfo) Reset() {
	*x = CVideo_GameRecordingSegmentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVideo_GameRecordingSegmentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVideo_GameRecordingSegmentInfo) ProtoMessage() {}

func (x *CVideo_GameRecordingSegmentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVideo_GameRecordingSegmentInfo.ProtoReflect.Descriptor instead.
func (*CVideo_GameRecordingSegmentInfo) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_video_proto_rawDescGZIP(), []int{3}
}

func (x *CVideo_GameRecordingSegmentInfo) GetSegmentNumber() uint32 {
	if x != nil && x.SegmentNumber != nil {
		return *x.SegmentNumber
	}
	return 0
}

func (x *CVideo_GameRecordingSegmentInfo) GetSegmentSizeBytes() uint64 {
	if x != nil && x.SegmentSizeBytes != nil {
		return *x.SegmentSizeBytes
	}
	return 0
}

func (x *CVideo_GameRecordingSegmentInfo) GetComponentName() string {
	if x != nil && x.ComponentName != nil {
		return *x.ComponentName
	}
	return ""
}

func (x *CVideo_GameRecordingSegmentInfo) GetRepresentationName() string {
	if x != nil && x.RepresentationName != nil {
		return *x.RepresentationName
	}
	return ""
}

type CVideo_GameRecordingSegmentUploadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SegmentInfo    *CVideo_GameRecordingSegmentInfo                     `protobuf:"bytes,1,opt,name=segment_info,json=segmentInfo" json:"segment_info,omitempty"`
	UrlHost        *string                                              `protobuf:"bytes,2,opt,name=url_host,json=urlHost" json:"url_host,omitempty"`
	UrlPath        *string                                              `protobuf:"bytes,3,opt,name=url_path,json=urlPath" json:"url_path,omitempty"`
	UseHttps       *bool                                                `protobuf:"varint,4,opt,name=use_https,json=useHttps" json:"use_https,omitempty"`
	RequestHeaders []*CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders `protobuf:"bytes,5,rep,name=request_headers,json=requestHeaders" json:"request_headers,omitempty"`
}

func (x *CVideo_GameRecordingSegmentUploadInfo) Reset() {
	*x = CVideo_GameRecordingSegmentUploadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVideo_GameRecordingSegmentUploadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVideo_GameRecordingSegmentUploadInfo) ProtoMessage() {}

func (x *CVideo_GameRecordingSegmentUploadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVideo_GameRecordingSegmentUploadInfo.ProtoReflect.Descriptor instead.
func (*CVideo_GameRecordingSegmentUploadInfo) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_video_proto_rawDescGZIP(), []int{4}
}

func (x *CVideo_GameRecordingSegmentUploadInfo) GetSegmentInfo() *CVideo_GameRecordingSegmentInfo {
	if x != nil {
		return x.SegmentInfo
	}
	return nil
}

func (x *CVideo_GameRecordingSegmentUploadInfo) GetUrlHost() string {
	if x != nil && x.UrlHost != nil {
		return *x.UrlHost
	}
	return ""
}

func (x *CVideo_GameRecordingSegmentUploadInfo) GetUrlPath() string {
	if x != nil && x.UrlPath != nil {
		return *x.UrlPath
	}
	return ""
}

func (x *CVideo_GameRecordingSegmentUploadInfo) GetUseHttps() bool {
	if x != nil && x.UseHttps != nil {
		return *x.UseHttps
	}
	return false
}

func (x *CVideo_GameRecordingSegmentUploadInfo) GetRequestHeaders() []*CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders {
	if x != nil {
		return x.RequestHeaders
	}
	return nil
}

type CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders) Reset() {
	*x = CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_steammessages_clientserver_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders) ProtoMessage() {}

func (x *CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders) ProtoReflect() protoreflect.Message {
	mi := &file_steammessages_clientserver_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders.ProtoReflect.Descriptor instead.
func (*CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders) Descriptor() ([]byte, []int) {
	return file_steammessages_clientserver_video_proto_rawDescGZIP(), []int{4, 0}
}

func (x *CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_steammessages_clientserver_video_proto protoreflect.FileDescriptor

var file_steammessages_clientserver_video_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdf, 0x03, 0x0a, 0x24, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x72,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x15,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x2a, 0x0a, 0x11, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x30, 0x0a,
	0x14, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x43, 0x0a, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x43, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x1f, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x0f, 0x72,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xad, 0x03, 0x0a,
	0x19, 0x43, 0x4d, 0x73, 0x67, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x65, 0x61, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x65,
	0x61, 0x6d, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e,
	0x75, 0x6d, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f,
	0x0a, 0x13, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x6c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x3c, 0x0a, 0x1a, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x18, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x43,
	0x4d, 0x73, 0x67, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x3c,
	0x0a, 0x1b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x17, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x49, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x73, 0x22, 0xce, 0x01, 0x0a,
	0x1f, 0x43, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13,
	0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd5, 0x02,
	0x0a, 0x25, 0x43, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x43, 0x0a, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x43, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08,
	0x75, 0x72, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x75, 0x72, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x72, 0x6c, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x73, 0x65, 0x48, 0x74, 0x74, 0x70, 0x73, 0x12,
	0x5b, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x43, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x0e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x37, 0x0a, 0x0b,
	0x48, 0x54, 0x54, 0x50, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x05, 0x48, 0x01, 0x80, 0x01, 0x00,
}

var (
	file_steammessages_clientserver_video_proto_rawDescOnce sync.Once
	file_steammessages_clientserver_video_proto_rawDescData = file_steammessages_clientserver_video_proto_rawDesc
)

func file_steammessages_clientserver_video_proto_rawDescGZIP() []byte {
	file_steammessages_clientserver_video_proto_rawDescOnce.Do(func() {
		file_steammessages_clientserver_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_steammessages_clientserver_video_proto_rawDescData)
	})
	return file_steammessages_clientserver_video_proto_rawDescData
}

var file_steammessages_clientserver_video_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_steammessages_clientserver_video_proto_goTypes = []any{
	(*CMsgVideoGameRecordingRepresentation)(nil),              // 0: CMsgVideoGameRecordingRepresentation
	(*CMsgVideoGameRecordingComponent)(nil),                   // 1: CMsgVideoGameRecordingComponent
	(*CMsgVideoGameRecordingDef)(nil),                         // 2: CMsgVideoGameRecordingDef
	(*CVideo_GameRecordingSegmentInfo)(nil),                   // 3: CVideo_GameRecordingSegmentInfo
	(*CVideo_GameRecordingSegmentUploadInfo)(nil),             // 4: CVideo_GameRecordingSegmentUploadInfo
	(*CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders)(nil), // 5: CVideo_GameRecordingSegmentUploadInfo.HTTPHeaders
}
var file_steammessages_clientserver_video_proto_depIdxs = []int32{
	3, // 0: CMsgVideoGameRecordingRepresentation.segment_info:type_name -> CVideo_GameRecordingSegmentInfo
	0, // 1: CMsgVideoGameRecordingComponent.representations:type_name -> CMsgVideoGameRecordingRepresentation
	1, // 2: CMsgVideoGameRecordingDef.components:type_name -> CMsgVideoGameRecordingComponent
	3, // 3: CVideo_GameRecordingSegmentUploadInfo.segment_info:type_name -> CVideo_GameRecordingSegmentInfo
	5, // 4: CVideo_GameRecordingSegmentUploadInfo.request_headers:type_name -> CVideo_GameRecordingSegmentUploadInfo.HTTPHeaders
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_steammessages_clientserver_video_proto_init() }
func file_steammessages_clientserver_video_proto_init() {
	if File_steammessages_clientserver_video_proto != nil {
		return
	}
	file_steammessages_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_steammessages_clientserver_video_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgVideoGameRecordingRepresentation); i {
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
		file_steammessages_clientserver_video_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgVideoGameRecordingComponent); i {
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
		file_steammessages_clientserver_video_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CMsgVideoGameRecordingDef); i {
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
		file_steammessages_clientserver_video_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CVideo_GameRecordingSegmentInfo); i {
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
		file_steammessages_clientserver_video_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CVideo_GameRecordingSegmentUploadInfo); i {
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
		file_steammessages_clientserver_video_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CVideo_GameRecordingSegmentUploadInfo_HTTPHeaders); i {
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
			RawDescriptor: file_steammessages_clientserver_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_steammessages_clientserver_video_proto_goTypes,
		DependencyIndexes: file_steammessages_clientserver_video_proto_depIdxs,
		MessageInfos:      file_steammessages_clientserver_video_proto_msgTypes,
	}.Build()
	File_steammessages_clientserver_video_proto = out.File
	file_steammessages_clientserver_video_proto_rawDesc = nil
	file_steammessages_clientserver_video_proto_goTypes = nil
	file_steammessages_clientserver_video_proto_depIdxs = nil
}
