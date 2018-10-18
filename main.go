package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		encode := determineEncod(resp.Body)
		reader := transform.NewReader(resp.Body, encode.NewDecoder())
		bytes, err := ioutil.ReadAll(reader)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%s\n", bytes)
		printCity(bytes)
	}
}

//对determineEncod包装
func determineEncod(r io.Reader) encoding.Encoding {
	//使用net/http包查看页面的编码格式,需要传入1024bytes,但是直接读取后面就不能再读了,使用bufio偷窥一下
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "")
	return encode
}

func printCity(contents []byte) {
	//regexp.MustCompile(`<a href="http://www.zhenai.com/zhengh/un/ganzi" data-v-3ae1cdb9>甘孜</a>`)
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(contents, -1)
	for _,v:=range matchs{
		fmt.Printf("Url:%s City:%s\n",v[1],v[2])
	}

}
