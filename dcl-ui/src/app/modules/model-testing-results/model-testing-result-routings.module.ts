import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TestingResultAddComponent } from './pages/testing-result-add/testing-result-add.component';

const routes: Routes = [
  { path: 'model-testing-results/:vid/:pid', component: TestingResultAddComponent },
];

@NgModule({
  imports: [
    RouterModule.forChild(routes)
  ],
  exports: [RouterModule]
})
export class ModelTestingResultRoutingModule {
}
