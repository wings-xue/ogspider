package setting

const (
	NumGoroutine = 30

	// req
	CrawlerRstKey = "req_id"

	// field
	TableName = "table_name"

	BaseCSS = "BaseCSS"
	// 字段中用来生成新的request字段
	URLName = "url"

	StartURL     = "start_url"
	Host         = "host"
	Download     = "download"
	SaveResponse = "save_response"
	BeginURL     = "start_url"
	PageTotal    = "pagetotal"

	// component
	Downloader = "downloader"
	Scraper    = "scraper"
	Pipeliner  = "pipeliner"
)
