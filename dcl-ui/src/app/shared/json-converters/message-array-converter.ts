import { JsonConverter, JsonCustomConvert } from 'json2typescript';
import { Message } from '../models/tx/message';
import { MessageConverter } from './message-converter';

// Library is strange...
@JsonConverter
export class MessageArrayConverter implements JsonCustomConvert<Message[]> {

  messageConverter: MessageConverter = new MessageConverter();

  serialize(messages: Message[]): any {
    // Js is beautiful...
    return messages.map(i => this.messageConverter.serialize(i));
  }

  deserialize(messages: any[]): Message[] {
    return messages.map(i => this.messageConverter.deserialize(i));
  }
}
