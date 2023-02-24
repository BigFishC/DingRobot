package main

import (
	"os"

	"github.com/liuchong/chat/src/server"
	"github.com/liuchong/chat/src/server/config"
	"github.com/urfave/cli/v2"
)

func main() {
	tn, at := config.ParseConfig(".")
	app := cli.NewApp()
	app.Name = "chat-dingtalk"
	app.Version = "v0.1.0"
	app.Usage = "Realize intelligent operation and maintenance"
	app.Commands = []*cli.Command{
		serverCMD(tn, at),
	}
	app.Run(os.Args)
}

func serverCMD(token string, as string) *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Run server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "conf",
				Aliases: []string{"c"},
				Usage:   "configuration file(.conf,.yaml,.toml)",
			},
		},
		Action: func(c *cli.Context) error {
			server.ServerStart(token, as)
			return nil
		},
	}
}
