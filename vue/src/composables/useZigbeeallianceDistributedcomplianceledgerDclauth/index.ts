/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useZigbeeallianceDistributedcomplianceledgerDclauth() {
  const client = useClient();
  const QueryAccount = (address: string,  options: any) => {
    const key = { type: 'QueryAccount',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryAccount(address).then( res => res.data );
    }, options);
  }
  
  const QueryAccountAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryAccountAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryAccountAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPendingAccount = (address: string,  options: any) => {
    const key = { type: 'QueryPendingAccount',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryPendingAccount(address).then( res => res.data );
    }, options);
  }
  
  const QueryPendingAccountAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryPendingAccountAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryPendingAccountAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPendingAccountRevocation = (address: string,  options: any) => {
    const key = { type: 'QueryPendingAccountRevocation',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryPendingAccountRevocation(address).then( res => res.data );
    }, options);
  }
  
  const QueryPendingAccountRevocationAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryPendingAccountRevocationAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryPendingAccountRevocationAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryAccountStat = ( options: any) => {
    const key = { type: 'QueryAccountStat',  };    
    return useQuery([key], () => {
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryAccountStat().then( res => res.data );
    }, options);
  }
  
  const QueryRevokedAccount = (address: string,  options: any) => {
    const key = { type: 'QueryRevokedAccount',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryRevokedAccount(address).then( res => res.data );
    }, options);
  }
  
  const QueryRevokedAccountAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRevokedAccountAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryRevokedAccountAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRejectedAccount = (address: string,  options: any) => {
    const key = { type: 'QueryRejectedAccount',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryRejectedAccount(address).then( res => res.data );
    }, options);
  }
  
  const QueryRejectedAccountAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRejectedAccountAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclauth.query.queryRejectedAccountAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryAccount,QueryAccountAll,QueryPendingAccount,QueryPendingAccountAll,QueryPendingAccountRevocation,QueryPendingAccountRevocationAll,QueryAccountStat,QueryRevokedAccount,QueryRevokedAccountAll,QueryRejectedAccount,QueryRejectedAccountAll,
  }
}