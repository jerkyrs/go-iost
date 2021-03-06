// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/merkletree/merkle_tree.proto

package merkletree

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MerkleTree struct {
	HashList             [][]byte         `protobuf:"bytes,1,rep,name=hash_list,json=hashList,proto3" json:"hash_list,omitempty"`
	Hash2Idx             map[string]int32 `protobuf:"bytes,2,rep,name=hash2_idx,json=hash2Idx,proto3" json:"hash2_idx,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	LeafNum              int32            `protobuf:"varint,3,opt,name=leaf_num,json=leafNum,proto3" json:"leaf_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MerkleTree) Reset()         { *m = MerkleTree{} }
func (m *MerkleTree) String() string { return proto.CompactTextString(m) }
func (*MerkleTree) ProtoMessage()    {}
func (*MerkleTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafd901455e59c2f, []int{0}
}

func (m *MerkleTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerkleTree.Unmarshal(m, b)
}
func (m *MerkleTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerkleTree.Marshal(b, m, deterministic)
}
func (m *MerkleTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleTree.Merge(m, src)
}
func (m *MerkleTree) XXX_Size() int {
	return xxx_messageInfo_MerkleTree.Size(m)
}
func (m *MerkleTree) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleTree.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleTree proto.InternalMessageInfo

func (m *MerkleTree) GetHashList() [][]byte {
	if m != nil {
		return m.HashList
	}
	return nil
}

func (m *MerkleTree) GetHash2Idx() map[string]int32 {
	if m != nil {
		return m.Hash2Idx
	}
	return nil
}

func (m *MerkleTree) GetLeafNum() int32 {
	if m != nil {
		return m.LeafNum
	}
	return 0
}

type TXRMerkleTree struct {
	Mt                   *MerkleTree       `protobuf:"bytes,1,opt,name=mt,proto3" json:"mt,omitempty"`
	Tx2Txr               map[string][]byte `protobuf:"bytes,2,rep,name=tx2_txr,json=tx2Txr,proto3" json:"tx2_txr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TXRMerkleTree) Reset()         { *m = TXRMerkleTree{} }
func (m *TXRMerkleTree) String() string { return proto.CompactTextString(m) }
func (*TXRMerkleTree) ProtoMessage()    {}
func (*TXRMerkleTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafd901455e59c2f, []int{1}
}

func (m *TXRMerkleTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TXRMerkleTree.Unmarshal(m, b)
}
func (m *TXRMerkleTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TXRMerkleTree.Marshal(b, m, deterministic)
}
func (m *TXRMerkleTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TXRMerkleTree.Merge(m, src)
}
func (m *TXRMerkleTree) XXX_Size() int {
	return xxx_messageInfo_TXRMerkleTree.Size(m)
}
func (m *TXRMerkleTree) XXX_DiscardUnknown() {
	xxx_messageInfo_TXRMerkleTree.DiscardUnknown(m)
}

var xxx_messageInfo_TXRMerkleTree proto.InternalMessageInfo

func (m *TXRMerkleTree) GetMt() *MerkleTree {
	if m != nil {
		return m.Mt
	}
	return nil
}

func (m *TXRMerkleTree) GetTx2Txr() map[string][]byte {
	if m != nil {
		return m.Tx2Txr
	}
	return nil
}

func init() {
	proto.RegisterType((*MerkleTree)(nil), "merkletree.MerkleTree")
	proto.RegisterMapType((map[string]int32)(nil), "merkletree.MerkleTree.Hash2IdxEntry")
	proto.RegisterType((*TXRMerkleTree)(nil), "merkletree.TXRMerkleTree")
	proto.RegisterMapType((map[string][]byte)(nil), "merkletree.TXRMerkleTree.Tx2TxrEntry")
}

func init() { proto.RegisterFile("core/merkletree/merkle_tree.proto", fileDescriptor_cafd901455e59c2f) }

var fileDescriptor_cafd901455e59c2f = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xce, 0x2f, 0x4a,
	0xd5, 0xcf, 0x4d, 0x2d, 0xca, 0xce, 0x49, 0x2d, 0x29, 0x4a, 0x85, 0x31, 0xe3, 0x41, 0x6c, 0xbd,
	0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x2e, 0x84, 0xac, 0xd2, 0x11, 0x46, 0x2e, 0x2e, 0x5f, 0x30,
	0x37, 0xa4, 0x28, 0x35, 0x55, 0x48, 0x9a, 0x8b, 0x33, 0x23, 0xb1, 0x38, 0x23, 0x3e, 0x27, 0xb3,
	0xb8, 0x44, 0x82, 0x51, 0x81, 0x59, 0x83, 0x27, 0x88, 0x03, 0x24, 0xe0, 0x93, 0x59, 0x5c, 0x22,
	0xe4, 0x08, 0x91, 0x34, 0x8a, 0xcf, 0x4c, 0xa9, 0x90, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x52,
	0xd1, 0x43, 0x98, 0xa5, 0x87, 0x30, 0x47, 0xcf, 0x03, 0xa4, 0xce, 0x33, 0xa5, 0xc2, 0x35, 0xaf,
	0xa4, 0xa8, 0x12, 0x62, 0x04, 0x88, 0x2b, 0x24, 0xc9, 0xc5, 0x91, 0x93, 0x9a, 0x98, 0x16, 0x9f,
	0x57, 0x9a, 0x2b, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x0e, 0xe2, 0xfb, 0x95, 0xe6, 0x4a,
	0x59, 0x73, 0xf1, 0xa2, 0xe8, 0x12, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98,
	0xc0, 0x5a, 0x21, 0x1c, 0x2b, 0x26, 0x0b, 0x46, 0xa5, 0x4d, 0x8c, 0x5c, 0xbc, 0x21, 0x11, 0x41,
	0x48, 0x3e, 0x51, 0xe3, 0x62, 0xca, 0x2d, 0x01, 0x6b, 0xe6, 0x36, 0x12, 0xc3, 0xee, 0xca, 0x20,
	0xa6, 0xdc, 0x12, 0x21, 0x3b, 0x2e, 0xf6, 0x92, 0x0a, 0xa3, 0xf8, 0x92, 0x8a, 0x22, 0xa8, 0x97,
	0x54, 0x91, 0x15, 0xa3, 0x98, 0xa9, 0x17, 0x52, 0x61, 0x14, 0x52, 0x51, 0x04, 0xf1, 0x13, 0x5b,
	0x09, 0x98, 0x23, 0x65, 0xc9, 0xc5, 0x8d, 0x24, 0x4c, 0xc8, 0xd1, 0x3c, 0x48, 0x8e, 0x4e, 0x62,
	0x03, 0x47, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x63, 0xc7, 0x17, 0xba, 0xb3, 0x01, 0x00,
	0x00,
}
