import { Component, OnInit, Query } from '@angular/core';
import {Apollo} from 'apollo-angular';
import gql from 'graphql-tag';
import {faEllipsisV} from '@fortawesome/free-solid-svg-icons';
import { Router } from '@angular/router';
@Component({
  selector: 'app-video-thumb',
  templateUrl: './video-thumb.component.html',
  styleUrls: ['./video-thumb.component.scss']
})
export class VideoThumbComponent implements OnInit {

  constructor(private apollo: Apollo, private router: Router) { }
  hovered=false
  thumbnail : any
  pict : any
  title: any
  name: any
  date :Date
  showDate: string
  showView: string
  views :number
  videoid = 53
  elipsis = faEllipsisV
  curLink : any
  ngOnInit(): void {
    this.curLink = this.router.url
   this.apollo.watchQuery({
     query:gql `
     query getvideo($videoid: ID! =0){
  getVideoDetail(id:$videoid){
    video{
      videos
      uploader_id
			thumbnail
      title
      view_count
      upload_date
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
     `,variables:{
       videoid:this.videoid
     }
   }).valueChanges.subscribe(x =>{
    // this.thumbnail = x.data.getVideoDetail.user.picture
    this.thumbnail = x.data.getVideoDetail.video.thumbnail
    this.pict = x.data.getVideoDetail.user.picture
    this.title = x.data.getVideoDetail.video.title
    this.name = x.data.getVideoDetail.user.name
    this.date = new Date(x.data.getVideoDetail.video.upload_date)
    this.views = (x.data.getVideoDetail.view_count) ? x.data.getVideoDetail.view_count : 0
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

    if (this.views > 999 && this.views < 1000000){
      var temp = this.views/1000
      this.showView = temp.toFixed(4) + "k Views"
    }else if (this.views >= 1000000)
    {
      var temp = this.views/1000000
      
      this.showView = temp.toFixed(4) + "M Views"
    }
    else{
      this.showView = this.views + " views"
    }

    //  console.log(this.imgUrl)
   })

    
  }
  toggleElipsis(){
    this.hovered = !this.hovered
  }

}
