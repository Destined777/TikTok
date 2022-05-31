package controller

import (
	"TikTok/http_param"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	http_param.Response
	UserList []http_param.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, http_param.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, http_param.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: http_param.Response{
			StatusCode: 0,
		},
		UserList: []http_param.User{DemoUser},
	})
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
