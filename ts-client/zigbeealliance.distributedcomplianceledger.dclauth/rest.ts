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

export enum RevokedAccountReason {
  TrusteeVoting = "TrusteeVoting",
  MaliciousValidator = "MaliciousValidator",
}

export interface CommonUint16Range {
  /** @format int32 */
  min?: number;

  /** @format int32 */
  max?: number;
}

export interface DclauthGrant {
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

export type DclauthMsgApproveAddAccountResponse = object;

export type DclauthMsgApproveRevokeAccountResponse = object;

export type DclauthMsgProposeAddAccountResponse = object;

export type DclauthMsgProposeRevokeAccountResponse = object;

export type DclauthMsgRejectAddAccountResponse = object;

export interface DclauthQueryAllAccountResponse {
  account?: DistributedcomplianceledgerdclauthAccount[];

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

export interface DclauthQueryAllPendingAccountResponse {
  pendingAccount?: DistributedcomplianceledgerdclauthPendingAccount[];

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

export interface DclauthQueryAllPendingAccountRevocationResponse {
  pendingAccountRevocation?: DistributedcomplianceledgerdclauthPendingAccountRevocation[];

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

export interface DclauthQueryAllRejectedAccountResponse {
  rejectedAccount?: DistributedcomplianceledgerdclauthRejectedAccount[];

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

export interface DclauthQueryAllRevokedAccountResponse {
  revokedAccount?: DistributedcomplianceledgerdclauthRevokedAccount[];

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

export interface DclauthQueryGetAccountResponse {
  account?: DistributedcomplianceledgerdclauthAccount;
}

export interface DclauthQueryGetAccountStatResponse {
  AccountStat?: DistributedcomplianceledgerdclauthAccountStat;
}

export interface DclauthQueryGetPendingAccountResponse {
  pendingAccount?: DistributedcomplianceledgerdclauthPendingAccount;
}

export interface DclauthQueryGetPendingAccountRevocationResponse {
  pendingAccountRevocation?: DistributedcomplianceledgerdclauthPendingAccountRevocation;
}

export interface DclauthQueryGetRejectedAccountResponse {
  rejectedAccount?: DistributedcomplianceledgerdclauthRejectedAccount;
}

export interface DclauthQueryGetRevokedAccountResponse {
  revokedAccount?: DistributedcomplianceledgerdclauthRevokedAccount;
}

export interface DistributedcomplianceledgerdclauthAccount {
  /**
   * BaseAccount defines a base account type. It contains all the necessary fields
   * for basic account functionality. Any custom account type should extend this
   * type for additional functionality (e.g. vesting).
   */
  base_account?: V1Beta1BaseAccount;

  /**
   * NOTE. we do not user AccountRoles casting here to preserve repeated form
   *       so protobuf takes care about repeated items in generated code,
   *       (but that might be not the final solution)
   */
  roles?: string[];
  approvals?: DclauthGrant[];

