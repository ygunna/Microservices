import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { SampleGuide } from '../models/sample-guide.model'
import { Tree } from '../models/tree.model'

import 'codemirror/mode/javascript/javascript.js';

declare const $: any;
declare let ace: any;

@Component({
  selector: 'app-micro-guide',
  templateUrl: './micro-guide.component.html',
  styleUrls: ['./micro-guide.component.css']
})
export class MicroGuideComponent implements OnInit {
  @ViewChild('explorer') explorer:ElementRef;
  text: string = "";
  sampleGuide: SampleGuide = new SampleGuide();
  tree: Tree = new Tree();
  searchName: string = "";
  codeconfig = {lineNumbers: true, theme: 'lesser-dark', mode: "javascript", readOnly: true};
  CF_INSTANCE_INTERNAL_IP = '${CF_INSTANCE_INTERNAL_IP}';
  PORT = '${PORT}';

  constructor() { }

  ngOnInit() {
    $('.menu .item').tab({});

    $('#explorer').jstree(this.tree.dirs);


    $(this.explorer.nativeElement).on('click', (e) => {
      //console.log(e)
      const src: string = e.target.innerText;
      const id: string = e.target.id;

      if(src == 'FrontApplication.java'){
        this.text = this.sampleGuide.frontApplication;
      }else if(src == 'FrontController.java') {
        this.text = this.sampleGuide.frontCongroller;
      }else if(src == 'BackApplication.java'){
        this.text = this.sampleGuide.backApplication;
      }else if(src == 'BackController.java'){
        this.text = this.sampleGuide.backCongroller;
      }else if(src == 'Data.java'){
        if(id == 'front_data_anchor'){
          this.text = this.sampleGuide.frontData;
        }else{
          this.text = this.sampleGuide.backData;
        }
      }else if(src == 'pom.xml'){
        if(id == 'front_pom_anchor'){
          this.text = this.sampleGuide.frontPom;
        }else{
          this.text = this.sampleGuide.backPom;
        }
      }else if(src == 'manifest.yml'){
        if(id == 'front_manfiest'){
          this.text = this.sampleGuide.frontManifest;
        }else{
          this.text = this.sampleGuide.backManifest;
        }
      }else if(src == 'application.properties'){
        if(id == 'front_application_properties_anchor'){
          this.text = this.sampleGuide.frontApplicationProperties;
        }else{
          this.text = this.sampleGuide.backApplicationProperties;
        }
      }else if(src == 'bootstrap.properties'){
        if(id == 'front_bootstrap_properties_anchor'){
          this.text = this.sampleGuide.frontBootstrapProperties;
        }else{
          this.text = this.sampleGuide.backBootstrapProperties;
        }
      }else if(src == 'index.html'){
        this.text = this.sampleGuide.indexHtml;
      }else if(src == 'albums.js'){
        this.text = this.sampleGuide.albumsJs;
      }else if(src == 'app.js'){
        this.text = this.sampleGuide.appJs;
      }else if(src == 'errors.js'){
        this.text = this.sampleGuide.errorsJs;
      }else if(src == 'albums.html'){
        this.text = this.sampleGuide.albumsHtml;
      }else if(src == 'errors.html'){
        this.text = this.sampleGuide.errorsHtml;
      }else{
        this.text = "";
      }
    });
  }


}
