// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/topic.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CreateTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicRequest) Reset()         { *m = CreateTopicRequest{} }
func (m *CreateTopicRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTopicRequest) ProtoMessage()    {}
func (*CreateTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{0}
}

func (m *CreateTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicRequest.Unmarshal(m, b)
}
func (m *CreateTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicRequest.Marshal(b, m, deterministic)
}
func (m *CreateTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicRequest.Merge(m, src)
}
func (m *CreateTopicRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTopicRequest.Size(m)
}
func (m *CreateTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicRequest proto.InternalMessageInfo

type CreateTopicResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTopicResponse) Reset()         { *m = CreateTopicResponse{} }
func (m *CreateTopicResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTopicResponse) ProtoMessage()    {}
func (*CreateTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{1}
}

func (m *CreateTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTopicResponse.Unmarshal(m, b)
}
func (m *CreateTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTopicResponse.Marshal(b, m, deterministic)
}
func (m *CreateTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTopicResponse.Merge(m, src)
}
func (m *CreateTopicResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTopicResponse.Size(m)
}
func (m *CreateTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTopicResponse proto.InternalMessageInfo

type ListTopicsRequest struct {
	Cluster              string   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsRequest) Reset()         { *m = ListTopicsRequest{} }
func (m *ListTopicsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTopicsRequest) ProtoMessage()    {}
func (*ListTopicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{2}
}

func (m *ListTopicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsRequest.Unmarshal(m, b)
}
func (m *ListTopicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsRequest.Marshal(b, m, deterministic)
}
func (m *ListTopicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsRequest.Merge(m, src)
}
func (m *ListTopicsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTopicsRequest.Size(m)
}
func (m *ListTopicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsRequest proto.InternalMessageInfo

func (m *ListTopicsRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type ListTopicsResponse struct {
	Topics               []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsResponse) Reset()         { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()    {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{3}
}

func (m *ListTopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsResponse.Unmarshal(m, b)
}
func (m *ListTopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsResponse.Marshal(b, m, deterministic)
}
func (m *ListTopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsResponse.Merge(m, src)
}
func (m *ListTopicsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTopicsResponse.Size(m)
}
func (m *ListTopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsResponse proto.InternalMessageInfo

func (m *ListTopicsResponse) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

type GetTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopicRequest) Reset()         { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()    {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{4}
}

func (m *GetTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopicRequest.Unmarshal(m, b)
}
func (m *GetTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopicRequest.Marshal(b, m, deterministic)
}
func (m *GetTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopicRequest.Merge(m, src)
}
func (m *GetTopicRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopicRequest.Size(m)
}
func (m *GetTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopicRequest proto.InternalMessageInfo

type DeleteTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTopicRequest) Reset()         { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()    {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{5}
}

func (m *DeleteTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTopicRequest.Unmarshal(m, b)
}
func (m *DeleteTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTopicRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTopicRequest.Merge(m, src)
}
func (m *DeleteTopicRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTopicRequest.Size(m)
}
func (m *DeleteTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTopicRequest proto.InternalMessageInfo

type UpdateTopicRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTopicRequest) Reset()         { *m = UpdateTopicRequest{} }
func (m *UpdateTopicRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTopicRequest) ProtoMessage()    {}
func (*UpdateTopicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{6}
}

func (m *UpdateTopicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTopicRequest.Unmarshal(m, b)
}
func (m *UpdateTopicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTopicRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTopicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTopicRequest.Merge(m, src)
}
func (m *UpdateTopicRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTopicRequest.Size(m)
}
func (m *UpdateTopicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTopicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTopicRequest proto.InternalMessageInfo

type Topic struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NumReplicas          int32          `protobuf:"varint,2,opt,name=numReplicas,proto3" json:"numReplicas,omitempty"`
	NumPartitions        int32          `protobuf:"varint,3,opt,name=numPartitions,proto3" json:"numPartitions,omitempty"`
	Partitions           []*Partition   `protobuf:"bytes,4,rep,name=partitions,proto3" json:"partitions,omitempty"`
	Config               []*TopicConfig `protobuf:"bytes,5,rep,name=config,proto3" json:"config,omitempty"`
	Messages             int64          `protobuf:"varint,6,opt,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{7}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Topic) GetNumReplicas() int32 {
	if m != nil {
		return m.NumReplicas
	}
	return 0
}

func (m *Topic) GetNumPartitions() int32 {
	if m != nil {
		return m.NumPartitions
	}
	return 0
}

func (m *Topic) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *Topic) GetConfig() []*TopicConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Topic) GetMessages() int64 {
	if m != nil {
		return m.Messages
	}
	return 0
}

type TopicConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ReadOnly             bool     `protobuf:"varint,3,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	Sensitive            bool     `protobuf:"varint,4,opt,name=sensitive,proto3" json:"sensitive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicConfig) Reset()         { *m = TopicConfig{} }
func (m *TopicConfig) String() string { return proto.CompactTextString(m) }
func (*TopicConfig) ProtoMessage()    {}
func (*TopicConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{8}
}

func (m *TopicConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicConfig.Unmarshal(m, b)
}
func (m *TopicConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicConfig.Marshal(b, m, deterministic)
}
func (m *TopicConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicConfig.Merge(m, src)
}
func (m *TopicConfig) XXX_Size() int {
	return xxx_messageInfo_TopicConfig.Size(m)
}
func (m *TopicConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TopicConfig proto.InternalMessageInfo

func (m *TopicConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopicConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TopicConfig) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *TopicConfig) GetSensitive() bool {
	if m != nil {
		return m.Sensitive
	}
	return false
}

type Partition struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	HighWatermark        int64    `protobuf:"varint,3,opt,name=high_watermark,json=highWatermark,proto3" json:"high_watermark,omitempty"`
	Leader               int32    `protobuf:"varint,4,opt,name=leader,proto3" json:"leader,omitempty"`
	Replicas             []int32  `protobuf:"varint,5,rep,packed,name=replicas,proto3" json:"replicas,omitempty"`
	Isr                  []int32  `protobuf:"varint,6,rep,packed,name=isr,proto3" json:"isr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_577e618acc61581b, []int{9}
}

func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (m *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(m, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Partition) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Partition) GetHighWatermark() int64 {
	if m != nil {
		return m.HighWatermark
	}
	return 0
}

func (m *Partition) GetLeader() int32 {
	if m != nil {
		return m.Leader
	}
	return 0
}

func (m *Partition) GetReplicas() []int32 {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *Partition) GetIsr() []int32 {
	if m != nil {
		return m.Isr
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTopicRequest)(nil), "kaf.api.CreateTopicRequest")
	proto.RegisterType((*CreateTopicResponse)(nil), "kaf.api.CreateTopicResponse")
	proto.RegisterType((*ListTopicsRequest)(nil), "kaf.api.ListTopicsRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "kaf.api.ListTopicsResponse")
	proto.RegisterType((*GetTopicRequest)(nil), "kaf.api.GetTopicRequest")
	proto.RegisterType((*DeleteTopicRequest)(nil), "kaf.api.DeleteTopicRequest")
	proto.RegisterType((*UpdateTopicRequest)(nil), "kaf.api.UpdateTopicRequest")
	proto.RegisterType((*Topic)(nil), "kaf.api.Topic")
	proto.RegisterType((*TopicConfig)(nil), "kaf.api.TopicConfig")
	proto.RegisterType((*Partition)(nil), "kaf.api.Partition")
}

func init() { proto.RegisterFile("api/topic.proto", fileDescriptor_577e618acc61581b) }

var fileDescriptor_577e618acc61581b = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xeb, 0xd8, 0x4d, 0xc6, 0xb4, 0xa5, 0x43, 0xa8, 0x2c, 0x87, 0x43, 0x64, 0x01, 0xca,
	0x01, 0x1c, 0x29, 0x70, 0xaa, 0x38, 0xb5, 0x54, 0x5c, 0x90, 0x40, 0x0b, 0x08, 0x89, 0x4b, 0xb5,
	0x49, 0x26, 0xe9, 0xaa, 0xfe, 0x63, 0x77, 0x1d, 0x94, 0x37, 0xe2, 0x4d, 0x78, 0x0e, 0xde, 0xa4,
	0xf2, 0xc6, 0x71, 0x9c, 0xb8, 0x37, 0xcf, 0x37, 0xdf, 0x78, 0xe6, 0x9b, 0x9d, 0x19, 0x38, 0xe3,
	0xb9, 0x18, 0xeb, 0x2c, 0x17, 0xb3, 0x28, 0x97, 0x99, 0xce, 0xf0, 0xf8, 0x9e, 0x2f, 0x22, 0x9e,
	0x8b, 0x60, 0xb0, 0xcc, 0xb2, 0x65, 0x4c, 0x63, 0x03, 0x4f, 0x8b, 0xc5, 0x98, 0x92, 0x5c, 0xaf,
	0x37, 0xac, 0xb0, 0x0f, 0x78, 0x2d, 0x89, 0x6b, 0xfa, 0x5e, 0x86, 0x32, 0xfa, 0x5d, 0x90, 0xd2,
	0xe1, 0x73, 0x78, 0xb6, 0x87, 0xaa, 0x3c, 0x4b, 0x15, 0x85, 0x6f, 0xe1, 0xfc, 0xb3, 0x50, 0xda,
	0x80, 0xaa, 0xe2, 0xa2, 0x0f, 0xc7, 0xb3, 0xb8, 0x50, 0x9a, 0xa4, 0x6f, 0x0d, 0xad, 0x51, 0x8f,
	0x6d, 0xcd, 0xf0, 0x03, 0x60, 0x93, 0xbe, 0xf9, 0x09, 0xbe, 0x06, 0xd7, 0x94, 0xa9, 0x7c, 0x6b,
	0x68, 0x8f, 0xbc, 0xc9, 0x69, 0x54, 0x15, 0x1a, 0x6d, 0x92, 0x55, 0xde, 0xf0, 0x1c, 0xce, 0x3e,
	0x91, 0xde, 0x2b, 0xab, 0x0f, 0xf8, 0x91, 0x62, 0x3a, 0x28, 0xb6, 0x0f, 0xf8, 0x23, 0x9f, 0x1f,
	0x4a, 0xf8, 0x6f, 0x81, 0x63, 0x00, 0x44, 0xe8, 0xa4, 0x3c, 0xa1, 0xaa, 0x3a, 0xf3, 0x8d, 0x43,
	0xf0, 0xd2, 0x22, 0x61, 0x94, 0xc7, 0x62, 0xc6, 0x95, 0x7f, 0x34, 0xb4, 0x46, 0x0e, 0x6b, 0x42,
	0xf8, 0x12, 0x4e, 0xd2, 0x22, 0xf9, 0xca, 0xa5, 0x16, 0x5a, 0x64, 0xa9, 0xf2, 0x6d, 0xc3, 0xd9,
	0x07, 0x71, 0x02, 0x90, 0xef, 0x28, 0x1d, 0x23, 0x08, 0x6b, 0x41, 0x35, 0x91, 0x35, 0x58, 0xf8,
	0x06, 0xdc, 0x59, 0x96, 0x2e, 0xc4, 0xd2, 0x77, 0x0c, 0xbf, 0xbf, 0xdf, 0x80, 0x6b, 0xe3, 0x63,
	0x15, 0x07, 0x03, 0xe8, 0x26, 0xa4, 0x14, 0x5f, 0x92, 0xf2, 0xdd, 0xa1, 0x35, 0xb2, 0x59, 0x6d,
	0x87, 0x12, 0xbc, 0x46, 0xc8, 0xa3, 0x42, 0xfb, 0xe0, 0xac, 0x78, 0x5c, 0x90, 0x91, 0xd8, 0x63,
	0x1b, 0x03, 0x07, 0xd0, 0x93, 0xc4, 0xe7, 0xb7, 0x59, 0x1a, 0xaf, 0x8d, 0xb0, 0x2e, 0xeb, 0x96,
	0xc0, 0x97, 0x34, 0x5e, 0xe3, 0x0b, 0xe8, 0x29, 0x4a, 0x95, 0xd0, 0x62, 0x45, 0x7e, 0xc7, 0x38,
	0x77, 0x40, 0xf8, 0xd7, 0x82, 0x5e, 0xad, 0x0b, 0x2f, 0xc0, 0x4d, 0x8b, 0x64, 0x5a, 0xbd, 0xbd,
	0xcd, 0x2a, 0xab, 0xc4, 0xb3, 0xc5, 0x42, 0x91, 0x36, 0x79, 0x6d, 0x56, 0x59, 0xf8, 0x0a, 0x4e,
	0xef, 0xc4, 0xf2, 0xee, 0xf6, 0x0f, 0xd7, 0x24, 0x13, 0x2e, 0xef, 0x4d, 0x76, 0x9b, 0x9d, 0x94,
	0xe8, 0xcf, 0x2d, 0x58, 0x86, 0xc7, 0xc4, 0xe7, 0x24, 0x4d, 0x7e, 0x87, 0x55, 0x56, 0xd9, 0x0c,
	0xb9, 0x7d, 0xb3, 0xb2, 0x79, 0x0e, 0xab, 0x6d, 0x7c, 0x0a, 0xb6, 0x50, 0xd2, 0x77, 0x0d, 0x5c,
	0x7e, 0x4e, 0xfe, 0x1d, 0xc1, 0x13, 0xd3, 0x9f, 0x6f, 0x24, 0x57, 0x62, 0x46, 0x78, 0x09, 0x5e,
	0x63, 0xac, 0x71, 0x50, 0x37, 0xbe, 0xbd, 0x02, 0xc1, 0xc1, 0x58, 0xe2, 0x7b, 0xe8, 0x6e, 0xc7,
	0x11, 0xfd, 0xda, 0x77, 0x30, 0xa1, 0xad, 0xa8, 0x4b, 0xf0, 0x1a, 0xb3, 0xd9, 0xc8, 0xd8, 0x9e,
	0xd8, 0x56, 0xec, 0x0d, 0xc0, 0x6e, 0x7d, 0x30, 0xa8, 0xbd, 0xad, 0x15, 0x0c, 0x06, 0x8f, 0xfa,
	0xaa, 0x7d, 0xbb, 0x02, 0xaf, 0xb1, 0x34, 0x8d, 0x12, 0xda, 0xab, 0x14, 0x5c, 0x44, 0x9b, 0x5b,
	0x11, 0x6d, 0x6f, 0x45, 0x74, 0x53, 0xde, 0x8a, 0x2b, 0xe7, 0x97, 0xcd, 0x73, 0x31, 0x75, 0x0d,
	0xfc, 0xee, 0x21, 0x00, 0x00, 0xff, 0xff, 0x96, 0x9f, 0xe1, 0x9f, 0x6c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TopicServiceClient is the client API for TopicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicServiceClient interface {
	CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*Topic, error)
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
	DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type topicServiceClient struct {
	cc *grpc.ClientConn
}

func NewTopicServiceClient(cc *grpc.ClientConn) TopicServiceClient {
	return &topicServiceClient{cc}
}

func (c *topicServiceClient) CreateTopic(ctx context.Context, in *CreateTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/CreateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) GetTopic(ctx context.Context, in *GetTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/GetTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) UpdateTopic(ctx context.Context, in *UpdateTopicRequest, opts ...grpc.CallOption) (*Topic, error) {
	out := new(Topic)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/UpdateTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicServiceClient) DeleteTopic(ctx context.Context, in *DeleteTopicRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/kaf.api.TopicService/DeleteTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicServiceServer is the server API for TopicService service.
type TopicServiceServer interface {
	CreateTopic(context.Context, *CreateTopicRequest) (*Topic, error)
	GetTopic(context.Context, *GetTopicRequest) (*Topic, error)
	UpdateTopic(context.Context, *UpdateTopicRequest) (*Topic, error)
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
	DeleteTopic(context.Context, *DeleteTopicRequest) (*empty.Empty, error)
}

// UnimplementedTopicServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicServiceServer struct {
}

func (*UnimplementedTopicServiceServer) CreateTopic(ctx context.Context, req *CreateTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTopic not implemented")
}
func (*UnimplementedTopicServiceServer) GetTopic(ctx context.Context, req *GetTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopic not implemented")
}
func (*UnimplementedTopicServiceServer) UpdateTopic(ctx context.Context, req *UpdateTopicRequest) (*Topic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTopic not implemented")
}
func (*UnimplementedTopicServiceServer) ListTopics(ctx context.Context, req *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}
func (*UnimplementedTopicServiceServer) DeleteTopic(ctx context.Context, req *DeleteTopicRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTopic not implemented")
}

func RegisterTopicServiceServer(s *grpc.Server, srv TopicServiceServer) {
	s.RegisterService(&_TopicService_serviceDesc, srv)
}

func _TopicService_CreateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).CreateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/CreateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).CreateTopic(ctx, req.(*CreateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_GetTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).GetTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/GetTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).GetTopic(ctx, req.(*GetTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_UpdateTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).UpdateTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/UpdateTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).UpdateTopic(ctx, req.(*UpdateTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicService_DeleteTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicServiceServer).DeleteTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kaf.api.TopicService/DeleteTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicServiceServer).DeleteTopic(ctx, req.(*DeleteTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kaf.api.TopicService",
	HandlerType: (*TopicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTopic",
			Handler:    _TopicService_CreateTopic_Handler,
		},
		{
			MethodName: "GetTopic",
			Handler:    _TopicService_GetTopic_Handler,
		},
		{
			MethodName: "UpdateTopic",
			Handler:    _TopicService_UpdateTopic_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _TopicService_ListTopics_Handler,
		},
		{
			MethodName: "DeleteTopic",
			Handler:    _TopicService_DeleteTopic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/topic.proto",
}
