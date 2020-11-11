package response

import (
	"og/item"
	req "og/reqeuest"
	"og/setting"
)

type Response struct {
	Req        *req.Request
	StatusCode int
	URL        string
	Page       string
}

func New(req *req.Request) *Response {
	return &Response{
		Req:        req,
		StatusCode: 500,
	}
}

func (self *Response) Extract() []*req.Request {
	out := make([]*req.Request, 0)
	rowField := item.Filter(
		self.Req.Datas,
		item.HasAttr(setting.BaseCSS, true),
	)
	rowsField := item.Filter(
		self.Req.Datas,
		item.HasAttr(setting.BaseCSS, false),
	)

	field := item.Append(self.ExtractRows(rowsField), self.ExtractRow(rowField))
	for _, f := range field {
		out = append(out, req.ToRequest(f))
	}
	return out
}

func (self *Response) ExtractRow(field []*item.Field) []*item.Field {

	return field
}

func (self *Response) ExtractRows(field []*item.Field) [][]*item.Field {
	return [][]*item.Field{}
}
