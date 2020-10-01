// Copyright 2019 Thibault NORMAND
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.8.0
// source: spotigraph/events/v1/event.proto

package eventsv1

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// EventType enumerates all event type values.
type EventType int32

const (
	// Default value when no enumeration is specified.
	EventType_EVENT_TYPE_INVALID EventType = 0
	// Explicitly Unknown object value.
	EventType_EVENT_TYPE_UNKNOWN                EventType = 1
	EventType_EVENT_TYPE_CHAPTER_CREATED        EventType = 2
	EventType_EVENT_TYPE_CHAPTER_DELETED        EventType = 3
	EventType_EVENT_TYPE_CHAPTER_LABEL_UPDATED  EventType = 4
	EventType_EVENT_TYPE_CHAPTER_LEADER_UPDATED EventType = 5
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_INVALID",
		1: "EVENT_TYPE_UNKNOWN",
		2: "EVENT_TYPE_CHAPTER_CREATED",
		3: "EVENT_TYPE_CHAPTER_DELETED",
		4: "EVENT_TYPE_CHAPTER_LABEL_UPDATED",
		5: "EVENT_TYPE_CHAPTER_LEADER_UPDATED",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_INVALID":                0,
		"EVENT_TYPE_UNKNOWN":                1,
		"EVENT_TYPE_CHAPTER_CREATED":        2,
		"EVENT_TYPE_CHAPTER_DELETED":        3,
		"EVENT_TYPE_CHAPTER_LABEL_UPDATED":  4,
		"EVENT_TYPE_CHAPTER_LEADER_UPDATED": 5,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_spotigraph_events_v1_event_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_spotigraph_events_v1_event_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_spotigraph_events_v1_event_proto_rawDescGZIP(), []int{0}
}

// Event describes event contract.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType     EventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=spotigraph.events.v1.EventType" json:"event_type,omitempty"`
	EventId       string    `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AggregateType string    `protobuf:"bytes,3,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	AggregateId   string    `protobuf:"bytes,4,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	Meta          *any.Any  `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	// Types that are assignable to Payload:
	//	*Event_ChapterCreated
	//	*Event_ChapterDeleted
	//	*Event_ChapterLabelUpdated
	//	*Event_ChapterLeaderUpdated
	Payload isEvent_Payload `protobuf_oneof:"payload"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_events_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_events_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_spotigraph_events_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_EVENT_TYPE_INVALID
}

func (x *Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Event) GetAggregateType() string {
	if x != nil {
		return x.AggregateType
	}
	return ""
}

func (x *Event) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *Event) GetMeta() *any.Any {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (m *Event) GetPayload() isEvent_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Event) GetChapterCreated() *ChapterCreated {
	if x, ok := x.GetPayload().(*Event_ChapterCreated); ok {
		return x.ChapterCreated
	}
	return nil
}

func (x *Event) GetChapterDeleted() *ChapterDeleted {
	if x, ok := x.GetPayload().(*Event_ChapterDeleted); ok {
		return x.ChapterDeleted
	}
	return nil
}

func (x *Event) GetChapterLabelUpdated() *ChapterLabelUpdated {
	if x, ok := x.GetPayload().(*Event_ChapterLabelUpdated); ok {
		return x.ChapterLabelUpdated
	}
	return nil
}

func (x *Event) GetChapterLeaderUpdated() *ChapterLeaderUpdated {
	if x, ok := x.GetPayload().(*Event_ChapterLeaderUpdated); ok {
		return x.ChapterLeaderUpdated
	}
	return nil
}

type isEvent_Payload interface {
	isEvent_Payload()
}

type Event_ChapterCreated struct {
	ChapterCreated *ChapterCreated `protobuf:"bytes,10,opt,name=chapter_created,json=chapterCreated,proto3,oneof"`
}

type Event_ChapterDeleted struct {
	ChapterDeleted *ChapterDeleted `protobuf:"bytes,11,opt,name=chapter_deleted,json=chapterDeleted,proto3,oneof"`
}

type Event_ChapterLabelUpdated struct {
	ChapterLabelUpdated *ChapterLabelUpdated `protobuf:"bytes,12,opt,name=chapter_label_updated,json=chapterLabelUpdated,proto3,oneof"`
}

