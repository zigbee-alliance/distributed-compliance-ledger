import { JsonObject, JsonProperty } from 'json2typescript';
import { PemCertificate } from './pem-certificate';
import { Message } from '../tx/message';
import { MessageType } from '../tx/message-type';

@JsonObject('MessageAddX509Cert')
export class MessageAddX509Cert extends Message {

  @JsonProperty('cert', String)
  cert = '';

  @JsonProperty('signer', String)
  signer = '';

  constructor(cert?: PemCertificate, signer?: string) {
    super(MessageType.AddX509Cert);

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
