/* eslint-disable */
import { Account } from '../dclauth/account';
import { PendingAccount } from '../dclauth/pending_account';
import { PendingAccountRevocation } from '../dclauth/pending_account_revocation';
import { AccountStat } from '../dclauth/account_stat';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.accountList) {
            Account.encode(v, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.pendingAccountList) {
            PendingAccount.encode(v, writer.uint32(18).fork()).ldelim();
        }
        for (const v of message.pendingAccountRevocationList) {
            PendingAccountRevocation.encode(v, writer.uint32(26).fork()).ldelim();
        }
        if (message.accountStat !== undefined) {
            AccountStat.encode(message.accountStat, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.accountList = [];
        message.pendingAccountList = [];
        message.pendingAccountRevocationList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.accountList.push(Account.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pendingAccountList.push(PendingAccount.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.pendingAccountRevocationList.push(PendingAccountRevocation.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.accountStat = AccountStat.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.accountList = [];
        message.pendingAccountList = [];
        message.pendingAccountRevocationList = [];
        if (object.accountList !== undefined && object.accountList !== null) {
            for (const e of object.accountList) {
                message.accountList.push(Account.fromJSON(e));
            }
        }
        if (object.pendingAccountList !== undefined && object.pendingAccountList !== null) {
            for (const e of object.pendingAccountList) {
                message.pendingAccountList.push(PendingAccount.fromJSON(e));
            }
        }
        if (object.pendingAccountRevocationList !== undefined && object.pendingAccountRevocationList !== null) {
            for (const e of object.pendingAccountRevocationList) {
                message.pendingAccountRevocationList.push(PendingAccountRevocation.fromJSON(e));
            }
        }
        if (object.accountStat !== undefined && object.accountStat !== null) {
            message.accountStat = AccountStat.fromJSON(object.accountStat);
        }
        else {
            message.accountStat = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.accountList) {
            obj.accountList = message.accountList.map((e) => (e ? Account.toJSON(e) : undefined));
        }
        else {
            obj.accountList = [];
        }
        if (message.pendingAccountList) {
            obj.pendingAccountList = message.pendingAccountList.map((e) => (e ? PendingAccount.toJSON(e) : undefined));
        }
        else {
            obj.pendingAccountList = [];
        }
        if (message.pendingAccountRevocationList) {
            obj.pendingAccountRevocationList = message.pendingAccountRevocationList.map((e) => (e ? PendingAccountRevocation.toJSON(e) : undefined));
        }
        else {
            obj.pendingAccountRevocationList = [];
        }
        message.accountStat !== undefined && (obj.accountStat = message.accountStat ? AccountStat.toJSON(message.accountStat) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.accountList = [];
        message.pendingAccountList = [];
        message.pendingAccountRevocationList = [];
        if (object.accountList !== undefined && object.accountList !== null) {
            for (const e of object.accountList) {
                message.accountList.push(Account.fromPartial(e));
            }
        }
        if (object.pendingAccountList !== undefined && object.pendingAccountList !== null) {
            for (const e of object.pendingAccountList) {
                message.pendingAccountList.push(PendingAccount.fromPartial(e));
            }
        }
        if (object.pendingAccountRevocationList !== undefined && object.pendingAccountRevocationList !== null) {
            for (const e of object.pendingAccountRevocationList) {
                message.pendingAccountRevocationList.push(PendingAccountRevocation.fromPartial(e));
            }
        }
        if (object.accountStat !== undefined && object.accountStat !== null) {
            message.accountStat = AccountStat.fromPartial(object.accountStat);
        }
        else {
            message.accountStat = undefined;
        }
        return message;
    }
};
