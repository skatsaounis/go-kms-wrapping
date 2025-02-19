// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package plugin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WrappingClient is the client API for Wrapping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WrappingClient interface {
	Type(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeResponse, error)
	KeyId(ctx context.Context, in *KeyIdRequest, opts ...grpc.CallOption) (*KeyIdResponse, error)
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
	// Init & finalize functions
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Finalize(ctx context.Context, in *FinalizeRequest, opts ...grpc.CallOption) (*FinalizeResponse, error)
	// HMAC related functions
	HmacKeyId(ctx context.Context, in *HmacKeyIdRequest, opts ...grpc.CallOption) (*HmacKeyIdResponse, error)
	// KeyBytes function
	KeyBytes(ctx context.Context, in *KeyBytesRequest, opts ...grpc.CallOption) (*KeyBytesResponse, error)
}

type wrappingClient struct {
	cc grpc.ClientConnInterface
}

func NewWrappingClient(cc grpc.ClientConnInterface) WrappingClient {
	return &wrappingClient{cc}
}

func (c *wrappingClient) Type(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeResponse, error) {
	out := new(TypeResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Type", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) KeyId(ctx context.Context, in *KeyIdRequest, opts ...grpc.CallOption) (*KeyIdResponse, error) {
	out := new(KeyIdResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/KeyId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) Finalize(ctx context.Context, in *FinalizeRequest, opts ...grpc.CallOption) (*FinalizeResponse, error) {
	out := new(FinalizeResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Finalize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) HmacKeyId(ctx context.Context, in *HmacKeyIdRequest, opts ...grpc.CallOption) (*HmacKeyIdResponse, error) {
	out := new(HmacKeyIdResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/HmacKeyId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrappingClient) KeyBytes(ctx context.Context, in *KeyBytesRequest, opts ...grpc.CallOption) (*KeyBytesResponse, error) {
	out := new(KeyBytesResponse)
	err := c.cc.Invoke(ctx, "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/KeyBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WrappingServer is the server API for Wrapping service.
// All implementations must embed UnimplementedWrappingServer
// for forward compatibility
type WrappingServer interface {
	Type(context.Context, *TypeRequest) (*TypeResponse, error)
	KeyId(context.Context, *KeyIdRequest) (*KeyIdResponse, error)
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	// Init & finalize functions
	Init(context.Context, *InitRequest) (*InitResponse, error)
	Finalize(context.Context, *FinalizeRequest) (*FinalizeResponse, error)
	// HMAC related functions
	HmacKeyId(context.Context, *HmacKeyIdRequest) (*HmacKeyIdResponse, error)
	// KeyBytes function
	KeyBytes(context.Context, *KeyBytesRequest) (*KeyBytesResponse, error)
	mustEmbedUnimplementedWrappingServer()
}

// UnimplementedWrappingServer must be embedded to have forward compatible implementations.
type UnimplementedWrappingServer struct {
}

func (UnimplementedWrappingServer) Type(context.Context, *TypeRequest) (*TypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Type not implemented")
}
func (UnimplementedWrappingServer) KeyId(context.Context, *KeyIdRequest) (*KeyIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyId not implemented")
}
func (UnimplementedWrappingServer) SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedWrappingServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedWrappingServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedWrappingServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedWrappingServer) Finalize(context.Context, *FinalizeRequest) (*FinalizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finalize not implemented")
}
func (UnimplementedWrappingServer) HmacKeyId(context.Context, *HmacKeyIdRequest) (*HmacKeyIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HmacKeyId not implemented")
}
func (UnimplementedWrappingServer) KeyBytes(context.Context, *KeyBytesRequest) (*KeyBytesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyBytes not implemented")
}
func (UnimplementedWrappingServer) mustEmbedUnimplementedWrappingServer() {}

// UnsafeWrappingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WrappingServer will
// result in compilation errors.
type UnsafeWrappingServer interface {
	mustEmbedUnimplementedWrappingServer()
}

func RegisterWrappingServer(s grpc.ServiceRegistrar, srv WrappingServer) {
	s.RegisterService(&Wrapping_ServiceDesc, srv)
}

func _Wrapping_Type_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).Type(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Type",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).Type(ctx, req.(*TypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_KeyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).KeyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/KeyId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).KeyId(ctx, req.(*KeyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_Finalize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).Finalize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/Finalize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).Finalize(ctx, req.(*FinalizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_HmacKeyId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HmacKeyIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).HmacKeyId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/HmacKeyId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).HmacKeyId(ctx, req.(*HmacKeyIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wrapping_KeyBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyBytesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrappingServer).KeyBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping/KeyBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrappingServer).KeyBytes(ctx, req.(*KeyBytesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wrapping_ServiceDesc is the grpc.ServiceDesc for Wrapping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wrapping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.hashicorp.go.kms.wrapping.plugin.v2.Wrapping",
	HandlerType: (*WrappingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Type",
			Handler:    _Wrapping_Type_Handler,
		},
		{
			MethodName: "KeyId",
			Handler:    _Wrapping_KeyId_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _Wrapping_SetConfig_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _Wrapping_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _Wrapping_Decrypt_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Wrapping_Init_Handler,
		},
		{
			MethodName: "Finalize",
			Handler:    _Wrapping_Finalize_Handler,
		},
		{
			MethodName: "HmacKeyId",
			Handler:    _Wrapping_HmacKeyId_Handler,
		},
		{
			MethodName: "KeyBytes",
			Handler:    _Wrapping_KeyBytes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin/github.com.hashicorp.go.kms.wrapping.plugin.v2.proto",
}
