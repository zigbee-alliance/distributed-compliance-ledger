// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateNewVendorInfo } from "./types/vendorinfo/tx";
import { MsgUpdateNewVendorInfo } from "./types/vendorinfo/tx";
import { MsgDeleteNewVendorInfo } from "./types/vendorinfo/tx";
const types = [
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgCreateNewVendorInfo", MsgCreateNewVendorInfo],
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgUpdateNewVendorInfo", MsgUpdateNewVendorInfo],
    ["/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgDeleteNewVendorInfo", MsgDeleteNewVendorInfo],
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
        msgCreateNewVendorInfo: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgCreateNewVendorInfo", value: data }),
        msgUpdateNewVendorInfo: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgUpdateNewVendorInfo", value: data }),
        msgDeleteNewVendorInfo: (data) => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgDeleteNewVendorInfo", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
