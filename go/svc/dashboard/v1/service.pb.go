// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: svc/dashboard/v1/service.proto

package dashboardv1

import (
	v1 "github.com/humanlogio/api/go/types/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDashboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EnvironmentId int64                  `protobuf:"varint,101,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ProjectName   string                 `protobuf:"bytes,102,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	IsReadonly    bool                   `protobuf:"varint,3,opt,name=is_readonly,json=isReadonly,proto3" json:"is_readonly,omitempty"`
	PersesJson    []byte                 `protobuf:"bytes,4,opt,name=perses_json,json=persesJson,proto3" json:"perses_json,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDashboardRequest) Reset() {
	*x = CreateDashboardRequest{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDashboardRequest) ProtoMessage() {}

func (x *CreateDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDashboardRequest.ProtoReflect.Descriptor instead.
func (*CreateDashboardRequest) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDashboardRequest) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *CreateDashboardRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *CreateDashboardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDashboardRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateDashboardRequest) GetIsReadonly() bool {
	if x != nil {
		return x.IsReadonly
	}
	return false
}

func (x *CreateDashboardRequest) GetPersesJson() []byte {
	if x != nil {
		return x.PersesJson
	}
	return nil
}

type CreateDashboardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dashboard     *v1.Dashboard          `protobuf:"bytes,1,opt,name=dashboard,proto3" json:"dashboard,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDashboardResponse) Reset() {
	*x = CreateDashboardResponse{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDashboardResponse) ProtoMessage() {}

func (x *CreateDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDashboardResponse.ProtoReflect.Descriptor instead.
func (*CreateDashboardResponse) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDashboardResponse) GetDashboard() *v1.Dashboard {
	if x != nil {
		return x.Dashboard
	}
	return nil
}

type GetDashboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EnvironmentId int64                  `protobuf:"varint,101,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ProjectName   string                 `protobuf:"bytes,102,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDashboardRequest) Reset() {
	*x = GetDashboardRequest{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardRequest) ProtoMessage() {}

func (x *GetDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardRequest.ProtoReflect.Descriptor instead.
func (*GetDashboardRequest) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetDashboardRequest) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *GetDashboardRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *GetDashboardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDashboardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dashboard     *v1.Dashboard          `protobuf:"bytes,1,opt,name=dashboard,proto3" json:"dashboard,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDashboardResponse) Reset() {
	*x = GetDashboardResponse{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardResponse) ProtoMessage() {}

func (x *GetDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDashboardResponse.ProtoReflect.Descriptor instead.
func (*GetDashboardResponse) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetDashboardResponse) GetDashboard() *v1.Dashboard {
	if x != nil {
		return x.Dashboard
	}
	return nil
}

type UpdateDashboardRequest struct {
	state         protoimpl.MessageState             `protogen:"open.v1"`
	EnvironmentId int64                              `protobuf:"varint,101,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ProjectName   string                             `protobuf:"bytes,102,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Id            string                             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Mutations     []*UpdateDashboardRequest_Mutation `protobuf:"bytes,2,rep,name=mutations,proto3" json:"mutations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDashboardRequest) Reset() {
	*x = UpdateDashboardRequest{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDashboardRequest) ProtoMessage() {}

func (x *UpdateDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDashboardRequest.ProtoReflect.Descriptor instead.
func (*UpdateDashboardRequest) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDashboardRequest) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *UpdateDashboardRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *UpdateDashboardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDashboardRequest) GetMutations() []*UpdateDashboardRequest_Mutation {
	if x != nil {
		return x.Mutations
	}
	return nil
}

type UpdateDashboardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dashboard     *v1.Dashboard          `protobuf:"bytes,1,opt,name=dashboard,proto3" json:"dashboard,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDashboardResponse) Reset() {
	*x = UpdateDashboardResponse{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDashboardResponse) ProtoMessage() {}

func (x *UpdateDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDashboardResponse.ProtoReflect.Descriptor instead.
func (*UpdateDashboardResponse) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDashboardResponse) GetDashboard() *v1.Dashboard {
	if x != nil {
		return x.Dashboard
	}
	return nil
}

type DeleteDashboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EnvironmentId int64                  `protobuf:"varint,101,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ProjectName   string                 `protobuf:"bytes,102,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDashboardRequest) Reset() {
	*x = DeleteDashboardRequest{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDashboardRequest) ProtoMessage() {}

func (x *DeleteDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDashboardRequest.ProtoReflect.Descriptor instead.
func (*DeleteDashboardRequest) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDashboardRequest) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *DeleteDashboardRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *DeleteDashboardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteDashboardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDashboardResponse) Reset() {
	*x = DeleteDashboardResponse{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDashboardResponse) ProtoMessage() {}

func (x *DeleteDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDashboardResponse.ProtoReflect.Descriptor instead.
func (*DeleteDashboardResponse) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{7}
}

type ListDashboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EnvironmentId int64                  `protobuf:"varint,101,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	ProjectName   string                 `protobuf:"bytes,102,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	Cursor        *v1.Cursor             `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	DashboardId   int64                  `protobuf:"varint,3,opt,name=dashboard_id,json=dashboardId,proto3" json:"dashboard_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDashboardRequest) Reset() {
	*x = ListDashboardRequest{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDashboardRequest) ProtoMessage() {}

func (x *ListDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDashboardRequest.ProtoReflect.Descriptor instead.
func (*ListDashboardRequest) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListDashboardRequest) GetEnvironmentId() int64 {
	if x != nil {
		return x.EnvironmentId
	}
	return 0
}

func (x *ListDashboardRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *ListDashboardRequest) GetCursor() *v1.Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ListDashboardRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListDashboardRequest) GetDashboardId() int64 {
	if x != nil {
		return x.DashboardId
	}
	return 0
}

