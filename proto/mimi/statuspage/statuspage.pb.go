// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: statuspage.proto

package statuspage

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

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Unsubscribe   string `protobuf:"bytes,1,opt,name=unsubscribe,proto3" json:"unsubscribe,omitempty"`
	Documentation string `protobuf:"bytes,2,opt,name=documentation,proto3" json:"documentation,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statuspage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_statuspage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_statuspage_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetUnsubscribe() string {
	if x != nil {
		return x.Unsubscribe
	}
	return ""
}

func (x *Meta) GetDocumentation() string {
	if x != nil {
		return x.Documentation
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StatusIndicator   string `protobuf:"bytes,2,opt,name=status_indicator,json=statusIndicator,proto3" json:"status_indicator,omitempty"`
	StatusDescription string `protobuf:"bytes,3,opt,name=status_description,json=statusDescription,proto3" json:"status_description,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statuspage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_statuspage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_statuspage_proto_rawDescGZIP(), []int{1}
}

func (x *Page) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Page) GetStatusIndicator() string {
	if x != nil {
		return x.StatusIndicator
	}
	return ""
}

func (x *Page) GetStatusDescription() string {
	if x != nil {
		return x.StatusDescription
	}
	return ""
}

type ComponentUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt   string `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	NewStatus   string `protobuf:"bytes,2,opt,name=new_status,json=newStatus,proto3" json:"new_status,omitempty"`
	OldStatus   string `protobuf:"bytes,3,opt,name=old_status,json=oldStatus,proto3" json:"old_status,omitempty"`
	Id          string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	ComponentId string `protobuf:"bytes,5,opt,name=component_id,json=componentId,proto3" json:"component_id,omitempty"`
}

func (x *ComponentUpdate) Reset() {
	*x = ComponentUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statuspage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentUpdate) ProtoMessage() {}

func (x *ComponentUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_statuspage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentUpdate.ProtoReflect.Descriptor instead.
func (*ComponentUpdate) Descriptor() ([]byte, []int) {
	return file_statuspage_proto_rawDescGZIP(), []int{2}
}

func (x *ComponentUpdate) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ComponentUpdate) GetNewStatus() string {
	if x != nil {
		return x.NewStatus
	}
	return ""
}

func (x *ComponentUpdate) GetOldStatus() string {
	if x != nil {
		return x.OldStatus
	}
	return ""
}

func (x *ComponentUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ComponentUpdate) GetComponentId() string {
	if x != nil {
		return x.ComponentId
	}
	return ""
}

type Component struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt string `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Component) Reset() {
	*x = Component{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statuspage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Component) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Component) ProtoMessage() {}

