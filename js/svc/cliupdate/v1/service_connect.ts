// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts,import_extension=none"
// @generated from file svc/cliupdate/v1/service.proto (package svc.cliupdate.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetNextUpdateRequest, GetNextUpdateResponse } from "./service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service svc.cliupdate.v1.UpdateService
 */
export const UpdateService = {
  typeName: "svc.cliupdate.v1.UpdateService",
  methods: {
    /**
     * @generated from rpc svc.cliupdate.v1.UpdateService.GetNextUpdate
     */
    getNextUpdate: {
      name: "GetNextUpdate",
      I: GetNextUpdateRequest,
      O: GetNextUpdateResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
