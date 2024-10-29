package pkg

import (
	"errors"
	"fmt"
	"module/app/constant"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicHandler(c *gin.Context) {
	// 패닉 복구
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		// 400, 401, etc 응답 처리
		switch key {
		case constant.DataNotFound.GetResponseStatus():
			// 에러 응답 반환
			c.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			// 이후 미들웨어 처리 중단
			c.Abort()
		case constant.Unauthorized.GetResponseStatus():
			c.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			c.Abort()
		}
	}
}
