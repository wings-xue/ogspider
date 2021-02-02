package setting

import "og/middle"

// 基础配置
const Retry int = 6
const Headless bool = true
const QueryConut int = 100

// 填写配置
type PipelineSet struct {
	SaveTable string
}

type CrawlerSet struct {
	SpiderloadMiddleware map[string][]middle.SpiderMiddle
	SpiderParse          map[string][]middle.Parse
	DownloadMiddleware   map[string][]middle.DownloadMiddle
	PipelineSetting      map[string]PipelineSet
}

func New() CrawlerSet {
	return CrawlerSet{
		SpiderloadMiddleware: map[string][]middle.SpiderMiddle{},
		SpiderParse:          map[string][]middle.Parse{},
		DownloadMiddleware:   map[string][]middle.DownloadMiddle{},
		PipelineSetting:      map[string]PipelineSet{},
	}
}
