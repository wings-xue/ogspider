package req

import (
	ogconfig "og/const"
	"og/hash"
	"og/item"
	"regexp"
)

const (
	StatusWait  = "waitting"
	StatusSche  = "scheduler"
	StatusSuc   = "succeed"
	StatusFail  = "fail"
	StatusRetry = "retry"
)

type Request struct {
	tableName struct{} `pg:"job,alias:job,discard_unknown_columns"`
	UUID      string
	URL       string
	Host      string
	Cookie    string
	Download  string
	Datas     []*item.Field
	Status    string // waitting， scheduler， succeed， fail， retry
	Retry     int
	Log       string
	Seed      bool `pg:"-"` // 是否做为种子爬取新的request对象

}

// New 创建一个Request对象, 可以传入任何对象
func New(URL string) *Request {
	return &Request{
		URL:  URL,
		Seed: false,
	}
}

func (request *Request) AddDatas(datas []*item.Field) {
	request.Datas = datas
}

func ToCrawlerRst(request *Request) map[string]interface{} {
	rst := make(map[string]interface{})
	for _, item := range request.Datas {
		rst[item.Name] = item.Value
	}
	rst[ogconfig.CrawlerRstKey] = request.UUID
	return rst
}

func ToTableSchema(tablename string, request *Request) string {

	column := ""
	for _, item := range request.Datas {
		column += item.Name + " text,\n"
	}
	column += ogconfig.CrawlerRstKey + " text,\n"
	return "create table if not exists " + tablename + "(" + column + " UNIQUE(req_id)) ;"

}

func ToRequest(fields []*item.Field) *Request {

	url := item.FindKey(ogconfig.URLName, fields).Value
	request := New(url)
	request.Datas = fields
	request.Host = item.FindKey(ogconfig.Host, fields).Value
	request.Status = StatusWait
	request.UUID = hash.Hash(url)
	request.Download = item.FindKey(ogconfig.Download, fields).Value
	request.Retry = 1
	request.Seed = false

	return request
}

func (self *Request) MatchBool(reg string) bool {
	if reg == "*" {
		return true
	}
	r, _ := regexp.Compile(reg)
	return r.FindString(self.URL) != ""
}
