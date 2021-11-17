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

  /** reverse is set to true if results are to be returned in the descending order. */
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

export interface ValidatorDescription {
  name?: string;
  identity?: string;
  website?: string;
  details?: string;
}

export interface ValidatorLastValidatorPower {
  consensusAddress?: string;

  /** @format int32 */
  power?: number;
}

export type ValidatorMsgCreateValidatorResponse = object;

export interface ValidatorQueryAllLastValidatorPowerResponse {
  lastValidatorPower?: ValidatorLastValidatorPower[];

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

export interface ValidatorQueryAllValidatorMissedBlockBitArrayResponse {
  validatorMissedBlockBitArray?: ValidatorValidatorMissedBlockBitArray[];

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

export interface ValidatorQueryAllValidatorOwnerResponse {
  validatorOwner?: ValidatorValidatorOwner[];

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
  validator?: ValidatorValidator[];

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

export interface ValidatorQueryAllValidatorSigningInfoResponse {
  validatorSigningInfo?: ValidatorValidatorSigningInfo[];

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

export interface ValidatorQueryGetLastValidatorPowerResponse {
  lastValidatorPower?: ValidatorLastValidatorPower;
}

export interface ValidatorQueryGetValidatorMissedBlockBitArrayResponse {
  validatorMissedBlockBitArray?: ValidatorValidatorMissedBlockBitArray;
}

export interface ValidatorQueryGetValidatorOwnerResponse {
  validatorOwner?: ValidatorValidatorOwner;
}

export interface ValidatorQueryGetValidatorResponse {
  validator?: ValidatorValidator;
}

export interface ValidatorQueryGetValidatorSigningInfoResponse {
  validatorSigningInfo?: ValidatorValidatorSigningInfo;
}

export interface ValidatorValidator {
  address?: string;
  description?: ValidatorDescription;
  pubKey?: string;

  /** @format int32 */
  power?: number;
  jailed?: boolean;
  jailedReason?: string;
  owner?: string;
}

export interface ValidatorValidatorMissedBlockBitArray {
  address?: string;

  /** @format uint64 */
  index?: string;
}

export interface ValidatorValidatorOwner {
  address?: string;
}

export interface ValidatorValidatorSigningInfo {
  address?: string;

  /** @format uint64 */
  startHeight?: string;

  /** @format uint64 */
  indexOffset?: string;

  /** @format uint64 */
  missedBlocksCounter?: string;
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
 * @title validator/description.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryLastValidatorPowerAll
   * @summary Queries a list of lastValidatorPower items.
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/lastValidatorPower
   */
  queryLastValidatorPowerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllLastValidatorPowerResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/lastValidatorPower`,
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
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/lastValidatorPower/{consensusAddress}
   */
  queryLastValidatorPower = (consensusAddress: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetLastValidatorPowerResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/lastValidatorPower/${consensusAddress}`,
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
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validator
   */
  queryValidatorAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllValidatorResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validator`,
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
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validator/{address}
   */
  queryValidator = (address: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetValidatorResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validator/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidatorMissedBlockBitArrayAll
   * @summary Queries a list of validatorMissedBlockBitArray items.
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorMissedBlockBitArray
   */
  queryValidatorMissedBlockBitArrayAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllValidatorMissedBlockBitArrayResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validatorMissedBlockBitArray`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidatorMissedBlockBitArray
   * @summary Queries a validatorMissedBlockBitArray by index.
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorMissedBlockBitArray/{address}/{index}
   */
  queryValidatorMissedBlockBitArray = (address: string, index: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetValidatorMissedBlockBitArrayResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validatorMissedBlockBitArray/${address}/${index}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidatorOwnerAll
   * @summary Queries a list of validatorOwner items.
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorOwner
   */
  queryValidatorOwnerAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllValidatorOwnerResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validatorOwner`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidatorOwner
   * @summary Queries a validatorOwner by index.
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorOwner/{address}
   */
  queryValidatorOwner = (address: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetValidatorOwnerResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validatorOwner/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidatorSigningInfoAll
   * @summary Queries a list of validatorSigningInfo items.
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorSigningInfo
   */
  queryValidatorSigningInfoAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.countTotal"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<ValidatorQueryAllValidatorSigningInfoResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validatorSigningInfo`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryValidatorSigningInfo
   * @summary Queries a validatorSigningInfo by index.
   * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorSigningInfo/{address}
   */
  queryValidatorSigningInfo = (address: string, params: RequestParams = {}) =>
    this.request<ValidatorQueryGetValidatorSigningInfoResponse, RpcStatus>({
      path: `/zigbee-alliance/distributedcomplianceledger/validator/validatorSigningInfo/${address}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
