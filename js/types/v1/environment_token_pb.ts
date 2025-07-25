// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/environment_token.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum types.v1.EnvironmentRole
 */
export enum EnvironmentRole {
  /**
   * @generated from enum value: EnvironmentRole_Invalid = 0;
   */
  EnvironmentRole_Invalid = 0,

  /**
   * @generated from enum value: EnvironmentRole_Ingestor = 1;
   */
  EnvironmentRole_Ingestor = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(EnvironmentRole)
proto3.util.setEnumType(EnvironmentRole, "types.v1.EnvironmentRole", [
  { no: 0, name: "EnvironmentRole_Invalid" },
  { no: 1, name: "EnvironmentRole_Ingestor" },
]);

/**
 * @generated from message types.v1.EnvironmentToken
 */
export class EnvironmentToken extends Message<EnvironmentToken> {
  /**
   * @generated from field: int64 environment_id = 1;
   */
  environmentId = protoInt64.zero;

  /**
   * @generated from field: int64 token_id = 2;
   */
  tokenId = protoInt64.zero;

  /**
   * @generated from field: string token = 3;
   */
  token = "";

  /**
   * @generated from field: repeated types.v1.EnvironmentRole roles = 4;
   */
  roles: EnvironmentRole[] = [];

  /**
   * @generated from field: google.protobuf.Timestamp expires_at = 5;
   */
  expiresAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp revoked_at = 6;
   */
  revokedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp last_used_at = 7;
   */
  lastUsedAt?: Timestamp;

  constructor(data?: PartialMessage<EnvironmentToken>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.EnvironmentToken";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "environment_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "token_id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "roles", kind: "enum", T: proto3.getEnumType(EnvironmentRole), repeated: true },
    { no: 5, name: "expires_at", kind: "message", T: Timestamp },
    { no: 6, name: "revoked_at", kind: "message", T: Timestamp },
    { no: 7, name: "last_used_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EnvironmentToken {
    return new EnvironmentToken().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EnvironmentToken {
    return new EnvironmentToken().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EnvironmentToken {
    return new EnvironmentToken().fromJsonString(jsonString, options);
  }

  static equals(a: EnvironmentToken | PlainMessage<EnvironmentToken> | undefined, b: EnvironmentToken | PlainMessage<EnvironmentToken> | undefined): boolean {
    return proto3.util.equals(EnvironmentToken, a, b);
  }
}

