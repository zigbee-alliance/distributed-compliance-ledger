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

export interface DistributedcomplianceledgervalidatorDisabledValidator {
  address?: string;
  creator?: string;
  approvals?: ValidatorGrant[];
  disabledByNodeAdmin?: boolean;
  rejects?: ValidatorGrant[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgervalidatorLastValidatorPower {
  owner?: string;

  /** @format int32 */
  power?: number;

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgervalidatorProposedDisableValidator {
  address?: string;
  creator?: string;
  approvals?: ValidatorGrant[];
  rejects?: ValidatorGrant[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgervalidatorRejectedDisableValidator {
  address?: string;
  creator?: string;
  approvals?: ValidatorGrant[];
  rejects?: ValidatorGrant[];

  /** @format int64 */
  schemaVersion?: number;
}

export interface DistributedcomplianceledgervalidatorValidator {
  /** the account address of validator owner */
  owner?: string;

  /** description of the validator */
  description?: ValidatorDescription;

  /**
   * the consensus public key of the tendermint validator
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
  pubKey?: ProtobufAny;

  /**
   * validator consensus power
   * @format int32
   */
  power?: number;

  /** has the validator been removed from validator set */
  jailed?: boolean;

  /** the reason of validator jailing */
  jailedReason?: string;

  /** @format int64 */
  schemaVersion?: number;
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

export interface ValidatorDescription {
  /** a human-readable name for the validator. */
  moniker?: string;

  /** optional identity signature. */
  identity?: string;

  /** optional website link. */
  website?: string;

  /** optional details. */
  details?: string;
}

export interface ValidatorGrant {
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

export type ValidatorMsgApproveDisableValidatorResponse = object;

export type ValidatorMsgCreateValidatorResponse = object;

export type ValidatorMsgDisableValidatorResponse = object;

export type ValidatorMsgEnableValidatorResponse = object;

export type ValidatorMsgProposeDisableValidatorResponse = object;

export type ValidatorMsgRejectDisableValidatorResponse = object;

export interface ValidatorQueryAllDisabledValidatorResponse {
  disabledValidator?: DistributedcomplianceledgervalidatorDisabledValidator[];

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

export interface ValidatorQueryAllLastValidatorPowerResponse {
  lastValidatorPower?: DistributedcomplianceledgervalidatorLastValidatorPower[];

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

export interface ValidatorQueryAllProposedDisableValidatorResponse {
  proposedDisableValidator?: DistributedcomplianceledgervalidatorProposedDisableValidator[];

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

export interface ValidatorQueryAllRejectedDisableValidatorResponse {
  rejectedValidator?: DistributedcomplianceledgervalidatorRejectedDisableValidator[];

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

export interface ValidatorQueryAllValidatorResponse {
  validator?: DistributedcomplianceledgervalidatorValidator[];

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

export interface ValidatorQueryGetDisabledValidatorResponse {
  disabledValidator?: DistributedcomplianceledgervalidatorDisabledValidator;
}

export interface ValidatorQueryGetLastValidatorPowerResponse {
  lastValidatorPower?: DistributedcomplianceledgervalidatorLastValidatorPower;
}

export interface ValidatorQueryGetProposedDisableValidatorResponse {
  proposedDisableValidator?: DistributedcomplianceledgervalidatorProposedDisableValidator;
}

export interface ValidatorQueryGetRejectedDisableValidatorResponse {
  rejectedValidator?: DistributedcomplianceledgervalidatorRejectedDisableValidator;
}

export interface ValidatorQueryGetValidatorResponse {
  validator?: DistributedcomplianceledgervalidatorValidator;
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
 * @title zigbeealliance/distributedcomplianceledger/validator/description.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryDisabledValidatorAll
   * @summary Queries a list of DisabledValidator items.
   * @request GET:/dcl/validator/disabled-nodes
   */
  queryDisabledValidatorAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllDisabledValidatorResponse, RpcStatus>({
      path: `/dcl/validator/disabled-nodes`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDisabledValidator
   * @summary Queries a DisabledValidator by index.
   * @request GET:/dcl/validator/disabled-nodes/{address}
   */
  queryDisabledValidator = (address: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetDisabledValidatorResponse, RpcStatus>({
      path: `/dcl/validator/disabled-nodes/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLastValidatorPowerAll
   * @summary Queries a list of lastValidatorPower items.
   * @request GET:/dcl/validator/last-powers
   */
  queryLastValidatorPowerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllLastValidatorPowerResponse, RpcStatus>({
      path: `/dcl/validator/last-powers`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryLastValidatorPower
   * @summary Queries a lastValidatorPower by index.
   * @request GET:/dcl/validator/last-powers/{owner}
   */
  queryLastValidatorPower = (owner: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetLastValidatorPowerResponse, RpcStatus>({
      path: `/dcl/validator/last-powers/${owner}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidatorAll
   * @summary Queries a list of validator items.
   * @request GET:/dcl/validator/nodes
   */
  queryValidatorAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllValidatorResponse, RpcStatus>({
      path: `/dcl/validator/nodes`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidator
   * @summary Queries a validator by index.
   * @request GET:/dcl/validator/nodes/{owner}
   */
  queryValidator = (owner: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetValidatorResponse, RpcStatus>({
      path: `/dcl/validator/nodes/${owner}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProposedDisableValidatorAll
   * @summary Queries a list of ProposedDisableValidator items.
   * @request GET:/dcl/validator/proposed-disable-nodes
   */
  queryProposedDisableValidatorAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllProposedDisableValidatorResponse, RpcStatus>({
      path: `/dcl/validator/proposed-disable-nodes`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryProposedDisableValidator
   * @summary Queries a ProposedDisableValidator by index.
   * @request GET:/dcl/validator/proposed-disable-nodes/{address}
   */
  queryProposedDisableValidator = (address: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetProposedDisableValidatorResponse, RpcStatus>({
      path: `/dcl/validator/proposed-disable-nodes/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRejectedDisableValidatorAll
   * @summary Queries a list of RejectedNode items.
   * @request GET:/dcl/validator/rejected-disable-nodes
   */
  queryRejectedDisableValidatorAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllRejectedDisableValidatorResponse, RpcStatus>({
      path: `/dcl/validator/rejected-disable-nodes`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryRejectedDisableValidator
   * @summary Queries a RejectedNode by index.
   * @request GET:/dcl/validator/rejected-disable-nodes/{owner}
   */
  queryRejectedDisableValidator = (owner: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetRejectedDisableValidatorResponse, RpcStatus>({
      path: `/dcl/validator/rejected-disable-nodes/${owner}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
