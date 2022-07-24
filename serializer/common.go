package serializer

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
    Code  int         `json:"code" default:"200"`
	Data  interface{} `json:"data,omitempty"`
    Msg   string      `json:"message" default:"ok"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeServiceErr 服务器发生错误
	CodeServiceErr = 502
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	// CodeEncryptError 加密失败
	CodeJudgeError = 50003
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}

// PermissionDeny 权限不足
func PermissionDeny() Response {
	return Response{
		Code: CodeNoRightErr,
		Msg:  "无权操作",
	}
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}

// ServiceErr 服务错误
func ServiceErr(msg string, err error) Response {
	if msg == "" {
		msg = "服务错误"
	}
	return Err(CodeServiceErr, msg, err)
}

// JudgeErr 评测错误
func JudgeErr(msg string, err error) Response {
	if msg == "" {
		msg = "评测错误"
	}
	return Err(CodeJudgeError, msg, err)
}

// OK 无误
func OK() Response {
	return Response{
		Code: 0,
		Msg:  "ok",
	}
}
