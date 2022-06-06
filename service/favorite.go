package service

import (
	"TikTok/dao"
	"TikTok/http_param"
	"errors"
)

func Favorite(params http_param.FavoriteParams) (err error) {
	id, err := dao.GetIDByToken(params.Token)
	if err != nil {
		return
	}

	//fmt.Println("id" + strconv.FormatInt(id, 10))
	if params.ActionType == 1 {
		err = dao.Favorite(id, params.VideoID)
	} else if params.ActionType == 2 {
		err = dao.UnFavorite(id, params.VideoID)
	} else {
		err = errors.New("action_type is wrong")
	}
	return
}

func GetFavorite(params http_param.GetUser) (res []http_param.Video, err error) {
	isExist, err := dao.ExistUserByID(params.ID)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("user doesn't exist")
		return
	}

	_, isMatch := dao.IsIDAndTokenMatch(params.ID, params.Token)
	if !isMatch {
		err = errors.New("ID and Token are not matched")
	}

	Vids, err := dao.GetFavoriteIds(params.ID)
	for i := 0; i < len(Vids); i++ {
		var video http_param.Video
		temp, err := dao.GetVideoByID(Vids[i])
		if err != nil {
			return nil, err
		}
		user, err := dao.GetUserByID(temp.UserId)
		if err != nil {
			return nil, err
		}
		video.Author = http_param.User{}
		video.Author.Id = user.ID
		video.Author.Name = user.Username
		video.Author.FollowCount = user.FollowNum
		video.Author.FollowerCount = user.FollowerNum
		video.Author.IsFollow, err = dao.IsFollow(params.ID, user.ID)
		if err != nil {
			return nil, err
		}
		video.IsFavorite = true
		video.Title = temp.Title
		video.FavoriteCount = temp.FavouriteNum
		video.CommentCount = temp.CommentNum
		video.PlayUrl = temp.PlayUrl
		video.CoverUrl = temp.CoverUrl
		video.Id = temp.Id
		res = append(res, video)
	}
	return
}
