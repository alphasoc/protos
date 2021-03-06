// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: fqdn/fqdn.proto

package fqdn

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

type ResolveRequest_Src int32

const (
	ResolveRequest_UNKNOWN ResolveRequest_Src = 0
	ResolveRequest_TIP     ResolveRequest_Src = 1
	ResolveRequest_WISDOM  ResolveRequest_Src = 2
)

// Enum value maps for ResolveRequest_Src.
var (
	ResolveRequest_Src_name = map[int32]string{
		0: "UNKNOWN",
		1: "TIP",
		2: "WISDOM",
	}
	ResolveRequest_Src_value = map[string]int32{
		"UNKNOWN": 0,
		"TIP":     1,
		"WISDOM":  2,
	}
)

func (x ResolveRequest_Src) Enum() *ResolveRequest_Src {
	p := new(ResolveRequest_Src)
	*p = x
	return p
}

func (x ResolveRequest_Src) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResolveRequest_Src) Descriptor() protoreflect.EnumDescriptor {
	return file_fqdn_fqdn_proto_enumTypes[0].Descriptor()
}

func (ResolveRequest_Src) Type() protoreflect.EnumType {
	return &file_fqdn_fqdn_proto_enumTypes[0]
}

func (x ResolveRequest_Src) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResolveRequest_Src.Descriptor instead.
func (ResolveRequest_Src) EnumDescriptor() ([]byte, []int) {
	return file_fqdn_fqdn_proto_rawDescGZIP(), []int{1, 0}
}

type ResolveRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reqs []*ResolveRequest `protobuf:"bytes,1,rep,name=reqs,proto3" json:"reqs,omitempty"`
}

func (x *ResolveRequests) Reset() {
	*x = ResolveRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fqdn_fqdn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequests) ProtoMessage() {}

func (x *ResolveRequests) ProtoReflect() protoreflect.Message {
	mi := &file_fqdn_fqdn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequests.ProtoReflect.Descriptor instead.
func (*ResolveRequests) Descriptor() ([]byte, []int) {
	return file_fqdn_fqdn_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveRequests) GetReqs() []*ResolveRequest {
	if x != nil {
		return x.Reqs
	}
	return nil
}

type ResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fqdn string             `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Src  ResolveRequest_Src `protobuf:"varint,2,opt,name=src,proto3,enum=fqdn.ResolveRequest_Src" json:"src,omitempty"` // default == UNKNOWN
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fqdn_fqdn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fqdn_fqdn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_fqdn_fqdn_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveRequest) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *ResolveRequest) GetSrc() ResolveRequest_Src {
	if x != nil {
		return x.Src
	}
	return ResolveRequest_UNKNOWN
}

var File_fqdn_fqdn_proto protoreflect.FileDescriptor

var file_fqdn_fqdn_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x71, 0x64, 0x6e, 0x2f, 0x66, 0x71, 0x64, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x22, 0x3b, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x72, 0x65,
	0x71, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x66, 0x71, 0x64, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x72, 0x65, 0x71, 0x73, 0x22, 0x79, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x2a, 0x0a, 0x03, 0x73, 0x72,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x66, 0x71, 0x64, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x72,
	0x63, 0x52, 0x03, 0x73, 0x72, 0x63, 0x22, 0x27, 0x0a, 0x03, 0x53, 0x72, 0x63, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x49,
	0x50, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x49, 0x53, 0x44, 0x4f, 0x4d, 0x10, 0x02, 0x42,
	0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x73, 0x6f, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x71,
	0x64, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fqdn_fqdn_proto_rawDescOnce sync.Once
	file_fqdn_fqdn_proto_rawDescData = file_fqdn_fqdn_proto_rawDesc
)

func file_fqdn_fqdn_proto_rawDescGZIP() []byte {
	file_fqdn_fqdn_proto_rawDescOnce.Do(func() {
		file_fqdn_fqdn_proto_rawDescData = protoimpl.X.CompressGZIP(file_fqdn_fqdn_proto_rawDescData)
	})
	return file_fqdn_fqdn_proto_rawDescData
}

var file_fqdn_fqdn_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fqdn_fqdn_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fqdn_fqdn_proto_goTypes = []interface{}{
	(ResolveRequest_Src)(0), // 0: fqdn.ResolveRequest.Src
	(*ResolveRequests)(nil), // 1: fqdn.ResolveRequests
	(*ResolveRequest)(nil),  // 2: fqdn.ResolveRequest
}
var file_fqdn_fqdn_proto_depIdxs = []int32{
	2, // 0: fqdn.ResolveRequests.reqs:type_name -> fqdn.ResolveRequest
	0, // 1: fqdn.ResolveRequest.src:type_name -> fqdn.ResolveRequest.Src
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_fqdn_fqdn_proto_init() }
func file_fqdn_fqdn_proto_init() {
	if File_fqdn_fqdn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fqdn_fqdn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequests); i {
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
		file_fqdn_fqdn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequest); i {
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
			RawDescriptor: file_fqdn_fqdn_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fqdn_fqdn_proto_goTypes,
		DependencyIndexes: file_fqdn_fqdn_proto_depIdxs,
		EnumInfos:         file_fqdn_fqdn_proto_enumTypes,
		MessageInfos:      file_fqdn_fqdn_proto_msgTypes,
	}.Build()
	File_fqdn_fqdn_proto = out.File
	file_fqdn_fqdn_proto_rawDesc = nil
	file_fqdn_fqdn_proto_goTypes = nil
	file_fqdn_fqdn_proto_depIdxs = nil
}
