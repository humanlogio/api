// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/data.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Table } from "./types_pb";
import { Query } from "./query_pb";
import { Log } from "./otel_logging_pb";
import { Span } from "./otel_tracing_pb";

/**
 * @generated from message types.v1.Data
 */
export class Data extends Message<Data> {
  /**
   * @generated from oneof types.v1.Data.shape
   */
  shape: {
    /**
     * @generated from field: types.v1.Subqueries subqueries = 100;
     */
    value: Subqueries;
    case: "subqueries";
  } | {
    /**
     * @generated from field: types.v1.Table free_form = 101;
     */
    value: Table;
    case: "freeForm";
  } | {
    /**
     * @generated from field: types.v1.Logs logs = 102;
     */
    value: Logs;
    case: "logs";
  } | {
    /**
     * @generated from field: types.v1.Spans spans = 103;
     */
    value: Spans;
    case: "spans";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Data>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Data";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 100, name: "subqueries", kind: "message", T: Subqueries, oneof: "shape" },
    { no: 101, name: "free_form", kind: "message", T: Table, oneof: "shape" },
    { no: 102, name: "logs", kind: "message", T: Logs, oneof: "shape" },
    { no: 103, name: "spans", kind: "message", T: Spans, oneof: "shape" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Data {
    return new Data().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Data {
    return new Data().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Data {
    return new Data().fromJsonString(jsonString, options);
  }

  static equals(a: Data | PlainMessage<Data> | undefined, b: Data | PlainMessage<Data> | undefined): boolean {
    return proto3.util.equals(Data, a, b);
  }
}

/**
 * @generated from message types.v1.Subqueries
 */
export class Subqueries extends Message<Subqueries> {
  /**
   * @generated from field: repeated types.v1.Query queries = 1;
   */
  queries: Query[] = [];

  constructor(data?: PartialMessage<Subqueries>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Subqueries";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "queries", kind: "message", T: Query, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Subqueries {
    return new Subqueries().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Subqueries {
    return new Subqueries().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Subqueries {
    return new Subqueries().fromJsonString(jsonString, options);
  }

  static equals(a: Subqueries | PlainMessage<Subqueries> | undefined, b: Subqueries | PlainMessage<Subqueries> | undefined): boolean {
    return proto3.util.equals(Subqueries, a, b);
  }
}

/**
 * @generated from message types.v1.Logs
 */
export class Logs extends Message<Logs> {
  /**
   * @generated from field: repeated types.v1.Log logs = 1;
   */
  logs: Log[] = [];

  constructor(data?: PartialMessage<Logs>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Logs";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "logs", kind: "message", T: Log, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Logs {
    return new Logs().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Logs {
    return new Logs().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Logs {
    return new Logs().fromJsonString(jsonString, options);
  }

  static equals(a: Logs | PlainMessage<Logs> | undefined, b: Logs | PlainMessage<Logs> | undefined): boolean {
    return proto3.util.equals(Logs, a, b);
  }
}

/**
 * @generated from message types.v1.Spans
 */
export class Spans extends Message<Spans> {
  /**
   * @generated from field: repeated types.v1.Span spans = 1;
   */
  spans: Span[] = [];

  constructor(data?: PartialMessage<Spans>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Spans";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "spans", kind: "message", T: Span, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Spans {
    return new Spans().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Spans {
    return new Spans().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Spans {
    return new Spans().fromJsonString(jsonString, options);
  }

  static equals(a: Spans | PlainMessage<Spans> | undefined, b: Spans | PlainMessage<Spans> | undefined): boolean {
    return proto3.util.equals(Spans, a, b);
  }
}

