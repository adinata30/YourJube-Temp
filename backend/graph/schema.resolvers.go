package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gql/graph/generated"
	"gql/graph/model"
	"time"

	pg "github.com/go-pg/pg/v10"
)

func (r *mutationResolver) CreateNewComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		UserID:  input.UserID,
		Content: input.Content,
	}
	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *mutationResolver) RemoveComment(ctx context.Context, commentID int) (bool, error) {
	var comment model.Comment

	err := r.DB.Model(&comment).Where("id = ?", commentID).First()

	if err != nil {
		return false, err
	}

	_, deleteError := r.DB.Model(&comment).Where("id = ?", commentID).Delete()

	if deleteError != nil {
		return false, deleteError
	}

	return true, nil
}

func (r *mutationResolver) EditComment(ctx context.Context, commentID int, input *model.NewComment) (*model.Comment, error) {
	var comment model.Comment

	err := r.DB.Model(&comment).Where("id = ?", input.UserID).First()

	if err != nil {
		return nil, err
	}

	comment.Content = input.Content

	_, updateErr := r.DB.Model(&comment).Where("id = ?", input.UserID).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &comment, nil
}

func (r *mutationResolver) LikeComment(ctx context.Context, input model.DoLikeOrDislike) (bool, error) {
	var comment model.Comment

	//coba delete dlu kalau ada dislike
	r.DB.Model(&comment).Where("target_id = ? AND user_id = ? AND target_type = 'Comment' AND Actions = 'Dislike'", input.ID, input.UserID).Delete()

	//kalau sudah didelete, create like baru
	newLike := model.LikeOrDislike{
		UserID:     input.UserID,
		TargetID:   input.ID,
		TargetType: "Comment",
		Actions:    "Like",
	}

	///kalau gagal insert gara" uda ada like, delete like dia (anggepnya unlike tapi bukan dislike)
	_, insertErr := r.DB.Model(&newLike).Insert()
	if insertErr != nil {
		r.DB.Model(&comment).Where("target_id = ? AND user_id = ? AND target_type = 'Comment' AND Actions = 'Like'", input.ID, input.UserID).Delete()
		return false, nil
	}
	return true, nil

	// panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DislikeComment(ctx context.Context, input model.DoLikeOrDislike) (bool, error) {
	var comment model.LikeOrDislike

	//coba delete dlu kalau ada like
	_, deleteErr := r.DB.Model(&comment).Where("user_id = ? AND target_id = ? AND target_type = 'Comment' AND Actions = 'Like'", input.UserID, input.ID).Delete()
	if deleteErr != nil {
		return false, deleteErr
	}
	//kalau sudah didelete, create dislike baru
	newLike := model.LikeOrDislike{
		UserID:     input.UserID,
		TargetID:   input.ID,
		TargetType: "Comment",
		Actions:    "Dislike",
	}

	///kalau gagal insert gara" uda ada dislike, delete like dia (anggepnya un-dislike tapi bukan like)
	_, insertErr := r.DB.Model(&newLike).Insert()
	if insertErr != nil {
		r.DB.Model(&comment).Where("target_id = ? AND user_id = ? AND target_type = 'Comment' AND Actions = 'Dislike'", input.ID, input.UserID).Delete()
		return false, insertErr
	}
	return true, nil
}

func (r *mutationResolver) ReplyComment(ctx context.Context, commentID int, input *model.NewComment) (*model.Comment, error) {
	var comment = model.Comment{
		UserID:    input.UserID,
		RepliedID: commentID,
		Content:   input.Content,
	}
	_, err := r.DB.Model(&comment).Insert()
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *mutationResolver) CreateNewPlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	var playlist = model.Playlist{
		UserID:      input.UserID,
		Name:        input.Name,
		URL:         input.URL,
		Type:        input.Type,
		Description: input.Description,
	}
	_, err := r.DB.Model(&playlist).Insert()
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *mutationResolver) AddToPlaylist(ctx context.Context, input *model.AddToPlaylist) (bool, error) {
	var detail = model.PlaylistContent{
		PlaylistID: input.ID,
		VideoID:    input.Videos,
	}
	_, err := r.DB.Model(&detail).Insert()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) RemoveFromPlaylist(ctx context.Context, input *model.AddToPlaylist) (bool, error) {
	var playlist model.PlaylistContent

	_, err := r.DB.Model(&playlist).Where("playlist_id = ? AND video_id = ? ", input.ID, input.Videos).Delete()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) EditPlaylist(ctx context.Context, playlistID int, input *model.NewPlaylist) (*model.Playlist, error) {
	var playlist model.Playlist
	findErr := r.DB.Model(&playlist).Where("id = ?", playlistID).First()
	if findErr != nil {
		return nil, findErr
	}
	playlist.Name = input.Name
	playlist.URL = input.URL
	playlist.Type = input.Type
	playlist.Description = input.Description
	_, updateErr := r.DB.Model(&playlist).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &playlist, nil
}

