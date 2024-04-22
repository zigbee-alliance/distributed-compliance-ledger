/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useZigbeeallianceDistributedcomplianceledgerPki() {
  const client = useClient();
  const QueryApprovedCertificates = (subject: string, subjectKeyId: string,  options: any) => {
    const key = { type: 'QueryApprovedCertificates',  subject,  subjectKeyId };    
    return useQuery([key], () => {
      const { subject,  subjectKeyId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedCertificates(subject, subjectKeyId).then( res => res.data );
    }, options);
  }
  
  const QueryApprovedCertificatesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryApprovedCertificatesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedCertificatesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryProposedCertificate = (subject: string, subjectKeyId: string,  options: any) => {
    const key = { type: 'QueryProposedCertificate',  subject,  subjectKeyId };    
    return useQuery([key], () => {
      const { subject,  subjectKeyId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificate(subject, subjectKeyId).then( res => res.data );
    }, options);
  }
  
  const QueryProposedCertificateAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryProposedCertificateAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryChildCertificates = (issuer: string, authorityKeyId: string,  options: any) => {
    const key = { type: 'QueryChildCertificates',  issuer,  authorityKeyId };    
    return useQuery([key], () => {
      const { issuer,  authorityKeyId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryChildCertificates(issuer, authorityKeyId).then( res => res.data );
    }, options);
  }
  
  const QueryProposedCertificateRevocation = (subject: string, subjectKeyId: string, query: any, options: any) => {
    const key = { type: 'QueryProposedCertificateRevocation',  subject,  subjectKeyId, query };    
    return useQuery([key], () => {
      const { subject,  subjectKeyId,query } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateRevocation(subject, subjectKeyId, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryProposedCertificateRevocationAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryProposedCertificateRevocationAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryProposedCertificateRevocationAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRevokedCertificates = (subject: string, subjectKeyId: string,  options: any) => {
    const key = { type: 'QueryRevokedCertificates',  subject,  subjectKeyId };    
    return useQuery([key], () => {
      const { subject,  subjectKeyId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedCertificates(subject, subjectKeyId).then( res => res.data );
    }, options);
  }
  
  const QueryRevokedCertificatesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRevokedCertificatesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedCertificatesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryApprovedRootCertificates = ( options: any) => {
    const key = { type: 'QueryApprovedRootCertificates',  };    
    return useQuery([key], () => {
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedRootCertificates().then( res => res.data );
    }, options);
  }
  
  const QueryRevokedRootCertificates = ( options: any) => {
    const key = { type: 'QueryRevokedRootCertificates',  };    
    return useQuery([key], () => {
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedRootCertificates().then( res => res.data );
    }, options);
  }
  
  const QueryApprovedCertificatesBySubject = (subject: string,  options: any) => {
    const key = { type: 'QueryApprovedCertificatesBySubject',  subject };    
    return useQuery([key], () => {
      const { subject } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryApprovedCertificatesBySubject(subject).then( res => res.data );
    }, options);
  }
  
  const QueryRejectedCertificate = (subject: string, subjectKeyId: string,  options: any) => {
    const key = { type: 'QueryRejectedCertificate',  subject,  subjectKeyId };    
    return useQuery([key], () => {
      const { subject,  subjectKeyId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRejectedCertificate(subject, subjectKeyId).then( res => res.data );
    }, options);
  }
  
  const QueryRejectedCertificateAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRejectedCertificateAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRejectedCertificateAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPkiRevocationDistributionPoint = (issuerSubjectKeyID: string, vid: string, label: string,  options: any) => {
    const key = { type: 'QueryPkiRevocationDistributionPoint',  issuerSubjectKeyID,  vid,  label };    
    return useQuery([key], () => {
      const { issuerSubjectKeyID,  vid,  label } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryPkiRevocationDistributionPoint(issuerSubjectKeyID, vid, label).then( res => res.data );
    }, options);
  }
  
  const QueryPkiRevocationDistributionPointAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryPkiRevocationDistributionPointAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryPkiRevocationDistributionPointAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPkiRevocationDistributionPointsByIssuerSubjectKeyID = (issuerSubjectKeyID: string,  options: any) => {
    const key = { type: 'QueryPkiRevocationDistributionPointsByIssuerSubjectKeyID',  issuerSubjectKeyID };    
    return useQuery([key], () => {
      const { issuerSubjectKeyID } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryPkiRevocationDistributionPointsByIssuerSubjectKeyID(issuerSubjectKeyID).then( res => res.data );
    }, options);
  }
  
  const QueryNocRootCertificates = (vid: string,  options: any) => {
    const key = { type: 'QueryNocRootCertificates',  vid };    
    return useQuery([key], () => {
      const { vid } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocRootCertificates(vid).then( res => res.data );
    }, options);
  }
  
  const QueryNocRootCertificatesByVidAndSkid = (vid: string, subjectKeyId: string,  options: any) => {
    const key = { type: 'QueryNocRootCertificatesByVidAndSkid',  vid,  subjectKeyId };    
    return useQuery([key], () => {
      const { vid,  subjectKeyId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocRootCertificatesByVidAndSkid(vid, subjectKeyId).then( res => res.data );
    }, options);
  }
  
  const QueryNocRootCertificatesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryNocRootCertificatesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocRootCertificatesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryNocIcaCertificates = (vid: string,  options: any) => {
    const key = { type: 'QueryNocIcaCertificates',  vid };    
    return useQuery([key], () => {
      const { vid } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocIcaCertificates(vid).then( res => res.data );
    }, options);
  }
  
  const QueryNocIcaCertificatesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryNocIcaCertificatesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryNocIcaCertificatesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRevokedNocRootCertificates = (subject: string, subjectKeyId: string,  options: any) => {
    const key = { type: 'QueryRevokedNocRootCertificates',  subject,  subjectKeyId };    
    return useQuery([key], () => {
      const { subject,  subjectKeyId } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocRootCertificates(subject, subjectKeyId).then( res => res.data );
    }, options);
  }
  
  const QueryRevokedNocRootCertificatesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRevokedNocRootCertificatesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerPki.query.queryRevokedNocRootCertificatesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryApprovedCertificates,QueryApprovedCertificatesAll,QueryProposedCertificate,QueryProposedCertificateAll,QueryChildCertificates,QueryProposedCertificateRevocation,QueryProposedCertificateRevocationAll,QueryRevokedCertificates,QueryRevokedCertificatesAll,QueryApprovedRootCertificates,QueryRevokedRootCertificates,QueryApprovedCertificatesBySubject,QueryRejectedCertificate,QueryRejectedCertificateAll,QueryPkiRevocationDistributionPoint,QueryPkiRevocationDistributionPointAll,QueryPkiRevocationDistributionPointsByIssuerSubjectKeyID,QueryNocRootCertificates,QueryNocRootCertificatesByVidAndSkid,QueryNocRootCertificatesAll,QueryNocIcaCertificates,QueryNocIcaCertificatesAll,QueryRevokedNocRootCertificates,QueryRevokedNocRootCertificatesAll,
  }
}