// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/price.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message types.v1.Price
 */
export class Price extends Message<Price> {
  /**
   * @generated from field: google.protobuf.Timestamp created_at = 100;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string stripe_id = 1;
   */
  stripeId = "";

  /**
   * @generated from field: string currency = 2;
   */
  currency = "";

  /**
   * @generated from field: int64 unit_amount = 3;
   */
  unitAmount = protoInt64.zero;

  /**
   * @generated from field: types.v1.Price.Recurring recurring = 4;
   */
  recurring?: Price_Recurring;

  /**
   * @generated from field: string lookup_key = 5;
   */
  lookupKey = "";

  constructor(data?: PartialMessage<Price>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Price";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 100, name: "created_at", kind: "message", T: Timestamp },
    { no: 1, name: "stripe_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "unit_amount", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "recurring", kind: "message", T: Price_Recurring },
    { no: 5, name: "lookup_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Price {
    return new Price().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Price {
    return new Price().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Price {
    return new Price().fromJsonString(jsonString, options);
  }

  static equals(a: Price | PlainMessage<Price> | undefined, b: Price | PlainMessage<Price> | undefined): boolean {
    return proto3.util.equals(Price, a, b);
  }
}

/**
 * @generated from message types.v1.Price.Recurring
 */
export class Price_Recurring extends Message<Price_Recurring> {
  /**
   * @generated from field: string interval = 1;
   */
  interval = "";

  /**
   * @generated from field: int64 interval_count = 2;
   */
  intervalCount = protoInt64.zero;

  /**
   * @generated from field: int64 trial_period_days = 3;
   */
  trialPeriodDays = protoInt64.zero;

  constructor(data?: PartialMessage<Price_Recurring>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Price.Recurring";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "interval", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "interval_count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "trial_period_days", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Price_Recurring {
    return new Price_Recurring().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Price_Recurring {
    return new Price_Recurring().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Price_Recurring {
    return new Price_Recurring().fromJsonString(jsonString, options);
  }

  static equals(a: Price_Recurring | PlainMessage<Price_Recurring> | undefined, b: Price_Recurring | PlainMessage<Price_Recurring> | undefined): boolean {
    return proto3.util.equals(Price_Recurring, a, b);
  }
}
