import { Component, OnInit } from '@angular/core';
import { Micro } from '../micro-list/micro.model';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

import { environment } from '../../../environments/environment';
import { ApiService } from '../../shared/api.service'

declare const $: any;

@Component({
  selector: 'app-micro-api',
  templateUrl: './micro-api.component.html',
  styleUrls: ['./micro-api.component.css']
})
export class MicroApiComponent implements OnInit {
  swaggerApiUrl: string = environment.swaggerApiUrl;
  api: string = environment.apiUrl;
  apiUrl: string = 'microservices/api/list?offset=';
  offset: number = 0;
  micros: Micro[];
  isEnd: boolean = false;
  searchName: string = "";
  iframeSrc: SafeResourceUrl;

  constructor(private apiService: ApiService, private sanitizer: DomSanitizer) {
    this.iframeSrc = this.sanitizer.bypassSecurityTrustResourceUrl('about:blank');
  }

  ngOnInit() {
    this.listMicroserviceApi();
  }

  listMicroserviceApi() {
    this.apiService.get<Micro[]>(`${this.apiUrl}${this.offset}`).subscribe(
      data => {
        this.micros = data;

      }

    );
  }

  more() {
    this.offset += 6;
    this.apiService.get<Micro[]>(`${this.apiUrl}${this.offset}`).subscribe(
      data => {
        data.forEach(d => {this.micros.push(d)});
        if(this.micros.length <= this.offset){
          this.isEnd = true;
        }
      }
    );
  }

  popModalFront(id){
    $('.ui.modal').modal({detachable: false}).modal('show');
    //console.log(this.swaggerApiUrl +'/entry/?id='+id+'&domain='+ this.api)
    this.iframeSrc = this.sanitizer.bypassSecurityTrustResourceUrl(this.swaggerApiUrl +'/entry/?id='+id+'&domain='+ this.api);
  }

}
