// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: huddle/profiles/v3/msgs_chain_links.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgLinkChainAccount represents a message to link an account to a profile.
type MsgLinkChainAccount struct {
	// ChainAddress contains the details of the external chain address to be
	// linked
	ChainAddress *types.Any `protobuf:"bytes,1,opt,name=chain_address,json=chainAddress,proto3" json:"chain_address,omitempty" yaml:"source_address"`
	// Proof contains the proof of ownership of the external chain address
	Proof Proof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof" yaml:"source_proof"`
	// ChainConfig contains the configuration of the external chain
	ChainConfig ChainConfig `protobuf:"bytes,3,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config" yaml:"source_chain_config"`
	// Signer represents the Huddle address associated with the
	// profile to which link the external account
	Signer string `protobuf:"bytes,4,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
}

func (m *MsgLinkChainAccount) Reset()         { *m = MsgLinkChainAccount{} }
func (m *MsgLinkChainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgLinkChainAccount) ProtoMessage()    {}
func (*MsgLinkChainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{0}
}
func (m *MsgLinkChainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkChainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkChainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkChainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkChainAccount.Merge(m, src)
}
func (m *MsgLinkChainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkChainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkChainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkChainAccount proto.InternalMessageInfo

func (m *MsgLinkChainAccount) GetChainAddress() *types.Any {
	if m != nil {
		return m.ChainAddress
	}
	return nil
}

func (m *MsgLinkChainAccount) GetProof() Proof {
	if m != nil {
		return m.Proof
	}
	return Proof{}
}

func (m *MsgLinkChainAccount) GetChainConfig() ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return ChainConfig{}
}

func (m *MsgLinkChainAccount) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

// MsgLinkChainAccountResponse defines the Msg/LinkAccount response type.
type MsgLinkChainAccountResponse struct {
}

func (m *MsgLinkChainAccountResponse) Reset()         { *m = MsgLinkChainAccountResponse{} }
func (m *MsgLinkChainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLinkChainAccountResponse) ProtoMessage()    {}
func (*MsgLinkChainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{1}
}
func (m *MsgLinkChainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkChainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkChainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkChainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkChainAccountResponse.Merge(m, src)
}
func (m *MsgLinkChainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkChainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkChainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkChainAccountResponse proto.InternalMessageInfo

// MsgUnlinkChainAccount represents a message to unlink an account from a
// profile.
type MsgUnlinkChainAccount struct {
	// Owner represents the Huddle profile from which to remove the link
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// ChainName represents the name of the chain to which the link to remove is
	// associated
	ChainName string `protobuf:"bytes,2,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty" yaml:"chain_name"`
	// Target represents the external address to be removed
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty" yaml:"target"`
}

func (m *MsgUnlinkChainAccount) Reset()         { *m = MsgUnlinkChainAccount{} }
func (m *MsgUnlinkChainAccount) String() string { return proto.CompactTextString(m) }
func (*MsgUnlinkChainAccount) ProtoMessage()    {}
func (*MsgUnlinkChainAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{2}
}
func (m *MsgUnlinkChainAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlinkChainAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlinkChainAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlinkChainAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlinkChainAccount.Merge(m, src)
}
func (m *MsgUnlinkChainAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlinkChainAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlinkChainAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlinkChainAccount proto.InternalMessageInfo

func (m *MsgUnlinkChainAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgUnlinkChainAccount) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *MsgUnlinkChainAccount) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

// MsgUnlinkChainAccountResponse defines the Msg/UnlinkAccount response type.
type MsgUnlinkChainAccountResponse struct {
}

