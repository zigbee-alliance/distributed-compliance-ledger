/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.model";

export interface Model {
  vid: number;
  pid: number;
  deviceTypeId: number;
  productName: string;
  productLabel: string;
  partNumber: string;
  commissioningCustomFlow: number;
  commissioningCustomFlowUrl: string;
  commissioningModeInitialStepsHint: number;
  commissioningModeInitialStepsInstruction: string;
  commissioningModeSecondaryStepsHint: number;
  commissioningModeSecondaryStepsInstruction: string;
  userManualUrl: string;
  supportUrl: string;
  productUrl: string;
  lsfUrl: string;
  lsfRevision: number;
  creator: string;
  schemaVersion: number;
  enhancedSetupFlowOptions: number;
  enhancedSetupFlowTCUrl: string;
  enhancedSetupFlowTCRevision: number;
  enhancedSetupFlowTCDigest: string;
  enhancedSetupFlowTCFileSize: number;
  maintenanceUrl: string;
  discoveryCapabilitiesBitmask: number;
  commissioningFallbackUrl: string;
  icdUserActiveModeTriggerHint: number;
  icdUserActiveModeTriggerInstruction: string;
  factoryResetStepsHint: number;
  factoryResetStepsInstruction: string;
}

function createBaseModel(): Model {
  return {
    vid: 0,
    pid: 0,
    deviceTypeId: 0,
    productName: "",
    productLabel: "",
    partNumber: "",
    commissioningCustomFlow: 0,
    commissioningCustomFlowUrl: "",
    commissioningModeInitialStepsHint: 0,
    commissioningModeInitialStepsInstruction: "",
    commissioningModeSecondaryStepsHint: 0,
    commissioningModeSecondaryStepsInstruction: "",
    userManualUrl: "",
    supportUrl: "",
    productUrl: "",
    lsfUrl: "",
    lsfRevision: 0,
    creator: "",
    schemaVersion: 0,
    enhancedSetupFlowOptions: 0,
    enhancedSetupFlowTCUrl: "",
    enhancedSetupFlowTCRevision: 0,
    enhancedSetupFlowTCDigest: "",
    enhancedSetupFlowTCFileSize: 0,
    maintenanceUrl: "",
    discoveryCapabilitiesBitmask: 0,
    commissioningFallbackUrl: "",
    icdUserActiveModeTriggerHint: 0,
    icdUserActiveModeTriggerInstruction: "",
    factoryResetStepsHint: 0,
    factoryResetStepsInstruction: "",
  };
}

