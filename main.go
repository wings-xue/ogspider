package main

import (
	og "og/crawl"
	"og/spider"
)

func main() {
	og.Crawl(
		// 1. 创建BaseSpider对象
		// 2. 加载Name
		spider.SpiderNew(spider.ZTBName).
			// 3. 加载Host
			SetHost(spider.ZTBHost).
			// 4. 加载Fields
			SetFields(spider.ZTBField).
			// 5. 加载StartURL
			SetStartURL(spider.ZTBStartURL).
			SetStartURLFunc("").
			// 6. 加载Setting
			SetSetting("").
			SetDownloadMiddleware(spider.Zhaotoubiao()),
	)

}
