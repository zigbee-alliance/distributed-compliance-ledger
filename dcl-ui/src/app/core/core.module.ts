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

import { CoreRoutingModule } from './core-routing.module';

import { PageNotFoundComponent } from './page-not-found/page-not-found.component';
import { HeaderComponent } from './header/header.component';
import { ToastsComponent } from './toasts/toasts.component';
import { ModalWindowComponent } from './modal/modal.component';
import { NgbToastModule } from '@ng-bootstrap/ng-bootstrap';


@NgModule({
  declarations: [
    PageNotFoundComponent,
    HeaderComponent,
    ToastsComponent,
    ModalWindowComponent
  ],
  exports: [
    HeaderComponent,
    ToastsComponent,
    ModalWindowComponent
  ],
  imports: [
    CommonModule,
    CoreRoutingModule,
    NgbToastModule
  ],
  entryComponents: [
    ModalWindowComponent
  ]
})
export class CoreModule {
}
