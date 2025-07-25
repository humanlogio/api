// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/shared_result.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";
import { QueryHistoryEntry } from "./query_history_entry_pb";
import { Data } from "./data_pb";

/**
 * @generated from enum types.v1.SharedResultVisibility
 */
export enum SharedResultVisibility {
  /**
   * @generated from enum value: INVALID = 0;
   */
  INVALID = 0,

  /**
   * @generated from enum value: PUBLIC = 1;
   */
  PUBLIC = 1,

  /**
   * @generated from enum value: ANYONE_WITH_LINK = 2;
   */
  ANYONE_WITH_LINK = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(SharedResultVisibility)
proto3.util.setEnumType(SharedResultVisibility, "types.v1.SharedResultVisibility", [
  { no: 0, name: "INVALID" },
  { no: 1, name: "PUBLIC" },
  { no: 2, name: "ANYONE_WITH_LINK" },
]);

/**
 * @generated from message types.v1.SharedResult
 */
export class SharedResult extends Message<SharedResult> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: types.v1.QueryHistoryEntry query = 2;
   */
  query?: QueryHistoryEntry;

  /**
   * @generated from field: types.v1.Data result = 3;
   */
  result?: Data;

  /**
   * @generated from oneof types.v1.SharedResult.visibility
   */
  visibility: {
    /**
     * @generated from field: types.v1.SharedResult.Public public = 400;
     */
    value: SharedResult_Public;
    case: "public";
  } | {
    /**
     * @generated from field: types.v1.SharedResult.AnyoneWithLink anyone_with_link = 401;
     */
    value: SharedResult_AnyoneWithLink;
    case: "anyoneWithLink";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  constructor(data?: PartialMessage<SharedResult>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.SharedResult";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "query", kind: "message", T: QueryHistoryEntry },
    { no: 3, name: "result", kind: "message", T: Data },
    { no: 400, name: "public", kind: "message", T: SharedResult_Public, oneof: "visibility" },
    { no: 401, name: "anyone_with_link", kind: "message", T: SharedResult_AnyoneWithLink, oneof: "visibility" },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SharedResult {
    return new SharedResult().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SharedResult {
    return new SharedResult().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SharedResult {
    return new SharedResult().fromJsonString(jsonString, options);
  }

  static equals(a: SharedResult | PlainMessage<SharedResult> | undefined, b: SharedResult | PlainMessage<SharedResult> | undefined): boolean {
    return proto3.util.equals(SharedResult, a, b);
  }
}

/**
 * @generated from message types.v1.SharedResult.Public
 */
export class SharedResult_Public extends Message<SharedResult_Public> {
  /**
   * @generated from field: string share_id = 1;
   */
  shareId = "";

  constructor(data?: PartialMessage<SharedResult_Public>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.SharedResult.Public";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "share_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SharedResult_Public {
    return new SharedResult_Public().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SharedResult_Public {
    return new SharedResult_Public().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SharedResult_Public {
    return new SharedResult_Public().fromJsonString(jsonString, options);
  }

  static equals(a: SharedResult_Public | PlainMessage<SharedResult_Public> | undefined, b: SharedResult_Public | PlainMessage<SharedResult_Public> | undefined): boolean {
    return proto3.util.equals(SharedResult_Public, a, b);
  }
}

/**
 * @generated from message types.v1.SharedResult.AnyoneWithLink
 */
export class SharedResult_AnyoneWithLink extends Message<SharedResult_AnyoneWithLink> {
  /**
   * @generated from field: string random_prefix = 1;
   */
  randomPrefix = "";

  /**
   * @generated from field: string share_id = 2;
   */
  shareId = "";

  constructor(data?: PartialMessage<SharedResult_AnyoneWithLink>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.SharedResult.AnyoneWithLink";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "random_prefix", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "share_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SharedResult_AnyoneWithLink {
    return new SharedResult_AnyoneWithLink().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SharedResult_AnyoneWithLink {
    return new SharedResult_AnyoneWithLink().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SharedResult_AnyoneWithLink {
    return new SharedResult_AnyoneWithLink().fromJsonString(jsonString, options);
  }

  static equals(a: SharedResult_AnyoneWithLink | PlainMessage<SharedResult_AnyoneWithLink> | undefined, b: SharedResult_AnyoneWithLink | PlainMessage<SharedResult_AnyoneWithLink> | undefined): boolean {
    return proto3.util.equals(SharedResult_AnyoneWithLink, a, b);
  }
}