func (x *Component) ProtoReflect() protoreflect.Message {
	mi := &file_statuspage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Component.ProtoReflect.Descriptor instead.
func (*Component) Descriptor() ([]byte, []int) {
	return file_statuspage_proto_rawDescGZIP(), []int{3}
}

func (x *Component) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Component) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Component) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Component) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type IncidentUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body               string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt          string `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	DisplayAt          string `protobuf:"bytes,3,opt,name=display_at,json=displayAt,proto3" json:"display_at,omitempty"`
	Status             string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	TwitterUpdatedAt   string `protobuf:"bytes,5,opt,name=twitter_updated_at,json=twitterUpdatedAt,proto3" json:"twitter_updated_at,omitempty"`
	UpdatedAt          string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	WantsTwitterUpdate bool   `protobuf:"varint,7,opt,name=wants_twitter_update,json=wantsTwitterUpdate,proto3" json:"wants_twitter_update,omitempty"`
	Id                 string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
	IncidentId         string `protobuf:"bytes,9,opt,name=incident_id,json=incidentId,proto3" json:"incident_id,omitempty"`
}

func (x *IncidentUpdate) Reset() {
	*x = IncidentUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statuspage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncidentUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncidentUpdate) ProtoMessage() {}

func (x *IncidentUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_statuspage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncidentUpdate.ProtoReflect.Descriptor instead.
func (*IncidentUpdate) Descriptor() ([]byte, []int) {
	return file_statuspage_proto_rawDescGZIP(), []int{4}
}

func (x *IncidentUpdate) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *IncidentUpdate) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *IncidentUpdate) GetDisplayAt() string {
	if x != nil {
		return x.DisplayAt
	}
	return ""
}

func (x *IncidentUpdate) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IncidentUpdate) GetTwitterUpdatedAt() string {
	if x != nil {
		return x.TwitterUpdatedAt
	}
	return ""
}

func (x *IncidentUpdate) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *IncidentUpdate) GetWantsTwitterUpdate() bool {
	if x != nil {
		return x.WantsTwitterUpdate
	}
	return false
}

func (x *IncidentUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IncidentUpdate) GetIncidentId() string {
	if x != nil {
		return x.IncidentId
	}
	return ""
}

type Incident struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Backfilled                    bool              `protobuf:"varint,1,opt,name=backfilled,proto3" json:"backfilled,omitempty"`
	CreatedAt                     string            `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Impact                        string            `protobuf:"bytes,3,opt,name=impact,proto3" json:"impact,omitempty"`
	ImpactOverride                string            `protobuf:"bytes,4,opt,name=impact_override,json=impactOverride,proto3" json:"impact_override,omitempty"`
	MonitoringAt                  string            `protobuf:"bytes,5,opt,name=monitoring_at,json=monitoringAt,proto3" json:"monitoring_at,omitempty"`
	PostmortemBody                string            `protobuf:"bytes,6,opt,name=postmortem_body,json=postmortemBody,proto3" json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt   string            `protobuf:"bytes,7,opt,name=postmortem_body_last_updated_at,json=postmortemBodyLastUpdatedAt,proto3" json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored             bool              `protobuf:"varint,8,opt,name=postmortem_ignored,json=postmortemIgnored,proto3" json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers bool              `protobuf:"varint,9,opt,name=postmortem_notified_subscribers,json=postmortemNotifiedSubscribers,proto3" json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter     bool              `protobuf:"varint,10,opt,name=postmortem_notified_twitter,json=postmortemNotifiedTwitter,proto3" json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt         string            `protobuf:"bytes,11,opt,name=postmortem_published_at,json=postmortemPublishedAt,proto3" json:"postmortem_published_at,omitempty"`
	ResolvedAt                    string            `protobuf:"bytes,12,opt,name=resolved_at,json=resolvedAt,proto3" json:"resolved_at,omitempty"`
	ScheduledAutoTransition       bool              `protobuf:"varint,13,opt,name=scheduled_auto_transition,json=scheduledAutoTransition,proto3" json:"scheduled_auto_transition,omitempty"`
	ScheduledFor                  string            `protobuf:"bytes,14,opt,name=scheduled_for,json=scheduledFor,proto3" json:"scheduled_for,omitempty"`
	ScheduledRemindPrior          bool              `protobuf:"varint,15,opt,name=scheduled_remind_prior,json=scheduledRemindPrior,proto3" json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt           string            `protobuf:"bytes,16,opt,name=scheduled_reminded_at,json=scheduledRemindedAt,proto3" json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil                string            `protobuf:"bytes,17,opt,name=scheduled_until,json=scheduledUntil,proto3" json:"scheduled_until,omitempty"`
	Shortlink                     string            `protobuf:"bytes,18,opt,name=shortlink,proto3" json:"shortlink,omitempty"`
	Status                        string            `protobuf:"bytes,19,opt,name=status,proto3" json:"status,omitempty"`
	UpdatedAt                     string            `protobuf:"bytes,20,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id                            string            `protobuf:"bytes,21,opt,name=id,proto3" json:"id,omitempty"`
	OrganizationId                string            `protobuf:"bytes,22,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	IncidentUpdates               []*IncidentUpdate `protobuf:"bytes,23,rep,name=incident_updates,json=incidentUpdates,proto3" json:"incident_updates,omitempty"`
	Name                          string            `protobuf:"bytes,24,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Incident) Reset() {
	*x = Incident{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statuspage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Incident) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Incident) ProtoMessage() {}

func (x *Incident) ProtoReflect() protoreflect.Message {
	mi := &file_statuspage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Incident.ProtoReflect.Descriptor instead.
func (*Incident) Descriptor() ([]byte, []int) {
	return file_statuspage_proto_rawDescGZIP(), []int{5}
}

func (x *Incident) GetBackfilled() bool {
	if x != nil {
		return x.Backfilled
	}
	return false
}

func (x *Incident) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Incident) GetImpact() string {
	if x != nil {
		return x.Impact
	}
	return ""
}

func (x *Incident) GetImpactOverride() string {
	if x != nil {
		return x.ImpactOverride
	}
	return ""
}

func (x *Incident) GetMonitoringAt() string {
	if x != nil {
		return x.MonitoringAt
	}
	return ""
}

func (x *Incident) GetPostmortemBody() string {
	if x != nil {
		return x.PostmortemBody
	}
	return ""
}

