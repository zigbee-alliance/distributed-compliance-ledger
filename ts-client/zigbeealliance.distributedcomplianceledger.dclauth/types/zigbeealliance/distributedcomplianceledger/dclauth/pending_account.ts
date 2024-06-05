/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Account } from "./account";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.dclauth";

/**
 * TODO issue 99: do we need that ???
 * option (gogoproto.goproto_getters)  = false;
 * option (gogoproto.goproto_stringer) = false;
 */
export interface PendingAccount {
  account: Account | undefined;
  pendingAccountSchemaVersion: number;
}

function createBasePendingAccount(): PendingAccount {
  return { account: undefined, pendingAccountSchemaVersion: 0 };
}

export const PendingAccount = {
  encode(message: PendingAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== undefined) {
      Account.encode(message.account, writer.uint32(10).fork()).ldelim();
    }
    if (message.pendingAccountSchemaVersion !== 0) {
      writer.uint32(16).uint32(message.pendingAccountSchemaVersion);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PendingAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePendingAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.account = Account.decode(reader, reader.uint32());
          break;
        case 2:
          message.pendingAccountSchemaVersion = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PendingAccount {
    return {
      account: isSet(object.account) ? Account.fromJSON(object.account) : undefined,
      pendingAccountSchemaVersion: isSet(object.pendingAccountSchemaVersion)
        ? Number(object.pendingAccountSchemaVersion)
        : 0,
    };
  },

  toJSON(message: PendingAccount): unknown {
    const obj: any = {};
    message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
    message.pendingAccountSchemaVersion !== undefined
      && (obj.pendingAccountSchemaVersion = Math.round(message.pendingAccountSchemaVersion));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PendingAccount>, I>>(object: I): PendingAccount {
    const message = createBasePendingAccount();
    message.account = (object.account !== undefined && object.account !== null)
      ? Account.fromPartial(object.account)
      : undefined;
    message.pendingAccountSchemaVersion = object.pendingAccountSchemaVersion ?? 0;
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
