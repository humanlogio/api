// @generated by protoc-gen-connect-es v1.6.1 with parameter "target=ts,import_extension=none"
// @generated from file svc/stack/v1/service.proto (package svc.stack.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateStackRequest, CreateStackResponse, DeleteStackRequest, DeleteStackResponse, GetStackRequest, GetStackResponse, ListStackRequest, ListStackResponse, UpdateStackRequest, UpdateStackResponse } from "./service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service svc.stack.v1.StackService
 */
export const StackService = {
  typeName: "svc.stack.v1.StackService",
  methods: {
    /**
     * @generated from rpc svc.stack.v1.StackService.CreateStack
     */
    createStack: {
      name: "CreateStack",
      I: CreateStackRequest,
      O: CreateStackResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.stack.v1.StackService.GetStack
     */
    getStack: {
      name: "GetStack",
      I: GetStackRequest,
      O: GetStackResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.stack.v1.StackService.UpdateStack
     */
    updateStack: {
      name: "UpdateStack",
      I: UpdateStackRequest,
      O: UpdateStackResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.stack.v1.StackService.DeleteStack
     */
    deleteStack: {
      name: "DeleteStack",
      I: DeleteStackRequest,
      O: DeleteStackResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc svc.stack.v1.StackService.ListStack
     */
    listStack: {
      name: "ListStack",
      I: ListStackRequest,
      O: ListStackResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