export const Model = {
  encode(message: Model, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.vid !== 0) {
      writer.uint32(8).int32(message.vid);
    }
    if (message.pid !== 0) {
      writer.uint32(16).int32(message.pid);
    }
    if (message.deviceTypeId !== 0) {
      writer.uint32(24).int32(message.deviceTypeId);
    }
    if (message.productName !== "") {
      writer.uint32(34).string(message.productName);
    }
    if (message.productLabel !== "") {
      writer.uint32(42).string(message.productLabel);
    }
    if (message.partNumber !== "") {
      writer.uint32(50).string(message.partNumber);
    }
    if (message.commissioningCustomFlow !== 0) {
      writer.uint32(56).int32(message.commissioningCustomFlow);
    }
    if (message.commissioningCustomFlowUrl !== "") {
      writer.uint32(66).string(message.commissioningCustomFlowUrl);
    }
    if (message.commissioningModeInitialStepsHint !== 0) {
      writer.uint32(72).uint32(message.commissioningModeInitialStepsHint);
    }
    if (message.commissioningModeInitialStepsInstruction !== "") {
      writer.uint32(82).string(message.commissioningModeInitialStepsInstruction);
    }
    if (message.commissioningModeSecondaryStepsHint !== 0) {
      writer.uint32(88).uint32(message.commissioningModeSecondaryStepsHint);
    }
    if (message.commissioningModeSecondaryStepsInstruction !== "") {
      writer.uint32(98).string(message.commissioningModeSecondaryStepsInstruction);
    }
    if (message.userManualUrl !== "") {
      writer.uint32(106).string(message.userManualUrl);
    }
    if (message.supportUrl !== "") {
      writer.uint32(114).string(message.supportUrl);
    }
    if (message.productUrl !== "") {
      writer.uint32(122).string(message.productUrl);
    }
    if (message.lsfUrl !== "") {
      writer.uint32(130).string(message.lsfUrl);
    }
    if (message.lsfRevision !== 0) {
      writer.uint32(136).int32(message.lsfRevision);
    }
    if (message.creator !== "") {
      writer.uint32(146).string(message.creator);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(152).uint32(message.schemaVersion);
    }
    if (message.enhancedSetupFlowOptions !== 0) {
      writer.uint32(160).int32(message.enhancedSetupFlowOptions);
    }
    if (message.enhancedSetupFlowTCUrl !== "") {
      writer.uint32(170).string(message.enhancedSetupFlowTCUrl);
    }
    if (message.enhancedSetupFlowTCRevision !== 0) {
      writer.uint32(176).int32(message.enhancedSetupFlowTCRevision);
    }
    if (message.enhancedSetupFlowTCDigest !== "") {
      writer.uint32(186).string(message.enhancedSetupFlowTCDigest);
    }
    if (message.enhancedSetupFlowTCFileSize !== 0) {
      writer.uint32(192).uint32(message.enhancedSetupFlowTCFileSize);
    }
    if (message.maintenanceUrl !== "") {
      writer.uint32(202).string(message.maintenanceUrl);
    }
    if (message.discoveryCapabilitiesBitmask !== 0) {
      writer.uint32(208).uint32(message.discoveryCapabilitiesBitmask);
    }
    if (message.commissioningFallbackUrl !== "") {
      writer.uint32(218).string(message.commissioningFallbackUrl);
    }
    if (message.icdUserActiveModeTriggerHint !== 0) {
      writer.uint32(224).uint32(message.icdUserActiveModeTriggerHint);
    }
    if (message.icdUserActiveModeTriggerInstruction !== "") {
      writer.uint32(234).string(message.icdUserActiveModeTriggerInstruction);
    }
    if (message.factoryResetStepsHint !== 0) {
      writer.uint32(240).uint32(message.factoryResetStepsHint);
    }
    if (message.factoryResetStepsInstruction !== "") {
      writer.uint32(250).string(message.factoryResetStepsInstruction);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Model {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseModel();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vid = reader.int32();
          break;
        case 2:
          message.pid = reader.int32();
          break;
        case 3:
          message.deviceTypeId = reader.int32();
          break;
        case 4:
          message.productName = reader.string();
          break;
        case 5:
          message.productLabel = reader.string();
          break;
        case 6:
          message.partNumber = reader.string();
          break;
        case 7:
          message.commissioningCustomFlow = reader.int32();
          break;
        case 8:
          message.commissioningCustomFlowUrl = reader.string();
          break;
        case 9:
          message.commissioningModeInitialStepsHint = reader.uint32();
          break;
        case 10:
          message.commissioningModeInitialStepsInstruction = reader.string();
          break;
        case 11:
          message.commissioningModeSecondaryStepsHint = reader.uint32();
          break;
        case 12:
          message.commissioningModeSecondaryStepsInstruction = reader.string();
          break;
        case 13:
          message.userManualUrl = reader.string();
          break;
        case 14:
          message.supportUrl = reader.string();
          break;
        case 15:
          message.productUrl = reader.string();
          break;
        case 16:
          message.lsfUrl = reader.string();
          break;
        case 17:
          message.lsfRevision = reader.int32();
          break;
        case 18:
          message.creator = reader.string();
          break;
        case 19:
          message.schemaVersion = reader.uint32();
          break;
        case 20:
          message.enhancedSetupFlowOptions = reader.int32();
          break;
        case 21:
          message.enhancedSetupFlowTCUrl = reader.string();
          break;
        case 22:
          message.enhancedSetupFlowTCRevision = reader.int32();
          break;
        case 23:
          message.enhancedSetupFlowTCDigest = reader.string();
          break;
        case 24:
          message.enhancedSetupFlowTCFileSize = reader.uint32();
          break;
        case 25:
          message.maintenanceUrl = reader.string();
          break;
        case 26:
          message.discoveryCapabilitiesBitmask = reader.uint32();
          break;
        case 27:
          message.commissioningFallbackUrl = reader.string();
          break;
        case 28:
          message.icdUserActiveModeTriggerHint = reader.uint32();
          break;
        case 29:
          message.icdUserActiveModeTriggerInstruction = reader.string();
          break;
        case 30:
          message.factoryResetStepsHint = reader.uint32();
          break;
        case 31:
          message.factoryResetStepsInstruction = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Model {
    return {
      vid: isSet(object.vid) ? Number(object.vid) : 0,
      pid: isSet(object.pid) ? Number(object.pid) : 0,
      deviceTypeId: isSet(object.deviceTypeId) ? Number(object.deviceTypeId) : 0,
      productName: isSet(object.productName) ? String(object.productName) : "",
      productLabel: isSet(object.productLabel) ? String(object.productLabel) : "",
      partNumber: isSet(object.partNumber) ? String(object.partNumber) : "",
      commissioningCustomFlow: isSet(object.commissioningCustomFlow) ? Number(object.commissioningCustomFlow) : 0,
      commissioningCustomFlowUrl: isSet(object.commissioningCustomFlowUrl)
        ? String(object.commissioningCustomFlowUrl)
        : "",
      commissioningModeInitialStepsHint: isSet(object.commissioningModeInitialStepsHint)
        ? Number(object.commissioningModeInitialStepsHint)
        : 0,
      commissioningModeInitialStepsInstruction: isSet(object.commissioningModeInitialStepsInstruction)
        ? String(object.commissioningModeInitialStepsInstruction)
        : "",
      commissioningModeSecondaryStepsHint: isSet(object.commissioningModeSecondaryStepsHint)
        ? Number(object.commissioningModeSecondaryStepsHint)
        : 0,
      commissioningModeSecondaryStepsInstruction: isSet(object.commissioningModeSecondaryStepsInstruction)
        ? String(object.commissioningModeSecondaryStepsInstruction)
        : "",
      userManualUrl: isSet(object.userManualUrl) ? String(object.userManualUrl) : "",
      supportUrl: isSet(object.supportUrl) ? String(object.supportUrl) : "",
      productUrl: isSet(object.productUrl) ? String(object.productUrl) : "",
      lsfUrl: isSet(object.lsfUrl) ? String(object.lsfUrl) : "",
      lsfRevision: isSet(object.lsfRevision) ? Number(object.lsfRevision) : 0,
      creator: isSet(object.creator) ? String(object.creator) : "",
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
      enhancedSetupFlowOptions: isSet(object.enhancedSetupFlowOptions) ? Number(object.enhancedSetupFlowOptions) : 0,
      enhancedSetupFlowTCUrl: isSet(object.enhancedSetupFlowTCUrl) ? String(object.enhancedSetupFlowTCUrl) : "",
      enhancedSetupFlowTCRevision: isSet(object.enhancedSetupFlowTCRevision)
        ? Number(object.enhancedSetupFlowTCRevision)
        : 0,
      enhancedSetupFlowTCDigest: isSet(object.enhancedSetupFlowTCDigest)
        ? String(object.enhancedSetupFlowTCDigest)
        : "",
      enhancedSetupFlowTCFileSize: isSet(object.enhancedSetupFlowTCFileSize)
        ? Number(object.enhancedSetupFlowTCFileSize)
        : 0,
      maintenanceUrl: isSet(object.maintenanceUrl) ? String(object.maintenanceUrl) : "",
      discoveryCapabilitiesBitmask: isSet(object.discoveryCapabilitiesBitmask)
        ? Number(object.discoveryCapabilitiesBitmask)
        : 0,
      commissioningFallbackUrl: isSet(object.commissioningFallbackUrl) ? String(object.commissioningFallbackUrl) : "",
      icdUserActiveModeTriggerHint: isSet(object.icdUserActiveModeTriggerHint)
        ? Number(object.icdUserActiveModeTriggerHint)
        : 0,
      icdUserActiveModeTriggerInstruction: isSet(object.icdUserActiveModeTriggerInstruction)
        ? String(object.icdUserActiveModeTriggerInstruction)
        : "",
      factoryResetStepsHint: isSet(object.factoryResetStepsHint) ? Number(object.factoryResetStepsHint) : 0,
      factoryResetStepsInstruction: isSet(object.factoryResetStepsInstruction)
        ? String(object.factoryResetStepsInstruction)
        : "",
    };
  },

  toJSON(message: Model): unknown {
    const obj: any = {};
    message.vid !== undefined && (obj.vid = Math.round(message.vid));
    message.pid !== undefined && (obj.pid = Math.round(message.pid));
    message.deviceTypeId !== undefined && (obj.deviceTypeId = Math.round(message.deviceTypeId));
    message.productName !== undefined && (obj.productName = message.productName);
    message.productLabel !== undefined && (obj.productLabel = message.productLabel);
    message.partNumber !== undefined && (obj.partNumber = message.partNumber);
    message.commissioningCustomFlow !== undefined
      && (obj.commissioningCustomFlow = Math.round(message.commissioningCustomFlow));
    message.commissioningCustomFlowUrl !== undefined
      && (obj.commissioningCustomFlowUrl = message.commissioningCustomFlowUrl);
    message.commissioningModeInitialStepsHint !== undefined
      && (obj.commissioningModeInitialStepsHint = Math.round(message.commissioningModeInitialStepsHint));
    message.commissioningModeInitialStepsInstruction !== undefined
      && (obj.commissioningModeInitialStepsInstruction = message.commissioningModeInitialStepsInstruction);
    message.commissioningModeSecondaryStepsHint !== undefined
      && (obj.commissioningModeSecondaryStepsHint = Math.round(message.commissioningModeSecondaryStepsHint));
    message.commissioningModeSecondaryStepsInstruction !== undefined
      && (obj.commissioningModeSecondaryStepsInstruction = message.commissioningModeSecondaryStepsInstruction);
    message.userManualUrl !== undefined && (obj.userManualUrl = message.userManualUrl);
    message.supportUrl !== undefined && (obj.supportUrl = message.supportUrl);
    message.productUrl !== undefined && (obj.productUrl = message.productUrl);
    message.lsfUrl !== undefined && (obj.lsfUrl = message.lsfUrl);
    message.lsfRevision !== undefined && (obj.lsfRevision = Math.round(message.lsfRevision));
    message.creator !== undefined && (obj.creator = message.creator);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    message.enhancedSetupFlowOptions !== undefined
      && (obj.enhancedSetupFlowOptions = Math.round(message.enhancedSetupFlowOptions));
    message.enhancedSetupFlowTCUrl !== undefined && (obj.enhancedSetupFlowTCUrl = message.enhancedSetupFlowTCUrl);
    message.enhancedSetupFlowTCRevision !== undefined
      && (obj.enhancedSetupFlowTCRevision = Math.round(message.enhancedSetupFlowTCRevision));
    message.enhancedSetupFlowTCDigest !== undefined
      && (obj.enhancedSetupFlowTCDigest = message.enhancedSetupFlowTCDigest);
    message.enhancedSetupFlowTCFileSize !== undefined
      && (obj.enhancedSetupFlowTCFileSize = Math.round(message.enhancedSetupFlowTCFileSize));
    message.maintenanceUrl !== undefined && (obj.maintenanceUrl = message.maintenanceUrl);
    message.discoveryCapabilitiesBitmask !== undefined
      && (obj.discoveryCapabilitiesBitmask = Math.round(message.discoveryCapabilitiesBitmask));
    message.commissioningFallbackUrl !== undefined && (obj.commissioningFallbackUrl = message.commissioningFallbackUrl);
    message.icdUserActiveModeTriggerHint !== undefined
      && (obj.icdUserActiveModeTriggerHint = Math.round(message.icdUserActiveModeTriggerHint));
    message.icdUserActiveModeTriggerInstruction !== undefined
      && (obj.icdUserActiveModeTriggerInstruction = message.icdUserActiveModeTriggerInstruction);
    message.factoryResetStepsHint !== undefined
      && (obj.factoryResetStepsHint = Math.round(message.factoryResetStepsHint));
    message.factoryResetStepsInstruction !== undefined
      && (obj.factoryResetStepsInstruction = message.factoryResetStepsInstruction);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Model>, I>>(object: I): Model {
    const message = createBaseModel();
    message.vid = object.vid ?? 0;
    message.pid = object.pid ?? 0;
    message.deviceTypeId = object.deviceTypeId ?? 0;
    message.productName = object.productName ?? "";
    message.productLabel = object.productLabel ?? "";
    message.partNumber = object.partNumber ?? "";
    message.commissioningCustomFlow = object.commissioningCustomFlow ?? 0;
    message.commissioningCustomFlowUrl = object.commissioningCustomFlowUrl ?? "";
    message.commissioningModeInitialStepsHint = object.commissioningModeInitialStepsHint ?? 0;
    message.commissioningModeInitialStepsInstruction = object.commissioningModeInitialStepsInstruction ?? "";
    message.commissioningModeSecondaryStepsHint = object.commissioningModeSecondaryStepsHint ?? 0;
    message.commissioningModeSecondaryStepsInstruction = object.commissioningModeSecondaryStepsInstruction ?? "";
    message.userManualUrl = object.userManualUrl ?? "";
    message.supportUrl = object.supportUrl ?? "";
    message.productUrl = object.productUrl ?? "";
    message.lsfUrl = object.lsfUrl ?? "";
    message.lsfRevision = object.lsfRevision ?? 0;
    message.creator = object.creator ?? "";
    message.schemaVersion = object.schemaVersion ?? 0;
    message.enhancedSetupFlowOptions = object.enhancedSetupFlowOptions ?? 0;
    message.enhancedSetupFlowTCUrl = object.enhancedSetupFlowTCUrl ?? "";
    message.enhancedSetupFlowTCRevision = object.enhancedSetupFlowTCRevision ?? 0;
    message.enhancedSetupFlowTCDigest = object.enhancedSetupFlowTCDigest ?? "";
    message.enhancedSetupFlowTCFileSize = object.enhancedSetupFlowTCFileSize ?? 0;
    message.maintenanceUrl = object.maintenanceUrl ?? "";
    message.discoveryCapabilitiesBitmask = object.discoveryCapabilitiesBitmask ?? 0;
    message.commissioningFallbackUrl = object.commissioningFallbackUrl ?? "";
    message.icdUserActiveModeTriggerHint = object.icdUserActiveModeTriggerHint ?? 0;
    message.icdUserActiveModeTriggerInstruction = object.icdUserActiveModeTriggerInstruction ?? "";
    message.factoryResetStepsHint = object.factoryResetStepsHint ?? 0;
    message.factoryResetStepsInstruction = object.factoryResetStepsInstruction ?? "";
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
