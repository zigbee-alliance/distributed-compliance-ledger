// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateVendorInfoType: (data: MsgCreateVendorInfoType): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgCreateVendorInfoType", value: data }),
    msgUpdateVendorInfoType: (data: MsgUpdateVendorInfoType): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgUpdateVendorInfoType", value: data }),
    msgDeleteVendorInfoType: (data: MsgDeleteVendorInfoType): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.vendorinfo.MsgDeleteVendorInfoType", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
