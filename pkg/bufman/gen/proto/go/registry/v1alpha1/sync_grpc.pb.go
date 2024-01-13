// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: registry/v1alpha1/sync.proto

package registryv1alpha1

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
	SyncService_GetGitSyncPoint_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.SyncService/GetGitSyncPoint"
	SyncService_SyncGitCommit_FullMethodName   = "/bufman.dubbo.apache.org.registry.v1alpha1.SyncService/SyncGitCommit"
)

// SyncServiceClient is the client API for SyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncServiceClient interface {
	// GetGitSyncPoint retrieves the Git sync point for the named repository
	// on the specified branch.
	GetGitSyncPoint(ctx context.Context, in *GetGitSyncPointRequest, opts ...grpc.CallOption) (*GetGitSyncPointResponse, error)
	// SyncGitCommit syncs a Git commit containing a module to a named repository.
	SyncGitCommit(ctx context.Context, in *SyncGitCommitRequest, opts ...grpc.CallOption) (*SyncGitCommitResponse, error)
}

type syncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncServiceClient(cc grpc.ClientConnInterface) SyncServiceClient {
	return &syncServiceClient{cc}
}

func (c *syncServiceClient) GetGitSyncPoint(ctx context.Context, in *GetGitSyncPointRequest, opts ...grpc.CallOption) (*GetGitSyncPointResponse, error) {
	out := new(GetGitSyncPointResponse)
	err := c.cc.Invoke(ctx, SyncService_GetGitSyncPoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *syncServiceClient) SyncGitCommit(ctx context.Context, in *SyncGitCommitRequest, opts ...grpc.CallOption) (*SyncGitCommitResponse, error) {
	out := new(SyncGitCommitResponse)
	err := c.cc.Invoke(ctx, SyncService_SyncGitCommit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncServiceServer is the server API for SyncService service.
// All implementations must embed UnimplementedSyncServiceServer
// for forward compatibility
type SyncServiceServer interface {
	// GetGitSyncPoint retrieves the Git sync point for the named repository
	// on the specified branch.
	GetGitSyncPoint(context.Context, *GetGitSyncPointRequest) (*GetGitSyncPointResponse, error)
	// SyncGitCommit syncs a Git commit containing a module to a named repository.
	SyncGitCommit(context.Context, *SyncGitCommitRequest) (*SyncGitCommitResponse, error)
	mustEmbedUnimplementedSyncServiceServer()
}

// UnimplementedSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSyncServiceServer struct {
}

func (UnimplementedSyncServiceServer) GetGitSyncPoint(context.Context, *GetGitSyncPointRequest) (*GetGitSyncPointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGitSyncPoint not implemented")
}
func (UnimplementedSyncServiceServer) SyncGitCommit(context.Context, *SyncGitCommitRequest) (*SyncGitCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncGitCommit not implemented")
}
func (UnimplementedSyncServiceServer) mustEmbedUnimplementedSyncServiceServer() {}

// UnsafeSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncServiceServer will
// result in compilation errors.
type UnsafeSyncServiceServer interface {
	mustEmbedUnimplementedSyncServiceServer()
}

func RegisterSyncServiceServer(s grpc.ServiceRegistrar, srv SyncServiceServer) {
	s.RegisterService(&SyncService_ServiceDesc, srv)
}

func _SyncService_GetGitSyncPoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGitSyncPointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).GetGitSyncPoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_GetGitSyncPoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).GetGitSyncPoint(ctx, req.(*GetGitSyncPointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SyncService_SyncGitCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncGitCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncServiceServer).SyncGitCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncService_SyncGitCommit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncServiceServer).SyncGitCommit(ctx, req.(*SyncGitCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncService_ServiceDesc is the grpc.ServiceDesc for SyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.SyncService",
	HandlerType: (*SyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGitSyncPoint",
			Handler:    _SyncService_GetGitSyncPoint_Handler,
		},
		{
			MethodName: "SyncGitCommit",
			Handler:    _SyncService_SyncGitCommit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/sync.proto",
}
