package controller

import (
	"TikTok/http_param"
	"TikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(c *gin.Context) {
	params := http_param.FavoriteParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	err := service.Favorite(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, http_param.Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, http_param.Response{StatusCode: 0})
	}
}

func FavoriteList(c *gin.Context) {
	params := http_param.GetUser{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, VideoListResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
	}

	res, err := service.GetFavorite(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideoListResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: http_param.Response{
				StatusCode: 0,
			},
			VideoList: res,
		})
	}
	/*c.JSON(http.StatusOK, VideoListResponse{
		Response: http_param.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})*/
}
