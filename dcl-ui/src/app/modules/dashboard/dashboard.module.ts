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
