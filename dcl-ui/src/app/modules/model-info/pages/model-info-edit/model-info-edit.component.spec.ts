import { async, ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ModelInfoEditComponent } from './model-info-edit.component';

describe('ModelInfoEditComponent', () => {
  let component: ModelInfoEditComponent;
  let fixture: ComponentFixture<ModelInfoEditComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ ModelInfoEditComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ModelInfoEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
