package service

import (
	"TikTok/dao"
	"TikTok/model"
)

func CreateVideo(title string, ID int64, name string) (err error) {
	video := model.Video{
		UserId:       ID,
		PlayUrl:      "http://10.21.191.27:8080/public/"+name,
		CoverUrl:     "",
		FavouriteNum: 0,
		CommentNum:   0,
		Title:        title,
	}
	err = dao.CreateVideo(video)
	return
}
