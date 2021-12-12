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
