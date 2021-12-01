import { JsonObject, JsonProperty } from 'json2typescript';
import { Description } from './validator-description';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('Validator')
export class Validator {
  @JsonProperty('description', Description)
  description = new Description();

  @JsonProperty('validator_address', String)
  address = '';

  @JsonProperty('validator_pubkey', String)
  pubkey = '';

  @JsonProperty('owner', String)
  owner = '';

  @JsonProperty('power', AminoNumberConverter)
  power = 0;

  @JsonProperty('jailed', Boolean)
  jailed = false;

  @JsonProperty('jailed_reason', String, true)
  jailedReason = null;

  constructor() {}
}
