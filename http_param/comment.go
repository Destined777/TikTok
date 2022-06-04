package http_param

import "TikTok/util"

type CommentParams struct {
	Token 		string`form:"token" binding:"required"`
	VideoId		int64 `form:"video_id" binding:"required"`
	ActionType	int32 `form:"action_type" binding:"required"`//1-评论，2-删除评论
	Text		string`form:"comment_text" binding:"omitempty"`//1-关注，2-取消关注
	CommentId	int64 `form:"comment_id" binding:"omitempty"`//1-关注，2-取消关注
}

func (r *CommentParams) GetError(err error) string {
	m := map[string]string{
		"Token": 		"用户鉴权token",
		"VideoId":  	"视频ID",
		"ActionType":	"是否评论",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}

type CommentList struct {
	Token 		string`form:"token" binding:"required"`
	VideoId		int64 `form:"video_id" binding:"required"`
}

func (r *CommentList) GetError(err error) string {
	m := map[string]string{
		"Token": 		"用户鉴权token",
		"VideoId":  	"视频ID",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}
