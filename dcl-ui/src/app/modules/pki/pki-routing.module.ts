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
