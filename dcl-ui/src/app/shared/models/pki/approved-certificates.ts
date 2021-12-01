import { JsonObject, JsonProperty } from 'json2typescript';
import { ApprovedCertificate } from './approved-certificate';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('ApprovedCertificates')
export class ApprovedCertificates {
  @JsonProperty('total', AminoNumberConverter, true)
  total = 0;

  @JsonProperty('items', [ApprovedCertificate])
  items: ApprovedCertificate[] = [];
}
