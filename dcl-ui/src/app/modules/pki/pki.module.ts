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

import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PkiRoutingModule } from './pki-routing.module';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { FormsModule } from '@angular/forms';
import { SharedModule } from '../../shared/shared.module';
import { AddCertificateComponent } from './pages/add-certificate/add-certificate.component';
import { ApprovedCertificateDetailsComponent } from './pages/approved-certificate-details/approved-certificate-details.component';
import { ApprovedCertificatesListComponent } from './pages/approved-certificates-list/approved-certificates-list.component';
import { ProposedCertificateDetailsComponent } from './pages/proposed-certificate-details/proposed-certificate-details.component';
import { ProposedCertificatesListComponent } from './pages/proposed-certificates-list/proposed-certificates-list.component';
import { ProposeCertificateComponent } from './pages/propose-certificate/propose-certificate.component';
import { CertificatesSearchComponent } from './components/certificates-search/certificates-search.component';

@NgModule({
  declarations: [
    AddCertificateComponent,
    ApprovedCertificateDetailsComponent,
    ApprovedCertificatesListComponent,
    ProposeCertificateComponent,
    ProposedCertificateDetailsComponent,
    ProposedCertificatesListComponent,
    CertificatesSearchComponent
  ],
  exports: [
    AddCertificateComponent,
    ApprovedCertificateDetailsComponent,
    ApprovedCertificatesListComponent,
    ProposeCertificateComponent,
    ProposedCertificateDetailsComponent,
    ProposedCertificatesListComponent
  ],
  imports: [
    CommonModule,
    PkiRoutingModule,
    NgbModule,
    FormsModule,
    SharedModule,
  ],
  entryComponents: [
    CertificatesSearchComponent
  ]
})
export class PkiModule {
}
