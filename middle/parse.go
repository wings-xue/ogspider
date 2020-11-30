package middle

import (
	req "og/reqeuest"
	"og/response"
)

// SpiderMiddle 中间件接口, 用来检查response是否正确
type Parse interface {
	Hook(resp *response.Response) []*req.Request
}
