// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/user/v1/service.proto

package userv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/humanlogio/api/go/svc/user/v1"
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
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "svc.user.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceWhoamiProcedure is the fully-qualified name of the UserService's Whoami RPC.
	UserServiceWhoamiProcedure = "/svc.user.v1.UserService/Whoami"
	// UserServiceGetLogoutURLProcedure is the fully-qualified name of the UserService's GetLogoutURL
	// RPC.
	UserServiceGetLogoutURLProcedure = "/svc.user.v1.UserService/GetLogoutURL"
	// UserServiceRefreshUserTokenProcedure is the fully-qualified name of the UserService's
	// RefreshUserToken RPC.
	UserServiceRefreshUserTokenProcedure = "/svc.user.v1.UserService/RefreshUserToken"
	// UserServiceCreateOrganizationProcedure is the fully-qualified name of the UserService's
	// CreateOrganization RPC.
	UserServiceCreateOrganizationProcedure = "/svc.user.v1.UserService/CreateOrganization"
	// UserServiceListOrganizationProcedure is the fully-qualified name of the UserService's
	// ListOrganization RPC.
	UserServiceListOrganizationProcedure = "/svc.user.v1.UserService/ListOrganization"
)

// UserServiceClient is a client for the svc.user.v1.UserService service.
type UserServiceClient interface {
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
	GetLogoutURL(context.Context, *connect.Request[v1.GetLogoutURLRequest]) (*connect.Response[v1.GetLogoutURLResponse], error)
	RefreshUserToken(context.Context, *connect.Request[v1.RefreshUserTokenRequest]) (*connect.Response[v1.RefreshUserTokenResponse], error)
	CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error)
	ListOrganization(context.Context, *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error)
}

