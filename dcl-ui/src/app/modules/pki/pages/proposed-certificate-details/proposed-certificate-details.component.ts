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
import { ProposedCertificate } from '../../../../shared/models/pki/proposed-certificate';
import { PkiService } from '../../pki.service';
import { ActivatedRoute } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ModalWindowComponent } from '../../../../core/modal/modal.component';

@Component({
  selector: 'app-proposed-certificate-details',
  templateUrl: './proposed-certificate-details.component.html',
  styleUrls: ['./proposed-certificate-details.component.scss']
})
export class ProposedCertificateDetailsComponent implements OnInit {

  item$: Observable<ProposedCertificate>;

  constructor(public pkiService: PkiService,
              private route: ActivatedRoute,
              private modalService: NgbModal) {
  }

  ngOnInit() {
    this.item$ = this.route.paramMap.pipe(
      switchMap(params => this.pkiService.getProposedX509RootCert(params.get('subject'), params.get('subjectKeyId')))
    );
  }

  viewCertificate(pemCert: string) {
    const modalRef = this.modalService.open(ModalWindowComponent, { size: 'lg' });
    modalRef.componentInstance.content = pemCert;
    modalRef.componentInstance.header = 'Proposed Root Certificate';
  }
}
