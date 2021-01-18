package req

import (
	ogconfig "og/const"
	"og/item"
	"regexp"
	"time"
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
	Datas     []*item.Field
	Status    string // waitting， scheduler， succeed， fail， retry
	Retry     int
	Log       string // 主要是网络错误

	Seed       bool      `pg:"-"` // 是否做为种子爬取新的request对象
	InsertDate time.Time `pg:"insert_date,alias:insert_date"`
	UpdateDate time.Time `pg:"update_date,alias:update_date"`
	// 生命周期
	FreshLife int `pg:"fresh_life,alias:fresh_life"`
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

func (self *Request) MatchBool(reg string) bool {
	if reg == "*" {
		return true
	}
	r, _ := regexp.Compile(reg)
	return r.FindString(self.URL) != ""
}
