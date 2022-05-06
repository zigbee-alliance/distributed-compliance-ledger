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
 * @title validator/description.proto
 * @version version not set
 */
export class Api extends HttpClient {
    constructor() {
        super(...arguments);
        /**
         * No description
         *
         * @tags Query
         * @name QueryDisabledValidatorAll
         * @summary Queries a list of DisabledValidator items.
         * @request GET:/dcl/validator/disabled-nodes
         */
        this.queryDisabledValidatorAll = (query, params = {}) => this.request({
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
        this.queryDisabledValidator = (address, params = {}) => this.request({
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
        this.queryLastValidatorPowerAll = (query, params = {}) => this.request({
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
        this.queryLastValidatorPower = (owner, params = {}) => this.request({
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
        this.queryValidatorAll = (query, params = {}) => this.request({
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
        this.queryValidator = (owner, params = {}) => this.request({
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
        this.queryProposedDisableValidatorAll = (query, params = {}) => this.request({
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
        this.queryProposedDisableValidator = (address, params = {}) => this.request({
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
        this.queryRejectedDisableValidatorAll = (query, params = {}) => this.request({
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
        this.queryRejectedDisableValidator = (owner, params = {}) => this.request({
            path: `/dcl/validator/rejected-disable-nodes/${owner}`,
            method: "GET",
            format: "json",
            ...params,
        });
    }
}
