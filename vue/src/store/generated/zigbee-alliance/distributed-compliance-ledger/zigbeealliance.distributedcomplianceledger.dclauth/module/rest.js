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
export var RevokedAccountReason;
(function (RevokedAccountReason) {
    RevokedAccountReason["TrusteeVoting"] = "TrusteeVoting";
    RevokedAccountReason["MaliciousValidator"] = "MaliciousValidator";
})(RevokedAccountReason || (RevokedAccountReason = {}));
export var ContentType;
(function (ContentType) {
    ContentType["Json"] = "application/json";
    ContentType["FormData"] = "multipart/form-data";
    ContentType["UrlEncoded"] = "application/x-www-form-urlencoded";
})(ContentType || (ContentType = {}));
export class HttpClient {
    constructor(apiConfig = {}) {
        this.baseUrl = "";
        this.securityData = null;
        this.securityWorker = null;
        this.abortControllers = new Map();
        this.baseApiParams = {
            credentials: "same-origin",
            headers: {},
            redirect: "follow",
            referrerPolicy: "no-referrer",
        };
        this.setSecurityData = (data) => {
            this.securityData = data;
        };
        this.contentFormatters = {
            [ContentType.Json]: (input) => input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
            [ContentType.FormData]: (input) => Object.keys(input || {}).reduce((data, key) => {
                data.append(key, input[key]);
                return data;
            }, new FormData()),
            [ContentType.UrlEncoded]: (input) => this.toQueryString(input),
        };
        this.createAbortSignal = (cancelToken) => {
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
        this.abortRequest = (cancelToken) => {
            const abortController = this.abortControllers.get(cancelToken);
            if (abortController) {
                abortController.abort();
                this.abortControllers.delete(cancelToken);
            }
        };
        this.request = ({ body, secure, path, type, query, format = "json", baseUrl, cancelToken, ...params }) => {
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
                const r = response;
                r.data = null;
                r.error = null;
                const data = await response[format]()
                    .then((data) => {
                    if (r.ok) {
                        r.data = data;
                    }
                    else {
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
                if (!response.ok)
                    throw data;
                return data;
            });
        };
        Object.assign(this, apiConfig);
    }
    addQueryParam(query, key) {
        const value = query[key];
        return (encodeURIComponent(key) +
            "=" +
            encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`));
    }
    toQueryString(rawQuery) {
        const query = rawQuery || {};
        const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
        return keys
            .map((key) => typeof query[key] === "object" && !Array.isArray(query[key])
            ? this.toQueryString(query[key])
            : this.addQueryParam(query, key))
            .join("&");
    }
    addQueryParams(rawQuery) {
        const queryString = this.toQueryString(rawQuery);
        return queryString ? `?${queryString}` : "";
    }
    mergeRequestParams(params1, params2) {
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
}
/**
 * @title dclauth/account.proto
 * @version version not set
 */
export class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryAccountAll
         * @summary Queries a list of account items.
         * @request GET:/dcl/auth/accounts
         */
        this.queryAccountAll = (query, params = {}) => this.request({
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
        this.queryAccountStat = (params = {}) => this.request({
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
        this.queryAccount = (address, params = {}) => this.request({
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
        this.queryPendingAccountAll = (query, params = {}) => this.request({
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
        this.queryPendingAccount = (address, params = {}) => this.request({
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
        this.queryPendingAccountRevocationAll = (query, params = {}) => this.request({
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
        this.queryPendingAccountRevocation = (address, params = {}) => this.request({
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
        this.queryRejectedAccountAll = (query, params = {}) => this.request({
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
        this.queryRejectedAccount = (address, params = {}) => this.request({
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
        this.queryRevokedAccountAll = (query, params = {}) => this.request({
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
        this.queryRevokedAccount = (address, params = {}) => this.request({
            path: `/dcl/auth/revoked-accounts/${address}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
