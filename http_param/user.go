package http_param

import "TikTok/util"

type UserLogIn struct {
	Username	string `form:"username" binding:"required,min=1,max=32"`
	Password	string `form:"Password" binding:"required,min=1,max=32"`
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
