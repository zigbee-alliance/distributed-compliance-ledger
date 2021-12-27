import { Reader, Writer } from 'protobufjs/minimal';
import { ApprovedCertificates } from '../pki/approved_certificates';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { ProposedCertificate } from '../pki/proposed_certificate';
import { ChildCertificates } from '../pki/child_certificates';
import { ProposedCertificateRevocation } from '../pki/proposed_certificate_revocation';
import { RevokedCertificates } from '../pki/revoked_certificates';
import { ApprovedRootCertificates } from '../pki/approved_root_certificates';
import { RevokedRootCertificates } from '../pki/revoked_root_certificates';
import { ApprovedCertificatesBySubject } from '../pki/approved_certificates_by_subject';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface QueryGetApprovedCertificatesRequest {
    subject: string;
    subjectKeyId: string;
}
export interface QueryGetApprovedCertificatesResponse {
    approvedCertificates: ApprovedCertificates | undefined;
}
export interface QueryAllApprovedCertificatesRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllApprovedCertificatesResponse {
    approvedCertificates: ApprovedCertificates[];
    pagination: PageResponse | undefined;
}
export interface QueryGetProposedCertificateRequest {
    subject: string;
    subjectKeyId: string;
}
export interface QueryGetProposedCertificateResponse {
    proposedCertificate: ProposedCertificate | undefined;
}
export interface QueryAllProposedCertificateRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllProposedCertificateResponse {
    proposedCertificate: ProposedCertificate[];
    pagination: PageResponse | undefined;
}
export interface QueryGetChildCertificatesRequest {
    issuer: string;
    authorityKeyId: string;
}
export interface QueryGetChildCertificatesResponse {
    childCertificates: ChildCertificates | undefined;
}
export interface QueryGetProposedCertificateRevocationRequest {
    subject: string;
    subjectKeyId: string;
}
export interface QueryGetProposedCertificateRevocationResponse {
    proposedCertificateRevocation: ProposedCertificateRevocation | undefined;
}
export interface QueryAllProposedCertificateRevocationRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllProposedCertificateRevocationResponse {
    proposedCertificateRevocation: ProposedCertificateRevocation[];
    pagination: PageResponse | undefined;
}
export interface QueryGetRevokedCertificatesRequest {
    subject: string;
    subjectKeyId: string;
}
export interface QueryGetRevokedCertificatesResponse {
    revokedCertificates: RevokedCertificates | undefined;
}
export interface QueryAllRevokedCertificatesRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllRevokedCertificatesResponse {
    revokedCertificates: RevokedCertificates[];
    pagination: PageResponse | undefined;
}
export interface QueryGetApprovedRootCertificatesRequest {
}
export interface QueryGetApprovedRootCertificatesResponse {
    ApprovedRootCertificates: ApprovedRootCertificates | undefined;
}
export interface QueryGetRevokedRootCertificatesRequest {
}
export interface QueryGetRevokedRootCertificatesResponse {
    RevokedRootCertificates: RevokedRootCertificates | undefined;
}
export interface QueryGetApprovedCertificatesBySubjectRequest {
    subject: string;
}
export interface QueryGetApprovedCertificatesBySubjectResponse {
    approvedCertificatesBySubject: ApprovedCertificatesBySubject | undefined;
}
export declare const QueryGetApprovedCertificatesRequest: {
    encode(message: QueryGetApprovedCertificatesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesRequest;
    fromJSON(object: any): QueryGetApprovedCertificatesRequest;
    toJSON(message: QueryGetApprovedCertificatesRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetApprovedCertificatesRequest>): QueryGetApprovedCertificatesRequest;
};
export declare const QueryGetApprovedCertificatesResponse: {
    encode(message: QueryGetApprovedCertificatesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesResponse;
    fromJSON(object: any): QueryGetApprovedCertificatesResponse;
    toJSON(message: QueryGetApprovedCertificatesResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetApprovedCertificatesResponse>): QueryGetApprovedCertificatesResponse;
};
export declare const QueryAllApprovedCertificatesRequest: {
    encode(message: QueryAllApprovedCertificatesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllApprovedCertificatesRequest;
    fromJSON(object: any): QueryAllApprovedCertificatesRequest;
    toJSON(message: QueryAllApprovedCertificatesRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllApprovedCertificatesRequest>): QueryAllApprovedCertificatesRequest;
};
export declare const QueryAllApprovedCertificatesResponse: {
    encode(message: QueryAllApprovedCertificatesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllApprovedCertificatesResponse;
    fromJSON(object: any): QueryAllApprovedCertificatesResponse;
    toJSON(message: QueryAllApprovedCertificatesResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllApprovedCertificatesResponse>): QueryAllApprovedCertificatesResponse;
};
export declare const QueryGetProposedCertificateRequest: {
    encode(message: QueryGetProposedCertificateRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateRequest;
    fromJSON(object: any): QueryGetProposedCertificateRequest;
    toJSON(message: QueryGetProposedCertificateRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedCertificateRequest>): QueryGetProposedCertificateRequest;
};
export declare const QueryGetProposedCertificateResponse: {
    encode(message: QueryGetProposedCertificateResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateResponse;
    fromJSON(object: any): QueryGetProposedCertificateResponse;
    toJSON(message: QueryGetProposedCertificateResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedCertificateResponse>): QueryGetProposedCertificateResponse;
};
export declare const QueryAllProposedCertificateRequest: {
    encode(message: QueryAllProposedCertificateRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateRequest;
    fromJSON(object: any): QueryAllProposedCertificateRequest;
    toJSON(message: QueryAllProposedCertificateRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedCertificateRequest>): QueryAllProposedCertificateRequest;
};
export declare const QueryAllProposedCertificateResponse: {
    encode(message: QueryAllProposedCertificateResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateResponse;
    fromJSON(object: any): QueryAllProposedCertificateResponse;
    toJSON(message: QueryAllProposedCertificateResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedCertificateResponse>): QueryAllProposedCertificateResponse;
};
export declare const QueryGetChildCertificatesRequest: {
    encode(message: QueryGetChildCertificatesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChildCertificatesRequest;
    fromJSON(object: any): QueryGetChildCertificatesRequest;
    toJSON(message: QueryGetChildCertificatesRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetChildCertificatesRequest>): QueryGetChildCertificatesRequest;
};
export declare const QueryGetChildCertificatesResponse: {
    encode(message: QueryGetChildCertificatesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetChildCertificatesResponse;
    fromJSON(object: any): QueryGetChildCertificatesResponse;
    toJSON(message: QueryGetChildCertificatesResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetChildCertificatesResponse>): QueryGetChildCertificatesResponse;
};
export declare const QueryGetProposedCertificateRevocationRequest: {
    encode(message: QueryGetProposedCertificateRevocationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateRevocationRequest;
    fromJSON(object: any): QueryGetProposedCertificateRevocationRequest;
    toJSON(message: QueryGetProposedCertificateRevocationRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedCertificateRevocationRequest>): QueryGetProposedCertificateRevocationRequest;
};
export declare const QueryGetProposedCertificateRevocationResponse: {
    encode(message: QueryGetProposedCertificateRevocationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProposedCertificateRevocationResponse;
    fromJSON(object: any): QueryGetProposedCertificateRevocationResponse;
    toJSON(message: QueryGetProposedCertificateRevocationResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetProposedCertificateRevocationResponse>): QueryGetProposedCertificateRevocationResponse;
};
export declare const QueryAllProposedCertificateRevocationRequest: {
    encode(message: QueryAllProposedCertificateRevocationRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateRevocationRequest;
    fromJSON(object: any): QueryAllProposedCertificateRevocationRequest;
    toJSON(message: QueryAllProposedCertificateRevocationRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedCertificateRevocationRequest>): QueryAllProposedCertificateRevocationRequest;
};
export declare const QueryAllProposedCertificateRevocationResponse: {
    encode(message: QueryAllProposedCertificateRevocationResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProposedCertificateRevocationResponse;
    fromJSON(object: any): QueryAllProposedCertificateRevocationResponse;
    toJSON(message: QueryAllProposedCertificateRevocationResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllProposedCertificateRevocationResponse>): QueryAllProposedCertificateRevocationResponse;
};
export declare const QueryGetRevokedCertificatesRequest: {
    encode(message: QueryGetRevokedCertificatesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedCertificatesRequest;
    fromJSON(object: any): QueryGetRevokedCertificatesRequest;
    toJSON(message: QueryGetRevokedCertificatesRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetRevokedCertificatesRequest>): QueryGetRevokedCertificatesRequest;
};
export declare const QueryGetRevokedCertificatesResponse: {
    encode(message: QueryGetRevokedCertificatesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedCertificatesResponse;
    fromJSON(object: any): QueryGetRevokedCertificatesResponse;
    toJSON(message: QueryGetRevokedCertificatesResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRevokedCertificatesResponse>): QueryGetRevokedCertificatesResponse;
};
export declare const QueryAllRevokedCertificatesRequest: {
    encode(message: QueryAllRevokedCertificatesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedCertificatesRequest;
    fromJSON(object: any): QueryAllRevokedCertificatesRequest;
    toJSON(message: QueryAllRevokedCertificatesRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllRevokedCertificatesRequest>): QueryAllRevokedCertificatesRequest;
};
export declare const QueryAllRevokedCertificatesResponse: {
    encode(message: QueryAllRevokedCertificatesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllRevokedCertificatesResponse;
    fromJSON(object: any): QueryAllRevokedCertificatesResponse;
    toJSON(message: QueryAllRevokedCertificatesResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllRevokedCertificatesResponse>): QueryAllRevokedCertificatesResponse;
};
export declare const QueryGetApprovedRootCertificatesRequest: {
    encode(_: QueryGetApprovedRootCertificatesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedRootCertificatesRequest;
    fromJSON(_: any): QueryGetApprovedRootCertificatesRequest;
    toJSON(_: QueryGetApprovedRootCertificatesRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetApprovedRootCertificatesRequest>): QueryGetApprovedRootCertificatesRequest;
};
export declare const QueryGetApprovedRootCertificatesResponse: {
    encode(message: QueryGetApprovedRootCertificatesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedRootCertificatesResponse;
    fromJSON(object: any): QueryGetApprovedRootCertificatesResponse;
    toJSON(message: QueryGetApprovedRootCertificatesResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetApprovedRootCertificatesResponse>): QueryGetApprovedRootCertificatesResponse;
};
export declare const QueryGetRevokedRootCertificatesRequest: {
    encode(_: QueryGetRevokedRootCertificatesRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedRootCertificatesRequest;
    fromJSON(_: any): QueryGetRevokedRootCertificatesRequest;
    toJSON(_: QueryGetRevokedRootCertificatesRequest): unknown;
    fromPartial(_: DeepPartial<QueryGetRevokedRootCertificatesRequest>): QueryGetRevokedRootCertificatesRequest;
};
export declare const QueryGetRevokedRootCertificatesResponse: {
    encode(message: QueryGetRevokedRootCertificatesResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetRevokedRootCertificatesResponse;
    fromJSON(object: any): QueryGetRevokedRootCertificatesResponse;
    toJSON(message: QueryGetRevokedRootCertificatesResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetRevokedRootCertificatesResponse>): QueryGetRevokedRootCertificatesResponse;
};
export declare const QueryGetApprovedCertificatesBySubjectRequest: {
    encode(message: QueryGetApprovedCertificatesBySubjectRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesBySubjectRequest;
    fromJSON(object: any): QueryGetApprovedCertificatesBySubjectRequest;
    toJSON(message: QueryGetApprovedCertificatesBySubjectRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetApprovedCertificatesBySubjectRequest>): QueryGetApprovedCertificatesBySubjectRequest;
};
export declare const QueryGetApprovedCertificatesBySubjectResponse: {
    encode(message: QueryGetApprovedCertificatesBySubjectResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetApprovedCertificatesBySubjectResponse;
    fromJSON(object: any): QueryGetApprovedCertificatesBySubjectResponse;
    toJSON(message: QueryGetApprovedCertificatesBySubjectResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetApprovedCertificatesBySubjectResponse>): QueryGetApprovedCertificatesBySubjectResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a ApprovedCertificates by index. */
    ApprovedCertificates(request: QueryGetApprovedCertificatesRequest): Promise<QueryGetApprovedCertificatesResponse>;
    /** Queries a list of ApprovedCertificates items. */
    ApprovedCertificatesAll(request: QueryAllApprovedCertificatesRequest): Promise<QueryAllApprovedCertificatesResponse>;
    /** Queries a ProposedCertificate by index. */
    ProposedCertificate(request: QueryGetProposedCertificateRequest): Promise<QueryGetProposedCertificateResponse>;
    /** Queries a list of ProposedCertificate items. */
    ProposedCertificateAll(request: QueryAllProposedCertificateRequest): Promise<QueryAllProposedCertificateResponse>;
    /** Queries a ChildCertificates by index. */
    ChildCertificates(request: QueryGetChildCertificatesRequest): Promise<QueryGetChildCertificatesResponse>;
    /** Queries a ProposedCertificateRevocation by index. */
    ProposedCertificateRevocation(request: QueryGetProposedCertificateRevocationRequest): Promise<QueryGetProposedCertificateRevocationResponse>;
    /** Queries a list of ProposedCertificateRevocation items. */
    ProposedCertificateRevocationAll(request: QueryAllProposedCertificateRevocationRequest): Promise<QueryAllProposedCertificateRevocationResponse>;
    /** Queries a RevokedCertificates by index. */
    RevokedCertificates(request: QueryGetRevokedCertificatesRequest): Promise<QueryGetRevokedCertificatesResponse>;
    /** Queries a list of RevokedCertificates items. */
    RevokedCertificatesAll(request: QueryAllRevokedCertificatesRequest): Promise<QueryAllRevokedCertificatesResponse>;
    /** Queries a ApprovedRootCertificates by index. */
    ApprovedRootCertificates(request: QueryGetApprovedRootCertificatesRequest): Promise<QueryGetApprovedRootCertificatesResponse>;
    /** Queries a RevokedRootCertificates by index. */
    RevokedRootCertificates(request: QueryGetRevokedRootCertificatesRequest): Promise<QueryGetRevokedRootCertificatesResponse>;
    /** Queries a ApprovedCertificatesBySubject by index. */
    ApprovedCertificatesBySubject(request: QueryGetApprovedCertificatesBySubjectRequest): Promise<QueryGetApprovedCertificatesBySubjectResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    ApprovedCertificates(request: QueryGetApprovedCertificatesRequest): Promise<QueryGetApprovedCertificatesResponse>;
    ApprovedCertificatesAll(request: QueryAllApprovedCertificatesRequest): Promise<QueryAllApprovedCertificatesResponse>;
    ProposedCertificate(request: QueryGetProposedCertificateRequest): Promise<QueryGetProposedCertificateResponse>;
    ProposedCertificateAll(request: QueryAllProposedCertificateRequest): Promise<QueryAllProposedCertificateResponse>;
    ChildCertificates(request: QueryGetChildCertificatesRequest): Promise<QueryGetChildCertificatesResponse>;
    ProposedCertificateRevocation(request: QueryGetProposedCertificateRevocationRequest): Promise<QueryGetProposedCertificateRevocationResponse>;
    ProposedCertificateRevocationAll(request: QueryAllProposedCertificateRevocationRequest): Promise<QueryAllProposedCertificateRevocationResponse>;
    RevokedCertificates(request: QueryGetRevokedCertificatesRequest): Promise<QueryGetRevokedCertificatesResponse>;
    RevokedCertificatesAll(request: QueryAllRevokedCertificatesRequest): Promise<QueryAllRevokedCertificatesResponse>;
    ApprovedRootCertificates(request: QueryGetApprovedRootCertificatesRequest): Promise<QueryGetApprovedRootCertificatesResponse>;
    RevokedRootCertificates(request: QueryGetRevokedRootCertificatesRequest): Promise<QueryGetRevokedRootCertificatesResponse>;
    ApprovedCertificatesBySubject(request: QueryGetApprovedCertificatesBySubjectRequest): Promise<QueryGetApprovedCertificatesBySubjectResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
