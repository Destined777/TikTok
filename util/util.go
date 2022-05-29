package util

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

// ErrorHandler 处理表单错误
func ErrorHandler(err error, m map[string]string) (msg string) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		return "输入参数错误"
	}
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, val := range validationErrs {
			switch val.Tag() {
			case "required":
				return m[val.Field()] + "不能为空哦"
			}
		}
	}
	return ""
}

//GenerateTokenByJwt 使用Jwt生成token作身份认证使用
//其中存储了username,password
func GenerateTokenByJwt(username, password string) (tokenString string, err error) {
	claims := jwt.MapClaims{
		"username":     username,
		"password": 	password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("userToken"))
	return
}