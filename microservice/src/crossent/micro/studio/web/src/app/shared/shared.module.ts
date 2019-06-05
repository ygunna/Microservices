import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { ApiService } from './api.service';
import { LoaderService } from './loader.service';
import { WindowService } from './window.service';
import { AuthService } from './auth.service';
import { AuthGuard } from './auth.guard';
import { RoleGuard } from './role.guard';

@NgModule({
  imports: [
    CommonModule,
    HttpClientModule,
  ],
  declarations: [],
  providers: [
    ApiService, LoaderService, WindowService, AuthService, AuthGuard, RoleGuard
  ],
})
export class SharedModule { }
