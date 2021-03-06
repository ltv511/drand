// Code generated by protoc-gen-go. DO NOT EDIT.
// source: drand/protocol.proto

package drand

import (
	context "context"
	dkg "github.com/drand/drand/protobuf/crypto/dkg"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// PrepareDKGPacket is the packet nodes send to a coordinator that collects all
// keys and setups the group and sends them back to the nodes such that they can
// start the DKG automatically.
type PrepareDKGPacket struct {
	Node *Identity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// following parameters are helpful if leader and participants did not run
	// with the same parameters, so the leader can reply if it's not consistent.
	Expected    uint32 `protobuf:"varint,2,opt,name=expected,proto3" json:"expected,omitempty"`
	Threshold   uint32 `protobuf:"varint,3,opt,name=threshold,proto3" json:"threshold,omitempty"`
	DkgTimeout  uint64 `protobuf:"varint,4,opt,name=dkg_timeout,json=dkgTimeout,proto3" json:"dkg_timeout,omitempty"`
	SecretProof string `protobuf:"bytes,5,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// In resharing cases, previous_group_hash is the hash of the previous group.
	// It is to make sure the nodes build on top of the correct previous group.
	PreviousGroupHash    string   `protobuf:"bytes,6,opt,name=previous_group_hash,json=previousGroupHash,proto3" json:"previous_group_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareDKGPacket) Reset()         { *m = PrepareDKGPacket{} }
func (m *PrepareDKGPacket) String() string { return proto.CompactTextString(m) }
func (*PrepareDKGPacket) ProtoMessage()    {}
func (*PrepareDKGPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{0}
}

func (m *PrepareDKGPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareDKGPacket.Unmarshal(m, b)
}
func (m *PrepareDKGPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareDKGPacket.Marshal(b, m, deterministic)
}
func (m *PrepareDKGPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareDKGPacket.Merge(m, src)
}
func (m *PrepareDKGPacket) XXX_Size() int {
	return xxx_messageInfo_PrepareDKGPacket.Size(m)
}
func (m *PrepareDKGPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareDKGPacket.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareDKGPacket proto.InternalMessageInfo

func (m *PrepareDKGPacket) GetNode() *Identity {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *PrepareDKGPacket) GetExpected() uint32 {
	if m != nil {
		return m.Expected
	}
	return 0
}

func (m *PrepareDKGPacket) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

func (m *PrepareDKGPacket) GetDkgTimeout() uint64 {
	if m != nil {
		return m.DkgTimeout
	}
	return 0
}

func (m *PrepareDKGPacket) GetSecretProof() string {
	if m != nil {
		return m.SecretProof
	}
	return ""
}

func (m *PrepareDKGPacket) GetPreviousGroupHash() string {
	if m != nil {
		return m.PreviousGroupHash
	}
	return ""
}

type BeaconPacket struct {
	// Round is the round for which the beacon will be created from the partial
	// signatures
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	// PreviousRound is the round for which the beacon is building on top of
	// from.
	PreviousRound        uint64   `protobuf:"varint,2,opt,name=previous_round,json=previousRound,proto3" json:"previous_round,omitempty"`
	PartialSig           []byte   `protobuf:"bytes,3,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
	PreviousSig          []byte   `protobuf:"bytes,4,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconPacket) Reset()         { *m = BeaconPacket{} }
func (m *BeaconPacket) String() string { return proto.CompactTextString(m) }
func (*BeaconPacket) ProtoMessage()    {}
func (*BeaconPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{1}
}

func (m *BeaconPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconPacket.Unmarshal(m, b)
}
func (m *BeaconPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconPacket.Marshal(b, m, deterministic)
}
func (m *BeaconPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconPacket.Merge(m, src)
}
func (m *BeaconPacket) XXX_Size() int {
	return xxx_messageInfo_BeaconPacket.Size(m)
}
func (m *BeaconPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconPacket.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconPacket proto.InternalMessageInfo

func (m *BeaconPacket) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *BeaconPacket) GetPreviousRound() uint64 {
	if m != nil {
		return m.PreviousRound
	}
	return 0
}

func (m *BeaconPacket) GetPartialSig() []byte {
	if m != nil {
		return m.PartialSig
	}
	return nil
}

func (m *BeaconPacket) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

type DKGPacket struct {
	Dkg                  *dkg.Packet `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DKGPacket) Reset()         { *m = DKGPacket{} }
func (m *DKGPacket) String() string { return proto.CompactTextString(m) }
func (*DKGPacket) ProtoMessage()    {}
func (*DKGPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{2}
}

func (m *DKGPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DKGPacket.Unmarshal(m, b)
}
func (m *DKGPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DKGPacket.Marshal(b, m, deterministic)
}
func (m *DKGPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DKGPacket.Merge(m, src)
}
func (m *DKGPacket) XXX_Size() int {
	return xxx_messageInfo_DKGPacket.Size(m)
}
func (m *DKGPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_DKGPacket.DiscardUnknown(m)
}

var xxx_messageInfo_DKGPacket proto.InternalMessageInfo

func (m *DKGPacket) GetDkg() *dkg.Packet {
	if m != nil {
		return m.Dkg
	}
	return nil
}

// Reshare is a wrapper around a Setup packet for resharing operation that
// serves two purposes: - indicate to non-leader old nodes that they should
// generate and send their deals - indicate to which new group are we resharing.
// drand should keep a list of new ready-to-operate groups allowed.
type ResharePacket struct {
	Dkg                  *dkg.Packet `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	GroupHash            string      `protobuf:"bytes,2,opt,name=group_hash,json=groupHash,proto3" json:"group_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResharePacket) Reset()         { *m = ResharePacket{} }
func (m *ResharePacket) String() string { return proto.CompactTextString(m) }
func (*ResharePacket) ProtoMessage()    {}
func (*ResharePacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{3}
}

func (m *ResharePacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResharePacket.Unmarshal(m, b)
}
func (m *ResharePacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResharePacket.Marshal(b, m, deterministic)
}
func (m *ResharePacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResharePacket.Merge(m, src)
}
func (m *ResharePacket) XXX_Size() int {
	return xxx_messageInfo_ResharePacket.Size(m)
}
func (m *ResharePacket) XXX_DiscardUnknown() {
	xxx_messageInfo_ResharePacket.DiscardUnknown(m)
}

var xxx_messageInfo_ResharePacket proto.InternalMessageInfo

func (m *ResharePacket) GetDkg() *dkg.Packet {
	if m != nil {
		return m.Dkg
	}
	return nil
}

func (m *ResharePacket) GetGroupHash() string {
	if m != nil {
		return m.GroupHash
	}
	return ""
}

// SyncRequest is from a node that needs to sync up with the current head of the
// chain
type SyncRequest struct {
	FromRound            uint64   `protobuf:"varint,1,opt,name=from_round,json=fromRound,proto3" json:"from_round,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRequest) Reset()         { *m = SyncRequest{} }
func (m *SyncRequest) String() string { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()    {}
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{4}
}

func (m *SyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRequest.Unmarshal(m, b)
}
func (m *SyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRequest.Marshal(b, m, deterministic)
}
func (m *SyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRequest.Merge(m, src)
}
func (m *SyncRequest) XXX_Size() int {
	return xxx_messageInfo_SyncRequest.Size(m)
}
func (m *SyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRequest proto.InternalMessageInfo

func (m *SyncRequest) GetFromRound() uint64 {
	if m != nil {
		return m.FromRound
	}
	return 0
}

// SyncResponse is basically a chain of beacon response
type SyncResponse struct {
	PreviousRound        uint64   `protobuf:"varint,1,opt,name=previous_round,json=previousRound,proto3" json:"previous_round,omitempty"`
	PreviousSig          []byte   `protobuf:"bytes,2,opt,name=previous_sig,json=previousSig,proto3" json:"previous_sig,omitempty"`
	Round                uint64   `protobuf:"varint,3,opt,name=round,proto3" json:"round,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncResponse) Reset()         { *m = SyncResponse{} }
func (m *SyncResponse) String() string { return proto.CompactTextString(m) }
func (*SyncResponse) ProtoMessage()    {}
func (*SyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e344a98fea1e2f3a, []int{5}
}

func (m *SyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncResponse.Unmarshal(m, b)
}
func (m *SyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncResponse.Marshal(b, m, deterministic)
}
func (m *SyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncResponse.Merge(m, src)
}
func (m *SyncResponse) XXX_Size() int {
	return xxx_messageInfo_SyncResponse.Size(m)
}
func (m *SyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncResponse proto.InternalMessageInfo

func (m *SyncResponse) GetPreviousRound() uint64 {
	if m != nil {
		return m.PreviousRound
	}
	return 0
}

func (m *SyncResponse) GetPreviousSig() []byte {
	if m != nil {
		return m.PreviousSig
	}
	return nil
}

func (m *SyncResponse) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *SyncResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*PrepareDKGPacket)(nil), "drand.PrepareDKGPacket")
	proto.RegisterType((*BeaconPacket)(nil), "drand.BeaconPacket")
	proto.RegisterType((*DKGPacket)(nil), "drand.DKGPacket")
	proto.RegisterType((*ResharePacket)(nil), "drand.ResharePacket")
	proto.RegisterType((*SyncRequest)(nil), "drand.SyncRequest")
	proto.RegisterType((*SyncResponse)(nil), "drand.SyncResponse")
}

func init() {
	proto.RegisterFile("drand/protocol.proto", fileDescriptor_e344a98fea1e2f3a)
}

var fileDescriptor_e344a98fea1e2f3a = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x14, 0x94, 0xf3, 0x51, 0xe2, 0xe7, 0x84, 0x96, 0x4d, 0x24, 0xac, 0xa8, 0x55, 0x43, 0x10, 0x52,
	0x54, 0x21, 0xb7, 0x2a, 0x12, 0x47, 0x0e, 0xa5, 0x10, 0x50, 0x05, 0x8a, 0x1c, 0x4e, 0x5c, 0x2c,
	0xe3, 0x7d, 0xb5, 0xad, 0x24, 0x5e, 0xb3, 0xbb, 0x06, 0xf2, 0x1f, 0x90, 0xe0, 0x4f, 0xf2, 0x3f,
	0xd0, 0x7e, 0x24, 0x4e, 0x43, 0x0f, 0x1c, 0x22, 0x65, 0x67, 0xe6, 0xad, 0xfc, 0x66, 0x66, 0x61,
	0x40, 0x79, 0x5c, 0xd0, 0xf3, 0x92, 0x33, 0xc9, 0x12, 0xb6, 0x0c, 0xf4, 0x1f, 0xd2, 0xd6, 0xe8,
	0x70, 0x90, 0xf0, 0x75, 0x29, 0xd9, 0x39, 0x5d, 0xa4, 0xea, 0x67, 0xc8, 0x21, 0x31, 0x23, 0x09,
	0x5b, 0xad, 0x58, 0x61, 0xb0, 0xf1, 0x1f, 0x07, 0x8e, 0x66, 0x1c, 0xcb, 0x98, 0xe3, 0xf5, 0xcd,
	0x74, 0x16, 0x27, 0x0b, 0x94, 0xe4, 0x29, 0xb4, 0x0a, 0x46, 0xd1, 0x77, 0x46, 0xce, 0xc4, 0xbb,
	0x3c, 0x0c, 0xf4, 0x5c, 0xf0, 0x9e, 0x62, 0x21, 0x73, 0xb9, 0x0e, 0x35, 0x49, 0x86, 0xd0, 0xc1,
	0x1f, 0x25, 0x26, 0x12, 0xa9, 0xdf, 0x18, 0x39, 0x93, 0x5e, 0xb8, 0x3d, 0x93, 0x63, 0x70, 0x65,
	0xc6, 0x51, 0x64, 0x6c, 0x49, 0xfd, 0xa6, 0x26, 0x6b, 0x80, 0x9c, 0x82, 0x47, 0x17, 0x69, 0x24,
	0xf3, 0x15, 0xb2, 0x4a, 0xfa, 0xad, 0x91, 0x33, 0x69, 0x85, 0x40, 0x17, 0xe9, 0x27, 0x83, 0x90,
	0x27, 0xd0, 0x15, 0x98, 0x70, 0x94, 0x51, 0xc9, 0x19, 0xbb, 0xf5, 0xdb, 0x23, 0x67, 0xe2, 0x86,
	0x9e, 0xc1, 0x66, 0x0a, 0x22, 0x01, 0xf4, 0x4b, 0x8e, 0xdf, 0x72, 0x56, 0x89, 0x28, 0xe5, 0xac,
	0x2a, 0xa3, 0x2c, 0x16, 0x99, 0x7f, 0xa0, 0x95, 0x8f, 0x36, 0xd4, 0x54, 0x31, 0xef, 0x62, 0x91,
	0x8d, 0x7f, 0x39, 0xd0, 0xbd, 0xc2, 0x38, 0x61, 0x85, 0xdd, 0x71, 0x00, 0x6d, 0xce, 0xaa, 0x82,
	0xea, 0x25, 0x5b, 0xa1, 0x39, 0x90, 0x67, 0xf0, 0x70, 0x7b, 0xad, 0xa1, 0x1b, 0x9a, 0xee, 0x6d,
	0xd0, 0x50, 0xcb, 0x4e, 0xc1, 0x2b, 0x63, 0x2e, 0xf3, 0x78, 0x19, 0x89, 0x3c, 0xd5, 0x1b, 0x76,
	0x43, 0xb0, 0xd0, 0x3c, 0x4f, 0xd5, 0x06, 0xdb, 0x7b, 0x94, 0xa2, 0xa5, 0x15, 0xde, 0x06, 0x9b,
	0xe7, 0xe9, 0xf8, 0x0c, 0xdc, 0xda, 0xf1, 0x13, 0x68, 0xd2, 0x45, 0x6a, 0x0d, 0xf7, 0x02, 0x95,
	0x99, 0x61, 0x42, 0x85, 0x8f, 0x3f, 0x40, 0x2f, 0x44, 0x91, 0xc5, 0x1c, 0xff, 0x4b, 0x4f, 0x4e,
	0x00, 0x76, 0x4c, 0x69, 0x68, 0x53, 0xdc, 0x74, 0x6b, 0xc6, 0x73, 0xf0, 0xe6, 0xeb, 0x22, 0x09,
	0xf1, 0x6b, 0x85, 0x42, 0x5d, 0x06, 0xb7, 0x9c, 0xad, 0xa2, 0x5d, 0x3f, 0x5c, 0x85, 0xe8, 0x65,
	0xc7, 0x3f, 0x1d, 0xe8, 0x1a, 0xb9, 0x28, 0x59, 0x21, 0xf0, 0x1e, 0x93, 0x9c, 0xfb, 0x4c, 0xda,
	0xf7, 0xa0, 0xf1, 0x8f, 0x07, 0x75, 0x08, 0xcd, 0xdd, 0x10, 0x8e, 0xc1, 0x15, 0x79, 0x5a, 0xc4,
	0xb2, 0xe2, 0x68, 0x9d, 0xab, 0x81, 0xcb, 0xdf, 0x0d, 0xe8, 0xcc, 0x6c, 0xeb, 0xc9, 0x2b, 0x38,
	0xac, 0xdb, 0xab, 0xd3, 0x26, 0x8f, 0x6d, 0x5d, 0xf7, 0x5b, 0x3d, 0x24, 0x96, 0xd0, 0x32, 0xeb,
	0xe3, 0x19, 0x74, 0xde, 0xaa, 0x5a, 0x5e, 0xdf, 0x4c, 0xc9, 0x91, 0xe5, 0xeb, 0x89, 0xae, 0x45,
	0xde, 0xac, 0x4a, 0xb9, 0x26, 0x17, 0x00, 0x36, 0x04, 0xa5, 0x1e, 0x58, 0xee, 0x4e, 0x2e, 0x7b,
	0x13, 0x01, 0xb8, 0x1f, 0xf1, 0xbb, 0xa9, 0x1d, 0xe9, 0x5b, 0x6a, 0xb7, 0x85, 0x7b, 0xfa, 0x97,
	0xe0, 0x2a, 0xa3, 0x5f, 0x67, 0x71, 0x5e, 0x90, 0xcd, 0xe7, 0xee, 0x24, 0x35, 0xec, 0xdf, 0xc1,
	0x4c, 0x1c, 0x17, 0xce, 0xd5, 0x83, 0xcf, 0xe6, 0xdd, 0x7f, 0x39, 0xd0, 0x8f, 0xfa, 0xc5, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xcc, 0xcc, 0x98, 0x1d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProtocolClient is the client API for Protocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtocolClient interface {
	PrepareDKGGroup(ctx context.Context, in *PrepareDKGPacket, opts ...grpc.CallOption) (*GroupPacket, error)
	// Setup is doing the DKG setup phase
	FreshDKG(ctx context.Context, in *DKGPacket, opts ...grpc.CallOption) (*Empty, error)
	// Reshare performs the resharing phase
	ReshareDKG(ctx context.Context, in *ResharePacket, opts ...grpc.CallOption) (*Empty, error)
	// NewBeacon asks for a partial signature to another node
	NewBeacon(ctx context.Context, in *BeaconPacket, opts ...grpc.CallOption) (*Empty, error)
	SyncChain(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (Protocol_SyncChainClient, error)
}

type protocolClient struct {
	cc grpc.ClientConnInterface
}

func NewProtocolClient(cc grpc.ClientConnInterface) ProtocolClient {
	return &protocolClient{cc}
}

func (c *protocolClient) PrepareDKGGroup(ctx context.Context, in *PrepareDKGPacket, opts ...grpc.CallOption) (*GroupPacket, error) {
	out := new(GroupPacket)
	err := c.cc.Invoke(ctx, "/drand.Protocol/PrepareDKGGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) FreshDKG(ctx context.Context, in *DKGPacket, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/drand.Protocol/FreshDKG", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) ReshareDKG(ctx context.Context, in *ResharePacket, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/drand.Protocol/ReshareDKG", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) NewBeacon(ctx context.Context, in *BeaconPacket, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/drand.Protocol/NewBeacon", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolClient) SyncChain(ctx context.Context, in *SyncRequest, opts ...grpc.CallOption) (Protocol_SyncChainClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Protocol_serviceDesc.Streams[0], "/drand.Protocol/SyncChain", opts...)
	if err != nil {
		return nil, err
	}
	x := &protocolSyncChainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Protocol_SyncChainClient interface {
	Recv() (*SyncResponse, error)
	grpc.ClientStream
}

type protocolSyncChainClient struct {
	grpc.ClientStream
}

func (x *protocolSyncChainClient) Recv() (*SyncResponse, error) {
	m := new(SyncResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProtocolServer is the server API for Protocol service.
type ProtocolServer interface {
	PrepareDKGGroup(context.Context, *PrepareDKGPacket) (*GroupPacket, error)
	// Setup is doing the DKG setup phase
	FreshDKG(context.Context, *DKGPacket) (*Empty, error)
	// Reshare performs the resharing phase
	ReshareDKG(context.Context, *ResharePacket) (*Empty, error)
	// NewBeacon asks for a partial signature to another node
	NewBeacon(context.Context, *BeaconPacket) (*Empty, error)
	SyncChain(*SyncRequest, Protocol_SyncChainServer) error
}

// UnimplementedProtocolServer can be embedded to have forward compatible implementations.
type UnimplementedProtocolServer struct {
}

func (*UnimplementedProtocolServer) PrepareDKGGroup(ctx context.Context, req *PrepareDKGPacket) (*GroupPacket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareDKGGroup not implemented")
}
func (*UnimplementedProtocolServer) FreshDKG(ctx context.Context, req *DKGPacket) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreshDKG not implemented")
}
func (*UnimplementedProtocolServer) ReshareDKG(ctx context.Context, req *ResharePacket) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReshareDKG not implemented")
}
func (*UnimplementedProtocolServer) NewBeacon(ctx context.Context, req *BeaconPacket) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewBeacon not implemented")
}
func (*UnimplementedProtocolServer) SyncChain(req *SyncRequest, srv Protocol_SyncChainServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncChain not implemented")
}

func RegisterProtocolServer(s *grpc.Server, srv ProtocolServer) {
	s.RegisterService(&_Protocol_serviceDesc, srv)
}

func _Protocol_PrepareDKGGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareDKGPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).PrepareDKGGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Protocol/PrepareDKGGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).PrepareDKGGroup(ctx, req.(*PrepareDKGPacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_FreshDKG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DKGPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).FreshDKG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Protocol/FreshDKG",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).FreshDKG(ctx, req.(*DKGPacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_ReshareDKG_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResharePacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).ReshareDKG(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Protocol/ReshareDKG",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).ReshareDKG(ctx, req.(*ResharePacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_NewBeacon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeaconPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServer).NewBeacon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/drand.Protocol/NewBeacon",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServer).NewBeacon(ctx, req.(*BeaconPacket))
	}
	return interceptor(ctx, in, info, handler)
}

func _Protocol_SyncChain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProtocolServer).SyncChain(m, &protocolSyncChainServer{stream})
}

type Protocol_SyncChainServer interface {
	Send(*SyncResponse) error
	grpc.ServerStream
}

type protocolSyncChainServer struct {
	grpc.ServerStream
}

func (x *protocolSyncChainServer) Send(m *SyncResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Protocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "drand.Protocol",
	HandlerType: (*ProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareDKGGroup",
			Handler:    _Protocol_PrepareDKGGroup_Handler,
		},
		{
			MethodName: "FreshDKG",
			Handler:    _Protocol_FreshDKG_Handler,
		},
		{
			MethodName: "ReshareDKG",
			Handler:    _Protocol_ReshareDKG_Handler,
		},
		{
			MethodName: "NewBeacon",
			Handler:    _Protocol_NewBeacon_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncChain",
			Handler:       _Protocol_SyncChain_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "drand/protocol.proto",
}
