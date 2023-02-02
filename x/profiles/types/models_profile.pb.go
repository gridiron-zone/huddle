// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: huddle/profiles/v3/models_profile.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Profile represents a generic first on Huddle, containing the information of a
// single user
type Profile struct {
	// Account represents the base Cosmos account associated with this profile
	Account *types.Any `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// DTag represents the unique tag of this profile
	DTag string `protobuf:"bytes,2,opt,name=dtag,proto3" json:"dtag,omitempty" yaml:"dtag"`
	// Nickname contains the custom human readable name of the profile
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty" yaml:"nickname"`
	// Bio contains the biography of the profile
	Bio string `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty" yaml:"bio"`
	// Pictures contains the data about the pictures associated with he profile
	Pictures Pictures `protobuf:"bytes,5,opt,name=pictures,proto3" json:"pictures" yaml:"pictures"`
	// CreationTime represents the time in which the profile has been created
	CreationDate time.Time `protobuf:"bytes,6,opt,name=creation_date,json=creationDate,proto3,stdtime" json:"creation_date" yaml:"creation_date"`
}

func (m *Profile) Reset()      { *m = Profile{} }
func (*Profile) ProtoMessage() {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ad67db56683fb7e, []int{0}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

// Pictures contains the data of a user profile's related pictures
type Pictures struct {
	// Profile contains the URL to the profile picture
	Profile string `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty" yaml:"profile"`
	// Cover contains the URL to the cover picture
	Cover string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty" yaml:"cover"`
}

func (m *Pictures) Reset()         { *m = Pictures{} }
func (m *Pictures) String() string { return proto.CompactTextString(m) }
func (*Pictures) ProtoMessage()    {}
func (*Pictures) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ad67db56683fb7e, []int{1}
}
func (m *Pictures) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pictures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pictures.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pictures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pictures.Merge(m, src)
}
func (m *Pictures) XXX_Size() int {
	return m.Size()
}
func (m *Pictures) XXX_DiscardUnknown() {
	xxx_messageInfo_Pictures.DiscardUnknown(m)
}

var xxx_messageInfo_Pictures proto.InternalMessageInfo

func (m *Pictures) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *Pictures) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func init() {
	proto.RegisterType((*Profile)(nil), "huddle.profiles.v3.Profile")
	proto.RegisterType((*Pictures)(nil), "huddle.profiles.v3.Pictures")
}

func init() {
	proto.RegisterFile("huddle/profiles/v3/models_profile.proto", fileDescriptor_1ad67db56683fb7e)
}

