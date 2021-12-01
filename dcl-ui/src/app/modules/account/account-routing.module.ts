import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccountListComponent } from './pages/account-list/account-list.component';
import { AccountDetailsComponent } from './pages/account-details/account-details.component';


const routes: Routes = [
  { path: 'account', component: AccountListComponent },
  { path: 'account/:addr', component: AccountDetailsComponent }
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class AccountRoutingModule {
}
