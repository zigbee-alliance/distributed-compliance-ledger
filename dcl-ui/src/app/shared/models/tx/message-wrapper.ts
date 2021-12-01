import { JsonObject, JsonProperty } from 'json2typescript';
import { Message } from './message';
import { MessageConverter } from '../../json-converters/message-converter';

// Used to pass message object between pages. Otherwise it's impossible to use MessageConverter.
@JsonObject('MessageWrapper')
export class MessageWrapper {
  @JsonProperty('message', MessageConverter)
  message: Message = undefined;

  constructor(message?: Message) {
    this.message = message;
  }
}
