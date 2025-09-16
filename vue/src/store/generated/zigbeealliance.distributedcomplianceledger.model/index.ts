import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { Model } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.model/types"
import { ModelVersion } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.model/types"
import { ModelVersions } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.model/types"
import { Product } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.model/types"
import { VendorProducts } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.model/types"


export { Model, ModelVersion, ModelVersions, Product, VendorProducts };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				VendorProducts: {},
				Model: {},
				ModelAll: {},
				ModelVersion: {},
				ModelVersions: {},
				
				_Structure: {
						Model: getStructure(Model.fromPartial({})),
						ModelVersion: getStructure(ModelVersion.fromPartial({})),
						ModelVersions: getStructure(ModelVersions.fromPartial({})),
						Product: getStructure(Product.fromPartial({})),
						VendorProducts: getStructure(VendorProducts.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getVendorProducts: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.VendorProducts[JSON.stringify(params)] ?? {}
		},
				getModel: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Model[JSON.stringify(params)] ?? {}
		},
				getModelAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ModelAll[JSON.stringify(params)] ?? {}
		},
				getModelVersion: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ModelVersion[JSON.stringify(params)] ?? {}
		},
				getModelVersions: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ModelVersions[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: zigbeealliance.distributedcomplianceledger.model initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryVendorProducts({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryVendorProducts( key.vid)).data
				
					
				commit('QUERY', { query: 'VendorProducts', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVendorProducts', payload: { options: { all }, params: {...key},query }})
				return getters['getVendorProducts']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryVendorProducts API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModel({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModel( key.vid,  key.pid)).data
				
					
				commit('QUERY', { query: 'Model', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModel', payload: { options: { all }, params: {...key},query }})
				return getters['getModel']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryModel API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModelAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModelAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModelAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ModelAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModelAll', payload: { options: { all }, params: {...key},query }})
				return getters['getModelAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryModelAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModelVersion({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModelVersion( key.vid,  key.pid,  key.softwareVersion)).data
				
					
				commit('QUERY', { query: 'ModelVersion', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModelVersion', payload: { options: { all }, params: {...key},query }})
				return getters['getModelVersion']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryModelVersion API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModelVersions({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModelVersions( key.vid,  key.pid)).data
				
					
				commit('QUERY', { query: 'ModelVersions', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModelVersions', payload: { options: { all }, params: {...key},query }})
				return getters['getModelVersions']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryModelVersions API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgUpdateModelVersion({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.sendMsgUpdateModelVersion({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateModelVersion:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateModelVersion:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateModelVersion({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.sendMsgCreateModelVersion({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateModelVersion:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateModelVersion:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateModel({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.sendMsgCreateModel({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateModel:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateModel:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteModel({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.sendMsgDeleteModel({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteModel:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteModel:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteModelVersion({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.sendMsgDeleteModelVersion({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteModelVersion:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeleteModelVersion:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateModel({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.sendMsgUpdateModel({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateModel:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdateModel:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgUpdateModelVersion({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.msgUpdateModelVersion({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateModelVersion:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateModelVersion:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateModelVersion({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.msgCreateModelVersion({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateModelVersion:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateModelVersion:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateModel({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.msgCreateModel({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateModel:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateModel:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteModel({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.msgDeleteModel({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteModel:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteModel:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeleteModelVersion({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.msgDeleteModelVersion({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeleteModelVersion:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeleteModelVersion:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdateModel({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerModel.tx.msgUpdateModel({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdateModel:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdateModel:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}