package main

import (
	"fmt"
)

func main() {
	fmt.Println("OK")
	kvs := map[string]string{
		"test": "test",
	}
	fmt.Println("kvs:", kvs)
	kkvs := map[string]map[string]string{
		"zh": map[string]string{
			"test": "tsss",
		},
	}
	if v, isExist := kkvs["zh"]["test"]; !isExist {
		fmt.Println("no exist")
	} else {
		fmt.Println("exist:", v)
	}

}