type Event_ChapterLeaderUpdated struct {
	ChapterLeaderUpdated *ChapterLeaderUpdated `protobuf:"bytes,13,opt,name=chapter_leader_updated,json=chapterLeaderUpdated,proto3,oneof"`
}

func (*Event_ChapterCreated) isEvent_Payload() {}

func (*Event_ChapterDeleted) isEvent_Payload() {}

func (*Event_ChapterLabelUpdated) isEvent_Payload() {}

func (*Event_ChapterLeaderUpdated) isEvent_Payload() {}

// ChapterCreated is raised on chapter entity creation.
type ChapterCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn      string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Label    string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	LeaderId string `protobuf:"bytes,3,opt,name=leader_id,json=leaderId,proto3" json:"leader_id,omitempty"`
}

func (x *ChapterCreated) Reset() {
	*x = ChapterCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_events_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterCreated) ProtoMessage() {}

func (x *ChapterCreated) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_events_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterCreated.ProtoReflect.Descriptor instead.
func (*ChapterCreated) Descriptor() ([]byte, []int) {
	return file_spotigraph_events_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *ChapterCreated) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *ChapterCreated) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ChapterCreated) GetLeaderId() string {
	if x != nil {
		return x.LeaderId
	}
	return ""
}

// ChapterDeleted is raised on chapter entity deletion.
type ChapterDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *ChapterDeleted) Reset() {
	*x = ChapterDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_events_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterDeleted) ProtoMessage() {}

func (x *ChapterDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_events_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterDeleted.ProtoReflect.Descriptor instead.
func (*ChapterDeleted) Descriptor() ([]byte, []int) {
	return file_spotigraph_events_v1_event_proto_rawDescGZIP(), []int{2}
}

func (x *ChapterDeleted) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

// ChapterLabelUpdated is raised on chapter entity label updates.
type ChapterLabelUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Old string `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	New string `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *ChapterLabelUpdated) Reset() {
	*x = ChapterLabelUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_events_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterLabelUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterLabelUpdated) ProtoMessage() {}

func (x *ChapterLabelUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_events_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterLabelUpdated.ProtoReflect.Descriptor instead.
func (*ChapterLabelUpdated) Descriptor() ([]byte, []int) {
	return file_spotigraph_events_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *ChapterLabelUpdated) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *ChapterLabelUpdated) GetOld() string {
	if x != nil {
		return x.Old
	}
	return ""
}

func (x *ChapterLabelUpdated) GetNew() string {
	if x != nil {
		return x.New
	}
	return ""
}

// ChapterLeaderUpdated is raised on chapter entity leader updates.
type ChapterLeaderUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Old string `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	New string `protobuf:"bytes,3,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *ChapterLeaderUpdated) Reset() {
	*x = ChapterLeaderUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spotigraph_events_v1_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterLeaderUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterLeaderUpdated) ProtoMessage() {}

func (x *ChapterLeaderUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_spotigraph_events_v1_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterLeaderUpdated.ProtoReflect.Descriptor instead.
func (*ChapterLeaderUpdated) Descriptor() ([]byte, []int) {
	return file_spotigraph_events_v1_event_proto_rawDescGZIP(), []int{4}
}

func (x *ChapterLeaderUpdated) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *ChapterLeaderUpdated) GetOld() string {
	if x != nil {
		return x.Old
	}
	return ""
}

func (x *ChapterLeaderUpdated) GetNew() string {
	if x != nil {
		return x.New
	}
	return ""
}

var File_spotigraph_events_v1_event_proto protoreflect.FileDescriptor

