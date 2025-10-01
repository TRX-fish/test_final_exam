package utils

import "github.com/gin-gonic/gin"

const (
	OK           = "2000"
	NODATA       = "2001"
	DATAEXIST    = "2002"
	DATAERR      = "2003"
	PARAMERR     = "4000"
	SESSIONERR   = "4001"
	LOGINERR     = "4002"
	USERERR      = "4003"
	ROLEERR      = "4004"
	PWDERR       = "4005"
	REQERR       = "4006"
	IPERR        = "4007"
	URLNOTFOUND  = "4040"
	INTERNALERR  = "5000"
	DBERR        = "5001"
	THIRDERR     = "5002"
	IOERR        = "5003"
	UNKNOWNERR   = "5004"
)

var errorMapCN = map[string]string{
	OK:          "成功",
	NODATA:      "无数据",
	DATAEXIST:   "数据已存在",
	DATAERR:     "数据错误",
	PARAMERR:    "参数错误",
	SESSIONERR:  "用户未登录",
	LOGINERR:    "用户登录失败",
	USERERR:     "用户不存在或未激活",
	ROLEERR:     "用户身份错误",
	PWDERR:      "密码错误",
	REQERR:      "非法请求或请求次数受限",
	IPERR:       "IP受限",
	DBERR:       "数据查询错误",
	THIRDERR:    "第三方系统错误",
	IOERR:       "文件读写错误",
	UNKNOWNERR:  "未知错误",
}

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func MakeResponse(code string, msg string, data interface{}) Response {
	if msg == "" {
		msg = errorMapCN[code]
	}
	return Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, MakeResponse(OK, "", data))
}

func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, MakeResponse(OK, msg, data))
}

func Error(c *gin.Context, code int, errCode string, msg string) {
	c.JSON(code, MakeResponse(errCode, msg, nil))
} 