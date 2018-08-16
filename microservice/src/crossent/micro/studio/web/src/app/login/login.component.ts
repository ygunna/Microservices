import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { ApiService } from '../shared/api.service';
import { Observable } from 'rxjs/Rx';
import { CookieService } from 'ngx-cookie-service';

class TokenRequest {
  public clientId: string;
  public clientSecret: string;
  public grantType: string;
  public responseType: string;
  public username: string;
  public password: string;

  constructor(clientId: string, clientSecret: string, grantType: string, responseType: string, username: string, password: string) {
    this.clientId = clientId;
    this.clientSecret = clientSecret;
    this.grantType = grantType;
    this.responseType = responseType;
    this.username = username;
    this.password = password;
  }
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  httpOptions = {headers: {}};
  username: string;
  password: string;

  constructor(private http: HttpClient,
              private cookieService: CookieService,
              private apiService: ApiService,
              private router:Router) { }

  ngOnInit() {
    // this.getCsrfToken();
  }

  login() {
    var tokenRequest = new TokenRequest('', '', 'password', 'token', this.username, this.password);
    this.apiService.post('login', tokenRequest).subscribe(
      res => {
        localStorage.setItem('username', this.username);
        this.router.navigate(['/list']);
      },
      err => {
        if(err.status == 401) {
          alert('아이디 또는 비밀번호가 일치하지 않습니다.');
        } else {
          alert(err.message);
        }
      }
    );
    //noinspection TypeScriptValidateTypes
    /*this.postAutoLogin('https://uaa.bosh-lite.com/oauth/token', headers, tokenRequest)
      .subscribe(
        res => {
          console.log(res.headers.get('Set-Cookie'));
          console.log(res.headers.get('Content-Type'));
          console.log(res.body);
          console.log('X-Uaa-Csrf:'+this.cookieService.get('X-Uaa-Csrf'));
          console.log('X-UAA-CSRF:'+this.cookieService.get('X-UAA-CSRF'));
        },
        err => {
          console.log(err);
        }
      );*/
  }

  postAutoLogin(url: string, headers: HttpHeaders, data: any): Observable<HttpResponse<any>> {
    return this.http.post(url, data, {
      headers: headers,
      withCredentials: true,
      observe: 'response'
    });
  }


/*  getCsrfToken() {
    const headers = new HttpHeaders({'Authorization': 'application/x-www-form-urlencoded'});
    //noinspection TypeScriptValidateTypes
    this.getFullResponseForWriter('http://uaa.bosh-lite.com/oauth/authorize?response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A4200%2Flist&client_id=portal-id', headers)
      .subscribe(
        res => {
          console.log(res.headers.get('Set-Cookie'));
          console.log(res.headers.get('Content-Type'));
          console.log(res.body);
          console.log('X-Uaa-Csrf:'+this.cookieService.get('X-Uaa-Csrf'));
          console.log('X-UAA-CSRF:'+this.cookieService.get('X-UAA-CSRF'));
        },
        err => {
          console.log(err);
        }
    );
  }

  getFullResponseForWriter(url: string, headers: HttpHeaders): Observable<HttpResponse<any>> {
    return this.http.get(url, {
      headers: headers,
      // withCredentials: true,
      observe: 'response'
    });
  }*/

}
