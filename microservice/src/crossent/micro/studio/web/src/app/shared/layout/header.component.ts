import { Component, OnInit } from '@angular/core';
import { Router, NavigationEnd } from '@angular/router';
import { ApiService } from '../../shared/api.service';
import { LoaderService } from '../../shared/loader.service'
import { AuthService } from '../../shared/auth.service';

declare const $: any;

@Component({
  selector: 'app-layout-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  isActive: boolean = false;
  buttonHide: boolean = false;
  username: any;

  constructor(private apiService: ApiService,
              private loaderService: LoaderService,
              private router:Router,
              public authService: AuthService) { }

  ngOnInit() {
    this.loaderService.changeActive.subscribe((d: boolean) => { this.isActive = d; });

    this.router.events.subscribe((event) => {
      this.username = localStorage.getItem('username');

      if (event instanceof NavigationEnd) {
        if(event.url.split('/')[1] == 'login') {
          this.buttonHide = true;
        } else {
          this.buttonHide = false;
        }
      }
    });

    try {
      $.fn.transition.settings.silent = true;
    }catch(e){}
    $('.ui.dropdown.semantic-dropdown').dropdown();
  }

  logout() {
    localStorage.clear();
    this.apiService.post('logout', null).subscribe(
      res => {
        this.router.navigate(['/login']);
      }
    );
  }
}
