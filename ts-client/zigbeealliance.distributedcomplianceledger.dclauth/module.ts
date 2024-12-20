// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgApproveRevokeAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgProposeAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgProposeRevokeAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgRejectAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";
import { MsgApproveAddAccount } from "./types/zigbeealliance/distributedcomplianceledger/dclauth/tx";

import { Account as typeAccount} from "./types"
import { AccountStat as typeAccountStat} from "./types"
import { Grant as typeGrant} from "./types"
import { PendingAccount as typePendingAccount} from "./types"
import { PendingAccountRevocation as typePendingAccountRevocation} from "./types"
import { RejectedAccount as typeRejectedAccount} from "./types"
import { RevokedAccount as typeRevokedAccount} from "./types"

export { MsgApproveRevokeAccount, MsgProposeAddAccount, MsgProposeRevokeAccount, MsgRejectAddAccount, MsgApproveAddAccount };

type sendMsgApproveRevokeAccountParams = {
  value: MsgApproveRevokeAccount,
  fee?: StdFee,
  memo?: string
};

type sendMsgProposeAddAccountParams = {
  value: MsgProposeAddAccount,
  fee?: StdFee,
  memo?: string
};

type sendMsgProposeRevokeAccountParams = {
  value: MsgProposeRevokeAccount,
  fee?: StdFee,
  memo?: string
};

type sendMsgRejectAddAccountParams = {
  value: MsgRejectAddAccount,
  fee?: StdFee,
  memo?: string
};

type sendMsgApproveAddAccountParams = {
  value: MsgApproveAddAccount,
  fee?: StdFee,
  memo?: string
};


type msgApproveRevokeAccountParams = {
  value: MsgApproveRevokeAccount,
};

type msgProposeAddAccountParams = {
  value: MsgProposeAddAccount,
};

type msgProposeRevokeAccountParams = {
  value: MsgProposeRevokeAccount,
};

type msgRejectAddAccountParams = {
  value: MsgRejectAddAccount,
};

type msgApproveAddAccountParams = {
  value: MsgApproveAddAccount,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgApproveRevokeAccount({ value, fee, memo }: sendMsgApproveRevokeAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgApproveRevokeAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgApproveRevokeAccount({ value: MsgApproveRevokeAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgApproveRevokeAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgProposeAddAccount({ value, fee, memo }: sendMsgProposeAddAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgProposeAddAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgProposeAddAccount({ value: MsgProposeAddAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgProposeAddAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgProposeRevokeAccount({ value, fee, memo }: sendMsgProposeRevokeAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgProposeRevokeAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgProposeRevokeAccount({ value: MsgProposeRevokeAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgProposeRevokeAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRejectAddAccount({ value, fee, memo }: sendMsgRejectAddAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRejectAddAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRejectAddAccount({ value: MsgRejectAddAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRejectAddAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgApproveAddAccount({ value, fee, memo }: sendMsgApproveAddAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgApproveAddAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgApproveAddAccount({ value: MsgApproveAddAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgApproveAddAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgApproveRevokeAccount({ value }: msgApproveRevokeAccountParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveRevokeAccount", value: MsgApproveRevokeAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgApproveRevokeAccount: Could not create message: ' + e.message)
			}
		},
		
		msgProposeAddAccount({ value }: msgProposeAddAccountParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount", value: MsgProposeAddAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgProposeAddAccount: Could not create message: ' + e.message)
			}
		},
		
		msgProposeRevokeAccount({ value }: msgProposeRevokeAccountParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeRevokeAccount", value: MsgProposeRevokeAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgProposeRevokeAccount: Could not create message: ' + e.message)
			}
		},
		
		msgRejectAddAccount({ value }: msgRejectAddAccountParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgRejectAddAccount", value: MsgRejectAddAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRejectAddAccount: Could not create message: ' + e.message)
			}
		},
		
		msgApproveAddAccount({ value }: msgApproveAddAccountParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount", value: MsgApproveAddAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgApproveAddAccount: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						Account: getStructure(typeAccount.fromPartial({})),
						AccountStat: getStructure(typeAccountStat.fromPartial({})),
						Grant: getStructure(typeGrant.fromPartial({})),
						PendingAccount: getStructure(typePendingAccount.fromPartial({})),
						PendingAccountRevocation: getStructure(typePendingAccountRevocation.fromPartial({})),
						RejectedAccount: getStructure(typeRejectedAccount.fromPartial({})),
						RevokedAccount: getStructure(typeRevokedAccount.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			ZigbeeallianceDistributedcomplianceledgerDclauth: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;