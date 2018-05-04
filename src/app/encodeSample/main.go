package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("initialize.....")
	a := Sample{
		Aint:     123,
		Astring:  "this is a string",
		Afloat32: 1.00,
		Abyte:    0x32,
		Abytes:   []byte("gogogo"),
	}
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("got a error:", err)
	}
	fmt.Println("encoding....")
	fmt.Println("the json string:")
	fmt.Println(string(data))
	fmt.Println("decoding....")
	v := &Sample{}
	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println("got a error:", err)
	}
	fmt.Println("the truct:")
	fmt.Println(v)
}

type Sample struct {
	Aint     int     `json:"aint, omitempty"`
	Astring  string  `json:"astring, omitempty"`
	Afloat32 float32 `json:"afloat, omitempty"`
	Abyte    byte    `json:"abyte, omitempty"`
	Abytes   []byte  `json:"abytes, omitempty"`
}