func (x *Incident) GetPostmortemBodyLastUpdatedAt() string {
	if x != nil {
		return x.PostmortemBodyLastUpdatedAt
	}
	return ""
}

func (x *Incident) GetPostmortemIgnored() bool {
	if x != nil {
		return x.PostmortemIgnored
	}
	return false
}

func (x *Incident) GetPostmortemNotifiedSubscribers() bool {
	if x != nil {
		return x.PostmortemNotifiedSubscribers
	}
	return false
}

func (x *Incident) GetPostmortemNotifiedTwitter() bool {
	if x != nil {
		return x.PostmortemNotifiedTwitter
	}
	return false
}

func (x *Incident) GetPostmortemPublishedAt() string {
	if x != nil {
		return x.PostmortemPublishedAt
	}
	return ""
}

func (x *Incident) GetResolvedAt() string {
	if x != nil {
		return x.ResolvedAt
	}
	return ""
}

func (x *Incident) GetScheduledAutoTransition() bool {
	if x != nil {
		return x.ScheduledAutoTransition
	}
	return false
}

func (x *Incident) GetScheduledFor() string {
	if x != nil {
		return x.ScheduledFor
	}
	return ""
}

func (x *Incident) GetScheduledRemindPrior() bool {
	if x != nil {
		return x.ScheduledRemindPrior
	}
	return false
}

func (x *Incident) GetScheduledRemindedAt() string {
	if x != nil {
		return x.ScheduledRemindedAt
	}
	return ""
}

func (x *Incident) GetScheduledUntil() string {
	if x != nil {
		return x.ScheduledUntil
	}
	return ""
}

func (x *Incident) GetShortlink() string {
	if x != nil {
		return x.Shortlink
	}
	return ""
}

func (x *Incident) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Incident) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Incident) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Incident) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Incident) GetIncidentUpdates() []*IncidentUpdate {
	if x != nil {
		return x.IncidentUpdates
	}
	return nil
}

func (x *Incident) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StatusUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta            *Meta            `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Page            *Page            `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Incident        *Incident        `protobuf:"bytes,3,opt,name=incident,proto3" json:"incident,omitempty"`
	Component       *Component       `protobuf:"bytes,4,opt,name=component,proto3" json:"component,omitempty"`
	ComponentUpdate *ComponentUpdate `protobuf:"bytes,5,opt,name=component_update,json=componentUpdate,proto3" json:"component_update,omitempty"`
}

func (x *StatusUpdate) Reset() {
	*x = StatusUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_statuspage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUpdate) ProtoMessage() {}

func (x *StatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_statuspage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUpdate.ProtoReflect.Descriptor instead.
func (*StatusUpdate) Descriptor() ([]byte, []int) {
	return file_statuspage_proto_rawDescGZIP(), []int{6}
}

func (x *StatusUpdate) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *StatusUpdate) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *StatusUpdate) GetIncident() *Incident {
	if x != nil {
		return x.Incident
	}
	return nil
}

func (x *StatusUpdate) GetComponent() *Component {
	if x != nil {
		return x.Component
	}
	return nil
}

func (x *StatusUpdate) GetComponentUpdate() *ComponentUpdate {
	if x != nil {
		return x.ComponentUpdate
	}
	return nil
}

var File_statuspage_proto protoreflect.FileDescriptor