var file_spotigraph_events_v1_event_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x04, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x4f, 0x0a, 0x0f,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0e, 0x63,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x4f, 0x0a,
	0x0f, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0e,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x5f,
	0x0a, 0x15, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x13, 0x63, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x62, 0x0a, 0x16, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x14, 0x63,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x55,
	0x0a, 0x0e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0x4b, 0x0a, 0x13, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6f, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6e, 0x65, 0x77, 0x22, 0x4c, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65,
	0x72, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6e, 0x65, 0x77, 0x2a, 0xc8, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x48, 0x41, 0x50, 0x54, 0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x48, 0x41, 0x50, 0x54, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x48, 0x41, 0x50, 0x54, 0x45, 0x52, 0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x25, 0x0a, 0x21, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x50, 0x54, 0x45, 0x52, 0x5f, 0x4c,
	0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x05, 0x42,
	0x84, 0x01, 0x0a, 0x21, 0x6f, 0x72, 0x67, 0x2e, 0x7a, 0x65, 0x6e, 0x69, 0x74, 0x68, 0x61, 0x72,
	0x2e, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x73, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x45, 0x58, 0xaa, 0x02, 0x14, 0x53, 0x70, 0x6f, 0x74, 0x69,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x14, 0x53, 0x70, 0x6f, 0x74, 0x69, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5c, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spotigraph_events_v1_event_proto_rawDescOnce sync.Once
	file_spotigraph_events_v1_event_proto_rawDescData = file_spotigraph_events_v1_event_proto_rawDesc
)

func file_spotigraph_events_v1_event_proto_rawDescGZIP() []byte {
	file_spotigraph_events_v1_event_proto_rawDescOnce.Do(func() {
		file_spotigraph_events_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_spotigraph_events_v1_event_proto_rawDescData)
	})
	return file_spotigraph_events_v1_event_proto_rawDescData
}

var file_spotigraph_events_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spotigraph_events_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spotigraph_events_v1_event_proto_goTypes = []interface{}{
	(EventType)(0),               // 0: spotigraph.events.v1.EventType
	(*Event)(nil),                // 1: spotigraph.events.v1.Event
	(*ChapterCreated)(nil),       // 2: spotigraph.events.v1.ChapterCreated
	(*ChapterDeleted)(nil),       // 3: spotigraph.events.v1.ChapterDeleted
	(*ChapterLabelUpdated)(nil),  // 4: spotigraph.events.v1.ChapterLabelUpdated
	(*ChapterLeaderUpdated)(nil), // 5: spotigraph.events.v1.ChapterLeaderUpdated
	(*any.Any)(nil),              // 6: google.protobuf.Any
}
var file_spotigraph_events_v1_event_proto_depIdxs = []int32{
	0, // 0: spotigraph.events.v1.Event.event_type:type_name -> spotigraph.events.v1.EventType
	6, // 1: spotigraph.events.v1.Event.meta:type_name -> google.protobuf.Any
	2, // 2: spotigraph.events.v1.Event.chapter_created:type_name -> spotigraph.events.v1.ChapterCreated
	3, // 3: spotigraph.events.v1.Event.chapter_deleted:type_name -> spotigraph.events.v1.ChapterDeleted
	4, // 4: spotigraph.events.v1.Event.chapter_label_updated:type_name -> spotigraph.events.v1.ChapterLabelUpdated
	5, // 5: spotigraph.events.v1.Event.chapter_leader_updated:type_name -> spotigraph.events.v1.ChapterLeaderUpdated
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_spotigraph_events_v1_event_proto_init() }
func file_spotigraph_events_v1_event_proto_init() {
	if File_spotigraph_events_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spotigraph_events_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_spotigraph_events_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterCreated); i {
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
		file_spotigraph_events_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterDeleted); i {
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
		file_spotigraph_events_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterLabelUpdated); i {
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
		file_spotigraph_events_v1_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterLeaderUpdated); i {
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
	file_spotigraph_events_v1_event_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_ChapterCreated)(nil),
		(*Event_ChapterDeleted)(nil),
		(*Event_ChapterLabelUpdated)(nil),
		(*Event_ChapterLeaderUpdated)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spotigraph_events_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spotigraph_events_v1_event_proto_goTypes,
		DependencyIndexes: file_spotigraph_events_v1_event_proto_depIdxs,
		EnumInfos:         file_spotigraph_events_v1_event_proto_enumTypes,
		MessageInfos:      file_spotigraph_events_v1_event_proto_msgTypes,
	}.Build()
	File_spotigraph_events_v1_event_proto = out.File
	file_spotigraph_events_v1_event_proto_rawDesc = nil
	file_spotigraph_events_v1_event_proto_goTypes = nil
	file_spotigraph_events_v1_event_proto_depIdxs = nil
}