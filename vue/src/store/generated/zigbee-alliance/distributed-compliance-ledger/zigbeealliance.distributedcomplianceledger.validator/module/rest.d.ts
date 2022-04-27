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
     *
     * * If no scheme is provided, `https` is assumed.
     * * An HTTP GET on the URL must yield a [google.protobuf.Type][]
     *   value in binary format, or produce an error.
     * * Applications are allowed to cache lookup results based on the
     *   URL, or have them precompiled into a binary to avoid any
     *   lookup. Therefore, binary compatibility needs to be preserved
     *   on changes to types. (Use versioned type names to manage
     *   breaking changes.)
     *
     * Note: this functionality is not currently available in the official
     * protobuf release, and it is not used for type URLs beginning with
     * type.googleapis.com.
     *
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
export interface ValidatorDisabledValidator {
    address?: string;
    creator?: string;
    approvals?: ValidatorGrant[];
    disabledByNodeAdmin?: boolean;
    rejects?: ValidatorGrant[];
}
export interface ValidatorGrant {
    address?: string;
    /** @format int64 */
    time?: string;
    info?: string;
}
export interface ValidatorLastValidatorPower {
    owner?: string;
    /** @format int32 */
    power?: number;
}
export declare type ValidatorMsgApproveDisableValidatorResponse = object;
export declare type ValidatorMsgCreateValidatorResponse = object;
export declare type ValidatorMsgDisableValidatorResponse = object;
export declare type ValidatorMsgEnableValidatorResponse = object;
export declare type ValidatorMsgProposeDisableValidatorResponse = object;
export declare type ValidatorMsgRejectDisableValidatorResponse = object;
export interface ValidatorProposedDisableValidator {
    address?: string;
    creator?: string;
    approvals?: ValidatorGrant[];
    rejects?: ValidatorGrant[];
}
export interface ValidatorQueryAllDisabledValidatorResponse {
    disabledValidator?: ValidatorDisabledValidator[];
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
export interface ValidatorQueryAllProposedDisableValidatorResponse {
    proposedDisableValidator?: ValidatorProposedDisableValidator[];
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
    rejectedValidator?: ValidatorRejectedDisableValidator[];
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
export interface ValidatorQueryGetDisabledValidatorResponse {
    disabledValidator?: ValidatorDisabledValidator;
}
export interface ValidatorQueryGetLastValidatorPowerResponse {
    lastValidatorPower?: ValidatorLastValidatorPower;
}
export interface ValidatorQueryGetProposedDisableValidatorResponse {
    proposedDisableValidator?: ValidatorProposedDisableValidator;
}
export interface ValidatorQueryGetRejectedDisableValidatorResponse {
    rejectedValidator?: ValidatorRejectedDisableValidator;
}
export interface ValidatorQueryGetValidatorResponse {
    validator?: ValidatorValidator;
}
export interface ValidatorRejectedDisableValidator {
    address?: string;
    creator?: string;
    approvals?: ValidatorGrant[];
    rejects?: ValidatorGrant[];
}
export interface ValidatorValidator {
    owner?: string;
    description?: ValidatorDescription;
    /**
     * `Any` contains an arbitrary serialized protocol buffer message along with a
     * URL that describes the type of the serialized message.
     *
     * Protobuf library provides support to pack/unpack Any values in the form
     * of utility functions or additional generated methods of the Any type.
     *
     * Example 1: Pack and unpack a message in C++.
     *
     *     Foo foo = ...;
     *     Any any;
     *     any.PackFrom(foo);
     *     ...
     *     if (any.UnpackTo(&foo)) {
     *       ...
     *     }
     *
     * Example 2: Pack and unpack a message in Java.
     *
     *     Foo foo = ...;
     *     Any any = Any.pack(foo);
     *     ...
     *     if (any.is(Foo.class)) {
     *       foo = any.unpack(Foo.class);
     *     }
     *
     *  Example 3: Pack and unpack a message in Python.
     *
     *     foo = Foo(...)
     *     any = Any()
     *     any.Pack(foo)
     *     ...
     *     if any.Is(Foo.DESCRIPTOR):
     *       any.Unpack(foo)
     *       ...
     *
     *  Example 4: Pack and unpack a message in Go
     *
     *      foo := &pb.Foo{...}
     *      any, err := anypb.New(foo)
     *      if err != nil {
     *        ...
     *      }
     *      ...
     *      foo := &pb.Foo{}
     *      if err := any.UnmarshalTo(foo); err != nil {
     *        ...
     *      }
     *
     * The pack methods provided by protobuf library will by default use
     * 'type.googleapis.com/full.type.name' as the type URL and the unpack
     * methods only use the fully qualified type name after the last '/'
     * in the type URL, for example "foo.bar.com/x/y.z" will yield type
     * name "y.z".
     *
     *
     * JSON
     * ====
     * The JSON representation of an `Any` value uses the regular
     * representation of the deserialized, embedded message, with an
     * additional field `@type` which contains the type URL. Example:
     *
     *     package google.profile;
     *     message Person {
     *       string first_name = 1;
     *       string last_name = 2;
     *     }
     *
     *     {
     *       "@type": "type.googleapis.com/google.profile.Person",
     *       "firstName": <string>,
     *       "lastName": <string>
     *     }
     *
     * If the embedded message type is well-known and has a custom JSON
     * representation, that representation will be embedded adding a field
     * `value` which holds the custom JSON in addition to the `@type`
     * field. Example (for message [google.protobuf.Duration][]):
     *
     *     {
     *       "@type": "type.googleapis.com/google.protobuf.Duration",
     *       "value": "1.212s"
     *     }
     */
    pubKey?: ProtobufAny;
    /** @format int32 */
    power?: number;
    jailed?: boolean;
    jailedReason?: string;
}
export declare type QueryParamsType = Record<string | number, any>;
export declare type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;
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
export declare type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;
export interface ApiConfig<SecurityDataType = unknown> {
    baseUrl?: string;
    baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
    securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}
export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
    data: D;
    error: E;
}
declare type CancelToken = Symbol | string | number;
export declare enum ContentType {
    Json = "application/json",
    FormData = "multipart/form-data",
    UrlEncoded = "application/x-www-form-urlencoded"
}
export declare class HttpClient<SecurityDataType = unknown> {
    baseUrl: string;
    private securityData;
    private securityWorker;
    private abortControllers;
    private baseApiParams;
    constructor(apiConfig?: ApiConfig<SecurityDataType>);
    setSecurityData: (data: SecurityDataType) => void;
    private addQueryParam;
    protected toQueryString(rawQuery?: QueryParamsType): string;
    protected addQueryParams(rawQuery?: QueryParamsType): string;
    private contentFormatters;
    private mergeRequestParams;
    private createAbortSignal;
    abortRequest: (cancelToken: CancelToken) => void;
    request: <T = any, E = any>({ body, secure, path, type, query, format, baseUrl, cancelToken, ...params }: FullRequestParams) => Promise<HttpResponse<T, E>>;
}
/**
 * @title validator/description.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryDisabledValidatorAll
     * @summary Queries a list of DisabledValidator items.
     * @request GET:/dcl/validator/disabled-nodes
     */
    queryDisabledValidatorAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllDisabledValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryDisabledValidator
     * @summary Queries a DisabledValidator by index.
     * @request GET:/dcl/validator/disabled-nodes/{address}
     */
    queryDisabledValidator: (address: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetDisabledValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryLastValidatorPowerAll
     * @summary Queries a list of lastValidatorPower items.
     * @request GET:/dcl/validator/last-powers
     */
    queryLastValidatorPowerAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllLastValidatorPowerResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryLastValidatorPower
     * @summary Queries a lastValidatorPower by index.
     * @request GET:/dcl/validator/last-powers/{owner}
     */
    queryLastValidatorPower: (owner: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetLastValidatorPowerResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorAll
     * @summary Queries a list of validator items.
     * @request GET:/dcl/validator/nodes
     */
    queryValidatorAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidator
     * @summary Queries a validator by index.
     * @request GET:/dcl/validator/nodes/{owner}
     */
    queryValidator: (owner: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProposedDisableValidatorAll
     * @summary Queries a list of ProposedDisableValidator items.
     * @request GET:/dcl/validator/proposed-disable-nodes
     */
    queryProposedDisableValidatorAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllProposedDisableValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProposedDisableValidator
     * @summary Queries a ProposedDisableValidator by index.
     * @request GET:/dcl/validator/proposed-disable-nodes/{address}
     */
    queryProposedDisableValidator: (address: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetProposedDisableValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRejectedDisableValidatorAll
     * @summary Queries a list of RejectedNode items.
     * @request GET:/dcl/validator/rejected-disable-nodes
     */
    queryRejectedDisableValidatorAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllRejectedDisableValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRejectedDisableValidator
     * @summary Queries a RejectedNode by index.
     * @request GET:/dcl/validator/rejected-disable-nodes/{owner}
     */
    queryRejectedDisableValidator: (owner: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetRejectedDisableValidatorResponse, RpcStatus>>;
}
export {};
