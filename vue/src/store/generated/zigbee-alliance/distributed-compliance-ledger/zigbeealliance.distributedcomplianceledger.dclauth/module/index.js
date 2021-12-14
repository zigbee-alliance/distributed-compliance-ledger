// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgProposeAddAccount } from "./types/dclauth/tx";
import { MsgApproveRevokeAccount } from "./types/dclauth/tx";
import { MsgProposeRevokeAccount } from "./types/dclauth/tx";
import { MsgApproveAddAccount } from "./types/dclauth/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", MsgProposeAddAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", MsgApproveRevokeAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", MsgProposeRevokeAccount],
    ["/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", MsgApproveAddAccount],
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
        msgProposeAddAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", value: MsgProposeAddAccount.fromPartial(data) }),
        msgApproveRevokeAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", value: MsgApproveRevokeAccount.fromPartial(data) }),
        msgProposeRevokeAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", value: MsgProposeRevokeAccount.fromPartial(data) }),
        msgApproveAddAccount: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", value: MsgApproveAddAccount.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
