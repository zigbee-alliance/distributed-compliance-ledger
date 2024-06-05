/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

export interface UniqueCertificate {
  issuer: string;
  serialNumber: string;
  present: boolean;
  schemaVersion: number;
}

function createBaseUniqueCertificate(): UniqueCertificate {
  return { issuer: "", serialNumber: "", present: false, schemaVersion: 0 };
}

export const UniqueCertificate = {
  encode(message: UniqueCertificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.issuer !== "") {
      writer.uint32(10).string(message.issuer);
    }
    if (message.serialNumber !== "") {
      writer.uint32(18).string(message.serialNumber);
    }
    if (message.present === true) {
      writer.uint32(24).bool(message.present);
    }
    if (message.schemaVersion !== 0) {
      writer.uint32(32).uint32(message.schemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UniqueCertificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUniqueCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.issuer = reader.string();
          break;
        case 2:
          message.serialNumber = reader.string();
          break;
        case 3:
          message.present = reader.bool();
          break;
        case 4:
          message.schemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UniqueCertificate {
    return {
      issuer: isSet(object.issuer) ? String(object.issuer) : "",
      serialNumber: isSet(object.serialNumber) ? String(object.serialNumber) : "",
      present: isSet(object.present) ? Boolean(object.present) : false,
      schemaVersion: isSet(object.schemaVersion) ? Number(object.schemaVersion) : 0,
    };
  },

  toJSON(message: UniqueCertificate): unknown {
    const obj: any = {};
    message.issuer !== undefined && (obj.issuer = message.issuer);
    message.serialNumber !== undefined && (obj.serialNumber = message.serialNumber);
    message.present !== undefined && (obj.present = message.present);
    message.schemaVersion !== undefined && (obj.schemaVersion = Math.round(message.schemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UniqueCertificate>, I>>(object: I): UniqueCertificate {
    const message = createBaseUniqueCertificate();
    message.issuer = object.issuer ?? "";
    message.serialNumber = object.serialNumber ?? "";
    message.present = object.present ?? false;
    message.schemaVersion = object.schemaVersion ?? 0;
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
