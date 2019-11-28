package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	option := ""
	var cpuProfile *os.File
	for {
		fmt.Println("input opttion")
		fmt.Scanln(&option)
		fmt.Println("option:", option)
		switch option {
		case "heap":
			p := pprof.Lookup("heap")
			p.WriteTo(os.Stdout, 2)
		case "goroutine":
			p := pprof.Lookup("goroutine")
			p.WriteTo(os.Stdout, 2)
		case "threadcreate":
			p := pprof.Lookup("threadcreate")
			p.WriteTo(os.Stdout, 2)
		case "block":
			p := pprof.Lookup("block")
			p.WriteTo(os.Stdout, 2)
		case "mutex":
			p := pprof.Lookup("mutex")
			p.WriteTo(os.Stdout, 2)
		case "start":
			if cpuProfile == nil {
				if f, err := os.Create("game_server.cpuprof"); err != nil {
					log.Printf("start cpu profile failed: %v", err)
				} else {
					log.Print("start cpu profile")
					pprof.StartCPUProfile(f)
					cpuProfile = f
				}
			}
		case "stop":
			if cpuProfile != nil {
				pprof.StopCPUProfile()
				cpuProfile.Close()
				cpuProfile = nil
				log.Print("stop cpu profile")
			}
		case "memprof":
			if f, err := os.Create("game_server.memprof"); err != nil {
				log.Printf("record memory profile failed: %v", err)
			} else {
				runtime.GC()
				pprof.WriteHeapProfile(f)
				f.Close()
				log.Print("record memory profile")
			}
		}
	}
}
