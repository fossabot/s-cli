// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/network/v1/sxsp/event.proto

package sxsp

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	events "skynx.io/s-api-go/grpc/resources/events"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventPDU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *events.Event `protobuf:"bytes,21,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventPDU) Reset() {
	*x = EventPDU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_network_v1_sxsp_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPDU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPDU) ProtoMessage() {}

func (x *EventPDU) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_network_v1_sxsp_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPDU.ProtoReflect.Descriptor instead.
func (*EventPDU) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventPDU) GetEvent() *events.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_skynx_protobuf_network_v1_sxsp_event_proto protoreflect.FileDescriptor

var file_skynx_protobuf_network_v1_sxsp_event_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x78, 0x73, 0x70,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x78,
	0x73, 0x70, 0x1a, 0x2e, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x08, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x44, 0x55, 0x12, 0x23,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x42, 0x25, 0x5a, 0x23, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f,
	0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x78, 0x73, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescOnce sync.Once
	file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescData = file_skynx_protobuf_network_v1_sxsp_event_proto_rawDesc
)

func file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescData)
	})
	return file_skynx_protobuf_network_v1_sxsp_event_proto_rawDescData
}

var file_skynx_protobuf_network_v1_sxsp_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_skynx_protobuf_network_v1_sxsp_event_proto_goTypes = []interface{}{
	(*EventPDU)(nil),     // 0: sxsp.EventPDU
	(*events.Event)(nil), // 1: events.Event
}
var file_skynx_protobuf_network_v1_sxsp_event_proto_depIdxs = []int32{
	1, // 0: sxsp.EventPDU.event:type_name -> events.Event
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_network_v1_sxsp_event_proto_init() }
func file_skynx_protobuf_network_v1_sxsp_event_proto_init() {
	if File_skynx_protobuf_network_v1_sxsp_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_network_v1_sxsp_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPDU); i {
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
			RawDescriptor: file_skynx_protobuf_network_v1_sxsp_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_network_v1_sxsp_event_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_network_v1_sxsp_event_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_network_v1_sxsp_event_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_network_v1_sxsp_event_proto = out.File
	file_skynx_protobuf_network_v1_sxsp_event_proto_rawDesc = nil
	file_skynx_protobuf_network_v1_sxsp_event_proto_goTypes = nil
	file_skynx_protobuf_network_v1_sxsp_event_proto_depIdxs = nil
}
