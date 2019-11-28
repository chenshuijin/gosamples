package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/go-ray/logging"
	cli "gopkg.in/urfave/cli.v1"
)

var app *cli.App
var conf = &Config{}

func init() {
	app = cli.NewApp()
	app.Action = run
	app.Before = before()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "csj"
	app.Email = "785795635@qq.com"
	app.Version = "1.0"
	app.Usage = "Secp256Key tool command line interface"
	app.Commands = []cli.Command{
		loopKeyscmd,
		qbalcmd,
	}

	app.Flags = []cli.Flag{
		ConfFlag,
	}
}

func before() cli.BeforeFunc {
	return func(ctx *cli.Context) error {
		configFile := ctx.String(ConfFlag.Name)
		configFile, _ = filepath.Abs(configFile)
		fmt.Println("config file:", configFile)
		loadConf(configFile)
		logging.InitLogger(conf.Log.Path, conf.Log.File, conf.Log.Level, conf.Log.Format)
		useDb := false
		for key, dbc := range conf.DBs {
			logging.Debug("key:", key)
			logging.Debug("dbc:", dbc)
			if dbc.Enable {
				if err := InitDatabaseConfig(key, dbc); err != nil {
					logging.Error("init database failed:", err)
				} else {
					continue
				}
				useDb = true
			}
		}
		if useDb {
			createAllTables()
		}
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		logging.Fatal("app run err:", err)
	}
}

func run(ctx *cli.Context) error {
	return nil
}

func loadConf(cpath string) {
	data, err := ioutil.ReadFile(cpath)
	if err != nil {
		log.Fatal("load config file failed:", err)
	}

	err = json.Unmarshal(data, conf)
	if err != nil {
		log.Fatal("load config file failed:", err)

	}
}
