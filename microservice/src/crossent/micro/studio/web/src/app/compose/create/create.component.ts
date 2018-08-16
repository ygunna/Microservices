import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ApiService } from '../../shared/api.service';
import { Org } from '../../shared/models/org.model';
import { Space } from '../../shared/models/space.model';
import { Service } from '../../shared/models/service.model';
import { Microservice } from '../../shared/models/microservice.model';
import { SampleApp } from '../models/sample-app.model';
import { environment } from '../../../environments/environment';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.css']
})
export class CreateComponent implements OnInit {
  orgs : Org[];
  spaces : Space[];
  microservice = new Microservice(0,null,null,'','',null,true);
  step : number;
  essential : number;
  frontend: SampleApp = new SampleApp();
  backend: SampleApp = new SampleApp();
  config: Service;
  registry: Service;
  gateway: Service;

  constructor(private apiService: ApiService,
              private router: Router) { }

  ngOnInit() {
    this.step = 1;
    this.listOrgs();
    this.essential = 2;

    this.apiService.get('marketplace').subscribe(
      data => {
        for(var service of data['resources']) {
          service.entity.service_plan_guid = service.entity.service_plans[0];
          if(service.entity.label == 'micro-config-server') {
            this.config = service;
            this.config.entity.name = 'config-server';
          } else if(service.entity.label == 'micro-registry-server') {
            this.registry = service;
            this.registry.entity.name = 'registry-server';
          } else if(service.entity.label == 'micro-gateway-server') {
            this.gateway = service;
            this.gateway.entity.name = 'gateway-server';
          }
        }
      }
    );
    this.apiService.get(encodeURI('apps?q=name%20IN%20' + environment.sampleApps)).subscribe(
      data => {
        for(var app of data['resources']) {
          if(app.entity.name == 'front') {
            this.frontend.app = app;
          } else if(app.entity.name == 'back') {
            this.backend.app = app;
          }
        }
      }
    );
  }

  listOrgs() {
    this.apiService.get('orgs').subscribe(
      data => {
        this.orgs = data['resources'];
        if(this.orgs.length > 0) {
          this.microservice.org = this.orgs[0];
          this.onChangeOrg();
        }
      }
    );
  }
  listOrgSpaces(guid) {
    this.apiService.get('orgs/'+guid+'/spaces').subscribe(
      data => {
        this.spaces = data['resources'];
        this.microservice.space = this.spaces[0];
      }
    );
  }

  onChangeOrg() {
    var guid = this.microservice.org.metadata.guid;
    this.listOrgSpaces(guid);
  }

  Next() {
    this.apiService.get('microservices?name='+this.microservice.name).subscribe(
      data => {
        if(data == 0) {
          this.step = 2;
        } else {
          alert("입력하신 마이크로서비스명이 이미 존재합니다.");
          this.microservice.name = '';
        }
      }
    );
  }
  Previous() {
    this.step = 1;
  }

  Create() {
    this.microservice['orgGuid'] = this.microservice.org.metadata.guid;
    this.microservice['orgName'] = this.microservice.org.entity.name;
    this.microservice['spaceGuid'] = this.microservice.space.metadata.guid;
    this.microservice['spaceName'] = this.microservice.space.entity.name;
    this.microservice.services = {resources: [this.config, this.registry]};
    var apps = [];
    if(this.frontend.checked) apps.push(this.frontend.app);
    if(this.backend.checked) apps.push(this.backend.app);
    this.microservice.apps = {resources: apps};

    this.apiService.post('microservices', this.microservice).subscribe(
      data => {
        this.router.navigate(['/edit/'+data['id']]);
      },
      err => {
        console.log('err >>>>>>>>>>> ');
        console.log(err);
      }
    );
  }


}
