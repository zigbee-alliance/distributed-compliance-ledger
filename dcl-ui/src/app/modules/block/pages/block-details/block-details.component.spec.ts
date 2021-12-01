import { async, ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';

import { BlockDetailsComponent } from './block-details.component';

describe('BlockDetailsComponent', () => {
  let component: BlockDetailsComponent;
  let fixture: ComponentFixture<BlockDetailsComponent>;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      declarations: [ BlockDetailsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(BlockDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
