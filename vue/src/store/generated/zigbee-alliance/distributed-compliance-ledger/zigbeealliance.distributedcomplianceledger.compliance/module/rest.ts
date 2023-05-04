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

export interface ComplianceCertifiedModel {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  certificationType?: string;
  value?: boolean;
}

export interface ComplianceComplianceHistoryItem {
  /** @format int64 */
  softwareVersionCertificationStatus?: number;
  date?: string;
  reason?: string;

  /** @format int64 */
  cDVersionNumber?: number;
}

export interface ComplianceComplianceInfo {
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
}

export interface ComplianceDeviceSoftwareCompliance {
  cDCertificateId?: string;
  complianceInfo?: ComplianceComplianceInfo[];
}

export type ComplianceMsgCertifyModelResponse = object;

export type ComplianceMsgDeleteComplianceInfoResponse = object;

export type ComplianceMsgProvisionModelResponse = object;

export type ComplianceMsgRevokeModelResponse = object;

export type ComplianceMsgUpdateComplianceInfoResponse = object;

export interface ComplianceProvisionalModel {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  certificationType?: string;
  value?: boolean;
}

export interface ComplianceQueryAllCertifiedModelResponse {
  certifiedModel?: ComplianceCertifiedModel[];

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
  complianceInfo?: ComplianceComplianceInfo[];

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
  deviceSoftwareCompliance?: ComplianceDeviceSoftwareCompliance[];

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
  provisionalModel?: ComplianceProvisionalModel[];

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
  revokedModel?: ComplianceRevokedModel[];

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
  certifiedModel?: ComplianceCertifiedModel;
}

export interface ComplianceQueryGetComplianceInfoResponse {
  complianceInfo?: ComplianceComplianceInfo;
}

export interface ComplianceQueryGetDeviceSoftwareComplianceResponse {
  deviceSoftwareCompliance?: ComplianceDeviceSoftwareCompliance;
}

export interface ComplianceQueryGetProvisionalModelResponse {
  provisionalModel?: ComplianceProvisionalModel;
}

export interface ComplianceQueryGetRevokedModelResponse {
  revokedModel?: ComplianceRevokedModel;
}

export interface ComplianceRevokedModel {
  /** @format int32 */
  vid?: number;

  /** @format int32 */
  pid?: number;

  /** @format int64 */
  softwareVersion?: number;
  certificationType?: string;
  value?: boolean;
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
  countTotal?: boolean;

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
  /** @format byte */
  nextKey?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title compliance/certified_model.proto
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
      "pagination.countTotal"?: boolean;
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
      "pagination.countTotal"?: boolean;
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
      "pagination.countTotal"?: boolean;
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
      "pagination.countTotal"?: boolean;
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
      "pagination.countTotal"?: boolean;
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
