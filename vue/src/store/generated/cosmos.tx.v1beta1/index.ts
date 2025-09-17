import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { Tx } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { TxRaw } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { SignDoc } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { SignDocDirectAux } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { TxBody } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { AuthInfo } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { SignerInfo } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { ModeInfo } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { ModeInfo_Single } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { ModeInfo_Multi } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { Fee } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { Tip } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"
import { AuxSignerData } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.tx.v1beta1/types"


export { Tx, TxRaw, SignDoc, SignDocDirectAux, TxBody, AuthInfo, SignerInfo, ModeInfo, ModeInfo_Single, ModeInfo_Multi, Fee, Tip, AuxSignerData };

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
				Simulate: {},
				GetTx: {},
				BroadcastTx: {},
				GetTxsEvent: {},
				GetBlockWithTxs: {},
				TxDecode: {},
				TxEncode: {},
				TxEncodeAmino: {},
				TxDecodeAmino: {},
				
				_Structure: {
						Tx: getStructure(Tx.fromPartial({})),
						TxRaw: getStructure(TxRaw.fromPartial({})),
						SignDoc: getStructure(SignDoc.fromPartial({})),
						SignDocDirectAux: getStructure(SignDocDirectAux.fromPartial({})),
						TxBody: getStructure(TxBody.fromPartial({})),
						AuthInfo: getStructure(AuthInfo.fromPartial({})),
						SignerInfo: getStructure(SignerInfo.fromPartial({})),
						ModeInfo: getStructure(ModeInfo.fromPartial({})),
						ModeInfo_Single: getStructure(ModeInfo_Single.fromPartial({})),
						ModeInfo_Multi: getStructure(ModeInfo_Multi.fromPartial({})),
						Fee: getStructure(Fee.fromPartial({})),
						Tip: getStructure(Tip.fromPartial({})),
						AuxSignerData: getStructure(AuxSignerData.fromPartial({})),
						
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
				getSimulate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Simulate[JSON.stringify(params)] ?? {}
		},
				getGetTx: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetTx[JSON.stringify(params)] ?? {}
		},
				getBroadcastTx: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.BroadcastTx[JSON.stringify(params)] ?? {}
		},
				getGetTxsEvent: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetTxsEvent[JSON.stringify(params)] ?? {}
		},
				getGetBlockWithTxs: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetBlockWithTxs[JSON.stringify(params)] ?? {}
		},
				getTxDecode: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TxDecode[JSON.stringify(params)] ?? {}
		},
				getTxEncode: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TxEncode[JSON.stringify(params)] ?? {}
		},
				getTxEncodeAmino: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TxEncodeAmino[JSON.stringify(params)] ?? {}
		},
				getTxDecodeAmino: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.TxDecodeAmino[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: cosmos.tx.v1beta1 initialized!')
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
		
		
		
		 		
		
		
		async ServiceSimulate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceSimulate({...key})).data
				
					
				commit('QUERY', { query: 'Simulate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceSimulate', payload: { options: { all }, params: {...key},query }})
				return getters['getSimulate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceSimulate API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetTx({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceGetTx( key.hash)).data
				
					
				commit('QUERY', { query: 'GetTx', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetTx', payload: { options: { all }, params: {...key},query }})
				return getters['getGetTx']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetTx API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceBroadcastTx({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceBroadcastTx({...key})).data
				
					
				commit('QUERY', { query: 'BroadcastTx', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceBroadcastTx', payload: { options: { all }, params: {...key},query }})
				return getters['getBroadcastTx']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceBroadcastTx API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetTxsEvent({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceGetTxsEvent(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosTxV1Beta1.query.serviceGetTxsEvent({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetTxsEvent', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetTxsEvent', payload: { options: { all }, params: {...key},query }})
				return getters['getGetTxsEvent']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetTxsEvent API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetBlockWithTxs({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceGetBlockWithTxs( key.height, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosTxV1Beta1.query.serviceGetBlockWithTxs( key.height, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetBlockWithTxs', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetBlockWithTxs', payload: { options: { all }, params: {...key},query }})
				return getters['getGetBlockWithTxs']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetBlockWithTxs API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceTxDecode({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceTxDecode({...key})).data
				
					
				commit('QUERY', { query: 'TxDecode', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceTxDecode', payload: { options: { all }, params: {...key},query }})
				return getters['getTxDecode']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceTxDecode API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceTxEncode({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceTxEncode({...key})).data
				
					
				commit('QUERY', { query: 'TxEncode', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceTxEncode', payload: { options: { all }, params: {...key},query }})
				return getters['getTxEncode']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceTxEncode API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceTxEncodeAmino({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceTxEncodeAmino({...key})).data
				
					
				commit('QUERY', { query: 'TxEncodeAmino', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceTxEncodeAmino', payload: { options: { all }, params: {...key},query }})
				return getters['getTxEncodeAmino']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceTxEncodeAmino API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceTxDecodeAmino({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosTxV1Beta1.query.serviceTxDecodeAmino({...key})).data
				
					
				commit('QUERY', { query: 'TxDecodeAmino', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceTxDecodeAmino', payload: { options: { all }, params: {...key},query }})
				return getters['getTxDecodeAmino']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceTxDecodeAmino API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
	}
}