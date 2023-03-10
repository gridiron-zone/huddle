// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: huddle/profiles/v3/models_packets.proto

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

// LinkChainAccountPacketData defines the object that should be sent inside a
// MsgSendPacket when wanting to link an external chain to a Huddle profile
// using IBC
type LinkChainAccountPacketData struct {
	// SourceAddress contains the details of the external chain address
	SourceAddress *types.Any `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty" yaml:"source_address"`
	// SourceProof represents the proof of ownership of the source address
	SourceProof Proof `protobuf:"bytes,2,opt,name=source_proof,json=sourceProof,proto3" json:"source_proof" yaml:"source_proof"`
	// SourceChainConfig contains the details of the source chain
	SourceChainConfig ChainConfig `protobuf:"bytes,3,opt,name=source_chain_config,json=sourceChainConfig,proto3" json:"source_chain_config" yaml:"source_chain_config"`
	// DestinationAddress represents the Huddle address of the profile that should
	// be linked with the external account
	DestinationAddress string `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty" yaml:"destination_address"`
	// DestinationProof contains the proof of ownership of the DestinationAddress
	DestinationProof Proof `protobuf:"bytes,5,opt,name=destination_proof,json=destinationProof,proto3" json:"destination_proof" yaml:"destination_proof"`
}

func (m *LinkChainAccountPacketData) Reset()         { *m = LinkChainAccountPacketData{} }
func (m *LinkChainAccountPacketData) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketData) ProtoMessage()    {}
func (*LinkChainAccountPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_943e1eb0c01292c0, []int{0}
}
func (m *LinkChainAccountPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketData.Merge(m, src)
}
func (m *LinkChainAccountPacketData) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketData proto.InternalMessageInfo

// LinkChainAccountPacketAck defines a struct for the packet acknowledgment
type LinkChainAccountPacketAck struct {
	// SourceAddress contains the external address that has been linked properly
	// with the profile
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
}

func (m *LinkChainAccountPacketAck) Reset()         { *m = LinkChainAccountPacketAck{} }
func (m *LinkChainAccountPacketAck) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketAck) ProtoMessage()    {}
func (*LinkChainAccountPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_943e1eb0c01292c0, []int{1}
}
func (m *LinkChainAccountPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketAck.Merge(m, src)
}
func (m *LinkChainAccountPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LinkChainAccountPacketData)(nil), "huddle.profiles.v3.LinkChainAccountPacketData")
	proto.RegisterType((*LinkChainAccountPacketAck)(nil), "huddle.profiles.v3.LinkChainAccountPacketAck")
}

func init() {
	proto.RegisterFile("huddle/profiles/v3/models_packets.proto", fileDescriptor_943e1eb0c01292c0)
}

