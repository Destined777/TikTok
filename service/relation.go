package service

import (
	"TikTok/dao"
	"TikTok/http_param"
	"TikTok/model"
	"errors"
)

func Follow(params http_param.Follow) (err error) {
	userId, err := dao.GetIDByToken(params.Token)
	if err != nil {
		return
	}
	isExist, err := dao.ExistUserByID(userId)
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

	_, isMatch := dao.IsIDAndTokenMatch(userId, params.Token)
	if !isMatch {
		err = errors.New("ID and Token are not matched")
	}

	if params.ActionType == 1 {
		err = dao.Follow(userId, params.ToID)
		if err != nil {
			err = dao.AddFollowNum(userId)
		}
		if err != nil {
			err = dao.AddFollowerNum(params.ToID)
		}
	} else if params.ActionType == 2 {
		err = dao.Unfollow(userId, params.ToID)
		if err != nil {
			err = dao.ReduceFollowNum(userId)
		}
		if err != nil {
			err = dao.ReduceFollowerNum(params.ToID)
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
		var logUser model.LogUser
		logUser, err = dao.GetUserByID(IDs[i])
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

func GetFollowers(params http_param.GetUser) (res []http_param.User, err error) {
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

	IDs, err := dao.FindFollowers(params.ID)
	if err != nil {
		return
	}

	for i := 0; i < len(IDs); i++ {
		var logUser model.LogUser
		logUser, err = dao.GetUserByID(IDs[i])
		if err != nil {
			return
		}
		var isFollow bool
		isFollow, err = dao.IsFollow(params.ID, logUser.ID)
		user := http_param.User{
			Id: 			logUser.ID,
			Name: 			logUser.Username,
			FollowCount: 	logUser.FollowNum,
			FollowerCount: 	logUser.FollowerNum,
			IsFollow:  		isFollow,
		}
		res = append(res, user)
	}
	return
}