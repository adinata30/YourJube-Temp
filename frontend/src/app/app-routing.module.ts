import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { HomeComponentComponent } from './home-component/home-component.component';
import { UploadVideoComponent } from './upload-video/upload-video.component';
import { VideoThumbComponent} from './video-thumb/video-thumb.component';
import { VideoComponent } from './video/video.component';
import {UserComponent} from "./user/user.component";
import { userInfo } from 'os';
const routes: Routes = [
   {path:"home",component:HomeComponentComponent },
   {path:"",redirectTo:"/home",pathMatch:'full'},
   {path:"upload",component:UploadVideoComponent},
   {
    path:"video",component:VideoComponent,
    children:[
      {
        path:':id',
        component:VideoComponent
      }
    ]
   },
   {path:"test",component:VideoThumbComponent},
   {
     path:"user",component:UserComponent,
     children:[
       {
         path:'user/id',
         component:UserComponent
       }
     ]
   }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
