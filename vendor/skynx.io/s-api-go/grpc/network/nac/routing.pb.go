// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: skynx/protobuf/network/v1/nac/routing.proto

package nac

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

type RoutingScope int32

const (
	RoutingScope_SUBNET  RoutingScope = 0
	RoutingScope_NETWORK RoutingScope = 1
)

// Enum value maps for RoutingScope.
var (
	RoutingScope_name = map[int32]string{
		0: "SUBNET",
		1: "NETWORK",
	}
	RoutingScope_value = map[string]int32{
		"SUBNET":  0,
		"NETWORK": 1,
	}
)

func (x RoutingScope) Enum() *RoutingScope {
	p := new(RoutingScope)
	*p = x
	return p
}

func (x RoutingScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoutingScope) Descriptor() protoreflect.EnumDescriptor {
	return file_skynx_protobuf_network_v1_nac_routing_proto_enumTypes[0].Descriptor()
}

func (RoutingScope) Type() protoreflect.EnumType {
	return &file_skynx_protobuf_network_v1_nac_routing_proto_enumTypes[0]
}

func (x RoutingScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoutingScope.Descriptor instead.
func (RoutingScope) EnumDescriptor() ([]byte, []int) {
	return file_skynx_protobuf_network_v1_nac_routing_proto_rawDescGZIP(), []int{0}
}

// routing-level access
type RoutingDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationID  string       `protobuf:"bytes,1,opt,name=locationID,proto3" json:"locationID,omitempty"`
	SubnetID    string       `protobuf:"bytes,11,opt,name=subnetID,proto3" json:"subnetID,omitempty"`
	Scope       RoutingScope `protobuf:"varint,21,opt,name=scope,proto3,enum=nac.RoutingScope" json:"scope,omitempty"`
	NetworkCIDR string       `protobuf:"bytes,31,opt,name=networkCIDR,proto3" json:"networkCIDR,omitempty"`
	SubnetCIDR  string       `protobuf:"bytes,41,opt,name=subnetCIDR,proto3" json:"subnetCIDR,omitempty"`
}

func (x *RoutingDomain) Reset() {
	*x = RoutingDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_skynx_protobuf_network_v1_nac_routing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoutingDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingDomain) ProtoMessage() {}

func (x *RoutingDomain) ProtoReflect() protoreflect.Message {
	mi := &file_skynx_protobuf_network_v1_nac_routing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingDomain.ProtoReflect.Descriptor instead.
func (*RoutingDomain) Descriptor() ([]byte, []int) {
	return file_skynx_protobuf_network_v1_nac_routing_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingDomain) GetLocationID() string {
	if x != nil {
		return x.LocationID
	}
	return ""
}

func (x *RoutingDomain) GetSubnetID() string {
	if x != nil {
		return x.SubnetID
	}
	return ""
}

func (x *RoutingDomain) GetScope() RoutingScope {
	if x != nil {
		return x.Scope
	}
	return RoutingScope_SUBNET
}

func (x *RoutingDomain) GetNetworkCIDR() string {
	if x != nil {
		return x.NetworkCIDR
	}
	return ""
}

func (x *RoutingDomain) GetSubnetCIDR() string {
	if x != nil {
		return x.SubnetCIDR
	}
	return ""
}

var File_skynx_protobuf_network_v1_nac_routing_proto protoreflect.FileDescriptor

var file_skynx_protobuf_network_v1_nac_routing_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x63, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e,
	0x61, 0x63, 0x22, 0xb6, 0x01, 0x0a, 0x0d, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x44,
	0x12, 0x27, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x6e, 0x61, 0x63, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44, 0x52, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x49, 0x44, 0x52, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x43, 0x49, 0x44, 0x52, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x43, 0x49, 0x44, 0x52, 0x2a, 0x27, 0x0a, 0x0c, 0x52,
	0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53,
	0x55, 0x42, 0x4e, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x45, 0x54, 0x57, 0x4f,
	0x52, 0x4b, 0x10, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x73, 0x6b, 0x79, 0x6e, 0x78, 0x2e, 0x69, 0x6f,
	0x2f, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6e, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_skynx_protobuf_network_v1_nac_routing_proto_rawDescOnce sync.Once
	file_skynx_protobuf_network_v1_nac_routing_proto_rawDescData = file_skynx_protobuf_network_v1_nac_routing_proto_rawDesc
)

func file_skynx_protobuf_network_v1_nac_routing_proto_rawDescGZIP() []byte {
	file_skynx_protobuf_network_v1_nac_routing_proto_rawDescOnce.Do(func() {
		file_skynx_protobuf_network_v1_nac_routing_proto_rawDescData = protoimpl.X.CompressGZIP(file_skynx_protobuf_network_v1_nac_routing_proto_rawDescData)
	})
	return file_skynx_protobuf_network_v1_nac_routing_proto_rawDescData
}

var file_skynx_protobuf_network_v1_nac_routing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_skynx_protobuf_network_v1_nac_routing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_skynx_protobuf_network_v1_nac_routing_proto_goTypes = []interface{}{
	(RoutingScope)(0),     // 0: nac.RoutingScope
	(*RoutingDomain)(nil), // 1: nac.RoutingDomain
}
var file_skynx_protobuf_network_v1_nac_routing_proto_depIdxs = []int32{
	0, // 0: nac.RoutingDomain.scope:type_name -> nac.RoutingScope
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_skynx_protobuf_network_v1_nac_routing_proto_init() }
func file_skynx_protobuf_network_v1_nac_routing_proto_init() {
	if File_skynx_protobuf_network_v1_nac_routing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_skynx_protobuf_network_v1_nac_routing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoutingDomain); i {
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
			RawDescriptor: file_skynx_protobuf_network_v1_nac_routing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_skynx_protobuf_network_v1_nac_routing_proto_goTypes,
		DependencyIndexes: file_skynx_protobuf_network_v1_nac_routing_proto_depIdxs,
		EnumInfos:         file_skynx_protobuf_network_v1_nac_routing_proto_enumTypes,
		MessageInfos:      file_skynx_protobuf_network_v1_nac_routing_proto_msgTypes,
	}.Build()
	File_skynx_protobuf_network_v1_nac_routing_proto = out.File
	file_skynx_protobuf_network_v1_nac_routing_proto_rawDesc = nil
	file_skynx_protobuf_network_v1_nac_routing_proto_goTypes = nil
	file_skynx_protobuf_network_v1_nac_routing_proto_depIdxs = nil
}
