// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: registry/v1alpha1/search.proto

package registryv1alpha1connect

import (
	context "context"
	errors "errors"
	v1alpha1 "github.com/apache/dubbo-kubernetes/pkg/bufman/gen/proto/go/registry/v1alpha1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion1_7_0

const (
	// SearchServiceName is the fully-qualified name of the SearchService service.
	SearchServiceName = "bufman.dubbo.apache.org.registry.v1alpha1.SearchService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SearchServiceSearchUserProcedure is the fully-qualified name of the SearchService's SearchUser
	// RPC.
	SearchServiceSearchUserProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchUser"
	// SearchServiceSearchRepositoryProcedure is the fully-qualified name of the SearchService's
	// SearchRepository RPC.
	SearchServiceSearchRepositoryProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchRepository"
	// SearchServiceSearchLastCommitByContentProcedure is the fully-qualified name of the
	// SearchService's SearchLastCommitByContent RPC.
	SearchServiceSearchLastCommitByContentProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchLastCommitByContent"
	// SearchServiceSearchCurationPluginProcedure is the fully-qualified name of the SearchService's
	// SearchCurationPlugin RPC.
	SearchServiceSearchCurationPluginProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchCurationPlugin"
	// SearchServiceSearchTagProcedure is the fully-qualified name of the SearchService's SearchTag RPC.
	SearchServiceSearchTagProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchTag"
	// SearchServiceSearchDraftProcedure is the fully-qualified name of the SearchService's SearchDraft
	// RPC.
	SearchServiceSearchDraftProcedure = "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/SearchDraft"
)

// SearchServiceClient is a client for the bufman.dubbo.apache.org.registry.v1alpha1.SearchService
// service.
type SearchServiceClient interface {
	// SearchUser searches users by username
	SearchUser(context.Context, *connect_go.Request[v1alpha1.SearchUserRequest]) (*connect_go.Response[v1alpha1.SearchUserResponse], error)
	// SearchRepository searches repositories by name or description
	SearchRepository(context.Context, *connect_go.Request[v1alpha1.SearchRepositoryRequest]) (*connect_go.Response[v1alpha1.SearchRepositoryResponse], error)
	// SearchCommitByContent searches last commit in same repo by idl content
	// that means, for a repo, search results only record last matched commit
	SearchLastCommitByContent(context.Context, *connect_go.Request[v1alpha1.SearchLastCommitByContentRequest]) (*connect_go.Response[v1alpha1.SearchLastCommitByContentResponse], error)
	// SearchCurationPlugin search plugins by name or description
	SearchCurationPlugin(context.Context, *connect_go.Request[v1alpha1.SearchCuratedPluginRequest]) (*connect_go.Response[v1alpha1.SearchCuratedPluginResponse], error)
	// SearchTag searches for tags in a repository
	SearchTag(context.Context, *connect_go.Request[v1alpha1.SearchTagRequest]) (*connect_go.Response[v1alpha1.SearchTagResponse], error)
	// SearchDraft searches for drafts in a repository
	SearchDraft(context.Context, *connect_go.Request[v1alpha1.SearchDraftRequest]) (*connect_go.Response[v1alpha1.SearchDraftResponse], error)
}

// NewSearchServiceClient constructs a client for the
// bufman.dubbo.apache.org.registry.v1alpha1.SearchService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSearchServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) SearchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &searchServiceClient{
		searchUser: connect_go.NewClient[v1alpha1.SearchUserRequest, v1alpha1.SearchUserResponse](
			httpClient,
			baseURL+SearchServiceSearchUserProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		searchRepository: connect_go.NewClient[v1alpha1.SearchRepositoryRequest, v1alpha1.SearchRepositoryResponse](
			httpClient,
			baseURL+SearchServiceSearchRepositoryProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		searchLastCommitByContent: connect_go.NewClient[v1alpha1.SearchLastCommitByContentRequest, v1alpha1.SearchLastCommitByContentResponse](
			httpClient,
			baseURL+SearchServiceSearchLastCommitByContentProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		searchCurationPlugin: connect_go.NewClient[v1alpha1.SearchCuratedPluginRequest, v1alpha1.SearchCuratedPluginResponse](
			httpClient,
			baseURL+SearchServiceSearchCurationPluginProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		searchTag: connect_go.NewClient[v1alpha1.SearchTagRequest, v1alpha1.SearchTagResponse](
			httpClient,
			baseURL+SearchServiceSearchTagProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
		searchDraft: connect_go.NewClient[v1alpha1.SearchDraftRequest, v1alpha1.SearchDraftResponse](
			httpClient,
			baseURL+SearchServiceSearchDraftProcedure,
			connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
			connect_go.WithClientOptions(opts...),
		),
	}
}

// searchServiceClient implements SearchServiceClient.
type searchServiceClient struct {
	searchUser                *connect_go.Client[v1alpha1.SearchUserRequest, v1alpha1.SearchUserResponse]
	searchRepository          *connect_go.Client[v1alpha1.SearchRepositoryRequest, v1alpha1.SearchRepositoryResponse]
	searchLastCommitByContent *connect_go.Client[v1alpha1.SearchLastCommitByContentRequest, v1alpha1.SearchLastCommitByContentResponse]
	searchCurationPlugin      *connect_go.Client[v1alpha1.SearchCuratedPluginRequest, v1alpha1.SearchCuratedPluginResponse]
	searchTag                 *connect_go.Client[v1alpha1.SearchTagRequest, v1alpha1.SearchTagResponse]
	searchDraft               *connect_go.Client[v1alpha1.SearchDraftRequest, v1alpha1.SearchDraftResponse]
}

// SearchUser calls bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchUser.
func (c *searchServiceClient) SearchUser(ctx context.Context, req *connect_go.Request[v1alpha1.SearchUserRequest]) (*connect_go.Response[v1alpha1.SearchUserResponse], error) {
	return c.searchUser.CallUnary(ctx, req)
}

// SearchRepository calls bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchRepository.
func (c *searchServiceClient) SearchRepository(ctx context.Context, req *connect_go.Request[v1alpha1.SearchRepositoryRequest]) (*connect_go.Response[v1alpha1.SearchRepositoryResponse], error) {
	return c.searchRepository.CallUnary(ctx, req)
}

// SearchLastCommitByContent calls
// bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchLastCommitByContent.
func (c *searchServiceClient) SearchLastCommitByContent(ctx context.Context, req *connect_go.Request[v1alpha1.SearchLastCommitByContentRequest]) (*connect_go.Response[v1alpha1.SearchLastCommitByContentResponse], error) {
	return c.searchLastCommitByContent.CallUnary(ctx, req)
}

// SearchCurationPlugin calls
// bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchCurationPlugin.
func (c *searchServiceClient) SearchCurationPlugin(ctx context.Context, req *connect_go.Request[v1alpha1.SearchCuratedPluginRequest]) (*connect_go.Response[v1alpha1.SearchCuratedPluginResponse], error) {
	return c.searchCurationPlugin.CallUnary(ctx, req)
}

// SearchTag calls bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchTag.
func (c *searchServiceClient) SearchTag(ctx context.Context, req *connect_go.Request[v1alpha1.SearchTagRequest]) (*connect_go.Response[v1alpha1.SearchTagResponse], error) {
	return c.searchTag.CallUnary(ctx, req)
}

// SearchDraft calls bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchDraft.
func (c *searchServiceClient) SearchDraft(ctx context.Context, req *connect_go.Request[v1alpha1.SearchDraftRequest]) (*connect_go.Response[v1alpha1.SearchDraftResponse], error) {
	return c.searchDraft.CallUnary(ctx, req)
}

// SearchServiceHandler is an implementation of the
// bufman.dubbo.apache.org.registry.v1alpha1.SearchService service.
type SearchServiceHandler interface {
	// SearchUser searches users by username
	SearchUser(context.Context, *connect_go.Request[v1alpha1.SearchUserRequest]) (*connect_go.Response[v1alpha1.SearchUserResponse], error)
	// SearchRepository searches repositories by name or description
	SearchRepository(context.Context, *connect_go.Request[v1alpha1.SearchRepositoryRequest]) (*connect_go.Response[v1alpha1.SearchRepositoryResponse], error)
	// SearchCommitByContent searches last commit in same repo by idl content
	// that means, for a repo, search results only record last matched commit
	SearchLastCommitByContent(context.Context, *connect_go.Request[v1alpha1.SearchLastCommitByContentRequest]) (*connect_go.Response[v1alpha1.SearchLastCommitByContentResponse], error)
	// SearchCurationPlugin search plugins by name or description
	SearchCurationPlugin(context.Context, *connect_go.Request[v1alpha1.SearchCuratedPluginRequest]) (*connect_go.Response[v1alpha1.SearchCuratedPluginResponse], error)
	// SearchTag searches for tags in a repository
	SearchTag(context.Context, *connect_go.Request[v1alpha1.SearchTagRequest]) (*connect_go.Response[v1alpha1.SearchTagResponse], error)
	// SearchDraft searches for drafts in a repository
	SearchDraft(context.Context, *connect_go.Request[v1alpha1.SearchDraftRequest]) (*connect_go.Response[v1alpha1.SearchDraftResponse], error)
}

// NewSearchServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSearchServiceHandler(svc SearchServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	searchServiceSearchUserHandler := connect_go.NewUnaryHandler(
		SearchServiceSearchUserProcedure,
		svc.SearchUser,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	searchServiceSearchRepositoryHandler := connect_go.NewUnaryHandler(
		SearchServiceSearchRepositoryProcedure,
		svc.SearchRepository,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	searchServiceSearchLastCommitByContentHandler := connect_go.NewUnaryHandler(
		SearchServiceSearchLastCommitByContentProcedure,
		svc.SearchLastCommitByContent,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	searchServiceSearchCurationPluginHandler := connect_go.NewUnaryHandler(
		SearchServiceSearchCurationPluginProcedure,
		svc.SearchCurationPlugin,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	searchServiceSearchTagHandler := connect_go.NewUnaryHandler(
		SearchServiceSearchTagProcedure,
		svc.SearchTag,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	searchServiceSearchDraftHandler := connect_go.NewUnaryHandler(
		SearchServiceSearchDraftProcedure,
		svc.SearchDraft,
		connect_go.WithIdempotency(connect_go.IdempotencyNoSideEffects),
		connect_go.WithHandlerOptions(opts...),
	)
	return "/bufman.dubbo.apache.org.registry.v1alpha1.SearchService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SearchServiceSearchUserProcedure:
			searchServiceSearchUserHandler.ServeHTTP(w, r)
		case SearchServiceSearchRepositoryProcedure:
			searchServiceSearchRepositoryHandler.ServeHTTP(w, r)
		case SearchServiceSearchLastCommitByContentProcedure:
			searchServiceSearchLastCommitByContentHandler.ServeHTTP(w, r)
		case SearchServiceSearchCurationPluginProcedure:
			searchServiceSearchCurationPluginHandler.ServeHTTP(w, r)
		case SearchServiceSearchTagProcedure:
			searchServiceSearchTagHandler.ServeHTTP(w, r)
		case SearchServiceSearchDraftProcedure:
			searchServiceSearchDraftHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSearchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSearchServiceHandler struct{}

func (UnimplementedSearchServiceHandler) SearchUser(context.Context, *connect_go.Request[v1alpha1.SearchUserRequest]) (*connect_go.Response[v1alpha1.SearchUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchUser is not implemented"))
}

func (UnimplementedSearchServiceHandler) SearchRepository(context.Context, *connect_go.Request[v1alpha1.SearchRepositoryRequest]) (*connect_go.Response[v1alpha1.SearchRepositoryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchRepository is not implemented"))
}

func (UnimplementedSearchServiceHandler) SearchLastCommitByContent(context.Context, *connect_go.Request[v1alpha1.SearchLastCommitByContentRequest]) (*connect_go.Response[v1alpha1.SearchLastCommitByContentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchLastCommitByContent is not implemented"))
}

func (UnimplementedSearchServiceHandler) SearchCurationPlugin(context.Context, *connect_go.Request[v1alpha1.SearchCuratedPluginRequest]) (*connect_go.Response[v1alpha1.SearchCuratedPluginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchCurationPlugin is not implemented"))
}

func (UnimplementedSearchServiceHandler) SearchTag(context.Context, *connect_go.Request[v1alpha1.SearchTagRequest]) (*connect_go.Response[v1alpha1.SearchTagResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchTag is not implemented"))
}

func (UnimplementedSearchServiceHandler) SearchDraft(context.Context, *connect_go.Request[v1alpha1.SearchDraftRequest]) (*connect_go.Response[v1alpha1.SearchDraftResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("bufman.dubbo.apache.org.registry.v1alpha1.SearchService.SearchDraft is not implemented"))
}
