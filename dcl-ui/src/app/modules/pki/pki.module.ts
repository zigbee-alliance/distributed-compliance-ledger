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
