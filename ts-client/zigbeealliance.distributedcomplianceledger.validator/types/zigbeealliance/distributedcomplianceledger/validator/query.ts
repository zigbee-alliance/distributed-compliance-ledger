/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { DisabledValidator } from "./disabled_validator";
import { LastValidatorPower } from "./last_validator_power";
import { ProposedDisableValidator } from "./proposed_disable_validator";
import { RejectedDisableValidator } from "./rejected_validator";
import { Validator } from "./validator";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.validator";

export interface QueryGetValidatorRequest {
  owner: string;
}

export interface QueryGetValidatorResponse {
  validator: Validator | undefined;
}

export interface QueryAllValidatorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllValidatorResponse {
  validator: Validator[];
  pagination: PageResponse | undefined;
}

export interface QueryGetLastValidatorPowerRequest {
  owner: string;
}

export interface QueryGetLastValidatorPowerResponse {
  lastValidatorPower: LastValidatorPower | undefined;
}

export interface QueryAllLastValidatorPowerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllLastValidatorPowerResponse {
  lastValidatorPower: LastValidatorPower[];
  pagination: PageResponse | undefined;
}

export interface QueryGetProposedDisableValidatorRequest {
  address: string;
}

export interface QueryGetProposedDisableValidatorResponse {
  proposedDisableValidator: ProposedDisableValidator | undefined;
}

export interface QueryAllProposedDisableValidatorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllProposedDisableValidatorResponse {
  proposedDisableValidator: ProposedDisableValidator[];
  pagination: PageResponse | undefined;
}

export interface QueryGetDisabledValidatorRequest {
  address: string;
}

export interface QueryGetDisabledValidatorResponse {
  disabledValidator: DisabledValidator | undefined;
}

export interface QueryAllDisabledValidatorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDisabledValidatorResponse {
  disabledValidator: DisabledValidator[];
  pagination: PageResponse | undefined;
}

export interface QueryGetRejectedDisableValidatorRequest {
  owner: string;
}

export interface QueryGetRejectedDisableValidatorResponse {
  rejectedValidator: RejectedDisableValidator | undefined;
}

export interface QueryAllRejectedDisableValidatorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRejectedDisableValidatorResponse {
  rejectedValidator: RejectedDisableValidator[];
  pagination: PageResponse | undefined;
}

function createBaseQueryGetValidatorRequest(): QueryGetValidatorRequest {
  return { owner: "" };
}

