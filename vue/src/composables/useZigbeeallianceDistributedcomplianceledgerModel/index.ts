/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useZigbeeallianceDistributedcomplianceledgerModel() {
  const client = useClient();
  const QueryVendorProducts = (vid: string,  options: any) => {
    const key = { type: 'QueryVendorProducts',  vid };    
    return useQuery([key], () => {
      const { vid } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryVendorProducts(vid).then( res => res.data );
    }, options);
  }
  
  const QueryModel = (vid: string, pid: string,  options: any) => {
    const key = { type: 'QueryModel',  vid,  pid };    
    return useQuery([key], () => {
      const { vid,  pid } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModel(vid, pid).then( res => res.data );
    }, options);
  }
  
  const QueryModelAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryModelAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModelAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryModelVersion = (vid: string, pid: string, softwareVersion: string,  options: any) => {
    const key = { type: 'QueryModelVersion',  vid,  pid,  softwareVersion };    
    return useQuery([key], () => {
      const { vid,  pid,  softwareVersion } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModelVersion(vid, pid, softwareVersion).then( res => res.data );
    }, options);
  }
  
  const QueryModelVersions = (vid: string, pid: string,  options: any) => {
    const key = { type: 'QueryModelVersions',  vid,  pid };    
    return useQuery([key], () => {
      const { vid,  pid } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerModel.query.queryModelVersions(vid, pid).then( res => res.data );
    }, options);
  }
  
  return {QueryVendorProducts,QueryModel,QueryModelAll,QueryModelVersion,QueryModelVersions,
  }
}