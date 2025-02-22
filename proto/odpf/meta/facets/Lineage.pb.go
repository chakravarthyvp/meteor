// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: odpf/meta/facets/Lineage.proto

package facets

import (
	common "github.com/odpf/meteor/proto/odpf/meta/common"
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

type Lineage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Upstreams   []*common.Resource `protobuf:"bytes,1,rep,name=upstreams,proto3" json:"upstreams,omitempty"`
	Downstreams []*common.Resource `protobuf:"bytes,2,rep,name=downstreams,proto3" json:"downstreams,omitempty"`
}

func (x *Lineage) Reset() {
	*x = Lineage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_meta_facets_Lineage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lineage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lineage) ProtoMessage() {}

func (x *Lineage) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_meta_facets_Lineage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lineage.ProtoReflect.Descriptor instead.
func (*Lineage) Descriptor() ([]byte, []int) {
	return file_odpf_meta_facets_Lineage_proto_rawDescGZIP(), []int{0}
}

func (x *Lineage) GetUpstreams() []*common.Resource {
	if x != nil {
		return x.Upstreams
	}
	return nil
}

func (x *Lineage) GetDownstreams() []*common.Resource {
	if x != nil {
		return x.Downstreams
	}
	return nil
}

var File_odpf_meta_facets_Lineage_proto protoreflect.FileDescriptor

var file_odpf_meta_facets_Lineage_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x66, 0x61, 0x63, 0x65,
	0x74, 0x73, 0x2f, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x66, 0x61, 0x63, 0x65,
	0x74, 0x73, 0x1a, 0x1f, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x07, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x64, 0x6f, 0x77,
	0x6e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x77, 0x6e,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x42, 0x49, 0x0a, 0x1a, 0x69, 0x6f, 0x2e, 0x6f, 0x64,
	0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x66,
	0x61, 0x63, 0x65, 0x74, 0x73, 0x42, 0x07, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x67, 0x65, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x64, 0x70, 0x66, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x66, 0x61, 0x63, 0x65,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_odpf_meta_facets_Lineage_proto_rawDescOnce sync.Once
	file_odpf_meta_facets_Lineage_proto_rawDescData = file_odpf_meta_facets_Lineage_proto_rawDesc
)

func file_odpf_meta_facets_Lineage_proto_rawDescGZIP() []byte {
	file_odpf_meta_facets_Lineage_proto_rawDescOnce.Do(func() {
		file_odpf_meta_facets_Lineage_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_meta_facets_Lineage_proto_rawDescData)
	})
	return file_odpf_meta_facets_Lineage_proto_rawDescData
}

var file_odpf_meta_facets_Lineage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_odpf_meta_facets_Lineage_proto_goTypes = []interface{}{
	(*Lineage)(nil),         // 0: odpf.meta.facets.Lineage
	(*common.Resource)(nil), // 1: odpf.meta.common.Resource
}
var file_odpf_meta_facets_Lineage_proto_depIdxs = []int32{
	1, // 0: odpf.meta.facets.Lineage.upstreams:type_name -> odpf.meta.common.Resource
	1, // 1: odpf.meta.facets.Lineage.downstreams:type_name -> odpf.meta.common.Resource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_odpf_meta_facets_Lineage_proto_init() }
func file_odpf_meta_facets_Lineage_proto_init() {
	if File_odpf_meta_facets_Lineage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_odpf_meta_facets_Lineage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lineage); i {
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
			RawDescriptor: file_odpf_meta_facets_Lineage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_meta_facets_Lineage_proto_goTypes,
		DependencyIndexes: file_odpf_meta_facets_Lineage_proto_depIdxs,
		MessageInfos:      file_odpf_meta_facets_Lineage_proto_msgTypes,
	}.Build()
	File_odpf_meta_facets_Lineage_proto = out.File
	file_odpf_meta_facets_Lineage_proto_rawDesc = nil
	file_odpf_meta_facets_Lineage_proto_goTypes = nil
	file_odpf_meta_facets_Lineage_proto_depIdxs = nil
}
