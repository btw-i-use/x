// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: uplodr.proto

package uplodr

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

type UploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Folder   string `protobuf:"bytes,3,opt,name=folder,proto3" json:"folder,omitempty"`
}

func (x *UploadReq) Reset() {
	*x = UploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uplodr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReq) ProtoMessage() {}

func (x *UploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_uplodr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReq.ProtoReflect.Descriptor instead.
func (*UploadReq) Descriptor() ([]byte, []int) {
	return file_uplodr_proto_rawDescGZIP(), []int{0}
}

func (x *UploadReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UploadReq) GetFolder() string {
	if x != nil {
		return x.Folder
	}
	return ""
}

type UploadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Variants []*Variant `protobuf:"bytes,1,rep,name=variants,proto3" json:"variants,omitempty"`
}

func (x *UploadResp) Reset() {
	*x = UploadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uplodr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResp) ProtoMessage() {}

func (x *UploadResp) ProtoReflect() protoreflect.Message {
	mi := &file_uplodr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResp.ProtoReflect.Descriptor instead.
func (*UploadResp) Descriptor() ([]byte, []int) {
	return file_uplodr_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResp) GetVariants() []*Variant {
	if x != nil {
		return x.Variants
	}
	return nil
}

type Variant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url      string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *Variant) Reset() {
	*x = Variant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uplodr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Variant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Variant) ProtoMessage() {}

func (x *Variant) ProtoReflect() protoreflect.Message {
	mi := &file_uplodr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Variant.ProtoReflect.Descriptor instead.
func (*Variant) Descriptor() ([]byte, []int) {
	return file_uplodr_proto_rawDescGZIP(), []int{2}
}

func (x *Variant) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Variant) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type Echo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce string `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *Echo) Reset() {
	*x = Echo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uplodr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Echo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Echo) ProtoMessage() {}

func (x *Echo) ProtoReflect() protoreflect.Message {
	mi := &file_uplodr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Echo.ProtoReflect.Descriptor instead.
func (*Echo) Descriptor() ([]byte, []int) {
	return file_uplodr_proto_rawDescGZIP(), []int{3}
}

func (x *Echo) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

var File_uplodr_proto protoreflect.FileDescriptor

var file_uplodr_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x64, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c,
	0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78,
	0x2e, 0x78, 0x65, 0x64, 0x6e, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x64, 0x72, 0x22, 0x54, 0x0a, 0x09,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x22, 0x4f, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x41, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x78, 0x65, 0x64, 0x6e, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x64,
	0x72, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x73, 0x22, 0x38, 0x0a, 0x07, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x1c, 0x0a,
	0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x32, 0xb4, 0x01, 0x0a, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x4e, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e,
	0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78,
	0x2e, 0x78, 0x65, 0x64, 0x6e, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x64, 0x72, 0x2e, 0x45, 0x63, 0x68,
	0x6f, 0x1a, 0x22, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x2e, 0x78, 0x2e, 0x78, 0x65, 0x64, 0x6e, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x64, 0x72,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x5b, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x27, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x2e, 0x78, 0x2e, 0x78, 0x65, 0x64, 0x6e, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x64, 0x72, 0x2e, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69,
	0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x78, 0x65, 0x64, 0x6e,
	0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x64, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x1f, 0x5a, 0x1d, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x2f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x70, 0x6c,
	0x6f, 0x64, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uplodr_proto_rawDescOnce sync.Once
	file_uplodr_proto_rawDescData = file_uplodr_proto_rawDesc
)

func file_uplodr_proto_rawDescGZIP() []byte {
	file_uplodr_proto_rawDescOnce.Do(func() {
		file_uplodr_proto_rawDescData = protoimpl.X.CompressGZIP(file_uplodr_proto_rawDescData)
	})
	return file_uplodr_proto_rawDescData
}

var file_uplodr_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_uplodr_proto_goTypes = []interface{}{
	(*UploadReq)(nil),  // 0: within.website.x.xedn.uplodr.UploadReq
	(*UploadResp)(nil), // 1: within.website.x.xedn.uplodr.UploadResp
	(*Variant)(nil),    // 2: within.website.x.xedn.uplodr.Variant
	(*Echo)(nil),       // 3: within.website.x.xedn.uplodr.Echo
}
var file_uplodr_proto_depIdxs = []int32{
	2, // 0: within.website.x.xedn.uplodr.UploadResp.variants:type_name -> within.website.x.xedn.uplodr.Variant
	3, // 1: within.website.x.xedn.uplodr.Image.Ping:input_type -> within.website.x.xedn.uplodr.Echo
	0, // 2: within.website.x.xedn.uplodr.Image.Upload:input_type -> within.website.x.xedn.uplodr.UploadReq
	3, // 3: within.website.x.xedn.uplodr.Image.Ping:output_type -> within.website.x.xedn.uplodr.Echo
	1, // 4: within.website.x.xedn.uplodr.Image.Upload:output_type -> within.website.x.xedn.uplodr.UploadResp
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_uplodr_proto_init() }
func file_uplodr_proto_init() {
	if File_uplodr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uplodr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReq); i {
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
		file_uplodr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResp); i {
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
		file_uplodr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Variant); i {
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
		file_uplodr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Echo); i {
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
			RawDescriptor: file_uplodr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uplodr_proto_goTypes,
		DependencyIndexes: file_uplodr_proto_depIdxs,
		MessageInfos:      file_uplodr_proto_msgTypes,
	}.Build()
	File_uplodr_proto = out.File
	file_uplodr_proto_rawDesc = nil
	file_uplodr_proto_goTypes = nil
	file_uplodr_proto_depIdxs = nil
}
