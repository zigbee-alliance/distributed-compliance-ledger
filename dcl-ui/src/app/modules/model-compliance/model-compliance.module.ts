import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ModelComplianceRoutingModule } from './model-compliance-routing.module';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { CertifiedModelAddComponent } from './pages/certified-model-add/certified-model-add.component';
import { RevokedModelAddComponent } from './pages/revoked-model-add/revoked-model-add.component';
import { ComplianceInfoListComponent } from './pages/compliance-info-list/compliance-info-list.component';
import { FormsModule } from '@angular/forms';
import { SharedModule } from '../../shared/shared.module';

@NgModule({
  declarations: [
    CertifiedModelAddComponent,
    ComplianceInfoListComponent,
    RevokedModelAddComponent,
  ],
  exports: [
    CertifiedModelAddComponent,
    ComplianceInfoListComponent,
    RevokedModelAddComponent,
  ],
  imports: [
    CommonModule,
    ModelComplianceRoutingModule,
    NgbModule,
    FormsModule,
    SharedModule
  ]
})
export class ModelComplianceModule {
}
