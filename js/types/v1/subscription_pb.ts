// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/subscription.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { PaymentMethod } from "./payment_method_pb";

/**
 * @generated from message types.v1.Subscription
 */
export class Subscription extends Message<Subscription> {
  /**
   * @generated from field: google.protobuf.Timestamp created_at = 100;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string stripe_id = 1;
   */
  stripeId = "";

  /**
   * @generated from field: types.v1.PaymentMethod default_payment_method = 2;
   */
  defaultPaymentMethod?: PaymentMethod;

  /**
   * @generated from field: types.v1.Subscription.Status status = 3;
   */
  status = Subscription_Status.Unknown;

  /**
   * @generated from field: google.protobuf.Timestamp current_period_start_at = 201;
   */
  currentPeriodStartAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp current_period_end_at = 202;
   */
  currentPeriodEndAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp cancel_at = 301;
   */
  cancelAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp canceled_at = 302;
   */
  canceledAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp ended_at = 401;
   */
  endedAt?: Timestamp;

  constructor(data?: PartialMessage<Subscription>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Subscription";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 100, name: "created_at", kind: "message", T: Timestamp },
    { no: 1, name: "stripe_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "default_payment_method", kind: "message", T: PaymentMethod },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(Subscription_Status) },
    { no: 201, name: "current_period_start_at", kind: "message", T: Timestamp },
    { no: 202, name: "current_period_end_at", kind: "message", T: Timestamp },
    { no: 301, name: "cancel_at", kind: "message", T: Timestamp },
    { no: 302, name: "canceled_at", kind: "message", T: Timestamp },
    { no: 401, name: "ended_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Subscription {
    return new Subscription().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Subscription {
    return new Subscription().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Subscription {
    return new Subscription().fromJsonString(jsonString, options);
  }

  static equals(a: Subscription | PlainMessage<Subscription> | undefined, b: Subscription | PlainMessage<Subscription> | undefined): boolean {
    return proto3.util.equals(Subscription, a, b);
  }
}

/**
 * @generated from enum types.v1.Subscription.Status
 */
export enum Subscription_Status {
  /**
   * @generated from enum value: Unknown = 0;
   */
  Unknown = 0,

  /**
   * @generated from enum value: Active = 1;
   */
  Active = 1,

  /**
   * @generated from enum value: Canceled = 2;
   */
  Canceled = 2,

  /**
   * @generated from enum value: Incomplete = 3;
   */
  Incomplete = 3,

  /**
   * @generated from enum value: IncompleteExpired = 4;
   */
  IncompleteExpired = 4,

  /**
   * @generated from enum value: PastDue = 5;
   */
  PastDue = 5,

  /**
   * @generated from enum value: Paused = 6;
   */
  Paused = 6,

  /**
   * @generated from enum value: Trialing = 7;
   */
  Trialing = 7,

  /**
   * @generated from enum value: Unpaid = 8;
   */
  Unpaid = 8,
}
// Retrieve enum metadata with: proto3.getEnumType(Subscription_Status)
proto3.util.setEnumType(Subscription_Status, "types.v1.Subscription.Status", [
  { no: 0, name: "Unknown" },
  { no: 1, name: "Active" },
  { no: 2, name: "Canceled" },
  { no: 3, name: "Incomplete" },
  { no: 4, name: "IncompleteExpired" },
  { no: 5, name: "PastDue" },
  { no: 6, name: "Paused" },
  { no: 7, name: "Trialing" },
  { no: 8, name: "Unpaid" },
]);

