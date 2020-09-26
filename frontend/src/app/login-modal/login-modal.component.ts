import { Component, OnInit } from '@angular/core';
import {faGoogle} from '@fortawesome/free-brands-svg-icons';
import { EventEmitterService } from '../event-emitter.service';   
import {trigger,state,style,transition,animate} from '@angular/animations';
@Component({
  selector: 'app-login-modal',
  templateUrl: './login-modal.component.html',
  styleUrls: ['./login-modal.component.scss'],
  animations:[
    trigger('toggleModal',[
      state('void',style({
        transform: 'scaleY(0)'
      })),
      state('*',style({
        transform: 'scaleY(1)'
      })),
      transition(':leave',[
        animate('0.1s')
      ]),
      transition(':enter',[
        animate('0.1s')
      ])
    ])
  ],
})
export class LoginModalComponent implements OnInit {
  faGoogle= faGoogle;
  toggled = false;
  constructor(private eventEmitterService: EventEmitterService) { }

  ngOnInit(): void {
    if (this.eventEmitterService.subsVar2==undefined) {    
      
      this.eventEmitterService.subsVar2 = this.eventEmitterService.    
      toggleLoginModal.subscribe(() => {    
        
        this.toggleFlyInOut();    
      });    
    } 
    
  }
  toggleFlyInOut(){
    
    this.toggled = !this.toggled;
  }
  doLogin(){
    this.eventEmitterService.doLoginClick();
    this.toggleFlyInOut();
  }

}
