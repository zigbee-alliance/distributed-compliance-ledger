import { JsonObject, JsonProperty } from 'json2typescript';
import { KeyInfo } from './key-info';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('KeyInfosResponse')
export class ResultKeyInfos {
  @JsonProperty('total', AminoNumberConverter)
  total: number = undefined;

  @JsonProperty('items', [KeyInfo])
  items: KeyInfo[] = undefined;
}
