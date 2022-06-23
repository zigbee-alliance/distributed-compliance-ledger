import { Reader, Writer } from 'protobufjs/minimal';
import { ComplianceInfo } from '../compliance/compliance_info';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { CertifiedModel } from '../compliance/certified_model';
import { RevokedModel } from '../compliance/revoked_model';
import { ProvisionalModel } from '../compliance/provisional_model';
import { DeviceSoftwareCompliance } from '../compliance/device_software_compliance';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";
export interface QueryGetComplianceInfoRequest {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
}
export interface QueryGetComplianceInfoResponse {
    complianceInfo: ComplianceInfo | undefined;
}
export interface QueryAllComplianceInfoRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllComplianceInfoResponse {
    complianceInfo: ComplianceInfo[];
    pagination: PageResponse | undefined;
}
export interface QueryGetCertifiedModelRequest {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
}
export interface QueryGetCertifiedModelResponse {
    certifiedModel: CertifiedModel | undefined;
}
export interface QueryAllCertifiedModelRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllCertifiedModelResponse {
    certifiedModel: CertifiedModel[];
    pagination: PageResponse | undefined;
}
export interface QueryGetRevokedModelRequest {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
}
export interface QueryGetRevokedModelResponse {
    revokedModel: RevokedModel | undefined;
}
export interface QueryAllRevokedModelRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllRevokedModelResponse {
    revokedModel: RevokedModel[];
    pagination: PageResponse | undefined;
}
export interface QueryGetProvisionalModelRequest {
    vid: number;
    pid: number;
    softwareVersion: number;
    certificationType: string;
}
export interface QueryGetProvisionalModelResponse {
    provisionalModel: ProvisionalModel | undefined;
}
export interface QueryAllProvisionalModelRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllProvisionalModelResponse {
    provisionalModel: ProvisionalModel[];
    pagination: PageResponse | undefined;
}
export interface QueryGetDeviceSoftwareComplianceRequest {
    cDCertificateId: string;
}
export interface QueryGetDeviceSoftwareComplianceResponse {
    deviceSoftwareCompliance: DeviceSoftwareCompliance | undefined;
}
export interface QueryAllDeviceSoftwareComplianceRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllDeviceSoftwareComplianceResponse {
    deviceSoftwareCompliance: DeviceSoftwareCompliance[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetComplianceInfoRequest: {
    encode(message: QueryGetComplianceInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoRequest;
    fromJSON(object: any): QueryGetComplianceInfoRequest;
    toJSON(message: QueryGetComplianceInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetComplianceInfoRequest>): QueryGetComplianceInfoRequest;
};
export declare const QueryGetComplianceInfoResponse: {
    encode(message: QueryGetComplianceInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetComplianceInfoResponse;
    fromJSON(object: any): QueryGetComplianceInfoResponse;
    toJSON(message: QueryGetComplianceInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetComplianceInfoResponse>): QueryGetComplianceInfoResponse;
};
export declare const QueryAllComplianceInfoRequest: {
    encode(message: QueryAllComplianceInfoRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoRequest;
    fromJSON(object: any): QueryAllComplianceInfoRequest;
    toJSON(message: QueryAllComplianceInfoRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllComplianceInfoRequest>): QueryAllComplianceInfoRequest;
};
export declare const QueryAllComplianceInfoResponse: {
    encode(message: QueryAllComplianceInfoResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllComplianceInfoResponse;
    fromJSON(object: any): QueryAllComplianceInfoResponse;
    toJSON(message: QueryAllComplianceInfoResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllComplianceInfoResponse>): QueryAllComplianceInfoResponse;
};
export declare const QueryGetCertifiedModelRequest: {
    encode(message: QueryGetCertifiedModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCertifiedModelRequest;
    fromJSON(object: any): QueryGetCertifiedModelRequest;
    toJSON(message: QueryGetCertifiedModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetCertifiedModelRequest>): QueryGetCertifiedModelRequest;
};
export declare const QueryGetCertifiedModelResponse: {
    encode(message: QueryGetCertifiedModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetCertifiedModelResponse;
    fromJSON(object: any): QueryGetCertifiedModelResponse;
    toJSON(message: QueryGetCertifiedModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetCertifiedModelResponse>): QueryGetCertifiedModelResponse;
};
export declare const QueryAllCertifiedModelRequest: {
    encode(message: QueryAllCertifiedModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCertifiedModelRequest;
    fromJSON(object: any): QueryAllCertifiedModelRequest;
    toJSON(message: QueryAllCertifiedModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllCertifiedModelRequest>): QueryAllCertifiedModelRequest;
};
export declare const QueryAllCertifiedModelResponse: {
    encode(message: QueryAllCertifiedModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllCertifiedModelResponse;
    fromJSON(object: any): QueryAllCertifiedModelResponse;
    toJSON(message: QueryAllCertifiedModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllCertifiedModelResponse>): QueryAllCertifiedModelResponse;
};
export declare const QueryGetRevokedModelRequest: {
    encode(message: QueryGetRevokedModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedModelRequest;
    fromJSON(object: any): QueryGetRevokedModelRequest;
    toJSON(message: QueryGetRevokedModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRevokedModelRequest>): QueryGetRevokedModelRequest;
};
export declare const QueryGetRevokedModelResponse: {
    encode(message: QueryGetRevokedModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedModelResponse;
    fromJSON(object: any): QueryGetRevokedModelResponse;
    toJSON(message: QueryGetRevokedModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRevokedModelResponse>): QueryGetRevokedModelResponse;
};
export declare const QueryAllRevokedModelRequest: {
    encode(message: QueryAllRevokedModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedModelRequest;
    fromJSON(object: any): QueryAllRevokedModelRequest;
    toJSON(message: QueryAllRevokedModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRevokedModelRequest>): QueryAllRevokedModelRequest;
};
export declare const QueryAllRevokedModelResponse: {
    encode(message: QueryAllRevokedModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedModelResponse;
    fromJSON(object: any): QueryAllRevokedModelResponse;
    toJSON(message: QueryAllRevokedModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRevokedModelResponse>): QueryAllRevokedModelResponse;
};
export declare const QueryGetProvisionalModelRequest: {
    encode(message: QueryGetProvisionalModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProvisionalModelRequest;
    fromJSON(object: any): QueryGetProvisionalModelRequest;
    toJSON(message: QueryGetProvisionalModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetProvisionalModelRequest>): QueryGetProvisionalModelRequest;
};
export declare const QueryGetProvisionalModelResponse: {
    encode(message: QueryGetProvisionalModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProvisionalModelResponse;
    fromJSON(object: any): QueryGetProvisionalModelResponse;
    toJSON(message: QueryGetProvisionalModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetProvisionalModelResponse>): QueryGetProvisionalModelResponse;
};
export declare const QueryAllProvisionalModelRequest: {
    encode(message: QueryAllProvisionalModelRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProvisionalModelRequest;
    fromJSON(object: any): QueryAllProvisionalModelRequest;
    toJSON(message: QueryAllProvisionalModelRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllProvisionalModelRequest>): QueryAllProvisionalModelRequest;
};
export declare const QueryAllProvisionalModelResponse: {
    encode(message: QueryAllProvisionalModelResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProvisionalModelResponse;
    fromJSON(object: any): QueryAllProvisionalModelResponse;
    toJSON(message: QueryAllProvisionalModelResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllProvisionalModelResponse>): QueryAllProvisionalModelResponse;
};
export declare const QueryGetDeviceSoftwareComplianceRequest: {
    encode(message: QueryGetDeviceSoftwareComplianceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDeviceSoftwareComplianceRequest;
    fromJSON(object: any): QueryGetDeviceSoftwareComplianceRequest;
    toJSON(message: QueryGetDeviceSoftwareComplianceRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetDeviceSoftwareComplianceRequest>): QueryGetDeviceSoftwareComplianceRequest;
};
export declare const QueryGetDeviceSoftwareComplianceResponse: {
    encode(message: QueryGetDeviceSoftwareComplianceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetDeviceSoftwareComplianceResponse;
    fromJSON(object: any): QueryGetDeviceSoftwareComplianceResponse;
    toJSON(message: QueryGetDeviceSoftwareComplianceResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetDeviceSoftwareComplianceResponse>): QueryGetDeviceSoftwareComplianceResponse;
};
export declare const QueryAllDeviceSoftwareComplianceRequest: {
    encode(message: QueryAllDeviceSoftwareComplianceRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDeviceSoftwareComplianceRequest;
    fromJSON(object: any): QueryAllDeviceSoftwareComplianceRequest;
    toJSON(message: QueryAllDeviceSoftwareComplianceRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllDeviceSoftwareComplianceRequest>): QueryAllDeviceSoftwareComplianceRequest;
};
export declare const QueryAllDeviceSoftwareComplianceResponse: {
    encode(message: QueryAllDeviceSoftwareComplianceResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllDeviceSoftwareComplianceResponse;
    fromJSON(object: any): QueryAllDeviceSoftwareComplianceResponse;
    toJSON(message: QueryAllDeviceSoftwareComplianceResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllDeviceSoftwareComplianceResponse>): QueryAllDeviceSoftwareComplianceResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a ComplianceInfo by index. */
    ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse>;
    /** Queries a list of ComplianceInfo items. */
    ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse>;
    /** Queries a CertifiedModel by index. */
    CertifiedModel(request: QueryGetCertifiedModelRequest): Promise<QueryGetCertifiedModelResponse>;
    /** Queries a list of CertifiedModel items. */
    CertifiedModelAll(request: QueryAllCertifiedModelRequest): Promise<QueryAllCertifiedModelResponse>;
    /** Queries a RevokedModel by index. */
    RevokedModel(request: QueryGetRevokedModelRequest): Promise<QueryGetRevokedModelResponse>;
    /** Queries a list of RevokedModel items. */
    RevokedModelAll(request: QueryAllRevokedModelRequest): Promise<QueryAllRevokedModelResponse>;
    /** Queries a ProvisionalModel by index. */
    ProvisionalModel(request: QueryGetProvisionalModelRequest): Promise<QueryGetProvisionalModelResponse>;
    /** Queries a list of ProvisionalModel items. */
    ProvisionalModelAll(request: QueryAllProvisionalModelRequest): Promise<QueryAllProvisionalModelResponse>;
    /** Queries a DeviceSoftwareCompliance by index. */
    DeviceSoftwareCompliance(request: QueryGetDeviceSoftwareComplianceRequest): Promise<QueryGetDeviceSoftwareComplianceResponse>;
    /** Queries a list of DeviceSoftwareCompliance items. */
    DeviceSoftwareComplianceAll(request: QueryAllDeviceSoftwareComplianceRequest): Promise<QueryAllDeviceSoftwareComplianceResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    ComplianceInfo(request: QueryGetComplianceInfoRequest): Promise<QueryGetComplianceInfoResponse>;
    ComplianceInfoAll(request: QueryAllComplianceInfoRequest): Promise<QueryAllComplianceInfoResponse>;
    CertifiedModel(request: QueryGetCertifiedModelRequest): Promise<QueryGetCertifiedModelResponse>;
    CertifiedModelAll(request: QueryAllCertifiedModelRequest): Promise<QueryAllCertifiedModelResponse>;
    RevokedModel(request: QueryGetRevokedModelRequest): Promise<QueryGetRevokedModelResponse>;
    RevokedModelAll(request: QueryAllRevokedModelRequest): Promise<QueryAllRevokedModelResponse>;
    ProvisionalModel(request: QueryGetProvisionalModelRequest): Promise<QueryGetProvisionalModelResponse>;
    ProvisionalModelAll(request: QueryAllProvisionalModelRequest): Promise<QueryAllProvisionalModelResponse>;
    DeviceSoftwareCompliance(request: QueryGetDeviceSoftwareComplianceRequest): Promise<QueryGetDeviceSoftwareComplianceResponse>;
    DeviceSoftwareComplianceAll(request: QueryAllDeviceSoftwareComplianceRequest): Promise<QueryAllDeviceSoftwareComplianceResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
