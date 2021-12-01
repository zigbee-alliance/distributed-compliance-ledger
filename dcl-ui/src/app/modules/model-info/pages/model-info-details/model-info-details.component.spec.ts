import { async, ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ModelInfoDetailsComponent } from './model-info-details.component';

describe('ModelInfoDetailsComponent', () => {
  let component: ModelInfoDetailsComponent;
  let fixture: ComponentFixture<ModelInfoDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ ModelInfoDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ModelInfoDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