  /** @format int32 */
  vendorID?: number;
  rejects?: DclauthGrant[];
  productIDs?: CommonUint16Range[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerdclauthAccountStat {
  /** @format uint64 */
  number?: string;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerdclauthPendingAccount {
  account?: DistributedcomplianceledgerdclauthAccount;

  /** @format int64 */
  pendingAccountSchemaVersion?: number;
}

export interface DistributedcomplianceledgerdclauthPendingAccountRevocation {
  address?: string;
  approvals?: DclauthGrant[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgerdclauthRejectedAccount {
  account?: DistributedcomplianceledgerdclauthAccount;

  /** @format int64 */
  rejectedAccountSchemaVersion?: number;
}

export interface DistributedcomplianceledgerdclauthRevokedAccount {
  account?: DistributedcomplianceledgerdclauthAccount;
  revokeApprovals?: DclauthGrant[];
  reason?: RevokedAccountReason;

  /** @format int64 */
  revokedAccountSchemaVersion?: number;
}

/**
* `Any` contains an arbitrary serialized protocol buffer message along with a
URL that describes the type of the serialized message.

Protobuf library provides support to pack/unpack Any values in the form
of utility functions or additional generated methods of the Any type.

Example 1: Pack and unpack a message in C++.

    Foo foo = ...;
    Any any;
    any.PackFrom(foo);
    ...
    if (any.UnpackTo(&foo)) {
      ...
    }

Example 2: Pack and unpack a message in Java.

    Foo foo = ...;
    Any any = Any.pack(foo);
    ...
    if (any.is(Foo.class)) {
      foo = any.unpack(Foo.class);
    }

 Example 3: Pack and unpack a message in Python.

    foo = Foo(...)
    any = Any()
    any.Pack(foo)
    ...
    if any.Is(Foo.DESCRIPTOR):
      any.Unpack(foo)
      ...

 Example 4: Pack and unpack a message in Go

     foo := &pb.Foo{...}
     any, err := anypb.New(foo)
     if err != nil {
       ...
     }
     ...
     foo := &pb.Foo{}
     if err := any.UnmarshalTo(foo); err != nil {
       ...
     }

The pack methods provided by protobuf library will by default use
'type.googleapis.com/full.type.name' as the type URL and the unpack
methods only use the fully qualified type name after the last '/'
in the type URL, for example "foo.bar.com/x/y.z" will yield type
name "y.z".


JSON
====
The JSON representation of an `Any` value uses the regular
representation of the deserialized, embedded message, with an
additional field `@type` which contains the type URL. Example:

    package google.profile;
    message Person {
      string first_name = 1;
      string last_name = 2;
    }

    {
      "@type": "type.googleapis.com/google.profile.Person",
      "firstName": <string>,
      "lastName": <string>
    }

If the embedded message type is well-known and has a custom JSON
representation, that representation will be embedded adding a field
`value` which holds the custom JSON in addition to the `@type`
field. Example (for message [google.protobuf.Duration][]):

    {
      "@type": "type.googleapis.com/google.protobuf.Duration",
      "value": "1.212s"
    }
*/
export interface ProtobufAny {
  /**
   * A URL/resource name that uniquely identifies the type of the serialized
   * protocol buffer message. This string must contain at least
   * one "/" character. The last segment of the URL's path must represent
   * the fully qualified name of the type (as in
   * `path/google.protobuf.Duration`). The name should be in a canonical form
   * (e.g., leading "." is not accepted).
   *
   * In practice, teams usually precompile into the binary all types that they
   * expect it to use in the context of Any. However, for URLs which use the
   * scheme `http`, `https`, or no scheme, one can optionally set up a type
   * server that maps type URLs to message definitions as follows:
   * * If no scheme is provided, `https` is assumed.
   * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
   *   value in binary format, or produce an error.
   * * Applications are allowed to cache lookup results based on the
   *   URL, or have them precompiled into a binary to avoid any
   *   lookup. Therefore, binary compatibility needs to be preserved
   *   on changes to types. (Use versioned type names to manage
   *   breaking changes.)
   * Note: this functionality is not currently available in the official
   * protobuf release, and it is not used for type URLs beginning with
   * type.googleapis.com.
   * Schemes other than `http`, `https` (or the empty scheme) might be
   * used with implementation specific semantics.
   */
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* BaseAccount defines a base account type. It contains all the necessary fields
for basic account functionality. Any custom account type should extend this
type for additional functionality (e.g. vesting).
*/
export interface V1Beta1BaseAccount {
  address?: string;

  /**
   * `Any` contains an arbitrary serialized protocol buffer message along with a
   * URL that describes the type of the serialized message.
   *
   * Protobuf library provides support to pack/unpack Any values in the form
   * of utility functions or additional generated methods of the Any type.
   * Example 1: Pack and unpack a message in C++.
   *     Foo foo = ...;
   *     Any any;
   *     any.PackFrom(foo);
   *     ...
   *     if (any.UnpackTo(&foo)) {
   *       ...
   *     }
   * Example 2: Pack and unpack a message in Java.
   *     Any any = Any.pack(foo);
   *     if (any.is(Foo.class)) {
   *       foo = any.unpack(Foo.class);
   *  Example 3: Pack and unpack a message in Python.
   *     foo = Foo(...)
   *     any = Any()
   *     any.Pack(foo)
   *     if any.Is(Foo.DESCRIPTOR):
   *       any.Unpack(foo)
   *  Example 4: Pack and unpack a message in Go
   *      foo := &pb.Foo{...}
   *      any, err := anypb.New(foo)
   *      if err != nil {
   *        ...
   *      }
   *      ...
   *      foo := &pb.Foo{}
   *      if err := any.UnmarshalTo(foo); err != nil {
   * The pack methods provided by protobuf library will by default use
   * 'type.googleapis.com/full.type.name' as the type URL and the unpack
   * methods only use the fully qualified type name after the last '/'
   * in the type URL, for example "foo.bar.com/x/y.z" will yield type
   * name "y.z".
   * JSON
   * ====
   * The JSON representation of an `Any` value uses the regular
   * representation of the deserialized, embedded message, with an
   * additional field `@type` which contains the type URL. Example:
   *     package google.profile;
   *     message Person {
   *       string first_name = 1;
   *       string last_name = 2;
   *     {
   *       "@type": "type.googleapis.com/google.profile.Person",
   *       "firstName": <string>,
   *       "lastName": <string>
   * If the embedded message type is well-known and has a custom JSON
   * representation, that representation will be embedded adding a field
   * `value` which holds the custom JSON in addition to the `@type`
   * field. Example (for message [google.protobuf.Duration][]):
   *       "@type": "type.googleapis.com/google.protobuf.Duration",
   *       "value": "1.212s"
   */
  pub_key?: ProtobufAny;

  /** @format uint64 */
  account_number?: string;

  /** @format uint64 */
  sequence?: string;
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
 * @title zigbeealliance/distributedcomplianceledger/dclauth/account.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAccountAll
   * @summary Queries a list of account items.
   * @request GET:/dcl/auth/accounts
   */
  queryAccountAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DclauthQueryAllAccountResponse, RpcStatus>({
      path: `/dcl/auth/accounts`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAccountStat
   * @summary Queries a accountStat by index.
   * @request GET:/dcl/auth/accounts/stat
   */
  queryAccountStat = (params: RequestParams = {}) =>
    this.request<DclauthQueryGetAccountStatResponse, RpcStatus>({
      path: `/dcl/auth/accounts/stat`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAccount
   * @summary Queries a account by index.
   * @request GET:/dcl/auth/accounts/{address}
   */
  queryAccount = (address: string, params: RequestParams = {}) =>
    this.request<DclauthQueryGetAccountResponse, RpcStatus>({
      path: `/dcl/auth/accounts/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPendingAccountAll
   * @summary Queries a list of pendingAccount items.
   * @request GET:/dcl/auth/proposed-accounts
   */
  queryPendingAccountAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DclauthQueryAllPendingAccountResponse, RpcStatus>({
      path: `/dcl/auth/proposed-accounts`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPendingAccount
   * @summary Queries a pendingAccount by index.
   * @request GET:/dcl/auth/proposed-accounts/{address}
   */
  queryPendingAccount = (address: string, params: RequestParams = {}) =>
    this.request<DclauthQueryGetPendingAccountResponse, RpcStatus>({
      path: `/dcl/auth/proposed-accounts/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPendingAccountRevocationAll
   * @summary Queries a list of pendingAccountRevocation items.
   * @request GET:/dcl/auth/proposed-revocation-accounts
   */
  queryPendingAccountRevocationAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DclauthQueryAllPendingAccountRevocationResponse, RpcStatus>({
      path: `/dcl/auth/proposed-revocation-accounts`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPendingAccountRevocation
   * @summary Queries a pendingAccountRevocation by index.
   * @request GET:/dcl/auth/proposed-revocation-accounts/{address}
   */
  queryPendingAccountRevocation = (address: string, params: RequestParams = {}) =>
    this.request<DclauthQueryGetPendingAccountRevocationResponse, RpcStatus>({
      path: `/dcl/auth/proposed-revocation-accounts/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRejectedAccountAll
   * @summary Queries a list of RejectedAccount items.
   * @request GET:/dcl/auth/rejected-accounts
   */
  queryRejectedAccountAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DclauthQueryAllRejectedAccountResponse, RpcStatus>({
      path: `/dcl/auth/rejected-accounts`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRejectedAccount
   * @summary Queries a RejectedAccount by index.
   * @request GET:/dcl/auth/rejected-accounts/{address}
   */
  queryRejectedAccount = (address: string, params: RequestParams = {}) =>
    this.request<DclauthQueryGetRejectedAccountResponse, RpcStatus>({
      path: `/dcl/auth/rejected-accounts/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedAccountAll
   * @summary Queries a list of RevokedAccount items.
   * @request GET:/dcl/auth/revoked-accounts
   */
  queryRevokedAccountAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<DclauthQueryAllRevokedAccountResponse, RpcStatus>({
      path: `/dcl/auth/revoked-accounts`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRevokedAccount
   * @summary Queries a RevokedAccount by index.
   * @request GET:/dcl/auth/revoked-accounts/{address}
   */
  queryRevokedAccount = (address: string, params: RequestParams = {}) =>
    this.request<DclauthQueryGetRevokedAccountResponse, RpcStatus>({
      path: `/dcl/auth/revoked-accounts/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
