package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

var app *cli.App

func init() {
	//cli.CommandHelpTemplate = CommandHelpTemplate
	cli.AppHelpTemplate = AppHelpTemplate
	app = cli.NewApp()

	app.Action = calcCmd
	app.Name = filepath.Base(os.Args[0])
	app.Author = "csj"
	app.Email = "785795635@qq.com"
	app.Version = "1.0"
	app.Usage = "the countlines command line interface for calculate code file lines"

}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal("app run err:%v", err)
	}
}

func calcCmd(ctx *cli.Context) error {
	calcLinesInDir(ctx.String(dirFlag.Name), ctx.Bool(isAsyncFlag.Name))
	return nil
}

//var dir = flag.String("d", "./", "The path of the code folder")
//var isAsync = flag.Bool("a", false, "Is count multi-routine")
var fileTypes = []string{".go", ".java", ".c", ".js", ".cpp", ".h"}
var reg = regexp.MustCompile(`[*.go|*.java|*.c|*.js|*.cpp|*.h]$`)

func calcLinesInDir(dir string, isAsync bool) {
	start := time.Now()
	fullPath, err := filepath.Abs(dir)
	if err != nil {
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
	if isAsync {
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
		num := calcLinesInFile(v)
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
			num := calcLinesInFile(v)
			atomic.AddUint64(&allLines, uint64(num))
			wg.Done()
		}(v)
	}
	wg.Wait()
	return allLines
}

func calcLinesInFile(fullPath string) int {
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
