import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('BlockId')
export class BlockId {
  @JsonProperty('hash', String)
  hash: string = undefined;
}
