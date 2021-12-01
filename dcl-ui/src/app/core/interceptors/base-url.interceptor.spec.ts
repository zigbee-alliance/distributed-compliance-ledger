import { TestBed } from '@angular/core/testing';

import { BaseUrlInterceptor } from './base-url.interceptor';

describe('BaseUrlInterceptorService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: BaseUrlInterceptor = TestBed.inject(BaseUrlInterceptor);
    expect(service).toBeTruthy();
  });
});
