package spider

import (
	"og/item"
	req "og/reqeuest"
)

const StartURL = "start_url"

func Zhaotoubiao() []*item.Field {

	ListURL := `https:\/\/www.chinabidding.cn\/search\/searchzbw\/search2\?areaid=\d+&keywords=&page=\d+&categoryid=&rp=22&table_type=0&b_date=`
	DetailURL := `www.chinabidding.cn/.*?/.*?.html`

	var Field = []*item.Field{
		{
			Name: "start_url",
			StartURL: []string{
				"https://www.chinabidding.cn/search/searchzbw/search2?areaid=17&keywords=&page=1&categoryid=&rp=22&table_type=0&b_date=",
				"https://www.chinabidding.cn/search/searchzbw/search2?areaid=18&keywords=&page=1&categoryid=&rp=22&table_type=0&b_date=",
			},
		},
		{
			Name:   "source_url",
			Value:  "https://www.chinabidding.cn/cgxx/",
			UrlReg: ListURL,
		},
		{
			Name:   "web",
			Value:  "采购与招标网",
			UrlReg: ListURL,
		},
		{
			Name:    "title",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(2) a",
			Attr:    "innerText",
			UrlReg:  ListURL,
		},
		{
			Name:    "url",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(2) a",
			Attr:    "href",
			UrlReg:  ListURL,
		},
		{
			Name:    "address",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(5)",
			Attr:    "innerText",
			UrlReg:  ListURL,
		},
		{
			Name:    "publish_date",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(7)",
			Attr:    "innerText",
			UrlReg:  ListURL,
		},
		{
			Name:    "doc_type",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(4)",
			Attr:    "innerText",
			UrlReg:  ListURL,
		},
		{
			Name:    "industry",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(6)",
			Attr:    "innerText",
			UrlReg:  ListURL,
		},
		{
			Name:   "pagetotal",
			CSS:    "#pages>span > a:last-child",
			Attr:   "href",
			Do:     `(\d+)`,
			UrlReg: ListURL,
		},
		{
			Name:   "content",
			CSS:    ".xq_nr",
			Attr:   "innerText",
			UrlReg: DetailURL,
		},
		{
			Name:   "doc_html",
			CSS:    ".xq_nr",
			Attr:   "innerHTML",
			UrlReg: DetailURL,
		},
	}
	return Field
}

func FindKey(key string, field []*item.Field) *item.Field {
	for _, f := range field {
		if f.Name == key {
			return f
		}
	}
	return nil
}

func ToSpider(item []*item.Field) []*req.Request {
	request := make([]*req.Request, 0)
	for _, url := range FindKey(StartURL, item).StartURL {
		startReq := req.New(url)
		startReq.Datas = item
		request = append(request, startReq)
	}
	return request
}
