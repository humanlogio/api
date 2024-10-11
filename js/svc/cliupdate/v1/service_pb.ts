// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file svc/cliupdate/v1/service.proto (package svc.cliupdate.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Version } from "../../../types/v1/version_pb";
import { ReqMeta, ResMeta } from "../../../types/v1/meta_pb";
import { VersionArtifact } from "../../../types/v1/version_artifact_pb";

/**
 * @generated from message svc.cliupdate.v1.GetNextUpdateRequest
 */
export class GetNextUpdateRequest extends Message<GetNextUpdateRequest> {
  /**
   * @generated from field: string project_name = 1;
   */
  projectName = "";

  /**
   * @generated from field: types.v1.Version current_version = 2;
   */
  currentVersion?: Version;

  /**
   * @generated from field: string machine_architecture = 3;
   */
  machineArchitecture = "";

  /**
   * @generated from field: string machine_operating_system = 4;
   */
  machineOperatingSystem = "";

  /**
   * @generated from field: optional string release_channel_name = 5;
   */
  releaseChannelName?: string;

  /**
   * @generated from field: types.v1.ReqMeta meta = 1000;
   */
  meta?: ReqMeta;

  constructor(data?: PartialMessage<GetNextUpdateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.cliupdate.v1.GetNextUpdateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "project_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "current_version", kind: "message", T: Version },
    { no: 3, name: "machine_architecture", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "machine_operating_system", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "release_channel_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 1000, name: "meta", kind: "message", T: ReqMeta },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNextUpdateRequest {
    return new GetNextUpdateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNextUpdateRequest {
    return new GetNextUpdateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNextUpdateRequest {
    return new GetNextUpdateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetNextUpdateRequest | PlainMessage<GetNextUpdateRequest> | undefined, b: GetNextUpdateRequest | PlainMessage<GetNextUpdateRequest> | undefined): boolean {
    return proto3.util.equals(GetNextUpdateRequest, a, b);
  }
}

/**
 * @generated from message svc.cliupdate.v1.GetNextUpdateResponse
 */
export class GetNextUpdateResponse extends Message<GetNextUpdateResponse> {
  /**
   * @generated from field: types.v1.Version next_version = 1;
   */
  nextVersion?: Version;

  /**
   * @generated from field: types.v1.VersionArtifact next_artifact = 2;
   */
  nextArtifact?: VersionArtifact;

  /**
   * @generated from field: types.v1.ResMeta meta = 1000;
   */
  meta?: ResMeta;

  constructor(data?: PartialMessage<GetNextUpdateResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "svc.cliupdate.v1.GetNextUpdateResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "next_version", kind: "message", T: Version },
    { no: 2, name: "next_artifact", kind: "message", T: VersionArtifact },
    { no: 1000, name: "meta", kind: "message", T: ResMeta },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetNextUpdateResponse {
    return new GetNextUpdateResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetNextUpdateResponse {
    return new GetNextUpdateResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetNextUpdateResponse {
    return new GetNextUpdateResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetNextUpdateResponse | PlainMessage<GetNextUpdateResponse> | undefined, b: GetNextUpdateResponse | PlainMessage<GetNextUpdateResponse> | undefined): boolean {
    return proto3.util.equals(GetNextUpdateResponse, a, b);
  }
}
