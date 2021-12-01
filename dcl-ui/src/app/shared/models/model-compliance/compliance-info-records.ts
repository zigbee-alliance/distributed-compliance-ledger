import { JsonObject, JsonProperty } from 'json2typescript';
import { ComplianceInfo } from './compliance-info';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('ComplianceInfoRecords')
export class ComplianceInfoRecords {
  @JsonProperty('total', AminoNumberConverter)
  total: number = undefined;

  @JsonProperty('items', [ComplianceInfo])
  items: ComplianceInfo[] = undefined;
}
