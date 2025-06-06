// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/query/v1/trace_service.proto (package svc.query.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetSpanRequest, GetSpanResponse, GetTraceRequest, GetTraceResponse } from "./trace_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service svc.query.v1.TraceService
 */
export const TraceService = {
  typeName: "svc.query.v1.TraceService",
  methods: {
    /**
     * @generated from rpc svc.query.v1.TraceService.GetTrace
     */
    getTrace: {
      name: "GetTrace",
      I: GetTraceRequest,
      O: GetTraceResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.query.v1.TraceService.GetSpan
     */
    getSpan: {
      name: "GetSpan",
      I: GetSpanRequest,
      O: GetSpanResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

