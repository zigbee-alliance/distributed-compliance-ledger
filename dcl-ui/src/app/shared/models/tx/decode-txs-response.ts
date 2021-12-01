import { JsonObject, JsonProperty, Any } from 'json2typescript';

@JsonObject('DecodeTxsResponse')
export class DecodeTxsResponse {
  @JsonProperty('txs', [Any])
  txs: any[] = undefined;
}
