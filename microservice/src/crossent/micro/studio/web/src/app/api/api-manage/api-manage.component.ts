import { Component, OnInit } from '@angular/core';
import { MicroApi } from '../models/microapi.model';
import { ApiService } from '../../shared/api.service';
import { PartService } from '../api-create/part.service';
import { Part } from '../models/part.model';
import { Org } from '../../shared/models/org.model';

@Component({
  selector: 'app-api-manage',
  templateUrl: './api-manage.component.html',
  styleUrls: ['./api-manage.component.css']
})
export class ApiManageComponent implements OnInit {
  apiUrl: string = 'apigateway?offset=';
  offset: number = 0;
  microapis: MicroApi[];
  parts: Part[];
  searchName: string = "";
  searchPath: string = "";
  isEnd: boolean = false;
  part: string;
  orgs : Org[];
  org: Org;

  constructor(private apiService: ApiService,
              private partService: PartService) { }

  ngOnInit() {
    this.parts = this.partService.partsList();
    this.listOrgs();
    //this.listApis();
  }

  listOrgs() {
    this.apiService.get('manageorgs').subscribe(
      data => {
        this.orgs = data['resources'];
        if(this.orgs && this.orgs.length > 0) {
          this.org = this.orgs[0];
          this.onChangeOrg();
        }
      }
    );
  }

  listApis() {
    this.apiService.get<MicroApi[]>(`${this.apiUrl}${this.offset}&orgguid=${this.org.metadata.guid}`).subscribe(
      data => {
        console.log(data)
        this.microapis = data;
      }
    );
  }

  onChangeOrg() {
    var guid = this.org.metadata.guid;
    this.microapis = [];
    this.offset = 0;
    this.isEnd = false;
    this.listApis();
  }

  more() {
    this.offset += 6;
    this.apiService.get<MicroApi[]>(`${this.apiUrl}${this.offset}&orgguid=${this.org.metadata.guid}`).subscribe(
      data => {
        data.forEach(d => {this.microapis.push(d)});
        if(this.microapis.length <= this.offset){
          this.isEnd = true;
        }
      }
    );
  }

  deleteApi(id: number) {
    if(!confirm("삭제하시겠습니까?")) {
      return;
    }

    this.apiService.delete<MicroApi>(`apigateway/${id}`).subscribe(
      data => {
        console.log(data)
        alert("삭제되었습니다.");
        this.listApis();
      }
    );
  }

}
