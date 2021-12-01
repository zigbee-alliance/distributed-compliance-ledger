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