var fileDescriptor_943e1eb0c01292c0 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0xb5, 0x21, 0x20, 0xd5, 0x01, 0x44, 0x9d, 0x22, 0x25, 0x41, 0xb2, 0xa3, 0x91, 0x10, 0x91,
	0x50, 0xc7, 0x88, 0xec, 0xba, 0x8b, 0xcb, 0x02, 0x21, 0x24, 0xaa, 0xec, 0x60, 0x13, 0x4d, 0xc6,
	0x13, 0x67, 0x14, 0x67, 0xbe, 0xe5, 0x19, 0x57, 0x84, 0x13, 0xb0, 0x64, 0xc1, 0x01, 0x38, 0x04,
	0x87, 0xa8, 0x58, 0x75, 0xc9, 0x2a, 0x42, 0xc9, 0x0d, 0x7a, 0x02, 0x94, 0x99, 0x09, 0x75, 0xa8,
	0x51, 0x77, 0xf3, 0xff, 0x7b, 0xef, 0xbf, 0xef, 0xe7, 0xef, 0x3d, 0x9f, 0x95, 0x49, 0x92, 0xb1,
	0x28, 0x2f, 0x60, 0xca, 0x33, 0x26, 0xa3, 0xf3, 0x41, 0xb4, 0x80, 0x84, 0x65, 0x72, 0x9c, 0x13,
	0x3a, 0x67, 0x4a, 0xe2, 0xbc, 0x00, 0x05, 0xbe, 0x6f, 0x88, 0x78, 0x47, 0xc4, 0xe7, 0x83, 0xee,
	0x51, 0x0a, 0x29, 0x68, 0x38, 0xda, 0xbe, 0x0c, 0xb3, 0xdb, 0x49, 0x01, 0x52, 0x33, 0x52, 0xc1,
	0xa4, 0x9c, 0x46, 0x44, 0x2c, 0x77, 0x10, 0x05, 0xb9, 0x00, 0x39, 0x36, 0x1a, 0x53, 0x58, 0xe8,
	0xc5, 0xff, 0x17, 0xa1, 0x33, 0xc2, 0xc5, 0x38, 0xe3, 0x62, 0x6e, 0xc9, 0xe8, 0x5b, 0xc3, 0xeb,
	0xbe, 0xe3, 0x62, 0x7e, 0xba, 0x45, 0x86, 0x94, 0x42, 0x29, 0xd4, 0x99, 0x5e, 0xf7, 0x35, 0x51,
	0xc4, 0x67, 0xde, 0x23, 0x09, 0x65, 0x41, 0xd9, 0x98, 0x24, 0x49, 0xc1, 0xa4, 0x6c, 0xbb, 0x3d,
	0xb7, 0xdf, 0x7c, 0x75, 0x84, 0xcd, 0x6a, 0x78, 0xb7, 0x1a, 0x1e, 0x8a, 0x65, 0xdc, 0xbf, 0x5a,
	0x85, 0x4f, 0x96, 0x64, 0x91, 0x9d, 0xa0, 0x7d, 0x15, 0xfa, 0xf9, 0xe3, 0xb8, 0x39, 0x34, 0xef,
	0xed, 0xdc, 0xd1, 0x43, 0x83, 0xdb, 0x96, 0xff, 0xc1, 0x7b, 0x60, 0x05, 0x79, 0x01, 0x30, 0x6d,
	0xdf, 0xd1, 0x26, 0x1d, 0x7c, 0x33, 0x29, 0x7c, 0xb6, 0x25, 0xc4, 0x4f, 0x2f, 0x56, 0xa1, 0x73,
	0xb5, 0x0a, 0x5b, 0x7b, 0x6e, 0x5a, 0x8c, 0x46, 0x4d, 0x53, 0x6a, 0xa6, 0x2f, 0xbd, 0x96, 0x45,
	0xcd, 0xc7, 0x53, 0x10, 0x53, 0x9e, 0xb6, 0xef, 0x6a, 0x87, 0xb0, 0xce, 0x41, 0x47, 0x71, 0xaa,
	0x69, 0x31, 0xb2, 0x3e, 0xdd, 0x3d, 0x9f, 0xea, 0x24, 0x34, 0x3a, 0x34, 0xdd, 0x8a, 0xcc, 0x7f,
	0xef, 0xb5, 0x12, 0x26, 0x15, 0x17, 0x44, 0x71, 0x10, 0x7f, 0xb3, 0x6b, 0xf4, 0xdc, 0xfe, 0x41,
	0x1c, 0x5c, 0xcf, 0xab, 0x21, 0xa1, 0x91, 0x5f, 0xe9, 0xee, 0x02, 0x9a, 0x79, 0x87, 0x55, 0xae,
	0x49, 0xe9, 0xde, 0x6d, 0x29, 0xf5, 0xec, 0xf6, 0xed, 0x9b, 0x6e, 0x36, 0xaa, 0xc7, 0x95, 0x9e,
	0xd6, 0x9c, 0x34, 0xbe, 0x7c, 0x0f, 0x1d, 0xf4, 0xc6, 0xeb, 0xd4, 0x5f, 0xc5, 0x90, 0xce, 0xfd,
	0x67, 0xb5, 0x47, 0x71, 0xf0, 0xcf, 0x4f, 0x35, 0x93, 0xe2, 0xb7, 0x17, 0xeb, 0xc0, 0xbd, 0x5c,
	0x07, 0xee, 0xef, 0x75, 0xe0, 0x7e, 0xdd, 0x04, 0xce, 0xe5, 0x26, 0x70, 0x7e, 0x6d, 0x02, 0xe7,
	0xe3, 0xcb, 0x94, 0xab, 0x59, 0x39, 0xc1, 0x14, 0x16, 0x51, 0x5a, 0xf0, 0x84, 0x17, 0x20, 0x8e,
	0x3f, 0x83, 0x60, 0x91, 0x3d, 0xe0, 0x4f, 0xd7, 0x27, 0xac, 0x96, 0x39, 0x93, 0x93, 0xfb, 0xfa,
	0xda, 0x06, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05, 0xcb, 0xc7, 0xbf, 0x6b, 0x03, 0x00, 0x00,
}

func (m *LinkChainAccountPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DestinationProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintModelsPackets(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.SourceChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SourceProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SourceAddress != nil {
		{
			size, err := m.SourceAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsPackets(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LinkChainAccountPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintModelsPackets(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsPackets(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsPackets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LinkChainAccountPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceAddress != nil {
		l = m.SourceAddress.Size()
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	l = m.SourceProof.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	l = m.SourceChainConfig.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	l = m.DestinationProof.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	return n
}

func (m *LinkChainAccountPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	return n
}

func sovModelsPackets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsPackets(x uint64) (n int) {
	return sovModelsPackets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LinkChainAccountPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsPackets
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
			return fmt.Errorf("proto: LinkChainAccountPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SourceAddress == nil {
				m.SourceAddress = &types.Any{}
			}
			if err := m.SourceAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SourceProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SourceChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DestinationProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsPackets
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
func (m *LinkChainAccountPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsPackets
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
			return fmt.Errorf("proto: LinkChainAccountPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsPackets
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
func skipModelsPackets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsPackets
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
					return 0, ErrIntOverflowModelsPackets
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
					return 0, ErrIntOverflowModelsPackets
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
				return 0, ErrInvalidLengthModelsPackets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsPackets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsPackets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsPackets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsPackets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsPackets = fmt.Errorf("proto: unexpected end of group")
)
