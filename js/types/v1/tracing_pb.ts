// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/tracing.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { KV } from "./types_pb";

/**
 * @generated from message types.v1.Span
 */
export class Span extends Message<Span> {
  /**
   * @generated from field: bytes trace_id = 1;
   */
  traceId = new Uint8Array(0);

  /**
   * @generated from field: bytes span_id = 2;
   */
  spanId = new Uint8Array(0);

  /**
   * @generated from field: string trace_state = 3;
   */
  traceState = "";

  /**
   * @generated from field: bytes parent_span_id = 4;
   */
  parentSpanId = new Uint8Array(0);

  /**
   * @generated from field: uint32 flags = 5;
   */
  flags = 0;

  /**
   * @generated from field: string name = 6;
   */
  name = "";

  /**
   * @generated from field: types.v1.Span.SpanKind kind = 7;
   */
  kind = Span_SpanKind.UNSPECIFIED;

  /**
   * @generated from field: google.protobuf.Timestamp start_time = 8;
   */
  startTime?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end_time = 9;
   */
  endTime?: Timestamp;

  /**
   * @generated from field: repeated types.v1.KV kvs = 10;
   */
  kvs: KV[] = [];

  /**
   * @generated from field: repeated types.v1.Span.Event events = 11;
   */
  events: Span_Event[] = [];

  /**
   * @generated from field: repeated types.v1.Span.Link links = 12;
   */
  links: Span_Link[] = [];

  /**
   * @generated from field: types.v1.Span.Status status = 13;
   */
  status?: Span_Status;

  constructor(data?: PartialMessage<Span>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "trace_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "span_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "trace_state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "parent_span_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 5, name: "flags", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 6, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "kind", kind: "enum", T: proto3.getEnumType(Span_SpanKind) },
    { no: 8, name: "start_time", kind: "message", T: Timestamp },
    { no: 9, name: "end_time", kind: "message", T: Timestamp },
    { no: 10, name: "kvs", kind: "message", T: KV, repeated: true },
    { no: 11, name: "events", kind: "message", T: Span_Event, repeated: true },
    { no: 12, name: "links", kind: "message", T: Span_Link, repeated: true },
    { no: 13, name: "status", kind: "message", T: Span_Status },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Span {
    return new Span().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Span {
    return new Span().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Span {
    return new Span().fromJsonString(jsonString, options);
  }

  static equals(a: Span | PlainMessage<Span> | undefined, b: Span | PlainMessage<Span> | undefined): boolean {
    return proto3.util.equals(Span, a, b);
  }
}

/**
 * @generated from enum types.v1.Span.SpanKind
 */
export enum Span_SpanKind {
  /**
   * @generated from enum value: UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: INTERNAL = 1;
   */
  INTERNAL = 1,

  /**
   * @generated from enum value: SERVER = 2;
   */
  SERVER = 2,

  /**
   * @generated from enum value: CLIENT = 3;
   */
  CLIENT = 3,

  /**
   * @generated from enum value: PRODUCER = 4;
   */
  PRODUCER = 4,

  /**
   * @generated from enum value: CONSUMER = 5;
   */
  CONSUMER = 5,
}
// Retrieve enum metadata with: proto3.getEnumType(Span_SpanKind)
proto3.util.setEnumType(Span_SpanKind, "types.v1.Span.SpanKind", [
  { no: 0, name: "UNSPECIFIED" },
  { no: 1, name: "INTERNAL" },
  { no: 2, name: "SERVER" },
  { no: 3, name: "CLIENT" },
  { no: 4, name: "PRODUCER" },
  { no: 5, name: "CONSUMER" },
]);

/**
 * @generated from message types.v1.Span.Event
 */
export class Span_Event extends Message<Span_Event> {
  /**
   * @generated from field: google.protobuf.Timestamp timestamp = 1;
   */
  timestamp?: Timestamp;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: repeated types.v1.KV kvs = 3;
   */
  kvs: KV[] = [];

  constructor(data?: PartialMessage<Span_Event>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span.Event";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "timestamp", kind: "message", T: Timestamp },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "kvs", kind: "message", T: KV, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Span_Event {
    return new Span_Event().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Span_Event {
    return new Span_Event().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Span_Event {
    return new Span_Event().fromJsonString(jsonString, options);
  }

  static equals(a: Span_Event | PlainMessage<Span_Event> | undefined, b: Span_Event | PlainMessage<Span_Event> | undefined): boolean {
    return proto3.util.equals(Span_Event, a, b);
  }
}

/**
 * @generated from message types.v1.Span.Link
 */
export class Span_Link extends Message<Span_Link> {
  /**
   * @generated from field: bytes trace_id = 1;
   */
  traceId = new Uint8Array(0);

  /**
   * @generated from field: bytes span_id = 2;
   */
  spanId = new Uint8Array(0);

  /**
   * @generated from field: string trace_state = 3;
   */
  traceState = "";

  /**
   * @generated from field: repeated types.v1.KV kvs = 4;
   */
  kvs: KV[] = [];

  /**
   * @generated from field: uint32 flags = 5;
   */
  flags = 0;

  constructor(data?: PartialMessage<Span_Link>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span.Link";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "trace_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "span_id", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "trace_state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "kvs", kind: "message", T: KV, repeated: true },
    { no: 5, name: "flags", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Span_Link {
    return new Span_Link().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Span_Link {
    return new Span_Link().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Span_Link {
    return new Span_Link().fromJsonString(jsonString, options);
  }

  static equals(a: Span_Link | PlainMessage<Span_Link> | undefined, b: Span_Link | PlainMessage<Span_Link> | undefined): boolean {
    return proto3.util.equals(Span_Link, a, b);
  }
}

/**
 * @generated from message types.v1.Span.Status
 */
export class Span_Status extends Message<Span_Status> {
  /**
   * @generated from field: string message = 1;
   */
  message = "";

  /**
   * @generated from field: int32 code = 2;
   */
  code = 0;

  constructor(data?: PartialMessage<Span_Status>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span.Status";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "code", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Span_Status {
    return new Span_Status().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Span_Status {
    return new Span_Status().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Span_Status {
    return new Span_Status().fromJsonString(jsonString, options);
  }

  static equals(a: Span_Status | PlainMessage<Span_Status> | undefined, b: Span_Status | PlainMessage<Span_Status> | undefined): boolean {
    return proto3.util.equals(Span_Status, a, b);
  }
}

