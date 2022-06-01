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
		err = errors.New("ID and Token are not matched")
	}

	if params.ActionType == 1 {
		err = dao.Follow(params.UserID, params.ToID)
		if err != nil {
			err = dao.AddFollowNum(params.UserID)
		}
	} else if params.ActionType == 2 {
		err = dao.Unfollow(params.UserID, params.ToID)
		if err != nil {
			err = dao.ReduceFollowNum(params.UserID)
		}
	} else {
		err = errors.New("action_type is not correct")
	}
	return
}

func GetFollowings(params http_param.GetUser) (res []http_param.User, err error) {
	isExist, err := dao.ExistUserByID(params.ID)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("user doesn't exist")
		return
	}

	_, isMatch := dao.IsIDAndTokenMatch(params.ID, params.Token)
	if !isMatch {
		err = errors.New("ID and Token are not matched")
	}

	IDs, err := dao.FindFollowings(params.ID)
	if err != nil {
		return
	}

	for i := 0; i < len(IDs); i++ {
		logUser, err := dao.GetUserByID(IDs[i])
		if err != nil {
			return
		}
		user := http_param.User{
			Id: 			logUser.ID,
			Name: 			logUser.Username,
			FollowCount: 	logUser.FollowNum,
			FollowerCount: 	logUser.FollowerNum,
			IsFollow:  		true,
		}
		res = append(res, user)
	}
	return
}
