package controller

import (
	"TikTok/http_param"
	"TikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	http_param.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	http_param.Response
	User http_param.User `json:"user"`
}

func Register(c *gin.Context) {
	params := http_param.UserLogIn{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}
	ID, token, err := service.Register(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: http_param.Response{StatusCode: 0},
			UserId:   ID,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	params := http_param.UserLogIn{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	ID, token, err := service.Login(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: 	http_param.Response{StatusCode: 0},
			UserId: 	ID,
			Token: 		token,
		})
	}
}

func UserInfo(c *gin.Context) {
	params := http_param.GetUser{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	userInfo, err := service.UserInfo(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: http_param.Response{StatusCode: 0},
			User:     userInfo,
		})
	}
}
