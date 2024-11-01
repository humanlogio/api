// @generated by protoc-gen-connect-query v1.4.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/query/v1/service.proto (package svc.query.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { SummarizeEventsRequest, SummarizeEventsResponse } from "./service_pb";

/**
 * @generated from rpc svc.query.v1.QueryService.SummarizeEvents
 */
export const summarizeEvents = {
  localName: "summarizeEvents",
  name: "SummarizeEvents",
  kind: MethodKind.Unary,
  I: SummarizeEventsRequest,
  O: SummarizeEventsResponse,
  service: {
    typeName: "svc.query.v1.QueryService"
  }
} as const;
