import {Injectable, EventEmitter, Output} from '@angular/core';
import { ApiService } from './api.service';

@Injectable()
export class AuthService {

  constructor(private apiService: ApiService) { }

  public isAuthenticated(): boolean {
    const token = localStorage.getItem('username');
    if (token) {
      return true;
    } else {
      return false;
    }
  }

  public getAuthRole(auth: string): boolean {
    const token = localStorage.getItem('auth');
    const local = atob(token);

    if (auth == local) {
      return true;
    } else {
      return false;
    }
  }

}