type ListDashboardResponse struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	Next          *v1.Cursor                        `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items         []*ListDashboardResponse_ListItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDashboardResponse) Reset() {
	*x = ListDashboardResponse{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDashboardResponse) ProtoMessage() {}

func (x *ListDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDashboardResponse.ProtoReflect.Descriptor instead.
func (*ListDashboardResponse) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListDashboardResponse) GetNext() *v1.Cursor {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *ListDashboardResponse) GetItems() []*ListDashboardResponse_ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type UpdateDashboardRequest_Mutation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Do:
	//
	//	*UpdateDashboardRequest_Mutation_SetName
	//	*UpdateDashboardRequest_Mutation_SetDescription
	//	*UpdateDashboardRequest_Mutation_SetReadonly
	//	*UpdateDashboardRequest_Mutation_SetSourceFile
	//	*UpdateDashboardRequest_Mutation_SetPersesJson
	Do            isUpdateDashboardRequest_Mutation_Do `protobuf_oneof:"do"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDashboardRequest_Mutation) Reset() {
	*x = UpdateDashboardRequest_Mutation{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDashboardRequest_Mutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDashboardRequest_Mutation) ProtoMessage() {}

func (x *UpdateDashboardRequest_Mutation) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDashboardRequest_Mutation.ProtoReflect.Descriptor instead.
func (*UpdateDashboardRequest_Mutation) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *UpdateDashboardRequest_Mutation) GetDo() isUpdateDashboardRequest_Mutation_Do {
	if x != nil {
		return x.Do
	}
	return nil
}

func (x *UpdateDashboardRequest_Mutation) GetSetName() string {
	if x != nil {
		if x, ok := x.Do.(*UpdateDashboardRequest_Mutation_SetName); ok {
			return x.SetName
		}
	}
	return ""
}

func (x *UpdateDashboardRequest_Mutation) GetSetDescription() string {
	if x != nil {
		if x, ok := x.Do.(*UpdateDashboardRequest_Mutation_SetDescription); ok {
			return x.SetDescription
		}
	}
	return ""
}

func (x *UpdateDashboardRequest_Mutation) GetSetReadonly() bool {
	if x != nil {
		if x, ok := x.Do.(*UpdateDashboardRequest_Mutation_SetReadonly); ok {
			return x.SetReadonly
		}
	}
	return false
}

func (x *UpdateDashboardRequest_Mutation) GetSetSourceFile() string {
	if x != nil {
		if x, ok := x.Do.(*UpdateDashboardRequest_Mutation_SetSourceFile); ok {
			return x.SetSourceFile
		}
	}
	return ""
}

func (x *UpdateDashboardRequest_Mutation) GetSetPersesJson() []byte {
	if x != nil {
		if x, ok := x.Do.(*UpdateDashboardRequest_Mutation_SetPersesJson); ok {
			return x.SetPersesJson
		}
	}
	return nil
}

type isUpdateDashboardRequest_Mutation_Do interface {
	isUpdateDashboardRequest_Mutation_Do()
}

type UpdateDashboardRequest_Mutation_SetName struct {
	SetName string `protobuf:"bytes,1,opt,name=set_name,json=setName,proto3,oneof"`
}

