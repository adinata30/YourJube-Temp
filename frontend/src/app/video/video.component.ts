import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {Apollo} from 'apollo-angular';
import gql from 'graphql-tag';
import {map} from 'rxjs/operators';
import {Observable} from'rxjs';
import {faThumbsUp,faThumbsDown,faFileDownload,faSubscript} from '@fortawesome/free-solid-svg-icons'
import {trigger,state,style,transition,animate} from '@angular/animations';
import {User} from '../user';
import {VgDashDirective} from'@videogular/ngx-videogular/streaming'
import {VgApiService, BitrateOptions} from '@videogular/ngx-videogular/core';
import {AngularFireStorage} from '@angular/fire/storage'
const getVideoQuery = gql `
    query getvideo($videoid: ID! =0){
    getVideoDetail(id:$videoid){
    video{
      videos
      uploader_id
      thumbnail
      title
      view_count
      upload_date
      descriptions
    }
    user{
      id
      picture
      name
    }
    likeCount
    dislikeCount
    }
    }
`

const checkForLikeOrDislike = gql`
 query checkForLike($uid:ID!, $vid:ID!){
      isVideoLikedorDislikedByUser(uid:$uid,vid:$vid)
          
    }
`

@Component({
  selector: 'app-video',
  templateUrl: './video.component.html',
  styleUrls: ['./video.component.scss'],
  animations:[
    trigger('likeOrDislike',[
      state('active',style({
        color:"lightblue",
      })),
      state('inactive',style({
        color:"grey",
      }))
    ])
  ]
})
export class VideoComponent implements OnInit {
  xhr = new XMLHttpRequest();
  dislikeActive = false
  likeActive = false
  isVideoReady = false
  test : Observable<any>
  curLink : string
  uploader_id: number
  uid: number
  pict : any
  title: any
  name: any
  date :Date
  showDate: string
  showView: string
  views :number
  downloadUrl :string
  likeCount : string
  dislikeCount : string
  description : string
  faSubscript = faSubscript
  faFileDownload= faFileDownload
  faThumbsUp = faThumbsUp
  faThumbsDown = faThumbsDown
  api: VgApiService
  videoId: string
  constructor(private router:Router, private apollo:Apollo, private isloggedIn: User, private storage : AngularFireStorage){
    isloggedIn.loggedIn
  }
  ngOnInit(){
    this.videoId = (this.router.url.split('/')[2])
    this.apollo.mutate({
      mutation: gql `
      mutation watch($vid:Int!){
	increaseVideoView(input:{id:$vid})
  {
    title
  }
}
      `,variables:{
        vid:this.videoId
      }
    }).subscribe(x =>{
      
    })
    this.apollo.watchQuery({
      query:getVideoQuery,variables:{
        videoid:this.videoId
      }
    }).valueChanges.
    subscribe(x =>{
      
     // this.thumbnail = x.data.getVideoDetail.user.picture
     this.uploader_id = x.data.getVideoDetail.video.uploader_id
     
     this.downloadUrl = x.data.getVideoDetail.video.videos
     this.pict = x.data.getVideoDetail.user.picture
     this.title = x.data.getVideoDetail.video.title
     this.name = x.data.getVideoDetail.user.name
     this.date = new Date(x.data.getVideoDetail.video.upload_date)
     this.description = x.data.getVideoDetail.video.descriptions
    //  this.views = (x.data.getVideoDetail.view_count) ? x.data.getVideoDetail.view_count : 0
    this.views = x.data.getVideoDetail.video.view_count
    
     const diffYear = new Date().getFullYear() - this.date.getFullYear()
     if (diffYear > 0) {
       this.showDate = diffYear + " Year Ago"
     } else{
       const diffMonth = new Date().getMonth() - this.date.getMonth()
       if (diffMonth > 0){
         this.showDate = diffMonth + " Month Ago"
       }
       else{
         const diffDay = new Date().getDate() - this.date.getDate()
         if (diffDay > 0){
           this.showDate = diffDay + " Day Ago"
         }else{
           const diffHour = new Date().getHours() - this.date.getHours()
           this.showDate = diffHour + " Hour Ago"
         }
       }
     }
 
    if (this.views >= 1000000)
    {
      var temp = this.views/1000000
       
      this.showView = temp.toFixed(4) + "M Views"
    }
    else{
      this.showView = this.views + " views"
    }
    const like = x.data.getVideoDetail.likeCount
    const dislike = x.data.getVideoDetail.dislikeCount
    if(like > 999 && like < 1000000){
      this.likeCount  = like/1000 + "K"
    }
    else if(like >= 1000000){
      this.likeCount = like/1000000 +"M"
    }
    else this.likeCount = like
    if(dislike > 999 && dislike < 1000000){
      this.dislikeCount  = dislike/1000 + "K"
    }
    else if(dislike >= 1000000){
      this.dislikeCount = dislike/1000000 +"M"
    }
    else this.dislikeCount = dislike
    
    this.isVideoReady=true
     //  console.log(this.imgUrl)
    })

    
    if(localStorage.getItem("uid")){
      this.uid = parseInt(localStorage.getItem("uid"))
      this.apollo.watchQuery({
        query: checkForLikeOrDislike,
        variables:{
          uid: localStorage.getItem("uid"),
          vid: this.videoId
        }
      }).valueChanges.subscribe(x => {
        this.likeActive = x.data.isVideoLikedorDislikedByUser[0]
        this.dislikeActive = x.data.isVideoLikedorDislikedByUser[1]
        console.log(this.likeActive)
        console.log(this.dislikeActive)

      })
      
    }
    

    
  }

