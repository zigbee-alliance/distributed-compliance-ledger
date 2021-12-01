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

import { DashboardRoutingModule } from './dashboard-routing.module';
import {ModelInfoModule} from '../model-info/model-info.module';
import {BlockModule} from '../block/block.module';
import {ModelTestingResultModule} from '../model-testing-results/model-testing-results.module';
import {ModelComplianceModule} from '../model-compliance/model-compliance.module';

import { AccountModule } from '../account/account.module';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { KeyModule } from '../key/key.module';
import { TxModule } from '../tx/tx.module';
import { PkiModule } from '../pki/pki.module';
import { ValidatorModule } from '../validator/validator.module';


@NgModule({
    declarations: [
        DashboardComponent
    ],
    exports: [
        DashboardComponent
    ],
    imports: [
        CommonModule,
        DashboardRoutingModule,
        AccountModule,
        ModelInfoModule,
        BlockModule,
        KeyModule,
        TxModule,
        ModelTestingResultModule,
        ModelComplianceModule,
        PkiModule,
        ValidatorModule
    ]
})
export class DashboardModule { }
