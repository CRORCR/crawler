package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(sends ...Request) {
	//存储所有需要解析的请求
	var requests []Request
	requests = append(requests, sends...)
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		//request去fatch 获得text文本
		log.Printf("url:%v\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetcher:error:%v url:%s\n", err, r.Url)
			continue
		}
		//解析文本
		parseResult := r.ParserFunc(body)
		//获得request加载到request队列,item输出即可
		requests = append(requests, parseResult.Requests...)

		//parseResult := parser.ParseCityList(body)
		for _, item := range parseResult.Items {
			log.Printf("got item:%v\n", item)
		}
	}
}
