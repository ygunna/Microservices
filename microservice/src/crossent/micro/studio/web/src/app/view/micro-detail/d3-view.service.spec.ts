import { TestBed, inject } from '@angular/core/testing';

import { D3ViewService } from './d3-view.service';

describe('D3ViewService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [D3ViewService]
    });
  });

  it('should be created', inject([D3ViewService], (service: D3ViewService) => {
    expect(service).toBeTruthy();
  }));
});
