import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ValidatorRoutingModule } from './validator-routing.module';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { ValidatorListComponent } from './pages/validator-list/validator-list.component';
import { FormsModule } from '@angular/forms';
import { SharedModule } from '../../shared/shared.module';

@NgModule({
  declarations: [
    ValidatorListComponent,
  ],
  exports: [
    ValidatorListComponent,
  ],
  imports: [
    CommonModule,
    ValidatorRoutingModule,
    NgbModule,
    FormsModule,
    SharedModule
  ]
})
export class ValidatorModule {
}
