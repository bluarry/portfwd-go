package main

import (
	"errors"
	"git.bluarry.top/bluarry/port-forward-cli/model"
	"git.bluarry.top/bluarry/port-forward-cli/service"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
)

/*
cli 端口转发工具，用法如下:
fwd -t tcp|udp 0.0.0.0:80 192.168.1.1:80
*/
func main() {
	log.SetLevel(log.DebugLevel)
	app := cli.App{
		Name:   "portfwd",
		Usage:  "port forward cli tool",
		Writer: os.Stdout,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "type",
				Aliases:  []string{"t"},
				Usage:    "tcp|udp",
				Required: true,
			},
			&cli.BoolFlag{
				Name:    "daemon",
				Aliases: []string{"d"},
				Value:   false,
			},
			&cli.StringFlag{
				Name:     "log",
				Aliases:  []string{"l"},
				Value:    "./run.log",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			prot := ctx.String("type")
			daemon := ctx.Bool("daemon")
			logPath := ctx.String("log")
			if len(ctx.Args().Slice()) != 2 {
				return errors.New("params is error")
			}

			sourceHosePort := ctx.Args().Get(0)
			destHostPort := ctx.Args().Get(1)
			cliArgs := &model.FwdArgs{
				Type:           prot,
				SourceHostPort: sourceHosePort,
				DestHostPort:   destHostPort,
			}
			svc := service.NewForwardService(cliArgs)
			f := func() error {
				if err := svc.Serve(); err != nil {
					log.Printf("service run failed,error is %v", err)
					return err
				}
				return nil
			}
			logfile, _ := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
			log.SetOutput(logfile)
			if daemon {
				RunInDaemon(f)
			} else {
				return f()
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func RunInDaemon(f func() error) {
	if os.Getppid() == 1 {
		err := f()
		if err != nil {
			log.Errorln("error to run tcp forward", err)
		}

	} else {
		args := make([]string, 0)
		for _, arg := range os.Args[1:] {
			if arg != "-d" && arg != "--daemon" {
				args = append(args, arg)
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
	}
}
