package dao

import (
	"TikTok/global"
	"TikTok/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func Favorite(id, Vid int64) (err error) {
	favorite:= model.Favorite{}
	err = global.DB.Where("user_id = ? and video_id = ?", id, Vid).First(&favorite).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("a new one")
		favorite.UserId= id
		favorite.VideoId = Vid
		favorite.IsFavorite = true
		fmt.Println(favorite.VideoId)
		err = global.DB.Create(&favorite).Error
		return
	} else if err != nil {
		return
	} else {
		if favorite.IsFavorite == true {
			err = errors.New("it has been favorite")
		} else {
			favorite.IsFavorite = true
			err = global.DB.Save(&favorite).Error
		}
		return
	}
}

func UnFavorite(id, Vid int64) (err error) {
	favorite:= model.Favorite{}
	err = global.DB.Where("user_id = ? and video_id = ?", id, Vid).First(&favorite).Error
	if err == gorm.ErrRecordNotFound || favorite.IsFavorite == false {
		err = errors.New("it has not been favorite")
		return
	} else if err != nil {
		return
	} else {
		favorite.IsFavorite = false
		err = global.DB.Save(&favorite).Error
		return
	}
}

func GetFavoriteIds(id int64) (ids []int64, err error) {
	err = global.DB.Model(&model.Favorite{}).Where("user_id = ? and is_favorite = ?", id, true).Pluck("video_id", &ids).Error
	return
}

func IsFavorite(id, Vid int64) (Is bool, err error) {
	var favorite model.Favorite
	err = global.DB.Where("user_id = ? and video_id = ?", id, Vid).First(&favorite).Error
	if err == gorm.ErrRecordNotFound || favorite.IsFavorite == false {
		Is = false
		err = nil
	} else if err == nil {
		Is = true
	}
	return
}