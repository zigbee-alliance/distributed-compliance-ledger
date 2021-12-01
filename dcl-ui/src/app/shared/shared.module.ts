import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { KeyInfoPipe } from './pipes/key-info.pipe';
import { AccountPipe } from './pipes/account.pipe';



@NgModule({
  declarations: [
    KeyInfoPipe,
    AccountPipe
  ],
  exports: [
    KeyInfoPipe,
    AccountPipe
  ],
  imports: [
    CommonModule
  ]
})
export class SharedModule { }
