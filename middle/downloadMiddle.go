package middle

import req "og/reqeuest"

// SpiderMiddle 中间件接口, 用来检查response是否正确
type DownloadMiddle interface {
	Hook(req *req.Request) *req.Request
}

type CookieDownloadMiddle struct {
	Cookie string
}

func (cookie *CookieDownloadMiddle) Hook(req *req.Request) *req.Request {
	req.Cookie = cookie.Cookie
	return req
}
