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
import { ModelInfoListComponent } from './pages/model-info-list/model-info-list.component';
import { ModelInfoEditComponent } from './pages/model-info-edit/model-info-edit.component';
import { ModelInfoDetailsComponent } from './pages/model-info-details/model-info-details.component';
import { TestingResultAddComponent } from '../model-testing-results/pages/testing-result-add/testing-result-add.component';
import { CertifiedModelAddComponent } from '../model-compliance/pages/certified-model-add/certified-model-add.component';
import { RevokedModelAddComponent } from '../model-compliance/pages/revoked-model-add/revoked-model-add.component';


const routes: Routes = [
  { path: 'model-info', component: ModelInfoListComponent },
  { path: 'model-info/new', component: ModelInfoEditComponent },
  { path: 'model-info/:vid/:pid', component: ModelInfoDetailsComponent },
  { path: 'model-info/:vid/:pid/update', component: ModelInfoEditComponent },
  { path: 'model-info/:vid/:pid/add-testing-result', component: TestingResultAddComponent },
  { path: 'model-info/:vid/:pid/certify', component: CertifiedModelAddComponent },
  { path: 'model-info/:vid/:pid/revoke', component: RevokedModelAddComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class ModelInfoRoutingModule {
}
