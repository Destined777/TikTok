package service

import (
	"TikTok/dao"
	"TikTok/http_param"
	"TikTok/model"
	"TikTok/util"
	"errors"
)

func CreateVideo(title string, ID int64, name string) (err error) {
	video := model.Video{
		UserId:       	ID,
		PlayUrl:      	"http://10.21.191.27:8080/static/"+name,
		CoverUrl:     	"https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavouriteNum: 	0,
		CommentNum:   	0,
		Title:        	title,
		CreatedAt: 		util.GetTimeStamp(),
	}
	err = dao.CreateVideo(video)
	return
}

func GetVideos(time1 int64, token string) (res []http_param.Video, next int64, err error) {
	if time1 == 0 {
		time1 = util.GetTimeStamp()
	}
	//用户处于登录状态
	if token != "" {
		id, err := dao.GetIDByToken(token)
		if err != nil {
			return nil, 0, err
		}
		//fmt.Println("id" + strconv.FormatInt(id, 10))
		videos, err := dao.GetVideos(time1)
		for i := 0; i < len(videos); i++ {
			var video http_param.Video
			user, err := dao.GetUserByID(videos[i].UserId)
			if err != nil {
				return nil, 0, err
			}
			video.Author = http_param.User{}
			video.Author.Id = user.ID
			video.Author.Name = user.Username
			video.Author.FollowCount = user.FollowNum
			video.Author.FollowerCount = user.FollowerNum
			video.Author.IsFollow, err = dao.IsFollow(id, user.ID)
			if err != nil {
				return nil, 0, err
			}
			video.Id = videos[i].Id
			video.CoverUrl = videos[i].CoverUrl
			video.PlayUrl = videos[i].PlayUrl
			video.CommentCount = videos[i].CommentNum
			video.FavoriteCount = videos[i].FavouriteNum
			video.IsFavorite = false
			video.Title = videos[i].Title
			res = append(res, video)
		}
		if len(res) > 0 {
			next = videos[len(res)-1].CreatedAt
		}
		return res, next, err
	}
	//用户处于未登录状态
	videos, err := dao.GetVideos(time1)
	for i := 0; i < len(videos); i++ {
		var video http_param.Video
		user, err := dao.GetUserByID(videos[i].UserId)
		if err != nil {
			return nil, 0, err
		}
		video.Author = http_param.User{}
		video.Author.Id = user.ID
		video.Author.Name = user.Username
		video.Author.FollowCount = user.FollowNum
		video.Author.FollowerCount = user.FollowerNum
		video.Author.IsFollow = false
		video.Id = videos[i].Id
		video.CoverUrl = videos[i].CoverUrl
		video.PlayUrl = videos[i].PlayUrl
		video.CommentCount = videos[i].CommentNum
		video.FavoriteCount = videos[i].FavouriteNum
		video.IsFavorite = false
		video.Title = videos[i].Title
		res = append(res, video)
	}
	if len(res) > 0 {
		next = videos[len(res)-1].CreatedAt
	}
	return res, next, err
}

func GetVideosOfUser(id int64, token string) (res []http_param.Video, err error) {
	user, isMatch := dao.IsIDAndTokenMatch(id, token)
	if !isMatch {
		err = errors.New("id and token are not matched")
		return
	}
	videos, err := dao.GetVideosOfUser(id)
	for i := 0; i < len(videos); i++ {
		var video http_param.Video
		video.Author = http_param.User{}
		video.Author.Id = user.ID
		video.Author.Name = user.Username
		video.Author.FollowCount = user.FollowNum
		video.Author.FollowerCount = user.FollowerNum
		video.Author.IsFollow, err = dao.IsFollow(id, user.ID)
		if err != nil {
			return nil, err
		}
		video.Id = videos[i].Id
		video.CoverUrl = videos[i].CoverUrl
		video.PlayUrl = videos[i].PlayUrl
		video.CommentCount = videos[i].CommentNum
		video.FavoriteCount = videos[i].FavouriteNum
		video.IsFavorite, err = dao.IsFavorite(id, videos[i].Id)
		video.Title = videos[i].Title
		res = append(res, video)
	}
	return
}