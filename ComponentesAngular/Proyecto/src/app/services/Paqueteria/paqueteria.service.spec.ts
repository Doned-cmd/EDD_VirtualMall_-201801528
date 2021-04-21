import { TestBed } from '@angular/core/testing';

import { PaqueteriaService } from './paqueteria.service';

describe('PaqueteriaService', () => {
  let service: PaqueteriaService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PaqueteriaService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
