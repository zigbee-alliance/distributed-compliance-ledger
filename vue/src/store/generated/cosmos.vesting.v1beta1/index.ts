import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { BaseVestingAccount } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.vesting.v1beta1/types"
import { ContinuousVestingAccount } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.vesting.v1beta1/types"
import { DelayedVestingAccount } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.vesting.v1beta1/types"
import { Period } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.vesting.v1beta1/types"
import { PeriodicVestingAccount } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.vesting.v1beta1/types"
import { PermanentLockedAccount } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.vesting.v1beta1/types"


export { BaseVestingAccount, ContinuousVestingAccount, DelayedVestingAccount, Period, PeriodicVestingAccount, PermanentLockedAccount };

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
				
				_Structure: {
						BaseVestingAccount: getStructure(BaseVestingAccount.fromPartial({})),
						ContinuousVestingAccount: getStructure(ContinuousVestingAccount.fromPartial({})),
						DelayedVestingAccount: getStructure(DelayedVestingAccount.fromPartial({})),
						Period: getStructure(Period.fromPartial({})),
						PeriodicVestingAccount: getStructure(PeriodicVestingAccount.fromPartial({})),
						PermanentLockedAccount: getStructure(PermanentLockedAccount.fromPartial({})),
						
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
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: cosmos.vesting.v1beta1 initialized!')
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
		
		async sendMsgCreatePermanentLockedAccount({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.CosmosVestingV1Beta1.tx.sendMsgCreatePermanentLockedAccount({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreatePermanentLockedAccount:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreatePermanentLockedAccount:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreatePeriodicVestingAccount({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.CosmosVestingV1Beta1.tx.sendMsgCreatePeriodicVestingAccount({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreatePeriodicVestingAccount:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreatePeriodicVestingAccount:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateVestingAccount({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.CosmosVestingV1Beta1.tx.sendMsgCreateVestingAccount({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVestingAccount:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateVestingAccount:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreatePermanentLockedAccount({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosVestingV1Beta1.tx.msgCreatePermanentLockedAccount({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreatePermanentLockedAccount:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreatePermanentLockedAccount:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreatePeriodicVestingAccount({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosVestingV1Beta1.tx.msgCreatePeriodicVestingAccount({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreatePeriodicVestingAccount:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreatePeriodicVestingAccount:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateVestingAccount({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.CosmosVestingV1Beta1.tx.msgCreateVestingAccount({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateVestingAccount:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateVestingAccount:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}