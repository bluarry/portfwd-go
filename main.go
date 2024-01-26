package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

/*
cli 端口转发工具，用法如下:
fwd -t tcp|udp 0.0.0.0:80 192.168.1.1:80
*/
func main() {
	app := cli.App{
		Name:   "fwd",
		Usage:  "port forward cli tool",
		Writer: os.Stdout,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "type",
				Aliases:  []string{"t"},
				Usage:    "tcp|udp",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			prot := ctx.String("type")
			if len(ctx.Args().Slice()) != 2 {
				return errors.New("params num is error")
			}

			soureIpPort := ctx.Args().Get(0)
			destIpPort := ctx.Args().Get(1)

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
