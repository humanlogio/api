// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/dashboard/v1/service.proto (package svc.dashboard.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Dashboard } from "../../../types/v1/dashboard_pb";
import { Cursor } from "../../../types/v1/cursor_pb";

/**
 * @generated from message svc.dashboard.v1.CreateDashboardRequest
 */
export class CreateDashboardRequest extends Message<CreateDashboardRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string project_name = 102;
   */
  projectName = "";

  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string description = 2;
   */
  description = "";

  /**
   * @generated from field: bool is_readonly = 3;
   */
  isReadonly = false;

  /**
   * @generated from field: bytes perses_json = 4;
   */
  persesJson = new Uint8Array(0);

  constructor(data?: PartialMessage<CreateDashboardRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.CreateDashboardRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 102, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "is_readonly", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "perses_json", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDashboardRequest {
    return new CreateDashboardRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDashboardRequest {
    return new CreateDashboardRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDashboardRequest {
    return new CreateDashboardRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDashboardRequest | PlainMessage<CreateDashboardRequest> | undefined, b: CreateDashboardRequest | PlainMessage<CreateDashboardRequest> | undefined): boolean {
    return proto3.util.equals(CreateDashboardRequest, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.CreateDashboardResponse
 */
export class CreateDashboardResponse extends Message<CreateDashboardResponse> {
  /**
   * @generated from field: types.v1.Dashboard dashboard = 1;
   */
  dashboard?: Dashboard;

  constructor(data?: PartialMessage<CreateDashboardResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.CreateDashboardResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dashboard", kind: "message", T: Dashboard },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateDashboardResponse {
    return new CreateDashboardResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateDashboardResponse {
    return new CreateDashboardResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateDashboardResponse {
    return new CreateDashboardResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateDashboardResponse | PlainMessage<CreateDashboardResponse> | undefined, b: CreateDashboardResponse | PlainMessage<CreateDashboardResponse> | undefined): boolean {
    return proto3.util.equals(CreateDashboardResponse, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.GetDashboardRequest
 */
export class GetDashboardRequest extends Message<GetDashboardRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string project_name = 102;
   */
  projectName = "";

  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetDashboardRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.GetDashboardRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 102, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDashboardRequest {
    return new GetDashboardRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDashboardRequest {
    return new GetDashboardRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDashboardRequest {
    return new GetDashboardRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetDashboardRequest | PlainMessage<GetDashboardRequest> | undefined, b: GetDashboardRequest | PlainMessage<GetDashboardRequest> | undefined): boolean {
    return proto3.util.equals(GetDashboardRequest, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.GetDashboardResponse
 */
export class GetDashboardResponse extends Message<GetDashboardResponse> {
  /**
   * @generated from field: types.v1.Dashboard dashboard = 1;
   */
  dashboard?: Dashboard;

  constructor(data?: PartialMessage<GetDashboardResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.GetDashboardResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dashboard", kind: "message", T: Dashboard },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDashboardResponse {
    return new GetDashboardResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDashboardResponse {
    return new GetDashboardResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDashboardResponse {
    return new GetDashboardResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetDashboardResponse | PlainMessage<GetDashboardResponse> | undefined, b: GetDashboardResponse | PlainMessage<GetDashboardResponse> | undefined): boolean {
    return proto3.util.equals(GetDashboardResponse, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.UpdateDashboardRequest
 */
export class UpdateDashboardRequest extends Message<UpdateDashboardRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string project_name = 102;
   */
  projectName = "";

  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: repeated svc.dashboard.v1.UpdateDashboardRequest.Mutation mutations = 2;
   */
  mutations: UpdateDashboardRequest_Mutation[] = [];

  constructor(data?: PartialMessage<UpdateDashboardRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.UpdateDashboardRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 102, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "mutations", kind: "message", T: UpdateDashboardRequest_Mutation, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateDashboardRequest {
    return new UpdateDashboardRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateDashboardRequest {
    return new UpdateDashboardRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateDashboardRequest {
    return new UpdateDashboardRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateDashboardRequest | PlainMessage<UpdateDashboardRequest> | undefined, b: UpdateDashboardRequest | PlainMessage<UpdateDashboardRequest> | undefined): boolean {
    return proto3.util.equals(UpdateDashboardRequest, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.UpdateDashboardRequest.Mutation
 */
export class UpdateDashboardRequest_Mutation extends Message<UpdateDashboardRequest_Mutation> {
  /**
   * @generated from oneof svc.dashboard.v1.UpdateDashboardRequest.Mutation.do
   */
  do: {
    /**
     * @generated from field: string set_name = 1;
     */
    value: string;
    case: "setName";
  } | {
    /**
     * @generated from field: string set_description = 2;
     */
    value: string;
    case: "setDescription";
  } | {
    /**
     * @generated from field: bool set_readonly = 3;
     */
    value: boolean;
    case: "setReadonly";
  } | {
    /**
     * @generated from field: string set_source_file = 401;
     */
    value: string;
    case: "setSourceFile";
  } | {
    /**
     * @generated from field: bytes set_perses_json = 5;
     */
    value: Uint8Array;
    case: "setPersesJson";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<UpdateDashboardRequest_Mutation>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.UpdateDashboardRequest.Mutation";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "set_name", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "do" },
    { no: 2, name: "set_description", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "do" },
    { no: 3, name: "set_readonly", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "do" },
    { no: 401, name: "set_source_file", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "do" },
    { no: 5, name: "set_perses_json", kind: "scalar", T: 12 /* ScalarType.BYTES */, oneof: "do" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateDashboardRequest_Mutation {
    return new UpdateDashboardRequest_Mutation().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateDashboardRequest_Mutation {
    return new UpdateDashboardRequest_Mutation().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateDashboardRequest_Mutation {
    return new UpdateDashboardRequest_Mutation().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateDashboardRequest_Mutation | PlainMessage<UpdateDashboardRequest_Mutation> | undefined, b: UpdateDashboardRequest_Mutation | PlainMessage<UpdateDashboardRequest_Mutation> | undefined): boolean {
    return proto3.util.equals(UpdateDashboardRequest_Mutation, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.UpdateDashboardResponse
 */
export class UpdateDashboardResponse extends Message<UpdateDashboardResponse> {
  /**
   * @generated from field: types.v1.Dashboard dashboard = 1;
   */
  dashboard?: Dashboard;

  constructor(data?: PartialMessage<UpdateDashboardResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.UpdateDashboardResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dashboard", kind: "message", T: Dashboard },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateDashboardResponse {
    return new UpdateDashboardResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateDashboardResponse {
    return new UpdateDashboardResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateDashboardResponse {
    return new UpdateDashboardResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateDashboardResponse | PlainMessage<UpdateDashboardResponse> | undefined, b: UpdateDashboardResponse | PlainMessage<UpdateDashboardResponse> | undefined): boolean {
    return proto3.util.equals(UpdateDashboardResponse, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.DeleteDashboardRequest
 */
export class DeleteDashboardRequest extends Message<DeleteDashboardRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string project_name = 102;
   */
  projectName = "";

  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<DeleteDashboardRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.DeleteDashboardRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 102, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteDashboardRequest {
    return new DeleteDashboardRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteDashboardRequest {
    return new DeleteDashboardRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteDashboardRequest {
    return new DeleteDashboardRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteDashboardRequest | PlainMessage<DeleteDashboardRequest> | undefined, b: DeleteDashboardRequest | PlainMessage<DeleteDashboardRequest> | undefined): boolean {
    return proto3.util.equals(DeleteDashboardRequest, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.DeleteDashboardResponse
 */
export class DeleteDashboardResponse extends Message<DeleteDashboardResponse> {
  constructor(data?: PartialMessage<DeleteDashboardResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.DeleteDashboardResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteDashboardResponse {
    return new DeleteDashboardResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteDashboardResponse {
    return new DeleteDashboardResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteDashboardResponse {
    return new DeleteDashboardResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteDashboardResponse | PlainMessage<DeleteDashboardResponse> | undefined, b: DeleteDashboardResponse | PlainMessage<DeleteDashboardResponse> | undefined): boolean {
    return proto3.util.equals(DeleteDashboardResponse, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.ListDashboardRequest
 */
export class ListDashboardRequest extends Message<ListDashboardRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string project_name = 102;
   */
  projectName = "";

  /**
   * @generated from field: types.v1.Cursor cursor = 1;
   */
  cursor?: Cursor;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  /**
   * @generated from field: int64 dashboard_id = 3;
   */
  dashboardId = protoInt64.zero;

  constructor(data?: PartialMessage<ListDashboardRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.ListDashboardRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 102, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 1, name: "cursor", kind: "message", T: Cursor },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "dashboard_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListDashboardRequest {
    return new ListDashboardRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListDashboardRequest {
    return new ListDashboardRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListDashboardRequest {
    return new ListDashboardRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListDashboardRequest | PlainMessage<ListDashboardRequest> | undefined, b: ListDashboardRequest | PlainMessage<ListDashboardRequest> | undefined): boolean {
    return proto3.util.equals(ListDashboardRequest, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.ListDashboardResponse
 */
export class ListDashboardResponse extends Message<ListDashboardResponse> {
  /**
   * @generated from field: types.v1.Cursor next = 1;
   */
  next?: Cursor;

  /**
   * @generated from field: repeated svc.dashboard.v1.ListDashboardResponse.ListItem items = 2;
   */
  items: ListDashboardResponse_ListItem[] = [];

  constructor(data?: PartialMessage<ListDashboardResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.ListDashboardResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "message", T: Cursor },
    { no: 2, name: "items", kind: "message", T: ListDashboardResponse_ListItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListDashboardResponse {
    return new ListDashboardResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListDashboardResponse {
    return new ListDashboardResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListDashboardResponse {
    return new ListDashboardResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListDashboardResponse | PlainMessage<ListDashboardResponse> | undefined, b: ListDashboardResponse | PlainMessage<ListDashboardResponse> | undefined): boolean {
    return proto3.util.equals(ListDashboardResponse, a, b);
  }
}

/**
 * @generated from message svc.dashboard.v1.ListDashboardResponse.ListItem
 */
export class ListDashboardResponse_ListItem extends Message<ListDashboardResponse_ListItem> {
  /**
   * @generated from field: types.v1.Dashboard dashboard = 1;
   */
  dashboard?: Dashboard;

  constructor(data?: PartialMessage<ListDashboardResponse_ListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.dashboard.v1.ListDashboardResponse.ListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "dashboard", kind: "message", T: Dashboard },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListDashboardResponse_ListItem {
    return new ListDashboardResponse_ListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListDashboardResponse_ListItem {
    return new ListDashboardResponse_ListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListDashboardResponse_ListItem {
    return new ListDashboardResponse_ListItem().fromJsonString(jsonString, options);
  }

  static equals(a: ListDashboardResponse_ListItem | PlainMessage<ListDashboardResponse_ListItem> | undefined, b: ListDashboardResponse_ListItem | PlainMessage<ListDashboardResponse_ListItem> | undefined): boolean {
    return proto3.util.equals(ListDashboardResponse_ListItem, a, b);
  }
}

