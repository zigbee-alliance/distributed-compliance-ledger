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

import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { NgForm, NgModel } from '@angular/forms';
import { TxService } from '../../../tx/tx.service';
import { PemCertificate } from 'src/app/shared/models/pki/pem-certificate';
import { MessageProposeAddX509RootCert } from 'src/app/shared/models/pki/message-propose-x509-root-certificate';

@Component({
  selector: 'app-propose-certificate',
  templateUrl: './propose-certificate.component.html',
  styleUrls: ['./propose-certificate.component.scss']
})
export class ProposeCertificateComponent implements OnInit {

  item: PemCertificate = new PemCertificate();

  showValidation = false;

  constructor(private txService: TxService) {}

  ngOnInit() {}

  onSubmit(form: NgForm, pemCertificate: PemCertificate) {
    if (!form.valid) {
      this.showValidation = true;
      return;
    }

    const message = new MessageProposeAddX509RootCert(pemCertificate);
    this.txService.goPreview(message, '/proposed-certificates');
  }

  getValidityClasses(model: NgModel) {
    return {
      'is-valid': (model.touched || this.showValidation) && model.valid,
      'is-invalid': (model.touched || this.showValidation) && model.invalid
    };
  }
}
