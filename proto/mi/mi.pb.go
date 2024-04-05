// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: mi.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MembersResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"` // required
}

func (x *MembersResp) Reset() {
	*x = MembersResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembersResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembersResp) ProtoMessage() {}

func (x *MembersResp) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembersResp.ProtoReflect.Descriptor instead.
func (*MembersResp) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{0}
}

func (x *MembersResp) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                               // required
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                            // required
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // required
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type Switch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                // required
	MemberId  string `protobuf:"bytes,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`    // required
	StartedAt string `protobuf:"bytes,3,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"` // RFC 3339, required
	EndedAt   string `protobuf:"bytes,4,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`       // RFC 3339, optional if switch is current
}

func (x *Switch) Reset() {
	*x = Switch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Switch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Switch) ProtoMessage() {}

func (x *Switch) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Switch.ProtoReflect.Descriptor instead.
func (*Switch) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{2}
}

func (x *Switch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Switch) GetMemberId() string {
	if x != nil {
		return x.MemberId
	}
	return ""
}

func (x *Switch) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

func (x *Switch) GetEndedAt() string {
	if x != nil {
		return x.EndedAt
	}
	return ""
}

type SwitchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberName string `protobuf:"bytes,1,opt,name=member_name,json=memberName,proto3" json:"member_name,omitempty"` // required
}

func (x *SwitchReq) Reset() {
	*x = SwitchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchReq) ProtoMessage() {}

func (x *SwitchReq) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchReq.ProtoReflect.Descriptor instead.
func (*SwitchReq) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{3}
}

func (x *SwitchReq) GetMemberName() string {
	if x != nil {
		return x.MemberName
	}
	return ""
}

type SwitchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Old     *Switch `protobuf:"bytes,1,opt,name=old,proto3" json:"old,omitempty"`         // required
	Current *Switch `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"` // required
}

func (x *SwitchResp) Reset() {
	*x = SwitchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchResp) ProtoMessage() {}

func (x *SwitchResp) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchResp.ProtoReflect.Descriptor instead.
func (*SwitchResp) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{4}
}

func (x *SwitchResp) GetOld() *Switch {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *SwitchResp) GetCurrent() *Switch {
	if x != nil {
		return x.Current
	}
	return nil
}

type GetSwitchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // required
}

func (x *GetSwitchReq) Reset() {
	*x = GetSwitchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSwitchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSwitchReq) ProtoMessage() {}

func (x *GetSwitchReq) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSwitchReq.ProtoReflect.Descriptor instead.
func (*GetSwitchReq) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{5}
}

func (x *GetSwitchReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FrontChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Switch *Switch `protobuf:"bytes,1,opt,name=switch,proto3" json:"switch,omitempty"` // required
	Member *Member `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"` // required
}

func (x *FrontChange) Reset() {
	*x = FrontChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontChange) ProtoMessage() {}

func (x *FrontChange) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontChange.ProtoReflect.Descriptor instead.
func (*FrontChange) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{6}
}

func (x *FrontChange) GetSwitch() *Switch {
	if x != nil {
		return x.Switch
	}
	return nil
}

