import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ModelInfoDetailsComponent } from '../model-info/pages/model-info-details/model-info-details.component';
import { CertifiedModelAddComponent } from './pages/certified-model-add/certified-model-add.component';
import { ComplianceInfoListComponent } from './pages/compliance-info-list/compliance-info-list.component';
import { RevokedModelAddComponent } from './pages/revoked-model-add/revoked-model-add.component';

const routes: Routes = [
  { path: 'compliance-info', component: ComplianceInfoListComponent },
  { path: 'compliance-info/certify', component: CertifiedModelAddComponent },
  { path: 'compliance-info/revoke', component: RevokedModelAddComponent },
  { path: 'compliance-info/:vid/:pid/revoke', component: RevokedModelAddComponent },
  { path: 'compliance-info/:vid/:pid/certify', component: CertifiedModelAddComponent },
  { path: 'compliance-info/:vid/:pid', component: ModelInfoDetailsComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class ModelComplianceRoutingModule {
}
