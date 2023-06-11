import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";
export interface MsgProposeAddX509RootCert {
    signer: string;
    cert: string;
    info: string;
    time: number;
}
export interface MsgProposeAddX509RootCertResponse {
}
export interface MsgApproveAddX509RootCert {
    signer: string;
    subject: string;
    subjectKeyId: string;
    info: string;
    time: number;
}
export interface MsgApproveAddX509RootCertResponse {
}
export interface MsgAddX509Cert {
    signer: string;
    cert: string;
    info: string;
    time: number;
}
export interface MsgAddX509CertResponse {
}
export interface MsgProposeRevokeX509RootCert {
    signer: string;
    subject: string;
    subjectKeyId: string;
    info: string;
    time: number;
}
export interface MsgProposeRevokeX509RootCertResponse {
}
export interface MsgApproveRevokeX509RootCert {
    signer: string;
    subject: string;
    subjectKeyId: string;
    info: string;
    time: number;
}
export interface MsgApproveRevokeX509RootCertResponse {
}
export interface MsgRevokeX509Cert {
    signer: string;
    subject: string;
    subjectKeyId: string;
    info: string;
    time: number;
}
export interface MsgRevokeX509CertResponse {
}
export interface MsgRejectAddX509RootCert {
    signer: string;
    subject: string;
    subjectKeyId: string;
    info: string;
    time: number;
}
export interface MsgRejectAddX509RootCertResponse {
}
export interface MsgAddPkiRevocationDistributionPoint {
    signer: string;
    vid: number;
    pid: number;
    isPAA: boolean;
    label: string;
    crlSignerCertificate: string;
    issuerSubjectKeyID: string;
    dataURL: string;
    dataFileSize: number;
    dataDigest: string;
    dataDigestType: number;
    revocationType: number;
}
export interface MsgAddPkiRevocationDistributionPointResponse {
}
export interface MsgUpdatePkiRevocationDistributionPoint {
    signer: string;
    vid: number;
    label: string;
    crlSignerCertificate: string;
    issuerSubjectKeyID: string;
    dataURL: string;
    dataFileSize: number;
    dataDigest: string;
    dataDigestType: number;
}
export interface MsgUpdatePkiRevocationDistributionPointResponse {
}
export interface MsgDeletePkiRevocationDistributionPoint {
    signer: string;
    vid: number;
    label: string;
    issuerSubjectKeyID: string;
}
export interface MsgDeletePkiRevocationDistributionPointResponse {
}
export declare const MsgProposeAddX509RootCert: {
    encode(message: MsgProposeAddX509RootCert, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeAddX509RootCert;
    fromJSON(object: any): MsgProposeAddX509RootCert;
    toJSON(message: MsgProposeAddX509RootCert): unknown;
    fromPartial(object: DeepPartial<MsgProposeAddX509RootCert>): MsgProposeAddX509RootCert;
};
export declare const MsgProposeAddX509RootCertResponse: {
    encode(_: MsgProposeAddX509RootCertResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeAddX509RootCertResponse;
    fromJSON(_: any): MsgProposeAddX509RootCertResponse;
    toJSON(_: MsgProposeAddX509RootCertResponse): unknown;
    fromPartial(_: DeepPartial<MsgProposeAddX509RootCertResponse>): MsgProposeAddX509RootCertResponse;
};
export declare const MsgApproveAddX509RootCert: {
    encode(message: MsgApproveAddX509RootCert, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveAddX509RootCert;
    fromJSON(object: any): MsgApproveAddX509RootCert;
    toJSON(message: MsgApproveAddX509RootCert): unknown;
    fromPartial(object: DeepPartial<MsgApproveAddX509RootCert>): MsgApproveAddX509RootCert;
};
export declare const MsgApproveAddX509RootCertResponse: {
    encode(_: MsgApproveAddX509RootCertResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveAddX509RootCertResponse;
    fromJSON(_: any): MsgApproveAddX509RootCertResponse;
    toJSON(_: MsgApproveAddX509RootCertResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveAddX509RootCertResponse>): MsgApproveAddX509RootCertResponse;
};
export declare const MsgAddX509Cert: {
    encode(message: MsgAddX509Cert, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddX509Cert;
    fromJSON(object: any): MsgAddX509Cert;
    toJSON(message: MsgAddX509Cert): unknown;
    fromPartial(object: DeepPartial<MsgAddX509Cert>): MsgAddX509Cert;
};
export declare const MsgAddX509CertResponse: {
    encode(_: MsgAddX509CertResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddX509CertResponse;
    fromJSON(_: any): MsgAddX509CertResponse;
    toJSON(_: MsgAddX509CertResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddX509CertResponse>): MsgAddX509CertResponse;
};
export declare const MsgProposeRevokeX509RootCert: {
    encode(message: MsgProposeRevokeX509RootCert, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeX509RootCert;
    fromJSON(object: any): MsgProposeRevokeX509RootCert;
    toJSON(message: MsgProposeRevokeX509RootCert): unknown;
    fromPartial(object: DeepPartial<MsgProposeRevokeX509RootCert>): MsgProposeRevokeX509RootCert;
};
export declare const MsgProposeRevokeX509RootCertResponse: {
    encode(_: MsgProposeRevokeX509RootCertResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgProposeRevokeX509RootCertResponse;
    fromJSON(_: any): MsgProposeRevokeX509RootCertResponse;
    toJSON(_: MsgProposeRevokeX509RootCertResponse): unknown;
    fromPartial(_: DeepPartial<MsgProposeRevokeX509RootCertResponse>): MsgProposeRevokeX509RootCertResponse;
};
export declare const MsgApproveRevokeX509RootCert: {
    encode(message: MsgApproveRevokeX509RootCert, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeX509RootCert;
    fromJSON(object: any): MsgApproveRevokeX509RootCert;
    toJSON(message: MsgApproveRevokeX509RootCert): unknown;
    fromPartial(object: DeepPartial<MsgApproveRevokeX509RootCert>): MsgApproveRevokeX509RootCert;
};
export declare const MsgApproveRevokeX509RootCertResponse: {
    encode(_: MsgApproveRevokeX509RootCertResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgApproveRevokeX509RootCertResponse;
    fromJSON(_: any): MsgApproveRevokeX509RootCertResponse;
    toJSON(_: MsgApproveRevokeX509RootCertResponse): unknown;
    fromPartial(_: DeepPartial<MsgApproveRevokeX509RootCertResponse>): MsgApproveRevokeX509RootCertResponse;
};
export declare const MsgRevokeX509Cert: {
    encode(message: MsgRevokeX509Cert, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRevokeX509Cert;
    fromJSON(object: any): MsgRevokeX509Cert;
    toJSON(message: MsgRevokeX509Cert): unknown;
    fromPartial(object: DeepPartial<MsgRevokeX509Cert>): MsgRevokeX509Cert;
};
export declare const MsgRevokeX509CertResponse: {
    encode(_: MsgRevokeX509CertResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRevokeX509CertResponse;
    fromJSON(_: any): MsgRevokeX509CertResponse;
    toJSON(_: MsgRevokeX509CertResponse): unknown;
    fromPartial(_: DeepPartial<MsgRevokeX509CertResponse>): MsgRevokeX509CertResponse;
};
export declare const MsgRejectAddX509RootCert: {
    encode(message: MsgRejectAddX509RootCert, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectAddX509RootCert;
    fromJSON(object: any): MsgRejectAddX509RootCert;
    toJSON(message: MsgRejectAddX509RootCert): unknown;
    fromPartial(object: DeepPartial<MsgRejectAddX509RootCert>): MsgRejectAddX509RootCert;
};
export declare const MsgRejectAddX509RootCertResponse: {
    encode(_: MsgRejectAddX509RootCertResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRejectAddX509RootCertResponse;
    fromJSON(_: any): MsgRejectAddX509RootCertResponse;
    toJSON(_: MsgRejectAddX509RootCertResponse): unknown;
    fromPartial(_: DeepPartial<MsgRejectAddX509RootCertResponse>): MsgRejectAddX509RootCertResponse;
};
export declare const MsgAddPkiRevocationDistributionPoint: {
    encode(message: MsgAddPkiRevocationDistributionPoint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddPkiRevocationDistributionPoint;
    fromJSON(object: any): MsgAddPkiRevocationDistributionPoint;
    toJSON(message: MsgAddPkiRevocationDistributionPoint): unknown;
    fromPartial(object: DeepPartial<MsgAddPkiRevocationDistributionPoint>): MsgAddPkiRevocationDistributionPoint;
};
export declare const MsgAddPkiRevocationDistributionPointResponse: {
    encode(_: MsgAddPkiRevocationDistributionPointResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAddPkiRevocationDistributionPointResponse;
    fromJSON(_: any): MsgAddPkiRevocationDistributionPointResponse;
    toJSON(_: MsgAddPkiRevocationDistributionPointResponse): unknown;
    fromPartial(_: DeepPartial<MsgAddPkiRevocationDistributionPointResponse>): MsgAddPkiRevocationDistributionPointResponse;
};
export declare const MsgUpdatePkiRevocationDistributionPoint: {
    encode(message: MsgUpdatePkiRevocationDistributionPoint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePkiRevocationDistributionPoint;
    fromJSON(object: any): MsgUpdatePkiRevocationDistributionPoint;
    toJSON(message: MsgUpdatePkiRevocationDistributionPoint): unknown;
    fromPartial(object: DeepPartial<MsgUpdatePkiRevocationDistributionPoint>): MsgUpdatePkiRevocationDistributionPoint;
};
export declare const MsgUpdatePkiRevocationDistributionPointResponse: {
    encode(_: MsgUpdatePkiRevocationDistributionPointResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdatePkiRevocationDistributionPointResponse;
    fromJSON(_: any): MsgUpdatePkiRevocationDistributionPointResponse;
    toJSON(_: MsgUpdatePkiRevocationDistributionPointResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdatePkiRevocationDistributionPointResponse>): MsgUpdatePkiRevocationDistributionPointResponse;
};
export declare const MsgDeletePkiRevocationDistributionPoint: {
    encode(message: MsgDeletePkiRevocationDistributionPoint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePkiRevocationDistributionPoint;
    fromJSON(object: any): MsgDeletePkiRevocationDistributionPoint;
    toJSON(message: MsgDeletePkiRevocationDistributionPoint): unknown;
    fromPartial(object: DeepPartial<MsgDeletePkiRevocationDistributionPoint>): MsgDeletePkiRevocationDistributionPoint;
};
export declare const MsgDeletePkiRevocationDistributionPointResponse: {
    encode(_: MsgDeletePkiRevocationDistributionPointResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeletePkiRevocationDistributionPointResponse;
    fromJSON(_: any): MsgDeletePkiRevocationDistributionPointResponse;
    toJSON(_: MsgDeletePkiRevocationDistributionPointResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeletePkiRevocationDistributionPointResponse>): MsgDeletePkiRevocationDistributionPointResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    ProposeAddX509RootCert(request: MsgProposeAddX509RootCert): Promise<MsgProposeAddX509RootCertResponse>;
    ApproveAddX509RootCert(request: MsgApproveAddX509RootCert): Promise<MsgApproveAddX509RootCertResponse>;
    AddX509Cert(request: MsgAddX509Cert): Promise<MsgAddX509CertResponse>;
    ProposeRevokeX509RootCert(request: MsgProposeRevokeX509RootCert): Promise<MsgProposeRevokeX509RootCertResponse>;
    ApproveRevokeX509RootCert(request: MsgApproveRevokeX509RootCert): Promise<MsgApproveRevokeX509RootCertResponse>;
    RevokeX509Cert(request: MsgRevokeX509Cert): Promise<MsgRevokeX509CertResponse>;
    RejectAddX509RootCert(request: MsgRejectAddX509RootCert): Promise<MsgRejectAddX509RootCertResponse>;
    AddPkiRevocationDistributionPoint(request: MsgAddPkiRevocationDistributionPoint): Promise<MsgAddPkiRevocationDistributionPointResponse>;
    UpdatePkiRevocationDistributionPoint(request: MsgUpdatePkiRevocationDistributionPoint): Promise<MsgUpdatePkiRevocationDistributionPointResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeletePkiRevocationDistributionPoint(request: MsgDeletePkiRevocationDistributionPoint): Promise<MsgDeletePkiRevocationDistributionPointResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    ProposeAddX509RootCert(request: MsgProposeAddX509RootCert): Promise<MsgProposeAddX509RootCertResponse>;
    ApproveAddX509RootCert(request: MsgApproveAddX509RootCert): Promise<MsgApproveAddX509RootCertResponse>;
    AddX509Cert(request: MsgAddX509Cert): Promise<MsgAddX509CertResponse>;
    ProposeRevokeX509RootCert(request: MsgProposeRevokeX509RootCert): Promise<MsgProposeRevokeX509RootCertResponse>;
    ApproveRevokeX509RootCert(request: MsgApproveRevokeX509RootCert): Promise<MsgApproveRevokeX509RootCertResponse>;
    RevokeX509Cert(request: MsgRevokeX509Cert): Promise<MsgRevokeX509CertResponse>;
    RejectAddX509RootCert(request: MsgRejectAddX509RootCert): Promise<MsgRejectAddX509RootCertResponse>;
    AddPkiRevocationDistributionPoint(request: MsgAddPkiRevocationDistributionPoint): Promise<MsgAddPkiRevocationDistributionPointResponse>;
    UpdatePkiRevocationDistributionPoint(request: MsgUpdatePkiRevocationDistributionPoint): Promise<MsgUpdatePkiRevocationDistributionPointResponse>;
    DeletePkiRevocationDistributionPoint(request: MsgDeletePkiRevocationDistributionPoint): Promise<MsgDeletePkiRevocationDistributionPointResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
