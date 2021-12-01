import { JsonObject, JsonProperty } from 'json2typescript';
import { Account } from './acount';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('ResultAccountHeaders')
export class ResultAccountHeaders {
  @JsonProperty('total', AminoNumberConverter)
  total: number = undefined;

  @JsonProperty('items', [Account])
  items: Account[] = undefined;
}
