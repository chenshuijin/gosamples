package main

import (
	"fmt"
	"golang.org/x/text/language"
)

func main() {
	t, err := language.Parse("ja-JP")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(t)
	tags := []language.Tag{
		language.Make("en-US"),
		language.Make("zh-CN"),
		language.Make("ko-KR"),
		language.Make("ja-JP"),
	}
	m := language.NewMatcher(tags)
	for {
		fmt.Println("input language string:")
		lang := ""
		fmt.Scanln(&lang)
		fmt.Println(m.Match(language.Make(lang)))
	}
	/*
		fmt.Println("begin match en-US")
		fmt.Println(m.Match(language.Make("")))
		fmt.Println(m.Match(language.Make("en_US")))
		fmt.Println(m.Match(language.Make("en")))
		fmt.Println(m.Match(language.Make("en_cn")))
		fmt.Println(m.Match(language.Make("en-cn")))
		fmt.Println(m.Match(language.Make("en_uk")))
		fmt.Println(m.Match(language.Make("en-cc")))
		fmt.Println("begin match ja-JP")
		fmt.Println(m.Match(language.Make("ja-JP")))
		fmt.Println(m.Match(language.Make("ja")))
		fmt.Println(m.Match(language.Make("ja_JP")))
		fmt.Println(m.Match(language.Make("ja-cc")))
		fmt.Println("begin match zh-CN")
		fmt.Println(m.Match(language.Make("zh")))
		fmt.Println(m.Match(language.Make("zh-cc")))
		fmt.Println(m.Match(language.Make("zh_cn")))
		fmt.Println(m.Match(language.Make("zh-us")))
		fmt.Println("begin match ko-KR")
		fmt.Println(m.Match(language.Make("ko")))
		fmt.Println(m.Match(language.Make("ko-cc")))
		fmt.Println(m.Match(language.Make("ko_cn")))
		fmt.Println(m.Match(language.Make("ko-KR")))

		/*	for _, tmp := range tags {
				fmt.Println(tmp.String())
			}
			fmt.Println("end print tags")
	*/
}
