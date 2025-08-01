// @generated by protoc-gen-es v1.10.1 with parameter "target=ts,import_extension=none"
// @generated from file types/v1/project.proto (package types.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message types.v1.ProjectPointer
 */
export class ProjectPointer extends Message<ProjectPointer> {
  /**
   * @generated from oneof types.v1.ProjectPointer.scheme
   */
  scheme: {
    /**
     * @generated from field: types.v1.ProjectPointer.RemoteGit remote = 1;
     */
    value: ProjectPointer_RemoteGit;
    case: "remote";
  } | {
    /**
     * @generated from field: types.v1.ProjectPointer.LocalGit localhost = 2;
     */
    value: ProjectPointer_LocalGit;
    case: "localhost";
  } | {
    /**
     * @generated from field: types.v1.ProjectPointer.Virtual db = 3;
     */
    value: ProjectPointer_Virtual;
    case: "db";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<ProjectPointer>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ProjectPointer";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "remote", kind: "message", T: ProjectPointer_RemoteGit, oneof: "scheme" },
    { no: 2, name: "localhost", kind: "message", T: ProjectPointer_LocalGit, oneof: "scheme" },
    { no: 3, name: "db", kind: "message", T: ProjectPointer_Virtual, oneof: "scheme" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectPointer {
    return new ProjectPointer().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectPointer {
    return new ProjectPointer().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectPointer {
    return new ProjectPointer().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectPointer | PlainMessage<ProjectPointer> | undefined, b: ProjectPointer | PlainMessage<ProjectPointer> | undefined): boolean {
    return proto3.util.equals(ProjectPointer, a, b);
  }
}

/**
 * @generated from message types.v1.ProjectPointer.RemoteGit
 */
export class ProjectPointer_RemoteGit extends Message<ProjectPointer_RemoteGit> {
  /**
   * @generated from field: string remote_url = 1;
   */
  remoteUrl = "";

  /**
   * @generated from field: string ref = 2;
   */
  ref = "";

  /**
   * @generated from field: string dashboard_dir = 3;
   */
  dashboardDir = "";

  /**
   * @generated from field: string alert_dir = 4;
   */
  alertDir = "";

  constructor(data?: PartialMessage<ProjectPointer_RemoteGit>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ProjectPointer.RemoteGit";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "remote_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "ref", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "dashboard_dir", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "alert_dir", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectPointer_RemoteGit {
    return new ProjectPointer_RemoteGit().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectPointer_RemoteGit {
    return new ProjectPointer_RemoteGit().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectPointer_RemoteGit {
    return new ProjectPointer_RemoteGit().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectPointer_RemoteGit | PlainMessage<ProjectPointer_RemoteGit> | undefined, b: ProjectPointer_RemoteGit | PlainMessage<ProjectPointer_RemoteGit> | undefined): boolean {
    return proto3.util.equals(ProjectPointer_RemoteGit, a, b);
  }
}

/**
 * @generated from message types.v1.ProjectPointer.LocalGit
 */
export class ProjectPointer_LocalGit extends Message<ProjectPointer_LocalGit> {
  /**
   * @generated from field: string path = 1;
   */
  path = "";

  /**
   * @generated from field: string dashboard_dir = 2;
   */
  dashboardDir = "";

  /**
   * @generated from field: string alert_dir = 3;
   */
  alertDir = "";

  /**
   * @generated from field: bool read_only = 4;
   */
  readOnly = false;

  constructor(data?: PartialMessage<ProjectPointer_LocalGit>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ProjectPointer.LocalGit";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "dashboard_dir", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "alert_dir", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "read_only", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectPointer_LocalGit {
    return new ProjectPointer_LocalGit().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectPointer_LocalGit {
    return new ProjectPointer_LocalGit().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectPointer_LocalGit {
    return new ProjectPointer_LocalGit().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectPointer_LocalGit | PlainMessage<ProjectPointer_LocalGit> | undefined, b: ProjectPointer_LocalGit | PlainMessage<ProjectPointer_LocalGit> | undefined): boolean {
    return proto3.util.equals(ProjectPointer_LocalGit, a, b);
  }
}

/**
 * @generated from message types.v1.ProjectPointer.Virtual
 */
export class ProjectPointer_Virtual extends Message<ProjectPointer_Virtual> {
  /**
   * @generated from field: string uri = 1;
   */
  uri = "";

  constructor(data?: PartialMessage<ProjectPointer_Virtual>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.ProjectPointer.Virtual";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "uri", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProjectPointer_Virtual {
    return new ProjectPointer_Virtual().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProjectPointer_Virtual {
    return new ProjectPointer_Virtual().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProjectPointer_Virtual {
    return new ProjectPointer_Virtual().fromJsonString(jsonString, options);
  }

  static equals(a: ProjectPointer_Virtual | PlainMessage<ProjectPointer_Virtual> | undefined, b: ProjectPointer_Virtual | PlainMessage<ProjectPointer_Virtual> | undefined): boolean {
    return proto3.util.equals(ProjectPointer_Virtual, a, b);
  }
}

/**
 * @generated from message types.v1.Project
 */
export class Project extends Message<Project> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: types.v1.ProjectPointer pointer = 2;
   */
  pointer?: ProjectPointer;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 300;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 301;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Project>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "types.v1.Project";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "pointer", kind: "message", T: ProjectPointer },
    { no: 300, name: "created_at", kind: "message", T: Timestamp },
    { no: 301, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Project {
    return new Project().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Project {
    return new Project().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Project {
    return new Project().fromJsonString(jsonString, options);
  }

  static equals(a: Project | PlainMessage<Project> | undefined, b: Project | PlainMessage<Project> | undefined): boolean {
    return proto3.util.equals(Project, a, b);
  }
}

