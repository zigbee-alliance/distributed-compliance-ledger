// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgUpdateComplianceInfo } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgCertifyModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgDeleteComplianceInfo } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgProvisionModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";
import { MsgRevokeModel } from "./types/zigbeealliance/distributedcomplianceledger/compliance/tx";

import { CertifiedModel as typeCertifiedModel} from "./types"
import { ComplianceHistoryItem as typeComplianceHistoryItem} from "./types"
import { ComplianceInfo as typeComplianceInfo} from "./types"
import { DeviceSoftwareCompliance as typeDeviceSoftwareCompliance} from "./types"
import { ProvisionalModel as typeProvisionalModel} from "./types"
import { RevokedModel as typeRevokedModel} from "./types"

export { MsgUpdateComplianceInfo, MsgCertifyModel, MsgDeleteComplianceInfo, MsgProvisionModel, MsgRevokeModel };

type sendMsgUpdateComplianceInfoParams = {
  value: MsgUpdateComplianceInfo,
  fee?: StdFee,
  memo?: string
};

type sendMsgCertifyModelParams = {
  value: MsgCertifyModel,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeleteComplianceInfoParams = {
  value: MsgDeleteComplianceInfo,
  fee?: StdFee,
  memo?: string
};

type sendMsgProvisionModelParams = {
  value: MsgProvisionModel,
  fee?: StdFee,
  memo?: string
};

type sendMsgRevokeModelParams = {
  value: MsgRevokeModel,
  fee?: StdFee,
  memo?: string
};


type msgUpdateComplianceInfoParams = {
  value: MsgUpdateComplianceInfo,
};

type msgCertifyModelParams = {
  value: MsgCertifyModel,
};

type msgDeleteComplianceInfoParams = {
  value: MsgDeleteComplianceInfo,
};

type msgProvisionModelParams = {
  value: MsgProvisionModel,
};

type msgRevokeModelParams = {
  value: MsgRevokeModel,
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
		
		async sendMsgUpdateComplianceInfo({ value, fee, memo }: sendMsgUpdateComplianceInfoParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdateComplianceInfo: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdateComplianceInfo({ value: MsgUpdateComplianceInfo.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdateComplianceInfo: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCertifyModel({ value, fee, memo }: sendMsgCertifyModelParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCertifyModel: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCertifyModel({ value: MsgCertifyModel.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCertifyModel: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeleteComplianceInfo({ value, fee, memo }: sendMsgDeleteComplianceInfoParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeleteComplianceInfo: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeleteComplianceInfo({ value: MsgDeleteComplianceInfo.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeleteComplianceInfo: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgProvisionModel({ value, fee, memo }: sendMsgProvisionModelParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgProvisionModel: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgProvisionModel({ value: MsgProvisionModel.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgProvisionModel: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRevokeModel({ value, fee, memo }: sendMsgRevokeModelParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRevokeModel: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRevokeModel({ value: MsgRevokeModel.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRevokeModel: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgUpdateComplianceInfo({ value }: msgUpdateComplianceInfoParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgUpdateComplianceInfo", value: MsgUpdateComplianceInfo.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdateComplianceInfo: Could not create message: ' + e.message)
			}
		},
		
		msgCertifyModel({ value }: msgCertifyModelParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgCertifyModel", value: MsgCertifyModel.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCertifyModel: Could not create message: ' + e.message)
			}
		},
		
		msgDeleteComplianceInfo({ value }: msgDeleteComplianceInfoParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgDeleteComplianceInfo", value: MsgDeleteComplianceInfo.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeleteComplianceInfo: Could not create message: ' + e.message)
			}
		},
		
		msgProvisionModel({ value }: msgProvisionModelParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgProvisionModel", value: MsgProvisionModel.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgProvisionModel: Could not create message: ' + e.message)
			}
		},
		
		msgRevokeModel({ value }: msgRevokeModelParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.compliance.MsgRevokeModel", value: MsgRevokeModel.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRevokeModel: Could not create message: ' + e.message)
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
						CertifiedModel: getStructure(typeCertifiedModel.fromPartial({})),
						ComplianceHistoryItem: getStructure(typeComplianceHistoryItem.fromPartial({})),
						ComplianceInfo: getStructure(typeComplianceInfo.fromPartial({})),
						DeviceSoftwareCompliance: getStructure(typeDeviceSoftwareCompliance.fromPartial({})),
						ProvisionalModel: getStructure(typeProvisionalModel.fromPartial({})),
						RevokedModel: getStructure(typeRevokedModel.fromPartial({})),
						
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
			ZigbeeallianceDistributedcomplianceledgerCompliance: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;