// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateModelVersion } from "./types/model/tx";
import { MsgUpdateModelVersion } from "./types/model/tx";
import { MsgCreateModel } from "./types/model/tx";
import { MsgUpdateModel } from "./types/model/tx";
import { MsgDeleteModel } from "./types/model/tx";


const types = [
  ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModelVersion", MsgCreateModelVersion],
  ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModelVersion", MsgUpdateModelVersion],
  ["/zigbeealliance.distributedcomplianceledger.model.MsgCreateModel", MsgCreateModel],
  ["/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModel", MsgUpdateModel],
  ["/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModel", MsgDeleteModel],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

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
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateModelVersion: (data: MsgCreateModelVersion): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgCreateModelVersion", value: MsgCreateModelVersion.fromPartial( data ) }),
    msgUpdateModelVersion: (data: MsgUpdateModelVersion): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModelVersion", value: MsgUpdateModelVersion.fromPartial( data ) }),
    msgCreateModel: (data: MsgCreateModel): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgCreateModel", value: MsgCreateModel.fromPartial( data ) }),
    msgUpdateModel: (data: MsgUpdateModel): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgUpdateModel", value: MsgUpdateModel.fromPartial( data ) }),
    msgDeleteModel: (data: MsgDeleteModel): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.model.MsgDeleteModel", value: MsgDeleteModel.fromPartial( data ) }),
    
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
