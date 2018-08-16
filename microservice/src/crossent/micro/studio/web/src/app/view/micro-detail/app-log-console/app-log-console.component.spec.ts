import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AppLogConsoleComponent } from './app-log-console.component';

describe('AppLogConsoleComponent', () => {
  let component: AppLogConsoleComponent;
  let fixture: ComponentFixture<AppLogConsoleComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AppLogConsoleComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AppLogConsoleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
