import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { MicroListComponent } from './view/micro-list/micro-list.component'
import { MicroDetailComponent } from './view/micro-detail/micro-detail.component'
import { CreateComponent } from './compose/create/create.component';
import { EditComponent } from './compose/edit/edit.component';
import { LoginComponent } from './login/login.component';
import { MicroGuideComponent } from './view/micro-guide/micro-guide.component'
import { MicroApiComponent } from './view/micro-api/micro-api.component'

const routes: Routes = [
  { path: 'list', component: MicroListComponent },
  { path: 'guide', component: MicroGuideComponent },
  { path: 'api', component: MicroApiComponent },
  { path: 'detail/:id', component: MicroDetailComponent },
  { path: 'create', component: CreateComponent },
  { path: 'edit/:id', component: EditComponent },
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: 'list', pathMatch: 'full' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
