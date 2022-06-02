package dao

import (
	"TikTok/global"
	"TikTok/model"
)

func CreateVideo(video model.Video) (err error) {
	err = global.DB.Create(&video).Error
	return
}
