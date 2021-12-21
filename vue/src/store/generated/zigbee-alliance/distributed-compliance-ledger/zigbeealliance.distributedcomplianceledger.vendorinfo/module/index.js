// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateVendorInfoType } from "./types/vendorinfo/tx";
import { MsgUpdateVendorInfoType } from "./types/vendorinfo/tx";
import { MsgDeleteVendorInfoType } from "./types/vendorinfo/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgCreateVendorInfoType", MsgCreateVendorInfoType],
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgUpdateVendorInfoType", MsgUpdateVendorInfoType],
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgDeleteVendorInfoType", MsgDeleteVendorInfoType],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgCreateVendorInfoType: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgCreateVendorInfoType", value: data }),
        msgUpdateVendorInfoType: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgUpdateVendorInfoType", value: data }),
        msgDeleteVendorInfoType: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgDeleteVendorInfoType", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
