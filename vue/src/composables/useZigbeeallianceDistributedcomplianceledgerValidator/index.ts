/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useZigbeeallianceDistributedcomplianceledgerValidator() {
  const client = useClient();
  const QueryValidator = (owner: string,  options: any) => {
    const key = { type: 'QueryValidator',  owner };    
    return useQuery([key], () => {
      const { owner } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryValidator(owner).then( res => res.data );
    }, options);
  }
  
  const QueryValidatorAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryValidatorAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryValidatorAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryLastValidatorPower = (owner: string,  options: any) => {
    const key = { type: 'QueryLastValidatorPower',  owner };    
    return useQuery([key], () => {
      const { owner } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryLastValidatorPower(owner).then( res => res.data );
    }, options);
  }
  
  const QueryLastValidatorPowerAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryLastValidatorPowerAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryLastValidatorPowerAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryProposedDisableValidator = (address: string,  options: any) => {
    const key = { type: 'QueryProposedDisableValidator',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryProposedDisableValidator(address).then( res => res.data );
    }, options);
  }
  
  const QueryProposedDisableValidatorAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryProposedDisableValidatorAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryProposedDisableValidatorAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryDisabledValidator = (address: string,  options: any) => {
    const key = { type: 'QueryDisabledValidator',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryDisabledValidator(address).then( res => res.data );
    }, options);
  }
  
  const QueryDisabledValidatorAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryDisabledValidatorAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryDisabledValidatorAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRejectedDisableValidator = (owner: string,  options: any) => {
    const key = { type: 'QueryRejectedDisableValidator',  owner };    
    return useQuery([key], () => {
      const { owner } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryRejectedDisableValidator(owner).then( res => res.data );
    }, options);
  }
  
  const QueryRejectedDisableValidatorAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRejectedDisableValidatorAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerValidator.query.queryRejectedDisableValidatorAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryValidator,QueryValidatorAll,QueryLastValidatorPower,QueryLastValidatorPowerAll,QueryProposedDisableValidator,QueryProposedDisableValidatorAll,QueryDisabledValidator,QueryDisabledValidatorAll,QueryRejectedDisableValidator,QueryRejectedDisableValidatorAll,
  }
}