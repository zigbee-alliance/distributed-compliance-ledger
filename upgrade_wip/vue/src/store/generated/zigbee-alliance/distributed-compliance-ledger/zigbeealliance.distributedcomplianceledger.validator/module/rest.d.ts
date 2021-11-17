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
export declare type ValidatorMsgCreateValidatorResponse = object;
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
     * @name QueryLastValidatorPowerAll
     * @summary Queries a list of lastValidatorPower items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/lastValidatorPower
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
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/lastValidatorPower/{consensusAddress}
     */
    queryLastValidatorPower: (consensusAddress: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetLastValidatorPowerResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorAll
     * @summary Queries a list of validator items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validator
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
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validator/{address}
     */
    queryValidator: (address: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetValidatorResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorMissedBlockBitArrayAll
     * @summary Queries a list of validatorMissedBlockBitArray items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorMissedBlockBitArray
     */
    queryValidatorMissedBlockBitArrayAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllValidatorMissedBlockBitArrayResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorMissedBlockBitArray
     * @summary Queries a validatorMissedBlockBitArray by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorMissedBlockBitArray/{address}/{index}
     */
    queryValidatorMissedBlockBitArray: (address: string, index: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetValidatorMissedBlockBitArrayResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorOwnerAll
     * @summary Queries a list of validatorOwner items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorOwner
     */
    queryValidatorOwnerAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllValidatorOwnerResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorOwner
     * @summary Queries a validatorOwner by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorOwner/{address}
     */
    queryValidatorOwner: (address: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetValidatorOwnerResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorSigningInfoAll
     * @summary Queries a list of validatorSigningInfo items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorSigningInfo
     */
    queryValidatorSigningInfoAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryAllValidatorSigningInfoResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryValidatorSigningInfo
     * @summary Queries a validatorSigningInfo by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/validator/validatorSigningInfo/{address}
     */
    queryValidatorSigningInfo: (address: string, params?: RequestParams) => Promise<HttpResponse<ValidatorQueryGetValidatorSigningInfoResponse, RpcStatus>>;
}
export {};
