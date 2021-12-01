import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ValidatorListComponent } from './pages/validator-list/validator-list.component';

const routes: Routes = [
  { path: 'validators', component: ValidatorListComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class ValidatorRoutingModule {
}
