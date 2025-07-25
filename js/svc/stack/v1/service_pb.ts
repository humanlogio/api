// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/stack/v1/service.proto (package svc.stack.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Stack, StackPointer } from "../../../types/v1/stack_pb";
import { Dashboard } from "../../../types/v1/dashboard_pb";
import { AlertGroup } from "../../../types/v1/alert_pb";
import { Cursor } from "../../../types/v1/cursor_pb";

/**
 * @generated from message svc.stack.v1.CreateStackRequest
 */
export class CreateStackRequest extends Message<CreateStackRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: types.v1.StackPointer pointer = 2;
   */
  pointer?: StackPointer;

  constructor(data?: PartialMessage<CreateStackRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.CreateStackRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "pointer", kind: "message", T: StackPointer },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateStackRequest {
    return new CreateStackRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateStackRequest {
    return new CreateStackRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateStackRequest {
    return new CreateStackRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateStackRequest | PlainMessage<CreateStackRequest> | undefined, b: CreateStackRequest | PlainMessage<CreateStackRequest> | undefined): boolean {
    return proto3.util.equals(CreateStackRequest, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.CreateStackResponse
 */
export class CreateStackResponse extends Message<CreateStackResponse> {
  /**
   * @generated from field: types.v1.Stack stack = 1;
   */
  stack?: Stack;

  constructor(data?: PartialMessage<CreateStackResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.CreateStackResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "stack", kind: "message", T: Stack },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateStackResponse {
    return new CreateStackResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateStackResponse {
    return new CreateStackResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateStackResponse {
    return new CreateStackResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateStackResponse | PlainMessage<CreateStackResponse> | undefined, b: CreateStackResponse | PlainMessage<CreateStackResponse> | undefined): boolean {
    return proto3.util.equals(CreateStackResponse, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.GetStackRequest
 */
export class GetStackRequest extends Message<GetStackRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<GetStackRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.GetStackRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStackRequest {
    return new GetStackRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStackRequest {
    return new GetStackRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStackRequest {
    return new GetStackRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetStackRequest | PlainMessage<GetStackRequest> | undefined, b: GetStackRequest | PlainMessage<GetStackRequest> | undefined): boolean {
    return proto3.util.equals(GetStackRequest, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.GetStackResponse
 */
export class GetStackResponse extends Message<GetStackResponse> {
  /**
   * @generated from field: types.v1.Stack stack = 1;
   */
  stack?: Stack;

  /**
   * @generated from field: repeated types.v1.Dashboard dashboards = 2;
   */
  dashboards: Dashboard[] = [];

  /**
   * @generated from field: repeated types.v1.AlertGroup alert_groups = 3;
   */
  alertGroups: AlertGroup[] = [];

  constructor(data?: PartialMessage<GetStackResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.GetStackResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "stack", kind: "message", T: Stack },
    { no: 2, name: "dashboards", kind: "message", T: Dashboard, repeated: true },
    { no: 3, name: "alert_groups", kind: "message", T: AlertGroup, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetStackResponse {
    return new GetStackResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetStackResponse {
    return new GetStackResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetStackResponse {
    return new GetStackResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetStackResponse | PlainMessage<GetStackResponse> | undefined, b: GetStackResponse | PlainMessage<GetStackResponse> | undefined): boolean {
    return proto3.util.equals(GetStackResponse, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.UpdateStackRequest
 */
export class UpdateStackRequest extends Message<UpdateStackRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: repeated svc.stack.v1.UpdateStackRequest.Mutation mutations = 2;
   */
  mutations: UpdateStackRequest_Mutation[] = [];

  constructor(data?: PartialMessage<UpdateStackRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.UpdateStackRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "mutations", kind: "message", T: UpdateStackRequest_Mutation, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateStackRequest {
    return new UpdateStackRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateStackRequest {
    return new UpdateStackRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateStackRequest {
    return new UpdateStackRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateStackRequest | PlainMessage<UpdateStackRequest> | undefined, b: UpdateStackRequest | PlainMessage<UpdateStackRequest> | undefined): boolean {
    return proto3.util.equals(UpdateStackRequest, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.UpdateStackRequest.Mutation
 */
export class UpdateStackRequest_Mutation extends Message<UpdateStackRequest_Mutation> {
  /**
   * @generated from oneof svc.stack.v1.UpdateStackRequest.Mutation.do
   */
  do: {
    /**
     * @generated from field: string set_name = 1;
     */
    value: string;
    case: "setName";
  } | {
    /**
     * @generated from field: types.v1.StackPointer set_pointer = 2;
     */
    value: StackPointer;
    case: "setPointer";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<UpdateStackRequest_Mutation>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.UpdateStackRequest.Mutation";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "set_name", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "do" },
    { no: 2, name: "set_pointer", kind: "message", T: StackPointer, oneof: "do" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateStackRequest_Mutation {
    return new UpdateStackRequest_Mutation().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateStackRequest_Mutation {
    return new UpdateStackRequest_Mutation().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateStackRequest_Mutation {
    return new UpdateStackRequest_Mutation().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateStackRequest_Mutation | PlainMessage<UpdateStackRequest_Mutation> | undefined, b: UpdateStackRequest_Mutation | PlainMessage<UpdateStackRequest_Mutation> | undefined): boolean {
    return proto3.util.equals(UpdateStackRequest_Mutation, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.UpdateStackResponse
 */
export class UpdateStackResponse extends Message<UpdateStackResponse> {
  /**
   * @generated from field: types.v1.Stack stack = 1;
   */
  stack?: Stack;

  constructor(data?: PartialMessage<UpdateStackResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.UpdateStackResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "stack", kind: "message", T: Stack },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateStackResponse {
    return new UpdateStackResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateStackResponse {
    return new UpdateStackResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateStackResponse {
    return new UpdateStackResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateStackResponse | PlainMessage<UpdateStackResponse> | undefined, b: UpdateStackResponse | PlainMessage<UpdateStackResponse> | undefined): boolean {
    return proto3.util.equals(UpdateStackResponse, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.DeleteStackRequest
 */
export class DeleteStackRequest extends Message<DeleteStackRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<DeleteStackRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.DeleteStackRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteStackRequest {
    return new DeleteStackRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteStackRequest {
    return new DeleteStackRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteStackRequest {
    return new DeleteStackRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteStackRequest | PlainMessage<DeleteStackRequest> | undefined, b: DeleteStackRequest | PlainMessage<DeleteStackRequest> | undefined): boolean {
    return proto3.util.equals(DeleteStackRequest, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.DeleteStackResponse
 */
export class DeleteStackResponse extends Message<DeleteStackResponse> {
  constructor(data?: PartialMessage<DeleteStackResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.DeleteStackResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteStackResponse {
    return new DeleteStackResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteStackResponse {
    return new DeleteStackResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteStackResponse {
    return new DeleteStackResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteStackResponse | PlainMessage<DeleteStackResponse> | undefined, b: DeleteStackResponse | PlainMessage<DeleteStackResponse> | undefined): boolean {
    return proto3.util.equals(DeleteStackResponse, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.ListStackRequest
 */
export class ListStackRequest extends Message<ListStackRequest> {
  /**
   * @generated from field: int64 environment_id = 101;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: types.v1.Cursor cursor = 1;
   */
  cursor?: Cursor;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  constructor(data?: PartialMessage<ListStackRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.ListStackRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 101, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 1, name: "cursor", kind: "message", T: Cursor },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListStackRequest {
    return new ListStackRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListStackRequest {
    return new ListStackRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListStackRequest {
    return new ListStackRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListStackRequest | PlainMessage<ListStackRequest> | undefined, b: ListStackRequest | PlainMessage<ListStackRequest> | undefined): boolean {
    return proto3.util.equals(ListStackRequest, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.ListStackResponse
 */
export class ListStackResponse extends Message<ListStackResponse> {
  /**
   * @generated from field: types.v1.Cursor next = 1;
   */
  next?: Cursor;

  /**
   * @generated from field: repeated svc.stack.v1.ListStackResponse.ListItem items = 2;
   */
  items: ListStackResponse_ListItem[] = [];

  constructor(data?: PartialMessage<ListStackResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.ListStackResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "message", T: Cursor },
    { no: 2, name: "items", kind: "message", T: ListStackResponse_ListItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListStackResponse {
    return new ListStackResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListStackResponse {
    return new ListStackResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListStackResponse {
    return new ListStackResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListStackResponse | PlainMessage<ListStackResponse> | undefined, b: ListStackResponse | PlainMessage<ListStackResponse> | undefined): boolean {
    return proto3.util.equals(ListStackResponse, a, b);
  }
}

/**
 * @generated from message svc.stack.v1.ListStackResponse.ListItem
 */
export class ListStackResponse_ListItem extends Message<ListStackResponse_ListItem> {
  /**
   * @generated from field: types.v1.Stack stack = 1;
   */
  stack?: Stack;

  constructor(data?: PartialMessage<ListStackResponse_ListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.stack.v1.ListStackResponse.ListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "stack", kind: "message", T: Stack },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListStackResponse_ListItem {
    return new ListStackResponse_ListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListStackResponse_ListItem {
    return new ListStackResponse_ListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListStackResponse_ListItem {
    return new ListStackResponse_ListItem().fromJsonString(jsonString, options);
  }

  static equals(a: ListStackResponse_ListItem | PlainMessage<ListStackResponse_ListItem> | undefined, b: ListStackResponse_ListItem | PlainMessage<ListStackResponse_ListItem> | undefined): boolean {
    return proto3.util.equals(ListStackResponse_ListItem, a, b);
  }
}

