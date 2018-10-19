package engine

type Request struct{
	Url string
	ParserFunc func([]byte)ParseResult
}

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

func NilParse([]byte)ParseResult{
	return ParseResult{}
}

type User struct {
	Name string
	Age string
	High string
	Income string //收入
	Marriage string
	Record string //学历
	Work string //职业
	Address string //地址
}