import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { KeyRoutingModule } from './key-routing.module';
import { KeyListComponent } from './pages/key-list/key-list.component';
import { KeyDetailsComponent } from './pages/key-details/key-details.component';


@NgModule({
  declarations: [
    KeyListComponent,
    KeyDetailsComponent,
  ],
  imports: [
    CommonModule,
    KeyRoutingModule
  ]
})
export class KeyModule {
}
