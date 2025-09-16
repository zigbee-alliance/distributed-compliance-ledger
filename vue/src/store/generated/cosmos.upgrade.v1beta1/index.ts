import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { Plan } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.upgrade.v1beta1/types"
import { SoftwareUpgradeProposal } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.upgrade.v1beta1/types"
import { CancelSoftwareUpgradeProposal } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.upgrade.v1beta1/types"
import { ModuleVersion } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.upgrade.v1beta1/types"


export { Plan, SoftwareUpgradeProposal, CancelSoftwareUpgradeProposal, ModuleVersion };

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
				CurrentPlan: {},
				AppliedPlan: {},
				UpgradedConsensusState: {},
				ModuleVersions: {},
				Authority: {},
				
				_Structure: {
						Plan: getStructure(Plan.fromPartial({})),
						SoftwareUpgradeProposal: getStructure(SoftwareUpgradeProposal.fromPartial({})),
						CancelSoftwareUpgradeProposal: getStructure(CancelSoftwareUpgradeProposal.fromPartial({})),
						ModuleVersion: getStructure(ModuleVersion.fromPartial({})),
						
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
				getCurrentPlan: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CurrentPlan[JSON.stringify(params)] ?? {}
		},
				getAppliedPlan: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AppliedPlan[JSON.stringify(params)] ?? {}
		},
				getUpgradedConsensusState: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.UpgradedConsensusState[JSON.stringify(params)] ?? {}
		},
				getModuleVersions: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ModuleVersions[JSON.stringify(params)] ?? {}
		},
				getAuthority: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Authority[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: cosmos.upgrade.v1beta1 initialized!')
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
		
		
		
		 		
		
		
		async QueryCurrentPlan({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosUpgradeV1Beta1.query.queryCurrentPlan()).data
				
					
				commit('QUERY', { query: 'CurrentPlan', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCurrentPlan', payload: { options: { all }, params: {...key},query }})
				return getters['getCurrentPlan']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCurrentPlan API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAppliedPlan({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosUpgradeV1Beta1.query.queryAppliedPlan( key.name)).data
				
					
				commit('QUERY', { query: 'AppliedPlan', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAppliedPlan', payload: { options: { all }, params: {...key},query }})
				return getters['getAppliedPlan']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAppliedPlan API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryUpgradedConsensusState({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosUpgradeV1Beta1.query.queryUpgradedConsensusState( key.last_height)).data
				
					
				commit('QUERY', { query: 'UpgradedConsensusState', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryUpgradedConsensusState', payload: { options: { all }, params: {...key},query }})
				return getters['getUpgradedConsensusState']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryUpgradedConsensusState API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModuleVersions({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosUpgradeV1Beta1.query.queryModuleVersions(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosUpgradeV1Beta1.query.queryModuleVersions({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ModuleVersions', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModuleVersions', payload: { options: { all }, params: {...key},query }})
				return getters['getModuleVersions']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryModuleVersions API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAuthority({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosUpgradeV1Beta1.query.queryAuthority()).data
				
					
				commit('QUERY', { query: 'Authority', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAuthority', payload: { options: { all }, params: {...key},query }})
				return getters['getAuthority']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAuthority API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSoftwareUpgrade({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.CosmosUpgradeV1Beta1.tx.sendMsgSoftwareUpgrade({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSoftwareUpgrade:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSoftwareUpgrade:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCancelUpgrade({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.CosmosUpgradeV1Beta1.tx.sendMsgCancelUpgrade({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelUpgrade:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCancelUpgrade:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSoftwareUpgrade({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosUpgradeV1Beta1.tx.msgSoftwareUpgrade({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSoftwareUpgrade:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSoftwareUpgrade:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCancelUpgrade({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosUpgradeV1Beta1.tx.msgCancelUpgrade({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCancelUpgrade:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCancelUpgrade:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}