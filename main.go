package main

import (
	og "og/crawl"
	"og/example"
	"og/middle"
	"og/setting"
	"og/spider"
)

const ()

var s = setting.CrawlerSet{
	// PipelineSet设置，存入表
	PipelineSetting: map[string]setting.PipelineSet{
		example.DetailURL: {SaveTable: example.SaveTable},
	},
	// SpiderloadMiddleware 处理代理失败
	SpiderloadMiddleware: map[string][]middle.SpiderMiddle{
		"*": {middle.ContentErrorMiddleware{Code: 405, Msg: example.ErrorText}},
	},
	SpiderParse: map[string][]middle.Parse{
		example.ListURL: {example.Parse{}},
	},
}

func main() {
	og.Crawl(
		// 1. 创建BaseSpider对象
		// 2. 加载Name
		spider.SpiderNew(example.ZTBName).
			// 3. 加载Host
			SetHost(example.ZTBHost).
			// 4. 加载Fields
			SetFields(example.ZTBField).
			// 5. 加载StartURL
			SetStartURL(example.ZTBStartURL).
			// 6. 加载Setting
			SetSetting(s),
	)

}
