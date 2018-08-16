import { Injectable } from '@angular/core';
import { HttpRequest, HttpResponse, HttpHandler, HttpEvent, HttpInterceptor, HttpErrorResponse } from '@angular/common/http';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Observable';

@Injectable()
export class AppInterceptor implements HttpInterceptor {
  constructor(private router:Router) {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {

    request = request.clone({
      /*setHeaders: {
        Authorization: `Bearer ${this.auth.getToken()}`
      }*/
    });

    //noinspection TypeScriptValidateTypes
    return next.handle(request).do((event: HttpEvent<any>) => {
      if (event instanceof HttpResponse) {
        // do stuff with response if you want
      }
    }, (err: any) => {
      if (err instanceof HttpErrorResponse) {
        if (err.status === 0) {
          alert('알 수 없는 오류가 발생하였습니다.\n관리자에게 문의하시기 바랍니다.');
        } else if (err.status === 401) {
          this.router.navigate(['/login']);
        }
      }
    });
  }
}
