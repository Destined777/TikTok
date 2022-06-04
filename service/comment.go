package service

import (
	"TikTok/dao"
	"TikTok/http_param"
	"TikTok/model"
	"errors"
	"time"
)

func Comment(params http_param.CommentParams) (res http_param.Comment, err error) {
	id, err := dao.GetIDByToken(params.Token)
	if err != nil {
		return
	}
	isExist, err := dao.ExistUserByID(id)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("user doesn't exist")
		return
	}

	video, err := dao.GetVideoByID(params.VideoId)
	if err != nil {
		return
	}
	user, err := dao.GetUserByID(video.UserId)
	if err != nil {
		return
	}

	res.User = http_param.User{}
	res.User.Id = user.ID
	res.User.Name = user.Username
	res.User.FollowCount = user.FollowNum
	res.User.FollowerCount = user.FollowerNum
	res.User.IsFollow, err = dao.IsFollow(id, user.ID)
	if err != nil {
		return
	}

	if params.ActionType == 1 { //评论
		if params.Text == "" {
			err = errors.New("the content of the comment is empty")
			return
		}
		month := time.Now().Month()
		day := time.Now().Day()
		var date string
		if month < 10 {
			date += "0"
		}
		date += string(rune(month))
		if day < 10 {
			date += "0"
		}
		date += string(rune(day))

		comment := model.Comment{
			UserId:    id,
			VideoId:   params.VideoId,
			Content:   params.Text,
			CreatedAt: date,
			IsDeleted: false,
		}
		id, err = dao.CreateComment(comment)
		if err != nil {
			return
		}
		res.Id = id
		res.Content = comment.Content
		res.CreateDate = comment.CreatedAt
	} else if params.ActionType == 2 { //删除评论
		if params.CommentId == 0 {
			err = errors.New("id is wrong")
		}
		err = dao.DeleteComment(params.CommentId)
		res.Id = params.CommentId
	} else {
		err = errors.New("action_type is wrong")
	}
	return
}

func CommentList(params http_param.CommentList) (res []http_param.Comment, err error) {
	id, err := dao.GetIDByToken(params.Token)
	if err != nil {
		return
	}
	isExist, err := dao.ExistUserByID(id)
	if err != nil {
		return
	}
	if !isExist {
		err = errors.New("user doesn't exist")
		return
	}

	video, err := dao.GetVideoByID(params.VideoId)
	if err != nil {
		return
	}
	user, err := dao.GetUserByID(video.UserId)
	if err != nil {
		return
	}

	comments, err := dao.GetComments(params.VideoId)
	for i := 0; i < len(comments); i++ {
		var temp http_param.Comment
		temp.User = http_param.User{}
		temp.User.Id = user.ID
		temp.User.Name = user.Username
		temp.User.FollowCount = user.FollowNum
		temp.User.FollowerCount = user.FollowerNum
		temp.User.IsFollow, err = dao.IsFollow(id, user.ID)
		if err != nil {
			return
		}
		temp.Id = comments[i].ID
		temp.Content = comments[i].Content
		temp.CreateDate = comments[i].CreatedAt
		res = append(res, temp)
	}
	return
}