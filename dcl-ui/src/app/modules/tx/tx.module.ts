import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TxRoutingModule } from './tx-routing.module';
import { TxPreviewComponent } from './pages/tx-preview/tx-preview.component';
import { FormsModule } from '@angular/forms';
import { SharedModule } from '../../shared/shared.module';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';


@NgModule({
  declarations: [
    TxPreviewComponent
  ],
  imports: [
    CommonModule,
    TxRoutingModule,
    FormsModule,
    SharedModule,
    NgbModule
  ]
})
export class TxModule {
}
