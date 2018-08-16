import { Component, AfterViewInit, Inject, ViewChild, ElementRef, Input } from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { MessageService } from '../log-console/message.service';
import { Message } from '../log-console/message.service';
import { App } from '../../models/app.model';

@Component({
  selector: 'app-log-console',
  templateUrl: './app-log-console.component.html',
  styleUrls: ['./app-log-console.component.css'],
  providers: [ MessageService ]
})
export class AppLogConsoleComponent implements AfterViewInit {
  @ViewChild('console') console: ElementRef;
  @Input() apps: App[];
  messages: Message[] = [];
  filterAppNames: string[] = [];
  domain: string;
  isFullScreen: boolean = false;
  consoleIcon = {'bordered': true, 'link': true, 'expand': true, 'arrows': true, 'alternate': true, 'icon':  true, 'compress': false};
  selectApp: string;
  subscription;
  isStop: boolean = false;

  constructor(@Inject(DOCUMENT) private document: any, private elRef: ElementRef, private messageService: MessageService) { }

  ngAfterViewInit() {
    let observer = new MutationObserver(() => {
      this.elRef.nativeElement.scrollTop = this.console.nativeElement.offsetHeight;
    });
    observer.observe(this.console.nativeElement, { childList: true });
  }

  public start() {
    //let apps = this.apps;
    if ( this.selectApp == "" ) return;

    let location = this.document.location;
    let domain = location.hostname;
    let httpsEnabled = location.protocol == "https:";
    let url_ws = (httpsEnabled ? 'wss://' : 'ws://');
    let url = url_ws + domain + ':8082' + '/stream?id='+this.selectApp;

    this.stop();
    this.messageService.connect_stream(url);
    this.subscription = this.messageService.messages.subscribe(msg => {
      if (msg) {
        msg.timestamp = Math.round(msg.timestamp / 1000000);
        msg.message = atob(msg.message)
        this.messages.push(msg);
        this.console.nativeElement.scrollTop = this.console.nativeElement.scrollHeight - this.console.nativeElement.clientHeight;
      }
    });
    this.isStop = false;
  }

  public stop() {
    console.log('websocket stop');
    if (this.subscription) {
      this.subscription.unsubscribe();
      this.isStop = true;
    }
  }

  public stopNstart() {
    this.stop();
    this.messages = [];
    this.start();
  }

  toggleConsole() {
    this.isFullScreen = !this.isFullScreen;
    this.consoleIcon.expand = !this.isFullScreen;
    this.consoleIcon.arrows = !this.isFullScreen;
    this.consoleIcon.compress = this.isFullScreen;
  }

}
