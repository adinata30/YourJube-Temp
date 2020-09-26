import { Injectable, EventEmitter } from '@angular/core';
import { Subscription } from 'rxjs/internal/Subscription';
@Injectable({
  providedIn: 'root'
})
export class EventEmitterService {

  constructor() { }
  toggleSideBar = new EventEmitter();  
  toggleLoginModal = new EventEmitter();  
  doLogin = new EventEmitter();
  subsVar: Subscription; 
  subsVar2: Subscription;
  subsVar3: Subscription;     
  toggleSideBarButtonClick() {    
    this.toggleSideBar.emit(null);    
  }
  toggleLoginModalClick(){
    
    this.toggleLoginModal.emit(null);
  }   
  doLoginClick(){
    this.doLogin.emit(null);
  }
}
