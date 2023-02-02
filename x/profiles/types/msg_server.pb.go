// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: huddle/profiles/v3/msg_server.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

func init() {
	proto.RegisterFile("huddle/profiles/v3/msg_server.proto", fileDescriptor_7e9e5831a8a59c64)
}

var fileDescriptor_7e9e5831a8a59c64 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x5b, 0x21, 0x71, 0x30, 0x42, 0x80, 0xe1, 0xb2, 0x1c, 0x72, 0xe0, 0xdf, 0x18, 0xb0,
	0xa4, 0xac, 0x17, 0xae, 0x65, 0xe5, 0x82, 0x98, 0x84, 0xba, 0x71, 0x41, 0x42, 0x95, 0xeb, 0xbc,
	0x75, 0xad, 0xb9, 0x76, 0xb0, 0x9d, 0x6a, 0xe3, 0x0b, 0x70, 0xe5, 0x53, 0xf0, 0x59, 0x38, 0xee,
	0xc8, 0x11, 0xb5, 0x5f, 0x04, 0xb5, 0x69, 0x43, 0xd2, 0xf5, 0xed, 0xd2, 0x5b, 0x92, 0xf7, 0xf7,
	0xbc, 0xcf, 0xe3, 0x47, 0x8a, 0xc9, 0x93, 0x51, 0x96, 0x24, 0x0a, 0xe2, 0xd4, 0x9a, 0xa1, 0x54,
	0xe0, 0xe2, 0x49, 0x3b, 0x1e, 0x3b, 0xd1, 0x77, 0x60, 0x27, 0x60, 0xa3, 0xd4, 0x1a, 0x6f, 0x28,
	0xcd, 0xa1, 0x68, 0x05, 0x45, 0x93, 0x76, 0xf0, 0x48, 0x18, 0x61, 0x16, 0xe3, 0x78, 0xfe, 0x94,
	0x93, 0xc1, 0x9e, 0x30, 0x46, 0xe4, 0xeb, 0xbc, 0x19, 0x64, 0xc3, 0x98, 0xe9, 0xcb, 0xd5, 0x88,
	0x1b, 0x37, 0x36, 0xae, 0x9f, 0x6b, 0xf2, 0x97, 0xe5, 0x68, 0x7f, 0x53, 0x08, 0x93, 0x80, 0x5a,
	0xd0, 0xf3, 0x4f, 0x4b, 0xf0, 0x10, 0x07, 0x13, 0xcf, 0x44, 0xdf, 0xc2, 0xb7, 0x0c, 0x9c, 0x5f,
	0xed, 0x7d, 0xb6, 0xf9, 0x70, 0xeb, 0x5b, 0x5f, 0x61, 0xd8, 0xa6, 0x9d, 0x07, 0x18, 0xcc, 0x47,
	0x4c, 0xea, 0xbe, 0x92, 0xfa, 0x7c, 0xeb, 0xb1, 0xe6, 0x28, 0x4b, 0xd3, 0x32, 0x78, 0xf4, 0x8b,
	0x90, 0x5b, 0x27, 0x4e, 0xd0, 0xaf, 0xe4, 0xce, 0x29, 0x9b, 0xc0, 0xa7, 0x9c, 0xa7, 0x8f, 0xa3,
	0xeb, 0xbd, 0x47, 0x27, 0x4e, 0x94, 0x98, 0xe0, 0xe5, 0xcd, 0x4c, 0x0f, 0x5c, 0x6a, 0xb4, 0x03,
	0xca, 0xc9, 0xdd, 0x2e, 0x28, 0xf0, 0x85, 0xc1, 0x53, 0x44, 0x5c, 0xa1, 0x82, 0xd7, 0x75, 0xa8,
	0xc2, 0x24, 0x23, 0x0f, 0x7b, 0x79, 0x63, 0xdd, 0x33, 0x26, 0xce, 0x2c, 0xd3, 0x6e, 0x08, 0x96,
	0x62, 0x39, 0x37, 0xb0, 0xc1, 0x51, 0x7d, 0xb6, 0xb0, 0xfd, 0xd1, 0x24, 0x7b, 0xc7, 0x4c, 0x73,
	0x50, 0xd5, 0xf1, 0x42, 0x41, 0x5b, 0xc8, 0x46, 0x54, 0x11, 0xbc, 0xdd, 0x55, 0x51, 0x49, 0xd2,
	0xe1, 0x1c, 0x52, 0xbf, 0x4b, 0x12, 0x54, 0x81, 0x26, 0x41, 0x15, 0x95, 0x24, 0x3d, 0x18, 0x66,
	0x0e, 0x76, 0x49, 0x82, 0x2a, 0xd0, 0x24, 0xa8, 0xa2, 0x48, 0xa2, 0xc8, 0xfd, 0x8f, 0x52, 0x9f,
	0x1f, 0xcf, 0x7f, 0x91, 0x0e, 0xe7, 0x26, 0xd3, 0x9e, 0xee, 0x23, 0xdb, 0xd6, 0xc1, 0x20, 0xae,
	0x09, 0x16, 0x6e, 0x96, 0xd0, 0xcf, 0x5a, 0xad, 0xfb, 0x1d, 0x20, 0x6b, 0xae, 0xa3, 0xc1, 0x9b,
	0xda, 0x68, 0xa5, 0xeb, 0x53, 0xf0, 0x5d, 0x18, 0xb2, 0x4c, 0xf9, 0xf7, 0x17, 0x1e, 0xac, 0x66,
	0xaa, 0x93, 0x24, 0x16, 0x9c, 0x43, 0xbb, 0x46, 0x15, 0x68, 0xd7, 0xa8, 0xa2, 0x48, 0x22, 0xc9,
	0xbd, 0x79, 0x33, 0x9d, 0x34, 0x55, 0x92, 0x33, 0x2f, 0x8d, 0xa6, 0xcf, 0xb7, 0x34, 0x58, 0xe2,
	0x82, 0xa8, 0x1e, 0x57, 0x58, 0x19, 0xf2, 0x20, 0xaf, 0xa4, 0x6c, 0xf6, 0x62, 0x6b, 0x79, 0x65,
	0xbb, 0x56, 0x5d, 0x72, 0x65, 0xf8, 0xee, 0xc3, 0xef, 0x69, 0xd8, 0xbc, 0x9a, 0x86, 0xcd, 0xbf,
	0xd3, 0xb0, 0xf9, 0x73, 0x16, 0x36, 0xae, 0x66, 0x61, 0xe3, 0xcf, 0x2c, 0x6c, 0x7c, 0x69, 0x09,
	0xe9, 0x47, 0xd9, 0x20, 0xe2, 0x66, 0x1c, 0x0b, 0x2b, 0x13, 0x69, 0x8d, 0x3e, 0xfc, 0x6e, 0x34,
	0xc4, 0xcb, 0x4b, 0xf8, 0xe2, 0xff, 0x35, 0xec, 0x2f, 0x53, 0x70, 0x83, 0xdb, 0x8b, 0xbb, 0xb7,
	0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xcc, 0x68, 0x3e, 0x02, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SaveProfile defines the method to save a profile
	SaveProfile(ctx context.Context, in *MsgSaveProfile, opts ...grpc.CallOption) (*MsgSaveProfileResponse, error)
	// DeleteProfile defines the method to delete an existing profile
	DeleteProfile(ctx context.Context, in *MsgDeleteProfile, opts ...grpc.CallOption) (*MsgDeleteProfileResponse, error)
	// RequestDTagTransfer defines the method to request another user to transfer
	// their DTag to you
	RequestDTagTransfer(ctx context.Context, in *MsgRequestDTagTransfer, opts ...grpc.CallOption) (*MsgRequestDTagTransferResponse, error)
	// CancelDTagTransferRequest defines the method to cancel an outgoing DTag
	// transfer request
	CancelDTagTransferRequest(ctx context.Context, in *MsgCancelDTagTransferRequest, opts ...grpc.CallOption) (*MsgCancelDTagTransferRequestResponse, error)
	// AcceptDTagTransferRequest defines the method to accept an incoming DTag
	// transfer request
	AcceptDTagTransferRequest(ctx context.Context, in *MsgAcceptDTagTransferRequest, opts ...grpc.CallOption) (*MsgAcceptDTagTransferRequestResponse, error)
	// RefuseDTagTransferRequest defines the method to refuse an incoming DTag
	// transfer request
	RefuseDTagTransferRequest(ctx context.Context, in *MsgRefuseDTagTransferRequest, opts ...grpc.CallOption) (*MsgRefuseDTagTransferRequestResponse, error)
	// LinkChainAccount defines a method to link an external chain account to a
	// profile
	LinkChainAccount(ctx context.Context, in *MsgLinkChainAccount, opts ...grpc.CallOption) (*MsgLinkChainAccountResponse, error)
	// UnlinkChainAccount defines a method to unlink an external chain account
	// from a profile
	UnlinkChainAccount(ctx context.Context, in *MsgUnlinkChainAccount, opts ...grpc.CallOption) (*MsgUnlinkChainAccountResponse, error)
	// SetDefaultExternalAddress allows to set a specific external address as the
	// default one for a given chain
	SetDefaultExternalAddress(ctx context.Context, in *MsgSetDefaultExternalAddress, opts ...grpc.CallOption) (*MsgSetDefaultExternalAddressResponse, error)
	// LinkApplication defines a method to create a centralized application
	// link
	LinkApplication(ctx context.Context, in *MsgLinkApplication, opts ...grpc.CallOption) (*MsgLinkApplicationResponse, error)
	// UnlinkApplication defines a method to remove a centralized application
	UnlinkApplication(ctx context.Context, in *MsgUnlinkApplication, opts ...grpc.CallOption) (*MsgUnlinkApplicationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SaveProfile(ctx context.Context, in *MsgSaveProfile, opts ...grpc.CallOption) (*MsgSaveProfileResponse, error) {
	out := new(MsgSaveProfileResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/SaveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteProfile(ctx context.Context, in *MsgDeleteProfile, opts ...grpc.CallOption) (*MsgDeleteProfileResponse, error) {
	out := new(MsgDeleteProfileResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/DeleteProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestDTagTransfer(ctx context.Context, in *MsgRequestDTagTransfer, opts ...grpc.CallOption) (*MsgRequestDTagTransferResponse, error) {
	out := new(MsgRequestDTagTransferResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/RequestDTagTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelDTagTransferRequest(ctx context.Context, in *MsgCancelDTagTransferRequest, opts ...grpc.CallOption) (*MsgCancelDTagTransferRequestResponse, error) {
	out := new(MsgCancelDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/CancelDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptDTagTransferRequest(ctx context.Context, in *MsgAcceptDTagTransferRequest, opts ...grpc.CallOption) (*MsgAcceptDTagTransferRequestResponse, error) {
	out := new(MsgAcceptDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/AcceptDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RefuseDTagTransferRequest(ctx context.Context, in *MsgRefuseDTagTransferRequest, opts ...grpc.CallOption) (*MsgRefuseDTagTransferRequestResponse, error) {
	out := new(MsgRefuseDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/RefuseDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LinkChainAccount(ctx context.Context, in *MsgLinkChainAccount, opts ...grpc.CallOption) (*MsgLinkChainAccountResponse, error) {
	out := new(MsgLinkChainAccountResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/LinkChainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnlinkChainAccount(ctx context.Context, in *MsgUnlinkChainAccount, opts ...grpc.CallOption) (*MsgUnlinkChainAccountResponse, error) {
	out := new(MsgUnlinkChainAccountResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/UnlinkChainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetDefaultExternalAddress(ctx context.Context, in *MsgSetDefaultExternalAddress, opts ...grpc.CallOption) (*MsgSetDefaultExternalAddressResponse, error) {
	out := new(MsgSetDefaultExternalAddressResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/SetDefaultExternalAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LinkApplication(ctx context.Context, in *MsgLinkApplication, opts ...grpc.CallOption) (*MsgLinkApplicationResponse, error) {
	out := new(MsgLinkApplicationResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/LinkApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnlinkApplication(ctx context.Context, in *MsgUnlinkApplication, opts ...grpc.CallOption) (*MsgUnlinkApplicationResponse, error) {
	out := new(MsgUnlinkApplicationResponse)
	err := c.cc.Invoke(ctx, "/huddle.profiles.v3.Msg/UnlinkApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SaveProfile defines the method to save a profile
	SaveProfile(context.Context, *MsgSaveProfile) (*MsgSaveProfileResponse, error)
	// DeleteProfile defines the method to delete an existing profile
	DeleteProfile(context.Context, *MsgDeleteProfile) (*MsgDeleteProfileResponse, error)
	// RequestDTagTransfer defines the method to request another user to transfer
	// their DTag to you
	RequestDTagTransfer(context.Context, *MsgRequestDTagTransfer) (*MsgRequestDTagTransferResponse, error)
	// CancelDTagTransferRequest defines the method to cancel an outgoing DTag
	// transfer request
	CancelDTagTransferRequest(context.Context, *MsgCancelDTagTransferRequest) (*MsgCancelDTagTransferRequestResponse, error)
	// AcceptDTagTransferRequest defines the method to accept an incoming DTag
	// transfer request
	AcceptDTagTransferRequest(context.Context, *MsgAcceptDTagTransferRequest) (*MsgAcceptDTagTransferRequestResponse, error)
	// RefuseDTagTransferRequest defines the method to refuse an incoming DTag
	// transfer request
	RefuseDTagTransferRequest(context.Context, *MsgRefuseDTagTransferRequest) (*MsgRefuseDTagTransferRequestResponse, error)
	// LinkChainAccount defines a method to link an external chain account to a
	// profile
	LinkChainAccount(context.Context, *MsgLinkChainAccount) (*MsgLinkChainAccountResponse, error)
	// UnlinkChainAccount defines a method to unlink an external chain account
	// from a profile
	UnlinkChainAccount(context.Context, *MsgUnlinkChainAccount) (*MsgUnlinkChainAccountResponse, error)
	// SetDefaultExternalAddress allows to set a specific external address as the
	// default one for a given chain
	SetDefaultExternalAddress(context.Context, *MsgSetDefaultExternalAddress) (*MsgSetDefaultExternalAddressResponse, error)
	// LinkApplication defines a method to create a centralized application
	// link
	LinkApplication(context.Context, *MsgLinkApplication) (*MsgLinkApplicationResponse, error)
	// UnlinkApplication defines a method to remove a centralized application
	UnlinkApplication(context.Context, *MsgUnlinkApplication) (*MsgUnlinkApplicationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SaveProfile(ctx context.Context, req *MsgSaveProfile) (*MsgSaveProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProfile not implemented")
}
func (*UnimplementedMsgServer) DeleteProfile(ctx context.Context, req *MsgDeleteProfile) (*MsgDeleteProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (*UnimplementedMsgServer) RequestDTagTransfer(ctx context.Context, req *MsgRequestDTagTransfer) (*MsgRequestDTagTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDTagTransfer not implemented")
}
func (*UnimplementedMsgServer) CancelDTagTransferRequest(ctx context.Context, req *MsgCancelDTagTransferRequest) (*MsgCancelDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) AcceptDTagTransferRequest(ctx context.Context, req *MsgAcceptDTagTransferRequest) (*MsgAcceptDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) RefuseDTagTransferRequest(ctx context.Context, req *MsgRefuseDTagTransferRequest) (*MsgRefuseDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefuseDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) LinkChainAccount(ctx context.Context, req *MsgLinkChainAccount) (*MsgLinkChainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkChainAccount not implemented")
}
func (*UnimplementedMsgServer) UnlinkChainAccount(ctx context.Context, req *MsgUnlinkChainAccount) (*MsgUnlinkChainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkChainAccount not implemented")
}
func (*UnimplementedMsgServer) SetDefaultExternalAddress(ctx context.Context, req *MsgSetDefaultExternalAddress) (*MsgSetDefaultExternalAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDefaultExternalAddress not implemented")
}
func (*UnimplementedMsgServer) LinkApplication(ctx context.Context, req *MsgLinkApplication) (*MsgLinkApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkApplication not implemented")
}
func (*UnimplementedMsgServer) UnlinkApplication(ctx context.Context, req *MsgUnlinkApplication) (*MsgUnlinkApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkApplication not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SaveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSaveProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SaveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/SaveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SaveProfile(ctx, req.(*MsgSaveProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/DeleteProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteProfile(ctx, req.(*MsgDeleteProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestDTagTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestDTagTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestDTagTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/RequestDTagTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestDTagTransfer(ctx, req.(*MsgRequestDTagTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/CancelDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelDTagTransferRequest(ctx, req.(*MsgCancelDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/AcceptDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptDTagTransferRequest(ctx, req.(*MsgAcceptDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RefuseDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRefuseDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RefuseDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/RefuseDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RefuseDTagTransferRequest(ctx, req.(*MsgRefuseDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LinkChainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLinkChainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LinkChainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/LinkChainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LinkChainAccount(ctx, req.(*MsgLinkChainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnlinkChainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnlinkChainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnlinkChainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/UnlinkChainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnlinkChainAccount(ctx, req.(*MsgUnlinkChainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetDefaultExternalAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetDefaultExternalAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetDefaultExternalAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/SetDefaultExternalAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetDefaultExternalAddress(ctx, req.(*MsgSetDefaultExternalAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LinkApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLinkApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LinkApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/LinkApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LinkApplication(ctx, req.(*MsgLinkApplication))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnlinkApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnlinkApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnlinkApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huddle.profiles.v3.Msg/UnlinkApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnlinkApplication(ctx, req.(*MsgUnlinkApplication))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "huddle.profiles.v3.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveProfile",
			Handler:    _Msg_SaveProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _Msg_DeleteProfile_Handler,
		},
		{
			MethodName: "RequestDTagTransfer",
			Handler:    _Msg_RequestDTagTransfer_Handler,
		},
		{
			MethodName: "CancelDTagTransferRequest",
			Handler:    _Msg_CancelDTagTransferRequest_Handler,
		},
		{
			MethodName: "AcceptDTagTransferRequest",
			Handler:    _Msg_AcceptDTagTransferRequest_Handler,
		},
		{
			MethodName: "RefuseDTagTransferRequest",
			Handler:    _Msg_RefuseDTagTransferRequest_Handler,
		},
		{
			MethodName: "LinkChainAccount",
			Handler:    _Msg_LinkChainAccount_Handler,
		},
		{
			MethodName: "UnlinkChainAccount",
			Handler:    _Msg_UnlinkChainAccount_Handler,
		},
		{
			MethodName: "SetDefaultExternalAddress",
			Handler:    _Msg_SetDefaultExternalAddress_Handler,
		},
		{
			MethodName: "LinkApplication",
			Handler:    _Msg_LinkApplication_Handler,
		},
		{
			MethodName: "UnlinkApplication",
			Handler:    _Msg_UnlinkApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "huddle/profiles/v3/msg_server.proto",
}
