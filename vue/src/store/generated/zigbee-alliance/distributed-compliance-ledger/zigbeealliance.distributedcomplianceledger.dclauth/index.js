import { txClient, queryClient, MissingWalletError, registry } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { Account } from "./module/types/dclauth/account";
import { AccountStat } from "./module/types/dclauth/account_stat";
import { PendingAccount } from "./module/types/dclauth/pending_account";
import { PendingAccountRevocation } from "./module/types/dclauth/pending_account_revocation";
export { Account, AccountStat, PendingAccount, PendingAccountRevocation };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        Account: {},
        AccountAll: {},
        PendingAccount: {},
        PendingAccountAll: {},
        PendingAccountRevocation: {},
        PendingAccountRevocationAll: {},
        AccountStat: {},
        _Structure: {
            Account: getStructure(Account.fromPartial({})),
            AccountStat: getStructure(AccountStat.fromPartial({})),
            PendingAccount: getStructure(PendingAccount.fromPartial({})),
            PendingAccountRevocation: getStructure(PendingAccountRevocation.fromPartial({})),
        },
        _Registry: registry,
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(JSON.stringify(subscription));
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(JSON.stringify(subscription));
        }
    },
    getters: {
        getAccount: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Account[JSON.stringify(params)] ?? {};
        },
        getAccountAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.AccountAll[JSON.stringify(params)] ?? {};
        },
        getPendingAccount: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PendingAccount[JSON.stringify(params)] ?? {};
        },
        getPendingAccountAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PendingAccountAll[JSON.stringify(params)] ?? {};
        },
        getPendingAccountRevocation: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PendingAccountRevocation[JSON.stringify(params)] ?? {};
        },
        getPendingAccountRevocationAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PendingAccountRevocationAll[JSON.stringify(params)] ?? {};
        },
        getAccountStat: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.AccountStat[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        },
        getRegistry: (state) => {
            return state._Registry;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: zigbeealliance.distributedcomplianceledger.dclauth initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    const sub = JSON.parse(subscription);
                    await dispatch(sub.action, sub.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryAccount({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryAccount(key.address)).data;
                commit('QUERY', { query: 'Account', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryAccount', payload: { options: { all }, params: { ...key }, query } });
                return getters['getAccount']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryAccount', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryAccountAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryAccountAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryAccountAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'AccountAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryAccountAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getAccountAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryAccountAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryPendingAccount({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryPendingAccount(key.address)).data;
                commit('QUERY', { query: 'PendingAccount', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryPendingAccount', payload: { options: { all }, params: { ...key }, query } });
                return getters['getPendingAccount']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryPendingAccount', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryPendingAccountAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryPendingAccountAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryPendingAccountAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'PendingAccountAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryPendingAccountAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getPendingAccountAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryPendingAccountAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryPendingAccountRevocation({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryPendingAccountRevocation(key.address)).data;
                commit('QUERY', { query: 'PendingAccountRevocation', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryPendingAccountRevocation', payload: { options: { all }, params: { ...key }, query } });
                return getters['getPendingAccountRevocation']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryPendingAccountRevocation', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryPendingAccountRevocationAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryPendingAccountRevocationAll(query)).data;
                while (all && value.pagination && value.pagination.next_key != null) {
                    let next_values = (await queryClient.queryPendingAccountRevocationAll({ ...query, 'pagination.key': value.pagination.next_key })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'PendingAccountRevocationAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryPendingAccountRevocationAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getPendingAccountRevocationAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryPendingAccountRevocationAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryAccountStat({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params, query = null }) {
            try {
                const key = params ?? {};
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryAccountStat()).data;
                commit('QUERY', { query: 'AccountStat', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryAccountStat', payload: { options: { all }, params: { ...key }, query } });
                return getters['getAccountStat']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryAccountStat', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgProposeAddAccount({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeAddAccount(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeAddAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeAddAccount:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgApproveRevokeAccount({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveRevokeAccount(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveRevokeAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveRevokeAccount:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgProposeRevokeAccount({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeRevokeAccount(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeRevokeAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeRevokeAccount:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgApproveAddAccount({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveAddAccount(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveAddAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveAddAccount:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgProposeAddAccount({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeAddAccount(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeAddAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeAddAccount:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgApproveRevokeAccount({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveRevokeAccount(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveRevokeAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveRevokeAccount:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgProposeRevokeAccount({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgProposeRevokeAccount(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgProposeRevokeAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgProposeRevokeAccount:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgApproveAddAccount({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgApproveAddAccount(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgApproveAddAccount:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgApproveAddAccount:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
