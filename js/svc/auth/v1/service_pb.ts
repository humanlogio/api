// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/auth/v1/service.proto (package svc.auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Version } from "../../../types/v1/version_pb";
import { UserToken } from "../../../types/v1/user_token_pb";

/**
 * @generated from message svc.auth.v1.GetAuthURLRequest
 */
export class GetAuthURLRequest extends Message<GetAuthURLRequest> {
  /**
   * optional: if an org is specified
   *
   * @generated from oneof svc.auth.v1.GetAuthURLRequest.organization
   */
  organization: {
    /**
     * @generated from field: int64 by_id = 100;
     */
    value: bigint;
    case: "byId";
  } | {
    /**
     * @generated from field: string by_name = 101;
     */
    value: string;
    case: "byName";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: string return_to_url = 2;
   */
  returnToUrl = "";

  /**
   * @generated from field: svc.auth.v1.LocalhostViaBrowser localhost = 3;
   */
  localhost?: LocalhostViaBrowser;

  constructor(data?: PartialMessage<GetAuthURLRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.auth.v1.GetAuthURLRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 100, name: "by_id", kind: "scalar", T: 3 /* ScalarType.INT64 */, oneof: "organization" },
    { no: 101, name: "by_name", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "organization" },
    { no: 2, name: "return_to_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "localhost", kind: "message", T: LocalhostViaBrowser },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAuthURLRequest {
    return new GetAuthURLRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAuthURLRequest {
    return new GetAuthURLRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAuthURLRequest {
    return new GetAuthURLRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetAuthURLRequest | PlainMessage<GetAuthURLRequest> | undefined, b: GetAuthURLRequest | PlainMessage<GetAuthURLRequest> | undefined): boolean {
    return proto3.util.equals(GetAuthURLRequest, a, b);
  }
}

/**
 * @generated from message svc.auth.v1.LocalhostViaBrowser
 */
export class LocalhostViaBrowser extends Message<LocalhostViaBrowser> {
  /**
   * @generated from field: string architecture = 1;
   */
  architecture = "";

  /**
   * @generated from field: string operating_system = 2;
   */
  operatingSystem = "";

  /**
   * @generated from field: types.v1.Version using_version = 3;
   */
  usingVersion?: Version;

  constructor(data?: PartialMessage<LocalhostViaBrowser>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.auth.v1.LocalhostViaBrowser";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "architecture", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "operating_system", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "using_version", kind: "message", T: Version },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LocalhostViaBrowser {
    return new LocalhostViaBrowser().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LocalhostViaBrowser {
    return new LocalhostViaBrowser().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LocalhostViaBrowser {
    return new LocalhostViaBrowser().fromJsonString(jsonString, options);
  }

  static equals(a: LocalhostViaBrowser | PlainMessage<LocalhostViaBrowser> | undefined, b: LocalhostViaBrowser | PlainMessage<LocalhostViaBrowser> | undefined): boolean {
    return proto3.util.equals(LocalhostViaBrowser, a, b);
  }
}

/**
 * @generated from message svc.auth.v1.GetAuthURLResponse
 */
export class GetAuthURLResponse extends Message<GetAuthURLResponse> {
  /**
   * @generated from field: string auth_url = 1;
   */
  authUrl = "";

  constructor(data?: PartialMessage<GetAuthURLResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.auth.v1.GetAuthURLResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "auth_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetAuthURLResponse {
    return new GetAuthURLResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetAuthURLResponse {
    return new GetAuthURLResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetAuthURLResponse {
    return new GetAuthURLResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetAuthURLResponse | PlainMessage<GetAuthURLResponse> | undefined, b: GetAuthURLResponse | PlainMessage<GetAuthURLResponse> | undefined): boolean {
    return proto3.util.equals(GetAuthURLResponse, a, b);
  }
}

/**
 * @generated from message svc.auth.v1.BeginDeviceAuthRequest
 */
export class BeginDeviceAuthRequest extends Message<BeginDeviceAuthRequest> {
  /**
   * @generated from field: string organization = 1;
   */
  organization = "";

  /**
   * @generated from field: string return_to_url = 2;
   */
  returnToUrl = "";

  constructor(data?: PartialMessage<BeginDeviceAuthRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.auth.v1.BeginDeviceAuthRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "organization", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "return_to_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginDeviceAuthRequest {
    return new BeginDeviceAuthRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginDeviceAuthRequest {
    return new BeginDeviceAuthRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginDeviceAuthRequest {
    return new BeginDeviceAuthRequest().fromJsonString(jsonString, options);
  }

  static equals(a: BeginDeviceAuthRequest | PlainMessage<BeginDeviceAuthRequest> | undefined, b: BeginDeviceAuthRequest | PlainMessage<BeginDeviceAuthRequest> | undefined): boolean {
    return proto3.util.equals(BeginDeviceAuthRequest, a, b);
  }
}

/**
 * @generated from message svc.auth.v1.BeginDeviceAuthResponse
 */
export class BeginDeviceAuthResponse extends Message<BeginDeviceAuthResponse> {
  /**
   * @generated from field: string url = 1;
   */
  url = "";

  /**
   * device_code must be used to retrieve an authenticated user token with `CompleteDeviceAuth`
   *
   * @generated from field: string device_code = 2;
   */
  deviceCode = "";

  /**
   * @generated from field: string user_code = 3;
   */
  userCode = "";

  /**
   * @generated from field: google.protobuf.Timestamp expires_at = 4;
   */
  expiresAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Duration poll_interval = 5;
   */
  pollInterval?: Duration;

  constructor(data?: PartialMessage<BeginDeviceAuthResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.auth.v1.BeginDeviceAuthResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "device_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "user_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "expires_at", kind: "message", T: Timestamp },
    { no: 5, name: "poll_interval", kind: "message", T: Duration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BeginDeviceAuthResponse {
    return new BeginDeviceAuthResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BeginDeviceAuthResponse {
    return new BeginDeviceAuthResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BeginDeviceAuthResponse {
    return new BeginDeviceAuthResponse().fromJsonString(jsonString, options);
  }

  static equals(a: BeginDeviceAuthResponse | PlainMessage<BeginDeviceAuthResponse> | undefined, b: BeginDeviceAuthResponse | PlainMessage<BeginDeviceAuthResponse> | undefined): boolean {
    return proto3.util.equals(BeginDeviceAuthResponse, a, b);
  }
}

/**
 * @generated from message svc.auth.v1.CompleteDeviceAuthRequest
 */
export class CompleteDeviceAuthRequest extends Message<CompleteDeviceAuthRequest> {
  /**
   * device_code is returned when calling `BeginDeviceAuth` and allows an auth flow
   * to retrieve the user token that is generated when a user is authenticated
   *
   * @generated from field: string device_code = 1;
   */
  deviceCode = "";

  /**
   * @generated from field: string user_code = 2;
   */
  userCode = "";

  /**
   * @generated from field: string architecture = 5;
   */
  architecture = "";

  /**
   * @generated from field: string operating_system = 6;
   */
  operatingSystem = "";

  constructor(data?: PartialMessage<CompleteDeviceAuthRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.auth.v1.CompleteDeviceAuthRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "device_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "architecture", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "operating_system", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompleteDeviceAuthRequest {
    return new CompleteDeviceAuthRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompleteDeviceAuthRequest {
    return new CompleteDeviceAuthRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompleteDeviceAuthRequest {
    return new CompleteDeviceAuthRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CompleteDeviceAuthRequest | PlainMessage<CompleteDeviceAuthRequest> | undefined, b: CompleteDeviceAuthRequest | PlainMessage<CompleteDeviceAuthRequest> | undefined): boolean {
    return proto3.util.equals(CompleteDeviceAuthRequest, a, b);
  }
}

/**
 * @generated from message svc.auth.v1.CompleteDeviceAuthResponse
 */
export class CompleteDeviceAuthResponse extends Message<CompleteDeviceAuthResponse> {
  /**
   * @generated from field: types.v1.UserToken token = 1;
   */
  token?: UserToken;

  constructor(data?: PartialMessage<CompleteDeviceAuthResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.auth.v1.CompleteDeviceAuthResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "token", kind: "message", T: UserToken },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CompleteDeviceAuthResponse {
    return new CompleteDeviceAuthResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CompleteDeviceAuthResponse {
    return new CompleteDeviceAuthResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CompleteDeviceAuthResponse {
    return new CompleteDeviceAuthResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CompleteDeviceAuthResponse | PlainMessage<CompleteDeviceAuthResponse> | undefined, b: CompleteDeviceAuthResponse | PlainMessage<CompleteDeviceAuthResponse> | undefined): boolean {
    return proto3.util.equals(CompleteDeviceAuthResponse, a, b);
  }
}

