// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consensus/scheme/rolldpos/endorsementpb/endorsementmanager.proto

package endorsementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	iotextypes "github.com/iotexproject/iotex-proto/golang/iotextypes"
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

type EndorserEndorsementCollection struct {
	Endorser             string                    `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	Topics               []uint32                  `protobuf:"varint,2,rep,packed,name=topics,proto3" json:"topics,omitempty"`
	Endorsements         []*iotextypes.Endorsement `protobuf:"bytes,3,rep,name=endorsements,proto3" json:"endorsements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EndorserEndorsementCollection) Reset()         { *m = EndorserEndorsementCollection{} }
func (m *EndorserEndorsementCollection) String() string { return proto.CompactTextString(m) }
func (*EndorserEndorsementCollection) ProtoMessage()    {}
func (*EndorserEndorsementCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0850e6b6c2493f2, []int{0}
}

func (m *EndorserEndorsementCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorserEndorsementCollection.Unmarshal(m, b)
}
func (m *EndorserEndorsementCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorserEndorsementCollection.Marshal(b, m, deterministic)
}
func (m *EndorserEndorsementCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorserEndorsementCollection.Merge(m, src)
}
func (m *EndorserEndorsementCollection) XXX_Size() int {
	return xxx_messageInfo_EndorserEndorsementCollection.Size(m)
}
func (m *EndorserEndorsementCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorserEndorsementCollection.DiscardUnknown(m)
}

var xxx_messageInfo_EndorserEndorsementCollection proto.InternalMessageInfo

func (m *EndorserEndorsementCollection) GetEndorser() string {
	if m != nil {
		return m.Endorser
	}
	return ""
}

func (m *EndorserEndorsementCollection) GetTopics() []uint32 {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *EndorserEndorsementCollection) GetEndorsements() []*iotextypes.Endorsement {
	if m != nil {
		return m.Endorsements
	}
	return nil
}

type BlockEndorsementCollection struct {
	Blk                  *iotextypes.Block                `protobuf:"bytes,1,opt,name=blk,proto3" json:"blk,omitempty"`
	BlockMap             []*EndorserEndorsementCollection `protobuf:"bytes,2,rep,name=blockMap,proto3" json:"blockMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *BlockEndorsementCollection) Reset()         { *m = BlockEndorsementCollection{} }
func (m *BlockEndorsementCollection) String() string { return proto.CompactTextString(m) }
func (*BlockEndorsementCollection) ProtoMessage()    {}
func (*BlockEndorsementCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0850e6b6c2493f2, []int{1}
}

func (m *BlockEndorsementCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockEndorsementCollection.Unmarshal(m, b)
}
func (m *BlockEndorsementCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockEndorsementCollection.Marshal(b, m, deterministic)
}
func (m *BlockEndorsementCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockEndorsementCollection.Merge(m, src)
}
func (m *BlockEndorsementCollection) XXX_Size() int {
	return xxx_messageInfo_BlockEndorsementCollection.Size(m)
}
func (m *BlockEndorsementCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockEndorsementCollection.DiscardUnknown(m)
}

var xxx_messageInfo_BlockEndorsementCollection proto.InternalMessageInfo

func (m *BlockEndorsementCollection) GetBlk() *iotextypes.Block {
	if m != nil {
		return m.Blk
	}
	return nil
}

func (m *BlockEndorsementCollection) GetBlockMap() []*EndorserEndorsementCollection {
	if m != nil {
		return m.BlockMap
	}
	return nil
}

type EndorsementManager struct {
	BlkHash              []string                      `protobuf:"bytes,1,rep,name=blkHash,proto3" json:"blkHash,omitempty"`
	BlockEndorsements    []*BlockEndorsementCollection `protobuf:"bytes,2,rep,name=blockEndorsements,proto3" json:"blockEndorsements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *EndorsementManager) Reset()         { *m = EndorsementManager{} }
func (m *EndorsementManager) String() string { return proto.CompactTextString(m) }
func (*EndorsementManager) ProtoMessage()    {}
func (*EndorsementManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0850e6b6c2493f2, []int{2}
}

func (m *EndorsementManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndorsementManager.Unmarshal(m, b)
}
func (m *EndorsementManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndorsementManager.Marshal(b, m, deterministic)
}
func (m *EndorsementManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndorsementManager.Merge(m, src)
}
func (m *EndorsementManager) XXX_Size() int {
	return xxx_messageInfo_EndorsementManager.Size(m)
}
func (m *EndorsementManager) XXX_DiscardUnknown() {
	xxx_messageInfo_EndorsementManager.DiscardUnknown(m)
}

var xxx_messageInfo_EndorsementManager proto.InternalMessageInfo

func (m *EndorsementManager) GetBlkHash() []string {
	if m != nil {
		return m.BlkHash
	}
	return nil
}

func (m *EndorsementManager) GetBlockEndorsements() []*BlockEndorsementCollection {
	if m != nil {
		return m.BlockEndorsements
	}
	return nil
}

func init() {
	proto.RegisterType((*EndorserEndorsementCollection)(nil), "endorsementpb.endorserEndorsementCollection")
	proto.RegisterType((*BlockEndorsementCollection)(nil), "endorsementpb.blockEndorsementCollection")
	proto.RegisterType((*EndorsementManager)(nil), "endorsementpb.endorsementManager")
}

func init() {
	proto.RegisterFile("consensus/scheme/rolldpos/endorsementpb/endorsementmanager.proto", fileDescriptor_b0850e6b6c2493f2)
}

var fileDescriptor_b0850e6b6c2493f2 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x4b, 0xc3, 0x30,
	0x18, 0xa5, 0x16, 0xe6, 0x96, 0xb9, 0xc3, 0x72, 0xd0, 0x52, 0x1c, 0x94, 0x79, 0xa9, 0xa0, 0x0d,
	0xcc, 0xa3, 0x17, 0x99, 0x08, 0x43, 0xdc, 0x25, 0x17, 0xc1, 0x5b, 0x9b, 0x85, 0xb5, 0x2e, 0xcd,
	0x17, 0x92, 0x0c, 0xf4, 0x2f, 0xf0, 0xe0, 0xc5, 0x3f, 0x59, 0x96, 0x6e, 0x33, 0x9b, 0x3f, 0xf0,
	0xf8, 0xf2, 0xde, 0xfb, 0xbe, 0xf7, 0xfa, 0x15, 0xdd, 0x30, 0x90, 0x86, 0x4b, 0xb3, 0x34, 0xc4,
	0xb0, 0x92, 0xd7, 0x9c, 0x68, 0x10, 0x62, 0xa6, 0xc0, 0x10, 0x2e, 0x67, 0xa0, 0x0d, 0xaf, 0xb9,
	0xb4, 0xaa, 0xf0, 0x51, 0x9d, 0xcb, 0x7c, 0xce, 0x75, 0xa6, 0x34, 0x58, 0xc0, 0xbd, 0x1d, 0x5d,
	0x7c, 0xea, 0x5e, 0x89, 0x7d, 0x55, 0xdc, 0x90, 0x42, 0x00, 0x5b, 0xb0, 0x32, 0xaf, 0x64, 0x23,
	0x8e, 0x07, 0x3e, 0xeb, 0x19, 0x1b, 0x7a, 0xf8, 0x11, 0xa0, 0xc1, 0xfa, 0x55, 0xdf, 0x7d, 0xb1,
	0xb7, 0x20, 0x04, 0x67, 0xb6, 0x02, 0x89, 0x63, 0xd4, 0xde, 0x08, 0xa2, 0x20, 0x09, 0xd2, 0x0e,
	0xdd, 0x62, 0x7c, 0x8c, 0x5a, 0x16, 0x54, 0xc5, 0x4c, 0x74, 0x90, 0x84, 0x69, 0x8f, 0xae, 0x11,
	0xbe, 0x46, 0x47, 0xde, 0x2a, 0x13, 0x85, 0x49, 0x98, 0x76, 0x47, 0x27, 0x59, 0x05, 0x96, 0xbf,
	0xb8, 0x28, 0x99, 0xb7, 0x8c, 0xee, 0x88, 0x87, 0xef, 0x01, 0x8a, 0x5d, 0x8d, 0x9f, 0xf3, 0x9c,
	0xa1, 0xb0, 0x10, 0x0b, 0x17, 0xa5, 0x3b, 0xea, 0xfb, 0x23, 0xc7, 0x2b, 0x13, 0x5d, 0xb1, 0x78,
	0x82, 0xda, 0x6e, 0xc4, 0x34, 0x57, 0x2e, 0x5a, 0x77, 0x74, 0x91, 0xed, 0x7c, 0xb5, 0xec, 0xcf,
	0xd2, 0x74, 0xeb, 0x1e, 0xbe, 0x05, 0x08, 0x7b, 0xce, 0x69, 0x73, 0x09, 0x1c, 0xa1, 0xc3, 0x42,
	0x2c, 0x26, 0xb9, 0x29, 0xa3, 0x20, 0x09, 0xd3, 0x0e, 0xdd, 0x40, 0xfc, 0x88, 0xfa, 0xfb, 0xe9,
	0xcd, 0x3a, 0xc3, 0xf9, 0x5e, 0x86, 0xdf, 0x5b, 0xd2, 0xef, 0x33, 0xc6, 0x0f, 0x4f, 0xf7, 0xf3,
	0xca, 0x96, 0xcb, 0x22, 0x63, 0x50, 0x13, 0xd7, 0x5b, 0x69, 0x78, 0xe6, 0xcc, 0x36, 0xe0, 0x92,
	0x81, 0xe6, 0xe4, 0x9f, 0x7f, 0x57, 0xd1, 0x72, 0xf7, 0xbf, 0xfa, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0x60, 0xff, 0x04, 0x8f, 0x02, 0x00, 0x00,
}