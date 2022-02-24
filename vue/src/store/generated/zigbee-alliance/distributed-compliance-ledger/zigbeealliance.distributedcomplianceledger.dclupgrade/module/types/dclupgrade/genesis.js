/* eslint-disable */
import { ProposedUpgrade } from '../dclupgrade/proposed_upgrade';
import { ApprovedUpgrade } from '../dclupgrade/approved_upgrade';
import { Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'zigbeealliance.distributedcomplianceledger.dclupgrade';
const baseGenesisState = {};
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.proposedUpgradeList) {
            ProposedUpgrade.encode(v, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.approvedUpgradeList) {
            ApprovedUpgrade.encode(v, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.proposedUpgradeList = [];
        message.approvedUpgradeList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.proposedUpgradeList.push(ProposedUpgrade.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.approvedUpgradeList.push(ApprovedUpgrade.decode(reader, reader.uint32()));
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
        message.proposedUpgradeList = [];
        message.approvedUpgradeList = [];
        if (object.proposedUpgradeList !== undefined && object.proposedUpgradeList !== null) {
            for (const e of object.proposedUpgradeList) {
                message.proposedUpgradeList.push(ProposedUpgrade.fromJSON(e));
            }
        }
        if (object.approvedUpgradeList !== undefined && object.approvedUpgradeList !== null) {
            for (const e of object.approvedUpgradeList) {
                message.approvedUpgradeList.push(ApprovedUpgrade.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.proposedUpgradeList) {
            obj.proposedUpgradeList = message.proposedUpgradeList.map((e) => (e ? ProposedUpgrade.toJSON(e) : undefined));
        }
        else {
            obj.proposedUpgradeList = [];
        }
        if (message.approvedUpgradeList) {
            obj.approvedUpgradeList = message.approvedUpgradeList.map((e) => (e ? ApprovedUpgrade.toJSON(e) : undefined));
        }
        else {
            obj.approvedUpgradeList = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.proposedUpgradeList = [];
        message.approvedUpgradeList = [];
        if (object.proposedUpgradeList !== undefined && object.proposedUpgradeList !== null) {
            for (const e of object.proposedUpgradeList) {
                message.proposedUpgradeList.push(ProposedUpgrade.fromPartial(e));
            }
        }
        if (object.approvedUpgradeList !== undefined && object.approvedUpgradeList !== null) {
            for (const e of object.approvedUpgradeList) {
                message.approvedUpgradeList.push(ApprovedUpgrade.fromPartial(e));
            }
        }
        return message;
    }
};
