package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"golang.org/x/net/html/charset"
	encoding "golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string)([]byte,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil,fmt.Errorf("wrong status:%d",resp.StatusCode)
	}
	encode := determineEncod(resp.Body)
	reader := transform.NewReader(resp.Body, encode.NewDecoder())
	return ioutil.ReadAll(reader)
}

//对determineEncod包装
func determineEncod(r io.Reader) encoding.Encoding {
	//使用net/http包查看页面的编码格式,需要传入1024bytes,但是直接读取后面就不能再读了,使用bufio偷窥一下
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		//如果不知道编码格式返回默认
		return unicode.UTF8
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "")
	return encode
}