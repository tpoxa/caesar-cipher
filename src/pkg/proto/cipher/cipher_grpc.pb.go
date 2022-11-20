// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cipher

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

// CipherClient is the client API for Cipher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CipherClient interface {
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type cipherClient struct {
	cc grpc.ClientConnInterface
}

func NewCipherClient(cc grpc.ClientConnInterface) CipherClient {
	return &cipherClient{cc}
}

func (c *cipherClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/api.Cipher/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cipherClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/api.Cipher/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CipherServer is the server API for Cipher service.
// All implementations must embed UnimplementedCipherServer
// for forward compatibility
type CipherServer interface {
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	mustEmbedUnimplementedCipherServer()
}

// UnimplementedCipherServer must be embedded to have forward compatible implementations.
type UnimplementedCipherServer struct {
}

func (UnimplementedCipherServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedCipherServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedCipherServer) mustEmbedUnimplementedCipherServer() {}

// UnsafeCipherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CipherServer will
// result in compilation errors.
type UnsafeCipherServer interface {
	mustEmbedUnimplementedCipherServer()
}

func RegisterCipherServer(s grpc.ServiceRegistrar, srv CipherServer) {
	s.RegisterService(&Cipher_ServiceDesc, srv)
}

func _Cipher_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Cipher/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cipher_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CipherServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Cipher/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CipherServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cipher_ServiceDesc is the grpc.ServiceDesc for Cipher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cipher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Cipher",
	HandlerType: (*CipherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _Cipher_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _Cipher_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/cipher/cipher.proto",
}
