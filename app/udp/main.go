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

var CommandHelpTemplate = `{{.cmd.Name}}{{if .cmd.Subcommands}} command{{end}}{{if .cmd.Flags}} [command options]{{end}} [arguments...]
{{if .cmd.Description}}{{.cmd.Description}}
{{end}}{{if .cmd.Subcommands}}
SUBCOMMANDS:
	{{range .cmd.Subcommands}}{{.cmd.Name}}{{with .cmd.ShortName}}, {{.cmd}}{{end}}{{ "\t" }}{{.cmd.Usage}}
	{{end}}{{end}}{{if .categorizedFlags}}
{{range $idx, $categorized := .categorizedFlags}}{{$categorized.Name}} OPTIONS:
{{range $categorized.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}{{end}}`

var app *cli.App

func init() {
	//cli.CommandHelpTemplate = CommandHelpTemplate
	cli.AppHelpTemplate = AppHelpTemplate
	app = cli.NewApp()

	app.Action = udpCmd
	app.Name = filepath.Base(os.Args[0])
	app.Author = "csj"
	app.Email = "785795635@qq.com"
	app.Version = "1.0"
	app.Usage = "the udp command line interface"
	app.Commands = []cli.Command{
		udpCliCmd,
		udpSvrCmd,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal("app run err:", err)
	}
}

func udpCmd(ctx *cli.Context) error {
	return nil
}
