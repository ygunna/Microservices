import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MicroApiComponent } from './micro-api.component';

describe('MicroApiComponent', () => {
  let component: MicroApiComponent;
  let fixture: ComponentFixture<MicroApiComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MicroApiComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MicroApiComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