func (r *mutationResolver) IncreasePlaylistView(ctx context.Context, input *model.IncreaseView) (*model.Playlist, error) {
	var playlist model.Playlist
	findErr := r.DB.Model(&playlist).Where("id = ?", input.ID).First()
	if findErr != nil {
		return nil, findErr
	}
	playlist.TotalViews = playlist.TotalViews + 1

	_, updateErr := r.DB.Model(&playlist).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &playlist, nil
}

func (r *mutationResolver) AddToPopularList(ctx context.Context, input *model.AddToPlaylist) (bool, error) {
	var popularContent = model.PopularContent{
		PopularID: input.ID,
		VideoID:   input.Videos,
	}

	_, insertErr := r.DB.Model(&popularContent).Insert()
	if insertErr != nil {
		return false, insertErr
	}
	return true, nil
}

func (r *mutationResolver) CreateNewSession(ctx context.Context, input *model.CreateNewSession) (*model.Session, error) {
	var session = model.Session{
		IP:       input.IP,
		Location: input.Location,
		Type:     input.Type,
	}
	_, insertErr := r.DB.Model(&session).Insert()
	if insertErr != nil {
		return nil, insertErr
	}
	return &session, nil
}

func (r *mutationResolver) RemoveSession(ctx context.Context, sessionID int) (bool, error) {
	var session model.Session

	err := r.DB.Model(&session).Where("id = ?", sessionID).First()
	if err != nil {
		return false, err
	}
	_, deleteErr := r.DB.Model(&session).Where("id = ?", sessionID).Delete()
	if deleteErr != nil {
		return false, deleteErr
	}
	return true, nil
}

func (r *mutationResolver) RestrictSession(ctx context.Context, sessionID int) (bool, error) {
	var session model.Session
	err := r.DB.Model(&session).Where("ID = ?", sessionID).First()
	if err != nil {
		return false, err
	}
	session.Restriction = true
	_, updateErr := r.DB.Model(&session).Update()
	if updateErr != nil {
		return false, updateErr
	}
	return true, nil
}

func (r *mutationResolver) ChangeSessionType(ctx context.Context, sessionID int, typeArg string) (*model.Session, error) {
	var session model.Session
	err := r.DB.Model(&session).Where("ID = ?", sessionID).First()
	if err != nil {
		return nil, err
	}
	session.Type = typeArg
	_, updateErr := r.DB.Model(&session).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &session, nil
}

func (r *mutationResolver) AddQueueToSession(ctx context.Context, input *model.AddToQueue) (bool, error) {
	var sessQueue = model.SessionQueue{
		SessionID: input.SessionID,
		VideoID:   input.Queue,
	}
	_, insertErr := r.DB.Model(&sessQueue).Insert()
	if insertErr != nil {
		return false, insertErr
	}
	return true, nil
}

func (r *mutationResolver) RemoveQueueFromSession(ctx context.Context, input *model.AddToQueue) (bool, error) {
	var sessQ model.Session

	err := r.DB.Model(&sessQ).Where("session_id = ? AND video_id = ?", input.SessionID, input.Queue).First()
	if err != nil {
		return false, err
	}
	_, deleteErr := r.DB.Model(&sessQ).Where("session_id = ? AND video_id = ?", input.SessionID, input.Queue).Delete()
	if deleteErr != nil {
		return false, deleteErr
	}
	return true, nil
}

