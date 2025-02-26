// @generated by protoc-gen-es v1.9.0 with parameter "target=ts"
// @generated from file rill/admin/v1/telemetry.proto (package rill.admin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Struct } from "@bufbuild/protobuf";

/**
 * @generated from message rill.admin.v1.RecordEventsRequest
 */
export class RecordEventsRequest extends Message<RecordEventsRequest> {
  /**
   * @generated from field: repeated google.protobuf.Struct events = 1;
   */
  events: Struct[] = [];

  constructor(data?: PartialMessage<RecordEventsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rill.admin.v1.RecordEventsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "events", kind: "message", T: Struct, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RecordEventsRequest {
    return new RecordEventsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RecordEventsRequest {
    return new RecordEventsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RecordEventsRequest {
    return new RecordEventsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RecordEventsRequest | PlainMessage<RecordEventsRequest> | undefined, b: RecordEventsRequest | PlainMessage<RecordEventsRequest> | undefined): boolean {
    return proto3.util.equals(RecordEventsRequest, a, b);
  }
}

/**
 * @generated from message rill.admin.v1.RecordEventsResponse
 */
export class RecordEventsResponse extends Message<RecordEventsResponse> {
  constructor(data?: PartialMessage<RecordEventsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "rill.admin.v1.RecordEventsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RecordEventsResponse {
    return new RecordEventsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RecordEventsResponse {
    return new RecordEventsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RecordEventsResponse {
    return new RecordEventsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RecordEventsResponse | PlainMessage<RecordEventsResponse> | undefined, b: RecordEventsResponse | PlainMessage<RecordEventsResponse> | undefined): boolean {
    return proto3.util.equals(RecordEventsResponse, a, b);
  }
}

