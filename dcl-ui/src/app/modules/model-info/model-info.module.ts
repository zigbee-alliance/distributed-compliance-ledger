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

import { ModelInfoRoutingModule } from './model-info-routing.module';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { ModelInfoListComponent } from './pages/model-info-list/model-info-list.component';
import { ModelInfoDetailsComponent } from './pages/model-info-details/model-info-details.component';
import { ModelInfoEditComponent } from './pages/model-info-edit/model-info-edit.component';
import { ModelTestingResultModule } from '../model-testing-results/model-testing-results.module';
import { FormsModule } from '@angular/forms';
import { SharedModule } from '../../shared/shared.module';
import { ModelInfoSearchComponent } from './components/certificates-search/model-info-search.component';


@NgModule({
  declarations: [
    ModelInfoListComponent,
    ModelInfoDetailsComponent,
    ModelInfoEditComponent,
    ModelInfoSearchComponent
  ],
  exports: [
    ModelInfoListComponent
  ],
  imports: [
    CommonModule,
    ModelInfoRoutingModule,
    NgbModule,
    FormsModule,
    SharedModule,
    ModelTestingResultModule
  ],
  entryComponents: [
    ModelInfoSearchComponent
  ]
})
export class ModelInfoModule {
}
