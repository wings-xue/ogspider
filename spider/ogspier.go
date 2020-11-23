package spider

import (
	"og/db"
	"og/item"
	req "og/reqeuest"
	"og/response"
	"og/setting"
	"strings"
)

// Middler 中间件接口
type Middler interface {
	ProcessSpiderIn(resp *response.Response) *response.Response
	ProcessSpiderOut(resp *response.Response, result interface{}) []*item.Field
}

// ContentErrorMiddleware 基于response对象内容处理错误
type ContentErrorMiddleware struct {
	Code int
	Msg  string
}

// NewContentError 创建基于正文的错误处理
func (err ContentErrorMiddleware) NewContentError(msg string) ContentErrorMiddleware {
	return ContentErrorMiddleware{
		Code: 401,
		Msg:  msg,
	}
}

// ProcessSpiderIn 基于关键词的结果处理
func (err ContentErrorMiddleware) ProcessSpiderIn(resp *response.Response) *response.Response {
	if strings.Contains(resp.Page, err.Msg) {
		resp.StatusCode = err.Code
	}
	return resp
}

// OGSpider 爬虫接口
type OGSpider interface {
	StartRequest() []*req.Request
	Parse(resp *response.Response, r *req.Request) []*req.Request
	CreateTable(db *db.PgSQL)
}

// BaseSpier 基础的爬虫配置
type BaseSpier struct {
	Middler  []Middler
	Host     string
	Fields   []*item.Field
	StartURL []string
	Crawler  map[string]string
	Name     string
}

// OpenSpider 初始化spider时调用
func OpenSpider(name string, startURL []string, field []*item.Field, host string, crawler map[string]string) *BaseSpier {
	spider := &BaseSpier{
		Middler:  []Middler{},
		Host:     host,
		Fields:   field,
		StartURL: startURL,
		Crawler:  crawler,
		Name:     name,
	}
	return spider
}

// AddSpiderMiddler 添加中间件到函数
func (spider *BaseSpier) AddSpiderMiddler(m Middler) {
	spider.Middler = append(spider.Middler, m)
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

func createTable(tablename string, field []*item.Field) string {
	column := ""
	for _, item := range field {
		column += item.Name + " text,\n"
	}
	column += setting.CrawlerRstKey + " text,\n"
	return "create table if not exists " + tablename + "(" + column + " UNIQUE(req_id)) ;"
}

// AddSpiderMiddler 创建数据库
func (spider *BaseSpier) CreateTable(db *db.PgSQL) {
	tablename := spider.Name

	_, err := db.Conn.Exec(createTable(tablename, spider.Fields))
	if err != nil {
		panic(err)
	}
}
