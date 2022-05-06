/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { Account } from '../dclauth/account';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
import { PendingAccount } from '../dclauth/pending_account';
import { PendingAccountRevocation } from '../dclauth/pending_account_revocation';
import { AccountStat } from '../dclauth/account_stat';
import { RevokedAccount } from '../dclauth/revoked_account';
import { RejectedAccount } from '../dclauth/rejected_account';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
const baseQueryGetAccountRequest = { address: '' };
export const QueryGetAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetAccountResponse = {};
export const QueryGetAccountResponse = {
    encode(message, writer = Writer.create()) {
        if (message.account !== undefined) {
            Account.encode(message.account, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetAccountResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.account = Account.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetAccountResponse };
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromJSON(object.account);
        }
        else {
            message.account = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetAccountResponse };
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromPartial(object.account);
        }
        else {
            message.account = undefined;
        }
        return message;
    }
};
const baseQueryAllAccountRequest = {};
export const QueryAllAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllAccountResponse = {};
export const QueryAllAccountResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.account) {
            Account.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllAccountResponse };
        message.account = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.account.push(Account.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllAccountResponse };
        message.account = [];
        if (object.account !== undefined && object.account !== null) {
            for (const e of object.account) {
                message.account.push(Account.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.account) {
            obj.account = message.account.map((e) => (e ? Account.toJSON(e) : undefined));
        }
        else {
            obj.account = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllAccountResponse };
        message.account = [];
        if (object.account !== undefined && object.account !== null) {
            for (const e of object.account) {
                message.account.push(Account.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetPendingAccountRequest = { address: '' };
export const QueryGetPendingAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPendingAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPendingAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPendingAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetPendingAccountResponse = {};
export const QueryGetPendingAccountResponse = {
    encode(message, writer = Writer.create()) {
        if (message.pendingAccount !== undefined) {
            PendingAccount.encode(message.pendingAccount, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPendingAccountResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pendingAccount = PendingAccount.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPendingAccountResponse };
        if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
            message.pendingAccount = PendingAccount.fromJSON(object.pendingAccount);
        }
        else {
            message.pendingAccount = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pendingAccount !== undefined && (obj.pendingAccount = message.pendingAccount ? PendingAccount.toJSON(message.pendingAccount) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPendingAccountResponse };
        if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
            message.pendingAccount = PendingAccount.fromPartial(object.pendingAccount);
        }
        else {
            message.pendingAccount = undefined;
        }
        return message;
    }
};
const baseQueryAllPendingAccountRequest = {};
export const QueryAllPendingAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllPendingAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllPendingAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllPendingAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllPendingAccountResponse = {};
export const QueryAllPendingAccountResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.pendingAccount) {
            PendingAccount.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllPendingAccountResponse };
        message.pendingAccount = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pendingAccount.push(PendingAccount.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllPendingAccountResponse };
        message.pendingAccount = [];
        if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
            for (const e of object.pendingAccount) {
                message.pendingAccount.push(PendingAccount.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.pendingAccount) {
            obj.pendingAccount = message.pendingAccount.map((e) => (e ? PendingAccount.toJSON(e) : undefined));
        }
        else {
            obj.pendingAccount = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllPendingAccountResponse };
        message.pendingAccount = [];
        if (object.pendingAccount !== undefined && object.pendingAccount !== null) {
            for (const e of object.pendingAccount) {
                message.pendingAccount.push(PendingAccount.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetPendingAccountRevocationRequest = { address: '' };
export const QueryGetPendingAccountRevocationRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPendingAccountRevocationRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPendingAccountRevocationRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPendingAccountRevocationRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetPendingAccountRevocationResponse = {};
export const QueryGetPendingAccountRevocationResponse = {
    encode(message, writer = Writer.create()) {
        if (message.pendingAccountRevocation !== undefined) {
            PendingAccountRevocation.encode(message.pendingAccountRevocation, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetPendingAccountRevocationResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pendingAccountRevocation = PendingAccountRevocation.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetPendingAccountRevocationResponse };
        if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
            message.pendingAccountRevocation = PendingAccountRevocation.fromJSON(object.pendingAccountRevocation);
        }
        else {
            message.pendingAccountRevocation = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pendingAccountRevocation !== undefined &&
            (obj.pendingAccountRevocation = message.pendingAccountRevocation ? PendingAccountRevocation.toJSON(message.pendingAccountRevocation) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetPendingAccountRevocationResponse };
        if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
            message.pendingAccountRevocation = PendingAccountRevocation.fromPartial(object.pendingAccountRevocation);
        }
        else {
            message.pendingAccountRevocation = undefined;
        }
        return message;
    }
};
const baseQueryAllPendingAccountRevocationRequest = {};
export const QueryAllPendingAccountRevocationRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllPendingAccountRevocationRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllPendingAccountRevocationRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllPendingAccountRevocationRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllPendingAccountRevocationResponse = {};
export const QueryAllPendingAccountRevocationResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.pendingAccountRevocation) {
            PendingAccountRevocation.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllPendingAccountRevocationResponse };
        message.pendingAccountRevocation = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pendingAccountRevocation.push(PendingAccountRevocation.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllPendingAccountRevocationResponse };
        message.pendingAccountRevocation = [];
        if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
            for (const e of object.pendingAccountRevocation) {
                message.pendingAccountRevocation.push(PendingAccountRevocation.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.pendingAccountRevocation) {
            obj.pendingAccountRevocation = message.pendingAccountRevocation.map((e) => (e ? PendingAccountRevocation.toJSON(e) : undefined));
        }
        else {
            obj.pendingAccountRevocation = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllPendingAccountRevocationResponse };
        message.pendingAccountRevocation = [];
        if (object.pendingAccountRevocation !== undefined && object.pendingAccountRevocation !== null) {
            for (const e of object.pendingAccountRevocation) {
                message.pendingAccountRevocation.push(PendingAccountRevocation.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetAccountStatRequest = {};
export const QueryGetAccountStatRequest = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetAccountStatRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(_) {
        const message = { ...baseQueryGetAccountStatRequest };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseQueryGetAccountStatRequest };
        return message;
    }
};
const baseQueryGetAccountStatResponse = {};
export const QueryGetAccountStatResponse = {
    encode(message, writer = Writer.create()) {
        if (message.AccountStat !== undefined) {
            AccountStat.encode(message.AccountStat, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetAccountStatResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.AccountStat = AccountStat.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetAccountStatResponse };
        if (object.AccountStat !== undefined && object.AccountStat !== null) {
            message.AccountStat = AccountStat.fromJSON(object.AccountStat);
        }
        else {
            message.AccountStat = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.AccountStat !== undefined && (obj.AccountStat = message.AccountStat ? AccountStat.toJSON(message.AccountStat) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetAccountStatResponse };
        if (object.AccountStat !== undefined && object.AccountStat !== null) {
            message.AccountStat = AccountStat.fromPartial(object.AccountStat);
        }
        else {
            message.AccountStat = undefined;
        }
        return message;
    }
};
const baseQueryGetRevokedAccountRequest = { address: '' };
export const QueryGetRevokedAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRevokedAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetRevokedAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRevokedAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetRevokedAccountResponse = {};
export const QueryGetRevokedAccountResponse = {
    encode(message, writer = Writer.create()) {
        if (message.revokedAccount !== undefined) {
            RevokedAccount.encode(message.revokedAccount, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRevokedAccountResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.revokedAccount = RevokedAccount.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetRevokedAccountResponse };
        if (object.revokedAccount !== undefined && object.revokedAccount !== null) {
            message.revokedAccount = RevokedAccount.fromJSON(object.revokedAccount);
        }
        else {
            message.revokedAccount = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.revokedAccount !== undefined && (obj.revokedAccount = message.revokedAccount ? RevokedAccount.toJSON(message.revokedAccount) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRevokedAccountResponse };
        if (object.revokedAccount !== undefined && object.revokedAccount !== null) {
            message.revokedAccount = RevokedAccount.fromPartial(object.revokedAccount);
        }
        else {
            message.revokedAccount = undefined;
        }
        return message;
    }
};
const baseQueryAllRevokedAccountRequest = {};
export const QueryAllRevokedAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRevokedAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllRevokedAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRevokedAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllRevokedAccountResponse = {};
export const QueryAllRevokedAccountResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.revokedAccount) {
            RevokedAccount.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRevokedAccountResponse };
        message.revokedAccount = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.revokedAccount.push(RevokedAccount.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllRevokedAccountResponse };
        message.revokedAccount = [];
        if (object.revokedAccount !== undefined && object.revokedAccount !== null) {
            for (const e of object.revokedAccount) {
                message.revokedAccount.push(RevokedAccount.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.revokedAccount) {
            obj.revokedAccount = message.revokedAccount.map((e) => (e ? RevokedAccount.toJSON(e) : undefined));
        }
        else {
            obj.revokedAccount = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRevokedAccountResponse };
        message.revokedAccount = [];
        if (object.revokedAccount !== undefined && object.revokedAccount !== null) {
            for (const e of object.revokedAccount) {
                message.revokedAccount.push(RevokedAccount.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryGetRejectedAccountRequest = { address: '' };
export const QueryGetRejectedAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== '') {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRejectedAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetRejectedAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRejectedAccountRequest };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseQueryGetRejectedAccountResponse = {};
export const QueryGetRejectedAccountResponse = {
    encode(message, writer = Writer.create()) {
        if (message.rejectedAccount !== undefined) {
            RejectedAccount.encode(message.rejectedAccount, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetRejectedAccountResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.rejectedAccount = RejectedAccount.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetRejectedAccountResponse };
        if (object.rejectedAccount !== undefined && object.rejectedAccount !== null) {
            message.rejectedAccount = RejectedAccount.fromJSON(object.rejectedAccount);
        }
        else {
            message.rejectedAccount = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.rejectedAccount !== undefined && (obj.rejectedAccount = message.rejectedAccount ? RejectedAccount.toJSON(message.rejectedAccount) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetRejectedAccountResponse };
        if (object.rejectedAccount !== undefined && object.rejectedAccount !== null) {
            message.rejectedAccount = RejectedAccount.fromPartial(object.rejectedAccount);
        }
        else {
            message.rejectedAccount = undefined;
        }
        return message;
    }
};
const baseQueryAllRejectedAccountRequest = {};
export const QueryAllRejectedAccountRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRejectedAccountRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllRejectedAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRejectedAccountRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllRejectedAccountResponse = {};
export const QueryAllRejectedAccountResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.rejectedAccount) {
            RejectedAccount.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllRejectedAccountResponse };
        message.rejectedAccount = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.rejectedAccount.push(RejectedAccount.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllRejectedAccountResponse };
        message.rejectedAccount = [];
        if (object.rejectedAccount !== undefined && object.rejectedAccount !== null) {
            for (const e of object.rejectedAccount) {
                message.rejectedAccount.push(RejectedAccount.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.rejectedAccount) {
            obj.rejectedAccount = message.rejectedAccount.map((e) => (e ? RejectedAccount.toJSON(e) : undefined));
        }
        else {
            obj.rejectedAccount = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllRejectedAccountResponse };
        message.rejectedAccount = [];
        if (object.rejectedAccount !== undefined && object.rejectedAccount !== null) {
            for (const e of object.rejectedAccount) {
                message.rejectedAccount.push(RejectedAccount.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Account(request) {
        const data = QueryGetAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'Account', data);
        return promise.then((data) => QueryGetAccountResponse.decode(new Reader(data)));
    }
    AccountAll(request) {
        const data = QueryAllAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'AccountAll', data);
        return promise.then((data) => QueryAllAccountResponse.decode(new Reader(data)));
    }
    PendingAccount(request) {
        const data = QueryGetPendingAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccount', data);
        return promise.then((data) => QueryGetPendingAccountResponse.decode(new Reader(data)));
    }
    PendingAccountAll(request) {
        const data = QueryAllPendingAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccountAll', data);
        return promise.then((data) => QueryAllPendingAccountResponse.decode(new Reader(data)));
    }
    PendingAccountRevocation(request) {
        const data = QueryGetPendingAccountRevocationRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccountRevocation', data);
        return promise.then((data) => QueryGetPendingAccountRevocationResponse.decode(new Reader(data)));
    }
    PendingAccountRevocationAll(request) {
        const data = QueryAllPendingAccountRevocationRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'PendingAccountRevocationAll', data);
        return promise.then((data) => QueryAllPendingAccountRevocationResponse.decode(new Reader(data)));
    }
    AccountStat(request) {
        const data = QueryGetAccountStatRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'AccountStat', data);
        return promise.then((data) => QueryGetAccountStatResponse.decode(new Reader(data)));
    }
    RevokedAccount(request) {
        const data = QueryGetRevokedAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'RevokedAccount', data);
        return promise.then((data) => QueryGetRevokedAccountResponse.decode(new Reader(data)));
    }
    RevokedAccountAll(request) {
        const data = QueryAllRevokedAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'RevokedAccountAll', data);
        return promise.then((data) => QueryAllRevokedAccountResponse.decode(new Reader(data)));
    }
    RejectedAccount(request) {
        const data = QueryGetRejectedAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'RejectedAccount', data);
        return promise.then((data) => QueryGetRejectedAccountResponse.decode(new Reader(data)));
    }
    RejectedAccountAll(request) {
        const data = QueryAllRejectedAccountRequest.encode(request).finish();
        const promise = this.rpc.request('zigbeealliance.distributedcomplianceledger.dclauth.Query', 'RejectedAccountAll', data);
        return promise.then((data) => QueryAllRejectedAccountResponse.decode(new Reader(data)));
    }
}
