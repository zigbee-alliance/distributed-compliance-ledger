import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { Description } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.validator/types"
import { DisabledValidator } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.validator/types"
import { Grant } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.validator/types"
import { LastValidatorPower } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.validator/types"
import { ProposedDisableValidator } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.validator/types"
import { RejectedDisableValidator } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.validator/types"
import { Validator } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.validator/types"


export { Description, DisabledValidator, Grant, LastValidatorPower, ProposedDisableValidator, RejectedDisableValidator, Validator };

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
				Validator: {},
				ValidatorAll: {},
				LastValidatorPower: {},
				LastValidatorPowerAll: {},
				ProposedDisableValidator: {},
				ProposedDisableValidatorAll: {},
				DisabledValidator: {},
				DisabledValidatorAll: {},
				RejectedDisableValidator: {},
				RejectedDisableValidatorAll: {},
				
				_Structure: {
						Description: getStructure(Description.fromPartial({})),
						DisabledValidator: getStructure(DisabledValidator.fromPartial({})),
						Grant: getStructure(Grant.fromPartial({})),
						LastValidatorPower: getStructure(LastValidatorPower.fromPartial({})),
						ProposedDisableValidator: getStructure(ProposedDisableValidator.fromPartial({})),
						RejectedDisableValidator: getStructure(RejectedDisableValidator.fromPartial({})),
						Validator: getStructure(Validator.fromPartial({})),
						
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
				getValidator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Validator[JSON.stringify(params)] ?? {}
		},
				getValidatorAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ValidatorAll[JSON.stringify(params)] ?? {}
		},
				getLastValidatorPower: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LastValidatorPower[JSON.stringify(params)] ?? {}
		},
				getLastValidatorPowerAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LastValidatorPowerAll[JSON.stringify(params)] ?? {}
		},
				getProposedDisableValidator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedDisableValidator[JSON.stringify(params)] ?? {}
		},
				getProposedDisableValidatorAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedDisableValidatorAll[JSON.stringify(params)] ?? {}
		},
				getDisabledValidator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DisabledValidator[JSON.stringify(params)] ?? {}
		},
				getDisabledValidatorAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.DisabledValidatorAll[JSON.stringify(params)] ?? {}
		},
				getRejectedDisableValidator: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RejectedDisableValidator[JSON.stringify(params)] ?? {}
		},
				getRejectedDisableValidatorAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RejectedDisableValidatorAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: zigbeealliance.distributedcomplianceledger.validator initialized!')
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
		
		
		
		 		
		
		
		async QueryValidator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryValidator( key.owner)).data
				
					
				commit('QUERY', { query: 'Validator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidator', payload: { options: { all }, params: {...key},query }})
				return getters['getValidator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryValidatorAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryValidatorAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ValidatorAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryValidatorAll', payload: { options: { all }, params: {...key},query }})
				return getters['getValidatorAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryValidatorAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLastValidatorPower({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryLastValidatorPower( key.owner)).data
				
					
				commit('QUERY', { query: 'LastValidatorPower', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLastValidatorPower', payload: { options: { all }, params: {...key},query }})
				return getters['getLastValidatorPower']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLastValidatorPower API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLastValidatorPowerAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryLastValidatorPowerAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryLastValidatorPowerAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LastValidatorPowerAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLastValidatorPowerAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLastValidatorPowerAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLastValidatorPowerAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposedDisableValidator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryProposedDisableValidator( key.address)).data
				
					
				commit('QUERY', { query: 'ProposedDisableValidator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedDisableValidator', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedDisableValidator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedDisableValidator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposedDisableValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryProposedDisableValidatorAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryProposedDisableValidatorAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ProposedDisableValidatorAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedDisableValidatorAll', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedDisableValidatorAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedDisableValidatorAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDisabledValidator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryDisabledValidator( key.address)).data
				
					
				commit('QUERY', { query: 'DisabledValidator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDisabledValidator', payload: { options: { all }, params: {...key},query }})
				return getters['getDisabledValidator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDisabledValidator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryDisabledValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryDisabledValidatorAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryDisabledValidatorAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'DisabledValidatorAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryDisabledValidatorAll', payload: { options: { all }, params: {...key},query }})
				return getters['getDisabledValidatorAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryDisabledValidatorAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRejectedDisableValidator({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryRejectedDisableValidator( key.owner)).data
				
					
				commit('QUERY', { query: 'RejectedDisableValidator', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRejectedDisableValidator', payload: { options: { all }, params: {...key},query }})
				return getters['getRejectedDisableValidator']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRejectedDisableValidator API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRejectedDisableValidatorAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryRejectedDisableValidatorAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryRejectedDisableValidatorAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RejectedDisableValidatorAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRejectedDisableValidatorAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRejectedDisableValidatorAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRejectedDisableValidatorAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateValidator({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.sendMsgCreateValidator({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgProposeDisableValidator({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.sendMsgProposeDisableValidator({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeDisableValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgProposeDisableValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgEnableValidator({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.sendMsgEnableValidator({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEnableValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgEnableValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgApproveDisableValidator({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.sendMsgApproveDisableValidator({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveDisableValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgApproveDisableValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDisableValidator({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.sendMsgDisableValidator({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDisableValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDisableValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRejectDisableValidator({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.sendMsgRejectDisableValidator({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectDisableValidator:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRejectDisableValidator:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.msgCreateValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgProposeDisableValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.msgProposeDisableValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeDisableValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgProposeDisableValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgEnableValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.msgEnableValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgEnableValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgEnableValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgApproveDisableValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.msgApproveDisableValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveDisableValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgApproveDisableValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDisableValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.msgDisableValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDisableValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDisableValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRejectDisableValidator({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerValidator.tx.msgRejectDisableValidator({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectDisableValidator:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRejectDisableValidator:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}