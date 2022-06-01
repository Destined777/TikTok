package http_param

import "TikTok/util"

type Follow struct {
	UserID 		int64 `form:"user_id" binding:"required"`
	Token 		string`form:"token" binding:"required"`
	ToID		int64 `form:"to_user_id" binding:"required"`
	ActionType  int32 `form:"action_type" binding:"required"`//1-关注，2-取消关注
}

func (r *Follow) GetError(err error) string {
	m := map[string]string{
		"UserID": 		"用户ID",
		"Token": 		"用户鉴权token",
		"ToID":  		"关注的用户ID",
		"ActionType":	"是否关注",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
