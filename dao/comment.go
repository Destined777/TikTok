package dao

import (
	"TikTok/global"
	"TikTok/model"
	"errors"
)

func CreateComment(comment model.Comment) (ID int64, err error) {
	err = global.DB.Create(&comment).Error
	if err != nil {
		return
	}
	ID = comment.ID
	return
}

func DeleteComment(id int64) (err error) {
	var comment model.Comment
	err = global.DB.Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return
	}
	if comment.IsDeleted == true {
		err = errors.New("the comment has been deleted")
	} else {
		comment.IsDeleted = true
	}
	err = global.DB.Save(&comment).Error
	return
}

func GetComments(vid int64) (res []model.Comment, err error) {
	err = global.DB.Where("video_id = ? and is_deleted = ?", vid, false).Find(&res).Error
	return
}