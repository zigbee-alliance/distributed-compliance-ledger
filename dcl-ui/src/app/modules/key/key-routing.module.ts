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
import { KeyDetailsComponent } from './pages/key-details/key-details.component';
import { KeyListComponent } from './pages/key-list/key-list.component';


const routes: Routes = [
  { path: 'key', component: KeyListComponent },
  { path: 'key/:name', component: KeyDetailsComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class KeyRoutingModule {
}
