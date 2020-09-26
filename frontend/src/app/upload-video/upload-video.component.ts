import { Component, OnInit, LOCALE_ID } from '@angular/core';
import {HeaderComponent} from '../header/header.component';
import {NgxDropzoneComponent} from 'ngx-dropzone';
import {AngularFireStorage} from '@angular/fire/storage';
import {AngularFireUploadTask} from '@angular/fire/storage/task'
import { Observable } from 'rxjs';
import {Apollo} from 'apollo-angular';
import gql from 'graphql-tag';
import {trigger,state,style,transition,animate} from '@angular/animations';
import { formatDate, getLocaleId } from '@angular/common';
import { prototype } from 'events';
import { ValueConverter } from '@angular/compiler/src/render3/view/template';
import { toBase64String } from '@angular/compiler/src/output/source_map';
import { ReadVarExpr } from '@angular/compiler';
import {Category} from "../model";

@Component({
  selector: 'app-upload-video',
  templateUrl: './upload-video.component.html',
  styleUrls: ['./upload-video.component.scss'],
  animations:[
    trigger('restrict',[
      state('off',style({
        width: '1.8vw',
        height: '3.6vh',
        top: '-21px',
        left:'-2px'
        
      })),
      state('on',style({
        top: '-23px',
        left:'25px',
        width: '2vw',
        height: '4vh',
        backgroundColor: 'rgba(135,122,122,1)'
      })),
      transition('off <=> on',[
        animate('0.2s')
      ])
    ])
  ]
})

export class UploadVideoComponent implements OnInit {


  fileName ="-"
  // }
  percentage : string ="0%"
  a = false
  progress:Observable<number>
  filePath:string
  task :AngularFireUploadTask
  constructor(private storage : AngularFireStorage, private apollo:Apollo) { }
  path : any
  pathUrl:any
  downloadURL: Observable<String>;
  fileExists = false
  categories: Category[]
  thisFile :File
  

  ngOnInit(): void {
  
      
  console.log("Category")

    this.apollo.query({
      query:gql `
    query getCategory{
    categories{
      name
    }
  }
      `
    }).subscribe((result:any)=>{
      console.log("Category Done")
      this.categories=result.data.categories
      
    })
  }
  update(event){
    console.log(event)
    this.thisFile = <File>event.addedFiles[0]
    
  
    
    this.fileName = this.thisFile.name
    this.title=this.fileName.substr(0,this.fileName.length-4)
    this.fileExists = true
    this.uploadVid()
  }
  upload(event : any){
    
    if (event.target.files){
      this.path = event.target.files[0]
      
      var reader = new FileReader();
      reader.readAsDataURL(this.path);
      reader.onload = (event)=>{
        this.pathUrl = (<FileReader>event.target).result
      }
      this.fileName = this.path.name;
    }
    
    
  }
  uploadVid(){
    
    
    // this.task  = this.storage.upload("/Video/"+generate(16),this.path)
    // const filePath = "Video/"+generate(4)
    
    
console.log("Uploading")
   
    this.filePath = "Video/"+this.generateString()
    
    const fileRef = this.storage.ref(this.filePath)
    console.log(this.fileName)
    const uploadTask = this.storage.upload(this.filePath,this.thisFile)
    
    uploadTask.percentageChanges().subscribe(x => {this.percentage = Math.round(x)+"%"
    console.log(x)
  }) 
    
    uploadTask.then((x) =>
    {
      
      fileRef.getDownloadURL().subscribe(url => {
        console.log("Uploading Done")
        this.downloadURL = url
        this.a=true
        this.checkUploadable()
      })
    }

    )
    // this.progress = this.task.percentageChanges()
    // this.task.snapshotChanges().pipe(
    //   finalize(() => {this.downloadURL = fileRef.getDownloadURL()
    //   this.a= true
    //   console.log(this.downloadURL)
    // }
    //   )
    // ).subscribe()
    
    // console.log(generate(4))
    
    // console.log(filePath)
    
    // console.log(this.downloadURL)
  }
  

