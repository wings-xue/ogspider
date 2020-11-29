package main

import (
	og "og/crawl"
	"og/spider"
)

func main() {
	og.Crawl(
		// 1. 创建BaseSpider对象
		// 2. 加载Name
		spider.SpiderNew("zhaotoubiao").
			// 3. 加载Host
			SetHost("").
			// 4. 加载Fields
			SetFields(spider.Zhaotoubiao()).
			// 5. 加载StartURL
			SetStartURL("").
			SetStartURLFunc("").
			// 6. 加载Setting
			SetSetting("").
			SetDownloadMiddleware(spider.Zhaotoubiao()),
	)

}
