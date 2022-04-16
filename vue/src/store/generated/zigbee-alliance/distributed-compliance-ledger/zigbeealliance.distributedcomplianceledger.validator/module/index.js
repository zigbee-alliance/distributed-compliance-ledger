// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateValidator } from "./types/validator/tx";
import { MsgProposeDisableValidator } from "./types/validator/tx";
import { MsgEnableValidator } from "./types/validator/tx";
import { MsgApproveDisableValidator } from "./types/validator/tx";
import { MsgDisableValidator } from "./types/validator/tx";
import { MsgRejectDisableNode } from "./types/validator/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgCreateValidator", MsgCreateValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgProposeDisableValidator", MsgProposeDisableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgEnableValidator", MsgEnableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgApproveDisableValidator", MsgApproveDisableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgDisableValidator", MsgDisableValidator],
    ["/zigbeealliance.distributedcomplianceledger.validator.MsgRejectDisableNode", MsgRejectDisableNode],
];
export const MissingWalletError = new Error("wallet is required");
export const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    let client;
    if (addr) {
        client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    }
    else {
        client = await SigningStargateClient.offline(wallet, { registry });
    }
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgCreateValidator: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgCreateValidator", value: MsgCreateValidator.fromPartial(data) }),
        msgProposeDisableValidator: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgProposeDisableValidator", value: MsgProposeDisableValidator.fromPartial(data) }),
        msgEnableValidator: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgEnableValidator", value: MsgEnableValidator.fromPartial(data) }),
        msgApproveDisableValidator: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgApproveDisableValidator", value: MsgApproveDisableValidator.fromPartial(data) }),
        msgDisableValidator: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgDisableValidator", value: MsgDisableValidator.fromPartial(data) }),
        msgRejectDisableNode: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgRejectDisableNode", value: MsgRejectDisableNode.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
