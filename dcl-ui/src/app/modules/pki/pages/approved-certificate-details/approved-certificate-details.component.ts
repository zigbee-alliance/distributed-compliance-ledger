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
import { Observable, from } from 'rxjs';
import { ProposedCertificate } from '../../../../shared/models/pki/proposed-certificate';
import { PkiService } from '../../pki.service';
import { ActivatedRoute } from '@angular/router';
import { switchMap, map, pluck } from 'rxjs/operators';
import { ApprovedCertificate } from 'src/app/shared/models/pki/approved-certificate';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ModalWindowComponent } from '../../../../core/modal/modal.component';

@Component({
  selector: 'app-approved-certificate-details',
  templateUrl: './approved-certificate-details.component.html',
  styleUrls: ['./approved-certificate-details.component.scss']
})
export class ApprovedCertificateDetailsComponent implements OnInit {

  items$: Observable<ApprovedCertificate[]>;

  constructor(public pkiService: PkiService,
              private route: ActivatedRoute,
              private modalService: NgbModal) {
  }

  ngOnInit() {
    const source = this.route.paramMap.pipe(
      switchMap(params =>
        this.pkiService.getX509Cert(params.get('subject'), params.get('subjectKeyId'))
      )
    );

    this.items$ = source.pipe(pluck('items'));
  }

  viewCertificate(pemCert: string) {
    const modalRef = this.modalService.open(ModalWindowComponent, { size: 'lg' });
    modalRef.componentInstance.content = pemCert;
    modalRef.componentInstance.header = 'Certificate';
  }
}
