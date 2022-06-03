package http_param

import (
	"TikTok/util"
)

type GetVideos struct {
	LatestTime  	int64 	`form:"latest_time" binding:"omitempty"`
	Token   		string 	`form:"token" binding:"omitempty"`
}

func (r *GetVideos) GetError(err error) string {
	m := map[string]string{
		"LatestTime": 	"限制最新时间",
		"Token": 		"用户鉴权token",
	}

	s := util.ErrorHandler(err, m)
	if s != "" {
		return s
	}
	return "参数错误"
}