func (r *mutationResolver) Subscribe(ctx context.Context, input *model.NewSubscription) (bool, error) {
	var subscriptions = model.Subscriptions{
		UserID:       input.UserID,
		SubscribedID: input.SubscribedID,
	}
	_, insertErr := r.DB.Model(&subscriptions).Insert()
	if insertErr != nil {
		return false, insertErr
	}
	return true, nil
}

func (r *mutationResolver) Unsubscribe(ctx context.Context, input *model.NewSubscription) (bool, error) {
	var subscriptions model.Subscriptions
	err := r.DB.Model(&subscriptions).Where("user_id = ? AND subscribed_id = ? ", input.UserID, input.SubscribedID).First()
	if err != nil {
		return false, err
	}
	_, deleteErr := r.DB.Model(&subscriptions).Where("user_id = ? AND subscribed_id = ? ", input.UserID, input.SubscribedID).Delete()
	if deleteErr != nil {
		return false, deleteErr
	}
	return true, nil
}

func (r *mutationResolver) AddNotification(ctx context.Context, input *model.NewSubscription) (bool, error) {
	var subscriptions model.Subscriptions
	err := r.DB.Model(&subscriptions).Where("user_id = ? AND subscribed_id = ? ", input.UserID, input.SubscribedID).First()
	if err != nil {
		return false, err
	}
	subscriptions.Notification = true
	_, updateErr := r.DB.Model(&subscriptions).Update()
	if updateErr != nil {
		return false, updateErr
	}
	return true, nil
}

func (r *mutationResolver) UnNotify(ctx context.Context, input *model.NewSubscription) (bool, error) {
	var subscriptions model.Subscriptions
	err := r.DB.Model(&subscriptions).Where("user_id = ? AND subscribed_id = ? ", input.UserID, input.SubscribedID).First()
	if err != nil {
		return false, err
	}
	subscriptions.Notification = false
	_, updateErr := r.DB.Model(&subscriptions).Update()
	if updateErr != nil {
		return false, updateErr
	}
	return true, nil
}

func (r *mutationResolver) NewUploadQueue(ctx context.Context, input *model.NewUploadQueue) (*model.UploadQueue, error) {
	var uploadQueue = model.UploadQueue{
		VideoID:    input.VideoID,
		UploadDate: input.UploadDate,
	}
	var video model.Video
	err := r.DB.Model(&video).Where("id = ?", input.VideoID).First()
	if err != nil {
		return nil, err
	}
	video.UploadDate = input.UploadDate
	r.DB.Model(&video).Update()
	_, insertErr := r.DB.Model(&uploadQueue).Insert()
	if insertErr != nil {
		return nil, insertErr
	}
	return &uploadQueue, nil
}

func (r *mutationResolver) RemoveUploadQueue(ctx context.Context, uploadQueueID int) (bool, error) {
	var uploadQueue model.UploadQueue
	err := r.DB.Model(&uploadQueue).Where("id = ?", uploadQueueID).First()
	if err != nil {
		return false, err
	}
	_, deleteErr := r.DB.Model(&uploadQueueID).Where("id = ?", uploadQueueID).Delete()
	if deleteErr != nil {
		return false, deleteErr
	}
	return true, nil
}

func (r *mutationResolver) CreateNewUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	var temp model.User
	r.DB.Model(&temp).Where("email = ?", input.Email).First()
	fmt.Println("Name = " + temp.Name)
	if temp.Name != "" {
		return &temp, nil
	}
	user := model.User{
		Name:    input.Name,
		Email:   input.Email,
		Picture: input.Picture,
	}
	_, err := r.DB.Model(&user).Insert()

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *mutationResolver) ChangeProfilePicture(ctx context.Context, input *model.ChangePicture) (*model.User, error) {
	var user model.User

	err := r.DB.Model(&user).Where("id = ?", input.ID).First()

	if err != nil {
		return nil, err
	}
	user.Picture = input.Pict

	_, updateErr := r.DB.Model(&user).Where("id = ?", input.ID).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &user, nil
}

