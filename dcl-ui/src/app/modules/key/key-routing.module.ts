import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { KeyDetailsComponent } from './pages/key-details/key-details.component';
import { KeyListComponent } from './pages/key-list/key-list.component';


const routes: Routes = [
  { path: 'key', component: KeyListComponent },
  { path: 'key/:name', component: KeyDetailsComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class KeyRoutingModule {
}
