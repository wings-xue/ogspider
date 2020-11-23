package scraped

import (
	"log"
	"og/db"
	req "og/reqeuest"
	"og/spider"
	"sync"
)

type scraper struct {
	spiders   map[string]spider.OGSpider
	mx        sync.Mutex
	scheduler chan *req.Request
}

func (scraper *scraper) OpenSpider(db *db.PgSQL) {
	for _, spider := range scraper.spiders {
		spider.CreateTable(db)
		request := spider.StartRequest()
		log.Printf("[spider] init spider :%d", len(request))
		for _, r := range request {
			db.Update(r, true)
			scraper.scheduler <- r
		}

	}

}
