import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';

@JsonObject('ABCIMessageLog')
export class ABCIMessageLog {
  @JsonProperty('msg_index', AminoNumberConverter)
  msgIndex: number = undefined;

  @JsonProperty('success', Boolean)
  success: boolean = undefined;

  @JsonProperty('log', String)
  log: string = undefined;
}
