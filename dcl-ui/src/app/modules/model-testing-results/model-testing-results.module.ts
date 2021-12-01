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
