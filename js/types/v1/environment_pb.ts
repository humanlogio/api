// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/environment.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * @generated from message types.v1.Environment
 */
export class Environment extends Message<Environment> {
  /**
   * @generated from field: int64 id = 1;
   */
  id = protoInt64.zero;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<Environment>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Environment";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Environment {
    return new Environment().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Environment {
    return new Environment().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Environment {
    return new Environment().fromJsonString(jsonString, options);
  }

  static equals(a: Environment | PlainMessage<Environment> | undefined, b: Environment | PlainMessage<Environment> | undefined): boolean {
    return proto3.util.equals(Environment, a, b);
  }
}
