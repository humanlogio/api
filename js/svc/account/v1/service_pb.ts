// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file svc/account/v1/service.proto (package svc.account.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Cursor } from "../../../types/v1/cursor_pb";
import { Machine } from "../../../types/v1/machine_pb";

/**
 * @generated from message svc.account.v1.ListMachineRequest
 */
export class ListMachineRequest extends Message<ListMachineRequest> {
  /**
   * @generated from field: types.v1.Cursor cursor = 1;
   */
  cursor?: Cursor;

  /**
   * @generated from field: int32 limit = 2;
   */
  limit = 0;

  /**
   * @generated from field: int64 account_id = 3;
   */
  accountId = protoInt64.zero;

  constructor(data?: PartialMessage<ListMachineRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.account.v1.ListMachineRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "cursor", kind: "message", T: Cursor },
    { no: 2, name: "limit", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "account_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListMachineRequest {
    return new ListMachineRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListMachineRequest {
    return new ListMachineRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListMachineRequest {
    return new ListMachineRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListMachineRequest | PlainMessage<ListMachineRequest> | undefined, b: ListMachineRequest | PlainMessage<ListMachineRequest> | undefined): boolean {
    return proto3.util.equals(ListMachineRequest, a, b);
  }
}

/**
 * @generated from message svc.account.v1.ListMachineResponse
 */
export class ListMachineResponse extends Message<ListMachineResponse> {
  /**
   * @generated from field: types.v1.Cursor next = 1;
   */
  next?: Cursor;

  /**
   * @generated from field: repeated svc.account.v1.ListMachineResponse.ListItem items = 2;
   */
  items: ListMachineResponse_ListItem[] = [];

  constructor(data?: PartialMessage<ListMachineResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.account.v1.ListMachineResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next", kind: "message", T: Cursor },
    { no: 2, name: "items", kind: "message", T: ListMachineResponse_ListItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListMachineResponse {
    return new ListMachineResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListMachineResponse {
    return new ListMachineResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListMachineResponse {
    return new ListMachineResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListMachineResponse | PlainMessage<ListMachineResponse> | undefined, b: ListMachineResponse | PlainMessage<ListMachineResponse> | undefined): boolean {
    return proto3.util.equals(ListMachineResponse, a, b);
  }
}

/**
 * @generated from message svc.account.v1.ListMachineResponse.ListItem
 */
export class ListMachineResponse_ListItem extends Message<ListMachineResponse_ListItem> {
  /**
   * @generated from field: types.v1.Machine machine = 1;
   */
  machine?: Machine;

  constructor(data?: PartialMessage<ListMachineResponse_ListItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.account.v1.ListMachineResponse.ListItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "machine", kind: "message", T: Machine },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListMachineResponse_ListItem {
    return new ListMachineResponse_ListItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListMachineResponse_ListItem {
    return new ListMachineResponse_ListItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListMachineResponse_ListItem {
    return new ListMachineResponse_ListItem().fromJsonString(jsonString, options);
  }

  static equals(a: ListMachineResponse_ListItem | PlainMessage<ListMachineResponse_ListItem> | undefined, b: ListMachineResponse_ListItem | PlainMessage<ListMachineResponse_ListItem> | undefined): boolean {
    return proto3.util.equals(ListMachineResponse_ListItem, a, b);
  }
}
