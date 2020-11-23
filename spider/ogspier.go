package spider

import (
	"og/db"
	"og/item"
	req "og/reqeuest"
	"og/response"
	"og/setting"
)

// OGSpider 爬虫接口
type OGSpider interface {
	StartRequest() []*req.Request
	Parse(resp *response.Response, r *req.Request) []*req.Request
}

// BaseSpier 基础的爬虫配置
type BaseSpier struct {
	Host     string
	Fields   []*item.Field
	StartURL []string
	Crawler  map[string]string
	Name     string
}

// OpenSpider 初始化spider时调用
func OpenSpider(name string, startURL []string, field []*item.Field, host string, crawler map[string]string) *BaseSpier {
	spider := &BaseSpier{

		Host:     host,
		Fields:   field,
		StartURL: startURL,
		Crawler:  crawler,
		Name:     name,
	}
	return spider
}

// StartRequest 入口函数
func (spider *BaseSpier) StartRequest() []*req.Request {
	out := make([]*req.Request, 0)
	for _, url := range spider.StartURL {
		startReq := req.New(url)
		startReq.Datas = spider.Fields
		startReq.Host = spider.Host
		startReq.Status = req.StatusWait
		startReq.UUID = HashK(url)
		startReq.Retry = 1
		startReq.Seed = false
		out = append(out, startReq)
	}
	return out
}

// CreateTable 创建数据库
func (spider *BaseSpier) CreateTable(db *db.PgSQL) {
	tablename := spider.Name

	_, err := db.Conn.Exec(createTable(tablename, spider.Fields))
	if err != nil {
		panic(err)
	}
}

func createTable(tablename string, field []*item.Field) string {
	column := ""
	for _, item := range field {
		column += item.Name + " text,\n"
	}
	column += setting.CrawlerRstKey + " text,\n"
	return "create table if not exists " + tablename + "(" + column + " UNIQUE(req_id)) ;"
}
