package req

import (
	"og/hash"
	"og/item"
	"og/setting"
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

func ToCrawlerRst(request *Request) map[string]interface{} {
	rst := make(map[string]interface{})
	for _, item := range request.Datas {
		rst[item.Name] = rst[item.Value]
	}
	rst[setting.CrawlerRstKey] = request.UUID
	return rst
}

func ToTableSchema(tablename string, request *Request) string {

	column := ""
	for _, item := range request.Datas {
		column += item.Name + " text,\n"
	}
	column += setting.CrawlerRstKey + " text,\n"
	return "create table if not exists " + tablename + "(" + column + " UNIQUE(req_id)) ;"

}

func ToRequest(fields []*item.Field) []*Request {
	out := make([]*Request, 0)
	for _, field := range item.FindReq(fields) {
		url := field.Value
		request := New(url)
		request.Datas = fields
		request.Host = item.FindKey(setting.Host, fields).Value
		request.Status = StatusWait
		request.UUID = hash.Hash(url)
		request.Download = item.FindKey(setting.Download, fields).Value
		request.Retry = 1
		request.Seed = false
	}

	return out
}
