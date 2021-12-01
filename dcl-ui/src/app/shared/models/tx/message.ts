import { MessageType } from './message-type';

export abstract class Message {

  protected constructor(public codecType: MessageType) {
  }

  abstract setSigner(signer: string);
}
