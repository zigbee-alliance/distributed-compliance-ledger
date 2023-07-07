// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgEnableValidator } from "./types/validator/tx";
import { MsgProposeDisableValidator } from "./types/validator/tx";
import { MsgCreateValidator } from "./types/validator/tx";
import { MsgRejectDisableValidator } from "./types/validator/tx";
import { MsgApproveDisableValidator } from "./types/validator/tx";
import { MsgDisableValidator } from "./types/validator/tx";


const types = [
  ["/zigbeealliance.distributedcomplianceledger.validator.MsgEnableValidator", MsgEnableValidator],
  ["/zigbeealliance.distributedcomplianceledger.validator.MsgProposeDisableValidator", MsgProposeDisableValidator],
  ["/zigbeealliance.distributedcomplianceledger.validator.MsgCreateValidator", MsgCreateValidator],
  ["/zigbeealliance.distributedcomplianceledger.validator.MsgRejectDisableValidator", MsgRejectDisableValidator],
  ["/zigbeealliance.distributedcomplianceledger.validator.MsgApproveDisableValidator", MsgApproveDisableValidator],
  ["/zigbeealliance.distributedcomplianceledger.validator.MsgDisableValidator", MsgDisableValidator],
  
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
    msgEnableValidator: (data: MsgEnableValidator): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgEnableValidator", value: MsgEnableValidator.fromPartial( data ) }),
    msgProposeDisableValidator: (data: MsgProposeDisableValidator): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgProposeDisableValidator", value: MsgProposeDisableValidator.fromPartial( data ) }),
    msgCreateValidator: (data: MsgCreateValidator): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgCreateValidator", value: MsgCreateValidator.fromPartial( data ) }),
    msgRejectDisableValidator: (data: MsgRejectDisableValidator): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgRejectDisableValidator", value: MsgRejectDisableValidator.fromPartial( data ) }),
    msgApproveDisableValidator: (data: MsgApproveDisableValidator): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgApproveDisableValidator", value: MsgApproveDisableValidator.fromPartial( data ) }),
    msgDisableValidator: (data: MsgDisableValidator): EncodeObject => ({ typeUrl: "/zigbeealliance.distributedcomplianceledger.validator.MsgDisableValidator", value: MsgDisableValidator.fromPartial( data ) }),
    
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
