// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/release/v1/service.proto (package svc.release.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { CreateReleaseChannelRequest, CreateReleaseChannelResponse, CreateVersionArtifactRequest, CreateVersionArtifactResponse, DeleteVersionArtifactRequest, DeleteVersionArtifactResponse, ListReleaseChannelRequest, ListReleaseChannelResponse, ListVersionArtifactRequest, ListVersionArtifactResponse, PublishVersionRequest, PublishVersionResponse, UnpublishVersionRequest, UnpublishVersionResponse } from "./service_pb";

/**
 * @generated from rpc svc.release.v1.ReleaseService.CreateReleaseChannel
 */
export const createReleaseChannel = {
  localName: "createReleaseChannel",
  name: "CreateReleaseChannel",
  kind: MethodKind.Unary,
  I: CreateReleaseChannelRequest,
  O: CreateReleaseChannelResponse,
  service: {
    typeName: "svc.release.v1.ReleaseService"
  }
} as const;

/**
 * @generated from rpc svc.release.v1.ReleaseService.ListReleaseChannel
 */
export const listReleaseChannel = {
  localName: "listReleaseChannel",
  name: "ListReleaseChannel",
  kind: MethodKind.Unary,
  I: ListReleaseChannelRequest,
  O: ListReleaseChannelResponse,
  service: {
    typeName: "svc.release.v1.ReleaseService"
  }
} as const;

/**
 * @generated from rpc svc.release.v1.ReleaseService.PublishVersion
 */
export const publishVersion = {
  localName: "publishVersion",
  name: "PublishVersion",
  kind: MethodKind.Unary,
  I: PublishVersionRequest,
  O: PublishVersionResponse,
  service: {
    typeName: "svc.release.v1.ReleaseService"
  }
} as const;

/**
 * @generated from rpc svc.release.v1.ReleaseService.UnpublishVersion
 */
export const unpublishVersion = {
  localName: "unpublishVersion",
  name: "UnpublishVersion",
  kind: MethodKind.Unary,
  I: UnpublishVersionRequest,
  O: UnpublishVersionResponse,
  service: {
    typeName: "svc.release.v1.ReleaseService"
  }
} as const;

/**
 * @generated from rpc svc.release.v1.ReleaseService.CreateVersionArtifact
 */
export const createVersionArtifact = {
  localName: "createVersionArtifact",
  name: "CreateVersionArtifact",
  kind: MethodKind.Unary,
  I: CreateVersionArtifactRequest,
  O: CreateVersionArtifactResponse,
  service: {
    typeName: "svc.release.v1.ReleaseService"
  }
} as const;

/**
 * @generated from rpc svc.release.v1.ReleaseService.DeleteVersionArtifact
 */
export const deleteVersionArtifact = {
  localName: "deleteVersionArtifact",
  name: "DeleteVersionArtifact",
  kind: MethodKind.Unary,
  I: DeleteVersionArtifactRequest,
  O: DeleteVersionArtifactResponse,
  service: {
    typeName: "svc.release.v1.ReleaseService"
  }
} as const;

/**
 * @generated from rpc svc.release.v1.ReleaseService.ListVersionArtifact
 */
export const listVersionArtifact = {
  localName: "listVersionArtifact",
  name: "ListVersionArtifact",
  kind: MethodKind.Unary,
  I: ListVersionArtifactRequest,
  O: ListVersionArtifactResponse,
  service: {
    typeName: "svc.release.v1.ReleaseService"
  }
} as const;
