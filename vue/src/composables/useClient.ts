import { Client } from 'zigbee-alliance-distributed-compliance-ledger-client-ts'
import { env } from '../env';

const useClientInstance = () => {
  const client = new Client(env);
  return client;
};
let clientInstance: ReturnType<typeof useClientInstance>;

export const useClient = () => {
  if (!clientInstance) {
    clientInstance = useClientInstance();
  }
  return clientInstance;
};