import { Component, OnInit } from '@angular/core';
import {trigger,state,style,transition,animate} from '@angular/animations';
import { EventEmitterService } from '../event-emitter.service';   
@Component({
  selector: 'app-side-bar',
  templateUrl: './side-bar.component.html',
  styleUrls: ['./side-bar.component.scss'],
  animations: [
    trigger('flyInOut',[
      state('void',style({
        transform:'translateX(-50%)'
      })),
      state('*',style({
        transform:'translateX(0%)'
      })),
      transition(':leave', [
        style({
          // transform: 'translateX(-200%)',
          // left:'400px'
        }),
        animate('0.1s')
      ]),
      transition(':enter',[
        style({
          // transform:'translateX(-300%)',
          // left:'400px'
        }),
        animate('0.1s')
      ])
    ])
  ],
})
export class SideBarComponent implements OnInit {

  flyIn=false;
  constructor(private eventEmitterService: EventEmitterService    ) { }

  ngOnInit(): void {
    if (this.eventEmitterService.subsVar==undefined) {    
      this.eventEmitterService.subsVar = this.eventEmitterService.    
      toggleSideBar.subscribe(() => {    
        this.toggleFlyInOut();    
      });    
    } 
  }
  toggleFlyInOut(){
    
    this.flyIn = !this.flyIn;
  }

  

}