func (x *FrontChange) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type ListSwitchesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"` // required
	Page  int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`   // required
}

func (x *ListSwitchesReq) Reset() {
	*x = ListSwitchesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSwitchesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSwitchesReq) ProtoMessage() {}

func (x *ListSwitchesReq) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSwitchesReq.ProtoReflect.Descriptor instead.
func (*ListSwitchesReq) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{7}
}

func (x *ListSwitchesReq) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListSwitchesReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListSwitchesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Switches []*Switch `protobuf:"bytes,1,rep,name=switches,proto3" json:"switches,omitempty"`
}

func (x *ListSwitchesResp) Reset() {
	*x = ListSwitchesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSwitchesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSwitchesResp) ProtoMessage() {}

func (x *ListSwitchesResp) ProtoReflect() protoreflect.Message {
	mi := &file_mi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSwitchesResp.ProtoReflect.Descriptor instead.
func (*ListSwitchesResp) Descriptor() ([]byte, []int) {
	return file_mi_proto_rawDescGZIP(), []int{8}
}

func (x *ListSwitchesResp) GetSwitches() []*Switch {
	if x != nil {
		return x.Switches
	}
	return nil
}

var File_mi_proto protoreflect.FileDescriptor

var file_mi_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x77, 0x69, 0x74, 0x68,
	0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0b,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77,
	0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e,
	0x6d, 0x69, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x22, 0x4b, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x22,
	0x6f, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x2c, 0x0a, 0x09, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x72,
	0x0a, 0x0a, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x03,
	0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x74, 0x68,
	0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x2e,
	0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x35, 0x0a, 0x07, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77,
	0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e,
	0x6d, 0x69, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x77, 0x0a, 0x0b, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x06,
	0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x33, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x08,
	0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e,
	0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x08, 0x73, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x32, 0x96, 0x03, 0x0a, 0x0d, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x77, 0x69, 0x74,
	0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0a,
	0x57, 0x68, 0x6f, 0x49, 0x73, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x20, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x1e,
	0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e,
	0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e,
	0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x21, 0x2e, 0x77,
	0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e,
	0x6d, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x20, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x5b, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x24, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x77, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e,
	0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x32, 0x46,
	0x0a, 0x05, 0x50, 0x4f, 0x53, 0x53, 0x45, 0x12, 0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1c, 0x5a, 0x1a, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e,
	0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x78, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x6d,
	0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mi_proto_rawDescOnce sync.Once
	file_mi_proto_rawDescData = file_mi_proto_rawDesc
)

func file_mi_proto_rawDescGZIP() []byte {
	file_mi_proto_rawDescOnce.Do(func() {
		file_mi_proto_rawDescData = protoimpl.X.CompressGZIP(file_mi_proto_rawDescData)
	})
	return file_mi_proto_rawDescData
}

var file_mi_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_mi_proto_goTypes = []interface{}{
	(*MembersResp)(nil),      // 0: within.website.x.mi.MembersResp
	(*Member)(nil),           // 1: within.website.x.mi.Member
	(*Switch)(nil),           // 2: within.website.x.mi.Switch
	(*SwitchReq)(nil),        // 3: within.website.x.mi.SwitchReq
	(*SwitchResp)(nil),       // 4: within.website.x.mi.SwitchResp
	(*GetSwitchReq)(nil),     // 5: within.website.x.mi.GetSwitchReq
	(*FrontChange)(nil),      // 6: within.website.x.mi.FrontChange
	(*ListSwitchesReq)(nil),  // 7: within.website.x.mi.ListSwitchesReq
	(*ListSwitchesResp)(nil), // 8: within.website.x.mi.ListSwitchesResp
	(*emptypb.Empty)(nil),    // 9: google.protobuf.Empty
}
var file_mi_proto_depIdxs = []int32{
	1,  // 0: within.website.x.mi.MembersResp.members:type_name -> within.website.x.mi.Member
	2,  // 1: within.website.x.mi.SwitchResp.old:type_name -> within.website.x.mi.Switch
	2,  // 2: within.website.x.mi.SwitchResp.current:type_name -> within.website.x.mi.Switch
	2,  // 3: within.website.x.mi.FrontChange.switch:type_name -> within.website.x.mi.Switch
	1,  // 4: within.website.x.mi.FrontChange.member:type_name -> within.website.x.mi.Member
	2,  // 5: within.website.x.mi.ListSwitchesResp.switches:type_name -> within.website.x.mi.Switch
	9,  // 6: within.website.x.mi.SwitchTracker.Members:input_type -> google.protobuf.Empty
	9,  // 7: within.website.x.mi.SwitchTracker.WhoIsFront:input_type -> google.protobuf.Empty
	3,  // 8: within.website.x.mi.SwitchTracker.Switch:input_type -> within.website.x.mi.SwitchReq
	5,  // 9: within.website.x.mi.SwitchTracker.GetSwitch:input_type -> within.website.x.mi.GetSwitchReq
	7,  // 10: within.website.x.mi.SwitchTracker.ListSwitches:input_type -> within.website.x.mi.ListSwitchesReq
	9,  // 11: within.website.x.mi.POSSE.RefreshBlog:input_type -> google.protobuf.Empty
	0,  // 12: within.website.x.mi.SwitchTracker.Members:output_type -> within.website.x.mi.MembersResp
	6,  // 13: within.website.x.mi.SwitchTracker.WhoIsFront:output_type -> within.website.x.mi.FrontChange
	4,  // 14: within.website.x.mi.SwitchTracker.Switch:output_type -> within.website.x.mi.SwitchResp
	6,  // 15: within.website.x.mi.SwitchTracker.GetSwitch:output_type -> within.website.x.mi.FrontChange
	8,  // 16: within.website.x.mi.SwitchTracker.ListSwitches:output_type -> within.website.x.mi.ListSwitchesResp
	9,  // 17: within.website.x.mi.POSSE.RefreshBlog:output_type -> google.protobuf.Empty
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_mi_proto_init() }
func file_mi_proto_init() {
	if File_mi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembersResp); i {
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
		file_mi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_mi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Switch); i {
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
		file_mi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchReq); i {
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
		file_mi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchResp); i {
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
		file_mi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSwitchReq); i {
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
		file_mi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontChange); i {
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
		file_mi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSwitchesReq); i {
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
		file_mi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSwitchesResp); i {
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
			RawDescriptor: file_mi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_mi_proto_goTypes,
		DependencyIndexes: file_mi_proto_depIdxs,
		MessageInfos:      file_mi_proto_msgTypes,
	}.Build()
	File_mi_proto = out.File
	file_mi_proto_rawDesc = nil
	file_mi_proto_goTypes = nil
	file_mi_proto_depIdxs = nil
}
