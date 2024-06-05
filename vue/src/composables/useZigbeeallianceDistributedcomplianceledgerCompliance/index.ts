/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useZigbeeallianceDistributedcomplianceledgerCompliance() {
  const client = useClient();
  const QueryComplianceInfo = (vid: string, pid: string, softwareVersion: string, certificationType: string,  options: any) => {
    const key = { type: 'QueryComplianceInfo',  vid,  pid,  softwareVersion,  certificationType };    
    return useQuery([key], () => {
      const { vid,  pid,  softwareVersion,  certificationType } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryComplianceInfo(vid, pid, softwareVersion, certificationType).then( res => res.data );
    }, options);
  }
  
  const QueryComplianceInfoAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryComplianceInfoAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryComplianceInfoAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryCertifiedModel = (vid: string, pid: string, softwareVersion: string, certificationType: string,  options: any) => {
    const key = { type: 'QueryCertifiedModel',  vid,  pid,  softwareVersion,  certificationType };    
    return useQuery([key], () => {
      const { vid,  pid,  softwareVersion,  certificationType } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryCertifiedModel(vid, pid, softwareVersion, certificationType).then( res => res.data );
    }, options);
  }
  
  const QueryCertifiedModelAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryCertifiedModelAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryCertifiedModelAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRevokedModel = (vid: string, pid: string, softwareVersion: string, certificationType: string,  options: any) => {
    const key = { type: 'QueryRevokedModel',  vid,  pid,  softwareVersion,  certificationType };    
    return useQuery([key], () => {
      const { vid,  pid,  softwareVersion,  certificationType } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryRevokedModel(vid, pid, softwareVersion, certificationType).then( res => res.data );
    }, options);
  }
  
  const QueryRevokedModelAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRevokedModelAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryRevokedModelAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryProvisionalModel = (vid: string, pid: string, softwareVersion: string, certificationType: string,  options: any) => {
    const key = { type: 'QueryProvisionalModel',  vid,  pid,  softwareVersion,  certificationType };    
    return useQuery([key], () => {
      const { vid,  pid,  softwareVersion,  certificationType } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryProvisionalModel(vid, pid, softwareVersion, certificationType).then( res => res.data );
    }, options);
  }
  
  const QueryProvisionalModelAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryProvisionalModelAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryProvisionalModelAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryDeviceSoftwareCompliance = (cDCertificateId: string,  options: any) => {
    const key = { type: 'QueryDeviceSoftwareCompliance',  cDCertificateId };    
    return useQuery([key], () => {
      const { cDCertificateId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryDeviceSoftwareCompliance(cDCertificateId).then( res => res.data );
    }, options);
  }
  
  const QueryDeviceSoftwareComplianceAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryDeviceSoftwareComplianceAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerCompliance.query.queryDeviceSoftwareComplianceAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryComplianceInfo,QueryComplianceInfoAll,QueryCertifiedModel,QueryCertifiedModelAll,QueryRevokedModel,QueryRevokedModelAll,QueryProvisionalModel,QueryProvisionalModelAll,QueryDeviceSoftwareCompliance,QueryDeviceSoftwareComplianceAll,
  }
}