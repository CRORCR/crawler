package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe  =`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte)engine.ParseResult{
	//regexp.MustCompile(`<a href="http://www.zhenai.com/zhengh/un/ganzi" data-v-3ae1cdb9>甘孜</a>`)
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result:=engine.ParseResult{}
	for _,v:=range matchs{
		result.Items=append(result.Items,string(v[2]))
		result.Requests=
			append(result.Requests,engine.Request{string(v[1]),engine.NilParse})
		//fmt.Printf("Url:%s City:%s\n",v[1],v[2])
	}
	return result
}