type UpdateDashboardRequest_Mutation_SetDescription struct {
	SetDescription string `protobuf:"bytes,2,opt,name=set_description,json=setDescription,proto3,oneof"`
}

type UpdateDashboardRequest_Mutation_SetReadonly struct {
	SetReadonly bool `protobuf:"varint,3,opt,name=set_readonly,json=setReadonly,proto3,oneof"`
}

type UpdateDashboardRequest_Mutation_SetSourceFile struct {
	SetSourceFile string `protobuf:"bytes,401,opt,name=set_source_file,json=setSourceFile,proto3,oneof"`
}

type UpdateDashboardRequest_Mutation_SetPersesJson struct {
	SetPersesJson []byte `protobuf:"bytes,5,opt,name=set_perses_json,json=setPersesJson,proto3,oneof"`
}

func (*UpdateDashboardRequest_Mutation_SetName) isUpdateDashboardRequest_Mutation_Do() {}

func (*UpdateDashboardRequest_Mutation_SetDescription) isUpdateDashboardRequest_Mutation_Do() {}

func (*UpdateDashboardRequest_Mutation_SetReadonly) isUpdateDashboardRequest_Mutation_Do() {}

func (*UpdateDashboardRequest_Mutation_SetSourceFile) isUpdateDashboardRequest_Mutation_Do() {}

func (*UpdateDashboardRequest_Mutation_SetPersesJson) isUpdateDashboardRequest_Mutation_Do() {}

