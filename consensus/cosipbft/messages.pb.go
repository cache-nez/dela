// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package cosipbft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type ChangeSet struct {
	Leader               uint32   `protobuf:"varint,1,opt,name=leader,proto3" json:"leader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeSet) Reset()         { *m = ChangeSet{} }
func (m *ChangeSet) String() string { return proto.CompactTextString(m) }
func (*ChangeSet) ProtoMessage()    {}
func (*ChangeSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *ChangeSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeSet.Unmarshal(m, b)
}
func (m *ChangeSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeSet.Marshal(b, m, deterministic)
}
func (m *ChangeSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeSet.Merge(m, src)
}
func (m *ChangeSet) XXX_Size() int {
	return xxx_messageInfo_ChangeSet.Size(m)
}
func (m *ChangeSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeSet.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeSet proto.InternalMessageInfo

func (m *ChangeSet) GetLeader() uint32 {
	if m != nil {
		return m.Leader
	}
	return 0
}

// ForwardLinkProto is the message representing a forward link between two
// proposals. It contains both hash and the prepare and commit signatures.
type ForwardLinkProto struct {
	From                 []byte     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []byte     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Prepare              *any.Any   `protobuf:"bytes,3,opt,name=prepare,proto3" json:"prepare,omitempty"`
	Commit               *any.Any   `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	ChangeSet            *ChangeSet `protobuf:"bytes,5,opt,name=changeSet,proto3" json:"changeSet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ForwardLinkProto) Reset()         { *m = ForwardLinkProto{} }
func (m *ForwardLinkProto) String() string { return proto.CompactTextString(m) }
func (*ForwardLinkProto) ProtoMessage()    {}
func (*ForwardLinkProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *ForwardLinkProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardLinkProto.Unmarshal(m, b)
}
func (m *ForwardLinkProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardLinkProto.Marshal(b, m, deterministic)
}
func (m *ForwardLinkProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardLinkProto.Merge(m, src)
}
func (m *ForwardLinkProto) XXX_Size() int {
	return xxx_messageInfo_ForwardLinkProto.Size(m)
}
func (m *ForwardLinkProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardLinkProto.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardLinkProto proto.InternalMessageInfo

func (m *ForwardLinkProto) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ForwardLinkProto) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *ForwardLinkProto) GetPrepare() *any.Any {
	if m != nil {
		return m.Prepare
	}
	return nil
}

func (m *ForwardLinkProto) GetCommit() *any.Any {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *ForwardLinkProto) GetChangeSet() *ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

// ChainProto is the message representing a list of forward links that creates
// a verifiable chain.
type ChainProto struct {
	Links                []*ForwardLinkProto `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ChainProto) Reset()         { *m = ChainProto{} }
func (m *ChainProto) String() string { return proto.CompactTextString(m) }
func (*ChainProto) ProtoMessage()    {}
func (*ChainProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *ChainProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainProto.Unmarshal(m, b)
}
func (m *ChainProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainProto.Marshal(b, m, deterministic)
}
func (m *ChainProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainProto.Merge(m, src)
}
func (m *ChainProto) XXX_Size() int {
	return xxx_messageInfo_ChainProto.Size(m)
}
func (m *ChainProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainProto.DiscardUnknown(m)
}

var xxx_messageInfo_ChainProto proto.InternalMessageInfo

func (m *ChainProto) GetLinks() []*ForwardLinkProto {
	if m != nil {
		return m.Links
	}
	return nil
}

