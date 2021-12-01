import { TestBed } from '@angular/core/testing';

import { KeyService } from './key.service';

describe('KeyService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: KeyService = TestBed.inject(KeyService);
    expect(service).toBeTruthy();
  });
});
