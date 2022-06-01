package controller

import (
	"TikTok/http_param"
	"TikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	http_param.Response
	UserList []http_param.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	params := http_param.Follow{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	err := service.Follow(params)

	if err != nil {
		c.JSON(http.StatusOK, http_param.Response{StatusCode: 1, StatusMsg: err.Error()})
	} else {
		c.JSON(http.StatusOK, http_param.Response{StatusCode: 0})
	}
}

// FollowList get all the users followed
func FollowList(c *gin.Context) {
	params := http_param.GetUser{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	res, err := service.GetFollowings(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: http_param.Response{
				StatusCode: 0,
			},
			UserList: res,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: http_param.Response{
			StatusCode: 0,
		},
		UserList: []http_param.User{DemoUser},
	})
}