func (r *mutationResolver) ChangeMembership(ctx context.Context, input *model.NewMembership) (*model.User, error) {
	var user model.User
	err := r.DB.Model(&user).Where("id = ?", input.ID).First()
	if err != nil {
		return nil, err
	}
	user.MembershipStatus = "Premium"
	user.BillingType = input.BillingType
	if input.BillingType == "Monthly" {
		user.MembershipExpDate = time.Now().AddDate(0, 1, 0).String()
	} else {
		user.MembershipExpDate = time.Now().AddDate(1, 0, 0).String() //yearly
	}
	_, updateErr := r.DB.Model(&user).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &user, nil
}

func (r *mutationResolver) CancelMembership(ctx context.Context, userID int) (bool, error) {
	var user model.User
	err := r.DB.Model(&user).Where("id = ?", userID).First()
	if err != nil {
		return false, err
	}
	user.MembershipStatus = "Free"
	_, updateErr := r.DB.Model(&user).Update()
	if updateErr != nil {
		return false, updateErr
	}
	return true, nil
}

func (r *mutationResolver) CreateNewVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	var video = model.Video{
		Videos:       input.Video,
		Title:        input.Title,
		Descriptions: input.Descriptions,
		UploaderID:   input.UploaderID,
		Label:        input.Label,
		Privacy:      input.Privacy,
		Type:         input.Type,
		Thumbnail:    input.Thumbnail,
		Location:     input.Location,
		Category:     input.Category,
	}

	_, insertErr := r.DB.Model(&video).Insert()
	if insertErr != nil {
		return nil, insertErr
	}
	return &video, nil
}

func (r *mutationResolver) ChangeVideo(ctx context.Context, input *model.ChangeVideoDetail) (*model.Video, error) {
	var videoDetail model.Video
	err := r.DB.Model(&videoDetail).Where("id = ?", input.ID).First()
	if err != nil {
		return nil, err
	}
	videoDetail.Title = input.Title
	videoDetail.Descriptions = input.Descriptions
	_, updateErr := r.DB.Model(&videoDetail).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &videoDetail, nil
}

func (r *mutationResolver) LikeVideo(ctx context.Context, input model.DoLikeOrDislike) ([]bool, error) {
	var comment model.LikeOrDislike

	//coba delete dlu kalau ada like
	_, deleteErr := r.DB.Model(&comment).Where("user_id = ? AND target_id = ? AND target_type = 'Video' AND Actions = 'Dislike'", input.UserID, input.ID).Delete()
	if deleteErr != nil {
		return nil, nil
	}
	//kalau sudah didelete, create dislike baru
	newLike := model.LikeOrDislike{
		UserID:     input.UserID,
		TargetID:   input.ID,
		TargetType: "Video",
		Actions:    "Like",
	}

	///kalau gagal insert gara" uda ada dislike, delete like dia (anggepnya un-dislike tapi bukan like)
	_, insertErr := r.DB.Model(&newLike).Insert()
	if insertErr != nil {
		r.DB.Model(&comment).Where("target_id = ? AND user_id = ? AND target_type = 'Video' AND Actions = 'Like'", input.ID, input.UserID).Delete()
		return nil, nil
	}
	var lod model.LikeOrDislike
	likeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Like' AND user_id = ? ", input.ID, input.UserID).Count()
	dislikeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Dislike'AND user_id = ? ", input.ID, input.UserID).Count()
	var isLiked = (likeCount == 1)
	var isDisliked = (dislikeCount == 1)
	toReturn := []bool{isLiked, isDisliked}
	// toReturn[0] = isLiked
	// toReturn[1] = isDisliked
	return toReturn, nil
}

