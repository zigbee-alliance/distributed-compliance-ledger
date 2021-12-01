import { JsonObject, JsonProperty } from 'json2typescript';
import { Validator } from './validator';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('ValidatorRecords')
export class ValidatorRecords {
  @JsonProperty('total', AminoNumberConverter)
  total: number = undefined;

  @JsonProperty('items', [Validator])
  items: Validator[] = undefined;
}
