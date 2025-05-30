// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: svc/user/v1/service_private.proto

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
	// UserServiceUpdateUserProcedure is the fully-qualified name of the UserService's UpdateUser RPC.
	UserServiceUpdateUserProcedure = "/svc.user.v1.UserService/UpdateUser"
	// UserServiceCreateOrganizationProcedure is the fully-qualified name of the UserService's
	// CreateOrganization RPC.
	UserServiceCreateOrganizationProcedure = "/svc.user.v1.UserService/CreateOrganization"
	// UserServiceListOrganizationProcedure is the fully-qualified name of the UserService's
	// ListOrganization RPC.
	UserServiceListOrganizationProcedure = "/svc.user.v1.UserService/ListOrganization"
	// UserServiceSaveLocalhostConfigProcedure is the fully-qualified name of the UserService's
	// SaveLocalhostConfig RPC.
	UserServiceSaveLocalhostConfigProcedure = "/svc.user.v1.UserService/SaveLocalhostConfig"
	// UserServiceGetLocalhostConfigProcedure is the fully-qualified name of the UserService's
	// GetLocalhostConfig RPC.
	UserServiceGetLocalhostConfigProcedure = "/svc.user.v1.UserService/GetLocalhostConfig"
	// UserServiceRecordQueryHistoryProcedure is the fully-qualified name of the UserService's
	// RecordQueryHistory RPC.
	UserServiceRecordQueryHistoryProcedure = "/svc.user.v1.UserService/RecordQueryHistory"
	// UserServiceGetQueryHistoryProcedure is the fully-qualified name of the UserService's
	// GetQueryHistory RPC.
	UserServiceGetQueryHistoryProcedure = "/svc.user.v1.UserService/GetQueryHistory"
	// UserServiceListQueryHistoryProcedure is the fully-qualified name of the UserService's
	// ListQueryHistory RPC.
	UserServiceListQueryHistoryProcedure = "/svc.user.v1.UserService/ListQueryHistory"
	// UserServiceDeleteQueryHistoryProcedure is the fully-qualified name of the UserService's
	// DeleteQueryHistory RPC.
	UserServiceDeleteQueryHistoryProcedure = "/svc.user.v1.UserService/DeleteQueryHistory"
	// UserServiceCreateFavoriteQueryProcedure is the fully-qualified name of the UserService's
	// CreateFavoriteQuery RPC.
	UserServiceCreateFavoriteQueryProcedure = "/svc.user.v1.UserService/CreateFavoriteQuery"
	// UserServiceGetFavoriteQueryProcedure is the fully-qualified name of the UserService's
	// GetFavoriteQuery RPC.
	UserServiceGetFavoriteQueryProcedure = "/svc.user.v1.UserService/GetFavoriteQuery"
	// UserServiceUpdateFavoriteQueryProcedure is the fully-qualified name of the UserService's
	// UpdateFavoriteQuery RPC.
	UserServiceUpdateFavoriteQueryProcedure = "/svc.user.v1.UserService/UpdateFavoriteQuery"
	// UserServiceListFavoriteQueryProcedure is the fully-qualified name of the UserService's
	// ListFavoriteQuery RPC.
	UserServiceListFavoriteQueryProcedure = "/svc.user.v1.UserService/ListFavoriteQuery"
	// UserServiceDeleteFavoriteQueryProcedure is the fully-qualified name of the UserService's
	// DeleteFavoriteQuery RPC.
	UserServiceDeleteFavoriteQueryProcedure = "/svc.user.v1.UserService/DeleteFavoriteQuery"
)

