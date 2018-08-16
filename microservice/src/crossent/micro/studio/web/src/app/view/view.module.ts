import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MicroListComponent } from './micro-list/micro-list.component';
import { MicroDetailComponent } from './micro-detail/micro-detail.component';
import { SearchFilterPipe } from './micro-list/search-filter.pipe';
import { D3StudioModule } from '../d3-studio/d3-studio.module';
import { D3ViewService } from './micro-detail/d3-view.service'
import { WebsocketService } from './micro-detail/log-console/websocket.service';
import { MessageService } from './micro-detail/log-console/message.service';

import { CodemirrorModule } from 'ng2-codemirror';
import { LogConsoleComponent } from './micro-detail/log-console/log-console.component';
import { MicroGuideComponent } from './micro-guide/micro-guide.component';
import { MicroApiComponent } from './micro-api/micro-api.component';
import { AppLogConsoleComponent } from './micro-detail/app-log-console/app-log-console.component';

@NgModule({
  imports: [
    FormsModule, CommonModule, RouterModule, CodemirrorModule, D3StudioModule
  ],
  declarations: [MicroListComponent, MicroDetailComponent, SearchFilterPipe, LogConsoleComponent, MicroGuideComponent, MicroApiComponent, AppLogConsoleComponent],
  schemas: [ CUSTOM_ELEMENTS_SCHEMA ],
  providers: [ D3ViewService, WebsocketService, MessageService ]
})
export class ViewModule { }
