import { JsonObject, JsonProperty } from 'json2typescript';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';

@JsonObject('MessageDeleteModelInfo')
export class MessageDeleteModelInfo extends Message {

  @JsonProperty('vid', Number)
  vid = 0;

  @JsonProperty('pid', Number)
  pid = 0;

  @JsonProperty('signer', String)
  signer = '';

  constructor(init?: Partial<MessageDeleteModelInfo>) {
    super(MessageType.DeleteModelInfo);

    Object.assign(this, init);
  }

  setSigner(signer: string) {
    this.signer = signer;
  }
}
