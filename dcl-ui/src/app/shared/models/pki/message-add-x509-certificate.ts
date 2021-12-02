/**
 * Copyright 2020 DSR Corporation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