// PrepareRequest is the message sent to start a consensus for a proposal.
type PrepareRequest struct {
	Proposal             *any.Any `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareRequest) Reset()         { *m = PrepareRequest{} }
func (m *PrepareRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareRequest) ProtoMessage()    {}
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *PrepareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareRequest.Unmarshal(m, b)
}
func (m *PrepareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareRequest.Marshal(b, m, deterministic)
}
func (m *PrepareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareRequest.Merge(m, src)
}
func (m *PrepareRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareRequest.Size(m)
}
func (m *PrepareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareRequest proto.InternalMessageInfo

func (m *PrepareRequest) GetProposal() *any.Any {
	if m != nil {
		return m.Proposal
	}
	return nil
}

// CommitRequest is the message sent to commit to a proposal.
type CommitRequest struct {
	To                   []byte   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Prepare              *any.Any `protobuf:"bytes,2,opt,name=prepare,proto3" json:"prepare,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommitRequest) Reset()         { *m = CommitRequest{} }
func (m *CommitRequest) String() string { return proto.CompactTextString(m) }
func (*CommitRequest) ProtoMessage()    {}
func (*CommitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *CommitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitRequest.Unmarshal(m, b)
}
func (m *CommitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitRequest.Marshal(b, m, deterministic)
}
func (m *CommitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitRequest.Merge(m, src)
}
func (m *CommitRequest) XXX_Size() int {
	return xxx_messageInfo_CommitRequest.Size(m)
}
func (m *CommitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommitRequest proto.InternalMessageInfo

func (m *CommitRequest) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *CommitRequest) GetPrepare() *any.Any {
	if m != nil {
		return m.Prepare
	}
	return nil
}

// PropagateRequest is the last message of a consensus process to send the valid
// forward link to participants.
type PropagateRequest struct {
	To                   []byte   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Commit               *any.Any `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PropagateRequest) Reset()         { *m = PropagateRequest{} }
func (m *PropagateRequest) String() string { return proto.CompactTextString(m) }
func (*PropagateRequest) ProtoMessage()    {}
func (*PropagateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{5}
}

func (m *PropagateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PropagateRequest.Unmarshal(m, b)
}
func (m *PropagateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PropagateRequest.Marshal(b, m, deterministic)
}
func (m *PropagateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropagateRequest.Merge(m, src)
}
func (m *PropagateRequest) XXX_Size() int {
	return xxx_messageInfo_PropagateRequest.Size(m)
}
func (m *PropagateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PropagateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PropagateRequest proto.InternalMessageInfo

func (m *PropagateRequest) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PropagateRequest) GetCommit() *any.Any {
	if m != nil {
		return m.Commit
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeSet)(nil), "cosipbft.ChangeSet")
	proto.RegisterType((*ForwardLinkProto)(nil), "cosipbft.ForwardLinkProto")
	proto.RegisterType((*ChainProto)(nil), "cosipbft.ChainProto")
	proto.RegisterType((*PrepareRequest)(nil), "cosipbft.PrepareRequest")
	proto.RegisterType((*CommitRequest)(nil), "cosipbft.CommitRequest")
	proto.RegisterType((*PropagateRequest)(nil), "cosipbft.PropagateRequest")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xfa, 0x61, 0x3b, 0xfd, 0xa0, 0xac, 0x22, 0xb1, 0xa7, 0x12, 0x2f, 0x3d, 0xc8,
	0xb6, 0xd6, 0xa3, 0x20, 0xd8, 0x82, 0x27, 0xc1, 0x10, 0x6f, 0xde, 0xb6, 0xed, 0x34, 0x0d, 0x4d,
	0x32, 0xeb, 0xee, 0x16, 0xe9, 0x2f, 0xf4, 0x6f, 0x09, 0x9b, 0x6e, 0x0b, 0x62, 0xd5, 0xdb, 0xee,
	0xf0, 0xcc, 0xf0, 0xbe, 0x0f, 0x74, 0x73, 0xd4, 0x5a, 0x24, 0xa8, 0xb9, 0x54, 0x64, 0x88, 0x35,
	0x16, 0xa4, 0x53, 0x39, 0x5f, 0x99, 0xfe, 0x55, 0x42, 0x94, 0x64, 0x38, 0xb2, 0xf3, 0xf9, 0x76,
	0x35, 0x12, 0xc5, 0xae, 0x84, 0xc2, 0x6b, 0x68, 0xce, 0xd6, 0xa2, 0x48, 0xf0, 0x15, 0x0d, 0xbb,
	0x84, 0x7a, 0x86, 0x62, 0x89, 0x2a, 0xf0, 0x06, 0xde, 0xb0, 0x13, 0xef, 0x7f, 0xe1, 0xa7, 0x07,
	0xbd, 0x27, 0x52, 0x1f, 0x42, 0x2d, 0x9f, 0xd3, 0x62, 0x13, 0xd9, 0xf3, 0x0c, 0xaa, 0x2b, 0x45,
	0xb9, 0x45, 0xdb, 0xb1, 0x7d, 0xb3, 0x2e, 0xf8, 0x86, 0x02, 0xdf, 0x4e, 0x7c, 0x43, 0x8c, 0xc3,
	0x99, 0x54, 0x28, 0x85, 0xc2, 0xa0, 0x32, 0xf0, 0x86, 0xad, 0xc9, 0x05, 0x2f, 0xa3, 0x70, 0x17,
	0x85, 0x3f, 0x16, 0xbb, 0xd8, 0x41, 0xec, 0x06, 0xea, 0x0b, 0xca, 0xf3, 0xd4, 0x04, 0xd5, 0x5f,
	0xf0, 0x3d, 0xc3, 0x6e, 0xa1, 0xb9, 0x70, 0xd9, 0x83, 0x9a, 0x5d, 0x38, 0xe7, 0xae, 0x34, 0x3f,
	0xd4, 0x8a, 0x8f, 0x54, 0xf8, 0x00, 0x30, 0x5b, 0x8b, 0xb4, 0x28, 0x2b, 0x8c, 0xa1, 0x96, 0xa5,
	0xc5, 0x46, 0x07, 0xde, 0xa0, 0x32, 0x6c, 0x4d, 0xfa, 0xc7, 0xe5, 0xef, 0x6d, 0xe3, 0x12, 0x0c,
	0xa7, 0xd0, 0x8d, 0xca, 0xac, 0x31, 0xbe, 0x6f, 0x51, 0x1b, 0x36, 0x86, 0x86, 0x54, 0x24, 0x49,
	0x8b, 0xcc, 0xaa, 0x38, 0x15, 0xfa, 0x40, 0x85, 0x2f, 0xd0, 0x99, 0xd9, 0x02, 0xee, 0x44, 0x69,
	0xcd, 0xfb, 0xc9, 0x9a, 0xff, 0x0f, 0x6b, 0x61, 0x04, 0xbd, 0x48, 0x91, 0x14, 0x89, 0x30, 0x78,
	0xea, 0xe6, 0xd1, 0xac, 0xff, 0xb7, 0xd9, 0x69, 0xfb, 0x0d, 0xf8, 0xbd, 0x93, 0x31, 0xaf, 0x5b,
	0xe6, 0xee, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x1c, 0xee, 0xa2, 0x61, 0x02, 0x00, 0x00,
}
