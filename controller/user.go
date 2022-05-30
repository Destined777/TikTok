package controller

import (
	"TikTok/dao"
	"TikTok/http_param"
	"TikTok/model"
	"TikTok/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	params := http_param.UserLogIn{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	token, _ := util.GenerateTokenByJwt(params.Username, params.Password)

	isExist, err := dao.ExistUser(params.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Something went wrong"},
		})
	}
	if isExist {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}

	user := model.LogUser{
		Username:	params.Username,
		Password: 	params.Password,
		Token: 		token,
		FollowNum:  0,
		FollowerNum:0,
		IsFollow: 	false,
	}
	ID, err := dao.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Something went wrong"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   ID,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	params := http_param.UserLogIn{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	isExist, err := dao.ExistUser(params.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Something went wrong"},
		})
	}
	if !isExist {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

	ID, token, isMatch := dao.IsUsernameAndPasswordMatch(params.Username, params.Password)
	if !isMatch {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Password is not matched with the username"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: 	Response{StatusCode: 0},
			UserId: 	ID,
			Token: 		token,
		})
	}
}

func UserInfo(c *gin.Context) {
	params := http_param.GetUser{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	isExist, err := dao.ExistUserByID(params.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Something went wrong"},
		})
	}
	if !isExist {
		c.JSON(http.StatusBadRequest, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

	if user, match := dao.IsIDAndTokenMatch(params.ID, params.Token); match {
		userInfo := User{
			Id: 			user.ID,
			Name: 			user.Username,
			FollowCount: 	user.FollowNum,
			FollowerCount: 	user.FollowerNum,
			IsFollow: 		user.IsFollow,
		}
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     userInfo,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
