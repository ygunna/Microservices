import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ApiavatarComponent } from './apiavatar.component';

describe('ApiavatarComponent', () => {
  let component: ApiavatarComponent;
  let fixture: ComponentFixture<ApiavatarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ApiavatarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ApiavatarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
