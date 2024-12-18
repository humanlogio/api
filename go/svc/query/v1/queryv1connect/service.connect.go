// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/query/v1/service.proto

package queryv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/humanlogio/api/go/svc/query/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// QueryServiceName is the fully-qualified name of the QueryService service.
	QueryServiceName = "svc.query.v1.QueryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QueryServiceSummarizeEventsProcedure is the fully-qualified name of the QueryService's
	// SummarizeEvents RPC.
	QueryServiceSummarizeEventsProcedure = "/svc.query.v1.QueryService/SummarizeEvents"
	// QueryServiceWatchQueryProcedure is the fully-qualified name of the QueryService's WatchQuery RPC.
	QueryServiceWatchQueryProcedure = "/svc.query.v1.QueryService/WatchQuery"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	queryServiceServiceDescriptor               = v1.File_svc_query_v1_service_proto.Services().ByName("QueryService")
	queryServiceSummarizeEventsMethodDescriptor = queryServiceServiceDescriptor.Methods().ByName("SummarizeEvents")
	queryServiceWatchQueryMethodDescriptor      = queryServiceServiceDescriptor.Methods().ByName("WatchQuery")
)

// QueryServiceClient is a client for the svc.query.v1.QueryService service.
type QueryServiceClient interface {
	SummarizeEvents(context.Context, *connect.Request[v1.SummarizeEventsRequest]) (*connect.Response[v1.SummarizeEventsResponse], error)
	WatchQuery(context.Context, *connect.Request[v1.WatchQueryRequest]) (*connect.ServerStreamForClient[v1.WatchQueryResponse], error)
}

// NewQueryServiceClient constructs a client for the svc.query.v1.QueryService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQueryServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) QueryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &queryServiceClient{
		summarizeEvents: connect.NewClient[v1.SummarizeEventsRequest, v1.SummarizeEventsResponse](
			httpClient,
			baseURL+QueryServiceSummarizeEventsProcedure,
			connect.WithSchema(queryServiceSummarizeEventsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		watchQuery: connect.NewClient[v1.WatchQueryRequest, v1.WatchQueryResponse](
			httpClient,
			baseURL+QueryServiceWatchQueryProcedure,
			connect.WithSchema(queryServiceWatchQueryMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// queryServiceClient implements QueryServiceClient.
type queryServiceClient struct {
	summarizeEvents *connect.Client[v1.SummarizeEventsRequest, v1.SummarizeEventsResponse]
	watchQuery      *connect.Client[v1.WatchQueryRequest, v1.WatchQueryResponse]
}

// SummarizeEvents calls svc.query.v1.QueryService.SummarizeEvents.
func (c *queryServiceClient) SummarizeEvents(ctx context.Context, req *connect.Request[v1.SummarizeEventsRequest]) (*connect.Response[v1.SummarizeEventsResponse], error) {
	return c.summarizeEvents.CallUnary(ctx, req)
}

// WatchQuery calls svc.query.v1.QueryService.WatchQuery.
func (c *queryServiceClient) WatchQuery(ctx context.Context, req *connect.Request[v1.WatchQueryRequest]) (*connect.ServerStreamForClient[v1.WatchQueryResponse], error) {
	return c.watchQuery.CallServerStream(ctx, req)
}

// QueryServiceHandler is an implementation of the svc.query.v1.QueryService service.
type QueryServiceHandler interface {
	SummarizeEvents(context.Context, *connect.Request[v1.SummarizeEventsRequest]) (*connect.Response[v1.SummarizeEventsResponse], error)
	WatchQuery(context.Context, *connect.Request[v1.WatchQueryRequest], *connect.ServerStream[v1.WatchQueryResponse]) error
}

// NewQueryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueryServiceHandler(svc QueryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queryServiceSummarizeEventsHandler := connect.NewUnaryHandler(
		QueryServiceSummarizeEventsProcedure,
		svc.SummarizeEvents,
		connect.WithSchema(queryServiceSummarizeEventsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceWatchQueryHandler := connect.NewServerStreamHandler(
		QueryServiceWatchQueryProcedure,
		svc.WatchQuery,
		connect.WithSchema(queryServiceWatchQueryMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/svc.query.v1.QueryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueryServiceSummarizeEventsProcedure:
			queryServiceSummarizeEventsHandler.ServeHTTP(w, r)
		case QueryServiceWatchQueryProcedure:
			queryServiceWatchQueryHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQueryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedQueryServiceHandler struct{}

func (UnimplementedQueryServiceHandler) SummarizeEvents(context.Context, *connect.Request[v1.SummarizeEventsRequest]) (*connect.Response[v1.SummarizeEventsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.query.v1.QueryService.SummarizeEvents is not implemented"))
}

func (UnimplementedQueryServiceHandler) WatchQuery(context.Context, *connect.Request[v1.WatchQueryRequest], *connect.ServerStream[v1.WatchQueryResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("svc.query.v1.QueryService.WatchQuery is not implemented"))
}
