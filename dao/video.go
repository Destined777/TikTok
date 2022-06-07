package dao

import (
	"TikTok/global"
	"TikTok/model"
)

func CreateVideo(video model.Video) (err error) {
	err = global.DB.Create(&video).Error
	return
}

func GetVideos(time1 int64) (videos []model.Video, err error) {
	err = global.DB.Where("created_at < ?", time1).Order("created_at desc").Limit(30).Find(&videos).Error
	return
}

func GetVideosOfUser(id int64) (videos []model.Video, err error) {
	err = global.DB.Where("user_id = ?", id).Order("created_at desc").Find(&videos).Error
	return
}

func GetVideoByID(id int64) (video model.Video, err error) {
	err = global.DB.Where("id = ?", id).Find(&video).Error
	return
}