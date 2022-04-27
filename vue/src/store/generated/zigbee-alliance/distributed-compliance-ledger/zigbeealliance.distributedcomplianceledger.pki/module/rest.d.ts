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
export declare type PkiMsgAddX509CertResponse = object;
export declare type PkiMsgApproveAddX509RootCertResponse = object;
export declare type PkiMsgApproveRevokeX509RootCertResponse = object;
export declare type PkiMsgProposeAddX509RootCertResponse = object;
export declare type PkiMsgProposeRevokeX509RootCertResponse = object;
export declare type PkiMsgRejectAddX509RootCertResponse = object;
export declare type PkiMsgRevokeX509CertResponse = object;
export interface PkiProposedCertificate {
    subject?: string;
    subjectKeyId?: string;
    pemCert?: string;
    serialNumber?: string;
    owner?: string;
    approvals?: PkiGrant[];
    subjectAsText?: string;
    rejects?: PkiGrant[];
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
 * @title pki/approved_certificates.proto
 * @version version not set
 */
export declare class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
    /**
     * No description
     *
     * @tags Query
     * @name QueryApprovedCertificatesAll
     * @summary Queries a list of ApprovedCertificates items.
     * @request GET:/dcl/pki/certificates
     */
    queryApprovedCertificatesAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<PkiQueryAllApprovedCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryApprovedCertificatesBySubject
     * @summary Queries a ApprovedCertificatesBySubject by index.
     * @request GET:/dcl/pki/certificates/{subject}
     */
    queryApprovedCertificatesBySubject: (subject: string, params?: RequestParams) => Promise<HttpResponse<PkiQueryGetApprovedCertificatesBySubjectResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryApprovedCertificates
     * @summary Queries a ApprovedCertificates by index.
     * @request GET:/dcl/pki/certificates/{subject}/{subjectKeyId}
     */
    queryApprovedCertificates: (subject: string, subjectKeyId: string, params?: RequestParams) => Promise<HttpResponse<PkiQueryGetApprovedCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryChildCertificates
     * @summary Queries a ChildCertificates by index.
     * @request GET:/dcl/pki/child-certificates/{issuer}/{authorityKeyId}
     */
    queryChildCertificates: (issuer: string, authorityKeyId: string, params?: RequestParams) => Promise<HttpResponse<PkiQueryGetChildCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProposedCertificateAll
     * @summary Queries a list of ProposedCertificate items.
     * @request GET:/dcl/pki/proposed-certificates
     */
    queryProposedCertificateAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<PkiQueryAllProposedCertificateResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProposedCertificate
     * @summary Queries a ProposedCertificate by index.
     * @request GET:/dcl/pki/proposed-certificates/{subject}/{subjectKeyId}
     */
    queryProposedCertificate: (subject: string, subjectKeyId: string, params?: RequestParams) => Promise<HttpResponse<PkiQueryGetProposedCertificateResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProposedCertificateRevocationAll
     * @summary Queries a list of ProposedCertificateRevocation items.
     * @request GET:/dcl/pki/proposed-revocation-certificates
     */
    queryProposedCertificateRevocationAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<PkiQueryAllProposedCertificateRevocationResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryProposedCertificateRevocation
     * @summary Queries a ProposedCertificateRevocation by index.
     * @request GET:/dcl/pki/proposed-revocation-certificates/{subject}/{subjectKeyId}
     */
    queryProposedCertificateRevocation: (subject: string, subjectKeyId: string, params?: RequestParams) => Promise<HttpResponse<PkiQueryGetProposedCertificateRevocationResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRejectedCertificateAll
     * @summary Queries a list of RejectedCertificate items.
     * @request GET:/dcl/pki/rejected-certificates
     */
    queryRejectedCertificateAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<PkiQueryAllRejectedCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRejectedCertificate
     * @summary Queries a RejectedCertificate by index.
     * @request GET:/dcl/pki/rejected-certificates/{subject}/{subjectKeyId}
     */
    queryRejectedCertificate: (subject: string, subjectKeyId: string, params?: RequestParams) => Promise<HttpResponse<PkiQueryGetRejectedCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRevokedCertificatesAll
     * @summary Queries a list of RevokedCertificates items.
     * @request GET:/dcl/pki/revoked-certificates
     */
    queryRevokedCertificatesAll: (query?: {
        "pagination.key"?: string;
        "pagination.offset"?: string;
        "pagination.limit"?: string;
        "pagination.countTotal"?: boolean;
        "pagination.reverse"?: boolean;
    }, params?: RequestParams) => Promise<HttpResponse<PkiQueryAllRevokedCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRevokedCertificates
     * @summary Queries a RevokedCertificates by index.
     * @request GET:/dcl/pki/revoked-certificates/{subject}/{subjectKeyId}
     */
    queryRevokedCertificates: (subject: string, subjectKeyId: string, params?: RequestParams) => Promise<HttpResponse<PkiQueryGetRevokedCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryRevokedRootCertificates
     * @summary Queries a RevokedRootCertificates by index.
     * @request GET:/dcl/pki/revoked-root-certificates
     */
    queryRevokedRootCertificates: (params?: RequestParams) => Promise<HttpResponse<PkiQueryGetRevokedRootCertificatesResponse, RpcStatus>>;
    /**
     * No description
     *
     * @tags Query
     * @name QueryApprovedRootCertificates
     * @summary Queries a ApprovedRootCertificates by index.
     * @request GET:/dcl/pki/root-certificates
     */
    queryApprovedRootCertificates: (params?: RequestParams) => Promise<HttpResponse<PkiQueryGetApprovedRootCertificatesResponse, RpcStatus>>;
}
export {};
