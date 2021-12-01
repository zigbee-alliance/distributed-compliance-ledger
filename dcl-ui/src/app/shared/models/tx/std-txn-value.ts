import { JsonObject, JsonProperty } from 'json2typescript';
import { StdFee } from './std-fee';
import { MessageArrayConverter } from '../../json-converters/message-array-converter';
import { Message } from './message';

@JsonObject('StdTxnValue')
export class StdTxnValue {

  @JsonProperty('fee', StdFee)
  fee: StdFee = new StdFee();

  @JsonProperty('msg', MessageArrayConverter)
  messages: Message[] = [];

  @JsonProperty('memo', String)
  memo = '';

  constructor(init?: Partial<StdTxnValue>) {
    Object.assign(this, init);
  }
}
