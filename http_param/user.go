package http_param

import "TikTok/util"

type UserLogIn struct {
	Username	string `form:"username" binding:"required,min=1,max=32"`
	Password	string `form:"password" binding:"required,min=1,max=32"`
}

func (r *UserLogIn) GetError(err error) string {
	m := map[string]string{
		"Username": "用户名",
		"Password": "密码",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type GetUser struct {
	ID  	int64 	`form:"user_id" binding:"required"`
	Token   string 	`form:"token" binding:"required"`
}

func (r *GetUser) GetError(err error) string {
	m := map[string]string{
		"ID": "用户ID",
		"Token": "用户鉴权token",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}