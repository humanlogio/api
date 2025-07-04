// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/tracing.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Duration, Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { ULID } from "./ulid_pb";
import { KV } from "./types_pb";

/**
 * @generated from message types.v1.Trace
 */
export class Trace extends Message<Trace> {
  /**
   * @generated from field: string trace_id = 1;
   */
  traceId = "";

  /**
   * @generated from field: repeated types.v1.Span spans = 2;
   */
  spans: Span[] = [];

  constructor(data?: PartialMessage<Trace>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Trace";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "trace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "spans", kind: "message", T: Span, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Trace {
    return new Trace().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Trace {
    return new Trace().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Trace {
    return new Trace().fromJsonString(jsonString, options);
  }

  static equals(a: Trace | PlainMessage<Trace> | undefined, b: Trace | PlainMessage<Trace> | undefined): boolean {
    return proto3.util.equals(Trace, a, b);
  }
}

/**
 * @generated from message types.v1.Span
 */
export class Span extends Message<Span> {
  /**
   * @generated from field: types.v1.ULID ulid = 100;
   */
  ulid?: ULID;

  /**
   * @generated from field: string trace_id = 1;
   */
  traceId = "";

  /**
   * @generated from field: string span_id = 2;
   */
  spanId = "";

  /**
   * @generated from field: string trace_state = 3;
   */
  traceState = "";

  /**
   * @generated from field: string parent_span_id = 4;
   */
  parentSpanId = "";

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
   * @generated from field: string service_name = 8;
   */
  serviceName = "";

  /**
   * @generated from field: types.v1.Span.Scope scope = 9;
   */
  scope?: Span_Scope;

  /**
   * @generated from field: types.v1.Span.Timing timing = 10;
   */
  timing?: Span_Timing;

  /**
   * @generated from field: repeated types.v1.KV resource_attributes = 11;
   */
  resourceAttributes: KV[] = [];

  /**
   * @generated from field: repeated types.v1.KV span_attributes = 12;
   */
  spanAttributes: KV[] = [];

  /**
   * @generated from field: repeated types.v1.Span.Event events = 13;
   */
  events: Span_Event[] = [];

  /**
   * @generated from field: repeated types.v1.Span.Link links = 14;
   */
  links: Span_Link[] = [];

  /**
   * @generated from field: types.v1.Span.Status status = 15;
   */
  status?: Span_Status;

  constructor(data?: PartialMessage<Span>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 100, name: "ulid", kind: "message", T: ULID },
    { no: 1, name: "trace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "span_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "trace_state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "parent_span_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "flags", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 6, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "kind", kind: "enum", T: proto3.getEnumType(Span_SpanKind) },
    { no: 8, name: "service_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "scope", kind: "message", T: Span_Scope },
    { no: 10, name: "timing", kind: "message", T: Span_Timing },
    { no: 11, name: "resource_attributes", kind: "message", T: KV, repeated: true },
    { no: 12, name: "span_attributes", kind: "message", T: KV, repeated: true },
    { no: 13, name: "events", kind: "message", T: Span_Event, repeated: true },
    { no: 14, name: "links", kind: "message", T: Span_Link, repeated: true },
    { no: 15, name: "status", kind: "message", T: Span_Status },
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
   * @generated from field: string trace_id = 1;
   */
  traceId = "";

  /**
   * @generated from field: string span_id = 2;
   */
  spanId = "";

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
    { no: 1, name: "trace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "span_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
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
   * @generated from field: types.v1.Span.Status.Code code = 2;
   */
  code = Span_Status_Code.UNSET;

  constructor(data?: PartialMessage<Span_Status>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span.Status";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "code", kind: "enum", T: proto3.getEnumType(Span_Status_Code) },
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

/**
 * @generated from enum types.v1.Span.Status.Code
 */
export enum Span_Status_Code {
  /**
   * @generated from enum value: UNSET = 0;
   */
  UNSET = 0,

  /**
   * @generated from enum value: OK = 1;
   */
  OK = 1,

  /**
   * @generated from enum value: ERROR = 2;
   */
  ERROR = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(Span_Status_Code)
proto3.util.setEnumType(Span_Status_Code, "types.v1.Span.Status.Code", [
  { no: 0, name: "UNSET" },
  { no: 1, name: "OK" },
  { no: 2, name: "ERROR" },
]);

/**
 * @generated from message types.v1.Span.Scope
 */
export class Span_Scope extends Message<Span_Scope> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: string version = 2;
   */
  version = "";

  constructor(data?: PartialMessage<Span_Scope>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span.Scope";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Span_Scope {
    return new Span_Scope().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Span_Scope {
    return new Span_Scope().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Span_Scope {
    return new Span_Scope().fromJsonString(jsonString, options);
  }

  static equals(a: Span_Scope | PlainMessage<Span_Scope> | undefined, b: Span_Scope | PlainMessage<Span_Scope> | undefined): boolean {
    return proto3.util.equals(Span_Scope, a, b);
  }
}

/**
 * @generated from message types.v1.Span.Timing
 */
export class Span_Timing extends Message<Span_Timing> {
  /**
   * @generated from field: google.protobuf.Timestamp start = 1;
   */
  start?: Timestamp;

  /**
   * @generated from field: google.protobuf.Duration duration = 2;
   */
  duration?: Duration;

  constructor(data?: PartialMessage<Span_Timing>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Span.Timing";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "start", kind: "message", T: Timestamp },
    { no: 2, name: "duration", kind: "message", T: Duration },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Span_Timing {
    return new Span_Timing().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Span_Timing {
    return new Span_Timing().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Span_Timing {
    return new Span_Timing().fromJsonString(jsonString, options);
  }

  static equals(a: Span_Timing | PlainMessage<Span_Timing> | undefined, b: Span_Timing | PlainMessage<Span_Timing> | undefined): boolean {
    return proto3.util.equals(Span_Timing, a, b);
  }
}

