package config

import (
	"log"
)

type config struct {
	Path string
}

var Conf config

func Parse(path string) {
	log.Println("path:", path)
	Conf.Path = path
}
