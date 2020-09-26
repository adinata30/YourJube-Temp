import { Component, OnInit } from '@angular/core';
import { AngularFireStorage} from "@angular/fire/storage";
import { runInThisContext } from 'vm';
import {AngularFireUploadTask} from '@angular/fire/storage/task'
import { Observable } from 'rxjs';
import { finalize } from 'rxjs/operators';

@Component({
  selector: 'app-home-component',
  templateUrl: './home-component.component.html',
  styleUrls: ['./home-component.component.scss']
})


export class HomeComponentComponent implements OnInit {
  


  constructor(){}
  ngOnInit():void{

  }
}
