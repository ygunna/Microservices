import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MicroGuideComponent } from './micro-guide.component';

describe('MicroGuideComponent', () => {
  let component: MicroGuideComponent;
  let fixture: ComponentFixture<MicroGuideComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MicroGuideComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MicroGuideComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
