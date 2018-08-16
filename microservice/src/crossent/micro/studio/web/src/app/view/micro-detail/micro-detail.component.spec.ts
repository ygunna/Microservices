import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MicroDetailComponent } from './micro-detail.component';

describe('MicroDetailComponent', () => {
  let component: MicroDetailComponent;
  let fixture: ComponentFixture<MicroDetailComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MicroDetailComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MicroDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
