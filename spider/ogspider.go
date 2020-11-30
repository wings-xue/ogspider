package spider

import (
	"crypto/md5"
	"encoding/hex"
	ogconfig "og/const"
	"og/db"
	"og/item"
	"og/middle"
	req "og/reqeuest"
	"og/response"
	"og/setting"
)

// OGSpider 爬虫接口
// type OGSpider interface {
// 	StartRequest() []*req.Request
// 	Parse(resp *response.Response, r *req.Request) []*req.Request
// 	CreateTable(*db.PgSQL)
// 	CheckSpider()
// }

// BaseSpier 基础的爬虫配置
type BaseSpider struct {
	Host     string
	Fields   []*item.Field
	StartURL []string
	Name     string
	Setting  setting.CralwerSet
}

func SpiderMiddleware() map[string]middle.SpiderMiddle {
	out := make(map[string]middle.SpiderMiddle)
	out["ContentErrorMiddleware"] = middle.ContentErrorMiddleware{}
	return out
}

// SpiderNew 初始化一个新的spider
func SpiderNew(name string) *BaseSpider {
	return &BaseSpider{
		Name: name,
	}
}

func (spider *BaseSpider) CheckSpider() {
	if spider.Name == "" {
		panic("爬虫名称没有配置")
	}
	if spider.Host == "" {
		panic(spider.Name + "没有配置Host")
	}
	if spider.Fields == nil || len(spider.Fields) == 0 {
		panic(spider.Name + "提取字段(Field)没有配置")
	}
	if spider.StartURL == nil || len(spider.StartURL) == 0 {
		panic(spider.Name + "初始url没有配置")
	}
	if spider.Setting.PipelineSetting == nil {
		panic(spider.Name + "没有指定存入的表名")
	}
}

func (spider *BaseSpider) SetStartURL(startURL []string) *BaseSpider {
	spider.StartURL = startURL
	return spider
}

func (spider *BaseSpider) SetStartURLFunc(host string) *BaseSpider {

	return spider
}

func (spider *BaseSpider) SetHost(host string) *BaseSpider {
	spider.Host = host
	return spider
}

func (spider *BaseSpider) SetFields(field []*item.Field) *BaseSpider {
	spider.Fields = field
	return spider
}

func (spider *BaseSpider) SetSetting(field string) *BaseSpider {

	return spider
}

func (spider *BaseSpider) SetDownloadMiddleware() *BaseSpider {

	return spider
}

func (spider *BaseSpider) SetPipelineSet() *BaseSpider {

	return spider
}

func (spider *BaseSpider) SetDownloadMiddlewareFunc(field []*item.Field) *BaseSpider {
	spider.Fields = field
	return spider
}

func (spider *BaseSpider) SetPipelineMiddlewareFunc(field []*item.Field) *BaseSpider {
	spider.Fields = field
	return spider
}

func (spider *BaseSpider) SetPipelineMiddleware(field []*item.Field) *BaseSpider {
	spider.Fields = field
	return spider
}

func (spider *BaseSpider) SetSpiderMiddleware(field []*item.Field) *BaseSpider {
	spider.Fields = field
	return spider
}

func (spider *BaseSpider) SetSpiderMiddlewareFunc(field []*item.Field) *BaseSpider {
	spider.Fields = field
	return spider
}

func (spider *BaseSpider) Parse(resp *response.Response, r *req.Request) []*req.Request {
	return []*req.Request{}
}

func FindKey(key string, field []*item.Field) *item.Field {
	for _, f := range field {
		if f.Name == key {
			return f
		}
	}
	return &item.Field{}
}

func HashK(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])

}

// StartRequest 入口函数
func (spider *BaseSpider) StartRequest() []*req.Request {
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
func (spider *BaseSpider) CreateTable(db *db.PgSQL) {
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
	column += ogconfig.CrawlerRstKey + " text,\n"
	return "create table if not exists " + tablename + "(" + column + " UNIQUE(req_id)) ;"
}
