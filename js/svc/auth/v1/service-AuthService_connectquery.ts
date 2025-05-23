// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/auth/v1/service.proto (package svc.auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { BeginDeviceAuthRequest, BeginDeviceAuthResponse, CompleteDeviceAuthRequest, CompleteDeviceAuthResponse, GetAuthURLRequest, GetAuthURLResponse } from "./service_pb";

/**
 * @generated from rpc svc.auth.v1.AuthService.GetAuthURL
 */
export const getAuthURL = {
  localName: "getAuthURL",
  name: "GetAuthURL",
  kind: MethodKind.Unary,
  I: GetAuthURLRequest,
  O: GetAuthURLResponse,
  service: {
    typeName: "svc.auth.v1.AuthService"
  }
} as const;

/**
 * @generated from rpc svc.auth.v1.AuthService.BeginDeviceAuth
 */
export const beginDeviceAuth = {
  localName: "beginDeviceAuth",
  name: "BeginDeviceAuth",
  kind: MethodKind.Unary,
  I: BeginDeviceAuthRequest,
  O: BeginDeviceAuthResponse,
  service: {
    typeName: "svc.auth.v1.AuthService"
  }
} as const;

/**
 * @generated from rpc svc.auth.v1.AuthService.CompleteDeviceAuth
 */
export const completeDeviceAuth = {
  localName: "completeDeviceAuth",
  name: "CompleteDeviceAuth",
  kind: MethodKind.Unary,
  I: CompleteDeviceAuthRequest,
  O: CompleteDeviceAuthResponse,
  service: {
    typeName: "svc.auth.v1.AuthService"
  }
} as const;
