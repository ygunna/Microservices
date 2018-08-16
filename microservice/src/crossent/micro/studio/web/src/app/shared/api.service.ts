import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpErrorResponse, HttpResponse } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs/Observable';
import { ErrorObservable } from 'rxjs/observable/ErrorObservable';
import { catchError, tap } from 'rxjs/operators';
import { environment } from '../../environments/environment';
import { Router } from '@angular/router';

import { LoaderService } from './loader.service'

const api: string = environment.apiUrl;

// reference : https://angular.io/guide/http

@Injectable()
export class ApiService {
  httpOptions = {headers: {}, withCredentials: true};

  constructor(private http: HttpClient, private loaderService: LoaderService, private router: Router) { this.defaultHttpOption(); }

  defaultHttpOption(){
    const headers = new HttpHeaders({'Authorization': 'x-msxpert-token'});
    this.httpOptions.headers = headers;
  }

  get<T>(url: string): Observable<T> {
    this.loaderService.show();
    return this.http.get<T>(`${api}/${url}`, this.httpOptions)
      .pipe(
        tap(_ => this.loaderService.hide(), error => { this.loaderService.hide(), this.checkCode(error) }),
        catchError(this.handleError)
      );
  }

  post<T>(url: string, data: any): Observable<T> {
    this.loaderService.show();
    return this.http.post<T>(`${api}/${url}`, data, this.httpOptions)
      .pipe(
        tap(_ => this.loaderService.hide(), error => { this.loaderService.hide(), this.checkCode(error) }),
        catchError(this.handleError)
      );
  }

  put<T>(url: string, data: any): Observable<T> {
    this.loaderService.show();
    return this.http.put<T>(`${api}/${url}`, data, this.httpOptions)
      .pipe(
        tap(_ => this.loaderService.hide(), error => { this.loaderService.hide(), this.checkCode(error) }),
        catchError(this.handleError)
      );
  }

  delete<T>(url: string): Observable<T> {
    this.loaderService.show();
    return this.http.delete<T>(`${api}/${url}`, this.httpOptions)
      .pipe(
        tap(_ => this.loaderService.hide(), error => { this.loaderService.hide(), this.checkCode(error) }),
        catchError(this.handleError)
      );
  }

  checkCode(error: HttpErrorResponse) {
    if (error.status == 401) {
      this.router.navigate(['login']);
    }
  }


  private handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong,
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`);

    }
    // return an ErrorObservable with a user-facing error message
    return new ErrorObservable(error);
  };


}
