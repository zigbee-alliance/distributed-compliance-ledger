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

export interface PkiApprovedCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];
}

export interface PkiApprovedCertificatesBySubject {
  subject?: string;
  subjectKeyIds?: string[];
}

export interface PkiApprovedRootCertificates {
  certs?: PkiCertificateIdentifier[];
}

export interface PkiCertificate {
  pemCert?: string;
  serialNumber?: string;
  issuer?: string;
  authorityKeyId?: string;
  rootSubject?: string;
  rootSubjectKeyId?: string;
  isRoot?: boolean;
  owner?: string;
  subject?: string;
  subjectKeyId?: string;
  approvals?: PkiGrant[];
  subjectAsText?: string;
  rejects?: PkiGrant[];

  /** @format int32 */
  vid?: number;
}

export interface PkiCertificateIdentifier {
  subject?: string;
  subjectKeyId?: string;
}

export interface PkiChildCertificates {
  issuer?: string;
  authorityKeyId?: string;
  certIds?: PkiCertificateIdentifier[];
}

export interface PkiGrant {
  address?: string;

  /** @format int64 */
  time?: string;
  info?: string;
}

export type PkiMsgAddPkiRevocationDistributionPointResponse = object;

export type PkiMsgAddX509CertResponse = object;

export type PkiMsgApproveAddX509RootCertResponse = object;

export type PkiMsgApproveRevokeX509RootCertResponse = object;

export type PkiMsgAssignVidResponse = object;

export type PkiMsgDeletePkiRevocationDistributionPointResponse = object;

export type PkiMsgProposeAddX509RootCertResponse = object;

export type PkiMsgProposeRevokeX509RootCertResponse = object;

export type PkiMsgRejectAddX509RootCertResponse = object;

export type PkiMsgRevokeX509CertResponse = object;

export type PkiMsgUpdatePkiRevocationDistributionPointResponse = object;

export interface PkiPkiRevocationDistributionPoint {
  /** @format int32 */
  vid?: number;
  label?: string;
  issuerSubjectKeyID?: string;

  /** @format int32 */
  pid?: number;
  isPAA?: boolean;
  crlSignerCertificate?: string;
  dataURL?: string;

  /** @format uint64 */
  dataFileSize?: string;
  dataDigest?: string;

  /** @format int64 */
  dataDigestType?: number;

  /** @format int64 */
  revocationType?: number;
}

export interface PkiPkiRevocationDistributionPointsByIssuerSubjectKeyID {
  issuerSubjectKeyID?: string;
  points?: PkiPkiRevocationDistributionPoint[];
}

export interface PkiProposedCertificate {
  subject?: string;
  subjectKeyId?: string;
  pemCert?: string;
  serialNumber?: string;
  owner?: string;
  approvals?: PkiGrant[];
  subjectAsText?: string;
  rejects?: PkiGrant[];

  /** @format int32 */
  vid?: number;
}

export interface PkiProposedCertificateRevocation {
  subject?: string;
  subjectKeyId?: string;
  approvals?: PkiGrant[];
  subjectAsText?: string;
}

