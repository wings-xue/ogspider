package setting

import "og/middle"

type PipelineSet struct {
	SaveTable string
}

type CralwerSet struct {
	SpiderloadMiddleware map[string][]middle.SpiderMiddle
	SpiderParse          map[string][]middle.Parse
	DownloadMiddleware   map[string][]middle.DownloadMiddle
	PipelineSetting      map[string]PipelineSet
}
