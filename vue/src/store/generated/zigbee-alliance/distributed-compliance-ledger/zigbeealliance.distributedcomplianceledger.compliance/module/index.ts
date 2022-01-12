// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRevokeModel } from "./types/compliance/tx";
import { MsgCertifyModel } from "./types/compliance/tx";
import { MsgProvisionModel } from "./types/compliance/tx";


const types = [
  ["/zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel", MsgRevokeModel],
  ["/zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel", MsgCertifyModel],
  ["/zigbeealliance.distributedcomplianceledger.compliance.MsgProvisionModel", MsgProvisionModel],
  
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
    msgRevokeModel: (data: MsgRevokeModel): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel", value: MsgRevokeModel.fromPartial( data ) }),
    msgCertifyModel: (data: MsgCertifyModel): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel", value: MsgCertifyModel.fromPartial( data ) }),
    msgProvisionModel: (data: MsgProvisionModel): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgProvisionModel", value: MsgProvisionModel.fromPartial( data ) }),
    
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
