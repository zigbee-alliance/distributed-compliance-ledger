import { txClient, queryClient, MissingWalletError , registry} from './module'
// @ts-ignore
import { SpVuexError } from '@starport/vuex'

import { Model } from "./module/types/model/model"
import { ModelVersion } from "./module/types/model/model_version"
import { ModelVersions } from "./module/types/model/model_versions"
import { Product } from "./module/types/model/product"
import { VendorProducts } from "./module/types/model/vendor_products"


export { Model, ModelVersion, ModelVersions, Product, VendorProducts };

async function initTxClient(vuexGetters) {
	return await txClient(vuexGetters['common/wallet/signer'], {
		addr: vuexGetters['common/env/apiTendermint']
	})
}

async function initQueryClient(vuexGetters) {
	return await queryClient({
		addr: vuexGetters['common/env/apiCosmos']
	})
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

function getStructure(template) {
	let structure = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field: any = {}
		field.name = key
		field.type = typeof value
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
					throw new SpVuexError('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryVendorProducts({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryVendorProducts( key.vid)).data
				
					
				commit('QUERY', { query: 'VendorProducts', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryVendorProducts', payload: { options: { all }, params: {...key},query }})
				return getters['getVendorProducts']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryVendorProducts', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModel({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryModel( key.vid,  key.pid)).data
				
					
				commit('QUERY', { query: 'Model', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModel', payload: { options: { all }, params: {...key},query }})
				return getters['getModel']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryModel', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModelAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryModelAll(query)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await queryClient.queryModelAll({...query, 'pagination.key':(<any> value).pagination.next_key})).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ModelAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModelAll', payload: { options: { all }, params: {...key},query }})
				return getters['getModelAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryModelAll', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModelVersion({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryModelVersion( key.vid,  key.pid,  key.software_version)).data
				
					
				commit('QUERY', { query: 'ModelVersion', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModelVersion', payload: { options: { all }, params: {...key},query }})
				return getters['getModelVersion']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryModelVersion', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryModelVersions({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const queryClient=await initQueryClient(rootGetters)
				let value= (await queryClient.queryModelVersions( key.vid,  key.pid)).data
				
					
				commit('QUERY', { query: 'ModelVersions', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryModelVersions', payload: { options: { all }, params: {...key},query }})
				return getters['getModelVersions']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new SpVuexError('QueryClient:QueryModelVersions', 'API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgCreateModelVersion({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateModelVersion(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateModelVersion:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateModelVersion:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateModelVersion({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateModelVersion(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateModelVersion:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateModelVersion:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateModel({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateModel(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateModel:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateModel:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdateModel({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateModel(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateModel:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateModel:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeleteModel({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteModel(value)
				const result = await txClient.signAndBroadcast([msg], {fee: { amount: fee, 
	gas: "200000" }, memo})
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDeleteModel:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteModel:Send', 'Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgCreateModelVersion({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateModelVersion(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateModelVersion:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateModelVersion:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgUpdateModelVersion({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateModelVersion(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateModelVersion:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateModelVersion:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgCreateModel({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgCreateModel(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgCreateModel:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgCreateModel:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgUpdateModel({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgUpdateModel(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgUpdateModel:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgUpdateModel:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		async MsgDeleteModel({ rootGetters }, { value }) {
			try {
				const txClient=await initTxClient(rootGetters)
				const msg = await txClient.msgDeleteModel(value)
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new SpVuexError('TxClient:MsgDeleteModel:Init', 'Could not initialize signing client. Wallet is required.')
				}else{
					throw new SpVuexError('TxClient:MsgDeleteModel:Create', 'Could not create message: ' + e.message)
					
				}
			}
		},
		
	}
}