var fileDescriptor_1ad67db56683fb7e = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x31, 0x6f, 0xd3, 0x4e,
	0x18, 0xc6, 0xed, 0x7f, 0xd2, 0xc6, 0xbd, 0xe6, 0x5f, 0xd0, 0x11, 0xa9, 0x26, 0x42, 0xbe, 0xe8,
	0x06, 0xa8, 0x04, 0xb5, 0x11, 0xd9, 0x22, 0x96, 0x5a, 0x5d, 0x60, 0x2a, 0x56, 0x27, 0x96, 0xea,
	0x62, 0x5f, 0xdd, 0x83, 0xd8, 0xaf, 0x65, 0x9f, 0x23, 0xc2, 0x17, 0x80, 0xb1, 0x63, 0xc7, 0x7c,
	0x08, 0x3e, 0x44, 0xc5, 0xd4, 0x91, 0xc9, 0xa0, 0x64, 0x61, 0xf6, 0x27, 0x40, 0xf1, 0xf9, 0x5a,
	0x41, 0xb7, 0xe4, 0x7d, 0x7e, 0xcf, 0xf3, 0xbe, 0x7a, 0x7c, 0xe8, 0xd9, 0x45, 0x19, 0x45, 0x33,
	0xee, 0x65, 0x39, 0x9c, 0x8b, 0x19, 0x2f, 0xbc, 0xf9, 0xd8, 0x4b, 0x20, 0xe2, 0xb3, 0xe2, 0xac,
	0x1d, 0xb9, 0x59, 0x0e, 0x12, 0x30, 0x56, 0xa0, 0xab, 0x41, 0x77, 0x3e, 0x1e, 0x0e, 0x62, 0x88,
	0xa1, 0x91, 0xbd, 0xcd, 0x2f, 0x45, 0x0e, 0x1f, 0xc7, 0x00, 0xb1, 0x8a, 0x94, 0x30, 0x2d, 0xcf,
	0x3d, 0x96, 0x2e, 0x5a, 0x89, 0xfc, 0x2b, 0x49, 0x91, 0xf0, 0x42, 0xb2, 0x24, 0xd3, 0xde, 0x10,
	0x8a, 0x04, 0x9a, 0xdd, 0x12, 0x3c, 0xf5, 0x47, 0x49, 0xf4, 0x4b, 0x07, 0xf5, 0x4e, 0xd4, 0x72,
	0xfc, 0x1a, 0xf5, 0x58, 0x18, 0x42, 0x99, 0x4a, 0xdb, 0x1c, 0x99, 0x07, 0xbb, 0xaf, 0x06, 0xae,
	0x4a, 0x76, 0x75, 0xb2, 0x7b, 0x94, 0x2e, 0xfc, 0xfe, 0xf7, 0x6f, 0x87, 0xd6, 0x91, 0x02, 0xdf,
	0x04, 0xda, 0x82, 0x9f, 0xa3, 0x6e, 0x24, 0x59, 0x6c, 0xff, 0x37, 0x32, 0x0f, 0x76, 0xfc, 0xfd,
	0x55, 0x45, 0xba, 0xc7, 0xa7, 0x2c, 0xae, 0x2b, 0xb2, 0xbb, 0x60, 0xc9, 0x6c, 0x42, 0x37, 0x2a,
	0x0d, 0x1a, 0x08, 0x7b, 0xc8, 0x4a, 0x45, 0xf8, 0x31, 0x65, 0x09, 0xb7, 0x3b, 0x8d, 0xe1, 0x51,
	0x5d, 0x91, 0x07, 0x0a, 0xd4, 0x0a, 0x0d, 0x6e, 0x21, 0x3c, 0x42, 0x9d, 0xa9, 0x00, 0xbb, 0xdb,
	0xb0, 0x7b, 0x75, 0x45, 0x90, 0x62, 0xa7, 0x02, 0x68, 0xb0, 0x91, 0xf0, 0x3b, 0x64, 0x65, 0x22,
	0x94, 0x65, 0xce, 0x0b, 0x7b, 0xab, 0x39, 0xff, 0x89, 0x7b, 0xbf, 0x5d, 0xf7, 0xa4, 0x65, 0xfc,
	0xfd, 0xeb, 0x8a, 0x18, 0x77, 0x4b, 0xb5, 0x97, 0x06, 0xb7, 0x31, 0x98, 0xa1, 0xff, 0xc3, 0x9c,
	0x33, 0x29, 0x20, 0x3d, 0x8b, 0x98, 0xe4, 0xf6, 0x76, 0x93, 0x3b, 0xbc, 0x57, 0xcb, 0xa9, 0x2e,
	0xdc, 0x1f, 0xb5, 0xa9, 0x03, 0x95, 0xfa, 0x97, 0x9d, 0x5e, 0xfe, 0x24, 0x66, 0xd0, 0xd7, 0xb3,
	0x63, 0x26, 0xf9, 0xc4, 0xfa, 0xba, 0x24, 0xc6, 0xd5, 0x92, 0x18, 0xf4, 0x03, 0xb2, 0xf4, 0x6d,
	0xf8, 0x05, 0xea, 0xb5, 0x37, 0x37, 0x5f, 0x62, 0xc7, 0xc7, 0x75, 0x45, 0xf6, 0xda, 0x43, 0x95,
	0x40, 0x03, 0x8d, 0xe0, 0xa7, 0x68, 0x2b, 0x84, 0x39, 0xcf, 0xdb, 0xea, 0x1f, 0xd6, 0x15, 0xe9,
	0xb7, 0xeb, 0x37, 0x63, 0x1a, 0x28, 0x79, 0x62, 0x5d, 0x2d, 0x89, 0xf9, 0x7b, 0x49, 0x4c, 0xff,
	0xed, 0xf5, 0xca, 0x31, 0x6f, 0x56, 0x8e, 0xf9, 0x6b, 0xe5, 0x98, 0x97, 0x6b, 0xc7, 0xb8, 0x59,
	0x3b, 0xc6, 0x8f, 0xb5, 0x63, 0xbc, 0x7f, 0x19, 0x0b, 0x79, 0x51, 0x4e, 0xdd, 0x10, 0x12, 0x2f,
	0xce, 0x45, 0x24, 0x72, 0x48, 0x0f, 0x3f, 0x43, 0xca, 0xbd, 0xf6, 0x49, 0x7f, 0xba, 0x7b, 0xd4,
	0x72, 0x91, 0xf1, 0x62, 0xba, 0xdd, 0xb4, 0x30, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x30, 0x0c,
	0x24, 0x28, 0xf4, 0x02, 0x00, 0x00,
}

func (this *Pictures) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pictures)
	if !ok {
		that2, ok := that.(Pictures)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Profile != that1.Profile {
		return false
	}
	if this.Cover != that1.Cover {
		return false
	}
	return true
}
func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreationDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationDate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintModelsProfile(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Pictures.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsProfile(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Bio) > 0 {
		i -= len(m.Bio)
		copy(dAtA[i:], m.Bio)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Bio)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DTag) > 0 {
		i -= len(m.DTag)
		copy(dAtA[i:], m.DTag)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.DTag)))
		i--
		dAtA[i] = 0x12
	}
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsProfile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pictures) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pictures) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pictures) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cover) > 0 {
		i -= len(m.Cover)
		copy(dAtA[i:], m.Cover)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Cover)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Profile) > 0 {
		i -= len(m.Profile)
		copy(dAtA[i:], m.Profile)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Profile)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.DTag)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.Bio)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = m.Pictures.Size()
	n += 1 + l + sovModelsProfile(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationDate)
	n += 1 + l + sovModelsProfile(uint64(l))
	return n
}

func (m *Pictures) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Profile)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.Cover)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	return n
}

func sovModelsProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsProfile(x uint64) (n int) {
	return sovModelsProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsProfile
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
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &types.Any{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pictures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pictures.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreationDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsProfile
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
func (m *Pictures) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsProfile
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
			return fmt.Errorf("proto: Pictures: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pictures: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cover", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cover = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsProfile
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
func skipModelsProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsProfile
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
					return 0, ErrIntOverflowModelsProfile
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
					return 0, ErrIntOverflowModelsProfile
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
				return 0, ErrInvalidLengthModelsProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsProfile = fmt.Errorf("proto: unexpected end of group")
)
