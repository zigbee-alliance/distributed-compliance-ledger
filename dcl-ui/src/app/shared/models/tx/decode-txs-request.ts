import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('DecodeTxsRequest')
export class DecodeTxsRequest {
  @JsonProperty('txs', [String])
  txs: string[] = undefined;
}
