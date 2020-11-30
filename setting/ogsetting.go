package setting

import "og/middle"

type CralwerSet struct {
	SpiderloadMiddleware map[string][]middle.SpiderMiddle
	SpiderParse          map[string][]middle.Parse
	DownloadMiddleware   map[string][]middle.DownloadMiddle
}