func (m *MsgUnlinkChainAccountResponse) Reset()         { *m = MsgUnlinkChainAccountResponse{} }
func (m *MsgUnlinkChainAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnlinkChainAccountResponse) ProtoMessage()    {}
func (*MsgUnlinkChainAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{3}
}
func (m *MsgUnlinkChainAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlinkChainAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlinkChainAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlinkChainAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlinkChainAccountResponse.Merge(m, src)
}
func (m *MsgUnlinkChainAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlinkChainAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlinkChainAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlinkChainAccountResponse proto.InternalMessageInfo

// MsgSetDefaultExternalAddress represents the message used to set a default
// address for a specific chain
type MsgSetDefaultExternalAddress struct {
	// Name of the chain for which to set the default address
	ChainName string `protobuf:"bytes,1,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	// Address to be set as the default one
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// User signing the message
	Signer string `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *MsgSetDefaultExternalAddress) Reset()         { *m = MsgSetDefaultExternalAddress{} }
func (m *MsgSetDefaultExternalAddress) String() string { return proto.CompactTextString(m) }
func (*MsgSetDefaultExternalAddress) ProtoMessage()    {}
func (*MsgSetDefaultExternalAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{4}
}
func (m *MsgSetDefaultExternalAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetDefaultExternalAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDefaultExternalAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetDefaultExternalAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDefaultExternalAddress.Merge(m, src)
}
func (m *MsgSetDefaultExternalAddress) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetDefaultExternalAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDefaultExternalAddress.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDefaultExternalAddress proto.InternalMessageInfo

func (m *MsgSetDefaultExternalAddress) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *MsgSetDefaultExternalAddress) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *MsgSetDefaultExternalAddress) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

// MsgSetDefaultExternalAddressResponse represents the
// Msg/SetDefaultExternalAddress response type
type MsgSetDefaultExternalAddressResponse struct {
}

func (m *MsgSetDefaultExternalAddressResponse) Reset()         { *m = MsgSetDefaultExternalAddressResponse{} }
func (m *MsgSetDefaultExternalAddressResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetDefaultExternalAddressResponse) ProtoMessage()    {}
func (*MsgSetDefaultExternalAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52cd1f5b825f2f4e, []int{5}
}
func (m *MsgSetDefaultExternalAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetDefaultExternalAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDefaultExternalAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetDefaultExternalAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDefaultExternalAddressResponse.Merge(m, src)
}
func (m *MsgSetDefaultExternalAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetDefaultExternalAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDefaultExternalAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDefaultExternalAddressResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgLinkChainAccount)(nil), "huddle.profiles.v3.MsgLinkChainAccount")
	proto.RegisterType((*MsgLinkChainAccountResponse)(nil), "huddle.profiles.v3.MsgLinkChainAccountResponse")
	proto.RegisterType((*MsgUnlinkChainAccount)(nil), "huddle.profiles.v3.MsgUnlinkChainAccount")
	proto.RegisterType((*MsgUnlinkChainAccountResponse)(nil), "huddle.profiles.v3.MsgUnlinkChainAccountResponse")
	proto.RegisterType((*MsgSetDefaultExternalAddress)(nil), "huddle.profiles.v3.MsgSetDefaultExternalAddress")
	proto.RegisterType((*MsgSetDefaultExternalAddressResponse)(nil), "huddle.profiles.v3.MsgSetDefaultExternalAddressResponse")
}

func init() {
	proto.RegisterFile("huddle/profiles/v3/msgs_chain_links.proto", fileDescriptor_52cd1f5b825f2f4e)
}

