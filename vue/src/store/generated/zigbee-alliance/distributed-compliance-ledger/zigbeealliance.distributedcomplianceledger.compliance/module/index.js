// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgProvisionModel } from "./types/compliance/tx";
import { MsgCertifyModel } from "./types/compliance/tx";
import { MsgRevokeModel } from "./types/compliance/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgProvisionModel", MsgProvisionModel],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel", MsgCertifyModel],
    ["/zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel", MsgRevokeModel],
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
        msgProvisionModel: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgProvisionModel", value: MsgProvisionModel.fromPartial(data) }),
        msgCertifyModel: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel", value: MsgCertifyModel.fromPartial(data) }),
        msgRevokeModel: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel", value: MsgRevokeModel.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
