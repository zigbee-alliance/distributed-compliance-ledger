// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgUpdateModel } from "./types/model/tx";
import { MsgUpdateModelVersion } from "./types/model/tx";
import { MsgCreateModel } from "./types/model/tx";
import { MsgCreateModelVersion } from "./types/model/tx";
import { MsgDeleteModel } from "./types/model/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModel", MsgUpdateModel],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModelVersion", MsgUpdateModelVersion],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModel", MsgCreateModel],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModelVersion", MsgCreateModelVersion],
    ["/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModel", MsgDeleteModel],
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
        msgUpdateModel: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModel", value: MsgUpdateModel.fromPartial(data) }),
        msgUpdateModelVersion: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModelVersion", value: MsgUpdateModelVersion.fromPartial(data) }),
        msgCreateModel: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgCreateModel", value: MsgCreateModel.fromPartial(data) }),
        msgCreateModelVersion: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgCreateModelVersion", value: MsgCreateModelVersion.fromPartial(data) }),
        msgDeleteModel: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModel", value: MsgDeleteModel.fromPartial(data) }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
