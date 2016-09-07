package bcd

import (
	"config"
)

func Get() string {
	return config.Conf.Path
}
