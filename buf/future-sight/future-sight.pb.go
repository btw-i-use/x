// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: future-sight/future-sight.proto

package future_sight

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

type NewVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *NewVersion) Reset() {
	*x = NewVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_future_sight_future_sight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewVersion) ProtoMessage() {}

func (x *NewVersion) ProtoReflect() protoreflect.Message {
	mi := &file_future_sight_future_sight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewVersion.ProtoReflect.Descriptor instead.
func (*NewVersion) Descriptor() ([]byte, []int) {
	return file_future_sight_future_sight_proto_rawDescGZIP(), []int{0}
}

func (x *NewVersion) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_future_sight_future_sight_proto protoreflect.FileDescriptor

var file_future_sight_future_sight_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73, 0x69, 0x67, 0x68, 0x74, 0x2f, 0x66,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x73, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x2e, 0x78, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x20, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x42, 0x23, 0x5a, 0x21, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x2f, 0x78, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x2d, 0x73, 0x69, 0x67, 0x68, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_future_sight_future_sight_proto_rawDescOnce sync.Once
	file_future_sight_future_sight_proto_rawDescData = file_future_sight_future_sight_proto_rawDesc
)

func file_future_sight_future_sight_proto_rawDescGZIP() []byte {
	file_future_sight_future_sight_proto_rawDescOnce.Do(func() {
		file_future_sight_future_sight_proto_rawDescData = protoimpl.X.CompressGZIP(file_future_sight_future_sight_proto_rawDescData)
	})
	return file_future_sight_future_sight_proto_rawDescData
}

var file_future_sight_future_sight_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_future_sight_future_sight_proto_goTypes = []any{
	(*NewVersion)(nil), // 0: within.website.x.future_sight.NewVersion
}
var file_future_sight_future_sight_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_future_sight_future_sight_proto_init() }
func file_future_sight_future_sight_proto_init() {
	if File_future_sight_future_sight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_future_sight_future_sight_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NewVersion); i {
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
			RawDescriptor: file_future_sight_future_sight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_future_sight_future_sight_proto_goTypes,
		DependencyIndexes: file_future_sight_future_sight_proto_depIdxs,
		MessageInfos:      file_future_sight_future_sight_proto_msgTypes,
	}.Build()
	File_future_sight_future_sight_proto = out.File
	file_future_sight_future_sight_proto_rawDesc = nil
	file_future_sight_future_sight_proto_goTypes = nil
	file_future_sight_future_sight_proto_depIdxs = nil
}