// NewUserServiceClient constructs a client for the svc.user.v1.UserService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	userServiceMethods := v1.File_svc_user_v1_service_proto.Services().ByName("UserService").Methods()
	return &userServiceClient{
		whoami: connect.NewClient[v1.WhoamiRequest, v1.WhoamiResponse](
			httpClient,
			baseURL+UserServiceWhoamiProcedure,
			connect.WithSchema(userServiceMethods.ByName("Whoami")),
			connect.WithClientOptions(opts...),
		),
		getLogoutURL: connect.NewClient[v1.GetLogoutURLRequest, v1.GetLogoutURLResponse](
			httpClient,
			baseURL+UserServiceGetLogoutURLProcedure,
			connect.WithSchema(userServiceMethods.ByName("GetLogoutURL")),
			connect.WithClientOptions(opts...),
		),
		refreshUserToken: connect.NewClient[v1.RefreshUserTokenRequest, v1.RefreshUserTokenResponse](
			httpClient,
			baseURL+UserServiceRefreshUserTokenProcedure,
			connect.WithSchema(userServiceMethods.ByName("RefreshUserToken")),
			connect.WithClientOptions(opts...),
		),
		createOrganization: connect.NewClient[v1.CreateOrganizationRequest, v1.CreateOrganizationResponse](
			httpClient,
			baseURL+UserServiceCreateOrganizationProcedure,
			connect.WithSchema(userServiceMethods.ByName("CreateOrganization")),
			connect.WithClientOptions(opts...),
		),
		listOrganization: connect.NewClient[v1.ListOrganizationRequest, v1.ListOrganizationResponse](
			httpClient,
			baseURL+UserServiceListOrganizationProcedure,
			connect.WithSchema(userServiceMethods.ByName("ListOrganization")),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	whoami             *connect.Client[v1.WhoamiRequest, v1.WhoamiResponse]
	getLogoutURL       *connect.Client[v1.GetLogoutURLRequest, v1.GetLogoutURLResponse]
	refreshUserToken   *connect.Client[v1.RefreshUserTokenRequest, v1.RefreshUserTokenResponse]
	createOrganization *connect.Client[v1.CreateOrganizationRequest, v1.CreateOrganizationResponse]
	listOrganization   *connect.Client[v1.ListOrganizationRequest, v1.ListOrganizationResponse]
}

// Whoami calls svc.user.v1.UserService.Whoami.
func (c *userServiceClient) Whoami(ctx context.Context, req *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error) {
	return c.whoami.CallUnary(ctx, req)
}

// GetLogoutURL calls svc.user.v1.UserService.GetLogoutURL.
func (c *userServiceClient) GetLogoutURL(ctx context.Context, req *connect.Request[v1.GetLogoutURLRequest]) (*connect.Response[v1.GetLogoutURLResponse], error) {
	return c.getLogoutURL.CallUnary(ctx, req)
}

// RefreshUserToken calls svc.user.v1.UserService.RefreshUserToken.
func (c *userServiceClient) RefreshUserToken(ctx context.Context, req *connect.Request[v1.RefreshUserTokenRequest]) (*connect.Response[v1.RefreshUserTokenResponse], error) {
	return c.refreshUserToken.CallUnary(ctx, req)
}

// CreateOrganization calls svc.user.v1.UserService.CreateOrganization.
func (c *userServiceClient) CreateOrganization(ctx context.Context, req *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error) {
	return c.createOrganization.CallUnary(ctx, req)
}

// ListOrganization calls svc.user.v1.UserService.ListOrganization.
func (c *userServiceClient) ListOrganization(ctx context.Context, req *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error) {
	return c.listOrganization.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the svc.user.v1.UserService service.
type UserServiceHandler interface {
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
	GetLogoutURL(context.Context, *connect.Request[v1.GetLogoutURLRequest]) (*connect.Response[v1.GetLogoutURLResponse], error)
	RefreshUserToken(context.Context, *connect.Request[v1.RefreshUserTokenRequest]) (*connect.Response[v1.RefreshUserTokenResponse], error)
	CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error)
	ListOrganization(context.Context, *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceMethods := v1.File_svc_user_v1_service_proto.Services().ByName("UserService").Methods()
	userServiceWhoamiHandler := connect.NewUnaryHandler(
		UserServiceWhoamiProcedure,
		svc.Whoami,
		connect.WithSchema(userServiceMethods.ByName("Whoami")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetLogoutURLHandler := connect.NewUnaryHandler(
		UserServiceGetLogoutURLProcedure,
		svc.GetLogoutURL,
		connect.WithSchema(userServiceMethods.ByName("GetLogoutURL")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceRefreshUserTokenHandler := connect.NewUnaryHandler(
		UserServiceRefreshUserTokenProcedure,
		svc.RefreshUserToken,
		connect.WithSchema(userServiceMethods.ByName("RefreshUserToken")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateOrganizationHandler := connect.NewUnaryHandler(
		UserServiceCreateOrganizationProcedure,
		svc.CreateOrganization,
		connect.WithSchema(userServiceMethods.ByName("CreateOrganization")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceListOrganizationHandler := connect.NewUnaryHandler(
		UserServiceListOrganizationProcedure,
		svc.ListOrganization,
		connect.WithSchema(userServiceMethods.ByName("ListOrganization")),
		connect.WithHandlerOptions(opts...),
	)
	return "/svc.user.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceWhoamiProcedure:
			userServiceWhoamiHandler.ServeHTTP(w, r)
		case UserServiceGetLogoutURLProcedure:
			userServiceGetLogoutURLHandler.ServeHTTP(w, r)
		case UserServiceRefreshUserTokenProcedure:
			userServiceRefreshUserTokenHandler.ServeHTTP(w, r)
		case UserServiceCreateOrganizationProcedure:
			userServiceCreateOrganizationHandler.ServeHTTP(w, r)
		case UserServiceListOrganizationProcedure:
			userServiceListOrganizationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.Whoami is not implemented"))
}

func (UnimplementedUserServiceHandler) GetLogoutURL(context.Context, *connect.Request[v1.GetLogoutURLRequest]) (*connect.Response[v1.GetLogoutURLResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.GetLogoutURL is not implemented"))
}

func (UnimplementedUserServiceHandler) RefreshUserToken(context.Context, *connect.Request[v1.RefreshUserTokenRequest]) (*connect.Response[v1.RefreshUserTokenResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.RefreshUserToken is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.CreateOrganization is not implemented"))
}

func (UnimplementedUserServiceHandler) ListOrganization(context.Context, *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.ListOrganization is not implemented"))
}
