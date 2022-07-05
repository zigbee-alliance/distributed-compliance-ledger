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
export declare type ComplianceMsgCertifyModelResponse = object;
export declare type ComplianceMsgProvisionModelResponse = object;
export declare type ComplianceMsgRevokeModelResponse = object;
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
 * @title compliance/certified_model.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryCertifiedModelAll
     * @summary Queries a list of CertifiedModel items.
     * @request GET:/dcl/compliance/certified-models
     */
    queryCertifiedModelAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryAllCertifiedModelResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryCertifiedModel
     * @summary Queries a CertifiedModel by index.
     * @request GET:/dcl/compliance/certified-models/{vid}/{pid}/{softwareVersion}/{certificationType}
     */
    queryCertifiedModel: (vid: number, pid: number, softwareVersion: number, certificationType: string, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryGetCertifiedModelResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryComplianceInfoAll
     * @summary Queries a list of ComplianceInfo items.
     * @request GET:/dcl/compliance/compliance-info
     */
    queryComplianceInfoAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryAllComplianceInfoResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryComplianceInfo
     * @summary Queries a ComplianceInfo by index.
     * @request GET:/dcl/compliance/compliance-info/{vid}/{pid}/{softwareVersion}/{certificationType}
     */
    queryComplianceInfo: (vid: number, pid: number, softwareVersion: number, certificationType: string, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryGetComplianceInfoResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryDeviceSoftwareComplianceAll
     * @summary Queries a list of DeviceSoftwareCompliance items.
     * @request GET:/dcl/compliance/device-software-compliance
     */
    queryDeviceSoftwareComplianceAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryAllDeviceSoftwareComplianceResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryDeviceSoftwareCompliance
     * @summary Queries a DeviceSoftwareCompliance by index.
     * @request GET:/dcl/compliance/device-software-compliance/{cDCertificateId}
     */
    queryDeviceSoftwareCompliance: (cDCertificateId: string, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryGetDeviceSoftwareComplianceResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProvisionalModelAll
     * @summary Queries a list of ProvisionalModel items.
     * @request GET:/dcl/compliance/provisional-models
     */
    queryProvisionalModelAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryAllProvisionalModelResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProvisionalModel
     * @summary Queries a ProvisionalModel by index.
     * @request GET:/dcl/compliance/provisional-models/{vid}/{pid}/{softwareVersion}/{certificationType}
     */
    queryProvisionalModel: (vid: number, pid: number, softwareVersion: number, certificationType: string, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryGetProvisionalModelResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRevokedModelAll
     * @summary Queries a list of RevokedModel items.
     * @request GET:/dcl/compliance/revoked-models
     */
    queryRevokedModelAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryAllRevokedModelResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRevokedModel
     * @summary Queries a RevokedModel by index.
     * @request GET:/dcl/compliance/revoked-models/{vid}/{pid}/{softwareVersion}/{certificationType}
     */
    queryRevokedModel: (vid: number, pid: number, softwareVersion: number, certificationType: string, params?: RequestParams) => Promise<HttpResponse<ComplianceQueryGetRevokedModelResponse, RpcStatus>>;
}
export {};
