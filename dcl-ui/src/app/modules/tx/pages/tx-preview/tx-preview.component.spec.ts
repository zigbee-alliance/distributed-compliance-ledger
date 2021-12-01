import { async, ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { TxPreviewComponent } from './tx-preview.component';

describe('TxPreviewComponent', () => {
  let component: TxPreviewComponent;
  let fixture: ComponentFixture<TxPreviewComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ TxPreviewComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TxPreviewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