func (r *mutationResolver) DislikeVideo(ctx context.Context, input model.DoLikeOrDislike) ([]bool, error) {
	var comment model.LikeOrDislike

	//coba delete dlu kalau ada like
	_, deleteErr := r.DB.Model(&comment).Where("user_id = ? AND target_id = ? AND target_type = 'Video' AND Actions = 'Like'", input.UserID, input.ID).Delete()
	if deleteErr != nil {
		return nil, nil
	}
	//kalau sudah didelete, create dislike baru
	newLike := model.LikeOrDislike{
		UserID:     input.UserID,
		TargetID:   input.ID,
		TargetType: "Video",
		Actions:    "Dislike",
	}

	///kalau gagal insert gara" uda ada dislike, delete like dia (anggepnya un-dislike tapi bukan like)
	_, insertErr := r.DB.Model(&newLike).Insert()
	if insertErr != nil {
		r.DB.Model(&comment).Where("target_id = ? AND user_id = ? AND target_type = 'Video' AND Actions = 'Dislike'", input.ID, input.UserID).Delete()
		return nil, nil
	}
	var lod model.LikeOrDislike
	likeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Like' AND user_id = ? ", input.ID, input.UserID).Count()
	dislikeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Dislike'AND user_id = ? ", input.ID, input.UserID).Count()
	var isLiked = (likeCount == 1)
	var isDisliked = (dislikeCount == 1)
	toReturn := []bool{isLiked, isDisliked}
	// toReturn[0] = isLiked
	// toReturn[1] = isDisliked
	return toReturn, nil
}

func (r *mutationResolver) IncreaseVideoView(ctx context.Context, input *model.IncreaseView) (*model.Video, error) {
	var video model.Video
	err := r.DB.Model(&video).Where("id = ?", input.ID).First()
	if err != nil {
		return nil, err
	}
	video.ViewCount = video.ViewCount + 1
	_, updateErr := r.DB.Model(&video).Where("id = ?", input.ID).Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &video, nil
}

func (r *mutationResolver) AddNewPopularVideo(ctx context.Context, input *model.NewPopularVideo) (bool, error) {
	var popularVideo = model.PopularContent{
		PopularID: input.PopularID,
		VideoID:   input.VideoID,
	}
	_, insertErr := r.DB.Model(&popularVideo).Insert()
	if insertErr != nil {
		return false, insertErr
	}
	return true, nil
}

func (r *mutationResolver) AddNewCategoryVideo(ctx context.Context, input *model.NewVideoCategory) (bool, error) {
	var categoryVideo = model.VideoCategory{
		VideoID:    input.VideoID,
		CategoryID: input.CategoryID,
	}
	_, insertErr := r.DB.Model(&categoryVideo).Insert()
	if insertErr != nil {
		return false, insertErr
	}
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Model(&users).Select()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Select()

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Select()

	if err != nil {
		return nil, err
	}
	return playlists, nil
}

func (r *queryResolver) Populars(ctx context.Context) ([]*model.Popular, error) {
	var populars []*model.Popular

	err := r.DB.Model(&populars).Select()

	if err != nil {
		return nil, err
	}
	return populars, nil
}

func (r *queryResolver) Sessions(ctx context.Context) ([]*model.Session, error) {
	var sessions []*model.Session

	err := r.DB.Model(&sessions).Select()

	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *queryResolver) Subscriptions(ctx context.Context) ([]*model.Subscriptions, error) {
	var subscriptions []*model.Subscriptions

	err := r.DB.Model(&subscriptions).Select()

	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Select()

	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *queryResolver) PlaylistContents(ctx context.Context) ([]*model.PlaylistContent, error) {
	var playlistContents []*model.PlaylistContent

	err := r.DB.Model(&playlistContents).Select()

	if err != nil {
		return nil, err
	}
	return playlistContents, nil
}

func (r *queryResolver) PopularContents(ctx context.Context) ([]*model.PopularContent, error) {
	var popularContents []*model.PopularContent

	err := r.DB.Model(&popularContents).Select()

	if err != nil {
		return nil, err
	}
	return popularContents, nil
}

func (r *queryResolver) SessionQueues(ctx context.Context) ([]*model.SessionQueue, error) {
	var sessionQueues []*model.SessionQueue

	err := r.DB.Model(&sessionQueues).Select()

	if err != nil {
		return nil, err
	}
	return sessionQueues, nil
}

