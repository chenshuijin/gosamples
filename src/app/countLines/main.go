package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var dir = flag.String("d", "./", "The path of the code folder")
var isAsync = flag.Bool("a", false, "Is count multi-routine")
var fileTypes = []string{".go", ".java", ".c", ".js", ".cpp", ".h"}
var reg = regexp.MustCompile(`[*.go|*.java|*.c|*.js|*.cpp|*.h]$`)

func main() {
	flag.Parse()
	start := time.Now()
	fullPath, err := filepath.Abs(*dir)
	if err != nil {
		panic(err)
	}
	argPath := flag.Arg(flag.NArg() - 1)
	if fullPath, err = filepath.Abs(argPath); err != nil && argPath != "" {
		panic(err)
	}
	fmt.Println("path:", fullPath)
	allFiles := []string{}
	err = filepath.Walk(fullPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && reg.MatchString(path) {
			allFiles = append(allFiles, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	allLines := uint64(0)
	if *isAsync {
		fmt.Println("count async...")
		allLines = AsyncCount(allFiles)
	} else {
		fmt.Println("count sync...")
		allLines = SyncCount(allFiles)
	}

	end := time.Now()
	fmt.Println("lines amount:", allLines)
	fmt.Println("files amount:", len(allFiles))
	fmt.Printf("cost time:%v\n", end.Sub(start))
}

func SyncCount(allFiles []string) uint64 {
	allLines := uint64(0)
	for _, v := range allFiles {
		num := CountFile(v)
		allLines += uint64(num)
	}
	return allLines
}

func AsyncCount(allFiles []string) uint64 {
	wg := sync.WaitGroup{}
	allLines := uint64(0)
	for _, v := range allFiles {
		wg.Add(1)
		go func(v string) {
			num := CountFile(v)
			atomic.AddUint64(&allLines, uint64(num))
			wg.Done()
		}(v)
	}
	wg.Wait()
	return allLines
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
