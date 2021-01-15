// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kernel/engines/xuperos/xpb/xpb.proto

package xpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	xldgpb "github.com/xuperchain/xupercore/bcs/ledger/xledger/xldgpb"
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

type Transactions struct {
	Txs                  []*xldgpb.Transaction `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Transactions) Reset()         { *m = Transactions{} }
func (m *Transactions) String() string { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()    {}
func (*Transactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9685bde11a1952e, []int{0}
}

func (m *Transactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transactions.Unmarshal(m, b)
}
func (m *Transactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transactions.Marshal(b, m, deterministic)
}
func (m *Transactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transactions.Merge(m, src)
}
func (m *Transactions) XXX_Size() int {
	return xxx_messageInfo_Transactions.Size(m)
}
func (m *Transactions) XXX_DiscardUnknown() {
	xxx_messageInfo_Transactions.DiscardUnknown(m)
}

var xxx_messageInfo_Transactions proto.InternalMessageInfo

func (m *Transactions) GetTxs() []*xldgpb.Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxInfo struct {
	// 当前状态
	Status xldgpb.TransactionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=xldgpb.TransactionStatus" json:"status,omitempty"`
	// 离主干末端的距离（如果在主干上)
	Distance             int64               `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	Tx                   *xldgpb.Transaction `protobuf:"bytes,3,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TxInfo) Reset()         { *m = TxInfo{} }
func (m *TxInfo) String() string { return proto.CompactTextString(m) }
func (*TxInfo) ProtoMessage()    {}
func (*TxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9685bde11a1952e, []int{1}
}

func (m *TxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxInfo.Unmarshal(m, b)
}
func (m *TxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxInfo.Marshal(b, m, deterministic)
}
func (m *TxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInfo.Merge(m, src)
}
func (m *TxInfo) XXX_Size() int {
	return xxx_messageInfo_TxInfo.Size(m)
}
func (m *TxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TxInfo proto.InternalMessageInfo

func (m *TxInfo) GetStatus() xldgpb.TransactionStatus {
	if m != nil {
		return m.Status
	}
	return xldgpb.TransactionStatus_TX_UNDEFINE
}

func (m *TxInfo) GetDistance() int64 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *TxInfo) GetTx() *xldgpb.Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

type BlockInfo struct {
	Status               xldgpb.BlockStatus    `protobuf:"varint,1,opt,name=status,proto3,enum=xldgpb.BlockStatus" json:"status,omitempty"`
	Block                *xldgpb.InternalBlock `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BlockInfo) Reset()         { *m = BlockInfo{} }
func (m *BlockInfo) String() string { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()    {}
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9685bde11a1952e, []int{2}
}

func (m *BlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockInfo.Unmarshal(m, b)
}
func (m *BlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockInfo.Marshal(b, m, deterministic)
}
func (m *BlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockInfo.Merge(m, src)
}
func (m *BlockInfo) XXX_Size() int {
	return xxx_messageInfo_BlockInfo.Size(m)
}
func (m *BlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockInfo proto.InternalMessageInfo

func (m *BlockInfo) GetStatus() xldgpb.BlockStatus {
	if m != nil {
		return m.Status
	}
	return xldgpb.BlockStatus_BLOCK_ERROR
}

func (m *BlockInfo) GetBlock() *xldgpb.InternalBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

type ChainStatus struct {
	LedgerMeta           *xldgpb.LedgerMeta    `protobuf:"bytes,1,opt,name=ledger_meta,json=ledgerMeta,proto3" json:"ledger_meta,omitempty"`
	UtxoMeta             *xldgpb.UtxoMeta      `protobuf:"bytes,2,opt,name=utxo_meta,json=utxoMeta,proto3" json:"utxo_meta,omitempty"`
	Block                *xldgpb.InternalBlock `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
	BranchIds            []string              `protobuf:"bytes,4,rep,name=branch_ids,json=branchIds,proto3" json:"branch_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChainStatus) Reset()         { *m = ChainStatus{} }
func (m *ChainStatus) String() string { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()    {}
func (*ChainStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9685bde11a1952e, []int{3}
}

func (m *ChainStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainStatus.Unmarshal(m, b)
}
func (m *ChainStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainStatus.Marshal(b, m, deterministic)
}
func (m *ChainStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainStatus.Merge(m, src)
}
func (m *ChainStatus) XXX_Size() int {
	return xxx_messageInfo_ChainStatus.Size(m)
}
func (m *ChainStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ChainStatus proto.InternalMessageInfo

func (m *ChainStatus) GetLedgerMeta() *xldgpb.LedgerMeta {
	if m != nil {
		return m.LedgerMeta
	}
	return nil
}

func (m *ChainStatus) GetUtxoMeta() *xldgpb.UtxoMeta {
	if m != nil {
		return m.UtxoMeta
	}
	return nil
}

func (m *ChainStatus) GetBlock() *xldgpb.InternalBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *ChainStatus) GetBranchIds() []string {
	if m != nil {
		return m.BranchIds
	}
	return nil
}

type SystemStatus struct {
	ChainStatus          *ChainStatus `protobuf:"bytes,1,opt,name=chain_status,json=chainStatus,proto3" json:"chain_status,omitempty"`
	PeerUrls             []string     `protobuf:"bytes,2,rep,name=peer_urls,json=peerUrls,proto3" json:"peer_urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SystemStatus) Reset()         { *m = SystemStatus{} }
func (m *SystemStatus) String() string { return proto.CompactTextString(m) }
func (*SystemStatus) ProtoMessage()    {}
func (*SystemStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9685bde11a1952e, []int{4}
}

func (m *SystemStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemStatus.Unmarshal(m, b)
}
func (m *SystemStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemStatus.Marshal(b, m, deterministic)
}
func (m *SystemStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemStatus.Merge(m, src)
}
func (m *SystemStatus) XXX_Size() int {
	return xxx_messageInfo_SystemStatus.Size(m)
}
func (m *SystemStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SystemStatus proto.InternalMessageInfo

func (m *SystemStatus) GetChainStatus() *ChainStatus {
	if m != nil {
		return m.ChainStatus
	}
	return nil
}

func (m *SystemStatus) GetPeerUrls() []string {
	if m != nil {
		return m.PeerUrls
	}
	return nil
}

type TipStatus struct {
	IsTrunkTip           bool     `protobuf:"varint,1,opt,name=is_trunk_tip,json=isTrunkTip,proto3" json:"is_trunk_tip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TipStatus) Reset()         { *m = TipStatus{} }
func (m *TipStatus) String() string { return proto.CompactTextString(m) }
func (*TipStatus) ProtoMessage()    {}
func (*TipStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9685bde11a1952e, []int{5}
}

func (m *TipStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TipStatus.Unmarshal(m, b)
}
func (m *TipStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TipStatus.Marshal(b, m, deterministic)
}
func (m *TipStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TipStatus.Merge(m, src)
}
func (m *TipStatus) XXX_Size() int {
	return xxx_messageInfo_TipStatus.Size(m)
}
func (m *TipStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TipStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TipStatus proto.InternalMessageInfo

func (m *TipStatus) GetIsTrunkTip() bool {
	if m != nil {
		return m.IsTrunkTip
	}
	return false
}

type BlockID struct {
	Bcname  string `protobuf:"bytes,1,opt,name=bcname,proto3" json:"bcname,omitempty"`
	Blockid []byte `protobuf:"bytes,2,opt,name=blockid,proto3" json:"blockid,omitempty"`
	// if need content
	NeedContent          bool     `protobuf:"varint,3,opt,name=need_content,json=needContent,proto3" json:"need_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9685bde11a1952e, []int{6}
}

func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetBcname() string {
	if m != nil {
		return m.Bcname
	}
	return ""
}

func (m *BlockID) GetBlockid() []byte {
	if m != nil {
		return m.Blockid
	}
	return nil
}

func (m *BlockID) GetNeedContent() bool {
	if m != nil {
		return m.NeedContent
	}
	return false
}

func init() {
	proto.RegisterType((*Transactions)(nil), "protos.Transactions")
	proto.RegisterType((*TxInfo)(nil), "protos.TxInfo")
	proto.RegisterType((*BlockInfo)(nil), "protos.BlockInfo")
	proto.RegisterType((*ChainStatus)(nil), "protos.ChainStatus")
	proto.RegisterType((*SystemStatus)(nil), "protos.SystemStatus")
	proto.RegisterType((*TipStatus)(nil), "protos.TipStatus")
	proto.RegisterType((*BlockID)(nil), "protos.BlockID")
}

func init() {
	proto.RegisterFile("kernel/engines/xuperos/xpb/xpb.proto", fileDescriptor_e9685bde11a1952e)
}

var fileDescriptor_e9685bde11a1952e = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xa6, 0x8d, 0xf6, 0x9a, 0x49, 0x10, 0x59, 0x51, 0xe2, 0x89, 0x50, 0xa3, 0x42, 0xe1, 0xb8,
	0x86, 0xeb, 0xa1, 0x0f, 0xe2, 0xd3, 0x9d, 0x2f, 0x05, 0x7d, 0xd9, 0xcb, 0xbd, 0xf8, 0x60, 0x48,
	0x36, 0x63, 0xbb, 0x34, 0xdd, 0x0d, 0xbb, 0x1b, 0x88, 0xe0, 0x5f, 0xf3, 0xbf, 0xc9, 0xee, 0xa6,
	0xbd, 0x72, 0x78, 0xf8, 0x10, 0x76, 0x67, 0xe6, 0xfb, 0xf2, 0x7d, 0x33, 0xc9, 0xc0, 0xbb, 0x2d,
	0x2a, 0x81, 0x4d, 0x86, 0x62, 0xcd, 0x05, 0xea, 0xac, 0xef, 0x5a, 0x54, 0x52, 0x67, 0x7d, 0x5b,
	0xd9, 0x67, 0xd1, 0x2a, 0x69, 0x24, 0x99, 0xb8, 0x43, 0x9f, 0x5e, 0xb8, 0x32, 0x93, 0x0a, 0xb3,
	0x8a, 0xe9, 0xac, 0xc1, 0x7a, 0x8d, 0x2a, 0xeb, 0x0f, 0x67, 0xbd, 0xb6, 0x34, 0x1f, 0x7a, 0x6a,
	0xfa, 0x01, 0xe2, 0x5c, 0x95, 0x42, 0x97, 0xcc, 0x70, 0x29, 0x34, 0x79, 0x0f, 0x81, 0xe9, 0x75,
	0x32, 0x9a, 0x05, 0xf3, 0x68, 0xf9, 0x6c, 0xe1, 0x39, 0x8b, 0x23, 0x08, 0xb5, 0xf5, 0xf4, 0x37,
	0x4c, 0xf2, 0x7e, 0x25, 0x7e, 0x4a, 0x72, 0x01, 0x13, 0x6d, 0x4a, 0xd3, 0x59, 0xce, 0x68, 0xfe,
	0x64, 0xf9, 0xf2, 0x1f, 0x9c, 0x1b, 0x07, 0xa0, 0x03, 0x90, 0x9c, 0xc2, 0xb4, 0xe6, 0xda, 0x94,
	0x82, 0x61, 0x32, 0x9e, 0x8d, 0xe6, 0x01, 0x3d, 0xc4, 0xe4, 0x2d, 0x8c, 0x4d, 0x9f, 0x04, 0xb3,
	0xd1, 0x43, 0xf2, 0x63, 0xd3, 0xa7, 0x08, 0xe1, 0x55, 0x23, 0xd9, 0xd6, 0x19, 0x38, 0xbb, 0x67,
	0xe0, 0xc0, 0x72, 0x90, 0x7b, 0xd2, 0x67, 0xf0, 0xb8, 0xb2, 0x69, 0xa7, 0x1b, 0x2d, 0x9f, 0xef,
	0xb1, 0x2b, 0x61, 0x50, 0x89, 0xb2, 0x71, 0x1c, 0xea, 0x31, 0xe9, 0x9f, 0x11, 0x44, 0xd7, 0x9b,
	0x92, 0x0f, 0xfe, 0xc9, 0x25, 0x44, 0x7e, 0x76, 0xc5, 0x0e, 0x4d, 0xe9, 0xe4, 0xa2, 0x25, 0xd9,
	0xbf, 0xe2, 0xab, 0x2b, 0x7d, 0x43, 0x53, 0x52, 0x68, 0x0e, 0x77, 0x72, 0x0e, 0x61, 0x67, 0x7a,
	0xe9, 0x29, 0x5e, 0xf5, 0xe9, 0x9e, 0x72, 0x6b, 0x7a, 0xe9, 0x08, 0xd3, 0x6e, 0xb8, 0xdd, 0x19,
	0x0c, 0xfe, 0x6f, 0x90, 0xbc, 0x06, 0xa8, 0x54, 0x29, 0xd8, 0xa6, 0xe0, 0xb5, 0x4e, 0x1e, 0xcd,
	0x82, 0x79, 0x48, 0x43, 0x9f, 0x59, 0xd5, 0x3a, 0x65, 0x10, 0xdf, 0xfc, 0xd2, 0x06, 0x77, 0x83,
	0xff, 0x8f, 0x10, 0x33, 0xdb, 0x4e, 0x71, 0x34, 0x2f, 0x3b, 0x65, 0xff, 0xf7, 0x2c, 0x8e, 0x5a,
	0xa5, 0x11, 0x3b, 0xea, 0xfb, 0x15, 0x84, 0x2d, 0xa2, 0x2a, 0x3a, 0xd5, 0xe8, 0x64, 0xec, 0x54,
	0xa6, 0x36, 0x71, 0xab, 0x1a, 0x9d, 0x9e, 0x43, 0x98, 0xf3, 0x76, 0x40, 0xce, 0x20, 0xe6, 0xba,
	0x30, 0xaa, 0x13, 0xdb, 0xc2, 0xf0, 0xd6, 0x29, 0x4c, 0x29, 0x70, 0x9d, 0xdb, 0x54, 0xce, 0xdb,
	0xf4, 0x07, 0x9c, 0xf8, 0x4f, 0xf7, 0x85, 0xbc, 0x80, 0x49, 0xc5, 0x44, 0xb9, 0x43, 0x07, 0x0b,
	0xe9, 0x10, 0x91, 0x04, 0x4e, 0x5c, 0x7b, 0xbc, 0x76, 0xf3, 0x8a, 0xe9, 0x3e, 0x24, 0x6f, 0x20,
	0x16, 0x88, 0x75, 0xc1, 0xa4, 0x30, 0x28, 0x8c, 0x9b, 0xd1, 0x94, 0x46, 0x36, 0x77, 0xed, 0x53,
	0x57, 0x9f, 0xbf, 0x7f, 0x5a, 0x73, 0xb3, 0xe9, 0xaa, 0x05, 0x93, 0x3b, 0xbf, 0x2e, 0xae, 0x95,
	0xec, 0x6e, 0x35, 0x1e, 0x5e, 0xa9, 0xca, 0x2f, 0xd2, 0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5a, 0x69, 0x82, 0x3c, 0x77, 0x03, 0x00, 0x00,
}