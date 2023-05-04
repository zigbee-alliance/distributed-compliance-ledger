// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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
    msgApproveRevokeAccount: (data: MsgApproveRevokeAccount): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", value: MsgApproveRevokeAccount.fromPartial( data ) }),
    msgRejectAddAccount: (data: MsgRejectAddAccount): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgRejectAddAccount", value: MsgRejectAddAccount.fromPartial( data ) }),
    msgProposeAddAccount: (data: MsgProposeAddAccount): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", value: MsgProposeAddAccount.fromPartial( data ) }),
    msgApproveAddAccount: (data: MsgApproveAddAccount): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", value: MsgApproveAddAccount.fromPartial( data ) }),
    msgProposeRevokeAccount: (data: MsgProposeRevokeAccount): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", value: MsgProposeRevokeAccount.fromPartial( data ) }),
    
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
