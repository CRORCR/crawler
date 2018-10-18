package regex
//package main
//
//import (
//	"fmt"
//	"regexp"
//)
//
//const text  = "my email address xingzhiheyi_code@163.com"
//func main() {
//	//re, _ := regexp.Compile("email")
//	//匹配邮箱 .:任意字符  +:0个或多个
//	re, _ := regexp.Compile(`[a-zA-Z_0-9]+@.+\..+`)
//	//匹配
//	match := re.FindAllStringSubmatch(text, -1)
//	fmt.Println(match)
//
//	//子匹配,所有匹配项分割开(把需要匹配的子项用()包裹)
//	re = regexp.MustCompile(`([a-zA-Z_0-9]+)@([a-zA-Z_0-9.]+)\.([a-zA-Z_0-9]+)`)
//	 submatch := re.FindAllStringSubmatch(text,-1)
//	 fmt.Println(submatch)
//
//}