import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs/Rx';
import 'rxjs/add/operator/map'
import { WebsocketService } from './websocket.service';
import { App } from '../../models/app.model';

export interface Message {
  app_name: string,
  app_id: string,
  message: string,
  message_type: string,
  source_instance: string,
  source_type: string,
  timestamp: number
}

@Injectable()
export class MessageService {

  public messages: Subject<Message>;

  constructor(private wsService: WebsocketService) {}

  connect_firehose(url: string, apps: App[]) {
    this.messages = <Subject<Message>>this.wsService
      .connect(url)
      .map((response: MessageEvent): Message => {
        let data = JSON.parse(response.data);
        if (data['logMessage']){
          // error message type = 2
          if (data.logMessage.message_type === 2 || data.logMessage.source_type.indexOf('APP/') != -1) {
            for (let app of apps) {
              if (app.appGuid == data.logMessage.app_id) {
                return {
                  app_name: app.appName,
                  app_id: data.logMessage.app_id,
                  message: data.logMessage.message,
                  message_type: data.logMessage.message_type,
                  source_instance: data.logMessage.source_instance,
                  source_type: data.logMessage.source_type,
                  timestamp: data.logMessage.timestamp
                }
              }
            }
          }
        }
      });
  }

  connect_stream(url: string) {
    this.messages = <Subject<any>>this.wsService
      .connect(url)
      .map(response => {
        //return JSON.parse(response.data);
        let data = JSON.parse(response.data);
        // error message type = 2
        if (data.message_type === 2 || data.source_type.indexOf('APP/') != -1) {
          return {
            app_id: data.app_id,
            message: data.message,
            message_type: data.message_type,
            source_instance: data.source_instance,
            source_type: data.source_type,
            timestamp: data.timestamp
          }
        }

      });
  }

}
