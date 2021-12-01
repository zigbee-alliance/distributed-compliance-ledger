import { JsonObject, JsonProperty } from 'json2typescript';
import { PemCertificate } from './pem-certificate';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';

@JsonObject('MessageProposeAddX509RootCert')
export class MessageProposeAddX509RootCert extends Message {

  @JsonProperty('cert', String)
  cert = '';

  @JsonProperty('signer', String)
  signer = '';

  constructor(cert?: PemCertificate, signer?: string) {
    super(MessageType.ProposeAddX509RootCert);

    if (cert) {
      this.cert = cert.cert;
    }

    if (signer) {
      this.signer = signer;
    }
  }

  setSigner(signer: string) {
    this.signer = signer;
  }
}
