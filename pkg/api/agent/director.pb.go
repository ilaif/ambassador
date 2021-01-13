//*
// Communication between the Ambassador Agent and the Director service
// to populate the Central Edge Policy Console, which is a cloud service
// run by Datawire.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.8.0
// source: agent/director.proto

package agent

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// How Ambassador's Agent identifies itself to the CEPC
type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account ID assigned by the CEPC
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// Ambassador version
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// This Ambassador's hostname
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// License information
	License string `protobuf:"bytes,4,opt,name=license,proto3" json:"license,omitempty"`
	// The cluster ID, as determined by Ambassador
	ClusterId string `protobuf:"bytes,5,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Label or description for the user
	Label string `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_director_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_agent_director_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_agent_director_proto_rawDescGZIP(), []int{0}
}

func (x *Identity) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Identity) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Identity) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Identity) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *Identity) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Identity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

// Information that Ambassador's Agent can send to the Director
// component of the CEPC
type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity *Identity `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// okay but like what is this used for...
	// or what was it intended to be used for...
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// no longer used. leave this all behind and create a v2 api?
	Services    []*Service `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
	RawSnapshot []byte     `protobuf:"bytes,4,opt,name=raw_snapshot,json=rawSnapshot,proto3" json:"raw_snapshot,omitempty"`
	// described how the raw_snapshot was encoded
	ContentType string `protobuf:"bytes,5,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// api version of the raw snapshot
	ApiVersion string `protobuf:"bytes,6,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// when the snapshot was taken.
	// so the agentcom can throw out older snapshots
	SnapshotTs *timestamp.Timestamp `protobuf:"bytes,7,opt,name=snapshot_ts,json=snapshotTs,proto3" json:"snapshot_ts,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_director_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_agent_director_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_agent_director_proto_rawDescGZIP(), []int{1}
}

func (x *Snapshot) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Snapshot) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Snapshot) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Snapshot) GetRawSnapshot() []byte {
	if x != nil {
		return x.RawSnapshot
	}
	return nil
}

func (x *Snapshot) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Snapshot) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *Snapshot) GetSnapshotTs() *timestamp.Timestamp {
	if x != nil {
		return x.SnapshotTs
	}
	return nil
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace   string            `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Labels      map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_director_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_agent_director_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_agent_director_proto_rawDescGZIP(), []int{2}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Service) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Service) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// The Director's response to a Snapshot from the Agent
type SnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SnapshotResponse) Reset() {
	*x = SnapshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_director_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotResponse) ProtoMessage() {}

func (x *SnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_director_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotResponse.ProtoReflect.Descriptor instead.
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return file_agent_director_proto_rawDescGZIP(), []int{3}
}

