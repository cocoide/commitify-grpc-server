// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: generate_message_service.proto

package api

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

const (
	GenerateMessageService_GenerateJapaneseCommitMessage_FullMethodName     = "/generate_message.GenerateMessageService/GenerateJapaneseCommitMessage"
	GenerateMessageService_GenerateEnglishCommitMessage_FullMethodName      = "/generate_message.GenerateMessageService/GenerateEnglishCommitMessage"
	GenerateMessageService_GeneratePrefixFormatCommitMessage_FullMethodName = "/generate_message.GenerateMessageService/GeneratePrefixFormatCommitMessage"
)

// GenerateMessageServiceClient is the client API for GenerateMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenerateMessageServiceClient interface {
	GenerateJapaneseCommitMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error)
	GenerateEnglishCommitMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error)
	GeneratePrefixFormatCommitMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error)
}

type generateMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGenerateMessageServiceClient(cc grpc.ClientConnInterface) GenerateMessageServiceClient {
	return &generateMessageServiceClient{cc}
}

func (c *generateMessageServiceClient) GenerateJapaneseCommitMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error) {
	out := new(GenerateMessageResponse)
	err := c.cc.Invoke(ctx, GenerateMessageService_GenerateJapaneseCommitMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateMessageServiceClient) GenerateEnglishCommitMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error) {
	out := new(GenerateMessageResponse)
	err := c.cc.Invoke(ctx, GenerateMessageService_GenerateEnglishCommitMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generateMessageServiceClient) GeneratePrefixFormatCommitMessage(ctx context.Context, in *GenerateMessageRequest, opts ...grpc.CallOption) (*GenerateMessageResponse, error) {
	out := new(GenerateMessageResponse)
	err := c.cc.Invoke(ctx, GenerateMessageService_GeneratePrefixFormatCommitMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateMessageServiceServer is the server API for GenerateMessageService service.
// All implementations must embed UnimplementedGenerateMessageServiceServer
// for forward compatibility
type GenerateMessageServiceServer interface {
	GenerateJapaneseCommitMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error)
	GenerateEnglishCommitMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error)
	GeneratePrefixFormatCommitMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error)
	mustEmbedUnimplementedGenerateMessageServiceServer()
}

// UnimplementedGenerateMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGenerateMessageServiceServer struct {
}

func (UnimplementedGenerateMessageServiceServer) GenerateJapaneseCommitMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateJapaneseCommitMessage not implemented")
}
func (UnimplementedGenerateMessageServiceServer) GenerateEnglishCommitMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateEnglishCommitMessage not implemented")
}
func (UnimplementedGenerateMessageServiceServer) GeneratePrefixFormatCommitMessage(context.Context, *GenerateMessageRequest) (*GenerateMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratePrefixFormatCommitMessage not implemented")
}
func (UnimplementedGenerateMessageServiceServer) mustEmbedUnimplementedGenerateMessageServiceServer() {
}

// UnsafeGenerateMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenerateMessageServiceServer will
// result in compilation errors.
type UnsafeGenerateMessageServiceServer interface {
	mustEmbedUnimplementedGenerateMessageServiceServer()
}

func RegisterGenerateMessageServiceServer(s grpc.ServiceRegistrar, srv GenerateMessageServiceServer) {
	s.RegisterService(&GenerateMessageService_ServiceDesc, srv)
}

func _GenerateMessageService_GenerateJapaneseCommitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateMessageServiceServer).GenerateJapaneseCommitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerateMessageService_GenerateJapaneseCommitMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateMessageServiceServer).GenerateJapaneseCommitMessage(ctx, req.(*GenerateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenerateMessageService_GenerateEnglishCommitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateMessageServiceServer).GenerateEnglishCommitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerateMessageService_GenerateEnglishCommitMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateMessageServiceServer).GenerateEnglishCommitMessage(ctx, req.(*GenerateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GenerateMessageService_GeneratePrefixFormatCommitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateMessageServiceServer).GeneratePrefixFormatCommitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenerateMessageService_GeneratePrefixFormatCommitMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateMessageServiceServer).GeneratePrefixFormatCommitMessage(ctx, req.(*GenerateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenerateMessageService_ServiceDesc is the grpc.ServiceDesc for GenerateMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenerateMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generate_message.GenerateMessageService",
	HandlerType: (*GenerateMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateJapaneseCommitMessage",
			Handler:    _GenerateMessageService_GenerateJapaneseCommitMessage_Handler,
		},
		{
			MethodName: "GenerateEnglishCommitMessage",
			Handler:    _GenerateMessageService_GenerateEnglishCommitMessage_Handler,
		},
		{
			MethodName: "GeneratePrefixFormatCommitMessage",
			Handler:    _GenerateMessageService_GeneratePrefixFormatCommitMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generate_message_service.proto",
}
