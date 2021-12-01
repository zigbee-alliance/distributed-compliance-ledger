import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { BlockListComponent } from './pages/block-list/block-list.component';
import { BlockDetailsComponent } from './pages/block-details/block-details.component';


const routes: Routes = [
  { path: 'block', component: BlockListComponent },
  { path: 'block/:height', component: BlockDetailsComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class BlockRoutingModule {
}
