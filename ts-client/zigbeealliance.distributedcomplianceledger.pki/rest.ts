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

export interface DistributedcomplianceledgerpkiAllCertificatesBySubject {
  subject?: string;
  subjectKeyIds?: string[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiApprovedCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiApprovedCertificatesBySubject {
  subject?: string;
  subjectKeyIds?: string[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiApprovedRootCertificates {
  certs?: PkiCertificateIdentifier[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiChildCertificates {
  issuer?: string;
  authorityKeyId?: string;
  certIds?: PkiCertificateIdentifier[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiNocCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format float */
  tq?: number;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiNocCertificatesBySubject {
  subject?: string;
  subjectKeyIds?: string[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiNocCertificatesByVidAndSkid {
  /** @format int32 */
  vid?: number;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format float */
  tq?: number;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiNocIcaCertificates {
  /** @format int32 */
  vid?: number;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiNocRootCertificates {
  /** @format int32 */
  vid?: number;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiPkiRevocationDistributionPoint {
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

  /** @format int64 */
  schemaVersion?: number;
  crlSignerDelegator?: string;
}

export interface DistributedcomplianceledgerpkiPkiRevocationDistributionPointsByIssuerSubjectKeyID {
  issuerSubjectKeyID?: string;
  points?: DistributedcomplianceledgerpkiPkiRevocationDistributionPoint[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiProposedCertificate {
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

  /** @format int64 */
  certSchemaVersion?: number;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiProposedCertificateRevocation {
  subject?: string;
  subjectKeyId?: string;
  approvals?: PkiGrant[];
  subjectAsText?: string;
  serialNumber?: string;
  revokeChild?: boolean;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiRejectedCertificate {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiRevokedCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiRevokedNocIcaCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiRevokedNocRootCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerpkiRevokedRootCertificates {
  certs?: PkiCertificateIdentifier[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface PkiAllCertificates {
  subject?: string;
  subjectKeyId?: string;
  certs?: PkiCertificate[];

  /** @format int64 */
  schemaVersion?: number;
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
  certificateType?: PkiCertificateType;

  /** @format int64 */
  schemaVersion?: number;
}

export interface PkiCertificateIdentifier {
  subject?: string;
  subjectKeyId?: string;
}

export enum PkiCertificateType {
  DeviceAttestationPKI = "DeviceAttestationPKI",
  OperationalPKI = "OperationalPKI",
  VIDSignerPKI = "VIDSignerPKI",
}

export interface PkiGrant {
  address?: string;

  /**
   * number of nanoseconds elapsed since January 1, 1970 UTC
   * @format int64
   */
  time?: string;
  info?: string;

  /** @format int64 */
  schemaVersion?: number;
}

export type PkiMsgAddNocX509IcaCertResponse = object;

export type PkiMsgAddNocX509RootCertResponse = object;

export type PkiMsgAddPkiRevocationDistributionPointResponse = object;

export type PkiMsgAddX509CertResponse = object;

export type PkiMsgApproveAddX509RootCertResponse = object;

export type PkiMsgApproveRevokeX509RootCertResponse = object;

export type PkiMsgAssignVidResponse = object;

export type PkiMsgDeletePkiRevocationDistributionPointResponse = object;

export type PkiMsgProposeAddX509RootCertResponse = object;

export type PkiMsgProposeRevokeX509RootCertResponse = object;

export type PkiMsgRejectAddX509RootCertResponse = object;

export type PkiMsgRemoveNocX509IcaCertResponse = object;

export type PkiMsgRemoveNocX509RootCertResponse = object;

export type PkiMsgRemoveX509CertResponse = object;

export type PkiMsgRevokeNocX509IcaCertResponse = object;

export type PkiMsgRevokeNocX509RootCertResponse = object;

export type PkiMsgRevokeX509CertResponse = object;

export type PkiMsgUpdatePkiRevocationDistributionPointResponse = object;

export interface PkiQueryAllApprovedCertificatesResponse {
  approvedCertificates?: DistributedcomplianceledgerpkiApprovedCertificates[];

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

export interface PkiQueryAllCertificatesResponse {
  certificates?: PkiAllCertificates[];

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

export interface PkiQueryAllNocIcaCertificatesResponse {
  nocIcaCertificates?: DistributedcomplianceledgerpkiNocIcaCertificates[];

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

export interface PkiQueryAllNocRootCertificatesResponse {
  nocRootCertificates?: DistributedcomplianceledgerpkiNocRootCertificates[];

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
  PkiRevocationDistributionPoint?: DistributedcomplianceledgerpkiPkiRevocationDistributionPoint[];

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
  proposedCertificate?: DistributedcomplianceledgerpkiProposedCertificate[];

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
  proposedCertificateRevocation?: DistributedcomplianceledgerpkiProposedCertificateRevocation[];

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
  rejectedCertificate?: DistributedcomplianceledgerpkiRejectedCertificate[];

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
  revokedCertificates?: DistributedcomplianceledgerpkiRevokedCertificates[];

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

export interface PkiQueryAllRevokedNocIcaCertificatesResponse {
  revokedNocIcaCertificates?: DistributedcomplianceledgerpkiRevokedNocIcaCertificates[];

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

export interface PkiQueryAllRevokedNocRootCertificatesResponse {
  revokedNocRootCertificates?: DistributedcomplianceledgerpkiRevokedNocRootCertificates[];

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

export interface PkiQueryGetAllCertificatesBySubjectResponse {
  allCertificatesBySubject?: DistributedcomplianceledgerpkiAllCertificatesBySubject;
}

export interface PkiQueryGetApprovedCertificatesBySubjectResponse {
  approvedCertificatesBySubject?: DistributedcomplianceledgerpkiApprovedCertificatesBySubject;
}

export interface PkiQueryGetApprovedCertificatesResponse {
  approvedCertificates?: DistributedcomplianceledgerpkiApprovedCertificates;
}

export interface PkiQueryGetApprovedRootCertificatesResponse {
  approvedRootCertificates?: DistributedcomplianceledgerpkiApprovedRootCertificates;
}

export interface PkiQueryGetCertificatesResponse {
  certificates?: PkiAllCertificates;
}

export interface PkiQueryGetChildCertificatesResponse {
  childCertificates?: DistributedcomplianceledgerpkiChildCertificates;
}

export interface PkiQueryGetNocCertificatesBySubjectResponse {
  nocCertificatesBySubject?: DistributedcomplianceledgerpkiNocCertificatesBySubject;
}

export interface PkiQueryGetNocCertificatesByVidAndSkidResponse {
  nocCertificatesByVidAndSkid?: DistributedcomplianceledgerpkiNocCertificatesByVidAndSkid;
}

export interface PkiQueryGetNocCertificatesResponse {
  nocCertificates?: DistributedcomplianceledgerpkiNocCertificates;
}

export interface PkiQueryGetNocIcaCertificatesResponse {
  nocIcaCertificates?: DistributedcomplianceledgerpkiNocIcaCertificates;
}

export interface PkiQueryGetNocRootCertificatesResponse {
  nocRootCertificates?: DistributedcomplianceledgerpkiNocRootCertificates;
}

export interface PkiQueryGetPkiRevocationDistributionPointResponse {
  PkiRevocationDistributionPoint?: DistributedcomplianceledgerpkiPkiRevocationDistributionPoint;
}

export interface PkiQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse {
  pkiRevocationDistributionPointsByIssuerSubjectKeyID?: DistributedcomplianceledgerpkiPkiRevocationDistributionPointsByIssuerSubjectKeyID;
}

export interface PkiQueryGetProposedCertificateResponse {
  proposedCertificate?: DistributedcomplianceledgerpkiProposedCertificate;
}

export interface PkiQueryGetProposedCertificateRevocationResponse {
  proposedCertificateRevocation?: DistributedcomplianceledgerpkiProposedCertificateRevocation;
}

export interface PkiQueryGetRejectedCertificatesResponse {
  rejectedCertificate?: DistributedcomplianceledgerpkiRejectedCertificate;
}

export interface PkiQueryGetRevokedCertificatesResponse {
  revokedCertificates?: DistributedcomplianceledgerpkiRevokedCertificates;
}

export interface PkiQueryGetRevokedNocIcaCertificatesResponse {
  revokedNocIcaCertificates?: DistributedcomplianceledgerpkiRevokedNocIcaCertificates;
}

export interface PkiQueryGetRevokedNocRootCertificatesResponse {
  revokedNocRootCertificates?: DistributedcomplianceledgerpkiRevokedNocRootCertificates;
}

export interface PkiQueryGetRevokedRootCertificatesResponse {
  revokedRootCertificates?: DistributedcomplianceledgerpkiRevokedRootCertificates;
}

export interface PkiQueryNocCertificatesResponse {
  nocCertificates?: DistributedcomplianceledgerpkiNocCertificates[];

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
 * @title zigbeealliance/distributedcomplianceledger/pki/all_certificates.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryCertificatesAll
   * @summary Queries a list of Certificates items.
   * @request GET:/dcl/pki/all-certificates
   */
  queryCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
      subjectKeyId?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/all-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAllCertificatesBySubject
   * @summary Queries a AllCertificatesBySubject by index.
   * @request GET:/dcl/pki/all-certificates/{subject}
   */
  queryAllCertificatesBySubject = (subject: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetAllCertificatesBySubjectResponse, RpcStatus>({
      path: `/dcl/pki/all-certificates/${subject}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryCertificates
   * @summary Queries a Certificates by index.
   * @request GET:/dcl/pki/all-certificates/{subject}/{subjectKeyId}
   */
  queryCertificates = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/all-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNocCertificatesAll
   * @summary Queries a list of NocCertificates items.
   * @request GET:/dcl/pki/all-noc-certificates
   */
  queryNocCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
      subjectKeyId?: string;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryNocCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/all-noc-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNocCertificatesBySubject
   * @summary Queries a NocCertificatesBySubject by index.
   * @request GET:/dcl/pki/all-noc-certificates/{subject}
   */
  queryNocCertificatesBySubject = (subject: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetNocCertificatesBySubjectResponse, RpcStatus>({
      path: `/dcl/pki/all-noc-certificates/${subject}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNocCertificates
   * @summary Queries a NocCertificates by index.
   * @request GET:/dcl/pki/all-noc-certificates/{subject}/{subjectKeyId}
   */
  queryNocCertificates = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetNocCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/all-noc-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

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
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
      subjectKeyId?: string;
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
   * @name QueryNocIcaCertificatesAll
   * @summary Queries a list of NocIcaCertificates items.
   * @request GET:/dcl/pki/noc-ica-certificates
   */
  queryNocIcaCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllNocIcaCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/noc-ica-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNocRootCertificatesAll
   * @summary Queries a list of NocRootCertificates items.
   * @request GET:/dcl/pki/noc-root-certificates
   */
  queryNocRootCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllNocRootCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/noc-root-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNocCertificatesByVidAndSkid
   * @summary Queries a NocCertificatesByVidAndSkid by index.
   * @request GET:/dcl/pki/noc-vid-certificates/{vid}/{subjectKeyId}
   */
  queryNocCertificatesByVidAndSkid = (vid: number, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetNocCertificatesByVidAndSkidResponse, RpcStatus>({
      path: `/dcl/pki/noc-vid-certificates/${vid}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNocIcaCertificates
   * @summary Queries a NocIcaCertificates by index.
   * @request GET:/dcl/pki/noc-vid-ica-certificates/{vid}
   */
  queryNocIcaCertificates = (vid: number, params: RequestParams = {}) =>
    this.request<PkiQueryGetNocIcaCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/noc-vid-ica-certificates/${vid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryNocRootCertificates
   * @summary Queries a NocRootCertificates by index.
   * @request GET:/dcl/pki/noc-vid-root-certificates/{vid}
   */
  queryNocRootCertificates = (vid: number, params: RequestParams = {}) =>
    this.request<PkiQueryGetNocRootCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/noc-vid-root-certificates/${vid}`,
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
      "pagination.count_total"?: boolean;
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
      "pagination.count_total"?: boolean;
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
  queryProposedCertificateRevocation = (
    subject: string,
    subjectKeyId: string,
    query?: { serialNumber?: string },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryGetProposedCertificateRevocationResponse, RpcStatus>({
      path: `/dcl/pki/proposed-revocation-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      query: query,
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
      "pagination.count_total"?: boolean;
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
      "pagination.count_total"?: boolean;
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
   * @name QueryPkiRevocationDistributionPointsByIssuerSubjectKeyId
   * @summary Queries a PkiRevocationDistributionPointsByIssuerSubjectKeyID by index.
   * @request GET:/dcl/pki/revocation-points/{issuerSubjectKeyID}
   */
  queryPkiRevocationDistributionPointsByIssuerSubjectKeyID = (issuerSubjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetPkiRevocationDistributionPointsByIssuerSubjectKeyIDResponse, RpcStatus>({
      path: `/dcl/pki/revocation-points/${issuerSubjectKeyId}`,
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
    issuerSubjectKeyId: string,
    vid: number,
    label: string,
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryGetPkiRevocationDistributionPointResponse, RpcStatus>({
      path: `/dcl/pki/revocation-points/${issuerSubjectKeyId}/${vid}/${label}`,
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
      "pagination.count_total"?: boolean;
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
   * @name QueryRevokedNocIcaCertificatesAll
   * @summary Queries a list of RevokedNocIcaCertificates items.
   * @request GET:/dcl/pki/revoked-noc-ica-certificates
   */
  queryRevokedNocIcaCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllRevokedNocIcaCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/revoked-noc-ica-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedNocIcaCertificates
   * @summary Queries a RevokedNocIcaCertificates by index.
   * @request GET:/dcl/pki/revoked-noc-ica-certificates/{subject}/{subjectKeyId}
   */
  queryRevokedNocIcaCertificates = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetRevokedNocIcaCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/revoked-noc-ica-certificates/${subject}/${subjectKeyId}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedNocRootCertificatesAll
   * @summary Queries a list of RevokedNocRootCertificates items.
   * @request GET:/dcl/pki/revoked-noc-root-certificates
   */
  queryRevokedNocRootCertificatesAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<PkiQueryAllRevokedNocRootCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/revoked-noc-root-certificates`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedNocRootCertificates
   * @summary Queries a RevokedNocRootCertificates by index.
   * @request GET:/dcl/pki/revoked-noc-root-certificates/{subject}/{subjectKeyId}
   */
  queryRevokedNocRootCertificates = (subject: string, subjectKeyId: string, params: RequestParams = {}) =>
    this.request<PkiQueryGetRevokedNocRootCertificatesResponse, RpcStatus>({
      path: `/dcl/pki/revoked-noc-root-certificates/${subject}/${subjectKeyId}`,
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
