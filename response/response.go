package response

import (
	req "og/reqeuest"
)

type Response struct {
	Req        *req.Request
	StatusCode int
}

func New(req *req.Request) *Response {
	return &Response{
		Req:        req,
		StatusCode: 500,
	}
}
