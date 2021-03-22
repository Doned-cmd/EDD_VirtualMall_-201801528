import { TestBed } from '@angular/core/testing';

import { GettiendasService } from './gettiendas.service';

describe('GettiendasService', () => {
  let service: GettiendasService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GettiendasService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
