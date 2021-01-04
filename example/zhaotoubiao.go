package example

import (
	"og/item"
	req "og/reqeuest"
	"og/response"
	"og/spider"
	"strconv"
	"strings"
)

const (
	ZTBName   = "zhaotoubiao"
	ZTBHost   = "www.chinabidding.cn"
	ListURL   = `https:\/\/www.chinabidding.cn\/search\/searchzbw\/search2\?areaid=\d+&keywords=&page=\d+&categoryid=&rp=22&table_type=0&b_date=`
	DetailURL = `www.chinabidding.cn/.*?/.*?.html`
	SaveTable = "zhaotoubiao"
	ErrorText = "很抱歉，由于您访问的URL有可能对网站造成安全威胁，您的访问被阻断"
)

var (
	ZTBStartURL = []string{

		"https://www.chinabidding.cn/search/searchzbw/search2?areaid=17&keywords=&page=1&categoryid=&rp=22&table_type=0&b_date=",
		"https://www.chinabidding.cn/search/searchzbw/search2?areaid=18&keywords=&page=1&categoryid=&rp=22&table_type=0&b_date=",
	}
	ZTBField = []*item.Field{
		{
			Name:  "web",
			Value: "采购与招标网",
		},
		{
			Name:    "title",
			BaseCSS: "[class*=listrow]",
			CSS:     "td:nth-child(2) a",
			Attr:    "innerText",
		},
		{
			Name:    "url",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(2) a",
			Attr:    "href",
		},
		{
			Name:    "address",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(5)",
			Attr:    "innerText",
		},
		{
			Name:    "publish_date",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(7)",
			Attr:    "innerText",
		},
		{
			Name:    "doc_type",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(4)",
			Attr:    "innerText",
		},
		{
			Name:    "industry",
			BaseCSS: "[class*=listrow]",
			CSS:     "[class*=listrow] td:nth-child(6)",
			Attr:    "innerText",
		},
		{
			Name: "pagetotal",
			CSS:  "#pages>span > a:last-child",
			Attr: "href",
			Do:   `(\d+)`,
		},
		{
			Name: "content",
			CSS:  ".xq_nr",
			Attr: "innerText",
		},
		{
			Name: "doc_html",
			CSS:  ".xq_nr",
			Attr: "innerHTML",
		},
		{
			Name:  "table_name",
			Value: "zhaotoubiao",
		},
	}
)

type Parse struct{}

func (p Parse) Hook(resp *response.Response) []*req.Request {
	out := make([]*req.Request, 0)
	for _, url := range ZTBStartURL {
		if resp.Req.URL == url {
			_total := spider.FindKey("pagetotal", resp.Req.Datas).Value
			total, _ := strconv.Atoi(_total)
			for i := 1; i < total; i++ {
				newPage := "page=" + strconv.Itoa(i)
				newURL := strings.Replace(url, "page=1", newPage, -1)
				q := *resp.Req
				q.URL = newURL
				q.UUID = spider.HashK(newURL)
				q.Seed = false
				out = append(out, &q)
			}
		}
	}
	return out

}
