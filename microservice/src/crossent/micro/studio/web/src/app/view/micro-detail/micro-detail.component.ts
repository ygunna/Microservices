import { Component, AfterViewInit, Input, ViewChild, ElementRef, ViewChildren, QueryList } from '@angular/core';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';
import { ActivatedRoute, Router  } from "@angular/router";
import { ApiService } from '../../shared/api.service';
import { LoaderService } from '../../shared/loader.service';
import { Micro } from '../micro-list/micro.model';
import { App } from '../models/app.model';
import { Service } from '../models/service.model';
import { Policy } from '../models/policy.model';
import { Node } from '../models/node.model';
import { Link } from '../models/link.model';
import { environment } from '../../../environments/environment';
import { D3ViewService } from './d3-view.service';
import { LogConsoleComponent } from './log-console/log-console.component'
import { ZoomableDirective } from '../../d3-studio/directives/zoomable.directive';

declare const $: any;

@Component({
  selector: 'app-micro-detail',
  templateUrl: './micro-detail.component.html',
  styleUrls: ['./micro-detail.component.css']
})
export class MicroDetailComponent implements AfterViewInit {
  apiUrl: string = 'microservices';
  swaggerApiUrl: string = environment.swaggerApiUrl;
  api: string = environment.apiUrl;
  id: string;
  micro: Micro = new Micro(0, "", "", "", "");
  apps: App[] = [];
  services: Service[];
  policies: Policy[];
  servicesApps: Service[];
  routings: Service[];
  registries: Service[];
  configurations: Service[];
  iframeSrc: SafeResourceUrl;
  gatewayapp: string;
  nodes: Node[];
  links: Link[];
  filter: string = "0";
  frontend: string;
  swaggers: Array<Swagger> = [];

  @Input('droppedNodes') droppedNodes: any = [];
  //@ViewChildren(LogConsoleComponent) ref: QueryList<LogConsoleComponent>;
  @ViewChild(ZoomableDirective) directive = null;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private apiService: ApiService,
              private sanitizer: DomSanitizer,
              private loaderService: LoaderService,
              private d3ViewService: D3ViewService) {
    this.iframeSrc = this.sanitizer.bypassSecurityTrustResourceUrl('about:blank');
    this.id = this.route.snapshot.params['id'];
  }


  ngAfterViewInit() {
    $('#viewer > .ui.accordion').accordion({animateChildren: true});
    $('#details > .ui.accordion').accordion({animateChildren: true, onOpen: function () {
      $('#viewer').prop('class', 'eight wide column');
      $('#details').prop('class', 'eight wide column');
    }, onClose: function (e) {
      $('#viewer').prop('class', 'fourteen wide column');
      $('#details').prop('class', 'two wide column');
    }});
    // this.ref.changes.subscribe((comps: QueryList<LogConsoleComponent>) =>
    // {
    //   this.ref.first.start(this.apps)
    // });

    this.getMicroserviceDetail();
    this.getMicroservice();
    this.getMicroserviceLink();

  }

  getMicroservice() {
    this.swaggers = [];
    this.apiService.get<Micro>(`${this.apiUrl}/${this.id}`).subscribe(
      data => {
        this.micro = data['microservice'];
        if (this.micro.swagger) {
          let obj = JSON.parse(this.micro.swagger);
          for (let s in obj['paths']){
            for (let p in obj['paths'][s]) {
              this.swaggers.push(new Swagger(p, s));
            }
          }
        }
      }
    );
  }

  getMicroserviceDetail() {
    this.apiService.get<any>(`${this.apiUrl}/detail/${this.id}`).subscribe(
      data => {
        if (data['apps']) {
          let a = data['apps'];
          this.apps = a['app'];
          this.services = a['service'];
          this.policies = a['policy'];
          this.servicesApps = a['serviceApp'];

          this.apps.forEach(d => {
            if (d.appName.startsWith("gatewayapp")) {
              this.gatewayapp = d.url;
            }
          });
        }
        this.routings = data['routes'];
        this.registries = data['registries'];
        this.configurations = data['properties'];

      }
    );

  }



  getMicroserviceLink() {
    this.apiService.get<Service[]>(`${this.apiUrl}/link/${this.id}`).subscribe(
      data => {
        this.loaderService.forceHide();
        this.d3ViewService.updatePath(data['nodes'], data['links']);
      }
    );
  }

  onChange(event: any) {
    if (this.filter == "0") {
      this.d3ViewService.updatePath(this.nodes, this.links);
    } else {
      let newNodes: Node[] = [];
      let newLinks: Link[] = [];
      this.nodes.forEach(d => {
        if (d.type == 'App') {
          newNodes.push(d);
        }
      });
      this.links.forEach(d => {
        if (d.type == 'App') {
          newLinks.push(d);
        }
      });

      this.d3ViewService.updatePath(newNodes, newLinks);
    }
  }

  redoSvg(){
    this.getMicroserviceLink();
  }

  zoom(direction) {
    this.directive.zoomClick(direction);
  }

  popModal(service){
    $('.ui.modal').modal({detachable: false}).modal('show');
    this.iframeSrc = this.sanitizer.bypassSecurityTrustResourceUrl(this.swaggerApiUrl +'?service='+service+'&gateway='+ this.gatewayapp);
  }


  popModalFront(url){
    $('.ui.modal').modal({detachable: false}).modal('show');
    this.iframeSrc = this.sanitizer.bypassSecurityTrustResourceUrl(this.swaggerApiUrl +'/entry/?id='+this.micro.id+'&domain='+ this.api);
  }

  makeApi() {
    if (this.frontend) {
      this.micro.url = this.frontend;
      this.apiService.put<Micro[]>(`${this.apiUrl}/${this.id}/api`, this.micro).subscribe(
        data => {
          alert('save ok.');
          this.getMicroservice();
        }
      );
    }else{
      alert('entrypoint가 없습니다.');
      return
    }
  }

  start() {
    let data = {
      name: this.micro['name'],
      spaceGuid: this.micro.spaceGuid,
      status: 'STARTED'
    };
    this.apiService.put('microservices/'+this.micro.id+'/state', data).subscribe(
      res => {
        alert("저장되었습니다.");
      }, err => {
        console.log(JSON.stringify(err.headers));
        console.log(err.status+" "+err.message);
      }
    );
  }

  stop() {
    let data = {
      name: this.micro['name'],
      spaceGuid: this.micro.spaceGuid,
      status: 'STOPPED'
    };
    this.apiService.put('microservices/'+this.micro.id+'/state', data).subscribe(
      res => {
        alert("저장되었습니다.");
      }, err => {
        console.log(JSON.stringify(err.headers));
        console.log(err.status+" "+err.message);
      }
    );
  }

  delete() {
    let data = {
      spaceGuid: this.micro.spaceGuid,
      status: 'STOPPED'
    };
    if( confirm("삭제하시겠습니까 ?") ) {
      this.apiService.delete('microservices/' + this.micro.id).subscribe(
        res => {
          alert("삭제되었습니다.");
          this.router.navigate(['list']);
        }, err => {
          console.log(JSON.stringify(err.headers));
          console.log(err.status + " " + err.message);
        }
      );
    }
  }
}

class Swagger {
  constructor(
    public method: string,
    public path: string
  ){}
}
