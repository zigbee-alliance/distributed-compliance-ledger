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

export interface ComplianceComplianceHistoryItem {
  /** @format int64 */
  softwareVersionCertificationStatus?: number;
  date?: string;
  reason?: string;

  /** @format int64 */
  cDVersionNumber?: number;

  /** @format int64 */
  schemaVersion?: number;
}

export type ComplianceMsgCertifyModelResponse = object;

export type ComplianceMsgDeleteComplianceInfoResponse = object;

export type ComplianceMsgProvisionModelResponse = object;

export type ComplianceMsgRevokeModelResponse = object;

export type ComplianceMsgUpdateComplianceInfoResponse = object;

export interface ComplianceQueryAllCertifiedModelResponse {
  certifiedModel?: DistributedcomplianceledgercomplianceCertifiedModel[];

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

export interface ComplianceQueryAllComplianceInfoResponse {
  complianceInfo?: DistributedcomplianceledgercomplianceComplianceInfo[];

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

export interface ComplianceQueryAllDeviceSoftwareComplianceResponse {
  deviceSoftwareCompliance?: DistributedcomplianceledgercomplianceDeviceSoftwareCompliance[];

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

export interface ComplianceQueryAllProvisionalModelResponse {
  provisionalModel?: DistributedcomplianceledgercomplianceProvisionalModel[];

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

export interface ComplianceQueryAllRevokedModelResponse {
  revokedModel?: DistributedcomplianceledgercomplianceRevokedModel[];

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

export interface ComplianceQueryGetCertifiedModelResponse {
  certifiedModel?: DistributedcomplianceledgercomplianceCertifiedModel;
}

export interface ComplianceQueryGetComplianceInfoResponse {
  complianceInfo?: DistributedcomplianceledgercomplianceComplianceInfo;
}

export interface ComplianceQueryGetDeviceSoftwareComplianceResponse {
  deviceSoftwareCompliance?: DistributedcomplianceledgercomplianceDeviceSoftwareCompliance;
}

export interface ComplianceQueryGetProvisionalModelResponse {
  provisionalModel?: DistributedcomplianceledgercomplianceProvisionalModel;
}

export interface ComplianceQueryGetRevokedModelResponse {
  revokedModel?: DistributedcomplianceledgercomplianceRevokedModel;
}

export interface DistributedcomplianceledgercomplianceCertifiedModel {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  certificationType?: string;
  value?: boolean;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgercomplianceComplianceInfo {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  certificationType?: string;
  softwareVersionString?: string;

  /** @format int64 */
  cDVersionNumber?: number;

  /** @format int64 */
  softwareVersionCertificationStatus?: number;
  date?: string;
  reason?: string;
  owner?: string;
  history?: ComplianceComplianceHistoryItem[];
  cDCertificateId?: string;
  certificationRoute?: string;
  programType?: string;
  programTypeVersion?: string;
  compliantPlatformUsed?: string;
  compliantPlatformVersion?: string;
  transport?: string;
  familyId?: string;
  supportedClusters?: string;
  OSVersion?: string;
  parentChild?: string;
  certificationIdOfSoftwareComponent?: string;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgercomplianceDeviceSoftwareCompliance {
  cDCertificateId?: string;
  complianceInfo?: DistributedcomplianceledgercomplianceComplianceInfo[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgercomplianceProvisionalModel {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  certificationType?: string;
  value?: boolean;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgercomplianceRevokedModel {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  certificationType?: string;
  value?: boolean;

  /** @format int64 */
  schemaVersion?: number;
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
 * @title zigbeealliance/distributedcomplianceledger/compliance/certified_model.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCertifiedModelAll
   * @summary Queries a list of CertifiedModel items.
   * @request GET:/dcl/compliance/certified-models
   */
  queryCertifiedModelAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryAllCertifiedModelResponse, RpcStatus>({
      path: `/dcl/compliance/certified-models`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCertifiedModel
   * @summary Queries a CertifiedModel by index.
   * @request GET:/dcl/compliance/certified-models/{vid}/{pid}/{softwareVersion}/{certificationType}
   */
  queryCertifiedModel = (
    vid: number,
    pid: number,
    softwareVersion: number,
    certificationType: string,
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryGetCertifiedModelResponse, RpcStatus>({
      path: `/dcl/compliance/certified-models/${vid}/${pid}/${softwareVersion}/${certificationType}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryComplianceInfoAll
   * @summary Queries a list of ComplianceInfo items.
   * @request GET:/dcl/compliance/compliance-info
   */
  queryComplianceInfoAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryAllComplianceInfoResponse, RpcStatus>({
      path: `/dcl/compliance/compliance-info`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryComplianceInfo
   * @summary Queries a ComplianceInfo by index.
   * @request GET:/dcl/compliance/compliance-info/{vid}/{pid}/{softwareVersion}/{certificationType}
   */
  queryComplianceInfo = (
    vid: number,
    pid: number,
    softwareVersion: number,
    certificationType: string,
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryGetComplianceInfoResponse, RpcStatus>({
      path: `/dcl/compliance/compliance-info/${vid}/${pid}/${softwareVersion}/${certificationType}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDeviceSoftwareComplianceAll
   * @summary Queries a list of DeviceSoftwareCompliance items.
   * @request GET:/dcl/compliance/device-software-compliance
   */
  queryDeviceSoftwareComplianceAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryAllDeviceSoftwareComplianceResponse, RpcStatus>({
      path: `/dcl/compliance/device-software-compliance`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDeviceSoftwareCompliance
   * @summary Queries a DeviceSoftwareCompliance by index.
   * @request GET:/dcl/compliance/device-software-compliance/{cDCertificateId}
   */
  queryDeviceSoftwareCompliance = (cDCertificateId: string, params: RequestParams = {}) =>
    this.request<ComplianceQueryGetDeviceSoftwareComplianceResponse, RpcStatus>({
      path: `/dcl/compliance/device-software-compliance/${cDCertificateId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProvisionalModelAll
   * @summary Queries a list of ProvisionalModel items.
   * @request GET:/dcl/compliance/provisional-models
   */
  queryProvisionalModelAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryAllProvisionalModelResponse, RpcStatus>({
      path: `/dcl/compliance/provisional-models`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProvisionalModel
   * @summary Queries a ProvisionalModel by index.
   * @request GET:/dcl/compliance/provisional-models/{vid}/{pid}/{softwareVersion}/{certificationType}
   */
  queryProvisionalModel = (
    vid: number,
    pid: number,
    softwareVersion: number,
    certificationType: string,
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryGetProvisionalModelResponse, RpcStatus>({
      path: `/dcl/compliance/provisional-models/${vid}/${pid}/${softwareVersion}/${certificationType}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedModelAll
   * @summary Queries a list of RevokedModel items.
   * @request GET:/dcl/compliance/revoked-models
   */
  queryRevokedModelAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryAllRevokedModelResponse, RpcStatus>({
      path: `/dcl/compliance/revoked-models`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedModel
   * @summary Queries a RevokedModel by index.
   * @request GET:/dcl/compliance/revoked-models/{vid}/{pid}/{softwareVersion}/{certificationType}
   */
  queryRevokedModel = (
    vid: number,
    pid: number,
    softwareVersion: number,
    certificationType: string,
    params: RequestParams = {},
  ) =>
    this.request<ComplianceQueryGetRevokedModelResponse, RpcStatus>({
      path: `/dcl/compliance/revoked-models/${vid}/${pid}/${softwareVersion}/${certificationType}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
