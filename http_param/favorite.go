package http_param

import "TikTok/util"

type FavoriteParams struct {
	Token 		string`form:"token" binding:"required"`
	VideoID		int64 `form:"video_id" binding:"required"`
	ActionType  int32 `form:"action_type" binding:"required"`//1-关注，2-取消关注
}

func (r *FavoriteParams) GetError(err error) string {
	m := map[string]string{
		"Token": 		"用户鉴权token",
		"VideoID":  	"视频ID",
		"ActionType":	"是否点赞",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
