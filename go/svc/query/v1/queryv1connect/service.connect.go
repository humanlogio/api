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
	// QueryServiceParseProcedure is the fully-qualified name of the QueryService's Parse RPC.
	QueryServiceParseProcedure = "/svc.query.v1.QueryService/Parse"
	// QueryServiceFormatProcedure is the fully-qualified name of the QueryService's Format RPC.
	QueryServiceFormatProcedure = "/svc.query.v1.QueryService/Format"
	// QueryServiceQueryProcedure is the fully-qualified name of the QueryService's Query RPC.
	QueryServiceQueryProcedure = "/svc.query.v1.QueryService/Query"
	// QueryServiceStreamProcedure is the fully-qualified name of the QueryService's Stream RPC.
	QueryServiceStreamProcedure = "/svc.query.v1.QueryService/Stream"
	// QueryServiceListSymbolsProcedure is the fully-qualified name of the QueryService's ListSymbols
	// RPC.
	QueryServiceListSymbolsProcedure = "/svc.query.v1.QueryService/ListSymbols"
)

// QueryServiceClient is a client for the svc.query.v1.QueryService service.
type QueryServiceClient interface {
	SummarizeEvents(context.Context, *connect.Request[v1.SummarizeEventsRequest]) (*connect.Response[v1.SummarizeEventsResponse], error)
	Parse(context.Context, *connect.Request[v1.ParseRequest]) (*connect.Response[v1.ParseResponse], error)
	Format(context.Context, *connect.Request[v1.FormatRequest]) (*connect.Response[v1.FormatResponse], error)
	Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error)
	Stream(context.Context, *connect.Request[v1.StreamRequest]) (*connect.ServerStreamForClient[v1.StreamResponse], error)
	ListSymbols(context.Context, *connect.Request[v1.ListSymbolsRequest]) (*connect.Response[v1.ListSymbolsResponse], error)
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
	queryServiceMethods := v1.File_svc_query_v1_service_proto.Services().ByName("QueryService").Methods()
	return &queryServiceClient{
		summarizeEvents: connect.NewClient[v1.SummarizeEventsRequest, v1.SummarizeEventsResponse](
			httpClient,
			baseURL+QueryServiceSummarizeEventsProcedure,
			connect.WithSchema(queryServiceMethods.ByName("SummarizeEvents")),
			connect.WithClientOptions(opts...),
		),
		parse: connect.NewClient[v1.ParseRequest, v1.ParseResponse](
			httpClient,
			baseURL+QueryServiceParseProcedure,
			connect.WithSchema(queryServiceMethods.ByName("Parse")),
			connect.WithClientOptions(opts...),
		),
		format: connect.NewClient[v1.FormatRequest, v1.FormatResponse](
			httpClient,
			baseURL+QueryServiceFormatProcedure,
			connect.WithSchema(queryServiceMethods.ByName("Format")),
			connect.WithClientOptions(opts...),
		),
		query: connect.NewClient[v1.QueryRequest, v1.QueryResponse](
			httpClient,
			baseURL+QueryServiceQueryProcedure,
			connect.WithSchema(queryServiceMethods.ByName("Query")),
			connect.WithClientOptions(opts...),
		),
		stream: connect.NewClient[v1.StreamRequest, v1.StreamResponse](
			httpClient,
			baseURL+QueryServiceStreamProcedure,
			connect.WithSchema(queryServiceMethods.ByName("Stream")),
			connect.WithClientOptions(opts...),
		),
		listSymbols: connect.NewClient[v1.ListSymbolsRequest, v1.ListSymbolsResponse](
			httpClient,
			baseURL+QueryServiceListSymbolsProcedure,
			connect.WithSchema(queryServiceMethods.ByName("ListSymbols")),
			connect.WithClientOptions(opts...),
		),
	}
}

// queryServiceClient implements QueryServiceClient.
type queryServiceClient struct {
	summarizeEvents *connect.Client[v1.SummarizeEventsRequest, v1.SummarizeEventsResponse]
	parse           *connect.Client[v1.ParseRequest, v1.ParseResponse]
	format          *connect.Client[v1.FormatRequest, v1.FormatResponse]
	query           *connect.Client[v1.QueryRequest, v1.QueryResponse]
	stream          *connect.Client[v1.StreamRequest, v1.StreamResponse]
	listSymbols     *connect.Client[v1.ListSymbolsRequest, v1.ListSymbolsResponse]
}

// SummarizeEvents calls svc.query.v1.QueryService.SummarizeEvents.
func (c *queryServiceClient) SummarizeEvents(ctx context.Context, req *connect.Request[v1.SummarizeEventsRequest]) (*connect.Response[v1.SummarizeEventsResponse], error) {
	return c.summarizeEvents.CallUnary(ctx, req)
}

// Parse calls svc.query.v1.QueryService.Parse.
func (c *queryServiceClient) Parse(ctx context.Context, req *connect.Request[v1.ParseRequest]) (*connect.Response[v1.ParseResponse], error) {
	return c.parse.CallUnary(ctx, req)
}

