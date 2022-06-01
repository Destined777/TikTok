package dao

import (
	"TikTok/global"
	"TikTok/model"
	"gorm.io/gorm"
)

// ExistUser 确认用户是否存在，对未查询到与查询到用户均不报错而只是分类
func ExistUser(username string) (isExist bool, err error) {
	user := model.LogUser{}
	err = global.DB.Where("Username = ?", username).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func CreateUser(user model.LogUser) (ID int64, err error) {
	err = global.DB.Create(&user).Error
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

func IsUsernameAndPasswordMatch(username, password string) (ID int64, token string, isMatch bool) {
	u, _ := GetUserByName(username)
	if u.Password != password {
		return u.ID, u.Token, false
	}
	return u.ID, u.Token, true
}

func ExistUserByID(ID int64) (isExist bool, err error) {
	user := model.LogUser{}
	err = global.DB.Where("ID = ?", ID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func IsIDAndTokenMatch(ID int64, token string) (user model.LogUser, match bool) {
	global.DB.Where("ID = ?", ID).First(&user)
	if user.Token != token {
		match = false
	} else {
		match = true
	}
	return
}

func GetUserByID(id int64) (user model.LogUser, err error) {
	err = global.DB.Where("ID=?", id).Find(&user).Error
	return
}

func AddFollowNum(id int64) (err error) {
	var user model.LogUser
	err = global.DB.Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return
	}
	user.FollowNum++
	err = global.DB.Save(&user).Error
	return
}

func ReduceFollowNum(id int64) (err error) {
	var user model.LogUser
	err = global.DB.Where("ID = ?", id).Find(&user).Error
	if err != nil {
		return
	}
	if user.FollowerNum > 0 {
		user.FollowNum--
	}
	err = global.DB.Save(&user).Error
	return
}