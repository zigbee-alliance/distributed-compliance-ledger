import { JsonObject, JsonProperty } from 'json2typescript';

@JsonObject('KeyInfo')
export class KeyInfo {
  @JsonProperty('type', String)
  type: string = undefined;

  @JsonProperty('name', String)
  name: string = undefined;

  @JsonProperty('pubkey', String)
  publicKey: string = undefined;

  @JsonProperty('address', String)
  address: string = undefined;
}
