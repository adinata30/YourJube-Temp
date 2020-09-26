
export type User={
    id:string;
    name:string;
    picture:any;
    membership_status:string;
    billing_type:string;
    membership_exp_date:Date;
    join_date:Date;
    email:string;
    Password:string;
}
export type Category={
    id:number;
    name:string;
}

export type VideoDetail={
    // video: Video;
    user: User;
    likeCount: number;
    dislikeCount: number;
    
}

export type allVideoDetails={
    allVideoDetails: [VideoDetail];
}

