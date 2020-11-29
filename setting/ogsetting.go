package setting

import "og/middle"

type CralwerSet struct {
	Download map[string]string
	Scrape   map[string]string
	Pipeline map[string]string
	Schedule map[string]string

	SpiderloadMiddleware []middle.Middler
	DonwloadMiddleware   []middle.Middler
}
