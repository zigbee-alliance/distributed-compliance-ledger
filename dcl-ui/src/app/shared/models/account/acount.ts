import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';
import { PubKey } from './pub-key';

@JsonObject('Account')
export class Account {
  @JsonProperty('address', String)
  address: string = undefined;

  @JsonProperty('public_key', PubKey)
  pubKey: PubKey = undefined;

  @JsonProperty('roles', [String])
  roles: string[] = undefined;

  @JsonProperty('account_number', AminoNumberConverter)
  accountNumber: number = undefined;

  @JsonProperty('sequence', AminoNumberConverter)
  sequence: number = undefined;
}
