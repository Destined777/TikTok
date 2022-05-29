package dao

import (
	"TikTok/global"
	"TikTok/model"
	"gorm.io/gorm"
)

// ExistUser 确认用户是否存在，对未查询到与查询到用户均不报错而只是分类
func ExistUser(username string) (isExist bool, err error) {
	user := model.LogUser{}
	err = global.DB.Where("Username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func CreateUser(user model.LogUser) (ID int64, err error) {
	err = global.DB.Create(user).Error
	if err != nil {
		return
	}
	logUser, err := GetUserByName(user.Username)
	if err != nil {
		return
	}
	ID = logUser.ID
	return
}

func GetUserByName(username string) (user model.LogUser, err error) {
	err = global.DB.Where("username=?", username).Find(&user).Error
	return
}