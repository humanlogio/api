// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/ingest/v1/service.proto (package svc.ingest.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Resource } from "../../../types/v1/otel_resource_pb";
import { Scope } from "../../../types/v1/otel_scope_pb";
import { Log } from "../../../types/v1/otel_logging_pb";

/**
 * @generated from message svc.ingest.v1.IngestRequest
 */
export class IngestRequest extends Message<IngestRequest> {
  /**
   * @generated from field: types.v1.Resource resource = 1;
   */
  resource?: Resource;

  /**
   * @generated from field: types.v1.Scope scope = 2;
   */
  scope?: Scope;

  /**
   * @generated from field: repeated types.v1.Log logs = 3;
   */
  logs: Log[] = [];

  constructor(data?: PartialMessage<IngestRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.ingest.v1.IngestRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "resource", kind: "message", T: Resource },
    { no: 2, name: "scope", kind: "message", T: Scope },
    { no: 3, name: "logs", kind: "message", T: Log, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngestRequest {
    return new IngestRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngestRequest {
    return new IngestRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngestRequest {
    return new IngestRequest().fromJsonString(jsonString, options);
  }

  static equals(a: IngestRequest | PlainMessage<IngestRequest> | undefined, b: IngestRequest | PlainMessage<IngestRequest> | undefined): boolean {
    return proto3.util.equals(IngestRequest, a, b);
  }
}

/**
 * @generated from message svc.ingest.v1.IngestResponse
 */
export class IngestResponse extends Message<IngestResponse> {
  constructor(data?: PartialMessage<IngestResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.ingest.v1.IngestResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngestResponse {
    return new IngestResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngestResponse {
    return new IngestResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngestResponse {
    return new IngestResponse().fromJsonString(jsonString, options);
  }

  static equals(a: IngestResponse | PlainMessage<IngestResponse> | undefined, b: IngestResponse | PlainMessage<IngestResponse> | undefined): boolean {
    return proto3.util.equals(IngestResponse, a, b);
  }
}

/**
 * @generated from message svc.ingest.v1.IngestStreamRequest
 */
export class IngestStreamRequest extends Message<IngestStreamRequest> {
  /**
   * only evaluated on the first message
   *
   * @generated from field: types.v1.Resource resource = 1;
   */
  resource?: Resource;

  /**
   * @generated from field: types.v1.Scope scope = 2;
   */
  scope?: Scope;

  /**
   * @generated from field: repeated types.v1.Log logs = 3;
   */
  logs: Log[] = [];

  constructor(data?: PartialMessage<IngestStreamRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.ingest.v1.IngestStreamRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "resource", kind: "message", T: Resource },
    { no: 2, name: "scope", kind: "message", T: Scope },
    { no: 3, name: "logs", kind: "message", T: Log, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngestStreamRequest {
    return new IngestStreamRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngestStreamRequest {
    return new IngestStreamRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngestStreamRequest {
    return new IngestStreamRequest().fromJsonString(jsonString, options);
  }

  static equals(a: IngestStreamRequest | PlainMessage<IngestStreamRequest> | undefined, b: IngestStreamRequest | PlainMessage<IngestStreamRequest> | undefined): boolean {
    return proto3.util.equals(IngestStreamRequest, a, b);
  }
}

/**
 * @generated from message svc.ingest.v1.IngestStreamResponse
 */
export class IngestStreamResponse extends Message<IngestStreamResponse> {
  constructor(data?: PartialMessage<IngestStreamResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.ingest.v1.IngestStreamResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IngestStreamResponse {
    return new IngestStreamResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IngestStreamResponse {
    return new IngestStreamResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IngestStreamResponse {
    return new IngestStreamResponse().fromJsonString(jsonString, options);
  }

  static equals(a: IngestStreamResponse | PlainMessage<IngestStreamResponse> | undefined, b: IngestStreamResponse | PlainMessage<IngestStreamResponse> | undefined): boolean {
    return proto3.util.equals(IngestStreamResponse, a, b);
  }
}

