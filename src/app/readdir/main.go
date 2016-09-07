// main package
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	path := "/Users/chenshuijin/Downloads/"
	files, _ := ioutil.ReadDir(path)
	for _, fi := range files {
		if fi.IsDir() {
			fmt.Println("dir:"+path+"/"+fi.Name(), strings.HasPrefix(fi.Name(), "b"))
		} else {
			fmt.Println("file:" + path + "/" + fi.Name())
		}
	}
}
