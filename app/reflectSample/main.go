package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("yes")

	tmp := new(pushMessageText)
	fmt.Println("tmp before:", tmp)
	tv := reflect.ValueOf(tmp).Elem()
	v := reflect.ValueOf(&text).Elem()
	for i := 0; i < v.NumField(); i++ {
		tv.Field(i).SetString(v.Field(i).String())
	}
	//fmt.Println(v.Type())
	//fmt.Println(v.NumField())
	//fmt.Println("this is v before:" + v.Field(0).String())
	// add test
	//fmt.Println("this is text before:", text)
	//v.Field(0).SetString("dfsdf")
	//fmt.Println("this is v after: " + v.Field(0).String())
	//fmt.Println("this is text after:", text)
	//fmt.Println(v.NumField())
	fmt.Println("tmp after:", tmp)
}

var (
	text pushMessageText = pushMessageText{
		SecretLikeTicker:      "Someone you know has a crush on you! Find out who it is!",
		SecretLikeTickerShort: "Crush notification",
		SecretLikeTitle:       "You received a secret crush!",
		SecretLikeValue:       "Find out now",
	}
)

type pushMessageText struct {
	SecretLikeTicker      string
	SecretLikeTickerShort string
	SecretLikeTitle       string
	SecretLikeValue       string
}
