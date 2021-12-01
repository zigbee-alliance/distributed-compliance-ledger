import { async, ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { KeyListComponent } from './key-list.component';

describe('KeyListComponent', () => {
  let component: KeyListComponent;
  let fixture: ComponentFixture<KeyListComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ KeyListComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(KeyListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
