package middle

import (
	"og/response"
	"strings"
)

// SpiderMiddle 中间件接口, 用来检查response是否正确
type SpiderMiddle interface {
	Hook(resp *response.Response) *response.Response
}

// ContentErrorMiddleware 基于response对象内容处理错误
type ContentErrorMiddleware struct {
	Code int
	Msg  string
}

// NewContentError 创建基于正文的错误处理
func (err ContentErrorMiddleware) NewContentError(msg string) ContentErrorMiddleware {
	return ContentErrorMiddleware{
		Code: 401,
		Msg:  msg,
	}
}

// ProcessSpiderIn 基于关键词的结果处理
func (err ContentErrorMiddleware) Hook(resp *response.Response) *response.Response {
	if strings.Contains(resp.Page, err.Msg) {
		resp.StatusCode = err.Code
		resp.StatusMsg = err.Msg
	}
	return resp
}

// AddSpiderMiddler 添加中间件到函数
// func (spider *BaseSpier) AddSpiderMiddler(m Middler) {
// 	spider.Middler = append(spider.Middler, m)
// }
