package respone

import "net/http"

var StatusText = map[int]string{
	http.StatusOK:        			"ok",
	http.StatusBadRequest:			"failed",
	http.StatusUnauthorized:		"未登录或已失效",
	http.StatusMethodNotAllowed:	"没有访问权限",

	INVALID_PARAMS:              "请求参数出错",
	USER_PASSWORD_INCORRECT:     "用户帐号或登录密码不正确",
	USER_AUTHORIZATION:          "登录认证失败",
	USER_CREATED_FAILED:         "创建用户失败",
	USER_OLD_PASSWORD_INCORRENT: "旧密码不正确",
	USER_NOT_EXISTS:             "用户不存在",
}

const (
	INVALID_PARAMS 			= 4000
	USER_AUTHORIZATION 		= 4001
	USER_NOT_EXISTS 		= 4003
	USER_PASSWORD_INCORRECT = 4004

	USER_CREATED_FAILED 	= 4005
	USER_OLD_PASSWORD_INCORRENT = 4006
)
