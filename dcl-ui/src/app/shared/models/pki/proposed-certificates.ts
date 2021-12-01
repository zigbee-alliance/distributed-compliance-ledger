import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';
import { ProposedCertificate } from './proposed-certificate';

@JsonObject('ProposedCertificates')
export class ProposedCertificates {
  @JsonProperty('total', AminoNumberConverter)
  total = 0;

  @JsonProperty('items', [ProposedCertificate])
  items: ProposedCertificate[] = [];
}
