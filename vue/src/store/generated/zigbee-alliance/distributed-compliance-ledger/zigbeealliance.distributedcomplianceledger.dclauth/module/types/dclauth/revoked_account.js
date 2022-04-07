/* eslint-disable */
import { Account } from '../dclauth/account';
import { Grant } from '../dclauth/grant';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclauth';
export var RevokedAccount_Reason;
(function (RevokedAccount_Reason) {
    RevokedAccount_Reason[RevokedAccount_Reason["TrusteeVoting"] = 0] = "TrusteeVoting";
    RevokedAccount_Reason[RevokedAccount_Reason["MaliciousValidator"] = 1] = "MaliciousValidator";
    RevokedAccount_Reason[RevokedAccount_Reason["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(RevokedAccount_Reason || (RevokedAccount_Reason = {}));
export function revokedAccount_ReasonFromJSON(object) {
    switch (object) {
        case 0:
        case 'TrusteeVoting':
            return RevokedAccount_Reason.TrusteeVoting;
        case 1:
        case 'MaliciousValidator':
            return RevokedAccount_Reason.MaliciousValidator;
        case -1:
        case 'UNRECOGNIZED':
        default:
            return RevokedAccount_Reason.UNRECOGNIZED;
    }
}
export function revokedAccount_ReasonToJSON(object) {
    switch (object) {
        case RevokedAccount_Reason.TrusteeVoting:
            return 'TrusteeVoting';
        case RevokedAccount_Reason.MaliciousValidator:
            return 'MaliciousValidator';
        default:
            return 'UNKNOWN';
    }
}
const baseRevokedAccount = { reason: 0 };
export const RevokedAccount = {
    encode(message, writer = Writer.create()) {
        if (message.account !== undefined) {
            Account.encode(message.account, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.revokeApprovals) {
            Grant.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.reason !== 0) {
            writer.uint32(24).int32(message.reason);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRevokedAccount };
        message.revokeApprovals = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.account = Account.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.revokeApprovals.push(Grant.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.reason = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseRevokedAccount };
        message.revokeApprovals = [];
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromJSON(object.account);
        }
        else {
            message.account = undefined;
        }
        if (object.revokeApprovals !== undefined && object.revokeApprovals !== null) {
            for (const e of object.revokeApprovals) {
                message.revokeApprovals.push(Grant.fromJSON(e));
            }
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = revokedAccount_ReasonFromJSON(object.reason);
        }
        else {
            message.reason = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.account !== undefined && (obj.account = message.account ? Account.toJSON(message.account) : undefined);
        if (message.revokeApprovals) {
            obj.revokeApprovals = message.revokeApprovals.map((e) => (e ? Grant.toJSON(e) : undefined));
        }
        else {
            obj.revokeApprovals = [];
        }
        message.reason !== undefined && (obj.reason = revokedAccount_ReasonToJSON(message.reason));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRevokedAccount };
        message.revokeApprovals = [];
        if (object.account !== undefined && object.account !== null) {
            message.account = Account.fromPartial(object.account);
        }
        else {
            message.account = undefined;
        }
        if (object.revokeApprovals !== undefined && object.revokeApprovals !== null) {
            for (const e of object.revokeApprovals) {
                message.revokeApprovals.push(Grant.fromPartial(e));
            }
        }
        if (object.reason !== undefined && object.reason !== null) {
            message.reason = object.reason;
        }
        else {
            message.reason = 0;
        }
        return message;
    }
};