func (r *queryResolver) VideoCategories(ctx context.Context) ([]*model.VideoCategory, error) {
	var videoCategories []*model.VideoCategory

	err := r.DB.Model(&videoCategories).Select()

	if err != nil {
		return nil, err
	}
	return videoCategories, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	var categories []*model.Category

	err := r.DB.Model(&categories).Select()

	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *queryResolver) LikesOrDislikes(ctx context.Context) ([]*model.LikeOrDislike, error) {
	var likesOrDislikes []*model.LikeOrDislike

	err := r.DB.Model(&likesOrDislikes).Select()

	if err != nil {
		return nil, err
	}
	return likesOrDislikes, nil
}

func (r *queryResolver) UploadQueues(ctx context.Context) ([]*model.UploadQueue, error) {
	var uploadQueues []*model.UploadQueue

	err := r.DB.Model(&uploadQueues).Select()

	if err != nil {
		return nil, err
	}
	return uploadQueues, nil
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.DB.Model(&user).Where("id = ? ", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *queryResolver) GetPlaylistVideos(ctx context.Context, id string) (*model.PlaylistVideoDetail, error) {
	var playlist model.Playlist
	// fmt.Println("String = %s", id)
	err := r.DB.Model(&playlist).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	var videoID []string
	var added_date []string
	var pc model.PlaylistContent
	// err = r.DB.Model(&pc).ColumnExpr("video_id::varchar", "added_date::varchar").Where("playlist_id = ?", id).Select(&videoID, &added_date)
	err = r.DB.Model(&pc).ColumnExpr("added_date::varchar").Where("playlist_id = ?", id).Select(&added_date)
	err = r.DB.Model(&pc).ColumnExpr("video_id::varchar").Where("playlist_id = ?", id).Select(&videoID)
	if err != nil {
		return nil, err
	}

	var video []*model.Video
	err = r.DB.Model(&video).Where("id IN (?)", pg.In(videoID)).Select()

	var addedVideo = model.AddedVideo{
		Video:     video,
		AddedDate: added_date,
	}
	var pvd = model.PlaylistVideoDetail{
		Playlist: &playlist,
		Videos:   &addedVideo,
	}
	return &pvd, nil
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPopularVideos(ctx context.Context, id string) (*model.PopularVideoDetail, error) {
	var popular model.Popular
	// fmt.Println("String = %s", id)
	err := r.DB.Model(&popular).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	var videoID []string
	var pc model.PlaylistContent
	err = r.DB.Model(&pc).Column("video_id").Where("popular_id = ?", id).Select(&videoID)

	var video []*model.Video
	err = r.DB.Model(&video).Where("id IN (?)", pg.In(videoID)).Select()

	var pvd = model.PopularVideoDetail{
		Popular: &popular,
		Videos:  video,
	}
	return &pvd, nil
}

func (r *queryResolver) GetQueueVideos(ctx context.Context, id string) (*model.QueueVideoDetail, error) {
	var session model.Session
	// fmt.Println("String = %s", id)
	err := r.DB.Model(&session).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	var videoID []string
	var pc model.SessionQueue
	err = r.DB.Model(&pc).Column("video_id").Where("session_id = ?", id).Select(&videoID)

	var video []*model.Video
	err = r.DB.Model(&video).Where("id IN (?)", pg.In(videoID)).Select()

	var pvd = model.QueueVideoDetail{
		Session: &session,
		Videos:  video,
	}
	return &pvd, nil
}

func (r *queryResolver) GetUserVideos(ctx context.Context, id string) (*model.UserVideoDetail, error) {
	var user model.User
	// fmt.Println("String = %s", id)
	err := r.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	var video []*model.Video
	err = r.DB.Model(&video).Where("uploader_id = ? ", id).Select()

	var uvd = model.UserVideoDetail{
		User:   &user,
		Videos: video,
	}
	return &uvd, nil
}

func (r *queryResolver) GetVideoDetail(ctx context.Context, id string) (*model.VideoDetail, error) {
	var video model.Video
	r.DB.Model(&video).Where("id = ?", id).First()
	var user model.User
	r.DB.Model(&user).Where("id = ?", video.UploaderID).First()
	var lod model.LikeOrDislike
	likeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Like'", id).Count()
	dislikeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Dislike'", id).Count()
	var videoDetail = model.VideoDetail{
		Video:        &video,
		User:         &user,
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
	}
	return &videoDetail, nil
}

func (r *queryResolver) GetAllVideoDetail(ctx context.Context) ([]*model.VideoDetail, error) {
	var videoDetails []*model.VideoDetail
	var videos []*model.Video
	r.DB.Model(&videos).Select()
	for _, vid := range videos {
		var user model.User
		r.DB.Model(&user).Where("id = ?", vid.UploaderID).First()
		var lod model.LikeOrDislike
		likeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Like'", vid.ID).Count()
		dislikeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Disike'", vid.ID).Count()
		details := model.VideoDetail{
			Video:        vid,
			User:         &user,
			LikeCount:    likeCount,
			DislikeCount: dislikeCount,
		}
		videoDetails = append(videoDetails, &details)
	}

	// fmt.Println(vid)
	return videoDetails, nil
}

func (r *queryResolver) IsVideoLikedorDislikedByUser(ctx context.Context, uid string, vid string) ([]bool, error) {
	var lod model.LikeOrDislike
	likeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Like' AND user_id = ? ", vid, uid).Count()
	dislikeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Video' and actions = 'Dislike'AND user_id = ? ", vid, uid).Count()
	var isLiked = (likeCount == 1)
	var isDisliked = (dislikeCount == 1)
	toReturn := []bool{isLiked, isDisliked}
	// toReturn[0] = isLiked
	// toReturn[1] = isDisliked
	return toReturn, nil
}

func (r *queryResolver) IsCommentLikedorDislikeByUser(ctx context.Context, uid string, cid string) ([]bool, error) {
	var lod model.LikeOrDislike
	likeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Comment' and actions = 'Like' AND user_id = ? ", cid, uid).Count()
	dislikeCount, _ := r.DB.Model(&lod).Where("target_id = ? and target_type = 'Comment' and actions = 'Dislike'AND user_id = ? ", cid, uid).Count()
	var isLiked = (likeCount == 1)
	var isDisliked = (dislikeCount == 1)
	toReturn := []bool{isLiked, isDisliked}
	// toReturn[0] = isLiked
	// toReturn[1] = isDisliked
	return toReturn, nil
}

func (r *queryResolver) GetTotalVideoLikes(ctx context.Context, id string) (int, error) {
	var lod model.LikeOrDislike
	totalLikes, err := r.DB.Model(&lod).Where("target_id = ? AND actions = 'Like' AND target_type = 'Video'", id).Count()
	if err != nil {
		return -1, err
	}
	return totalLikes, nil
}

func (r *queryResolver) GetTotalVideoDislikes(ctx context.Context, id string) (int, error) {
	var lod model.LikeOrDislike
	totalDislike, err := r.DB.Model(&lod).Where("target_id = ? AND actions = 'Dislike' AND target_type = 'Video'", id).Count()
	if err != nil {
		return -1, err
	}
	return totalDislike, nil
}

func (r *queryResolver) GetTotalCommentLikes(ctx context.Context, id string) (int, error) {
	var lod model.LikeOrDislike
	totalLikes, err := r.DB.Model(&lod).Where("target_id = ? AND actions = 'Like' AND target_type = 'Comment'", id).Count()
	if err != nil {
		return -1, err
	}
	return totalLikes, nil
}

func (r *queryResolver) GetTotalCommentDislikes(ctx context.Context, id string) (int, error) {
	var lod model.LikeOrDislike
	totalDislike, err := r.DB.Model(&lod).Where("target_id = ? AND actions = 'Dislike' AND target_type = 'Comment'", id).Count()
	if err != nil {
		return -1, err
	}
	return totalDislike, nil
}

func (r *queryResolver) GetUserSubscriptions(ctx context.Context, id string) (*model.UserSubscriptionDetail, error) {
	var user model.User
	err := r.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	var subscriptions []*model.Subscriptions
	err = r.DB.Model(&subscriptions).Where("user_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	var usd = model.UserSubscriptionDetail{
		User:          &user,
		Subscriptions: subscriptions,
	}
	return &usd, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func getIndex(arrayToSearch []int, key int) int {
	for idx, val := range arrayToSearch {
		if val == key {
			return idx
		}
	}
	fmt.Println("hello")
	return -1
}
