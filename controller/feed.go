package controller

import (
	"TikTok/http_param"
	"TikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	http_param.Response
	VideoList []http_param.Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	params := http_param.GetVideos{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	videos, nextTime, err := service.GetVideos(params.LatestTime, params.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	} else {
		c.JSON(http.StatusOK, FeedResponse{
			Response: http_param.Response{
				StatusCode: 0,
			},
			VideoList: videos,
			NextTime:  nextTime,
		})
	}
	/*c.JSON(http.StatusOK, FeedResponse{
		Response:  http_param.Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})*/
}
