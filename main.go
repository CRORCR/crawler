package main

import (
	"crawler/engine"
	"crawler/parser"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
 engine.Run(engine.Request{URL,parser.ParseCityList})
}