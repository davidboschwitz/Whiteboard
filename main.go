package main

import (
	"github.com/codegangsta/cli"
	"github.com/hunterpraska/Whiteboard/cmd"
	"os"
)

const APP_VER = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "Whiteboard"
	app.Usage = "Go Education Software"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
	}
	app.Run(os.args)
}
