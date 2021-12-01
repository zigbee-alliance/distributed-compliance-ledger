import { JsonConvert, JsonConverter, JsonCustomConvert } from 'json2typescript';
import { Message } from '../models/tx/message';
import { MessageAddModelInfo } from '../models/model-info/message-add-model-info';
import { MessageType } from '../models/tx/message-type';
import { MessageUpdateModelInfo } from '../models/model-info/message-update-model-info';
import { MessageDeleteModelInfo } from '../models/model-info/message-delete-model-info';
import { MsgAddTestingResult } from '../models/model-testing-results/message-add-testing-result';
import { MsgCertifyModel } from '../models/model-compliance/message-certify-model';
import { MsgRevokeModel } from '../models/model-compliance/message-revoke-model';
import { MessageApproveAddX509RootCert } from '../models/pki/message-approve-add-x509-root-certificate';
import { MessageAddX509Cert } from '../models/pki/message-add-x509-certificate';
import { MessageProposeAddX509RootCert } from '../models/pki/message-propose-x509-root-certificate';
import { StdTxnValue } from '../models/tx/std-txn-value';

@JsonConverter
export class MessageConverter implements JsonCustomConvert<Message> {

  jsonConvert = new JsonConvert();

  serialize(message: Message): any {
    const result: any = {};

    switch (message.codecType) {
      case MessageType.AddModelInfo:
        result.type = MessageType.AddModelInfo.toString();
        result.value = this.jsonConvert.serialize(message as MessageAddModelInfo);
        break;
      case MessageType.UpdateModelInfo:
        result.type = MessageType.UpdateModelInfo.toString();
        result.value = this.jsonConvert.serialize(message as MessageUpdateModelInfo);
        break;
      case MessageType.DeleteModelInfo:
        result.type = MessageType.DeleteModelInfo.toString();
        result.value = this.jsonConvert.serialize(message as MessageDeleteModelInfo);
        break;
      case MessageType.AddTestingResult:
        result.type = MessageType.AddTestingResult.toString();
        result.value = this.jsonConvert.serialize(message as MsgAddTestingResult);
        break;
      case MessageType.CertifyModel:
        result.type = MessageType.CertifyModel.toString();
        result.value = this.jsonConvert.serialize(message as MsgCertifyModel);
        break;
      case MessageType.RevokeModel:
        result.type = MessageType.RevokeModel.toString();
        result.value = this.jsonConvert.serialize(message as MsgRevokeModel);
        break;
      case MessageType.ProposeAddX509RootCert:
        result.type = MessageType.ProposeAddX509RootCert.toString();
        result.value = this.jsonConvert.serialize(message as MessageProposeAddX509RootCert);
        break;
      case MessageType.ApproveAddX509RootCert:
        result.type = MessageType.ApproveAddX509RootCert.toString();
        result.value = this.jsonConvert.serialize(message as MessageApproveAddX509RootCert);
        break;
      case MessageType.AddX509Cert:
        result.type = MessageType.AddX509Cert.toString();
        result.value = this.jsonConvert.serialize(message as MessageAddX509Cert);
        break;
      default:
        throw new Error('Unknown message type');
    }

    return result;
  }

  deserialize(message: any): Message {
    let result: Message;

    switch (message.type) {
      case MessageType.AddModelInfo.toString():
        result = this.jsonConvert.deserializeObject(message.value, MessageAddModelInfo);
        break;
      case MessageType.UpdateModelInfo.toString():
        result = this.jsonConvert.deserializeObject(message.value, MessageUpdateModelInfo);
        break;
      case MessageType.DeleteModelInfo.toString():
        result = this.jsonConvert.deserializeObject(message.value, MessageDeleteModelInfo);
        break;
      case MessageType.AddTestingResult.toString():
        result = this.jsonConvert.deserializeObject(message.value, MsgAddTestingResult);
        break;
      case MessageType.CertifyModel.toString():
        result = this.jsonConvert.deserializeObject(message.value, MsgCertifyModel);
        break;
      case MessageType.RevokeModel.toString():
        result = this.jsonConvert.deserializeObject(message.value, MsgRevokeModel);
        break;
      case MessageType.ProposeAddX509RootCert.toString():
        result = this.jsonConvert.deserializeObject(message.value, MessageProposeAddX509RootCert);
        break;
      case MessageType.ApproveAddX509RootCert.toString():
        result = this.jsonConvert.deserializeObject(message.value, MessageApproveAddX509RootCert);
        break;
      case MessageType.AddX509Cert.toString():
        result = this.jsonConvert.deserializeObject(message.value, MessageAddX509Cert);
        break;
      default:
        throw new Error('Unknown message type');
    }

    return result;
  }
}
