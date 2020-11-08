package req

import "og/item"

const (
	StatusWait  = "waitting"
	StatusSche  = "scheduler"
	StatusSuc   = "succeed"
	StatusFail  = "fail"
	StatusRetry = "retry"
)

type Request struct {
	tableName struct{} `pg:"job,alias:job"`
	UUID      string
	URL       string
	Host      string
	Download  string
	Datas     []*item.Field
	Status    string // waitting， scheduler， succeed， fail， retry
	Retry     int
	Log       string
}

// New 创建一个Request对象, 可以传入任何对象
func New(URL string) *Request {
	return &Request{
		URL: URL,
	}
}
