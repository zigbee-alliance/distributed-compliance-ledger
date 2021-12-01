import { TestBed } from '@angular/core/testing';

import { TxService } from './tx.service';

describe('TxService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: TxService = TestBed.inject(TxService);
    expect(service).toBeTruthy();
  });
});
