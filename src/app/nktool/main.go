package main

import (
	"log"
	"os"
	"path/filepath"

	cli "gopkg.in/urfave/cli.v1"
)

var AppHelpTemplate = `NAME:
   {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
   {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

COMMANDS:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}
{{range .VisibleCommands}}{{if .VisibleFlags}}{{if .Name}}
{{.Name}} OPTIONS:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}
{{end}}{{end}}{{if .VisibleFlags}}
global OPTIONS:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}
{{if .Copyright}}
COPYRIGHT:
   {{.Copyright}}{{end}}
`

var app *cli.App

func init() {
	cli.AppHelpTemplate = AppHelpTemplate
	app = cli.NewApp()

	app.Name = filepath.Base(os.Args[0])
	app.Author = "csj"
	app.Email = "785795635@qq.com"
	app.Version = "1.0"
	app.Usage = "nas account manage interface"
	app.Commands = []cli.Command{
		newEthAccountCmd,
		listEthAccountsCmd,
		decEthKeyCmd,
		newNasAccountCmd,
		eth2nasKeyCmd,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal("app run err:%v", err)
	}
}