var file_statuspage_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x70, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x75, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x70, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x64, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6c, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xaa, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x12,
	0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x77, 0x61, 0x6e,
	0x74, 0x73, 0x5f, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x77, 0x61, 0x6e, 0x74, 0x73, 0x54, 0x77,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xa1, 0x08, 0x0a,
	0x08, 0x49, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x63,
	0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x62,
	0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x4f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74, 0x65, 0x6d, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72,
	0x74, 0x65, 0x6d, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x44, 0x0a, 0x1f, 0x70, 0x6f, 0x73, 0x74, 0x6d,
	0x6f, 0x72, 0x74, 0x65, 0x6d, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1b, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74, 0x65, 0x6d, 0x42, 0x6f, 0x64, 0x79,
	0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a,
	0x12, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x67, 0x6e, 0x6f,
	0x72, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x6f, 0x73, 0x74, 0x6d,
	0x6f, 0x72, 0x74, 0x65, 0x6d, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x1f,
	0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1d, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74, 0x65,
	0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74,
	0x65, 0x6d, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x77, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x70, 0x6f, 0x73, 0x74, 0x6d,
	0x6f, 0x72, 0x74, 0x65, 0x6d, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x54, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x17, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74,
	0x65, 0x6d, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x70, 0x6f, 0x73, 0x74, 0x6d, 0x6f, 0x72, 0x74, 0x65,
	0x6d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3a, 0x0a,
	0x19, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x41, 0x75, 0x74, 0x6f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x34,
	0x0a, 0x16, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x50,
	0x72, 0x69, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x55, 0x6e, 0x74, 0x69,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x5b, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x77, 0x69, 0x74, 0x68,
	0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d,
	0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xf7, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x3a, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x3a, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d,
	0x69, 0x6d, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x08, 0x69, 0x6e, 0x63,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d,
	0x69, 0x6d, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x49,
	0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x69, 0x6e, 0x63, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x12, 0x49, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65,
	0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e,
	0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e, 0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x32, 0x58, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x04, 0x50, 0x6f, 0x6b, 0x65, 0x12, 0x2e, 0x2e, 0x77,
	0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x78, 0x2e,
	0x6d, 0x69, 0x6d, 0x69, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x28, 0x5a, 0x26, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2e, 0x77,
	0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x69, 0x6d, 0x69, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x70, 0x61, 0x67, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_statuspage_proto_rawDescOnce sync.Once
	file_statuspage_proto_rawDescData = file_statuspage_proto_rawDesc
)

func file_statuspage_proto_rawDescGZIP() []byte {
	file_statuspage_proto_rawDescOnce.Do(func() {
		file_statuspage_proto_rawDescData = protoimpl.X.CompressGZIP(file_statuspage_proto_rawDescData)
	})
	return file_statuspage_proto_rawDescData
}

var file_statuspage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_statuspage_proto_goTypes = []any{
	(*Meta)(nil),            // 0: within.website.x.mimi.statuspage.Meta
	(*Page)(nil),            // 1: within.website.x.mimi.statuspage.Page
	(*ComponentUpdate)(nil), // 2: within.website.x.mimi.statuspage.ComponentUpdate
	(*Component)(nil),       // 3: within.website.x.mimi.statuspage.Component
	(*IncidentUpdate)(nil),  // 4: within.website.x.mimi.statuspage.IncidentUpdate
	(*Incident)(nil),        // 5: within.website.x.mimi.statuspage.Incident
	(*StatusUpdate)(nil),    // 6: within.website.x.mimi.statuspage.StatusUpdate
	(*emptypb.Empty)(nil),   // 7: google.protobuf.Empty
}
var file_statuspage_proto_depIdxs = []int32{
	4, // 0: within.website.x.mimi.statuspage.Incident.incident_updates:type_name -> within.website.x.mimi.statuspage.IncidentUpdate
	0, // 1: within.website.x.mimi.statuspage.StatusUpdate.meta:type_name -> within.website.x.mimi.statuspage.Meta
	1, // 2: within.website.x.mimi.statuspage.StatusUpdate.page:type_name -> within.website.x.mimi.statuspage.Page
	5, // 3: within.website.x.mimi.statuspage.StatusUpdate.incident:type_name -> within.website.x.mimi.statuspage.Incident
	3, // 4: within.website.x.mimi.statuspage.StatusUpdate.component:type_name -> within.website.x.mimi.statuspage.Component
	2, // 5: within.website.x.mimi.statuspage.StatusUpdate.component_update:type_name -> within.website.x.mimi.statuspage.ComponentUpdate
	6, // 6: within.website.x.mimi.statuspage.Update.Poke:input_type -> within.website.x.mimi.statuspage.StatusUpdate
	7, // 7: within.website.x.mimi.statuspage.Update.Poke:output_type -> google.protobuf.Empty
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_statuspage_proto_init() }
func file_statuspage_proto_init() {
	if File_statuspage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_statuspage_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
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
		file_statuspage_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Page); i {
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
		file_statuspage_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ComponentUpdate); i {
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
		file_statuspage_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Component); i {
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
		file_statuspage_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IncidentUpdate); i {
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
		file_statuspage_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Incident); i {
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
		file_statuspage_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*StatusUpdate); i {
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
			RawDescriptor: file_statuspage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_statuspage_proto_goTypes,
		DependencyIndexes: file_statuspage_proto_depIdxs,
		MessageInfos:      file_statuspage_proto_msgTypes,
	}.Build()
	File_statuspage_proto = out.File
	file_statuspage_proto_rawDesc = nil
	file_statuspage_proto_goTypes = nil
	file_statuspage_proto_depIdxs = nil
}