// Format calls svc.query.v1.QueryService.Format.
func (c *queryServiceClient) Format(ctx context.Context, req *connect.Request[v1.FormatRequest]) (*connect.Response[v1.FormatResponse], error) {
	return c.format.CallUnary(ctx, req)
}

// Query calls svc.query.v1.QueryService.Query.
func (c *queryServiceClient) Query(ctx context.Context, req *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	return c.query.CallUnary(ctx, req)
}

// Stream calls svc.query.v1.QueryService.Stream.
func (c *queryServiceClient) Stream(ctx context.Context, req *connect.Request[v1.StreamRequest]) (*connect.ServerStreamForClient[v1.StreamResponse], error) {
	return c.stream.CallServerStream(ctx, req)
}

// ListSymbols calls svc.query.v1.QueryService.ListSymbols.
func (c *queryServiceClient) ListSymbols(ctx context.Context, req *connect.Request[v1.ListSymbolsRequest]) (*connect.Response[v1.ListSymbolsResponse], error) {
	return c.listSymbols.CallUnary(ctx, req)
}

// QueryServiceHandler is an implementation of the svc.query.v1.QueryService service.
type QueryServiceHandler interface {
	SummarizeEvents(context.Context, *connect.Request[v1.SummarizeEventsRequest]) (*connect.Response[v1.SummarizeEventsResponse], error)
	Parse(context.Context, *connect.Request[v1.ParseRequest]) (*connect.Response[v1.ParseResponse], error)
	Format(context.Context, *connect.Request[v1.FormatRequest]) (*connect.Response[v1.FormatResponse], error)
	Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error)
	Stream(context.Context, *connect.Request[v1.StreamRequest], *connect.ServerStream[v1.StreamResponse]) error
	ListSymbols(context.Context, *connect.Request[v1.ListSymbolsRequest]) (*connect.Response[v1.ListSymbolsResponse], error)
}

// NewQueryServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueryServiceHandler(svc QueryServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queryServiceMethods := v1.File_svc_query_v1_service_proto.Services().ByName("QueryService").Methods()
	queryServiceSummarizeEventsHandler := connect.NewUnaryHandler(
		QueryServiceSummarizeEventsProcedure,
		svc.SummarizeEvents,
		connect.WithSchema(queryServiceMethods.ByName("SummarizeEvents")),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceParseHandler := connect.NewUnaryHandler(
		QueryServiceParseProcedure,
		svc.Parse,
		connect.WithSchema(queryServiceMethods.ByName("Parse")),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceFormatHandler := connect.NewUnaryHandler(
		QueryServiceFormatProcedure,
		svc.Format,
		connect.WithSchema(queryServiceMethods.ByName("Format")),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceQueryHandler := connect.NewUnaryHandler(
		QueryServiceQueryProcedure,
		svc.Query,
		connect.WithSchema(queryServiceMethods.ByName("Query")),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceStreamHandler := connect.NewServerStreamHandler(
		QueryServiceStreamProcedure,
		svc.Stream,
		connect.WithSchema(queryServiceMethods.ByName("Stream")),
		connect.WithHandlerOptions(opts...),
	)
	queryServiceListSymbolsHandler := connect.NewUnaryHandler(
		QueryServiceListSymbolsProcedure,
		svc.ListSymbols,
		connect.WithSchema(queryServiceMethods.ByName("ListSymbols")),
		connect.WithHandlerOptions(opts...),
	)
	return "/svc.query.v1.QueryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueryServiceSummarizeEventsProcedure:
			queryServiceSummarizeEventsHandler.ServeHTTP(w, r)
		case QueryServiceParseProcedure:
			queryServiceParseHandler.ServeHTTP(w, r)
		case QueryServiceFormatProcedure:
			queryServiceFormatHandler.ServeHTTP(w, r)
		case QueryServiceQueryProcedure:
			queryServiceQueryHandler.ServeHTTP(w, r)
		case QueryServiceStreamProcedure:
			queryServiceStreamHandler.ServeHTTP(w, r)
		case QueryServiceListSymbolsProcedure:
			queryServiceListSymbolsHandler.ServeHTTP(w, r)
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

func (UnimplementedQueryServiceHandler) Parse(context.Context, *connect.Request[v1.ParseRequest]) (*connect.Response[v1.ParseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.query.v1.QueryService.Parse is not implemented"))
}

func (UnimplementedQueryServiceHandler) Format(context.Context, *connect.Request[v1.FormatRequest]) (*connect.Response[v1.FormatResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.query.v1.QueryService.Format is not implemented"))
}

func (UnimplementedQueryServiceHandler) Query(context.Context, *connect.Request[v1.QueryRequest]) (*connect.Response[v1.QueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.query.v1.QueryService.Query is not implemented"))
}

func (UnimplementedQueryServiceHandler) Stream(context.Context, *connect.Request[v1.StreamRequest], *connect.ServerStream[v1.StreamResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("svc.query.v1.QueryService.Stream is not implemented"))
}

func (UnimplementedQueryServiceHandler) ListSymbols(context.Context, *connect.Request[v1.ListSymbolsRequest]) (*connect.Response[v1.ListSymbolsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.query.v1.QueryService.ListSymbols is not implemented"))
}
