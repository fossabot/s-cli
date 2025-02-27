// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/common/v1/datetime/datetime.proto

package datetime

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

type DateTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year   int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month  int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day    int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Hour   int32 `protobuf:"varint,4,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute int32 `protobuf:"varint,5,opt,name=minute,proto3" json:"minute,omitempty"`
	Second int32 `protobuf:"varint,6,opt,name=second,proto3" json:"second,omitempty"`
}

func (x *DateTime) Reset() {
	*x = DateTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_common_v1_datetime_datetime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateTime) ProtoMessage() {}

func (x *DateTime) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_common_v1_datetime_datetime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateTime.ProtoReflect.Descriptor instead.
func (*DateTime) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescGZIP(), []int{0}
}

func (x *DateTime) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *DateTime) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *DateTime) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *DateTime) GetHour() int32 {
	if x != nil {
		return x.Hour
	}
	return 0
}

func (x *DateTime) GetMinute() int32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

func (x *DateTime) GetSecond() int32 {
	if x != nil {
		return x.Second
	}
	return 0
}

var File_skynx_protobuf_common_v1_datetime_datetime_proto protoreflect.FileDescriptor

var file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x8a, 0x01, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x42, 0x28, 0x5a, 0x26, 0x73, 0x6b, 0x79,
	0x6e, 0x78, 0x2e, 0x69, 0x6f, 0x2f, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescOnce sync.Once
	file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescData = file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDesc
)

func file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescData)
	})
	return file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDescData
}

var file_skynx_protobuf_common_v1_datetime_datetime_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_skynx_protobuf_common_v1_datetime_datetime_proto_goTypes = []interface{}{
	(*DateTime)(nil), // 0: datetime.DateTime
}
var file_skynx_protobuf_common_v1_datetime_datetime_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_common_v1_datetime_datetime_proto_init() }
func file_skynx_protobuf_common_v1_datetime_datetime_proto_init() {
	if File_skynx_protobuf_common_v1_datetime_datetime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_common_v1_datetime_datetime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateTime); i {
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
			RawDescriptor: file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_common_v1_datetime_datetime_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_common_v1_datetime_datetime_proto_depIdxs,
		MessageInfos:      file_skynx_protobuf_common_v1_datetime_datetime_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_common_v1_datetime_datetime_proto = out.File
	file_skynx_protobuf_common_v1_datetime_datetime_proto_rawDesc = nil
	file_skynx_protobuf_common_v1_datetime_datetime_proto_goTypes = nil
	file_skynx_protobuf_common_v1_datetime_datetime_proto_depIdxs = nil
}
