import { Component, AfterViewInit, Inject, ViewChild, ElementRef, Input } from '@angular/core';
import { DOCUMENT } from '@angular/common';
import { MessageService } from './message.service';
import { Message } from './message.service';
import { App } from '../../models/app.model';

@Component({
  selector: 'log-console',
  templateUrl: './log-console.component.html',
  styleUrls: ['./log-console.component.css'],
  providers: [ MessageService ]
})
export class LogConsoleComponent implements AfterViewInit {
  @ViewChild('console') console: ElementRef;
  @Input() apps: App[];
  messages: Message[] = [];
  filterAppNames: string[] = [];
  domain: string;
  isFullScreen: boolean = false;
  consoleIcon = {'bordered': true, 'link': true, 'expand': true, 'arrows': true, 'alternate': true, 'icon':  true, 'compress': false};
  subscription;
  isStop: boolean = false;

  constructor(@Inject(DOCUMENT) private document: any, private elRef: ElementRef, private messageService: MessageService) {
  }

  public start() {
    let apps = this.apps;

    if (apps.length == 0) return;

    let location = this.document.location;
    let domain = location.hostname;
    let httpsEnabled = location.protocol == "https:";
    let url_ws = (httpsEnabled ? 'wss://' : 'ws://');
    let url = url_ws + domain + ':8082' + '/firehose?id='+this.appendAppGuid(apps);

    this.stop();

    this.messageService.connect_firehose(url, apps);
    this.subscription = this.messageService.messages.subscribe(msg => {
      if (msg) {
        for (let filterAppName of this.filterAppNames) {
          if (filterAppName == msg.app_name) {
            msg.timestamp = Math.round(msg.timestamp / 1000000);
            msg.message = atob(msg.message)
            this.messages.push(msg);
            this.console.nativeElement.scrollTop = this.console.nativeElement.scrollHeight - this.console.nativeElement.clientHeight;
          }
        }
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


  // reference : https://github.com/lwojciechowski/mildchat-client

  ngAfterViewInit() {
    let observer = new MutationObserver(() => {
      this.elRef.nativeElement.scrollTop = this.console.nativeElement.offsetHeight;
    });
    observer.observe(this.console.nativeElement, { childList: true });
  }


  logFilter(ev, appName: string) {
    if (ev.target.checked) {
      this.filterAppNames.push(appName);
    } else {
      this.filterAppNames = this.filterAppNames.filter(app => app != appName)
    }
    // console.log(this.filterAppNames)
  }

  toggleConsole() {
    this.isFullScreen = !this.isFullScreen;
    this.consoleIcon.expand = !this.isFullScreen;
    this.consoleIcon.arrows = !this.isFullScreen;
    this.consoleIcon.compress = this.isFullScreen;
  }

  private appendAppGuid(apps: App[]) {
    let guid: string = "";
    for (let app of apps) {
      guid = guid + ':' + app.appGuid;
    }
    return guid;
  }

}
