package response

import (
	"log"
	"net/url"
	ogconfig "og/const"
	"og/hash"
	"og/item"
	req "og/reqeuest"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Response struct {
	Req        *req.Request
	StatusCode int
	StatusMsg  string
	URL        string
	Page       string
	NewReq     []*req.Request
}

func NewFail(r *req.Request) *Response {
	newReq := make([]*req.Request, 0)
	return &Response{
		Req:        r,
		URL:        r.URL,
		StatusCode: 500,
		NewReq:     newReq,
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
		item.HasAttr(ogconfig.BaseCSS, true),
	)
	rowsField := item.Filter(
		self.Req.Datas,
		item.HasAttr(ogconfig.BaseCSS, false),
	)

	field := item.Append(self.ExtractRows(rowsField), self.ExtractRow(rowField))
	for _, f := range field {
		if f != nil {
			q := self.ToRequest(item.FindKey(ogconfig.URLName, f).Value)
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

func (self *Response) MatchBool(reg string) bool {
	if reg == "*" {
		return true
	}
	r, _ := regexp.Compile(reg)
	return r.FindString(self.URL) != ""
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

func (self *Response) newURL(oldURL string) string {
	return ""
}

func (self *Response) ToRequest(url string) *req.Request {
	url = self.ParseUrl(url, self.Req.Host)
	request := req.New(url)
	request.Datas = self.Req.Datas
	request.Host = self.Req.Host
	request.Status = req.StatusWait
	request.UUID = hash.Hash(url)
	request.Retry = 1
	request.Seed = false
	request.UpdateDate = time.Now()
	request.InsertDate = time.Now()
	request.FreshLife = self.Req.FreshLife
	return request
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
