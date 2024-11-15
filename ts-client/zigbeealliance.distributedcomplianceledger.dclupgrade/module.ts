// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgRejectUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgProposeUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";
import { MsgApproveUpgrade } from "./types/zigbeealliance/distributedcomplianceledger/dclupgrade/tx";

import { ApprovedUpgrade as typeApprovedUpgrade} from "./types"
import { Grant as typeGrant} from "./types"
import { ProposedUpgrade as typeProposedUpgrade} from "./types"
import { RejectedUpgrade as typeRejectedUpgrade} from "./types"

export { MsgRejectUpgrade, MsgProposeUpgrade, MsgApproveUpgrade };

type sendMsgRejectUpgradeParams = {
  value: MsgRejectUpgrade,
  fee?: StdFee,
  memo?: string
};

type sendMsgProposeUpgradeParams = {
  value: MsgProposeUpgrade,
  fee?: StdFee,
  memo?: string
};

type sendMsgApproveUpgradeParams = {
  value: MsgApproveUpgrade,
  fee?: StdFee,
  memo?: string
};


type msgRejectUpgradeParams = {
  value: MsgRejectUpgrade,
};

type msgProposeUpgradeParams = {
  value: MsgProposeUpgrade,
};

type msgApproveUpgradeParams = {
  value: MsgApproveUpgrade,
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
		
		async sendMsgRejectUpgrade({ value, fee, memo }: sendMsgRejectUpgradeParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRejectUpgrade: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRejectUpgrade({ value: MsgRejectUpgrade.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRejectUpgrade: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgProposeUpgrade({ value, fee, memo }: sendMsgProposeUpgradeParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgProposeUpgrade: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgProposeUpgrade({ value: MsgProposeUpgrade.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgProposeUpgrade: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgApproveUpgrade({ value, fee, memo }: sendMsgApproveUpgradeParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgApproveUpgrade: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgApproveUpgrade({ value: MsgApproveUpgrade.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgApproveUpgrade: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgRejectUpgrade({ value }: msgRejectUpgradeParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgRejectUpgrade", value: MsgRejectUpgrade.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRejectUpgrade: Could not create message: ' + e.message)
			}
		},
		
		msgProposeUpgrade({ value }: msgProposeUpgradeParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgProposeUpgrade", value: MsgProposeUpgrade.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgProposeUpgrade: Could not create message: ' + e.message)
			}
		},
		
		msgApproveUpgrade({ value }: msgApproveUpgradeParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.dclupgrade.MsgApproveUpgrade", value: MsgApproveUpgrade.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgApproveUpgrade: Could not create message: ' + e.message)
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
						ApprovedUpgrade: getStructure(typeApprovedUpgrade.fromPartial({})),
						Grant: getStructure(typeGrant.fromPartial({})),
						ProposedUpgrade: getStructure(typeProposedUpgrade.fromPartial({})),
						RejectedUpgrade: getStructure(typeRejectedUpgrade.fromPartial({})),
						
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
			ZigbeeallianceDistributedcomplianceledgerDclupgrade: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;