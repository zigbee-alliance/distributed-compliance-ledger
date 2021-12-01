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

import { Inject, Injectable } from '@angular/core';
import { HttpExtensionsService } from '../../core/services/http-extensions.service';
import { Observable } from 'rxjs';
import { DecodeTxsRequest } from '../../shared/models/tx/decode-txs-request';
import { DecodeTxsResponse } from '../../shared/models/tx/decode-txs-response';
import { Message } from '../../shared/models/tx/message';
import { Router } from '@angular/router';
import { JsonConvert } from 'json2typescript';
import { MessageWrapper } from '../../shared/models/tx/message-wrapper';
import { StdTxnValue } from '../../shared/models/tx/std-txn-value';
import { HttpHeaders } from '@angular/common/http';
import { StdFee } from '../../shared/models/tx/std-fee';
import { map } from 'rxjs/operators';
import { TxResponse } from '../../shared/models/tx/tx-response';
import { StdSignMsg } from 'src/app/shared/models/tx/sign-request';
import { BaseReq } from 'src/app/shared/models/tx/base-request';
import { StdTxn } from 'src/app/shared/models/tx/std-txn';
import { Account } from '../../shared/models/account/acount';
import { KeyInfo } from '../../shared/models/key/key-info';

@Injectable({
  providedIn: 'root'
})
export class TxService {

  private gasAmount = 200000;
  private passphrase = 'test1234';

  private txBaseUrl = 'tx';

  private jsonConvert: JsonConvert = new JsonConvert();

  constructor(private http: HttpExtensionsService,
              private router: Router,
              @Inject('CHAIN_ID') private chainId: string) {
  }

  goPreview(message: Message, backRoute: string): Promise<boolean> {
    return this.router.navigate(['tx', 'preview', { message: this.encodeMessage(message), back: backRoute }]);
  }

  encodeMessage(message: Message) {
    const wrapper = new MessageWrapper(message);
    const json = this.jsonConvert.serialize(wrapper);
    return JSON.stringify(json);
  }

  decodeMessage(str: string): Message {
    const json = JSON.parse(str);
    const wrapper = this.jsonConvert.deserializeObject(json, MessageWrapper);
    return wrapper.message;
  }

  makeSignMsg(message: Message, signer: Account): StdSignMsg {
    return new StdSignMsg({
      baseReq: new BaseReq({
        chainId: this.chainId,
        accountNumber: signer.accountNumber,
        sequence: signer.sequence,
        from: signer.address
      }),
      txn: new StdTxn({
        value: new StdTxnValue({
          messages: [message],
          fee: new StdFee({
            gas: this.gasAmount
          })
        })
      })
    });
  }

  decodeTx(req: DecodeTxsRequest): Observable<DecodeTxsResponse> {
    return this.http.post(`${this.txBaseUrl}/decode`, req, DecodeTxsResponse);
  }

  signTx(signMsg: StdSignMsg, keyInfo: KeyInfo): Observable<any> {
    const httpHeaders = new HttpHeaders()
      .append('Authorization', 'Basic ' + btoa(`${keyInfo.name}:${this.passphrase}`));

    const encodedMsg = this.jsonConvert.serializeObject(signMsg);
    return this.http.postAny(`${this.txBaseUrl}/sign`, encodedMsg, undefined, httpHeaders);
  }

  broadcastTx(tx: any): Observable<TxResponse> {
    return this.http.postAny(`${this.txBaseUrl}/broadcast`, tx).pipe(
      map(resp => this.jsonConvert.deserializeObject(resp, TxResponse))
    );
  }
}
