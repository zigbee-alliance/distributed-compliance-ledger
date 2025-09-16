import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { Validator } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.base.tendermint.v1beta1/types"
import { VersionInfo } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.base.tendermint.v1beta1/types"
import { Module } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.base.tendermint.v1beta1/types"
import { ProofOp } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.base.tendermint.v1beta1/types"
import { ProofOps } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.base.tendermint.v1beta1/types"
import { Block } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.base.tendermint.v1beta1/types"
import { Header } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.base.tendermint.v1beta1/types"


export { Validator, VersionInfo, Module, ProofOp, ProofOps, Block, Header };

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
				GetNodeInfo: {},
				GetSyncing: {},
				GetLatestBlock: {},
				GetBlockByHeight: {},
				GetLatestValidatorSet: {},
				GetValidatorSetByHeight: {},
				ABCIQuery: {},
				
				_Structure: {
						Validator: getStructure(Validator.fromPartial({})),
						VersionInfo: getStructure(VersionInfo.fromPartial({})),
						Module: getStructure(Module.fromPartial({})),
						ProofOp: getStructure(ProofOp.fromPartial({})),
						ProofOps: getStructure(ProofOps.fromPartial({})),
						Block: getStructure(Block.fromPartial({})),
						Header: getStructure(Header.fromPartial({})),
						
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
				getGetNodeInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetNodeInfo[JSON.stringify(params)] ?? {}
		},
				getGetSyncing: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetSyncing[JSON.stringify(params)] ?? {}
		},
				getGetLatestBlock: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetLatestBlock[JSON.stringify(params)] ?? {}
		},
				getGetBlockByHeight: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetBlockByHeight[JSON.stringify(params)] ?? {}
		},
				getGetLatestValidatorSet: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetLatestValidatorSet[JSON.stringify(params)] ?? {}
		},
				getGetValidatorSetByHeight: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.GetValidatorSetByHeight[JSON.stringify(params)] ?? {}
		},
				getABCIQuery: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ABCIQuery[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: cosmos.base.tendermint.v1beta1 initialized!')
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
		
		
		
		 		
		
		
		async ServiceGetNodeInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosBaseTendermintV1Beta1.query.serviceGetNodeInfo()).data
				
					
				commit('QUERY', { query: 'GetNodeInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetNodeInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getGetNodeInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetNodeInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetSyncing({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosBaseTendermintV1Beta1.query.serviceGetSyncing()).data
				
					
				commit('QUERY', { query: 'GetSyncing', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetSyncing', payload: { options: { all }, params: {...key},query }})
				return getters['getGetSyncing']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetSyncing API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetLatestBlock({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosBaseTendermintV1Beta1.query.serviceGetLatestBlock()).data
				
					
				commit('QUERY', { query: 'GetLatestBlock', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetLatestBlock', payload: { options: { all }, params: {...key},query }})
				return getters['getGetLatestBlock']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetLatestBlock API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetBlockByHeight({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosBaseTendermintV1Beta1.query.serviceGetBlockByHeight( key.height)).data
				
					
				commit('QUERY', { query: 'GetBlockByHeight', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetBlockByHeight', payload: { options: { all }, params: {...key},query }})
				return getters['getGetBlockByHeight']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetBlockByHeight API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetLatestValidatorSet({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosBaseTendermintV1Beta1.query.serviceGetLatestValidatorSet(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosBaseTendermintV1Beta1.query.serviceGetLatestValidatorSet({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetLatestValidatorSet', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetLatestValidatorSet', payload: { options: { all }, params: {...key},query }})
				return getters['getGetLatestValidatorSet']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetLatestValidatorSet API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceGetValidatorSetByHeight({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosBaseTendermintV1Beta1.query.serviceGetValidatorSetByHeight( key.height, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosBaseTendermintV1Beta1.query.serviceGetValidatorSetByHeight( key.height, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'GetValidatorSetByHeight', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceGetValidatorSetByHeight', payload: { options: { all }, params: {...key},query }})
				return getters['getGetValidatorSetByHeight']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceGetValidatorSetByHeight API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async ServiceABCIQuery({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosBaseTendermintV1Beta1.query.serviceABCIQuery(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosBaseTendermintV1Beta1.query.serviceABCIQuery({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ABCIQuery', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'ServiceABCIQuery', payload: { options: { all }, params: {...key},query }})
				return getters['getABCIQuery']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:ServiceABCIQuery API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
	}
}