// UserServiceClient is a client for the svc.user.v1.UserService service.
type UserServiceClient interface {
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
	GetLogoutURL(context.Context, *connect.Request[v1.GetLogoutURLRequest]) (*connect.Response[v1.GetLogoutURLResponse], error)
	RefreshUserToken(context.Context, *connect.Request[v1.RefreshUserTokenRequest]) (*connect.Response[v1.RefreshUserTokenResponse], error)
	UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error)
	CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error)
	ListOrganization(context.Context, *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error)
	SaveLocalhostConfig(context.Context, *connect.Request[v1.SaveLocalhostConfigRequest]) (*connect.Response[v1.SaveLocalhostConfigResponse], error)
	GetLocalhostConfig(context.Context, *connect.Request[v1.GetLocalhostConfigRequest]) (*connect.Response[v1.GetLocalhostConfigResponse], error)
	RecordQueryHistory(context.Context, *connect.Request[v1.RecordQueryHistoryRequest]) (*connect.Response[v1.RecordQueryHistoryResponse], error)
	GetQueryHistory(context.Context, *connect.Request[v1.GetQueryHistoryRequest]) (*connect.Response[v1.GetQueryHistoryResponse], error)
	ListQueryHistory(context.Context, *connect.Request[v1.ListQueryHistoryRequest]) (*connect.Response[v1.ListQueryHistoryResponse], error)
	DeleteQueryHistory(context.Context, *connect.Request[v1.DeleteQueryHistoryRequest]) (*connect.Response[v1.DeleteQueryHistoryResponse], error)
	CreateFavoriteQuery(context.Context, *connect.Request[v1.CreateFavoriteQueryRequest]) (*connect.Response[v1.CreateFavoriteQueryResponse], error)
	GetFavoriteQuery(context.Context, *connect.Request[v1.GetFavoriteQueryRequest]) (*connect.Response[v1.GetFavoriteQueryResponse], error)
	UpdateFavoriteQuery(context.Context, *connect.Request[v1.UpdateFavoriteQueryRequest]) (*connect.Response[v1.UpdateFavoriteQueryResponse], error)
	ListFavoriteQuery(context.Context, *connect.Request[v1.ListFavoriteQueryRequest]) (*connect.Response[v1.ListFavoriteQueryResponse], error)
	DeleteFavoriteQuery(context.Context, *connect.Request[v1.DeleteFavoriteQueryRequest]) (*connect.Response[v1.DeleteFavoriteQueryResponse], error)
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
	userServiceMethods := v1.File_svc_user_v1_service_private_proto.Services().ByName("UserService").Methods()
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
		updateUser: connect.NewClient[v1.UpdateUserRequest, v1.UpdateUserResponse](
			httpClient,
			baseURL+UserServiceUpdateUserProcedure,
			connect.WithSchema(userServiceMethods.ByName("UpdateUser")),
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
		saveLocalhostConfig: connect.NewClient[v1.SaveLocalhostConfigRequest, v1.SaveLocalhostConfigResponse](
			httpClient,
			baseURL+UserServiceSaveLocalhostConfigProcedure,
			connect.WithSchema(userServiceMethods.ByName("SaveLocalhostConfig")),
			connect.WithClientOptions(opts...),
		),
		getLocalhostConfig: connect.NewClient[v1.GetLocalhostConfigRequest, v1.GetLocalhostConfigResponse](
			httpClient,
			baseURL+UserServiceGetLocalhostConfigProcedure,
			connect.WithSchema(userServiceMethods.ByName("GetLocalhostConfig")),
			connect.WithClientOptions(opts...),
		),
		recordQueryHistory: connect.NewClient[v1.RecordQueryHistoryRequest, v1.RecordQueryHistoryResponse](
			httpClient,
			baseURL+UserServiceRecordQueryHistoryProcedure,
			connect.WithSchema(userServiceMethods.ByName("RecordQueryHistory")),
			connect.WithClientOptions(opts...),
		),
		getQueryHistory: connect.NewClient[v1.GetQueryHistoryRequest, v1.GetQueryHistoryResponse](
			httpClient,
			baseURL+UserServiceGetQueryHistoryProcedure,
			connect.WithSchema(userServiceMethods.ByName("GetQueryHistory")),
			connect.WithClientOptions(opts...),
		),
		listQueryHistory: connect.NewClient[v1.ListQueryHistoryRequest, v1.ListQueryHistoryResponse](
			httpClient,
			baseURL+UserServiceListQueryHistoryProcedure,
			connect.WithSchema(userServiceMethods.ByName("ListQueryHistory")),
			connect.WithClientOptions(opts...),
		),
		deleteQueryHistory: connect.NewClient[v1.DeleteQueryHistoryRequest, v1.DeleteQueryHistoryResponse](
			httpClient,
			baseURL+UserServiceDeleteQueryHistoryProcedure,
			connect.WithSchema(userServiceMethods.ByName("DeleteQueryHistory")),
			connect.WithClientOptions(opts...),
		),
		createFavoriteQuery: connect.NewClient[v1.CreateFavoriteQueryRequest, v1.CreateFavoriteQueryResponse](
			httpClient,
			baseURL+UserServiceCreateFavoriteQueryProcedure,
			connect.WithSchema(userServiceMethods.ByName("CreateFavoriteQuery")),
			connect.WithClientOptions(opts...),
		),
		getFavoriteQuery: connect.NewClient[v1.GetFavoriteQueryRequest, v1.GetFavoriteQueryResponse](
			httpClient,
			baseURL+UserServiceGetFavoriteQueryProcedure,
			connect.WithSchema(userServiceMethods.ByName("GetFavoriteQuery")),
			connect.WithClientOptions(opts...),
		),
		updateFavoriteQuery: connect.NewClient[v1.UpdateFavoriteQueryRequest, v1.UpdateFavoriteQueryResponse](
			httpClient,
			baseURL+UserServiceUpdateFavoriteQueryProcedure,
			connect.WithSchema(userServiceMethods.ByName("UpdateFavoriteQuery")),
			connect.WithClientOptions(opts...),
		),
		listFavoriteQuery: connect.NewClient[v1.ListFavoriteQueryRequest, v1.ListFavoriteQueryResponse](
			httpClient,
			baseURL+UserServiceListFavoriteQueryProcedure,
			connect.WithSchema(userServiceMethods.ByName("ListFavoriteQuery")),
			connect.WithClientOptions(opts...),
		),
		deleteFavoriteQuery: connect.NewClient[v1.DeleteFavoriteQueryRequest, v1.DeleteFavoriteQueryResponse](
			httpClient,
			baseURL+UserServiceDeleteFavoriteQueryProcedure,
			connect.WithSchema(userServiceMethods.ByName("DeleteFavoriteQuery")),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	whoami              *connect.Client[v1.WhoamiRequest, v1.WhoamiResponse]
	getLogoutURL        *connect.Client[v1.GetLogoutURLRequest, v1.GetLogoutURLResponse]
	refreshUserToken    *connect.Client[v1.RefreshUserTokenRequest, v1.RefreshUserTokenResponse]
	updateUser          *connect.Client[v1.UpdateUserRequest, v1.UpdateUserResponse]
	createOrganization  *connect.Client[v1.CreateOrganizationRequest, v1.CreateOrganizationResponse]
	listOrganization    *connect.Client[v1.ListOrganizationRequest, v1.ListOrganizationResponse]
	saveLocalhostConfig *connect.Client[v1.SaveLocalhostConfigRequest, v1.SaveLocalhostConfigResponse]
	getLocalhostConfig  *connect.Client[v1.GetLocalhostConfigRequest, v1.GetLocalhostConfigResponse]
	recordQueryHistory  *connect.Client[v1.RecordQueryHistoryRequest, v1.RecordQueryHistoryResponse]
	getQueryHistory     *connect.Client[v1.GetQueryHistoryRequest, v1.GetQueryHistoryResponse]
	listQueryHistory    *connect.Client[v1.ListQueryHistoryRequest, v1.ListQueryHistoryResponse]
	deleteQueryHistory  *connect.Client[v1.DeleteQueryHistoryRequest, v1.DeleteQueryHistoryResponse]
	createFavoriteQuery *connect.Client[v1.CreateFavoriteQueryRequest, v1.CreateFavoriteQueryResponse]
	getFavoriteQuery    *connect.Client[v1.GetFavoriteQueryRequest, v1.GetFavoriteQueryResponse]
	updateFavoriteQuery *connect.Client[v1.UpdateFavoriteQueryRequest, v1.UpdateFavoriteQueryResponse]
	listFavoriteQuery   *connect.Client[v1.ListFavoriteQueryRequest, v1.ListFavoriteQueryResponse]
	deleteFavoriteQuery *connect.Client[v1.DeleteFavoriteQueryRequest, v1.DeleteFavoriteQueryResponse]
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

// UpdateUser calls svc.user.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// CreateOrganization calls svc.user.v1.UserService.CreateOrganization.
func (c *userServiceClient) CreateOrganization(ctx context.Context, req *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error) {
	return c.createOrganization.CallUnary(ctx, req)
}

// ListOrganization calls svc.user.v1.UserService.ListOrganization.
func (c *userServiceClient) ListOrganization(ctx context.Context, req *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error) {
	return c.listOrganization.CallUnary(ctx, req)
}

// SaveLocalhostConfig calls svc.user.v1.UserService.SaveLocalhostConfig.
func (c *userServiceClient) SaveLocalhostConfig(ctx context.Context, req *connect.Request[v1.SaveLocalhostConfigRequest]) (*connect.Response[v1.SaveLocalhostConfigResponse], error) {
	return c.saveLocalhostConfig.CallUnary(ctx, req)
}

// GetLocalhostConfig calls svc.user.v1.UserService.GetLocalhostConfig.
func (c *userServiceClient) GetLocalhostConfig(ctx context.Context, req *connect.Request[v1.GetLocalhostConfigRequest]) (*connect.Response[v1.GetLocalhostConfigResponse], error) {
	return c.getLocalhostConfig.CallUnary(ctx, req)
}

// RecordQueryHistory calls svc.user.v1.UserService.RecordQueryHistory.
func (c *userServiceClient) RecordQueryHistory(ctx context.Context, req *connect.Request[v1.RecordQueryHistoryRequest]) (*connect.Response[v1.RecordQueryHistoryResponse], error) {
	return c.recordQueryHistory.CallUnary(ctx, req)
}

// GetQueryHistory calls svc.user.v1.UserService.GetQueryHistory.
func (c *userServiceClient) GetQueryHistory(ctx context.Context, req *connect.Request[v1.GetQueryHistoryRequest]) (*connect.Response[v1.GetQueryHistoryResponse], error) {
	return c.getQueryHistory.CallUnary(ctx, req)
}

// ListQueryHistory calls svc.user.v1.UserService.ListQueryHistory.
func (c *userServiceClient) ListQueryHistory(ctx context.Context, req *connect.Request[v1.ListQueryHistoryRequest]) (*connect.Response[v1.ListQueryHistoryResponse], error) {
	return c.listQueryHistory.CallUnary(ctx, req)
}

// DeleteQueryHistory calls svc.user.v1.UserService.DeleteQueryHistory.
func (c *userServiceClient) DeleteQueryHistory(ctx context.Context, req *connect.Request[v1.DeleteQueryHistoryRequest]) (*connect.Response[v1.DeleteQueryHistoryResponse], error) {
	return c.deleteQueryHistory.CallUnary(ctx, req)
}

// CreateFavoriteQuery calls svc.user.v1.UserService.CreateFavoriteQuery.
func (c *userServiceClient) CreateFavoriteQuery(ctx context.Context, req *connect.Request[v1.CreateFavoriteQueryRequest]) (*connect.Response[v1.CreateFavoriteQueryResponse], error) {
	return c.createFavoriteQuery.CallUnary(ctx, req)
}

// GetFavoriteQuery calls svc.user.v1.UserService.GetFavoriteQuery.
func (c *userServiceClient) GetFavoriteQuery(ctx context.Context, req *connect.Request[v1.GetFavoriteQueryRequest]) (*connect.Response[v1.GetFavoriteQueryResponse], error) {
	return c.getFavoriteQuery.CallUnary(ctx, req)
}

// UpdateFavoriteQuery calls svc.user.v1.UserService.UpdateFavoriteQuery.
func (c *userServiceClient) UpdateFavoriteQuery(ctx context.Context, req *connect.Request[v1.UpdateFavoriteQueryRequest]) (*connect.Response[v1.UpdateFavoriteQueryResponse], error) {
	return c.updateFavoriteQuery.CallUnary(ctx, req)
}

// ListFavoriteQuery calls svc.user.v1.UserService.ListFavoriteQuery.
func (c *userServiceClient) ListFavoriteQuery(ctx context.Context, req *connect.Request[v1.ListFavoriteQueryRequest]) (*connect.Response[v1.ListFavoriteQueryResponse], error) {
	return c.listFavoriteQuery.CallUnary(ctx, req)
}

// DeleteFavoriteQuery calls svc.user.v1.UserService.DeleteFavoriteQuery.
func (c *userServiceClient) DeleteFavoriteQuery(ctx context.Context, req *connect.Request[v1.DeleteFavoriteQueryRequest]) (*connect.Response[v1.DeleteFavoriteQueryResponse], error) {
	return c.deleteFavoriteQuery.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the svc.user.v1.UserService service.
type UserServiceHandler interface {
	Whoami(context.Context, *connect.Request[v1.WhoamiRequest]) (*connect.Response[v1.WhoamiResponse], error)
	GetLogoutURL(context.Context, *connect.Request[v1.GetLogoutURLRequest]) (*connect.Response[v1.GetLogoutURLResponse], error)
	RefreshUserToken(context.Context, *connect.Request[v1.RefreshUserTokenRequest]) (*connect.Response[v1.RefreshUserTokenResponse], error)
	UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error)
	CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error)
	ListOrganization(context.Context, *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error)
	SaveLocalhostConfig(context.Context, *connect.Request[v1.SaveLocalhostConfigRequest]) (*connect.Response[v1.SaveLocalhostConfigResponse], error)
	GetLocalhostConfig(context.Context, *connect.Request[v1.GetLocalhostConfigRequest]) (*connect.Response[v1.GetLocalhostConfigResponse], error)
	RecordQueryHistory(context.Context, *connect.Request[v1.RecordQueryHistoryRequest]) (*connect.Response[v1.RecordQueryHistoryResponse], error)
	GetQueryHistory(context.Context, *connect.Request[v1.GetQueryHistoryRequest]) (*connect.Response[v1.GetQueryHistoryResponse], error)
	ListQueryHistory(context.Context, *connect.Request[v1.ListQueryHistoryRequest]) (*connect.Response[v1.ListQueryHistoryResponse], error)
	DeleteQueryHistory(context.Context, *connect.Request[v1.DeleteQueryHistoryRequest]) (*connect.Response[v1.DeleteQueryHistoryResponse], error)
	CreateFavoriteQuery(context.Context, *connect.Request[v1.CreateFavoriteQueryRequest]) (*connect.Response[v1.CreateFavoriteQueryResponse], error)
	GetFavoriteQuery(context.Context, *connect.Request[v1.GetFavoriteQueryRequest]) (*connect.Response[v1.GetFavoriteQueryResponse], error)
	UpdateFavoriteQuery(context.Context, *connect.Request[v1.UpdateFavoriteQueryRequest]) (*connect.Response[v1.UpdateFavoriteQueryResponse], error)
	ListFavoriteQuery(context.Context, *connect.Request[v1.ListFavoriteQueryRequest]) (*connect.Response[v1.ListFavoriteQueryResponse], error)
	DeleteFavoriteQuery(context.Context, *connect.Request[v1.DeleteFavoriteQueryRequest]) (*connect.Response[v1.DeleteFavoriteQueryResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceMethods := v1.File_svc_user_v1_service_private_proto.Services().ByName("UserService").Methods()
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
	userServiceUpdateUserHandler := connect.NewUnaryHandler(
		UserServiceUpdateUserProcedure,
		svc.UpdateUser,
		connect.WithSchema(userServiceMethods.ByName("UpdateUser")),
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
	userServiceSaveLocalhostConfigHandler := connect.NewUnaryHandler(
		UserServiceSaveLocalhostConfigProcedure,
		svc.SaveLocalhostConfig,
		connect.WithSchema(userServiceMethods.ByName("SaveLocalhostConfig")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetLocalhostConfigHandler := connect.NewUnaryHandler(
		UserServiceGetLocalhostConfigProcedure,
		svc.GetLocalhostConfig,
		connect.WithSchema(userServiceMethods.ByName("GetLocalhostConfig")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceRecordQueryHistoryHandler := connect.NewUnaryHandler(
		UserServiceRecordQueryHistoryProcedure,
		svc.RecordQueryHistory,
		connect.WithSchema(userServiceMethods.ByName("RecordQueryHistory")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetQueryHistoryHandler := connect.NewUnaryHandler(
		UserServiceGetQueryHistoryProcedure,
		svc.GetQueryHistory,
		connect.WithSchema(userServiceMethods.ByName("GetQueryHistory")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceListQueryHistoryHandler := connect.NewUnaryHandler(
		UserServiceListQueryHistoryProcedure,
		svc.ListQueryHistory,
		connect.WithSchema(userServiceMethods.ByName("ListQueryHistory")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteQueryHistoryHandler := connect.NewUnaryHandler(
		UserServiceDeleteQueryHistoryProcedure,
		svc.DeleteQueryHistory,
		connect.WithSchema(userServiceMethods.ByName("DeleteQueryHistory")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateFavoriteQueryHandler := connect.NewUnaryHandler(
		UserServiceCreateFavoriteQueryProcedure,
		svc.CreateFavoriteQuery,
		connect.WithSchema(userServiceMethods.ByName("CreateFavoriteQuery")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetFavoriteQueryHandler := connect.NewUnaryHandler(
		UserServiceGetFavoriteQueryProcedure,
		svc.GetFavoriteQuery,
		connect.WithSchema(userServiceMethods.ByName("GetFavoriteQuery")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceUpdateFavoriteQueryHandler := connect.NewUnaryHandler(
		UserServiceUpdateFavoriteQueryProcedure,
		svc.UpdateFavoriteQuery,
		connect.WithSchema(userServiceMethods.ByName("UpdateFavoriteQuery")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceListFavoriteQueryHandler := connect.NewUnaryHandler(
		UserServiceListFavoriteQueryProcedure,
		svc.ListFavoriteQuery,
		connect.WithSchema(userServiceMethods.ByName("ListFavoriteQuery")),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteFavoriteQueryHandler := connect.NewUnaryHandler(
		UserServiceDeleteFavoriteQueryProcedure,
		svc.DeleteFavoriteQuery,
		connect.WithSchema(userServiceMethods.ByName("DeleteFavoriteQuery")),
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
		case UserServiceUpdateUserProcedure:
			userServiceUpdateUserHandler.ServeHTTP(w, r)
		case UserServiceCreateOrganizationProcedure:
			userServiceCreateOrganizationHandler.ServeHTTP(w, r)
		case UserServiceListOrganizationProcedure:
			userServiceListOrganizationHandler.ServeHTTP(w, r)
		case UserServiceSaveLocalhostConfigProcedure:
			userServiceSaveLocalhostConfigHandler.ServeHTTP(w, r)
		case UserServiceGetLocalhostConfigProcedure:
			userServiceGetLocalhostConfigHandler.ServeHTTP(w, r)
		case UserServiceRecordQueryHistoryProcedure:
			userServiceRecordQueryHistoryHandler.ServeHTTP(w, r)
		case UserServiceGetQueryHistoryProcedure:
			userServiceGetQueryHistoryHandler.ServeHTTP(w, r)
		case UserServiceListQueryHistoryProcedure:
			userServiceListQueryHistoryHandler.ServeHTTP(w, r)
		case UserServiceDeleteQueryHistoryProcedure:
			userServiceDeleteQueryHistoryHandler.ServeHTTP(w, r)
		case UserServiceCreateFavoriteQueryProcedure:
			userServiceCreateFavoriteQueryHandler.ServeHTTP(w, r)
		case UserServiceGetFavoriteQueryProcedure:
			userServiceGetFavoriteQueryHandler.ServeHTTP(w, r)
		case UserServiceUpdateFavoriteQueryProcedure:
			userServiceUpdateFavoriteQueryHandler.ServeHTTP(w, r)
		case UserServiceListFavoriteQueryProcedure:
			userServiceListFavoriteQueryHandler.ServeHTTP(w, r)
		case UserServiceDeleteFavoriteQueryProcedure:
			userServiceDeleteFavoriteQueryHandler.ServeHTTP(w, r)
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

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.UpdateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateOrganization(context.Context, *connect.Request[v1.CreateOrganizationRequest]) (*connect.Response[v1.CreateOrganizationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.CreateOrganization is not implemented"))
}

func (UnimplementedUserServiceHandler) ListOrganization(context.Context, *connect.Request[v1.ListOrganizationRequest]) (*connect.Response[v1.ListOrganizationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.ListOrganization is not implemented"))
}

func (UnimplementedUserServiceHandler) SaveLocalhostConfig(context.Context, *connect.Request[v1.SaveLocalhostConfigRequest]) (*connect.Response[v1.SaveLocalhostConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.SaveLocalhostConfig is not implemented"))
}

func (UnimplementedUserServiceHandler) GetLocalhostConfig(context.Context, *connect.Request[v1.GetLocalhostConfigRequest]) (*connect.Response[v1.GetLocalhostConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.GetLocalhostConfig is not implemented"))
}

func (UnimplementedUserServiceHandler) RecordQueryHistory(context.Context, *connect.Request[v1.RecordQueryHistoryRequest]) (*connect.Response[v1.RecordQueryHistoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.RecordQueryHistory is not implemented"))
}

func (UnimplementedUserServiceHandler) GetQueryHistory(context.Context, *connect.Request[v1.GetQueryHistoryRequest]) (*connect.Response[v1.GetQueryHistoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.GetQueryHistory is not implemented"))
}

func (UnimplementedUserServiceHandler) ListQueryHistory(context.Context, *connect.Request[v1.ListQueryHistoryRequest]) (*connect.Response[v1.ListQueryHistoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.ListQueryHistory is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteQueryHistory(context.Context, *connect.Request[v1.DeleteQueryHistoryRequest]) (*connect.Response[v1.DeleteQueryHistoryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.DeleteQueryHistory is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateFavoriteQuery(context.Context, *connect.Request[v1.CreateFavoriteQueryRequest]) (*connect.Response[v1.CreateFavoriteQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.CreateFavoriteQuery is not implemented"))
}

func (UnimplementedUserServiceHandler) GetFavoriteQuery(context.Context, *connect.Request[v1.GetFavoriteQueryRequest]) (*connect.Response[v1.GetFavoriteQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.GetFavoriteQuery is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateFavoriteQuery(context.Context, *connect.Request[v1.UpdateFavoriteQueryRequest]) (*connect.Response[v1.UpdateFavoriteQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.UpdateFavoriteQuery is not implemented"))
}

func (UnimplementedUserServiceHandler) ListFavoriteQuery(context.Context, *connect.Request[v1.ListFavoriteQueryRequest]) (*connect.Response[v1.ListFavoriteQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.ListFavoriteQuery is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteFavoriteQuery(context.Context, *connect.Request[v1.DeleteFavoriteQueryRequest]) (*connect.Response[v1.DeleteFavoriteQueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("svc.user.v1.UserService.DeleteFavoriteQuery is not implemented"))
}
