package bcd

import (
	"config"
	"log"
	"testing"
)

var b = func() int {
	log.Println("auto run")
	return 0
}()

func TestGet(t *testing.T) {
	config.Parse("yes")
	tr := Get()
	if tr != "yes" {
		t.Errorf("expected 'yes' but get '%s'", tr)
	}
}
