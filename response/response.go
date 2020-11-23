package response

import (
	"log"
	"net/url"
	"og/hash"
	"og/item"
	req "og/reqeuest"
	"og/setting"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Response struct {
	Req        *req.Request
	StatusCode int
	URL        string
	Page       string
}

func NewFail(req *req.Request) *Response {
	return &Response{
		Req:        req,
		URL:        req.URL,
		StatusCode: 500,
	}
}

func (self *Response) Extract() []*req.Request {
	// 1. 直接提取[]*field -> []*Field,并且修改原req为携带数据的新req
	// 2. 产生更多行的field[]*field -> [][]*Field
	// 3. 合并field ->[][]*Field
	// 4. 转为req.request
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
		if f != nil {
			q := self.ToRequest(item.FindKey(setting.URLName, f).Value)
			q.AddDatas(f)
			out = append(out, q)
		}

	}
	return out
}

func Selector(doc *goquery.Selection, css, attr string) string {
	switch {
	case attr == "innerHTML":
		s, _ := doc.Find(css).Html()
		return s
	case attr == "innerText":
		return doc.Find(css).Text()
	default:
		s, _ := doc.Find(css).Attr(attr)
		return s
	}
}

func (self *Response) Selector(css string, attr string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(self.Page))
	if err != nil {
		log.Printf("[response] %s\n", err.Error())
	}
	switch {
	case attr == "innerHTML":
		s, _ := doc.Find(css).Html()
		return s
	case attr == "innerText":
		return doc.Find(css).Text()
	default:
		s, _ := doc.Find(css).Attr(attr)
		return s
	}
}

func (self *Response) Match(reg string, s string) string {
	if reg == "" {
		return s
	}
	r, _ := regexp.Compile(reg)
	return r.FindString(s)
}

func (self *Response) ExtractRow(field []*item.Field) []*item.Field {
	for _, each := range field {
		if each.Value == "" {
			each.ExtractValue = self.Selector(each.CSS, each.Attr)
			each.Value = self.Match(each.Do, each.ExtractValue)

		}
	}
	return field
}

func (self *Response) ExtractRows(field []*item.Field) [][]*item.Field {
	out := make([][]*item.Field, 0)
	if len(field) == 0 {
		return out
	}
	base := field[0].BaseCSS
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(self.Page))
	if err != nil {
		log.Printf("[response] %s", err.Error())
	}
	doc.Find(base).Each(func(r int, s *goquery.Selection) {
		row := make([]*item.Field, 0)
		for i := 0; i < len(field); i++ {
			tmp := *field[i]
			tmp.ExtractValue = Selector(s, tmp.CSS, tmp.Attr)
			tmp.Value = self.Match(tmp.Do, tmp.ExtractValue)
			row = append(row, &tmp)
		}
		out = append(out, row)
	})
	return out
}

func (self *Response) newHeader() {}
func (self *Response) newBody()   {}
func (self *Response) newURL(oldURL string) string {
	return ""
}

func (self *Response) ToRequest(url string) *req.Request {
	url = self.ParseUrl(url, item.FindKey(setting.Host, self.Req.Datas).Value)
	request := req.New(url)
	request.Datas = self.Req.Datas
	request.Host = item.FindKey(setting.Host, self.Req.Datas).Value
	request.Status = req.StatusWait
	request.UUID = hash.Hash(url)
	request.Download = item.FindKey(setting.Download, self.Req.Datas).Value
	request.Retry = 1
	request.Seed = false

	return request
}

func (self *Response) ToPageReq(page int) []*req.Request {
	// 1. 生成header
	// 2. 生成data
	// 3. 生成url
	out := make([]*req.Request, 0)
	for i := 0; i < page; i++ {
		url := self.newURL(self.Req.URL)
		out = append(out, self.ToRequest(url))
	}
	return out
}

func (self *Response) ParseUrl(s, host string) string {
	u, err := url.Parse(s)
	if err != nil {
		log.Printf("[response] %s", err.Error())
	}
	if !u.IsAbs() {
		u.Host = host
		u.Scheme = "http"
	}

	return u.String()
}
