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
import { RouterModule, Routes } from '@angular/router';
import { ProposedCertificatesListComponent } from './pages/proposed-certificates-list/proposed-certificates-list.component';
import { ProposedCertificateDetailsComponent } from './pages/proposed-certificate-details/proposed-certificate-details.component';
import { ApprovedCertificatesListComponent } from './pages/approved-certificates-list/approved-certificates-list.component';
import { ApprovedCertificateDetailsComponent } from './pages/approved-certificate-details/approved-certificate-details.component';
import { ProposeCertificateComponent } from './pages/propose-certificate/propose-certificate.component';
import { AddCertificateComponent } from './pages/add-certificate/add-certificate.component';

const routes: Routes = [
  { path: 'proposed-certificates', component: ProposedCertificatesListComponent },
  { path: 'proposed-certificates/propose', component: ProposeCertificateComponent },
  { path: 'proposed-certificates/:subject/:subjectKeyId', component: ProposedCertificateDetailsComponent },
  { path: 'certificates', component: ApprovedCertificatesListComponent },
  { path: 'certificates/propose', component: ProposeCertificateComponent },
  { path: 'certificates/add', component: AddCertificateComponent },
  { path: 'certificates/:subject/:subjectKeyId', component: ApprovedCertificateDetailsComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class PkiRoutingModule {
}
