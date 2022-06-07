package util

import (
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"math/rand"
	"os"
	"time"
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

// GetTimeStamp 获取当前时间的int64形式
func GetTimeStamp() (t int64) {
	t = time.Now().Unix()
	return
}

// ReadFrameAsJpeg 从视频中抽帧
func ReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

// GenerateVerificationCode 生成6位数随机数
func GenerateVerificationCode() (code string) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}