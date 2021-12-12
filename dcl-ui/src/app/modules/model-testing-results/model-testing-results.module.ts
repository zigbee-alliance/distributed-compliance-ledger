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

import { ModelTestingResultRoutingModule } from './model-testing-result-routings.module';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { TestingResultAddComponent } from './pages/testing-result-add/testing-result-add.component';
import { TestingResultListComponent } from './pages/testing-result-list/testing-result-list.component';
import { FormsModule } from '@angular/forms';
import { SharedModule } from '../../shared/shared.module';


@NgModule({
  declarations: [
    TestingResultAddComponent,
    TestingResultListComponent
  ],
  exports: [
    TestingResultAddComponent,
    TestingResultListComponent
  ],
  imports: [
    CommonModule,
    ModelTestingResultRoutingModule,
    NgbModule,
    FormsModule,
    SharedModule
  ]
})
export class ModelTestingResultModule {
}
