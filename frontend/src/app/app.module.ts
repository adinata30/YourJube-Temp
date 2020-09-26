import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { SideBarComponent } from './side-bar/side-bar.component';
import {LoginModalComponent} from './login-modal/login-modal.component';
import {HomeComponentComponent} from'./home-component/home-component.component';
import {UploadVideoComponent} from "./upload-video/upload-video.component";
import {VideoThumbComponent} from "./video-thumb/video-thumb.component";
import {UserComponent} from "./user/user.component";

import { EventEmitterService} from './event-emitter.service';
import { VideoComponent } from './video/video.component';
import {HttpClientModule} from '@angular/common/http';
import {ApolloModule, Apollo} from 'apollo-angular'
import {HttpLinkModule, HttpLink} from 'apollo-angular-link-http';
import { IfRolesDirective } from './if-roles.directive';
import { GraphQLModule } from './graphql.module';
import {MatVideoModule} from 'mat-video';
import { OAuthModule } from 'angular-oauth2-oidc';
import { SocialLoginModule, SocialAuthServiceConfig } from "angularx-social-login";
import { GoogleLoginProvider} from "angularx-social-login";
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import {AngularFireModule} from '@angular/fire';
import {AngularFireStorageModule} from '@angular/fire/storage';
import {environment} from '../environments/environment';
import {AngularFirestoreModule} from '@angular/fire/firestore';
import {DomSanitizer} from '@angular/platform-browser'
import { VgCoreModule } from '@videogular/ngx-videogular/core';
import { VgControlsModule } from '@videogular/ngx-videogular/controls';
import "@angular/compiler";
import {NgxDropzoneComponent,NgxDropzoneModule} from 'ngx-dropzone';
import { FormsModule } from '@angular/forms';
import {User} from './user';
import {VgStreamingModule} from '@videogular/ngx-videogular/streaming';
import {VgApiService} from '@videogular/ngx-videogular/core';
import {VgBufferingModule} from '@videogular/ngx-videogular/buffering';
import {VgOverlayPlayModule} from'@videogular/ngx-videogular/overlay-play';
@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    SideBarComponent,
    VideoComponent,
    LoginModalComponent,
    HomeComponentComponent,
    IfRolesDirective,
    UploadVideoComponent,
    VideoThumbComponent,
    UserComponent,
    // NgxDropzoneComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    ApolloModule,
    HttpLinkModule,
    GraphQLModule,
    MatVideoModule,
    SocialLoginModule,
    OAuthModule.forRoot(),
    FontAwesomeModule,
    AngularFireModule.initializeApp(environment.firebaseConfig),
    AngularFireStorageModule,
    VgCoreModule,
    VgControlsModule,
    NgxDropzoneModule,
    FormsModule,
    VgStreamingModule,
    // VgApiService ,
    VgBufferingModule,
    VgOverlayPlayModule
    
  ],
  providers: [
    EventEmitterService,
    {
      provide: 'SocialAuthServiceConfig',
      useValue: {
        autoLogin: false,
        providers: [
          {
            id: GoogleLoginProvider.PROVIDER_ID,
            provider: new GoogleLoginProvider(
              '283290995297-nlbm1uc2mejb28fdmrqqpphlvod0r1ua.apps.googleusercontent.com'
            ),
          },  
        ],
      } as SocialAuthServiceConfig,
    },User
    
  ],
  bootstrap: [AppComponent]
})
// @Injectable()
export class AppModule { 
  constructor(
    apollo: Apollo,
    httpLink: HttpLink
  ){

  }
}
