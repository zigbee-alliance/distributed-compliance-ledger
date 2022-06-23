// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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
    msgApproveUpgrade: (data: MsgApproveUpgrade): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgApproveUpgrade", value: MsgApproveUpgrade.fromPartial( data ) }),
    msgProposeUpgrade: (data: MsgProposeUpgrade): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgProposeUpgrade", value: MsgProposeUpgrade.fromPartial( data ) }),
    msgRejectUpgrade: (data: MsgRejectUpgrade): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgRejectUpgrade", value: MsgRejectUpgrade.fromPartial( data ) }),
    
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
