// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgApproveUpgrade } from "./types/dclupgrade/tx";
import { MsgProposeUpgrade } from "./types/dclupgrade/tx";
import { MsgRejectUpgrade } from "./types/dclupgrade/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgApproveUpgrade", MsgApproveUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgProposeUpgrade", MsgProposeUpgrade],
    ["/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgRejectUpgrade", MsgRejectUpgrade],
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
        msgApproveUpgrade: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgApproveUpgrade", value: MsgApproveUpgrade.fromPartial(data) }),
        msgProposeUpgrade: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgProposeUpgrade", value: MsgProposeUpgrade.fromPartial(data) }),
        msgRejectUpgrade: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgRejectUpgrade", value: MsgRejectUpgrade.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
