package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var dir = flag.String("d", "./", "The path of the code folder")
var fileTypes = []string{".go", ".jave", ".c"}

func main() {
	flag.Parse()
	start := time.Now()
	fullPath, _ := filepath.Abs(*dir)
	allFiles := ReadDir(fullPath)
	allLines := 0
	for _, v := range allFiles {
		num := CountFile(v)
		allLines += num

	}
	end := time.Now()
	fmt.Println("all lines:", allLines)
	fmt.Printf("cost time:%v\n", end.Sub(start))
}

func ReadDir(fullPath string) []string {
	fi, err := os.Stat(fullPath)
	if err != nil {
		panic(err)
	}

	if !fi.IsDir() {
		return []string{fullPath}
	}

	dirList, err := ioutil.ReadDir(fullPath)
	if err != nil {

		panic(err)
	}
	files := []string{}
	for _, v := range dirList {
		if v.Name() == "./" || v.Name() == "../" {
			continue
		}
		pathSep := string(os.PathSeparator)
		filepath := fullPath + pathSep + v.Name()

		_, err := os.Stat(filepath)
		if err != nil {
			panic(err)
		}

		if v.IsDir() {
			for _, f := range ReadDir(filepath) {
				files = append(files, f)
			}
		} else {
			for _, t := range fileTypes {
				if strings.HasSuffix(filepath, t) {
					files = append(files, filepath)
					break
				}
			}
		}
	}
	return files
}

func CountFile(fullPath string) int {
	fi, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	data, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}

	allText := fmt.Sprintf("%s", data)
	allText = strings.Replace(allText, " ", "", -1)
	allText = strings.Replace(allText, "\t", "", -1)
	allText = strings.Replace(allText, "\r", "", -1)

	allLines := strings.Split(allText, "\n")

	unCount := 0
	isUnCountBlock := false

	for _, line := range allLines {
		line = strings.Replace(line, "\n", "", -1)
		switch isUnCountBlock {
		case false:
			switch {
			case IsBlankLine(line):
				unCount += 1
			case strings.HasPrefix(line, "/*"):
				unCount += 1
				isUnCountBlock = true
			case strings.Contains(line, "/*"):
				isUnCountBlock = true
			case strings.HasSuffix(line, "*/"):
				unCount += 1
				isUnCountBlock = false
			case strings.Contains(line, "*/"):
				isUnCountBlock = false
			}
		case true:
			switch {
			case strings.HasSuffix(line, "*/"):
				isUnCountBlock = false
				unCount += 1
			case strings.Contains(line, "*/"):
				isUnCountBlock = false
			default:
				unCount += 1
			}
		}
	}
	return len(allLines) - unCount
}

func IsBlankLine(line string) bool {
	switch {
	case line == "", line == "\r", strings.HasPrefix(line, "//"):
		return true
	}
	return false
}
