export interface DclauthAccount {
    address?: string;
    roles?: string[];
    /** @format uint64 */
    vendorID?: string;
}
export interface DclauthAccountStat {
    /** @format uint64 */
    number?: string;
}
export declare type DclauthMsgApproveAddAccountResponse = object;
export declare type DclauthMsgApproveRevokeAccountResponse = object;
export declare type DclauthMsgProposeAddAccountResponse = object;
export declare type DclauthMsgProposeRevokeAccountResponse = object;
export interface DclauthPendingAccount {
    address?: string;
    approvals?: string[];
}
export interface DclauthPendingAccountRevocation {
    address?: string;
    approvals?: string[];
}
export interface DclauthQueryAllAccountResponse {
    account?: DclauthAccount[];
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
    pendingAccount?: DclauthPendingAccount[];
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
    pendingAccountRevocation?: DclauthPendingAccountRevocation[];
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
    account?: DclauthAccount;
}
export interface DclauthQueryGetAccountStatResponse {
    AccountStat?: DclauthAccountStat;
}
export interface DclauthQueryGetPendingAccountResponse {
    pendingAccount?: DclauthPendingAccount;
}
export interface DclauthQueryGetPendingAccountRevocationResponse {
    pendingAccountRevocation?: DclauthPendingAccountRevocation;
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
 * @title dclauth/account.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryAccountAll
     * @summary Queries a list of account items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/dclauth/account
     */
    queryAccountAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<DclauthQueryAllAccountResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryAccount
     * @summary Queries a account by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/dclauth/account/{address}
     */
    queryAccount: (address: string, params?: RequestParams) => Promise<HttpResponse<DclauthQueryGetAccountResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryAccountStat
     * @summary Queries a accountStat by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/dclauth/accountStat
     */
    queryAccountStat: (params?: RequestParams) => Promise<HttpResponse<DclauthQueryGetAccountStatResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryPendingAccountAll
     * @summary Queries a list of pendingAccount items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/dclauth/pendingAccount
     */
    queryPendingAccountAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<DclauthQueryAllPendingAccountResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryPendingAccount
     * @summary Queries a pendingAccount by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/dclauth/pendingAccount/{address}
     */
    queryPendingAccount: (address: string, params?: RequestParams) => Promise<HttpResponse<DclauthQueryGetPendingAccountResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryPendingAccountRevocationAll
     * @summary Queries a list of pendingAccountRevocation items.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/dclauth/pendingAccountRevocation
     */
    queryPendingAccountRevocationAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<DclauthQueryAllPendingAccountRevocationResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryPendingAccountRevocation
     * @summary Queries a pendingAccountRevocation by index.
     * @request GET:/zigbee-alliance/distributedcomplianceledger/dclauth/pendingAccountRevocation/{address}
     */
    queryPendingAccountRevocation: (address: string, params?: RequestParams) => Promise<HttpResponse<DclauthQueryGetPendingAccountRevocationResponse, RpcStatus>>;
}
export {};
