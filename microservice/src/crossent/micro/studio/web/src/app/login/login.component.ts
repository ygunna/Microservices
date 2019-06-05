import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient, HttpHeaders, HttpResponse } from '@angular/common/http';
import { ApiService } from '../shared/api.service';
import { Observable } from 'rxjs/Rx';
import { CookieService } from 'ngx-cookie-service';
import { AuthService } from '../shared/auth.service';

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
              private router:Router,
              private authService: AuthService) { }

  ngOnInit() {
    // this.getCsrfToken();
  }

  login() {
    var tokenRequest = new TokenRequest('', '', 'password', 'token', this.username, this.password);
    this.apiService.post<any>('login', tokenRequest).subscribe(
      res => {
        localStorage.setItem('username', this.username);
        localStorage.setItem('auth', res.result);
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

  }

  postAutoLogin(url: string, headers: HttpHeaders, data: any): Observable<HttpResponse<any>> {
    return this.http.post(url, data, {
      headers: headers,
      withCredentials: true,
      observe: 'response'
    });
  }




}
