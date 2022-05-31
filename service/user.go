package service

import (
	"TikTok/dao"
	"TikTok/http_param"
	"TikTok/model"
	"TikTok/util"
	"errors"
)

func Register(params http_param.UserLogIn) (ID int64, token string, err error) {
	token, _ = util.GenerateTokenByJwt(params.Username, params.Password)

	isExist, err := dao.ExistUser(params.Username)
	if err != nil {
		return
	}
	if isExist {
		err = errors.New("user has already existed")
	}

	user := model.LogUser{
		Username:	params.Username,
		Password: 	params.Password,
		Token: 		token,
		FollowNum:  0,
		FollowerNum:0,
		IsFollow: 	false,
	}
	ID, err = dao.CreateUser(user)
	return
}

func Login(params http_param.UserLogIn)(ID int64, token string, err error) {
	isExist, err := dao.ExistUser(params.Username)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("user doesn't exist")
		return
	}

	ID, token, isMatch := dao.IsUsernameAndPasswordMatch(params.Username, params.Password)
	if !isMatch {
		err = errors.New("password is not matched with the username")
	}
	return
}

func UserInfo(params http_param.GetUser)(userInfo http_param.User, err error) {
	isExist, err := dao.ExistUserByID(params.ID)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("user doesn't exist")
	}

	if user, match := dao.IsIDAndTokenMatch(params.ID, params.Token); match {
		userInfo = http_param.User{
			Id: 			user.ID,
			Name: 			user.Username,
			FollowCount: 	user.FollowNum,
			FollowerCount: 	user.FollowerNum,
			IsFollow: 		user.IsFollow,
		}
	} else {
		err = errors.New("user doesn't exist")
	}
	return
}
