import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { BaseAccount } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.auth.v1beta1/types"
import { ModuleAccount } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.auth.v1beta1/types"
import { ModuleCredential } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.auth.v1beta1/types"
import { Params } from "zigbee-alliance-distributed-compliance-ledger-client-ts/cosmos.auth.v1beta1/types"


export { BaseAccount, ModuleAccount, ModuleCredential, Params };

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
				Accounts: {},
				Account: {},
				AccountAddressByID: {},
				Params: {},
				ModuleAccounts: {},
				ModuleAccountByName: {},
				Bech32Prefix: {},
				AddressBytesToString: {},
				AddressStringToBytes: {},
				AccountInfo: {},
				
				_Structure: {
						BaseAccount: getStructure(BaseAccount.fromPartial({})),
						ModuleAccount: getStructure(ModuleAccount.fromPartial({})),
						ModuleCredential: getStructure(ModuleCredential.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
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
				getAccounts: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Accounts[JSON.stringify(params)] ?? {}
		},
				getAccount: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Account[JSON.stringify(params)] ?? {}
		},
				getAccountAddressByID: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AccountAddressByID[JSON.stringify(params)] ?? {}
		},
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getModuleAccounts: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ModuleAccounts[JSON.stringify(params)] ?? {}
		},
				getModuleAccountByName: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ModuleAccountByName[JSON.stringify(params)] ?? {}
		},
				getBech32Prefix: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Bech32Prefix[JSON.stringify(params)] ?? {}
		},
				getAddressBytesToString: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AddressBytesToString[JSON.stringify(params)] ?? {}
		},
				getAddressStringToBytes: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AddressStringToBytes[JSON.stringify(params)] ?? {}
		},
				getAccountInfo: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AccountInfo[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: cosmos.auth.v1beta1 initialized!')
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
		
		
		
		 		
		
		
		async QueryAccounts({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryAccounts(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosAuthV1Beta1.query.queryAccounts({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'Accounts', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccounts', payload: { options: { all }, params: {...key},query }})
				return getters['getAccounts']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccounts API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccount({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryAccount( key.address)).data
				
					
				commit('QUERY', { query: 'Account', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccount', payload: { options: { all }, params: {...key},query }})
				return getters['getAccount']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccount API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccountAddressByID({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryAccountAddressByID( key.id, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.CosmosAuthV1Beta1.query.queryAccountAddressByID( key.id, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'AccountAddressByID', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccountAddressByID', payload: { options: { all }, params: {...key},query }})
				return getters['getAccountAddressByID']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccountAddressByID API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModuleAccounts({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryModuleAccounts()).data
				
					
				commit('QUERY', { query: 'ModuleAccounts', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModuleAccounts', payload: { options: { all }, params: {...key},query }})
				return getters['getModuleAccounts']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryModuleAccounts API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModuleAccountByName({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryModuleAccountByName( key.name)).data
				
					
				commit('QUERY', { query: 'ModuleAccountByName', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModuleAccountByName', payload: { options: { all }, params: {...key},query }})
				return getters['getModuleAccountByName']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryModuleAccountByName API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryBech32Prefix({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryBech32Prefix()).data
				
					
				commit('QUERY', { query: 'Bech32Prefix', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryBech32Prefix', payload: { options: { all }, params: {...key},query }})
				return getters['getBech32Prefix']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryBech32Prefix API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAddressBytesToString({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryAddressBytesToString( key.address_bytes)).data
				
					
				commit('QUERY', { query: 'AddressBytesToString', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAddressBytesToString', payload: { options: { all }, params: {...key},query }})
				return getters['getAddressBytesToString']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAddressBytesToString API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAddressStringToBytes({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryAddressStringToBytes( key.address_string)).data
				
					
				commit('QUERY', { query: 'AddressStringToBytes', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAddressStringToBytes', payload: { options: { all }, params: {...key},query }})
				return getters['getAddressStringToBytes']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAddressStringToBytes API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAccountInfo({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.CosmosAuthV1Beta1.query.queryAccountInfo( key.address)).data
				
					
				commit('QUERY', { query: 'AccountInfo', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAccountInfo', payload: { options: { all }, params: {...key},query }})
				return getters['getAccountInfo']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAccountInfo API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
	}
}