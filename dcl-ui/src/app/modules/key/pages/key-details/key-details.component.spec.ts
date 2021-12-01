import { async, ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { KeyDetailsComponent } from './key-details.component';

describe('KeyDetailsComponent', () => {
  let component: KeyDetailsComponent;
  let fixture: ComponentFixture<KeyDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ KeyDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KeyDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