var fileDescriptor_52cd1f5b825f2f4e = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x6d, 0x36, 0x36, 0xa9, 0xee, 0x26, 0xb1, 0x6c, 0x45, 0x5d, 0x47, 0x93, 0xc9, 0x42, 0x53,
	0x27, 0xb4, 0x44, 0xd0, 0x3d, 0xf1, 0xd6, 0x6c, 0x3c, 0xb1, 0x22, 0x14, 0xc4, 0x0b, 0x2f, 0x95,
	0x9b, 0x3a, 0x5e, 0xb4, 0xc4, 0x8e, 0x62, 0xa7, 0xac, 0x7f, 0xc1, 0x27, 0xf0, 0x11, 0x7c, 0xc4,
	0xc4, 0xd3, 0x1e, 0x91, 0x90, 0x22, 0xd4, 0xfe, 0x41, 0xbe, 0x00, 0xc5, 0x4e, 0xd6, 0x15, 0x22,
	0xde, 0xec, 0x7b, 0xcf, 0x3d, 0xe7, 0xdc, 0x13, 0x07, 0x9c, 0x4e, 0x31, 0x8f, 0x18, 0xb7, 0xe3,
	0x84, 0xf9, 0x41, 0x88, 0xb9, 0x3d, 0x1b, 0xd8, 0x11, 0x27, 0x7c, 0xec, 0x5d, 0xa3, 0x80, 0x8e,
	0xc3, 0x80, 0xde, 0x70, 0x2b, 0x4e, 0x98, 0x60, 0xba, 0xae, 0xa0, 0x56, 0x05, 0xb5, 0x66, 0x83,
	0xee, 0x01, 0x61, 0x84, 0xc9, 0xb6, 0x5d, 0x9c, 0x14, 0xb2, 0x7b, 0x48, 0x18, 0x23, 0x21, 0xb6,
	0xe5, 0x6d, 0x92, 0xfa, 0x36, 0xa2, 0xf3, 0xaa, 0xe5, 0xb1, 0x82, 0x64, 0xac, 0x66, 0xd4, 0xa5,
	0x6c, 0xbd, 0xac, 0xb3, 0xc2, 0xa6, 0x38, 0xac, 0x31, 0x03, 0x7f, 0x6d, 0x80, 0xfd, 0x11, 0x27,
	0x57, 0x01, 0xbd, 0xb9, 0x28, 0x9a, 0x43, 0xcf, 0x63, 0x29, 0x15, 0xba, 0x07, 0x76, 0x15, 0x18,
	0x4d, 0xa7, 0x09, 0xe6, 0xbc, 0xa3, 0x1d, 0x6b, 0xfd, 0xd6, 0xeb, 0x03, 0x4b, 0x59, 0xb2, 0x2a,
	0x4b, 0xd6, 0x90, 0xce, 0x9d, 0x7e, 0x9e, 0x99, 0xed, 0x39, 0x8a, 0xc2, 0x37, 0x90, 0xb3, 0x34,
	0xf1, 0x70, 0x35, 0x05, 0x7f, 0x7c, 0x3f, 0x6b, 0x0d, 0xd5, 0xf9, 0x12, 0x09, 0xe4, 0xee, 0x48,
	0xd2, 0xb2, 0xa2, 0x5f, 0x81, 0xad, 0x38, 0x61, 0xcc, 0xef, 0x6c, 0x48, 0xf2, 0x43, 0xeb, 0xdf,
	0x64, 0xac, 0x0f, 0x05, 0xc0, 0x39, 0xba, 0xcb, 0xcc, 0x46, 0x9e, 0x99, 0xfb, 0x6b, 0x2a, 0x72,
	0x18, 0xba, 0x8a, 0x44, 0xf7, 0x81, 0x62, 0x1f, 0x7b, 0x8c, 0xfa, 0x01, 0xe9, 0x6c, 0x4a, 0x52,
	0xb3, 0x8e, 0x54, 0xae, 0x7a, 0x21, 0x61, 0x0e, 0x2c, 0xa9, 0xbb, 0x6b, 0xd4, 0x8f, 0x99, 0xa0,
	0xdb, 0xf2, 0x56, 0x03, 0xfa, 0x29, 0xd8, 0xe6, 0x01, 0xa1, 0x38, 0xe9, 0x3c, 0x39, 0xd6, 0xfa,
	0x4d, 0x67, 0x2f, 0xcf, 0xcc, 0xdd, 0x72, 0x58, 0xd6, 0xa1, 0x5b, 0x02, 0x60, 0x0f, 0x1c, 0xd5,
	0x84, 0xeb, 0x62, 0x1e, 0x33, 0xca, 0x31, 0xfc, 0xa6, 0x81, 0xf6, 0x88, 0x93, 0x4f, 0x34, 0xfc,
	0x3b, 0xfe, 0x13, 0xb0, 0xc5, 0xbe, 0x14, 0x12, 0x9a, 0x94, 0x78, 0x9a, 0x67, 0xe6, 0x8e, 0x92,
	0x90, 0x65, 0xe8, 0xaa, 0xb6, 0x7e, 0x0e, 0x80, 0x72, 0x4a, 0x51, 0x84, 0x65, 0x8c, 0x4d, 0xa7,
	0x9d, 0x67, 0xe6, 0x9e, 0x02, 0xaf, 0x7a, 0xd0, 0x6d, 0xca, 0xcb, 0x7b, 0x14, 0xe1, 0x62, 0x03,
	0x81, 0x12, 0x82, 0x85, 0xcc, 0x68, 0x6d, 0x03, 0x55, 0x87, 0x6e, 0x09, 0x80, 0x26, 0xe8, 0xd5,
	0x3a, 0x7c, 0xd8, 0x21, 0x02, 0xcf, 0x47, 0x9c, 0x7c, 0xc4, 0xe2, 0x12, 0xfb, 0x28, 0x0d, 0xc5,
	0xdb, 0x5b, 0x81, 0x13, 0x8a, 0xc2, 0xea, 0x1b, 0xf7, 0xd6, 0x1c, 0xca, 0x75, 0x1e, 0x5b, 0x79,
	0xf6, 0x60, 0x45, 0x9a, 0xaf, 0x74, 0x8b, 0x7a, 0x19, 0xf2, 0xa6, 0xaa, 0x97, 0x89, 0x9e, 0x80,
	0x17, 0xff, 0x93, 0xab, 0x6c, 0x39, 0xef, 0xee, 0x16, 0x86, 0x76, 0xbf, 0x30, 0xb4, 0xdf, 0x0b,
	0x43, 0xfb, 0xba, 0x34, 0x1a, 0xf7, 0x4b, 0xa3, 0xf1, 0x73, 0x69, 0x34, 0x3e, 0xbf, 0x22, 0x81,
	0xb8, 0x4e, 0x27, 0x96, 0xc7, 0x22, 0x5b, 0x3d, 0x8d, 0xb3, 0x10, 0x4d, 0x78, 0x79, 0xb6, 0x67,
	0xe7, 0xf6, 0xed, 0xea, 0xd7, 0x11, 0xf3, 0x18, 0xf3, 0xc9, 0xb6, 0x7c, 0xed, 0x83, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0x45, 0xc0, 0x28, 0xe5, 0x03, 0x00, 0x00,
}

