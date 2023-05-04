// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgApproveRevokeAccount } from "./types/dclauth/tx";
import { MsgRejectAddAccount } from "./types/dclauth/tx";
import { MsgProposeAddAccount } from "./types/dclauth/tx";
import { MsgApproveAddAccount } from "./types/dclauth/tx";
import { MsgProposeRevokeAccount } from "./types/dclauth/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", MsgApproveRevokeAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgRejectAddAccount", MsgRejectAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", MsgProposeAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", MsgApproveAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", MsgProposeRevokeAccount],
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
        msgApproveRevokeAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", value: MsgApproveRevokeAccount.fromPartial(data) }),
        msgRejectAddAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgRejectAddAccount", value: MsgRejectAddAccount.fromPartial(data) }),
        msgProposeAddAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", value: MsgProposeAddAccount.fromPartial(data) }),
        msgApproveAddAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", value: MsgApproveAddAccount.fromPartial(data) }),
        msgProposeRevokeAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", value: MsgProposeRevokeAccount.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
