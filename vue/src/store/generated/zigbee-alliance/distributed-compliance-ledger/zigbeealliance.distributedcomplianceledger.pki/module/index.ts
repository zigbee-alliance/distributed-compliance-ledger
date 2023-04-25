// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgApproveRevokeX509RootCert } from "./types/pki/tx";
import { MsgRejectAddX509RootCert } from "./types/pki/tx";
import { MsgProposeAddX509RootCert } from "./types/pki/tx";
import { MsgRevokeX509Cert } from "./types/pki/tx";
import { MsgProposeRevokeX509RootCert } from "./types/pki/tx";
import { MsgApproveAddX509RootCert } from "./types/pki/tx";
import { MsgAddX509Cert } from "./types/pki/tx";


const types = [
  ["/zigbeealliance.distributedcomplianceledger.pki.MsgApproveRevokeX509RootCert", MsgApproveRevokeX509RootCert],
  ["/zigbeealliance.distributedcomplianceledger.pki.MsgRejectAddX509RootCert", MsgRejectAddX509RootCert],
  ["/zigbeealliance.distributedcomplianceledger.pki.MsgProposeAddX509RootCert", MsgProposeAddX509RootCert],
  ["/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeX509Cert", MsgRevokeX509Cert],
  ["/zigbeealliance.distributedcomplianceledger.pki.MsgProposeRevokeX509RootCert", MsgProposeRevokeX509RootCert],
  ["/zigbeealliance.distributedcomplianceledger.pki.MsgApproveAddX509RootCert", MsgApproveAddX509RootCert],
  ["/zigbeealliance.distributedcomplianceledger.pki.MsgAddX509Cert", MsgAddX509Cert],
  
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
    msgApproveRevokeX509RootCert: (data: MsgApproveRevokeX509RootCert): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgApproveRevokeX509RootCert", value: MsgApproveRevokeX509RootCert.fromPartial( data ) }),
    msgRejectAddX509RootCert: (data: MsgRejectAddX509RootCert): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRejectAddX509RootCert", value: MsgRejectAddX509RootCert.fromPartial( data ) }),
    msgProposeAddX509RootCert: (data: MsgProposeAddX509RootCert): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgProposeAddX509RootCert", value: MsgProposeAddX509RootCert.fromPartial( data ) }),
    msgRevokeX509Cert: (data: MsgRevokeX509Cert): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeX509Cert", value: MsgRevokeX509Cert.fromPartial( data ) }),
    msgProposeRevokeX509RootCert: (data: MsgProposeRevokeX509RootCert): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgProposeRevokeX509RootCert", value: MsgProposeRevokeX509RootCert.fromPartial( data ) }),
    msgApproveAddX509RootCert: (data: MsgApproveAddX509RootCert): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgApproveAddX509RootCert", value: MsgApproveAddX509RootCert.fromPartial( data ) }),
    msgAddX509Cert: (data: MsgAddX509Cert): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgAddX509Cert", value: MsgAddX509Cert.fromPartial( data ) }),
    
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