func (m *MsgLinkChainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkChainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkChainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ChainAddress != nil {
		{
			size, err := m.ChainAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgsChainLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLinkChainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkChainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkChainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnlinkChainAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlinkChainAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlinkChainAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnlinkChainAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlinkChainAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlinkChainAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetDefaultExternalAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDefaultExternalAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDefaultExternalAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintMsgsChainLinks(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetDefaultExternalAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDefaultExternalAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDefaultExternalAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsChainLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsChainLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLinkChainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainAddress != nil {
		l = m.ChainAddress.Size()
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = m.Proof.Size()
	n += 1 + l + sovMsgsChainLinks(uint64(l))
	l = m.ChainConfig.Size()
	n += 1 + l + sovMsgsChainLinks(uint64(l))
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	return n
}

func (m *MsgLinkChainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnlinkChainAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	return n
}

func (m *MsgUnlinkChainAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetDefaultExternalAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgsChainLinks(uint64(l))
	}
	return n
}

func (m *MsgSetDefaultExternalAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsChainLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsChainLinks(x uint64) (n int) {
	return sovMsgsChainLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLinkChainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgLinkChainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkChainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ChainAddress == nil {
				m.ChainAddress = &types.Any{}
			}
			if err := m.ChainAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgLinkChainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgLinkChainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkChainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnlinkChainAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnlinkChainAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlinkChainAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnlinkChainAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnlinkChainAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlinkChainAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetDefaultExternalAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetDefaultExternalAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDefaultExternalAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetDefaultExternalAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetDefaultExternalAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDefaultExternalAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsChainLinks
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMsgsChainLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsChainLinks
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMsgsChainLinks
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMsgsChainLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsChainLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsChainLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsChainLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsChainLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsChainLinks = fmt.Errorf("proto: unexpected end of group")
)
