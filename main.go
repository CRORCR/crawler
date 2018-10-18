package main

import (
	"fmt"
	"regexp"

)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
fetcher
}


func printCity(contents []byte) {
	//regexp.MustCompile(`<a href="http://www.zhenai.com/zhengh/un/ganzi" data-v-3ae1cdb9>甘孜</a>`)
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(contents, -1)
	for _,v:=range matchs{
		fmt.Printf("Url:%s City:%s\n",v[1],v[2])
	}

}
