package spider

type ZtbSpider struct {
	ogspdier BaseSpier
}

func (spider ZtbSpider) Parse() {
	// spider.StartRequest()
}

// func ZhaotoubiaoSpider() {
// 	pgsql := db.New()
// 	startURL := []string{
// 		// "http://www.baidu.com",
// 		"https://www.chinabidding.cn/search/searchzbw/search2?areaid=17&keywords=&page=1&categoryid=&rp=22&table_type=0&b_date=",
// 		// "https://www.chinabidding.cn/search/searchzbw/search2?areaid=18&keywords=&page=1&categoryid=&rp=22&table_type=0&b_date=",
// 	}
// 	field := []*item.Field{
// 		{
// 			Name:  "web",
// 			Value: "采购与招标网",
// 		},
// 		{
// 			Name:    "title",
// 			BaseCSS: "[class*=listrow]",
// 			CSS:     "td:nth-child(2) a",
// 			Attr:    "innerText",
// 		},
// 		{
// 			Name:    "url",
// 			BaseCSS: "[class*=listrow]",
// 			CSS:     "[class*=listrow] td:nth-child(2) a",
// 			Attr:    "href",
// 		},
// 		{
// 			Name:    "address",
// 			BaseCSS: "[class*=listrow]",
// 			CSS:     "[class*=listrow] td:nth-child(5)",
// 			Attr:    "innerText",
// 		},
// 		{
// 			Name:    "publish_date",
// 			BaseCSS: "[class*=listrow]",
// 			CSS:     "[class*=listrow] td:nth-child(7)",
// 			Attr:    "innerText",
// 		},
// 		{
// 			Name:    "doc_type",
// 			BaseCSS: "[class*=listrow]",
// 			CSS:     "[class*=listrow] td:nth-child(4)",
// 			Attr:    "innerText",
// 		},
// 		{
// 			Name:    "industry",
// 			BaseCSS: "[class*=listrow]",
// 			CSS:     "[class*=listrow] td:nth-child(6)",
// 			Attr:    "innerText",
// 		},
// 		{
// 			Name: "pagetotal",
// 			CSS:  "#pages>span > a:last-child",
// 			Attr: "href",
// 			Do:   `(\d+)`,
// 		},
// 		{
// 			Name: "content",
// 			CSS:  ".xq_nr",
// 			Attr: "innerText",
// 		},
// 		{
// 			Name: "doc_html",
// 			CSS:  ".xq_nr",
// 			Attr: "innerHTML",
// 		},
// 		{
// 			Name:  "table_name",
// 			Value: "zhaotoubiao",
// 		},
// 	}
// 	host := "www.chinabidding.cn"
// 	crawler := map[string]string{}
// 	name := "zhaotoubiao"
// 	ztb := OpenSpider(name, startURL, field, host, crawler)

// }
