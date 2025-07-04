// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: types/v1/logevent.proto

package typesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogEventGroup struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceId    int64                  `protobuf:"varint,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ScopeId       int64                  `protobuf:"varint,2,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty"`
	Logs          []*LogEvent            `protobuf:"bytes,3,rep,name=logs,proto3" json:"logs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogEventGroup) Reset() {
	*x = LogEventGroup{}
	mi := &file_types_v1_logevent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEventGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEventGroup) ProtoMessage() {}

func (x *LogEventGroup) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logevent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEventGroup.ProtoReflect.Descriptor instead.
func (*LogEventGroup) Descriptor() ([]byte, []int) {
	return file_types_v1_logevent_proto_rawDescGZIP(), []int{0}
}

func (x *LogEventGroup) GetResourceId() int64 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *LogEventGroup) GetScopeId() int64 {
	if x != nil {
		return x.ScopeId
	}
	return 0
}

func (x *LogEventGroup) GetLogs() []*LogEvent {
	if x != nil {
		return x.Logs
	}
	return nil
}

type IngestedLogEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UlidBase32    string                 `protobuf:"bytes,100,opt,name=ulid_base32,json=ulidBase32,proto3" json:"ulid_base32,omitempty"`
	ResourceId    int64                  `protobuf:"varint,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	ScopeId       int64                  `protobuf:"varint,2,opt,name=scope_id,json=scopeId,proto3" json:"scope_id,omitempty"`
	EventId       int64                  `protobuf:"varint,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	ParsedAt      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=parsed_at,json=parsedAt,proto3" json:"parsed_at,omitempty"`
	Raw           []byte                 `protobuf:"bytes,5,opt,name=raw,proto3" json:"raw,omitempty"`
	Structured    *StructuredLogEvent    `protobuf:"bytes,6,opt,name=structured,proto3" json:"structured,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IngestedLogEvent) Reset() {
	*x = IngestedLogEvent{}
	mi := &file_types_v1_logevent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IngestedLogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestedLogEvent) ProtoMessage() {}

func (x *IngestedLogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logevent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestedLogEvent.ProtoReflect.Descriptor instead.
func (*IngestedLogEvent) Descriptor() ([]byte, []int) {
	return file_types_v1_logevent_proto_rawDescGZIP(), []int{1}
}

func (x *IngestedLogEvent) GetUlidBase32() string {
	if x != nil {
		return x.UlidBase32
	}
	return ""
}

func (x *IngestedLogEvent) GetResourceId() int64 {
	if x != nil {
		return x.ResourceId
	}
	return 0
}

func (x *IngestedLogEvent) GetScopeId() int64 {
	if x != nil {
		return x.ScopeId
	}
	return 0
}

func (x *IngestedLogEvent) GetEventId() int64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *IngestedLogEvent) GetParsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ParsedAt
	}
	return nil
}

func (x *IngestedLogEvent) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *IngestedLogEvent) GetStructured() *StructuredLogEvent {
	if x != nil {
		return x.Structured
	}
	return nil
}

type LogEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ParsedAt      *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=parsed_at,json=parsedAt,proto3" json:"parsed_at,omitempty"`
	Raw           []byte                 `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	Structured    *StructuredLogEvent    `protobuf:"bytes,3,opt,name=structured,proto3" json:"structured,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogEvent) Reset() {
	*x = LogEvent{}
	mi := &file_types_v1_logevent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEvent) ProtoMessage() {}

func (x *LogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logevent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEvent.ProtoReflect.Descriptor instead.
func (*LogEvent) Descriptor() ([]byte, []int) {
	return file_types_v1_logevent_proto_rawDescGZIP(), []int{2}
}

func (x *LogEvent) GetParsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ParsedAt
	}
	return nil
}

func (x *LogEvent) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *LogEvent) GetStructured() *StructuredLogEvent {
	if x != nil {
		return x.Structured
	}
	return nil
}

type StructuredLogEvent struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Lvl           string                 `protobuf:"bytes,2,opt,name=lvl,proto3" json:"lvl,omitempty"`
	LvlNumber     int32                  `protobuf:"varint,201,opt,name=lvl_number,json=lvlNumber,proto3" json:"lvl_number,omitempty"`
	Msg           string                 `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Kvs           []*KV                  `protobuf:"bytes,4,rep,name=kvs,proto3" json:"kvs,omitempty"`
	TraceId       []byte                 `protobuf:"bytes,500,opt,name=trace_id,json=traceId,proto3,oneof" json:"trace_id,omitempty"`
	SpanId        []byte                 `protobuf:"bytes,501,opt,name=span_id,json=spanId,proto3,oneof" json:"span_id,omitempty"`
	OtlpFlags     *uint32                `protobuf:"varint,502,opt,name=otlp_flags,json=otlpFlags,proto3,oneof" json:"otlp_flags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StructuredLogEvent) Reset() {
	*x = StructuredLogEvent{}
	mi := &file_types_v1_logevent_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StructuredLogEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructuredLogEvent) ProtoMessage() {}

