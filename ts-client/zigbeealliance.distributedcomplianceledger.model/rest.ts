/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface DistributedcomplianceledgermodelModel {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int32 */
  deviceTypeId?: number;
  productName?: string;
  productLabel?: string;
  partNumber?: string;

  /** @format int32 */
  commissioningCustomFlow?: number;
  commissioningCustomFlowUrl?: string;

  /** @format int64 */
  commissioningModeInitialStepsHint?: number;
  commissioningModeInitialStepsInstruction?: string;

  /** @format int64 */
  commissioningModeSecondaryStepsHint?: number;
  commissioningModeSecondaryStepsInstruction?: string;
  userManualUrl?: string;
  supportUrl?: string;
  productUrl?: string;
  lsfUrl?: string;

  /** @format int32 */
  lsfRevision?: number;
  creator?: string;

  /** @format int64 */
  schemaVersion?: number;

  /** @format int32 */
  enhancedSetupFlowOptions?: number;
  enhancedSetupFlowTCUrl?: string;

  /** @format int32 */
  enhancedSetupFlowTCRevision?: number;
  enhancedSetupFlowTCDigest?: string;

  /** @format int64 */
  enhancedSetupFlowTCFileSize?: number;
  maintenanceUrl?: string;

  /** @format int64 */
  discoveryCapabilitiesBitmask?: number;
  commissioningFallbackUrl?: string;

  /** @format int64 */
  icdUserActiveModeTriggerHint?: number;
  icdUserActiveModeTriggerInstruction?: string;

  /** @format int64 */
  factoryResetStepsHint?: number;
  factoryResetStepsInstruction?: string;
}

export interface DistributedcomplianceledgermodelModelVersion {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  softwareVersionString?: string;

  /** @format int32 */
  cdVersionNumber?: number;
  firmwareInformation?: string;
  softwareVersionValid?: boolean;
  otaUrl?: string;

  /** @format uint64 */
  otaFileSize?: string;
  otaChecksum?: string;

  /** @format int32 */
  otaChecksumType?: number;

  /** @format int64 */
  minApplicableSoftwareVersion?: number;

  /** @format int64 */
  maxApplicableSoftwareVersion?: number;
  releaseNotesUrl?: string;
  creator?: string;

  /** @format int64 */
  schemaVersion?: number;

  /** @format int64 */
  specificationVersion?: number;
}

export interface DistributedcomplianceledgermodelModelVersions {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;
  softwareVersions?: number[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgermodelVendorProducts {
  /** @format int32 */
  vid?: number;
  products?: ModelProduct[];

  /** @format int64 */
  schemaVersion?: number;
}

export type ModelMsgCreateModelResponse = object;

export type ModelMsgCreateModelVersionResponse = object;

export type ModelMsgDeleteModelResponse = object;

export type ModelMsgDeleteModelVersionResponse = object;

export type ModelMsgUpdateModelResponse = object;

export type ModelMsgUpdateModelVersionResponse = object;

export interface ModelProduct {
  /** @format int32 */
  pid?: number;
  name?: string;
  partNumber?: string;

  /** @format int64 */
  schemaVersion?: number;
}

export interface ModelQueryAllModelResponse {
  model?: DistributedcomplianceledgermodelModel[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface ModelQueryGetModelResponse {
  model?: DistributedcomplianceledgermodelModel;
}

export interface ModelQueryGetModelVersionResponse {
  modelVersion?: DistributedcomplianceledgermodelModelVersion;
}

export interface ModelQueryGetModelVersionsResponse {
  modelVersions?: DistributedcomplianceledgermodelModelVersions;
}

export interface ModelQueryGetVendorProductsResponse {
  vendorProducts?: DistributedcomplianceledgermodelVendorProducts;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /**
   * next_key is the key to be passed to PageRequest.key to
   * query the next page most efficiently. It will be empty if
   * there are no more results.
   * @format byte
   */
  next_key?: string;

  /**
   * total is total number of results available if PageRequest.count_total
   * was set, its value is undefined otherwise
   * @format uint64
   */
  total?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title zigbeealliance/distributedcomplianceledger/model/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryModelAll
   * @summary Queries a list of all Model items.
   * @request GET:/dcl/model/models
   */
  queryModelAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ModelQueryAllModelResponse, RpcStatus>({
      path: `/dcl/model/models`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryVendorProducts
   * @summary Queries VendorProducts by index.
   * @request GET:/dcl/model/models/{vid}
   */
  queryVendorProducts = (vid: number, params: RequestParams = {}) =>
    this.request<ModelQueryGetVendorProductsResponse, RpcStatus>({
      path: `/dcl/model/models/${vid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryModel
   * @summary Queries a Model by index.
   * @request GET:/dcl/model/models/{vid}/{pid}
   */
  queryModel = (vid: number, pid: number, params: RequestParams = {}) =>
    this.request<ModelQueryGetModelResponse, RpcStatus>({
      path: `/dcl/model/models/${vid}/${pid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryModelVersions
   * @summary Queries ModelVersions by index.
   * @request GET:/dcl/model/versions/{vid}/{pid}
   */
  queryModelVersions = (vid: number, pid: number, params: RequestParams = {}) =>
    this.request<ModelQueryGetModelVersionsResponse, RpcStatus>({
      path: `/dcl/model/versions/${vid}/${pid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryModelVersion
   * @summary Queries a ModelVersion by index.
   * @request GET:/dcl/model/versions/{vid}/{pid}/{softwareVersion}
   */
  queryModelVersion = (vid: number, pid: number, softwareVersion: number, params: RequestParams = {}) =>
    this.request<ModelQueryGetModelVersionResponse, RpcStatus>({
      path: `/dcl/model/versions/${vid}/${pid}/${softwareVersion}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
