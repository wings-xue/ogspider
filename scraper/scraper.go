package scraper

import (
	"log"
	"og/db"
	req "og/reqeuest"
	"og/spider"
)

type Scraper struct {
	spiders   map[string]spider.OGSpider
	scheduler chan *req.Request
}

func (scraper *Scraper) OpenSpider(db *db.PgSQL) {
	for _, spider := range scraper.spiders {

		request := spider.StartRequest()
		log.Printf("[spider] init spider :%d", len(request))
		for _, r := range request {
			db.Update(r, true)
			scraper.scheduler <- r
		}

	}

}
