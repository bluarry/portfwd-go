package service

import (
	"git.bluarry.top/bluarry/port-forward-cli/model"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type ForwardService struct {
	CliArgs           *model.FwdArgs
	udpForwardService *UdpForward
	tcpForwardService *TcpForward
}

func NewForwardJob(args *model.FwdArgs) *ForwardService {
	return &ForwardService{
		CliArgs:           args,
		udpForwardService: NewUdpForward(),
		tcpForwardService: NewTcpForWard(),
	}
}

func (s *ForwardService) Serve() error {
	switch s.CliArgs.Type {
	case "udp":
		err := s.udpForwardService.DoUdpForward(s.CliArgs.SourceHostPort, s.CliArgs.DestHostPort)
		if err != nil {
			log.Errorf("do udp forward failed ,error is %v", err)
			return err
		}
	case "tcp":
		err := s.tcpForwardService.DoTcpForward(s.CliArgs.SourceHostPort, s.CliArgs.DestHostPort)
		if err != nil {
			log.Errorf("do tcp forward failed ,error is %v", err)
			return err
		}
	default:
	}
	// 等待退出信号
	s.WaitExit()
	return nil
}

func (s *ForwardService) WaitExit() {
	// 等待退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	log.Println("udp listener started success")
	<-quit
	s.udpForwardService.Close()
	log.Println("udp listener will exit")
}
