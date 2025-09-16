import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { ApprovedUpgrade } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.dclupgrade/types"
import { Grant } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.dclupgrade/types"
import { ProposedUpgrade } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.dclupgrade/types"
import { RejectedUpgrade } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.dclupgrade/types"


export { ApprovedUpgrade, Grant, ProposedUpgrade, RejectedUpgrade };

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
				ProposedUpgrade: {},
				ProposedUpgradeAll: {},
				ApprovedUpgrade: {},
				ApprovedUpgradeAll: {},
				RejectedUpgrade: {},
				RejectedUpgradeAll: {},
				
				_Structure: {
						ApprovedUpgrade: getStructure(ApprovedUpgrade.fromPartial({})),
						Grant: getStructure(Grant.fromPartial({})),
						ProposedUpgrade: getStructure(ProposedUpgrade.fromPartial({})),
						RejectedUpgrade: getStructure(RejectedUpgrade.fromPartial({})),
						
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
				getProposedUpgrade: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedUpgrade[JSON.stringify(params)] ?? {}
		},
				getProposedUpgradeAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedUpgradeAll[JSON.stringify(params)] ?? {}
		},
				getApprovedUpgrade: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ApprovedUpgrade[JSON.stringify(params)] ?? {}
		},
				getApprovedUpgradeAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ApprovedUpgradeAll[JSON.stringify(params)] ?? {}
		},
				getRejectedUpgrade: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RejectedUpgrade[JSON.stringify(params)] ?? {}
		},
				getRejectedUpgradeAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RejectedUpgradeAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: zigbeealliance.distributedcomplianceledger.dclupgrade initialized!')
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
		
		
		
		 		
		
		
		async QueryProposedUpgrade({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryProposedUpgrade( key.name)).data
				
					
				commit('QUERY', { query: 'ProposedUpgrade', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedUpgrade', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedUpgrade']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedUpgrade API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposedUpgradeAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryProposedUpgradeAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryProposedUpgradeAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ProposedUpgradeAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedUpgradeAll', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedUpgradeAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedUpgradeAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryApprovedUpgrade({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryApprovedUpgrade( key.name)).data
				
					
				commit('QUERY', { query: 'ApprovedUpgrade', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryApprovedUpgrade', payload: { options: { all }, params: {...key},query }})
				return getters['getApprovedUpgrade']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryApprovedUpgrade API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryApprovedUpgradeAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryApprovedUpgradeAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryApprovedUpgradeAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ApprovedUpgradeAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryApprovedUpgradeAll', payload: { options: { all }, params: {...key},query }})
				return getters['getApprovedUpgradeAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryApprovedUpgradeAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRejectedUpgrade({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryRejectedUpgrade( key.name)).data
				
					
				commit('QUERY', { query: 'RejectedUpgrade', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRejectedUpgrade', payload: { options: { all }, params: {...key},query }})
				return getters['getRejectedUpgrade']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRejectedUpgrade API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRejectedUpgradeAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryRejectedUpgradeAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryRejectedUpgradeAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RejectedUpgradeAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRejectedUpgradeAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRejectedUpgradeAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRejectedUpgradeAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgApproveUpgrade({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.tx.sendMsgApproveUpgrade({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveUpgrade:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgApproveUpgrade:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgProposeUpgrade({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.tx.sendMsgProposeUpgrade({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeUpgrade:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgProposeUpgrade:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRejectUpgrade({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.tx.sendMsgRejectUpgrade({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectUpgrade:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRejectUpgrade:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgApproveUpgrade({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.tx.msgApproveUpgrade({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveUpgrade:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgApproveUpgrade:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgProposeUpgrade({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.tx.msgProposeUpgrade({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeUpgrade:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgProposeUpgrade:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRejectUpgrade({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.tx.msgRejectUpgrade({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectUpgrade:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRejectUpgrade:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}