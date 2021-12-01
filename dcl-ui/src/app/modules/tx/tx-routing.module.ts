import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TxPreviewComponent } from './pages/tx-preview/tx-preview.component';


const routes: Routes = [
  { path: 'tx/preview', component: TxPreviewComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class TxRoutingModule {
}
