/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { VendorProducts } from '../model/vendor_products'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'
import { Model } from '../model/model'
import { ModelVersion } from '../model/model_version'
import { ModelVersions } from '../model/model_versions'

export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.model'

export interface QueryGetVendorProductsRequest {
  vid: number
}

export interface QueryGetVendorProductsResponse {
  vendorProducts: VendorProducts | undefined
}

export interface QueryAllVendorProductsRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllVendorProductsResponse {
  vendorProducts: VendorProducts[]
  pagination: PageResponse | undefined
}

export interface QueryGetModelRequest {
  vid: number
  pid: number
}

export interface QueryGetModelResponse {
  model: Model | undefined
}

export interface QueryAllModelRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllModelResponse {
  model: Model[]
  pagination: PageResponse | undefined
}

export interface QueryGetModelVersionRequest {
  vid: number
  pid: number
  softwareVersion: number
}

export interface QueryGetModelVersionResponse {
  modelVersion: ModelVersion | undefined
}

export interface QueryAllModelVersionRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllModelVersionResponse {
  modelVersion: ModelVersion[]
  pagination: PageResponse | undefined
}

export interface QueryGetModelVersionsRequest {
  vid: number
  pid: number
}

export interface QueryGetModelVersionsResponse {
  modelVersions: ModelVersions | undefined
}

export interface QueryAllModelVersionsRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllModelVersionsResponse {
  modelVersions: ModelVersions[]
  pagination: PageResponse | undefined
}

const baseQueryGetVendorProductsRequest: object = { vid: 0 }

export const QueryGetVendorProductsRequest = {
  encode(message: QueryGetVendorProductsRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetVendorProductsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetVendorProductsRequest } as QueryGetVendorProductsRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetVendorProductsRequest {
    const message = { ...baseQueryGetVendorProductsRequest } as QueryGetVendorProductsRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    return message
  },

  toJSON(message: QueryGetVendorProductsRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetVendorProductsRequest>): QueryGetVendorProductsRequest {
    const message = { ...baseQueryGetVendorProductsRequest } as QueryGetVendorProductsRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    return message
  }
}

const baseQueryGetVendorProductsResponse: object = {}

