package controller

import (
	"TikTok/http_param"
	"TikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	http_param.Response
	CommentList []http_param.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	http_param.Response
	Comment http_param.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	params := http_param.CommentParams{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	comment, err := service.Comment(params)

	if err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: 	http_param.Response{StatusCode: 0},
			Comment: 	comment,
		})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	params := http_param.CommentList{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, CommentActionResponse{
			Response: http_param.Response{StatusCode: 1, StatusMsg: params.GetError(err)},
		})
		return
	}

	res, err := service.CommentList(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentListResponse{
			Response:    http_param.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    http_param.Response{StatusCode: 0},
		CommentList: res,
	})
}
