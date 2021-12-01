import { JsonObject, JsonProperty } from 'json2typescript';
import { AminoNumberConverter } from '../../json-converters/amino-number-converter';
import { ABCIMessageLog } from './abci-message-log';

@JsonObject('TxResponse')
export class TxResponse {
  @JsonProperty('height', AminoNumberConverter)
  height: number = undefined;

  @JsonProperty('txhash', String)
  txHash: string = undefined;

  @JsonProperty('code', AminoNumberConverter, true)
  code: number = undefined;

  @JsonProperty('codespace', String, true)
  codespace: string = undefined;

  @JsonProperty('logs', [ABCIMessageLog], true)
  logs: ABCIMessageLog[] = undefined;

  @JsonProperty('raw_log', String, true)
  rawLog: string = undefined;
}
