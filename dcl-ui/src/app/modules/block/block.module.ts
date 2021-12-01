import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BlockRoutingModule } from './block-routing.module';

import { BlockListComponent } from './pages/block-list/block-list.component';
import { BlockDetailsComponent } from './pages/block-details/block-details.component';
import { NgbPaginationModule } from '@ng-bootstrap/ng-bootstrap';


@NgModule({
  declarations: [
    BlockListComponent,
    BlockDetailsComponent
  ],
  imports: [
    CommonModule,
    BlockRoutingModule,
    NgbPaginationModule
  ]
})
export class BlockModule {
}
