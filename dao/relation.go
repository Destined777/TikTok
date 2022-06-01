package dao

import (
	"TikTok/global"
	"TikTok/model"
	"errors"
	"gorm.io/gorm"
)

func Follow(userID, toID int64) (err error) {
	follow := model.Follow{}
	err = global.DB.Where("UserID = ? and ToID = ?", userID, toID).Find(&follow).Error
	if err == gorm.ErrRecordNotFound {
		follow.UserID = userID
		follow.ToID = toID
		follow.IsFollow = true
		err = global.DB.Create(&follow).Error
		return
	} else if err != nil {
		return
	} else {
		if follow.IsFollow == true {
			err = errors.New("he has been followed")
		} else {
			follow.IsFollow = true
			err = global.DB.Save(&follow).Error
		}
		return
	}
}

func Unfollow(userID, toID int64) (err error) {
	follow := model.Follow{}
	err = global.DB.Where("UserID = ? and ToID = ?", userID, toID).Find(&follow).Error
	if err == gorm.ErrRecordNotFound || follow.IsFollow == false {
		err = errors.New("he has not been followed")
		return
	} else if err != nil {
		return
	} else {
		follow.IsFollow = false
		err = global.DB.Save(&follow).Error
		return
	}
}

func FindFollowings(id int64) (IDs []int64, err error) {
	err = global.DB.Model(&model.Follow{}).Where("UserID = ? and IsFollow = ?", id, true).Pluck("ToID", &IDs).Error
	return
}

func FindFollowers(id int64) (IDs []int64, err error) {
	err = global.DB.Model(&model.Follow{}).Where("ToID = ? and IsFollow = ?", id, true).Pluck("UserID", &IDs).Error
	return
}

func IsFollow(UserID, ToID int64) (isFollow bool, err error) {
	var follow model.Follow
	err = global.DB.Where("UserID = ? and ToID = ?", UserID, ToID).Find(&follow).Error
	if err == gorm.ErrRecordNotFound || follow.IsFollow == false {
		isFollow = false
		err = nil
	} else if err == nil {
		isFollow = true
	}
	return
}