package parser

import (
	"io/ioutil"
	"log"
	"regexp"
	"testing"
)
const resultSum  = 470
var cityList=[]string{"阿坝","阿克苏","阿拉善盟"}
var urlList=[]string{
	"http://www.zhenai.com/zhenghun/aba",
	"http://www.zhenai.com/zhenghun/akesu",
	"http://www.zhenai.com/zhenghun/alashanmeng"}
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityListData.html")
	if err!=nil{
		t.Fatalf("parse url failed:%v",err)
		return
	}
	result := ParseCityList(contents)

	if len(result.Requests)!=resultSum{
		t.Logf("num failed,should have:%d but:%d",resultSum,result.Requests)
	}

	for i:=0;i<3;i++{
		if result.Requests[i].Url != urlList[i]{
			log.Fatalf("fail,should url:%s,but:%s\n",urlList[i],result.Requests[i].Url)
		}
		if result.Items[i] != cityList[i]{
			log.Fatalf("fail,should city:%s,but:%s\n",cityList[i],result.Items[i])
		}
	}
}

func TestParseCityList2(t *testing.T) {
	text:=`<td><span class="label">年龄：</span>41岁</td>`
	re := regexp.MustCompile(`<span class="label">年龄：</span>([\d]+)岁</td>`)
	string := re.FindAllStringSubmatch(text,-1)
	t.Log(string)


	text=`<td><span class="label">月收入：</span>5001-8000元</td>`
	re = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	t.Log(re.FindAllStringSubmatch(text, -1))

	text=`<a href="http://album.zhenai.com/u/1572218980" target="_blank">安然</a>`
	re  =regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`)
	t.Log(re.FindAllStringSubmatch(text,-1))

}
