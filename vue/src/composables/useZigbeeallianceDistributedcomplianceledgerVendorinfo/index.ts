/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useZigbeeallianceDistributedcomplianceledgerVendorinfo() {
  const client = useClient();
  const QueryVendorInfo = (vendorID: string,  options: any) => {
    const key = { type: 'QueryVendorInfo',  vendorID };    
    return useQuery([key], () => {
      const { vendorID } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerVendorinfo.query.queryVendorInfo(vendorID).then( res => res.data );
    }, options);
  }
  
  const QueryVendorInfoAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryVendorInfoAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerVendorinfo.query.queryVendorInfoAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryVendorInfo,QueryVendorInfoAll,
  }
}