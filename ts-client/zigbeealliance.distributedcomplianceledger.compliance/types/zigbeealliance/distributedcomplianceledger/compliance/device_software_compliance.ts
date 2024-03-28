/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { ComplianceInfo } from "./compliance_info";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";

export interface DeviceSoftwareCompliance {
  cDCertificateId: string;
  complianceInfo: ComplianceInfo[];
}

function createBaseDeviceSoftwareCompliance(): DeviceSoftwareCompliance {
  return { cDCertificateId: "", complianceInfo: [] };
}

export const DeviceSoftwareCompliance = {
  encode(message: DeviceSoftwareCompliance, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.cDCertificateId !== "") {
      writer.uint32(10).string(message.cDCertificateId);
    }
    for (const v of message.complianceInfo) {
      ComplianceInfo.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeviceSoftwareCompliance {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeviceSoftwareCompliance();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cDCertificateId = reader.string();
          break;
        case 2:
          message.complianceInfo.push(ComplianceInfo.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeviceSoftwareCompliance {
    return {
      cDCertificateId: isSet(object.cDCertificateId) ? String(object.cDCertificateId) : "",
      complianceInfo: Array.isArray(object?.complianceInfo)
        ? object.complianceInfo.map((e: any) => ComplianceInfo.fromJSON(e))
        : [],
    };
  },

  toJSON(message: DeviceSoftwareCompliance): unknown {
    const obj: any = {};
    message.cDCertificateId !== undefined && (obj.cDCertificateId = message.cDCertificateId);
    if (message.complianceInfo) {
      obj.complianceInfo = message.complianceInfo.map((e) => e ? ComplianceInfo.toJSON(e) : undefined);
    } else {
      obj.complianceInfo = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeviceSoftwareCompliance>, I>>(object: I): DeviceSoftwareCompliance {
    const message = createBaseDeviceSoftwareCompliance();
    message.cDCertificateId = object.cDCertificateId ?? "";
    message.complianceInfo = object.complianceInfo?.map((e) => ComplianceInfo.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
