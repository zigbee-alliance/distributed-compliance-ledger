import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { ModelInfoListComponent } from './model-info-list.component';

describe('ModelInfoListComponent', () => {
  let component: ModelInfoListComponent;
  let fixture: ComponentFixture<ModelInfoListComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ ModelInfoListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ModelInfoListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
