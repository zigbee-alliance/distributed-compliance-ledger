import { Client, registry, MissingWalletError } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'

import { AllCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { AllCertificatesBySubject } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { AllCertificatesBySubjectKeyId } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { ApprovedCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { ApprovedCertificatesBySubject } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { ApprovedCertificatesBySubjectKeyId } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { ApprovedRootCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { Certificate } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { CertificateIdentifier } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { ChildCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { Grant } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { NocCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { NocCertificatesBySubject } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { NocCertificatesBySubjectKeyID } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { NocCertificatesByVidAndSkid } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { NocIcaCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { NocRootCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { PkiRevocationDistributionPoint } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { PkiRevocationDistributionPointsByIssuerSubjectKeyID } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { ProposedCertificate } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { ProposedCertificateRevocation } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { RejectedCertificate } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { RevokedCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { RevokedNocIcaCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { RevokedNocRootCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { RevokedRootCertificates } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"
import { UniqueCertificate } from "zigbee-alliance-distributed-compliance-ledger-client-ts/zigbeealliance.distributedcomplianceledger.pki/types"


export { AllCertificates, AllCertificatesBySubject, AllCertificatesBySubjectKeyId, ApprovedCertificates, ApprovedCertificatesBySubject, ApprovedCertificatesBySubjectKeyId, ApprovedRootCertificates, Certificate, CertificateIdentifier, ChildCertificates, Grant, NocCertificates, NocCertificatesBySubject, NocCertificatesBySubjectKeyID, NocCertificatesByVidAndSkid, NocIcaCertificates, NocRootCertificates, PkiRevocationDistributionPoint, PkiRevocationDistributionPointsByIssuerSubjectKeyID, ProposedCertificate, ProposedCertificateRevocation, RejectedCertificate, RevokedCertificates, RevokedNocIcaCertificates, RevokedNocRootCertificates, RevokedRootCertificates, UniqueCertificate };

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
				CertificatesAll: {},
				AllCertificatesBySubject: {},
				Certificates: {},
				ApprovedCertificatesAll: {},
				ApprovedCertificatesBySubject: {},
				ApprovedCertificates: {},
				ProposedCertificate: {},
				ProposedCertificateAll: {},
				ChildCertificates: {},
				ProposedCertificateRevocation: {},
				ProposedCertificateRevocationAll: {},
				RevokedCertificates: {},
				RevokedCertificatesAll: {},
				ApprovedRootCertificates: {},
				RevokedRootCertificates: {},
				RejectedCertificate: {},
				RejectedCertificateAll: {},
				PkiRevocationDistributionPoint: {},
				PkiRevocationDistributionPointAll: {},
				PkiRevocationDistributionPointsByIssuerSubjectKeyID: {},
				NocCertificatesAll: {},
				NocCertificatesBySubject: {},
				NocCertificates: {},
				NocCertificatesByVidAndSkid: {},
				NocRootCertificates: {},
				NocRootCertificatesAll: {},
				NocIcaCertificates: {},
				NocIcaCertificatesAll: {},
				RevokedNocRootCertificates: {},
				RevokedNocRootCertificatesAll: {},
				RevokedNocIcaCertificates: {},
				RevokedNocIcaCertificatesAll: {},
				
				_Structure: {
						AllCertificates: getStructure(AllCertificates.fromPartial({})),
						AllCertificatesBySubject: getStructure(AllCertificatesBySubject.fromPartial({})),
						AllCertificatesBySubjectKeyId: getStructure(AllCertificatesBySubjectKeyId.fromPartial({})),
						ApprovedCertificates: getStructure(ApprovedCertificates.fromPartial({})),
						ApprovedCertificatesBySubject: getStructure(ApprovedCertificatesBySubject.fromPartial({})),
						ApprovedCertificatesBySubjectKeyId: getStructure(ApprovedCertificatesBySubjectKeyId.fromPartial({})),
						ApprovedRootCertificates: getStructure(ApprovedRootCertificates.fromPartial({})),
						Certificate: getStructure(Certificate.fromPartial({})),
						CertificateIdentifier: getStructure(CertificateIdentifier.fromPartial({})),
						ChildCertificates: getStructure(ChildCertificates.fromPartial({})),
						Grant: getStructure(Grant.fromPartial({})),
						NocCertificates: getStructure(NocCertificates.fromPartial({})),
						NocCertificatesBySubject: getStructure(NocCertificatesBySubject.fromPartial({})),
						NocCertificatesBySubjectKeyID: getStructure(NocCertificatesBySubjectKeyID.fromPartial({})),
						NocCertificatesByVidAndSkid: getStructure(NocCertificatesByVidAndSkid.fromPartial({})),
						NocIcaCertificates: getStructure(NocIcaCertificates.fromPartial({})),
						NocRootCertificates: getStructure(NocRootCertificates.fromPartial({})),
						PkiRevocationDistributionPoint: getStructure(PkiRevocationDistributionPoint.fromPartial({})),
						PkiRevocationDistributionPointsByIssuerSubjectKeyID: getStructure(PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromPartial({})),
						ProposedCertificate: getStructure(ProposedCertificate.fromPartial({})),
						ProposedCertificateRevocation: getStructure(ProposedCertificateRevocation.fromPartial({})),
						RejectedCertificate: getStructure(RejectedCertificate.fromPartial({})),
						RevokedCertificates: getStructure(RevokedCertificates.fromPartial({})),
						RevokedNocIcaCertificates: getStructure(RevokedNocIcaCertificates.fromPartial({})),
						RevokedNocRootCertificates: getStructure(RevokedNocRootCertificates.fromPartial({})),
						RevokedRootCertificates: getStructure(RevokedRootCertificates.fromPartial({})),
						UniqueCertificate: getStructure(UniqueCertificate.fromPartial({})),
						
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
				getCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.CertificatesAll[JSON.stringify(params)] ?? {}
		},
				getAllCertificatesBySubject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.AllCertificatesBySubject[JSON.stringify(params)] ?? {}
		},
				getCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Certificates[JSON.stringify(params)] ?? {}
		},
				getApprovedCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ApprovedCertificatesAll[JSON.stringify(params)] ?? {}
		},
				getApprovedCertificatesBySubject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ApprovedCertificatesBySubject[JSON.stringify(params)] ?? {}
		},
				getApprovedCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ApprovedCertificates[JSON.stringify(params)] ?? {}
		},
				getProposedCertificate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedCertificate[JSON.stringify(params)] ?? {}
		},
				getProposedCertificateAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedCertificateAll[JSON.stringify(params)] ?? {}
		},
				getChildCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ChildCertificates[JSON.stringify(params)] ?? {}
		},
				getProposedCertificateRevocation: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedCertificateRevocation[JSON.stringify(params)] ?? {}
		},
				getProposedCertificateRevocationAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ProposedCertificateRevocationAll[JSON.stringify(params)] ?? {}
		},
				getRevokedCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RevokedCertificates[JSON.stringify(params)] ?? {}
		},
				getRevokedCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RevokedCertificatesAll[JSON.stringify(params)] ?? {}
		},
				getApprovedRootCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ApprovedRootCertificates[JSON.stringify(params)] ?? {}
		},
				getRevokedRootCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RevokedRootCertificates[JSON.stringify(params)] ?? {}
		},
				getRejectedCertificate: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RejectedCertificate[JSON.stringify(params)] ?? {}
		},
				getRejectedCertificateAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RejectedCertificateAll[JSON.stringify(params)] ?? {}
		},
				getPkiRevocationDistributionPoint: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PkiRevocationDistributionPoint[JSON.stringify(params)] ?? {}
		},
				getPkiRevocationDistributionPointAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PkiRevocationDistributionPointAll[JSON.stringify(params)] ?? {}
		},
				getPkiRevocationDistributionPointsByIssuerSubjectKeyID: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.PkiRevocationDistributionPointsByIssuerSubjectKeyID[JSON.stringify(params)] ?? {}
		},
				getNocCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocCertificatesAll[JSON.stringify(params)] ?? {}
		},
				getNocCertificatesBySubject: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocCertificatesBySubject[JSON.stringify(params)] ?? {}
		},
				getNocCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocCertificates[JSON.stringify(params)] ?? {}
		},
				getNocCertificatesByVidAndSkid: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocCertificatesByVidAndSkid[JSON.stringify(params)] ?? {}
		},
				getNocRootCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocRootCertificates[JSON.stringify(params)] ?? {}
		},
				getNocRootCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocRootCertificatesAll[JSON.stringify(params)] ?? {}
		},
				getNocIcaCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocIcaCertificates[JSON.stringify(params)] ?? {}
		},
				getNocIcaCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.NocIcaCertificatesAll[JSON.stringify(params)] ?? {}
		},
				getRevokedNocRootCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RevokedNocRootCertificates[JSON.stringify(params)] ?? {}
		},
				getRevokedNocRootCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RevokedNocRootCertificatesAll[JSON.stringify(params)] ?? {}
		},
				getRevokedNocIcaCertificates: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RevokedNocIcaCertificates[JSON.stringify(params)] ?? {}
		},
				getRevokedNocIcaCertificatesAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.RevokedNocIcaCertificatesAll[JSON.stringify(params)] ?? {}
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
			console.log('Vuex module: zigbeealliance.distributedcomplianceledger.pki initialized!')
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
		
		
		
		 		
		
		
		async QueryCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'CertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryAllCertificatesBySubject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryAllCertificatesBySubject( key.subject)).data
				
					
				commit('QUERY', { query: 'AllCertificatesBySubject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryAllCertificatesBySubject', payload: { options: { all }, params: {...key},query }})
				return getters['getAllCertificatesBySubject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryAllCertificatesBySubject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryCertificates( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'Certificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryApprovedCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ApprovedCertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryApprovedCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getApprovedCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryApprovedCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryApprovedCertificatesBySubject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedCertificatesBySubject( key.subject)).data
				
					
				commit('QUERY', { query: 'ApprovedCertificatesBySubject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryApprovedCertificatesBySubject', payload: { options: { all }, params: {...key},query }})
				return getters['getApprovedCertificatesBySubject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryApprovedCertificatesBySubject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryApprovedCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedCertificates( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'ApprovedCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryApprovedCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getApprovedCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryApprovedCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposedCertificate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificate( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'ProposedCertificate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedCertificate', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedCertificate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedCertificate API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposedCertificateAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ProposedCertificateAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedCertificateAll', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedCertificateAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedCertificateAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryChildCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryChildCertificates( key.issuer,  key.authorityKeyId)).data
				
					
				commit('QUERY', { query: 'ChildCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryChildCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getChildCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryChildCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposedCertificateRevocation({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateRevocation( key.subject,  key.subjectKeyId, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateRevocation( key.subject,  key.subjectKeyId, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ProposedCertificateRevocation', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedCertificateRevocation', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedCertificateRevocation']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedCertificateRevocation API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryProposedCertificateRevocationAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateRevocationAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateRevocationAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'ProposedCertificateRevocationAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryProposedCertificateRevocationAll', payload: { options: { all }, params: {...key},query }})
				return getters['getProposedCertificateRevocationAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryProposedCertificateRevocationAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRevokedCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedCertificates( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'RevokedCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRevokedCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getRevokedCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRevokedCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRevokedCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RevokedCertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRevokedCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRevokedCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRevokedCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryApprovedRootCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedRootCertificates()).data
				
					
				commit('QUERY', { query: 'ApprovedRootCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryApprovedRootCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getApprovedRootCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryApprovedRootCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRevokedRootCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedRootCertificates()).data
				
					
				commit('QUERY', { query: 'RevokedRootCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRevokedRootCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getRevokedRootCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRevokedRootCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRejectedCertificate({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRejectedCertificate( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'RejectedCertificate', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRejectedCertificate', payload: { options: { all }, params: {...key},query }})
				return getters['getRejectedCertificate']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRejectedCertificate API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRejectedCertificateAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRejectedCertificateAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRejectedCertificateAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RejectedCertificateAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRejectedCertificateAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRejectedCertificateAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRejectedCertificateAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPkiRevocationDistributionPoint({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryPkiRevocationDistributionPoint( key.issuerSubjectKeyID,  key.vid,  key.label)).data
				
					
				commit('QUERY', { query: 'PkiRevocationDistributionPoint', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPkiRevocationDistributionPoint', payload: { options: { all }, params: {...key},query }})
				return getters['getPkiRevocationDistributionPoint']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPkiRevocationDistributionPoint API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPkiRevocationDistributionPointAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryPkiRevocationDistributionPointAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryPkiRevocationDistributionPointAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'PkiRevocationDistributionPointAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPkiRevocationDistributionPointAll', payload: { options: { all }, params: {...key},query }})
				return getters['getPkiRevocationDistributionPointAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPkiRevocationDistributionPointAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryPkiRevocationDistributionPointsByIssuerSubjectKeyID({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryPkiRevocationDistributionPointsByIssuerSubjectKeyID( key.issuerSubjectKeyID)).data
				
					
				commit('QUERY', { query: 'PkiRevocationDistributionPointsByIssuerSubjectKeyID', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryPkiRevocationDistributionPointsByIssuerSubjectKeyID', payload: { options: { all }, params: {...key},query }})
				return getters['getPkiRevocationDistributionPointsByIssuerSubjectKeyID']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryPkiRevocationDistributionPointsByIssuerSubjectKeyID API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NocCertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getNocCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocCertificatesBySubject({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocCertificatesBySubject( key.subject)).data
				
					
				commit('QUERY', { query: 'NocCertificatesBySubject', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocCertificatesBySubject', payload: { options: { all }, params: {...key},query }})
				return getters['getNocCertificatesBySubject']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocCertificatesBySubject API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocCertificates( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'NocCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getNocCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocCertificatesByVidAndSkid({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocCertificatesByVidAndSkid( key.vid,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'NocCertificatesByVidAndSkid', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocCertificatesByVidAndSkid', payload: { options: { all }, params: {...key},query }})
				return getters['getNocCertificatesByVidAndSkid']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocCertificatesByVidAndSkid API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocRootCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocRootCertificates( key.vid)).data
				
					
				commit('QUERY', { query: 'NocRootCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocRootCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getNocRootCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocRootCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocRootCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocRootCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocRootCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NocRootCertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocRootCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getNocRootCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocRootCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocIcaCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocIcaCertificates( key.vid)).data
				
					
				commit('QUERY', { query: 'NocIcaCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocIcaCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getNocIcaCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocIcaCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryNocIcaCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocIcaCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocIcaCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'NocIcaCertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryNocIcaCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getNocIcaCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryNocIcaCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRevokedNocRootCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocRootCertificates( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'RevokedNocRootCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRevokedNocRootCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getRevokedNocRootCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRevokedNocRootCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRevokedNocRootCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocRootCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocRootCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RevokedNocRootCertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRevokedNocRootCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRevokedNocRootCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRevokedNocRootCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRevokedNocIcaCertificates({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocIcaCertificates( key.subject,  key.subjectKeyId)).data
				
					
				commit('QUERY', { query: 'RevokedNocIcaCertificates', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRevokedNocIcaCertificates', payload: { options: { all }, params: {...key},query }})
				return getters['getRevokedNocIcaCertificates']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRevokedNocIcaCertificates API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryRevokedNocIcaCertificatesAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocIcaCertificatesAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocIcaCertificatesAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'RevokedNocIcaCertificatesAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryRevokedNocIcaCertificatesAll', payload: { options: { all }, params: {...key},query }})
				return getters['getRevokedNocIcaCertificatesAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryRevokedNocIcaCertificatesAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgApproveRevokeX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgApproveRevokeX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveRevokeX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgApproveRevokeX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgUpdatePkiRevocationDistributionPoint({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgUpdatePkiRevocationDistributionPoint({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdatePkiRevocationDistributionPoint:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgUpdatePkiRevocationDistributionPoint:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRemoveX509Cert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgRemoveX509Cert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveX509Cert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRemoveX509Cert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAssignVid({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgAssignVid({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAssignVid:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAssignVid:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddNocX509IcaCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgAddNocX509IcaCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddNocX509IcaCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddNocX509IcaCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRemoveNocX509IcaCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgRemoveNocX509IcaCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveNocX509IcaCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRemoveNocX509IcaCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgProposeAddX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgProposeAddX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeAddX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgProposeAddX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRevokeNocX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgRevokeNocX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeNocX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRevokeNocX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDeletePkiRevocationDistributionPoint({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgDeletePkiRevocationDistributionPoint({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeletePkiRevocationDistributionPoint:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDeletePkiRevocationDistributionPoint:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRevokeNocX509IcaCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgRevokeNocX509IcaCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeNocX509IcaCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRevokeNocX509IcaCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgApproveAddX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgApproveAddX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveAddX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgApproveAddX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRemoveNocX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgRemoveNocX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveNocX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRemoveNocX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgProposeRevokeX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgProposeRevokeX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeRevokeX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgProposeRevokeX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddPkiRevocationDistributionPoint({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgAddPkiRevocationDistributionPoint({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddPkiRevocationDistributionPoint:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddPkiRevocationDistributionPoint:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRejectAddX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgRejectAddX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectAddX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRejectAddX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddX509Cert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgAddX509Cert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddX509Cert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddX509Cert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgAddNocX509RootCert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgAddNocX509RootCert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddNocX509RootCert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgAddNocX509RootCert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgRevokeX509Cert({ rootGetters }, { value, fee = {amount: [], gas: "200000"}, memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const fullFee = Array.isArray(fee)  ? {amount: fee, gas: "200000"} :fee;
				const result = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.sendMsgRevokeX509Cert({ value, fee: fullFee, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeX509Cert:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgRevokeX509Cert:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgApproveRevokeX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgApproveRevokeX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveRevokeX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgApproveRevokeX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgUpdatePkiRevocationDistributionPoint({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgUpdatePkiRevocationDistributionPoint({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgUpdatePkiRevocationDistributionPoint:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgUpdatePkiRevocationDistributionPoint:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRemoveX509Cert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgRemoveX509Cert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveX509Cert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRemoveX509Cert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAssignVid({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgAssignVid({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAssignVid:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAssignVid:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddNocX509IcaCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgAddNocX509IcaCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddNocX509IcaCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddNocX509IcaCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRemoveNocX509IcaCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgRemoveNocX509IcaCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveNocX509IcaCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRemoveNocX509IcaCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgProposeAddX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgProposeAddX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeAddX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgProposeAddX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRevokeNocX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgRevokeNocX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeNocX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRevokeNocX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDeletePkiRevocationDistributionPoint({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgDeletePkiRevocationDistributionPoint({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDeletePkiRevocationDistributionPoint:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDeletePkiRevocationDistributionPoint:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRevokeNocX509IcaCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgRevokeNocX509IcaCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeNocX509IcaCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRevokeNocX509IcaCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgApproveAddX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgApproveAddX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgApproveAddX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgApproveAddX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRemoveNocX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgRemoveNocX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRemoveNocX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRemoveNocX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgProposeRevokeX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgProposeRevokeX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgProposeRevokeX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgProposeRevokeX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddPkiRevocationDistributionPoint({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgAddPkiRevocationDistributionPoint({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddPkiRevocationDistributionPoint:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddPkiRevocationDistributionPoint:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRejectAddX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgRejectAddX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRejectAddX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRejectAddX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddX509Cert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgAddX509Cert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddX509Cert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddX509Cert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgAddNocX509RootCert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgAddNocX509RootCert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgAddNocX509RootCert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgAddNocX509RootCert:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgRevokeX509Cert({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.ZigbeeallianceDistributedcomplianceledgerPki.tx.msgRevokeX509Cert({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgRevokeX509Cert:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgRevokeX509Cert:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}