export interface PkiQueryAllApprovedCertificatesResponse {
  approvedCertificates?: PkiApprovedCertificates[];

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

export interface PkiQueryAllPkiRevocationDistributionPointResponse {
  PkiRevocationDistributionPoint?: PkiPkiRevocationDistributionPoint[];

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

export interface PkiQueryAllProposedCertificateResponse {
  proposedCertificate?: PkiProposedCertificate[];

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

export interface PkiQueryAllProposedCertificateRevocationResponse {
  proposedCertificateRevocation?: PkiProposedCertificateRevocation[];

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

export interface PkiQueryAllRejectedCertificatesResponse {
  rejectedCertificate?: PkiRejectedCertificate[];

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

export interface PkiQueryAllRevokedCertificatesResponse {
  revokedCertificates?: PkiRevokedCertificates[];

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

export interface PkiQueryGetApprovedCertificatesBySubjectResponse {
  approvedCertificatesBySubject?: PkiApprovedCertificatesBySubject;
}

export interface PkiQueryGetApprovedCertificatesResponse {
  approvedCertificates?: PkiApprovedCertificates;
}

export interface PkiQueryGetApprovedRootCertificatesResponse {
  approvedRootCertificates?: PkiApprovedRootCertificates;
}

export interface PkiQueryGetChildCertificatesResponse {
  childCertificates?: PkiChildCertificates;
}

export interface PkiQueryGetPkiRevocationDistributionPointResponse {
  PkiRevocationDistributionPoint?: PkiPkiRevocationDistributionPoint;
}

export interface PkiQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
  pkiRevocationDistributionPointsByIssuerSubjectKeyID?: PkiPkiRevocationDistributionPointsByIssuerSubjectKeyID;
}

export interface PkiQueryGetProposedCertificateResponse {
  proposedCertificate?: PkiProposedCertificate;
}

export interface PkiQueryGetProposedCertificateRevocationResponse {
  proposedCertificateRevocation?: PkiProposedCertificateRevocation;
}

export interface PkiQueryGetRejectedCertificatesResponse {
  rejectedCertificate?: PkiRejectedCertificate;
}

export interface PkiQueryGetRevokedCertificatesResponse {
  revokedCertificates?: PkiRevokedCertificates;
}

export interface PkiQueryGetRevokedRootCertificatesResponse {
  revokedRootCertificates?: PkiRevokedRootCertificates;
}

export interface PkiRejectedCertificate {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];
}

export interface PkiRevokedCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];
}

export interface PkiRevokedRootCertificates {
  certs?: PkiCertificateIdentifier[];
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
 * @title pki/approved_certificates.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryApprovedCertificatesAll
   * @summary Queries a list of ApprovedCertificates items.
   * @request GET:/dcl/pki/certificates
   */
  queryApprovedCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllApprovedCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryApprovedCertificatesBySubject
   * @summary Queries a ApprovedCertificatesBySubject by index.
   * @request GET:/dcl/pki/certificates/{subject}
   */
  queryApprovedCertificatesBySubject = (subject: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetApprovedCertificatesBySubjectResponse, RpcStatus>({
      path: `/dcl/pki/certificates/${subject}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryApprovedCertificates
   * @summary Queries a ApprovedCertificates by index.
   * @request GET:/dcl/pki/certificates/{subject}/{subjectKeyId}
   */
  queryApprovedCertificates = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetApprovedCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryChildCertificates
   * @summary Queries a ChildCertificates by index.
   * @request GET:/dcl/pki/child-certificates/{issuer}/{authorityKeyId}
   */
  queryChildCertificates = (issuer: string, authorityKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetChildCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/child-certificates/${issuer}/${authorityKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProposedCertificateAll
   * @summary Queries a list of ProposedCertificate items.
   * @request GET:/dcl/pki/proposed-certificates
   */
  queryProposedCertificateAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllProposedCertificateResponse, RpcStatus>({
      path: `/dcl/pki/proposed-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProposedCertificate
   * @summary Queries a ProposedCertificate by index.
   * @request GET:/dcl/pki/proposed-certificates/{subject}/{subjectKeyId}
   */
  queryProposedCertificate = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetProposedCertificateResponse, RpcStatus>({
      path: `/dcl/pki/proposed-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProposedCertificateRevocationAll
   * @summary Queries a list of ProposedCertificateRevocation items.
   * @request GET:/dcl/pki/proposed-revocation-certificates
   */
  queryProposedCertificateRevocationAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllProposedCertificateRevocationResponse, RpcStatus>({
      path: `/dcl/pki/proposed-revocation-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProposedCertificateRevocation
   * @summary Queries a ProposedCertificateRevocation by index.
   * @request GET:/dcl/pki/proposed-revocation-certificates/{subject}/{subjectKeyId}
   */
  queryProposedCertificateRevocation = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetProposedCertificateRevocationResponse, RpcStatus>({
      path: `/dcl/pki/proposed-revocation-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRejectedCertificateAll
   * @summary Queries a list of RejectedCertificate items.
   * @request GET:/dcl/pki/rejected-certificates
   */
  queryRejectedCertificateAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllRejectedCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/rejected-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRejectedCertificate
   * @summary Queries a RejectedCertificate by index.
   * @request GET:/dcl/pki/rejected-certificates/{subject}/{subjectKeyId}
   */
  queryRejectedCertificate = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetRejectedCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/rejected-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPkiRevocationDistributionPointAll
   * @summary Queries a list of PkiRevocationDistributionPoint items.
   * @request GET:/dcl/pki/revocation-points
   */
  queryPkiRevocationDistributionPointAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllPkiRevocationDistributionPointResponse, RpcStatus>({
      path: `/dcl/pki/revocation-points`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPkiRevocationDistributionPointsByIssuerSubjectKeyID
   * @summary Queries a PkiRevocationDistributionPointsByIssuerSubjectKeyID by index.
   * @request GET:/dcl/pki/revocation-points/{issuerSubjectKeyID}
   */
  queryPkiRevocationDistributionPointsByIssuerSubjectKeyID = (issuerSubjectKeyID: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse, RpcStatus>({
      path: `/dcl/pki/revocation-points/${issuerSubjectKeyID}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPkiRevocationDistributionPoint
   * @summary Queries a PkiRevocationDistributionPoint by index.
   * @request GET:/dcl/pki/revocation-points/{issuerSubjectKeyID}/{vid}/{label}
   */
  queryPkiRevocationDistributionPoint = (
    issuerSubjectKeyID: string,
    vid: number,
    label: string,
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryGetPkiRevocationDistributionPointResponse, RpcStatus>({
      path: `/dcl/pki/revocation-points/${issuerSubjectKeyID}/${vid}/${label}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedCertificatesAll
   * @summary Queries a list of RevokedCertificates items.
   * @request GET:/dcl/pki/revoked-certificates
   */
  queryRevokedCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllRevokedCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/revoked-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedCertificates
   * @summary Queries a RevokedCertificates by index.
   * @request GET:/dcl/pki/revoked-certificates/{subject}/{subjectKeyId}
   */
  queryRevokedCertificates = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetRevokedCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/revoked-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedRootCertificates
   * @summary Queries a RevokedRootCertificates by index.
   * @request GET:/dcl/pki/revoked-root-certificates
   */
  queryRevokedRootCertificates = (params: RequestParams = {}) =>
    this.request<PkiQueryGetRevokedRootCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/revoked-root-certificates`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryApprovedRootCertificates
   * @summary Queries a ApprovedRootCertificates by index.
   * @request GET:/dcl/pki/root-certificates
   */
  queryApprovedRootCertificates = (params: RequestParams = {}) =>
    this.request<PkiQueryGetApprovedRootCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/root-certificates`,
      method: "GET",
      format: "json",
      ...params,
    });
}
