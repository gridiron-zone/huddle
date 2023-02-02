// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: huddle/profiles/v3/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("huddle/profiles/v3/query.proto", fileDescriptor_bcbdebc2a1cf2f2b) }

var fileDescriptor_bcbdebc2a1cf2f2b = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x4f, 0x13, 0x4f,
	0x18, 0xc7, 0xd9, 0x5f, 0xf2, 0xc3, 0x64, 0x62, 0xa2, 0x99, 0x68, 0x22, 0x0d, 0x2e, 0x11, 0xa5,
	0x10, 0x60, 0x77, 0x28, 0x55, 0x22, 0x46, 0x0f, 0xfc, 0xf1, 0x40, 0x34, 0x11, 0x0d, 0x27, 0x2f,
	0xcd, 0xec, 0x76, 0xba, 0x6c, 0xdc, 0xce, 0x0c, 0x3b, 0xb3, 0x15, 0x42, 0xb8, 0x78, 0xf3, 0x42,
	0x4c, 0xbc, 0x7b, 0xf2, 0x25, 0xf8, 0x1e, 0xf4, 0x62, 0x24, 0xf1, 0xe2, 0xd1, 0x80, 0x2f, 0xc4,
	0xec, 0xfc, 0xa1, 0xa1, 0xee, 0xb6, 0xa5, 0xb7, 0x9d, 0x3e, 0xdf, 0xef, 0x7c, 0x3f, 0xcf, 0xcc,
	0xd3, 0x5d, 0xe0, 0x36, 0x89, 0x68, 0x33, 0x81, 0x78, 0xca, 0x5a, 0x71, 0x42, 0x04, 0xea, 0xd4,
	0xd1, 0x5e, 0x46, 0xd2, 0x03, 0x9f, 0xa7, 0x4c, 0x32, 0x08, 0x75, 0xdd, 0xb7, 0x75, 0xbf, 0x53,
	0xaf, 0xdc, 0x88, 0x58, 0xc4, 0x54, 0x19, 0xe5, 0x4f, 0x5a, 0x59, 0x99, 0x8c, 0x18, 0x8b, 0x12,
	0x82, 0x30, 0x8f, 0x11, 0xa6, 0x94, 0x49, 0x2c, 0x63, 0x46, 0x85, 0xa9, 0x4e, 0x98, 0xaa, 0x5a,
	0x05, 0x59, 0x0b, 0x61, 0x6a, 0x22, 0x2a, 0xd5, 0x32, 0x84, 0x86, 0xf9, 0xc5, 0xe8, 0x16, 0x4b,
	0x75, 0x4d, 0x89, 0xa3, 0x46, 0x4a, 0xf6, 0x32, 0x22, 0xa4, 0x0d, 0x9c, 0x29, 0xdf, 0x15, 0xa7,
	0xb8, 0x6d, 0x65, 0xf3, 0xa5, 0xb2, 0x70, 0x17, 0xc7, 0xb4, 0x91, 0xc4, 0xf4, 0x8d, 0xd5, 0xce,
	0x95, 0x6a, 0x31, 0xe7, 0x17, 0x94, 0x13, 0x21, 0xcb, 0x95, 0x0d, 0x7d, 0x48, 0x7a, 0x61, 0x03,
	0xf5, 0x0a, 0x05, 0x58, 0x10, 0xed, 0x46, 0x9d, 0x5a, 0x40, 0x24, 0xae, 0x21, 0x8e, 0xa3, 0x98,
	0xaa, 0x53, 0xd3, 0xda, 0xe5, 0xe3, 0xab, 0xe0, 0xff, 0x97, 0xb9, 0x04, 0xbe, 0x77, 0xc0, 0x95,
	0x6d, 0x1d, 0x0b, 0x67, 0xfd, 0x7f, 0xef, 0xc4, 0x57, 0x32, 0xa3, 0x78, 0xa5, 0x4f, 0xa2, 0x32,
	0x37, 0x58, 0x28, 0x38, 0xa3, 0x82, 0x4c, 0x2f, 0xbc, 0xfb, 0xf9, 0xe7, 0xe3, 0x7f, 0x33, 0xf0,
	0x2e, 0x2a, 0x68, 0xf1, 0xfc, 0xf9, 0x30, 0x13, 0x24, 0x3d, 0x82, 0x3f, 0x1c, 0x30, 0xb9, 0x45,
	0x43, 0xd6, 0x8e, 0x69, 0xb4, 0xb9, 0x83, 0xa3, 0x9d, 0x14, 0x53, 0xd1, 0x22, 0xa9, 0x89, 0x15,
	0xf0, 0x71, 0x69, 0x6e, 0x3f, 0x9b, 0xa5, 0x7e, 0x32, 0xa2, 0xdb, 0xb4, 0xb2, 0xac, 0x5a, 0x59,
	0x84, 0xf3, 0x45, 0xad, 0xe4, 0x83, 0xe2, 0x49, 0x63, 0xf5, 0xec, 0xc4, 0xc0, 0x63, 0x07, 0x80,
	0x8d, 0xfc, 0xba, 0x9f, 0xe7, 0x77, 0x08, 0xe7, 0x4b, 0x09, 0xba, 0x22, 0x4b, 0xbb, 0x30, 0x94,
	0xd6, 0xb0, 0xcd, 0x2a, 0xb6, 0x3b, 0x70, 0xaa, 0x88, 0x4d, 0xcd, 0x9b, 0xa7, 0xa6, 0x08, 0x7e,
	0x76, 0xc0, 0xb5, 0x73, 0xff, 0x8b, 0xb7, 0x94, 0xa4, 0x02, 0xa2, 0xc1, 0x49, 0x5a, 0x69, 0xd1,
	0x96, 0x86, 0x37, 0x18, 0x3e, 0x5f, 0xf1, 0xcd, 0xc1, 0xea, 0x00, 0x3e, 0xc4, 0x34, 0xd2, 0x57,
	0x07, 0xdc, 0xda, 0x24, 0x2d, 0x9c, 0x25, 0xf2, 0xe9, 0xbe, 0x24, 0x29, 0xc5, 0xc9, 0x5a, 0xb3,
	0x99, 0x12, 0x21, 0x88, 0x80, 0x0f, 0x4b, 0xe3, 0xcb, 0x2c, 0x16, 0x7c, 0x75, 0x04, 0xa7, 0xe9,
	0x60, 0x45, 0x75, 0xb0, 0x04, 0xfd, 0xc2, 0xdb, 0xd7, 0x6e, 0x8f, 0x18, 0xbb, 0x87, 0xcf, 0x61,
	0x3f, 0x39, 0xe0, 0xfa, 0x1a, 0xe7, 0x49, 0x1c, 0xaa, 0xff, 0x9f, 0x9e, 0x83, 0xf2, 0x03, 0xec,
	0x95, 0x5a, 0xf2, 0xda, 0x25, 0x1c, 0x86, 0x78, 0x46, 0x11, 0x4f, 0xc1, 0xdb, 0x45, 0xc4, 0x98,
	0x73, 0x33, 0x11, 0xdf, 0x1d, 0x30, 0xd1, 0xb3, 0xc7, 0xfa, 0xc1, 0x46, 0x12, 0x13, 0x2a, 0xb7,
	0x36, 0xe1, 0xea, 0xb0, 0xb9, 0x5d, 0x8f, 0x45, 0x7e, 0x34, 0x8a, 0xd5, 0xb0, 0xaf, 0x2a, 0xf6,
	0x3a, 0xac, 0xf5, 0x65, 0x47, 0xa1, 0xf2, 0x09, 0x74, 0xa8, 0x1f, 0x1a, 0x71, 0xf3, 0x08, 0x7e,
	0x71, 0xc0, 0xcd, 0x9e, 0x00, 0x33, 0xe7, 0x0f, 0x86, 0x05, 0xba, 0x38, 0xed, 0x2b, 0x97, 0xb5,
	0x99, 0x1e, 0x16, 0x55, 0x0f, 0x55, 0x78, 0xaf, 0x7f, 0x0f, 0x66, 0xe2, 0x8f, 0xc0, 0xf8, 0xb6,
	0xfa, 0x7c, 0xc0, 0x6a, 0xf9, 0xcb, 0x55, 0x09, 0x2c, 0xd7, 0xec, 0x40, 0x9d, 0x01, 0x99, 0x56,
	0x20, 0x93, 0xb0, 0x52, 0xf8, 0x0e, 0x56, 0xda, 0xf5, 0x67, 0xdf, 0x4e, 0x5d, 0xe7, 0xe4, 0xd4,
	0x75, 0x7e, 0x9f, 0xba, 0xce, 0x87, 0x33, 0x77, 0xec, 0xe4, 0xcc, 0x1d, 0xfb, 0x75, 0xe6, 0x8e,
	0xbd, 0xae, 0x45, 0xb1, 0xdc, 0xcd, 0x02, 0x3f, 0x64, 0x6d, 0xe3, 0xf7, 0x12, 0x1c, 0x08, 0xbb,
	0x57, 0xe7, 0x3e, 0xda, 0xef, 0x6e, 0x28, 0x0f, 0x38, 0x11, 0xc1, 0xb8, 0xfa, 0xc8, 0xd4, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x04, 0xd5, 0x63, 0xc8, 0x03, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Profile queries the profile of a specific user given their DTag or address.
	// If the queried user does not have a profile, the returned response will
	// contain a null profile.
	Profile(ctx context.Context, in *QueryProfileRequest, opts ...grpc.CallOption) (*QueryProfileResponse, error)
	// IncomingDTagTransferRequests queries all the DTag transfers requests that
	// have been made towards the user with the given address
	IncomingDTagTransferRequests(ctx context.Context, in *QueryIncomingDTagTransferRequestsRequest, opts ...grpc.CallOption) (*QueryIncomingDTagTransferRequestsResponse, error)
	// ChainLinks queries the chain links associated to the given user, if
	// provided. Otherwise it queries all the chain links stored.
	ChainLinks(ctx context.Context, in *QueryChainLinksRequest, opts ...grpc.CallOption) (*QueryChainLinksResponse, error)
	// ChainLinkOwners queries for the owners of chain links, optionally searching
	// for a specific chain name and external address
	ChainLinkOwners(ctx context.Context, in *QueryChainLinkOwnersRequest, opts ...grpc.CallOption) (*QueryChainLinkOwnersResponse, error)
	// DefaultExternalAddresses queries the default addresses associated to the
	// given user and (optionally) chain name
	DefaultExternalAddresses(ctx context.Context, in *QueryDefaultExternalAddressesRequest, opts ...grpc.CallOption) (*QueryDefaultExternalAddressesResponse, error)
	// ApplicationLinks queries the applications links associated to the given
	// user, if provided. Otherwise, it queries all the application links stored.
	ApplicationLinks(ctx context.Context, in *QueryApplicationLinksRequest, opts ...grpc.CallOption) (*QueryApplicationLinksResponse, error)
	// ApplicationLinkByClientID queries a single application link for a given
	// client id.
	ApplicationLinkByClientID(ctx context.Context, in *QueryApplicationLinkByClientIDRequest, opts ...grpc.CallOption) (*QueryApplicationLinkByClientIDResponse, error)
	// ApplicationLinkOwners queries for the owners of applications links,
	// optionally searching for a specific application and username.
	ApplicationLinkOwners(ctx context.Context, in *QueryApplicationLinkOwnersRequest, opts ...grpc.CallOption) (*QueryApplicationLinkOwnersResponse, error)
	// Params queries the profiles module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Profile(ctx context.Context, in *QueryProfileRequest, opts ...grpc.CallOption) (*QueryProfileResponse, error) {
	out := new(QueryProfileResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IncomingDTagTransferRequests(ctx context.Context, in *QueryIncomingDTagTransferRequestsRequest, opts ...grpc.CallOption) (*QueryIncomingDTagTransferRequestsResponse, error) {
	out := new(QueryIncomingDTagTransferRequestsResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/IncomingDTagTransferRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainLinks(ctx context.Context, in *QueryChainLinksRequest, opts ...grpc.CallOption) (*QueryChainLinksResponse, error) {
	out := new(QueryChainLinksResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/ChainLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainLinkOwners(ctx context.Context, in *QueryChainLinkOwnersRequest, opts ...grpc.CallOption) (*QueryChainLinkOwnersResponse, error) {
	out := new(QueryChainLinkOwnersResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/ChainLinkOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DefaultExternalAddresses(ctx context.Context, in *QueryDefaultExternalAddressesRequest, opts ...grpc.CallOption) (*QueryDefaultExternalAddressesResponse, error) {
	out := new(QueryDefaultExternalAddressesResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/DefaultExternalAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinks(ctx context.Context, in *QueryApplicationLinksRequest, opts ...grpc.CallOption) (*QueryApplicationLinksResponse, error) {
	out := new(QueryApplicationLinksResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/ApplicationLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinkByClientID(ctx context.Context, in *QueryApplicationLinkByClientIDRequest, opts ...grpc.CallOption) (*QueryApplicationLinkByClientIDResponse, error) {
	out := new(QueryApplicationLinkByClientIDResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/ApplicationLinkByClientID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinkOwners(ctx context.Context, in *QueryApplicationLinkOwnersRequest, opts ...grpc.CallOption) (*QueryApplicationLinkOwnersResponse, error) {
	out := new(QueryApplicationLinkOwnersResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/ApplicationLinkOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Profile queries the profile of a specific user given their DTag or address.
	// If the queried user does not have a profile, the returned response will
	// contain a null profile.
	Profile(context.Context, *QueryProfileRequest) (*QueryProfileResponse, error)
	// IncomingDTagTransferRequests queries all the DTag transfers requests that
	// have been made towards the user with the given address
	IncomingDTagTransferRequests(context.Context, *QueryIncomingDTagTransferRequestsRequest) (*QueryIncomingDTagTransferRequestsResponse, error)
	// ChainLinks queries the chain links associated to the given user, if
	// provided. Otherwise it queries all the chain links stored.
	ChainLinks(context.Context, *QueryChainLinksRequest) (*QueryChainLinksResponse, error)
	// ChainLinkOwners queries for the owners of chain links, optionally searching
	// for a specific chain name and external address
	ChainLinkOwners(context.Context, *QueryChainLinkOwnersRequest) (*QueryChainLinkOwnersResponse, error)
	// DefaultExternalAddresses queries the default addresses associated to the
	// given user and (optionally) chain name
	DefaultExternalAddresses(context.Context, *QueryDefaultExternalAddressesRequest) (*QueryDefaultExternalAddressesResponse, error)
	// ApplicationLinks queries the applications links associated to the given
	// user, if provided. Otherwise, it queries all the application links stored.
	ApplicationLinks(context.Context, *QueryApplicationLinksRequest) (*QueryApplicationLinksResponse, error)
	// ApplicationLinkByClientID queries a single application link for a given
	// client id.
	ApplicationLinkByClientID(context.Context, *QueryApplicationLinkByClientIDRequest) (*QueryApplicationLinkByClientIDResponse, error)
	// ApplicationLinkOwners queries for the owners of applications links,
	// optionally searching for a specific application and username.
	ApplicationLinkOwners(context.Context, *QueryApplicationLinkOwnersRequest) (*QueryApplicationLinkOwnersResponse, error)
	// Params queries the profiles module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Profile(ctx context.Context, req *QueryProfileRequest) (*QueryProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedQueryServer) IncomingDTagTransferRequests(ctx context.Context, req *QueryIncomingDTagTransferRequestsRequest) (*QueryIncomingDTagTransferRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncomingDTagTransferRequests not implemented")
}
func (*UnimplementedQueryServer) ChainLinks(ctx context.Context, req *QueryChainLinksRequest) (*QueryChainLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainLinks not implemented")
}
func (*UnimplementedQueryServer) ChainLinkOwners(ctx context.Context, req *QueryChainLinkOwnersRequest) (*QueryChainLinkOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainLinkOwners not implemented")
}
func (*UnimplementedQueryServer) DefaultExternalAddresses(ctx context.Context, req *QueryDefaultExternalAddressesRequest) (*QueryDefaultExternalAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultExternalAddresses not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinks(ctx context.Context, req *QueryApplicationLinksRequest) (*QueryApplicationLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinks not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinkByClientID(ctx context.Context, req *QueryApplicationLinkByClientIDRequest) (*QueryApplicationLinkByClientIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinkByClientID not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinkOwners(ctx context.Context, req *QueryApplicationLinkOwnersRequest) (*QueryApplicationLinkOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinkOwners not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Profile(ctx, req.(*QueryProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IncomingDTagTransferRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIncomingDTagTransferRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IncomingDTagTransferRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/IncomingDTagTransferRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IncomingDTagTransferRequests(ctx, req.(*QueryIncomingDTagTransferRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChainLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChainLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChainLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/ChainLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainLinks(ctx, req.(*QueryChainLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChainLinkOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChainLinkOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChainLinkOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/ChainLinkOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainLinkOwners(ctx, req.(*QueryChainLinkOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DefaultExternalAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDefaultExternalAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DefaultExternalAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/DefaultExternalAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DefaultExternalAddresses(ctx, req.(*QueryDefaultExternalAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ApplicationLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ApplicationLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/ApplicationLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApplicationLinks(ctx, req.(*QueryApplicationLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ApplicationLinkByClientID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationLinkByClientIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ApplicationLinkByClientID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/ApplicationLinkByClientID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApplicationLinkByClientID(ctx, req.(*QueryApplicationLinkByClientIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ApplicationLinkOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationLinkOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ApplicationLinkOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/ApplicationLinkOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApplicationLinkOwners(ctx, req.(*QueryApplicationLinkOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "huddle.profiles.v3.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Profile",
			Handler:    _Query_Profile_Handler,
		},
		{
			MethodName: "IncomingDTagTransferRequests",
			Handler:    _Query_IncomingDTagTransferRequests_Handler,
		},
		{
			MethodName: "ChainLinks",
			Handler:    _Query_ChainLinks_Handler,
		},
		{
			MethodName: "ChainLinkOwners",
			Handler:    _Query_ChainLinkOwners_Handler,
		},
		{
			MethodName: "DefaultExternalAddresses",
			Handler:    _Query_DefaultExternalAddresses_Handler,
		},
		{
			MethodName: "ApplicationLinks",
			Handler:    _Query_ApplicationLinks_Handler,
		},
		{
			MethodName: "ApplicationLinkByClientID",
			Handler:    _Query_ApplicationLinkByClientID_Handler,
		},
		{
			MethodName: "ApplicationLinkOwners",
			Handler:    _Query_ApplicationLinkOwners_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "huddle/profiles/v3/query.proto",
}
