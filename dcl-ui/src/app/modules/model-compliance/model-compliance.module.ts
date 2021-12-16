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