type ListDashboardResponse_ListItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Dashboard     *v1.Dashboard          `protobuf:"bytes,1,opt,name=dashboard,proto3" json:"dashboard,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDashboardResponse_ListItem) Reset() {
	*x = ListDashboardResponse_ListItem{}
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDashboardResponse_ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDashboardResponse_ListItem) ProtoMessage() {}

func (x *ListDashboardResponse_ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_svc_dashboard_v1_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDashboardResponse_ListItem.ProtoReflect.Descriptor instead.
func (*ListDashboardResponse_ListItem) Descriptor() ([]byte, []int) {
	return file_svc_dashboard_v1_service_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ListDashboardResponse_ListItem) GetDashboard() *v1.Dashboard {
	if x != nil {
		return x.Dashboard
	}
	return nil
}

var File_svc_dashboard_v1_service_proto protoreflect.FileDescriptor

const file_svc_dashboard_v1_service_proto_rawDesc = "" +
	"\n" +
	"\x1esvc/dashboard/v1/service.proto\x12\x10svc.dashboard.v1\x1a\x15types/v1/cursor.proto\x1a\x18types/v1/dashboard.proto\"\xda\x01\n" +
	"\x16CreateDashboardRequest\x12%\n" +
	"\x0eenvironment_id\x18e \x01(\x03R\renvironmentId\x12!\n" +
	"\fproject_name\x18f \x01(\tR\vprojectName\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x1f\n" +
	"\vis_readonly\x18\x03 \x01(\bR\n" +
	"isReadonly\x12\x1f\n" +
	"\vperses_json\x18\x04 \x01(\fR\n" +
	"persesJson\"L\n" +
	"\x17CreateDashboardResponse\x121\n" +
	"\tdashboard\x18\x01 \x01(\v2\x13.types.v1.DashboardR\tdashboard\"o\n" +
	"\x13GetDashboardRequest\x12%\n" +
	"\x0eenvironment_id\x18e \x01(\x03R\renvironmentId\x12!\n" +
	"\fproject_name\x18f \x01(\tR\vprojectName\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"I\n" +
	"\x14GetDashboardResponse\x121\n" +
	"\tdashboard\x18\x01 \x01(\v2\x13.types.v1.DashboardR\tdashboard\"\x98\x03\n" +
	"\x16UpdateDashboardRequest\x12%\n" +
	"\x0eenvironment_id\x18e \x01(\x03R\renvironmentId\x12!\n" +
	"\fproject_name\x18f \x01(\tR\vprojectName\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12O\n" +
	"\tmutations\x18\x02 \x03(\v21.svc.dashboard.v1.UpdateDashboardRequest.MutationR\tmutations\x1a\xd2\x01\n" +
	"\bMutation\x12\x1b\n" +
	"\bset_name\x18\x01 \x01(\tH\x00R\asetName\x12)\n" +
	"\x0fset_description\x18\x02 \x01(\tH\x00R\x0esetDescription\x12#\n" +
	"\fset_readonly\x18\x03 \x01(\bH\x00R\vsetReadonly\x12)\n" +
	"\x0fset_source_file\x18\x91\x03 \x01(\tH\x00R\rsetSourceFile\x12(\n" +
	"\x0fset_perses_json\x18\x05 \x01(\fH\x00R\rsetPersesJsonB\x04\n" +
	"\x02do\"L\n" +
	"\x17UpdateDashboardResponse\x121\n" +
	"\tdashboard\x18\x01 \x01(\v2\x13.types.v1.DashboardR\tdashboard\"r\n" +
	"\x16DeleteDashboardRequest\x12%\n" +
	"\x0eenvironment_id\x18e \x01(\x03R\renvironmentId\x12!\n" +
	"\fproject_name\x18f \x01(\tR\vprojectName\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x19\n" +
	"\x17DeleteDashboardResponse\"\xc3\x01\n" +
	"\x14ListDashboardRequest\x12%\n" +
	"\x0eenvironment_id\x18e \x01(\x03R\renvironmentId\x12!\n" +
	"\fproject_name\x18f \x01(\tR\vprojectName\x12(\n" +
	"\x06cursor\x18\x01 \x01(\v2\x10.types.v1.CursorR\x06cursor\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12!\n" +
	"\fdashboard_id\x18\x03 \x01(\x03R\vdashboardId\"\xc4\x01\n" +
	"\x15ListDashboardResponse\x12$\n" +
	"\x04next\x18\x01 \x01(\v2\x10.types.v1.CursorR\x04next\x12F\n" +
	"\x05items\x18\x02 \x03(\v20.svc.dashboard.v1.ListDashboardResponse.ListItemR\x05items\x1a=\n" +
	"\bListItem\x121\n" +
	"\tdashboard\x18\x01 \x01(\v2\x13.types.v1.DashboardR\tdashboard2\x8b\x04\n" +
	"\x10DashboardService\x12f\n" +
	"\x0fCreateDashboard\x12(.svc.dashboard.v1.CreateDashboardRequest\x1a).svc.dashboard.v1.CreateDashboardResponse\x12]\n" +
	"\fGetDashboard\x12%.svc.dashboard.v1.GetDashboardRequest\x1a&.svc.dashboard.v1.GetDashboardResponse\x12f\n" +
	"\x0fUpdateDashboard\x12(.svc.dashboard.v1.UpdateDashboardRequest\x1a).svc.dashboard.v1.UpdateDashboardResponse\x12f\n" +
	"\x0fDeleteDashboard\x12(.svc.dashboard.v1.DeleteDashboardRequest\x1a).svc.dashboard.v1.DeleteDashboardResponse\x12`\n" +
	"\rListDashboard\x12&.svc.dashboard.v1.ListDashboardRequest\x1a'.svc.dashboard.v1.ListDashboardResponseB\xc1\x01\n" +
	"\x14com.svc.dashboard.v1B\fServiceProtoP\x01Z9github.com/humanlogio/api/go/svc/dashboard/v1;dashboardv1\xa2\x02\x03SDX\xaa\x02\x10Svc.Dashboard.V1\xca\x02\x10Svc\\Dashboard\\V1\xe2\x02\x1cSvc\\Dashboard\\V1\\GPBMetadata\xea\x02\x12Svc::Dashboard::V1b\x06proto3"

var (
	file_svc_dashboard_v1_service_proto_rawDescOnce sync.Once
	file_svc_dashboard_v1_service_proto_rawDescData []byte
)

func file_svc_dashboard_v1_service_proto_rawDescGZIP() []byte {
	file_svc_dashboard_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_dashboard_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_svc_dashboard_v1_service_proto_rawDesc), len(file_svc_dashboard_v1_service_proto_rawDesc)))
	})
	return file_svc_dashboard_v1_service_proto_rawDescData
}