  generateString(){
    const chars = "ABCDEFGHIJKLMNOPRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    var ret = ""
    for (let index = 0; index < 16; index++) {
      ret+= chars.charAt(parseInt(Math.random()*61+""))
    }
    return ret
  }
  thumbnail : File
  thumbnailExists = false
  img :any
  onSelect(event) {
    
        
    this.thumbnail = event.addedFiles[0]
    
    var fr = new FileReader()
    // fr.readAsText(this.thumbnail)
    fr.readAsDataURL(this.thumbnail)
    fr.onload = () => this.img= fr.result
    this.thumbnailExists = true
    this.checkUploadable()
	}

	removeThumbnail() {
		this.thumbnailExists = false
	}

  restricted = false
  toggleRestrictedMode(){
    this.restricted = !this.restricted
  }

  publishNow=true
  publish()
  {
    this.publishNow = !this.publishNow
    this.checkUploadable()
  }
  today = new Date()
  minTime = new Date().getTime()
  //input
  title: string =""
  desc: string ="No Description"
  ctg: Category
  privacy:string = ""
  videoType :string = ""
  publishTime = new Date()
  publishDate :Date
  refresh()
  {
    console.log(this.title)
    console.log(this.desc)
    console.log(this.ctg)
    console.log(this.privacy)
    console.log(this.videoType)
    console.log(this.publishTime)
  }
  changeCategory(event){
    
    // event.target.forEach(element => {
    //   if(element.selected)this.category = element.value
    // });
  }

  changePrivacy(event){
    this.privacy=event.target.value
  }
  changePremium(event)
  {
    this.videoType=event.target.value
  }
  changeDate(event){
    // this.publishTime = formatDate("30th August 2000","DD/MM/YY",LOCALE_ID.toString())
    // this.publishTime.setDate(event)
    // console.log(this.publishTime.setDate(new Date().getDate()+new Date(event).getDate()))
    this.publishTime.setTime(new Date(event).getTime())
    this.publishDate = new Date(event)
    // console.log(this.publishDate)
    if((this.today.getDate() == this.publishDate.getDate())&&(this.today.getMonth() == this.publishDate.getMonth())&&(this.today.getFullYear() == this.publishDate.getFullYear())){
      this.minTime = new Date().getTime()
    }
    this.checkUploadable()
  }
  changeTime(event){
    // this.publishTime.setTime(event)
    this.publishTime.setTime(this.publishDate.getTime()+event.target.valueAsNumber-25200000)
    this.checkUploadable()
    // console.log(new Date(this.publishTime))
  }
  
  checkUploadable()
  {
    this.uploadable = (
      (this.publishTime.getTime() - new Date().getTime() >= 0 || this.publishNow)&& 
      this.percentage == "100%" &&
      this.title.trim() != "" &&
      this.thumbnailExists &&
      this.ctg
    ) ? true : false

  }

  //button upload
  uploadable = false

  finishUpload(){
    console.log(this.desc)
    console.log(this.title)
    console.log("Uploading...")
    this.apollo.mutate({
      mutation:gql `
      mutation uploadVid(
        $title:String! = "",
        $link:String! = "",
        $desc:String! = "",
        $label:String! = "",
        $privacy:String! = "",
        $uploader_id:Int! = 0,
        $type:String! = "",
        $thumbnail:String! = "",
        $location:String! = "",
        $category:String! = "",
      ){
  createNewVideo(input:{
  video:$link,
  title:$title,
  descriptions:$desc,
  label:$label,
  privacy:$privacy,
  uploader_id:$uploader_id,
  type:$type,
  thumbnail:$thumbnail,
  location:$location,
  category:$category
  }){
    title
  }
  
}
      `,variables:{
        title:this.title,
        link:this.downloadURL,
        desc:this.desc,
        label:(this.restricted) ? "Mature":"Not Mature",
        privacy:this.privacy,
        uploader_id:localStorage.getItem("uid"),
        type:this.videoType,
        thumbnail:this.img,
        location:"Indonesia",//localStorage.getItem("loc")
        category:this.ctg.name
      }
    }).subscribe(res => {
      console.log("Uploaded!")
    })
  }
  test(event){
    console.log(event)
  }
}
