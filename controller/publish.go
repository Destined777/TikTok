package controller

import (
	"TikTok/dao"
	"TikTok/http_param"
	"TikTok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	http_param.Response
	VideoList []http_param.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	ID, err := dao.GetIDByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, http_param.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, http_param.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user, err := dao.GetUserByID(ID)
	if err != nil {
		c.JSON(http.StatusOK, http_param.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	finalName := fmt.Sprintf("%d_%s", user.ID, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, http_param.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	err = service.CreateVideo(title, ID, finalName)
	c.JSON(http.StatusOK, http_param.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: http_param.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