  like(){
      if (this.isloggedIn.loggedIn ==false)return;
      this.apollo.mutate({
        mutation: gql `
        mutation likeVideo($id : Int!,  $uid :Int!){
          likeVideo(input:{
            id:$id,
            user_id:$uid
          })
          
        }
        `,variables:{
          id: this.videoId,
          uid: localStorage.getItem("uid")
        },refetchQueries:[{
          query:getVideoQuery
          ,variables:{
            videoid:this.videoId
          }
        }
      ]
      }).subscribe(x => {
        if (x.data.likeVideo == null) {
          console.log("Null")
        } else {
          this.likeActive = x.data.likeVideo[0]
          this.dislikeActive = x.data.likeVideo[1]
        }
      })
      
    }

dislike(){
      if (this.isloggedIn.loggedIn ==false)return;
      
        this.apollo.mutate({
          mutation: gql `
          mutation dislikeVideo($id : Int!,  $uid :Int!){
            dislikeVideo(input:{
              id:$id,
              user_id:$uid
            })
            
          }
          `,variables:{
            id: this.videoId,
            uid: localStorage.getItem("uid")
          },refetchQueries:[{
            query:getVideoQuery
            ,variables:{
              videoid:this.videoId
            }
          }]
        }).subscribe(x => {
          if (x.data.dislikeVideo == null) {
            console.log("Null")
          } else {
            this.likeActive = x.data.dislikeVideo[0]
            this.dislikeActive = x.data.dislikeVideo[1]
          }
        })
        // this.checkforLikeorDislike() 
    }
    // const checkForLikeorDislike(){
    //   this.apollo.watchQuery({
    //     query: gql `
    //     query checkForLike($uid:ID!, $vid:ID!){
    //   isVideoLikedorDislikedByUser(uid:$uid,vid:$vid)
          
    //   }
    //     `,variables:{
    //       uid:localStorage.getItem("uid"),
    //       vid:this.videoId
    //     }
    //   }).valueChanges.subscribe(x =>{
    //     console.log(x.data)
    //     this.likeActive = x.data.isVideoLikedorDislikedByUser[0] 
    //     this.dislikeActive = x.data.isVideoLikedorDislikedByUser[1] 
    //   })
    // }
   b(){
  var xhr2 = new XMLHttpRequest()
  var temp = this.title
  xhr2.responseType = 'blob';
  xhr2.onload = function download(){
    var downloadUrl = URL.createObjectURL(xhr2.response);
    var a = document.createElement("a");
    document.body.appendChild(a);
  
    a.href = downloadUrl;
    a.download = temp;
    a.click();
  }
  xhr2.onerror = function () {
    console.log("** An error occurred during the transaction");
  };
  console.log(xhr2.readyState)
  xhr2.open('GET', this.downloadUrl);
  xhr2.send();
  
  // this.xhr.onreadystatechange = this.updateProgress
// progress on transfers from the server to the client (downloads)
}
   onPlayerReady(api:VgApiService) {
    this.api = api;
  }
  dashBitrates : BitrateOptions[]
  

}

