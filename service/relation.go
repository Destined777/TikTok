package service

import (
	"TikTok/dao"
	"TikTok/http_param"
	"errors"
)

func Follow(params http_param.Follow) (err error) {
	isExist, err := dao.ExistUserByID(params.UserID)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("user doesn't exist")
		return
	}

	isExist, err = dao.ExistUserByID(params.ToID)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("the one followed doesn't exist")
		return
	}

	_, isMatch := dao.IsIDAndTokenMatch(params.UserID, params.Token)
	if !isMatch {
		err = errors.New("ID and Token is not matched")
	}

	if params.ActionType == 1 {
		err = dao.Follow(params.UserID, params.ToID)
	} else if params.ActionType == 2 {
		err = dao.Unfollow(params.UserID, params.ToID)
	} else {
		err = errors.New("action_type is not correct")
	}
	return
}
