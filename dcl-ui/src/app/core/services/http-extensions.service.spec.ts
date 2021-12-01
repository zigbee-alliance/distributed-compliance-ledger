import { TestBed } from '@angular/core/testing';

import { HttpExtensionsService } from './http-extensions.service';

describe('HttpExtensionsService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: HttpExtensionsService = TestBed.inject(HttpExtensionsService);
    expect(service).toBeTruthy();
  });
});
