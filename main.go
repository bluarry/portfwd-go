package main

import (
	"errors"
	"git.bluarry.top/bluarry/port-forward-cli/model"
	"git.bluarry.top/bluarry/port-forward-cli/service"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

/*
cli 端口转发工具，用法如下:
fwd -t tcp|udp 0.0.0.0:80 192.168.1.1:80
*/
func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
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

			sourceHosePort := ctx.Args().Get(0)
			destHostPort := ctx.Args().Get(1)
			cliArgs := &model.FwdArgs{
				Type:           prot,
				SourceHostPort: sourceHosePort,
				DestHostPort:   destHostPort,
			}
			svc := service.NewForwardJob(cliArgs)
			if err := svc.Serve(); err != nil {
				log.Printf("service run failed,error is %v", err)
				return err
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