func (x *StructuredLogEvent) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logevent_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructuredLogEvent.ProtoReflect.Descriptor instead.
func (*StructuredLogEvent) Descriptor() ([]byte, []int) {
	return file_types_v1_logevent_proto_rawDescGZIP(), []int{3}
}

func (x *StructuredLogEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *StructuredLogEvent) GetLvl() string {
	if x != nil {
		return x.Lvl
	}
	return ""
}

func (x *StructuredLogEvent) GetLvlNumber() int32 {
	if x != nil {
		return x.LvlNumber
	}
	return 0
}

func (x *StructuredLogEvent) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *StructuredLogEvent) GetKvs() []*KV {
	if x != nil {
		return x.Kvs
	}
	return nil
}

func (x *StructuredLogEvent) GetTraceId() []byte {
	if x != nil {
		return x.TraceId
	}
	return nil
}

func (x *StructuredLogEvent) GetSpanId() []byte {
	if x != nil {
		return x.SpanId
	}
	return nil
}

func (x *StructuredLogEvent) GetOtlpFlags() uint32 {
	if x != nil && x.OtlpFlags != nil {
		return *x.OtlpFlags
	}
	return 0
}

var File_types_v1_logevent_proto protoreflect.FileDescriptor

var file_types_v1_logevent_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22,
	0x93, 0x02, 0x0a, 0x10, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6c, 0x69, 0x64, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x33, 0x32, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x6c, 0x69, 0x64, 0x42,
	0x61, 0x73, 0x65, 0x33, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x3c, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x64, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x37, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x08, 0x70, 0x61, 0x72, 0x73, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x3c, 0x0a,
	0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x22, 0xbf, 0x02, 0x0a, 0x12,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x76, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x76, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x6c, 0x76, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0xc9, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x76, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x1e, 0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x56, 0x52, 0x03, 0x6b, 0x76, 0x73,
	0x12, 0x1f, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xf4, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x1d, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0xf5, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x01, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x6f, 0x74, 0x6c, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0xf6,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x6f, 0x74, 0x6c, 0x70, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6f, 0x74, 0x6c, 0x70, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x42, 0x8d, 0x01,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x4c, 0x6f, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x09, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_types_v1_logevent_proto_rawDescOnce sync.Once
	file_types_v1_logevent_proto_rawDescData []byte
)

func file_types_v1_logevent_proto_rawDescGZIP() []byte {
	file_types_v1_logevent_proto_rawDescOnce.Do(func() {
		file_types_v1_logevent_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_types_v1_logevent_proto_rawDesc), len(file_types_v1_logevent_proto_rawDesc)))
	})
	return file_types_v1_logevent_proto_rawDescData
}

var file_types_v1_logevent_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_v1_logevent_proto_goTypes = []any{
	(*LogEventGroup)(nil),         // 0: types.v1.LogEventGroup
	(*IngestedLogEvent)(nil),      // 1: types.v1.IngestedLogEvent
	(*LogEvent)(nil),              // 2: types.v1.LogEvent
	(*StructuredLogEvent)(nil),    // 3: types.v1.StructuredLogEvent
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*KV)(nil),                    // 5: types.v1.KV
}
var file_types_v1_logevent_proto_depIdxs = []int32{
	2, // 0: types.v1.LogEventGroup.logs:type_name -> types.v1.LogEvent
	4, // 1: types.v1.IngestedLogEvent.parsed_at:type_name -> google.protobuf.Timestamp
	3, // 2: types.v1.IngestedLogEvent.structured:type_name -> types.v1.StructuredLogEvent
	4, // 3: types.v1.LogEvent.parsed_at:type_name -> google.protobuf.Timestamp
	3, // 4: types.v1.LogEvent.structured:type_name -> types.v1.StructuredLogEvent
	4, // 5: types.v1.StructuredLogEvent.timestamp:type_name -> google.protobuf.Timestamp
	5, // 6: types.v1.StructuredLogEvent.kvs:type_name -> types.v1.KV
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_types_v1_logevent_proto_init() }
func file_types_v1_logevent_proto_init() {
	if File_types_v1_logevent_proto != nil {
		return
	}
	file_types_v1_types_proto_init()
	file_types_v1_logevent_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_types_v1_logevent_proto_rawDesc), len(file_types_v1_logevent_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_logevent_proto_goTypes,
		DependencyIndexes: file_types_v1_logevent_proto_depIdxs,
		MessageInfos:      file_types_v1_logevent_proto_msgTypes,
	}.Build()
	File_types_v1_logevent_proto = out.File
	file_types_v1_logevent_proto_goTypes = nil
	file_types_v1_logevent_proto_depIdxs = nil
}
