import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AccountRoutingModule } from './account-routing.module';
import { AccountListComponent } from './pages/account-list/account-list.component';
import { AccountDetailsComponent } from './pages/account-details/account-details.component';
import {AccountRolesComponent} from './components/account-roles/account-roles.component';


@NgModule({
  declarations: [
    AccountListComponent,
    AccountDetailsComponent,
    AccountRolesComponent
  ],
  imports: [
    CommonModule,
    AccountRoutingModule
  ]
})
export class AccountModule {
}