var file_svc_dashboard_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_svc_dashboard_v1_service_proto_goTypes = []any{
	(*CreateDashboardRequest)(nil),          // 0: svc.dashboard.v1.CreateDashboardRequest
	(*CreateDashboardResponse)(nil),         // 1: svc.dashboard.v1.CreateDashboardResponse
	(*GetDashboardRequest)(nil),             // 2: svc.dashboard.v1.GetDashboardRequest
	(*GetDashboardResponse)(nil),            // 3: svc.dashboard.v1.GetDashboardResponse
	(*UpdateDashboardRequest)(nil),          // 4: svc.dashboard.v1.UpdateDashboardRequest
	(*UpdateDashboardResponse)(nil),         // 5: svc.dashboard.v1.UpdateDashboardResponse
	(*DeleteDashboardRequest)(nil),          // 6: svc.dashboard.v1.DeleteDashboardRequest
	(*DeleteDashboardResponse)(nil),         // 7: svc.dashboard.v1.DeleteDashboardResponse
	(*ListDashboardRequest)(nil),            // 8: svc.dashboard.v1.ListDashboardRequest
	(*ListDashboardResponse)(nil),           // 9: svc.dashboard.v1.ListDashboardResponse
	(*UpdateDashboardRequest_Mutation)(nil), // 10: svc.dashboard.v1.UpdateDashboardRequest.Mutation
	(*ListDashboardResponse_ListItem)(nil),  // 11: svc.dashboard.v1.ListDashboardResponse.ListItem
	(*v1.Dashboard)(nil),                    // 12: types.v1.Dashboard
	(*v1.Cursor)(nil),                       // 13: types.v1.Cursor
}
var file_svc_dashboard_v1_service_proto_depIdxs = []int32{
	12, // 0: svc.dashboard.v1.CreateDashboardResponse.dashboard:type_name -> types.v1.Dashboard
	12, // 1: svc.dashboard.v1.GetDashboardResponse.dashboard:type_name -> types.v1.Dashboard
	10, // 2: svc.dashboard.v1.UpdateDashboardRequest.mutations:type_name -> svc.dashboard.v1.UpdateDashboardRequest.Mutation
	12, // 3: svc.dashboard.v1.UpdateDashboardResponse.dashboard:type_name -> types.v1.Dashboard
	13, // 4: svc.dashboard.v1.ListDashboardRequest.cursor:type_name -> types.v1.Cursor
	13, // 5: svc.dashboard.v1.ListDashboardResponse.next:type_name -> types.v1.Cursor
	11, // 6: svc.dashboard.v1.ListDashboardResponse.items:type_name -> svc.dashboard.v1.ListDashboardResponse.ListItem
	12, // 7: svc.dashboard.v1.ListDashboardResponse.ListItem.dashboard:type_name -> types.v1.Dashboard
	0,  // 8: svc.dashboard.v1.DashboardService.CreateDashboard:input_type -> svc.dashboard.v1.CreateDashboardRequest
	2,  // 9: svc.dashboard.v1.DashboardService.GetDashboard:input_type -> svc.dashboard.v1.GetDashboardRequest
	4,  // 10: svc.dashboard.v1.DashboardService.UpdateDashboard:input_type -> svc.dashboard.v1.UpdateDashboardRequest
	6,  // 11: svc.dashboard.v1.DashboardService.DeleteDashboard:input_type -> svc.dashboard.v1.DeleteDashboardRequest
	8,  // 12: svc.dashboard.v1.DashboardService.ListDashboard:input_type -> svc.dashboard.v1.ListDashboardRequest
	1,  // 13: svc.dashboard.v1.DashboardService.CreateDashboard:output_type -> svc.dashboard.v1.CreateDashboardResponse
	3,  // 14: svc.dashboard.v1.DashboardService.GetDashboard:output_type -> svc.dashboard.v1.GetDashboardResponse
	5,  // 15: svc.dashboard.v1.DashboardService.UpdateDashboard:output_type -> svc.dashboard.v1.UpdateDashboardResponse
	7,  // 16: svc.dashboard.v1.DashboardService.DeleteDashboard:output_type -> svc.dashboard.v1.DeleteDashboardResponse
	9,  // 17: svc.dashboard.v1.DashboardService.ListDashboard:output_type -> svc.dashboard.v1.ListDashboardResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_svc_dashboard_v1_service_proto_init() }
func file_svc_dashboard_v1_service_proto_init() {
	if File_svc_dashboard_v1_service_proto != nil {
		return
	}
	file_svc_dashboard_v1_service_proto_msgTypes[10].OneofWrappers = []any{
		(*UpdateDashboardRequest_Mutation_SetName)(nil),
		(*UpdateDashboardRequest_Mutation_SetDescription)(nil),
		(*UpdateDashboardRequest_Mutation_SetReadonly)(nil),
		(*UpdateDashboardRequest_Mutation_SetSourceFile)(nil),
		(*UpdateDashboardRequest_Mutation_SetPersesJson)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_svc_dashboard_v1_service_proto_rawDesc), len(file_svc_dashboard_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_dashboard_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_dashboard_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_dashboard_v1_service_proto_msgTypes,
	}.Build()
	File_svc_dashboard_v1_service_proto = out.File
	file_svc_dashboard_v1_service_proto_goTypes = nil
	file_svc_dashboard_v1_service_proto_depIdxs = nil
}
