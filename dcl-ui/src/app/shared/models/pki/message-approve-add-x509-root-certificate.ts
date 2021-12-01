import { JsonObject, JsonProperty } from 'json2typescript';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';

@JsonObject('MessageApproveAddX509RootCert')
export class MessageApproveAddX509RootCert extends Message {

  @JsonProperty('subject', String)
  subject = '';

  @JsonProperty('subject_key_id', String)
  subjectKeyId = '';

  @JsonProperty('signer', String)
  signer = '';

  constructor(subject?: string, subjectKeyId?: string, signer?: string) {
    super(MessageType.ApproveAddX509RootCert);

    if (subject) {
      this.subject = subject;
    }

    if (subjectKeyId) {
      this.subjectKeyId = subjectKeyId;
    }

    if (signer) {
      this.signer = signer;
    }
  }

  setSigner(signer: string) {
    this.signer = signer;
  }
}
