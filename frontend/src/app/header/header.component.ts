import { Component, OnInit } from '@angular/core';
import {trigger,state,style,transition,animate} from '@angular/animations';
import { EventEmitterService } from '../event-emitter.service';  
import { SocialAuthService } from "angularx-social-login";
import { GoogleLoginProvider } from "angularx-social-login";
import { SocialUser } from "angularx-social-login";
import { faUser,faEllipsisV,faMapMarkerAlt,faKeyboard,faSearch} from '@fortawesome/free-solid-svg-icons';
import {faGoogle} from '@fortawesome/free-brands-svg-icons';
import {Apollo} from 'apollo-angular';
import gql from 'graphql-tag';
import {User} from '../user';
// import { User } from '../model';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss',
],
  animations: [
    trigger('click',[
      state('clicked',style({
        opacity:'0.5',
        
      })),
      state('unclicked',style({
        opacity:'1',
      
      })),
      state('return',style({
        opacity:'1',
      
      })),
      transition('unclicked => *',[
        animate('0.5s'),
        // animate('0.5s')
      ]),
      transition('clicked => unclicked',[
        animate('0.5s')
      ]),
    ]),
    trigger('input',[
      state('focused',style({
        boxShadow: '0.2px 0.2px 0.2px 0.2px black'
      })),
      state('idle',style({
        boxShadow: 'none'
      })),
    ]),
    trigger('restrict',[
      state('off',style({
        width: '1.8vw',
        height: '3.6vh',
        top: '-6px',
        left:'-2px'
        
      })),
      state('on',style({
        top: '-8px',
        left:'25px',
        width: '2vw',
        height: '4vh',
        backgroundColor: 'rgba(135,122,122,1)'
      })),
      transition('off <=> on',[
        animate('0.2s')
      ])
    ]),
    trigger('openClose',[
      state('void',style({
        display:'none'
      })),
      state("*",style({
        display:'block'
      }))
    ])
  ],
})
export class HeaderComponent implements OnInit {  

  user: SocialUser;
  
  faUser= faUser;
  faGoogle= faGoogle;
  constructor(private apollo:Apollo,private eventEmitterService: EventEmitterService,private authService: SocialAuthService, private isLoggedIn : User) { }
  clicked = false;
  focus = false;
  faEllipsisV = faEllipsisV;
  faKeyboard = faKeyboard;
  faMap = faMapMarkerAlt;
  faSearch = faSearch;
  restricted=false;
  opened = false;
  imgUrl :any
  ngOnInit(): void {
    if(localStorage.getItem("uid") != null){
      const id = localStorage.getItem("uid")
      this.apollo.watchQuery({
        query: gql `
        query getUserById($id :ID!){
          getUserByID(
            id: $id
          )
          {
            id
            name
            membership_status
            billing_type
            membership_exp_date
            join_date
            email
            picture
          }
        }
        `,variables:{
          id:id
        }
      }).valueChanges.subscribe(x => {
        this.user = new SocialUser()
        this.user.name = x.data.getUserByID.name;
        this.user.photoUrl = x.data.getUserByID.picture;
        
        this.isLoggedIn.loggedIn = true;
      })
    }
    this.authService.authState.subscribe((user) => {
      
      this.user = user;
      this.isLoggedIn.loggedIn = (user != null);
      this.tryInsert();
      });
    if(this.eventEmitterService.subsVar3 == undefined){
      this.eventEmitterService.subsVar3 = this.eventEmitterService.doLogin.subscribe(()=>{
        this.login();
      });
    }
  }
  turnInputField(){
    this.focus = true;
  }
  outfocus(){
    this.focus=false;
  }
  toggle(){
    this.clicked = !this.clicked;
  }

  callSideBar(){    
    this.eventEmitterService.toggleSideBarButtonClick();    
  }
  callLoginModal(){
    this.eventEmitterService.toggleLoginModalClick();
  }    
  login(){
    
    this.authService.signIn(GoogleLoginProvider.PROVIDER_ID);
    
  }
  logout(){
    this.authService.signOut();
  }
  toggleRestrictedMode(){
    this.restricted = !this.restricted
  }
  toggleOpenClose(){
    this.opened = !this.opened
  }
  tryInsert(){
    var fr = new FileReader()
    this.imgUrl = this.user.photoUrl;
    // fr.readAsBinaryString()
    this.apollo.mutate({
      mutation: gql `
        mutation create($name: String!, $email: String!, $asd: String!){
          createNewUser(input:{
          name:$name,
          email:$email,
          picture:$asd
          })
  				{
            id
						name
            email
            picture
          }   
}
      `,variables:{
        name: this.user.name,
        email: this.user.email,
        asd: this.user.photoUrl
      }
    }).subscribe(result =>{
      // console.log();
      localStorage.setItem("uid",result.data.createNewUser.id);
      //redirect 
    });
    
 }
 
}

