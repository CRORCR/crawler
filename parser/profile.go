package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	age = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
    high = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
	income = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	marriage = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
	record = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
    work= regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
	address = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
)

func Parseprofile(contents []byte,name string)engine.ParseResult{
	user:=&engine.User{}
	user.Name=name
	user.Age=extractString(contents,age)
	user.High = extractString(contents,high)
	user.Income = extractString(contents,income)
	user.Marriage = extractString(contents,marriage)
	user.Work = extractString(contents,work)
	user.Address = extractString(contents,address)
	result:=engine.ParseResult{Items:[]interface{}{user}}
	return result
}
func extractString(contents []byte, str *regexp.Regexp)string {
	submatch := str.FindSubmatch(contents)
	if len(submatch)==2{
		return string(submatch[1])
	}else{
		return ""
	}
}