// Instructions that the CEPC can send to Ambassador
type Directive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Stop sending snapshots. The default value (false) indicates that
	// snapshot should be sent.
	StopReporting bool `protobuf:"varint,2,opt,name=stop_reporting,json=stopReporting,proto3" json:"stop_reporting,omitempty"`
	// Minimum time to wait before pushing the next snapshot. The default
	// value (zero duration) indicates that the Agent should not modify
	// the existing report period.
	MinReportPeriod *duration.Duration `protobuf:"bytes,3,opt,name=min_report_period,json=minReportPeriod,proto3" json:"min_report_period,omitempty"`
	// Commands to execute
	Commands []*Command `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
}

func (x *Directive) Reset() {
	*x = Directive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_director_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Directive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directive) ProtoMessage() {}

func (x *Directive) ProtoReflect() protoreflect.Message {
	mi := &file_agent_director_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directive.ProtoReflect.Descriptor instead.
func (*Directive) Descriptor() ([]byte, []int) {
	return file_agent_director_proto_rawDescGZIP(), []int{4}
}

func (x *Directive) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Directive) GetStopReporting() bool {
	if x != nil {
		return x.StopReporting
	}
	return false
}

func (x *Directive) GetMinReportPeriod() *duration.Duration {
	if x != nil {
		return x.MinReportPeriod
	}
	return nil
}

func (x *Directive) GetCommands() []*Command {
	if x != nil {
		return x.Commands
	}
	return nil
}

// An individual instruction from the CEPC
type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Log this message if present
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_director_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_agent_director_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_agent_director_proto_rawDescGZIP(), []int{5}
}

func (x *Command) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_agent_director_proto protoreflect.FileDescriptor

var file_agent_director_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae,
	0x01, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22,
	0xa1, 0x02, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x2b, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x61, 0x77, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x54, 0x73, 0x22, 0xad, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x32, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x70, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x73,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x11,
	0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x22,
	0x23, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x73, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x34, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0f, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x12, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_agent_director_proto_rawDescOnce sync.Once
	file_agent_director_proto_rawDescData = file_agent_director_proto_rawDesc
)

func file_agent_director_proto_rawDescGZIP() []byte {
	file_agent_director_proto_rawDescOnce.Do(func() {
		file_agent_director_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_director_proto_rawDescData)
	})
	return file_agent_director_proto_rawDescData
}

var file_agent_director_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_agent_director_proto_goTypes = []interface{}{
	(*Identity)(nil),            // 0: agent.Identity
	(*Snapshot)(nil),            // 1: agent.Snapshot
	(*Service)(nil),             // 2: agent.Service
	(*SnapshotResponse)(nil),    // 3: agent.SnapshotResponse
	(*Directive)(nil),           // 4: agent.Directive
	(*Command)(nil),             // 5: agent.Command
	nil,                         // 6: agent.Service.LabelsEntry
	nil,                         // 7: agent.Service.AnnotationsEntry
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 9: google.protobuf.Duration
}
var file_agent_director_proto_depIdxs = []int32{
	0, // 0: agent.Snapshot.identity:type_name -> agent.Identity
	2, // 1: agent.Snapshot.services:type_name -> agent.Service
	8, // 2: agent.Snapshot.snapshot_ts:type_name -> google.protobuf.Timestamp
	6, // 3: agent.Service.labels:type_name -> agent.Service.LabelsEntry
	7, // 4: agent.Service.annotations:type_name -> agent.Service.AnnotationsEntry
	9, // 5: agent.Directive.min_report_period:type_name -> google.protobuf.Duration
	5, // 6: agent.Directive.commands:type_name -> agent.Command
	1, // 7: agent.Director.Report:input_type -> agent.Snapshot
	0, // 8: agent.Director.Retrieve:input_type -> agent.Identity
	3, // 9: agent.Director.Report:output_type -> agent.SnapshotResponse
	4, // 10: agent.Director.Retrieve:output_type -> agent.Directive
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_agent_director_proto_init() }
func file_agent_director_proto_init() {
	if File_agent_director_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_director_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_agent_director_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
		file_agent_director_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_agent_director_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotResponse); i {
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
		file_agent_director_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Directive); i {
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
		file_agent_director_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
			RawDescriptor: file_agent_director_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_director_proto_goTypes,
		DependencyIndexes: file_agent_director_proto_depIdxs,
		MessageInfos:      file_agent_director_proto_msgTypes,
	}.Build()
	File_agent_director_proto = out.File
	file_agent_director_proto_rawDesc = nil
	file_agent_director_proto_goTypes = nil
	file_agent_director_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DirectorClient is the client API for Director service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DirectorClient interface {
	// Report a consistent Snapshot of information to the CEPC
	Report(ctx context.Context, in *Snapshot, opts ...grpc.CallOption) (*SnapshotResponse, error)
	// Retrieve Directives from the CEPC
	Retrieve(ctx context.Context, in *Identity, opts ...grpc.CallOption) (Director_RetrieveClient, error)
}

type directorClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectorClient(cc grpc.ClientConnInterface) DirectorClient {
	return &directorClient{cc}
}

func (c *directorClient) Report(ctx context.Context, in *Snapshot, opts ...grpc.CallOption) (*SnapshotResponse, error) {
	out := new(SnapshotResponse)
	err := c.cc.Invoke(ctx, "/agent.Director/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directorClient) Retrieve(ctx context.Context, in *Identity, opts ...grpc.CallOption) (Director_RetrieveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Director_serviceDesc.Streams[0], "/agent.Director/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &directorRetrieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Director_RetrieveClient interface {
	Recv() (*Directive, error)
	grpc.ClientStream
}

type directorRetrieveClient struct {
	grpc.ClientStream
}

func (x *directorRetrieveClient) Recv() (*Directive, error) {
	m := new(Directive)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DirectorServer is the server API for Director service.
type DirectorServer interface {
	// Report a consistent Snapshot of information to the CEPC
	Report(context.Context, *Snapshot) (*SnapshotResponse, error)
	// Retrieve Directives from the CEPC
	Retrieve(*Identity, Director_RetrieveServer) error
}

// UnimplementedDirectorServer can be embedded to have forward compatible implementations.
type UnimplementedDirectorServer struct {
}

func (*UnimplementedDirectorServer) Report(context.Context, *Snapshot) (*SnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (*UnimplementedDirectorServer) Retrieve(*Identity, Director_RetrieveServer) error {
	return status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}

func RegisterDirectorServer(s *grpc.Server, srv DirectorServer) {
	s.RegisterService(&_Director_serviceDesc, srv)
}

func _Director_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Snapshot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectorServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.Director/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectorServer).Report(ctx, req.(*Snapshot))
	}
	return interceptor(ctx, in, info, handler)
}

func _Director_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Identity)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DirectorServer).Retrieve(m, &directorRetrieveServer{stream})
}

type Director_RetrieveServer interface {
	Send(*Directive) error
	grpc.ServerStream
}

type directorRetrieveServer struct {
	grpc.ServerStream
}

func (x *directorRetrieveServer) Send(m *Directive) error {
	return x.ServerStream.SendMsg(m)
}

var _Director_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.Director",
	HandlerType: (*DirectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _Director_Report_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Retrieve",
			Handler:       _Director_Retrieve_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "agent/director.proto",
}
