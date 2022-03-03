// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/query_app_links.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryUserApplicationLinkRequest represents the request used when querying an
// application link using an application name and username for a given user
type QueryApplicationLinksRequest struct {
	// (Optional) User contains the Desmos profile address associated for which
	// the link should be searched for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// (Optional) Application represents the application name associated with the
	// link. Used only if user is also set.
	Application string `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	// Username represents the username inside the application associated with the
	// link. Used only if application is also set.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// Pagination defines an optional pagination for the request
	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryApplicationLinksRequest) Reset()         { *m = QueryApplicationLinksRequest{} }
func (m *QueryApplicationLinksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryApplicationLinksRequest) ProtoMessage()    {}
func (*QueryApplicationLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b6ad0478f8a288, []int{0}
}
func (m *QueryApplicationLinksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryApplicationLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryApplicationLinksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryApplicationLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryApplicationLinksRequest.Merge(m, src)
}
func (m *QueryApplicationLinksRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryApplicationLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryApplicationLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryApplicationLinksRequest proto.InternalMessageInfo

func (m *QueryApplicationLinksRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryApplicationLinksRequest) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *QueryApplicationLinksRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *QueryApplicationLinksRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryApplicationLinksResponse represents the response to the query used
// to get the application links for a specific user
type QueryApplicationLinksResponse struct {
	Links []ApplicationLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links"`
	// Pagination defines the pagination response
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryApplicationLinksResponse) Reset()         { *m = QueryApplicationLinksResponse{} }
func (m *QueryApplicationLinksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryApplicationLinksResponse) ProtoMessage()    {}
func (*QueryApplicationLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b6ad0478f8a288, []int{1}
}
func (m *QueryApplicationLinksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryApplicationLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryApplicationLinksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryApplicationLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryApplicationLinksResponse.Merge(m, src)
}
func (m *QueryApplicationLinksResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryApplicationLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryApplicationLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryApplicationLinksResponse proto.InternalMessageInfo

func (m *QueryApplicationLinksResponse) GetLinks() []ApplicationLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *QueryApplicationLinksResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryApplicationLinkByClientIDRequest contains the data of the request that
// can be used to get an application link based on a client id
type QueryApplicationLinkByClientIDRequest struct {
	// ClientID represents the ID of the client to which search the link for
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *QueryApplicationLinkByClientIDRequest) Reset()         { *m = QueryApplicationLinkByClientIDRequest{} }
func (m *QueryApplicationLinkByClientIDRequest) String() string { return proto.CompactTextString(m) }
func (*QueryApplicationLinkByClientIDRequest) ProtoMessage()    {}
func (*QueryApplicationLinkByClientIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b6ad0478f8a288, []int{2}
}
func (m *QueryApplicationLinkByClientIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryApplicationLinkByClientIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryApplicationLinkByClientIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryApplicationLinkByClientIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryApplicationLinkByClientIDRequest.Merge(m, src)
}
func (m *QueryApplicationLinkByClientIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryApplicationLinkByClientIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryApplicationLinkByClientIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryApplicationLinkByClientIDRequest proto.InternalMessageInfo

func (m *QueryApplicationLinkByClientIDRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

// QueryApplicationLinkByClientIDResponse contains the data returned by the
// request allowing to get an application link using a client id
type QueryApplicationLinkByClientIDResponse struct {
	Link ApplicationLink `protobuf:"bytes,1,opt,name=link,proto3" json:"link"`
}

func (m *QueryApplicationLinkByClientIDResponse) Reset() {
	*m = QueryApplicationLinkByClientIDResponse{}
}
func (m *QueryApplicationLinkByClientIDResponse) String() string { return proto.CompactTextString(m) }
func (*QueryApplicationLinkByClientIDResponse) ProtoMessage()    {}
func (*QueryApplicationLinkByClientIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b6ad0478f8a288, []int{3}
}
func (m *QueryApplicationLinkByClientIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryApplicationLinkByClientIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryApplicationLinkByClientIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryApplicationLinkByClientIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryApplicationLinkByClientIDResponse.Merge(m, src)
}
func (m *QueryApplicationLinkByClientIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryApplicationLinkByClientIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryApplicationLinkByClientIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryApplicationLinkByClientIDResponse proto.InternalMessageInfo

func (m *QueryApplicationLinkByClientIDResponse) GetLink() ApplicationLink {
	if m != nil {
		return m.Link
	}
	return ApplicationLink{}
}

func init() {
	proto.RegisterType((*QueryApplicationLinksRequest)(nil), "desmos.profiles.v3.QueryApplicationLinksRequest")
	proto.RegisterType((*QueryApplicationLinksResponse)(nil), "desmos.profiles.v3.QueryApplicationLinksResponse")
	proto.RegisterType((*QueryApplicationLinkByClientIDRequest)(nil), "desmos.profiles.v3.QueryApplicationLinkByClientIDRequest")
	proto.RegisterType((*QueryApplicationLinkByClientIDResponse)(nil), "desmos.profiles.v3.QueryApplicationLinkByClientIDResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/query_app_links.proto", fileDescriptor_b3b6ad0478f8a288)
}

var fileDescriptor_b3b6ad0478f8a288 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xad, 0xa0, 0xcd, 0xbd, 0x59, 0x1c, 0xb2, 0x30, 0x42, 0x55, 0xc4, 0x28, 0x48,
	0xd8, 0x6a, 0x7b, 0x46, 0x88, 0x32, 0x81, 0x26, 0x38, 0x40, 0x8f, 0x5c, 0x2a, 0x27, 0xf1, 0x8c,
	0x45, 0x62, 0x7b, 0xb1, 0x13, 0x91, 0x6f, 0xc1, 0xd7, 0xe0, 0x1b, 0xf0, 0x11, 0x76, 0xdc, 0x91,
	0x13, 0x42, 0xed, 0x17, 0x41, 0xb1, 0x3d, 0xba, 0x8d, 0xa1, 0x69, 0x37, 0xbf, 0xf7, 0xff, 0xbf,
	0xf7, 0x7e, 0xef, 0x25, 0x70, 0x9c, 0x33, 0x53, 0x2a, 0x43, 0x74, 0xa5, 0x8e, 0x45, 0xc1, 0x0c,
	0x69, 0x66, 0xe4, 0xa4, 0x66, 0x55, 0xbb, 0xa4, 0x5a, 0x2f, 0x0b, 0x21, 0xbf, 0x18, 0xac, 0x2b,
	0x65, 0x15, 0x42, 0xde, 0x89, 0xcf, 0x9d, 0xb8, 0x99, 0xc5, 0xf7, 0xb8, 0xe2, 0xca, 0xc9, 0xa4,
	0x7b, 0x79, 0x67, 0xbc, 0xcf, 0x95, 0xe2, 0x05, 0x23, 0x54, 0x0b, 0x42, 0xa5, 0x54, 0x96, 0x5a,
	0xa1, 0x64, 0xe8, 0x13, 0xef, 0x05, 0xd5, 0x45, 0x69, 0x7d, 0x4c, 0xa8, 0x6c, 0x83, 0xf4, 0xf4,
	0x1a, 0x98, 0x52, 0xe5, 0xac, 0x30, 0x57, 0x69, 0xe2, 0xbd, 0x4c, 0x75, 0xd6, 0xa5, 0x1f, 0xee,
	0x83, 0x20, 0x3d, 0xf3, 0x11, 0x49, 0xa9, 0x61, 0x7e, 0x17, 0xd2, 0x4c, 0x52, 0x66, 0xe9, 0x84,
	0x68, 0xca, 0x85, 0x74, 0x34, 0xde, 0x3b, 0xfa, 0x01, 0xe0, 0xfe, 0xc7, 0xce, 0xf2, 0x4a, 0xeb,
	0x42, 0x64, 0x4e, 0x7a, 0xdf, 0x8d, 0x59, 0xb0, 0x93, 0x9a, 0x19, 0x8b, 0x10, 0xec, 0xd7, 0x86,
	0x55, 0x11, 0x18, 0x82, 0xf1, 0xee, 0xc2, 0xbd, 0xd1, 0x10, 0x0e, 0xe8, 0xc6, 0x1e, 0x6d, 0x39,
	0xe9, 0x62, 0x0a, 0xc5, 0x70, 0xa7, 0x73, 0x4a, 0x5a, 0xb2, 0x68, 0xdb, 0xc9, 0x7f, 0x63, 0xf4,
	0x06, 0xc2, 0x0d, 0x46, 0xd4, 0x1f, 0x82, 0xf1, 0x60, 0x7a, 0x80, 0xc3, 0x06, 0x1d, 0x33, 0x76,
	0xcc, 0x38, 0x30, 0xe3, 0x0f, 0x94, 0xb3, 0x40, 0xb3, 0xb8, 0x50, 0x39, 0xfa, 0x0e, 0xe0, 0x83,
	0xff, 0xa0, 0x1b, 0xad, 0xa4, 0x61, 0xe8, 0x25, 0xbc, 0xe3, 0x4e, 0x16, 0x81, 0xe1, 0xf6, 0x78,
	0x30, 0x7d, 0x84, 0xff, 0xfd, 0x82, 0xf8, 0x4a, 0xf1, 0xbc, 0x7f, 0xfa, 0xeb, 0x61, 0x6f, 0xe1,
	0xeb, 0xd0, 0xdb, 0x4b, 0xa8, 0x5b, 0x0e, 0xf5, 0xc9, 0x8d, 0xa8, 0x7e, 0xfa, 0x25, 0xd6, 0x43,
	0xf8, 0xf8, 0x3a, 0xd4, 0x79, 0xfb, 0xba, 0x10, 0x4c, 0xda, 0xa3, 0xc3, 0xf3, 0x73, 0xdf, 0x87,
	0xbb, 0x99, 0x4b, 0x2d, 0x45, 0x1e, 0x6e, 0xbe, 0xe3, 0x13, 0x47, 0xf9, 0x88, 0xc3, 0x83, 0x9b,
	0xba, 0x84, 0xcd, 0x5f, 0xc0, 0x7e, 0xb7, 0x81, 0xeb, 0x70, 0xab, 0xc5, 0x5d, 0xd9, 0xfc, 0xdd,
	0xe9, 0x2a, 0x01, 0x67, 0xab, 0x04, 0xfc, 0x5e, 0x25, 0xe0, 0xdb, 0x3a, 0xe9, 0x9d, 0xad, 0x93,
	0xde, 0xcf, 0x75, 0xd2, 0xfb, 0x34, 0xe1, 0xc2, 0x7e, 0xae, 0x53, 0x9c, 0xa9, 0x92, 0xf8, 0xa6,
	0xcf, 0x0b, 0x9a, 0x9a, 0xf0, 0x26, 0xcd, 0x94, 0x7c, 0xdd, 0xfc, 0xbd, 0xb6, 0xd5, 0xcc, 0xa4,
	0x77, 0xdd, 0x9f, 0x36, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x14, 0x05, 0x93, 0x6a, 0x03,
	0x00, 0x00,
}

func (m *QueryApplicationLinksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryApplicationLinksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryApplicationLinksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Application) > 0 {
		i -= len(m.Application)
		copy(dAtA[i:], m.Application)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.Application)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryApplicationLinksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryApplicationLinksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryApplicationLinksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryApplicationLinkByClientIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryApplicationLinkByClientIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryApplicationLinkByClientIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryApplicationLinkByClientIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryApplicationLinkByClientIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryApplicationLinkByClientIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Link.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQueryAppLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryAppLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryApplicationLinksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	l = len(m.Application)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func (m *QueryApplicationLinksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovQueryAppLinks(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func (m *QueryApplicationLinkByClientIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func (m *QueryApplicationLinkByClientIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Link.Size()
	n += 1 + l + sovQueryAppLinks(uint64(l))
	return n
}

func sovQueryAppLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryAppLinks(x uint64) (n int) {
	return sovQueryAppLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryApplicationLinksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryApplicationLinksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryApplicationLinksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Application = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func (m *QueryApplicationLinksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryApplicationLinksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryApplicationLinksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, ApplicationLink{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func (m *QueryApplicationLinkByClientIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryApplicationLinkByClientIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryApplicationLinkByClientIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func (m *QueryApplicationLinkByClientIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryApplicationLinkByClientIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryApplicationLinkByClientIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Link.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func skipQueryAppLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryAppLinks
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
					return 0, ErrIntOverflowQueryAppLinks
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
					return 0, ErrIntOverflowQueryAppLinks
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
				return 0, ErrInvalidLengthQueryAppLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryAppLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryAppLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryAppLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryAppLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryAppLinks = fmt.Errorf("proto: unexpected end of group")
)
