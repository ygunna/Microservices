import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiHealthComponent } from './api-health.component';

describe('ApiHealthComponent', () => {
  let component: ApiHealthComponent;
  let fixture: ComponentFixture<ApiHealthComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ApiHealthComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiHealthComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
