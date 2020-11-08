package req

import "og/item"

type Request struct {
<<<<<<< HEAD
	tableName struct{} `pg:"job,alias:job"`
	UUID      string
	URL       string
	Host      string
	Download  string
	Datas     []*item.Field
	Status    string // 等待， 调度中， 成功， 失败， 重试
	Retry     int
	Log       string
}

// New 创建一个Request对象, 可以传入任何对象
func New(URL string) *Request {
	return &Request{
		URL: URL,
=======
	URL string
}

// New 创建一个Request对象, 可以传入任何对象
func New(i string) *Request {
	return &Request{
		URL: i,
>>>>>>> 249a72f8d86994d610206485cda12f6fbeed323a
	}
}
