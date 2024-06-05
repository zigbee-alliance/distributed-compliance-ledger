/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { CertifiedModel } from "./certified_model";
import { ComplianceInfo } from "./compliance_info";
import { DeviceSoftwareCompliance } from "./device_software_compliance";
import { ProvisionalModel } from "./provisional_model";
import { RevokedModel } from "./revoked_model";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.compliance";

/** GenesisState defines the compliance module's genesis state. */
export interface GenesisState {
  complianceInfoList: ComplianceInfo[];
  certifiedModelList: CertifiedModel[];
  revokedModelList: RevokedModel[];
  provisionalModelList: ProvisionalModel[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  deviceSoftwareComplianceList: DeviceSoftwareCompliance[];
}

function createBaseGenesisState(): GenesisState {
  return {
    complianceInfoList: [],
    certifiedModelList: [],
    revokedModelList: [],
    provisionalModelList: [],
    deviceSoftwareComplianceList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.complianceInfoList) {
      ComplianceInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.certifiedModelList) {
      CertifiedModel.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.revokedModelList) {
      RevokedModel.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.provisionalModelList) {
      ProvisionalModel.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.deviceSoftwareComplianceList) {
      DeviceSoftwareCompliance.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.complianceInfoList.push(ComplianceInfo.decode(reader, reader.uint32()));
          break;
        case 2:
          message.certifiedModelList.push(CertifiedModel.decode(reader, reader.uint32()));
          break;
        case 3:
          message.revokedModelList.push(RevokedModel.decode(reader, reader.uint32()));
          break;
        case 4:
          message.provisionalModelList.push(ProvisionalModel.decode(reader, reader.uint32()));
          break;
        case 5:
          message.deviceSoftwareComplianceList.push(DeviceSoftwareCompliance.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      complianceInfoList: Array.isArray(object?.complianceInfoList)
        ? object.complianceInfoList.map((e: any) => ComplianceInfo.fromJSON(e))
        : [],
      certifiedModelList: Array.isArray(object?.certifiedModelList)
        ? object.certifiedModelList.map((e: any) => CertifiedModel.fromJSON(e))
        : [],
      revokedModelList: Array.isArray(object?.revokedModelList)
        ? object.revokedModelList.map((e: any) => RevokedModel.fromJSON(e))
        : [],
      provisionalModelList: Array.isArray(object?.provisionalModelList)
        ? object.provisionalModelList.map((e: any) => ProvisionalModel.fromJSON(e))
        : [],
      deviceSoftwareComplianceList: Array.isArray(object?.deviceSoftwareComplianceList)
        ? object.deviceSoftwareComplianceList.map((e: any) => DeviceSoftwareCompliance.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.complianceInfoList) {
      obj.complianceInfoList = message.complianceInfoList.map((e) => e ? ComplianceInfo.toJSON(e) : undefined);
    } else {
      obj.complianceInfoList = [];
    }
    if (message.certifiedModelList) {
      obj.certifiedModelList = message.certifiedModelList.map((e) => e ? CertifiedModel.toJSON(e) : undefined);
    } else {
      obj.certifiedModelList = [];
    }
    if (message.revokedModelList) {
      obj.revokedModelList = message.revokedModelList.map((e) => e ? RevokedModel.toJSON(e) : undefined);
    } else {
      obj.revokedModelList = [];
    }
    if (message.provisionalModelList) {
      obj.provisionalModelList = message.provisionalModelList.map((e) => e ? ProvisionalModel.toJSON(e) : undefined);
    } else {
      obj.provisionalModelList = [];
    }
    if (message.deviceSoftwareComplianceList) {
      obj.deviceSoftwareComplianceList = message.deviceSoftwareComplianceList.map((e) =>
        e ? DeviceSoftwareCompliance.toJSON(e) : undefined
      );
    } else {
      obj.deviceSoftwareComplianceList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.complianceInfoList = object.complianceInfoList?.map((e) => ComplianceInfo.fromPartial(e)) || [];
    message.certifiedModelList = object.certifiedModelList?.map((e) => CertifiedModel.fromPartial(e)) || [];
    message.revokedModelList = object.revokedModelList?.map((e) => RevokedModel.fromPartial(e)) || [];
    message.provisionalModelList = object.provisionalModelList?.map((e) => ProvisionalModel.fromPartial(e)) || [];
    message.deviceSoftwareComplianceList =
      object.deviceSoftwareComplianceList?.map((e) => DeviceSoftwareCompliance.fromPartial(e)) || [];
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