export const QueryGetVendorProductsResponse = {
  encode(message: QueryGetVendorProductsResponse, writer: Writer = Writer.create()): Writer {
    if (message.vendorProducts !== undefined) {
      VendorProducts.encode(message.vendorProducts, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetVendorProductsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetVendorProductsResponse } as QueryGetVendorProductsResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorProducts = VendorProducts.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetVendorProductsResponse {
    const message = { ...baseQueryGetVendorProductsResponse } as QueryGetVendorProductsResponse
    if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
      message.vendorProducts = VendorProducts.fromJSON(object.vendorProducts)
    } else {
      message.vendorProducts = undefined
    }
    return message
  },

  toJSON(message: QueryGetVendorProductsResponse): unknown {
    const obj: any = {}
    message.vendorProducts !== undefined && (obj.vendorProducts = message.vendorProducts ? VendorProducts.toJSON(message.vendorProducts) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetVendorProductsResponse>): QueryGetVendorProductsResponse {
    const message = { ...baseQueryGetVendorProductsResponse } as QueryGetVendorProductsResponse
    if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
      message.vendorProducts = VendorProducts.fromPartial(object.vendorProducts)
    } else {
      message.vendorProducts = undefined
    }
    return message
  }
}

const baseQueryAllVendorProductsRequest: object = {}

export const QueryAllVendorProductsRequest = {
  encode(message: QueryAllVendorProductsRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllVendorProductsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllVendorProductsRequest } as QueryAllVendorProductsRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllVendorProductsRequest {
    const message = { ...baseQueryAllVendorProductsRequest } as QueryAllVendorProductsRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllVendorProductsRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllVendorProductsRequest>): QueryAllVendorProductsRequest {
    const message = { ...baseQueryAllVendorProductsRequest } as QueryAllVendorProductsRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllVendorProductsResponse: object = {}

export const QueryAllVendorProductsResponse = {
  encode(message: QueryAllVendorProductsResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.vendorProducts) {
      VendorProducts.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllVendorProductsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllVendorProductsResponse } as QueryAllVendorProductsResponse
    message.vendorProducts = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vendorProducts.push(VendorProducts.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllVendorProductsResponse {
    const message = { ...baseQueryAllVendorProductsResponse } as QueryAllVendorProductsResponse
    message.vendorProducts = []
    if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
      for (const e of object.vendorProducts) {
        message.vendorProducts.push(VendorProducts.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllVendorProductsResponse): unknown {
    const obj: any = {}
    if (message.vendorProducts) {
      obj.vendorProducts = message.vendorProducts.map((e) => (e ? VendorProducts.toJSON(e) : undefined))
    } else {
      obj.vendorProducts = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllVendorProductsResponse>): QueryAllVendorProductsResponse {
    const message = { ...baseQueryAllVendorProductsResponse } as QueryAllVendorProductsResponse
    message.vendorProducts = []
    if (object.vendorProducts !== undefined && object.vendorProducts !== null) {
      for (const e of object.vendorProducts) {
        message.vendorProducts.push(VendorProducts.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetModelRequest: object = { vid: 0, pid: 0 }

export const QueryGetModelRequest = {
  encode(message: QueryGetModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetModelRequest } as QueryGetModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetModelRequest {
    const message = { ...baseQueryGetModelRequest } as QueryGetModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    return message
  },

  toJSON(message: QueryGetModelRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetModelRequest>): QueryGetModelRequest {
    const message = { ...baseQueryGetModelRequest } as QueryGetModelRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    return message
  }
}

const baseQueryGetModelResponse: object = {}

export const QueryGetModelResponse = {
  encode(message: QueryGetModelResponse, writer: Writer = Writer.create()): Writer {
    if (message.model !== undefined) {
      Model.encode(message.model, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetModelResponse } as QueryGetModelResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.model = Model.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetModelResponse {
    const message = { ...baseQueryGetModelResponse } as QueryGetModelResponse
    if (object.model !== undefined && object.model !== null) {
      message.model = Model.fromJSON(object.model)
    } else {
      message.model = undefined
    }
    return message
  },

  toJSON(message: QueryGetModelResponse): unknown {
    const obj: any = {}
    message.model !== undefined && (obj.model = message.model ? Model.toJSON(message.model) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetModelResponse>): QueryGetModelResponse {
    const message = { ...baseQueryGetModelResponse } as QueryGetModelResponse
    if (object.model !== undefined && object.model !== null) {
      message.model = Model.fromPartial(object.model)
    } else {
      message.model = undefined
    }
    return message
  }
}

const baseQueryAllModelRequest: object = {}

export const QueryAllModelRequest = {
  encode(message: QueryAllModelRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllModelRequest } as QueryAllModelRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllModelRequest {
    const message = { ...baseQueryAllModelRequest } as QueryAllModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllModelRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllModelRequest>): QueryAllModelRequest {
    const message = { ...baseQueryAllModelRequest } as QueryAllModelRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllModelResponse: object = {}

export const QueryAllModelResponse = {
  encode(message: QueryAllModelResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.model) {
      Model.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllModelResponse } as QueryAllModelResponse
    message.model = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.model.push(Model.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllModelResponse {
    const message = { ...baseQueryAllModelResponse } as QueryAllModelResponse
    message.model = []
    if (object.model !== undefined && object.model !== null) {
      for (const e of object.model) {
        message.model.push(Model.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllModelResponse): unknown {
    const obj: any = {}
    if (message.model) {
      obj.model = message.model.map((e) => (e ? Model.toJSON(e) : undefined))
    } else {
      obj.model = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllModelResponse>): QueryAllModelResponse {
    const message = { ...baseQueryAllModelResponse } as QueryAllModelResponse
    message.model = []
    if (object.model !== undefined && object.model !== null) {
      for (const e of object.model) {
        message.model.push(Model.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetModelVersionRequest: object = { vid: 0, pid: 0, softwareVersion: 0 }

export const QueryGetModelVersionRequest = {
  encode(message: QueryGetModelVersionRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    if (message.softwareVersion !== 0) {
      writer.uint32(24).uint32(message.softwareVersion)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetModelVersionRequest } as QueryGetModelVersionRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        case 3:
          message.softwareVersion = reader.uint32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetModelVersionRequest {
    const message = { ...baseQueryGetModelVersionRequest } as QueryGetModelVersionRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = Number(object.softwareVersion)
    } else {
      message.softwareVersion = 0
    }
    return message
  },

  toJSON(message: QueryGetModelVersionRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    message.softwareVersion !== undefined && (obj.softwareVersion = message.softwareVersion)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetModelVersionRequest>): QueryGetModelVersionRequest {
    const message = { ...baseQueryGetModelVersionRequest } as QueryGetModelVersionRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    if (object.softwareVersion !== undefined && object.softwareVersion !== null) {
      message.softwareVersion = object.softwareVersion
    } else {
      message.softwareVersion = 0
    }
    return message
  }
}

const baseQueryGetModelVersionResponse: object = {}

export const QueryGetModelVersionResponse = {
  encode(message: QueryGetModelVersionResponse, writer: Writer = Writer.create()): Writer {
    if (message.modelVersion !== undefined) {
      ModelVersion.encode(message.modelVersion, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetModelVersionResponse } as QueryGetModelVersionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.modelVersion = ModelVersion.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetModelVersionResponse {
    const message = { ...baseQueryGetModelVersionResponse } as QueryGetModelVersionResponse
    if (object.modelVersion !== undefined && object.modelVersion !== null) {
      message.modelVersion = ModelVersion.fromJSON(object.modelVersion)
    } else {
      message.modelVersion = undefined
    }
    return message
  },

  toJSON(message: QueryGetModelVersionResponse): unknown {
    const obj: any = {}
    message.modelVersion !== undefined && (obj.modelVersion = message.modelVersion ? ModelVersion.toJSON(message.modelVersion) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetModelVersionResponse>): QueryGetModelVersionResponse {
    const message = { ...baseQueryGetModelVersionResponse } as QueryGetModelVersionResponse
    if (object.modelVersion !== undefined && object.modelVersion !== null) {
      message.modelVersion = ModelVersion.fromPartial(object.modelVersion)
    } else {
      message.modelVersion = undefined
    }
    return message
  }
}

const baseQueryAllModelVersionRequest: object = {}

export const QueryAllModelVersionRequest = {
  encode(message: QueryAllModelVersionRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllModelVersionRequest } as QueryAllModelVersionRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllModelVersionRequest {
    const message = { ...baseQueryAllModelVersionRequest } as QueryAllModelVersionRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllModelVersionRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllModelVersionRequest>): QueryAllModelVersionRequest {
    const message = { ...baseQueryAllModelVersionRequest } as QueryAllModelVersionRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllModelVersionResponse: object = {}

export const QueryAllModelVersionResponse = {
  encode(message: QueryAllModelVersionResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.modelVersion) {
      ModelVersion.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllModelVersionResponse } as QueryAllModelVersionResponse
    message.modelVersion = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.modelVersion.push(ModelVersion.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllModelVersionResponse {
    const message = { ...baseQueryAllModelVersionResponse } as QueryAllModelVersionResponse
    message.modelVersion = []
    if (object.modelVersion !== undefined && object.modelVersion !== null) {
      for (const e of object.modelVersion) {
        message.modelVersion.push(ModelVersion.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllModelVersionResponse): unknown {
    const obj: any = {}
    if (message.modelVersion) {
      obj.modelVersion = message.modelVersion.map((e) => (e ? ModelVersion.toJSON(e) : undefined))
    } else {
      obj.modelVersion = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllModelVersionResponse>): QueryAllModelVersionResponse {
    const message = { ...baseQueryAllModelVersionResponse } as QueryAllModelVersionResponse
    message.modelVersion = []
    if (object.modelVersion !== undefined && object.modelVersion !== null) {
      for (const e of object.modelVersion) {
        message.modelVersion.push(ModelVersion.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryGetModelVersionsRequest: object = { vid: 0, pid: 0 }

export const QueryGetModelVersionsRequest = {
  encode(message: QueryGetModelVersionsRequest, writer: Writer = Writer.create()): Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid)
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetModelVersionsRequest } as QueryGetModelVersionsRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32()
          break
        case 2:
          message.pid = reader.int32()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetModelVersionsRequest {
    const message = { ...baseQueryGetModelVersionsRequest } as QueryGetModelVersionsRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = Number(object.vid)
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = Number(object.pid)
    } else {
      message.pid = 0
    }
    return message
  },

  toJSON(message: QueryGetModelVersionsRequest): unknown {
    const obj: any = {}
    message.vid !== undefined && (obj.vid = message.vid)
    message.pid !== undefined && (obj.pid = message.pid)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetModelVersionsRequest>): QueryGetModelVersionsRequest {
    const message = { ...baseQueryGetModelVersionsRequest } as QueryGetModelVersionsRequest
    if (object.vid !== undefined && object.vid !== null) {
      message.vid = object.vid
    } else {
      message.vid = 0
    }
    if (object.pid !== undefined && object.pid !== null) {
      message.pid = object.pid
    } else {
      message.pid = 0
    }
    return message
  }
}

const baseQueryGetModelVersionsResponse: object = {}

export const QueryGetModelVersionsResponse = {
  encode(message: QueryGetModelVersionsResponse, writer: Writer = Writer.create()): Writer {
    if (message.modelVersions !== undefined) {
      ModelVersions.encode(message.modelVersions, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetModelVersionsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetModelVersionsResponse } as QueryGetModelVersionsResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.modelVersions = ModelVersions.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetModelVersionsResponse {
    const message = { ...baseQueryGetModelVersionsResponse } as QueryGetModelVersionsResponse
    if (object.modelVersions !== undefined && object.modelVersions !== null) {
      message.modelVersions = ModelVersions.fromJSON(object.modelVersions)
    } else {
      message.modelVersions = undefined
    }
    return message
  },

  toJSON(message: QueryGetModelVersionsResponse): unknown {
    const obj: any = {}
    message.modelVersions !== undefined && (obj.modelVersions = message.modelVersions ? ModelVersions.toJSON(message.modelVersions) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetModelVersionsResponse>): QueryGetModelVersionsResponse {
    const message = { ...baseQueryGetModelVersionsResponse } as QueryGetModelVersionsResponse
    if (object.modelVersions !== undefined && object.modelVersions !== null) {
      message.modelVersions = ModelVersions.fromPartial(object.modelVersions)
    } else {
      message.modelVersions = undefined
    }
    return message
  }
}

const baseQueryAllModelVersionsRequest: object = {}

export const QueryAllModelVersionsRequest = {
  encode(message: QueryAllModelVersionsRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllModelVersionsRequest } as QueryAllModelVersionsRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllModelVersionsRequest {
    const message = { ...baseQueryAllModelVersionsRequest } as QueryAllModelVersionsRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllModelVersionsRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllModelVersionsRequest>): QueryAllModelVersionsRequest {
    const message = { ...baseQueryAllModelVersionsRequest } as QueryAllModelVersionsRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllModelVersionsResponse: object = {}

export const QueryAllModelVersionsResponse = {
  encode(message: QueryAllModelVersionsResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.modelVersions) {
      ModelVersions.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllModelVersionsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllModelVersionsResponse } as QueryAllModelVersionsResponse
    message.modelVersions = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.modelVersions.push(ModelVersions.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllModelVersionsResponse {
    const message = { ...baseQueryAllModelVersionsResponse } as QueryAllModelVersionsResponse
    message.modelVersions = []
    if (object.modelVersions !== undefined && object.modelVersions !== null) {
      for (const e of object.modelVersions) {
        message.modelVersions.push(ModelVersions.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllModelVersionsResponse): unknown {
    const obj: any = {}
    if (message.modelVersions) {
      obj.modelVersions = message.modelVersions.map((e) => (e ? ModelVersions.toJSON(e) : undefined))
    } else {
      obj.modelVersions = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllModelVersionsResponse>): QueryAllModelVersionsResponse {
    const message = { ...baseQueryAllModelVersionsResponse } as QueryAllModelVersionsResponse
    message.modelVersions = []
    if (object.modelVersions !== undefined && object.modelVersions !== null) {
      for (const e of object.modelVersions) {
        message.modelVersions.push(ModelVersions.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a VendorProducts by index. */
  VendorProducts(request: QueryGetVendorProductsRequest): Promise<QueryGetVendorProductsResponse>
  /** Queries a list of VendorProducts items. */
  VendorProductsAll(request: QueryAllVendorProductsRequest): Promise<QueryAllVendorProductsResponse>
  /** Queries a Model by index. */
  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse>
  /** Queries a list of Model items. */
  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse>
  /** Queries a ModelVersion by index. */
  ModelVersion(request: QueryGetModelVersionRequest): Promise<QueryGetModelVersionResponse>
  /** Queries a list of ModelVersion items. */
  ModelVersionAll(request: QueryAllModelVersionRequest): Promise<QueryAllModelVersionResponse>
  /** Queries a ModelVersions by index. */
  ModelVersions(request: QueryGetModelVersionsRequest): Promise<QueryGetModelVersionsResponse>
  /** Queries a list of ModelVersions items. */
  ModelVersionsAll(request: QueryAllModelVersionsRequest): Promise<QueryAllModelVersionsResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  VendorProducts(request: QueryGetVendorProductsRequest): Promise<QueryGetVendorProductsResponse> {
    const data = QueryGetVendorProductsRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'VendorProducts', data)
    return promise.then((data) => QueryGetVendorProductsResponse.decode(new Reader(data)))
  }

  VendorProductsAll(request: QueryAllVendorProductsRequest): Promise<QueryAllVendorProductsResponse> {
    const data = QueryAllVendorProductsRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'VendorProductsAll', data)
    return promise.then((data) => QueryAllVendorProductsResponse.decode(new Reader(data)))
  }

  Model(request: QueryGetModelRequest): Promise<QueryGetModelResponse> {
    const data = QueryGetModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'Model', data)
    return promise.then((data) => QueryGetModelResponse.decode(new Reader(data)))
  }

  ModelAll(request: QueryAllModelRequest): Promise<QueryAllModelResponse> {
    const data = QueryAllModelRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelAll', data)
    return promise.then((data) => QueryAllModelResponse.decode(new Reader(data)))
  }

  ModelVersion(request: QueryGetModelVersionRequest): Promise<QueryGetModelVersionResponse> {
    const data = QueryGetModelVersionRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersion', data)
    return promise.then((data) => QueryGetModelVersionResponse.decode(new Reader(data)))
  }

  ModelVersionAll(request: QueryAllModelVersionRequest): Promise<QueryAllModelVersionResponse> {
    const data = QueryAllModelVersionRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersionAll', data)
    return promise.then((data) => QueryAllModelVersionResponse.decode(new Reader(data)))
  }

  ModelVersions(request: QueryGetModelVersionsRequest): Promise<QueryGetModelVersionsResponse> {
    const data = QueryGetModelVersionsRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersions', data)
    return promise.then((data) => QueryGetModelVersionsResponse.decode(new Reader(data)))
  }

  ModelVersionsAll(request: QueryAllModelVersionsRequest): Promise<QueryAllModelVersionsResponse> {
    const data = QueryAllModelVersionsRequest.encode(request).finish()
    const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.model.Query', 'ModelVersionsAll', data)
    return promise.then((data) => QueryAllModelVersionsResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