export const QueryGetValidatorRequest = {
  encode(message: QueryGetValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValidatorRequest {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: QueryGetValidatorRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValidatorRequest>, I>>(object: I): QueryGetValidatorRequest {
    const message = createBaseQueryGetValidatorRequest();
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryGetValidatorResponse(): QueryGetValidatorResponse {
  return { validator: undefined };
}

export const QueryGetValidatorResponse = {
  encode(message: QueryGetValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.validator !== undefined) {
      Validator.encode(message.validator, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator = Validator.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetValidatorResponse {
    return { validator: isSet(object.validator) ? Validator.fromJSON(object.validator) : undefined };
  },

  toJSON(message: QueryGetValidatorResponse): unknown {
    const obj: any = {};
    message.validator !== undefined
      && (obj.validator = message.validator ? Validator.toJSON(message.validator) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetValidatorResponse>, I>>(object: I): QueryGetValidatorResponse {
    const message = createBaseQueryGetValidatorResponse();
    message.validator = (object.validator !== undefined && object.validator !== null)
      ? Validator.fromPartial(object.validator)
      : undefined;
    return message;
  },
};

function createBaseQueryAllValidatorRequest(): QueryAllValidatorRequest {
  return { pagination: undefined };
}

export const QueryAllValidatorRequest = {
  encode(message: QueryAllValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllValidatorRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllValidatorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllValidatorRequest>, I>>(object: I): QueryAllValidatorRequest {
    const message = createBaseQueryAllValidatorRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllValidatorResponse(): QueryAllValidatorResponse {
  return { validator: [], pagination: undefined };
}

export const QueryAllValidatorResponse = {
  encode(message: QueryAllValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.validator) {
      Validator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.validator.push(Validator.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllValidatorResponse {
    return {
      validator: Array.isArray(object?.validator) ? object.validator.map((e: any) => Validator.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllValidatorResponse): unknown {
    const obj: any = {};
    if (message.validator) {
      obj.validator = message.validator.map((e) => e ? Validator.toJSON(e) : undefined);
    } else {
      obj.validator = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllValidatorResponse>, I>>(object: I): QueryAllValidatorResponse {
    const message = createBaseQueryAllValidatorResponse();
    message.validator = object.validator?.map((e) => Validator.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetLastValidatorPowerRequest(): QueryGetLastValidatorPowerRequest {
  return { owner: "" };
}

export const QueryGetLastValidatorPowerRequest = {
  encode(message: QueryGetLastValidatorPowerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetLastValidatorPowerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetLastValidatorPowerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLastValidatorPowerRequest {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: QueryGetLastValidatorPowerRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetLastValidatorPowerRequest>, I>>(
    object: I,
  ): QueryGetLastValidatorPowerRequest {
    const message = createBaseQueryGetLastValidatorPowerRequest();
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryGetLastValidatorPowerResponse(): QueryGetLastValidatorPowerResponse {
  return { lastValidatorPower: undefined };
}

export const QueryGetLastValidatorPowerResponse = {
  encode(message: QueryGetLastValidatorPowerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.lastValidatorPower !== undefined) {
      LastValidatorPower.encode(message.lastValidatorPower, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetLastValidatorPowerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetLastValidatorPowerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lastValidatorPower = LastValidatorPower.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetLastValidatorPowerResponse {
    return {
      lastValidatorPower: isSet(object.lastValidatorPower)
        ? LastValidatorPower.fromJSON(object.lastValidatorPower)
        : undefined,
    };
  },

  toJSON(message: QueryGetLastValidatorPowerResponse): unknown {
    const obj: any = {};
    message.lastValidatorPower !== undefined && (obj.lastValidatorPower = message.lastValidatorPower
      ? LastValidatorPower.toJSON(message.lastValidatorPower)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetLastValidatorPowerResponse>, I>>(
    object: I,
  ): QueryGetLastValidatorPowerResponse {
    const message = createBaseQueryGetLastValidatorPowerResponse();
    message.lastValidatorPower = (object.lastValidatorPower !== undefined && object.lastValidatorPower !== null)
      ? LastValidatorPower.fromPartial(object.lastValidatorPower)
      : undefined;
    return message;
  },
};

function createBaseQueryAllLastValidatorPowerRequest(): QueryAllLastValidatorPowerRequest {
  return { pagination: undefined };
}

export const QueryAllLastValidatorPowerRequest = {
  encode(message: QueryAllLastValidatorPowerRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllLastValidatorPowerRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllLastValidatorPowerRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLastValidatorPowerRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllLastValidatorPowerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllLastValidatorPowerRequest>, I>>(
    object: I,
  ): QueryAllLastValidatorPowerRequest {
    const message = createBaseQueryAllLastValidatorPowerRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllLastValidatorPowerResponse(): QueryAllLastValidatorPowerResponse {
  return { lastValidatorPower: [], pagination: undefined };
}

export const QueryAllLastValidatorPowerResponse = {
  encode(message: QueryAllLastValidatorPowerResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.lastValidatorPower) {
      LastValidatorPower.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllLastValidatorPowerResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllLastValidatorPowerResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.lastValidatorPower.push(LastValidatorPower.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllLastValidatorPowerResponse {
    return {
      lastValidatorPower: Array.isArray(object?.lastValidatorPower)
        ? object.lastValidatorPower.map((e: any) => LastValidatorPower.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllLastValidatorPowerResponse): unknown {
    const obj: any = {};
    if (message.lastValidatorPower) {
      obj.lastValidatorPower = message.lastValidatorPower.map((e) => e ? LastValidatorPower.toJSON(e) : undefined);
    } else {
      obj.lastValidatorPower = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllLastValidatorPowerResponse>, I>>(
    object: I,
  ): QueryAllLastValidatorPowerResponse {
    const message = createBaseQueryAllLastValidatorPowerResponse();
    message.lastValidatorPower = object.lastValidatorPower?.map((e) => LastValidatorPower.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetProposedDisableValidatorRequest(): QueryGetProposedDisableValidatorRequest {
  return { address: "" };
}

export const QueryGetProposedDisableValidatorRequest = {
  encode(message: QueryGetProposedDisableValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProposedDisableValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProposedDisableValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProposedDisableValidatorRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryGetProposedDisableValidatorRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProposedDisableValidatorRequest>, I>>(
    object: I,
  ): QueryGetProposedDisableValidatorRequest {
    const message = createBaseQueryGetProposedDisableValidatorRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryGetProposedDisableValidatorResponse(): QueryGetProposedDisableValidatorResponse {
  return { proposedDisableValidator: undefined };
}

export const QueryGetProposedDisableValidatorResponse = {
  encode(message: QueryGetProposedDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.proposedDisableValidator !== undefined) {
      ProposedDisableValidator.encode(message.proposedDisableValidator, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetProposedDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetProposedDisableValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposedDisableValidator = ProposedDisableValidator.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetProposedDisableValidatorResponse {
    return {
      proposedDisableValidator: isSet(object.proposedDisableValidator)
        ? ProposedDisableValidator.fromJSON(object.proposedDisableValidator)
        : undefined,
    };
  },

  toJSON(message: QueryGetProposedDisableValidatorResponse): unknown {
    const obj: any = {};
    message.proposedDisableValidator !== undefined && (obj.proposedDisableValidator = message.proposedDisableValidator
      ? ProposedDisableValidator.toJSON(message.proposedDisableValidator)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetProposedDisableValidatorResponse>, I>>(
    object: I,
  ): QueryGetProposedDisableValidatorResponse {
    const message = createBaseQueryGetProposedDisableValidatorResponse();
    message.proposedDisableValidator =
      (object.proposedDisableValidator !== undefined && object.proposedDisableValidator !== null)
        ? ProposedDisableValidator.fromPartial(object.proposedDisableValidator)
        : undefined;
    return message;
  },
};

function createBaseQueryAllProposedDisableValidatorRequest(): QueryAllProposedDisableValidatorRequest {
  return { pagination: undefined };
}

export const QueryAllProposedDisableValidatorRequest = {
  encode(message: QueryAllProposedDisableValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProposedDisableValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProposedDisableValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProposedDisableValidatorRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllProposedDisableValidatorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProposedDisableValidatorRequest>, I>>(
    object: I,
  ): QueryAllProposedDisableValidatorRequest {
    const message = createBaseQueryAllProposedDisableValidatorRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllProposedDisableValidatorResponse(): QueryAllProposedDisableValidatorResponse {
  return { proposedDisableValidator: [], pagination: undefined };
}

export const QueryAllProposedDisableValidatorResponse = {
  encode(message: QueryAllProposedDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.proposedDisableValidator) {
      ProposedDisableValidator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllProposedDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllProposedDisableValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.proposedDisableValidator.push(ProposedDisableValidator.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllProposedDisableValidatorResponse {
    return {
      proposedDisableValidator: Array.isArray(object?.proposedDisableValidator)
        ? object.proposedDisableValidator.map((e: any) => ProposedDisableValidator.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllProposedDisableValidatorResponse): unknown {
    const obj: any = {};
    if (message.proposedDisableValidator) {
      obj.proposedDisableValidator = message.proposedDisableValidator.map((e) =>
        e ? ProposedDisableValidator.toJSON(e) : undefined
      );
    } else {
      obj.proposedDisableValidator = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllProposedDisableValidatorResponse>, I>>(
    object: I,
  ): QueryAllProposedDisableValidatorResponse {
    const message = createBaseQueryAllProposedDisableValidatorResponse();
    message.proposedDisableValidator =
      object.proposedDisableValidator?.map((e) => ProposedDisableValidator.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDisabledValidatorRequest(): QueryGetDisabledValidatorRequest {
  return { address: "" };
}

export const QueryGetDisabledValidatorRequest = {
  encode(message: QueryGetDisabledValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDisabledValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDisabledValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDisabledValidatorRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryGetDisabledValidatorRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDisabledValidatorRequest>, I>>(
    object: I,
  ): QueryGetDisabledValidatorRequest {
    const message = createBaseQueryGetDisabledValidatorRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryGetDisabledValidatorResponse(): QueryGetDisabledValidatorResponse {
  return { disabledValidator: undefined };
}

export const QueryGetDisabledValidatorResponse = {
  encode(message: QueryGetDisabledValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.disabledValidator !== undefined) {
      DisabledValidator.encode(message.disabledValidator, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDisabledValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDisabledValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.disabledValidator = DisabledValidator.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDisabledValidatorResponse {
    return {
      disabledValidator: isSet(object.disabledValidator)
        ? DisabledValidator.fromJSON(object.disabledValidator)
        : undefined,
    };
  },

  toJSON(message: QueryGetDisabledValidatorResponse): unknown {
    const obj: any = {};
    message.disabledValidator !== undefined && (obj.disabledValidator = message.disabledValidator
      ? DisabledValidator.toJSON(message.disabledValidator)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDisabledValidatorResponse>, I>>(
    object: I,
  ): QueryGetDisabledValidatorResponse {
    const message = createBaseQueryGetDisabledValidatorResponse();
    message.disabledValidator = (object.disabledValidator !== undefined && object.disabledValidator !== null)
      ? DisabledValidator.fromPartial(object.disabledValidator)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDisabledValidatorRequest(): QueryAllDisabledValidatorRequest {
  return { pagination: undefined };
}

export const QueryAllDisabledValidatorRequest = {
  encode(message: QueryAllDisabledValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDisabledValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDisabledValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDisabledValidatorRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDisabledValidatorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDisabledValidatorRequest>, I>>(
    object: I,
  ): QueryAllDisabledValidatorRequest {
    const message = createBaseQueryAllDisabledValidatorRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDisabledValidatorResponse(): QueryAllDisabledValidatorResponse {
  return { disabledValidator: [], pagination: undefined };
}

export const QueryAllDisabledValidatorResponse = {
  encode(message: QueryAllDisabledValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.disabledValidator) {
      DisabledValidator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDisabledValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDisabledValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.disabledValidator.push(DisabledValidator.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDisabledValidatorResponse {
    return {
      disabledValidator: Array.isArray(object?.disabledValidator)
        ? object.disabledValidator.map((e: any) => DisabledValidator.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDisabledValidatorResponse): unknown {
    const obj: any = {};
    if (message.disabledValidator) {
      obj.disabledValidator = message.disabledValidator.map((e) => e ? DisabledValidator.toJSON(e) : undefined);
    } else {
      obj.disabledValidator = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDisabledValidatorResponse>, I>>(
    object: I,
  ): QueryAllDisabledValidatorResponse {
    const message = createBaseQueryAllDisabledValidatorResponse();
    message.disabledValidator = object.disabledValidator?.map((e) => DisabledValidator.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetRejectedDisableValidatorRequest(): QueryGetRejectedDisableValidatorRequest {
  return { owner: "" };
}

export const QueryGetRejectedDisableValidatorRequest = {
  encode(message: QueryGetRejectedDisableValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRejectedDisableValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRejectedDisableValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRejectedDisableValidatorRequest {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: QueryGetRejectedDisableValidatorRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRejectedDisableValidatorRequest>, I>>(
    object: I,
  ): QueryGetRejectedDisableValidatorRequest {
    const message = createBaseQueryGetRejectedDisableValidatorRequest();
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryGetRejectedDisableValidatorResponse(): QueryGetRejectedDisableValidatorResponse {
  return { rejectedValidator: undefined };
}

export const QueryGetRejectedDisableValidatorResponse = {
  encode(message: QueryGetRejectedDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.rejectedValidator !== undefined) {
      RejectedDisableValidator.encode(message.rejectedValidator, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetRejectedDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetRejectedDisableValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rejectedValidator = RejectedDisableValidator.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRejectedDisableValidatorResponse {
    return {
      rejectedValidator: isSet(object.rejectedValidator)
        ? RejectedDisableValidator.fromJSON(object.rejectedValidator)
        : undefined,
    };
  },

  toJSON(message: QueryGetRejectedDisableValidatorResponse): unknown {
    const obj: any = {};
    message.rejectedValidator !== undefined && (obj.rejectedValidator = message.rejectedValidator
      ? RejectedDisableValidator.toJSON(message.rejectedValidator)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetRejectedDisableValidatorResponse>, I>>(
    object: I,
  ): QueryGetRejectedDisableValidatorResponse {
    const message = createBaseQueryGetRejectedDisableValidatorResponse();
    message.rejectedValidator = (object.rejectedValidator !== undefined && object.rejectedValidator !== null)
      ? RejectedDisableValidator.fromPartial(object.rejectedValidator)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRejectedDisableValidatorRequest(): QueryAllRejectedDisableValidatorRequest {
  return { pagination: undefined };
}

export const QueryAllRejectedDisableValidatorRequest = {
  encode(message: QueryAllRejectedDisableValidatorRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRejectedDisableValidatorRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRejectedDisableValidatorRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRejectedDisableValidatorRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllRejectedDisableValidatorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRejectedDisableValidatorRequest>, I>>(
    object: I,
  ): QueryAllRejectedDisableValidatorRequest {
    const message = createBaseQueryAllRejectedDisableValidatorRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllRejectedDisableValidatorResponse(): QueryAllRejectedDisableValidatorResponse {
  return { rejectedValidator: [], pagination: undefined };
}

export const QueryAllRejectedDisableValidatorResponse = {
  encode(message: QueryAllRejectedDisableValidatorResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.rejectedValidator) {
      RejectedDisableValidator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllRejectedDisableValidatorResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllRejectedDisableValidatorResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.rejectedValidator.push(RejectedDisableValidator.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRejectedDisableValidatorResponse {
    return {
      rejectedValidator: Array.isArray(object?.rejectedValidator)
        ? object.rejectedValidator.map((e: any) => RejectedDisableValidator.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllRejectedDisableValidatorResponse): unknown {
    const obj: any = {};
    if (message.rejectedValidator) {
      obj.rejectedValidator = message.rejectedValidator.map((e) => e ? RejectedDisableValidator.toJSON(e) : undefined);
    } else {
      obj.rejectedValidator = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllRejectedDisableValidatorResponse>, I>>(
    object: I,
  ): QueryAllRejectedDisableValidatorResponse {
    const message = createBaseQueryAllRejectedDisableValidatorResponse();
    message.rejectedValidator = object.rejectedValidator?.map((e) => RejectedDisableValidator.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a validator by index. */
  Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse>;
  /** Queries a list of validator items. */
  ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse>;
  /** Queries a lastValidatorPower by index. */
  LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse>;
  /** Queries a list of lastValidatorPower items. */
  LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse>;
  /** Queries a ProposedDisableValidator by index. */
  ProposedDisableValidator(
    request: QueryGetProposedDisableValidatorRequest,
  ): Promise<QueryGetProposedDisableValidatorResponse>;
  /** Queries a list of ProposedDisableValidator items. */
  ProposedDisableValidatorAll(
    request: QueryAllProposedDisableValidatorRequest,
  ): Promise<QueryAllProposedDisableValidatorResponse>;
  /** Queries a DisabledValidator by index. */
  DisabledValidator(request: QueryGetDisabledValidatorRequest): Promise<QueryGetDisabledValidatorResponse>;
  /** Queries a list of DisabledValidator items. */
  DisabledValidatorAll(request: QueryAllDisabledValidatorRequest): Promise<QueryAllDisabledValidatorResponse>;
  /** Queries a RejectedNode by index. */
  RejectedDisableValidator(
    request: QueryGetRejectedDisableValidatorRequest,
  ): Promise<QueryGetRejectedDisableValidatorResponse>;
  /** Queries a list of RejectedNode items. */
  RejectedDisableValidatorAll(
    request: QueryAllRejectedDisableValidatorRequest,
  ): Promise<QueryAllRejectedDisableValidatorResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Validator = this.Validator.bind(this);
    this.ValidatorAll = this.ValidatorAll.bind(this);
    this.LastValidatorPower = this.LastValidatorPower.bind(this);
    this.LastValidatorPowerAll = this.LastValidatorPowerAll.bind(this);
    this.ProposedDisableValidator = this.ProposedDisableValidator.bind(this);
    this.ProposedDisableValidatorAll = this.ProposedDisableValidatorAll.bind(this);
    this.DisabledValidator = this.DisabledValidator.bind(this);
    this.DisabledValidatorAll = this.DisabledValidatorAll.bind(this);
    this.RejectedDisableValidator = this.RejectedDisableValidator.bind(this);
    this.RejectedDisableValidatorAll = this.RejectedDisableValidatorAll.bind(this);
  }
  Validator(request: QueryGetValidatorRequest): Promise<QueryGetValidatorResponse> {
    const data = QueryGetValidatorRequest.encode(request).finish();
    const promise = this.rpc.request("zigbeealliance.distributedcomplianceledger.validator.Query", "Validator", data);
    return promise.then((data) => QueryGetValidatorResponse.decode(new _m0.Reader(data)));
  }

  ValidatorAll(request: QueryAllValidatorRequest): Promise<QueryAllValidatorResponse> {
    const data = QueryAllValidatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "ValidatorAll",
      data,
    );
    return promise.then((data) => QueryAllValidatorResponse.decode(new _m0.Reader(data)));
  }

  LastValidatorPower(request: QueryGetLastValidatorPowerRequest): Promise<QueryGetLastValidatorPowerResponse> {
    const data = QueryGetLastValidatorPowerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "LastValidatorPower",
      data,
    );
    return promise.then((data) => QueryGetLastValidatorPowerResponse.decode(new _m0.Reader(data)));
  }

  LastValidatorPowerAll(request: QueryAllLastValidatorPowerRequest): Promise<QueryAllLastValidatorPowerResponse> {
    const data = QueryAllLastValidatorPowerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "LastValidatorPowerAll",
      data,
    );
    return promise.then((data) => QueryAllLastValidatorPowerResponse.decode(new _m0.Reader(data)));
  }

  ProposedDisableValidator(
    request: QueryGetProposedDisableValidatorRequest,
  ): Promise<QueryGetProposedDisableValidatorResponse> {
    const data = QueryGetProposedDisableValidatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "ProposedDisableValidator",
      data,
    );
    return promise.then((data) => QueryGetProposedDisableValidatorResponse.decode(new _m0.Reader(data)));
  }

  ProposedDisableValidatorAll(
    request: QueryAllProposedDisableValidatorRequest,
  ): Promise<QueryAllProposedDisableValidatorResponse> {
    const data = QueryAllProposedDisableValidatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "ProposedDisableValidatorAll",
      data,
    );
    return promise.then((data) => QueryAllProposedDisableValidatorResponse.decode(new _m0.Reader(data)));
  }

  DisabledValidator(request: QueryGetDisabledValidatorRequest): Promise<QueryGetDisabledValidatorResponse> {
    const data = QueryGetDisabledValidatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "DisabledValidator",
      data,
    );
    return promise.then((data) => QueryGetDisabledValidatorResponse.decode(new _m0.Reader(data)));
  }

  DisabledValidatorAll(request: QueryAllDisabledValidatorRequest): Promise<QueryAllDisabledValidatorResponse> {
    const data = QueryAllDisabledValidatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "DisabledValidatorAll",
      data,
    );
    return promise.then((data) => QueryAllDisabledValidatorResponse.decode(new _m0.Reader(data)));
  }

  RejectedDisableValidator(
    request: QueryGetRejectedDisableValidatorRequest,
  ): Promise<QueryGetRejectedDisableValidatorResponse> {
    const data = QueryGetRejectedDisableValidatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "RejectedDisableValidator",
      data,
    );
    return promise.then((data) => QueryGetRejectedDisableValidatorResponse.decode(new _m0.Reader(data)));
  }

  RejectedDisableValidatorAll(
    request: QueryAllRejectedDisableValidatorRequest,
  ): Promise<QueryAllRejectedDisableValidatorResponse> {
    const data = QueryAllRejectedDisableValidatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "zigbeealliance.distributedcomplianceledger.validator.Query",
      "RejectedDisableValidatorAll",
      data,
    );
    return promise.then((data) => QueryAllRejectedDisableValidatorResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
